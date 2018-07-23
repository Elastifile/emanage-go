// package nfsx (type declarations) is AUTO GENERATED from nfsx.x by sunrpc_maker

package nfsx

import "nfs/sunrpc/basic"

const MAX_ARRAY_LENGTH = 0x7fffffff

// Global constants

const (
	V2_MAXDATA                       = 8192
	V2_MAXPATHLEN                    = 1024
	V2_MAXNAMLEN                     = 255
	V2_COOKIESIZE                    = 4
	V2_FHSIZE                        = 32
	V2_MNTPATHLEN                    = 1024
	V2_MNTNAMLEN                     = 255
	NFS_PORT                         = 2049
	MOUNTD_PORT                      = 644
	V3_FHSIZE                        = 64
	V3_COOKIEVERFSIZE                = 8
	V3_CREATEVERFSIZE                = 8
	V3_WRITEVERFSIZE                 = 8
	V3_MAX_DATA                      = 4294967295
	VARIABLE_LENGTH_OPAQUE_HEADER    = 4
	XDR_SIZE_FATTR3                  = 84
	ACCESS3_READ                     = 1
	ACCESS3_LOOKUP                   = 2
	ACCESS3_MODIFY                   = 4
	ACCESS3_EXTEND                   = 8
	ACCESS3_DELETE                   = 16
	ACCESS3_EXECUTE                  = 32
	XDR_SIZE_ENTRY3_WITHOUT_NAME     = 24
	XDR_SIZE_READDIR_RES_OK_BASE     = 20
	XDR_SIZE_ENTRYPLUS3_WITHOUT_NAME = 32
	XDR_SIZE_READDIRPLUS_RES_OK_BASE = 20
	FSF3_LINK                        = 1
	FSF3_SYMLINK                     = 2
	FSF3_HOMOGENEOUS                 = 8
	FSF3_CANSETTIME                  = 16
	MAXNETOBJ_SZ                     = 1024
	LM_MAXSTRLEN                     = 1024
	LM_MAXNAMELEN                    = 1025
	SM_MAXSTRLEN                     = 1024
	SM_PRIV_SZ                       = 16
)

////////////////////////////////
// Enums

// enum stat2

type stat2 basic.Enum

const (
	_OK             stat2 = 0
	ERR_PERM        stat2 = 1
	ERR_NOENT       stat2 = 2
	ERR_IO          stat2 = 5
	ERR_NXIO        stat2 = 6
	ERR_ACCES       stat2 = 13
	ERR_EXIST       stat2 = 17
	ERR_NODEV       stat2 = 19
	ERR_NOTDIR      stat2 = 20
	ERR_ISDIR       stat2 = 21
	ERR_FBIG        stat2 = 27
	ERR_NOSPC       stat2 = 28
	ERR_ROFS        stat2 = 30
	ERR_NAMETOOLONG stat2 = 63
	ERR_NOTEMPTY    stat2 = 66
	ERR_DQUOT       stat2 = 69
	ERR_STALE       stat2 = 70
	ERR_WFLUSH      stat2 = 99
)

// enum ftype2

type ftype2 basic.Enum

const (
	NFNON  ftype2 = 0
	NFREG  ftype2 = 1
	NFDIR  ftype2 = 2
	NFBLK  ftype2 = 3
	NFCHR  ftype2 = 4
	NFLNK  ftype2 = 5
	OCK    ftype2 = 6
	NFBAD  ftype2 = 7
	NFFIFO ftype2 = 8
)

// enum stat3

type stat3 basic.Enum

const (
	V3_OK             stat3 = 0
	V3ERR_PERM        stat3 = 1
	V3ERR_NOENT       stat3 = 2
	V3ERR_IO          stat3 = 5
	V3ERR_NXIO        stat3 = 6
	V3ERR_ACCES       stat3 = 13
	V3ERR_EXIST       stat3 = 17
	V3ERR_XDEV        stat3 = 18
	V3ERR_NODEV       stat3 = 19
	V3ERR_NOTDIR      stat3 = 20
	V3ERR_ISDIR       stat3 = 21
	V3ERR_INVAL       stat3 = 22
	V3ERR_FBIG        stat3 = 27
	V3ERR_NOSPC       stat3 = 28
	V3ERR_ROFS        stat3 = 30
	V3ERR_MLINK       stat3 = 31
	V3ERR_NAMETOOLONG stat3 = 63
	V3ERR_NOTEMPTY    stat3 = 66
	V3ERR_DQUOT       stat3 = 69
	V3ERR_STALE       stat3 = 70
	V3ERR_REMOTE      stat3 = 71
	V3ERR_BADHANDLE   stat3 = 10001
	V3ERR_NOT_SYNC    stat3 = 10002
	V3ERR_BAD_COOKIE  stat3 = 10003
	V3ERR_NOTSUPP     stat3 = 10004
	V3ERR_TOOSMALL    stat3 = 10005
	V3ERR_SERVERFAULT stat3 = 10006
	V3ERR_BADTYPE     stat3 = 10007
	V3ERR_JUKEBOX     stat3 = 10008
)

