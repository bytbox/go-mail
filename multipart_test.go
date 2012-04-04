package mail

import (
	"testing"
)

type parseBodyTest struct {
	ct   string
	body []byte
	rt   string
	tct  string
	rps  []Part
}

var parseBodyTests = []parseBodyTest{
	parseBodyTest{
		ct: "multipart/alternative; boundary=90e6ba1efd30b0013a04b8d4970f",
		body: []byte(`
--90e6ba1efd30b0013a04b8d4970f
Content-Type: text/plain; charset=ISO-8859-1

Some text.
--90e6ba1efd30b0013a04b8d4970f
Content-Type: text/html; charset=ISO-8859-1
Content-Transfer-Encoding: quoted-printable

Some other text.
--90e6ba1efd30b0013a04b8d4970f--
`),
	},
}

func TestParseBody(t *testing.T) {

}

