package anagram

import (
	"sort"
	"strings"
)

// Detect checks if a slice of strings contains an anagram of given subject.
// returns a slice of all anagrams found.
func Detect(subject string, candidates []string) (anagrams []string) {
	var want, have []rune

	// convert the subject to a slice of lower-case runes, and sort.
	subject = strings.ToLower(subject)
	want = []rune(subject)
	sort.Slice(want, func(i, j int) bool { return want[i] < want[j] })

outer:
	for i, s := range candidates {
		if len(subject) != len(s) {
			continue // skip if different number of characters
		}
		s = strings.ToLower(s)
		if s == subject {
			continue // skip if same string
		}
		have = []rune(s)
		sort.Slice(have, func(i, j int) bool { return have[i] < have[j] })
		for j, c := range want {
			if c != have[j] {
				continue outer // continue with the outer loop if no anagram
			}
		}
		// all characters match, append anagram
		anagrams = append(anagrams, candidates[i])
	}

	return anagrams
}

// // Detect check whether word s has an anagram in the candidates
// func Detect(s string, candidates []string) (rs []string) {
// 	cs := cleanString(s)
// 	for _, v := range candidates {
// 		if cleanString(v) == cs && (strings.ToLower(s) != strings.ToLower(v)) {
// 			rs = append(rs, v)
// 		}
// 	}
// 	return rs
// }

// func cleanString(s string) string {
// 	s = strings.ToLower(s)
// 	slice := strings.Split(s, "")
// 	sort.Strings(slice)
// 	return strings.Join(slice, "")
// }
// type counts [26]int

// func count(s string) (c counts) {
// 	for _, r := range s {
// 		if unicode.IsLetter(r) {
// 			c[unicode.ToLower(r)-'a']++
// 		}
// 	}
// 	return
// }

// // Detect selects the correct sublist given a word and a list of possible anagrams.
// func Detect(subject string, candidates []string) (sublist []string) {
// 	sc := count(subject)
// 	for _, c := range candidates {
// 		if len(subject) != len(c) || strings.EqualFold(subject, c) {
// 			continue
// 		}
// 		cc := count(c)
// 		if sc == cc {
// 			sublist = append(sublist, c)
// 		}
// 	}
// 	return
// }
