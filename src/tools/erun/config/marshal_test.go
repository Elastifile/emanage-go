package erun_config

import (
	"size"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMarshalIOProfile(t *testing.T) {
	conf := Config{
		DirSuffix: "dir_suffix",
		Profile: &ProfileIO{
			Duration:    time.Minute,
			MaxFileSize: size.MiB,
		},
		ConfigFile: "config.file",
	}
	data, err := json.Marshal(conf)
	if err != nil {
		t.Fatalf("Couldn't marshal config: %v", err)
	}
	fmt.Printf("Marshalled: %s\n", data)
	iconf := &Config{}
	if err = json.Unmarshal(data, iconf); err != nil {
		t.Fatalf("Unmarshalling failed: %v", err)
	}
	fmt.Printf("Unmarshalled: %+v, %+v\n", iconf, iconf.Profile)
}

func TestMarshalMetafuzzerProfile(t *testing.T) {
	conf := Config{
		DirSuffix: "dir_suffix",
		Profile: &ProfileMetafuzzer{
			LogFile:        "tesla_metafuzzer.log",
			MaxLogFileSize: 100 * size.MiB,
			Duration:       2 * time.Minute,
			Clients:        4,
			QueueSize:      20,
		},
		ConfigFile: "config.file",
	}
	data, err := json.Marshal(conf)
	if err != nil {
		t.Fatalf("Couldn't marshal config: %v", err)
	}
	fmt.Printf("Marshalled: %s\n", data)
	iconf := &Config{}
	if err = json.Unmarshal(data, iconf); err != nil {
		t.Fatalf("Unmarshalling failed: %v", err)
	}
	fmt.Printf("Unmarshalled: %+v, %+v\n", iconf, iconf.Profile)
}
