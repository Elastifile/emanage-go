package tools_api

import (
	"time"

	"messaging"
	"sysapi"
	"github.com/elastifile/emanage-go/src/helputils"
	"github.com/elastifile/emanage-go/src/logging"
	"github.com/elastifile/emanage-go/src/types"

	cthon_config "github.com/elastifile/emanage-go/src/tools/cthon/config"
	erun_config "github.com/elastifile/emanage-go/src/tools/erun/config"
	fio_config "github.com/elastifile/emanage-go/src/tools/fio/config"
	fstool_config "github.com/elastifile/emanage-go/src/tools/fstool/config"
	migration_config "github.com/elastifile/emanage-go/src/tools/migration/config"
	sfs2008_config "github.com/elastifile/emanage-go/src/tools/sfs2008/config"
	sfs2014_config "github.com/elastifile/emanage-go/src/tools/sfs2014/config"
	vdbench_config "github.com/elastifile/emanage-go/src/tools/vdbench/config"

	tooly_api "tooly/api"
)

var logger = logging.NewLogger("tools_api")

const (
	Timeout     = 5 * time.Minute
	TimeoutWait = 30 * 24 * time.Hour
)

////////////////////////////////////////////////////////////////////////////////

type Tools struct {
	conf   types.Config
	tooly  *tooly_api.Client
	system *sysapi.System
}

type resultWaiter interface {
	GetResults() []types.ToolUnifiedResults
}

type waiter interface {
	Wait() (resultWaiter, error)
	GetResults() []types.ToolUnifiedResults
}

type StopWaiter interface {
	Wait() (resultWaiter, error)
	Stop() (waiter, error)
	Status() (*types.ToolStatus, error)
}

type ToolPropsCollector interface {
	Start(...string) (StopWaiter, error)
	SetLoaders(...int) ToolPropsCollector
	SetExports(...*sysapi.Export) ToolPropsCollector
	SetSystem(*sysapi.System) ToolPropsCollector
}

type toolClient interface {
	ToolPropsCollector
	StopWaiter
	resultWaiter
}

func New(msg messaging.Client, conf types.Config, system *sysapi.System) *Tools {
	tooly := tooly_api.NewClient(msg, conf)
	return &Tools{
		conf:   conf,
		tooly:  tooly,
		system: system,
	}
}

func (t *Tools) NewErun(config ...*erun_config.Config) ToolPropsCollector {
	var conf *erun_config.Config
	if len(config) > 0 {
		conf = config[0]
	} else {
		conf = t.conf.Erun
	}
	logger.Info("New Erun", helputils.MustStructToKeyValueInterfaces(conf)...)

	result := &erun{
		toolClientBase: &toolClientBase{
			Tools: t,
			name:  types.Erun,
		},
		toolConfig: conf,
	}
	result.conf.Erun = result.toolConfig
	result.impl = result
	logger.Info("New erun", "cmd", conf.Profile.ToCmd(), "conf", result.toolConfig)
	return result
}

func (t *Tools) NewSfs2008(config ...*sfs2008_config.Config) ToolPropsCollector {
	var conf *sfs2008_config.Config
	if len(config) > 0 {
		conf = config[0]
	} else {
		conf = &t.conf.Sfs2008
	}
	logger.Info("New sfs2008", helputils.MustStructToKeyValueInterfaces(conf)...)

	result := &sfs2008{
		toolClientBase: &toolClientBase{
			Tools: t,
			name:  types.Sfs2008,
		},
		toolConfig: conf,
	}
	result.conf.Sfs2008 = *result.toolConfig
	result.impl = result
	return result
}

func (t *Tools) NewSfs2014(config ...*sfs2014_config.Config) ToolPropsCollector {
	var conf *sfs2014_config.Config
	if len(config) > 0 {
		conf = config[0]
	} else {
		conf = &t.conf.Sfs2014
	}
	logger.Info("New sfs2014", helputils.MustStructToKeyValueInterfaces(conf)...)

	result := &sfs2014{
		toolClientBase: &toolClientBase{
			Tools: t,
			name:  types.Sfs2014,
		},
		toolConfig: conf,
	}
	result.conf.Sfs2014 = *result.toolConfig
	result.impl = result
	return result
}

