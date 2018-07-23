package ioutil2

import (
	"bytes"
	"io"
)

type FilterReader struct {
	filter func(b rune) bool
	src    io.Reader
}

func NewFilterReader(src io.Reader, filter func(b rune) bool) *FilterReader {
	return &FilterReader{
		filter: filter,
		src:    src,
	}
}

func (fr *FilterReader) Read(p []byte) (int, error) {
	cpy := make([]byte, len(p))
	n, e := fr.src.Read(cpy)
	if n != 0 {
		d := bytes.Runes(cpy[:n])
		c := make([]rune, len(d))
		i := 0
		for _, n := range d {
			if fr.filter(n) {
				c[i] = n
				i++
			}
		}
		c = c[:i]
		n = i
		copy(p, []byte(string(c)))
	}
	return n, e
}
