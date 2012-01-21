// Package mail implements a parser for electronic mail messages as specified
// in RFC2822.
package mail

import (
	"strings"
)

type Message struct {

}

func Parse(s string) (m Message) {
	return
}

func getHeaders(s string) (hs []string) {
	return
}

func splitHeader(s string) (k, v string) {
	// remove all CRLFs and split on the first colon
	ps := strings.SplitN(s, ":", 2)
	k = ps[0]
	v = strings.Replace(strings.TrimSpace(ps[1]), "\r\n", "", -1)
	return
}
