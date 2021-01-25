// Package hamming implements the Hamming distance calculation
package hamming

import "errors"

// Distance - calculate the Hamming distance between two strands of DNA, a and b.
func Distance(a, b string) (int, error) {
	// ensure equal length
	if len(a) != len(b) {
		return 0, errors.New("strands of unequal length not allowed")
	}
	var dist int
	// compare the letters...
	for i, l := range a { // ranging over string gives letters as type rune
		if l != rune(b[i]) { // string at index yields byte; need to cast to rune for comp.
			dist++
		}
	}
	return dist, nil
}
