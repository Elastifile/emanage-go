package ioutil2

import (
	"bytes"
	"io"

	"github.com/go-errors/errors"
)

var Blocksize = 32 * 1024

var (
	ErrShortRead     = errors.New("Short read")
	ErrInequalData   = errors.New("Inequal data")
	ErrInequalLength = errors.New("Inequal length")
)

// ReadersAreEqual reads from the provided readers and returns an error if their data differs
func ReadersAreEqual(r1, r2 io.Reader) error {
	buf1 := make([]byte, Blocksize)
	buf2 := make([]byte, Blocksize)

	for {
		n1, err1 := r1.Read(buf1)
		n2, err2 := io.ReadFull(r2, buf2[:n1])

		if n2 != n1 {
			return ErrShortRead
		}

		if bytes.Compare(buf1, buf2) != 0 {
			return ErrInequalData
		}

		if err1 == io.EOF {
			_, err2 = r2.Read(buf2)
			if err2 == io.EOF {
				return nil
			}
			return ErrInequalLength
		}
	}
}
