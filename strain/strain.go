// Package strain implements KEEP and DISCARD functionality for collections
package strain

// Ints - a collection of integers
type Ints []int

// Lists - a collection of collections of integers
type Lists [][]int

// Strings - a collection of strings
type Strings []string

// Keep - a method of the Ints collection that returns a collection for which predicate is true
func (i Ints) Keep(predicate func(int) bool) Ints {
	var result Ints
	for _, val := range i {
		if predicate(val) {
			result = append(result, val)
		}
	}
	return result
}

// Discard - a method of the Ints collection that returns a collection for which predicate is false
func (i Ints) Discard(predicate func(int) bool) Ints {
	var result Ints
	for _, val := range i {
		if !predicate(val) {
			result = append(result, val)
		}
	}
	return result
}

// Keep - a method of the Lists collection that returns a collection for which predicate is true
func (l Lists) Keep(predicate func([]int) bool) Lists {
	var result Lists
	for _, list := range l {
		if predicate(list) {
			result = append(result, list)
		}
	}
	return result
}

// Keep - a method of the Strings collection that returns a collection for which predicate is true
func (s Strings) Keep(predicate func(string) bool) Strings {
	var result Strings
	for _, str := range s {
		if predicate(str) {
			result = append(result, str)
		}
	}
	return result
}
