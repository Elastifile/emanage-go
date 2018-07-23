package ioutil2_test

import (
	"bytes"
	"testing"

	. "elastifile/tesla/infra/ioutil2"
)

func TestReadersEqual(t *testing.T) {
	type testCase struct {
		a   string
		b   string
		err error
	}
	testData := []testCase{
		testCase{
			a:   "some test data",
			b:   "some test data",
			err: nil,
		},
		testCase{
			a:   "some test data",
			b:   "some foob data",
			err: ErrInequalData,
		},
		testCase{
			a:   "some test data",
			b:   "some te",
			err: ErrShortRead,
		},
		testCase{
			a:   "some te",
			b:   "some test data",
			err: ErrInequalLength,
		},
		testCase{
			a:   "some tex",
			b:   "some test data",
			err: ErrInequalData,
		},
	}
	for Blocksize = 1; Blocksize < 20; Blocksize *= 2 {
		for _, tc := range testData {
			r1 := bytes.NewBufferString(tc.a)
			r2 := bytes.NewBufferString(tc.b)
			actual := ReadersAreEqual(r1, r2)
			if tc.err != actual {
				t.Fatalf("Expected %s but got %s", tc.err, actual)
			}
		}
	}
}
