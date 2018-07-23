// package nfsx (implementation) is AUTO GENERATED from nfsx.x by sunrpc_maker
package nfsx

import (
	"fmt"

	xdr "github.com/davecgh/go-xdr/xdr2"
)
import (
	"nfs/sunrpc/basic"
	"nfs/sunrpc/client"
	"nfs/sunrpc/rpc2"
)

// enum stat2

func (value stat2) FieldCheck() error {
	switch value {
	case _OK:
	case ERR_PERM:
	case ERR_NOENT:
	case ERR_IO:
	case ERR_NXIO:
	case ERR_ACCES:
	case ERR_EXIST:
	case ERR_NODEV:
	case ERR_NOTDIR:
	case ERR_ISDIR:
	case ERR_FBIG:
	case ERR_NOSPC:
	case ERR_ROFS:
	case ERR_NAMETOOLONG:
	case ERR_NOTEMPTY:
	case ERR_DQUOT:
	case ERR_STALE:
	case ERR_WFLUSH:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value stat2) String() string {
	switch value {
	case _OK:
		return "_OK"
	case ERR_PERM:
		return "ERR_PERM"
	case ERR_NOENT:
		return "ERR_NOENT"
	case ERR_IO:
		return "ERR_IO"
	case ERR_NXIO:
		return "ERR_NXIO"
	case ERR_ACCES:
		return "ERR_ACCES"
	case ERR_EXIST:
		return "ERR_EXIST"
	case ERR_NODEV:
		return "ERR_NODEV"
	case ERR_NOTDIR:
		return "ERR_NOTDIR"
	case ERR_ISDIR:
		return "ERR_ISDIR"
	case ERR_FBIG:
		return "ERR_FBIG"
	case ERR_NOSPC:
		return "ERR_NOSPC"
	case ERR_ROFS:
		return "ERR_ROFS"
	case ERR_NAMETOOLONG:
		return "ERR_NAMETOOLONG"
	case ERR_NOTEMPTY:
		return "ERR_NOTEMPTY"
	case ERR_DQUOT:
		return "ERR_DQUOT"
	case ERR_STALE:
		return "ERR_STALE"
	case ERR_WFLUSH:
		return "ERR_WFLUSH"
	}
	return fmt.Sprintf("stat2(%d)", value)
}

func (value *stat2) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("stat2", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *stat2) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = stat2(intValue)
	basic.SniffDecode("stat2", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum ftype2

func (value ftype2) FieldCheck() error {
	switch value {
	case NFNON:
	case NFREG:
	case NFDIR:
	case NFBLK:
	case NFCHR:
	case NFLNK:
	case OCK:
	case NFBAD:
	case NFFIFO:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value ftype2) String() string {
	switch value {
	case NFNON:
		return "NFNON"
	case NFREG:
		return "NFREG"
	case NFDIR:
		return "NFDIR"
	case NFBLK:
		return "NFBLK"
	case NFCHR:
		return "NFCHR"
	case NFLNK:
		return "NFLNK"
	case OCK:
		return "OCK"
	case NFBAD:
		return "NFBAD"
	case NFFIFO:
		return "NFFIFO"
	}
	return fmt.Sprintf("ftype2(%d)", value)
}

func (value *ftype2) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("ftype2", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *ftype2) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = ftype2(intValue)
	basic.SniffDecode("ftype2", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum stat3

func (value stat3) FieldCheck() error {
	switch value {
	case V3_OK:
	case V3ERR_PERM:
	case V3ERR_NOENT:
	case V3ERR_IO:
	case V3ERR_NXIO:
	case V3ERR_ACCES:
	case V3ERR_EXIST:
	case V3ERR_XDEV:
	case V3ERR_NODEV:
	case V3ERR_NOTDIR:
	case V3ERR_ISDIR:
	case V3ERR_INVAL:
	case V3ERR_FBIG:
	case V3ERR_NOSPC:
	case V3ERR_ROFS:
	case V3ERR_MLINK:
	case V3ERR_NAMETOOLONG:
	case V3ERR_NOTEMPTY:
	case V3ERR_DQUOT:
	case V3ERR_STALE:
	case V3ERR_REMOTE:
	case V3ERR_BADHANDLE:
	case V3ERR_NOT_SYNC:
	case V3ERR_BAD_COOKIE:
	case V3ERR_NOTSUPP:
	case V3ERR_TOOSMALL:
	case V3ERR_SERVERFAULT:
	case V3ERR_BADTYPE:
	case V3ERR_JUKEBOX:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value stat3) String() string {
	switch value {
	case V3_OK:
		return "V3_OK"
	case V3ERR_PERM:
		return "V3ERR_PERM"
	case V3ERR_NOENT:
		return "V3ERR_NOENT"
	case V3ERR_IO:
		return "V3ERR_IO"
	case V3ERR_NXIO:
		return "V3ERR_NXIO"
	case V3ERR_ACCES:
		return "V3ERR_ACCES"
	case V3ERR_EXIST:
		return "V3ERR_EXIST"
	case V3ERR_XDEV:
		return "V3ERR_XDEV"
	case V3ERR_NODEV:
		return "V3ERR_NODEV"
	case V3ERR_NOTDIR:
		return "V3ERR_NOTDIR"
	case V3ERR_ISDIR:
		return "V3ERR_ISDIR"
	case V3ERR_INVAL:
		return "V3ERR_INVAL"
	case V3ERR_FBIG:
		return "V3ERR_FBIG"
	case V3ERR_NOSPC:
		return "V3ERR_NOSPC"
	case V3ERR_ROFS:
		return "V3ERR_ROFS"
	case V3ERR_MLINK:
		return "V3ERR_MLINK"
	case V3ERR_NAMETOOLONG:
		return "V3ERR_NAMETOOLONG"
	case V3ERR_NOTEMPTY:
		return "V3ERR_NOTEMPTY"
	case V3ERR_DQUOT:
		return "V3ERR_DQUOT"
	case V3ERR_STALE:
		return "V3ERR_STALE"
	case V3ERR_REMOTE:
		return "V3ERR_REMOTE"
	case V3ERR_BADHANDLE:
		return "V3ERR_BADHANDLE"
	case V3ERR_NOT_SYNC:
		return "V3ERR_NOT_SYNC"
	case V3ERR_BAD_COOKIE:
		return "V3ERR_BAD_COOKIE"
	case V3ERR_NOTSUPP:
		return "V3ERR_NOTSUPP"
	case V3ERR_TOOSMALL:
		return "V3ERR_TOOSMALL"
	case V3ERR_SERVERFAULT:
		return "V3ERR_SERVERFAULT"
	case V3ERR_BADTYPE:
		return "V3ERR_BADTYPE"
	case V3ERR_JUKEBOX:
		return "V3ERR_JUKEBOX"
	}
	return fmt.Sprintf("stat3(%d)", value)
}

func (value *stat3) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("stat3", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *stat3) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = stat3(intValue)
	basic.SniffDecode("stat3", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum Ftype3

func (value Ftype3) FieldCheck() error {
	switch value {
	case NF3REG:
	case NF3DIR:
	case NF3BLK:
	case NF3CHR:
	case NF3LNK:
	case NF3SOCK:
	case NF3FIFO:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value Ftype3) String() string {
	switch value {
	case NF3REG:
		return "NF3REG"
	case NF3DIR:
		return "NF3DIR"
	case NF3BLK:
		return "NF3BLK"
	case NF3CHR:
		return "NF3CHR"
	case NF3LNK:
		return "NF3LNK"
	case NF3SOCK:
		return "NF3SOCK"
	case NF3FIFO:
		return "NF3FIFO"
	}
	return fmt.Sprintf("Ftype3(%d)", value)
}

func Ftype3Enum(name string) Ftype3 {
	switch name {
	case "NF3REG":
		return NF3REG
	case "NF3DIR":
		return NF3DIR
	case "NF3BLK":
		return NF3BLK
	case "NF3CHR":
		return NF3CHR
	case "NF3LNK":
		return NF3LNK
	case "NF3SOCK":
		return NF3SOCK
	case "NF3FIFO":
		return NF3FIFO
	default:
		return 0
	}
}

func (value *Ftype3) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("Ftype3", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *Ftype3) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = Ftype3(intValue)
	basic.SniffDecode("Ftype3", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum fs_properties

func (value fsProperties) FieldCheck() error {
	switch value {
	case FSF_LINK:
	case FSF_SYMLINK:
	case FSF_HOMOGENEOUS:
	case FSF_CANSETTIME:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value fsProperties) String() string {
	switch value {
	case FSF_LINK:
		return "FSF_LINK"
	case FSF_SYMLINK:
		return "FSF_SYMLINK"
	case FSF_HOMOGENEOUS:
		return "FSF_HOMOGENEOUS"
	case FSF_CANSETTIME:
		return "FSF_CANSETTIME"
	}
	return fmt.Sprintf("fsProperties(%d)", value)
}

func (value *fsProperties) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("fsProperties", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *fsProperties) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = fsProperties(intValue)
	basic.SniffDecode("fsProperties", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum time_how

func (value timeHow) FieldCheck() error {
	switch value {
	case DONT_CHANGE:
	case SET_TO_SERVER_TIME:
	case SET_TO_CLIENT_TIME:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value timeHow) String() string {
	switch value {
	case DONT_CHANGE:
		return "DONT_CHANGE"
	case SET_TO_SERVER_TIME:
		return "SET_TO_SERVER_TIME"
	case SET_TO_CLIENT_TIME:
		return "SET_TO_CLIENT_TIME"
	}
	return fmt.Sprintf("timeHow(%d)", value)
}

func (value *timeHow) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("timeHow", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *timeHow) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = timeHow(intValue)
	basic.SniffDecode("timeHow", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum stable_how

func (value stableHow) FieldCheck() error {
	switch value {
	case UNSTABLE:
	case DATA_SYNC:
	case FILE_SYNC:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value stableHow) String() string {
	switch value {
	case UNSTABLE:
		return "UNSTABLE"
	case DATA_SYNC:
		return "DATA_SYNC"
	case FILE_SYNC:
		return "FILE_SYNC"
	}
	return fmt.Sprintf("stableHow(%d)", value)
}

func (value *stableHow) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("stableHow", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *stableHow) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = stableHow(intValue)
	basic.SniffDecode("stableHow", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum createmode3

func (value createmode3) FieldCheck() error {
	switch value {
	case UNCHECKED:
	case GUARDED:
	case EXCLUSIVE:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value createmode3) String() string {
	switch value {
	case UNCHECKED:
		return "UNCHECKED"
	case GUARDED:
		return "GUARDED"
	case EXCLUSIVE:
		return "EXCLUSIVE"
	}
	return fmt.Sprintf("createmode3(%d)", value)
}

func (value *createmode3) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("createmode3", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *createmode3) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = createmode3(intValue)
	basic.SniffDecode("createmode3", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum mountstat3

func (value mountstat3) FieldCheck() error {
	switch value {
	case MNT3_OK:
	case MNT3ERR_PERM:
	case MNT3ERR_NOENT:
	case MNT3ERR_IO:
	case MNT3ERR_ACCES:
	case MNT3ERR_NOTDIR:
	case MNT3ERR_INVAL:
	case MNT3ERR_NAMETOOLONG:
	case MNT3ERR_NOTSUPP:
	case MNT3ERR_SERVERFAULT:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value mountstat3) String() string {
	switch value {
	case MNT3_OK:
		return "MNT3_OK"
	case MNT3ERR_PERM:
		return "MNT3ERR_PERM"
	case MNT3ERR_NOENT:
		return "MNT3ERR_NOENT"
	case MNT3ERR_IO:
		return "MNT3ERR_IO"
	case MNT3ERR_ACCES:
		return "MNT3ERR_ACCES"
	case MNT3ERR_NOTDIR:
		return "MNT3ERR_NOTDIR"
	case MNT3ERR_INVAL:
		return "MNT3ERR_INVAL"
	case MNT3ERR_NAMETOOLONG:
		return "MNT3ERR_NAMETOOLONG"
	case MNT3ERR_NOTSUPP:
		return "MNT3ERR_NOTSUPP"
	case MNT3ERR_SERVERFAULT:
		return "MNT3ERR_SERVERFAULT"
	}
	return fmt.Sprintf("mountstat3(%d)", value)
}

func (value *mountstat3) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("mountstat3", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *mountstat3) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = mountstat3(intValue)
	basic.SniffDecode("mountstat3", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum nlm4_stats

func (value nlm4Stats) FieldCheck() error {
	switch value {
	case NLM4_GRANTED:
	case NLM4_DENIED:
	case NLM4_DENIED_NOLOCKS:
	case NLM4_BLOCKED:
	case NLM4_DENIED_GRACE_PERIOD:
	case NLM4_DEADLCK:
	case NLM4_ROFS:
	case NLM4_STALE_FH:
	case NLM4_FBIG:
	case NLM4_FAILED:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value nlm4Stats) String() string {
	switch value {
	case NLM4_GRANTED:
		return "NLM4_GRANTED"
	case NLM4_DENIED:
		return "NLM4_DENIED"
	case NLM4_DENIED_NOLOCKS:
		return "NLM4_DENIED_NOLOCKS"
	case NLM4_BLOCKED:
		return "NLM4_BLOCKED"
	case NLM4_DENIED_GRACE_PERIOD:
		return "NLM4_DENIED_GRACE_PERIOD"
	case NLM4_DEADLCK:
		return "NLM4_DEADLCK"
	case NLM4_ROFS:
		return "NLM4_ROFS"
	case NLM4_STALE_FH:
		return "NLM4_STALE_FH"
	case NLM4_FBIG:
		return "NLM4_FBIG"
	case NLM4_FAILED:
		return "NLM4_FAILED"
	}
	return fmt.Sprintf("nlm4Stats(%d)", value)
}

func (value *nlm4Stats) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("nlm4Stats", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *nlm4Stats) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = nlm4Stats(intValue)
	basic.SniffDecode("nlm4Stats", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum fsh4_mode

func (value fsh4Mode) FieldCheck() error {
	switch value {
	case fsm_DN:
	case fsm_DR:
	case fsm_DW:
	case fsm_DRW:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value fsh4Mode) String() string {
	switch value {
	case fsm_DN:
		return "fsm_DN"
	case fsm_DR:
		return "fsm_DR"
	case fsm_DW:
		return "fsm_DW"
	case fsm_DRW:
		return "fsm_DRW"
	}
	return fmt.Sprintf("fsh4Mode(%d)", value)
}

func (value *fsh4Mode) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("fsh4Mode", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *fsh4Mode) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = fsh4Mode(intValue)
	basic.SniffDecode("fsh4Mode", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum fsh4_access

func (value fsh4Access) FieldCheck() error {
	switch value {
	case fsa_NONE:
	case fsa_R:
	case fsa_W:
	case fsa_RW:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value fsh4Access) String() string {
	switch value {
	case fsa_NONE:
		return "fsa_NONE"
	case fsa_R:
		return "fsa_R"
	case fsa_W:
		return "fsa_W"
	case fsa_RW:
		return "fsa_RW"
	}
	return fmt.Sprintf("fsh4Access(%d)", value)
}

func (value *fsh4Access) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("fsh4Access", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *fsh4Access) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = fsh4Access(intValue)
	basic.SniffDecode("fsh4Access", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// type path2

func (t *Path2) Encode(e *xdr.Encoder) (err error) {
	value := string(*t)

	// value: string<V2_MAXPATHLEN>

	basic.SniffEncode("value", value)

	if len(value) > V2_MAXPATHLEN {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Path2) Decode(d *xdr.Decoder) (err error) {
	var value string

	// value: string<V2_MAXPATHLEN>

	*(*string)(&value), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("value", value)

	if len(value) > V2_MAXPATHLEN {
		err = basic.ErrArrayTooLarge
		return
	}
	*t = Path2(value)
	return nil
}

// type filename2

func (t *Filename2) Encode(e *xdr.Encoder) (err error) {
	value := string(*t)

	// value: string<V2_MAXNAMLEN>

	basic.SniffEncode("value", value)

	if len(value) > V2_MAXNAMLEN {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Filename2) Decode(d *xdr.Decoder) (err error) {
	var value string

	// value: string<V2_MAXNAMLEN>

	*(*string)(&value), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("value", value)

	if len(value) > V2_MAXNAMLEN {
		err = basic.ErrArrayTooLarge
		return
	}
	*t = Filename2(value)
	return nil
}

// type fhandle2

func (t *Fhandle2) Encode(e *xdr.Encoder) (err error) {
	value := [V2_FHSIZE]uint8(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeFixedOpaque((value)[:])
	if err != nil {
		return
	}
	return nil
}

func (t *Fhandle2) Decode(d *xdr.Decoder) (err error) {
	var value [V2_FHSIZE]uint8

	{
		var bytes []byte
		bytes, _, err = d.DecodeFixedOpaque(int32(len(value)))
		if err != nil {
			return
		}
		copy(value[:], bytes)
		basic.SniffDecode("value", value)
	}
	*t = Fhandle2(value)
	return nil
}

// type data2

func (t *Data2) Encode(e *xdr.Encoder) (err error) {
	value := []uint8(*t)

	// value: []uint8<V2_MAXDATA>

	basic.SniffEncode("value", value)

	if len(value) > V2_MAXDATA {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeOpaque(value)
	if err != nil {
		return
	}
	return nil
}

func (t *Data2) Decode(d *xdr.Decoder) (err error) {
	var value []uint8

	// value: []uint8<V2_MAXDATA>

	value, _, err = d.DecodeOpaque()
	if err != nil {
		return
	}

	basic.SniffDecode("value", value)

	if len(value) > V2_MAXDATA {
		err = basic.ErrArrayTooLarge
		return
	}
	*t = Data2(value)
	return nil
}

// type cookie2

func (t *Cookie2) Encode(e *xdr.Encoder) (err error) {
	value := [V2_COOKIESIZE]uint8(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeFixedOpaque((value)[:])
	if err != nil {
		return
	}
	return nil
}

func (t *Cookie2) Decode(d *xdr.Decoder) (err error) {
	var value [V2_COOKIESIZE]uint8

	{
		var bytes []byte
		bytes, _, err = d.DecodeFixedOpaque(int32(len(value)))
		if err != nil {
			return
		}
		copy(value[:], bytes)
		basic.SniffDecode("value", value)
	}
	*t = Cookie2(value)
	return nil
}

// struct time2

func (s *Time2) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Time2 {")
	basic.SniffIndent()
	basic.SniffEncode("s.Seconds", s.Seconds)
	_, err = e.EncodeUint(uint32(s.Seconds))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Useconds", s.Useconds)
	_, err = e.EncodeUint(uint32(s.Useconds))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Time2)")
	return nil
}

func (s *Time2) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Time2 {")
	basic.SniffIndent()
	s.Seconds, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Seconds", s.Seconds)
	s.Useconds, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Useconds", s.Useconds)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Time2)")
	return nil
}

// struct fattr2

func (s *Fattr2) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fattr2 {")
	basic.SniffIndent()
	err = s.Type.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Mode", s.Mode)
	_, err = e.EncodeUint(uint32(s.Mode))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Nlink", s.Nlink)
	_, err = e.EncodeUint(uint32(s.Nlink))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Uid", s.Uid)
	_, err = e.EncodeUint(uint32(s.Uid))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Gid", s.Gid)
	_, err = e.EncodeUint(uint32(s.Gid))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Size", s.Size)
	_, err = e.EncodeUint(uint32(s.Size))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Blocksize", s.Blocksize)
	_, err = e.EncodeUint(uint32(s.Blocksize))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Rdev", s.Rdev)
	_, err = e.EncodeUint(uint32(s.Rdev))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Blocks", s.Blocks)
	_, err = e.EncodeUint(uint32(s.Blocks))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Fsid", s.Fsid)
	_, err = e.EncodeUint(uint32(s.Fsid))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Fileid", s.Fileid)
	_, err = e.EncodeUint(uint32(s.Fileid))
	if err != nil {
		return
	}
	err = s.Atime.Encode(e)
	if err != nil {
		return
	}
	err = s.Mtime.Encode(e)
	if err != nil {
		return
	}
	err = s.Ctime.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fattr2)")
	return nil
}

func (s *Fattr2) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fattr2 {")
	basic.SniffIndent()
	err = s.Type.Decode(d)
	if err != nil {
		return
	}
	s.Mode, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Mode", s.Mode)
	s.Nlink, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Nlink", s.Nlink)
	s.Uid, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Uid", s.Uid)
	s.Gid, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Gid", s.Gid)
	s.Size, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Size", s.Size)
	s.Blocksize, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Blocksize", s.Blocksize)
	s.Rdev, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Rdev", s.Rdev)
	s.Blocks, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Blocks", s.Blocks)
	s.Fsid, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Fsid", s.Fsid)
	s.Fileid, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Fileid", s.Fileid)
	err = s.Atime.Decode(d)
	if err != nil {
		return
	}
	err = s.Mtime.Decode(d)
	if err != nil {
		return
	}
	err = s.Ctime.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fattr2)")
	return nil
}

// union fhstatus2

func (s *Fhstatus2) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Fhstatus2 {")
	basic.SniffIndent()
	basic.SniffEncode("s.Status", s.Status)
	_, err = e.EncodeUint(uint32(s.Status))
	if err != nil {
		return
	}
	switch s.Status {
	case 0:
		u, ok := s.Union.(Fhandle2)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Fhstatus2)")
	return nil
}

func (s *Fhstatus2) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Fhstatus2 {")
	basic.SniffIndent()
	s.Status, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Status", s.Status)
	switch s.Status {
	case 0:
		u := new(Fhandle2)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Fhstatus2)")
	return nil
}

// struct diropargs2

func (s *Diropargs2) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Diropargs2 {")
	basic.SniffIndent()
	err = s.Dir.Encode(e)
	if err != nil {
		return
	}
	err = s.Name.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Diropargs2)")
	return nil
}

func (s *Diropargs2) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Diropargs2 {")
	basic.SniffIndent()
	err = s.Dir.Decode(d)
	if err != nil {
		return
	}
	err = s.Name.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Diropargs2)")
	return nil
}

// struct DIROP2resok

func (s *Dirop2resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Dirop2resok {")
	basic.SniffIndent()
	err = s.File.Encode(e)
	if err != nil {
		return
	}
	err = s.Attributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Dirop2resok)")
	return nil
}

func (s *Dirop2resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Dirop2resok {")
	basic.SniffIndent()
	err = s.File.Decode(d)
	if err != nil {
		return
	}
	err = s.Attributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Dirop2resok)")
	return nil
}

// union DIROP2res

func (s *Dirop2res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Dirop2res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u, ok := s.Union.(Dirop2resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Dirop2res)")
	return nil
}

func (s *Dirop2res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Dirop2res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u := new(Dirop2resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Dirop2res)")
	return nil
}

// union ATTR2res

func (s *Attr2res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Attr2res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u, ok := s.Union.(Fattr2)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Attr2res)")
	return nil
}

func (s *Attr2res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Attr2res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u := new(Fattr2)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Attr2res)")
	return nil
}

// struct sattr2

func (s *Sattr2) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Sattr2 {")
	basic.SniffIndent()
	basic.SniffEncode("s.Mode", s.Mode)
	_, err = e.EncodeUint(uint32(s.Mode))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Uid", s.Uid)
	_, err = e.EncodeUint(uint32(s.Uid))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Gid", s.Gid)
	_, err = e.EncodeUint(uint32(s.Gid))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Size", s.Size)
	_, err = e.EncodeUint(uint32(s.Size))
	if err != nil {
		return
	}
	err = s.Atime.Encode(e)
	if err != nil {
		return
	}
	err = s.Mtime.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Sattr2)")
	return nil
}

func (s *Sattr2) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Sattr2 {")
	basic.SniffIndent()
	s.Mode, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Mode", s.Mode)
	s.Uid, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Uid", s.Uid)
	s.Gid, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Gid", s.Gid)
	s.Size, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Size", s.Size)
	err = s.Atime.Decode(d)
	if err != nil {
		return
	}
	err = s.Mtime.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Sattr2)")
	return nil
}

// struct statinfo2

func (s *Statinfo2) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Statinfo2 {")
	basic.SniffIndent()
	basic.SniffEncode("s.Tsize", s.Tsize)
	_, err = e.EncodeUint(uint32(s.Tsize))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Bsize", s.Bsize)
	_, err = e.EncodeUint(uint32(s.Bsize))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Blocks", s.Blocks)
	_, err = e.EncodeUint(uint32(s.Blocks))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Bfree", s.Bfree)
	_, err = e.EncodeUint(uint32(s.Bfree))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Bavail", s.Bavail)
	_, err = e.EncodeUint(uint32(s.Bavail))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Statinfo2)")
	return nil
}

func (s *Statinfo2) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Statinfo2 {")
	basic.SniffIndent()
	s.Tsize, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Tsize", s.Tsize)
	s.Bsize, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Bsize", s.Bsize)
	s.Blocks, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Blocks", s.Blocks)
	s.Bfree, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Bfree", s.Bfree)
	s.Bavail, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Bavail", s.Bavail)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Statinfo2)")
	return nil
}

// union STATFS2res

func (s *Statfs2res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Statfs2res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u, ok := s.Union.(Statinfo2)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Statfs2res)")
	return nil
}

func (s *Statfs2res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Statfs2res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u := new(Statinfo2)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Statfs2res)")
	return nil
}

// struct READDIR2args

func (s *Readdir2args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readdir2args {")
	basic.SniffIndent()
	err = s.Dir.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Count", s.Count)
	_, err = e.EncodeUint(uint32(s.Count))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readdir2args)")
	return nil
}

func (s *Readdir2args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readdir2args {")
	basic.SniffIndent()
	err = s.Dir.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	s.Count, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Count", s.Count)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readdir2args)")
	return nil
}

// struct entry2

func (s *Entry2) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Entry2 {")
	basic.SniffIndent()
	basic.SniffEncode("s.Fileid", s.Fileid)
	_, err = e.EncodeUint(uint32(s.Fileid))
	if err != nil {
		return
	}
	err = s.Name.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	// s.Nextentry: Optional
	{
		var notnull bool

		if s.Nextentry != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.Nextentry.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Entry2)")
	return nil
}

func (s *Entry2) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Entry2 {")
	basic.SniffIndent()
	s.Fileid, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Fileid", s.Fileid)
	err = s.Name.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	// s.Nextentry: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.Nextentry = &Entry2{}
			err = s.Nextentry.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.Nextentry = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Entry2)")
	return nil
}

// struct READDIR2resok

func (s *Readdir2resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readdir2resok {")
	basic.SniffIndent()
	// s.Entries: Optional
	{
		var notnull bool

		if s.Entries != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.Entries.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}
	basic.SniffEncode("s.Eof", s.Eof)
	_, err = e.EncodeBool(bool(s.Eof))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readdir2resok)")
	return nil
}

func (s *Readdir2resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readdir2resok {")
	basic.SniffIndent()
	// s.Entries: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.Entries = &Entry2{}
			err = s.Entries.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.Entries = nil
		}
	}
	s.Eof, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Eof", s.Eof)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readdir2resok)")
	return nil
}

// union READDIR2res

func (s *Readdir2res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Readdir2res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u, ok := s.Union.(Readdir2resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Readdir2res)")
	return nil
}

func (s *Readdir2res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Readdir2res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u := new(Readdir2resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Readdir2res)")
	return nil
}

// struct SYMLINK2args

func (s *Symlink2args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Symlink2args {")
	basic.SniffIndent()
	err = s.From.Encode(e)
	if err != nil {
		return
	}
	err = s.To.Encode(e)
	if err != nil {
		return
	}
	err = s.Attributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Symlink2args)")
	return nil
}

func (s *Symlink2args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Symlink2args {")
	basic.SniffIndent()
	err = s.From.Decode(d)
	if err != nil {
		return
	}
	err = s.To.Decode(d)
	if err != nil {
		return
	}
	err = s.Attributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Symlink2args)")
	return nil
}

// struct LINK2args

func (s *Link2args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Link2args {")
	basic.SniffIndent()
	err = s.From.Encode(e)
	if err != nil {
		return
	}
	err = s.To.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Link2args)")
	return nil
}

func (s *Link2args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Link2args {")
	basic.SniffIndent()
	err = s.From.Decode(d)
	if err != nil {
		return
	}
	err = s.To.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Link2args)")
	return nil
}

// struct RENAME2args

func (s *Rename2args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Rename2args {")
	basic.SniffIndent()
	err = s.From.Encode(e)
	if err != nil {
		return
	}
	err = s.To.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Rename2args)")
	return nil
}

func (s *Rename2args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Rename2args {")
	basic.SniffIndent()
	err = s.From.Decode(d)
	if err != nil {
		return
	}
	err = s.To.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Rename2args)")
	return nil
}

// struct CREATE2args

func (s *Create2args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Create2args {")
	basic.SniffIndent()
	err = s.Where.Encode(e)
	if err != nil {
		return
	}
	err = s.Attributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Create2args)")
	return nil
}

func (s *Create2args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Create2args {")
	basic.SniffIndent()
	err = s.Where.Decode(d)
	if err != nil {
		return
	}
	err = s.Attributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Create2args)")
	return nil
}

// struct WRITE2args

func (s *Write2args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Write2args {")
	basic.SniffIndent()
	err = s.File.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Beginoffset", s.Beginoffset)
	_, err = e.EncodeUint(uint32(s.Beginoffset))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Offset", s.Offset)
	_, err = e.EncodeUint(uint32(s.Offset))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Totalcount", s.Totalcount)
	_, err = e.EncodeUint(uint32(s.Totalcount))
	if err != nil {
		return
	}
	err = s.Data.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Write2args)")
	return nil
}

func (s *Write2args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Write2args {")
	basic.SniffIndent()
	err = s.File.Decode(d)
	if err != nil {
		return
	}
	s.Beginoffset, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Beginoffset", s.Beginoffset)
	s.Offset, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Offset", s.Offset)
	s.Totalcount, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Totalcount", s.Totalcount)
	err = s.Data.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Write2args)")
	return nil
}

// struct READ2resok

func (s *Read2resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Read2resok {")
	basic.SniffIndent()
	err = s.Attributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Data.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Read2resok)")
	return nil
}

func (s *Read2resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Read2resok {")
	basic.SniffIndent()
	err = s.Attributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Data.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Read2resok)")
	return nil
}

// union READ2res

func (s *Read2res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Read2res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u, ok := s.Union.(Read2resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Read2res)")
	return nil
}

func (s *Read2res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Read2res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u := new(Read2resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Read2res)")
	return nil
}

// struct READ2args

func (s *Read2args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Read2args {")
	basic.SniffIndent()
	err = s.File.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Offset", s.Offset)
	_, err = e.EncodeUint(uint32(s.Offset))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Count", s.Count)
	_, err = e.EncodeUint(uint32(s.Count))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Totalcount", s.Totalcount)
	_, err = e.EncodeUint(uint32(s.Totalcount))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Read2args)")
	return nil
}

func (s *Read2args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Read2args {")
	basic.SniffIndent()
	err = s.File.Decode(d)
	if err != nil {
		return
	}
	s.Offset, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Offset", s.Offset)
	s.Count, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Count", s.Count)
	s.Totalcount, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Totalcount", s.Totalcount)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Read2args)")
	return nil
}

// union READLINK2res

func (s *Readlink2res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Readlink2res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u, ok := s.Union.(Path2)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Readlink2res)")
	return nil
}

func (s *Readlink2res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Readlink2res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case _OK:
		u := new(Path2)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Readlink2res)")
	return nil
}

// struct SETATTR2args

func (s *Setattr2args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Setattr2args {")
	basic.SniffIndent()
	err = s.File.Encode(e)
	if err != nil {
		return
	}
	err = s.Attributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Setattr2args)")
	return nil
}

func (s *Setattr2args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Setattr2args {")
	basic.SniffIndent()
	err = s.File.Decode(d)
	if err != nil {
		return
	}
	err = s.Attributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Setattr2args)")
	return nil
}

// type uint64

func (t *Uint64) Encode(e *xdr.Encoder) (err error) {
	value := uint64(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeUhyper(uint64(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Uint64) Decode(d *xdr.Decoder) (err error) {
	var value uint64

	value, _, err = d.DecodeUhyper()

	if err != nil {
		return
	}
	basic.SniffDecode("value", value)
	*t = Uint64(value)
	return nil
}

// type int64

func (t *Int64) Encode(e *xdr.Encoder) (err error) {
	value := int64(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeHyper(int64(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Int64) Decode(d *xdr.Decoder) (err error) {
	var value int64

	value, _, err = d.DecodeHyper()

	if err != nil {
		return
	}
	basic.SniffDecode("value", value)
	*t = Int64(value)
	return nil
}

// type uint32

func (t *Uint32) Encode(e *xdr.Encoder) (err error) {
	value := uint32(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeUint(uint32(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Uint32) Decode(d *xdr.Decoder) (err error) {
	var value uint32

	value, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("value", value)
	*t = Uint32(value)
	return nil
}

// type int32

func (t *Int32) Encode(e *xdr.Encoder) (err error) {
	value := int32(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeInt(int32(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Int32) Decode(d *xdr.Decoder) (err error) {
	var value int32

	value, _, err = d.DecodeInt()

	if err != nil {
		return
	}
	basic.SniffDecode("value", value)
	*t = Int32(value)
	return nil
}

// type filename3

func (t *Filename3) Encode(e *xdr.Encoder) (err error) {
	value := string(*t)

	// value: string<MAX_ARRAY_LENGTH>

	basic.SniffEncode("value", value)

	if len(value) > MAX_ARRAY_LENGTH {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Filename3) Decode(d *xdr.Decoder) (err error) {
	var value string

	// value: string<MAX_ARRAY_LENGTH>

	*(*string)(&value), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("value", value)

	if len(value) > MAX_ARRAY_LENGTH {
		err = basic.ErrArrayTooLarge
		return
	}
	*t = Filename3(value)
	return nil
}

// type path3

func (t *Path3) Encode(e *xdr.Encoder) (err error) {
	value := []uint8(*t)

	// value: []uint8<V3_MAX_DATA>

	basic.SniffEncode("value", value)

	if len(value) > V3_MAX_DATA {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeOpaque(value)
	if err != nil {
		return
	}
	return nil
}

func (t *Path3) Decode(d *xdr.Decoder) (err error) {
	var value []uint8

	// value: []uint8<V3_MAX_DATA>

	value, _, err = d.DecodeOpaque()
	if err != nil {
		return
	}

	basic.SniffDecode("value", value)

	if len(value) > V3_MAX_DATA {
		err = basic.ErrArrayTooLarge
		return
	}
	*t = Path3(value)
	return nil
}

// type fileid3

func (t *Fileid3) Encode(e *xdr.Encoder) (err error) {
	value := Uint64(*t)

	err = value.Encode(e)
	if err != nil {
		return
	}
	return nil
}

func (t *Fileid3) Decode(d *xdr.Decoder) (err error) {
	var value Uint64

	err = value.Decode(d)
	if err != nil {
		return
	}
	*t = Fileid3(value)
	return nil
}

// type cookie3

func (t *Cookie3) Encode(e *xdr.Encoder) (err error) {
	value := Uint64(*t)

	err = value.Encode(e)
	if err != nil {
		return
	}
	return nil
}

func (t *Cookie3) Decode(d *xdr.Decoder) (err error) {
	var value Uint64

	err = value.Decode(d)
	if err != nil {
		return
	}
	*t = Cookie3(value)
	return nil
}

// type cookieverf3

func (t *Cookieverf3) Encode(e *xdr.Encoder) (err error) {
	value := [V3_COOKIEVERFSIZE]uint8(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeFixedOpaque((value)[:])
	if err != nil {
		return
	}
	return nil
}

func (t *Cookieverf3) Decode(d *xdr.Decoder) (err error) {
	var value [V3_COOKIEVERFSIZE]uint8

	{
		var bytes []byte
		bytes, _, err = d.DecodeFixedOpaque(int32(len(value)))
		if err != nil {
			return
		}
		copy(value[:], bytes)
		basic.SniffDecode("value", value)
	}
	*t = Cookieverf3(value)
	return nil
}

// type createverf3

func (t *Createverf3) Encode(e *xdr.Encoder) (err error) {
	value := [V3_CREATEVERFSIZE]uint8(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeFixedOpaque((value)[:])
	if err != nil {
		return
	}
	return nil
}

func (t *Createverf3) Decode(d *xdr.Decoder) (err error) {
	var value [V3_CREATEVERFSIZE]uint8

	{
		var bytes []byte
		bytes, _, err = d.DecodeFixedOpaque(int32(len(value)))
		if err != nil {
			return
		}
		copy(value[:], bytes)
		basic.SniffDecode("value", value)
	}
	*t = Createverf3(value)
	return nil
}

// type writeverf3

func (t *Writeverf3) Encode(e *xdr.Encoder) (err error) {
	value := [V3_WRITEVERFSIZE]uint8(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeFixedOpaque((value)[:])
	if err != nil {
		return
	}
	return nil
}

func (t *Writeverf3) Decode(d *xdr.Decoder) (err error) {
	var value [V3_WRITEVERFSIZE]uint8

	{
		var bytes []byte
		bytes, _, err = d.DecodeFixedOpaque(int32(len(value)))
		if err != nil {
			return
		}
		copy(value[:], bytes)
		basic.SniffDecode("value", value)
	}
	*t = Writeverf3(value)
	return nil
}

// type uid3

func (t *Uid3) Encode(e *xdr.Encoder) (err error) {
	value := Uint32(*t)

	err = value.Encode(e)
	if err != nil {
		return
	}
	return nil
}

func (t *Uid3) Decode(d *xdr.Decoder) (err error) {
	var value Uint32

	err = value.Decode(d)
	if err != nil {
		return
	}
	*t = Uid3(value)
	return nil
}

// type gid3

func (t *Gid3) Encode(e *xdr.Encoder) (err error) {
	value := Uint32(*t)

	err = value.Encode(e)
	if err != nil {
		return
	}
	return nil
}

func (t *Gid3) Decode(d *xdr.Decoder) (err error) {
	var value Uint32

	err = value.Decode(d)
	if err != nil {
		return
	}
	*t = Gid3(value)
	return nil
}

// type size3

func (t *Size3) Encode(e *xdr.Encoder) (err error) {
	value := Uint64(*t)

	err = value.Encode(e)
	if err != nil {
		return
	}
	return nil
}

func (t *Size3) Decode(d *xdr.Decoder) (err error) {
	var value Uint64

	err = value.Decode(d)
	if err != nil {
		return
	}
	*t = Size3(value)
	return nil
}

// type offset3

func (t *Offset3) Encode(e *xdr.Encoder) (err error) {
	value := Uint64(*t)

	err = value.Encode(e)
	if err != nil {
		return
	}
	return nil
}

func (t *Offset3) Decode(d *xdr.Decoder) (err error) {
	var value Uint64

	err = value.Decode(d)
	if err != nil {
		return
	}
	*t = Offset3(value)
	return nil
}

// type mode3

func (t *Mode3) Encode(e *xdr.Encoder) (err error) {
	value := Uint32(*t)

	err = value.Encode(e)
	if err != nil {
		return
	}
	return nil
}

func (t *Mode3) Decode(d *xdr.Decoder) (err error) {
	var value Uint32

	err = value.Decode(d)
	if err != nil {
		return
	}
	*t = Mode3(value)
	return nil
}

// type count3

func (t *Count3) Encode(e *xdr.Encoder) (err error) {
	value := Uint32(*t)

	err = value.Encode(e)
	if err != nil {
		return
	}
	return nil
}

func (t *Count3) Decode(d *xdr.Decoder) (err error) {
	var value Uint32

	err = value.Decode(d)
	if err != nil {
		return
	}
	*t = Count3(value)
	return nil
}

// struct specdata3

func (s *Specdata3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Specdata3 {")
	basic.SniffIndent()
	err = s.Specdata1.Encode(e)
	if err != nil {
		return
	}
	err = s.Specdata2.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Specdata3)")
	return nil
}

func (s *Specdata3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Specdata3 {")
	basic.SniffIndent()
	err = s.Specdata1.Decode(d)
	if err != nil {
		return
	}
	err = s.Specdata2.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Specdata3)")
	return nil
}

// struct _fh3

func (s *Fh3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fh3 {")
	basic.SniffIndent()

	// s.Data: []uint8<V3_FHSIZE>

	basic.SniffEncode("s.Data", s.Data)

	if len(s.Data) > V3_FHSIZE {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeOpaque(s.Data)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fh3)")
	return nil
}

func (s *Fh3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fh3 {")
	basic.SniffIndent()

	// s.Data: []uint8<V3_FHSIZE>

	s.Data, _, err = d.DecodeOpaque()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Data", s.Data)

	if len(s.Data) > V3_FHSIZE {
		err = basic.ErrArrayTooLarge
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fh3)")
	return nil
}

// struct time3

func (s *Time3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Time3 {")
	basic.SniffIndent()
	err = s.Seconds.Encode(e)
	if err != nil {
		return
	}
	err = s.Nseconds.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Time3)")
	return nil
}

func (s *Time3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Time3 {")
	basic.SniffIndent()
	err = s.Seconds.Decode(d)
	if err != nil {
		return
	}
	err = s.Nseconds.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Time3)")
	return nil
}

// struct fattr3

func (s *Fattr3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fattr3 {")
	basic.SniffIndent()
	err = s.Type.Encode(e)
	if err != nil {
		return
	}
	err = s.Mode.Encode(e)
	if err != nil {
		return
	}
	err = s.Nlink.Encode(e)
	if err != nil {
		return
	}
	err = s.Uid.Encode(e)
	if err != nil {
		return
	}
	err = s.Gid.Encode(e)
	if err != nil {
		return
	}
	err = s.Size.Encode(e)
	if err != nil {
		return
	}
	err = s.Used.Encode(e)
	if err != nil {
		return
	}
	err = s.Rdev.Encode(e)
	if err != nil {
		return
	}
	err = s.Fsid.Encode(e)
	if err != nil {
		return
	}
	err = s.Fileid.Encode(e)
	if err != nil {
		return
	}
	err = s.Atime.Encode(e)
	if err != nil {
		return
	}
	err = s.Mtime.Encode(e)
	if err != nil {
		return
	}
	err = s.Ctime.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fattr3)")
	return nil
}

func (s *Fattr3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fattr3 {")
	basic.SniffIndent()
	err = s.Type.Decode(d)
	if err != nil {
		return
	}
	err = s.Mode.Decode(d)
	if err != nil {
		return
	}
	err = s.Nlink.Decode(d)
	if err != nil {
		return
	}
	err = s.Uid.Decode(d)
	if err != nil {
		return
	}
	err = s.Gid.Decode(d)
	if err != nil {
		return
	}
	err = s.Size.Decode(d)
	if err != nil {
		return
	}
	err = s.Used.Decode(d)
	if err != nil {
		return
	}
	err = s.Rdev.Decode(d)
	if err != nil {
		return
	}
	err = s.Fsid.Decode(d)
	if err != nil {
		return
	}
	err = s.Fileid.Decode(d)
	if err != nil {
		return
	}
	err = s.Atime.Decode(d)
	if err != nil {
		return
	}
	err = s.Mtime.Decode(d)
	if err != nil {
		return
	}
	err = s.Ctime.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fattr3)")
	return nil
}

// union post_op_attr

func (s *PostOpAttr) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union PostOpAttr {")
	basic.SniffIndent()
	basic.SniffEncode("s.AttributesFollow", s.AttributesFollow)
	_, err = e.EncodeBool(bool(s.AttributesFollow))
	if err != nil {
		return
	}
	switch s.AttributesFollow {
	case true:
		u, ok := s.Union.(Fattr3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case false:
	// Empty

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union PostOpAttr)")
	return nil
}

func (s *PostOpAttr) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union PostOpAttr {")
	basic.SniffIndent()
	s.AttributesFollow, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.AttributesFollow", s.AttributesFollow)
	switch s.AttributesFollow {
	case true:
		u := new(Fattr3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case false:
	// Empty

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union PostOpAttr)")
	return nil
}

// struct wcc_attr

func (s *WccAttr) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct WccAttr {")
	basic.SniffIndent()
	err = s.Size.Encode(e)
	if err != nil {
		return
	}
	err = s.Mtime.Encode(e)
	if err != nil {
		return
	}
	err = s.Ctime.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct WccAttr)")
	return nil
}

func (s *WccAttr) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct WccAttr {")
	basic.SniffIndent()
	err = s.Size.Decode(d)
	if err != nil {
		return
	}
	err = s.Mtime.Decode(d)
	if err != nil {
		return
	}
	err = s.Ctime.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct WccAttr)")
	return nil
}

// union pre_op_attr

func (s *PreOpAttr) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union PreOpAttr {")
	basic.SniffIndent()
	basic.SniffEncode("s.AttributesFollow", s.AttributesFollow)
	_, err = e.EncodeBool(bool(s.AttributesFollow))
	if err != nil {
		return
	}
	switch s.AttributesFollow {
	case true:
		u, ok := s.Union.(WccAttr)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case false:
	// Empty

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union PreOpAttr)")
	return nil
}

func (s *PreOpAttr) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union PreOpAttr {")
	basic.SniffIndent()
	s.AttributesFollow, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.AttributesFollow", s.AttributesFollow)
	switch s.AttributesFollow {
	case true:
		u := new(WccAttr)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case false:
	// Empty

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union PreOpAttr)")
	return nil
}

// struct wcc_data

func (s *WccData) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct WccData {")
	basic.SniffIndent()
	err = s.Before.Encode(e)
	if err != nil {
		return
	}
	err = s.After.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct WccData)")
	return nil
}

func (s *WccData) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct WccData {")
	basic.SniffIndent()
	err = s.Before.Decode(d)
	if err != nil {
		return
	}
	err = s.After.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct WccData)")
	return nil
}

// union post_op_fh3

func (s *PostOpFh3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union PostOpFh3 {")
	basic.SniffIndent()
	basic.SniffEncode("s.HandleFollows", s.HandleFollows)
	_, err = e.EncodeBool(bool(s.HandleFollows))
	if err != nil {
		return
	}
	switch s.HandleFollows {
	case true:
		u, ok := s.Union.(Fh3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case false:
	// Empty

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union PostOpFh3)")
	return nil
}

func (s *PostOpFh3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union PostOpFh3 {")
	basic.SniffIndent()
	s.HandleFollows, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.HandleFollows", s.HandleFollows)
	switch s.HandleFollows {
	case true:
		u := new(Fh3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case false:
	// Empty

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union PostOpFh3)")
	return nil
}

// union set_mode3

func (s *SetMode3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union SetMode3 {")
	basic.SniffIndent()
	basic.SniffEncode("s.SetIt", s.SetIt)
	_, err = e.EncodeBool(bool(s.SetIt))
	if err != nil {
		return
	}
	switch s.SetIt {
	case true:
		u, ok := s.Union.(Mode3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union SetMode3)")
	return nil
}

func (s *SetMode3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union SetMode3 {")
	basic.SniffIndent()
	s.SetIt, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.SetIt", s.SetIt)
	switch s.SetIt {
	case true:
		u := new(Mode3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union SetMode3)")
	return nil
}

// union set_uid3

func (s *SetUid3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union SetUid3 {")
	basic.SniffIndent()
	basic.SniffEncode("s.SetIt", s.SetIt)
	_, err = e.EncodeBool(bool(s.SetIt))
	if err != nil {
		return
	}
	switch s.SetIt {
	case true:
		u, ok := s.Union.(Uid3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union SetUid3)")
	return nil
}

func (s *SetUid3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union SetUid3 {")
	basic.SniffIndent()
	s.SetIt, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.SetIt", s.SetIt)
	switch s.SetIt {
	case true:
		u := new(Uid3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union SetUid3)")
	return nil
}

// union set_gid3

func (s *SetGid3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union SetGid3 {")
	basic.SniffIndent()
	basic.SniffEncode("s.SetIt", s.SetIt)
	_, err = e.EncodeBool(bool(s.SetIt))
	if err != nil {
		return
	}
	switch s.SetIt {
	case true:
		u, ok := s.Union.(Gid3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union SetGid3)")
	return nil
}

func (s *SetGid3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union SetGid3 {")
	basic.SniffIndent()
	s.SetIt, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.SetIt", s.SetIt)
	switch s.SetIt {
	case true:
		u := new(Gid3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union SetGid3)")
	return nil
}

// union set_size3

func (s *SetSize3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union SetSize3 {")
	basic.SniffIndent()
	basic.SniffEncode("s.SetIt", s.SetIt)
	_, err = e.EncodeBool(bool(s.SetIt))
	if err != nil {
		return
	}
	switch s.SetIt {
	case true:
		u, ok := s.Union.(Size3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union SetSize3)")
	return nil
}

func (s *SetSize3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union SetSize3 {")
	basic.SniffIndent()
	s.SetIt, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.SetIt", s.SetIt)
	switch s.SetIt {
	case true:
		u := new(Size3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union SetSize3)")
	return nil
}

// union set_atime

func (s *SetAtime) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union SetAtime {")
	basic.SniffIndent()
	err = s.SetIt.Encode(e)
	if err != nil {
		return
	}
	switch s.SetIt {
	case SET_TO_CLIENT_TIME:
		u, ok := s.Union.(Time3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union SetAtime)")
	return nil
}

func (s *SetAtime) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union SetAtime {")
	basic.SniffIndent()
	err = s.SetIt.Decode(d)
	if err != nil {
		return
	}
	switch s.SetIt {
	case SET_TO_CLIENT_TIME:
		u := new(Time3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union SetAtime)")
	return nil
}

// union set_mtime

func (s *SetMtime) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union SetMtime {")
	basic.SniffIndent()
	err = s.SetIt.Encode(e)
	if err != nil {
		return
	}
	switch s.SetIt {
	case SET_TO_CLIENT_TIME:
		u, ok := s.Union.(Time3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union SetMtime)")
	return nil
}

func (s *SetMtime) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union SetMtime {")
	basic.SniffIndent()
	err = s.SetIt.Decode(d)
	if err != nil {
		return
	}
	switch s.SetIt {
	case SET_TO_CLIENT_TIME:
		u := new(Time3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union SetMtime)")
	return nil
}

// struct sattr3

func (s *Sattr3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Sattr3 {")
	basic.SniffIndent()
	err = s.Mode.Encode(e)
	if err != nil {
		return
	}
	err = s.Uid.Encode(e)
	if err != nil {
		return
	}
	err = s.Gid.Encode(e)
	if err != nil {
		return
	}
	err = s.Size.Encode(e)
	if err != nil {
		return
	}
	err = s.Atime.Encode(e)
	if err != nil {
		return
	}
	err = s.Mtime.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Sattr3)")
	return nil
}

func (s *Sattr3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Sattr3 {")
	basic.SniffIndent()
	err = s.Mode.Decode(d)
	if err != nil {
		return
	}
	err = s.Uid.Decode(d)
	if err != nil {
		return
	}
	err = s.Gid.Decode(d)
	if err != nil {
		return
	}
	err = s.Size.Decode(d)
	if err != nil {
		return
	}
	err = s.Atime.Decode(d)
	if err != nil {
		return
	}
	err = s.Mtime.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Sattr3)")
	return nil
}

// struct diropargs3

func (s *Diropargs3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Diropargs3 {")
	basic.SniffIndent()
	err = s.Dir.Encode(e)
	if err != nil {
		return
	}
	err = s.Name.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Diropargs3)")
	return nil
}

func (s *Diropargs3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Diropargs3 {")
	basic.SniffIndent()
	err = s.Dir.Decode(d)
	if err != nil {
		return
	}
	err = s.Name.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Diropargs3)")
	return nil
}

// struct GETATTR3args

func (s *Getattr3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Getattr3args {")
	basic.SniffIndent()
	err = s.Object.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Getattr3args)")
	return nil
}

func (s *Getattr3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Getattr3args {")
	basic.SniffIndent()
	err = s.Object.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Getattr3args)")
	return nil
}

// struct GETATTR3resok

func (s *Getattr3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Getattr3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Getattr3resok)")
	return nil
}

func (s *Getattr3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Getattr3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Getattr3resok)")
	return nil
}

// union GETATTR3res

func (s *Getattr3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Getattr3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Getattr3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Getattr3res)")
	return nil
}

func (s *Getattr3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Getattr3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Getattr3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Getattr3res)")
	return nil
}

// union sattrguard3

func (s *Sattrguard3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Sattrguard3 {")
	basic.SniffIndent()
	basic.SniffEncode("s.Check", s.Check)
	_, err = e.EncodeBool(bool(s.Check))
	if err != nil {
		return
	}
	switch s.Check {
	case true:
		u, ok := s.Union.(Time3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case false:
	// Empty

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Sattrguard3)")
	return nil
}

func (s *Sattrguard3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Sattrguard3 {")
	basic.SniffIndent()
	s.Check, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Check", s.Check)
	switch s.Check {
	case true:
		u := new(Time3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case false:
	// Empty

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Sattrguard3)")
	return nil
}

// struct SETATTR3args

func (s *Setattr3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Setattr3args {")
	basic.SniffIndent()
	err = s.Object.Encode(e)
	if err != nil {
		return
	}
	err = s.NewAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Guard.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Setattr3args)")
	return nil
}

func (s *Setattr3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Setattr3args {")
	basic.SniffIndent()
	err = s.Object.Decode(d)
	if err != nil {
		return
	}
	err = s.NewAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Guard.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Setattr3args)")
	return nil
}

// struct SETATTR3resok

func (s *Setattr3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Setattr3resok {")
	basic.SniffIndent()
	err = s.ObjWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Setattr3resok)")
	return nil
}

func (s *Setattr3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Setattr3resok {")
	basic.SniffIndent()
	err = s.ObjWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Setattr3resok)")
	return nil
}

// struct SETATTR3resfail

func (s *Setattr3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Setattr3resfail {")
	basic.SniffIndent()
	err = s.ObjWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Setattr3resfail)")
	return nil
}

func (s *Setattr3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Setattr3resfail {")
	basic.SniffIndent()
	err = s.ObjWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Setattr3resfail)")
	return nil
}

// union SETATTR3res

func (s *Setattr3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Setattr3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Setattr3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Setattr3res)")
	return nil
}

func (s *Setattr3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Setattr3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Setattr3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Setattr3res)")
	return nil
}

// struct LOOKUP3args

func (s *Lookup3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Lookup3args {")
	basic.SniffIndent()
	err = s.What.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Lookup3args)")
	return nil
}

func (s *Lookup3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Lookup3args {")
	basic.SniffIndent()
	err = s.What.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Lookup3args)")
	return nil
}

// struct LOOKUP3resok

func (s *Lookup3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Lookup3resok {")
	basic.SniffIndent()
	err = s.Object.Encode(e)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.DirAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Lookup3resok)")
	return nil
}

func (s *Lookup3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Lookup3resok {")
	basic.SniffIndent()
	err = s.Object.Decode(d)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.DirAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Lookup3resok)")
	return nil
}

// struct LOOKUP3resfail

func (s *Lookup3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Lookup3resfail {")
	basic.SniffIndent()
	err = s.DirAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Lookup3resfail)")
	return nil
}

func (s *Lookup3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Lookup3resfail {")
	basic.SniffIndent()
	err = s.DirAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Lookup3resfail)")
	return nil
}

// union LOOKUP3res

func (s *Lookup3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Lookup3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Lookup3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Lookup3res)")
	return nil
}

func (s *Lookup3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Lookup3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Lookup3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Lookup3res)")
	return nil
}

// struct ACCESS3args

func (s *Access3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Access3args {")
	basic.SniffIndent()
	err = s.Object.Encode(e)
	if err != nil {
		return
	}
	err = s.Access.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Access3args)")
	return nil
}

func (s *Access3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Access3args {")
	basic.SniffIndent()
	err = s.Object.Decode(d)
	if err != nil {
		return
	}
	err = s.Access.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Access3args)")
	return nil
}

// struct ACCESS3resok

func (s *Access3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Access3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Access.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Access3resok)")
	return nil
}

func (s *Access3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Access3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Access.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Access3resok)")
	return nil
}

// struct ACCESS3resfail

func (s *Access3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Access3resfail {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Access3resfail)")
	return nil
}

func (s *Access3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Access3resfail {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Access3resfail)")
	return nil
}

// union ACCESS3res

func (s *Access3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Access3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Access3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Access3res)")
	return nil
}

func (s *Access3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Access3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Access3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Access3res)")
	return nil
}

// struct READLINK3args

func (s *Readlink3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readlink3args {")
	basic.SniffIndent()
	err = s.Symlink.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readlink3args)")
	return nil
}

func (s *Readlink3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readlink3args {")
	basic.SniffIndent()
	err = s.Symlink.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readlink3args)")
	return nil
}

// struct READLINK3resok

func (s *Readlink3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readlink3resok {")
	basic.SniffIndent()
	err = s.SymlinkAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Data.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readlink3resok)")
	return nil
}

func (s *Readlink3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readlink3resok {")
	basic.SniffIndent()
	err = s.SymlinkAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Data.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readlink3resok)")
	return nil
}

// struct READLINK3resfail

func (s *Readlink3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readlink3resfail {")
	basic.SniffIndent()
	err = s.SymlinkAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readlink3resfail)")
	return nil
}

func (s *Readlink3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readlink3resfail {")
	basic.SniffIndent()
	err = s.SymlinkAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readlink3resfail)")
	return nil
}

// union READLINK3res

func (s *Readlink3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Readlink3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Readlink3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Readlink3res)")
	return nil
}

func (s *Readlink3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Readlink3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Readlink3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Readlink3res)")
	return nil
}

// struct READ3args

func (s *Read3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Read3args {")
	basic.SniffIndent()
	err = s.File.Encode(e)
	if err != nil {
		return
	}
	err = s.Offset.Encode(e)
	if err != nil {
		return
	}
	err = s.Count.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Read3args)")
	return nil
}

func (s *Read3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Read3args {")
	basic.SniffIndent()
	err = s.File.Decode(d)
	if err != nil {
		return
	}
	err = s.Offset.Decode(d)
	if err != nil {
		return
	}
	err = s.Count.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Read3args)")
	return nil
}

// struct READ3resok

func (s *Read3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Read3resok {")
	basic.SniffIndent()
	err = s.FileAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Count.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Eof", s.Eof)
	_, err = e.EncodeBool(bool(s.Eof))
	if err != nil {
		return
	}

	// s.Data: []uint8<V3_MAX_DATA>

	basic.SniffEncode("s.Data", s.Data)

	if len(s.Data) > V3_MAX_DATA {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeOpaque(s.Data)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Read3resok)")
	return nil
}

func (s *Read3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Read3resok {")
	basic.SniffIndent()
	err = s.FileAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Count.Decode(d)
	if err != nil {
		return
	}
	s.Eof, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Eof", s.Eof)

	// s.Data: []uint8<V3_MAX_DATA>

	s.Data, _, err = d.DecodeOpaque()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Data", s.Data)

	if len(s.Data) > V3_MAX_DATA {
		err = basic.ErrArrayTooLarge
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Read3resok)")
	return nil
}

// struct READ3resfail

func (s *Read3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Read3resfail {")
	basic.SniffIndent()
	err = s.FileAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Read3resfail)")
	return nil
}

func (s *Read3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Read3resfail {")
	basic.SniffIndent()
	err = s.FileAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Read3resfail)")
	return nil
}

// union READ3res

func (s *Read3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Read3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Read3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Read3res)")
	return nil
}

func (s *Read3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Read3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Read3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Read3res)")
	return nil
}

// struct WRITE3args

func (s *Write3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Write3args {")
	basic.SniffIndent()
	err = s.File.Encode(e)
	if err != nil {
		return
	}
	err = s.Offset.Encode(e)
	if err != nil {
		return
	}
	err = s.Count.Encode(e)
	if err != nil {
		return
	}
	err = s.Stable.Encode(e)
	if err != nil {
		return
	}

	// s.Data: []uint8<V3_MAX_DATA>

	basic.SniffEncode("s.Data", s.Data)

	if len(s.Data) > V3_MAX_DATA {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeOpaque(s.Data)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Write3args)")
	return nil
}

func (s *Write3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Write3args {")
	basic.SniffIndent()
	err = s.File.Decode(d)
	if err != nil {
		return
	}
	err = s.Offset.Decode(d)
	if err != nil {
		return
	}
	err = s.Count.Decode(d)
	if err != nil {
		return
	}
	err = s.Stable.Decode(d)
	if err != nil {
		return
	}

	// s.Data: []uint8<V3_MAX_DATA>

	s.Data, _, err = d.DecodeOpaque()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Data", s.Data)

	if len(s.Data) > V3_MAX_DATA {
		err = basic.ErrArrayTooLarge
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Write3args)")
	return nil
}

// struct WRITE3resok

func (s *Write3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Write3resok {")
	basic.SniffIndent()
	err = s.FileWcc.Encode(e)
	if err != nil {
		return
	}
	err = s.Count.Encode(e)
	if err != nil {
		return
	}
	err = s.Committed.Encode(e)
	if err != nil {
		return
	}
	err = s.Verf.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Write3resok)")
	return nil
}

func (s *Write3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Write3resok {")
	basic.SniffIndent()
	err = s.FileWcc.Decode(d)
	if err != nil {
		return
	}
	err = s.Count.Decode(d)
	if err != nil {
		return
	}
	err = s.Committed.Decode(d)
	if err != nil {
		return
	}
	err = s.Verf.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Write3resok)")
	return nil
}

// struct WRITE3resfail

func (s *Write3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Write3resfail {")
	basic.SniffIndent()
	err = s.FileWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Write3resfail)")
	return nil
}

func (s *Write3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Write3resfail {")
	basic.SniffIndent()
	err = s.FileWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Write3resfail)")
	return nil
}

// union WRITE3res

func (s *Write3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Write3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Write3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Write3res)")
	return nil
}

func (s *Write3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Write3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Write3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Write3res)")
	return nil
}

// union createhow3

func (s *Createhow3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Createhow3 {")
	basic.SniffIndent()
	err = s.Mode.Encode(e)
	if err != nil {
		return
	}
	switch s.Mode {
	case UNCHECKED:
		u, ok := s.Union.(Sattr3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case GUARDED:
		u, ok := s.Union.(Sattr3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case EXCLUSIVE:
		u, ok := s.Union.(Createverf3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Createhow3)")
	return nil
}

func (s *Createhow3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Createhow3 {")
	basic.SniffIndent()
	err = s.Mode.Decode(d)
	if err != nil {
		return
	}
	switch s.Mode {
	case UNCHECKED:
		u := new(Sattr3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case GUARDED:
		u := new(Sattr3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case EXCLUSIVE:
		u := new(Createverf3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Createhow3)")
	return nil
}

// struct CREATE3args

func (s *Create3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Create3args {")
	basic.SniffIndent()
	err = s.Where.Encode(e)
	if err != nil {
		return
	}
	err = s.How.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Create3args)")
	return nil
}

func (s *Create3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Create3args {")
	basic.SniffIndent()
	err = s.Where.Decode(d)
	if err != nil {
		return
	}
	err = s.How.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Create3args)")
	return nil
}

// struct CREATE3resok

func (s *Create3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Create3resok {")
	basic.SniffIndent()
	err = s.Obj.Encode(e)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Create3resok)")
	return nil
}

func (s *Create3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Create3resok {")
	basic.SniffIndent()
	err = s.Obj.Decode(d)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Create3resok)")
	return nil
}

// struct CREATE3resfail

func (s *Create3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Create3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Create3resfail)")
	return nil
}

func (s *Create3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Create3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Create3resfail)")
	return nil
}

// union CREATE3res

func (s *Create3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Create3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Create3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Create3res)")
	return nil
}

func (s *Create3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Create3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Create3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Create3res)")
	return nil
}

// struct MKDIR3args

func (s *Mkdir3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mkdir3args {")
	basic.SniffIndent()
	err = s.Where.Encode(e)
	if err != nil {
		return
	}
	err = s.Attributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mkdir3args)")
	return nil
}

func (s *Mkdir3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mkdir3args {")
	basic.SniffIndent()
	err = s.Where.Decode(d)
	if err != nil {
		return
	}
	err = s.Attributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mkdir3args)")
	return nil
}

// struct MKDIR3resok

func (s *Mkdir3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mkdir3resok {")
	basic.SniffIndent()
	err = s.Obj.Encode(e)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mkdir3resok)")
	return nil
}

func (s *Mkdir3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mkdir3resok {")
	basic.SniffIndent()
	err = s.Obj.Decode(d)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mkdir3resok)")
	return nil
}

// struct MKDIR3resfail

func (s *Mkdir3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mkdir3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mkdir3resfail)")
	return nil
}

func (s *Mkdir3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mkdir3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mkdir3resfail)")
	return nil
}

// union MKDIR3res

func (s *Mkdir3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Mkdir3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Mkdir3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Mkdir3res)")
	return nil
}

func (s *Mkdir3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Mkdir3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Mkdir3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Mkdir3res)")
	return nil
}

// struct symlinkdata3

func (s *Symlinkdata3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Symlinkdata3 {")
	basic.SniffIndent()
	err = s.SymlinkAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.SymlinkData.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Symlinkdata3)")
	return nil
}

func (s *Symlinkdata3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Symlinkdata3 {")
	basic.SniffIndent()
	err = s.SymlinkAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.SymlinkData.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Symlinkdata3)")
	return nil
}

// struct SYMLINK3args

func (s *Symlink3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Symlink3args {")
	basic.SniffIndent()
	err = s.Where.Encode(e)
	if err != nil {
		return
	}
	err = s.Symlink.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Symlink3args)")
	return nil
}

func (s *Symlink3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Symlink3args {")
	basic.SniffIndent()
	err = s.Where.Decode(d)
	if err != nil {
		return
	}
	err = s.Symlink.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Symlink3args)")
	return nil
}

// struct SYMLINK3resok

func (s *Symlink3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Symlink3resok {")
	basic.SniffIndent()
	err = s.Obj.Encode(e)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Symlink3resok)")
	return nil
}

func (s *Symlink3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Symlink3resok {")
	basic.SniffIndent()
	err = s.Obj.Decode(d)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Symlink3resok)")
	return nil
}

// struct SYMLINK3resfail

func (s *Symlink3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Symlink3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Symlink3resfail)")
	return nil
}

func (s *Symlink3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Symlink3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Symlink3resfail)")
	return nil
}

// union SYMLINK3res

func (s *Symlink3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Symlink3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Symlink3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Symlink3res)")
	return nil
}

func (s *Symlink3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Symlink3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Symlink3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Symlink3res)")
	return nil
}

// struct devicedata3

func (s *Devicedata3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Devicedata3 {")
	basic.SniffIndent()
	err = s.DevAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Spec.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Devicedata3)")
	return nil
}

func (s *Devicedata3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Devicedata3 {")
	basic.SniffIndent()
	err = s.DevAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Spec.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Devicedata3)")
	return nil
}

// union mknoddata3

func (s *Mknoddata3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Mknoddata3 {")
	basic.SniffIndent()
	err = s.Type.Encode(e)
	if err != nil {
		return
	}
	switch s.Type {
	case NF3CHR:
		// Empty
	case NF3BLK:
		u, ok := s.Union.(Devicedata3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case NF3SOCK:
		// Empty
	case NF3FIFO:
		u, ok := s.Union.(Sattr3)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Mknoddata3)")
	return nil
}

func (s *Mknoddata3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Mknoddata3 {")
	basic.SniffIndent()
	err = s.Type.Decode(d)
	if err != nil {
		return
	}
	switch s.Type {
	case NF3CHR:
		// Empty
	case NF3BLK:
		u := new(Devicedata3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case NF3SOCK:
		// Empty
	case NF3FIFO:
		u := new(Sattr3)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Mknoddata3)")
	return nil
}

// struct MKNOD3args

func (s *Mknod3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mknod3args {")
	basic.SniffIndent()
	err = s.Where.Encode(e)
	if err != nil {
		return
	}
	err = s.What.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mknod3args)")
	return nil
}

func (s *Mknod3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mknod3args {")
	basic.SniffIndent()
	err = s.Where.Decode(d)
	if err != nil {
		return
	}
	err = s.What.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mknod3args)")
	return nil
}

// struct MKNOD3resok

func (s *Mknod3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mknod3resok {")
	basic.SniffIndent()
	err = s.Obj.Encode(e)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mknod3resok)")
	return nil
}

func (s *Mknod3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mknod3resok {")
	basic.SniffIndent()
	err = s.Obj.Decode(d)
	if err != nil {
		return
	}
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mknod3resok)")
	return nil
}

// struct MKNOD3resfail

func (s *Mknod3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mknod3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mknod3resfail)")
	return nil
}

func (s *Mknod3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mknod3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mknod3resfail)")
	return nil
}

// union MKNOD3res

func (s *Mknod3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Mknod3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Mknod3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Mknod3res)")
	return nil
}

func (s *Mknod3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Mknod3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Mknod3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Mknod3res)")
	return nil
}

// struct REMOVE3args

func (s *Remove3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Remove3args {")
	basic.SniffIndent()
	err = s.Object.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Remove3args)")
	return nil
}

func (s *Remove3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Remove3args {")
	basic.SniffIndent()
	err = s.Object.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Remove3args)")
	return nil
}

// struct REMOVE3resok

func (s *Remove3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Remove3resok {")
	basic.SniffIndent()
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Remove3resok)")
	return nil
}

func (s *Remove3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Remove3resok {")
	basic.SniffIndent()
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Remove3resok)")
	return nil
}

// struct REMOVE3resfail

func (s *Remove3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Remove3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Remove3resfail)")
	return nil
}

func (s *Remove3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Remove3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Remove3resfail)")
	return nil
}

// union REMOVE3res

func (s *Remove3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Remove3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Remove3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Remove3res)")
	return nil
}

func (s *Remove3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Remove3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Remove3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Remove3res)")
	return nil
}

// struct RMDIR3args

func (s *Rmdir3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Rmdir3args {")
	basic.SniffIndent()
	err = s.Object.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Rmdir3args)")
	return nil
}

func (s *Rmdir3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Rmdir3args {")
	basic.SniffIndent()
	err = s.Object.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Rmdir3args)")
	return nil
}

// struct RMDIR3resok

func (s *Rmdir3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Rmdir3resok {")
	basic.SniffIndent()
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Rmdir3resok)")
	return nil
}

func (s *Rmdir3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Rmdir3resok {")
	basic.SniffIndent()
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Rmdir3resok)")
	return nil
}

// struct RMDIR3resfail

func (s *Rmdir3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Rmdir3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Rmdir3resfail)")
	return nil
}

func (s *Rmdir3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Rmdir3resfail {")
	basic.SniffIndent()
	err = s.DirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Rmdir3resfail)")
	return nil
}

// union RMDIR3res

func (s *Rmdir3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Rmdir3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Rmdir3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Rmdir3res)")
	return nil
}

func (s *Rmdir3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Rmdir3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Rmdir3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Rmdir3res)")
	return nil
}

// struct RENAME3args

func (s *Rename3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Rename3args {")
	basic.SniffIndent()
	err = s.From.Encode(e)
	if err != nil {
		return
	}
	err = s.To.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Rename3args)")
	return nil
}

func (s *Rename3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Rename3args {")
	basic.SniffIndent()
	err = s.From.Decode(d)
	if err != nil {
		return
	}
	err = s.To.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Rename3args)")
	return nil
}

// struct RENAME3resok

func (s *Rename3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Rename3resok {")
	basic.SniffIndent()
	err = s.FromdirWcc.Encode(e)
	if err != nil {
		return
	}
	err = s.TodirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Rename3resok)")
	return nil
}

func (s *Rename3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Rename3resok {")
	basic.SniffIndent()
	err = s.FromdirWcc.Decode(d)
	if err != nil {
		return
	}
	err = s.TodirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Rename3resok)")
	return nil
}

// struct RENAME3resfail

func (s *Rename3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Rename3resfail {")
	basic.SniffIndent()
	err = s.FromdirWcc.Encode(e)
	if err != nil {
		return
	}
	err = s.TodirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Rename3resfail)")
	return nil
}

func (s *Rename3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Rename3resfail {")
	basic.SniffIndent()
	err = s.FromdirWcc.Decode(d)
	if err != nil {
		return
	}
	err = s.TodirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Rename3resfail)")
	return nil
}

// union RENAME3res

func (s *Rename3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Rename3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Rename3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Rename3res)")
	return nil
}

func (s *Rename3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Rename3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Rename3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Rename3res)")
	return nil
}

// struct LINK3args

func (s *Link3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Link3args {")
	basic.SniffIndent()
	err = s.File.Encode(e)
	if err != nil {
		return
	}
	err = s.Link.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Link3args)")
	return nil
}

func (s *Link3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Link3args {")
	basic.SniffIndent()
	err = s.File.Decode(d)
	if err != nil {
		return
	}
	err = s.Link.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Link3args)")
	return nil
}

// struct LINK3resok

func (s *Link3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Link3resok {")
	basic.SniffIndent()
	err = s.FileAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.LinkdirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Link3resok)")
	return nil
}

func (s *Link3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Link3resok {")
	basic.SniffIndent()
	err = s.FileAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.LinkdirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Link3resok)")
	return nil
}

// struct LINK3resfail

func (s *Link3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Link3resfail {")
	basic.SniffIndent()
	err = s.FileAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.LinkdirWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Link3resfail)")
	return nil
}

func (s *Link3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Link3resfail {")
	basic.SniffIndent()
	err = s.FileAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.LinkdirWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Link3resfail)")
	return nil
}

// union LINK3res

func (s *Link3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Link3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Link3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Link3res)")
	return nil
}

func (s *Link3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Link3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Link3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Link3res)")
	return nil
}

// struct READDIR3args

func (s *Readdir3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readdir3args {")
	basic.SniffIndent()
	err = s.Dir.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookieverf.Encode(e)
	if err != nil {
		return
	}
	err = s.Count.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readdir3args)")
	return nil
}

func (s *Readdir3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readdir3args {")
	basic.SniffIndent()
	err = s.Dir.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookieverf.Decode(d)
	if err != nil {
		return
	}
	err = s.Count.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readdir3args)")
	return nil
}

// struct entry3

func (s *Entry3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Entry3 {")
	basic.SniffIndent()
	err = s.Fileid.Encode(e)
	if err != nil {
		return
	}
	err = s.Name.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	// s.Nextentry: Optional
	{
		var notnull bool

		if s.Nextentry != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.Nextentry.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Entry3)")
	return nil
}

func (s *Entry3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Entry3 {")
	basic.SniffIndent()
	err = s.Fileid.Decode(d)
	if err != nil {
		return
	}
	err = s.Name.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	// s.Nextentry: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.Nextentry = &Entry3{}
			err = s.Nextentry.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.Nextentry = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Entry3)")
	return nil
}

// struct dirlist3

func (s *Dirlist3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Dirlist3 {")
	basic.SniffIndent()
	// s.Entries: Optional
	{
		var notnull bool

		if s.Entries != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.Entries.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}
	basic.SniffEncode("s.Eof", s.Eof)
	_, err = e.EncodeBool(bool(s.Eof))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Dirlist3)")
	return nil
}

func (s *Dirlist3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Dirlist3 {")
	basic.SniffIndent()
	// s.Entries: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.Entries = &Entry3{}
			err = s.Entries.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.Entries = nil
		}
	}
	s.Eof, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Eof", s.Eof)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Dirlist3)")
	return nil
}

// struct READDIR3resok

func (s *Readdir3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readdir3resok {")
	basic.SniffIndent()
	err = s.DirAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookieverf.Encode(e)
	if err != nil {
		return
	}
	err = s.Reply.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readdir3resok)")
	return nil
}

func (s *Readdir3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readdir3resok {")
	basic.SniffIndent()
	err = s.DirAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookieverf.Decode(d)
	if err != nil {
		return
	}
	err = s.Reply.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readdir3resok)")
	return nil
}

// struct READDIR3resfail

func (s *Readdir3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readdir3resfail {")
	basic.SniffIndent()
	err = s.DirAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readdir3resfail)")
	return nil
}

func (s *Readdir3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readdir3resfail {")
	basic.SniffIndent()
	err = s.DirAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readdir3resfail)")
	return nil
}

// union READDIR3res

func (s *Readdir3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Readdir3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Readdir3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Readdir3res)")
	return nil
}

func (s *Readdir3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Readdir3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Readdir3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Readdir3res)")
	return nil
}

// struct READDIRPLUS3args

func (s *Readdirplus3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readdirplus3args {")
	basic.SniffIndent()
	err = s.Dir.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookieverf.Encode(e)
	if err != nil {
		return
	}
	err = s.Dircount.Encode(e)
	if err != nil {
		return
	}
	err = s.Maxcount.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readdirplus3args)")
	return nil
}

func (s *Readdirplus3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readdirplus3args {")
	basic.SniffIndent()
	err = s.Dir.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookieverf.Decode(d)
	if err != nil {
		return
	}
	err = s.Dircount.Decode(d)
	if err != nil {
		return
	}
	err = s.Maxcount.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readdirplus3args)")
	return nil
}

// struct entryplus3

func (s *Entryplus3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Entryplus3 {")
	basic.SniffIndent()
	err = s.Fileid.Encode(e)
	if err != nil {
		return
	}
	err = s.Name.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	err = s.NameAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.NameHandle.Encode(e)
	if err != nil {
		return
	}
	// s.Nextentry: Optional
	{
		var notnull bool

		if s.Nextentry != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.Nextentry.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Entryplus3)")
	return nil
}

func (s *Entryplus3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Entryplus3 {")
	basic.SniffIndent()
	err = s.Fileid.Decode(d)
	if err != nil {
		return
	}
	err = s.Name.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	err = s.NameAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.NameHandle.Decode(d)
	if err != nil {
		return
	}
	// s.Nextentry: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.Nextentry = &Entryplus3{}
			err = s.Nextentry.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.Nextentry = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Entryplus3)")
	return nil
}

// struct dirlistplus3

func (s *Dirlistplus3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Dirlistplus3 {")
	basic.SniffIndent()
	// s.Entries: Optional
	{
		var notnull bool

		if s.Entries != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.Entries.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}
	basic.SniffEncode("s.Eof", s.Eof)
	_, err = e.EncodeBool(bool(s.Eof))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Dirlistplus3)")
	return nil
}

func (s *Dirlistplus3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Dirlistplus3 {")
	basic.SniffIndent()
	// s.Entries: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.Entries = &Entryplus3{}
			err = s.Entries.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.Entries = nil
		}
	}
	s.Eof, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Eof", s.Eof)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Dirlistplus3)")
	return nil
}

// struct READDIRPLUS3resok

func (s *Readdirplus3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readdirplus3resok {")
	basic.SniffIndent()
	err = s.DirAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Cookieverf.Encode(e)
	if err != nil {
		return
	}
	err = s.Reply.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readdirplus3resok)")
	return nil
}

func (s *Readdirplus3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readdirplus3resok {")
	basic.SniffIndent()
	err = s.DirAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Cookieverf.Decode(d)
	if err != nil {
		return
	}
	err = s.Reply.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readdirplus3resok)")
	return nil
}

// struct READDIRPLUS3resfail

func (s *Readdirplus3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Readdirplus3resfail {")
	basic.SniffIndent()
	err = s.DirAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Readdirplus3resfail)")
	return nil
}

func (s *Readdirplus3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Readdirplus3resfail {")
	basic.SniffIndent()
	err = s.DirAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Readdirplus3resfail)")
	return nil
}

// union READDIRPLUS3res

func (s *Readdirplus3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Readdirplus3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Readdirplus3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Readdirplus3res)")
	return nil
}

func (s *Readdirplus3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Readdirplus3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Readdirplus3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Readdirplus3res)")
	return nil
}

// struct FSSTAT3args

func (s *Fsstat3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fsstat3args {")
	basic.SniffIndent()
	err = s.Fsroot.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fsstat3args)")
	return nil
}

func (s *Fsstat3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fsstat3args {")
	basic.SniffIndent()
	err = s.Fsroot.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fsstat3args)")
	return nil
}

// struct FSSTAT3resok

func (s *Fsstat3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fsstat3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Tbytes.Encode(e)
	if err != nil {
		return
	}
	err = s.Fbytes.Encode(e)
	if err != nil {
		return
	}
	err = s.Abytes.Encode(e)
	if err != nil {
		return
	}
	err = s.Tfiles.Encode(e)
	if err != nil {
		return
	}
	err = s.Ffiles.Encode(e)
	if err != nil {
		return
	}
	err = s.Afiles.Encode(e)
	if err != nil {
		return
	}
	err = s.Invarsec.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fsstat3resok)")
	return nil
}

func (s *Fsstat3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fsstat3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Tbytes.Decode(d)
	if err != nil {
		return
	}
	err = s.Fbytes.Decode(d)
	if err != nil {
		return
	}
	err = s.Abytes.Decode(d)
	if err != nil {
		return
	}
	err = s.Tfiles.Decode(d)
	if err != nil {
		return
	}
	err = s.Ffiles.Decode(d)
	if err != nil {
		return
	}
	err = s.Afiles.Decode(d)
	if err != nil {
		return
	}
	err = s.Invarsec.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fsstat3resok)")
	return nil
}

// struct FSSTAT3resfail

func (s *Fsstat3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fsstat3resfail {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fsstat3resfail)")
	return nil
}

func (s *Fsstat3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fsstat3resfail {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fsstat3resfail)")
	return nil
}

// union FSSTAT3res

func (s *Fsstat3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Fsstat3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Fsstat3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Fsstat3res)")
	return nil
}

func (s *Fsstat3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Fsstat3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Fsstat3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Fsstat3res)")
	return nil
}

// struct FSINFO3args

func (s *Fsinfo3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fsinfo3args {")
	basic.SniffIndent()
	err = s.Fsroot.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fsinfo3args)")
	return nil
}

func (s *Fsinfo3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fsinfo3args {")
	basic.SniffIndent()
	err = s.Fsroot.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fsinfo3args)")
	return nil
}

// struct FSINFO3resok

func (s *Fsinfo3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fsinfo3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Rtmax.Encode(e)
	if err != nil {
		return
	}
	err = s.Rtpref.Encode(e)
	if err != nil {
		return
	}
	err = s.Rtmult.Encode(e)
	if err != nil {
		return
	}
	err = s.Wtmax.Encode(e)
	if err != nil {
		return
	}
	err = s.Wtpref.Encode(e)
	if err != nil {
		return
	}
	err = s.Wtmult.Encode(e)
	if err != nil {
		return
	}
	err = s.Dtpref.Encode(e)
	if err != nil {
		return
	}
	err = s.Maxfilesize.Encode(e)
	if err != nil {
		return
	}
	err = s.TimeDelta.Encode(e)
	if err != nil {
		return
	}
	err = s.Properties.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fsinfo3resok)")
	return nil
}

func (s *Fsinfo3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fsinfo3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Rtmax.Decode(d)
	if err != nil {
		return
	}
	err = s.Rtpref.Decode(d)
	if err != nil {
		return
	}
	err = s.Rtmult.Decode(d)
	if err != nil {
		return
	}
	err = s.Wtmax.Decode(d)
	if err != nil {
		return
	}
	err = s.Wtpref.Decode(d)
	if err != nil {
		return
	}
	err = s.Wtmult.Decode(d)
	if err != nil {
		return
	}
	err = s.Dtpref.Decode(d)
	if err != nil {
		return
	}
	err = s.Maxfilesize.Decode(d)
	if err != nil {
		return
	}
	err = s.TimeDelta.Decode(d)
	if err != nil {
		return
	}
	err = s.Properties.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fsinfo3resok)")
	return nil
}

// struct FSINFO3resfail

func (s *Fsinfo3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Fsinfo3resfail {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Fsinfo3resfail)")
	return nil
}

func (s *Fsinfo3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Fsinfo3resfail {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Fsinfo3resfail)")
	return nil
}

// union FSINFO3res

func (s *Fsinfo3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Fsinfo3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Fsinfo3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Fsinfo3res)")
	return nil
}

func (s *Fsinfo3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Fsinfo3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Fsinfo3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Fsinfo3res)")
	return nil
}

// struct PATHCONF3args

func (s *Pathconf3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Pathconf3args {")
	basic.SniffIndent()
	err = s.Object.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Pathconf3args)")
	return nil
}

func (s *Pathconf3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Pathconf3args {")
	basic.SniffIndent()
	err = s.Object.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Pathconf3args)")
	return nil
}

// struct PATHCONF3resok

func (s *Pathconf3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Pathconf3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}
	err = s.Linkmax.Encode(e)
	if err != nil {
		return
	}
	err = s.NameMax.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.NoTrunc", s.NoTrunc)
	_, err = e.EncodeBool(bool(s.NoTrunc))
	if err != nil {
		return
	}
	basic.SniffEncode("s.ChownRestricted", s.ChownRestricted)
	_, err = e.EncodeBool(bool(s.ChownRestricted))
	if err != nil {
		return
	}
	basic.SniffEncode("s.CaseInsensitive", s.CaseInsensitive)
	_, err = e.EncodeBool(bool(s.CaseInsensitive))
	if err != nil {
		return
	}
	basic.SniffEncode("s.CasePreserving", s.CasePreserving)
	_, err = e.EncodeBool(bool(s.CasePreserving))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Pathconf3resok)")
	return nil
}

func (s *Pathconf3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Pathconf3resok {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}
	err = s.Linkmax.Decode(d)
	if err != nil {
		return
	}
	err = s.NameMax.Decode(d)
	if err != nil {
		return
	}
	s.NoTrunc, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.NoTrunc", s.NoTrunc)
	s.ChownRestricted, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.ChownRestricted", s.ChownRestricted)
	s.CaseInsensitive, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.CaseInsensitive", s.CaseInsensitive)
	s.CasePreserving, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.CasePreserving", s.CasePreserving)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Pathconf3resok)")
	return nil
}

// struct PATHCONF3resfail

func (s *Pathconf3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Pathconf3resfail {")
	basic.SniffIndent()
	err = s.ObjAttributes.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Pathconf3resfail)")
	return nil
}

func (s *Pathconf3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Pathconf3resfail {")
	basic.SniffIndent()
	err = s.ObjAttributes.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Pathconf3resfail)")
	return nil
}

// union PATHCONF3res

func (s *Pathconf3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Pathconf3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Pathconf3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Pathconf3res)")
	return nil
}

func (s *Pathconf3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Pathconf3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Pathconf3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Pathconf3res)")
	return nil
}

// struct COMMIT3args

func (s *Commit3args) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Commit3args {")
	basic.SniffIndent()
	err = s.File.Encode(e)
	if err != nil {
		return
	}
	err = s.Offset.Encode(e)
	if err != nil {
		return
	}
	err = s.Count.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Commit3args)")
	return nil
}

func (s *Commit3args) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Commit3args {")
	basic.SniffIndent()
	err = s.File.Decode(d)
	if err != nil {
		return
	}
	err = s.Offset.Decode(d)
	if err != nil {
		return
	}
	err = s.Count.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Commit3args)")
	return nil
}

// struct COMMIT3resok

func (s *Commit3resok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Commit3resok {")
	basic.SniffIndent()
	err = s.FileWcc.Encode(e)
	if err != nil {
		return
	}
	err = s.Verf.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Commit3resok)")
	return nil
}

func (s *Commit3resok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Commit3resok {")
	basic.SniffIndent()
	err = s.FileWcc.Decode(d)
	if err != nil {
		return
	}
	err = s.Verf.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Commit3resok)")
	return nil
}

// struct COMMIT3resfail

func (s *Commit3resfail) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Commit3resfail {")
	basic.SniffIndent()
	err = s.FileWcc.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Commit3resfail)")
	return nil
}

func (s *Commit3resfail) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Commit3resfail {")
	basic.SniffIndent()
	err = s.FileWcc.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Commit3resfail)")
	return nil
}

// union COMMIT3res

func (s *Commit3res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Commit3res {")
	basic.SniffIndent()
	err = s.Status.Encode(e)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u, ok := s.Union.(Commit3resok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Commit3res)")
	return nil
}

func (s *Commit3res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Commit3res {")
	basic.SniffIndent()
	err = s.Status.Decode(d)
	if err != nil {
		return
	}
	switch s.Status {
	case V3_OK:
		u := new(Commit3resok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Commit3res)")
	return nil
}

// type dirpath

func (t *Dirpath) Encode(e *xdr.Encoder) (err error) {
	value := string(*t)

	// value: string<V2_MNTPATHLEN>

	basic.SniffEncode("value", value)

	if len(value) > V2_MNTPATHLEN {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Dirpath) Decode(d *xdr.Decoder) (err error) {
	var value string

	// value: string<V2_MNTPATHLEN>

	*(*string)(&value), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("value", value)

	if len(value) > V2_MNTPATHLEN {
		err = basic.ErrArrayTooLarge
		return
	}
	*t = Dirpath(value)
	return nil
}

// type mntname

func (t *Mntname) Encode(e *xdr.Encoder) (err error) {
	value := string(*t)

	// value: string<V2_MNTNAMLEN>

	basic.SniffEncode("value", value)

	if len(value) > V2_MNTNAMLEN {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(value))
	if err != nil {
		return
	}
	return nil
}

func (t *Mntname) Decode(d *xdr.Decoder) (err error) {
	var value string

	// value: string<V2_MNTNAMLEN>

	*(*string)(&value), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("value", value)

	if len(value) > V2_MNTNAMLEN {
		err = basic.ErrArrayTooLarge
		return
	}
	*t = Mntname(value)
	return nil
}

// struct groupnode

func (s *Groupnode) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Groupnode {")
	basic.SniffIndent()
	err = s.GrName.Encode(e)
	if err != nil {
		return
	}
	// s.GrNext: Optional
	{
		var notnull bool

		if s.GrNext != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.GrNext.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Groupnode)")
	return nil
}

func (s *Groupnode) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Groupnode {")
	basic.SniffIndent()
	err = s.GrName.Decode(d)
	if err != nil {
		return
	}
	// s.GrNext: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.GrNext = &Groupnode{}
			err = s.GrNext.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.GrNext = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Groupnode)")
	return nil
}

// struct exportnode

func (s *Exportnode) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Exportnode {")
	basic.SniffIndent()
	err = s.ExDir.Encode(e)
	if err != nil {
		return
	}
	// s.ExGroups: Optional
	{
		var notnull bool

		if s.ExGroups != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.ExGroups.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}
	// s.ExNext: Optional
	{
		var notnull bool

		if s.ExNext != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.ExNext.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Exportnode)")
	return nil
}

func (s *Exportnode) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Exportnode {")
	basic.SniffIndent()
	err = s.ExDir.Decode(d)
	if err != nil {
		return
	}
	// s.ExGroups: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.ExGroups = &Groupnode{}
			err = s.ExGroups.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.ExGroups = nil
		}
	}
	// s.ExNext: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.ExNext = &Exportnode{}
			err = s.ExNext.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.ExNext = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Exportnode)")
	return nil
}

// struct mountbody

func (s *Mountbody) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mountbody {")
	basic.SniffIndent()
	err = s.MlHostname.Encode(e)
	if err != nil {
		return
	}
	err = s.MlDirectory.Encode(e)
	if err != nil {
		return
	}
	// s.MlNext: Optional
	{
		var notnull bool

		if s.MlNext != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.MlNext.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mountbody)")
	return nil
}

func (s *Mountbody) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mountbody {")
	basic.SniffIndent()
	err = s.MlHostname.Decode(d)
	if err != nil {
		return
	}
	err = s.MlDirectory.Decode(d)
	if err != nil {
		return
	}
	// s.MlNext: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.MlNext = &Mountbody{}
			err = s.MlNext.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.MlNext = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mountbody)")
	return nil
}

// struct mountlist_first

func (s *MountlistFirst) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct MountlistFirst {")
	basic.SniffIndent()
	// s.First: Optional
	{
		var notnull bool

		if s.First != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.First.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct MountlistFirst)")
	return nil
}

func (s *MountlistFirst) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct MountlistFirst {")
	basic.SniffIndent()
	// s.First: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.First = &Mountbody{}
			err = s.First.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.First = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct MountlistFirst)")
	return nil
}

// struct exports_first

func (s *ExportsFirst) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct ExportsFirst {")
	basic.SniffIndent()
	// s.First: Optional
	{
		var notnull bool

		if s.First != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.First.Encode(e)
			if err != nil {
				return
			}

		} else {
			notnull = false
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct ExportsFirst)")
	return nil
}

func (s *ExportsFirst) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct ExportsFirst {")
	basic.SniffIndent()
	// s.First: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.First = &Exportnode{}
			err = s.First.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.First = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct ExportsFirst)")
	return nil
}

