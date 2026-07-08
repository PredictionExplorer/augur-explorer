package cosmicgame

import (
	"strconv"
	"strings"
	"testing"
)

func TestParseOptionalIntQuery(t *testing.T) {
	cases := []struct {
		in         string
		defaultVal int
		want       int
	}{
		{"", 5, 5},
		{"7", 5, 7},
		{"-3", 5, -3},
		{"0", 5, 0},
		{"+7", 5, 7},
		{"abc", 5, 5},
		{"7.5", 5, 5},
		{" 7", 5, 5},                     // Atoi rejects surrounding whitespace
		{"999999999999999999999", 5, 5},  // out of int range
		{"-999999999999999999999", 5, 5}, // out of int range
	}
	for _, tc := range cases {
		if got := ParseOptionalIntQuery(tc.in, tc.defaultVal); got != tc.want {
			t.Errorf("ParseOptionalIntQuery(%q, %d) = %d, want %d", tc.in, tc.defaultVal, got, tc.want)
		}
	}
}

func FuzzParseOptionalIntQuery(f *testing.F) {
	f.Add("", 0)
	f.Add("42", 7)
	f.Add("-1", 7)
	f.Add("junk", 7)
	f.Add("999999999999999999999", 7)
	f.Fuzz(func(t *testing.T, s string, defaultVal int) {
		got := ParseOptionalIntQuery(s, defaultVal)
		want := defaultVal
		if v, err := strconv.Atoi(s); err == nil {
			want = v
		}
		if got != want {
			t.Fatalf("ParseOptionalIntQuery(%q, %d) = %d, want %d", s, defaultVal, got, want)
		}
	})
}

// TestBidFrequencySQLBranchSelection pins which sampling intervals get the
// UTC epoch-aligned bucket query (hour and day windows) versus the
// initTs-anchored one, and that each branch binds the parameters it
// documents: the anchored query takes the interval as a third bound
// parameter, the epoch-aligned one interpolates it as an integer literal.
func TestBidFrequencySQLBranchSelection(t *testing.T) {
	cases := []struct {
		intervalSecs     int
		wantEpochAligned bool
	}{
		{3600, true},
		{86400, true},
		{1, false},
		{900, false},
		{3599, false},
		{3601, false},
		{86401, false},
		{999999, false},
	}
	for _, tc := range cases {
		query, epochAligned := bidFrequencySQL(tc.intervalSecs)
		if epochAligned != tc.wantEpochAligned {
			t.Errorf("bidFrequencySQL(%d): epochAligned = %v, want %v",
				tc.intervalSecs, epochAligned, tc.wantEpochAligned)
		}
		if !strings.Contains(query, "$1") || !strings.Contains(query, "$2") {
			t.Errorf("bidFrequencySQL(%d): query must bind $1/$2 range params", tc.intervalSecs)
		}
		if epochAligned {
			if strings.Contains(query, "$3") {
				t.Errorf("bidFrequencySQL(%d): epoch-aligned query must not reference $3", tc.intervalSecs)
			}
		} else {
			if !strings.Contains(query, "($3 || ' seconds')::interval") {
				t.Errorf("bidFrequencySQL(%d): anchored query must bind the interval as $3", tc.intervalSecs)
			}
		}
		if !strings.Contains(query, "cg_round_stats") {
			t.Errorf("bidFrequencySQL(%d): first-hour-after-round-start exclusion missing", tc.intervalSecs)
		}
	}
}

// TestBidFrequencySQLInterpolationIsNumeric proves the only value ever
// interpolated into the epoch-aligned query text is the decimal rendering of
// the interval: strip the two known literals and the remaining text must be
// identical for any two intervals (i.e. nothing else varies with input).
func TestBidFrequencySQLInterpolationIsNumeric(t *testing.T) {
	a, _ := bidFrequencySQL(3600)
	b, _ := bidFrequencySQL(86400)
	excl := excludeFirstHourAfterRoundStartSQL() // shared literal, contains "3600" itself
	normalized := func(q, interval string) string {
		q = strings.ReplaceAll(q, excl, "<EXCL>")
		return strings.ReplaceAll(q, interval, "<N>")
	}
	if normalized(a, "3600") != normalized(b, "86400") {
		t.Error("epoch-aligned queries differ beyond the interval literal")
	}
}
