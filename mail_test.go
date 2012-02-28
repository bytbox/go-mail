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
	ret RawMessage
}

var parseTests = []parseTest{
	parseTest{
		msg: crlf(`
`),
		ret: RawMessage{
			RawHeaders: []RawHeader{},
			Body:       crlf(""),
		},
	},
	parseTest{
		msg: crlf(`
ab
c
`),
		ret: RawMessage{
			RawHeaders: []RawHeader{},
			Body:       crlf(`ab
c
`),
		},
	},
	parseTest{
		msg: crlf(`a: b

`),
		ret: RawMessage{
			RawHeaders: []RawHeader{RawHeader{crlf("a"), crlf("b")}},
			Body:       crlf(""),
		},
	},
	parseTest{
		msg: crlf(`a: b
c: def
 hi

`),
		ret: RawMessage{
			RawHeaders: []RawHeader{
				RawHeader{crlf("a"), crlf("b")},
				RawHeader{crlf("c"), crlf("def hi")},
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
		ret: RawMessage{
			RawHeaders: []RawHeader{
				RawHeader{crlf("a"), crlf("b")},
				RawHeader{crlf("c"), crlf("d fdsa")},
				RawHeader{crlf("ef"), crlf("as")},
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
		ret: RawMessage{
			RawHeaders: []RawHeader{
				RawHeader{[]byte("a"), []byte("b")},
				RawHeader{[]byte("c"), []byte("d fdsa")},
				RawHeader{[]byte("ef"), []byte("as")},
			},
			Body:       []byte(`hello, world
`),
		},
	},
}

func TestParseRaw(t *testing.T) {
	for _, pt := range parseTests {
		msg := pt.msg
		ret := pt.ret
		act, err := ParseRaw(msg)
		if err != nil {
			t.Errorf("ParseRaw returned error for %#v", string(msg))
		} else if !reflect.DeepEqual(act, ret) {
			t.Errorf("ParseRaw: incorrectly result from %#v as %#v; expected %#v", string(msg), act, ret)
		}
	}
}
