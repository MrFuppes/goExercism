// Package clock implements type and methods for a time-only clock.
package clock

import "fmt"

// Clock - stores time in integer minutes
type Clock struct {
	totalMinutes int
}

// minPerDay - maximum number of minutes in a day
const minPerDay = 1440

// New - a "constructor" for the clock type
func New(hour, minute int) Clock {
	return Clock{0}.Add(hour*60 + minute)
}

// Add - add minutes to the clock
func (c Clock) Add(minutes int) Clock {
	c.totalMinutes += minutes
	if c.totalMinutes < 0 {
		c.totalMinutes = minPerDay + c.totalMinutes%minPerDay
	}
	if c.totalMinutes >= minPerDay {
		c.totalMinutes -= minPerDay * (c.totalMinutes / minPerDay)
	}
	return c
}

// Subtract - remove minutes from the clock.
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(minutes * -1)
}

// String - a "stringer" for the clock type
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.totalMinutes/60, c.totalMinutes%60)
}
