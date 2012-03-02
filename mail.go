// Package mail implements a parser for electronic mail messages as specified
// in RFC5322.
//
// We allow both CRLF and LF to be used in the input, possibly mixed.
package mail

import (
	"bytes"
	"encoding/base64"
	"errors"
	"strings"
)

var benc = base64.URLEncoding

type Message struct {
	FullHeaders []Header // all headers
	OptHeaders  []Header // unprocessed headers

	MessageId   string
	Id          string
	Subject     string
	Comments    []string
	Keywords    []string

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
	m.FullHeaders = []Header{}
	m.OptHeaders = []Header{}
	m.Text = string(r.Body) // TODO mime
	for _, rh := range r.RawHeaders {
		h := Header{string(rh.Key), string(rh.Value)}
		m.FullHeaders = append(m.FullHeaders, h)
		switch string(rh.Key) {
		case `Message-ID`:
			m.MessageId = string(rh.Value)
			m.Id = benc.EncodeToString(rh.Value)
		case `Subject`:
			m.Subject = string(rh.Value)
		case `Comments`:
			m.Comments = append(m.Comments, string(rh.Value))
		case `Keywords`:
			ks := strings.Split(string(rh.Value), ",")
			for _, k := range ks {
				m.Keywords = append(m.Keywords, strings.TrimSpace(k))
			}
		default:
			m.OptHeaders = append(m.OptHeaders, h)
		}
	}
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
