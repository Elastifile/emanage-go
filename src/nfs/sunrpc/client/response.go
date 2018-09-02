package client

import (
	xdr "github.com/davecgh/go-xdr/xdr2"

	"github.com/elastifile/emanage-go/src/nfs/sunrpc/rpc2"
)

type Response struct {
	Decoder *xdr.Decoder
	message *rpc2.RpcMsg
}
