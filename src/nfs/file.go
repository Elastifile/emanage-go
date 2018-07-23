package nfs

import (
	"fmt"
	"io"

	xdr "github.com/davecgh/go-xdr/xdr2"
	multierror "github.com/hashicorp/go-multierror"
	log "gopkg.in/inconshreveable/log15.v2"

	"nfs/sunrpc/nfsx"
	"size"
)

const MaxWriteBufferSize = 10 * size.MiB

// NfsFile implements the File interface for an NFS file.
type NfsFile struct {
	attr     *nfsx.Fattr3 // Initial attr of file after create/lookup
	position uint64
	fh       *nfsx.Fh3
	name     string
	nfs      *NfsClient

	log.Logger
}

var _ File = (*NfsFile)(nil) // Ensure interface compliance

// NewNfsFile creates a new NfsFile value.
func NewNfsFile(fh *nfsx.Fh3, attr *nfsx.Fattr3, name string, nfsClient *NfsClient) *NfsFile {
	return &NfsFile{
		fh:     fh,
		attr:   attr,
		name:   name,
		nfs:    nfsClient,
		Logger: log.New("fh", fh),
	}
}

// Name returns the name of the file
func (f *NfsFile) Name() string {
	return f.name
}

// Attr returns the cached NFS attributes from the last NFS GETATTR operation.
func (f *NfsFile) Attr() *nfsx.Fattr3 {
	return f.attr
}

// Seek sets the offset for the next Read or Write on file to the specified
// offset, which is interpreted according to whence: 0 means relative to the
// origin of the file, 1 means relative to the current offset, and 2 means
// relative to the end.  It returns the new offset and an error, if any.
func (f *NfsFile) Seek(offset int64, whence int) (int64, error) {
	f.Debug("Seek", "offset", offset, "whence", whence)

	switch whence {
	case 0:
		f.position = uint64(offset)
	case 1:
		f.position += uint64(offset)
	case 2:
		// FIXME: The size may not be up to date, need to GetAttr() first
		f.position = uint64(f.attr.Size) + uint64(offset)
	}
	return int64(f.position), nil
}

// Tell returns the current position in the file
func (f *NfsFile) Tell() (offset int64) {
	return int64(f.position)
}

// Close the file (actually this does nothing for NFS)
func (f *NfsFile) Close() error {
	return nil
}

// Read reads up to len(b) bytes from the File.
// It returns the number of bytes read and an error, if any.
// EOF is signaled by a zero count with err set to io.EOF.
func (f *NfsFile) Read(b []byte) (n int, err error) {
	f.Debug("Read", "len", len(b), "file", f.name)
	return f.read(b, 0, true)
}

// ReadAt reads len(b) bytes from the File starting at byte offset offset.
// It returns the number of bytes read and the error, if any.
// ReadAt always returns a non-nil error when n < len(b).
// At end of file, that error is io.EOF.
func (f *NfsFile) ReadAt(b []byte, offset int64) (n int, err error) {
	defer func() { f.Debug("ReadAt ->", "len", len(b), "offset", offset, "n", n, "err", err) }()
	f.Debug("ReadAt", "len", len(b), "offset", offset)
	// This code is based on os.File.ReadAt()
	for len(b) > 0 {
		var m int
		m, err = f.read(b, uint64(offset), false)
		n += m
		if err != nil {
			return n, err
		}
		b = b[m:]
		offset += int64(m)
	}
	return
}

// According to RFC 1813, a read request that has a Count that is larger than
// what the server supports (rtmax) will result in a short read of rtmax bytes or less.
// This is perfectly compatible with the io.Reader interface description.
// Source: https://tools.ietf.org/html/rfc1813#section-3.3.6
func (f *NfsFile) read(b []byte, offset uint64, useCurrentPosition bool) (n int, err error) {
	if useCurrentPosition {
		offset = f.position
	}
	if f.nfs.Limits != nil {
		if len(b) >= int(f.nfs.Limits.MaxReadSize) {
			b = b[:f.nfs.Limits.MaxReadSize]
		}
	}

	result, err := f.nfs.Nfs3.Proc3Read(&nfsx.Read3args{
		File:   *f.fh,
		Offset: nfsx.Offset3(offset),
		Count:  nfsx.Count3(len(b)),
	})

	if err != nil {
		return 0, NewFSError(err, "Couldn't read from %s", f.name)
	}
	if result.Status != nfsx.V3_OK {
		return 0, NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Failed to read from %s", f.name,
		)
	}

	read := result.Union.(nfsx.Read3resok)
	n = int(read.Count)

	offset += uint64(n)

	if useCurrentPosition {
		f.position = offset
	}

	copy(b, read.Data)

	// Must return this error to implement io.Writer
	if read.Eof {
		return n, io.EOF
	}

	return n, nil
}

