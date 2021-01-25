package secret

// 1 = wink
// 10 = double blink
// 100 = close your eyes
// 1000 = jump
// 10000 = Reverse the order of the operations in the secret handshake.

// revStrSlice - a helper to reverse a slice
func revStrSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// operations - an array of defined handshake operations
var operations = [...]string{"wink", "double blink", "close your eyes", "jump"}

// Handshake converts a number to a series of operations
func Handshake(code uint) (h []string) {
	for i, op := range operations {
		// right-shift the bits of the input variable code by the index and check if the bit is high (& 1 retunrs 1)
		if code>>i&1 == 1 { // if so, append the corresponding operation
			h = append(h, op)
		}
	}
	if code > 16 {
		revStrSlice(h)
	}
	return h
}
