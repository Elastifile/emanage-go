package shadow

import (
	"bytes"
	"io"

	"github.com/go-errors/errors"
	log "gopkg.in/inconshreveable/log15.v2"

	"nfs"
	"nfs/sunrpc/nfsx"
	"size"
)

type ShadowFile struct {
	Refr nfs.File // Reference to compare against
	Elfs nfs.File // ELFS

	log.Logger
}

var _ nfs.File = (*ShadowFile)(nil) // Ensure interface compliance

func NewShadowFile(refr, elfs nfs.File) *ShadowFile {
	return &ShadowFile{
		Refr:   refr,
		Elfs:   elfs,
		Logger: logger.New("shadow", "ShadowFile"),
	}
}

func (sf *ShadowFile) Read(b []byte) (n int, err error) {
	sf.Debug("Read", "len", len(b))
	return shadowRead(sf, b, func(file nfs.File, buf []byte) (n int, initialOffset int64, err error) {
		n, err = file.Read(buf)
		if err != nil {
			initialOffset = file.Tell() - int64(n)
		}
		return
	})
}

func (sf *ShadowFile) ReadAt(b []byte, offset int64) (n int, err error) {
	sf.Debug("ReadAt", "len", len(b), "offset", offset)
	return shadowRead(sf, b, func(file nfs.File, buf []byte) (n int, initialOffset int64, err error) {
		n, err = file.ReadAt(buf, offset)
		initialOffset = offset
		return
	})
}

func shadowRead(sf *ShadowFile, b []byte,
	operation func(file nfs.File, buf []byte) (n int, initialOffset int64, err error),
) (n int, err error) {
	bufRefr := make([]byte, len(b))
	bufElfs := b

	nRefr, initialOffset, errRefr := operation(sf.Refr, bufRefr)
	nElfs, _, errElfs := operation(sf.Elfs, bufElfs)
	sf.Debug("shadowRead", "nRefr", nRefr, "nElfs", nElfs)
	n = nElfs

	bufRefr = bufRefr[:nRefr]
	bufElfs = bufElfs[:nElfs]

	makeDataError := func(message string) error {
		return &ShadowDataError{
			Message:        message,
			Offset:         initialOffset,
			LengthRefrData: nRefr,
			LengthElfsData: nElfs,
			RefrData:       bufRefr,
			ElfsData:       bufElfs,
			ElfsFile:       sf.Elfs,
		}
	}

	switch {
	case nRefr != nElfs:
		return n, makeDataError("Length of data read from ELFS does not equal reference")
	case !bytes.Equal(bufRefr, bufElfs):
		return n, makeDataError("Data read from ELFS does not equal reference")
	}

	if errRefr == io.EOF && errElfs == io.EOF {
		return n, io.EOF
	}

	if errRefr == io.EOF && errElfs == nil {
		logger.Warn("Encountered possible known issue, compensating: EL-491: " +
			"ELFS doesn't set the EOF flag when reading the last byte of the file")
		return n, io.EOF
	}

	err = checkShadow("Read/ReadAt", errRefr, errElfs)
	if err != nil {
		return n, err
	}

	return n, err
}

func (sf *ShadowFile) Write(b []byte) (n int, err error) {
	sf.Debug("Write", "len", len(b))
	return shadowWrite(sf, b, func(file nfs.File) (int, error) {
		return file.Write(b)
	})
}

func (sf *ShadowFile) WriteAt(b []byte, offset int64) (n int, err error) {
	sf.Debug("WriteAt", "len", len(b), "offset", offset)
	return shadowWrite(sf, b, func(file nfs.File) (int, error) {
		return file.WriteAt(b, offset)
	})
}

func shadowWrite(sf *ShadowFile, b []byte,
	operation func(file nfs.File) (n int, err error),
) (n int, err error) {
	nRefr, errRefr := operation(sf.Refr)
	nElfs, errElfs := operation(sf.Elfs)
	n = nRefr

	err = checkShadow("Write/WriteAt", errRefr, errElfs)
	if err != nil {
		return n, err
	}

	if nRefr != nElfs {
		return n, errors.Errorf("Length of data written to ELFS (%v) does not equal reference (%v)",
			nElfs, nRefr)
	}

	return n, err
}