// WriteSingle writes once limited to nfs.Wtmax buffer
// It returns the number of bytes written and an error, if any.
// Write does not loop due to nfs RFC specification
// https://tools.ietf.org/html/rfc1813#section-3.3.7
func (f *NfsFile) WriteSingle(b []byte) (n int, err error) {
	return f.write(b, 0, true, false)
}

// Write writes len(b) bytes to the File.
// It returns the number of bytes written and an error, if any.
// Write returns a non-nil error when n != len(b).
// Write does loop till all write is done, this is due to Go semantics
func (f *NfsFile) Write(b []byte) (n int, err error) {
	for len(b) >= 0 {
		m, err := f.write(b, 0, true, true)
		n += m
		if err != nil {
			return n, err
		}
		if m == 0 {
			return n, nil
		}
		b = b[m:]
	}
	return n, nil
}

// WriteAt writes len(b) bytes to the File starting at byte offset off.
// It returns the number of bytes written and an error, if any.
// WriteAt returns a non-nil error when n != len(b).
func (f *NfsFile) WriteAt(b []byte, offset int64) (n int, err error) {
	f.Debug("WriteAt", "len", len(b), "offset", offset)
	for len(b) >= 0 {
		m, err := f.write(b, uint64(offset), false, true)
		n += m
		if err != nil {
			return n, err
		}
		if m == 0 {
			return n, nil
		}
		b = b[m:]
		offset += int64(m)
	}
	return n, nil
}

// According to RFC 1813, a write request that has a Count that is larger than
// what the server supports (wtmax) may result in a short write.
// This is OK since our Write method has a loop.
// See: https://tools.ietf.org/html/rfc1813#section-3.3.7
func (f *NfsFile) write(b []byte, offset uint64, useCurrentPosition bool, useMaxWriteLimit bool) (n int, err error) {
	f.Debug("write", "size", len(b), "off", offset)

	if len(b) > int(MaxWriteBufferSize) {
		return 0, NewFSError(
			nil,
			"Couldn't write %d bytes into %s (max write size is %d)",
			len(b), f.name, MaxWriteBufferSize,
		)
	}

	if useCurrentPosition {
		offset = f.position
	}

	// Returned due to performance issues
	if useMaxWriteLimit && f.nfs.Limits != nil {
		if len(b) >= int(f.nfs.Limits.MaxWriteSize) {
			b = b[:f.nfs.Limits.MaxWriteSize]
		}
	}

	commitLevel := nfsx.FILE_SYNC

	result, err := f.nfs.Nfs3.Proc3Write(&nfsx.Write3args{
		File:   *f.fh,
		Offset: nfsx.Offset3(offset),
		Count:  nfsx.Count3(len(b)),
		Stable: commitLevel,
		Data:   b,
	})

	if err != nil {
		return 0, NewFSError(err, "Failed to write to %s", f.name)
	}
	if result.Status != nfsx.V3_OK {
		return 0, NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Failed to write to %s", f.name,
		)
	}

	written := result.Union.(nfsx.Write3resok)
	n = int(written.Count)
	f.Debug("write", "written", n)

	offset += uint64(n)

	if useCurrentPosition {
		f.position = offset
	}

	if written.Committed != commitLevel {
		return n, NewFSError(&nfsx.ComplianceError{
			Message: fmt.Sprintf(
				"Commit level was not as strong as requested (%v)",
				commitLevel,
			)},
			"Failed to write to file %s",
			f.name,
		)
	}

	// Returned due to performance issues
	if useMaxWriteLimit && n != len(b) {
		// This error must be returned because NfsFile needs to
		// implent io.Writer
		return n, io.ErrShortWrite
	}

	return n, nil
}

