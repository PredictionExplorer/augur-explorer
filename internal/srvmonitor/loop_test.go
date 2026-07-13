package srvmonitor

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunLoopChecksImmediatelyAndRepeats(t *testing.T) {
	t.Parallel()
	var checks atomic.Int64
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	go func() {
		runLoop(ctx, time.Millisecond, func(context.Context) {
			checks.Add(1)
		})
		close(done)
	}()

	waitFor(t, "three check cycles", func() bool { return checks.Load() >= 3 })
	cancel()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("runLoop did not stop")
	}
}

func TestRunLoopCancelDuringWaitInterruptsPromptly(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	start := time.Now()
	go func() {
		// A one-hour interval: only prompt cancellation lets this return.
		runLoop(ctx, time.Hour, func(context.Context) {})
		close(done)
	}()

	time.Sleep(10 * time.Millisecond) // let the first check complete
	cancel()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("runLoop waited out the interval despite cancellation")
	}
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Fatalf("cancellation took %v", elapsed)
	}
}

func TestSleepCtx(t *testing.T) {
	t.Parallel()

	// Normal expiry.
	start := time.Now()
	sleepCtx(context.Background(), 5*time.Millisecond)
	if time.Since(start) < 5*time.Millisecond {
		t.Fatal("sleepCtx returned early")
	}

	// Cancellation interrupts a long sleep.
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Millisecond)
		cancel()
	}()
	start = time.Now()
	sleepCtx(ctx, time.Hour)
	if time.Since(start) > time.Second {
		t.Fatal("sleepCtx ignored cancellation")
	}
}

func TestSendErr(t *testing.T) {
	t.Parallel()

	// Delivered when there is room.
	ch := make(chan string, 1)
	sendErr(context.Background(), ch, "msg")
	if got := <-ch; got != "msg" {
		t.Fatalf("got %q", got)
	}

	// A full channel with a cancelled context does not block.
	full := make(chan string) // unbuffered, no reader
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	done := make(chan struct{})
	go func() {
		sendErr(ctx, full, "dropped")
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("sendErr blocked on a full channel despite cancellation")
	}
}