func (t *Tools) NewFsTool(config ...fstool_config.Config) ToolPropsCollector {
	conf := fstool_config.NewConfig()
	if len(config) > 1 {
		panic("cannot receive more than one config")
	} else if len(config) == 1 {
		conf = config[0]
	}

	switch conf.Method {
	case "TreeCreate":
		logger.Info("New fstool."+conf.Method,
			append(helputils.MustStructToKeyValueInterfaces(conf.TreeCreate), helputils.MustStructToKeyValueInterfaces(conf.NfsClient)...)...,
		)
	case "TreeUpdate":
		logger.Info("New fstool."+conf.Method,
			append(helputils.MustStructToKeyValueInterfaces(conf.TreeUpdate), helputils.MustStructToKeyValueInterfaces(conf.NfsClient)...)...,
		)
	case "TreeDelete":
		logger.Info("New fstool."+conf.Method,
			append(helputils.MustStructToKeyValueInterfaces(conf.TreeDelete), helputils.MustStructToKeyValueInterfaces(conf.NfsClient)...)...,
		)
	case "TreeDiff":
		logger.Info("New fstool."+conf.Method,
			append(helputils.MustStructToKeyValueInterfaces(conf.TreeDiff), helputils.MustStructToKeyValueInterfaces(conf.NfsClient)...)...,
		)
	default:
		logger.Info("New fstool", helputils.MustStructToKeyValueInterfaces(conf)...)
	}

	result := &fsTool{
		toolClientBase: &toolClientBase{
			Tools: t,
			name:  types.FsTool,
		},
		toolConfig: &conf,
	}
	result.conf.FsTool = *result.toolConfig
	result.impl = result
	logger.Debug("Created fstool", "conf", result.conf)
	return result
}

func (t *Tools) NewMigration(config ...*migration_config.Config) ToolPropsCollector {
	var conf *migration_config.Config
	if len(config) > 0 {
		conf = config[0]
	} else {
		conf = &t.conf.Migration
	}
	result := &mig{
		toolClientBase: &toolClientBase{
			Tools: t,
			name:  types.Migration,
		},
		toolConfig: conf,
	}
	result.conf.Migration = *result.toolConfig
	result.impl = result
	return result
}

func (t *Tools) NewCthon(config ...*cthon_config.Config) ToolPropsCollector {
	var conf *cthon_config.Config
	if len(config) > 0 {
		conf = config[0]
	} else {
		conf = &t.conf.Cthon
	}
	logger.Info("New Cthon", helputils.MustStructToKeyValueInterfaces(conf)...)

	result := &cthon{
		toolClientBase: &toolClientBase{
			Tools: t,
			name:  types.Cthon,
		},
		toolConfig: conf,
	}
	result.conf.Cthon = *result.toolConfig
	result.impl = result
	return result
}

func (t *Tools) NewVdbench(config ...*vdbench_config.Config) ToolPropsCollector {
	var conf *vdbench_config.Config
	if len(config) > 0 {
		conf = config[0]
	} else {
		conf = &t.conf.Vdbench
	}
	logger.Info("New vdbench", helputils.MustStructToKeyValueInterfaces(conf)...)

	result := &vdbench{
		toolClientBase: &toolClientBase{
			Tools: t,
			name:  types.Vdbench,
		},
		toolConfig: conf,
	}
	result.conf.Vdbench = *result.toolConfig
	result.impl = result
	return result
}

func (t *Tools) NewFio(config ...*fio_config.Config) ToolPropsCollector {
	var conf *fio_config.Config
	if len(config) > 0 {
		conf = config[0]
	} else {
		conf = &t.conf.Fio
	}
	logger.Info("New fio", helputils.MustStructToKeyValueInterfaces(conf)...)

	result := &fio{
		toolClientBase: &toolClientBase{
			Tools: t,
			name:  types.Fio,
		},
		toolConfig: conf,
	}
	result.conf.Fio = *result.toolConfig
	result.impl = result
	return result
}
