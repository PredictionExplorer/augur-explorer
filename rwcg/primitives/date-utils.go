package primitives

import (
	"fmt"
	"time"
)
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

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

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
/*	fmt.Printf(
		"year = %v, month=%v, day = %v, hour=%v, min=%v, sec=%v\n",
		year,month,day,hour,min,sec,
	)*/
	return
}
func DurationToString(year, month, day, hour, min, sec int) string {

	var output string
	if year > 0 {
		if year == 1 {
			output = fmt.Sprintf("%v year",year)
		} else {
			output = fmt.Sprintf("%v years",year)
		}
	}
	if month > 0 {
		if len(output) >0 {
			output = output + ","
		}
		if month == 1 {
			output = fmt.Sprintf("%v %v month",output,month)
		} else {
			output = fmt.Sprintf("%v %v months",output,month)
		}
	}
	if day > 0 {
		if len(output)>0 {
			output = output + ","
		}
		if day == 1 {
			output = fmt.Sprintf("%v %v day",output,day)
		} else {
			output = fmt.Sprintf("%v %v days",output,day)
		}
	}
	if hour > 0 {
		if len(output) >0 {
			output = output + ","
		}
		if hour == 1 {
			output = fmt.Sprintf("%v %v hour",output,hour)
		} else {
			output = fmt.Sprintf("%v %v hours",output,hour)
		}
	}
	if min > 0 {
		if len(output) > 0 {
			output = output + ","
		}
		if min == 1 {
			output = fmt.Sprintf("%v %v minute",output,min)
		} else {
			output = fmt.Sprintf("%v %v minutes",output,min)
		}
	}

	return output
}
