package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Pair - a type to hold a pair of x-y-coordinates
type Pair [2]int

// Matrix - a type to hold an N x M matrix
type Matrix [][]int

// New makes a new matrix from a string
func New(s string) (*Matrix, error) {
	m := new(Matrix)
	for i, line := range strings.Split(s, "\n") {
		*m = append(*m, []int{})
		for _, part := range strings.Split(line, " ") {
			j, err := strconv.Atoi(part)
			if err != nil {
				return m, err
			}
			(*m)[i] = append((*m)[i], j)
		}
	}
	return m, nil
}

// maxInRow - a helper to get the maximum of a row
func maxInRow(r []int) int {
	result := r[0]
	for _, v := range r[1:len(r)] {
		if v > result {
			result = v
		}
	}
	return result
}

// minInCol - a helper to get the minimum of a column
func minInCol(c []int) int {
	result := c[0]
	for _, v := range c[1:len(c)] {
		if v < result {
			result = v
		}
	}
	return result
}

// getCol - a helper to extract a column
func (m *Matrix) getCol(i int) (col []int) {
	for _, r := range *m {
		col = append(col, r[i])
	}
	return col
}

// Saddle calculates saddle point coordinates of the given matrix
func (m *Matrix) Saddle() []Pair {

	var (
		nRows, nCols int
		maxima       [][]bool
		result       []Pair
	)

	nRows, nCols = len(*m), len((*m)[0])

	// greater than or equal to every element in its row
	for r, row := range *m {
		maxima = append(maxima, []bool{})
		max := maxInRow(row)
		for c := 0; c < nCols; c++ {
			maxima[r] = append(maxima[r], false)
			if (*m)[r][c] == max {
				maxima[r][c] = true
			}
		}
	}

	// ...and less than or equal to every element in its column.
	for c := 0; c < nCols; c++ {
		min := minInCol(m.getCol(c))
		for r := 0; r < nRows; r++ {
			if (*m)[r][c] == min && maxima[r][c] {
				fmt.Println(r, c, (*m)[r][c])
				result = append(result, Pair{r, c})
			}
		}
	}

	return result
}