// enum Ftype3

type Ftype3 basic.Enum

const (
	NF3REG  Ftype3 = 1
	NF3DIR  Ftype3 = 2
	NF3BLK  Ftype3 = 3
	NF3CHR  Ftype3 = 4
	NF3LNK  Ftype3 = 5
	NF3SOCK Ftype3 = 6
	NF3FIFO Ftype3 = 7
)

// enum fs_properties

type fsProperties basic.Enum

const (
	FSF_LINK        fsProperties = 1
	FSF_SYMLINK     fsProperties = 2
	FSF_HOMOGENEOUS fsProperties = 8
	FSF_CANSETTIME  fsProperties = 16
)

// enum time_how

type timeHow basic.Enum

const (
	DONT_CHANGE        timeHow = 0
	SET_TO_SERVER_TIME timeHow = 1
	SET_TO_CLIENT_TIME timeHow = 2
)

// enum stable_how

type stableHow basic.Enum

const (
	UNSTABLE  stableHow = 0
	DATA_SYNC stableHow = 1
	FILE_SYNC stableHow = 2
)

// enum createmode3

type createmode3 basic.Enum

const (
	UNCHECKED createmode3 = 0
	GUARDED   createmode3 = 1
	EXCLUSIVE createmode3 = 2
)

// enum mountstat3

type mountstat3 basic.Enum

const (
	MNT3_OK             mountstat3 = 0
	MNT3ERR_PERM        mountstat3 = 1
	MNT3ERR_NOENT       mountstat3 = 2
	MNT3ERR_IO          mountstat3 = 5
	MNT3ERR_ACCES       mountstat3 = 13
	MNT3ERR_NOTDIR      mountstat3 = 20
	MNT3ERR_INVAL       mountstat3 = 22
	MNT3ERR_NAMETOOLONG mountstat3 = 63
	MNT3ERR_NOTSUPP     mountstat3 = 10004
	MNT3ERR_SERVERFAULT mountstat3 = 10006
)

// enum nlm4_stats

type nlm4Stats basic.Enum

const (
	NLM4_GRANTED             nlm4Stats = 0
	NLM4_DENIED              nlm4Stats = 1
	NLM4_DENIED_NOLOCKS      nlm4Stats = 2
	NLM4_BLOCKED             nlm4Stats = 3
	NLM4_DENIED_GRACE_PERIOD nlm4Stats = 4
	NLM4_DEADLCK             nlm4Stats = 5
	NLM4_ROFS                nlm4Stats = 6
	NLM4_STALE_FH            nlm4Stats = 7
	NLM4_FBIG                nlm4Stats = 8
	NLM4_FAILED              nlm4Stats = 9
)

// enum fsh4_mode

type fsh4Mode basic.Enum

const (
	fsm_DN  fsh4Mode = 0
	fsm_DR  fsh4Mode = 1
	fsm_DW  fsh4Mode = 2
	fsm_DRW fsh4Mode = 3
)

// enum fsh4_access

type fsh4Access basic.Enum

const (
	fsa_NONE fsh4Access = 0
	fsa_R    fsh4Access = 1
	fsa_W    fsh4Access = 2
	fsa_RW   fsh4Access = 3
)

////////////////////////////////
// Structs

type Path2 string // Max length: V2_MAXPATHLEN

type Filename2 string // Max length: V2_MAXNAMLEN

type Fhandle2 [V2_FHSIZE]uint8

type Data2 []uint8 // Max length: V2_MAXDATA

type Cookie2 [V2_COOKIESIZE]uint8

// struct time2

type Time2 struct {
	Seconds  uint32
	Useconds uint32
}

// struct fattr2

type Fattr2 struct {
	Type      ftype2
	Mode      uint32
	Nlink     uint32
	Uid       uint32
	Gid       uint32
	Size      uint32
	Blocksize uint32
	Rdev      uint32
	Blocks    uint32
	Fsid      uint32
	Fileid    uint32
	Atime     Time2
	Mtime     Time2
	Ctime     Time2
}

