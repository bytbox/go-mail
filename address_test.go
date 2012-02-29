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
		nil,
	},
	parseAddressTest{
		`Mary Smith <mary@x.test>`,
		nil,
	},
	parseAddressTest{
		`jdoe@example.org`,
		nil,
	},
	parseAddressTest{
		`Who? <one@y.test>`,
		nil,
	},
	parseAddressTest{
		`<boss@nil.test>`,
		nil,
	},
	parseAddressTest{
		`"Giant; \"Big\" Box" <sysservices@example.net>`,
		nil,
	},
	parseAddressTest{
		`Pete <pete@silly.example>`,
		nil,
	},
	parseAddressTest{
		`A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;`,
		nil,
	},
	parseAddressTest{
		`Undisclosed recipients:;`,
		nil,
	},
}

func TestParseAddress(t *testing.T) {
	for _, pt := range parseAddressTests {
		address, err := ParseAddress(pt.addrStr)
		if err != nil {
			t.Errorf("ParseAddress returned error for %#v", address)
		} else if !reflect.DeepEqual(address, pt.addrRes) {
			t.Errorf(
				"ParseAddress: incorrect result for %#v: gave %#v; expected %#v",
				pt.addrStr, address, pt.addrRes)
		}
	}
}
