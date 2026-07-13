package timefmt

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

func TestTimeDifferenceMixedLocations(t *testing.T) {
	// The second instant is converted into the first one's location before
	// decomposing, so the same absolute span yields the same components.
	est := time.FixedZone("EST", -5*3600)
	a := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	b := time.Date(2024, 1, 1, 1, 0, 0, 0, est) // 06:00 UTC
	y, mo, d, h, mi, s := TimeDifference(a, b)
	if y != 0 || mo != 0 || d != 0 || h != 6 || mi != 0 || s != 0 {
		t.Errorf("TimeDifference across locations = (%d,%d,%d,%d,%d,%d), want (0,0,0,6,0,0)", y, mo, d, h, mi, s)
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

func FuzzDateUtils(f *testing.F) {
	f.Add(int64(0), int64(0))
	f.Add(int64(0), int64(1))
	f.Add(int64(0), int64(86400))
	f.Add(int64(1580428800), int64(1583020800)) // 2020-01-31 → 2020-03-01
	f.Add(int64(-62135596800), int64(253402300799))
	f.Fuzz(func(t *testing.T, aSec, bSec int64) {
		a := time.Unix(aSec, 0).UTC()
		b := time.Unix(bSec, 0).UTC()
		y, mo, d, h, mi, s := TimeDifference(a, b)
		// All components normalized and non-negative.
		if y < 0 || mo < 0 || d < 0 || h < 0 || mi < 0 || s < 0 {
			t.Fatalf("TimeDifference(%v, %v) = (%d,%d,%d,%d,%d,%d): negative component", a, b, y, mo, d, h, mi, s)
		}
		if mo > 11 || d > 30 || h > 23 || mi > 59 || s > 59 {
			t.Fatalf("TimeDifference(%v, %v) = (%d,%d,%d,%d,%d,%d): component out of range", a, b, y, mo, d, h, mi, s)
		}
		// Symmetric in argument order.
		y2, mo2, d2, h2, mi2, s2 := TimeDifference(b, a)
		if y != y2 || mo != mo2 || d != d2 || h != h2 || mi != mi2 || s != s2 {
			t.Fatalf("TimeDifference not symmetric for (%v, %v)", a, b)
		}
		if aSec == bSec && (y|mo|d|h|mi|s) != 0 {
			t.Fatalf("TimeDifference of equal times = (%d,%d,%d,%d,%d,%d), want all zero", y, mo, d, h, mi, s)
		}
		// DurationToString must never panic on TimeDifference output.
		_ = DurationToString(y, mo, d, h, mi, s)
	})
}
