package ocr

import "strings"

var digits = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

// Recognize converts OCR digits to string
func Recognize(input string) []string {
	var result []string
	// split input string to individual lines
	lines := strings.Split(input, "\n")

	// step through lines in blocks of 4
	for i := 1; i < len(lines); i += 4 {
		result = append(result, parseLines(lines[i:i+3]))
	}

	return result

}

// parseLines - a helper to parse lines with digits on the OCR
func parseLines(input []string) string {
	var result string
	// step trough lines in blocks of 3 columns
	for i := 0; i < len(input[0]); i += 3 {
		// make a string row-wise and map to number
		result += recognizeDigit(input[0][i:i+3] + input[1][i:i+3] + input[2][i:i+3])
	}

	return result
}

// recognizeDigit - a helper to parse a single digit
func recognizeDigit(s string) string {
	if d, ok := digits[s]; ok {
		return d
	}
	return "?"
}
