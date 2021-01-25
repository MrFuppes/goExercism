// Package collatzconjecture implements functionality to calculate the 3x+1 problem.
package collatzconjecture

import "errors"

// CollatzConjecture calculates the number of steps to reach 1.
func CollatzConjecture(i int) (int, error) {
	if i <= 0 {
		return 0, errors.New("zero input not allowed")
	}
	var n int
	for {
		if i == 1 {
			break
		}
		if i%2 == 0 {
			i /= 2
		} else {
			i = i*3 + 1
		}
		n++
	}
	return n, nil
}
