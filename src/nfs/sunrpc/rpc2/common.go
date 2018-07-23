package rpc2

import "time"

const (
	// Time to wait for a response from the server before timing out
	// (orenz:) - updated to 5 minute so ensure NodeFailure capacity tool test case will not failed on timeout
	ResponseTimeout = 300 * time.Second
	InitialTimeout  = 2 * time.Minute

	// RPC record marking
	RecordMarkingLastFragmentFlag uint32 = 0x80000000
	RecordMarkingLastFragmentMask uint32 = ^RecordMarkingLastFragmentFlag
)
