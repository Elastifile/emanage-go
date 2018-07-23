package common

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/go-errors/errors"
	multierror "github.com/hashicorp/go-multierror"

	"agent"
	"config"
	"containers"
	"filestore"
	filestore_api "filestore/api"
	filestore_types "filestore/types"
	"jobs"
	"logging"
	"orchestrator"
	"remote"
	tool_errors "tools/errors"
	"types"
)

var logger = logging.NewLogger("common")

type ExternalToolProperties struct {
	*ToolPropertiesBase
	Shell []string
}

type DefaultExternalTool struct {
	*DefaultTool
	Props types.ExternalProperties
}

type containerAction func(types.Host, string) error

func NewExternalProps(params *types.ToolParams) (result *ExternalToolProperties, err error) {
	base, err := NewProperties(params)
	if err != nil {
		return nil, err
	}
	result = &ExternalToolProperties{
		ToolPropertiesBase: base,
		Shell:              []string{"/bin/sh", "-x", "-c"},
	}

	if params.Config.Tesla.Devel.Fake.ToolSim {
		result.Shell = append([]string{"toolsim"}, result.Shell...)
	}

	return result, nil
}

func NewExternalTool(props types.ExternalProperties) *DefaultExternalTool {
	result := &DefaultExternalTool{
		DefaultTool: &DefaultTool{
			Props: props,
		},
		Props: props,
	}
	return result
}

func (t *ExternalToolProperties) GetDockerHostConfig() *docker.HostConfig {
	return &docker.HostConfig{
		CapAdd:      []string{"SYS_ADMIN"},
		NetworkMode: "host",
	}
}

func (t *ExternalToolProperties) WaitFor() []types.Host {
	return []types.Host{t.Master()}
}

func (t *ExternalToolProperties) Master() types.Host {
	return types.Host(t.params.Config.Loaders()[0])
}

func (t *ExternalToolProperties) GetImage() *types.Image {
	result := config.GetToolImage(t.params.ToolName)
	return &result
}

func (t *ExternalToolProperties) WaitStrategy(_ *types.Context) types.ErrorHandlingStrategy {
	return tool_errors.AccumulateStrategy
}

func (t *ExternalToolProperties) GetStdout(host string) string {
	conf := t.Params().Config
	jobID := t.Params().JobID
	toolName := t.Params().ToolName
	var prefix string
	if conf.Tesla.Jenkins.BuildUrl != "" {
		prefix = fmt.Sprintf("%s/artifact/", conf.Tesla.Jenkins.BuildUrl)
	}
	return fmt.Sprintf("%sresults/%s/%s/%s/",
		prefix, jobID, toolName, host)
}

func (t *DefaultExternalTool) Start(context *types.Context) (err error) {
	props := t.Props
	params := props.Params()
	tlogger := props.Logger()
	conf := agent.Config()
	loaders := conf.Loaders()
	jobID := params.JobID
	path := t.DataStorePath()

	filestoreHost := types.Host(params.System.Setup.Filestore)

	fs, err := filestore.New(filestoreHost, false)
	if err != nil {
		return err
	}
	api_fs, err := filestore_api.NewClient(agent.Msg(), filestoreHost, false)
	if err != nil {
		return err
	}

	srcBucket := params.ConfigFilesBucket
	// set to the job's dedicated bucket
	params.ConfigFilesBucket = filestore_types.Bucket(path).TrimDataVolume().ToBucketName()
	tlogger.Debug("Switching the job's config bucket from its parent to its dedicated bucket "+
		"(and copy its content to the dedicated job)", "parentBucket", srcBucket, "bucket", params.ConfigFilesBucket)

	// ensure bucket exists
	tlogger.Debug("Ensuring job's bucket exists", "bucket", params.ConfigFilesBucket)
	if e := fs.MakeBucket(params.ConfigFilesBucket); e != nil {
		return e
	}

	if srcBucket == "" {
		tlogger.Info("Missing config bucket to the tool's job (params.ConfigFilesBucket) might be sanity mode")
	} else {
		tlogger.Debug("Copying temporary/parent bucket content to the tool's job dedicated bucket",
			"srcBucket", srcBucket, "dstBucket", params.ConfigFilesBucket)
		if err = fs.CopyBucket(srcBucket, params.ConfigFilesBucket); err != nil {
			return err
		}
	}

	// will also ensure all local dir are created on remotes
	tlogger.Debug("Storing files to hosts from job's config bucket", "bucket", params.ConfigFilesBucket, "hosts", loaders)
	if err = api_fs.RequestStoreFiles(params.ConfigFilesBucket, path, loaders); err != nil {
		return err
	}

	containers, err := t.startContainers(context)
	if err != nil {
		return err
	}

	logger.Info("Starting loggers for containers")
	if err := t.startLoggers(containers); err != nil {
		return err
	}
	logger.Info("Loggers for containers started")

	for _, cont := range containers {
		logger.Debug("Adding container", "container", cont, "job", jobID, "tool", props.Name())
		jobs.AddContainer(jobID, *cont)
	}

	if conf.Tesla.ShellUser != "" {
		remote.DefaultUser = conf.Tesla.ShellUser
	}

	logger.Info("conf.Tesla.Monitoring.IpNeighbours", "enabled?", conf.Tesla.Monitoring.IpNeighbours)

	if conf.Tesla.Monitoring.IpNeighbours {
		logger.Debug("logging ip neighbours, before", "job", jobID)
		err := t.logIpNeighbours("before")
		if err != nil {
			logger.Warn("Failed logging ip neigh", "err", err)
		}
	}

	return nil
}

