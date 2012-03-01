package mail

import (
	"reflect"
	"testing"
)

type tokenTest struct {
	s string
	t []string
}

var tokenTests = []tokenTest{
	tokenTest{``, []string{}},
	tokenTest{`a`, []string{`a`}},
	tokenTest{`af&' al43`, []string{`af&'`, `al43`}},
	tokenTest{
		`"Joe Q. Public" <john.q.public@example.com>`,
		[]string{`"Joe Q. Public"`, `<`, `john.q.public`, `@`, `example.com`, `>`},
	},
	tokenTest{
		`"Giant; \"Big\" Box" <sysservices@example.net>`,
		[]string{
			`"Giant; \"Big\" Box"`,
			`<`,
			`sysservices`,
			`@`,
			`example.net`,
			`>`,
		},
	},
}

func TestTokenize(t *testing.T) {
	for _, tt := range tokenTests {
		o, e := tokenize([]byte(tt.s))
		if e != nil {
			t.Errorf("tokenize returned error for %#V", tt.s)
		} else {
			rt := []string{}
			for _, tok := range o {
				rt = append(rt, string(tok))
			}
			if !reflect.DeepEqual(rt, tt.t) {
				t.Errorf("tokenize(%#V) gave %#V; expected %#V", tt.s, rt, tt.t)
			}
		}
	}
}
