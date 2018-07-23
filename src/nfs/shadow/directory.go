package shadow

import (
	"os"

	log "gopkg.in/inconshreveable/log15.v2"

	"nfs"
	"nfs/sunrpc/nfsx"
)

var logger = log.New("package", "shadow")

type ShadowDirectory struct {
	Refr nfs.Directory // Reference to compare against
	Elfs nfs.Directory // ELFS

	log.Logger
}

var _ nfs.Directory = (*ShadowDirectory)(nil) // Ensure interface compliance

func NewShadowDirectory(refr, elfs nfs.Directory) *ShadowDirectory {
	return &ShadowDirectory{
		Refr:   refr,
		Elfs:   elfs,
		Logger: logger.New("shadow", "ShadowDirectory"),
	}
}

func (sd *ShadowDirectory) Name() string {
	return "<shadow_directory>"
}

func (sd *ShadowDirectory) SetLimits(limits *nfs.Limits) {
	sd.Debug("SetLimits", "limits", limits)
	sd.Refr.SetLimits(limits)
	sd.Elfs.SetLimits(limits)
}

func (sd *ShadowDirectory) Mkdir(name string, perm uint32) (dir nfs.Directory, err error) {
	sd.Debug("Mkdir", "name", name, "perm", perm)
	dirRefr, errRefr := sd.Refr.Mkdir(name, perm)
	dirElfs, errElfs := sd.Elfs.Mkdir(name, perm)

	err = checkShadow("Mkdir", errRefr, errElfs)
	if err != nil {
		return nil, err
	}

	return NewShadowDirectory(dirRefr, dirElfs), nil
}

func (sd *ShadowDirectory) LookupDir(name string) (dir nfs.Directory, err error) {
	sd.Debug("LookupDir", "name", name)
	dirRefr, errRefr := sd.Refr.LookupDir(name)
	dirElfs, errElfs := sd.Elfs.LookupDir(name)

	err = checkShadow("LookupDir", errRefr, errElfs)
	if err != nil {
		return nil, err
	}

	return NewShadowDirectory(dirRefr, dirElfs), nil
}

func (sd *ShadowDirectory) OpenFile(name string, flag int, perm os.FileMode) (file nfs.File, err error) {
	sd.Debug("OpenFile", "name", name)

	fileRefr, errRefr := sd.Refr.OpenFile(name, flag, perm)
	fileElfs, errElfs := sd.Elfs.OpenFile(name, flag, perm)

	err = checkShadow("OpenFile", errRefr, errElfs)
	if err != nil {
		return nil, err
	}

	return NewShadowFile(fileRefr, fileElfs), nil
}

func (sd *ShadowDirectory) Open(name string) (file nfs.File, err error) {
	sd.Debug("Open", "name", name)
	return sd.OpenFile(name, os.O_RDONLY, 0)
}

func (sd *ShadowDirectory) Create(name string) (file nfs.File, err error) {
	sd.Debug("Create", "name", name)
	return sd.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

func (sd *ShadowDirectory) Remove(name string) (err error) {
	sd.Debug("Remove", "name", name)
	return checkShadow("Remove", sd.Refr.Remove(name), sd.Elfs.Remove(name))
}

func (sd *ShadowDirectory) Rmdir(name string) (err error) {
	sd.Debug("Rmdir", "name", name)
	return checkShadow("Rmdir", sd.Refr.Rmdir(name), sd.Elfs.Rmdir(name))
}

func (sd *ShadowDirectory) RemoveAll() (err error) {
	sd.Debug("RemoveAll")
	return checkShadow("RemoveAll", sd.Refr.RemoveAll(), sd.Elfs.RemoveAll())
}

func (sd *ShadowDirectory) Rename(fromName string, toDir nfs.Directory, toName string) error {
	sd.Debug("Rename", "name", fromName, "toDir", toDir, "toName", toName)
	shadowToDir := toDir.(*ShadowDirectory)
	return checkShadow("Rename",
		sd.Refr.Rename(fromName, shadowToDir.Refr, toName),
		sd.Elfs.Rename(fromName, shadowToDir.Elfs, toName),
	)
}

func (sd *ShadowDirectory) Link(file nfs.File, name string) error {
	sd.Debug("Link", "name", name)
	shadowFile := file.(*ShadowFile)
	return checkShadow("Link",
		sd.Refr.Link(shadowFile.Refr, name),
		sd.Elfs.Link(shadowFile.Elfs, name),
	)
}

func (sd *ShadowDirectory) ReadDir(name string) (<-chan *nfs.DirEntry, error) {
	sd.Debug("ReadDir", "name", name)
	return shadowReadDir(
		func() (<-chan *nfs.DirEntry, error) { return sd.Refr.ReadDir(name) },
		func() (<-chan *nfs.DirEntry, error) { return sd.Elfs.ReadDir(name) },
	)
}

func (sd *ShadowDirectory) FsInfo() (*nfsx.Fsinfo3resok, error) {
	sd.Debug("FsInfo")
	infoRefr, errRefr := sd.Refr.FsInfo()
	infoElfs, errElfs := sd.Elfs.FsInfo()
	_ = infoRefr
	return infoElfs, checkShadow("FsInfo", errRefr, errElfs)
}

func (sd *ShadowDirectory) FsStat() (*nfsx.Fsstat3resok, error) {
	sd.Debug("FsStat")
	statRefr, errRefr := sd.Refr.FsStat()
	statElfs, errElfs := sd.Elfs.FsStat()
	_ = statRefr
	return statElfs, checkShadow("FsStat", errRefr, errElfs)
}

func (sd *ShadowDirectory) BeforeCreate(hook nfs.BeforeCreateHook) {
	sd.Debug("BeforeCreate")
	sd.Refr.BeforeCreate(hook)
	sd.Elfs.BeforeCreate(hook)
}

func (sd *ShadowDirectory) Limits() *nfs.Limits {
	sd.Debug("Limits")
	return nil
}
