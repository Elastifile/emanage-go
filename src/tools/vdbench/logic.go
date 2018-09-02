package vdbench

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"jobs"

	"github.com/elastifile/emanage-go/src/optional"
	"github.com/elastifile/emanage-go/src/runtimeutil"
	"github.com/elastifile/emanage-go/src/size"
	"github.com/elastifile/emanage-go/src/tools/common"
	tool_errors "github.com/elastifile/emanage-go/src/tools/errors"
	vdbench_config "github.com/elastifile/emanage-go/src/tools/vdbench/config"
	"github.com/elastifile/emanage-go/src/types"
)

type tool struct {
	*common.DefaultLoggingTool
	*common.ExternalToolProperties
	output string
}

type toolCmd struct {
	MountPoint   string
	RemoteFs     string
	ElfsFrontend string
	ConfigFile   string
	Loaders      []types.Host
	Results      string
}

var _ types.LoggingProperties = (*tool)(nil)

func populateConfig(params *types.ToolParams) *vdbench_config.Case {
	caseConf := &params.Config.Vdbench.Case
	if caseConf.DedupRatio == 0 {
		caseConf.DedupRatio = 10
	}
	if caseConf.DedupUnit == 0 {
		caseConf.DedupUnit = int(2 * size.KiB)
	}
	if caseConf.CompRatio == 0.0 {
		caseConf.CompRatio = 4.0
	}
	if caseConf.Debug == 0 {
		caseConf.Debug = 25
	}
	hds := make([]vdbench_config.Host, len(params.TargetLoaders))
	fdss := make([]vdbench_config.FileSystem, len(params.TargetLoaders))
	for i, loader := range params.TargetLoaders {
		hds[i] = vdbench_config.Host{
			User:   "root",
			Shell:  "ssh",
			Name:   loader,
			System: loader,
		}
		fdss[i] = vdbench_config.FileSystem{
			Name:   fmt.Sprintf("fsd%d", i),
			Anchor: fmt.Sprintf("/mnt/elfs/%s", loader),
			Depth:  2,
			Width:  8,
			Files:  150,
			Size: []int{
				int(128 * size.KiB),
				int(30 * size.Byte),
				int(512 * size.KiB),
				int(30 * size.KiB),
				int(1 * size.MiB),
				int(25 * size.Byte),
				int(10 * size.MiB),
				int(10 * size.Byte),
				int(20 * size.MiB),
				int(5 * size.Byte),
			},
			Shared:    true,
			OpenFlags: "o_direct",
		}
	}
	caseConf.Hosts = hds
	caseConf.FileSystems = fdss
	return caseConf
}

func NewTool(params *types.ToolParams) (types.Tool, error) {
	if err := common.ValidateExports(params); err != nil {
		return nil, err
	}

	// validate required config files are in minio bucket
	if err := common.ValidateConfigFiles(params.ConfigFilesBucket, []string{
		params.Config.Vdbench.ConfigFile}); err != nil {
		return nil, err
	}

	params.ToolName = types.ToolName(runtimeutil.PackageName())
	props, err := common.NewExternalProps(params)
	if err != nil {
		return nil, err
	}
	result := &tool{nil, props, "out.csv"}
	result.DefaultLoggingTool = common.NewLoggingTool(result)
	configFile, err := result.ParseTemplate(&common.ParseOpts{
		SourceBucket: params.ConfigFilesBucket,
		ConfigFile:   params.Config.Vdbench.ConfigFile,
		Template:     vdBenchTemplate,
		Data:         populateConfig(params),
	})
	if err != nil {
		return nil, err
	}
	params.Config.Vdbench.ConfigFile = configFile
	return result, nil
}

func (t *tool) GetResults(opts *types.ResultOpts) (err error) {
	host := t.Master()
	t.Logger().Info("Processing Vdbench results",
		"filesByHost", opts.FilesByHost, "host", host)
	results := opts.FilesByHost[host]
	if len(results) == 0 {
		return &tool_errors.NoResultsError{
			Host: host,
		}
	}

	summaries, err := t.populateResults(results[0])
	if err != nil {
		return err
	}
	err = jobs.UpdateToolResults(opts.JobID, summaries)
	if err != nil {
		return err
	}
	return err
}

