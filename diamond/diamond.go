package diamond

import (
	"errors"
	"strings"
)

// set to e.g. "." for debugging
const sep = " "

// Gen generates the diamond
func Gen(b byte) (string, error) {

	if b < 'A' || b > 'Z' {
		return "", errors.New("invalid input character")
	}

	var (
		rows  []string
		nRows = int(b-'A')*2 + 1
	)

	// generate row for each character:
	for char := byte('A'); char <= b; char++ {
		nFrameChars := int(b - char)
		row := strings.Repeat(sep, nFrameChars) + string(char)
		if char > byte('A') {
			row += strings.Repeat(sep, nRows-nFrameChars*2-2) + string(char)
		}
		row += strings.Repeat(sep, nFrameChars)
		rows = append(rows, row)
	}

	// mirror top rows to bottom
	if len(rows) > 1 {
		for i := len(rows) - 2; i >= 0; i-- {
			rows = append(rows, rows[i])
		}
	}

	return strings.Join(rows, "\n") + "\n", nil
}
