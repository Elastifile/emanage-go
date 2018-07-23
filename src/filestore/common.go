package filestore

import (
	"encoding/json"
	"fmt"
	"time"

	log "gopkg.in/inconshreveable/log15.v2"
	yaml "gopkg.in/yaml.v2"

	"remote"
	"types"
)

const (
	insecure = true
	port     = 9000

	AccessKey = "12345678901234567890"
	SecretKey = "1234567890123456789012345678901234567890"

	Image = "tesla/tools/minio"

	ToolsConfigBucket = "tools.config"

	downloadTimeout = 5 * time.Minute
)

func LoadFullConfig(env types.StartupEnvVars) (types.Config, error) {
	conf := types.NewConfig()

	data, err := SafeReadFiles(types.Host(env.FilestoreHost), types.FullConfig, nil, downloadTimeout,
		func(report *SafeTransferReport) {
			conf.Logging.PrintDebug(report.Text,
				"err", report.Err,
				"file", types.FullConfig,
				"grace", report.Grace,
			)
		},
	)

	if err != nil {
		conf.Logging.PrintError(err.Error())
		return *conf, err
	}

	if len(data) > 1 {
		err = fmt.Errorf("Expected a single config file in filestore, Host: %v, Bucket: %v, Count %d", env.FilestoreHost, types.FullConfig, len(data))
		conf.Logging.PrintError(err.Error())
		return *conf, err
	}

	if err := json.Unmarshal(data[0], conf); err != nil {
		conf.Logging.PrintError(err.Error())
		return *conf, err
	}
	conf.Logging.PrintDebug("Unmarshalled tesla full config",
		"data", conf,
	)

	if conf.Tesla.ShellUser != "" {
		remote.DefaultUser = conf.Tesla.ShellUser
	}

	return *conf, nil
}

func LoadConfElabSystems(conf *types.Config) error {
	// TODO: here FilestoreHost is already initialized (with external IP), but conf is empty - so we can't get the internal addr
	// It should be safe enough to use the existing value, since we should be passed the correct one via an env variable
	bodies, err := SafeReadFiles(conf.FilestoreHostInternal(), types.ConfElabSystemBucket, nil, downloadTimeout,
		func(report *SafeTransferReport) {
			conf.Logging.PrintDebug(report.Text,
				"err", report.Err,
				"file", types.ConfElabSystemBucket,
				"grace", report.Grace,
			)
		},
	)
	if err != nil {
		conf.Logging.PrintError(err.Error())
		return err
	}

	for _, body := range bodies {
		conf.System.Elab = types.ElabSystem{}
		if err := json.Unmarshal(body, &conf.System.Elab); err != nil {
			conf.Logging.PrintError(err.Error())
			return err
		}
		conf.Logging.PrintDebug("Unmarshalled system elab",
			"data", conf.System.Elab.Data,
		)

		conf.AddSystem()
	}
	conf.MakeFirstSystemReference()

	conf.Logging.PrintInfo("Primary system",
		"Name", conf.System.Elab.Data.Name,
		"MasterLoader", conf.MasterLoaderInternal(),
		"Loaders", conf.LoadersInternal(),
		"Vheads", conf.VHeads(),
		"Emanage", conf.System.Elab.Data.Emanage,
		"FilestoreHost", conf.FilestoreHostInternal(),
	)

	return nil
}

func LoadTestConfigs(conf *types.Config, test *types.TestConfigs, logger log.Logger) error {
	if conf.Tests.TestConfigsPath() == "" {
		logger.Info("Test configs file not found")
		test = nil
		return nil
	}

	logger.Info("Parsing test suite params")
	yamlBodies, err := SafeReadFiles(conf.FilestoreHostInternal(), types.ConfTestConfigs, []string{conf.Tests.TestConfigsPath()}, downloadTimeout,
		func(report *SafeTransferReport) {
			logger.Debug(report.Text,
				"err", report.Err,
				"file", types.ConfTestConfigs,
				"grace", report.Grace,
			)
		})
	if err != nil {
		logger.Error(err.Error())
		return err
	} else if err := yaml.Unmarshal(yamlBodies[0], &test); err != nil {
		logger.Error("Failed unmarshal test configs file")
		logger.Error(err.Error())
		return err
	}
	logger.Info("Test configs loaded")

	return nil
}

func LoadProductConfig(conf *types.Config, product *types.ProductParameters, logger log.Logger) error {
	if conf.Tests.ProductConfigFile == "" {
		logger.Info("Product config file not supplied")
		*product = types.NewProductDefaults()
		return nil
	}

	logger.Info("Parsing product params")
	yamlBodies, err := SafeReadFiles(conf.FilestoreHostInternal(), types.ConfProductConfig, []string{conf.Tests.ProductConfigFile}, downloadTimeout,
		func(report *SafeTransferReport) {
			logger.Warn(report.Text,
				"err", report.Err,
				"file", types.ConfProductConfig,
				"grace", report.Grace,
			)
		})
	if err != nil {
		logger.Error(err.Error())
		return err
	} else if err := yaml.Unmarshal(yamlBodies[0], &product); err != nil {
		logger.Error("Failed unmarshal product config file")
		logger.Error(err.Error())
		return err
	}
	logger.Info("Product configs loaded")

	return nil
}
