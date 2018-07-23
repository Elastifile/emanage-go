package sfs2008config

type Config struct {
	// TODO: It seems we should use int and not int64 here for all fields
	Latency            float64 `yaml:"Latency" default:"1.5"`  // Not used by SFS directly, but used for validating results
	LoadIo             int     `yaml:"LoadIo" default:"10000"` // Throughput (IOPs)
	NumberOfRuns       int     `yaml:"NumberOfRuns" default:"1"`
	ProcessesPerClient int     `yaml:"ProcessesPerClient" default:"10"` // type must be int for strings.join() (used by template SpecSfs2008ConfigTemplate)
	IncrLoad           int     `yaml:"IncrLoad" default:"0"`
	Runtime            int     `yaml:"Runtime" default:"300"`
	WarmupTime         int     `yaml:"WarmupTime" default:"300"`
	MixFile            string  `yaml:"MixFile" tesla:"filename"` // The files themselves are sent to filestore file server and later pulled by name.
	ConfigFile         string  `yaml:"ConfigFile" tesla:"filename"`
}
