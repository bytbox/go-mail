package mail

import (
	"strings"
	"testing"
)

// Converts all newlines to CRLFs.
func crlf(s string) string {
	return strings.Replace(s, "\n", "\r\n", -1)
}

type getHeadersTest struct {

}

var getHeadersTests = []getHeadersTest{

}

func TestGetHeaders(t *testing.T) {

}

type splitHeadersTest struct {
	orig, key, val string
}

var splitHeadersTests = []splitHeadersTest{
	{
		`a: b`,
		`a`, `b`,
	},
	{
		`A1: cD`,
		`A1`, `cD`,
	},
	{
		crlf(`ab: cd
 ef`),
		`ab`, `cd ef`,
	},
	{
		crlf(`ab: cd
	ef
	dh`),
		`ab`, `cd	ef	dh`,
	},
}

func TestSplitHeader(t *testing.T) {
	for i, ht := range splitHeadersTests {
		k, v := splitHeader(ht.orig)
		if k != ht.key || v != ht.val {
			t.Errorf(`%d. splitHeader gave ("%s", "%s"), wanted ("%s", "%s")`,
				i, k, v, ht.key, ht.val)
		}
	}
}

type hdr struct {
	key, val string
}

type parseTest struct {
	orig string
	hdrs []string
	cont string
}

var parseTests = []parseTest{

}

func TestParse(t *testing.T) {
	for _, pt := range parseTests {
		m := Parse(pt.orig)
		m = m
	}
}
