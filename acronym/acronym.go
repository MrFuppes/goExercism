// Package acronym implements functionality to build acronyms.
package acronym

import (
	"strings"
)

// Abbreviate returns and acronym for the input string (all starting letters capitalized).
func Abbreviate(s string) string {
	// need to split the input on multiple chars;
	// using from the docs: https://golang.org/pkg/strings/#FieldsFunc
	f := func(c rune) bool {
		return c == ' ' || c == '-'
	}
	parts := strings.FieldsFunc(s, f)

	acro := ""
	for _, v := range parts {
		// remove surrounding stuff...
		v = strings.Trim(v, "_- ")
		// check if something remains after trimming:
		if len(v) > 0 {
			acro += strings.ToUpper(string(v[0]))
		}
	}
	return acro
}
