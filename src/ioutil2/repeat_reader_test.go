package ioutil2

import (
	"bytes"
	"io"
	"log"
	"testing"
)

func TestRepeatReaderOnce(t *testing.T) {
	data := []byte("ABCDE/")
	r := bytes.NewReader(data)
	rr := NewRepeatReader(r)

	length := int64(len(data))

	var dst bytes.Buffer
	n, err := io.CopyN(&dst, rr, length)

	if err != nil {
		t.Fatal(err)
	}
	if n != length {
		t.Fatal("Copied wrong length")
	}

	if bytes.Compare(dst.Bytes(), data) != 0 {
		t.Fatal("Bad data")
	}
}

func TestRepeatReaderCount(t *testing.T) {
	data := []byte("ABCDE/")
	r := bytes.NewReader(data)
	rr := NewRepeatReader(r)

	for count := 0; count < 10; count++ {
		length := int64(count * len(data))

		var dst bytes.Buffer
		n, err := io.CopyN(&dst, rr, length)

		if err != nil {
			t.Fatal(err)
		}
		if n != length {
			t.Fatalf("Copied wrong length; count: %v", count)
		}

		log.Printf("data: %s", data)
		log.Printf("dst: %s", dst.Bytes())

		if bytes.Compare(dst.Bytes(), bytes.Repeat(data, count)) != 0 {
			t.Fatalf("Bad data; count: %v", count)
		}
	}
}

func TestRepeatReaderLength(t *testing.T) {
	data := []byte("ABCDE/")
	r := bytes.NewReader(data)
	rr := NewRepeatReader(r)

	for length := int64(0); length < 10*int64(len(data)); length++ {
		var dst bytes.Buffer
		n, err := io.CopyN(&dst, rr, length)

		if err != nil {
			t.Fatal(err)
		}
		if n != length {
			t.Fatalf("Copied wrong length; length: %v", length)
		}

		log.Printf("data: %s", data)
		log.Printf("dst: %s", dst.Bytes())

		// if bytes.Compare(dst.Bytes(), bytes.Repeat(data, count)) != 0 {
		// t.Fatal("Bad data; length: %v", length)
		// }
	}
}
