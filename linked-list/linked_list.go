package linkedlist

import "errors"

// Node of the linked list. Has data and a pointer to the next and previous value.
// the prev pointer of the first node points to the last.
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// List of linked elements.
type List struct {
	head *Node
	tail *Node
	size int
}

// ErrEmptyList for not allowed operations on empty lists.
var ErrEmptyList = errors.New("list is empty")

// NewList - create a new linked list
func NewList(args ...interface{}) *List {
	l := new(List)
	for _, v := range args {
		l.PushBack(v)
	}
	return l
}

// First returns the pointer to the first node
func (l *List) First() *Node { return l.head }

// Last returns the pointer to the last node
func (l *List) Last() *Node { return l.tail }

// Next returns the pointer to the next node
func (e *Node) Next() *Node { return e.next }

// Prev returns the pointer to the previous node
func (e *Node) Prev() *Node { return e.prev }

// PushFront - add a node at the head of the list
func (l *List) PushFront(v interface{}) {
	p := &Node{Val: v, next: l.head, prev: nil}
	if l.head == nil {
		l.tail = p
	} else {
		l.head.prev = p
	}
	l.head = p
	l.size++
}

// PushBack - add a node at the tail of the list
func (l *List) PushBack(v interface{}) {
	p := &Node{Val: v, next: nil, prev: l.tail}
	if l.tail == nil {
		l.head = p
	} else {
		l.tail.next = p
	}
	l.tail = p
	l.size++
}

// PopFront removes the node at the head of the list
func (l *List) PopFront() (interface{}, error) {
	pFirst := l.First()
	if pFirst == nil {
		return 0, ErrEmptyList
	}
	l.head = pFirst.Next()
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return pFirst.Val, nil
}

// PopBack removes the node at the tail of the list
func (l *List) PopBack() (interface{}, error) {
	pLast := l.Last()
	if pLast == nil {
		return 0, ErrEmptyList
	}
	l.tail = pLast.Prev()
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return pLast.Val, nil
}

// Reverse the order of the list
func (l *List) Reverse() {
	for p := l.First(); p != nil; p = p.Prev() {
		p.next, p.prev = p.prev, p.next
	}
	l.head, l.tail = l.tail, l.head
}
