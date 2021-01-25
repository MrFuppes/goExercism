// Package gigasecond implements functions dealing with gigaseconds.
package gigasecond

import (
	"time"
)

const gs = time.Second * 1e9

// AddGigasecond returns the time 1 Gs after the input time t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gs)
}
