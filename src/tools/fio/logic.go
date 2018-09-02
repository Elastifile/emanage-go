package fio

import (
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	text_template "text/template"

	"jobs"

	"github.com/elastifile/emanage-go/src/runtimeutil"
	"github.com/elastifile/emanage-go/src/tools/common"
	tool_errors "github.com/elastifile/emanage-go/src/tools/errors"
	"github.com/elastifile/emanage-go/src/types"
)

var (
	// Results file
	resultsFilePrefix = "fio"
	resultsFileMiddle = "results"
	resultsFileSuffix = "tesla"
	resultsFileExt    = "json"
	outputFormat      = "json"

	// Execution commands
	fioBaseCmd         = "/opt/fio/fio/fio"
	checkFioServersCmd = "/opt/fio/check_fio_servers.sh"
	mountPoint         = "/mnt"
)

type tool struct {
	*common.DefaultLoggingTool
	*common.ExternalToolProperties
}

type checkLoadersParams struct {
	Retries     int
	TimeToSleep int
}

type outputFileParams struct {
	Path      string
	Prefix    string
	Middle    string
	Suffix    string
	Extention string
	Format    string
}

type toolCmd struct {
	MountOptions       string
	MountPoint         string
	RemoteFs           string
	ElfsFrontend       string
	ConfigFile         string
	CurrentLoader      string
	BaseCmd            string
	Loaders            []types.Host
	CheckLoadersCmd    string
	CheckLoadersParams checkLoadersParams
	OutputFileParams   outputFileParams
	Jobs               []string
}

type runeReader interface {
	io.Reader
	io.RuneScanner
}

// Do avarge to all internal jobs fio job file to one line results
var averageAllResults = false

func NewTool(params *types.ToolParams) (types.Tool, error) {
	// validate required config files are in filestore
	if err := common.ValidateConfigFiles(params.ConfigFilesBucket, []string{
		params.Config.Fio.ConfigFile}); err != nil {
		return nil, err
	}

	params.ToolName = types.ToolName(runtimeutil.PackageName())

	props, err := common.NewExternalProps(params)
	if err != nil {
		return nil, err
	}

	result := &tool{
		ExternalToolProperties: props,
	}
	result.DefaultLoggingTool = common.NewLoggingTool(result)
	logger := result.Logger()

	workType := result.Params().Config.Fio.WorkType
	logger.Info("Fio - NewTool", "WorkType", workType, "WorkJobs", result.Params().Config.Fio.WorkJobs)
	fileTemplate := defaultFileTemplate
	if workType == "data_integrity" {
		logger.Debug("WorkType = data_integrity")
		fileTemplate = diFileTemplate
	}

	configFile, err := result.ParseTemplate(&common.ParseOpts{
		SourceBucket: params.ConfigFilesBucket,
		ConfigFile:   params.Config.Fio.ConfigFile,
		Template:     fileTemplate,
		Data:         params.Config,
	})
	if err != nil {
		return nil, err
	}

	params.Config.Fio.ConfigFile = configFile
	return result, nil
}

func (t *tool) GetMasterCommand() (cmd []string, err error) {
	// Update: a ticket was open about resolving these issues - http://jira.il.elastifile.com/browse/INF-313
	// REVIEW(gavriep): We must find a better way to write these templates. They are quite unreadable like this.
	// Maybe we should put the templates as cleanly formatted, separate files in the tool images and read and parse them from there?
	return t.commonCommand(t.Master(), `mkdir -p {{.MountPoint}} && \
	 mount {{.MountOptions}} {{.ElfsFrontend}}:{{.RemoteFs}} {{.MountPoint}} && \
	 { {{.BaseCmd}} --directory {{.MountPoint}} --server &} && \
	 {{.CheckLoadersCmd}}{{$chkParams := .CheckLoadersParams}} \
	 -r {{$chkParams.Retries}} -t {{$chkParams.TimeToSleep}} \
	 {{range .Loaders}} {{.}}{{end}} && \
	 {{.BaseCmd}}{{$outParams := .OutputFileParams}}\
	 --output-format={{$outParams.Format}}\
	 --output={{$outParams.Path}}/{{$outParams.Prefix}}_{{$outParams.Middle}}_{{$outParams.Suffix}}.{{$outParams.Extention}}\
	 {{.ConfigFile}}\
	 {{$jobs:=.Jobs}}{{range .Loaders}} --client={{.}}{{if $jobs}}{{range $jobs}} --section={{.}}{{end}}{{end}}{{end}}`)

}

