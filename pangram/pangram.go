package pangram

import (
	"strings"
	"unicode"
)

// IsPangram checks if a given string s is a pangram
func IsPangram(s string) bool {
	var seen = make(map[rune]bool, 26)
	for _, c := range strings.ToUpper(s) {
		if unicode.IsLetter(c) {
			seen[c] = true
		}
	}
	return len(seen) == 26
}
