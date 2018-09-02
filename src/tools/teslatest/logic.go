package teslatest

import (
	"config"
	"github.com/elastifile/emanage-go/src/runtimeutil"
	"github.com/elastifile/emanage-go/src/tools/common"
	"github.com/elastifile/emanage-go/src/types"
)

type tool struct {
	*common.DefaultLoggingTool
	*common.ExternalToolProperties
	junitXML string
}

var _ types.LoggingProperties = (*tool)(nil)

func NewTool(params *types.ToolParams) (types.Tool, error) {
	if err := common.ValidateArgs(params); err != nil {
		return nil, err
	}

	params.ToolName = types.ToolName(runtimeutil.PackageName())

	props, err := common.NewExternalProps(params)
	if err != nil {
		return nil, err
	}
	result := &tool{
		ExternalToolProperties: props,
		junitXML:               "junit.xml",
	}
	result.DefaultLoggingTool = common.NewLoggingTool(result)
	return result, nil
}

func (t *tool) WaitFor() (hosts []types.Host) {
	return t.Params().Config.Loaders()
}

func (t *tool) GetMasterCommand() (cmd []string, err error) {
	return t.Params().Args, nil
}

func (t *tool) GetSlaveCommand(host types.Host) (cmd []string, err error) {
	return nil, nil
}

func (t *tool) GetImage() *types.Image {
	result := config.GetImage(config.NewAPIImage)
	confDocker := &t.Props.Params().Config.Tesla.Docker
	if confDocker.PushPull.ImageTag != "" {
		result.SetTag(confDocker.PushPull.ImageTag)
	}
	return &result
}

func (t *tool) GetResultFilesPatterns() []string {
	return []string{t.junitXML, "*-failure-report.log", "tesla.*.log"}
}

func (t *tool) GetResults(opts *types.ResultOpts) error {
	return nil
}