// union fhstatus2

type Fhstatus2 struct {
	Status uint32 // Arbitrator
	Union  fhstatus2Union
}

type fhstatus2Union interface {
	isFhstatus2Union()
}

func (Fhandle2) isFhstatus2Union() {}

// struct diropargs2

type Diropargs2 struct {
	Dir  Fhandle2
	Name Filename2
}

// struct DIROP2resok

type Dirop2resok struct {
	File       Fhandle2
	Attributes Fattr2
}

// union DIROP2res

type Dirop2res struct {
	Status stat2 // Arbitrator
	Union  dirop2resUnion
}

type dirop2resUnion interface {
	isDirop2resUnion()
}

func (Dirop2resok) isDirop2resUnion() {}

// union ATTR2res

type Attr2res struct {
	Status stat2 // Arbitrator
	Union  attr2resUnion
}

type attr2resUnion interface {
	isAttr2resUnion()
}

func (Fattr2) isAttr2resUnion() {}

// struct sattr2

type Sattr2 struct {
	Mode  uint32
	Uid   uint32
	Gid   uint32
	Size  uint32
	Atime Time2
	Mtime Time2
}

// struct statinfo2

type Statinfo2 struct {
	Tsize  uint32
	Bsize  uint32
	Blocks uint32
	Bfree  uint32
	Bavail uint32
}

// union STATFS2res

type Statfs2res struct {
	Status stat2 // Arbitrator
	Union  statfs2resUnion
}

type statfs2resUnion interface {
	isStatfs2resUnion()
}

func (Statinfo2) isStatfs2resUnion() {}

// struct READDIR2args

type Readdir2args struct {
	Dir    Fhandle2
	Cookie Cookie2
	Count  uint32
}

// struct entry2

type Entry2 struct {
	Fileid    uint32
	Name      Filename2
	Cookie    Cookie2
	Nextentry *Entry2
}

// struct READDIR2resok

type Readdir2resok struct {
	Entries *Entry2
	Eof     bool
}

// union READDIR2res

type Readdir2res struct {
	Status stat2 // Arbitrator
	Union  readdir2resUnion
}

type readdir2resUnion interface {
	isReaddir2resUnion()
}

func (Readdir2resok) isReaddir2resUnion() {}

// struct SYMLINK2args

type Symlink2args struct {
	From       Diropargs2
	To         Path2
	Attributes Sattr2
}

// struct LINK2args

type Link2args struct {
	From Fhandle2
	To   Diropargs2
}

// struct RENAME2args

type Rename2args struct {
	From Diropargs2
	To   Diropargs2
}

// struct CREATE2args

type Create2args struct {
	Where      Diropargs2
	Attributes Sattr2
}

// struct WRITE2args

type Write2args struct {
	File        Fhandle2
	Beginoffset uint32
	Offset      uint32
	Totalcount  uint32
	Data        Data2
}

// struct READ2resok

type Read2resok struct {
	Attributes Fattr2
	Data       Data2
}

// union READ2res

type Read2res struct {
	Status stat2 // Arbitrator
	Union  read2resUnion
}

type read2resUnion interface {
	isRead2resUnion()
}

func (Read2resok) isRead2resUnion() {}

// struct READ2args

type Read2args struct {
	File       Fhandle2
	Offset     uint32
	Count      uint32
	Totalcount uint32
}

// union READLINK2res

type Readlink2res struct {
	Status stat2 // Arbitrator
	Union  readlink2resUnion
}

type readlink2resUnion interface {
	isReadlink2resUnion()
}

func (Path2) isReadlink2resUnion() {}

// struct SETATTR2args

type Setattr2args struct {
	File       Fhandle2
	Attributes Sattr2
}

type Uint64 uint64

type Int64 int64

type Uint32 uint32

type Int32 int32

type Filename3 string

type Path3 []uint8 // Max length: V3_MAX_DATA

type Fileid3 Uint64

type Cookie3 Uint64

type Cookieverf3 [V3_COOKIEVERFSIZE]uint8

type Createverf3 [V3_CREATEVERFSIZE]uint8

type Writeverf3 [V3_WRITEVERFSIZE]uint8

type Uid3 Uint32

type Gid3 Uint32

type Size3 Uint64

type Offset3 Uint64

type Mode3 Uint32

type Count3 Uint32

// struct specdata3

