package brackets

var (
	pairs = map[rune]rune{'(': ')', '[': ']', '{': '}'}
	open  = map[rune]bool{'(': true, '[': true, '{': true}
	close = map[rune]bool{')': true, ']': true, '}': true}
)

// Bracket checks if brackets in a string are opend and closed in the right order
func Bracket(s string) bool {
	var stack []rune
	for _, c := range s {
		if _, ok := open[c]; ok { // if is opening, append closing to stack
			stack = append(stack, pairs[c])
		}
		// check order:
		if _, ok := close[c]; ok { // if is closing, check if last stack element is same closing bracket
			if len(stack) > 0 && stack[len(stack)-1] == c {
				stack = stack[:len(stack)-1] // if true, remove last element from stack
			} else {
				return false
			}
		}
	}
	// if every opened bracket was closed in right order, stack must now be empty
	return len(stack) == 0
}
