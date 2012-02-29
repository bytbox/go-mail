// Address parsing

package mail

type Address interface {
	String() string
}

type MailboxAddr struct {

}

type GroupAddr struct {

}

func ParseAddress(str string) (Address, error) {
	return nil, nil
}
