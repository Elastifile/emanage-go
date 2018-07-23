package client

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	xdr "github.com/davecgh/go-xdr/xdr2"
	"github.com/go-errors/errors"
	multierror "github.com/hashicorp/go-multierror"
	log "gopkg.in/inconshreveable/log15.v2"

	"nfs/sunrpc/basic"
	"nfs/sunrpc/rpc2"
	"retry"
)

var liveClients = make(chan *Client)

func CloseAllClients() {
	for {
		select {
		case c := <-liveClients:
			c.Close()
		default:
			return
		}
	}
}

type requestMap map[Xid]*Request

type stopper struct {
	sync.Mutex
	timeout time.Duration
	timer   <-chan time.Time
	notify  chan struct{}
	closed  bool
}

func newStopper(timeout time.Duration) *stopper {
	return &stopper{
		timeout: timeout,
		notify:  make(chan struct{}),
	}
}

func (s *stopper) start() {
	s.Lock()
	defer s.Unlock()
	if s.timer == nil {
		logger.Info("stopper.start", "timeout", s.timeout)
		s.timer = time.After(s.timeout)
		go s.wait(s.timer)
	}
}

func (s *stopper) wait(clock <-chan time.Time) {
	<-clock
	s.Lock()
	defer s.Unlock()
	if s.timer != nil && !s.closed {
		s.closed = true
		close(s.notify)
	}
}

func (s *stopper) stop() {
	s.Lock()
	defer s.Unlock()
	s.timer = nil
}

// TODO(olegs): This could be refactored to use Go's rpc package.
// This should implement NFS traffic, while using rpc.Client
// as an interface for the users.
type Client struct {
	sync.Mutex
	AuthData       rpc2.OpaqueAuth
	name           string
	address        string
	numConns       int32
	nextXid        chan Xid
	connections    chan *connection
	chanRequestMap chan requestMap
	finish         chan bool
	reportFinish   chan bool
	open           bool
	lastLocalAddr  string
	clock          *stopper
	order          PoolOrder

	expectServerSideClose bool // TODO: Test with race detector
	log.Logger
}

type connection struct {
	sync.Mutex
	done     bool
	sock     net.Conn
	next     *connection
	prev     *connection
	numConns *int32
	reporter chan bool
}

func (c *connection) finish() {
	c.Lock()
	defer c.Unlock()
	if c.sock != nil && !c.done {
		_ = atomic.AddInt32(c.numConns, -1)
		_ = c.sock.Close()
		c.done = true
		if c.reporter != nil {
			c.reporter <- true
		}
	}
}

func (c *connection) wait(finish chan bool, reportFinish chan bool) {
	c.reporter = reportFinish
	<-finish
	c.finish()
}

func (c *connection) open() bool {
	return c.sock != nil && !c.done
}

func newConnection(numConns *int32) *connection {
	return &connection{
		numConns: numConns,
	}
}

var logger = log.New("package", "sunrpc/client")

type PoolOrder bool

const (
	LIFO PoolOrder = true
	FIFO PoolOrder = false
)

type RPCClientConfig struct {
	Name      string
	Host      string
	Port      uint32
	Mountport uint32
	Auth      rpc2.Auth
	NumConns  int
	XidSeed   uint32
	Timeout   time.Duration
	// The default (FIFO) order will cause connection pool to reuse
	// the connection which was most recently returned to the pull.
	// LIFO order will cause connections be be reused only after all
	// connections had been used at least once.  Thus, if you wanted
	// to create NumConns connections and perform an action on each
	// one of them, use LIFO order.  If you want to create minimum
	// number of connections, use FIFO order.
	Order PoolOrder
}

