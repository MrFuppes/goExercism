// Package proverb implements functionality to generate a famous proverb.
package proverb

import "fmt"

const (
	template = "For want of a %s the %s was lost."
	closure  = "And all for the want of a %s."
)

// Proverb takes a slice of words and converts them to a proverb.
func Proverb(rhyme []string) []string {
	var result []string
	if len(rhyme) == 0 {
		return result // somebody is joking...
	}

	// if rhyme has only one element, we only need the closure.
	if len(rhyme) == 1 {
		return append(result, fmt.Sprintf(closure, rhyme[0]))
	}

	// rhyme has multiple elements, we need the template.
	for i := 0; i < len(rhyme)-1; i++ {
		result = append(result, fmt.Sprintf(template, rhyme[i], rhyme[i+1]))
	}

	return append(result, fmt.Sprintf(closure, rhyme[0]))
}
