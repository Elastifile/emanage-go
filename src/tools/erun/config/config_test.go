package erun_config

import (
	"testing"
	"time"

	"size"
)

func TestProfileIODefaults(t *testing.T) {
	c := ProfileIO{}
	got := c.ToCmd()
	expected := "--duration 180 --max-file-size 104857600 --queue-size 16 --readwrites 70"
	if got != expected {
		t.Fatalf("Got:\n\t%v\nExpected:\n\t%v", got, expected)
	}
}

// func TestProfileIOBool(t *testing.T) {
// 	c := ProfileIO{
// 		Duration:    1 * time.Second,
// 		MaxFileSize: 2 * size.Bytes,
// 		QueueSize:   3,
// 		Readwrites:  "4",
// 		DataPayload: true,
// 	}
// 	got := c.ToCmd()
// 	expected := "--duration 1 --max-file-size 2 --queue-size 3 --readwrites 4 --data-payload"
// 	if got != expected {
// 		t.Fatalf("Got:\n\t%v\nExpected:\n\t%v", got, expected)
// 	}
// }

// func TestProfileIOBoolAll(t *testing.T) {
// 	c := ProfileIO{
// 		Duration:          1 * time.Second,
// 		MaxFileSize:       2 * size.Bytes,
// 		QueueSize:         3,
// 		Readwrites:        "4",
// 		DataPayload:       true,
// 		InitialWritePhase: true,
// 	}
// 	got := c.ToCmd()
// 	expected := "--duration 1 --max-file-size 2 --queue-size 3 --readwrites 4 --data-payload --initial-write-phase"
// 	if got != expected {
// 		t.Fatalf("Got:\n\t%v\nExpected:\n\t%v", got, expected)
// 	}
// }

// func TestProfileIOAll(t *testing.T) {
// 	c := ProfileIO{
// 		Duration:          1 * time.Hour,
// 		Clients:           2,
// 		NrFiles:           3,
// 		MaxFileSize:       4 * size.GiB,
// 		QueueSize:         5,
// 		Readwrites:        "6",
// 		DataPayload:       true,
// 		InitialWritePhase: true,
// 		MaxIoSize:         7 * size.KiB,
// 		MinIoSize:         8 * size.MiB,
// 		RecoveryTimeout:   9 * time.Minute,
// 	}
// 	got := c.ToCmd()
// 	// REVIEW(orenz): Are we sure that erun is not expected to parse a string after '--max-file-size' and parse it himself?
// 	// REVIEW(gavriep): Erun expects either a pure number or one with a K/M/G etc. suffix. It considers '4096' and '4K' to be exactly the same thing.
// 	expected := "--duration 3600 --clients 2 --nr-files 3 --max-file-size 4294967296 --queue-size 5 --readwrites 6 " +
// 		"--data-payload --initial-write-phase --max-io-size 7168 --min-io-size 8388608 --recovery-timeout 540"
// 	if got != expected {
// 		t.Fatalf("Got:\n\t%v\nExpected:\n\t%v", got, expected)
// 	}
// }

func TestProfileIOMixed(t *testing.T) {
	c := ProfileIO{
		Duration:    30 * time.Second,
		QueueSize:   4,
		MaxIoSize:   8 * size.KiB,
		DataPayload: true,
	}
	got := c.ToCmd()
	expected := "--duration 30 --max-file-size 104857600 --queue-size 4 --readwrites 70 --data-payload --max-io-size 8192"
	if got != expected {
		t.Fatalf("Got:\n\t%v\nExpected:\n\t%v", got, expected)
	}
}