func (t *DefaultExternalTool) logIpNeighbours(stage string) error {
	errChan := make(chan error)

	jid := t.Props.Params().JobID
	job, err := jobs.ByID(jid)
	if err != nil {
		return err
	}

	hosts := make([]string, len(job.Containers))

	for i, cont := range job.Containers {
		id := cont.ID
		hosts[i] = string(cont.Host)
		host := hosts[i]

		go func() {
			rem, err := remote.NewCancelableRemote(host, time.Minute)
			if err != nil {
				errChan <- err
				return
			}

			dir, err := t.getContainerDataDir(rem, id)
			if err != nil {
				errChan <- errors.Wrap(err, 0)
				return
			}

			dsSubpath := strings.Join(strings.Split(t.DataStorePath(), "/")[2:], "/")
			file := filepath.Join(dir, dsSubpath, "ip-neigh-"+stage+".log")
			out, err := rem.RunWithTimeout("sudo ip -s neigh > "+file, time.Minute)
			if err != nil {
				errChan <- errors.WrapPrefix(err, out, 0)
				return
			}

			errChan <- nil
		}()
	}

	err = tool_errors.CollectAll(errChan, len(hosts))
	if err != nil {
		return err
	}

	logger.Info("Logged ip neighbours", "stage", stage, "job", jid, "hosts", hosts)
	return nil
}

func (t *DefaultExternalTool) getContainerDataDir(rem *remote.Remote, id string) (string, error) {
	out, err := rem.RunWithTimeout(fmt.Sprintf("docker inspect %s | grep _data", id), time.Minute)
	if err != nil {
		return "", err
	}

	inspJson := "{" + strings.Replace(strings.TrimSpace(out), ",", "}", -1)
	type inspect struct {
		Source string `json:"Source"`
	}
	var insp inspect

	if err := json.Unmarshal([]byte(inspJson), &insp); err != nil {
		return "", errors.WrapPrefix(err, inspJson, 0)
	}

	return insp.Source, nil
}

func (t *DefaultExternalTool) Abort(context *types.Context) error {
	job, err := jobs.ByID(t.Props.Params().JobID)
	if err != nil {
		return err
	}
	go t.doContainers(orchestrator.Stop, job.Containers)
	errs := tool_errors.FoldErrors(tool_errors.AccumulateStrategy, t.Props.ErrorWatcher(), len(job.Containers))
	return tool_errors.Collect(errs...)
}

func (t *DefaultExternalTool) DataStorePath() string {
	return PkgDataStorePath(t.Props.Params().JobID, string(t.Props.Name()))
}

