package bookstore

import "sort"

// the base price of one book
const basePrice = 800

// relative discount for groups of n distinct books
var discount = map[int]float64{
	2: 0.05,
	3: 0.10,
	4: 0.20,
	5: 0.25,
}

// Cost calculates the cost for the books in the basket
func Cost(basket []int) (cost int) {
	// determine number of books for each of the five different
	counts := make(map[int]int)
	for _, item := range basket {
		counts[item]++
	}
	// make an ordered basket with five indices, each representing one of the five different
	orderedBasket := []int{}
	for i := 1; i <= 5; i++ {
		if c, ok := counts[i]; ok {
			orderedBasket = append(orderedBasket, c)
		} else {
			orderedBasket = append(orderedBasket, 0)
		}
	}
	return int(groupsCosts(orderedBasket))
}

// groupedsCosts - a helper to calculate prices for groups. recursively determines
// prices for possible group combinations. returns the minimum combined cost.
func groupsCosts(orderedBasket []int) (cost float64) {
	for _, v := range orderedBasket {
		cost += float64(v) * basePrice
	}
	for groupSize := 5; groupSize > 1; groupSize-- {
		// make a temporary basket to work with, do not modify original
		basketCopy := make([]int, len(orderedBasket))
		copy(basketCopy, orderedBasket)
		nUniq := 0
		for _, v := range basketCopy {
			if v > 0 {
				nUniq++
			}
		}
		if groupSize <= nUniq {
			var groupBasket []int
			sort.Sort(sort.Reverse(sort.IntSlice(basketCopy)))
			for _, v := range basketCopy[:groupSize] {
				groupBasket = append(groupBasket, v-1)
			}
			copy(basketCopy[:groupSize], groupBasket)
			// cost of current group
			groupCost := float64(basePrice) * float64(groupSize) * (1. - discount[groupSize])
			// recursively add costs of nested groups
			groupCost += groupsCosts(basketCopy)
			if groupCost < cost {
				cost = groupCost
			}
		}

	}
	return cost
}