// struct mountres3_ok

func (s *Mountres3Ok) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mountres3Ok {")
	basic.SniffIndent()
	err = s.Fhandle.Encode(e)
	if err != nil {
		return
	}

	// s.AuthFlavors: []int32<MAX_ARRAY_LENGTH>
	{
		dataLength := uint32(len(s.AuthFlavors))

		basic.SniffEncode("dataLength", dataLength)

		if dataLength > MAX_ARRAY_LENGTH {
			err = basic.ErrArrayTooLarge
			return
		}

		_, err = e.EncodeUint(dataLength)
		if err != nil {
			return
		}

		for _, value := range s.AuthFlavors {
			basic.SniffEncode("value", value)
			_, err = e.EncodeInt(int32(value))
			if err != nil {
				return
			}

		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mountres3Ok)")
	return nil
}

func (s *Mountres3Ok) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mountres3Ok {")
	basic.SniffIndent()
	err = s.Fhandle.Decode(d)
	if err != nil {
		return
	}

	// s.AuthFlavors: []int32<MAX_ARRAY_LENGTH>
	{
		var dataLength uint32
		dataLength, _, err = d.DecodeUint()
		if err != nil {
			return
		}

		basic.SniffDecode("dataLength", dataLength)

		values := make([]int32, dataLength)
		for i := range values {
			values[i], _, err = d.DecodeInt()

			if err != nil {
				return
			}
			basic.SniffDecode("values[i]", values[i])
		}
		*(*[]int32)(&s.AuthFlavors) = values

		if dataLength > MAX_ARRAY_LENGTH {
			err = basic.ErrArrayTooLarge
			return
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mountres3Ok)")
	return nil
}

// union mountres3

func (s *Mountres3) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Mountres3 {")
	basic.SniffIndent()
	err = s.FhsStatus.Encode(e)
	if err != nil {
		return
	}
	switch s.FhsStatus {
	case MNT3_OK:
		u, ok := s.Union.(Mountres3Ok)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Mountres3)")
	return nil
}

func (s *Mountres3) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Mountres3 {")
	basic.SniffIndent()
	err = s.FhsStatus.Decode(d)
	if err != nil {
		return
	}
	switch s.FhsStatus {
	case MNT3_OK:
		u := new(Mountres3Ok)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Mountres3)")
	return nil
}

// type netobj

func (t *Netobj) Encode(e *xdr.Encoder) (err error) {
	value := [MAXNETOBJ_SZ]uint8(*t)

	basic.SniffEncode("value", value)
	_, err = e.EncodeFixedOpaque((value)[:])
	if err != nil {
		return
	}
	return nil
}

func (t *Netobj) Decode(d *xdr.Decoder) (err error) {
	var value [MAXNETOBJ_SZ]uint8

	{
		var bytes []byte
		bytes, _, err = d.DecodeFixedOpaque(int32(len(value)))
		if err != nil {
			return
		}
		copy(value[:], bytes)
		basic.SniffDecode("value", value)
	}
	*t = Netobj(value)
	return nil
}

// struct nlm4_stat

func (s *Nlm4Stat) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Stat {")
	basic.SniffIndent()
	err = s.Stat.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Stat)")
	return nil
}

func (s *Nlm4Stat) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Stat {")
	basic.SniffIndent()
	err = s.Stat.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Stat)")
	return nil
}