// func TestProfileIOErunWriteLC(t *testing.T) {
// 	erunWriteLC := ProfileIO{
// 		Duration:          10 * time.Minute,
// 		Clients:           4,
// 		NrFiles:           4,
// 		MaxFileSize:       100 * size.MiB,
// 		QueueSize:         20,
// 		Readwrites:        "70",
// 		DataPayload:       true,
// 		InitialWritePhase: true,
// 		InitialWriteStop:  true,
// 		MinIoSize:         4 * size.KiB,
// 		MaxIoSize:         4 * size.KiB,
// 	}
// 	got := erunWriteLC.ToCmd()
// 	expected := `--duration 600 --clients 4 --nr-files 4 --max-file-size 104857600 --queue-size 20 --readwrites 70 --data-payload` +
// 		` --initial-write-phase --initial-write-stop --max-io-size 4096 --min-io-size 4096`
// 	if got != expected {
// 		t.Fatalf("Got:\n\t%v\nExpected:\n\t%v", got, expected)
// 	}
// }

// func TestErunConfigGeneric(t *testing.T) {
// 	table := []struct {
// 		conf Config
// 		json string
// 		cmd  string
// 	}{
// 		{
// 			Config{
// 				Profile: &ProfileIO{
// 					NrFiles: 7,
// 				},
// 			},
// 			`{"profile":{"Duration":0,"Clients":0,"NrFiles":7,"MaxFileSize":0,"QueueSize":0,"Readwrites":0,"DataPayload":false,"InitialWritePhase":false,"InitialWriteStop":false,"ReuseExistingFiles":false,"MaxIoSize":0,"MinIoSize":0,"RecoveryTimeout":0,"ShuffleFiles":0,"CncTraces":0,"MinUncomp":0,"MaxUncomp":0},"profile-name":"io"}`,
// 			`--duration 180 --nr-files 7 --max-file-size 104857600 --queue-size 16 --readwrites 70`,
// 		},
// 		{
// 			Config{
// 				Profile: &ProfileIO{
// 					Duration:          10 * time.Minute,
// 					Clients:           4,
// 					NrFiles:           4,
// 					MaxFileSize:       100 * size.MiB,
// 					QueueSize:         20,
// 					Readwrites:        "70",
// 					DataPayload:       true, // Oren says he didn't see these
// 					InitialWritePhase: true, // as above
// 					InitialWriteStop:  true,
// 					MinIoSize:         4 * size.KiB,
// 					MaxIoSize:         4 * size.KiB,
// 				},
// 			},
// 			`{"profile":{"Duration":600000000000,"Clients":4,"NrFiles":4,"MaxFileSize":104857600,"QueueSize":20,"Readwrites":70,"DataPayload":true,"InitialWritePhase":true,"InitialWriteStop":true,"ReuseExistingFiles":false,"MaxIoSize":4096,"MinIoSize":4096,"RecoveryTimeout":0,"ShuffleFiles":0,"CncTraces":0,"MinUncomp":0,"MaxUncomp":0},"profile-name":"io"}`,
// 			`--duration 600 --clients 4 --nr-files 4 --max-file-size 104857600 --queue-size 20 --readwrites 70 --data-payload --initial-write-phase --initial-write-stop --max-io-size 4096 --min-io-size 4096`,
// 		},
// 	}

// 	for _, test := range table {
// 		ec := test.conf
// 		expected := test.json

// 		// Marshal
// 		js, err := json.Marshal(ec)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		got := string(js)
// 		if got != expected {
// 			t.Fatalf("\nGot:\n\t%v\nExpected:\n\t%v", got, expected)
// 		}

// 		// Unmarshal
// 		var ec2 Config
// 		err = json.Unmarshal(js, &ec2)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if ec2.Profile.Name() != ec.Profile.Name() {
// 			t.Fatalf("\nGot:\n\t%v\nExpected:\n\t%v", ec2.Profile.Name(), ec.Profile.Name())
// 		}

// 		cmd := ec.Profile.ToCmd()
// 		if test.cmd != cmd {
// 			t.Fatalf("\nGot:\n\t%v\nExpected:\n\t%v", test.cmd, cmd)
// 		}
// 	}
// }
