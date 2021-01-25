package say

import "strings"

var names = map[int64]string{
	0:          "",
	1:          "one",
	2:          "two",
	3:          "three",
	4:          "four",
	5:          "five",
	6:          "six",
	7:          "seven",
	8:          "eight",
	9:          "nine",
	10:         "ten",
	11:         "eleven",
	12:         "twelve",
	13:         "thirteen",
	14:         "fourteen",
	15:         "fifteen",
	16:         "sixteen",
	17:         "seventeen",
	18:         "eighteen",
	19:         "nineteen",
	20:         "twenty",
	30:         "thirty",
	40:         "forty",
	50:         "fifty",
	60:         "sixty",
	70:         "seventy",
	80:         "eighty",
	90:         "ninety",
	100:        "hundred",
	1000:       "thousand",
	1000000:    "million",
	1000000000: "billion",
}

// Say returns the English word for an input number
func Say(input int64) (say string, ok bool) {
	switch {
	case input > 0 && input <= 999999999999:
		return strings.Trim(producer(input), " "), true
	case input == 0:
		return "zero", true
	default:
		return "", false
	}
}

// since Say returns ok as well, we need a separate function which can be called recursively
func producer(n int64) string {
	switch {
	case n <= 20:
		return names[n]
	case n < 1e2:
		return names[(n/10)*10] + "-" + names[n%10]
	case n < 1e3:
		return names[n/100] + " " + names[100] + " " + producer(n%100)
	case n < 1e6:
		return producer(n/1e3) + " " + names[1e3] + " " + producer(n%1e3)
	case n < 1e9:
		return producer(n/1e6) + " " + names[1e6] + " " + producer(n%1e6)
	case n < 1e12:
		return producer(n/1e9) + " " + names[1e9] + " " + producer(n%1e9)
	default:
		return ""
	}
}
