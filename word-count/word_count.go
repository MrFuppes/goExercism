package wordcount

import (
	"strings"
)

// Frequency - a map to count words
type Frequency map[string]int

// cleanString - a helper to remove everything but characters, spaces and apostrophes.
// all other characters are replaced with space.
func cleanString(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') || b == '\'' {
			result.WriteByte(b)
		} else {
			result.WriteByte(' ')
		}
	}
	return result.String()
}

// WordCount takes a string and counts the words in it.
func WordCount(input string) Frequency {
	var count = make(Frequency)
	// split string on space; use FieldsFunc to ignore repeated separators
	splitFn := func(c rune) bool { return c == ' ' }
	for _, s := range strings.FieldsFunc(cleanString(input), splitFn) {
		if strings.HasPrefix(s, "'") || strings.HasSuffix(s, "'") {
			s = strings.Replace(s, "'", "", 2)
		}
		count[strings.ToLower(s)]++
	}
	return count
}
