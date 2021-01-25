// Package darts implements scoring system for darts game.
package darts

import "math"

// Score - calculate the score for a dart thrown at coordinates (x, y).
// radius of hit circles are fixed to 10 (outer), 5 (middle), 1 (inner).
func Score(x, y float64) int {
	// calculate distance to origin
	dist := math.Hypot(x, y)
	// based on that distance, return the number of points
	switch {
	case dist <= 1:
		return 10
	case dist <= 5:
		return 5
	case dist <= 10:
		return 1
	default:
		return 0
	}
}
