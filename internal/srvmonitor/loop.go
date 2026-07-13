package srvmonitor

import (
	"context"
	"time"
)

// runLoop performs check immediately and then once per interval until ctx is
// cancelled. Unlike the legacy per-monitor loops, cancellation interrupts the
// wait instead of letting a full interval (up to an hour for SSL checks)
// elapse first.
func runLoop(ctx context.Context, interval time.Duration, check func(ctx context.Context)) {
	timer := time.NewTimer(0)
	defer timer.Stop()
	<-timer.C
	for {
		check(ctx)
		timer.Reset(interval)
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
		}
	}
}

// sleepCtx pauses for d or until ctx is cancelled, whichever comes first.
// Monitors use it for in-check waits (e.g. between two block-number reads).
func sleepCtx(ctx context.Context, d time.Duration) {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
	case <-timer.C:
	}
}

// sendErr delivers a monitor error to the manager's channel without blocking
// past cancellation: the legacy unconditional send could wedge a monitor
// goroutine forever once the error handler had exited.
func sendErr(ctx context.Context, ch chan<- string, msg string) {
	select {
	case ch <- msg:
	case <-ctx.Done():
	}
}
