package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode creates an encoded version of the input string using square code method
func Encode(s string) string {

	if len(s) <= 1 { // precautions: input too short,
		return s // just return it
	}

	var (
		idx, nRows, nCols int
		letters           []rune
	)

	// normalze and downcase the input
	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			letters = append(letters, c)
		}
	}

	// determine number of columns and rows for the input
	nRows = int(math.Round(math.Sqrt(float64(len(letters)))))
	nCols = int(math.Ceil(float64(len(letters)) / float64(nRows)))

	// create the output column-wise
	encoded := make([]string, nCols)
	for r := 0; r < nCols; r++ {
		for c := 0; c < nRows; c++ {
			idx = r + nCols*c
			if idx < len(letters) {
				encoded[r] += string(letters[idx])
			} else {
				encoded[r] += " " // pad with spaces if idx >= len(letters)
			}
		}
	}

	return strings.Join(encoded, " ")
}
