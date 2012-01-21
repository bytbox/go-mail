package mail

import (
	"testing"
)

type getHeadersTest struct {

}

func TestGetHeaders(t *testing.T) {

}

type splitHeadersTest struct {

}

func TestSplitHeader(t *testing.T) {

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
