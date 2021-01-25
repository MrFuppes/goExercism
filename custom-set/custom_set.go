// Package stringset implements a set type for datatype string.
package stringset

var seen struct{}

// Set - a custom set type for strings
type Set map[string]struct{}

// New returns a pointer to a new string set
func New() Set {
	s := make(Set)
	return s
}

// NewFromSlice returns a set generated from a slice of strings
func NewFromSlice(input []string) Set {
	s := New()
	for _, v := range input {
		s[v] = seen
	}
	return s
}

// IsEmpty - check if a set is empty
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Add an element to a set
func (s Set) Add(str string) {
	s[str] = seen
}

// Has checks if a set has a certain element
func (s Set) Has(str string) bool {
	_, ok := s[str]
	return ok
}

// Equal checks if two sets are equal
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

// Subset checks if s2 is a subset of s1
func Subset(s1, s2 Set) bool {
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			return false
		}
	}
	return true
}

// Disjoint checks if two sets share no elements
func Disjoint(s1, s2 Set) bool {
	for k1 := range s1 {
		if _, ok := s2[k1]; ok {
			return false
		}
	}
	return true
}

// Intersection returns the set containing all elements of s1 that also belong to s2
func Intersection(s1, s2 Set) Set {
	result := New()
	for k1 := range s1 {
		if _, ok := s2[k1]; ok {
			result.Add(k1)
		}
	}
	return result
}

// Difference returns the set containing all elements of s1 that are not in s2
func Difference(s1, s2 Set) Set {
	result := New()
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			result.Add(k1)
		}
	}
	return result
}

// Union returns the set of all elements from s1 and s2
func Union(s1, s2 Set) Set {
	result := New()
	for k1 := range s1 {
		result.Add(k1)
	}
	for k2 := range s2 {
		result.Add(k2)
	}
	return result
}

// String - the stringer
func (s Set) String() string {
	str := "{"
	for k := range s {
		str += `"` + k + `", `
	}
	if len(str) > 1 {
		str = str[:len(str)-2]
	}
	return str + "}"
}