func (t *DefaultExternalTool) Wait(context *types.Context) error {
	conf := agent.Config()
	props := t.Props
	params := props.Params()
	jobID := params.JobID

	errs := newWaitHandler(t, context).do()
	t.Props.Logger().Info("Wait finished", "errs", errs)

	if conf.Tesla.Monitoring.IpNeighbours {
		logger.Debug("logging ip neighbours, after", "job", jobID)
		err := t.logIpNeighbours("after")
		if err != nil {
			logger.Warn("Failed logging ip neigh", "err", err)
		}
	}

	return tool_errors.Collect(tool_errors.NonNil(errs...)...)
}

func (t *DefaultExternalTool) joinWait(containers []types.Container, context *types.Context, errHandler types.ErrorHandlingStrategy) []error {
	go t.doContainers(orchestrator.Wait, containers)
	return tool_errors.FoldErrors(errHandler, t.Props.ErrorWatcher(), len(containers))
}

func isIgnorable(err error) (result bool) {
	defer logger.Info("Will ignore", "error", err, "ignore", result)
	var (
		merror *multierror.Error
		ok     bool
	)
	if err == nil {
		return true
	}
	if oerror, ok := err.(*errors.Error); ok {
		return isIgnorable(oerror.Err)
	}
	if merror, ok = err.(*multierror.Error); !ok {
		merror = multierror.Append(err)
	}
	result = true
	for _, cErr := range merror.Errors {
		if _, ok = cErr.(*containers.ContainerExitError); !ok {
			result = false
		}
	}
	return result
}

func (t *DefaultExternalTool) jobContainers() ([]types.Container, error) {
	job, err := jobs.ByID(t.Props.Params().JobID)
	if err != nil {
		return nil, err
	}
	return job.Containers, nil
}

func (t *DefaultExternalTool) Cleanup(context *types.Context) error {
	containers, err := t.jobContainers()
	if err != nil {
		return err
	}
	go t.doContainers(orchestrator.Stop, containers)
	return tool_errors.CollectAll(t.Props.ErrorWatcher(), len(containers))
}

func (t *DefaultExternalTool) Stop(context *types.Context) error {
	containers, err := t.jobContainers()
	if err != nil {
		return err
	}
	if err := t.trySaveLogs(); err != nil {
		logger.Error("Couldn't save logs", "error", err)
	}
	go t.doContainers(orchestrator.Stop, containers)
	return tool_errors.CollectAll(t.Props.ErrorWatcher(), len(containers))
}

func (t *DefaultExternalTool) trySaveLogs() error {
	watcher := make(chan error)

	client, err := filestore.New(t.Props.Params().System.FilestoreHostInternal(), false)
	if err != nil {
		return err
	}

	path := t.DataStorePath()
	bucket := filestore_types.Bucket(path).TrimDataVolume().ToBucketName()
	logs := t.Props.Logs()
	logger.Info("Will save logs into", "bucket", bucket, "logs", len(logs))
	for host, file := range logs {
		if host == "" {
			logger.Warn("Empty host in logs", "host", host, "file", file)
		}
		host := host
		file := file
		go func() {
			b := bucket.ByHost(string(host))
			if err := client.MakeBucket(b); err != nil {
				watcher <- err
			} else {
				watcher <- client.Upload(b, file)
			}
		}()
	}
	return tool_errors.CollectAll(watcher, len(logs))
}

func waitingContainers(job *types.Job, hosts []types.Host) []types.Container {
	var result []types.Container
	hs := map[types.Host]bool{}
	for _, host := range hosts {
		hs[host] = true
	}
	for _, container := range job.Containers {
		if hs[container.Host] {
			result = append(result, container)
		}
	}
	return result
}

func (t *DefaultExternalTool) doContainers(action containerAction, containers []types.Container) {
	done := make(chan bool)

	for _, container := range containers {
		container := container
		go func() {
			t.Props.ErrorWatcher() <- action(container.Host, container.ID)
			done <- true
		}()
	}
	for range containers {
		<-done
	}
}

func (t *DefaultExternalTool) newContainerStartOpts(host types.Host, cmd []string) *types.ContainerStartOpts {
	return &types.ContainerStartOpts{
		Host:             host,
		Image:            t.Props.GetImage(),
		Cmd:              cmd,
		JobID:            t.Props.Params().JobID,
		DataStorePath:    t.DataStorePath(),
		DockerHostConfig: t.Props.GetDockerHostConfig(),
	}
}

