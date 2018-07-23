package tools_api

import (
	"sync"

	"github.com/go-errors/errors"

	"sysapi"
	cthon_config "tools/cthon/config"
	erun_config "tools/erun/config"
	fio_config "tools/fio/config"
	fstool_config "tools/fstool/config"
	migration_config "tools/migration/config"
	sfs2008_config "tools/sfs2008/config"
	sfs2014_config "tools/sfs2014/config"
	vdbench_config "tools/vdbench/config"
	"types"
)

type pendingJobs struct {
	sync.Mutex
	jobs []types.JobID
}

type resultsMap map[string][]types.ToolUnifiedResults

var (
	pending        *pendingJobs = &pendingJobs{}
	sessionResults resultsMap
	parentID       types.JobID
)

func init() {
	sessionResults = make(resultsMap)
}

func SetParentID(id types.JobID) {
	parentID = id
}

func SessionResults() resultsMap {
	return sessionResults
}

// Filter the loaders and return only those with matching indexes
func filterLoaders(loaders []string, indexes []int) (targets []string, err error) {
	if len(indexes) == 0 {
		return loaders, err
	}

	for _, idx := range indexes {
		if idx >= len(loaders) {
			err = errors.Errorf("filterLoaders: given index is out loaders range"+
				"loaders: %v, indexes: %v", loaders, indexes)
			break
		}
		targets = append(targets, loaders[idx])
	}

	return targets, err
}

type toolClientBase struct {
	*Tools
	loaders []int
	exports []*sysapi.Export
	params  *types.ToolParams
	jobID   types.JobID
	goal    float64
	name    types.ToolName
	impl    toolClient
	stopped bool

	Identifier string
}

func PendingJobs() []types.JobID {
	return pending.jobs
}

func (t *toolClientBase) Start(identifiers ...string) (StopWaiter, error) {
	if t.stopped {
		return nil, errors.Errorf("starting stopped tool?")
	}

	if len(t.exports) > 0 {
		t.conf.Tesla.Elfs.Export = t.exports[0].Name()
	} else if t.conf.Tesla.Elfs.Export == "" {
		return nil, errors.Errorf("invalid empty ELFS export")
	}

	t.Identifier = string(t.name)
	if len(identifiers) > 0 {
		t.Identifier = identifiers[0]
	}

	targets, err := filterLoaders(t.system.System.Elab.LoaderIpsInternal(), t.loaders)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	logger.Action("Start tool", "name", t.name, "system", t.system.Name(), "targets", targets, "export", t.conf.Tesla.Elfs.Export)

	toolParams := &types.ToolParams{
		Identifier:    t.Identifier,
		ToolName:      t.name,
		Config:        t.conf,
		TargetLoaders: targets,
		ParentID:      parentID,
		System:        t.conf.Systems[0],
	}
	context := types.NewContextWithTimeout(Timeout)
	logger.Debug("Requesting tooly start",
		"loaders", toolParams.TargetLoaders,
		"tool config", toolParams.Config)
	jobID, err := t.tooly.RequestStart(context, toolParams)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	t.jobID = jobID
	logger.Info("Tool started",
		"host", t.tooly.Host(),
		"Name", t.name,
		"JobID", t.jobID,
		"system", t.system.Name(),
		"ELFS.Export", t.conf.Tesla.Elfs.Export,
	)

	pending.Lock()
	pending.jobs = append(pending.jobs, jobID)
	pending.Unlock()

	return t.impl, nil
}

func (t *toolClientBase) Wait() (resultWaiter, error) {
	tooly := t.tooly
	context := types.NewContextWithTimeout(TimeoutWait)

	args := []string{}
	if t.params != nil {
		args = t.params.Args
	}
	logger.Action("Waiting tool to finish ...", "name", t.name, "args", args, "job", t.jobID)
	exitcode, err := tooly.RequestWait(context, t.jobID)
	logger.Info("Tool done", "id", t.Identifier, "job", t.jobID)
	sessionResults[string(t.jobID)] = t.GetResults()
	if err != nil {
		logger.Error("Tool has failed", "tool", t.name, "JobID", t.jobID, "Error", err)
		return nil, errors.Wrap(err, 0)
	}

	if exitcode != types.ExitCodeSuccess {
		return nil, errors.Errorf(
			"Tool has completed with nonzero exit code: %v, tool: %v, JobID: %v",
			exitcode,
			t.name,
			t.jobID,
		)
	}
	return t.impl, nil
}

