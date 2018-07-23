package rpc2

import (
	"fmt"
	"time"
)

//////////

type RejectedReplyError struct {
	Stat rejectStat
	*RejectedReply
}

func (e *RejectedReplyError) Error() string {
	return fmt.Sprintf("RejectedReplyError: %v, %#v", e.RejectedReply.Stat, e.RejectedReply.Union)
}

//////////

type AcceptedReplyError struct {
	Stat acceptStat
	*AcceptedReply
}

func (e *AcceptedReplyError) Error() string {
	return fmt.Sprintf("AcceptedReplyError: %v, %#v",
		e.AcceptedReply.ReplyData.Stat,
		e.AcceptedReply.ReplyData.Union,
	)
}

//////////

type TimeoutError struct {
	time.Duration
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("timed out after %v", e.Duration)
}

func (e *TimeoutError) Timeout() bool {
	return true
}

func (e *TimeoutError) Temporary() bool {
	return false
}

//////////

type IncompleteResponseError struct{}

func (e *IncompleteResponseError) Error() string {
	return fmt.Sprintf("IncompleteResponseError")
}

func (e *IncompleteResponseError) Temporary() bool {
	return true
}

//////////
