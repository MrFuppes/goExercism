// Package poker has functions to analyze and compare poker hands
package poker

import (
	"errors"
	"math"
	"regexp"
	"sort"
	"strings"
)

// card has a suit and a value
type card struct {
	suit, value int
}

// hand consists of a deck of cards
type hand struct {
	cards           []card
	rank, score     int
	input, category string
}

var (
	values = map[string]int{
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  11,
		"Q":  12,
		"K":  13,
		"A":  14,
	}
	suits = map[string]int{"♤": 1, "♧": 2, "♡": 4, "♢": 8}
)

var (
	errNoHands  = errors.New("no hands, no cookies")
	errIvdValue = errors.New("invalid value")
	errIvdSuit  = errors.New("invalid suit")
)

// BestHand selects the best hand from a slice of hands represented as string.
func BestHand(hands []string) ([]string, error) {
	if len(hands) == 0 {
		return hands, errNoHands
	}

	evaluated := []hand{}
	for _, handString := range hands {
		h, err := parseHandstring(handString)
		if err != nil {
			return []string{}, err
		}
		if len(hands) == 1 {
			return hands, nil
		}
		h = evalHand(h)
		evaluated = append(evaluated, h)
	}

	sort.SliceStable(evaluated, func(i, j int) bool {
		if evaluated[i].rank == evaluated[j].rank {
			return evaluated[i].score > evaluated[j].score
		}
		return evaluated[i].rank < evaluated[j].rank
	})

	winners := []string{evaluated[0].input}
	if evaluated[1].rank == evaluated[0].rank {
		tie := compareTie(evaluated[0], evaluated[1])
		switch tie {
		case 0:
			winners = []string{evaluated[0].input, evaluated[1].input}
		case 1:
			winners = []string{evaluated[0].input}
		case -1:
			winners = []string{evaluated[1].input}
		}
	}

	return winners, nil
}

// parseHandstring parses the string to hand type
func parseHandstring(handString string) (hand, error) {
	parts := strings.Fields(handString)
	if len(parts) != 5 {
		return hand{}, errors.New("must be 5 cards")
	}
	h := hand{}
	reRank := regexp.MustCompile(`^\d+|J|Q|K|A`)
	for _, c := range parts {
		newCard := card{}
		idxR := reRank.FindStringIndex(c)
		if len(idxR) != 2 {
			return hand{}, errIvdValue
		}
		r, s := c[idxR[0]:idxR[1]], c[idxR[1]:len(c)]
		v, ok := values[r]
		if !ok {
			return hand{}, errIvdValue
		}
		newCard.value = v
		v, ok = suits[s]
		if !ok {
			return hand{}, errIvdSuit
		}
		newCard.suit = v
		h.cards = append(h.cards, newCard)
	}
	h.input = handString

	return h, nil
}

// category2rank maps category calculated in evalHand to its rank. lower rank is better.
var (
	categories = []string{"4 of a Kind", "Straight Flush", "Straight", "Flush", "High Card",
		"1 Pair", "2 Pair", "Royal Flush", "3 of a Kind", "Full House"}
	category2rank = map[int]int{
		0: 3,  // 4 of a Kind
		1: 2,  // Straight Flush
		2: 6,  // Straight
		3: 5,  // Flush
		4: 10, // High Card
		5: 9,  // 1 Pair
		6: 8,  // 2 Pair
		7: 1,  // Royal Flush
		8: 7,  // 3 of a Kind
		9: 4,  // Full House
	}
)

// evalHand implements @subskybox's poker analyzer,
// https://www.codeproject.com/Articles/569271/A-Poker-hand-analyzer-in-JavaScript-using-bit-math
func evalHand(h hand) hand {
	var (
		v, o int
		s    = 1<<h.cards[0].value | 1<<h.cards[1].value | 1<<h.cards[2].value | 1<<h.cards[3].value | 1<<h.cards[4].value
	)
	for i := 0; i < 5; i++ {
		o = int(math.Pow(2, float64(h.cards[i].value)*4))
		v += o * ((v / o & 15) + 1)
	}
	v %= 15

	// check straight
	if (s/(s&-s) == 31) || (s == 0x403c) {
		v -= 3           // straight
		if s == 0x403c { // low ace makes the lowest ranked straight
			s = -1
		}
	} else {
		v-- // no straight
	}

	// check flush
	if h.cards[0].suit == (h.cards[1].suit | h.cards[2].suit | h.cards[3].suit | h.cards[4].suit) {
		if s == 0x7c00 {
			v += 5 // royal flush
		} else {
			v-- // flush
		}
	}

	h.score, h.rank, h.category = s, category2rank[v], categories[v]

	return h
}

// compareTie - 1 = this wins / 0 = equal / -1 = other wins
func compareTie(this, other hand) int {
	thisCounts, otherCounts := getCounts(this), getCounts(other)
	thisValue, otherValue := 0, 0
	// some categories require to calculate their value specifically:
	switch this.category {
	case "4 of a Kind":
		thisValue, otherValue = getValue(thisCounts, 4), getValue(otherCounts, 4)
	case "1 Pair", "2 Pair":
		thisValue, otherValue = getValue(thisCounts, 2), getValue(otherCounts, 2)
	case "3 of a Kind", "Full House":
		thisValue, otherValue = getValue(thisCounts, 3), getValue(otherCounts, 3)
	}
	// for all other categories, thisValue == otherValue, so score decides
	switch {
	case thisValue > otherValue:
		return 1
	case thisValue < otherValue:
		return -1
	case this.score > other.score:
		return 1
	case this.score < other.score:
		return -1
	}
	return 0
}

//********* +helpers ***********************************************************
func getCounts(h hand) map[int]int {
	counts := make(map[int]int)
	for _, v := range h.cards {
		counts[v.value]++
	}
	return counts
}

func getValue(c map[int]int, m int) (result int) {
	for k, v := range c {
		if v == m {
			result += m * k
		}
	}
	return result
}
