package cthon_config

type CthonTests struct {
	Basic       bool `yaml:"Basic"`
	General     bool `yaml:"General"`
	Special     bool `yaml:"Special"`
	Lock        bool `yaml:"Lock"`
	Functional  bool `yaml:"Functional"`
	SuppressDir bool `yaml:"SuppressDir"`
}

func (tests *CthonTests) Flags() []string {
	chunks := []string{}
	if tests.Basic {
		chunks = append(chunks, "-b")
	}
	if tests.General {
		chunks = append(chunks, "-g")
	}
	if tests.Special {
		chunks = append(chunks, "-s")
	}
	if tests.Lock {
		chunks = append(chunks, "-l")
	}
	if tests.Functional {
		chunks = append(chunks, "-f")
	}
	if tests.SuppressDir {
		chunks = append(chunks, "-n")
	}
	if len(chunks) == 4 {
		chunks = []string{"-a"}
	}
	return chunks
}

type Config struct {
	Tests CthonTests `yaml:"Tests"`
}

func New() *Config {
	return &Config{
		Tests: CthonTests{

			Basic:       true,
			General:     true,
			Special:     true,
			Lock:        true,
			Functional:  false,
			SuppressDir: false,
		},
	}
}
