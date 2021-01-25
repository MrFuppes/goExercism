package sieve

// Sieve uses the sieve of Eratosthenes to calculate all prime numbers up to a give limit.
func Sieve(limit int) (primes []int) {
	if limit == 1 {
		return primes
	}

	var (
		marked = make([]bool, limit+1)
		p      = 2
	)
	// initialize marekd; first two elements (0, 1) are false, rest is true
	for i := range marked[1 : len(marked)-1] {
		marked[i+2] = true
	}
	// the sieve: for each p, mark all multiples false
	for p*p <= limit {
		if marked[p] {
			for i := p * 2; i <= limit; i = i + p {
				marked[i] = false
			}
		}
		p++
	}

	for i, p := range marked {
		if p {
			primes = append(primes, i)
		}
	}

	return primes
}
