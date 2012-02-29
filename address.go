// Address parsing

package mail

type Address interface {
	String() string
}

type MailboxAddr struct {
	name   string
	local  string
	domain string
}

func (ma *MailboxAddr) String() string {
	return ""
}

type GroupAddr struct {
}

func ParseAddress(bs []byte) (Address, error) {
	toks, err := tokenize(bs)
	if err != nil {
		return nil, err
	}

	// If this is a group, it must end in a ";" token.
	ltok := toks[len(toks)-1]
	if len(ltok) == 1 && ltok[0] == ';' {
		return nil, nil
	}
	return parseMailboxAddr(toks)
}

func parseMailboxAddr(ts []token) (*MailboxAddr, error) {
	ma := &MailboxAddr{}
	return ma, nil
}
