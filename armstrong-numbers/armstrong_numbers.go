package armstrong

// IsNumber checks if a given integer is an Armstrong number
func IsNumber(n int) bool {
	p := nDigits(n)
	m := n
	for m > 0 {
		n -= intPow(m%10, p)
		m /= 10
	}
	return n == 0
}

// intPow - a helper to calculate powers of integers
func intPow(n, p int) int {
	m := n
	for i := 1; i < p; i++ {
		n *= m
	}
	return n
}

// nDigits - a helper to determine how many digits an integer is long
func nDigits(i int) int {
	if i < 0 {
		i *= -1
	}
	switch {
	case i < 10:
		return 1
	case i < 100:
		return 2
	case i < 1000:
		return 3
	case i < 10000:
		return 4
	case i < 100000:
		return 5
	case i < 1000000:
		return 6
	case i < 10000000:
		return 7
	case i < 100000000:
		return 8
	case i < 1000000000:
		return 9
	}
	return -1
}
