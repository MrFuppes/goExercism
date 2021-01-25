// Package diffsquares implements functionality to calculate sums and sums of squares.
package diffsquares

// SquareOfSum - a function that returns the square of the sum of n natural numbers.
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares - a function that returns the sum of n squared natural numbers.
func SumOfSquares(n int) int {
	var sumOfSquares int
	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
	}
	return sumOfSquares
}

// Difference between square of sum and s um of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
