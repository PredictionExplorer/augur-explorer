package main

import (
	"math"
	"strings"
	"testing"
)

// TestBlockNumFromWatermark pins the negative-watermark guard: a corrupt
// negative block number must abort the backfill instead of wrapping into an
// astronomically large scan end.
func TestBlockNumFromWatermark(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		in   int64
		want uint64
	}{
		{0, 0},
		{1, 1},
		{455_767_500, 455_767_500},
		{math.MaxInt64, math.MaxInt64},
	} {
		got, err := blockNumFromWatermark("processing status", tc.in)
		if err != nil {
			t.Errorf("blockNumFromWatermark(%d): unexpected error %v", tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("blockNumFromWatermark(%d) = %d, want %d", tc.in, got, tc.want)
		}
	}

	for _, in := range []int64{-1, math.MinInt64} {
		if _, err := blockNumFromWatermark("last_block", in); err == nil {
			t.Errorf("blockNumFromWatermark(%d): want error, got nil", in)
		} else if !strings.Contains(err.Error(), "last_block") || !strings.Contains(err.Error(), "negative") {
			t.Errorf("blockNumFromWatermark(%d): error %q must name the source and the defect", in, err)
		}
	}
}
