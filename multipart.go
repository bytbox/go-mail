// Handle multipart messages.

package mail

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
)

type Part struct {
	Type string
	Data []byte
	Headers map[string][]string
}

// Parse the body of a message, using the given content-type. If the content
// type is multipart, the parts slice will contain an entry for each part
// present; otherwise, it will contain a single entry, with the entire (raw)
// message contents.
func parseBody(ct string, body []byte) (parts []Part, err error) {
	mt, ps, err := mime.ParseMediaType(ct)
	if err != nil { return }
	if mt != "multipart/alternative" {
		parts = append(parts, Part{ct, body, nil})
		return
	}
	boundary, ok := ps["boundary"]
	if !ok {
		return nil, errors.New("multipart specified without boundary")
	}
	r := multipart.NewReader(bytes.NewReader(body), boundary)
	p, err := r.NextPart()
	for err == nil {
		data, _ := ioutil.ReadAll(p) // ignore error
		ct := p.Header["Content-Type"]

		part := Part{ct[0], data, p.Header}
		parts = append(parts, part)
		p, err = r.NextPart()
	}
	if err == io.EOF { err = nil }
	return
}
