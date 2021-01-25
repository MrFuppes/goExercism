package phonenumber

import (
	"errors"
	"strings"
	"unicode"
)

// Number cleans the input
func Number(s string) (string, error) {
	// slice country code
	if strings.HasPrefix(s, "+1") {
		s = s[2:len(s)]
	}
	if strings.HasPrefix(s, "1") {
		s = s[1:len(s)]
	}
	// remove separators
	r := strings.NewReplacer(" ", "", "(", "", ")", "", "-", "", ".", "")
	s = r.Replace(s)
	// without country code, must now be 10 characters
	if len(s) != 10 {
		return "", errors.New("invalid number of input characters")
	}
	// check if only numbers remain
	for _, c := range s {
		if !unicode.IsNumber(c) {
			return "", errors.New("invalid character in input")
		}
	}
	// check area code, must not be 0
	if s[0] == '0' || s[0] == '1' {
		return "", errors.New("area code must not be 0 or 1")
	}
	// check exchange code, must not be 0 or 1
	if s[3] == '0' || s[3] == '1' {
		return "", errors.New("exchange code must not be 0 or 1")
	}
	return s, nil
}

// Format calls the cleaner and formats the cleaned input to `(xxx) xxx-xxxx`
func Format(s string) (string, error) {
	s, e := Number(s)
	if e != nil {
		return s, e
	}
	return "(" + s[0:3] + ") " + s[3:6] + "-" + s[6:10], nil
}

// AreaCode extracts the area code
func AreaCode(s string) (string, error) {
	s, e := Number(s)
	if e != nil {
		return s, e
	}
	return s[0:3], nil
}
