package erun

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/elastifile/emanage-go/src/tools/common"
	erun_config "github.com/elastifile/emanage-go/src/tools/erun/config"
	"github.com/elastifile/emanage-go/src/types"
)

var jsonResult = `{"READS_min_latency_nanos" : 701625, "READS_max_latency_nanos" : 41404780, "WRITES_min_latency_nanos" : 1078933, "WRITES_max_latency_nanos" : 44840375, "COMBINED_min_latency_nanos" : 701625, "COMBINED_max_latency_nanos" : 44840375, "READS_total_ios" : 134188, "READS_total_xfer_bytes" : 549408768, "READS_total_accum_latency" : 910272414631, "WRITES_total_ios" : 467118, "WRITES_total_xfer_bytes" : 1913151488, "WRITES_total_accum_latency" : 4527826723865, "long_pending_ios" : 0, "time_running_secs" : 59}`

func TestParseResults(t *testing.T) {
	jsonBytes := []byte(jsonResult)
	params := &types.ToolParams{
		ToolName:      "erun",
		TargetLoaders: []string{"localhost"},
		Config: types.Config{
			Erun: &erun_config.Config{
				Profile: &erun_config.ProfileIO{},
			},
		},
	}
	props, err := common.NewExternalProps(params)
	if err != nil {
		t.Fatal(err)
	}

	erun := &tool{ExternalToolProperties: props, useHostNetworking: true}
	erun.DefaultLoggingTool = common.NewLoggingTool(erun)

	p, err := parseResults(erun, [][]byte{jsonBytes})
	if err != nil {
		t.Fatalf("Failed parsing: %s\n", err)
	}
	spew.Dump(p)
}