type Specdata3 struct {
	Specdata1 Uint32
	Specdata2 Uint32
}

// struct _fh3

type Fh3 struct {
	Data []uint8 // Max length: V3_FHSIZE
}

// struct time3

type Time3 struct {
	Seconds  Uint32
	Nseconds Uint32
}

// struct fattr3

type Fattr3 struct {
	Type   Ftype3
	Mode   Mode3
	Nlink  Uint32
	Uid    Uid3
	Gid    Gid3
	Size   Size3
	Used   Size3
	Rdev   Specdata3
	Fsid   Uint64
	Fileid Fileid3
	Atime  Time3
	Mtime  Time3
	Ctime  Time3
}

// union post_op_attr

type PostOpAttr struct {
	AttributesFollow bool // Arbitrator
	Union            postOpAttrUnion
}

type postOpAttrUnion interface {
	isPostOpAttrUnion()
}

func (Fattr3) isPostOpAttrUnion() {}

// struct wcc_attr

type WccAttr struct {
	Size  Size3
	Mtime Time3
	Ctime Time3
}

// union pre_op_attr

type PreOpAttr struct {
	AttributesFollow bool // Arbitrator
	Union            preOpAttrUnion
}

type preOpAttrUnion interface {
	isPreOpAttrUnion()
}

func (WccAttr) isPreOpAttrUnion() {}

// struct wcc_data

type WccData struct {
	Before PreOpAttr
	After  PostOpAttr
}

// union post_op_fh3

type PostOpFh3 struct {
	HandleFollows bool // Arbitrator
	Union         postOpFh3Union
}

type postOpFh3Union interface {
	isPostOpFh3Union()
}

func (Fh3) isPostOpFh3Union() {}

// union set_mode3

type SetMode3 struct {
	SetIt bool // Arbitrator
	Union setMode3Union
}

type setMode3Union interface {
	isSetMode3Union()
}

func (Mode3) isSetMode3Union() {}

// union set_uid3

type SetUid3 struct {
	SetIt bool // Arbitrator
	Union setUid3Union
}

type setUid3Union interface {
	isSetUid3Union()
}

func (Uid3) isSetUid3Union() {}

// union set_gid3

type SetGid3 struct {
	SetIt bool // Arbitrator
	Union setGid3Union
}

type setGid3Union interface {
	isSetGid3Union()
}

func (Gid3) isSetGid3Union() {}

// union set_size3

type SetSize3 struct {
	SetIt bool // Arbitrator
	Union setSize3Union
}

type setSize3Union interface {
	isSetSize3Union()
}

func (Size3) isSetSize3Union() {}

// union set_atime

type SetAtime struct {
	SetIt timeHow // Arbitrator
	Union setAtimeUnion
}

type setAtimeUnion interface {
	isSetAtimeUnion()
}

func (Time3) isSetAtimeUnion() {}

// union set_mtime

type SetMtime struct {
	SetIt timeHow // Arbitrator
	Union setMtimeUnion
}

type setMtimeUnion interface {
	isSetMtimeUnion()
}

func (Time3) isSetMtimeUnion() {}

// struct sattr3

type Sattr3 struct {
	Mode  SetMode3
	Uid   SetUid3
	Gid   SetGid3
	Size  SetSize3
	Atime SetAtime
	Mtime SetMtime
}

// struct diropargs3

type Diropargs3 struct {
	Dir  Fh3
	Name Filename3
}

// struct GETATTR3args

type Getattr3args struct {
	Object Fh3
}

// struct GETATTR3resok

type Getattr3resok struct {
	ObjAttributes Fattr3
}

// union GETATTR3res

type Getattr3res struct {
	Status stat3 // Arbitrator
	Union  getattr3resUnion
}

type getattr3resUnion interface {
	isGetattr3resUnion()
}

func (Getattr3resok) isGetattr3resUnion() {}

// union sattrguard3

type Sattrguard3 struct {
	Check bool // Arbitrator
	Union sattrguard3Union
}

type sattrguard3Union interface {
	isSattrguard3Union()
}

func (Time3) isSattrguard3Union() {}

// struct SETATTR3args

type Setattr3args struct {
	Object        Fh3
	NewAttributes Sattr3
	Guard         Sattrguard3
}

// struct SETATTR3resok

type Setattr3resok struct {
	ObjWcc WccData
}

// struct SETATTR3resfail

type Setattr3resfail struct {
	ObjWcc WccData
}

// union SETATTR3res

