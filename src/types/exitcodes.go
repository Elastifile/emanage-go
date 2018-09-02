package types

import tester_types "github.com/elastifile/emanage-go/src/tester/types"

// Tesla exit codes

//go:generate stringer -type=ExitCode

type ExitCode int

const (
	ExitCodeSuccess ExitCode = iota
	ExitCodeTeslaFailed
	ExitCodeToolFailed
	ExitCodeVerifyFailed
	ExitCodeFailed
	ExitCodeUnknown

	ExitCodeTesterFocus ExitCode = tester_types.TESTER_FOCUS_EXIT_CODE
)
