package erun

import (
	"encoding/json"
	"fmt"
	"time"

	"types"
)

type ErunResults struct {
	CombinedMaxLatencyNanos int `json:"COMBINED_max_latency_nanos"`
	CombinedMinLatencyNanos int `json:"COMBINED_min_latency_nanos"`
	ReadsMinLatencyNanos    int `json:"READS_min_latency_nanos"`
	ReadsTotalAccumLatency  int `json:"READS_total_accum_latency"`
	ReadsTotalIos           int `json:"READS_total_ios"`
	ReadsTotalXferBytes     int `json:"READS_total_xfer_bytes"`
	WritesMaxLatencyNanos   int `json:"WRITES_max_latency_nanos"`
	WritesMinLatencyNanos   int `json:"WRITES_min_latency_nanos"`
	WritesTotalAccumLatency int `json:"WRITES_total_accum_latency"`
	WritesTotalIos          int `json:"WRITES_total_ios"`
	WritesTotalXferBytes    int `json:"WRITES_total_xfer_bytes"`
	LongPendingIos          int `json:"long_pending_ios"`
	TimeRunningSecs         int `json:"time_running_secs"`
}

func nanoToMilli(nano int) int {
	return nano / int(time.Millisecond)
}

// REVIEW(gavriep): This function (and its siblings in other tools) should io.Reader values instead of []byte slices.
func parseResults(t *tool, data [][]byte) (r types.ToolUnifiedResults, err error) {
	var totalIOs, duration, totalIOps, totalAccumLatency int
	var avgLatencyMilli float64

	t.Logger().Info("Parsing results...")
	for _, d := range data {
		t.Logger().Debug("parsing data", "data", string(d))
		var res ErunResults
		err = json.Unmarshal(d, &res)
		if err != nil {
			return r, err
		}

		// sum hosts results
		totalIOs += (res.WritesTotalIos + res.ReadsTotalIos)
		totalAccumLatency += res.ReadsTotalAccumLatency + res.WritesTotalAccumLatency
		duration = res.TimeRunningSecs
	}

	// for latency calculations see EL-1675
	t.Logger().Debug("parseResults", "totalIOs", totalIOs, "totalTimeRunningSecs", duration)

	var latency int = nanoToMilli(totalAccumLatency)

	// protect dev zero
	if totalIOs > 0 {
		avgLatencyMilli = float64(latency) / float64(totalIOs)
	}
	if duration > 0 {
		totalIOps = totalIOs / duration
	}

	profileName := fmt.Sprintf("%v", t.Params().Config.Erun.Profile.Name())

	identifier := t.Params().Identifier
	name := string(t.Name())
	jobID := string(t.Params().JobID)
	return types.ToolUnifiedResults{
		Identifier:     &identifier,
		ToolName:       &name,
		JobID:          &jobID,
		IOps:           &totalIOps,
		AverageLatency: &avgLatencyMilli,
		Workload:       &profileName,
		// The rest of the fields aren't supported by erun and will automatically
		// convert to "N/A" by textutil.StructToStrings in results csv file.
	}, err
}
