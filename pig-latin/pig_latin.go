package piglatin

import (
	"regexp"
	"strings"
)

// what is considered a vowel vs. consonant: regex overkill. even order matters.
var (
	vwls   = regexp.MustCompile("^([aeiou]|y[^aeiou]|xr)[a-z]*")
	cnsnts = regexp.MustCompile("^([^aeiou]?qu|[^aeiou]+?y|[^aeiou]+)([a-z]*)")
)

// Sentence translates each word in the input string to pig latin
func Sentence(input string) string {
	var (
		words = strings.Fields(input)
		idx   int
	)
	for i, w := range words {
		if vwls.MatchString(w) {
			words[i] = w + "ay"
		} else if x := cnsnts.FindStringSubmatchIndex(w); x != nil {
			idx = x[3]
			// a hack for n consonants + y:
			if len(w[:idx]) > 1 && strings.HasSuffix(w[:idx], "y") {
				idx--
			}
			words[i] = w[idx:] + w[:idx] + "ay"
		}
	}
	return strings.Join(words, " ")
}
