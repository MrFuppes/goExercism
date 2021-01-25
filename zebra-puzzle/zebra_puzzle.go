// Package zebra implements a solver for the zebra puzzle
// source: https://rosettacode.org/wiki/Zebra_puzzle#Go
package zebra

// Solution holds the info who drinks water and who owns the zebra
type Solution struct {
	DrinksWater, OwnsZebra string
}

// House accomodates a person of certain nationality and so on
type House struct {
	n Nationality
	c Colour
	a Animal
	d Drink
	s Smoke
}

// HouseSet has all five houses
type HouseSet [5]*House

// the five properties of the houses +  their possible values
type Nationality uint8

const (
	English Nationality = iota
	Swede
	Dane
	Norwegian
	Japanese
)

type Colour uint8

const (
	Red Colour = iota
	Green
	White
	Yellow
	Blue
)

type Animal uint8

const (
	Dog Animal = iota
	Birds
	Cats
	Horse
	Zebra
)

type Drink uint8

const (
	Tea Drink = iota
	Coffee
	Milk
	Beer
	Water
)

type Smoke uint8

const (
	PallMall Smoke = iota
	Dunhill
	Blend
	BlueMaster
	Prince
)

// string representations of the properties
var (
	nationalities = [...]string{"English", "Swede", "Dane", "Norwegian", "Japanese"}
	colours       = [...]string{"red", "green", "white", "yellow", "blue"}
	animals       = [...]string{"dog", "birds", "cats", "horse", "zebra"}
	drinks        = [...]string{"tea", "coffee", "milk", "beer", "water"}
	smokes        = [...]string{"Pall Mall", "Dunhill", "Blend", "Blue Master", "Prince"}
)

// simpleBruteForce solution
func simpleBruteForce() (int, HouseSet) {
	var v []House
	for n := range nationalities {
		for c := range colours {
			for a := range animals {
				for d := range drinks {
					for s := range smokes {
						h := House{
							n: Nationality(n),
							c: Colour(c),
							a: Animal(a),
							d: Drink(d),
							s: Smoke(s),
						}
						if !h.Valid() {
							continue
						}
						v = append(v, h)
					}
				}
			}
		}
	}

	n := len(v)
	valid := 0
	var validSet HouseSet
	for a := 0; a < n; a++ {
		if v[a].n != Norwegian { // Condition 10:
			continue
		}
		for b := 0; b < n; b++ {
			if b == a {
				continue
			}
			if v[b].hasDupeAttr(&v[a]) {
				continue
			}
			for c := 0; c < n; c++ {
				if c == b || c == a {
					continue
				}
				if v[c].d != Milk { // Condition 9:
					continue
				}
				if v[c].hasDupeAttr(&v[b], &v[a]) {
					continue
				}
				for d := 0; d < n; d++ {
					if d == c || d == b || d == a {
						continue
					}
					if v[d].hasDupeAttr(&v[c], &v[b], &v[a]) {
						continue
					}
					for e := 0; e < n; e++ {
						if e == d || e == c || e == b || e == a {
							continue
						}
						if v[e].hasDupeAttr(&v[d], &v[c], &v[b], &v[a]) {
							continue
						}
						set := HouseSet{&v[a], &v[b], &v[c], &v[d], &v[e]}
						if set.Valid() {
							valid++

							validSet = set
						}
					}
				}
			}
		}
	}
	return valid, validSet
}

// hasDupeAttr returns true if House h as any duplicate attributes with any of the houses in list
func (h *House) hasDupeAttr(list ...*House) bool {
	for _, b := range list {
		if h.n == b.n || h.c == b.c || h.a == b.a || h.d == b.d || h.s == b.s {
			return true
		}
	}
	return false
}

// Valid checks if house h has valid attributes according to the given conditions
func (h *House) Valid() bool {
	// Condition 2:
	if h.n == English && h.c != Red || h.n != English && h.c == Red {
		return false
	}
	// Condition 3:
	if h.n == Swede && h.a != Dog || h.n != Swede && h.a == Dog {
		return false
	}
	// Condition 4:
	if h.n == Dane && h.d != Tea || h.n != Dane && h.d == Tea {
		return false
	}
	// Condition 6:
	if h.c == Green && h.d != Coffee || h.c != Green && h.d == Coffee {
		return false
	}
	// Condition 7:
	if h.a == Birds && h.s != PallMall || h.a != Birds && h.s == PallMall {
		return false
	}
	// Condition 8:
	if h.c == Yellow && h.s != Dunhill || h.c != Yellow && h.s == Dunhill {
		return false
	}
	// Condition 11:
	if h.a == Cats && h.s == Blend {
		return false
	}
	// Condition 12:
	if h.a == Horse && h.s == Dunhill {
		return false
	}
	// Condition 13:
	if h.d == Beer && h.s != BlueMaster || h.d != Beer && h.s == BlueMaster {
		return false
	}
	// Condition 14:
	if h.n == Japanese && h.s != Prince || h.n != Japanese && h.s == Prince {
		return false
	}
	// Condition 15:
	if h.n == Norwegian && h.c == Blue {
		return false
	}
	// Condition 16:
	if h.d == Water && h.s == Blend {
		return false
	}
	return true
}

// Valid checks if the set of houses hs has valid attributes according to the given conditions
func (hs *HouseSet) Valid() bool {
	ni := make(map[Nationality]int, 5)
	ci := make(map[Colour]int, 5)
	ai := make(map[Animal]int, 5)
	di := make(map[Drink]int, 5)
	si := make(map[Smoke]int, 5)
	for i, h := range hs {
		ni[h.n] = i
		ci[h.c] = i
		ai[h.a] = i
		di[h.d] = i
		si[h.s] = i
	}
	// Condition 5:
	if ci[Green]+1 != ci[White] {
		return false
	}
	// Condition 11:
	if dist(ai[Cats], si[Blend]) != 1 {
		return false
	}
	// Condition 12:
	if dist(ai[Horse], si[Dunhill]) != 1 {
		return false
	}
	// Condition 15:
	if dist(ni[Norwegian], ci[Blue]) != 1 {
		return false
	}
	// Condition 16:
	if dist(di[Water], si[Blend]) != 1 {
		return false
	}
	// Condition 9: (already tested elsewhere)
	if hs[2].d != Milk {
		return false
	}
	// Condition 10: (already tested elsewhere)
	if hs[0].n != Norwegian {
		return false
	}
	return true
}

// dist - a helper to return the distance of an attribute within a house set
func dist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// SolvePuzzle calls the solver
func SolvePuzzle() Solution {
	_, result := simpleBruteForce()
	solution := Solution{}
	for _, house := range result {
		if house.d == Water {
			solution.DrinksWater = nationalities[house.n]
		}
		if house.a == Zebra {
			solution.OwnsZebra = nationalities[house.n]
		}
	}
	return solution
}