type Setattr3res struct {
	Status stat3 // Arbitrator
	Union  setattr3resUnion
}

type setattr3resUnion interface {
	isSetattr3resUnion()
}

func (Setattr3resok) isSetattr3resUnion() {}

// struct LOOKUP3args

type Lookup3args struct {
	What Diropargs3
}

// struct LOOKUP3resok

type Lookup3resok struct {
	Object        Fh3
	ObjAttributes PostOpAttr
	DirAttributes PostOpAttr
}

// struct LOOKUP3resfail

type Lookup3resfail struct {
	DirAttributes PostOpAttr
}

// union LOOKUP3res

type Lookup3res struct {
	Status stat3 // Arbitrator
	Union  lookup3resUnion
}

type lookup3resUnion interface {
	isLookup3resUnion()
}

func (Lookup3resok) isLookup3resUnion() {}

// struct ACCESS3args

type Access3args struct {
	Object Fh3
	Access Uint32
}

// struct ACCESS3resok

type Access3resok struct {
	ObjAttributes PostOpAttr
	Access        Uint32
}

// struct ACCESS3resfail

type Access3resfail struct {
	ObjAttributes PostOpAttr
}

// union ACCESS3res

type Access3res struct {
	Status stat3 // Arbitrator
	Union  access3resUnion
}

type access3resUnion interface {
	isAccess3resUnion()
}

func (Access3resok) isAccess3resUnion() {}

// struct READLINK3args

type Readlink3args struct {
	Symlink Fh3
}

// struct READLINK3resok

type Readlink3resok struct {
	SymlinkAttributes PostOpAttr
	Data              Path3
}

// struct READLINK3resfail

type Readlink3resfail struct {
	SymlinkAttributes PostOpAttr
}

// union READLINK3res

type Readlink3res struct {
	Status stat3 // Arbitrator
	Union  readlink3resUnion
}

type readlink3resUnion interface {
	isReadlink3resUnion()
}

func (Readlink3resok) isReadlink3resUnion() {}

// struct READ3args

type Read3args struct {
	File   Fh3
	Offset Offset3
	Count  Count3
}

// struct READ3resok

type Read3resok struct {
	FileAttributes PostOpAttr
	Count          Count3
	Eof            bool
	Data           []uint8 // Max length: V3_MAX_DATA
}

// struct READ3resfail

type Read3resfail struct {
	FileAttributes PostOpAttr
}

// union READ3res

type Read3res struct {
	Status stat3 // Arbitrator
	Union  read3resUnion
}

type read3resUnion interface {
	isRead3resUnion()
}

func (Read3resok) isRead3resUnion() {}

// struct WRITE3args

type Write3args struct {
	File   Fh3
	Offset Offset3
	Count  Count3
	Stable stableHow
	Data   []uint8 // Max length: V3_MAX_DATA
}

// struct WRITE3resok

type Write3resok struct {
	FileWcc   WccData
	Count     Count3
	Committed stableHow
	Verf      Writeverf3
}

// struct WRITE3resfail

type Write3resfail struct {
	FileWcc WccData
}

// union WRITE3res

type Write3res struct {
	Status stat3 // Arbitrator
	Union  write3resUnion
}

type write3resUnion interface {
	isWrite3resUnion()
}

func (Write3resok) isWrite3resUnion() {}

// union createhow3

type Createhow3 struct {
	Mode  createmode3 // Arbitrator
	Union createhow3Union
}

type createhow3Union interface {
	isCreatehow3Union()
}

func (Sattr3) isCreatehow3Union() {}

func (Createverf3) isCreatehow3Union() {}

// struct CREATE3args

type Create3args struct {
	Where Diropargs3
	How   Createhow3
}

// struct CREATE3resok

type Create3resok struct {
	Obj           PostOpFh3
	ObjAttributes PostOpAttr
	DirWcc        WccData
}

// struct CREATE3resfail

type Create3resfail struct {
	DirWcc WccData
}

// union CREATE3res

type Create3res struct {
	Status stat3 // Arbitrator
	Union  create3resUnion
}

type create3resUnion interface {
	isCreate3resUnion()
}

func (Create3resok) isCreate3resUnion() {}

// struct MKDIR3args

type Mkdir3args struct {
	Where      Diropargs3
	Attributes Sattr3
}

// struct MKDIR3resok

type Mkdir3resok struct {
	Obj           PostOpFh3
	ObjAttributes PostOpAttr
	DirWcc        WccData
}

// struct MKDIR3resfail

