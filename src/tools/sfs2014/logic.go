package sfs2014

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/go-errors/errors"

	"jobs"
	"github.com/elastifile/emanage-go/src/runtimeutil"
	"github.com/elastifile/emanage-go/src/tools/common"
	"github.com/elastifile/emanage-go/src/tools/errors"
	"github.com/elastifile/emanage-go/src/types"
)

var resultSuffix = "tesla"
var resultFilePrefix = "sfssum"

// Execution commands
var baseCmd = []string{"/bin/sh", "-c"}
var symlinkResults = "ln -s %s /opt/sfs2014/sfs2014/results && "
var mountCmd = "mount -o soft %s:%s /mnt/ && "
var sshdBackground = "/usr/sbin/sshd &&"

var waitLoadersCmd = `/wait_ssh.sh "%s" && `
var masterCmd = "python /opt/sfs2014/sfs2014/bin/SfsManager --debug -r %s -s %s"

var slaveCmd = "/usr/sbin/sshd && sleep infinity"

type tool struct {
	*common.DefaultLoggingTool
	*common.ExternalToolProperties
}

func NewTool(params *types.ToolParams) (types.Tool, error) {
	if err := common.ValidateExports(params); err != nil {
		return nil, err
	}
	params.ToolName = types.ToolName(runtimeutil.PackageName())

	// validate required config files are in minio bucket
	if err := common.ValidateConfigFiles(params.ConfigFilesBucket, []string{
		params.Config.Sfs2014.ConfigFile}); err != nil {
		return nil, err
	}

	props, err := common.NewExternalProps(params)
	if err != nil {
		return nil, err
	}
	result := &tool{nil, props}
	result.DefaultLoggingTool = common.NewLoggingTool(result)
	configFile, err := result.ParseTemplate(&common.ParseOpts{
		SourceBucket: params.ConfigFilesBucket,
		ConfigFile:   params.Config.Sfs2014.ConfigFile,
		Template:     template,
		Data:         params.Config,
	})
	if err != nil {
		return nil, err
	}
	params.Config.Sfs2014.ConfigFile = configFile
	return result, nil
}

func (t *tool) GetMasterCommand() (cmd []string, err error) {
	params := t.Params()
	conf := params.Config
	path := t.DataStorePath()
	configFile := filepath.Join(path, conf.Sfs2014.ConfigFile)

	cmd = append(baseCmd, slaveCmd)
	ln := fmt.Sprintf(symlinkResults, path)
	mount := fmt.Sprintf(mountCmd, params.System.Frontend, conf.Tesla.Elfs.Export)
	waitLoaders := fmt.Sprintf(waitLoadersCmd, strings.Join(types.HostStrings(conf.Loaders())[1:], " "))
	exec := fmt.Sprintf(masterCmd, configFile, resultSuffix)
	cmd = append(baseCmd, ln+mount+sshdBackground+waitLoaders+exec)
	return cmd, nil
}

func (t *tool) GetSlaveCommand(host types.Host) (cmd []string, err error) {
	params := t.Params()
	conf := &params.Config
	cmd = append(baseCmd, slaveCmd)
	mount := fmt.Sprintf(mountCmd, params.System.Frontend, conf.Tesla.Elfs.Export)
	cmd = append(baseCmd, mount+slaveCmd)
	return cmd, nil
}

func (t *tool) GetResultFilesPatterns() []string {
	return []string{
		resultFilePrefix + "*",
	}
}

func (t *tool) GetDockerHostConfig() (c *docker.HostConfig) {
	hostConfig := t.ExternalToolProperties.GetDockerHostConfig()
	hostConfig.Binds = append(hostConfig.Binds, "/var/log:/var/log")
	return hostConfig
}

func (t *tool) GetResults(opts *types.ResultOpts) (err error) {
	master := t.Master()

	// TODO - also get the container's exit code (see at orchestrator/api 'RequestWait' func)
	files, ok := opts.FilesByHost[master]
	if !ok {
		return &tool_errors.NoSfsResultsError{
			Host:   master,
			Suffix: resultSuffix,
		}
	}

	t.Logger().Debug("opts.FilesByHost", "value", opts.FilesByHost)
	sumFile := &types.NamedReader{}
	for _, f := range files {
		if strings.HasSuffix(f.Name, "xml") {
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

	t.Logger().Debug("sfs2014 GetResults", "sumfile content", string(data))

	sumFileResults, err := unmarshalXML(data)
	if err != nil {
		return errors.Errorf("Failure, failed parsing sfs2014 sum file xml. err: %v", err)
	}

	var results []types.ToolUnifiedResults
	for _, run := range sumFileResults.Runs {
		var averageLatency float64
		var iops, requestedIops int

		for _, metric := range run.Metrics {
			switch metric.Name {
			case "op rate":
				iops = int(metric.Value)
			case "achieved rate":
				requestedIops = int(metric.Value)
			case "average latency":
				averageLatency = metric.Value
			}
		}
		// `run' is always a fresh struct in this loop,
		// so is the Benchmakr and the Name.
		// I don't think you need to copy here, but I'm curious to
		// know why you did this. (Oleg)
		benchmark := run.Benchmark.Name //create a copy

		identifier := t.Params().Identifier
		name := string(t.Name())
		jobID := string(t.Params().JobID)
		results = append(results, types.ToolUnifiedResults{
			Identifier:     &identifier,
			ToolName:       &name,
			IOps:           &iops,
			RequestedIOps:  &requestedIops,
			AverageLatency: &averageLatency,
			Workload:       &benchmark,
			JobID:          &jobID,
		})
	}

	err = jobs.UpdateToolResults(opts.JobID, results)
	if err != nil {
		return err
	}

	return nil
}
