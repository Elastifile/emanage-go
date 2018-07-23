package ioutil2

import (
	"fmt"
	"io"
	"time"
)

type TimedReader struct {
	ctx Context
	src io.Reader
}

type CancelFunc func()

var (
	Canceled         = fmt.Errorf("Canceled")
	DeadlineExceeded = fmt.Errorf("DeadlineExceeded")
)

type Context interface {
	Deadline() (time.Time, bool)
	Done() <-chan struct{}
	Err() error
	Value(interface{}) interface{}
}

type context struct {
	m        map[interface{}]interface{}
	deadline time.Time
	done     chan struct{}
	err      error
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	c := &context{
		m:        map[interface{}]interface{}{},
		done:     make(chan struct{}),
		deadline: deadline,
	}
	finish := make(chan struct{})
	f := func() {
		close(finish)
	}
	go func() {
		select {
		case <-time.After(deadline.Sub(time.Now())):
			c.err = DeadlineExceeded
			close(c.done)
		case <-finish:
			c.err = Canceled
			close(c.done)
		}
	}()
	return c, f
}

func (c *context) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

func (c *context) Done() <-chan struct{} { return c.done }

func (c *context) Err() error { return c.err }

func (c *context) Value(key interface{}) interface{} { return c.m[key] }

// TODO(olegs): Once we move to Go 1.7 use context.Context instead
func NewTimedReader(src io.Reader, ctx Context) *TimedReader {
	return &TimedReader{
		ctx: ctx,
		src: src,
	}
}

func (tr *TimedReader) Read(p []byte) (int, error) {
	type res struct {
		n int
		e error
	}
	finish := make(chan res)
	go func() {
		n, e := tr.src.Read(p)
		finish <- res{n: n, e: e}
	}()
	select {
	case <-tr.ctx.Done():
		return 0, io.EOF
	case r := <-finish:
		return r.n, r.e
	}
}
