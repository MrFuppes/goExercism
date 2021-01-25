package tree

import (
	"errors"
	"sort"
)

// root ID is always 0:
const rootID = 0

// Record - the basic type to hold ID and Parent
type Record struct {
	ID, Parent int
}

// isInvalid - check if a record is invalid. pos specifies the position in the sorted Records slice.
func (r Record) isInvalid(pos int) bool {
	return (r.ID < r.Parent) || (r.ID != pos) || (pos != rootID && r.ID == r.Parent)
}

// Records - a slice of Record; required to define Less method...
type Records []Record

// Len - a method so we can sort a slice of records, see https://golang.org/pkg/sort/#Sort
func (recs Records) Len() int {
	return len(recs)
}

// Less - a method so we can sort a slice of records
func (recs Records) Less(i, j int) bool {
	return recs[i].ID < recs[j].ID
}

// Swap - a method so we can sort a slice of records
func (recs Records) Swap(i, j int) {
	recs[i], recs[j] = recs[j], recs[i] // since recs is a slice, we can swap in-place
}

// Node of the tree. Each node has an id and children, which are kept in slices of nodes
type Node struct {
	ID       int
	Children []*Node
}

// Build builds the tree from a slice of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Sort(Records(records))

	if records[0].ID != rootID {
		return nil, errors.New("root ID must be zero")
	}
	if records[0].Parent != rootID {
		return nil, errors.New("root record must not have a parent")
	}

	tree := make(map[int]*Node)

	for i, r := range records {
		if r.isInvalid(i) {
			return nil, errors.New("encountered invalid record")
		}

		tree[i] = &Node{ID: r.ID}
		if i == 0 {
			continue // basis of the tree = root node, no child (branch)
		}

		// otherwise, append the current record as a child
		tree[r.Parent].Children = append(tree[r.Parent].Children, tree[i])
	}

	return tree[0], nil // the root points to all the rest...

}
