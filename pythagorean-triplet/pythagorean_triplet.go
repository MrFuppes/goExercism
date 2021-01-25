// Package pythagorean implements functions to calculate Pythagorean Triplets
package pythagorean

import "sort"

// Triplet - a 3-element array holding the Pythagorean triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	// with u, v being natural numbers and u > v > 0:
	// x = u**2 - v**2
	// y = 2 u v
	// z = u**2 + v**2
	// https://en.wikipedia.org/wiki/Pythagorean_triple
	var (
		u       = 2
		x, y, z int
		tMap    = make(map[Triplet]bool)
		result  []Triplet
	)

	for z < max {
		for v := 1; v <= u; v++ {
			for k := 1; k < max/(u*u+v*v)+1; k++ {
				x = k * (u*u - v*v)
				y = k * (2 * u * v)
				z = k * (u*u + v*v)
				if x < min || y < min {
					continue
				}
				if x < y {
					tMap[Triplet{x, y, z}] = true
				} else {
					tMap[Triplet{y, x, z}] = true
				}
			}
		}
		u++
	}

	for t := range tMap { // extract all the map's keys to a slice (desired output)
		result = append(result, t)
	}

	sort.Slice(result, func(i, j int) bool { // since map is unordered, we need to sort
		return result[i][0] < result[j][0]
	})

	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c (the perimeter) is equal to p.
// The three elements of each returned triplet are in order,
// t[0] <= t[1] <= t[2], and the list of triplets is in lexicographic order.
func Sum(p int) []Triplet {
	var t []Triplet
	// first element must be < p/3
	for x := 1; x < p/3+1; x++ {
		// second element must be <= p/2 and at least x+1
		for y := x + 1; y < p/2+1; y++ {
			z := p - (x + y)
			if x*x+y*y == z*z {
				t = append(t, Triplet{x, y, z})
			}
		}
	}
	return t
}
