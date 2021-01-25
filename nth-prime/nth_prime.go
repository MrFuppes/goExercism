package prime

import "math"

// Nth returns the nth prime number
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	var p, count = 2, 0
	for ; count < n; p++ {
		if isPrime(p) {
			count++
		}
	}
	return p - 1, true
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return n == 2
	}
	if n%3 == 0 {
		return n == 3
	}
	for i, step := 5, 4; i < int(math.Sqrt(float64(n)))+1; i = i + step {
		if n%i == 0 {
			return false
		}
		step = 6 - step
	}
	return true
}