func (t *tool) GetSlaveCommand(host types.Host) (cmd []string, err error) {
	return t.commonCommand(host, `mkdir -p {{.MountPoint}} && \
	 mount {{.MountOptions}} {{.ElfsFrontend}}:{{.RemoteFs}} {{.MountPoint}} && \
	 { {{.BaseCmd}} --directory {{.MountPoint}} --server &} && \
	 sleep infinity`)
}

func (t *tool) commonCommand(host types.Host, cmdTemplate string) (cmd []string, err error) {
	params := t.Params()
	conf := params.Config
	path := t.DataStorePath()
	configFile := filepath.Join(path, conf.Fio.ConfigFile)
	var chkLoadersParams = checkLoadersParams{
		Retries:     conf.Fio.CheckServersRetries,
		TimeToSleep: conf.Fio.CheckServersTimeToSleep,
	}
	var outFileParams = outputFileParams{
		Path:      path,
		Prefix:    resultsFilePrefix,
		Middle:    resultsFileMiddle,
		Suffix:    resultsFileSuffix,
		Extention: resultsFileExt,
		Format:    outputFormat,
	}
	var workJobs = conf.Fio.WorkJobs
	var jobs []string
	if workJobs != "" {
		jobs = strings.Split(workJobs, ",")
	}

	t.Logger().Info("Fio - commonCommand", "jobs", jobs)
	tpl := text_template.New("cmd")
	args := &toolCmd{
		MountPoint:         mountPoint,
		MountOptions:       conf.Tesla.Elfs.MountOptions,
		RemoteFs:           conf.Tesla.Elfs.Export,
		ElfsFrontend:       params.System.Frontend,
		ConfigFile:         configFile,
		CurrentLoader:      string(host),
		Loaders:            conf.Loaders(),
		BaseCmd:            fioBaseCmd,
		CheckLoadersCmd:    checkFioServersCmd,
		CheckLoadersParams: chkLoadersParams,
		OutputFileParams:   outFileParams,
		Jobs:               jobs,
	}
	parse, err := tpl.Parse(cmdTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = parse.Execute(&buf, args)
	cmd = append(t.Shell, buf.String())

	return cmd, err
}

func (t *tool) GetResultFilesPatterns() []string {
	return []string{
		resultsFilePrefix + "*",
	}
}

func (t *tool) GetResults(opts *types.ResultOpts) (err error) {
	var sumFile *types.NamedReader
	master := t.Master()

	files, ok := opts.FilesByHost[master]
	if !ok {
		return &tool_errors.NoResultsError{
			Host: master,
		}
	}

	for _, f := range files {
		if strings.HasPrefix(f.Name, resultsFilePrefix) {
			sumFile = f
			break
		}
	}

	data, err := ioutil.ReadAll(sumFile)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		t.Logger().Error("fio empty results", "filesByHost", opts.FilesByHost)
		return &tool_errors.EmptyResultsError{
			File: sumFile,
			Host: master,
		}
	}

	t.Logger().Debug("fio GetResults", "sumfile content", string(data))
	sumFileResults, err := t.parseAllResults(bytes.NewReader(data))

	var results []types.ToolUnifiedResults

	name := string(t.Name())
	jobID := string(t.Params().JobID)
	for _, r := range sumFileResults {
		r := r // Create new instance of r for fresh referencing, for more see: docs/coding_conventions.md#bug-in-go---for-loop
		identifier := t.Params().Identifier
		if r.jobname != "" {
			identifier += "-" + r.jobname
		}
		results = append(results, types.ToolUnifiedResults{
			Identifier:     &identifier,
			ToolName:       &name,
			Workload:       &r.jobname,
			Bandwidth:      &r.Bandwidth,
			IOps:           &r.IOPS,
			AverageLatency: &r.Latency,
			JobID:          &jobID,
		})
	}

	err = jobs.UpdateToolResults(opts.JobID, results)
	if err != nil {
		return err
	}

	t.Logger().Debug("Successfully got results")
	return nil
}
