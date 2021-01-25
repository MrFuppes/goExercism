package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix - a type to hold data of a matrix
type Matrix struct {
	data   []int
	nLines int // number of rows in the input; defines the shape
}

const (
	newline = "\n"
	del     = " "
)

// New - matrix constructor from string
func New(in string) (*Matrix, error) {
	m := new(Matrix)
	lines := strings.Split(in, newline)
	m.data = []int{}
	m.nLines = len(lines) // need to keep track of how input looked...
	for _, l := range lines {
		parts := strings.Split(strings.Trim(l, " \t\n"), del)
		for _, s := range parts {
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			m.data = append(m.data, i)
		}
	}
	if len(m.data)%m.nLines != 0 {
		return nil, errors.New("number of elements per line must be equal")
	}
	return m, nil
}

// Rows - returns the values row-wise
func (m *Matrix) Rows() (rows [][]int) {
	nRows := m.nLines
	nCols := len(m.data) / nRows
	for i := 0; i < nRows; i++ {
		row := []int{} // make a new row, so we get a copy and not a view
		for j := 0; j < nCols; j++ {
			row = append(row, m.data[i*nCols+j])
		}
		rows = append(rows, row)
	}
	return rows
}

// Cols - returns the values column-wise
func (m *Matrix) Cols() (cols [][]int) {
	nRows := len(m.data) / m.nLines
	nCols := m.nLines
	for i := 0; i < nRows; i++ {
		row := []int{}
		for j := i; j <= i+nRows*(nCols-1); j = j + nRows {
			row = append(row, m.data[j])
		}
		cols = append(cols, row)
	}
	return cols
}

// Set - set a value of the matrix at [row, col] to val
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= m.nLines || col >= len(m.data)/m.nLines {
		return false
	}
	m.data[row*(len(m.data)/m.nLines)+col] = val
	return true
}
