package erun

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/koding/multiconfig"

	"github.com/elastifile/emanage-go/src/size"
	erun_config "github.com/elastifile/emanage-go/src/tools/erun/config"
	"github.com/elastifile/emanage-go/src/types"
)

func TestErunRet(t *testing.T) {
	fmt.Printf("Erun error: %d (%[1]v)\n", ErunRetErrorDataVerifError)
}

func TestParseOptions(t *testing.T) {
	// --json-output \"/data/'d49156761a2c'/er\\ un/erun _full_log.json\"
	//
	// is not a profile option, but a runtime option.  Currently,
	// there is no way to pass runtime options.
	tomlFile := strings.NewReader(`
[erun]
profile = "io"
cmd = " --erun-dir loader5a-tesla --duration 6000 --clients 4 --nr-files 4 --max-file-size '100M' --queue-size 20 --readwrites 70 --data-payload --initial-write-phase --initial-write-stop --min-io-size 4K --max-io-size 4K"
`)
	tomlLoader := multiconfig.TOMLLoader{
		Reader: tomlFile,
	}
	conf := types.Config{}
	err := tomlLoader.Load(&conf)
	if err != nil {
		t.Fatalf("Loading toml failed: %+v", err)
	}

	io, ok := conf.Erun.Profile.(*erun_config.ProfileIO)
	if !ok {
		t.Fatalf("Expected IO profile but got: %T", conf.Erun)
	}
	if io.Duration != 6000*time.Second {
		t.Fatalf("Expected duration 6000s got: %s", io.Duration)
	}
	if io.Clients != 4 {
		t.Fatalf("Expected 4 clients got: %d", io.Clients)
	}
	if io.NrFiles != 4 {
		t.Fatalf("Expected 4 files got: %d", io.NrFiles)
	}
	if io.MaxFileSize != 100*size.MiB {
		t.Fatalf("Expected max file size %s got: %s", 100*size.MiB, io.MaxFileSize)
	}
	if io.QueueSize != 20 {
		t.Fatalf("Expected queue size 20 got: %d", io.QueueSize)
	}
	if io.Readwrites != "70" {
		t.Fatalf("Expected readwrite 70 got: %v", io.Readwrites)
	}
	if !io.DataPayload {
		t.Fatalf("Expected data payload to be true")
	}
	if !io.InitialWritePhase {
		t.Fatalf("Expected initial write phase to be true")
	}
	if !io.InitialWriteStop {
		t.Fatalf("Expected initial write stop to be true")
	}
	if io.MaxIoSize != 4*size.KiB {
		t.Fatalf("Expected max I/O size %s got: %s", 4*size.KiB, io.MaxIoSize)
	}
	if io.MinIoSize != 4*size.KiB {
		t.Fatalf("Expected min I/O size %s got: %s", 4*size.KiB, io.MinIoSize)
	}
}
