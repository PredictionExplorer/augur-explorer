package srvmonitor

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"
)

// fakeRunner records commands and scripts per-command errors.
type fakeRunner struct {
	mu    sync.Mutex
	calls [][]string
	fail  map[string]error // command name -> error
}

func (r *fakeRunner) run(_ context.Context, name string, args ...string) ([]byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls = append(r.calls, append([]string{name}, args...))
	if err, ok := r.fail[name]; ok {
		return nil, err
	}
	return nil, nil
}

func (r *fakeRunner) commandNames() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	names := make([]string, 0, len(r.calls))
	for _, call := range r.calls {
		names = append(names, call[0])
	}
	return names
}

// alarmClock is a controllable clock.
type alarmClock struct {
	mu sync.Mutex
	t  time.Time
}

func (c *alarmClock) now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.t
}

func (c *alarmClock) advance(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.t = c.t.Add(d)
}

func newAlarmTest(enabled bool) (*AlarmTracker, *fakeRunner, *alarmClock, *captureLogger) {
	logger, sink := newTestLogger()
	tracker := NewAlarmTracker(enabled, logger)
	runner := &fakeRunner{fail: map[string]error{}}
	clock := &alarmClock{t: time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)}
	tracker.run = runner.run
	tracker.now = clock.now
	return tracker, runner, clock, sink
}

func TestAlarmTrackerThresholdAndCooldown(t *testing.T) {
	t.Parallel()
	tracker, runner, clock, _ := newAlarmTest(true)
	ctx := context.Background()

	// Four occurrences: below threshold, no notification.
	for i := 0; i < AlarmThreshold-1; i++ {
		tracker.RecordAlarm(ctx, "db down")
		clock.advance(time.Minute)
	}
	if got := len(runner.commandNames()); got != 0 {
		t.Fatalf("notifications before threshold = %d", got)
	}

	// Fifth occurrence triggers exactly one notification.
	tracker.RecordAlarm(ctx, "db down")
	if names := runner.commandNames(); len(names) != 1 || names[0] != "termux-notification" {
		t.Fatalf("commands = %v, want one termux-notification", names)
	}

	// Further occurrences within the cooldown stay silent.
	clock.advance(time.Minute)
	tracker.RecordAlarm(ctx, "db down")
	if got := len(runner.commandNames()); got != 1 {
		t.Fatalf("notifications within cooldown = %d, want 1", got)
	}

	// After the cooldown, the still-firing alarm notifies again.
	clock.advance(NotificationCooldown)
	// Re-arm the window: occurrences must again reach the threshold within
	// AlarmTimeWindow of "now" (older ones were pruned).
	for i := 0; i < AlarmThreshold; i++ {
		tracker.RecordAlarm(ctx, "db down")
	}
	if got := len(runner.commandNames()); got != 2 {
		t.Fatalf("notifications after cooldown = %d, want 2", got)
	}
}

func TestAlarmTrackerWindowPruning(t *testing.T) {
	t.Parallel()
	tracker, runner, clock, _ := newAlarmTest(true)
	ctx := context.Background()

	// Four occurrences, then a long pause: the window empties, so the next
	// occurrence counts as 1 and no notification fires.
	for i := 0; i < AlarmThreshold-1; i++ {
		tracker.RecordAlarm(ctx, "flaky")
	}
	clock.advance(AlarmTimeWindow + time.Minute)
	tracker.RecordAlarm(ctx, "flaky")

	if got := len(runner.commandNames()); got != 0 {
		t.Fatalf("notifications = %d, want 0 after window reset", got)
	}
}

func TestAlarmTrackerDisabledAndEmpty(t *testing.T) {
	t.Parallel()
	tracker, runner, _, sink := newAlarmTest(false)
	ctx := context.Background()

	for i := 0; i < AlarmThreshold*2; i++ {
		tracker.RecordAlarm(ctx, "ignored")
	}
	tracker.SendTestNotification(ctx, "test")
	if got := len(runner.commandNames()); got != 0 {
		t.Fatalf("disabled tracker ran %d commands", got)
	}
	if !strings.Contains(sink.String(), "mobile notifications disabled") {
		t.Fatalf("log = %q", sink.String())
	}

	// Empty messages are ignored even when enabled.
	enabled, runner2, _, _ := newAlarmTest(true)
	for i := 0; i < AlarmThreshold*2; i++ {
		enabled.RecordAlarm(ctx, "")
	}
	if got := len(runner2.commandNames()); got != 0 {
		t.Fatalf("empty messages ran %d commands", got)
	}
}