// struct nlm4_res

func (s *Nlm4Res) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Res {")
	basic.SniffIndent()
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	err = s.Stat.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Res)")
	return nil
}

func (s *Nlm4Res) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Res {")
	basic.SniffIndent()
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	err = s.Stat.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Res)")
	return nil
}

// struct nlm4_holder

func (s *Nlm4Holder) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Holder {")
	basic.SniffIndent()
	basic.SniffEncode("s.Exclusive", s.Exclusive)
	_, err = e.EncodeBool(bool(s.Exclusive))
	if err != nil {
		return
	}
	err = s.Svid.Encode(e)
	if err != nil {
		return
	}
	err = s.Oh.Encode(e)
	if err != nil {
		return
	}
	err = s.LOffset.Encode(e)
	if err != nil {
		return
	}
	err = s.LLen.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Holder)")
	return nil
}

func (s *Nlm4Holder) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Holder {")
	basic.SniffIndent()
	s.Exclusive, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Exclusive", s.Exclusive)
	err = s.Svid.Decode(d)
	if err != nil {
		return
	}
	err = s.Oh.Decode(d)
	if err != nil {
		return
	}
	err = s.LOffset.Decode(d)
	if err != nil {
		return
	}
	err = s.LLen.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Holder)")
	return nil
}

