package queenattack

import (
	"errors"
	"strconv"
)

// pos - a struct to represent the position of a piece
type pos struct {
	x int
	y int
}

// posMap - mapping lower case characters to integers
var posMap = map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8}

// stringToPos - a helper to convert string to position
func stringToPos(s string) (p pos, err error) {
	r, c := rune(s[0]), rune(s[1])
	x, ok := posMap[r]
	if !ok {
		return p, errors.New("can't convert letter to position - must be [a-h]")
	}
	y, err := strconv.Atoi(string(c))
	if err != nil {
		return p, err
	}
	if y < 1 || y > 8 {
		return p, errors.New("column out of bounds")
	}

	return pos{x, y}, nil
}

// abs - a helper to get absolute value of an int
func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

// CanQueenAttack - given two positions an a chess board, check if one queen can attack the other
func CanQueenAttack(w, b string) (attack bool, err error) {
	if w == b {
		return false, errors.New("same square")
	}

	wPos, err := stringToPos(w)
	if err != nil {
		return false, err
	}

	bPos, err := stringToPos(b)
	if err != nil {
		return false, err
	}

	if wPos.x == bPos.x || wPos.y == bPos.y {
		return true, err
	}

	if abs(wPos.x-bPos.x) == abs(wPos.y-bPos.y) {
		return true, err
	}

	return false, err
}
