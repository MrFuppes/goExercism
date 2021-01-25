package twelve

import (
	"fmt"
	"strings"
)

var (
	start    = "On the %s day of Christmas my true love gave to me: "
	join     = ""
	nthDay   = [...]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	nthCount = [...]string{"a", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}
	what     = [...]string{"Partridge in a Pear Tree", "Turtle Doves", "French Hens", "Calling Birds", "Gold Rings", "Geese-a-Laying", "Swans-a-Swimming", "Maids-a-Milking", "Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming"}
)

// Verse creates the n-th verse of the song
func Verse(n int) string {
	verse := fmt.Sprintf(start, nthDay[n-1])
	for i := n - 1; i >= 0; i-- {
		switch i {
		case 0:
			join = "."
		case 1:
			join = ", and "
		default:
			join = ", "
		}
		verse += fmt.Sprintf("%s %s%s", nthCount[i], what[i], join)
	}
	return verse
}

// Song creates the entire song
func Song() string {
	song := make([]string, 12)
	for i := 1; i <= 12; i++ {
		song[i-1] = Verse(i)
	}
	return strings.Join(song, "\n")
}
