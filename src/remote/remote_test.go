package remote

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"

	"logging"
	logging_config "logging/config"
)

func TestRemote(t *testing.T) {
	remoteIp := os.Getenv("REMOTE_IP")
	if remoteIp == "" {
		t.Fatal("must predefine $REMOTE_IP")
	}

	rem, err := NewRemote(remoteIp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Created client %v for %s", *rem.Client, rem.address)

	out, err := rem.Run("ls")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ran 'ls':\n%s", out)

	out, err = rem.Run("whoami")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ran 'whoami':\n%s", out)

	err = rem.UploadFile("/etc/passwd", "/root/foobar")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Uploaded /etc/passwd to /root/foobar")
}

func TestKeyParse(t *testing.T) {
	b, err := extractPublicKey([]byte(privateKey))

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n", string(b))
}

func TestPingFromRemoteServer(t *testing.T) {
	// integration test (to be run manually)
	t.SkipNow()
	ok, err := PingFromRemoteServer("10.11.209.25", "127.0.0.1")
	if !ok {
		t.Errorf("Test Failed.  Error: %s", err)
	}
}

func TestValidHostName(t *testing.T) {
	t.SkipNow() // integration test

	host := "10.11.193.223"
	err := ValidateHostname(host)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseFreeMem(t *testing.T) {
	str := "              total        used        free      shared  buff/cache   available\n" +
		"Mem:            47G         38G        5.4G        2.2G        3.6G        6.5G\n" +
		"Swap:            0B          0B          0B"

	result, err := parseMemoryStats(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetMemStats(t *testing.T) {
	t.Skip() // integration test

	data, err := GetMemoryStats("10.11.181.249")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestReadDebugfsCounters(t *testing.T) {
	t.Skip() // integration test

	counters := []string{
		"allocator_memory",
		"traces_memory",
		"dstore_memory_required_ram",
		"pre_config_registration_size",
		"pre_config_allocation_size",
		"post_config_registration_size",
	}
	data, err := ReadDebugfsCounters("10.11.182.240", counters)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(data)
}

func TestSshTunnel(t *testing.T) {
	logging.Setup(&logging_config.Config{Level: "debug"})

	opts := NewSshTunnelConfig()

	opts.SourceIp = os.Getenv("SOURCE_IP")
	if opts.SourceIp == "" {
		t.Fatal("must predefine $SOURCE_IP")
	}
	t.Logf("source IP: %s\n", opts.SourceIp)

	opts.TargetIp = os.Getenv("TARGET_IP")
	if opts.TargetIp == "" {
		t.Fatal("must predefine $TARGET_IP")
	}
	t.Logf("target IP: %s\n", opts.TargetIp)

	closer, err := OpenSshTunnel(opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Opened SSH tunnel")

	sleep := 10 * time.Second
	sleepStr := os.Getenv("SLEEP")
	if sleepStr != "" {
		sleep, err = time.ParseDuration(sleepStr)
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("Sleep for %v", sleep)
	time.Sleep(sleep)

	err = closer()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Closed the tunnel.")
}
