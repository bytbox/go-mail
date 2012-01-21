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

func getHeaders(s string) (hs []string, body string) {
	// TODO this could be faster via a rewrite without strings
	ps := strings.SplitN(s, "\r\n\r\n", 2)
	if len(ps) == 2 {
		body = ps[1]
	}
	ls := strings.Split(strings.Trim(ps[0], "\r\n"), "\r\n")
	hs = ls
	return
}

func splitHeader(s string) (k, v string) {
	// remove all CRLFs and split on the first colon
	ps := strings.SplitN(s, ":", 2)
	k = ps[0]
	v = strings.Replace(strings.TrimSpace(ps[1]), "\r\n", "", -1)
	return
}
