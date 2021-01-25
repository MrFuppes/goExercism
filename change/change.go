package change

import (
	"errors"
)

var (
	errNegative    = errors.New("cannot use negative target value")
	errUnreachable = errors.New("cannot reach target with given coinset")
)

func sum(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return sum
}

func min(n, m int) int {
	if n < m {
		return n
	} else {
		return m
	}
}

// Change calculates the coins required to get a target amount
func Change(coins []int, target int) (selected []int, err error) {
	if target == 0 {
		return []int{}, nil
	}
	if target < 0 {
		return []int{}, errNegative
	}
	if target < coins[0] {
		return []int{}, errUnreachable
	}

	// try the naive approach first
	remains := target
	for i := len(coins) - 1; i > 0; i-- {
		quotient, remainder := remains/coins[i], remains%coins[i]
		if quotient > 0 {
			for j := quotient; j > 0; j-- {
				selected = append([]int{coins[i]}, selected...)
			}
		}
		remains = remainder
	}

	if remains == 0 {
		return selected, nil
	}

	// there could be other combinations... use a stolen brute-force approach
	// still fails "possible change without unit coins available"
	wallet := make(map[int][]int)
	for _, coin := range coins {
		for N := 1; N <= target; N++ {
			if coin == N {
				wallet[N] = []int{coin}
			}
			if coin < N {
				collection, _ := wallet[N-coin]
				collection = append(collection, coin)
				if sum(collection) != N {
					continue
				}
				if len(wallet[N]) == 0 || len(wallet[N]) > len(collection) {
					wallet[N] = collection
				}
			}
		}
	}
	if selected, ok := wallet[target]; ok {
		return selected, nil
	}
	return []int{}, errUnreachable
}
