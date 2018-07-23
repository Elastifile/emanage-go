package rpc2

import (
	"bytes"
	"log"
	"testing"

	xdr "github.com/davecgh/go-xdr/xdr2"
)

func TestAuthFlavor(*testing.T) {
	var in, out authFlavor
	var buf bytes.Buffer
	var n int
	var err error

	encoder := xdr.NewEncoder(&buf)
	decoder := xdr.NewDecoder(&buf)

	in = AUTH_UNIX

	log.Printf("encode")
	log.Printf("in: %+v", in)
	err = in.Encode(encoder)
	log.Printf("n: %v, err: %v", n, err)

	log.Printf("buf: %v", buf.Bytes())

	log.Printf("decode")
	err = out.Decode(decoder)
	log.Printf("n: %v, err: %v", n, err)
	log.Printf("out: %+v", out)
}

func TestAuthUnix(*testing.T) {
	var in, out AuthUnix
	var buf bytes.Buffer
	var n int
	var err error

	encoder := xdr.NewEncoder(&buf)
	decoder := xdr.NewDecoder(&buf)

	in = AuthUnix{
		Stamp:       0x41,
		Machinename: "foobarbaz",
		Uid:         0x42,
		Gid:         0x43,
		Gids:        []uint32{0x44, 0x45, 0x46, 0x47},
	}

	log.Printf("encode")
	log.Printf("in: %+v", in)
	err = in.Encode(encoder)
	log.Printf("n: %v, err: %v", n, err)

	log.Printf("buf: %v", buf.Bytes())

	log.Printf("decode")
	err = out.Decode(decoder)
	log.Printf("n: %v, err: %v", n, err)
	log.Printf("out: %+v", out)
}

func TestOpaqueAuth(t *testing.T) {
	var in, out OpaqueAuth
	var buf bytes.Buffer
	var n int
	var err error

	encoder := xdr.NewEncoder(&buf)
	decoder := xdr.NewDecoder(&buf)

	// Auth
	var authBuf bytes.Buffer
	authEncoder := xdr.NewEncoder(&authBuf)
	authBody := AuthUnix{
		Stamp:       0x41,
		Machinename: "foobarbaz",
		Uid:         0x42,
		Gid:         0x43,
		Gids:        []uint32{0x44, 0x45, 0x46, 0x47},
	}
	err = authBody.Encode(authEncoder)
	if err != nil {
		t.Fatal(err)
	}

	in = OpaqueAuth{
		Flavor: AUTH_UNIX,
		Body:   authBuf.Bytes(),
		// Body:   []uint8{65, 66, 67, 68, 69, 70},
	}

	log.Printf("encode")
	log.Printf("in authBody: %+v", authBody)
	log.Printf("in: %+v", in)
	err = in.Encode(encoder)
	log.Printf("n: %v, err: %v", n, err)

	log.Printf("buf: %v", buf.Bytes())

	log.Printf("decode")
	err = out.Decode(decoder)
	log.Printf("n: %v, err: %v", n, err)
	log.Printf("out: %+v", out)

	authBody = AuthUnix{}
	authDecoder := xdr.NewDecoder(bytes.NewBuffer(out.Body))
	err = authBody.Decode(authDecoder)
	log.Printf("out authBody: %+v", authBody)
}

func TestRpcMsg(t *testing.T) {
	var in, out RpcMsg
	var buf bytes.Buffer
	var n int
	var err error

	encoder := xdr.NewEncoder(&buf)
	decoder := xdr.NewDecoder(&buf)

	in = RpcMsg{
		Xid: 0x12345678,
		Body: RpcBody{
			Mtype: CALL, // Test what happens if wrong
			Union: CallBody{
				Rpcvers: 2,
				Prog:    100005, //nfsx.ProgramMountprogmountV3,
				Vers:    1,
				Proc:    5, //MOUNTPROC3_EXPORT,
				Cred:    OpaqueAuth{Flavor: AUTH_UNIX, Body: []uint8{0x41, 0x42, 0x43, 0x44}},
				Verf:    OpaqueAuth{Flavor: AUTH_NULL},
			},
		},
	}

	log.Printf("encode")
	log.Printf("in: %+v", in)
	err = in.Encode(encoder)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("n: %v, err: %v", n, err)

	log.Printf("buf: %v", buf.Bytes())

	log.Printf("decode")
	err = out.Decode(decoder)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("n: %v, err: %v", n, err)
	log.Printf("out: %+v", out)
}

/*
func TestRejectedReply(*testing.T) {
	var rr RejectedReply

	rr = MismatchInfoBody
}
*/

/*
func TestAuthFlavor(*testing.T) {
	var flavor authFlavor
}
*/
