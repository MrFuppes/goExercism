package series

// All the contiguous substrings of length n
func All(n int, s string) []string {
	r := []rune(s)
	var result []string
	for i := 0; i <= len(r)-n; i++ {
		result = append(result, string(r[i:i+n]))
	}
	return result
}

// First returns the first substring of s with length n,
// if s has less than n characters, false is returned as "ok" parameter.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return first, ok
	}
	return string([]rune(s)[0:n]), true
}

// UnsafeFirst returns the first substring of s with length n,
// if s has less than n characters, the function panics.
func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		panic("n >= len(s)")
	}
	return string([]rune(s)[0:n])
}