func TestAlarmTrackerFallbackChain(t *testing.T) {
	t.Parallel()

	t.Run("vibrate fallback", func(t *testing.T) {
		t.Parallel()
		tracker, runner, _, sink := newAlarmTest(true)
		runner.fail["termux-notification"] = errors.New("not installed")

		tracker.SendTestNotification(context.Background(), "hello")

		names := runner.commandNames()
		if len(names) != 2 || names[0] != "termux-notification" || names[1] != "termux-vibrate" {
			t.Fatalf("commands = %v", names)
		}
		if !strings.Contains(sink.String(), "Vibration alert sent successfully") {
			t.Fatalf("log = %q", sink.String())
		}
	})

	t.Run("terminal bell last resort", func(t *testing.T) {
		t.Parallel()
		tracker, runner, _, sink := newAlarmTest(true)
		runner.fail["termux-notification"] = errors.New("not installed")
		runner.fail["termux-vibrate"] = errors.New("also missing")

		tracker.SendTestNotification(context.Background(), "hello")

		if !strings.Contains(sink.String(), "ALERT (no notification sent): hello") {
			t.Fatalf("log = %q", sink.String())
		}
	})
}

func TestAlarmTrackerNotificationArguments(t *testing.T) {
	t.Parallel()
	tracker, runner, _, _ := newAlarmTest(true)
	tracker.SendTestNotification(context.Background(), "msg body")

	runner.mu.Lock()
	defer runner.mu.Unlock()
	call := runner.calls[0]
	want := []string{"termux-notification", "--title", "Server Monitor Alert", "--content", "msg body", "--priority", "high", "--sound"}
	if len(call) != len(want) {
		t.Fatalf("call = %v, want %v", call, want)
	}
	for i := range want {
		if call[i] != want[i] {
			t.Fatalf("arg %d = %q, want %q", i, call[i], want[i])
		}
	}
}

func TestAlarmTrackerCleanupOldData(t *testing.T) {
	t.Parallel()
	tracker, _, clock, _ := newAlarmTest(true)
	ctx := context.Background()

	for i := 0; i < AlarmThreshold; i++ {
		tracker.RecordAlarm(ctx, "old alarm") // fifth one notifies
	}
	tracker.RecordAlarm(ctx, "recent alarm")

	// After 25 hours everything is stale.
	clock.advance(25 * time.Hour)
	tracker.CleanupOldData()

	tracker.mutex.Lock()
	defer tracker.mutex.Unlock()
	if len(tracker.occurrences) != 0 {
		t.Fatalf("occurrences = %v, want pruned", tracker.occurrences)
	}
	if len(tracker.lastNotification) != 0 {
		t.Fatalf("lastNotification = %v, want pruned", tracker.lastNotification)
	}
}

func TestAlarmTrackerRunCleanupStopsOnCancel(t *testing.T) {
	t.Parallel()
	tracker, _, _, _ := newAlarmTest(true)

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		tracker.RunCleanup(ctx)
		close(done)
	}()

	cancel()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("RunCleanup did not stop on cancellation")
	}
}

func TestAlarmTrackerRunCleanupPrunesOnTick(t *testing.T) {
	t.Parallel()
	tracker, _, clock, _ := newAlarmTest(true)
	tracker.cleanupInterval = time.Millisecond
	ctx := context.Background()

	tracker.RecordAlarm(ctx, "stale")
	clock.advance(25 * time.Hour) // everything is now stale

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go tracker.RunCleanup(runCtx)

	waitFor(t, "cleanup tick", func() bool {
		tracker.mutex.Lock()
		defer tracker.mutex.Unlock()
		return len(tracker.occurrences) == 0
	})
}

func TestProductionCommandRunners(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// runStdout returns only stdout; stderr noise is dropped.
	out, err := runStdout(ctx, "sh", "-c", "echo visible; echo hidden 1>&2")
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.TrimSpace(string(out)); got != "visible" {
		t.Fatalf("runStdout = %q", got)
	}

	// runCombinedOutput interleaves both streams (failure text included).
	out, err = runCombinedOutput(ctx, "sh", "-c", "echo visible; echo hidden 1>&2; exit 3")
	if err == nil {
		t.Fatal("exit 3 must error")
	}
	combined := string(out)
	if !strings.Contains(combined, "visible") || !strings.Contains(combined, "hidden") {
		t.Fatalf("runCombinedOutput = %q", combined)
	}
}
