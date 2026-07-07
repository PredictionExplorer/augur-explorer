package primitives

import (
	"math"
	"testing"
)

func TestShortAddress(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"valid", "0x1234567890abcdef1234567890abcdef12345678", "123456…345678"},
		{"valid checksummed", "0xC0ffee254729296a45a3885639AC7E10F9d54979", "C0ffee…d54979"},
		{"empty", "", "inval_addr"},
		{"too short", "0x1234", "inval_addr"},
		{"41 chars", "0x123456789012345678901234567890123456789", "inval_addr"},
		{"43 chars", "0x12345678901234567890123456789012345678901", "inval_addr"},
	}
	for _, tc := range cases {
		if got := Short_address(tc.in); got != tc.want {
			t.Errorf("%s: Short_address(%q) = %q, want %q", tc.name, tc.in, got, tc.want)
		}
	}
}

func TestShortHash(t *testing.T) {
	// Note the existing quirk: the tail slice is [59:65], so the final
	// character of a 66-char hash is not included in the output.
	in := "0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"
	want := "abcdef…345678" // in[59:65] == "345678"; in[65] ('9', the last char) is dropped
	if got := Short_hash(in); got != want {
		t.Errorf("Short_hash(%q) = %q, want %q", in, got, want)
	}
	if got := Short_hash("0x1234"); got != "inval_hash: 6" {
		t.Errorf("Short_hash short input = %q, want %q", got, "inval_hash: 6")
	}
	if got := Short_hash(""); got != "inval_hash: 0" {
		t.Errorf("Short_hash empty input = %q, want %q", got, "inval_hash: 0")
	}
}

func TestThousandsFormat(t *testing.T) {
	cases := []struct {
		in   int64
		want string
	}{
		{0, "0"},
		{1, "1"},
		{999, "999"},
		{1000, "1,000"},
		{999999, "999,999"},
		{1000000, "1,000,000"},
		{1234567890, "1,234,567,890"},
		{-1, "-1"},
		{-999, "-999"},
		{-1000, "-1,000"},
		{-1234567, "-1,234,567"},
		{math.MaxInt64, "9,223,372,036,854,775,807"},
		{math.MinInt64, "-9,223,372,036,854,775,808"},
	}
	for _, tc := range cases {
		if got := ThousandsFormat(tc.in); got != tc.want {
			t.Errorf("ThousandsFormat(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestRemoveDuplicatesInt64(t *testing.T) {
	// The helper removes ADJACENT duplicates only (callers pass sorted input).
	cases := []struct {
		name string
		in   []int64
		want []int64
	}{
		{"empty", nil, nil},
		{"sorted with dups", []int64{1, 1, 2, 3, 3}, []int64{1, 2, 3}},
		{"no dups", []int64{1, 2, 3}, []int64{1, 2, 3}},
		{"all same", []int64{7, 7, 7}, []int64{7}},
		{"unsorted keeps non-adjacent dups", []int64{1, 2, 1}, []int64{1, 2, 1}},
	}
	for _, tc := range cases {
		n := Remove_duplicates_int64(tc.in)
		if n != len(tc.want) {
			t.Errorf("%s: Remove_duplicates_int64(%v) = %d, want %d", tc.name, tc.in, n, len(tc.want))
			continue
		}
		for i := 0; i < n; i++ {
			if tc.in[i] != tc.want[i] {
				t.Errorf("%s: element %d = %d, want %d", tc.name, i, tc.in[i], tc.want[i])
			}
		}
	}
}
