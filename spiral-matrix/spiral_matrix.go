package spiralmatrix

// SpiralMatrix returns a matrix m*m
func SpiralMatrix(size int) [][]int {
	var (
		fill               int = 1
		rowStart, colStart int
		rowEnd, colEnd     = size, size
	)

	arr := make([][]int, size)
	for i := 0; i < size; i++ {
		arr[i] = make([]int, size)
	}

	for rowStart < rowEnd && colStart < colEnd {
		// first row, fill columns
		for i := colStart; i < colEnd; i++ {
			arr[rowStart][i] = fill
			fill++
		}
		rowStart++

		// last column, fill rows
		for i := rowStart; i < rowEnd; i++ {
			arr[i][colEnd-1] = fill
			fill++
		}
		colEnd--

		// last row, fill columns in reverse order
		if rowStart < rowEnd {
			for i := colEnd - 1; i > colStart-1; i-- {
				arr[rowEnd-1][i] = fill
				fill++
			}
			rowEnd--
		}

		// first column, fill rows in reverse order
		if colStart < colEnd {
			for i := rowEnd - 1; i > rowStart-1; i-- {
				arr[i][colStart] = fill
				fill++
			}
			colStart++
		}
	}

	return arr
}