// union nlm4_testrply

func (s *Nlm4Testrply) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union Nlm4Testrply {")
	basic.SniffIndent()
	err = s.Stat.Encode(e)
	if err != nil {
		return
	}
	switch s.Stat {
	case NLM4_DENIED:
		u, ok := s.Union.(Nlm4Holder)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffEncode("} (union Nlm4Testrply)")
	return nil
}

func (s *Nlm4Testrply) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union Nlm4Testrply {")
	basic.SniffIndent()
	err = s.Stat.Decode(d)
	if err != nil {
		return
	}
	switch s.Stat {
	case NLM4_DENIED:
		u := new(Nlm4Holder)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)

	default:
		// Nothing to do here
	}
	basic.SniffUnindent()
	basic.SniffDecode("} (union Nlm4Testrply)")
	return nil
}

// struct nlm4_testres

func (s *Nlm4Testres) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Testres {")
	basic.SniffIndent()
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	err = s.TestStat.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Testres)")
	return nil
}

func (s *Nlm4Testres) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Testres {")
	basic.SniffIndent()
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	err = s.TestStat.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Testres)")
	return nil
}

// struct nlm4_lock

func (s *Nlm4Lock) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Lock {")
	basic.SniffIndent()

	// s.CallerName: string<LM_MAXSTRLEN>

	basic.SniffEncode("s.CallerName", s.CallerName)

	if len(s.CallerName) > LM_MAXSTRLEN {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(s.CallerName))
	if err != nil {
		return
	}
	err = s.Fh.Encode(e)
	if err != nil {
		return
	}
	err = s.Oh.Encode(e)
	if err != nil {
		return
	}
	err = s.Svid.Encode(e)
	if err != nil {
		return
	}
	err = s.LOffset.Encode(e)
	if err != nil {
		return
	}
	err = s.LLen.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Lock)")
	return nil
}

