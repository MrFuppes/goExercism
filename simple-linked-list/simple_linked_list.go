package linkedlist

import "errors"

// Element of the linked list. Has a value and a pointer to the next value.
type Element struct {
	data int
	next *Element
}

// List of linked elements.
type List struct {
	head *Element
	size int
}

// New returns a linked list containing the supplied values.
func New(values []int) *List {
	l := new(List)
	for _, v := range values {
		l.Push(v)
	}
	return l
}

// Size returns the number of elements in the linked list
func (l *List) Size() int { return (*l).size }

// Push adds a value to the list
func (l *List) Push(i int) {
	(*l).head = &Element{data: i, next: l.head}
	(*l).size++
}

// Pop removes an element from the list
func (l *List) Pop() (int, error) {
	if l.Size() == 0 {
		return 0, errors.New("cannot pop from empty list")
	}
	p, data := (*l).head, (*l).head.data
	(*l).head = p.next
	(*l).size--
	return data, nil
}

// Array converts the linked list to an int slice
func (l *List) Array() []int {
	arr := make([]int, (*l).size)
	p := (*l).head
	for i := (*l).size - 1; i >= 0; i-- {
		arr[i], p = p.data, p.next
	}
	return arr
}

// Reverse returns a reversly linked list
func (l *List) Reverse() *List {
	rev := new(List)
	for p := (*l).head; p != nil; p = p.next {
		rev.Push(p.data)
	}
	return rev
}
