package main

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

// watermarkFunc adapts a function to the blockWatermark seam.
type watermarkFunc func(ctx context.Context) (int64, error)

func (f watermarkFunc) LastBlock(ctx context.Context) (int64, error) { return f(ctx) }

var testImpl = ethcommon.HexToAddress("0x7700000000000000000000000000000000000077")

// discardLogger swallows the backfills' diagnostics.
var discardLogger = slog.New(slog.NewTextHandler(io.Discard, nil))

// shrinkCommitWait makes the commit-wait deadline elapse immediately and
// restores the production bound afterwards.
func shrinkCommitWait(t *testing.T) {
	t.Helper()
	old := implementationCommitWait
	implementationCommitWait = 0
	t.Cleanup(func() { implementationCommitWait = old })
}

// The backfills' guard arms need neither handlers nor an engine: they must
// return before touching them (nil values would panic otherwise).

func TestBackfillImplementationStopsOnCancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	behind := watermarkFunc(func(context.Context) (int64, error) { return 0, nil })

	done := make(chan struct{})
	go func() {
		backfillImplementationAfterCommit(ctx, discardLogger, behind, nil, nil, nil, testImpl, 100)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("cancelled backfill did not return")
	}
}

func TestBackfillImplementationSkipsWhenCommitWaitExpires(t *testing.T) {
	shrinkCommitWait(t)
	behind := watermarkFunc(func(context.Context) (int64, error) { return 0, nil })
	// Deadline of now: the first behind-watermark read trips the skip arm.
	backfillImplementationAfterCommit(context.Background(), discardLogger, behind, nil, nil, nil, testImpl, 100)
}

func TestBackfillImplementationIgnoresNegativeBlock(t *testing.T) {
	committed := watermarkFunc(func(context.Context) (int64, error) { return 10, nil })
	// blockNum -1 is the "unknown block" sentinel: the loop exits at once
	// (10 >= -1) and the negative guard returns before any backfill.
	backfillImplementationAfterCommit(context.Background(), discardLogger, committed, nil, nil, nil, testImpl, -1)
}

func TestBackfillPrizesWalletStopsOnCancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	behind := watermarkFunc(func(context.Context) (int64, error) { return 0, nil })

	done := make(chan struct{})
	go func() {
		backfillPrizesWalletAfterCommit(ctx, discardLogger, behind, nil, nil, nil, testImpl, 100)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("cancelled backfill did not return")
	}
}

func TestBackfillPrizesWalletSkipsWhenCommitWaitExpires(t *testing.T) {
	shrinkCommitWait(t)
	behind := watermarkFunc(func(context.Context) (int64, error) { return 0, nil })
	backfillPrizesWalletAfterCommit(context.Background(), discardLogger, behind, nil, nil, nil, testImpl, 100)
}

func TestBackfillPrizesWalletIgnoresNegativeBlock(t *testing.T) {
	// The negative guard runs before the first watermark read.
	failing := watermarkFunc(func(context.Context) (int64, error) {
		t.Fatal("watermark read despite negative block")
		return 0, nil
	})
	backfillPrizesWalletAfterCommit(context.Background(), discardLogger, failing, nil, nil, nil, testImpl, -1)
}
