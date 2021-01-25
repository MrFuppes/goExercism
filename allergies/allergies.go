package allergies

import "sort"

const ( // give a number to each allergen
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var scoreMap = map[string]uint{ // map allergen names to their numbers
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}

// Allergies returns matching allergies given a certain score
func Allergies(score uint) (allergicTo []string) {
	for allergen := range scoreMap {
		if AllergicTo(score, allergen) {
			allergicTo = append(allergicTo, allergen)
		}
	}
	// make sure the output is sorted based on the score of the individual allergenes
	sort.Slice(allergicTo, func(i, j int) bool { return scoreMap[allergicTo[i]] < scoreMap[allergicTo[j]] })
	return allergicTo
}

// AllergicTo checks if a certain score contains a given allergen
func AllergicTo(score uint, allergen string) bool {
	// use a bit-wise AND to check if the corresponding bits of the allergen are set to 1
	return score&scoreMap[allergen] > 0
}
