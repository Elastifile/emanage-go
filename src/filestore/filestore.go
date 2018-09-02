package filestore

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-errors/errors"
	minio "github.com/minio/minio-go"

	"github.com/elastifile/emanage-go/src/logging"
	"github.com/elastifile/emanage-go/src/types"
)

var logger = logging.NewLogger("filestore")

func New(endpoint types.Host, verbose bool) (*Filestore, error) {
	addr := fmt.Sprintf("%v:%v", endpoint, port)
	logger.Debug("New filestore client", "Endpoint", addr)
	c, err := minio.New(addr, AccessKey, SecretKey, insecure)
	if err != nil {
		buf := make([]byte, 1<<16)
		stackSize := runtime.Stack(buf, true)
		err = errors.WrapPrefix(err, fmt.Sprintf("addr=%s stack=%s", addr, string(buf[0:stackSize])), 0)
		return nil, err
	}

	// just another health check
	attempts := 10
	for _, err = c.ListBuckets(); err != nil && attempts > 0; attempts -= 1 {
		time.Sleep(time.Second)
	}
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	m := &Filestore{c, endpoint, verbose}
	if verbose {
		m.SetTraces(true)
	}
	return m, nil
}

type Filestore struct {
	Client   *minio.Client
	Endpoint types.Host
	verbose  bool
}

func (m *Filestore) SetTraces(on bool) {
	m.Client.TraceOn(os.Stdout)
}

func (m *Filestore) trace(i interface{}) {
	if m.verbose {
		spew.Dump(i)
	}
}
