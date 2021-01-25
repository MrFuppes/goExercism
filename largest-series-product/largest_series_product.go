package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct calculates the largest product for n (spatially) consequtive
// digits from a series of digits supplied as a string
func LargestSeriesProduct(s string, n int) (int64, error) {
	if n < 0 {
		return -1, errors.New("span must be greater than zero")
	}
	if n > len(s) {
		return -1, errors.New("span must be smaller than string length")
	}
	if len(s) == 0 || n == 0 {
		return 1, nil
	}

	var (
		nbrs []int64
		p    int64
	)

	// convert to slice of int
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return -1, errors.New("digits input must only contain digits")
		}
		nbrs = append(nbrs, int64(c-'0'))
	}

	// now check all possible products
	for i := 0; i <= len(nbrs)-n; i++ {
		nbrs[0] = nbrs[i] // we can use the first entry of nbrs slice as tmp storage
		for j := 1; j < n; j++ {
			nbrs[0] *= nbrs[i+j]
		}
		if nbrs[0] > p {
			p = nbrs[0]
		}
	}
	return p, nil
}
