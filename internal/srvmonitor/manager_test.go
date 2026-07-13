package srvmonitor

import (
	"context"
	"strings"
	"testing"
)

// stubMonitor is a Monitor pushing scripted errors when started.
type stubMonitor struct {
	name   string
	errors []string
}

func (s *stubMonitor) Name() string { return s.name }

func (s *stubMonitor) Start(ctx context.Context, _ Display, errorChan chan<- string) {
	for _, msg := range s.errors {
		sendErr(ctx, errorChan, msg)
	}
	<-ctx.Done()
}

func TestManagerHandlesErrors(t *testing.T) {
	t.Parallel()
	logger, sink := newTestLogger()
	tracker := NewAlarmTracker(true, logger)
	runner := &fakeRunner{}
	tracker.run = runner.run
	disp := newFakeDisplay()

	mgr := NewManager(disp, logger, tracker, 5, 40)
	mgr.Register(&stubMonitor{name: "stub", errors: []string{"", "first failure", "second failure", "third failure"}})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mgr.Start(ctx)

	// The first two non-empty errors land on the display's error area.
	waitFor(t, "both error lines", func() bool {
		return strings.Contains(disp.Row(40), "first failure") &&
			strings.Contains(disp.Row(41), "second failure")
	})
	// The third error is logged and alarm-tracked but not displayed.
	waitFor(t, "third failure logged", func() bool {
		return strings.Contains(sink.String(), "third failure")
	})
	if row := disp.Row(42); row != "" {
		t.Fatalf("row 42 = %q, want empty", row)
	}

	// All three errors were recorded with the alarm tracker.
	waitFor(t, "alarm occurrences", func() bool {
		tracker.mutex.Lock()
		defer tracker.mutex.Unlock()
		return len(tracker.occurrences) == 3
	})
}

func TestManagerMonitorNames(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	mgr := NewManager(newFakeDisplay(), logger, NewAlarmTracker(false, logger), 0, 0)
	mgr.Register(&stubMonitor{name: "a"})
	mgr.Register(&stubMonitor{name: "b"})

	names := mgr.MonitorNames()
	if len(names) != 2 || names[0] != "a" || names[1] != "b" {
		t.Fatalf("names = %v", names)
	}
}

func TestManagerSendTestNotification(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	tracker := NewAlarmTracker(true, logger)
	runner := &fakeRunner{}
	tracker.run = runner.run

	mgr := NewManager(newFakeDisplay(), logger, tracker, 0, 0)
	mgr.SendTestNotification(context.Background())

	names := runner.commandNames()
	if len(names) != 1 || names[0] != "termux-notification" {
		t.Fatalf("commands = %v", names)
	}
}
