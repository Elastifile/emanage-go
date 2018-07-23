package fio

import (
	"encoding/json"
	"strings"

	"github.com/go-errors/errors"
)

// Examples:

/*
{
  "fio version" : "fio-2.6",
  "global options" : {
    "unified_rw_reporting" : "1"
  },
  "client_stats" : [
    {
      "jobname" : "All clients",
      "groupid" : 0,
      "error" : 0,
      "mixed" : {
        "io_bytes" : 2021176,
        "bw" : 16653,
        "iops" : 2252.49,
        "runtime" : 121363,
        "total_ios" : 273369,
        "short_ios" : 0,
        "drop_ios" : 0,
        "slat" : {
          "min" : 0,
          "max" : 0,
          "mean" : 0.00,
          "stddev" : 0.00
        },
        "lat" : {
          "min" : 431,
          "max" : 3781314,
          "mean" : 1782.92,
          "stddev" : 19405.44
        },
        "bw_min" : 2,
        "bw_max" : 5160,
        "bw_agg" : 6.98,
        "bw_mean" : 2335.74,
        "bw_dev" : 835.52
      },
      "hostname" : "func12-loader2",
      "port" : 8765
    }
  ],
  "disk_util" : [
  ]
}
*/

type fioJSONFileResults struct {
	FioVersion    string `json:"fio version"`
	Header        string `json:"header"`
	HaveErrors    bool   `json:"have errors"`
	GlobalOptions struct {
		UnifiedRwReporting string `json:"unified_rw_reporting"`
	} `json:"global options"`
	ClientStats []struct {
		JobName string `json:"jobname"`
		Error   int    `json:"error"`
		Mixed   struct {
			Bandwidth float64 `json:"bw"`
			IOPS      float64 `json:"iops"`
			Latency   struct {
				Mean float64 `json:"mean"`
			} `json:"lat"`
		} `json:"mixed"`
		Read struct {
			Bandwidth float64 `json:"bw"`
			IOPS      float64 `json:"iops"`
			Latency   struct {
				Mean float64 `json:"mean"`
			} `json:"lat"`
		} `json:"read"`
		Write struct {
			Bandwidth float64 `json:"bw"`
			IOPS      float64 `json:"iops"`
			Latency   struct {
				Mean float64 `json:"mean"`
			} `json:"lat"`
		} `json:"write"`
	} `json:"client_stats"`
}

type result struct {
	IOPS      int
	Latency   float64
	Bandwidth float64
}

type resultList struct {
	IOPSList      []int
	LatencyList   []float64
	BandwidthList []float64
}

type jobResults struct {
	result
	resultList
	jobname string
}

func (t *tool) parseAllResults(input runeReader) (results []jobResults, err error) {
	resultsJSON, err := decodeJSON(input)
	if err != nil {
		return nil, errors.Errorf("parseAllResults: Failed to decode JSON, error=%s", err)
	}

	results, err = t.exportResults(resultsJSON)
	if err != nil {
		return nil, errors.Errorf("parseAllResults: Failed to export results from JSON, error=%s", err)
	}

	return results, nil
}

func decodeJSON(reader runeReader) (result *fioJSONFileResults, err error) {
	var (
		ch     rune
		nl     = true
		header []rune
	)
	for err == nil {
		ch, _, err = reader.ReadRune()
		if nl && ch == '{' {
			err = reader.UnreadRune()
			break
		} else if ch == '\n' {
			nl = true
		} else {
			nl = false
			// REVIEW(oleg): TODO: The original reader is bytes.Reader, this
			// means it has Seek() method and Read([]byte), you could
			// do all the reading in a single go.  When you append an
			// element to a slice, you might need to re-allocate the
			// array pointed to by this slice.  In other words, the
			// amount of work done by this function would be:
			// T(n) = T(n - 1) + n
			//      = T(1) + 2T(1) + 3T(1) + ... + nT(1)
			//      = n * (n + 2) / 2 = O(n^2)
			//      for T(1) = 1
			// in the worst case.  However, if you did something like:
			// for err == nil {
			//     i++
			//     if error {
			//         reader.Seek(0)
			//         header = make([]byte, i)
			//         reader.Read(header)
			//         break
			//     }
			// }
			// you would have done it in O(n) time.
			header = append(header, ch)
		}
	}
	if err != nil {
		return nil, errors.Errorf("decodeJSON: Not valid JSON input, error=%s", err)
	}
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, errors.Errorf("decodeJSON: Failed to decode JSON input, error=%s", err)
	}
	result.Header = string(header)

	return result, nil
}

