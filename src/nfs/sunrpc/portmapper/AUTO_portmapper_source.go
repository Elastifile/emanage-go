// package portmapper (implementation) is AUTO GENERATED from portmapper.x by sunrpc_maker
package portmapper

import xdr "github.com/davecgh/go-xdr/xdr2"
import (
	"nfs/sunrpc/basic"
	"nfs/sunrpc/client"
	"nfs/sunrpc/rpc2"
)

// struct mapping

func (s *Mapping) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Mapping {")
	basic.SniffIndent()
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
	basic.SniffEncode("s.Prot", s.Prot)
	_, err = e.EncodeUint(uint32(s.Prot))
	if err != nil {
		return
	}
	basic.SniffEncode("s.Port", s.Port)
	_, err = e.EncodeUint(uint32(s.Port))
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct Mapping)")
	return nil
}

func (s *Mapping) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Mapping {")
	basic.SniffIndent()
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
	s.Prot, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Prot", s.Prot)
	s.Port, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Port", s.Port)

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Mapping)")
	return nil
}

// struct pmaplist

func (s *Pmaplist) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct Pmaplist {")
	basic.SniffIndent()
	err = s.Map.Encode(e)
	if err != nil {
		return
	}
	// s.Next: Optional
	{
		var notnull bool

		if s.Next != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.Next.Encode(e)
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
	basic.SniffEncode("} (struct Pmaplist)")
	return nil
}

func (s *Pmaplist) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct Pmaplist {")
	basic.SniffIndent()
	err = s.Map.Decode(d)
	if err != nil {
		return
	}
	// s.Next: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.Next = &Pmaplist{}
			err = s.Next.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.Next = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct Pmaplist)")
	return nil
}

// struct pmaplist_first

func (s *PmaplistFirst) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct PmaplistFirst {")
	basic.SniffIndent()
	// s.Next: Optional
	{
		var notnull bool

		if s.Next != nil {
			notnull = true
			basic.SniffEncode("notnull", notnull)
			_, err = e.EncodeBool(notnull)
			if err != nil {
				return
			}
			err = s.Next.Encode(e)
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
	basic.SniffEncode("} (struct PmaplistFirst)")
	return nil
}

func (s *PmaplistFirst) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct PmaplistFirst {")
	basic.SniffIndent()
	// s.Next: Optional
	{
		var notnull bool

		notnull, _, err = d.DecodeBool()
		if err != nil {
			return
		}
		basic.SniffDecode("notnull", notnull)

		if notnull {
			s.Next = &Pmaplist{}
			err = s.Next.Decode(d)
			if err != nil {
				return
			}
		} else {
			s.Next = nil
		}
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct PmaplistFirst)")
	return nil
}

// struct call_args

func (s *CallArgs) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct CallArgs {")
	basic.SniffIndent()
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

	// s.Args: []uint8<MAX_ARRAY_LENGTH>

	basic.SniffEncode("s.Args", s.Args)

	if len(s.Args) > MAX_ARRAY_LENGTH {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeOpaque(s.Args)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct CallArgs)")
	return nil
}

func (s *CallArgs) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct CallArgs {")
	basic.SniffIndent()
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

	// s.Args: []uint8<MAX_ARRAY_LENGTH>

	s.Args, _, err = d.DecodeOpaque()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Args", s.Args)

	if len(s.Args) > MAX_ARRAY_LENGTH {
		err = basic.ErrArrayTooLarge
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct CallArgs)")
	return nil
}

// struct call_result

func (s *CallResult) Encode(e *xdr.Encoder) (err error) {
	basic.SniffEncode("struct CallResult {")
	basic.SniffIndent()
	basic.SniffEncode("s.Port", s.Port)
	_, err = e.EncodeUint(uint32(s.Port))
	if err != nil {
		return
	}

	// s.Res: []uint8<MAX_ARRAY_LENGTH>

	basic.SniffEncode("s.Res", s.Res)

	if len(s.Res) > MAX_ARRAY_LENGTH {
		err = basic.ErrArrayTooLarge
		return
	}

	_, err = e.EncodeOpaque(s.Res)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffEncode("} (struct CallResult)")
	return nil
}

func (s *CallResult) Decode(d *xdr.Decoder) (err error) {
	basic.SniffDecode("struct CallResult {")
	basic.SniffIndent()
	s.Port, _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("s.Port", s.Port)

	// s.Res: []uint8<MAX_ARRAY_LENGTH>

	s.Res, _, err = d.DecodeOpaque()
	if err != nil {
		return
	}

	basic.SniffDecode("s.Res", s.Res)

	if len(s.Res) > MAX_ARRAY_LENGTH {
		err = basic.ErrArrayTooLarge
		return
	}

	basic.SniffUnindent()
	basic.SniffDecode("} (struct CallResult)")
	return nil
}

// RPC client for PmapV2

const ProgramPmapV2 = 100000

type PmapV2 struct {
	Client *client.Client
}

func NewPmapV2(client *client.Client) *PmapV2 {
	return &PmapV2{
		Client: client,
	}
}

func (p *PmapV2) Null() (err error) {
	basic.SniffFunc("func PmapV2.Null {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100000,
				Vers:    2,
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
	basic.SniffFunc("} (func PmapV2.Null)")

	return
}

func (p *PmapV2) Set(params *Mapping) (result *bool, err error) {
	basic.SniffFunc("func PmapV2.Set {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100000,
				Vers:    2,
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
	result = new(bool)
	(*result), _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("(*result)", (*result))

	basic.SniffUnindent()
	basic.SniffFunc("} (func PmapV2.Set)")

	return
}

func (p *PmapV2) Unset(params *Mapping) (result *bool, err error) {
	basic.SniffFunc("func PmapV2.Unset {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100000,
				Vers:    2,
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
	result = new(bool)
	(*result), _, err = d.DecodeBool()

	if err != nil {
		return
	}
	basic.SniffDecode("(*result)", (*result))

	basic.SniffUnindent()
	basic.SniffFunc("} (func PmapV2.Unset)")

	return
}

func (p *PmapV2) Getport(params *Mapping) (result *uint32, err error) {
	basic.SniffFunc("func PmapV2.Getport {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100000,
				Vers:    2,
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
	result = new(uint32)
	(*result), _, err = d.DecodeUint()

	if err != nil {
		return
	}
	basic.SniffDecode("(*result)", (*result))

	basic.SniffUnindent()
	basic.SniffFunc("} (func PmapV2.Getport)")

	return
}

func (p *PmapV2) Dump() (result *PmaplistFirst, err error) {
	basic.SniffFunc("func PmapV2.Dump {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100000,
				Vers:    2,
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

	// Decode result
	d := response.Decoder
	result = new(PmaplistFirst)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func PmapV2.Dump)")

	return
}

func (p *PmapV2) Callit(params *CallArgs) (result *CallResult, err error) {
	basic.SniffFunc("func PmapV2.Callit {")
	basic.SniffIndent()

	message := &rpc2.RpcMsg{
		Body: rpc2.RpcBody{
			Mtype: rpc2.CALL,
			Union: rpc2.CallBody{
				Rpcvers: 2,
				Prog:    100000,
				Vers:    2,
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
	result = new(CallResult)
	err = (*result).Decode(d)
	if err != nil {
		return
	}

	basic.SniffUnindent()
	basic.SniffFunc("} (func PmapV2.Callit)")

	return
}
