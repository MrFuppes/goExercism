package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var (
	plantMap = map[byte]string{
		'G': "grass",
		'C': "clover",
		'R': "radishes",
		'V': "violets",
	}

	errDiagFmt         = errors.New("wrong diagram format")
	errDupeChild       = errors.New("child names must be unique")
	errInvalidCupCode  = errors.New("invalid cup code")
	errMismatchedRows  = errors.New("rows must match")
	errOddNumberOfCups = errors.New("number of cups must be even")
)

// Garden type, mapping child name to plants
type Garden map[string][]string

// NewGarden creates a new garde from diagram and children
func NewGarden(diagram string, children []string) (*Garden, error) {
	// make sure input look nice:
	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 {
		return nil, errDiagFmt
	}
	rows = rows[1:]
	if len(rows[0]) != 2*len(children) || len(rows[1]) != 2*len(children) {
		return nil, errMismatchedRows
	}

	// ensure names appear sorted without modifying input:
	decendants := make([]string, len(children))
	copy(decendants, children)
	sort.Strings(decendants)

	// populate the garden with children and give each child its plants
	garden := Garden(map[string][]string{})

	for i, child := range decendants {
		plant0, ok0 := plantMap[rows[0][2*i]]
		plant1, ok1 := plantMap[rows[0][2*i+1]]
		plant2, ok2 := plantMap[rows[1][2*i]]
		plant3, ok3 := plantMap[rows[1][2*i+1]]
		if !(ok0 && ok1 && ok2 && ok3) {
			return nil, errInvalidCupCode
		}
		garden[child] = []string{plant0, plant1, plant2, plant3}
	}

	if len(garden) != len(children) {
		return nil, errDupeChild
	}

	return &garden, nil
}

// Plants returns the plants of a certain child
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	return plants, ok
}
