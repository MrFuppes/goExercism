// Package atbash implements the atbash cypher.
package atbash

import (
	"strings"
)

// cleanString - a helper to remove everything but characters.
func cleanString(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return strings.ToLower(result.String())
}

// Atbash calculates the atbash cypher for a given string.
func Atbash(input string) string {
	var (
		collector []string
		pos       int
	)
	input = cleanString(input)
	for i, c := range input {
		if i%5 == 0 { // append a new string to the collector each 5 characters.
			collector = append(collector, "")
			pos++
		}
		if 'a' <= c && c <= 'z' {
			collector[pos-1] += string('z' - c + 'a')
		} else {
			collector[pos-1] += string(c)
		}
	}
	return strings.Join(collector, " ")
}