func (t *DefaultExternalTool) startContainers(context *types.Context) ([]*types.Container, error) {
	props := t.Props
	config := &props.Params().Config
	cchan := make(chan *types.Container)
	var (
		slaves     []*types.ContainerStartOpts
		containers []*types.Container
	)

	logger.Debug("startContainers", "Loaders", props.Params().TargetLoaders)
	for _, loader := range props.Params().TargetLoaders {
		host := types.Host(loader)
		if host == t.Props.Master() {
			continue
		}
		cmd, cerr := props.GetSlaveCommand(host)
		if cerr != nil {
			return nil, cerr
		}
		if cmd != nil {
			slaves = append(slaves, t.newContainerStartOpts(host, cmd))
		}
	}
	for _, slave := range slaves {
		slave := slave
		go func() {
			container, err := orchestrator.Start(config, slave)
			props.ErrorWatcher() <- err
			cchan <- container
		}()
	}
	for range slaves {
		err := <-props.ErrorWatcher()
		container := <-cchan
		containers = append(containers, container)
		if err != nil {
			return containers, err
		}
	}
	cmd, cerr := props.GetMasterCommand()
	logger.Info("Tooly Command to run", "cmd", cmd)
	if cerr != nil {
		return containers, cerr
	}
	if cmd != nil {
		container, err := orchestrator.Start(
			config,
			t.newContainerStartOpts(t.Props.Master(), cmd),
		)
		if err != nil {
			return containers, err
		}
		containers = append(containers, container)
	}
	return containers, nil
}

func (t *DefaultExternalTool) startLoggers(containers []*types.Container) error {
	type containerReader struct {
		c *types.Container
		r io.ReadCloser
	}
	watcher := make(chan error)
	logs := make(chan containerReader)
	for _, c := range containers {
		c := c
		go func() {
			rc, err := orchestrator.Logs(c.Host, c.ID)
			watcher <- err
			logs <- containerReader{
				r: rc,
				c: c,
			}
		}()
	}
	if err := tool_errors.CollectAll(watcher, len(containers)); err != nil {
		return err
	}
	for range containers {
		cr := <-logs
		la := t.Props.NewLogAggregator(cr.c)
		go func() {
			scanner := bufio.NewScanner(cr.r)
			for scanner.Scan() {
				la.ProcessLine(scanner.Text())
			}
		}()
	}
	return nil
}

type waitHandler struct {
	tool    *DefaultExternalTool
	context *types.Context
	next    func(*types.Job) []error
	errors  []error
}

func newWaitHandler(tool *DefaultExternalTool, context *types.Context) *waitHandler {
	result := &waitHandler{
		tool:    tool,
		context: context,
	}
	result.next = result.waitMaster
	return result
}

func (w *waitHandler) waitMaster(job *types.Job) []error {
	w.next = w.waitStop
	return w.joinWait(
		waitingContainers(job, w.tool.Props.WaitFor()),
		w.context,
		w.tool.Props.WaitStrategy(w.context),
	)
}

func (w *waitHandler) waitStop(_ *types.Job) []error {
	w.next = w.waitSlaves
	if tp, ok := w.tool.Props.(types.ToolProtocol); ok {
		return []error{tp.Stop(w.context)}
	}
	return []error{w.tool.Stop(w.context)}
}

func (w *waitHandler) waitSlaves(job *types.Job) []error {
	w.next = nil
	return w.joinWait(
		job.Containers,
		w.context,
		tool_errors.FilterStrategy(isIgnorable),
	)
}

func (w *waitHandler) joinWait(containers []types.Container, context *types.Context, errHandler types.ErrorHandlingStrategy) []error {
	go w.tool.doContainers(orchestrator.Wait, containers)
	return tool_errors.FoldErrors(errHandler, w.tool.Props.ErrorWatcher(), len(containers))
}

func (w *waitHandler) do() []error {
	for w.next != nil {
		job, err := jobs.ByID(w.tool.Props.Params().JobID)
		if err != nil {
			w.errors = append(w.errors, err)
			break
		}
		if job.Status == types.JobStatusAborted {
			break
		}
		w.errors = append(w.errors, w.next(&job)...)
	}
	return w.errors
}
