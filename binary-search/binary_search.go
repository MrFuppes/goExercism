package binarysearch

// SearchInts performs a binary search on slice for key.
// returns the index of key.
// slice must be sorted, ascending.
func SearchInts(slice []int, key int) int {
	var (
		left, middle int
		right        = len(slice) - 1
	)

	for left <= right {
		middle = (left + right) / 2
		if slice[middle] < key { // value in the middle is less than key;
			left = middle + 1 // key must be to the right
		} else if slice[middle] > key { // value in the middle is greater than key;
			right = middle - 1 // value must be to the left
		} else { // value in the middle is the key!
			return middle
		}
	}
	return -1 // search done, nothing found
}