func NewClient(conf RPCClientConfig) (c *Client, err error) {
	authData, err := newAuthData(conf.Auth)
	if err != nil {
		// newAuthData() already returns error with stack
		return nil, err
	}

	if conf.Host == "" {
		return nil, errors.New("cannot connect: Empty host provided")
	}

	if conf.NumConns < 1 {
		conf.NumConns = 17
	}

	if conf.Timeout == 0 {
		conf.Timeout = rpc2.ResponseTimeout * 10
	}

	if conf.XidSeed == 0 {
		conf.XidSeed = 0x77770000
	}

	address := fmt.Sprintf("%v:%v", conf.Host, conf.Port)

	c = &Client{
		AuthData:       *authData,
		name:           conf.Name,
		address:        address,
		nextXid:        generateXids(conf.XidSeed),
		connections:    make(chan *connection, 1),
		chanRequestMap: make(chan requestMap, 1),
		finish:         make(chan bool),
		reportFinish:   make(chan bool),
		open:           true,
		clock:          newStopper(conf.Timeout),
		order:          conf.Order,

		Logger: logger.New("name", conf.Name, "address", address),
	}

	c.chanRequestMap <- make(requestMap)
	c.createConnections(conf.NumConns)
	go func() { liveClients <- c }()
	go func() {
		<-c.clock.notify
		c.Close()
	}()
	return c, nil
}

func (c *Client) createConnections(ncons int) {
	head := newConnection(&c.numConns)
	conn := head
	var nconn *connection
	for i := 0; i < ncons; i++ {
		nconn = newConnection(&c.numConns)
		conn.next = nconn
		nconn.prev = conn
		conn = nconn
	}
	conn.next = head
	head.prev = conn
	c.connections <- head
}

func (c *Client) NewRequest(message *rpc2.RpcMsg) (request *Request, err error) {
	xid := <-c.nextXid
	message.Xid = uint32(xid)

	sendBuf := &bytes.Buffer{}

	request = &Request{
		Xid:          Xid(message.Xid),
		Message:      message,
		Encoder:      xdr.NewEncoder(sendBuf),
		sendBuf:      sendBuf,
		client:       c,
		responseChan: make(chan *Response),
		Timeout:      c.clock.timeout,
	}

	basic.SniffMessage("NewRequest: message:", message)
	err = message.Encode(request.Encoder)
	if err != nil {
		return nil, errors.Errorf(
			"Couldn't encode %+v with encoder %+v because: %s",
			message, request.Encoder, err,
		)
	}

	requestMap := <-c.chanRequestMap
	requestMap[xid] = request
	c.chanRequestMap <- requestMap

	return request, nil
}

func (c *Client) initConnection(conn *connection) error {
	sock, err := net.DialTimeout("tcp4", c.address, c.clock.timeout)
	if err != nil {
		return errors.Errorf("Coudn't create SunRPC connection: %v", err)
	}
	conn.sock = sock
	_ = atomic.AddInt32(&c.numConns, 1)
	go conn.wait(c.finish, c.reportFinish)
	c.lastLocalAddr = sock.LocalAddr().String()
	go c.handleReplies(conn)
	return nil
}

func (c *Client) getConnection() (*connection, error) {
	for conn := range c.connections {
		if conn.next == conn {
			c.connections <- conn
			time.Sleep(time.Millisecond)
		} else {
			result := conn.next
			if result.sock == nil {
				if err := c.initConnection(result); err != nil {
					c.clock.start()
					logger.Error("Discarding connection", "error", err)
					conn.next.next.prev = conn
					conn.next = conn.next.next
					c.connections <- conn
					continue
				}
			}
			conn.next.next.prev = conn
			conn.next = conn.next.next
			c.connections <- conn
			return result, nil
		}
	}
	return nil, fmt.Errorf("Connections pool exhausted")
}

func (c *Client) putConnection(conn *connection) {
	head := <-c.connections
	if c.order == FIFO {
		conn.next = head.next
		conn.prev = head
		head.next = conn
	} else {
		conn.next = head
		head.prev.next = conn
		conn.prev = head.prev
	}
	c.connections <- head
}

