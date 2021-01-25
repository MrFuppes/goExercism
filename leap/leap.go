// Package leap implements leap year functionality.
package leap

// IsLeapYear checks if a year is a leap year. Returns true if the year is either
// divisible (w/o remainder) by 400 OR by 4 but not 100.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}
