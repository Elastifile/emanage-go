package sfs2014

import (
	"bytes"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"textutil"
	sfs2014config "tools/sfs2014/config"
	"types"
)

var sfsresult = `
<?xml version="1.0" encoding="utf-8"?>
<summary id="c0fb2d3e">
  <run fingerprint="86f6" id="d284b811" time="1448530508.25">
    <business_metric>10</business_metric>
    <metric name="op rate" units="ops/s">2000.00</metric>
    <metric name="achieved rate" units="ops/s">1894.72</metric>
    <metric name="average latency" units="milliseconds">9.68</metric>
    <metric name="overall throughput" units="KB/s">27816.84</metric>
    <metric name="read throughput" units="KB/s">10067.23</metric>
    <metric name="write throughput" units="KB/s">17749.62</metric>
    <metric name="run time" units="seconds">300</metric>
    <metric name="clients">3</metric>
    <metric name="processes per client">6</metric>
    <metric name="file size" units="KB">10240</metric>
    <metric name="client data set size" units="MiB">733</metric>
    <metric name="starting data set size" units="MiB">2200</metric>
    <metric name="initial file space" units="MiB">2200</metric>
    <metric name="maximum file space" units="MiB">2400</metric>
    <benchmark name="VDI"/>
    <valid_run/>
  </run>
  <run fingerprint="86f6" id="d284b811" time="1448530508.25">
    <business_metric>11</business_metric>
    <metric name="op rate" units="ops/s">2200.00</metric>
    <metric name="achieved rate" units="ops/s">2004.98</metric>
    <metric name="average latency" units="milliseconds">10.09</metric>
    <metric name="overall throughput" units="KB/s">29384.10</metric>
    <metric name="read throughput" units="KB/s">10590.78</metric>
    <metric name="write throughput" units="KB/s">18793.33</metric>
    <metric name="run time" units="seconds">300</metric>
    <metric name="clients">3</metric>
    <metric name="processes per client">7</metric>
    <metric name="file size" units="KB">10240</metric>
    <metric name="client data set size" units="MiB">806</metric>
    <metric name="starting data set size" units="MiB">2420</metric>
    <metric name="initial file space" units="MiB">2420</metric>
    <metric name="maximum file space" units="MiB">2640</metric>
    <benchmark name="VDI"/>
    <valid_run/>
  </run>
  <run fingerprint="86f6" id="d284b811" time="1448530508.25">
    <business_metric>12</business_metric>
    <metric name="op rate" units="ops/s">2400.00</metric>
    <metric name="achieved rate" units="ops/s">2047.07</metric>
    <metric name="average latency" units="milliseconds">11.45</metric>
    <metric name="overall throughput" units="KB/s">29910.64</metric>
    <metric name="read throughput" units="KB/s">10943.82</metric>
    <metric name="write throughput" units="KB/s">18966.82</metric>
    <metric name="run time" units="seconds">300</metric>
    <metric name="clients">3</metric>
    <metric name="processes per client">8</metric>
    <metric name="file size" units="KB">10240</metric>
    <metric name="client data set size" units="MiB">880</metric>
    <metric name="starting data set size" units="MiB">2640</metric>
    <metric name="initial file space" units="MiB">2640</metric>
    <metric name="maximum file space" units="MiB">2880</metric>
    <benchmark name="VDI"/>
    <valid_run>INVALID_RUN</valid_run>
  </run>
  <run fingerprint="86f6" id="d284b811" time="1448530508.25">
    <business_metric>13</business_metric>
    <metric name="op rate" units="ops/s">2600.00</metric>
    <metric name="achieved rate" units="ops/s">2597.03</metric>
    <metric name="average latency" units="milliseconds">7.50</metric>
    <metric name="overall throughput" units="KB/s">37908.28</metric>
    <metric name="read throughput" units="KB/s">13845.89</metric>
    <metric name="write throughput" units="KB/s">24062.40</metric>
    <metric name="run time" units="seconds">300</metric>
    <metric name="clients">3</metric>
    <metric name="processes per client">8</metric>
    <metric name="file size" units="KB">10240</metric>
    <metric name="client data set size" units="MiB">953</metric>
    <metric name="starting data set size" units="MiB">2860</metric>
    <metric name="initial file space" units="MiB">2860</metric>
    <metric name="maximum file space" units="MiB">3120</metric>
    <benchmark name="VDI"/>
    <valid_run/>
  </run>
</summary>
`

func TestParseFull(t *testing.T) {
	sumFileResults, err := unmarshalXML([]byte(sfsresult))
	if err != nil {
		t.Fatal("Failed parsing...")
	}
	// spew.Dump(sumFileResults)

	var results []types.ToolUnifiedResults
	for _, run := range sumFileResults.Runs {
		var averageLatency float64
		var iops, requestedIops int

		for _, metric := range run.Metrics {
			switch metric.Name {
			case "op rate":
				iops = int(metric.Value)
			case "achieved rate":
				requestedIops = int(metric.Value)
			case "average latency":
				averageLatency = metric.Value
			}
		}
		benchmark := run.Benchmark.Name //create a copy

		toolName := "sfs2014"
		results = append(results, types.ToolUnifiedResults{
			ToolName:       &toolName,
			IOps:           &iops,
			RequestedIOps:  &requestedIops,
			AverageLatency: &averageLatency,
			Workload:       &benchmark,
		})
	}

	var lines [][]string
	for _, res := range results {
		lines = append(lines, textutil.StructToStrings(res))
	}
	spew.Dump(lines)
}

// The following tests ParseTemplate() by processing spec08 template
//
// To print out the processed template just run: 'go test -v [this_file]''
func TestParseSfs2014Template(t *testing.T) {
	conf := types.Config{}

	conf.SetLoaders("loader1", "loader2", "loader3")
	conf.System.Frontend = "192.168.0.1"
	conf.Tesla.Elfs.Export = "my_export/dir"
	conf.Sfs2014 = sfs2014config.Config{}

	opts := &textutil.ParseOpts{
		Data:     conf,
		Template: bytes.NewReader([]byte(template)),
		Title:    "TemplateTest",
	}

	text, err := textutil.ParseTemplate(opts)
	if err != nil {
		t.Errorf("Failed parse, err: %v", err)
	}

	t.Logf("Parsed text: \n%v", text)
}
