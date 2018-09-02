package cthon

import (
	"bytes"
	"text/template"

	"github.com/elastifile/emanage-go/src/runtimeutil"
	"github.com/elastifile/emanage-go/src/tools/common"
	cthon_config "github.com/elastifile/emanage-go/src/tools/cthon/config"
	tool_errors "github.com/elastifile/emanage-go/src/tools/errors"
	"github.com/elastifile/emanage-go/src/types"
)

type toolCmd struct {
	MountPoint    string
	RemoteFs      string
	ElfsFrontend  string
	Tests         *cthon_config.CthonTests
	CurrentLoader string
}

type tool struct {
	*common.DefaultExternalTool
	*common.ExternalToolProperties
}

var _ types.ExternalProperties = (*tool)(nil)

func NewTool(params *types.ToolParams) (types.Tool, error) {
	params.ToolName = types.ToolName(runtimeutil.PackageName())

	props, err := common.NewExternalProps(params)
	if err != nil {
		return nil, err
	}

	result := &tool{nil, props}
	result.DefaultExternalTool = common.NewExternalTool(result)
	return result, nil
}

func (t *tool) GetMasterCommand() (cmd []string, err error) {
	return t.GetSlaveCommand(t.Props.Master())
}

func (t *tool) GetSlaveCommand(host types.Host) (cmd []string, err error) {
	params := t.Params()
	conf := params.Config
	tests := t.ensureTests()
	tpl := template.New("cmd")
	args := &toolCmd{
		MountPoint:    "/mnt/elfs",
		RemoteFs:      conf.Tesla.Elfs.Export,
		ElfsFrontend:  params.System.Frontend,
		Tests:         tests,
		CurrentLoader: string(host),
	}
	if err = ensureCmdArgs(args); err != nil {
		return nil, err
	}
	parse, err := tpl.Parse(
		`mkdir -p {{.MountPoint}} && \
	     mount -o soft {{.ElfsFrontend}}:{{.RemoteFs}} {{.MountPoint}} && \
         {{range .Tests.Flags}}{{/*
         */}}bash -x ./runtests {{.}} -t {{$.MountPoint}}/{{$.CurrentLoader}} && \
         {{end}} true`)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = parse.Execute(&buf, args)
	return append(t.Shell, buf.String()), err
}

func (t *tool) WaitFor() (hosts []types.Host) {
	for _, loader := range t.Params().TargetLoaders {
		hosts = append(hosts, types.Host(loader))
	}
	return hosts
}

func (t *tool) ensureTests() *cthon_config.CthonTests {
	conf := t.Params().Config
	ownConfig := &conf.Cthon
	if ownConfig == nil {
		ownConfig = cthon_config.New()
	}
	tests := &ownConfig.Tests
	if tests == nil {
		tests = &cthon_config.New().Tests
	}
	return tests
}

func ensureCmdArgs(cmd *toolCmd) (err error) {
	if len(cmd.Tests.Flags()) == 0 {
		cmd.Tests.Basic = true
		cmd.Tests.General = true
		cmd.Tests.Lock = true
		cmd.Tests.Special = true
	}
	pass := cmd.MountPoint != "" &&
		cmd.RemoteFs != "" &&
		cmd.ElfsFrontend != "" &&
		cmd.CurrentLoader != ""
	if !pass {
		err = &tool_errors.MissingArgsError{
			Args: cmd,
		}
	}
	return err
}
