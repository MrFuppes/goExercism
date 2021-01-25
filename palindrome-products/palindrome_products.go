package palindrome

import (
	"errors"
	"strconv"
)

// Product - the palindrome product and its factors
type Product struct {
	Product        int
	Factorizations [][2]int
}

// IsPalindrome reports whether s reads the same forward and backward.
// A helper from the gopl book.
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Products returns the minimum and maximum palindrome product for a given range of numbers [from, to].
func Products(from, to int) (min, max Product, err error) {
	if from > to {
		return min, max, errors.New("fmin > fmax...") // seriously, a prescribed error string?
	}
	for i := from; i <= to; i++ {
		for j := i; j <= to; j++ {
			p := i * j
			if !IsPalindrome(strconv.Itoa(p)) {
				continue
			}
			if p > max.Product {
				max = Product{p, [][2]int{{i, j}}}
			} else if p == max.Product {
				max.Factorizations = append(max.Factorizations, [2]int{i, j})
			}
			if min.Product == 0 || p < min.Product {
				min = Product{p, [][2]int{{i, j}}}
			} else if p == min.Product {
				min.Factorizations = append(min.Factorizations, [2]int{i, j})
			}

		}
	}
	if min.Product == 0 && max.Product == 0 {
		return min, max, errors.New("no palindromes...") // seriously, another prescribed error string?
	}
	return min, max, nil
}
