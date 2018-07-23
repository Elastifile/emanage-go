package nfs

import (
	"os"
	"path/filepath"

	log "gopkg.in/inconshreveable/log15.v2"

	"nfs/sunrpc/nfsx"
	terrors "tools/errors"
)

// NfsDirectory implements the Directory interface for an NFS directory
type NfsDirectory struct {
	log.Logger

	Fh     *nfsx.Fh3
	name   string
	Export string
	Nfs    *NfsClient
	BeforeCreateHook
	Parent *NfsDirectory
}

var _ Directory = (*NfsDirectory)(nil) // Ensure interface compliance

// NewNfsDirectory creates a new NfsDirectory value.
func NewNfsDirectory(fh *nfsx.Fh3, name string, export string, nfs *NfsClient) NfsDirectory {
	return NfsDirectory{
		Fh:     fh,
		name:   name,
		Export: export,
		Nfs:    nfs,
		Logger: log.New("fh", fh),
	}
}

func (d *NfsDirectory) Name() string {
	return d.name
}

func (d *NfsDirectory) clone(fh *nfsx.Fh3, name string) *NfsDirectory {
	var dir = *d
	dir.Fh = fh
	dir.name = name
	dir.Logger = log.New("fh", fh)
	return &dir
}

func (d *NfsDirectory) inherit(fh *nfsx.Fh3, name string) *NfsDirectory {
	dir := d.clone(fh, name)
	dir.Parent = d
	return dir
}

// SetLimits stores the provided limits for the NFS client.
func (d *NfsDirectory) SetLimits(limits *Limits) {
	d.Nfs.Limits = limits
}

func (d *NfsDirectory) Limits() *Limits {
	return d.Nfs.Limits
}

func (d *NfsDirectory) createFile(name string, flag int, perm os.FileMode) (file File, err error) {
	args := nfsx.Create3args{
		Where: nfsx.Diropargs3{
			Dir:  *d.Fh,
			Name: nfsx.Filename3(name),
		},
		How: nfsx.Createhow3{
			Mode: nfsx.UNCHECKED,
			Union: nfsx.Sattr3{
				Mode: nfsx.SetMode3{
					SetIt: true,
					Union: nfsx.Mode3(perm),
				},
			},
		},
	}
	d.Debug("create file", "name", name, "flag", flag, "perm", perm, "args", args)

	if d.BeforeCreateHook != nil {
		d.BeforeCreateHook(&args)
	}

	result, err := d.Nfs.Nfs3.Proc3Create(&args)
	if err != nil {
		return nil, err
	}
	if result.Status != nfsx.V3_OK {
		return nil, NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Failed to create file: %s in: %s",
			name, d.Name(),
		)
	}
	resOk := result.Union.(nfsx.Create3resok)
	if !resOk.Obj.HandleFollows {
		return nil, NewFSError(
			nil, "File %s created in %s, but no file handle received",
			name, d.Name(),
		)
	}

	fh := resOk.Obj.Union.(nfsx.Fh3)

	var attr nfsx.Fattr3
	if resOk.ObjAttributes.AttributesFollow {
		attr = resOk.ObjAttributes.Union.(nfsx.Fattr3)
	}

	return NewNfsFile(&fh, &attr, name, d.Nfs), nil
}

// OpenFile opens the NFS file with the provided flags and permissions.
// Supported flags are a subset of POSIX flags:
//   O_CREATE (0100)
func (d *NfsDirectory) OpenFile(name string, flag int, perm os.FileMode) (file File, err error) {
	d.Debug("OpenFile", "name", name)

	if flag&os.O_CREATE != 0 {
		// FIXME: Create only if the file doesn't exist yet! Do the lookup first.
		// Update: We should use GUARDED instead of UNCHECKED in createFile.
		return d.createFile(name, flag, perm)
	}

	fh, attr, err := d.Nfs.Lookup(d.Fh, nfsx.Filename3(name))
	if err != nil {
		return nil, NewFSError(err, "Failed to open file %s in %s", name, d.Name())
	}

	return NewNfsFile(fh, attr, name, d.Nfs), nil
}

// Open a file in read only mode.
func (d *NfsDirectory) Open(name string) (file File, err error) {
	d.Debug("Open", "path", d.PathNames(), "name", name)
	return d.OpenFile(name, os.O_RDONLY, 0)
}