func (t *tool) GetResultFilesPatterns() (result []string) {
	return []string{t.output}
}

func (t *tool) getCommonCommand(host types.Host, commandTemplate string) (cmd []string, err error) {
	var parse *template.Template
	params := t.Params()
	conf := params.Config
	tpl := template.New("cmd")
	path := t.DataStorePath()
	args := &toolCmd{
		MountPoint:   "/mnt/elfs",
		RemoteFs:     conf.Tesla.Elfs.Export,
		ElfsFrontend: params.System.Frontend,
		ConfigFile:   filepath.Join(path, conf.Vdbench.ConfigFile),
		Loaders:      conf.Loaders(),
		Results:      filepath.Join(path, t.output),
	}
	if err = ensureCmdArgs(args); err != nil {
		return nil, err
	}
	parse, err = tpl.Parse(commandTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = parse.Execute(&buf, args)
	return append(t.Shell, buf.String()), err
}

func (t *tool) GetSlaveCommand(host types.Host) (cmd []string, err error) {
	return t.getCommonCommand(host,
		`mkdir -p {{.MountPoint}} && \
	     mount -o soft {{.ElfsFrontend}}:{{.RemoteFs}} {{.MountPoint}} && \
         /usr/sbin/sshd && sleep infinity`)
}

func (t *tool) GetMasterCommand() (cmd []string, err error) {
	return t.getCommonCommand(t.Props.Master(),
		`mkdir -p {{.MountPoint}} && \
	     mount -o soft {{.ElfsFrontend}}:{{.RemoteFs}} {{.MountPoint}} && \
         {{range .Loaders}}
         mkdir -p {{$.MountPoint}}/{{.}} && \
         rm -rf {{$.MountPoint}}/{{.}}/* && \
         {{end}}
         ./vdbench -f {{.ConfigFile}} && \
         ./vdbench parse -i ./output/flatfile.html -c run rate resp MB/sec -a -o {{.Results}}`)
}

func ensureCmdArgs(cmd *toolCmd) (err error) {
	pass := cmd.MountPoint != "" &&
		cmd.RemoteFs != "" &&
		cmd.ElfsFrontend != "" &&
		cmd.ConfigFile != "" &&
		len(cmd.Loaders) > 0 &&
		cmd.Results != ""
	if !pass {
		err = &tool_errors.MissingArgsError{
			Args: cmd,
		}
	}
	return err
}

func (t *tool) populateResults(results *types.NamedReader) (summaries []types.ToolUnifiedResults, err error) {
	var (
		row    []string
		parsed types.ToolUnifiedResults
	)

	r := csv.NewReader(results)
	_, err = r.Read()
	if err != nil {
		return nil, err
	}

	parseIntRow := func(i int) (optional.Int, error) {
		parsedf, e := strconv.ParseFloat(row[i], 64)
		parsedi := int(parsedf)
		return &parsedi, e
	}
	parseFloatRow := func(i int) (optional.Float64, error) {
		parsedf, e := strconv.ParseFloat(row[i], 64)
		return &parsedf, e
	}

	name := string(t.Name())
	jobID := string(t.Params().JobID)
	validators := []func(){
		func() { row, err = r.Read() },
		func() {
			identifier := t.Params().Identifier
			parsed = types.ToolUnifiedResults{
				Identifier: &identifier,
				ToolName:   &name,
				JobID:      &jobID,
			}
		},
		func() {
			if len(row) < 3 {
				err = &tool_errors.MissingArgsError{Args: row}
			}
		},
		func() { parsed.IOps, err = parseIntRow(1) },
		func() { parsed.AverageLatency, err = parseFloatRow(2) },
		func() { parsed.Bandwidth, err = parseFloatRow(3) },
		func() {
			parsed.Workload = &row[0]
			summaries = append(summaries, parsed)
		},
	}

	i := 0
	for ; err == nil; i++ {
		validator := validators[i%len(validators)]
		validator()
		t.Logger().Debug("Validator ran", "validator", i)
	}

	if err != io.EOF {
		data, err := ioutil.ReadAll(results)
		if err != nil {
			return nil, err
		}
		return summaries, &tool_errors.InvalidResultsError{
			Raw:   string(data),
			Cause: err,
			Pos:   i,
		}
	}

	return summaries, nil
}
