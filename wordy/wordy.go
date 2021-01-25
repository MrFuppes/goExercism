package wordy

import (
	"strconv"
	"strings"
)

var operations = []string{"plus", "minus", "multiplied", "divided"}
var forbidden = "cubed"

// Answer tries to answer the question
func Answer(q string) (int, bool) {

	var (
		result, n, m int
		oprs         []string
		nbrs         []int
	)

	parts := strings.Fields(strings.Trim(q, "?"))

	for j, p := range parts {
		if p == forbidden {
			return 0, false
		}
		i, err := strconv.Atoi(p)
		if err == nil {
			if n == j-1 { // double number
				return 0, false
			}
			nbrs = append(nbrs, i)
			n = j
			continue
		}
		for _, o := range operations {
			if o == p {
				if m == j-1 { // double operation
					return 0, false
				}
				oprs = append(oprs, o)
				m = j
			}
		}
	}

	// must be one more number than there are operations and at least one number
	if len(nbrs) != len(oprs)+1 || len(nbrs) == 0 {
		return 0, false
	}

	result = nbrs[0]
	for i, o := range oprs {
		switch {
		case o == "plus":
			result += nbrs[i+1]
		case o == "minus":
			result -= nbrs[i+1]
		case o == "multiplied":
			result *= nbrs[i+1]
		case o == "divided":
			result /= nbrs[i+1]
		}
	}

	return result, true
}
