// Package scrabble implements functionality for the scrabble game
package scrabble

import (
	"strings"
)

// scoreMap attributes scores to specific letter subsets
var scoreMap = map[int]string{
	1:  "AEIOULNRST",
	2:  "DG",
	3:  "BCMP",
	4:  "FHVWY",
	5:  "K",
	8:  "JX",
	10: "QZ",
}

// Score calculates the scrabble score for a given string s
func Score(s string) int {
	var result int
	for _, r := range strings.ToUpper(s) {
		for score, str := range scoreMap {
			if strings.ContainsRune(str, r) {
				result += score
				break // since a certain letter occures only once in the score map, we can skip further iterations
			}
		}
	}
	return result
}
