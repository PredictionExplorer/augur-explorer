package cosmicgame

import (
	"strconv"
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