func (s *Nlm4Lock) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Lock {")
	basic.SniffIndent()

	// s.CallerName: string<LM_MAXSTRLEN>

	*(*string)(&s.CallerName), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("s.CallerName", s.CallerName)

	if len(s.CallerName) > LM_MAXSTRLEN {
		err = basic.ErrArrayTooLarge
		return
	}
	err = s.Fh.Decode(d)
	if err != nil {
		return
	}
	err = s.Oh.Decode(d)
	if err != nil {
		return
	}
	err = s.Svid.Decode(d)
	if err != nil {
		return
	}
	err = s.LOffset.Decode(d)
	if err != nil {
		return
	}
	err = s.LLen.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Lock)")
	return nil
}

// struct nlm4_lockargs

func (s *Nlm4Lockargs) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Lockargs {")
	basic.SniffIndent()
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Block", s.Block)
	_, err = e.EncodeBool(bool(s.Block))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Exclusive", s.Exclusive)
	_, err = e.EncodeBool(bool(s.Exclusive))
	if err != nil {
		return
	}
	err = s.Alock.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Reclaim", s.Reclaim)
	_, err = e.EncodeBool(bool(s.Reclaim))
	if err != nil {
		return
	}
	err = s.State.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Lockargs)")
	return nil
}

func (s *Nlm4Lockargs) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Lockargs {")
	basic.SniffIndent()
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	s.Block, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Block", s.Block)
	s.Exclusive, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Exclusive", s.Exclusive)
	err = s.Alock.Decode(d)
	if err != nil {
		return
	}
	s.Reclaim, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Reclaim", s.Reclaim)
	err = s.State.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Lockargs)")
	return nil
}