// Create a new file. Truncate and open for read/write.
func (d *NfsDirectory) Create(name string) (file File, err error) {
	d.Debug("Create", "path", d.PathNames(), "name", name)
	return d.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

// BeforeCreate sets a hook that is executed before Create operations.
// This can be used to change the NFS Create arguments that are used to create the file.
func (d *NfsDirectory) BeforeCreate(hook BeforeCreateHook) {
	d.BeforeCreateHook = hook
}

// ReadDir reads the directory entries.
// It creates a new goroutine that sequenctially writes entries to a channel, and returns this channel.
func (d *NfsDirectory) ReadDir(name string) (entries <-chan *DirEntry, err error) {
	d.Debug("Read", "Dir", *d, "name", name)
	f, err := d.Open(name)
	if err != nil {
		return nil, err
	}

	file := f.(*NfsFile)

	if filetype := file.Attr().Type; filetype != nfsx.NF3DIR {
		return nil, NewFSError(
			nil,
			"%s (%v) is not a directory in %s",
			name, filetype, d.Name(),
		)
	}
	return file.ReadDir()
}

// Remove a file.
func (d *NfsDirectory) Remove(name string) (err error) {
	d.Debug("Remove", "path", d.PathNames(), "name", name)
	message := "Failed to remove file %s in %s because: %s"
	result, err := d.Nfs.Nfs3.Proc3Remove(&nfsx.Remove3args{
		Object: nfsx.Diropargs3{
			Dir:  *d.Fh,
			Name: nfsx.Filename3(name),
		},
	})
	if err != nil {
		return NewFSError(err, message, name, d.Name())
	}
	if result.Status != nfsx.V3_OK {
		return NewFSError(
			&nfsx.NfsError{Status: result.Status},
			message, name, d.Name(),
		)
	}

	// removeOk := result.Union.(nfsx.Remove3resok)
	return nil
}

// Rmdir removes a directory.
func (d *NfsDirectory) Rmdir(name string) (err error) {
	d.Debug("Rmdir", "path", d.PathNames(), "name", name)
	message := "Failed to remove directory %s"
	result, err := d.Nfs.Nfs3.Proc3Rmdir(&nfsx.Rmdir3args{
		Object: nfsx.Diropargs3{
			Dir:  *d.Fh,
			Name: nfsx.Filename3(name),
		},
	})
	if err != nil {
		return NewFSError(err, message, name)
	}
	if result.Status != nfsx.V3_OK {
		return NewFSError(&nfsx.NfsError{Status: result.Status}, message, name)
	}

	// removeOk := result.Union.(nfsx.Rmdir3resok)
	return nil
}

func (d *NfsDirectory) RemoveAll() (err error) {
	d.Debug("RemoveAll")
	entries, err := d.ReadDir(".")
	if err != nil {
		return NewFSError(
			err,
			"Failed to read directory %s while removing",
			d.Name(),
		)
	}

	for entry := range entries {
		if entry.Name != "." && entry.Name != ".." {
			f, err := d.Open(entry.Name)
			if err != nil {
				return NewFSError(
					err,
					"Failed to open file %s in %s while removing",
					entry.Name, d.Name(),
				)
			}

			file := f.(*NfsFile)

			filetype := file.Attr().Type
			if filetype == nfsx.NF3DIR {
				dir := NewNfsDirectory(file.fh, file.name, d.Export, d.Nfs)
				if err = dir.RemoveAll(); err != nil {
					return err
				}
				if err = d.Rmdir(dir.Name()); err != nil {
					return err
				}
			} else if err := d.Remove(entry.Name); err != nil {
				return err
			}
		}
	}
	return nil
}

// Rename a file and/or move it to another directory.
func (d *NfsDirectory) Rename(fromName string, toDir Directory, toName string) error {
	d.Debug("Rename", "fromName", fromName, "toDir", toDir.PathNames(), "toName", toName)

	nfsToDir, ok := toDir.(*NfsDirectory)
	if !ok {
		return NewFSError(nil, "%s must be an %T, but is %T", toDir.Name(), d, toDir)
	}

	result, err := d.Nfs.Nfs3.Proc3Rename(&nfsx.Rename3args{
		From: nfsx.Diropargs3{
			Dir:  *d.Fh,
			Name: nfsx.Filename3(fromName),
		},
		To: nfsx.Diropargs3{
			Dir:  *nfsToDir.Fh,
			Name: nfsx.Filename3(toName),
		},
	})
	if err != nil {
		return NewFSError(
			err, "Failed to rename directory %s/%s -> %s/%s",
			d.Name(), fromName, toDir.Name(), toName,
		)
	}
	if result.Status != nfsx.V3_OK {
		return NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Failed to rename directory %s/%s -> %s/%s",
			d.Name(), fromName, toDir.Name(), toName,
		)
	}

	return nil
}

