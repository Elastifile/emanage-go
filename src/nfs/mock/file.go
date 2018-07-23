package nfsmock

import (
	log "gopkg.in/inconshreveable/log15.v2"

	"nfs"
	"nfs/sunrpc/nfsx"
	"size"
)

type MockFile struct {
	name string
	log.Logger
}

var _ nfs.File = (*MockFile)(nil) // Ensure interface compliance

func NewMockFile(name string) *MockFile {
	return &MockFile{
		name:   name,
		Logger: logger.New("mock", "MockFile"),
	}
}

func (mf *MockFile) isMockEntry() {}

func (mf *MockFile) Name() string { return mf.name }

func (mf *MockFile) Read(b []byte) (n int, err error) {
	mf.Debug("Read", "len", len(b))
	return 0, nil
}

func (mf *MockFile) ReadAt(b []byte, offset int64) (n int, err error) {
	mf.Debug("ReadAt", "len", len(b), "offset", offset)
	return 0, nil
}

func (mf *MockFile) Write(b []byte) (n int, err error) {
	mf.Debug("Write", "len", len(b))
	return 0, nil
}

func (mf *MockFile) WriteAt(b []byte, offset int64) (n int, err error) {
	mf.Debug("WriteAt", "len", len(b), "offset", offset)
	return 0, nil
}

func (mf *MockFile) ReadDir() (<-chan *nfs.DirEntry, error) {
	mf.Debug("ReadDir")
	return nil, nil
}

func (mf *MockFile) Truncate(size size.Size) (err error) {
	mf.Debug("Truncate", "size", size)
	return nil
}

func (mf *MockFile) Attr() (attr *nfsx.Fattr3) {
	return nil
}

func (mf *MockFile) GetAttr() (attr *nfsx.Fattr3, err error) {
	mf.Debug("GetAttr")
	return nil, nil
}

func (mf *MockFile) SetAttr(newAttr *nfsx.Sattr3) error {
	mf.Debug("SetAttr")
	return nil
}

func (mf *MockFile) Close() error {
	return nil
}

func (mf *MockFile) Seek(offset int64, whence int) (ret int64, err error) {
	mf.Debug("Seek", "offset", offset, "whence", whence)
	return 0, nil
}

func (mf *MockFile) Tell() (offset int64) {
	return 0
}

func (mf *MockFile) Pathconf() (pathconf *nfsx.Pathconf3resok, err error) {
	mf.Debug("Pathconf")
	return nil, nil
}

func (mf *MockFile) WriteSingle(b []byte) (n int, err error) {
	mf.Debug("WriteSingle", "len", len(b))
	return 0, nil
}