// Truncate the file to the specified size.
// This may cause creation of a sparse file in case of truncation beyond the current size.
func (f *NfsFile) Truncate(size size.Size) error {
	f.Debug("Truncate", "size", size)
	result, err := f.nfs.Nfs3.Proc3Setattr(&nfsx.Setattr3args{
		Object: *f.fh,
		NewAttributes: nfsx.Sattr3{
			Size: nfsx.SetSize3{
				SetIt: true,
				Union: nfsx.Size3(size),
			},
		},
	})
	if err != nil {
		return NewFSError(err, "Couldn't truncate %s", f.name)
	}
	if result.Status != nfsx.V3_OK {
		return NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Couldn't truncate %s", f.name,
		)
	}

	return nil
}

// GetAttr returns the attributes of the NFS file.
func (f *NfsFile) GetAttr() (*nfsx.Fattr3, error) {
	f.Debug("GetAttr")
	result, err := f.nfs.Nfs3.Proc3Getattr(&nfsx.Getattr3args{
		Object: *f.fh,
	})
	if err != nil {
		return nil, NewFSError(err, "Couldn't get attribute of %s", f.name)
	}
	if result.Status != nfsx.V3_OK {
		return nil, NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Couldn't get attribute of %s", f.name,
		)
	}

	resOk := result.Union.(nfsx.Getattr3resok)
	attr := &resOk.ObjAttributes
	f.attr = attr
	return attr, nil
}

// SetAttr sets the specified attributes on the file.
func (f *NfsFile) SetAttr(sattr *nfsx.Sattr3) error {
	f.Debug("SetAttr", "sattr", *sattr)
	result, err := f.nfs.Nfs3.Proc3Setattr(&nfsx.Setattr3args{
		Object:        *f.fh,
		NewAttributes: *sattr,
	})
	if err != nil {
		return NewFSError(err, "Couldn't set attribute of %s", f.name)
	}
	if result.Status != nfsx.V3_OK {
		return NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Couldn't set attribute of %s", f.name,
		)
	}

	// If we got this far, the server didn't return an error.  However, some
	// servers (Linux, ELFS) silently ignore nonpermitted changes instead if
	// returning V3ERR_PERM or such.  Therefore, we verify that the changes we
	// requested were actually applied.

	fattr, err := f.GetAttr()
	if err != nil {
		return NewFSError(err, "Verifying attribute for %s was not successful", f.name)
	}

	var multi error
	if sattr.Mode.SetIt && fattr.Mode != sattr.Mode.Union.(nfsx.Mode3) {
		multi = multierror.Append(multi, &nfsx.AttributeMismatchError{
			Prefix: "Mode",
			Fattr3: fattr,
			Sattr3: sattr,
		})
	}
	if sattr.Uid.SetIt && fattr.Uid != sattr.Uid.Union.(nfsx.Uid3) {
		multi = multierror.Append(multi, &nfsx.AttributeMismatchError{
			Prefix: "Uid",
			Fattr3: fattr,
			Sattr3: sattr,
		})
	}
	if sattr.Gid.SetIt && fattr.Gid != sattr.Gid.Union.(nfsx.Gid3) {
		multi = multierror.Append(multi, &nfsx.AttributeMismatchError{
			Prefix: "Gid",
			Fattr3: fattr,
			Sattr3: sattr,
		})
	}
	if sattr.Size.SetIt && fattr.Size != sattr.Size.Union.(nfsx.Size3) {
		multi = multierror.Append(multi, &nfsx.AttributeMismatchError{
			Prefix: "Size",
			Fattr3: fattr,
			Sattr3: sattr,
		})
	}
	if sattr.Atime.SetIt == nfsx.SET_TO_CLIENT_TIME && fattr.Atime != sattr.Atime.Union.(nfsx.Time3) {
		multi = multierror.Append(multi, &nfsx.AttributeMismatchError{
			Prefix: "Atime",
			Fattr3: fattr,
			Sattr3: sattr,
		})
	}
	if sattr.Mtime.SetIt == nfsx.SET_TO_CLIENT_TIME && fattr.Mtime != sattr.Mtime.Union.(nfsx.Time3) {
		multi = multierror.Append(multi, &nfsx.AttributeMismatchError{
			Prefix: "Mtime",
			Fattr3: fattr,
			Sattr3: sattr,
		})
	}
	if multi != nil {
		multi = NewFSError(multi, "Attribute of %s was not updated", f.name)
	}

	return multi
}

