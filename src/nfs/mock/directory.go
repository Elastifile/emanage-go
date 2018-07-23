package nfsmock

import (
	"fmt"
	"os"
	"strings"

	log "gopkg.in/inconshreveable/log15.v2"

	"nfs"
	"nfs/sunrpc/nfsx"
)

var logger = log.New("package", "nfsmock")

type MockDirectory struct {
	name    string
	entries []MockEntry
	log.Logger
}

type MockEntry interface {
	isMockEntry()
	Name() string
}

var _ nfs.Directory = (*MockDirectory)(nil) // Ensure interface compliance

func NewMockDirectory(name string, entries ...MockEntry) *MockDirectory {
	return &MockDirectory{
		name:    name,
		entries: entries,
		Logger:  logger.New("mock", "MockDirectory"),
	}
}

func (md *MockDirectory) isMockEntry() {}

func (md *MockDirectory) Name() string { return md.name }

func (md *MockDirectory) String() string {
	var names []string
	for _, entry := range md.entries {
		names = append(names, entry.Name())
	}
	return fmt.Sprintf("%v [%v]", md.name, strings.Join(names, " "))
}

func (md *MockDirectory) ReadDir(name string) (<-chan *nfs.DirEntry, error) {
	md.Debug("ReadDir", "name", name)

	entries := []MockEntry{NewMockDirectory("."), NewMockDirectory("..")}
	entries = append(entries, md.entries...)

	dirEntries := make(chan *nfs.DirEntry, len(entries))
	for _, entry := range entries {
		md.Debug("ReadDir:", "name", name, "entry", entry.Name())
		dirEntries <- &nfs.DirEntry{
			Name: entry.Name(),
		}
	}
	close(dirEntries)

	return dirEntries, nil
}

func (md *MockDirectory) LookupDir(name string) (dir nfs.Directory, err error) {
	l := md.Logger.New("name", name)
	l.Debug("LookupDir")

	for _, entry := range md.entries {
		if entry.Name() == name {
			if dir, ok := entry.(*MockDirectory); ok {
				l.Debug("LookupDir: found directory")
				return dir, nil
			}
			l.Debug("LookupDir: found non-directory")
			return nil, nfs.ErrNotADirectory
		}
	}

	l.Debug("LookupDir: not found")
	return nil, &nfsx.NfsError{Status: nfsx.V3ERR_NOENT}
}

func (md *MockDirectory) SetLimits(limits *nfs.Limits) {
	md.Debug("SetLimits", "limits", limits)
}

func (md *MockDirectory) Mkdir(name string, perm uint32) (dir nfs.Directory, err error) {
	md.Debug("Mkdir", "name", name, "perm", perm)
	return NewMockDirectory(name), nil
}

func (md *MockDirectory) OpenFile(name string, flag int, perm os.FileMode) (file nfs.File, err error) {
	md.Debug("OpenFile", "name", name)
	return NewMockFile(name), nil
}

func (md *MockDirectory) Open(name string) (file nfs.File, err error) {
	md.Debug("Open", "name", name)
	return md.OpenFile(name, os.O_RDONLY, 0)
}

func (md *MockDirectory) Create(name string) (file nfs.File, err error) {
	md.Debug("Create", "name", name)
	return md.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

func (md *MockDirectory) Remove(name string) (err error) {
	md.Debug("Remove", "name", name)
	return nil
}

func (md *MockDirectory) Rmdir(name string) (err error) {
	md.Debug("Rmdir", "name", name)
	return nil
}

func (md *MockDirectory) RemoveAll() (err error) {
	md.Debug("RemoveAll")
	return nil
}

func (md *MockDirectory) Rename(fromName string, toDir nfs.Directory, toName string) error {
	md.Debug("Rename", "name", fromName, "toDir", toDir, "toName", toName)
	return nil
}

func (md *MockDirectory) Link(file nfs.File, name string) error {
	md.Debug("Link", "name", name)
	return nil
}

func (md *MockDirectory) FsInfo() (*nfsx.Fsinfo3resok, error) {
	md.Debug("FsInfo")
	return nil, nil
}

func (md *MockDirectory) FsStat() (*nfsx.Fsstat3resok, error) {
	md.Debug("FsStat")
	return nil, nil
}

func (md *MockDirectory) BeforeCreate(hook nfs.BeforeCreateHook) {
	md.Debug("BeforeCreate")
}

func (md *MockDirectory) Limits() *nfs.Limits {
	md.Debug("Limits")
	return nil
}