func (t *toolClientBase) Status() (*types.ToolStatus, error) {
	context := types.NewContextWithTimeout(Timeout)
	return t.tooly.RequestToolStatus(context, t.jobID)
}

func (t *toolClientBase) Stop() (waiter, error) {
	if t.stopped {
		return nil, errors.Errorf("stopping stopped tool?")
	}
	logger.Action("Client stops tool", "tool", t.name)

	t.stopped = true
	tooly := t.tooly
	context := types.NewContextWithTimeout(TimeoutWait)
	exitcode, err := tooly.RequestStop(context, t.jobID)
	if err != nil {
		logger.Error("Failed on tool", "tool", t.name, "JobID", t.jobID)
		return nil, errors.Wrap(err, 0)
	}

	if exitcode != types.ExitCodeSuccess {
		return nil, errors.Errorf(
			"Job has completed with nonzero exit code: %v, tool: %v, JobID: %v",
			exitcode,
			t.name,
			t.jobID,
		)
	}
	removeJob(t.jobID)
	return t.impl, nil
}

func (t *toolClientBase) SetLoaders(loaders ...int) ToolPropsCollector {
	t.loaders = loaders
	return t.impl
}

func (t *toolClientBase) SetSystem(system *sysapi.System) ToolPropsCollector {
	t.system = system
	t.conf.System = *t.system.System
	return t.impl
}

func (t *toolClientBase) SetExports(exports ...*sysapi.Export) ToolPropsCollector {
	if len(exports) != 1 {
		panic("must specify exactly one export")
	}

	t.exports = exports
	exp := t.exports[0]
	t.SetSystem(exp.System)
	logger.Info("Set tool",
		"name", t.name,
		"export", exp.EmanageExport().Name,
		"system", t.system.Name(),
	)
	logger.Debug("Set tool",
		"name", t.name,
		"export", *(exp.EmanageExport()),
		"elab", t.system.System.Elab,
	)

	return t.impl
}

func (t *toolClientBase) GetResults() []types.ToolUnifiedResults {
	var (
		err     error
		results []types.ToolUnifiedResults
	)
	context := types.NewContextWithTimeout(Timeout)
	// Some tools don't create jobs (e.g. Migration)
	if t.jobID != "" {
		results, err = t.tooly.RequestResults(context, t.jobID)
		if err != nil {
			panic(err)
		}
	}
	return results
}

func removeJob(id types.JobID) {
	pending.Lock()
	result := make([]types.JobID, len(pending.jobs)-1)
	j := 0
	for _, oid := range pending.jobs {
		if oid != id {
			result[j] = oid
			j++
		}
	}
	pending.jobs = result
	pending.Unlock()
}

////////////////////////////////////////////////////////////////////////////////

type erun struct {
	*toolClientBase
	toolConfig *erun_config.Config
}

////////////////////////////////////////////////////////////////////////////////

type sfs2008 struct {
	*toolClientBase
	toolConfig *sfs2008_config.Config
}

////////////////////////////////////////////////////////////////////////////////

type sfs2014 struct {
	*toolClientBase
	toolConfig *sfs2014_config.Config
}

////////////////////////////////////////////////////////////////////////////////

type fsTool struct {
	*toolClientBase
	toolConfig *fstool_config.Config
}

////////////////////////////////////////////////////////////////////////////////

type mig struct {
	*toolClientBase
	toolConfig *migration_config.Config
}

////////////////////////////////////////////////////////////////////////////////

type cthon struct {
	*toolClientBase
	toolConfig *cthon_config.Config
}

////////////////////////////////////////////////////////////////////////////////

type vdbench struct {
	*toolClientBase
	toolConfig *vdbench_config.Config
}

////////////////////////////////////////////////////////////////////////////////

type fio struct {
	*toolClientBase
	toolConfig *fio_config.Config
}
