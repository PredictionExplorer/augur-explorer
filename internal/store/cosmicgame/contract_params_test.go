package cosmicgame

import "testing"

func TestDecimalStringsEqual(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"100", "100", true},
		{"0100", "100", true},
		{"100000000000000000000", "100000000000000000000", true},
		{"100", "101", false},
		{"not-a-number", "not-a-number", true},
		{"not-a-number", "100", false},
	}
	for _, tc := range cases {
		got := DecimalStringsEqual(tc.a, tc.b)
		if got != tc.want {
			t.Fatalf("DecimalStringsEqual(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}
