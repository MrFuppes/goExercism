package yacht

// funcMap maps categories to functions which calculate results for a given category.
// interface type is used since the different functions have different signatures.
var funcMap = map[string]interface{}{
	"Ones":            nOfNumber,
	"Twos":            nOfNumber,
	"Threes":          nOfNumber,
	"Fours":           nOfNumber,
	"Fives":           nOfNumber,
	"Sixes":           nOfNumber,
	"Full House":      fullHouse,
	"Four of a Kind":  fourOfAKind,
	"Little Straight": straight,
	"Big Straight":    straight,
	"Choice":          choice,
	"Yacht":           yacht,
}

// get the counts for each of the thrown values
func getCounts(dice []int) map[int]int {
	counts := make(map[int]int)
	for _, v := range dice {
		counts[v]++
	}
	return counts
}

// calculate score for selected number n
func nOfNumber(n int, dice []int) (score int) {
	for _, v := range dice {
		if v == n {
			score += n
		}
	}
	return score
}

// calculate score for "full house"
func fullHouse(dice []int) (score int) {
	for v, n := range getCounts(dice) {
		if n < 2 || 3 < n {
			return 0
		}
		score += n * v
	}
	return score
}

// calculate score for "4 of a kind"
func fourOfAKind(dice []int) (score int) {
	for v, n := range getCounts(dice) {
		if n >= 4 {
			score = 4 * v
		}
	}
	return score
}

// calculate score for straigt, set "start" to 1 for a little straigt
func straight(dice []int, start int) (score int) {
	if len(getCounts(dice)) != 5 {
		return 0
	}
	min, max := dice[0], dice[0]
	for i := 0; i < len(dice); i++ {
		if dice[i] < min {
			min = dice[i]
		}
		if dice[i] > max {
			max = dice[i]
		}
	}
	if min == start && max-min == 4 {
		return 30
	}
	return 0
}

// just count all numbers
func choice(dice []int) (score int) {
	for _, v := range dice {
		score += v
	}
	return score
}

// yacht: all numbers equal give score 50
func yacht(dice []int) int {
	v := dice[0]
	for i := 0; i < len(dice); i++ {
		if dice[i] != v {
			return 0
		}
	}
	return 50
}

// Score returns results for a throw of dice
func Score(dice []int, category string) (score int) {
	switch category {
	case "ones":
		score = funcMap["Ones"].(func(int, []int) int)(1, dice)
	case "twos":
		score = funcMap["Twos"].(func(int, []int) int)(2, dice)
	case "threes":
		score = funcMap["Threes"].(func(int, []int) int)(3, dice)
	case "fours":
		score = funcMap["Fours"].(func(int, []int) int)(4, dice)
	case "fives":
		score = funcMap["Fives"].(func(int, []int) int)(5, dice)
	case "sixes":
		score = funcMap["Sixes"].(func(int, []int) int)(6, dice)
	case "full house":
		score = funcMap["Full House"].(func([]int) int)(dice)
	case "four of a kind":
		score = funcMap["Four of a Kind"].(func([]int) int)(dice)
	case "little straight":
		score = funcMap["Little Straight"].(func([]int, int) int)(dice, 1)
	case "big straight":
		score = funcMap["Big Straight"].(func([]int, int) int)(dice, 2)
	case "choice":
		score = funcMap["Choice"].(func([]int) int)(dice)
	case "yacht":
		score = funcMap["Yacht"].(func([]int) int)(dice)
	default:
		score = 0
	}
	return score
}
