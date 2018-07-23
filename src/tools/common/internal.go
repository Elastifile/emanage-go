package common

import "types"

type InternalToolProperties struct {
	*ToolPropertiesBase
	client types.InternalClient
}

type DefaultInternalTool struct {
	*DefaultTool
	Props types.InternalProperties
}

func NewInternalProps(params *types.ToolParams, client types.InternalClient) (result *InternalToolProperties, err error) {
	base, err := NewProperties(params)
	if err != nil {
		return nil, err
	}
	return &InternalToolProperties{
		ToolPropertiesBase: base,
		client:             client,
	}, nil
}

func NewInternalTool(props types.InternalProperties) *DefaultInternalTool {
	return &DefaultInternalTool{
		DefaultTool: &DefaultTool{
			Props: props,
		},
		Props: props,
	}
}

func (t *DefaultInternalTool) Start(context *types.Context) error {
	props := t.Props
	return props.Client().RequestStart(
		props.Master(), context, props.Params().JobID, props.Config(),
	)
}

func (t *DefaultInternalTool) Wait(context *types.Context) error {
	props := t.Props
	logger.Info(
		"DefaultInternalTool: Request Wait",
		"master", props.Master(),
		"jobid", props.Params().JobID,
	)
	return props.Client().RequestWait(props.Master(), context, props.Params().JobID)
}

func (t *DefaultInternalTool) Cleanup(context *types.Context) error {
	props := t.Props
	return props.Client().RequestCleanup(context, props.Params().JobID)
}

func (t *DefaultInternalTool) Stop(context *types.Context) error {
	props := t.Props
	return props.Client().RequestStop(props.Master(), context, props.Params().JobID)
}

func (p *InternalToolProperties) Client() types.InternalClient {
	return p.client
}

func (t *InternalToolProperties) Master() types.Host {
	return types.Host(t.params.Config.Loaders()[0])
}
