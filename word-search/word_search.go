package wordsearch

import (
	"errors"
	"strings"
)

var (
	errNoWords      = errors.New("no words or puzzle supplied")
	errPuzzleStrLen = errors.New("all strings in puzzle must be of equal length")
)

// Solve searches for words in a puzzle
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	if len(words) == 0 || len(puzzle) == 0 {
		return nil, errNoWords
	}
	l := len(puzzle[0])
	for _, s := range puzzle {
		if len(s) != l {
			return nil, errPuzzleStrLen
		}
		l = len(s)
	}

	found := make(map[string][2][2]int)

	// check horizontal, left to right and right to left
	for i, s := range puzzle {
		for _, w := range words {
			if j := strings.Index(s, w); j != -1 {
				found[w] = [2][2]int{{j, i}, {j + len(w) - 1, i}}
			}
			if j := strings.Index(s, reverse(w)); j != -1 {
				found[w] = [2][2]int{{j + len(w) - 1, i}, {j, i}}
			}
		}
	}

	// check vertical, top to bottom and bottom to top
	upDownPuzzle := make([]string, len(puzzle[0]))
	for i := range puzzle[0] {
		for j := range puzzle {
			upDownPuzzle[i] += string(puzzle[j][i])
		}
	}
	for i, s := range upDownPuzzle {
		for _, w := range words {
			if j := strings.Index(s, w); j != -1 {
				found[w] = [2][2]int{{i, j}, {i, j + len(w) - 1}}
			}
			if j := strings.Index(s, reverse(w)); j != -1 {
				found[w] = [2][2]int{{i, j + len(w) - 1}, {i, j}}
			}
		}
	}

	// check diagonals
	if len(puzzle) > 1 {
		diagPuzzle0, diagIdx0 := diagonals(puzzle)
		for i, s := range diagPuzzle0 {
			for _, w := range words {
				if len(s) >= len(w) {
					if j := strings.Index(s, w); j != -1 {
						found[w] = [2][2]int{diagIdx0[i][j], diagIdx0[i][j+len(w)-1]}
					}
					if j := strings.Index(s, reverse(w)); j != -1 {
						found[w] = [2][2]int{diagIdx0[i][j+len(w)-1], diagIdx0[i][j]}
					}
				}
			}
		}

		flippedPuzzle := make([]string, len(puzzle))
		for i, j := len(puzzle)-1, 0; i >= 0; i, j = i-1, j+1 {
			flippedPuzzle[j] = puzzle[i]
		}
		diagPuzzle1, diagIdx1 := diagonals(flippedPuzzle)
		for i, s := range diagPuzzle1 {
			for _, w := range words {
				if len(s) >= len(w) {
					if j := strings.Index(s, w); j != -1 {
						found[w] = [2][2]int{{diagIdx1[i][j][0], (len(puzzle) - 1) % diagIdx1[i][j][1]},
							{diagIdx1[i][j+len(w)-1][0], (len(puzzle) - 1) % diagIdx1[i][j+len(w)-1][1]}}
					}
					if j := strings.Index(s, reverse(w)); j != -1 {
						found[w] = [2][2]int{{diagIdx1[i][j+len(w)-1][0], (len(puzzle) - 1) - diagIdx1[i][j+len(w)-1][1]},
							{diagIdx1[i][j][0], (len(puzzle) - 1) - diagIdx1[i][j][1]}}
					}
				}
			}
		}
	}

	if len(found) != len(words) {
		errNotFound := "Not found:"
		for _, w := range words {
			if _, ok := found[w]; !ok {
				errNotFound += " " + w
			}
		}
		return found, errors.New(errNotFound)
	}

	return found, nil
}

// reverse - a helper to reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// diagonals - a helper to return the diagonals of the puzzle
func diagonals(m []string) (diagon []string, diagonIdx [][][2]int) {
	width, height := len(m[0]), len(m)
	for k := 0; k <= width+height-2; k++ {
		diagon = append(diagon, "")
		diagonIdx = append(diagonIdx, [][2]int{})
		for j := 0; j <= k; j++ {
			i := k - j
			if i < height && j < width {
				diagon[k] += string(m[i][j])
				diagonIdx[k] = append(diagonIdx[k], [2]int{j, i})
			}
		}
	}
	return diagon, diagonIdx
}