// struct nlm4_cancargs

func (s *Nlm4Cancargs) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Cancargs {")
	basic.SniffIndent()
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Block", s.Block)
	_, err = e.EncodeBool(bool(s.Block))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Exclusive", s.Exclusive)
	_, err = e.EncodeBool(bool(s.Exclusive))
	if err != nil {
		return
	}
	err = s.Alock.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Cancargs)")
	return nil
}

func (s *Nlm4Cancargs) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Cancargs {")
	basic.SniffIndent()
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	s.Block, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Block", s.Block)
	s.Exclusive, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Exclusive", s.Exclusive)
	err = s.Alock.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Cancargs)")
	return nil
}

// struct nlm4_testargs

func (s *Nlm4Testargs) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Testargs {")
	basic.SniffIndent()
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Exclusive", s.Exclusive)
	_, err = e.EncodeBool(bool(s.Exclusive))
	if err != nil {
		return
	}
	err = s.Alock.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Testargs)")
	return nil
}

func (s *Nlm4Testargs) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Testargs {")
	basic.SniffIndent()
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	s.Exclusive, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Exclusive", s.Exclusive)
	err = s.Alock.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Testargs)")
	return nil
}

// struct nlm4_unlockargs

func (s *Nlm4Unlockargs) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Unlockargs {")
	basic.SniffIndent()
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	err = s.Alock.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Unlockargs)")
	return nil
}

