package prime

// Factors calculates prime factors of a given integer n
func Factors(n int64) (factors []int64) {
	if n < 2 {
		return []int64{}
	}

	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}

	for i := int64(3); i*i <= n; i = i + 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}
