package isbn

import (
	"strings"
)

// IsValidISBN determins if an ISBN given as a string is valid
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false // must have 10 characters
	}
	if isbn[9] < '0' || ('9' < isbn[9] && isbn[9] != 'X') {
		return false // must end with digit or X
	}
	var checksum int
	for i, j := 0, 10; i < 9; i, j = i+1, j-1 {
		if isbn[i] < '0' || '9' < isbn[i] {
			return false
		}
		checksum += int(isbn[i]-48) * j // 48 == '0'
	}
	if isbn[len(isbn)-1] == 'X' {
		checksum += 10
	} else {
		checksum += int(isbn[9] - 48)
	}

	return checksum%11 == 0
}
