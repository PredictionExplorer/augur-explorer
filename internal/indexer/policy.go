// Batch-sizing and retry-backoff policies of the polling loop, kept as pure
// types so their transitions are unit-testable without a chain or database.

package indexer

import (
	"math/rand/v2"
	"time"
)

// batchPolicy adapts the FilterLogs block-range size: large while scanning
// empty history, small while events are flowing so failures re-process little.
type batchPolicy struct {
	size, min, max uint64
}

func newBatchPolicy(cfg BatchConfig) *batchPolicy {
	cfg = cfg.withDefaults()
	return &batchPolicy{size: cfg.Initial, min: cfg.Min, max: cfg.Max}
}

// onEvents shrinks to the minimum: the range carries activity, keep batches
// granular.
func (b *batchPolicy) onEvents() { b.size = b.min }

// onEmpty doubles the size (up to max) to scan empty history faster.
func (b *batchPolicy) onEmpty() {
	b.size *= 2
	if b.size > b.max {
		b.size = b.max
	}
}

// onFetchError halves the size (down to min): oversized ranges are the most
// common cause of FilterLogs failures.
func (b *batchPolicy) onFetchError() {
	b.size /= 2
	if b.size < b.min {
		b.size = b.min
	}
}

// onCaughtUp resets to the minimum for real-time tailing.
func (b *batchPolicy) onCaughtUp() { b.size = b.min }

// backoffDelay returns the delay before retry attempt (1-based): exponential
// growth from min capped at max, with ±25% jitter so restarting replicas do
// not synchronize against a struggling upstream.
func backoffDelay(attempt int, minDelay, maxDelay time.Duration, randFloat func() float64) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	d := minDelay
	for i := 1; i < attempt; i++ {
		d *= 2
		if d >= maxDelay {
			d = maxDelay
			break
		}
	}
	if d > maxDelay {
		d = maxDelay
	}
	// Jitter in [0.75, 1.25).
	jitter := 0.75 + 0.5*randFloat()
	d = time.Duration(float64(d) * jitter)
	if d > maxDelay {
		d = maxDelay
	}
	if d < 0 {
		d = 0
	}
	return d
}

// randFloat is the package's jitter source; tests pass their own into
// backoffDelay directly.
func randFloat() float64 { return rand.Float64() } // #nosec G404 -- retry jitter needs no cryptographic randomness
