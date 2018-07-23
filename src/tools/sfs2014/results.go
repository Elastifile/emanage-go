package sfs2014

import (
	xml "encoding/xml"
)

type Summary struct {
	Id   string `xml:"id,attr"`
	Runs []Run  `xml:"run"`
}

type Run struct {
	Fingerprint    string    `xml:"fingerprint,attr"`
	Id             string    `xml:"id,attr"`
	Time           string    `xml:"time,attr"`
	BusinessMetric BMetric   `xml:"business_metric"`
	Metrics        []Metric  `xml:"metric"`
	Benchmark      Benchmark `xml:"benchmark"`
}

type BMetric struct {
	Value float64 `xml:",chardata"`
}

type Benchmark struct {
	Name string `xml:"name,attr"`
}

type Metric struct {
	Name  string  `xml:"name,attr"`
	Units string  `xml:"units,attr"`
	Value float64 `xml:",chardata"`
}

type result struct {
	Latency    float64
	Throughput float64
}

type expectedResults struct {
	result
	numberOfRuns int
}

func unmarshalXML(input []byte) (summary Summary, err error) {
	err = xml.Unmarshal(input, &summary)
	return summary, err
}
