package mail

import (
	"reflect"
	"testing"
)

// We use examples from RFC5322 as our test suite.

type parseAddressTest struct {
	addrStr string
	addrRes Address
}

var parseAddressTests = []parseAddressTest{
	parseAddressTest{
		`"Joe Q. Public" <john.q.public@example.com>`,
		MailboxAddr{`"Joe Q. Public"`, `john.q.public`, `example.com`},
	},
	parseAddressTest{
		`Mary Smith <mary@x.test>`,
		MailboxAddr{`Mary Smith`, `mary`, `x.test`},
	},
	parseAddressTest{
		`jdoe@example.org`,
		MailboxAddr{``, `jdoe`, `example.org`},
	},
	parseAddressTest{
		`Who? <one@y.test>`,
		MailboxAddr{`Who?`, `one`, `y.test`},
	},
	parseAddressTest{
		`<boss@nil.test>`,
		MailboxAddr{``, `boss`, `nil.test`},
	},
	parseAddressTest{
		`"Giant; \"Big\" Box" <sysservices@example.net>`,
		MailboxAddr{`"Giant; \"Big\" Box"`, `sysservices`, `example.net`},
	},
	parseAddressTest{
		`Pete <pete@silly.example>`,
		MailboxAddr{`Pete`, `pete`, `silly.example`},
	},
	parseAddressTest{
		`A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;`,
		GroupAddr{
			`A Group`,
			[]MailboxAddr{
				MailboxAddr{`Ed Jones`, `c`, `a.test`},
				MailboxAddr{``, `joe`, `where.test`},
				MailboxAddr{`John`, `jdoe`, `one.test`},
			},
		},
	},
	parseAddressTest{
		`Undisclosed recipients:;`,
		GroupAddr{`Undisclosed recipients`, []MailboxAddr{}},
	},
	parseAddressTest{
		`Undisclosed recipients:      ;`,
		GroupAddr{`Undisclosed recipients`, []MailboxAddr{}},
	},
}

func TestParseAddress(t *testing.T) {
	for _, pt := range parseAddressTests {
		address, err := ParseAddress([]byte(pt.addrStr))
		if err != nil {
			t.Errorf("ParseAddress returned error for %#V", pt.addrStr)
		} else if !reflect.DeepEqual(address, pt.addrRes) {
			t.Errorf(
				"ParseAddress: incorrect result for %#V: gave %#V; expected %#V",
				pt.addrStr, address, pt.addrRes)
		}
	}
}
