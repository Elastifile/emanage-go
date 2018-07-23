package client

import (
	xdr "github.com/davecgh/go-xdr/xdr2"

	"nfs/sunrpc/rpc2"
)

type Response struct {
	Decoder *xdr.Decoder
	message *rpc2.RpcMsg
}
