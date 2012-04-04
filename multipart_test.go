package mail

import (
	"reflect"
	"testing"
)

type parseBodyTest struct {
	ct   string
	body []byte
	rps  []Part
}

var parseBodyTests = []parseBodyTest{
	parseBodyTest{
		ct: "text/plain",
		body: []byte(`This is some text.`),
		rps: []Part{
			Part{"text/plain", []byte("This is some text."), nil},
		},
	},
	parseBodyTest{
		ct: "multipart/alternative; boundary=90e6ba1efd30b0013a04b8d4970f",
		body: []byte(`--90e6ba1efd30b0013a04b8d4970f
Content-Type: text/plain; charset=ISO-8859-1

Some text.
--90e6ba1efd30b0013a04b8d4970f
Content-Type: text/html; charset=ISO-8859-1
Content-Transfer-Encoding: quoted-printable

Some other text.
--90e6ba1efd30b0013a04b8d4970f--
`),
		rps: []Part{
			Part{
				"text/plain; charset=ISO-8859-1",
				[]byte("Some text."),
				map[string][]string{
					"Content-Type": []string{
						"text/plain; charset=ISO-8859-1",
					},
				},
			},
			Part{
				"text/html; charset=ISO-8859-1",
				[]byte("Some other text."),
				map[string][]string{
					"Content-Type": []string{
						"text/html; charset=ISO-8859-1",
					},
					"Content-Transfer-Encoding": []string{
						"quoted-printable",
					},
				},
			},
		},
	},
}

func TestParseBody(t *testing.T) {
	for _, pt := range parseBodyTests {
		parts, e := parseBody(pt.ct, pt.body)
		if e != nil {
			t.Errorf("parseBody returned error for %#v: %#v", pt, e)
		} else if !reflect.DeepEqual(parts, pt.rps) {
			t.Errorf(
				"parseBody: incorrect result for %#V: \n%#V\nvs.\n%#V",
				pt, parts, pt.rps)
		}
	}
}

