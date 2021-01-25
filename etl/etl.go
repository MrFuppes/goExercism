// Package etl implements fancy score table transformations
package etl

import "strings"

// Transform - a function to switch keys for values for a mapping of int to string slices
// all strings are put to lower-case.
func Transform(given map[int][]string) map[string]int {
	result := map[string]int{}
	for k, v := range given {
		for _, s := range v {
			result[strings.ToLower(string(s))] = k
		}
	}
	return result
}