// ReadDir reads the directory entries.
// It creates a new goroutine that sequenctially writes entries to a channel, and returns this channel.
func (f *NfsFile) ReadDir() (<-chan *DirEntry, error) {
	f.Debug("ReadDir", "file", f)

	count := 64 * size.KiB
	if f.nfs.Limits != nil {
		count = f.nfs.Limits.PrefReadDirSize
	}

	entries := make(chan *DirEntry)
	go f.readDir(entries, count)

	return entries, nil
}

func (f *NfsFile) readDir(entries chan<- *DirEntry, count size.Size) {
	var cookie nfsx.Cookie3
	var cookieVerf nfsx.Cookieverf3

	for {
		args := nfsx.Readdir3args{
			Dir:        *f.fh,
			Cookie:     cookie,
			Cookieverf: cookieVerf,
			Count:      nfsx.Count3(count),
		}
		f.Debug("Reading dir from Nfs3...", "args", args)
		result, err := f.nfs.Nfs3.Proc3Readdir(&args)
		if err != nil {
			entries <- &DirEntry{Err: err}
			return
		}
		if result.Status != nfsx.V3_OK {
			f.Error("read dir failure", "args", args, "status", result.Status)
			entries <- &DirEntry{
				Err: &nfsx.NfsError{Status: result.Status},
			}
			return
		}

		readdir := result.Union.(nfsx.Readdir3resok)

		// FIXME: We should invalidate cookieVerf upon CREATE or REMOVE
		// See: http://tools.ietf.org/html/rfc1813#section-3.3.16
		cookieVerf = readdir.Cookieverf

		for entry := readdir.Reply.Entries; entry != nil; entry = entry.Nextentry {
			entries <- &DirEntry{
				Name:   string(entry.Name),
				Fileid: uint64(entry.Fileid),
				Cookie: uint64(entry.Cookie),
			}
			cookie = entry.Cookie
		}

		if readdir.Reply.Eof {
			close(entries)
			return
		}
	}
}

// Pathconf returns the NFS pathconf information for the file.
func (f *NfsFile) Pathconf() (pathconf *nfsx.Pathconf3resok, err error) {
	f.Debug("Pathconf")
	var result *nfsx.Pathconf3res
	result, err = f.nfs.Nfs3.Proc3Pathconf(&nfsx.Pathconf3args{
		Object: *f.fh,
	})

	// FIXME: OS X returns 200112 for the fields pathconf.NoTrunc and
	// pathconf.ChownRestricted.  While these may be valid values for
	// POSIX, they are invalid for NFS which requires those to be
	// bools.  This breaks the XDR decoder for bools... should we
	// follow Postel's Law here, and allow bools to be any number
	// where nonzero is true?  This is what Wireshark does.  Need to
	// think this through.
	if errUnmarshal, ok := err.(*xdr.UnmarshalError); ok {
		if errUnmarshal.ErrorCode == xdr.ErrBadEnumValue && errUnmarshal.Func == "DecodeBool" {
			if intValue, ok := errUnmarshal.Value.(int32); ok && intValue == 200112 {
				// Ignore OS X issue -- hardcode values sniffed from Wireshark on OS X
				f.nfs.Info("OS X broken pathconf response workaround")
				return &nfsx.Pathconf3resok{
					Linkmax:         32767,
					NameMax:         255,
					NoTrunc:         true,
					ChownRestricted: true,
					CaseInsensitive: true,
					CasePreserving:  true,
				}, nil
			}
		}
	}

	if err != nil {
		return pathconf, NewFSError(err, "Failed to read pathconf in %s", f.name)
	}

	if result.Status != nfsx.V3_OK {
		return pathconf, NewFSError(
			&nfsx.NfsError{Status: result.Status},
			"Failed to read pathconf in %s", f.name,
		)
	}

	resOk := result.Union.(nfsx.Pathconf3resok)
	pathconf = &resOk
	return pathconf, nil
}
