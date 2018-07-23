package ioutil2

import "io"

// RepeatReader reads from the Reader that it wraps, possibly repeatedly, until the buffer is filled.
type RepeatReader struct {
	source io.ReadSeeker
}

func (rr *RepeatReader) Read(b []byte) (n int, err error) {
	for len(b) > 0 {
		m, err := rr.source.Read(b)
		n += m
		b = b[m:]
		if err == io.EOF {
			if _, err = rr.source.Seek(0, 0); err != nil {
				return n, err
			}
		}
	}
	return n, nil
}

func NewRepeatReader(r io.ReadSeeker) *RepeatReader {
	return &RepeatReader{r}
}
