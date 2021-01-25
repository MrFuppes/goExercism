package binarysearchtree

// SearchTreeData - node of the binary tree
type SearchTreeData struct {
	data        int
	left, right *SearchTreeData
}

// Bst makes a new binary search tree with one node
func Bst(i int) *SearchTreeData { return &SearchTreeData{data: i} }

// Insert a value into the BST
func (bst *SearchTreeData) Insert(i int) {
	switch {
	case i <= bst.data && bst.left == nil:
		bst.left = &SearchTreeData{data: i}
	case i > bst.data && bst.right == nil:
		bst.right = &SearchTreeData{data: i}
	case i <= bst.data:
		bst.left.Insert(i)
	case i > bst.data:
		bst.right.Insert(i)
	}
}

// MapString maps the BST to a slice of string
func (bst *SearchTreeData) MapString(f func(int) string) []string {
	data := bst.Walk()
	result := make([]string, len(data))

	for i, value := range data {
		v, ok := value.(int)
		if ok {
			result[i] = f(v)
		} // could panic / log error if not ok
	}

	return result
}

// MapInt maps the BST to a slice of int
func (bst *SearchTreeData) MapInt(f func(int) int) []int {
	data := bst.Walk()
	result := make([]int, len(data))

	for i, value := range data {
		v, ok := value.(int)
		if ok {
			result[i] = f(v)
		} // could panic / log error if not ok
	}

	return result
}

// Walk the tree to return a slice of values
func (bst *SearchTreeData) Walk() []interface{} {
	data := []interface{}{}

	if bst.left != nil {
		data = append(data, bst.left.Walk()...)
	}
	data = append(data, bst.data)
	if bst.right != nil {
		data = append(data, bst.right.Walk()...)
	}

	return data
}
