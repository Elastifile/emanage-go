package sfs2008

import (
	"regexp"

	"github.com/go-errors/errors"

	"tools/parser"
)

// Example: `SFS NFS THROUGHPUT:  167352 Ops/Sec   AVG. RESPONSE TIME:     1.8 Msec/Op`
var reResFile = regexp.MustCompile(`(?m:^` + // multi-line flag
	`SFS NFS THROUGHPUT:\s+` +
	`(?P<IOps>[0-9.]+)` +
	`\s+Ops/Sec` +
	`\s+` +
	`AVG. RESPONSE TIME:\s+` +
	`(?P<Latency>[0-9.]+)` +
	`\sMsec/Op` +
	`)`,
)

// Examples:
//            2560    2571     5.8   768874  299 CIFS T 4  314628128   4 10  2  2 2008
// INVALID   10000   11523     0.9   345703   30 NFS3 T 4   18882360   2  5  2  2 2008
//                   thru      lat
var reSumFile = regexp.MustCompile(`(?m:^` + // multi-line flag
	`(?P<Invalid>INVALID)?\s+` +
	`(?P<RequestedIOPS>[0-9.]+)\s+` +
	`(?P<IOPS>[0-9.]+)\s+` +
	`(?P<Latency>[0-9.]+)` +
	`)`,
)

type result struct {
	RequestedIOPS int
	IOPS          int
	Latency       float64
}

type expectedResults struct {
	result
	numberOfRuns int
}

func parseAllResults(re *regexp.Regexp, input string) (results []result, err error) {
	var r result
	p := parser.New(re, input)

	for p.NextResult(&r) {
		if p.Err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}

func verify(input string, reference expectedResults) error {
	results, err := parseAllResults(reSumFile, input)

	if err != nil {
		return errors.Errorf("Failed to parse results. err: %v", err)
	}

	if len(results) != int(reference.numberOfRuns) {
		return errors.Errorf("Found results of %v runs, while expected %v runs.",
			len(results), reference.numberOfRuns)
	}

	avg := average(results)

	if avg.Latency > reference.Latency {
		return errors.Errorf("Latency is %v ms/op, which is worse than the maximum of %v ms/op",
			avg.Latency, reference.Latency)
	}

	// TODO: consider losing the expectedResults and comparing vs result.RequestedIOps
	relativeThroughput := float64(avg.IOPS) / float64(reference.IOPS)

	// TODO: Expose the following constants as parameters in the config
	if relativeThroughput < 0.95 {
		return errors.Errorf("Throughput is %v ops/s, which is only %.2f%% of the expected %v ops/s",
			avg.IOPS, 100.0*relativeThroughput, avg.RequestedIOPS)
	}

	return nil
}

// TODO: Make this work automatically on any data structure via reflection?
// It can then be reused for other tools.
func average(results []result) result {
	avg := result{}
	numResults := len(results)

	for _, result := range results {
		avg.Latency += result.Latency
		avg.IOPS += result.IOPS
	}

	avg.Latency = avg.Latency / float64(numResults)
	avg.IOPS = avg.IOPS / numResults

	return avg
}
