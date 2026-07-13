// Package timefmt renders human-readable calendar durations for API
// responses and notification texts (RandomWalk offer lifetimes, mint-gap
// tweets). The output quirks — a leading space when the largest unit is
// below a year, and seconds never rendered — are frozen v1 wire format,
// pinned by the store goldens and the rwbot format tests.
package timefmt

import (
	"fmt"
	"time"
)

// TimeDifference decomposes the span between two instants into calendar
// components (years, months, days, hours, minutes, seconds). Argument order
// does not matter; components are always non-negative.
func TimeDifference(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = y2 - y1
	month = int(M2 - M1)
	day = d2 - d1
	hour = h2 - h1
	min = m2 - m1
	sec = s2 - s1

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}
	return
}

// DurationToString renders the calendar components produced by
// TimeDifference as legacy display text ("1 year, 2 months"). Zero
// components are omitted and seconds are never rendered; when the largest
// non-zero unit is smaller than a year the result keeps its historical
// leading space.
func DurationToString(year, month, day, hour, min, sec int) string {
	var output string
	if year > 0 {
		if year == 1 {
			output = fmt.Sprintf("%v year", year)
		} else {
			output = fmt.Sprintf("%v years", year)
		}
	}
	if month > 0 {
		if len(output) > 0 {
			output = output + ","
		}
		if month == 1 {
			output = fmt.Sprintf("%v %v month", output, month)
		} else {
			output = fmt.Sprintf("%v %v months", output, month)
		}
	}
	if day > 0 {
		if len(output) > 0 {
			output = output + ","
		}
		if day == 1 {
			output = fmt.Sprintf("%v %v day", output, day)
		} else {
			output = fmt.Sprintf("%v %v days", output, day)
		}
	}
	if hour > 0 {
		if len(output) > 0 {
			output = output + ","
		}
		if hour == 1 {
			output = fmt.Sprintf("%v %v hour", output, hour)
		} else {
			output = fmt.Sprintf("%v %v hours", output, hour)
		}
	}
	if min > 0 {
		if len(output) > 0 {
			output = output + ","
		}
		if min == 1 {
			output = fmt.Sprintf("%v %v minute", output, min)
		} else {
			output = fmt.Sprintf("%v %v minutes", output, min)
		}
	}

	return output
}
