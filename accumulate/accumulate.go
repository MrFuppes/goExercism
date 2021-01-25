// Package accumulate implements accumulator functionality
package accumulate

// accFunc - a typ to use as mapping function, mapping a string to a (modified) string.
type accFunc func(string) string

// Accumulate applies a mapping function to each element in the given slice of strings.
func Accumulate(given []string, converter accFunc) []string {
	result := []string{}
	for _, s := range given {
		result = append(result, converter(s))
	}
	return result
}
