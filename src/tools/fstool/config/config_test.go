package fstool_config

import (
	"strings"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestConfigToCmd(t *testing.T) {
	conf := NewConfig()
	t.Logf("%+v", conf)
	confStr := strings.Join(conf.ToCmd(), " ")
	t.Log(confStr)
	err := yaml.Unmarshal([]byte(confStr), &conf)
	if err != nil {
		t.Fatal(err)
	}
}
