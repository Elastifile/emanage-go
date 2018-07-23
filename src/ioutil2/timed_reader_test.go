package ioutil2

import (
	"encoding/json"
	"io"
	"testing"
	"time"
)

type blockingReader struct{}

func (r *blockingReader) Read(p []byte) (n int, e error) {
	var c chan struct{}
	<-c
	return 0, nil
}

func TestTimedReader(t *testing.T) {
	ctx, _ := WithDeadline(nil, time.Now().Add(time.Second))
	tr := NewTimedReader(&blockingReader{}, ctx)
	dec := json.NewDecoder(tr)
	done := make(chan error)
	timeout := 2 * time.Second
	go func() {
		var s string
		done <- dec.Decode(&s)
	}()
	select {
	case e := <-done:
		if e != io.EOF {
			t.Fatalf("Received wrong error: %s", e)
		}
	case <-time.After(timeout):
		t.Fatalf("Failed to read in %s", timeout)
	}
}
