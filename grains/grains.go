// Package grains implements the Persian chess board
package grains

import "errors"

// Square returns the number of grains on the nth square of the board
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("n must be 1 <= n < 65")
	}
	return uint64(1) << (n - 1), nil
}

// Total returns the sum of grains on all 64 squares
func Total() uint64 {
	return ^uint64(0)
}
