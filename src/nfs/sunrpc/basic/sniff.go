package basic

import (
	"fmt"
	"strings"

	log "gopkg.in/inconshreveable/log15.v2"

	"github.com/davecgh/go-spew/spew"
)

var logger = log.New("package", "sunrpc/sniff")

var chanIndent chan int

var Sniff bool

/*

TODO: Add context for sniffing.

Currently the sniffs for all concurrent requests are interleaved.
We should add context information to allow filtering out sniffs for a specific request.
Of course the indentLevel should be managed per sniff context.

*/

func init() {
	chanIndent = make(chan int, 1)
	chanIndent <- 0
}

func SniffEncode(name string, a ...interface{}) {
	if Sniff {
		logger.Debug(fmt.Sprintf("Sniff: >>> %v%v %v", getIndent(), name, spew.Sdump(a...)))
	}
}

func SniffDecode(name string, a ...interface{}) {
	if Sniff {
		logger.Debug(fmt.Sprintf("Sniff: <<< %v%v %v", getIndent(), name, spew.Sdump(a...)))
	}
}

func SniffFunc(name string, a ...interface{}) {
	if Sniff {
		logger.Debug(fmt.Sprintf("Sniff: === %v%v %v", getIndent(), name, spew.Sdump(a...)))
	}
}

func SniffMessage(name string, a ...interface{}) {
	SniffFunc(name, a...)
}

func getIndent() string {
	indentLevel := <-chanIndent
	indent := strings.Repeat("    ", indentLevel)
	chanIndent <- indentLevel
	return indent
}

func SniffIndent() {
	if Sniff {
		indentLevel := <-chanIndent
		indentLevel++
		chanIndent <- indentLevel
	}
}

func SniffUnindent() {
	if Sniff {
		indentLevel := <-chanIndent
		indentLevel--
		chanIndent <- indentLevel
	}
}
