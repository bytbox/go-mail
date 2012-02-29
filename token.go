package mail

import (
	"bytes"
	"errors"
	"regexp"
)

// Regular expressions that correspond roughly to the syntax described by
// RFC5322. We're a bit loose here, so we might succeed in parsing material
// that the RFC considers invalid.
var (
	dotAtomR = regexp.MustCompile("^[a-zA-Z0-9!#$%&`*+\\-/=?^_'{|}~][a-zA-Z0-9.!#$%&`*+\\-/=?^_'{|}~]+")
	atomR    = regexp.MustCompile("^[a-zA-Z0-9!#$%&`*+\\-/=?^_'{|}~]+")
	specialR = regexp.MustCompile(`^[()<>\[\]:;@\,."]`)
	qStringR = regexp.MustCompilePOSIX(`^"([^"]|\\")*"`)
)

type token []byte

func try(re *regexp.Regexp, s []byte) int {
	is := re.FindIndex(s)
	if is == nil {
		return 0
	}
	if is[0] != 0 {
		return 0
	}
	return is[1]
}

func tokenize(s []byte) (ts []token, err error) {
Next:
	s = bytes.TrimSpace(s)
	if len(s) == 0 {
		return
	}
	for _, r := range []*regexp.Regexp{dotAtomR, atomR, qStringR, specialR} {
		i := try(r, s)
		if i > 0 {
			ts = append(ts, s[0:i])
			s = s[i:]
			goto Next
		}
	}
	return nil, errors.New("unidentifiable token")
}
