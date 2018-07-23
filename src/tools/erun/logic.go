package erun

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"sync"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/go-errors/errors"

	"containers"
	"helputils"
	"ioutil2"
	"jobs"
	"logging"
	"runtimeutil"
	"tools/common"
	erun_config "tools/erun/config"
	"types"
)

var logger = logging.NewLogger("erun")

const (
	resultsFile    = "results.json"
	resultsFullLog = "erun_full_log.json"
	logFileFmt     = "stdout.%s.log"
)

//go:generate stringer -type=ErunRet

type ErunRet int

const (
	ErunRetSuccess             ErunRet = 0
	ErunRetErrorCommandLine    ErunRet = 1
	ErunRetErrorPrep           ErunRet = 2
	ErunRetErrorClient         ErunRet = 3
	ErunRetErrorLongIo         ErunRet = 4
	ErunRetErrorParams         ErunRet = 5
	ErunRetErrorDataVerifError ErunRet = 6
)

type tool struct {
	sync.Mutex
	*common.DefaultLoggingTool
	*common.ExternalToolProperties

	erunPath          string
	useHostNetworking bool
	logAggregators    map[types.Host]map[string]types.LogAggregator
}

var _ types.LoggingProperties = (*tool)(nil)

func NewTool(params *types.ToolParams) (types.Tool, error) {
	if err := common.ValidateExports(params); err != nil {
		return nil, err
	}

	// validate required config files are in filestore
	if err := common.ValidateConfigFiles(params.ConfigFilesBucket, []string{
		params.Config.Erun.ConfigFile}); err != nil {
		return nil, err
	}
	params.ToolName = types.ToolName(runtimeutil.PackageName())

	// Most profiles should run only on one loader (they don't coordinate multiple instances,
	// and do not test performance).
	// We explicitly specify that some profiles should run on multiple loaders:
	singleLoader := true
	switch params.Config.Erun.Profile.Name() {
	case (*erun_config.ProfileIO).Name(nil):
		singleLoader = false
	}

	if singleLoader {
		// Setting TargetLoaders to a single random loader.
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := rand.Intn(len(params.Config.Loaders()))
		randLoader := string(params.Config.Loaders()[idx])
		fmt.Printf("Erun will run on random loader: %v (index: %v)\n", randLoader, idx)
		params.TargetLoaders = []string{randLoader}
		params.Config.SetLoaders(params.TargetLoaders...)
	}

	props, err := common.NewExternalProps(params)
	if err != nil {
		return nil, err
	}

	result := &tool{
		ExternalToolProperties: props,
		logAggregators:         map[types.Host]map[string]types.LogAggregator{},
		erunPath:               "/usr/bin/erun",
	}
	result.DefaultLoggingTool = common.NewLoggingTool(result)

	// Handle Metabench config file
	if mbParams, ok := params.Config.Erun.Profile.(*erun_config.ProfileMetabench); ok {
		if mbParams.Description == "" {
			mbParams.Description = "Tesla Metabench test"
		}

		configFile, e := result.ParseTemplate(&common.ParseOpts{
			SourceBucket: params.ConfigFilesBucket,
			ConfigFile:   params.Config.Erun.ConfigFile,
			Template:     metabenchTemplate,
			Data:         mbParams,
		})
		if e != nil {
			return nil, e
		}
		params.Config.Erun.ConfigFile = configFile
	}

	return result, nil
}

func (t *tool) GetMasterCommand() ([]string, error) {
	return t.GetSlaveCommand(t.Master())
}

func (t *tool) GetSlaveCommand(host types.Host) ([]string, error) {
	params := t.Params()
	conf := params.Config

	jobSessPath := filepath.Join(params.FilerSessionPath(), string(host))

	cmd := conf.Erun.Profile.ToFullCmd(&erun_config.ErunFormatCmdOpts{
		Conf:             conf.Erun,
		Host:             string(host),
		DataStorePath:    t.DataStorePath(),
		JobID:            string(params.JobID),
		Frontend:         params.System.Frontend,
		Export:           conf.Tesla.Elfs.Export,
		ResultFile:       resultsFile,
		ResultsFullLog:   resultsFullLog,
		ExecPath:         t.erunPath,
		FilerSessionPath: jobSessPath,
	})

	t.Logger().Info(fmt.Sprintf("erun cmd: %v", cmd),
		"FilerSessionPath", jobSessPath,
	)

	return append(t.Shell, cmd), nil
}

func (t *tool) WaitFor() (hosts []types.Host) {
	return t.Params().Config.Loaders()
}

func (t *tool) WaitStrategy(context *types.Context) types.ErrorHandlingStrategy {
	t.Logger().Info("Running WaitStrategy...")
	return func(acc []error, next error) ([]error, bool) {
		for {
			if e, ok := next.(*errors.Error); ok {
				next = e.Err
			} else {
				break
			}
		}
		if e, ok := next.(*containers.ContainerExitError); ok {
			e.Stdout = t.GetStdout(string(e.Host))
			e.SessionDir = t.Params().FilerSessionPath()
			erunRet := ErunRet(e.ExitCode)
			switch erunRet {
			case ErunRetErrorLongIo:
				// This is actually a warning, and not an error. The run should be considered successful.
				// TODO: Somehow report this to the test so the end user can see this message:
				t.Logger().Warn("Erun alert: long latency", "Host", e.Host)
				return acc, true
			case ErunRetErrorDataVerifError:
				e.Message = "Erun: data integrity error"
				all := append(acc, next)
				return []error{common.Abort(context, all...)}, false
			default:
				e.Message = fmt.Sprintf("Erun error: %d (%[1]v)", erunRet)
				return append(acc, next), true
			}
		}
		if next != nil {
			return append(acc, next), true
		}
		return acc, true
	}
}