func (t *tool) exportResults(fioResults *fioJSONFileResults) (jResults []jobResults, err error) {
	if fioResults == nil {
		return jResults, errors.New("exportResults: input fioResults is nil")
	}

	// Check if we need to expect a Mixed or Read & Write results
	unifiedRwReport := fioResults.GlobalOptions.UnifiedRwReporting == "1"

	jobs := make(map[string]jobResults)
	jobAllClientsStr := "All clients"

	for _, clientStats := range fioResults.ClientStats {
		// Skip summary job named "All clients"
		if clientStats.JobName != jobAllClientsStr {
			var job, ok = jobs[clientStats.JobName]
			if !ok {
				job.jobname = clientStats.JobName
			}
			if unifiedRwReport {
				job.BandwidthList = append(job.BandwidthList, clientStats.Mixed.Bandwidth)
				job.IOPSList = append(job.IOPSList, int(clientStats.Mixed.IOPS))
				job.LatencyList = append(job.LatencyList, clientStats.Mixed.Latency.Mean)
			} else {
				job.BandwidthList = append(job.BandwidthList, clientStats.Read.Bandwidth+clientStats.Write.Bandwidth)
				job.IOPSList = append(job.IOPSList, int(clientStats.Read.IOPS+clientStats.Write.IOPS))
				job.LatencyList = append(job.LatencyList, (clientStats.Read.Latency.Mean+clientStats.Write.Latency.Mean)/2)
			}
			jobs[clientStats.JobName] = job
		}
	}

	for _, job := range jobs {
		err = calcJobResults(&job)
		if err != nil {
			// log error
			t.Logger().Error("Failure calcluating job results", "job results", job.result, "error", err)
		} else {
			jResults = append(jResults, job)
		}
	}

	if averageAllResults {
		t.Logger().Info("Averaging the following jobs results", "jobs results", jResults)
		jResults, err = averageResults(jResults)
		t.Logger().Info("Averaged results", "average results", jResults)
	}

	return jResults, nil
}

func calcJobResults(job *jobResults) (err error) {
	// calc IOPSList to IOPS
	job.IOPS = 0
	for _, iops := range job.IOPSList {
		job.IOPS += iops
	}

	// calc BandwidthList to Bandwidth
	job.Bandwidth = 0
	for _, bw := range job.BandwidthList {
		job.Bandwidth += bw
	}
	// fio bandwidth units - KB/s, tesla bandwidth units - MB/s, div by 1000
	job.Bandwidth = job.Bandwidth / 1000

	// calc LatencyList to Latency
	// fio latency units - usec, tesla latancy units - msec, div by 1000
	job.Latency = float64(averageFloat(job.LatencyList) / 1000)

	return nil
}

func averageResults(jobsResults []jobResults) ([]jobResults, error) {
	avgResults := make([]jobResults, 1)
	var jobNames []string
	for _, jobResults := range jobsResults {
		jobNames = append(jobNames, jobResults.jobname)
		avgResults[0].IOPSList = append(avgResults[0].IOPSList, jobResults.IOPS)
		avgResults[0].LatencyList = append(avgResults[0].LatencyList, jobResults.Latency)
		avgResults[0].BandwidthList = append(avgResults[0].BandwidthList, jobResults.Bandwidth)
	}
	avgResults[0].jobname = strings.Join(jobNames, ";")
	avgResults[0].IOPS = int(averageInt(avgResults[0].IOPSList))
	avgResults[0].Latency = averageFloat(avgResults[0].LatencyList)
	avgResults[0].Bandwidth = averageFloat(avgResults[0].BandwidthList)

	return avgResults, nil
}

// REVIEW(oleg): TODO: Averaging floats is more complicated than this.
// Unless you are certain that numerical underflow will not occur, it
// is best to use: https://en.wikipedia.org/wiki/Kahan_summation_algorithm

// TODO: move this to general math package
func averageFloat(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// TODO: move this to general math package
func averageInt(xs []int) float64 {
	total := float64(0)
	for _, x := range xs {
		total += float64(x)
	}
	return total / float64(len(xs))
}
