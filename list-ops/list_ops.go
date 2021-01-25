// Package listops implements list functionality without using Go's built-ins.
package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// IntList - a type to hold a list of integers.
type IntList []int

// Length returns the number of elements in a slice without using the built-in len
func (l IntList) Length() int {
	n := 0
	for range l {
		n++
	}
	return n
}

// Foldl - reduce each item into the accumulator from the left
func (l IntList) Foldl(f binFunc, item int) int {
	for _, i := range l {
		item = f(item, i)
	}
	return item
}

// Foldr - reduce each item into the accumulator from the right
func (l IntList) Foldr(f binFunc, item int) int {
	for i := l.Length() - 1; i >= 0; i-- {
		item = f(l[i], item)
	}
	return item
}

// Filter filters an IntList for elements the return true when passed to predFunc
func (l IntList) Filter(f predFunc) IntList {
	result := IntList{}
	for _, i := range l {
		if f(i) {
			result = result.Append(IntList{i})
		}
	}
	return result
}

// Map maps an IntList to another IntList by passing the elements through unaryFunc
func (l IntList) Map(f unaryFunc) IntList {
	result := IntList{}
	for _, i := range l {
		result = result.Append(IntList{f(i)})
	}
	return result
}

// Append appends elements to IntList without using the built-ins append or len
func (l IntList) Append(m IntList) IntList {
	n := l.Length()
	result := make(IntList, n+m.Length())
	for i, v := range l {
		result[i] = v
	}
	for i, v := range m {
		result[i+n] = v
	}
	return result
}

// Reverse inverts the elements in an IntList without using built-in functions
func (l IntList) Reverse() IntList {
	for i, j := 0, l.Length()-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
	return l
}

// Concat concatenates list elements, returning a flattend list
func (l IntList) Concat(m []IntList) IntList {
	for _, v := range m {
		l = l.Append(v)
	}
	return l
}