// Link creates a hard link to a file.
func (d *NfsDirectory) Link(file File, name string) error {
	nfsFile, ok := file.(*NfsFile)
	if !ok {
		return NewFSError(nil, "%v must be %T, but is %T", file, &NfsFile{}, file)
	}

	args := nfsx.Link3args{
		File: *nfsFile.fh,
		Link: nfsx.Diropargs3{
			Dir:  *d.Fh,
			Name: nfsx.Filename3(name),
		},
	}
	d.Debug("Link", "file", file, "name", name, "args", args)

	result, err := d.Nfs.Nfs3.Proc3Link(&args)
	if err != nil {
		return NewFSError(err, "Linking %v to %s in %s failed", nfsFile, name, d)
	}
	if result.Status != nfsx.V3_OK {
		return NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Linking %v to %s in %s failed", nfsFile, name, d,
		)
	}

	return nil
}

// Link creates a hard link to a file.
func (d *NfsDirectory) Symlink(file File, name string) error {
	nfsFile, ok := file.(*NfsFile)
	if !ok {
		return NewFSError(nil, "%v must be %T, but is %T", file, &NfsFile{}, file)
	}

	args := nfsx.Symlink3args{
		Where: nfsx.Diropargs3{
			Dir:  *d.Fh,
			Name: nfsx.Filename3(name),
		},
		Symlink: nfsx.Symlinkdata3{
			SymlinkData: nfsx.Path3([]byte(file.Name())),
		},
	}
	d.Debug("Symlink", "file", file, "name", name, "args", args)

	result, err := d.Nfs.Nfs3.Proc3Symlink(&args)
	if err != nil {
		return NewFSError(err, "Failed Symlinking %v from %s in %s", nfsFile, name, d)
	}
	if result.Status != nfsx.V3_OK {
		return NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Failed Symlinking %v from %s in %s", nfsFile, name, d,
		)
	}

	return nil
}

// Mkdir creates a directory with the provided POSIX permissions.
func (d *NfsDirectory) Mkdir(name string, perm uint32) (Directory, error) {
	d.Debug("Mkdir", "path", d.PathNames(), "name", name, "perm", perm)
	message := "Failed to create directory %s"
	result, err := d.Nfs.Nfs3.Proc3Mkdir(&nfsx.Mkdir3args{
		Where: nfsx.Diropargs3{
			Dir:  *d.Fh,
			Name: nfsx.Filename3(name),
		},
		Attributes: nfsx.Sattr3{
			Mode: nfsx.SetMode3{
				SetIt: true,
				Union: nfsx.Mode3(perm), // FIXME: Use umask?
			},
		},
	})
	if err != nil {
		return nil, NewFSError(err, message, name)
	}
	if result.Status != nfsx.V3_OK {
		return nil, NewFSError(
			&nfsx.NfsError{Status: result.Status}, message, name,
		)
	}

	resOk := result.Union.(nfsx.Mkdir3resok)
	if !resOk.Obj.HandleFollows {
		return nil, NewFSError(nil,
			"Created %s in %s, but did not receive file handle",
			name, d.Name(),
		)
	}

	newFh := resOk.Obj.Union.(nfsx.Fh3)
	return d.inherit(&newFh, name), nil
}

func (d *NfsDirectory) MkdirAll(names []string, perm uint32) (Directory, error) {
	logger.Debug("MkdirAll", "names", names, "perm", perm)

	if len(names) == 0 {
		return d, nil
	}

	sub, err := d.LookupDir(names[0])
	if err != nil {
		sub, err = d.Mkdir(names[0], perm)
		if err != nil {
			return nil, err
		}
	}

	return sub.MkdirAll(names[1:], perm)
}

func (d *NfsDirectory) Root() Directory {
	if d.Parent == nil {
		return d
	}
	return d.Parent.Root()
}

func (d *NfsDirectory) Path() []Directory {
	if d.Parent == nil {
		return []Directory{}
	}
	if d.Name() != "." {
		return append(d.Parent.Path(), d)
	}
	return d.Parent.Path()
}

