package transpose

// Transpose returns a transposed version of a slice of strings.
func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	var maxLen int
	for _, str := range input {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	out := make([]string, maxLen)

	for r := range input {
		for c := range input[r] {
			for len(out[c]) < r {
				out[c] += " "
			}
			out[c] += string(input[r][c])
		}
	}

	return out
}
