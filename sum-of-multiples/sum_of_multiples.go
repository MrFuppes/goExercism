package summultiples

// SumMultiples - Given a number, find the sum of all the unique multiples
// of particular numbers up to but not including that number.
func SumMultiples(limit int, divisors ...int) (sum int) {
	var (
		m    int
		allM = make(map[int]bool)
	)

	for _, d := range divisors {
		if d == 0 {
			continue
		}
		for i := 1; ; i++ {
			m = i * d
			if m >= limit {
				break
			}
			allM[m] = true
		}
	}

	for m = range allM {
		sum += m
	}

	return sum
}
