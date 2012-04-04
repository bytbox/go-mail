// Handle multipart messages.

package mail

type Part struct {
	Type string
	Data []byte
}

// Parse the body of a message, using the given content-type. If the content
// type is multipart, the parts slice will contain an entry for each part
// present; otherwise, it will contain a single entry, with the entire (raw)
// message contents. The returned text will be an appropriate string
// representation of the first part of the message.
func ParseBody(ct string, body []byte) (text string, textct string, parts []Part) {
	return
}