func (d *NfsDirectory) PathNames() []string {
	path := d.Path()
	result := make([]string, len(path))
	for i, dir := range path {
		result[i] = dir.Name()
	}
	return result
}

func (d *NfsDirectory) PathString() string {
	return string(filepath.Separator) + filepath.Join(d.PathNames()...)
}

func (d *NfsDirectory) Node(typ nfsx.Ftype3, name string, perm os.FileMode) error {
	args := nfsx.Mknod3args{
		Where: nfsx.Diropargs3{
			Dir:  *d.Fh,
			Name: nfsx.Filename3(name),
		},
		What: nfsx.Mknoddata3{
			Type: typ,
		},
	}

	switch typ {
	case nfsx.NF3BLK:
		args.What.Union = nfsx.Devicedata3{
			Spec: nfsx.Specdata3{},
		}
	case nfsx.NF3FIFO:
		args.What.Union = nfsx.Sattr3{
			Mode: nfsx.SetMode3{
				SetIt: true,
				Union: nfsx.Mode3(perm),
			},
		}
	}

	d.Debug("Node", "type", typ, "name", name, "args", args)

	result, err := d.Nfs.Nfs3.Proc3Mknod(&args)
	if err != nil {
		return NewFSError(err, "Failed creating node of type %s in %s", typ, name, d)
	}
	if result.Status != nfsx.V3_OK {
		return NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Failed creating node of type %s in %s", typ, d,
		)
	}

	return nil
}

// LookupDir looks up a directory by name and returns a matching Directory object.
func (d *NfsDirectory) LookupDir(name string) (Directory, error) {
	logger.Debug("LookupDir", "path", d.PathNames(), "name", name)

	f, err := d.Open(name)
	if err != nil {
		return nil, NewFSError(err, "Lookup %s in %s failed", name, d.Name())
	}

	file := f.(*NfsFile)

	if filetype := file.Attr().Type; filetype != nfsx.NF3DIR {
		return nil, NewFSError(
			ErrNotADirectory,
			"Lookup %s in %s failed", name, d.Name(),
		)
	}

	return d.inherit(file.fh, name), nil
}

// FsInfo returns the NFS FSINFO data (static filesystem information).
func (d *NfsDirectory) FsInfo() (fsinfo *nfsx.Fsinfo3resok, err error) {
	d.Debug("FsInfo")
	result, err := d.Nfs.Nfs3.Proc3Fsinfo(&nfsx.Fsinfo3args{
		Fsroot: *d.Fh,
	})
	if err != nil {
		return nil, NewFSError(
			err,
			"Couldn't fetch file-system info from %s", d.Name(),
		)
	}
	if result.Status != nfsx.V3_OK {
		return nil, NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Couldn't fetch file-system info from %s", d.Name(),
		)
	}

	resOk := result.Union.(nfsx.Fsinfo3resok)
	return &resOk, nil
}

// FsStat returns the NFS FSSTAT data (dynamic file system information).
func (d *NfsDirectory) FsStat() (fsstat *nfsx.Fsstat3resok, err error) {
	d.Debug("FsStat")
	result, err := d.Nfs.Nfs3.Proc3Fsstat(&nfsx.Fsstat3args{
		Fsroot: *d.Fh,
	})
	if err != nil {
		return nil, NewFSError(
			err,
			"Couldn't fetch file-system statistics from %s", d.Name(),
		)
	}
	if result.Status != nfsx.V3_OK {
		return nil, NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Couldn't fetch file-system statistics from %s", d.Name(),
		)
	}

	resOk := result.Union.(nfsx.Fsstat3resok)
	return &resOk, nil
}

func (d *NfsDirectory) WalkDir(name string, fn func(*DirEntry) error) error {
	logger.Debug("Walking dir", "path", d.PathNames(), "name", name)

	entries, err := d.ReadDir(name)
	if err != nil {
		return err
	}

	i := 0
	errChan := make(chan error)
	n := d.Nfs.ConnectionCount()

	for entry := range entries {
		if entry.Err != nil {
			return entry.Err
		}

		if entry.Name == "." || entry.Name == ".." {
			continue
		}

		e := entry
		go func() {
			errChan <- fn(e)
		}()
		i += 1

		if i == n {
			err = terrors.CollectAll(errChan, i)
			if err != nil {
				return err
			}
			i = 0
		}
	}

	return terrors.CollectAll(errChan, i)
}
