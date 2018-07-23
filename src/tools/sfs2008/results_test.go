package sfs2008

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"types"
)

var sfsresult = `
Tue Oct 13 03:55:38 IDT 2015 ************************************************************************
Started on client (loader-esx9a): ../bin/sfs_prime -l 41750 -C ../result/sfssum.jenkins -t 300 -m ../result/mix_file -a 30 -A 70 -B 32 -R 2 -W 2 -Q -p 80 -w 300 -X 8000 loader-esx9a loader-esx9b loader-esx9c loader-esx9d

Aggregate Test Parameters:
    Number of processes = 320
    Requested Load (NFS operations/second) = 167000
    Maximum number of outstanding biod writes = 2
    Maximum number of outstanding biod reads = 2
    Warm-up time (seconds) = 300
    Run time (seconds) = 300
    NFS Mixfile = ../result/mix_file
    File Set = 640000 Files created for I/O operations
               192000 Files accessed for I/O operations
                12800 Files for non-I/O operations
                  320 Symlinks
                21120 Directories
                      Additional non-I/O files created as necessary

SFS Aggregate Results for 4 Client(s), Tue Oct 13 04:13:30 2015
NFS Protocol Version 3

-------------------------------------------------------------------------------------
NFS         Target Actual  NFS Op  NFS Op    NFS    Mean     Std Dev  Std Error Pcnt
Op          NFS    NFS     Logical Physical  Op     Response Response of Mean,  of
Type        Mix    Mix     Success Success   Error  Time     Time     95% Conf  Total
            Pcnt   Pcnt    Count   Count     Count  Msec/Op  Msec/Op  +-Msec/Op Time
-------------------------------------------------------------------------------------
getattr      26.0%  26.1% 13113761 13113761      0     0.66     1.00      0.00   9.3%
setattr       4.0%   4.0%  2018508  2018508      0     1.54     2.10      0.00   3.3%
lookup       24.0%  24.1% 12115120 12115120      0     1.03     1.40      0.00  13.5%
readlink      1.0%   1.0%   505442   505442      0     2.26     3.08      0.00   1.2%
read         18.0%  18.1%  9077788 11162846      0     3.35     3.85      0.00  33.0%
write        10.0%  10.0%  5045795  6106171      0     4.17     4.44      0.00  22.8%
create        1.0%   1.0%   503632   503632      0     3.70     4.30      0.01   2.0%
remove        1.0%   0.6%   296389   296389      0     1.35     2.53      0.01   0.4%
readdir       1.0%   1.0%   503161   503161      0     1.26     1.40      0.00   0.7%
fsstat        1.0%   1.0%   504881   504881      0    11.43    30.27      0.02   6.3%
access       11.0%  11.1%  5553425  5553425      0     0.64     1.04      0.00   3.8%
readdirplus   2.0%   2.0%  1010018  1010018      0     3.27     4.05      0.00   3.5%
-------------------------------------------------------------------------------------
INVALID RUN reported for Client 1 (loader-esx9a).
INVALID RUN, ILLEGAL PARAMETER: Non-standard Mix file
INVALID RUN reported for Client 2 (loader-esx9b).
INVALID RUN, ILLEGAL PARAMETER: Non-standard Mix file
INVALID RUN reported for Client 3 (loader-esx9c).
INVALID RUN, ILLEGAL PARAMETER: Non-standard Mix file
INVALID RUN reported for Client 4 (loader-esx9d).
INVALID RUN, ILLEGAL PARAMETER: Non-standard Mix file

        ---------------------------------------------
        |  SPEC SFS 2008 AGGREGATE RESULTS SUMMARY  |
        ---------------------------------------------
SFS NFS THROUGHPUT:  167352 Ops/Sec   AVG. RESPONSE TIME:     1.8 Msec/Op
TCP PROTOCOL (IPv4)
NFS MIXFILE:../result/mix_file
AGGREGATE REQUESTED LOAD: 167000 Ops/Sec
TOTAL LOGICAL NFS OPERATIONS:  50247920       TEST TIME: 300 Sec
TOTAL PHYSICAL NFS OPERATIONS: 53393354
PHYSICAL NFS IO THROUGHPUT: 177977 Ops/sec
NUMBER OF SFS CLIENTS: 4
TOTAL FILE SET SIZE CREATED: 331875.0 MB
TOTAL FILE SET SIZE ACCESSED: 99728.4 - 110533.3 MB  (100.00% to 110.83% of Base)

------------------------------------------------------------------------
`

// func TestParse(t *testing.T) {
// 	input := "SFS NFS THROUGHPUT:  167352 Ops/Sec   AVG. RESPONSE TIME:     1.8 Msec/Op\n"
// 	results, err := parseAllResults(reResFile, input)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if len(results) != 1 {
// 		t.Fatal("Expected 1 result instead got", len(results))
// 	}

// 	record := results[0]

