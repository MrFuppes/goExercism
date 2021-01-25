package foodchain

import (
	"fmt"
	"strings"
)

var (
	opening     = "I know an old lady who swallowed a %s.\n"
	creatures   = [...]string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
	creaturesDo = [...]string{"",
		"It wriggled and jiggled and tickled inside her.\n",
		"How absurd to swallow a bird!\n",
		"Imagine that, to swallow a cat!\n",
		"What a hog, to swallow a dog!\n",
		"Just opened her throat and swallowed a goat!\n",
		"I don't know how she swallowed a cow!\n",
		""}
	closing = "I don't know why she swallowed the fly. Perhaps she'll die."
)

const (
	maxVerses = 8
	end       = "She's dead, of course!"
)

// Verse returns the n-th verse
func Verse(n int) string {
	if n < 1 {
		return "" // too short
	}
	out := fmt.Sprintf(opening, creatures[n-1])

	if n >= maxVerses {
		return out + end // the end has come
	}

	out += creaturesDo[n-1] // unique line after first one...

	// rest can be sort-of auto-generated.
	for i := n; i > 1; i-- {
		out += fmt.Sprintf("She swallowed the %s to catch the %s.\n", creatures[i-1], creatures[i-2])
		if strings.HasSuffix(out, "spider.\n") { // spider needs a hack!
			out = strings.Replace(out, "spider.\n", "spider that wriggled and jiggled and tickled inside her.\n", -1)
		}
	}

	out += closing
	return out
}

// Verses returns verses n to m (inclusive)
func Verses(n, m int) string {
	var out []string
	for i := n; i <= m; i++ {
		out = append(out, Verse(i))
	}
	return strings.Join(out, "\n\n")
}

// Song returns the whole song
func Song() string {
	return Verses(1, 8)
}
