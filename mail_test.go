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

type parseTest struct {
	msg []byte
	ret Message
}

var parseTests = []parseTest{
	parseTest{
		msg: crlf(`
`),
		ret: Message{
			RawHeaders: []Header{},
			Body:       crlf(""),
		},
	},
	parseTest{
		msg: crlf(`
ab
c
`),
		ret: Message{
			RawHeaders: []Header{},
			Body:       crlf(`ab
c
`),
		},
	},
	parseTest{
		msg: crlf(`a: b

`),
		ret: Message{
			RawHeaders: []Header{Header{crlf("a"), crlf("b")}},
			Body:       crlf(""),
		},
	},
	parseTest{
		msg: crlf(`a: b
c: def
 hi

`),
		ret: Message{
			RawHeaders: []Header{
				Header{crlf("a"), crlf("b")},
				Header{crlf("c"), crlf("def hi")},
			},
			Body:       crlf(``),
		},
	},
	parseTest{
		msg: crlf(`a: b
c: d fdsa
ef:  as

hello, world
`),
		ret: Message{
			RawHeaders: []Header{
				Header{crlf("a"), crlf("b")},
				Header{crlf("c"), crlf("d fdsa")},
				Header{crlf("ef"), crlf("as")},
			},
			Body:       crlf(`hello, world
`),
		},
	},
	parseTest{
		msg: []byte(`a: b
c: d fdsa
ef:  as

hello, world
`),
		ret: Message{
			RawHeaders: []Header{
				Header{[]byte("a"), []byte("b")},
				Header{[]byte("c"), []byte("d fdsa")},
				Header{[]byte("ef"), []byte("as")},
			},
			Body:       []byte(`hello, world
`),
		},
	},
}

func TestParse(t *testing.T) {
	for _, pt := range parseTests {
		msg := pt.msg
		ret := pt.ret
		act, err := Parse(msg)
		if err != nil {
			t.Errorf("parse returned error for %#v", string(msg))
		} else if !reflect.DeepEqual(act, ret) {
			t.Errorf("incorrectly parsed message %#v as %#v; expected %#v", string(msg), act, ret)
		}
	}
}
