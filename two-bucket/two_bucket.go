package twobucket

import (
	"errors"
)

type bucket struct {
	cap, level int
	name       string
}

// buckets - a container to holder the buckets
type buckets []*bucket

func (b *bucket) fill() {
	b.level = b.cap
}

func (b *bucket) empty() {
	b.level = 0
}

// transfer from one to two
func (b *bucket) transferTo(to *bucket) (ok bool) {
	room := to.cap - to.level
	if room == 0 || b.level == 0 {
		return false
	}
	if b.level > room {
		to.level = to.cap
		b.level -= room
	} else {
		to.level += b.level
		b.level = 0
	}
	return true
}

// helper to calculate greatest common divisor
func gcd(a, b int) int {
	var tmp int
	for b != 0 {
		tmp = a
		a = b
		b = tmp % b
	}
	return a
}

// Solve calculates a solution for the jug-pouring problem with two jugs
func Solve(bucketOne, bucketTwo, goal int, startBucket string) (gBucket string, moves, other int, err error) {

	if bucketOne < 1 || bucketTwo < 1 || goal < 1 {
		return "", 0, 0, errors.New("invalid input")
	}

	if gcd(bucketOne, bucketTwo) != 1 {
		return "", 0, 0, errors.New("bucket sizes must be relatively prime")
	}

	var container buckets

	if startBucket == "one" {
		container = buckets{&bucket{cap: bucketOne, name: "one"}, &bucket{cap: bucketTwo, name: "two"}}
	} else if startBucket == "two" {
		container = buckets{&bucket{cap: bucketTwo, name: "two"}, &bucket{cap: bucketOne, name: "one"}}
	} else {
		return "", 0, 0, errors.New("invalid start bucket name")
	}

	for container[0].level != goal {
		switch {
		case container[0].level == 0:
			container[0].fill()
		case container[1].level == container[1].cap:
			container[1].empty()
		case container[1].cap == goal:
			container[1].fill()
			container[0], container[1] = container[1], container[0]
		default:
			container[0].transferTo(container[1])
		}
		moves++
	}

	return container[0].name, moves, container[1].level, nil
}
