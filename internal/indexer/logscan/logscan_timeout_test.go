package logscan

import (
	"context"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

// TestScanBoundsEachFetchAndRetries proves a black-holed FilterLogs call
// fails at the per-call deadline and is retried like any fetch error — the
// scan completes instead of hanging forever on one call.
func TestScanBoundsEachFetchAndRetries(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var calls int
		client := filtererFunc(func(ctx context.Context, _ ethereum.FilterQuery) ([]types.Log, error) {
			calls++
			if _, ok := ctx.Deadline(); !ok {
				t.Errorf("fetch %d carries no deadline", calls)
			}
			if calls == 1 {
				<-ctx.Done() // black hole until the per-call bound fires
				return nil, ctx.Err()
			}
			return nil, nil
		})

		opts := validOptions()
		opts.FetchTimeout = time.Second
		start := time.Now()
		stats, err := Scan(t.Context(), client, opts, func(context.Context, types.Log) error { return nil })
		if err != nil {
			t.Fatalf("Scan: %v (a timed-out fetch must be retried, not returned)", err)
		}
		if stats.FilterErrors != 1 {
			t.Fatalf("FilterErrors = %d, want the one bounded hang", stats.FilterErrors)
		}
		// One 1s bounded hang; the healthy retries are instant in fake time.
		if elapsed := time.Since(start); elapsed != time.Second {
			t.Fatalf("Scan took %v of fake time, want exactly the 1s bounded hang", elapsed)
		}
	})
}

// TestScanDefaultFetchTimeoutApplied pins that an unset FetchTimeout still
// bounds the fetch (with DefaultFetchTimeout) rather than running unbounded.
func TestScanDefaultFetchTimeoutApplied(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var sawDeadline time.Time
		client := filtererFunc(func(ctx context.Context, _ ethereum.FilterQuery) ([]types.Log, error) {
			d, ok := ctx.Deadline()
			if !ok {
				t.Error("fetch carries no deadline with FetchTimeout unset")
			}
			sawDeadline = d
			return nil, nil
		})

		start := time.Now()
		if _, err := Scan(t.Context(), client, validOptions(), func(context.Context, types.Log) error { return nil }); err != nil {
			t.Fatalf("Scan: %v", err)
		}
		if got := sawDeadline.Sub(start); got != DefaultFetchTimeout {
			t.Fatalf("per-fetch deadline = %v from start, want DefaultFetchTimeout %v (fake time is exact)", got, DefaultFetchTimeout)
		}
	})
}

func TestScanRejectsNegativeFetchTimeout(t *testing.T) {
	t.Parallel()
	opts := validOptions()
	opts.FetchTimeout = -time.Second
	_, err := Scan(context.Background(),
		filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) { return nil, nil }),
		opts, func(context.Context, types.Log) error { return nil })
	if err == nil || !strings.Contains(err.Error(), "fetch timeout") {
		t.Fatalf("err = %v, want the fetch-timeout validation error", err)
	}
}
