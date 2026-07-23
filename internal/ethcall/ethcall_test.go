package ethcall

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"testing/synctest"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// hangingCaller blocks every read until its context dies.
type hangingCaller struct{ calls int }

func (h *hangingCaller) CodeAt(ctx context.Context, _ common.Address, _ *big.Int) ([]byte, error) {
	h.calls++
	<-ctx.Done()
	return nil, ctx.Err()
}

func (h *hangingCaller) CallContract(ctx context.Context, _ ethereum.CallMsg, _ *big.Int) ([]byte, error) {
	h.calls++
	<-ctx.Done()
	return nil, ctx.Err()
}

// TestBoundedCallerBoundsBothReads proves CodeAt and CallContract return
// context.DeadlineExceeded after exactly the configured bound on a
// black-holed endpoint.
func TestBoundedCallerBoundsBothReads(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const bound = 7 * time.Second
		inner := &hangingCaller{}
		c := NewBoundedCaller(inner, bound)

		start := time.Now()
		if _, err := c.CodeAt(t.Context(), common.Address{0x01}, nil); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("CodeAt err = %v, want DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != bound {
			t.Fatalf("CodeAt returned after %v, want exactly %v", elapsed, bound)
		}

		start = time.Now()
		if _, err := c.CallContract(t.Context(), ethereum.CallMsg{}, nil); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("CallContract err = %v, want DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != bound {
			t.Fatalf("CallContract returned after %v, want exactly %v", elapsed, bound)
		}
		if inner.calls != 2 {
			t.Fatalf("inner caller saw %d calls, want 2", inner.calls)
		}
	})
}

// TestBoundedCallerNilContextInCallOpts pins the load-bearing production
// case: generated bindings pass context.Background() when CallOpts carries
// no context, and the wrapper must bound that too.
func TestBoundedCallerNilContextInCallOpts(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := NewBoundedCaller(&hangingCaller{}, time.Second)
		start := time.Now()
		// Deliberately context.Background(): the exact context bind's
		// ensureContext supplies when CallOpts.Context is nil.
		_, err := c.CallContract(context.Background(), ethereum.CallMsg{}, nil)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("err = %v, want DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != time.Second {
			t.Fatalf("returned after %v, want the 1s bound on a background context", elapsed)
		}
	})
}

func TestNewBoundedCallerAppliesDefaultTimeout(t *testing.T) {
	t.Parallel()
	if got := NewBoundedCaller(&hangingCaller{}, 0).timeout; got != DefaultTimeout {
		t.Fatalf("zero timeout = %v, want DefaultTimeout %v", got, DefaultTimeout)
	}
	if got := NewBoundedCaller(&hangingCaller{}, -time.Second).timeout; got != DefaultTimeout {
		t.Fatalf("negative timeout = %v, want DefaultTimeout %v", got, DefaultTimeout)
	}
	if got := NewBoundedCaller(&hangingCaller{}, 3*time.Second).timeout; got != 3*time.Second {
		t.Fatalf("explicit timeout = %v, want 3s", got)
	}
}

// hangingBackend is a full ContractBackend whose reads hang and whose
// transact surface records whether a deadline was imposed on it.
type hangingBackend struct {
	bind.ContractBackend // reads overridden below; unused methods panic via nil embed

	hangingCaller

	sendSawDeadline bool
}

func (h *hangingBackend) SendTransaction(ctx context.Context, _ *types.Transaction) error {
	_, h.sendSawDeadline = ctx.Deadline()
	return nil
}

func (h *hangingBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return h.hangingCaller.CodeAt(ctx, contract, blockNumber)
}

func (h *hangingBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return h.hangingCaller.CallContract(ctx, call, blockNumber)
}

// TestBoundedReadBackendBoundsReadsOnly proves the read surface gets the
// deadline while the transact surface passes through untouched (transaction
// bounding is the caller's contract — internal/ethtx owns receipt waits).
func TestBoundedReadBackendBoundsReadsOnly(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		inner := &hangingBackend{}
		b := NewBoundedReadBackend(inner, 2*time.Second)

		start := time.Now()
		if _, err := b.CallContract(t.Context(), ethereum.CallMsg{}, nil); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("CallContract err = %v, want DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != 2*time.Second {
			t.Fatalf("CallContract returned after %v, want exactly 2s", elapsed)
		}
		start = time.Now()
		if _, err := b.CodeAt(t.Context(), common.Address{0x01}, nil); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("CodeAt err = %v, want DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != 2*time.Second {
			t.Fatalf("CodeAt returned after %v, want exactly 2s", elapsed)
		}

		if err := b.SendTransaction(t.Context(), nil); err != nil {
			t.Fatalf("SendTransaction: %v", err)
		}
		if inner.sendSawDeadline {
			t.Fatal("SendTransaction must pass through without a wrapper-imposed deadline")
		}
	})
}

func TestNewBoundedReadBackendAppliesDefaultTimeout(t *testing.T) {
	t.Parallel()
	if got := NewBoundedReadBackend(&hangingBackend{}, 0).timeout; got != DefaultTimeout {
		t.Fatalf("zero timeout = %v, want DefaultTimeout %v", got, DefaultTimeout)
	}
}
