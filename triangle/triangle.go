// Package triangle implements triangle checks.
package triangle

import "math"

// Kind - a string to hold the type of a triangle.
type Kind string

const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// checkFinite a helper function to check if inputs are finite values.
func checkFinite(v float64) bool {
	return !(math.IsNaN(v) || math.IsInf(v, 0))
}

// KindFromSides determines the kind of a triangle based on its three sides a, b and c.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case !(checkFinite(a) && checkFinite(b) && checkFinite(c)):
		k = NaT
	case (a == 0 && b == 0 && c == 0):
		k = NaT
	case (a > (b+c) || b > (a+c) || c > (a+b)):
		k = NaT
	case (a == b && b == c):
		k = Equ
	case (a == b || a == c || b == c):
		k = Iso
	default:
		k = Sca
	}

	return k
}
