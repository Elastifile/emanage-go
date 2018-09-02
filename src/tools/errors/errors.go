package tool_errors

import (
	"fmt"

	"github.com/go-errors/errors"
	multierror "github.com/hashicorp/go-multierror"

	"github.com/elastifile/emanage-go/src/types"
)

var (
	ErrNoLoaders      = errors.New("At least one loader is required")
	ErrNotImplemented = errors.New("You must override this method.")
)

// REVIEW(gavriep): Why do we need so many specific error types?
// I don't see anywhere in the code that these are checked for.
// Why not just return general errors with errors.Errorf, and create specific errors only if/when they will be needed?
type MissingArgsError struct {
	Args interface{}
}

func (e *MissingArgsError) Error() string {
	return fmt.Sprintf("Missing required arguments in %s",
		e.Args)
}

type InvalidResultsError struct {
	Raw   string
	Cause error
	Pos   int
}

func (e *InvalidResultsError) Error() string {
	return fmt.Sprintf("Results didn't match the expected format: '%s', at '%d', caused by: %s",
		e.Raw, e.Pos, e.Cause)
}

type NoResultsError struct {
	Host types.Host
}

func (e *NoResultsError) Error() string {
	return fmt.Sprintf("No results found for host: '%s'", e.Host)
}

type NoSfsResultsError struct {
	Host   types.Host
	Suffix string
}

func (e *NoSfsResultsError) Error() string {
	return fmt.Sprintf("No results found for host %v, see sfslog.%v to debug failure.",
		e.Host, e.Suffix)
}

type EmptyResultsError struct {
	Host types.Host
	File *types.NamedReader
}

func (e *EmptyResultsError) Error() string {
	return fmt.Sprintf("Tool Failed!\n(produced empty results file '%s' for host: '%s')",
		e.File, e.Host)
}

type WrongThresholdError struct {
	Cause error
}

func (e *WrongThresholdError) Error() string {
	return fmt.Sprintf("Results dont match the expected thresholds.  err: %v",
		e.Cause)
}

// REVIEW(gavriep): The following "multierror" extensions are not specific to this package.
// We should consider creating our own multierror package and add these to it.
// Alternatively, we could extend the existing multierror package.
func NonNil(errs ...error) (result []error) {
	for _, err := range errs {
		result = MaybeAppend(result, err)
	}
	return result
}

func Collect(errs ...error) (err error) {
	for _, e := range errs {
		if e != nil {
			err = multierror.Append(err, e)
		}
	}
	return err
}

func MaybeAppend(s []error, e error) []error {
	if e == nil {
		return s
	}
	return append(s, e)
}

// This will be used by Connectathon once we decide that the change in
// Erun can be used by other tools.

// Used together with `FoldErrors', this error handling strategy will
// short-circuit as soon as the first non-empty error is seen
func BarfStrategy(acc []error, next error) ([]error, bool) {
	if next != nil {
		return []error{next}, false
	}
	return acc, true
}

// Used together with `FoldErrors', this error handling strategy will
// collect all errors.
func AccumulateStrategy(acc []error, next error) ([]error, bool) {
	if next == nil {
		return acc, true
	}
	return append(acc, next), true
}

// Used together with `FoldErrors', this error handling strategy will
// only collect the error matching `filter'.
func FilterStrategy(filter func(e error) bool) types.ErrorHandlingStrategy {
	return func(acc []error, next error) ([]error, bool) {
		if !filter(next) {
			return append(acc, next), true
		}
		return acc, true
	}
}

// TODO: Maybe return the number of unhandled errors?
func FoldErrors(fun types.ErrorHandlingStrategy, c chan error, n int) []error {
	// REVIEW(gavriep): No need to explicitly set a value
	var (
		result []error
		ok     bool
	)
	for i := 0; i < n; i++ {
		err := <-c
		result, ok = fun(result, err)
		if !ok {
			break
		}
	}
	return result
}

func CollectAll(watcher chan error, n int) error {
	if n == 0 {
		return nil
	}
	return Collect(FoldErrors(AccumulateStrategy, watcher, n)...)
}
