// Package mail implements a parser for electronic mail messages as specified
// in RFC2822.
//
// Although RFC2822 requires that all lines end in CRLF, this package will
// accept just LF for compatibility with bufio.ReadLine; however, all output
// from this package will use LF only.
package mail

import (
	"strings"
)

type Header struct {
	Key, Value string
}

type Message struct {
	RawHeaders []Header
	Body       string
}

func Parse(s string) (m Message) {
	hs, b := getHeaders(s)
	m.Body = b
	for _, hl := range hs {
		k, v := splitHeader(hl)
		h := Header{k, v}
		m.RawHeaders = append(m.RawHeaders, h)
	}
	return
}

func getHeaders(s string) (hs []string, body string) {
	// TODO this could be faster via a rewrite without `strings'

	// Use LF
	s = strings.Replace(s, "\r\n", "\n", -1)

	ps := strings.SplitN(s, "\n\n", 2)
	if len(ps) == 2 {
		body = ps[1]
	}
	ls := strings.Split(strings.Trim(ps[0], "\n"), "\n")
	i := 0
	for i < len(ls) {
		l := ls[i]
		if l[0] == ' ' || l[0] == '\t' {
			if len(hs) > 0 {
				hs[len(hs)-1] += "\n" + l
			} // TODO error
		} else {
			hs = append(hs, l)
		}
		i++
	}
	return
}

func splitHeader(s string) (k, v string) {
	// remove all CRLFs and split on the first colon
	ps := strings.SplitN(s, ":", 2)
	k = ps[0]
	v = strings.Replace(strings.TrimSpace(ps[1]), "\n", "", -1)
	v = strings.Replace(v, "\r", "", -1)
	return
}