func (c *Client) Send(request *Request) (err error) {
	if !c.open {
		return errors.New("Sending from closed client")
	}
	doSend := func() error {
		conn, serr := c.getConnection()
		if serr != nil {
			c.clock.start()
			return &retry.TemporaryError{
				Err: errors.Errorf(
					"Creating new connection failed because: %s",
					serr,
				),
			}
		}

		headerReader, serr := request.makeHeaderReader()
		if serr != nil {
			c.clock.start()
			c.putConnection(conn)
			// makeHeaderReader() already returns formatted error
			return &retry.TemporaryError{
				Err: serr,
			}
		}

		messageReader := io.MultiReader(headerReader, request.sendBuf)

		_, serr = io.Copy(conn.sock, messageReader)
		if serr != nil {
			conn.finish()
			c.clock.start()
			nconn := newConnection(&c.numConns)
			if cerr := c.initConnection(nconn); cerr == nil {
				c.putConnection(nconn)
			} else {
				serr = multierror.Append(serr, cerr)
			}
			return &retry.TemporaryError{
				Err: errors.Errorf("Writing request failed because: %s", serr),
			}
		}

		c.clock.stop()
		c.putConnection(conn)
		return nil
	}

	if err = retry.Do(c.clock.timeout, doSend); err != nil {
		if c.open {
			if terr, ok := err.(*retry.TemporaryError); ok {
				err = terr.Err
			}
			if ferr, ok := err.(*errors.Error); ok {
				return ferr
			}
			return errors.Errorf("Sending NFS request failed because: %s", err)
		}
	}

	return nil
}

func (c *Client) Close() {
	c.Lock()
	defer func() {
		c.open = false
		c.Unlock()
	}()
	if c.numConns == 0 {
		// We haven't been initialized yet...
		return
	}
	nconns := c.numConns
	close(c.finish)
	<-c.connections

	timer := time.After(time.Duration(nconns) * 10 * time.Millisecond)
	for nconns > 0 {
		select {
		case <-timer:
			logger.Warn("Not all connections were closed", "remaining", c.numConns)
			return
		case <-c.reportFinish:
			nconns--
		}
	}
}

func (c *Client) ExpectServerSideClose() {
	c.Debug("Expecting server-side close of connections")
	c.expectServerSideClose = true
}

func (c *Client) String() string {
	return fmt.Sprintf("&Client{%v: %v %v -> %v}",
		c.name,
		c.AuthData.Flavor,
		c.lastLocalAddr,
		c.address,
	)
}

func (c *Client) handleReplies(conn *connection) {
	for {
		select {
		case <-c.finish:
			conn.finish()
			return
		default:
			err := c.readAndHandleOneMessage(conn)
			if err != nil {
				conn.finish()
				return
			}
		}
	}
}

// TODO(olegs): This needs a lot of refactoring.  At least split this
// in two functions and deal with errors separately
func (c *Client) readAndHandleOneMessage(conn *connection) (err error) {
	responseSent := false

	defer func() {
		message := "Sending from closed client: %s"
		// If this was called after Client.Close(), it will panic
		if p := recover(); p != nil {
			err = errors.Errorf(message, p)
			return
		}
		if !responseSent {
			if err != nil {
				if !c.expectServerSideClose {
					logger.Debug("Error while handling message", "err", err)
				}
			}

			requestMap := <-c.chanRequestMap
			for xid, request := range requestMap {
				// Deleting while iterating is safe in Go.
				// See https://golang.org/ref/spec#For_statements
				delete(requestMap, xid)
				close(request.responseChan)
			}
			c.chanRequestMap <- requestMap
		}
		if !c.open {
			if ferr, ok := err.(*errors.Error); ok {
				ferr.Err = fmt.Errorf(message, ferr.Err)
			} else {
				err = errors.Errorf(message, err)
			}
		}
	}()

	messageReader, err := readMessage(conn.sock)
	if err != nil {
		// readMessage() should already return formatted error
		return err
	}

	decoder := xdr.NewDecoder(messageReader)
	responseMessage := &rpc2.RpcMsg{}
	err = responseMessage.Decode(decoder)
	if err != nil {
		return errors.Errorf(
			"Couldn't decode NFS message %+v because: %s",
			responseMessage, err,
		)
	}

	xid := Xid(responseMessage.Xid)

	basic.SniffMessage("responseMessage:", responseMessage)

	requestMap := <-c.chanRequestMap
	request, ok := requestMap[xid]
	c.chanRequestMap <- requestMap
	if !ok {
		return errors.Errorf("No outstanding request matching xid %v, dropping response", xid)
	}

	response := &Response{
		Decoder: decoder,
		message: responseMessage,
	}

	request.handleResponseInjections(response)
	request.responseChan <- response

	responseSent = true

	requestMap = <-c.chanRequestMap
	delete(requestMap, xid)
	c.chanRequestMap <- requestMap

	return nil
}
