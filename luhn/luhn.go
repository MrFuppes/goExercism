// Package luhn implements the Luhn algorithm
package luhn

import (
	"strconv"
	"strings"
)

// reverseString - a helper to reverse the input string
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Valid checks if a sequence of digits given as a string is valid according to
// the Luhn algorithm.
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1) // clean string from spaces
	if len(s) <= 1 {
		return false // too short...
	}

	var sum int

	for i, c := range reverseString(s) {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}

		if i%2 != 0 { // for every second element, check the double
			if v*2 > 9 {
				sum += v*2 - 9
			} else {
				sum += v * 2
			}
		} else {
			sum += v
		}

	}
	return sum%10 == 0
}
