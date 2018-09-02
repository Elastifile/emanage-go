package migration

import (
	"golang.org/x/crypto/ssh/agent"
	migration_api "migration/api"
	"github.com/elastifile/emanage-go/src/runtimeutil"
	"github.com/elastifile/emanage-go/src/tools/common"
	"github.com/elastifile/emanage-go/src/types"
)

type tool struct {
	*common.DefaultInternalTool
	*common.InternalToolProperties
}

var _ types.InternalProperties = (*tool)(nil)

func NewTool(params *types.ToolParams) (types.Tool, error) {
	params.ToolName = types.ToolName(runtimeutil.PackageName())

	props, err := common.NewInternalProps(
		params, migration_api.NewClient(agent.Msg()))
	if err != nil {
		return nil, err
	}

	result := &tool{nil, props}
	result.DefaultInternalTool = common.NewInternalTool(result)
	return result, nil
}

func (t *tool) Abort(ctx *types.Context) error {
	return t.Stop(ctx)
}

func (t *tool) Config() interface{} {
	return &t.Params().Config.Migration
}

func (t *tool) Name() types.ToolName {
	return types.ToolName("slave")
}
