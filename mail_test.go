package mail

import (
	"reflect"
	"strings"
	"testing"
)

// Converts all newlines to CRLFs.
func crlf(s string) []byte {
	return []byte(strings.Replace(s, "\n", "\r\n", -1))

}

type parseRawTest struct {
	msg []byte
	ret RawMessage
}

var parseRawTests = []parseRawTest{
	parseRawTest{
		msg: crlf(`
`),
		ret: RawMessage{
			RawHeaders: []RawHeader{},
			Body:       crlf(""),
		},
	},
	parseRawTest{
		msg: crlf(`
ab
c
`),
		ret: RawMessage{
			RawHeaders: []RawHeader{},
			Body: crlf(`ab
c
`),
		},
	},
	parseRawTest{
		msg: crlf(`a: b

`),
		ret: RawMessage{
			RawHeaders: []RawHeader{RawHeader{crlf("a"), crlf("b")}},
			Body:       crlf(""),
		},
	},
	parseRawTest{
		msg: crlf(`a: b
c: def
 hi

`),
		ret: RawMessage{
			RawHeaders: []RawHeader{
				RawHeader{crlf("a"), crlf("b")},
				RawHeader{crlf("c"), crlf("def hi")},
			},
			Body: crlf(``),
		},
	},
	parseRawTest{
		msg: crlf(`a: b
c: d fdsa
ef:  as

hello, world
`),
		ret: RawMessage{
			RawHeaders: []RawHeader{
				RawHeader{crlf("a"), crlf("b")},
				RawHeader{crlf("c"), crlf("d fdsa")},
				RawHeader{crlf("ef"), crlf("as")},
			},
			Body: crlf(`hello, world
`),
		},
	},
	parseRawTest{
		msg: []byte(`a: b
c: d fdsa
ef:  as

hello, world
`),
		ret: RawMessage{
			RawHeaders: []RawHeader{
				RawHeader{[]byte("a"), []byte("b")},
				RawHeader{[]byte("c"), []byte("d fdsa")},
				RawHeader{[]byte("ef"), []byte("as")},
			},
			Body: []byte(`hello, world
`),
		},
	},
}

func TestParseRaw(t *testing.T) {
	for _, pt := range parseRawTests {
		msg := pt.msg
		ret := pt.ret
		act, err := ParseRaw(msg)
		if err != nil {
			t.Errorf("ParseRaw returned error for %#V", string(msg))
		} else if !reflect.DeepEqual(act, ret) {
			t.Errorf("ParseRaw: incorrect result from %#V as %#V; expected %#V", string(msg), act, ret)
		}
	}
}

type processTest struct {
	name string
	raw  RawMessage
	ret  Message
}

var processTests = []processTest{}

func TestProcess(t *testing.T) {
	for _, pt := range processTests {
		act, err := Process(pt.raw)
		if err != nil {
			t.Errorf("Parse returned error for %s", pt.name)
		} else if !reflect.DeepEqual(act, pt.ret) {
			t.Errorf("Parse: incorrect result from %#V as %#V; expected %#V", pt.name, act, pt.ret)
		}
	}
}

type parseTest struct {
	msg []byte
	ret Message
}

var parseTests = []parseTest{
	parseTest{
		crlf(`

`),
		Message{
			FullHeaders: []Header{},
			OptHeaders: []Header{},
			Text: "\r\n",
		},
	},
	parseTest{
		crlf(`Subject: Hello, world

G'day, mate.
`),
		Message{
			FullHeaders: []Header{Header{"Subject", "Hello, world"}},
			OptHeaders: []Header{},
			Subject: "Hello, world",
			Text: "G'day, mate.\r\n",
		},
	},
}

func TestParse(t *testing.T) {
	for _, pt := range parseTests {
		msg := pt.msg
		ret := pt.ret
		act, err := Parse(msg)
		if err != nil {
			t.Errorf("Parse returned error for %#V", string(msg))
		} else if !reflect.DeepEqual(act, ret) {
			t.Errorf("Parse: incorrect result from %#V as %#V; expected %#V", string(msg), act, ret)
		}
	}
}
