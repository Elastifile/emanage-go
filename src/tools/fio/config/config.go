package fio_config

// Config examples:
//
// cmd: fio --directory /mnt jobfile.fio
type Config struct {
	Cmd                     string `yaml:"Cmd" default:"/opt/fio/fio/fio"`      // This will be exported as environment parameter like this: TESLA_FIO_CMD
	ConfigFile              string `yaml:"ConfigFile"`                          // This will be exported as environment parameter like this: TESLA_FIO_CONFIG_FILE
	CheckServersRetries     int    `yaml:"CheckServersRetries" default:"5"`     // Number of retries while checking if fio servers are up and running
	CheckServersTimeToSleep int    `yaml:"CheckServersTimeToSleep" default:"3"` // Time(sec) to sleep between each while checking if fio servers are up and running
	WorkType                string `yaml:"WorkType" default:"normal"`           // 'normal' - normal mode, 'data_integrity' - use data_integrity template and results
	WorkJobs                string `yaml:"WorkJobs"`                            // allow running specific sections in the job file (sep by ','), empty will run all jobs in job file
}
