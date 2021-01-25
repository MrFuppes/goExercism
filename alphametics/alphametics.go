package alphametics

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// PermutationNext generates the next permutation of a slice of int; modifies data in-place.
// Data must be sorted initially.
// Returns true if successfull, false if exhausted.
// http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func PermutationNext(data []int) bool {
	var k, l int
	// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	for k = len(data) - 2; ; k-- {
		if k < 0 {
			return false
		}
		if data[k] < data[k+1] {
			break
		}
	}
	// Find the largest index l greater than k such that a[k] < a[l].
	for l = len(data) - 1; !(data[k] < data[l]); l-- {
	}
	// Swap the value of a[k] with that of a[l].
	data[k], data[l] = data[l], data[k]
	// Reverse the sequence from a[k + 1] up to and including the final element a[n].
	for i, j := k+1, len(data)-1; i < j; i++ {
		data[i], data[j] = data[j], data[i]
		j--
	}
	return true
}

// getMapping - a helper to generate the string-to-number mapping
func getMapping(s []string, nbrs []int) map[string]int {
	mapping := make(map[string]int, len(s))
	for i, c := range s {
		mapping[c] = nbrs[i]
	}
	return mapping
}

// check Map - a helper to see if the guessed mapping is correct
func checkMapping(have []string, want string, mapping map[string]int) (ok bool) {
	var (
		haveSum int
		desired string
	)
	for _, s := range have {
		tmpS := ""
		for i, c := range s {
			// leading zero not allowed
			if i == 0 && mapping[string(c)] == 0 {
				return false
			}
			tmpS += fmt.Sprint(mapping[string(c)])
		}
		tmpI, _ := strconv.Atoi(tmpS)
		haveSum += tmpI
	}
	for i, c := range want {
		// leading zero not allowed
		if i == 0 && mapping[string(c)] == 0 {
			return false
		}
		desired += fmt.Sprint(mapping[string(c)])
	}
	desiredSum, _ := strconv.Atoi(desired)
	return desiredSum == haveSum
}

// Solve by brute force!
func Solve(s string) (result map[string]int, err error) {
	parts := strings.Split(strings.ReplaceAll(s, " ", ""), "==")
	have, want := strings.Split(parts[0], "+"), parts[1]
	// check if have and want are strings with one character each
	if len(have) == 1 {
		if len(have[0]) == 1 && len(want) == 1 && have[0] != want {
			return result, errors.New("not solvable")
		}
	}
	// check if solution would require leading zero
	for _, h := range have {
		if len(h) > len(want) {
			return result, errors.New("not solvable")
		}
	}

	var (
		unqCharMap = make(map[rune]int)
		unqChars   = []string{}
		nbrs       = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	)

	for _, r := range s {
		if r != ' ' && r != '=' && r != '+' {
			unqCharMap[r]++
		}
	}

	if len(unqCharMap) > 10 {
		return result, errors.New("number of unique letters must be <= 10")
	}

	for k := range unqCharMap {
		unqChars = append(unqChars, string(k))
	}

	result = getMapping(unqChars, nbrs)

	for i := 0; i <= 10e10; i++ {
		ok := checkMapping(have, want, result)
		if ok {
			return result, nil
		}
		PermutationNext(nbrs)
		result = getMapping(unqChars, nbrs)
	}
	return result, errors.New("not solvable")
}