type Mkdir3resfail struct {
	DirWcc WccData
}

// union MKDIR3res

type Mkdir3res struct {
	Status stat3 // Arbitrator
	Union  mkdir3resUnion
}

type mkdir3resUnion interface {
	isMkdir3resUnion()
}

func (Mkdir3resok) isMkdir3resUnion() {}

// struct symlinkdata3

type Symlinkdata3 struct {
	SymlinkAttributes Sattr3
	SymlinkData       Path3
}

// struct SYMLINK3args

type Symlink3args struct {
	Where   Diropargs3
	Symlink Symlinkdata3
}

// struct SYMLINK3resok

type Symlink3resok struct {
	Obj           PostOpFh3
	ObjAttributes PostOpAttr
	DirWcc        WccData
}

// struct SYMLINK3resfail

type Symlink3resfail struct {
	DirWcc WccData
}

// union SYMLINK3res

type Symlink3res struct {
	Status stat3 // Arbitrator
	Union  symlink3resUnion
}

type symlink3resUnion interface {
	isSymlink3resUnion()
}

func (Symlink3resok) isSymlink3resUnion() {}

// struct devicedata3

type Devicedata3 struct {
	DevAttributes Sattr3
	Spec          Specdata3
}

// union mknoddata3

type Mknoddata3 struct {
	Type  Ftype3 // Arbitrator
	Union mknoddata3Union
}

type mknoddata3Union interface {
	isMknoddata3Union()
}

func (Devicedata3) isMknoddata3Union() {}

func (Sattr3) isMknoddata3Union() {}

// struct MKNOD3args

type Mknod3args struct {
	Where Diropargs3
	What  Mknoddata3
}

// struct MKNOD3resok

type Mknod3resok struct {
	Obj           PostOpFh3
	ObjAttributes PostOpAttr
	DirWcc        WccData
}

// struct MKNOD3resfail

type Mknod3resfail struct {
	DirWcc WccData
}

// union MKNOD3res

type Mknod3res struct {
	Status stat3 // Arbitrator
	Union  mknod3resUnion
}

type mknod3resUnion interface {
	isMknod3resUnion()
}

func (Mknod3resok) isMknod3resUnion() {}

// struct REMOVE3args

type Remove3args struct {
	Object Diropargs3
}

// struct REMOVE3resok

type Remove3resok struct {
	DirWcc WccData
}

// struct REMOVE3resfail

type Remove3resfail struct {
	DirWcc WccData
}

// union REMOVE3res

type Remove3res struct {
	Status stat3 // Arbitrator
	Union  remove3resUnion
}

type remove3resUnion interface {
	isRemove3resUnion()
}

func (Remove3resok) isRemove3resUnion() {}

// struct RMDIR3args

type Rmdir3args struct {
	Object Diropargs3
}

// struct RMDIR3resok

type Rmdir3resok struct {
	DirWcc WccData
}

// struct RMDIR3resfail

type Rmdir3resfail struct {
	DirWcc WccData
}

// union RMDIR3res

type Rmdir3res struct {
	Status stat3 // Arbitrator
	Union  rmdir3resUnion
}

type rmdir3resUnion interface {
	isRmdir3resUnion()
}

func (Rmdir3resok) isRmdir3resUnion() {}

// struct RENAME3args

type Rename3args struct {
	From Diropargs3
	To   Diropargs3
}

// struct RENAME3resok

type Rename3resok struct {
	FromdirWcc WccData
	TodirWcc   WccData
}

// struct RENAME3resfail

type Rename3resfail struct {
	FromdirWcc WccData
	TodirWcc   WccData
}

// union RENAME3res

type Rename3res struct {
	Status stat3 // Arbitrator
	Union  rename3resUnion
}

type rename3resUnion interface {
	isRename3resUnion()
}

func (Rename3resok) isRename3resUnion() {}

// struct LINK3args

type Link3args struct {
	File Fh3
	Link Diropargs3
}

// struct LINK3resok

type Link3resok struct {
	FileAttributes PostOpAttr
	LinkdirWcc     WccData
}

// struct LINK3resfail

type Link3resfail struct {
	FileAttributes PostOpAttr
	LinkdirWcc     WccData
}

// union LINK3res

type Link3res struct {
	Status stat3 // Arbitrator
	Union  link3resUnion
}

type link3resUnion interface {
	isLink3resUnion()
}

func (Link3resok) isLink3resUnion() {}

// struct READDIR3args

