// The nfs package implements an NFS client.
//
// It implements the NFS, MOUNT and PORTMAP protocols on top of SunRPC.
package nfs

import (
	"io"
	"os"

	"nfs/sunrpc/nfsx"
	"size"
)

// The Directory interface includes operations that can be performed on a directory.
type Directory interface {
	Name() string
	BeforeCreate(hook BeforeCreateHook)
	Create(name string) (File, error)
	LookupDir(name string) (Directory, error)
	Mkdir(name string, perm uint32) (Directory, error)
	MkdirAll(names []string, perm uint32) (Directory, error)
	Root() Directory
	Path() []Directory
	PathString() string
	PathNames() []string
	Open(name string) (File, error)
	OpenFile(name string, flag int, perm os.FileMode) (File, error)
	ReadDir(name string) (<-chan *DirEntry, error)
	WalkDir(name string, fn func(*DirEntry) error) error
	Remove(name string) error
	Rmdir(name string) error
	RemoveAll() error
	Rename(fromName string, toDir Directory, toName string) error
	Link(file File, name string) error
	Symlink(file File, name string) error
	Node(typ nfsx.Ftype3, name string, perm os.FileMode) error
	FsInfo() (*nfsx.Fsinfo3resok, error)
	FsStat() (*nfsx.Fsstat3resok, error)
	SetLimits(*Limits)
	Limits() *Limits
}

// BeforeCreateHook describes a function to be called before creating a file.
type BeforeCreateHook func(*nfsx.Create3args)

// The File interface includes operations that can be performed on a file.
type File interface {
	io.ReadWriteSeeker
	io.ReaderAt
	io.WriterAt
	io.Closer
	// Stat() (fi os.FileInfo, err error)
	Truncate(size size.Size) error
	// Sync() error
	GetAttr() (*nfsx.Fattr3, error)
	SetAttr(*nfsx.Sattr3) error
	ReadDir() (<-chan *DirEntry, error)
	Name() string
	Attr() *nfsx.Fattr3
	Pathconf() (*nfsx.Pathconf3resok, error)
	Tell() int64
	WriteSingle([]byte) (int, error)
}

// DirEntry represents an NFS directory entry.
type DirEntry struct {
	Name   string
	Fileid uint64
	Cookie uint64
	Err    error
}

// Limits for file operations.
type Limits struct {
	MaxReadSize     size.Size
	MaxWriteSize    size.Size
	PrefReadDirSize size.Size
	MaxOffset       size.Size
}
