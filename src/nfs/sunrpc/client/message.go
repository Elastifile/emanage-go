package client

import (
	"bytes"
	"fmt"
	"io"

	xdr "github.com/davecgh/go-xdr/xdr2"
	"github.com/go-errors/errors"

	"nfs/sunrpc/basic"
	"nfs/sunrpc/rpc2"
)

// Inject error responses
var injectError = false
var injectCounter int

func readMessage(in io.Reader) (out io.Reader, err error) {
	decoder := xdr.NewDecoder(in)
	var readers []io.Reader

	for moreFragments := true; moreFragments; {
		var data []byte
		data, moreFragments, err = readFragment(in, decoder)
		readers = append(readers, bytes.NewReader(data))
		if err != nil {
			break
		}
	}
	return io.MultiReader(readers...), err
}

func readFragment(r io.Reader, decoder *xdr.Decoder) (data []byte, moreFragments bool, err error) {
	var header uint32
	header, _, err = decoder.DecodeUint()
	if err != nil {
		return nil, false, errors.Errorf("Couldn't decode NFS fragment: %s", err)
	}
	basic.SniffDecode("header", fmt.Sprintf("0x%x", header))

	length := header & rpc2.RecordMarkingLastFragmentMask
	data = make([]byte, length)
	_, err = io.ReadFull(r, data)

	if injectError {
		injectCounter++
		if injectCounter > 100 {
			err = &rpc2.IncompleteResponseError{}
			injectError = false
			// injectCounter = 0
		}
	}

	if err != nil {
		return data, false, errors.Errorf(
			"Couldn't read NFS fragment entirely: %s", err,
		)
	}

	moreFragments = (header & rpc2.RecordMarkingLastFragmentFlag) == 0

	return data, moreFragments, nil
}
