package migration_config

import "time"

type Config struct {
	Interval time.Duration `yaml:"Interval"`
	Duration time.Duration `yaml:"Duration"`
}
