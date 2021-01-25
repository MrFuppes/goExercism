// Package isogram implements functionality to check for isograms
package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram checks if a string is an isogram
func IsIsogram(s string) bool {
	if len(strings.Trim(s, "- ")) == 0 {
		return true
	}
	// make sure to compare only lower-case chars:
	s = strings.ToLower(s)

	// split the string in case there are spaces or hyphens:
	parts := regexp.MustCompile("[\\-\\s]+").Split(s, -1)

	for _, substr := range parts {
		// use a map to count occurances of each letter:
		runeCounts := map[rune]int{}
		for _, r := range substr {
			runeCounts[r]++
			if runeCounts[r] > 1 {
				return false
			}
		}
	}

	return true
}
