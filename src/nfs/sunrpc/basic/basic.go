package basic

import xdr "github.com/davecgh/go-xdr/xdr2"

/*
// basic types
type Empty struct{}

// We define our own basic types since we want to define
// methods (e.g. Encode/Decode) on them
type Uint8 uint8
type Uint32 uint32
type Uint64 uint64

type Int8 int8
type Int32 int32
type Int64 int64

type Float32 float32
type Float64 float64

type String string
type Bool bool

*/

type Enum int32 // according to RFC 1014, XDR ints and enums are int32

// Encoding/Decoding

type encoder interface {
	Encode(e *xdr.Encoder) error
}

type decoder interface {
	Decode(d *xdr.Decoder) error
}

type fieldChecker interface {
	FieldCheck() error
}