type Readdir3args struct {
	Dir        Fh3
	Cookie     Cookie3
	Cookieverf Cookieverf3
	Count      Count3
}

// struct entry3

type Entry3 struct {
	Fileid    Fileid3
	Name      Filename3
	Cookie    Cookie3
	Nextentry *Entry3
}

// struct dirlist3

type Dirlist3 struct {
	Entries *Entry3
	Eof     bool
}

// struct READDIR3resok

type Readdir3resok struct {
	DirAttributes PostOpAttr
	Cookieverf    Cookieverf3
	Reply         Dirlist3
}

// struct READDIR3resfail

type Readdir3resfail struct {
	DirAttributes PostOpAttr
}

// union READDIR3res

type Readdir3res struct {
	Status stat3 // Arbitrator
	Union  readdir3resUnion
}

type readdir3resUnion interface {
	isReaddir3resUnion()
}

func (Readdir3resok) isReaddir3resUnion() {}

// struct READDIRPLUS3args

type Readdirplus3args struct {
	Dir        Fh3
	Cookie     Cookie3
	Cookieverf Cookieverf3
	Dircount   Count3
	Maxcount   Count3
}

// struct entryplus3

type Entryplus3 struct {
	Fileid         Fileid3
	Name           Filename3
	Cookie         Cookie3
	NameAttributes PostOpAttr
	NameHandle     PostOpFh3
	Nextentry      *Entryplus3
}

// struct dirlistplus3

type Dirlistplus3 struct {
	Entries *Entryplus3
	Eof     bool
}

// struct READDIRPLUS3resok

type Readdirplus3resok struct {
	DirAttributes PostOpAttr
	Cookieverf    Cookieverf3
	Reply         Dirlistplus3
}

// struct READDIRPLUS3resfail

type Readdirplus3resfail struct {
	DirAttributes PostOpAttr
}

// union READDIRPLUS3res

type Readdirplus3res struct {
	Status stat3 // Arbitrator
	Union  readdirplus3resUnion
}

type readdirplus3resUnion interface {
	isReaddirplus3resUnion()
}

func (Readdirplus3resok) isReaddirplus3resUnion() {}

// struct FSSTAT3args

type Fsstat3args struct {
	Fsroot Fh3
}

// struct FSSTAT3resok

type Fsstat3resok struct {
	ObjAttributes PostOpAttr
	Tbytes        Size3
	Fbytes        Size3
	Abytes        Size3
	Tfiles        Size3
	Ffiles        Size3
	Afiles        Size3
	Invarsec      Uint32
}

// struct FSSTAT3resfail

type Fsstat3resfail struct {
	ObjAttributes PostOpAttr
}

// union FSSTAT3res

type Fsstat3res struct {
	Status stat3 // Arbitrator
	Union  fsstat3resUnion
}

type fsstat3resUnion interface {
	isFsstat3resUnion()
}

func (Fsstat3resok) isFsstat3resUnion() {}

// struct FSINFO3args

type Fsinfo3args struct {
	Fsroot Fh3
}

// struct FSINFO3resok

type Fsinfo3resok struct {
	ObjAttributes PostOpAttr
	Rtmax         Uint32
	Rtpref        Uint32
	Rtmult        Uint32
	Wtmax         Uint32
	Wtpref        Uint32
	Wtmult        Uint32
	Dtpref        Uint32
	Maxfilesize   Size3
	TimeDelta     Time3
	Properties    Uint32
}

// struct FSINFO3resfail

type Fsinfo3resfail struct {
	ObjAttributes PostOpAttr
}

// union FSINFO3res

type Fsinfo3res struct {
	Status stat3 // Arbitrator
	Union  fsinfo3resUnion
}

type fsinfo3resUnion interface {
	isFsinfo3resUnion()
}

func (Fsinfo3resok) isFsinfo3resUnion() {}

// struct PATHCONF3args

type Pathconf3args struct {
	Object Fh3
}

// struct PATHCONF3resok

type Pathconf3resok struct {
	ObjAttributes   PostOpAttr
	Linkmax         Uint32
	NameMax         Uint32
	NoTrunc         bool
	ChownRestricted bool
	CaseInsensitive bool
	CasePreserving  bool
}

// struct PATHCONF3resfail

type Pathconf3resfail struct {
	ObjAttributes PostOpAttr
}

// union PATHCONF3res

type Pathconf3res struct {
	Status stat3 // Arbitrator
	Union  pathconf3resUnion
}

type pathconf3resUnion interface {
	isPathconf3resUnion()
}

