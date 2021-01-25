package dominoes

// Domino is the building block of the domino chain
type Domino [2]int

// flip flips the orientation of the domino
func (d Domino) flip() Domino { return Domino{d[1], d[0]} }

// MakeChain attempts to create a chain of dominos
func MakeChain(input []Domino) (chain []Domino, ok bool) {
	if len(input) == 0 {
		return input, true
	}
	if len(input) == 1 {
		if input[0][0] == input[0][1] {
			return input, true
		}
		return input, false
	}

	// keep track of the stones we've used:
	used := make([]bool, len(input))

	// now call the chainer, trying each domino from the input as start, as second and so on
	for i, domino := range input {
		chain, used[i] = []Domino{domino}, true
		// if ok = chainer(1); ok {
		if ok = chainer(1, &input, &chain, &used); ok {
			break // found a chain!
		}
	}

	return chain, ok
}

// chainer appends a stone if posible, from the pool of not-used stones.
// need a function to build the chain so we can call it recursively
// n = index where a stone should be attached in the chain
func chainer(n int, input, chain *[]Domino, used *[]bool) bool {
	if n == len(*input) {
		return true // chain is complete
	}
	for i, domino := range *input {
		if !(*used)[i] {
			switch (*chain)[n-1][1] { // check previous stone...
			case domino[0]:
				*chain = append(*chain, domino)
			case domino[1]:
				*chain = append(*chain, domino.flip())
			default:
				continue // no match, try next
			}
			(*used)[i] = true                     // found match...
			if chainer(n+1, input, chain, used) { // recursion: find next domino
				return true // chain complete
			}
			(*used)[i] = false // reset for the next try
			// update chain, remove previously added domino since it did not result in complete chain:
			*chain = (*chain)[:n]
		}
	}

	return false // chain still incomplete...
}
