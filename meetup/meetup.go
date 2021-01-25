package meetup

import "time"

// WeekSchedule - made-up name for the n-th day
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth WeekSchedule = 13
)

// Day returns the day of the month of the meetup.
func Day(ws WeekSchedule, wd time.Weekday, m time.Month, y int) int {
	t0 := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	matchDoM := []int{}

	for t0.Month() == m {
		if t0.Weekday() == wd {
			matchDoM = append(matchDoM, t0.Day())
		}
		t0 = t0.Add(time.Hour * 24)
	}

	switch ws {
	case Last:
		return matchDoM[len(matchDoM)-1]
	case Teenth:
		for _, d := range matchDoM {
			if d >= int(ws) {
				return d
			}
		}
	}
	return matchDoM[ws]
}
