package nfs

import (
	"nfs/sunrpc/nfsx"
	"nfs/sunrpc/rpc2"
	"fmt"

	"github.com/go-errors/errors"
)

// FIXME(olegs): Remove this and its dependencies
// ErrNotADirectory is returned when a directory operation is performed on a non-directory file.
var ErrNotADirectory = errors.New("Entry is not a directory")

type FSError struct {
	cause   *errors.Error
	message string
	args    []interface{}
	stack   string
}

func (e *FSError) Error() string {
	var (
		args    []interface{}
		message string
	)
	if e.cause != nil {
		args = append(e.args, e.cause.Err, e.stack)
		message = e.message + "\nCaused by: %s\n%s"
	} else {
		args = append(e.args, e.stack)
		message = e.message + "\n%s"
	}
	return fmt.Sprintf(message, args...)
}

func (e *FSError) GoString() string {
	return e.Error()
}

func NewFSError(cause error, message string, args ...interface{}) *FSError {
	var (
		wrapped *errors.Error
		stack   string
	)
	if e, ok := cause.(*errors.Error); ok {
		wrapped = e
		stack = e.ErrorStack()
	} else if cause != nil {
		wrapped = errors.Wrap(cause, 2)
		stack = wrapped.ErrorStack()
	}
	return &FSError{
		cause:   wrapped,
		message: message,
		args:    args,
		stack:   stack,
	}
}

func (e *FSError) NFSError() error {
	if e.cause != nil {
		switch err := e.cause.Err.(type) {
		case *nfsx.NfsError:
			return err
		case *FSError:
			return err.NFSError()
		default:
			return nil
		}
	}
	return nil
}

func (e *FSError) NotADirectory() bool {
	return e.cause != nil && e.cause == ErrNotADirectory
}

func (e *FSError) CanIgnoreOnRemove() bool {
	if e.cause != nil {
		switch inner := e.cause.Err.(type) {
		case *nfsx.NfsError:
			switch inner.Status {
			case nfsx.V3ERR_ISDIR:
				// OK, Happens on Linux
				return true
			case nfsx.V3ERR_IO:
				// OK, Happens on OS X
				return true
			}
		}
		return false
	}
	return true
}

func (e *FSError) DoesNotExist() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.NfsError); ok {
			return inner.Status == nfsx.V3ERR_NOENT
		}
	}
	return false
}

func (e *FSError) Exists() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.NfsError); ok {
			return inner.Status == nfsx.V3ERR_EXIST
		}
	}
	return false
}

func (e *FSError) NotEmpty() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.NfsError); ok {
			return inner.Status == nfsx.V3ERR_NOTEMPTY
		}
	}
	return false
}

func (e *FSError) AccessViolation() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.NfsError); ok {
			return inner.Status == nfsx.V3ERR_ACCES
		}
	}
	return false
}

func (e *FSError) RPCGarbage() bool {
	// Happens on OS X. Does not match the spec.
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*rpc2.AcceptedReplyError); ok {
			return inner.Stat == rpc2.GARBAGE_ARGS
		}
	}
	return false
}

func (e *FSError) NameTooLong() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.NfsError); ok {
			return inner.Status == nfsx.V3ERR_NAMETOOLONG
		}
	}
	return false
}

func (e *FSError) IsNFS() bool {
	if e.cause != nil {
		_, ok := e.cause.Err.(*nfsx.NfsError)
		return ok
	}
	return false
}

func (e *FSError) IsRejectReply() bool {
	if e.cause != nil {
		_, ok := e.cause.Err.(*rpc2.RejectedReplyError)
		return ok
	}
	return false
}

func (e *FSError) IsAcceptedReply() bool {
	if e.cause != nil {
		_, ok := e.cause.Err.(*rpc2.AcceptedReplyError)
		return ok
	}
	return false
}

func (e *FSError) IsTimeout() bool {
	if e.cause != nil {
		_, ok := e.cause.Err.(*rpc2.TimeoutError)
		return ok
	}
	return false
}

func (e *FSError) IsIncompleteResponse() bool {
	if e.cause != nil {
		_, ok := e.cause.Err.(*rpc2.IncompleteResponseError)
		return ok
	}
	return false
}

func (e *FSError) IsAttributeMismatch() bool {
	if e.cause != nil {
		_, ok := e.cause.Err.(*nfsx.AttributeMismatchError)
		return ok
	}
	return false
}

func (e *FSError) Auth() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*rpc2.RejectedReplyError); ok {
			return inner.Stat == rpc2.AUTH_ERROR
		}
	}
	return false
}

func (e *FSError) TooBig() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.NfsError); ok {
			return inner.Status == nfsx.V3ERR_FBIG
		}
	}
	return false
}

func (e *FSError) Permissions() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.NfsError); ok {
			return inner.Status == nfsx.V3ERR_PERM
		}
	}
	return false
}

func (e *FSError) Stale() bool {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.NfsError); ok {
			return inner.Status == nfsx.V3ERR_STALE
		}
	}
	return false
}

func (e *FSError) Uid() nfsx.Uid3 {
	if e.cause != nil {
		if inner, ok := e.cause.Err.(*nfsx.AttributeMismatchError); ok {
			return inner.Fattr3.Uid
		}
	}
	return 0
}
