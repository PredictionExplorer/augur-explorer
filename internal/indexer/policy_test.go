package indexer

import (
	"testing"
	"time"
)

func TestBatchPolicyDefaults(t *testing.T) {
	b := newBatchPolicy(BatchConfig{})
	if b.size != 100_000 || b.min != 1_000 || b.max != 1_000_000 {
		t.Fatalf("defaults = %d/%d/%d, want 100000/1000/1000000", b.size, b.min, b.max)
	}
}

func TestBatchPolicyTransitions(t *testing.T) {
	b := newBatchPolicy(BatchConfig{Initial: 100, Min: 10, Max: 400})

	b.onEmpty()
	if b.size != 200 {
		t.Errorf("after onEmpty: size = %d, want 200 (doubled)", b.size)
	}
	b.onEmpty()
	b.onEmpty() // 400 -> capped at max
	if b.size != 400 {
		t.Errorf("after repeated onEmpty: size = %d, want max 400", b.size)
	}

	b.onFetchError()
	if b.size != 200 {
		t.Errorf("after onFetchError: size = %d, want 200 (halved)", b.size)
	}
	for range 10 {
		b.onFetchError()
	}
	if b.size != 10 {
		t.Errorf("after repeated onFetchError: size = %d, want min 10", b.size)
	}

	b.onEmpty()
	b.onEvents()
	if b.size != 10 {
		t.Errorf("after onEvents: size = %d, want min 10", b.size)
	}

	b.onEmpty()
	b.onCaughtUp()
	if b.size != 10 {
		t.Errorf("after onCaughtUp: size = %d, want min 10", b.size)
	}
}

func TestBackoffDelayGrowsAndCaps(t *testing.T) {
	noJitter := func() float64 { return 0.5 } // jitter factor exactly 1.0

	cases := []struct {
		attempt int
		want    time.Duration
	}{
		{1, time.Second},
		{2, 2 * time.Second},
		{3, 4 * time.Second},
		{4, 8 * time.Second},
		{7, 60 * time.Second}, // 64s uncapped
		{100, 60 * time.Second},
		{0, time.Second}, // clamped to attempt 1
	}
	for _, c := range cases {
		got := backoffDelay(c.attempt, time.Second, time.Minute, noJitter)
		if got != c.want {
			t.Errorf("backoffDelay(attempt=%d) = %v, want %v", c.attempt, got, c.want)
		}
	}
}

func TestBackoffDelayJitterBounds(t *testing.T) {
	low := func() float64 { return 0.0 }
	high := func() float64 { return 0.999999 }

	base := 4 * time.Second // attempt 3 with min 1s
	if got := backoffDelay(3, time.Second, time.Minute, low); got != 3*time.Second {
		t.Errorf("lowest jitter = %v, want %v (0.75x)", got, 3*time.Second)
	}
	if got := backoffDelay(3, time.Second, time.Minute, high); got < base || got >= time.Duration(float64(base)*1.25)+time.Millisecond {
		t.Errorf("highest jitter = %v, want just under %v (1.25x)", got, time.Duration(float64(base)*1.25))
	}
	// Jitter never exceeds the configured cap.
	if got := backoffDelay(50, time.Second, time.Minute, high); got > time.Minute {
		t.Errorf("capped delay with max jitter = %v, want <= 1m", got)
	}
}
