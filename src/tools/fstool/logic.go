package fstool

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"containers"

	"github.com/elastifile/emanage-go/src/helputils"
	"github.com/elastifile/emanage-go/src/runtimeutil"
	"github.com/elastifile/emanage-go/src/tools/common"
	fstool_config "github.com/elastifile/emanage-go/src/tools/fstool/config"
	"github.com/elastifile/emanage-go/src/types"
)

const (
	logFileFmt = "out.%s.log"
)

type tool struct {
	sync.Mutex
	*common.DefaultExternalTool
	*common.ExternalToolProperties

	logAggregators map[types.Host]map[string]types.LogAggregator
}

var _ types.ExternalProperties = (*tool)(nil)

func NewTool(params *types.ToolParams) (types.Tool, error) {
	params.ToolName = types.ToolName(runtimeutil.PackageName())

	props, err := common.NewExternalProps(params)
	if err != nil {
		return nil, err
	}

	result := &tool{
		ExternalToolProperties: props,
		logAggregators:         map[types.Host]map[string]types.LogAggregator{},
	}
	result.DefaultExternalTool = common.NewExternalTool(result)
	return result, nil
}

func (t *tool) GetMasterCommand() (cmd []string, err error) {
	return t.GetSlaveCommand(t.Props.Master())
}

func (t *tool) GetSlaveCommand(host types.Host) (cmd []string, err error) {
	params := t.Params()
	conf := params.Config
	loaders := params.TargetLoaders
	fstConf := &conf.FsTool

	ldInx := helputils.FindStr(loaders, string(host))

	switch fstConf.Method {

	case "FillCapacity":
		fstConf.FillCapacity.EmanageIp = conf.EmanageServer()

	case "TreeCreate", "TreeUpdate":
		fstConf.TreeCreate.FileCount /= len(loaders)
		fstConf.TreeCreate.LinkCount /= len(loaders)
		fstConf.TreeCreate.NodeCount /= len(loaders)

	case "TreeDiff":
		fstConf.TreeDiff.OtherNfsClient.Loader = "." + strconv.Itoa(ldInx)
	}

	confNfs := &fstConf.NfsClient
	confNfs.Loader = "." + strconv.Itoa(ldInx)
	if conf.Tesla.Elfs.Export != "" {
		confNfs.Export = conf.Tesla.Elfs.Export
	}
	if params.System.Frontend != "" {
		confNfs.Frontend = params.System.Frontend
	}

	return fstConf.ToCmd(), nil
}

func (t *tool) WaitFor() (hosts []types.Host) {
	for _, loader := range t.Params().TargetLoaders {
		hosts = append(hosts, types.Host(loader))
	}
	return hosts
}

func (t *tool) GetResultFilesPatterns() []string {
	return []string{"*.log"}
}

func (t *tool) WaitStrategy(_ *types.Context) types.ErrorHandlingStrategy {
	return func(acc []error, next error) ([]error, bool) {
		if next != nil {
			if e, ok := next.(*containers.ContainerExitError); ok {
				e.Stdout = t.GetStdout(string(e.Host))
				ret := fstool_config.FCRet(e.ExitCode)
				switch ret {
				case fstool_config.FCRetErrorBadConfig:
					e.Message = "Couldn't parse tesla configuration from environment"
				case fstool_config.FCRetErrorNFSClient:
					e.Message = "Couldn't create NFS client"
				case fstool_config.FCRetErrorRootDir:
					e.Message = "Couldn't create root directory"
				case fstool_config.FCRetErrorFSInfo:
					e.Message = "Couldn't fetch file-system info"
				case fstool_config.FCRetErrorWorkDir:
					e.Message = "Couldn't create working directory"
				case fstool_config.FCRetErrorTool:
					e.Message = "Tool failed"
				case fstool_config.FCRetErrorPanic:
					e.Message = "Tool paniced"
				default:
					e.Message = fmt.Sprintf("Unknown exit code: (%d)", e.ExitCode)
				}
			}
			return append(acc, next), true
		}
		return acc, true
	}
}

func (t *tool) NewLogAggregator(cnt *types.Container) types.LogAggregator {
	t.Lock()
	defer t.Unlock()
	hdAgg := common.NewHeadAggregator(0)
	if t.logAggregators[cnt.Host] == nil {
		t.logAggregators[cnt.Host] = map[string]types.LogAggregator{}
	}
	t.logAggregators[cnt.Host][cnt.ID] = hdAgg
	return hdAgg
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
