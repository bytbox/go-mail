// Package mail implements a parser for electronic mail messages as specified
// in RFC2822.
//
// We allow both CRLF and LF to be used in the input, possibly mixed.
package mail

import (
	"bytes"
	"errors"
)

type Message struct {
	FullHeaders []Header
	Text        string
}

type Header struct {
	Key, Value string
}

func Parse(s []byte) (m Message, e error) {
	r, e := ParseRaw(s)
	if e != nil {
		return
	}
	return Process(r)
}

func Process(r RawMessage) (m Message, e error) {
	return
}

type RawHeader struct {
	Key, Value []byte
}

type RawMessage struct {
	RawHeaders []RawHeader
	Body       []byte
}

func isWSP(b byte) bool {
	return b == ' ' || b == '\t'
}

func ParseRaw(s []byte) (m RawMessage, e error) {
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

	m.RawHeaders = []RawHeader{}

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
			if b == LF {
				m.Body = s[i+1:]
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
				hdr := RawHeader{s[kstart:kend], v}
				m.RawHeaders = append(m.RawHeaders, hdr)
				state = READY
				i++
			} else if b == LF && i < len(s)-1 && !isWSP(s[i+1]) {
				v := bytes.Replace(s[vstart:i], CRLF, nil, -1)
				hdr := RawHeader{s[kstart:kend], v}
				m.RawHeaders = append(m.RawHeaders, hdr)
				state = READY
			}
		}
	}
Done:
	if !done {
		e = errors.New("unexpected EOF")
	}
	return
}
