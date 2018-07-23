package teslatest_config

type Config struct {
	ConfigFile string `default:"" yaml:"ConfigFile"` // This will be exported as environment parameter like this: TESLA_TESLATEST_CONFIG_FILE
}
