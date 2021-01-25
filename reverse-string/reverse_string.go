package reverse

import "strings"

// Reverse inverts the order of letters in a string.
// Null-bytes are removed before output.
func Reverse(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		result[len(s)-(i+1)] = r
	}
	return strings.ReplaceAll(string(result), "\x00", "")
}
