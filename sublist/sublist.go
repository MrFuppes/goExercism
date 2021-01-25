package sublist

// a helper function to check the occurances (and indices of those occurances) of an int in an int slice.
func valueCounts(s []int, i int) (int, []int) {
	var (
		count int
		idx   []int
	)
	for j, v := range s {
		if v == i {
			count++
			idx = append(idx, j)
		}
	}
	return count, idx
}

// a helper function to check if an int is in an int slice. returns index if the value is found, -1 otherwise.
func containsInt(s []int, i int) int {
	for j, v := range s {
		if v == i {
			return j
		}
	}
	return -1
}

// Relation is returned by Sublist
type Relation string

// Sublist checks if a list l1 is a sublist or superlist of a list l2.
// If both lists are equal or unequal, return value is as such.
func Sublist(l1 []int, l2 []int) Relation {
	// trivial case #1: both slices are empty:
	if len(l1) == 0 && len(l2) == 0 {
		return "equal"
	}
	// trivial case #2: l1 is empty, l2 not
	if len(l1) == 0 && len(l2) > 0 {
		return "sublist"
	}
	// trivial case #3: l2 is empty, l1 not
	if len(l1) > 0 && len(l2) == 0 {
		return "superlist"
	}
	// both slices of equal length, compare element-wise:
	if len(l1) == len(l2) {
		for i := range l1 {
			if l1[i] != l2[i] {
				return "unequal"
			}
		}
		return "equal"
	}
	// l1 is shorter than l2 -> check if it is a sublist.
	if len(l1) < len(l2) {
		n, idx := valueCounts(l2, l1[0])
		if n > 0 {
			for _, i := range idx {
				for k, v := range l1[1:] {
					j := containsInt(l2[i+1:], v)
					if j != k {
						break
					}
					return "sublist"
				}

			}

		}
	}
	// l1 is longer than l2 -> check if it is a superlist.
	if len(l1) > len(l2) {
		n, idx := valueCounts(l1, l2[0])
		if n > 0 {
			for _, i := range idx {
				for k, v := range l2[1:] {
					j := containsInt(l1[i+1:], v)
					if j != k {
						break
					}
					return "superlist"
				}

			}

		}
	}

	return "unequal"
}
