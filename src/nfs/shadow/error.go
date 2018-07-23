package shadow

import (
	"fmt"
	"io"

	"nfs"
	"retry"
)

type ShadowError struct {
	message string
	ErrRefr error
	ErrElfs error
}

func (e *ShadowError) Error() string {
	return e.message
}

func (e *ShadowError) Temporary() bool {
	return retry.IsTemporary(e.ErrRefr) || retry.IsTemporary(e.ErrElfs)
}

func (e *ShadowError) Timeout() bool {
	return retry.IsTimeout(e.ErrRefr) || retry.IsTimeout(e.ErrElfs)
}

////////////////////////////////////////

func checkShadow(description string, errRefr, errElfs error) error {
	fmt.Printf("%v: Reference: %v, ELFS: %v", description, errRefr, errElfs)
	if errRefr == errElfs {
		// Both may be nil, or identical real errors
		return errElfs
	}

	if errRefr != nil && errElfs != nil {
		// They may be similar errors but not the same object,
		// so we compare the string values.
		// This may not be very efficient since it involves string generation.
		// It may be preferable to do a DeepEqual test or something similar.
		if errRefr.Error() == errElfs.Error() {
			return errElfs
		}
	}

	if errRefr == io.EOF && errElfs == nil {
		logger.Warn("Encountered possible known issue, compensating: EL-491: " +
			"ELFS doesn't set the EOF flag when reading the last byte of the file")
		return io.EOF
	}

	return &ShadowError{
		fmt.Sprintf("%v: Reference: %v, ELFS: %v", description, errRefr, errElfs),
		errRefr,
		errElfs,
	}
}

////////////////////////////////////////

type ShadowDataError struct {
	Message        string
	Offset         int64
	LengthRefrData int
	LengthElfsData int
	RefrData       []byte
	ElfsData       []byte
	ElfsFile       nfs.File
}

func (e *ShadowDataError) Error() string {
	return e.Message
}

////////////////////////////////////////

type ShadowDirectoryMismatchError struct {
	OnlyInElfs []string
	OnlyInRefr []string
}

func (e *ShadowDirectoryMismatchError) Error() string {
	return "ELFS and reference have different entries"
}

func (e *ShadowDirectoryMismatchError) Temporary() bool {
	// We consider this a temporary error, since it may occur when using
	// several concurrent goroutines to create directory entries.
	return true
}
