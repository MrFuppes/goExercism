// Package raindrops implements functionality to convert drops to sounds
package raindrops

import (
	"strconv"
)

// sounds and soundMul - a mapping of multiples of drops to sounds;
// not using a map here since order whould not be guaranteed with a map.
var (
	soundMul = [...]int{3, 5, 7}
	sounds   = [...]string{"Pling", "Plang", "Plong"}
)

// Convert returns the sound for a given number of drops
func Convert(n int) string {
	var out string
	for i, v := range soundMul {
		if n%v == 0 {
			out += sounds[i]
		}
	}

	if out == "" {
		return strconv.Itoa(n)
	}
	return out
}
