// Package scale implements functionality to generate the musical scales.
package scale

import "strings"

// idxIn - a helper function to get the index of a string in a slice of strings.
// returns -1 if the string is not found.
func idxIn(s string, list []string) int {
	for i, v := range list {
		if v == s {
			return i
		}
	}
	return -1
}

var (
	// sharp and flat sequences
	sharpSeq = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flatSeq  = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	// when to use the sharp sequence
	useSharp = []string{"A", "C", "G", "D", "A", "E", "B", "F#", "a", "c", "e", "b", "f#", "c#", "g#", "d#"}
)

const nNotes = 12 // 12 note system, so this is a constant

// Scale computes the musical scale for a given tonic and interva.
func Scale(tonic string, interval string) []string {
	// check from which sequence to grab notes
	src := sharpSeq
	if idxIn(tonic, useSharp) == -1 {
		src = flatSeq
	}

	// determine where to start
	idx := idxIn(strings.Title(tonic), src)

	// set default numbers of steps and notes. n notes is always 12 while n steps varies.
	var nSteps, steps = 12, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	if interval != "" {
		nSteps = len(interval)
		steps = make([]int, nSteps)
		for i, c := range interval {
			if c == 'm' {
				steps[i] = 1
			} else if c == 'M' {
				steps[i] = 2
			} else {
				steps[i] = 3
			}
		}
	}

	var result []string
	for i := 0; i <= nSteps-1; i++ {
		result = append(result, src[idx])
		idx += steps[i]
		if idx >= nNotes { // start from beginning if range is exceeded.
			idx -= nNotes
		}
	}

	return result
}