// 	if record.Latency != 1.8 {
// 		t.Fatal("Wrong latency")
// 	}
// 	if record.IOPS != 167352 {
// 		t.Fatal("Wrong throughput", record.IOPS)
// 	}
// }

func TestParseSum(t *testing.T) {
	sumFile := `
INVALID   10000   11523     0.9   345703   30 NFS3 T 4   18882360   2  5  2  2 2008
           2560    2571     5.8   768874  299 CIFS T 4  314628128   4 10  2  2 2008
INVALID   20000   20000     1.5   111111   30 NFS3 T 4   18882360   2  5  2  2 2008
          30000   20000     1.5   111111   30 NFS3 T 4   18882360   2  5  2  2 2008
`
	expectedResults := []result{
		{Latency: 0.9, IOPS: 11523, RequestedIOPS: 10000},
		{Latency: 5.8, IOPS: 2571, RequestedIOPS: 2560},
		{Latency: 1.5, IOPS: 20000, RequestedIOPS: 20000},
		{Latency: 1.5, IOPS: 20000, RequestedIOPS: 30000},
	}

	results, err := parseAllResults(reSumFile, sumFile)
	if err != nil {
		t.Fatal(err)
	}

	if len(results) != len(expectedResults) {
		t.Fatalf("Results count mismatch, expected %v instead got %v", len(expectedResults), len(results))
	}

	for i, r := range results {
		if r != expectedResults[i] {
			t.Fatalf("Mismatch: Got %v, expected %v", r, expectedResults[i])
		}
	}
}

// func TestParseMulti(t *testing.T) {
// 	input := "" +
// 		"SFS NFS THROUGHPUT:  100000 Ops/Sec   AVG. RESPONSE TIME:     1.0 Msec/Op\n" +
// 		"SFS NFS THROUGHPUT:  200000 Ops/Sec   AVG. RESPONSE TIME:     2.0 Msec/Op\n"

// 	results, err := parseAllResults(reResFile, input)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if len(results) != 2 {
// 		t.Fatalf("Expected 2 results, got %v", len(results))
// 	}

// 	avg := average(results)

// 	if avg.Latency != 1.5 {
// 		t.Fatal("Wrong latency")
// 	}
// 	if avg.IOPS != 150000 {
// 		t.Fatal("Wrong throughput")
// 	}
// }

// func TestParseFull(t *testing.T) {
// 	results, err := parseAllResults(reResFile, sfsresult)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if len(results) != 1 {
// 		t.Fatalf("Expected 1 results, got %v", len(results))
// 	}

// 	avg := average(results)

// 	if avg.Latency != 1.8 {
// 		t.Fatal("Wrong latency")
// 	}
// 	if avg.IOPS != 167352 {
// 		t.Fatal("Wrong throughput")
// 	}
// }

func XTestVerify(t *testing.T) {
	input := "" +
		"SFS NFS THROUGHPUT:  100000 Ops/Sec   AVG. RESPONSE TIME:     1.0 Msec/Op\n" +
		"SFS NFS THROUGHPUT:  200000 Ops/Sec   AVG. RESPONSE TIME:     2.0 Msec/Op\n"

	err := verify(input, expectedResults{numberOfRuns: 2, result: result{Latency: 1.5, IOPS: 155000}})
	if err != nil {
		t.Fatal(err)
	}
}

// The following tests ParseTemplate() by processing spec08 template
//
// To print out the processed template just run: 'go test -v [this_file]''
// func TestParseSfs2008Template(t *testing.T) {
// 	conf := types.Config{}

// 	conf.SetLoaders("loader1", "loader2", "loader3")
// 	conf.System.Frontend = "192.168.0.1"
// 	conf.Tesla.Elfs.Export = "my_export/dir"
// 	conf.Sfs2008 = sfs2008config.Config{
// 		LoadIo:             1,
// 		NumberOfRuns:       2,
// 		ProcessesPerClient: 3,
// 		IncrLoad:           4,
// 		Runtime:            5,
// 		WarmupTime:         6,
// 		MixFile:            "7",
// 		ConfigFile:         "8",
// 	}

// 	opts := &textutil.ParseOpts{
// 		Data:     conf,
// 		Template: bytes.NewReader([]byte(template)),
// 		Title:    "TemplateTest",
// 	}

// 	text, err := textutil.ParseTemplate(opts)
// 	if err != nil {
// 		t.Errorf("Failed parse, err: %v", err)
// 	}

// 	t.Logf("Parsed text: \n%v", text)
// }

func TestParseUnifiedResults(t *testing.T) {
	sumFileResults, err := parseAllResults(reResFile, sfsresult)
	if err != nil {
		t.Fatal(err)
	}

	var all []types.ToolUnifiedResults

	toolName := "sfs2008"
	for _, res := range sumFileResults {
		r := res // create a copy
		all = append(all, types.ToolUnifiedResults{
			ToolName:       &toolName,
			RequestedIOps:  &r.RequestedIOPS,
			IOps:           &r.IOPS,
			AverageLatency: &r.Latency,
		})
	}
	spew.Dump(all)
}
