// Package space implements functionality to convert an age in seconds to an age
// in years on different planets.
package space

// Planet - a type to hold the name of a Planet.
type Planet string

// YearFactor - a method for the Planet type that returns the relative
// years on earth for a given planet p. If the name of the planet is not known,
// the function returns nil.
func (p Planet) YearFactor() float64 {
	mapping := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return mapping[p]
}

// EarthDaysPerYear - days per year on earth according to NASA, see https://pumas.nasa.gov/files/04_21_97_1.pdf
const EarthDaysPerYear float64 = 365.2422

// Age - a function to derive the age in years on a planet for a given age in seconds.
func Age(seconds float64, planet Planet) float64 {
	return (seconds / (60 * 60 * 24 * EarthDaysPerYear)) / Planet.YearFactor(planet)
}
