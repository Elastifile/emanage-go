package erun_config

import (
	"size"
	"testing"
	"time"
)

func TestProfileMetafuzzerDefaults(t *testing.T) {
	c := ProfileMetafuzzer{}
	got := c.ToCmd()
	expected := ""
	if got != expected {
		t.Fatalf("\nGot:\n\t`%v`\nExpected:\n\t`%v`", got, expected)
	}
}

func TestProfileMetafuzzerAll(t *testing.T) {
	c := ProfileMetafuzzer{
		Duration:         1 * time.Minute,
		Clients:          2,
		QueueSize:        3,
		TargetNrDirs:     4,
		TargetNrFiles:    5,
		LogFile:          "6.txt",
		MaxLogFileSize:   7 * size.KiB,
		EndStateDumpFile: "8.foo",
	}
	got := c.ToCmd()
	expected := `--duration 60 --clients 2 --queue-size 3 --target-nr-dirs 4 --target-nr-files 5 --log-file "6.txt" --max-log-file-size 7168 --end-state-dump-file "8.foo"`
	if got != expected {
		t.Fatalf("Got:\n\t%v\nExpected:\n\t%v", got, expected)
	}
}