func (t *tool) GetDockerHostConfig() (c *docker.HostConfig) {
	hostConfig := &docker.HostConfig{
		CapAdd: []string{"SYS_ADMIN"},
	}

	if t.useHostNetworking {
		hostConfig.NetworkMode = "host"
	}

	hostConfig.Binds = append(
		hostConfig.Binds,
		"/var/log:/var/log",
		"/dev/shm:/dev/shm",
		fmt.Sprintf("%v:/usr/bin/erun:ro", t.erunPath),
	)

	return hostConfig
}

func (t *tool) GetResultFilesPatterns() []string {
	return []string{resultsFile, resultsFullLog, fmt.Sprintf(logFileFmt, "*")}
}

func readLastResult(bulk *types.NamedReader) (*types.NamedReader, error) {
	ctx, _ := ioutil2.WithDeadline(nil, time.Now().Add(5*time.Minute))
	timedReader := ioutil2.NewTimedReader(bulk, ctx)
	scanner := bufio.NewScanner(timedReader)
	last := ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 1 && line[0] == '{' && line[len(line)-1] == '}' {
			last = line
		}
	}
	return types.NewNamedReader(bulk.Name, strings.NewReader(last)), ctx.Err()
}

func selectResults(possibleResults []*types.NamedReader) (*types.NamedReader, error) {
	var (
		final        *types.NamedReader
		intermediate *types.NamedReader
		err          error
	)

	logger.Info("Selecting results to read...")
	for _, candidate := range possibleResults {
		if candidate.Name == resultsFile {
			final = candidate
		} else if candidate.Name == resultsFullLog {
			intermediate, err = readLastResult(candidate)
		}
	}
	if final == nil {
		return intermediate, err
	}
	return final, err
}

func (t *tool) GetResults(opts *types.ResultOpts) error {
	if !t.Params().Config.Erun.Profile.HasResults() {
		return nil
	}
	logger.Info(
		"Processing erun results",
		"profile", t.Params().Config.Erun.Profile.Name(),
		"profile-type", fmt.Sprintf("%T", t.Params().Config.Erun.Profile),
		"FilesByHost", opts.FilesByHost,
		"opts.Conf.Loaders()", opts.Conf.Loaders(),
	)

	var hosts []types.Host
	data := [][]byte{}

	for host, files := range opts.FilesByHost {
		if helputils.ContainsStr(types.HostStrings(opts.Conf.Loaders()), string(host)) { // refer only active loaders
			if len(files) == 0 {
				return errors.Errorf("Missing results on host: %v", host)
			}

			hosts = append(hosts, host)
			actualResults, err := selectResults(files)
			if actualResults == nil {
				return errors.Errorf("Results not found in files: %+v", files)
			} else if err != nil {
				return errors.Errorf("Couldn't read Erun's results because: %s", err)
			}
			d, e := ioutil.ReadAll(actualResults)
			if e != nil {
				return errors.Errorf("Failed reading file results for host: %v current file: %v", host, files[0])
			}
			data = append(data, d)
		}
	}

	loaders := helputils.StringSet{}
	loaders.Add(types.HostStrings(opts.Conf.Loaders())...)

	if len(hosts) != len(loaders) {
		return errors.Errorf("Missing results. got results only for hosts: %v expected: %v", hosts, opts.Conf.Loaders())
	}

	res, err := parseResults(t, data)
	if err != nil {
		return err
	}

	return jobs.UpdateToolResults(opts.JobID, []types.ToolUnifiedResults{res})
}

func (t *tool) NewLogAggregator(c *types.Container) types.LogAggregator {
	t.Lock()
	defer t.Unlock()
	const lines = 100
	ha := common.NewLabeledAggregator(
		fmt.Sprintf("---- Head: (%d)", lines),
		common.NewHeadAggregator(lines),
	)
	ta := common.NewLabeledAggregator(
		fmt.Sprintf("---- Tail: (%d)", lines),
		common.NewTailAggregator(lines),
	)
	ca := common.CombinedAggregator{ha, ta}
	if t.logAggregators[c.Host] == nil {
		t.logAggregators[c.Host] = map[string]types.LogAggregator{}
	}
	t.logAggregators[c.Host][c.ID] = ca
	return ca
}

func (t *tool) Logs() map[types.Host]*types.NamedReader {
	result := map[types.Host]*types.NamedReader{}
	for host, agg := range t.logAggregators {
		for id, ag := range agg {
			trunc := 6
			if len(id) < trunc {
				trunc = len(id)
			}
			result[host] = types.NewNamedReader(
				fmt.Sprintf(logFileFmt, "cid-"+id[:trunc]),
				strings.NewReader(ag.String()),
			)
		}
	}
	return result
}
