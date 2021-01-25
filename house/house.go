package house

import "strings"

var lines = [12]string{ // the trick is correct placement of newlines...
	"the house that Jack built.",
	"the malt\nthat lay in ",
	"the rat\nthat ate ",
	"the cat\nthat killed ",
	"the dog\nthat worried ",
	"the cow with the crumpled horn\nthat tossed ",
	"the maiden all forlorn\nthat milked ",
	"the man all tattered and torn\nthat kissed ",
	"the priest all shaven and shorn\nthat married ",
	"the rooster that crowed in the morn\nthat woke ",
	"the farmer sowing his corn\nthat kept ",
	"the horse and the hound and the horn\nthat belonged to ",
}

// Verse creates the n-th verse of the song
func Verse(n int) string {
	verse := "This is "
	for i := n - 1; i >= 0; i-- {
		verse += lines[i]
	}
	return verse
}

// Song creates the entire song
func Song() string {
	song := make([]string, 12)
	for i := 1; i <= 12; i++ {
		song[i-1] = Verse(i)
	}
	return strings.Join(song, "\n\n")
}
