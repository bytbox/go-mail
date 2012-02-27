package mail

import (
	"reflect"
	"strings"
	"testing"
)

// Converts all newlines to CRLFs.
func crlf(s string) string {
	return strings.Replace(s, "\n", "\r\n", -1)

}

func crlfb(s string) []byte {
	return []byte(crlf(s))
}

type parseTest struct {
	msg []byte
	ret Message
}

var parseTests = []parseTest{
	parseTest{
		msg: crlfb(`
`),
		ret: Message{
			RawHeaders: []Header{},
			Body:       "",
		},
	},
	parseTest{
		msg: crlfb(`
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
		msg: crlfb(`a: b

`),
		ret: Message{
			RawHeaders: []Header{Header{"a", "b"}},
			Body:       "",
		},
	},
	parseTest{
		msg: crlfb(`a: b
c: def
 hi

`),
		ret: Message{
			RawHeaders: []Header{Header{"a", "b"}, Header{"c", "def hi"}},
			Body:       "",
		},
	},
	parseTest{
		msg: crlfb(`a: b
c: d fdsa
ef:  as

hello, world
`),
		ret: Message{
			RawHeaders: []Header{
				Header{"a", "b"},
				Header{"c", "d fdsa"},
				Header{"ef", "as"},
			},
			Body:       "hello, world\r\n",
		},
	},
}

func TestParse(t *testing.T) {
	for _, pt := range parseTests {
		msg := pt.msg
		ret := pt.ret
		act, err := Parse(msg)
		if err != nil {
			t.Errorf("parse returned error")
		} else if !reflect.DeepEqual(act, ret) {
			t.Errorf("incorrectly parsed message %#v as %#v; expected %#v", string(msg), act, ret)
		}
	}
}
