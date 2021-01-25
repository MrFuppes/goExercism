package encode

import (
	"fmt"
	"strconv"
	"strings"
)

// RunLengthEncode returns a run-length-encoded version of the input string
func RunLengthEncode(s string) (encoded string) {
	if len(s) <= 1 {
		return s
	}

	chars := []rune(s)
	count := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[i-1] {
			prefix := ""
			if count > 1 {
				prefix = fmt.Sprintf("%v", count)
			}
			encoded += prefix + string(chars[i-1])
			count = 0
		}
		count++
		if i == len(chars)-1 && chars[i] != chars[i-1] { // edge case: last letter different
			encoded += string(chars[i])
		}
		if i == len(chars)-1 && chars[i] == chars[i-1] { // edge case: last letter as previous
			encoded += fmt.Sprintf("%v", count) + string(chars[i])
		}
	}

	return encoded
}

// RunLengthDecode decodes to a normal string
func RunLengthDecode(s string) (decoded string) {
	if len(s) <= 1 {
		return s
	}

	count := ""
	for _, c := range s {
		if '0' <= c && c <= '9' {
			count += string(c)
		} else { // can use a simple else here since numbers always precede chars
			n := 1
			if count != "" {
				n, _ = strconv.Atoi(count)
			}
			decoded += strings.Repeat(string(c), n) //string(c)
			count = ""
		}
	}

	return decoded
}
