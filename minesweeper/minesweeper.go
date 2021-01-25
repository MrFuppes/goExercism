package minesweeper

import (
	"errors"
)

const (
	edge  = '+'
	wall  = '-'
	side  = '|'
	mine  = '*'
	empty = ' '
)

// Count calculates the adjacent mine count matrix
func (b Board) Count() error {
	nRows, nCols := len(b), len((b)[0])
	if !b.valid(nRows, nCols) {
		return errors.New("invalid board")
	}

	counts := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		counts[i] = make([]int, nCols)
	}
	for i := 1; i < nRows-1; i++ {
		for j := 1; j < nCols-1; j++ {
			if b[i][j] == mine {
				counts[i-1][j-1]++
				counts[i][j-1]++
				counts[i+1][j-1]++
				counts[i-1][j]++
				counts[i+1][j]++
				counts[i-1][j+1]++
				counts[i][j+1]++
				counts[i+1][j+1]++
			}
		}
	}

	for i := 1; i < nRows-1; i++ {
		for j := 1; j < nCols-1; j++ {
			if b[i][j] != mine && counts[i][j] != 0 {
				b[i][j] = '0' + byte(counts[i][j])
			}
		}
	}

	return nil
}

// valid checks if the input is a valid minesweeper board drawing
func (b *Board) valid(nRows, nCols int) bool {
	if len((*b)[nRows-1]) != nCols {
		return false
	}
	// all edges must be +
	if (*b)[0][0] != edge || (*b)[0][nCols-1] != edge || (*b)[nRows-1][0] != edge || (*b)[nRows-1][nCols-1] != edge {
		return false
	}
	// upper and lower walls must be -
	for i := 1; i < nCols-1; i++ {
		if (*b)[0][i] != wall || (*b)[nRows-1][i] != wall {
			return false
		}
	}
	// side walls must be |, everything else either empty (space) or mine
	if nRows > 2 {
		for _, row := range (*b)[1 : nRows-1] {
			if len(row) != nCols || row[0] != side || row[nCols-1] != side {
				return false
			}
			for _, c := range row[1 : nCols-1] {
				if c != empty && c != mine {
					return false
				}
			}
		}
	}

	return true
}
