// Package connect - a modified version of sbfaulkner's nice solution
package connect

// Board to play the game
type Board []string

const (
	playerX byte = 'X' // plays left to right
	playerO byte = 'O' // plays top to bottom
)

// vectors to move on the board
var dxdy = []struct {
	dx, dy int
}{
	{-1, 1},
	{-1, 0},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
}

// ResultOf calculates the winner of a given board
func ResultOf(board []string) (string, error) {
	b := Board(board)
	switch {
	case b.leftRightPath(playerX):
		return string(playerX), nil
	case b.topBottomPath(playerO):
		return string(playerO), nil
	default:
		return "", nil
	}
}

// leftRightPath - check if there is a complete path from left to right for a given player
func (b *Board) leftRightPath(player byte) bool {
	for y := range *b {
		if b.followTrace(0, y, player, false) { // b.followAcross(0, r, player) {
			return true
		}
	}

	return false
}

// topDownValid - check if there is a complete path from top to bottom for a given player
func (b *Board) topBottomPath(player byte) bool {
	for x := range (*b)[0] {
		if b.followTrace(x, 0, player, true) { //b.followDown(c, 0, player) {
			return true
		}
	}

	return false
}

// followTrace - recursively follow a player's stones
func (b *Board) followTrace(x, y int, p byte, topDown bool) bool {
	if !b.validateField(x, y, p) {
		return false
	}

	if (topDown && y == len(*b)-1) || (!topDown && x == len((*b)[y])-1) {
		return true
	}

	for _, d := range dxdy {
		if b.followTrace(x+d.dx, y+d.dy, p, topDown) {
			return true
		}
	}

	return false
}

// validateField - returns true if coordinate is valid, color matches player p, has not been visited yet
func (b *Board) validateField(x, y int, p byte) bool {
	if y < 0 || y >= len(*b) {
		return false
	}

	if x < 0 || x >= len((*b)[y]) {
		return false
	}

	row := []byte((*b)[y])
	if row[x] != p {
		return false
	}

	row[x] = '*' // visited; field is now != p
	(*b)[y] = string(row)

	return true
}
