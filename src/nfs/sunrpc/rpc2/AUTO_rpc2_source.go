// package rpc2 (implementation) is AUTO GENERATED from rpc2.x by sunrpc_maker
package rpc2

import (
	"fmt"

	xdr "github.com/davecgh/go-xdr/xdr2"
)
import "nfs/sunrpc/basic"

// enum msg_type

func (value msgType) FieldCheck() error {
	switch value {
	case CALL:
	case REPLY:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value msgType) String() string {
	switch value {
	case CALL:
		return "CALL"
	case REPLY:
		return "REPLY"
	}
	return fmt.Sprintf("msgType(%d)", value)
}

func (value *msgType) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("msgType", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *msgType) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = msgType(intValue)
	basic.SniffDecode("msgType", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum reply_stat

func (value replyStat) FieldCheck() error {
	switch value {
	case MSG_ACCEPTED:
	case MSG_DENIED:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value replyStat) String() string {
	switch value {
	case MSG_ACCEPTED:
		return "MSG_ACCEPTED"
	case MSG_DENIED:
		return "MSG_DENIED"
	}
	return fmt.Sprintf("replyStat(%d)", value)
}

func (value *replyStat) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("replyStat", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *replyStat) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = replyStat(intValue)
	basic.SniffDecode("replyStat", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum accept_stat

func (value acceptStat) FieldCheck() error {
	switch value {
	case SUCCESS:
	case PROG_UNAVAIL:
	case PROG_MISMATCH:
	case PROC_UNAVAIL:
	case GARBAGE_ARGS:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value acceptStat) String() string {
	switch value {
	case SUCCESS:
		return "SUCCESS"
	case PROG_UNAVAIL:
		return "PROG_UNAVAIL"
	case PROG_MISMATCH:
		return "PROG_MISMATCH"
	case PROC_UNAVAIL:
		return "PROC_UNAVAIL"
	case GARBAGE_ARGS:
		return "GARBAGE_ARGS"
	}
	return fmt.Sprintf("acceptStat(%d)", value)
}

func (value *acceptStat) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("acceptStat", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *acceptStat) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = acceptStat(intValue)
	basic.SniffDecode("acceptStat", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum reject_stat

func (value rejectStat) FieldCheck() error {
	switch value {
	case RPC_MISMATCH:
	case AUTH_ERROR:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value rejectStat) String() string {
	switch value {
	case RPC_MISMATCH:
		return "RPC_MISMATCH"
	case AUTH_ERROR:
		return "AUTH_ERROR"
	}
	return fmt.Sprintf("rejectStat(%d)", value)
}

func (value *rejectStat) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("rejectStat", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *rejectStat) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = rejectStat(intValue)
	basic.SniffDecode("rejectStat", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum auth_stat

func (value authStat) FieldCheck() error {
	switch value {
	case AUTH_BADCRED:
	case AUTH_REJECTEDCRED:
	case AUTH_BADVERF:
	case AUTH_REJECTEDVERF:
	case AUTH_TOOWEAK:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value authStat) String() string {
	switch value {
	case AUTH_BADCRED:
		return "AUTH_BADCRED"
	case AUTH_REJECTEDCRED:
		return "AUTH_REJECTEDCRED"
	case AUTH_BADVERF:
		return "AUTH_BADVERF"
	case AUTH_REJECTEDVERF:
		return "AUTH_REJECTEDVERF"
	case AUTH_TOOWEAK:
		return "AUTH_TOOWEAK"
	}
	return fmt.Sprintf("authStat(%d)", value)
}

func (value *authStat) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("authStat", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *authStat) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = authStat(intValue)
	basic.SniffDecode("authStat", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// enum auth_flavor

func (value authFlavor) FieldCheck() error {
	switch value {
	case AUTH_NULL:
	case AUTH_UNIX:
	case AUTH_SHORT:
	case AUTH_DES:
	default:
		return basic.ErrInvalidEnumOpt
	}
	return nil
}

func (value authFlavor) String() string {
	switch value {
	case AUTH_NULL:
		return "AUTH_NULL"
	case AUTH_UNIX:
		return "AUTH_UNIX"
	case AUTH_SHORT:
		return "AUTH_SHORT"
	case AUTH_DES:
		return "AUTH_DES"
	}
	return fmt.Sprintf("authFlavor(%d)", value)
}

func (value *authFlavor) Encode(e *xdr.Encoder) (err error) {
	err = value.FieldCheck()
	if err != nil {
		return
	}

	basic.SniffEncode("authFlavor", *value)
	_, err = e.EncodeInt(int32(*value))
	if err != nil {
		return
	}

	return nil
}

func (value *authFlavor) Decode(d *xdr.Decoder) (err error) {
	var intValue int32

	intValue, _, err = d.DecodeInt()
	if err != nil {
		return
	}

	*value = authFlavor(intValue)
	basic.SniffDecode("authFlavor", *value)

	err = value.FieldCheck()
	if err != nil {
		return
	}

	return nil
}

// struct opaque_auth

func (s *OpaqueAuth) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct OpaqueAuth {")
	basic.SniffIndent()
	err = s.Flavor.Encode(e)
	if err != nil {
		return
	}

	// s.Body: []uint8<400>

	basic.SniffEncode("s.Body", s.Body)

	if len(s.Body) > 400 {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeOpaque(s.Body)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct OpaqueAuth)")
	return nil
}

func (s *OpaqueAuth) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct OpaqueAuth {")
	basic.SniffIndent()
	err = s.Flavor.Decode(d)
	if err != nil {
		return
	}

	// s.Body: []uint8<400>

	s.Body, _, err = d.DecodeOpaque()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Body", s.Body)

	if len(s.Body) > 400 {
		err = basic.ErrArrayTooLarge
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct OpaqueAuth)")
	return nil
}

// struct call_body

func (s *CallBody) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct CallBody {")
	basic.SniffIndent()
	basic.SniffEncode("s.Rpcvers", s.Rpcvers)
	_, err = e.EncodeUint(uint32(s.Rpcvers))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Prog", s.Prog)
	_, err = e.EncodeUint(uint32(s.Prog))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Vers", s.Vers)
	_, err = e.EncodeUint(uint32(s.Vers))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Proc", s.Proc)
	_, err = e.EncodeUint(uint32(s.Proc))
	if err != nil {
		return
	}
	err = s.Cred.Encode(e)
	if err != nil {
		return
	}
	err = s.Verf.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct CallBody)")
	return nil
}

func (s *CallBody) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct CallBody {")
	basic.SniffIndent()
	s.Rpcvers, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Rpcvers", s.Rpcvers)
	s.Prog, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Prog", s.Prog)
	s.Vers, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Vers", s.Vers)
	s.Proc, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Proc", s.Proc)
	err = s.Cred.Decode(d)
	if err != nil {
		return
	}
	err = s.Verf.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct CallBody)")
	return nil
}

// struct mismatch_info_body

func (s *MismatchInfoBody) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct MismatchInfoBody {")
	basic.SniffIndent()
	basic.SniffEncode("s.Low", s.Low)
	_, err = e.EncodeUint(uint32(s.Low))
	if err != nil {
		return
	}
	basic.SniffEncode("s.High", s.High)
	_, err = e.EncodeUint(uint32(s.High))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct MismatchInfoBody)")
	return nil
}

func (s *MismatchInfoBody) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct MismatchInfoBody {")
	basic.SniffIndent()
	s.Low, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Low", s.Low)
	s.High, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.High", s.High)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct MismatchInfoBody)")
	return nil
}

// union reply_data_body

func (s *ReplyDataBody) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union ReplyDataBody {")
	basic.SniffIndent()
	err = s.Stat.Encode(e)
	if err != nil {
		return
	}
	switch s.Stat {
	case SUCCESS:
		// Empty
	case PROG_MISMATCH:
		u, ok := s.Union.(MismatchInfoBody)
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
	basic.SniffEncode("} (union ReplyDataBody)")
	return nil
}

func (s *ReplyDataBody) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union ReplyDataBody {")
	basic.SniffIndent()
	err = s.Stat.Decode(d)
	if err != nil {
		return
	}
	switch s.Stat {
	case SUCCESS:
		// Empty
	case PROG_MISMATCH:
		u := new(MismatchInfoBody)
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
	basic.SniffDecode("} (union ReplyDataBody)")
	return nil
}

// struct accepted_reply

func (s *AcceptedReply) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct AcceptedReply {")
	basic.SniffIndent()
	err = s.Verf.Encode(e)
	if err != nil {
		return
	}
	err = s.ReplyData.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct AcceptedReply)")
	return nil
}

func (s *AcceptedReply) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct AcceptedReply {")
	basic.SniffIndent()
	err = s.Verf.Decode(d)
	if err != nil {
		return
	}
	err = s.ReplyData.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct AcceptedReply)")
	return nil
}

// union rejected_reply

func (s *RejectedReply) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union RejectedReply {")
	basic.SniffIndent()
	err = s.Stat.Encode(e)
	if err != nil {
		return
	}
	switch s.Stat {
	case RPC_MISMATCH:
		u, ok := s.Union.(MismatchInfoBody)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case AUTH_ERROR:
		u, ok := s.Union.(authStat)
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
	basic.SniffEncode("} (union RejectedReply)")
	return nil
}

func (s *RejectedReply) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union RejectedReply {")
	basic.SniffIndent()
	err = s.Stat.Decode(d)
	if err != nil {
		return
	}
	switch s.Stat {
	case RPC_MISMATCH:
		u := new(MismatchInfoBody)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case AUTH_ERROR:
		u := new(authStat)
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
	basic.SniffDecode("} (union RejectedReply)")
	return nil
}

// union reply_body

func (s *ReplyBody) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union ReplyBody {")
	basic.SniffIndent()
	err = s.Stat.Encode(e)
	if err != nil {
		return
	}
	switch s.Stat {
	case MSG_ACCEPTED:
		u, ok := s.Union.(AcceptedReply)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case MSG_DENIED:
		u, ok := s.Union.(RejectedReply)
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
	basic.SniffEncode("} (union ReplyBody)")
	return nil
}

func (s *ReplyBody) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union ReplyBody {")
	basic.SniffIndent()
	err = s.Stat.Decode(d)
	if err != nil {
		return
	}
	switch s.Stat {
	case MSG_ACCEPTED:
		u := new(AcceptedReply)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case MSG_DENIED:
		u := new(RejectedReply)
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
	basic.SniffDecode("} (union ReplyBody)")
	return nil
}

// union rpc_body

func (s *RpcBody) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("union RpcBody {")
	basic.SniffIndent()
	err = s.Mtype.Encode(e)
	if err != nil {
		return
	}
	switch s.Mtype {
	case CALL:
		u, ok := s.Union.(CallBody)
		if !ok {
			return basic.ErrArbitratorValueMismatch
		}
		err = u.Encode(e)
		if err != nil {
			return
		}
	case REPLY:
		u, ok := s.Union.(ReplyBody)
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
	basic.SniffEncode("} (union RpcBody)")
	return nil
}

func (s *RpcBody) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("union RpcBody {")
	basic.SniffIndent()
	err = s.Mtype.Decode(d)
	if err != nil {
		return
	}
	switch s.Mtype {
	case CALL:
		u := new(CallBody)
		err = u.Decode(d)
		if err != nil {
			return
		}

		s.Union = *u
		basic.SniffDecode("s.Union", s.Union)
	case REPLY:
		u := new(ReplyBody)
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
	basic.SniffDecode("} (union RpcBody)")
	return nil
}

// struct rpc_msg

func (s *RpcMsg) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct RpcMsg {")
	basic.SniffIndent()
	basic.SniffEncode("s.Xid", s.Xid)
	_, err = e.EncodeUint(uint32(s.Xid))
	if err != nil {
		return
	}
	err = s.Body.Encode(e)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct RpcMsg)")
	return nil
}

func (s *RpcMsg) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct RpcMsg {")
	basic.SniffIndent()
	s.Xid, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Xid", s.Xid)
	err = s.Body.Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct RpcMsg)")
	return nil
}

// struct auth_unix

func (s *AuthUnix) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct AuthUnix {")
	basic.SniffIndent()
	basic.SniffEncode("s.Stamp", s.Stamp)
	_, err = e.EncodeUint(uint32(s.Stamp))
	if err != nil {
		return
	}

	// s.Machinename: string<255>

	basic.SniffEncode("s.Machinename", s.Machinename)

	if len(s.Machinename) > 255 {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeString(string(s.Machinename))
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

	// s.Gids: []uint32<16>
	{
		dataLength := uint32(len(s.Gids))

		basic.SniffEncode("dataLength", dataLength)

		if dataLength > 16 {
			err = basic.ErrArrayTooLarge
			return
		}

		_, err = e.EncodeUint(dataLength)
		if err != nil {
			return
		}

		for _, value := range s.Gids {
			basic.SniffEncode("value", value)
			_, err = e.EncodeUint(uint32(value))
			if err != nil {
				return
			}

		}
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct AuthUnix)")
	return nil
}

func (s *AuthUnix) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct AuthUnix {")
	basic.SniffIndent()
	s.Stamp, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Stamp", s.Stamp)

	// s.Machinename: string<255>

	*(*string)(&s.Machinename), _, err = d.DecodeString()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Machinename", s.Machinename)

	if len(s.Machinename) > 255 {
		err = basic.ErrArrayTooLarge
		return
	}
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

	// s.Gids: []uint32<16>
	{
		var dataLength uint32
		dataLength, _, err = d.DecodeUint()
		if err != nil {
			return
		}

		basic.SniffDecode("dataLength", dataLength)

		values := make([]uint32, dataLength)
		for i := range values {
			values[i], _, err = d.DecodeUint()

			if err != nil {
				return
			}
			basic.SniffDecode("values[i]", values[i])
		}
		*(*[]uint32)(&s.Gids) = values

		if dataLength > 16 {
			err = basic.ErrArrayTooLarge
			return
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct AuthUnix)")
	return nil
}
