package sfs2014config

type Config struct {
	// TODO: It seems we should use int and not int64 here for all fields
	Latency            float64 `yaml:"Latency" default:"1.5"` // Not used by SFS directly, but used for validating results
	Benchmark          string  `yaml:"Benchmark" default:"VDI"`
	Load               int     `yaml:"Load" default:"10"` // Throughput (IOPs)
	IncrLoad           int     `yaml:"IncrLoad" default:"1"`
	NumberOfRuns       int     `yaml:"NumberOfRuns" default:"1"`
	ProcessesPerClient int     `yaml:"ProcessesPerClient" default:"10"` // type must be int for strings.join() (used by template SpecSfs2008ConfigTemplate)
	Runtime            int     `yaml:"Runtime" default:"300"`
	WarmupTime         int     `yaml:"WarmupTime" default:"300"`
	ConfigFile         string  `yaml:"ConfigFile" tesla:"filename"`
}
