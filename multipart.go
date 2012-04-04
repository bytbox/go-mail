// Handle multipart messages.

package mail

import (
	"errors"
	"mime"
	//"mime/multipart"
)

type Part struct {
	Type string
	Data []byte
}

// Parse the body of a message, using the given content-type. If the content
// type is multipart, the parts slice will contain an entry for each part
// present; otherwise, it will contain a single entry, with the entire (raw)
// message contents.
func parseBody(ct string, body []byte) (parts []Part, err error) {
	mt, ps, err := mime.ParseMediaType(ct)
	if err != nil { return }
	if mt != "multipart/alternative" {
		parts = append(parts, Part{ct, body})
		return
	}
	_, ok := ps["boundary"]
	if !ok {
		return nil, errors.New("multipart specified without boundary")
	}
	return
}
