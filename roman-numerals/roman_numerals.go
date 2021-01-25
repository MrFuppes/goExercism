// Package romannumerals implements conversion between arabic and roman numerals.
package romannumerals

import (
	"errors"
	"strings"
)

// roman, arabic: mappings of roman to arabic numerals
// I started to code the rules for combinations of the letters but it was too painful -
// ... so I hard-coded them.
var (
	roman  = [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabic = [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

// ToRomanNumeral converts a decimal number (arabic numerals) to roman numerals.
func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", errors.New("numbers less than one or greater than 3000 are not allowed")
	}
	var result string
	for i, v := range arabic {
		frac := n / v
		result += strings.Repeat(roman[i], frac)
		n -= v * frac
	}
	return result, nil
}
