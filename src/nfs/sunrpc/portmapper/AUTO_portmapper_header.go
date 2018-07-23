// package portmapper (type declarations) is AUTO GENERATED from portmapper.x by sunrpc_maker

package portmapper

const MAX_ARRAY_LENGTH = 0x7fffffff

// Global constants

const (
	PMAP_PORT   = 111
	IPPROTO_TCP = 6
	IPPROTO_UDP = 17
)

////////////////////////////////
// Enums

////////////////////////////////
// Structs

// struct mapping

type Mapping struct {
	Prog uint32
	Vers uint32
	Prot uint32
	Port uint32
}

// struct pmaplist

type Pmaplist struct {
	Map  Mapping
	Next *Pmaplist
}

// struct pmaplist_first

type PmaplistFirst struct {
	Next *Pmaplist
}

// struct call_args

type CallArgs struct {
	Prog uint32
	Vers uint32
	Proc uint32
	Args []uint8
}

// struct call_result

type CallResult struct {
	Port uint32
	Res  []uint8
}
