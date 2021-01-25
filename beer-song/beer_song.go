package beer

import (
	"errors"
	"fmt"
	"strings"
)

const (
	gotBeer    = "%v bottles of beer on the wall, %v bottles of beer.\n"
	drinkBeer  = "Take one down and pass it around, %v bottles of beer on the wall.\n"
	beerGone   = "Take it down and pass it around, no more bottles of beer on the wall.\n"
	needRefill = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
)

// Verse of the beer song
func Verse(n int) (string, error) {
	if 99 < n || n < 0 {
		return "", errors.New("invalid verse number")
	}
	if n == 0 {
		return needRefill, nil
	}
	var result, part string

	part = fmt.Sprintf(gotBeer, n, n)
	if n == 1 {
		part = strings.ReplaceAll(part, "bottles", "bottle")
	}
	result += part

	if n-1 == 0 {
		result += beerGone
	} else {
		part = fmt.Sprintf(drinkBeer, n-1)
		if n-1 == 1 {
			part = strings.ReplaceAll(part, "bottles", "bottle")
		}
		result += part
	}

	return result, nil
}

// Verses returns multiple verses
func Verses(upper, lower int) (string, error) {
	if 99 < upper || lower < 0 || lower > upper {
		return "", errors.New("invalid bounds")
	}
	var verses []string
	for i := upper; i >= lower; i-- {
		v, _ := Verse(i)
		verses = append(verses, v)
	}
	return strings.Join(verses, "\n") + "\n", nil
}

// Song returns all verses 99 to 0. skips possible error.
func Song() string {
	s, _ := Verses(99, 0)
	return s
}
