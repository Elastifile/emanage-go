package sfs2008

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"jobs"
	"runtimeutil"
	"tools/common"
	tool_errors "tools/errors"
	"types"
)

var resultsSuffix = "tesla"

var resultsFile = "sfssum"
var resultsFilePatterns = []string{
	resultsFile + "*",
	"sfsres*",
}

// var resultsFilePrefixes = []string{"sfssum", "sfsres"}
var resultsFilePrefixes = []string{}

// Execution commands
var ensureRPC = "(rpcinfo > /dev/null 2>&1 || rpcbind) && "
var symlinkResults = "ln -s %s /opt/sfs2008/sfs2008/result && "
var masterCmd = "java SfsManager -r %s -s %s"
var slaveCmd = "java -Djava.security.policy=java.policy -jar Manager.jar"

// to simulate sfs run with parsable results
var fakeMasterWithResultsCmd = `echo "INVALID   10000   11523     0.9   345703   30 NFS3 T 4   18882360   2  5  2  2 2008" > /opt/sfs2008/sfs2008/result/sfssum.tesla`

type tool struct {
	*common.DefaultLoggingTool
	*common.ExternalToolProperties
}

func NewTool(params *types.ToolParams) (types.Tool, error) {
	if err := common.ValidateExports(params); err != nil {
		return nil, err
	}

	// validate required config files are in filestore (minio) bucket
	if err := common.ValidateConfigFiles(params.ConfigFilesBucket, []string{
		params.Config.Sfs2008.ConfigFile,
		params.Config.Sfs2008.MixFile}); err != nil {
		return nil, err
	}

	params.ToolName = types.ToolName(runtimeutil.PackageName())

	props, err := common.NewExternalProps(params)
	if err != nil {
		return nil, err
	}
	result := &tool{nil, props}
	result.DefaultLoggingTool = common.NewLoggingTool(result)
	configFile, err := result.ParseTemplate(&common.ParseOpts{
		SourceBucket: params.ConfigFilesBucket,
		ConfigFile:   params.Config.Sfs2008.ConfigFile,
		Template:     template,
		Data:         params.Config,
	})
	if err != nil {
		return nil, err
	}
	params.Config.Sfs2008.ConfigFile = configFile
	return result, nil
}

func (t *tool) GetMasterCommand() (cmd []string, err error) {
	cmd, err = t.GetSlaveCommand(t.Props.Master())
	if err != nil {
		return nil, err
	}

	conf := &t.Params().Config.Sfs2008
	path := t.DataStorePath()
	configFile := filepath.Join(path, conf.ConfigFile)
	ln := fmt.Sprintf(symlinkResults, path)
	exec := fmt.Sprintf(masterCmd, configFile, resultsSuffix)

	// the following will simulate a successfull sfs2008 run. it will produce a parsable sfssum results.
	if false {
		exec = fakeMasterWithResultsCmd
	}

	cmd = append(t.Shell, ensureRPC+ln+exec)
	return cmd, nil
}

func (t *tool) GetSlaveCommand(host types.Host) (cmd []string, err error) {
	return append(t.Shell, ensureRPC+slaveCmd), nil
}

func (t *tool) GetResultFilesPatterns() []string {
	var patterns []string
	for _, name := range resultsFilePrefixes {
		patterns = append(patterns, fmt.Sprintf("%v*", name))
	}
	return patterns
}

func (t *tool) GetResults(opts *types.ResultOpts) (err error) {
	master := t.Master()
	files, ok := opts.FilesByHost[master]
	if !ok {
		return &tool_errors.NoSfsResultsError{
			Host:   master,
			Suffix: resultsSuffix,
		}
	}

	sumFile := &types.NamedReader{}

	for _, f := range files {
		if strings.HasPrefix(f.Name, "sfssum") {
			sumFile = f
			break
		}
	}

	if sumFile.Name == "" { // didn't find any sum file
		return &tool_errors.EmptyResultsError{
			File: sumFile,
			Host: master,
		}
	}

	data, err := ioutil.ReadAll(sumFile)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		t.Logger().Error("sfs2008 empty results", "filesByHost", opts.FilesByHost)
		return &tool_errors.EmptyResultsError{
			File: sumFile,
			Host: master,
		}
	}

	t.Logger().Info("sfs2008 GetResults", "sumfile content", string(data))

	sumFileResults, err := parseAllResults(reSumFile, string(data))
	if err != nil {
		return err
	}

	var results []types.ToolUnifiedResults

	identifier := t.Params().Identifier
	name := string(t.Name())
	jobID := string(t.Params().JobID)
	for _, r := range sumFileResults {
		r := r // Create new instance of r for fresh referencing, for more see: docs/coding_conventions.md#bug-in-go---for-loop
		results = append(results, types.ToolUnifiedResults{
			Identifier:     &identifier,
			ToolName:       &name,
			RequestedIOps:  &r.RequestedIOPS,
			IOps:           &r.IOPS,
			AverageLatency: &r.Latency,
			JobID:          &jobID,
		})
	}

	err = jobs.UpdateToolResults(opts.JobID, results)
	if err != nil {
		return err
	}

	if opts.Conf.Tests.FailOnToolValidation {
		t.Logger().Debug("Validating results...")
		expected := expectedResults{
			result: result{
				Latency: opts.Conf.Sfs2008.Latency,
				IOPS:    opts.Conf.Sfs2008.LoadIo,
			},
			numberOfRuns: opts.Conf.Sfs2008.NumberOfRuns,
		}

		err = verify(string(data), expected)
		if err != nil {
			return &tool_errors.WrongThresholdError{
				Cause: err,
			}
		}
	}
	t.Logger().Debug("Successfully got results")
	return nil
}