func (sf *ShadowFile) ReadDir() (<-chan *nfs.DirEntry, error) {
	sf.Debug("ReadDir")
	return shadowReadDir(
		sf.Refr.ReadDir,
		sf.Elfs.ReadDir,
	)
}

func (sf *ShadowFile) Truncate(size size.Size) (err error) {
	sf.Debug("Truncate", "size", size)
	return checkShadow("Truncate",
		sf.Refr.Truncate(size),
		sf.Elfs.Truncate(size),
	)
}

func (sf *ShadowFile) Attr() (attr *nfsx.Fattr3) {
	return sf.Elfs.Attr()
}

func (sf *ShadowFile) GetAttr() (attr *nfsx.Fattr3, err error) {
	sf.Debug("GetAttr")
	attrRefr, errRefr := sf.Refr.GetAttr()
	attrElfs, errElfs := sf.Elfs.GetAttr()
	err = checkShadow("GetAttr", errRefr, errElfs)
	if err != nil {
		return
	}
	return shadowAttr(attrRefr, attrElfs)
}

func (sf *ShadowFile) SetAttr(newAttr *nfsx.Sattr3) error {
	sf.Debug("SetAttr")
	return checkShadow("SetAttr",
		sf.Refr.SetAttr(newAttr),
		sf.Elfs.SetAttr(newAttr),
	)
}

func shadowAttr(attrRefr, attrElfs *nfsx.Fattr3) (attr *nfsx.Fattr3, err error) {

	attrError := func(attrName string, format string, fieldElfs, fieldRefr interface{}) (*nfsx.Fattr3, error) {
		return attrElfs, errors.Errorf("Attribute mismatch for Fattr3.%v: ELFS: "+format+", reference: "+format,
			attrName, fieldElfs, fieldRefr)
	}

	e, r := attrElfs, attrRefr

	switch {
	case e.Type != r.Type:
		return attrError("Type", "%s", e.Uid, r.Uid)
	case e.Mode&nfs.ModeMask != r.Mode&nfs.ModeMask:
		return attrError("Mode", "0%o", e.Mode, r.Mode)
	case e.Nlink != r.Nlink:
		return attrError("Nlink", "%d", e.Nlink, r.Nlink)
	case e.Uid != r.Uid:
		return attrError("Uid", "%d", e.Uid, r.Uid)
	case e.Gid != r.Gid:
		return attrError("Gid", "%d", e.Gid, r.Gid)
	case e.Size != r.Size:
		return attrError("Size", "%d", e.Size, r.Size)
	case e.Rdev != r.Rdev:
		return attrError("Rdev", "%v", e.Rdev, r.Rdev)
	}

	return attrElfs, nil
}

func (sf *ShadowFile) Close() error {
	return checkShadow("Close",
		sf.Refr.Close(),
		sf.Elfs.Close(),
	)
}

func (sf *ShadowFile) Seek(offset int64, whence int) (ret int64, err error) {
	sf.Debug("Seek", "offset", offset, "whence", whence)

	retRefr, errRefr := sf.Refr.Seek(offset, whence)
	retElfs, errElfs := sf.Elfs.Seek(offset, whence)
	ret = retRefr

	if retElfs != retRefr {
		return -1, errors.Errorf("Seek returned offset %v for shadow instead of expected %v",
			retElfs, retRefr)
	}

	return ret, checkShadow("Seek", errRefr, errElfs)
}

func (sf *ShadowFile) Tell() (offset int64) {
	return sf.Refr.Tell()
}

func (sf *ShadowFile) Pathconf() (pathconf *nfsx.Pathconf3resok, err error) {
	sf.Debug("Pathconf")
	_, errRefr := sf.Refr.Pathconf()
	pathconfElfs, errElfs := sf.Elfs.Pathconf()
	return pathconfElfs, checkShadow("Pathconf", errRefr, errElfs)
}

func (sf *ShadowFile) WriteSingle(b []byte) (n int, err error) {
	sf.Debug("WriteSingle", "len", len(b))
	return 0, nil
}

// Name added to make ShadowFile implement nfs.File (compilation error)
func (sf *ShadowFile) Name() string {
	return sf.Elfs.Name()
}
