package client

import (
	"bytes"
	"fmt"

	xdr "github.com/davecgh/go-xdr/xdr2"

	"nfs/sunrpc/rpc2"

	"github.com/go-errors/errors"
)

type Xid uint32

func (xid Xid) String() string {
	return fmt.Sprintf("%#08x", uint32(xid))
}

func generateXids(seed uint32) (xidChan chan Xid) {
	/*
		const xidPoolSegmentSize = 1000000
		const xidPoolSize = (^Xid(0)) / xidPoolSegmentSize
	*/

	xidChan = make(chan Xid)

	go func() {
		xid := Xid(seed)
		for {
			// for xid := Xid(0); xid < xidPoolSize; xid += xidPoolSegmentSize {
			xidChan <- xid
			xid++
			// }
		}
	}()

	return xidChan
}

func newAuthData(auth rpc2.Auth) (opaqueAuth *rpc2.OpaqueAuth, err error) {
	var authBuf bytes.Buffer
	authEncoder := xdr.NewEncoder(&authBuf)

	switch actualAuth := auth.(type) {
	case *rpc2.AuthUnix:
		err = actualAuth.Encode(authEncoder)
		opaqueAuth = &rpc2.OpaqueAuth{
			Flavor: rpc2.AUTH_UNIX,
			Body:   authBuf.Bytes(),
		}
		if err != nil {
			err = errors.Errorf(
				"Couldn't encode <%T>%+v because %s",
				auth, auth, err,
			)
		}
	case *rpc2.AuthNull:
		err = nil
		opaqueAuth = &rpc2.OpaqueAuth{
			Flavor: rpc2.AUTH_NULL,
			Body:   []uint8{},
		}
	default:
		return nil, errors.Errorf("Unsupported auth type: %T", actualAuth)
	}

	return opaqueAuth, err
}
