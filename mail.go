// Package mail implements a parser for electronic mail messages as specified
// in RFC2822.
package mail

import (
	"bytes"
	"errors"
)

type Header struct {
	Key, Value []byte
}

type Message struct {
	RawHeaders []Header
	Body       []byte
}

func isWSP(b byte) bool {
	return b == ' ' || b == '\t'
}

func Parse(s []byte) (m Message, e error) {
	// parser states
	const (
		READY = iota
		HKEY
		HVWS
		HVAL
	)

	const (
		CR = '\r'
		LF = '\n'
	)
	CRLF := []byte{CR, LF}

	state := READY
	kstart, kend, vstart := 0, 0, 0
	done := false

	m.RawHeaders = []Header{}

	for i := 0; i < len(s); i++ {
		b := s[i]
		switch state {
		case READY:
			if b == CR && i < len(s)-1 && s[i+1] == LF {
				// we are at the beginning of an empty header
				m.Body = s[i+2:]
				done = true
				goto Done
			}
			// otherwise this character is the first in a header
			// key
			kstart = i
			state = HKEY
		case HKEY:
			if b == ':' {
				kend = i
				state = HVWS
			}
		case HVWS:
			if !isWSP(b) {
				vstart = i
				state = HVAL
			}
		case HVAL:
			if b == CR && i < len(s)-2 && s[i+1] == LF && !isWSP(s[i+2]) {
				v := bytes.Replace(s[vstart:i], CRLF, nil, -1)
				hdr := Header{s[kstart:kend], v}
				m.RawHeaders = append(m.RawHeaders, hdr)
				state = READY
				i++
			}
		}
	}
Done:
	if !done {
		e = errors.New("unexpected EOF")
	}
	return
}