func (s *Nlm4Unlockargs) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Unlockargs {")
	basic.SniffIndent()
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	err = s.Alock.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Unlockargs)")
	return nil
}

// struct nlm4_share

func (s *Nlm4Share) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Share {")
	basic.SniffIndent()

	// s.CallerName: string<LM_MAXSTRLEN>

	basic.SniffEncode("s.CallerName", s.CallerName)

	if len(s.CallerName) > LM_MAXSTRLEN {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(s.CallerName))
	if err != nil {
		return
	}
	err = s.Fh.Encode(e)
	if err != nil {
		return
	}
	err = s.Oh.Encode(e)
	if err != nil {
		return
	}
	err = s.Mode.Encode(e)
	if err != nil {
		return
	}
	err = s.Access.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Share)")
	return nil
}

func (s *Nlm4Share) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Share {")
	basic.SniffIndent()

	// s.CallerName: string<LM_MAXSTRLEN>

	*(*string)(&s.CallerName), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("s.CallerName", s.CallerName)

	if len(s.CallerName) > LM_MAXSTRLEN {
		err = basic.ErrArrayTooLarge
		return
	}
	err = s.Fh.Decode(d)
	if err != nil {
		return
	}
	err = s.Oh.Decode(d)
	if err != nil {
		return
	}
	err = s.Mode.Decode(d)
	if err != nil {
		return
	}
	err = s.Access.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Share)")
	return nil
}

// struct nlm4_shareargs

func (s *Nlm4Shareargs) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Shareargs {")
	basic.SniffIndent()
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	err = s.Share.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Reclaim", s.Reclaim)
	_, err = e.EncodeBool(bool(s.Reclaim))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Shareargs)")
	return nil
}

func (s *Nlm4Shareargs) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Shareargs {")
	basic.SniffIndent()
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	err = s.Share.Decode(d)
	if err != nil {
		return
	}
	s.Reclaim, _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Reclaim", s.Reclaim)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Shareargs)")
	return nil
}

// struct nlm4_shareres

func (s *Nlm4Shareres) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Shareres {")
	basic.SniffIndent()
	err = s.Cookie.Encode(e)
	if err != nil {
		return
	}
	err = s.Stat.Encode(e)
	if err != nil {
		return
	}
	err = s.Sequence.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Shareres)")
	return nil
}

func (s *Nlm4Shareres) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Shareres {")
	basic.SniffIndent()
	err = s.Cookie.Decode(d)
	if err != nil {
		return
	}
	err = s.Stat.Decode(d)
	if err != nil {
		return
	}
	err = s.Sequence.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Shareres)")
	return nil
}

// struct nlm4_notify

func (s *Nlm4Notify) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4Notify {")
	basic.SniffIndent()

	// s.Name: string<LM_MAXNAMELEN>

	basic.SniffEncode("s.Name", s.Name)

	if len(s.Name) > LM_MAXNAMELEN {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(s.Name))
	if err != nil {
		return
	}
	err = s.State.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4Notify)")
	return nil
}

func (s *Nlm4Notify) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4Notify {")
	basic.SniffIndent()

	// s.Name: string<LM_MAXNAMELEN>

	*(*string)(&s.Name), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Name", s.Name)

	if len(s.Name) > LM_MAXNAMELEN {
		err = basic.ErrArrayTooLarge
		return
	}
	err = s.State.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4Notify)")
	return nil
}

// struct nlm4_sm_notifyargs

func (s *Nlm4SmNotifyargs) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Nlm4SmNotifyargs {")
	basic.SniffIndent()

	// s.Name: string<SM_MAXSTRLEN>

	basic.SniffEncode("s.Name", s.Name)

	if len(s.Name) > SM_MAXSTRLEN {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(s.Name))
	if err != nil {
		return
	}
	err = s.State.Encode(e)
	if err != nil {
		return
	}
	basic.SniffEncode("s.Priv", s.Priv)
	_, err = e.EncodeFixedOpaque((s.Priv)[:])
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Nlm4SmNotifyargs)")
	return nil
}

func (s *Nlm4SmNotifyargs) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Nlm4SmNotifyargs {")
	basic.SniffIndent()

	// s.Name: string<SM_MAXSTRLEN>

	*(*string)(&s.Name), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Name", s.Name)

	if len(s.Name) > SM_MAXSTRLEN {
		err = basic.ErrArrayTooLarge
		return
	}
	err = s.State.Decode(d)
	if err != nil {
		return
	}
	{
		var bytes []byte
		bytes, _, err = d.DecodeFixedOpaque(int32(len(s.Priv)))
		if err != nil {
			return
		}
		copy(s.Priv[:], bytes)
		basic.SniffDecode("s.Priv", s.Priv)
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Nlm4SmNotifyargs)")
	return nil
}