func (Pathconf3resok) isPathconf3resUnion() {}

// struct COMMIT3args

type Commit3args struct {
	File   Fh3
	Offset Offset3
	Count  Count3
}

// struct COMMIT3resok

type Commit3resok struct {
	FileWcc WccData
	Verf    Writeverf3
}

// struct COMMIT3resfail

type Commit3resfail struct {
	FileWcc WccData
}

// union COMMIT3res

type Commit3res struct {
	Status stat3 // Arbitrator
	Union  commit3resUnion
}

type commit3resUnion interface {
	isCommit3resUnion()
}

func (Commit3resok) isCommit3resUnion() {}

type Dirpath string // Max length: V2_MNTPATHLEN

type Mntname string // Max length: V2_MNTNAMLEN

// struct groupnode

type Groupnode struct {
	GrName Mntname
	GrNext *Groupnode
}

// struct exportnode

type Exportnode struct {
	ExDir    Dirpath
	ExGroups *Groupnode
	ExNext   *Exportnode
}

// struct mountbody

type Mountbody struct {
	MlHostname  Mntname
	MlDirectory Dirpath
	MlNext      *Mountbody
}

// struct mountlist_first

type MountlistFirst struct {
	First *Mountbody
}

// struct exports_first

type ExportsFirst struct {
	First *Exportnode
}

// struct mountres3_ok

type Mountres3Ok struct {
	Fhandle     Fh3
	AuthFlavors []int32
}

// union mountres3

type Mountres3 struct {
	FhsStatus mountstat3 // Arbitrator
	Union     mountres3Union
}

type mountres3Union interface {
	isMountres3Union()
}

func (Mountres3Ok) isMountres3Union() {}

type Netobj [MAXNETOBJ_SZ]uint8

// struct nlm4_stat

type Nlm4Stat struct {
	Stat nlm4Stats
}

// struct nlm4_res

type Nlm4Res struct {
	Cookie Netobj
	Stat   Nlm4Stat
}

// struct nlm4_holder

type Nlm4Holder struct {
	Exclusive bool
	Svid      Int32
	Oh        Netobj
	LOffset   Uint64
	LLen      Uint64
}

// union nlm4_testrply

type Nlm4Testrply struct {
	Stat  nlm4Stats // Arbitrator
	Union nlm4TestrplyUnion
}

type nlm4TestrplyUnion interface {
	isNlm4TestrplyUnion()
}

func (Nlm4Holder) isNlm4TestrplyUnion() {}

// struct nlm4_testres

type Nlm4Testres struct {
	Cookie   Netobj
	TestStat Nlm4Testrply
}

// struct nlm4_lock

type Nlm4Lock struct {
	CallerName string // Max length: LM_MAXSTRLEN
	Fh         Netobj
	Oh         Netobj
	Svid       Int32
	LOffset    Uint64
	LLen       Uint64
}

// struct nlm4_lockargs

type Nlm4Lockargs struct {
	Cookie    Netobj
	Block     bool
	Exclusive bool
	Alock     Nlm4Lock
	Reclaim   bool
	State     Int32
}

// struct nlm4_cancargs

type Nlm4Cancargs struct {
	Cookie    Netobj
	Block     bool
	Exclusive bool
	Alock     Nlm4Lock
}

// struct nlm4_testargs

type Nlm4Testargs struct {
	Cookie    Netobj
	Exclusive bool
	Alock     Nlm4Lock
}

// struct nlm4_unlockargs

type Nlm4Unlockargs struct {
	Cookie Netobj
	Alock  Nlm4Lock
}

// struct nlm4_share

type Nlm4Share struct {
	CallerName string // Max length: LM_MAXSTRLEN
	Fh         Netobj
	Oh         Netobj
	Mode       fsh4Mode
	Access     fsh4Access
}

// struct nlm4_shareargs

type Nlm4Shareargs struct {
	Cookie  Netobj
	Share   Nlm4Share
	Reclaim bool
}

// struct nlm4_shareres

type Nlm4Shareres struct {
	Cookie   Netobj
	Stat     nlm4Stats
	Sequence Int32
}

// struct nlm4_notify

type Nlm4Notify struct {
	Name  string // Max length: LM_MAXNAMELEN
	State Int64
}

// struct nlm4_sm_notifyargs

type Nlm4SmNotifyargs struct {
	Name  string // Max length: SM_MAXSTRLEN
	State Int32
	Priv  [SM_PRIV_SZ]uint8
}
