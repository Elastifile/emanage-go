package client

import (
	"bytes"
	"fmt"
	"io"
	"time"

	xdr "github.com/davecgh/go-xdr/xdr2"
	"github.com/go-errors/errors"

	"nfs/sunrpc/basic"
	"nfs/sunrpc/rpc2"
)

type Request struct {
	Xid                   Xid
	Message               *rpc2.RpcMsg
	Encoder               *xdr.Encoder
	sendBuf               *bytes.Buffer
	client                *Client
	responseChan          chan *Response
	injectedResponseDelay time.Duration
	Timeout               time.Duration
}

// Encode fragment header
func (request *Request) makeHeaderReader() (r io.Reader, err error) {
	headerBuf := &bytes.Buffer{}
	headerEnc := xdr.NewEncoder(headerBuf)

	header := rpc2.RecordMarkingLastFragmentFlag | uint32(request.sendBuf.Len())
	basic.SniffEncode("header", fmt.Sprintf("0x%x", header))
	_, err = headerEnc.EncodeUint(header)
	if err != nil {
		return nil, errors.Errorf(
			"Couldn't encode NFS message header %+v because: %s",
			header, err,
		)
	}

	return headerBuf, nil
}

// Send message and wait for response
func (request *Request) SendAndWaitForResponse() (response *Response, err error) {
	client := request.client

	timeStarted := time.Now()

	err = client.Send(request)
	if err != nil {
		// client.Send() already returns a formatted error
		return nil, err
	}

	response, err = request.waitForResponse()
	if err != nil {
		return nil, err
	}

	timeElapsed := time.Since(timeStarted)
	if timeElapsed > 50*time.Millisecond {
		logger.Debug("Slow operation",
			"operation", "SendAndWaitForResponse",
			"timeElapsed", timeElapsed,
			"xid", request.Xid,
			"proc", request.Message.Body.Union.(rpc2.CallBody).Proc,
		)
	}

	return response, nil
}

func (request *Request) waitForResponse() (response *Response, err error) {
	var ok bool
	timeStart := time.Now()

	select {
	case <-time.After(request.Timeout):
		return nil, errors.New(&rpc2.TimeoutError{Duration: time.Since(timeStart)})
	case response, ok = <-request.responseChan:
		if !ok {
			return nil, errors.New(&rpc2.IncompleteResponseError{})
		}
	}

	replyBody, ok := response.message.Body.Union.(rpc2.ReplyBody)
	if !ok {
		return response, errors.Errorf(
			"Expected a reply, but instead got %v",
			response.message.Body.Mtype,
		)
	}

	switch replyContents := replyBody.Union.(type) {
	case rpc2.AcceptedReply:
		if replyContents.ReplyData.Stat != rpc2.SUCCESS {
			return response, &rpc2.AcceptedReplyError{
				Stat:          replyContents.ReplyData.Stat,
				AcceptedReply: &replyContents,
			}
		}
	case rpc2.RejectedReply:
		return response, &rpc2.RejectedReplyError{
			Stat:          replyContents.Stat,
			RejectedReply: &replyContents,
		}
	default:
		return response, errors.Errorf(
			"Expected the reply to be accepted or denied, but instead got %v",
			replyBody.Stat,
		)
	}

	// Got a reply which was both accepted and successful
	return response, nil
}

func (request *Request) InjectResponseDelay(delay time.Duration) {
	request.injectedResponseDelay = delay
}

func (request *Request) handleResponseInjections(response *Response) {
	if request.injectedResponseDelay != 0 {
		logger.Debug("Injecting delay before returning response...", "delay", request.injectedResponseDelay)
		time.Sleep(request.injectedResponseDelay)
		logger.Debug("Done.")
	}
}
