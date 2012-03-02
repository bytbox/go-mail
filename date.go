package mail

import (
	"time"
)

var dateFormats = []string{
	`Mon, 02 Jan 2006 15:04 -0700`,
	`02 Jan 2006 15:04 -0700`,
	`Mon, 02 Jan 2006 15:04:05 -0700`,
	`02 Jan 2006 15:04:05 -0700`,

	`Mon, 02 Jan 2006 15:04 -0700 (MST)`,
	`02 Jan 2006 15:04 -0700 (MST)`,
	`Mon, 02 Jan 2006 15:04:05 -0700 (MST)`,
	`02 Jan 2006 15:04:05 -0700 (MST)`,

	`Mon, 2 Jan 2006 15:04 -0700`,
	`2 Jan 2006 15:04 -0700`,
	`Mon, 2 Jan 2006 15:04:05 -0700`,
	`2 Jan 2006 15:04:05 -0700`,

	`Mon, 2 Jan 2006 15:04 -0700 (MST)`,
	`2 Jan 2006 15:04 -0700 (MST)`,
	`Mon, 2 Jan 2006 15:04:05 -0700 (MST)`,
	`2 Jan 2006 15:04:05 -0700 (MST)`,
}

func ParseDate(s string) time.Time {
	for _, fmt := range dateFormats {
		t, e := time.Parse(fmt, s)
		if e == nil {
			return t
		}
	}
	return time.Now()
}
