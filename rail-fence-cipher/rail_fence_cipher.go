package railfence

// getRowIdx - a helper to generate a triangle wave
func getRowIdx(colIdx, n int) int {
	tmp := colIdx%(2*n) - n
	if tmp < 0 {
		return n - tmp*-1
	}
	return n - tmp
}

// Encode a string with rail fence cipher
func Encode(s string, nRails int) string {
	var (
		railfence = make([]string, nRails)
		result    string
	)
	for i, c := range s {
		railfence[getRowIdx(i, nRails-1)] += string(c)
	}
	for _, rail := range railfence {
		result += rail
	}
	return result
}

// Decode a string encoded with rail fence cipher
func Decode(s string, nRails int) string {
	var (
		l         = len(s)
		railLen   = make(map[int]int, nRails)
		railfence = make([]string, nRails)
		result    string
	)
	for i := range s {
		railLen[getRowIdx(i, nRails-1)]++
	}
	for i := 0; i < nRails; i++ {
		railfence[i] = s[0:railLen[i]]
		s = s[railLen[i]:len(s)]
	}
	for i := 0; i < l; i++ {
		idx := getRowIdx(i, nRails-1)
		result += string(railfence[idx][0])
		railfence[idx] = railfence[idx][1:len(railfence[idx])]
	}
	return result
}