// RPC client for NfsV2

const ProgramNfsV2 = 100003

type NfsV2 struct {
	client *client.Client
}

func NewNfsV2(client *client.Client) *NfsV2 {
	return &NfsV2{
		client: client,
	}
}

func (p *NfsV2) ProcNull() (err error) {
	basic.SniffFunc("func NfsV2.ProcNull {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    0,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcNull)")

	return
}

func (p *NfsV2) ProcGetattr(params *Fhandle2) (result *Attr2res, err error) {
	basic.SniffFunc("func NfsV2.ProcGetattr {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    1,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Attr2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcGetattr)")

	return
}

func (p *NfsV2) ProcSetattr(params *Setattr2args) (result *Attr2res, err error) {
	basic.SniffFunc("func NfsV2.ProcSetattr {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    2,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Attr2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcSetattr)")

	return
}

func (p *NfsV2) ProcRoot() (err error) {
	basic.SniffFunc("func NfsV2.ProcRoot {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    3,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcRoot)")

	return
}

func (p *NfsV2) ProcLookup(params *Diropargs2) (result *Dirop2res, err error) {
	basic.SniffFunc("func NfsV2.ProcLookup {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    4,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Dirop2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcLookup)")

	return
}

func (p *NfsV2) ProcReadlink(params *Fhandle2) (result *Readlink2res, err error) {
	basic.SniffFunc("func NfsV2.ProcReadlink {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    5,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Readlink2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcReadlink)")

	return
}

func (p *NfsV2) ProcRead(params *Read2args) (result *Read2res, err error) {
	basic.SniffFunc("func NfsV2.ProcRead {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    6,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Read2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcRead)")

	return
}

func (p *NfsV2) ProcWritecache() (err error) {
	basic.SniffFunc("func NfsV2.ProcWritecache {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    7,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcWritecache)")

	return
}

func (p *NfsV2) ProcWrite(params *Write2args) (result *Attr2res, err error) {
	basic.SniffFunc("func NfsV2.ProcWrite {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    8,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Attr2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcWrite)")

	return
}

func (p *NfsV2) ProcCreate(params *Create2args) (result *Dirop2res, err error) {
	basic.SniffFunc("func NfsV2.ProcCreate {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    9,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Dirop2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcCreate)")

	return
}

func (p *NfsV2) ProcRemove(params *Diropargs2) (result *stat2, err error) {
	basic.SniffFunc("func NfsV2.ProcRemove {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    10,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(stat2)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcRemove)")

	return
}

func (p *NfsV2) ProcRename(params *Rename2args) (result *stat2, err error) {
	basic.SniffFunc("func NfsV2.ProcRename {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    11,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(stat2)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcRename)")

	return
}

func (p *NfsV2) ProcLink(params *Link2args) (result *stat2, err error) {
	basic.SniffFunc("func NfsV2.ProcLink {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    12,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(stat2)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcLink)")

	return
}

func (p *NfsV2) ProcSymlink(params *Symlink2args) (result *stat2, err error) {
	basic.SniffFunc("func NfsV2.ProcSymlink {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    13,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(stat2)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcSymlink)")

	return
}

func (p *NfsV2) ProcMkdir(params *Create2args) (result *Dirop2res, err error) {
	basic.SniffFunc("func NfsV2.ProcMkdir {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    14,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Dirop2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcMkdir)")

	return
}

func (p *NfsV2) ProcRmdir(params *Diropargs2) (result *stat2, err error) {
	basic.SniffFunc("func NfsV2.ProcRmdir {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    15,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(stat2)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcRmdir)")

	return
}

func (p *NfsV2) ProcReaddir(params *Readdir2args) (result *Readdir2res, err error) {
	basic.SniffFunc("func NfsV2.ProcReaddir {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    16,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Readdir2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcReaddir)")

	return
}

func (p *NfsV2) ProcStatfs(params *Fhandle2) (result *Statfs2res, err error) {
	basic.SniffFunc("func NfsV2.ProcStatfs {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    2,
				Proc:    17,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Statfs2res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV2.ProcStatfs)")

	return
}

// RPC client for NfsV3

const ProgramNfsV3 = 100003

type NfsV3 struct {
	Client *client.Client
}

func NewNfsV3(client *client.Client) *NfsV3 {
	return &NfsV3{
		Client: client,
	}
}

func (p *NfsV3) Proc3Null() (err error) {
	basic.SniffFunc("func NfsV3.Proc3Null {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    0,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Null)")

	return
}

func (p *NfsV3) Proc3Getattr(params *Getattr3args) (result *Getattr3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Getattr {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    1,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Getattr3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Getattr)")

	return
}

func (p *NfsV3) Proc3Setattr(params *Setattr3args) (result *Setattr3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Setattr {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    2,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Setattr3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Setattr)")

	return
}

func (p *NfsV3) Proc3Lookup(params *Lookup3args) (result *Lookup3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Lookup {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    3,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Lookup3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Lookup)")

	return
}

func (p *NfsV3) Proc3Access(params *Access3args) (result *Access3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Access {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    4,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Access3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Access)")

	return
}

func (p *NfsV3) Proc3Readlink(params *Readlink3args) (result *Readlink3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Readlink {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    5,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Readlink3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Readlink)")

	return
}

func (p *NfsV3) Proc3Read(params *Read3args) (result *Read3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Read {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    6,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Read3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Read)")

	return
}

func (p *NfsV3) Proc3Write(params *Write3args) (result *Write3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Write {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    7,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Write3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Write)")

	return
}

func (p *NfsV3) Proc3Create(params *Create3args) (result *Create3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Create {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    8,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	/*
		// Simulate EL-462 issue
		if len(params.Where.Name) > 255 {
			request.InjectResponseDelay(60 * time.Second)
		}
	*/

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Create3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Create)")

	return
}

func (p *NfsV3) Proc3Mkdir(params *Mkdir3args) (result *Mkdir3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Mkdir {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    9,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Mkdir3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Mkdir)")

	return
}

func (p *NfsV3) Proc3Symlink(params *Symlink3args) (result *Symlink3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Symlink {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    10,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Symlink3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Symlink)")

	return
}

func (p *NfsV3) Proc3Mknod(params *Mknod3args) (result *Mknod3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Mknod {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    11,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Mknod3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Mknod)")

	return
}

func (p *NfsV3) Proc3Remove(params *Remove3args) (result *Remove3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Remove {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    12,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Remove3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Remove)")

	return
}

func (p *NfsV3) Proc3Rmdir(params *Rmdir3args) (result *Rmdir3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Rmdir {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    13,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Rmdir3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Rmdir)")

	return
}

func (p *NfsV3) Proc3Rename(params *Rename3args) (result *Rename3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Rename {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    14,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Rename3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Rename)")

	return
}

func (p *NfsV3) Proc3Link(params *Link3args) (result *Link3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Link {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    15,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Link3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Link)")

	return
}

func (p *NfsV3) Proc3Readdir(params *Readdir3args) (result *Readdir3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Readdir {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    16,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Readdir3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Readdir)")

	return
}

func (p *NfsV3) Proc3Readdirplus(params *Readdirplus3args) (result *Readdirplus3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Readdirplus {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    17,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Readdirplus3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Readdirplus)")

	return
}

func (p *NfsV3) Proc3Fsstat(params *Fsstat3args) (result *Fsstat3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Fsstat {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    18,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Fsstat3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Fsstat)")

	return
}

func (p *NfsV3) Proc3Fsinfo(params *Fsinfo3args) (result *Fsinfo3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Fsinfo {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    19,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Fsinfo3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Fsinfo)")

	return
}

func (p *NfsV3) Proc3Pathconf(params *Pathconf3args) (result *Pathconf3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Pathconf {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    20,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Pathconf3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Pathconf)")

	return
}

func (p *NfsV3) Proc3Commit(params *Commit3args) (result *Commit3res, err error) {
	basic.SniffFunc("func NfsV3.Proc3Commit {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100003,
				Vers:    3,
				Proc:    21,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Commit3res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func NfsV3.Proc3Commit)")

	return
}

// RPC client for MountprogmountV1

const ProgramMountprogmountV1 = 100005

type MountprogmountV1 struct {
	client *client.Client
}

func NewMountprogmountV1(client *client.Client) *MountprogmountV1 {
	return &MountprogmountV1{
		client: client,
	}
}

func (p *MountprogmountV1) Mountproc2Null() (err error) {
	basic.SniffFunc("func MountprogmountV1.Mountproc2Null {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    1,
				Proc:    0,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV1.Mountproc2Null)")

	return
}

func (p *MountprogmountV1) Mountproc2Mnt(params *Dirpath) (result *Fhstatus2, err error) {
	basic.SniffFunc("func MountprogmountV1.Mountproc2Mnt {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    1,
				Proc:    1,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Fhstatus2)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV1.Mountproc2Mnt)")

	return
}

func (p *MountprogmountV1) Mountproc2Dump() (result *MountlistFirst, err error) {
	basic.SniffFunc("func MountprogmountV1.Mountproc2Dump {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    1,
				Proc:    2,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(MountlistFirst)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV1.Mountproc2Dump)")

	return
}

func (p *MountprogmountV1) Mountproc2Umnt(params *Dirpath) (err error) {
	basic.SniffFunc("func MountprogmountV1.Mountproc2Umnt {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    1,
				Proc:    3,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV1.Mountproc2Umnt)")

	return
}

func (p *MountprogmountV1) Mountproc2Umntall() (err error) {
	basic.SniffFunc("func MountprogmountV1.Mountproc2Umntall {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    1,
				Proc:    4,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV1.Mountproc2Umntall)")

	return
}

func (p *MountprogmountV1) Mountproc2Export() (result *ExportsFirst, err error) {
	basic.SniffFunc("func MountprogmountV1.Mountproc2Export {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    1,
				Proc:    5,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(ExportsFirst)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV1.Mountproc2Export)")

	return
}

// RPC client for MountprogmountV3

const ProgramMountprogmountV3 = 100005

type MountprogmountV3 struct {
	Client *client.Client
}

func NewMountprogmountV3(client *client.Client) *MountprogmountV3 {
	return &MountprogmountV3{
		Client: client,
	}
}

func (p *MountprogmountV3) Mountproc3Null() (err error) {
	basic.SniffFunc("func MountprogmountV3.Mountproc3Null {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    3,
				Proc:    0,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV3.Mountproc3Null)")

	return
}

func (p *MountprogmountV3) Mountproc3Mnt(params *Dirpath) (result *Mountres3, err error) {
	basic.SniffFunc("func MountprogmountV3.Mountproc3Mnt {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    3,
				Proc:    1,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Mountres3)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV3.Mountproc3Mnt)")

	return
}

func (p *MountprogmountV3) Mountproc3Dump() (result *MountlistFirst, err error) {
	basic.SniffFunc("func MountprogmountV3.Mountproc3Dump {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    3,
				Proc:    2,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(MountlistFirst)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV3.Mountproc3Dump)")

	return
}

func (p *MountprogmountV3) Mountproc3Umnt(params *Dirpath) (err error) {
	basic.SniffFunc("func MountprogmountV3.Mountproc3Umnt {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    3,
				Proc:    3,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV3.Mountproc3Umnt)")

	return
}

func (p *MountprogmountV3) Mountproc3Umntall() (err error) {
	basic.SniffFunc("func MountprogmountV3.Mountproc3Umntall {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    3,
				Proc:    4,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV3.Mountproc3Umntall)")

	return
}

func (p *MountprogmountV3) Mountproc3Export() (result *ExportsFirst, err error) {
	basic.SniffFunc("func MountprogmountV3.Mountproc3Export {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100005,
				Vers:    3,
				Proc:    5,
				Cred:    p.Client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.Client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(ExportsFirst)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func MountprogmountV3.Mountproc3Export)")

	return
}

// RPC client for Nlmprognlm4Vers

const ProgramNlmprognlm4Vers = 100021

type Nlmprognlm4Vers struct {
	client *client.Client
}

func NewNlmprognlm4Vers(client *client.Client) *Nlmprognlm4Vers {
	return &Nlmprognlm4Vers{
		client: client,
	}
}

func (p *Nlmprognlm4Vers) Nlmproc4Null() (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4Null {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    0,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// No params to encode

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4Null)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4Test(params *Nlm4Testargs) (result *Nlm4Testres, err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4Test {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    1,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Nlm4Testres)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4Test)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4Lock(params *Nlm4Lockargs) (result *Nlm4Res, err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4Lock {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    2,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Nlm4Res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4Lock)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4Cancel(params *Nlm4Cancargs) (result *Nlm4Res, err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4Cancel {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    3,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Nlm4Res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4Cancel)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4Unlock(params *Nlm4Unlockargs) (result *Nlm4Res, err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4Unlock {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    4,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Nlm4Res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4Unlock)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4Granted(params *Nlm4Testargs) (result *Nlm4Res, err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4Granted {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    5,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Nlm4Res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4Granted)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4TestMsg(params *Nlm4Testargs) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4TestMsg {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    6,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4TestMsg)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4LockMsg(params *Nlm4Lockargs) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4LockMsg {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    7,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4LockMsg)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4CancelMsg(params *Nlm4Cancargs) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4CancelMsg {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    8,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4CancelMsg)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4UnlockMsg(params *Nlm4Unlockargs) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4UnlockMsg {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    9,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4UnlockMsg)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4GrantedMsg(params *Nlm4Testargs) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4GrantedMsg {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    10,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4GrantedMsg)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4TestRes(params *Nlm4Testres) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4TestRes {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    11,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4TestRes)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4LockRes(params *Nlm4Res) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4LockRes {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    12,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4LockRes)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4CancelRes(params *Nlm4Res) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4CancelRes {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    13,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4CancelRes)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4UnlockRes(params *Nlm4Res) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4UnlockRes {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    14,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4UnlockRes)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4GrantedRes(params *Nlm4Res) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4GrantedRes {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    15,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4GrantedRes)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4SmNotify(params *Nlm4SmNotifyargs) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4SmNotify {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    16,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4SmNotify)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4Share(params *Nlm4Shareargs) (result *Nlm4Shareres, err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4Share {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    20,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Nlm4Shareres)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4Share)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4Unshare(params *Nlm4Shareargs) (result *Nlm4Shareres, err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4Unshare {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    21,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Nlm4Shareres)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4Unshare)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4NmLock(params *Nlm4Lockargs) (result *Nlm4Res, err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4NmLock {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    22,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// Decode result
	d := response.Decoder
	result = new(Nlm4Res)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4NmLock)")

	return
}

func (p *Nlmprognlm4Vers) Nlmproc4FreeAll(params *Nlm4Notify) (err error) {
	basic.SniffFunc("func Nlmprognlm4Vers.Nlmproc4FreeAll {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100021,
				Vers:    4,
				Proc:    23,
				Cred:    p.client.AuthData,
				Verf:    rpc2.OpaqueAuth{Flavor: rpc2.AUTH_NULL},
			},
		},
	}

	// Create request and encode RPC message
	request, err := p.client.NewRequest(message)
	if err != nil {
		return
	}

	// Encode params
	e := request.Encoder
	err = (*params).Encode(e)
	if err != nil {
		return
	}

	// Wait for response
	response, err := request.SendAndWaitForResponse()
	if err != nil {
		return
	}

	// No result to decode
	_ = response

	basic.SniffUnindent()
	basic.SniffFunc("} (func Nlmprognlm4Vers.Nlmproc4FreeAll)")

	return
}
