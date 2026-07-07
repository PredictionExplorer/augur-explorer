package primitives

import (
	"testing"
	"time"
)

func TestTimeDifference(t *testing.T) {
	utc := func(y int, m time.Month, d, hh, mm, ss int) time.Time {
		return time.Date(y, m, d, hh, mm, ss, 0, time.UTC)
	}
	cases := []struct {
		name                             string
		a, b                             time.Time
		year, month, day, hour, min, sec int
	}{
		{"zero", utc(2024, 1, 1, 0, 0, 0), utc(2024, 1, 1, 0, 0, 0), 0, 0, 0, 0, 0, 0},
		{"one second", utc(2024, 1, 1, 0, 0, 0), utc(2024, 1, 1, 0, 0, 1), 0, 0, 0, 0, 0, 1},
		{"month rollover", utc(2020, 1, 31, 0, 0, 0), utc(2020, 3, 1, 0, 0, 0), 0, 1, 1, 0, 0, 0},
		{"across midnight", utc(2024, 1, 31, 23, 59, 59), utc(2024, 2, 1, 0, 0, 0), 0, 0, 0, 0, 0, 1},
		{"leap day span", utc(2024, 2, 28, 0, 0, 0), utc(2024, 3, 1, 0, 0, 0), 0, 0, 2, 0, 0, 0},
		{"full year", utc(2023, 6, 15, 12, 30, 45), utc(2024, 6, 15, 12, 30, 45), 1, 0, 0, 0, 0, 0},
	}
	for _, tc := range cases {
		y, mo, d, h, mi, s := TimeDifference(tc.a, tc.b)
		if y != tc.year || mo != tc.month || d != tc.day || h != tc.hour || mi != tc.min || s != tc.sec {
			t.Errorf("%s: TimeDifference = (%d,%d,%d,%d,%d,%d), want (%d,%d,%d,%d,%d,%d)",
				tc.name, y, mo, d, h, mi, s, tc.year, tc.month, tc.day, tc.hour, tc.min, tc.sec)
		}
		// Argument order must not matter (the function swaps internally).
		y2, mo2, d2, h2, mi2, s2 := TimeDifference(tc.b, tc.a)
		if y != y2 || mo != mo2 || d != d2 || h != h2 || mi != mi2 || s != s2 {
			t.Errorf("%s: TimeDifference is not symmetric", tc.name)
		}
	}
}

func TestDurationToString(t *testing.T) {
	cases := []struct {
		name                             string
		year, month, day, hour, min, sec int
		want                             string
	}{
		{"all zero", 0, 0, 0, 0, 0, 0, ""},
		{"one year", 1, 0, 0, 0, 0, 0, "1 year"},
		{"two years", 2, 0, 0, 0, 0, 0, "2 years"},
		{"year and months", 1, 2, 0, 0, 0, 0, "1 year, 2 months"},
		// Existing quirk: when the largest non-zero unit is smaller than a year,
		// the output starts with a leading space (Sprintf("%v %v ...", "", n)).
		{"months only", 0, 2, 0, 0, 0, 0, " 2 months"},
		{"single minute", 0, 0, 0, 0, 1, 0, " 1 minute"},
		{"full chain", 1, 1, 1, 1, 1, 0, "1 year, 1 month, 1 day, 1 hour, 1 minute"},
		// Existing quirk: seconds are never rendered.
		{"seconds ignored", 0, 0, 0, 0, 0, 30, ""},
	}
	for _, tc := range cases {
		got := DurationToString(tc.year, tc.month, tc.day, tc.hour, tc.min, tc.sec)
		if got != tc.want {
			t.Errorf("%s: DurationToString = %q, want %q", tc.name, got, tc.want)
		}
	}
}
