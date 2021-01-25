package minesweeper

import "bytes"

// Board to play minesweeper
type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}
