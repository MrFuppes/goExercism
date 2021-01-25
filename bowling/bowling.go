package bowling

import (
	"errors"
)

// Game holds results for all throws
type Game struct {
	throws []int
}

// NewGame returns a pointer to a new Game struct
func NewGame() *Game {
	return &Game{
		throws: []int{},
	}
}

// Roll adds pins to throws of the game
func (g *Game) Roll(pins int) error {
	if g.isDone() {
		return errors.New("Cannot roll after game is over")
	}
	if pins < 0 {
		return errors.New("Negative roll is invalid")
	}
	if pins > 10 {
		return errors.New("Pin count exceeds pins on the lane")
	}
	if len(g.throws)%2 == 1 && 10 < g.throws[len(g.throws)-1]+pins {
		return errors.New("Pin count exceeds pins on the lane")
	}
	g.throws = append(g.throws, pins)
	if pins == 10 {
		g.throws = append(g.throws, 0)
	}

	return nil
}

// Score evaluates the throws
func (g *Game) Score() (total int, err error) {
	if !g.isDone() {
		return total, errors.New("Score cannot be taken until the end of the game")
	}
	for t := 0; t < 20; t += 2 {
		first, second := g.throws[t], g.throws[t+1]
		total += first + second
		if first+second == 10 {
			total += g.throws[t+2]
		}
		if first == 10 {
			if g.throws[t+2] == 10 {
				total += g.throws[t+4]
			} else {
				total += g.throws[t+3]
			}
		}
	}

	return total, nil
}

func (g *Game) isDone() bool {
	n := len(g.throws)
	switch {
	case 0 <= n && n < 20:
		return false
	case n == 20:
		return g.throws[18]+g.throws[19] != 10
	case n == 21:
		return g.throws[18] != 10
	case n == 22 && g.throws[18] == 10 && g.throws[20] == 10:
		return false
	}

	return true
}
