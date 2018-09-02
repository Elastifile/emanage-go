// package rpc2 (type declarations) is AUTO GENERATED from rpc2.x by sunrpc_maker

package rpc2

import "github.com/elastifile/emanage-go/src/nfs/sunrpc/basic"

const MAX_ARRAY_LENGTH = 0x7fffffff

////////////////////////////////
// Enums

// enum msg_type

type msgType basic.Enum

const (
	CALL  msgType = 0
	REPLY msgType = 1
)

// enum reply_stat

type replyStat basic.Enum

const (
	MSG_ACCEPTED replyStat = 0
	MSG_DENIED   replyStat = 1
)

// enum accept_stat

type acceptStat basic.Enum

const (
	SUCCESS       acceptStat = 0
	PROG_UNAVAIL  acceptStat = 1
	PROG_MISMATCH acceptStat = 2
	PROC_UNAVAIL  acceptStat = 3
	GARBAGE_ARGS  acceptStat = 4
)

// enum reject_stat

type rejectStat basic.Enum

const (
	RPC_MISMATCH rejectStat = 0
	AUTH_ERROR   rejectStat = 1
)

// enum auth_stat

type authStat basic.Enum

const (
	AUTH_BADCRED      authStat = 1
	AUTH_REJECTEDCRED authStat = 2
	AUTH_BADVERF      authStat = 3
	AUTH_REJECTEDVERF authStat = 4
	AUTH_TOOWEAK      authStat = 5
)

// enum auth_flavor

type authFlavor basic.Enum

const (
	AUTH_NULL  authFlavor = 0
	AUTH_UNIX  authFlavor = 1
	AUTH_SHORT authFlavor = 2
	AUTH_DES   authFlavor = 3
)

////////////////////////////////
// Structs

// struct opaque_auth

type OpaqueAuth struct {
	Flavor authFlavor
	Body   []uint8 // Max length: 400
}

// struct call_body

type CallBody struct {
	Rpcvers uint32
	Prog    uint32
	Vers    uint32
	Proc    uint32
	Cred    OpaqueAuth
	Verf    OpaqueAuth
}

// struct mismatch_info_body

type MismatchInfoBody struct {
	Low  uint32
	High uint32
}

// union reply_data_body

type ReplyDataBody struct {
	Stat  acceptStat // Arbitrator
	Union replyDataBodyUnion
}

type replyDataBodyUnion interface {
	isReplyDataBodyUnion()
}

func (MismatchInfoBody) isReplyDataBodyUnion() {}

// struct accepted_reply

type AcceptedReply struct {
	Verf      OpaqueAuth
	ReplyData ReplyDataBody
}

// union rejected_reply

type RejectedReply struct {
	Stat  rejectStat // Arbitrator
	Union rejectedReplyUnion
}

type rejectedReplyUnion interface {
	isRejectedReplyUnion()
}

func (MismatchInfoBody) isRejectedReplyUnion() {}

func (authStat) isRejectedReplyUnion() {}

// union reply_body

type ReplyBody struct {
	Stat  replyStat // Arbitrator
	Union replyBodyUnion
}

type replyBodyUnion interface {
	isReplyBodyUnion()
}

func (AcceptedReply) isReplyBodyUnion() {}

func (RejectedReply) isReplyBodyUnion() {}

// union rpc_body

type RpcBody struct {
	Mtype msgType // Arbitrator
	Union rpcBodyUnion
}

type rpcBodyUnion interface {
	isRpcBodyUnion()
}

func (CallBody) isRpcBodyUnion() {}

func (ReplyBody) isRpcBodyUnion() {}

// struct rpc_msg

type RpcMsg struct {
	Xid  uint32
	Body RpcBody
}

// struct auth_unix

type AuthUnix struct {
	Stamp       uint32
	Machinename string // Max length: 255
	Uid         uint32
	Gid         uint32
	Gids        []uint32 // Max length: 16
}
