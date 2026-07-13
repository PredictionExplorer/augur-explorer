package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/srvmonitor"
	"github.com/nsf/termbox-go"
)

// recordingDisplay implements srvmonitor.Display, counting operations.
type recordingDisplay struct {
	mu      sync.Mutex
	clears  int
	flushes int
	texts   []string
}

func (d *recordingDisplay) Init() error  { return nil }
func (d *recordingDisplay) Close() error { return nil }
func (d *recordingDisplay) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.clears++
}

func (d *recordingDisplay) DrawText(_ srvmonitor.Position, text string, _, _ srvmonitor.Color) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.texts = append(d.texts, text)
}

func (d *recordingDisplay) Flush() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.flushes++
}

func (d *recordingDisplay) Size() (int, int) { return 80, 24 }

type sink struct {
	mu  sync.Mutex
	buf strings.Builder
}

func (s *sink) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.Write(p)
}

func (s *sink) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.String()
}

type fakeNotifier struct {
	mu    sync.Mutex
	calls int
}

func (f *fakeNotifier) SendTestNotification(context.Context) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls++
}

func (f *fakeNotifier) count() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.calls
}

func waitCond(t *testing.T, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(2 * time.Millisecond)
	}
	t.Fatalf("timed out waiting for %s", what)
}

func TestRunEventLoopExitKeys(t *testing.T) {
	t.Parallel()
	for name, ev := range map[string]termbox.Event{
		"q key":     {Type: termbox.EventKey, Ch: 'q'},
		"ctrl-c":    {Type: termbox.EventKey, Key: termbox.KeyCtrlC},
		"interrupt": {Type: termbox.EventInterrupt},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			s := &sink{}
			logger := log.New(s, "", 0)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			events := make(chan termbox.Event, 1)
			events <- ev

			done := make(chan struct{})
			go func() {
				runEventLoop(ctx, events, &recordingDisplay{}, logger, cancel)
				close(done)
			}()

			select {
			case <-done:
			case <-time.After(5 * time.Second):
				t.Fatal("event loop did not exit")
			}
			// The exit path must have cancelled the shared context.
			if ctx.Err() == nil {
				t.Fatal("context not cancelled")
			}
		})
	}
}

func TestRunEventLoopResizeRepaints(t *testing.T) {
	t.Parallel()
	s := &sink{}
	logger := log.New(s, "", 0)
	disp := &recordingDisplay{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	events := make(chan termbox.Event)

	go runEventLoop(ctx, events, disp, logger, cancel)

	events <- termbox.Event{Type: termbox.EventResize, Width: 80, Height: 24}
	waitCond(t, "resize repaint", func() bool {
		disp.mu.Lock()
		defer disp.mu.Unlock()
		return disp.clears == 2 && disp.flushes == 2 && len(disp.texts) == 1
	})
	disp.mu.Lock()
	text := disp.texts[0]
	disp.mu.Unlock()
	if text != "Refreshing display..." {
		t.Fatalf("text = %q", text)
	}

	// An error event only logs.
	events <- termbox.Event{Type: termbox.EventError, Err: errors.New("tty gone")}
	waitCond(t, "error logged", func() bool {
		return strings.Contains(s.String(), "Termbox error event: tty gone")
	})

	cancel()
	waitCond(t, "loop exit logged", func() bool {
		return strings.Contains(s.String(), "Context cancelled, exiting event loop")
	})
}

func TestRunEventLoopNonExitKeyIgnored(t *testing.T) {
	t.Parallel()
	s := &sink{}
	logger := log.New(s, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	events := make(chan termbox.Event)

	done := make(chan struct{})
	go func() {
		runEventLoop(ctx, events, &recordingDisplay{}, logger, cancel)
		close(done)
	}()

	events <- termbox.Event{Type: termbox.EventKey, Ch: 'x'}
	select {
	case <-done:
		t.Fatal("loop exited on a non-exit key")
	case <-time.After(20 * time.Millisecond):
	}

	cancel()
	<-done
}

func TestHandleSignals(t *testing.T) {
	t.Parallel()
	s := &sink{}
	logger := log.New(s, "", 0)
	mgr := &fakeNotifier{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal)

	done := make(chan struct{})
	go func() {
		handleSignals(ctx, sigs, mgr, logger, cancel)
		close(done)
	}()

	sigs <- syscall.SIGWINCH
	waitCond(t, "SIGWINCH logged", func() bool {
		return strings.Contains(s.String(), "Window resize signal received")
	})
	if mgr.count() != 0 {
		t.Fatal("SIGWINCH must not notify")
	}

	sigs <- syscall.SIGUSR1
	waitCond(t, "test notification", func() bool { return mgr.count() == 1 })

	sigs <- syscall.SIGTERM
	waitCond(t, "termination", func() bool { return ctx.Err() != nil })

	// The handler itself exits once the context is cancelled.
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("signal handler did not exit")
	}
	if !strings.Contains(s.String(), "Termination signal received, exiting") {
		t.Fatalf("log = %q", s.String())
	}
}

func TestSetup(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		tmp := t.TempDir()
		env := map[string]string{
			"TMPDIR":          tmp,
			"RPC0_NAME":       "node0",
			"RPC0_URL":        "http://rpc0.example",
			"DB_L1_NAME_SRV1": "l1",
			"DB_L1_HOST_SRV1": "db1.example:5432",
		}
		s, err := setup(func(k string) string { return env[k] })
		if err != nil {
			t.Fatal(err)
		}
		if s.tmpDir != tmp {
			t.Fatalf("tmpDir = %q", s.tmpDir)
		}
		if s.logPath != tmp+"/srvmonitor.log" || s.oldPath != tmp+"/srvmonitor-old.log" {
			t.Fatalf("paths = %q / %q", s.logPath, s.oldPath)
		}
		if len(s.cfg.RPCNodes) != 1 {
			t.Fatalf("cfg = %+v", s.cfg)
		}
		// The startup banner and config summary landed in the log file.
		data, err := os.ReadFile(s.logPath)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), "=== Server Monitor Starting ===") ||
			!strings.Contains(string(data), "1 RPC nodes") {
			t.Fatalf("log = %q", data)
		}
	})

	t.Run("defaults TMPDIR to /tmp", func(t *testing.T) {
		env := map[string]string{
			"RPC0_NAME":       "node0",
			"RPC0_URL":        "http://rpc0.example",
			"DB_L1_NAME_SRV1": "l1",
			"DB_L1_HOST_SRV1": "db1.example:5432",
		}
		s, err := setup(func(k string) string { return env[k] })
		if err != nil {
			t.Fatal(err)
		}
		if s.tmpDir != "/tmp" {
			t.Fatalf("tmpDir = %q", s.tmpDir)
		}
	})

	t.Run("unwritable log path", func(t *testing.T) {
		_, err := setup(func(k string) string {
			if k == "TMPDIR" {
				return "/dev/null/nope"
			}
			return ""
		})
		if err == nil || !strings.Contains(err.Error(), "error opening log file") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("configuration error", func(t *testing.T) {
		env := map[string]string{"TMPDIR": t.TempDir()}
		_, err := setup(func(k string) string { return env[k] })
		if err == nil || !strings.Contains(err.Error(), "configuration error") {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestRunEventLoopResizeClampsOnTinyScreen(t *testing.T) {
	t.Parallel()
	s := &sink{}
	logger := log.New(s, "", 0)
	disp := &tinyDisplay{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	events := make(chan termbox.Event)

	go runEventLoop(ctx, events, disp, logger, cancel)

	// A degenerate screen size makes both centering coordinates negative;
	// they clamp to zero instead of drawing off-screen.
	events <- termbox.Event{Type: termbox.EventResize, Width: 1, Height: 1}
	waitCond(t, "clamped draw", func() bool {
		disp.mu.Lock()
		defer disp.mu.Unlock()
		return len(disp.positions) == 1
	})
	disp.mu.Lock()
	pos := disp.positions[0]
	disp.mu.Unlock()
	if pos.X != 0 || pos.Y != 0 {
		t.Fatalf("draw position = %+v, want clamped to origin", pos)
	}
}

// tinyDisplay reports a degenerate screen size (as termbox does before the
// terminal is measured) and records draw positions.
type tinyDisplay struct {
	mu        sync.Mutex
	positions []srvmonitor.Position
}

func (d *tinyDisplay) Init() error  { return nil }
func (d *tinyDisplay) Close() error { return nil }
func (d *tinyDisplay) Clear()       {}
func (d *tinyDisplay) DrawText(pos srvmonitor.Position, _ string, _, _ srvmonitor.Color) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.positions = append(d.positions, pos)
}
func (d *tinyDisplay) Flush()           {}
func (d *tinyDisplay) Size() (int, int) { return 1, -1 }

func TestLogConfigSummary(t *testing.T) {
	t.Parallel()
	s := &sink{}
	logger := log.New(s, "", 0)

	cfg := &srvmonitor.Config{
		RPCNodes:  []srvmonitor.RPCConfig{{Name: "n"}},
		Layer1DBs: []srvmonitor.DatabaseConfig{{Name: "db"}},
		Anomaly:   srvmonitor.AnomalyConfig{User: "u", Host: "h", RemoteFile: "/f"},
	}
	logConfigSummary(logger, cfg)
	out := s.String()
	if !strings.Contains(out, "1 RPC nodes") || !strings.Contains(out, "Anomaly monitor: u@h:/f") {
		t.Fatalf("log = %q", out)
	}

	// Disabled anomaly renders the hint.
	s2 := &sink{}
	logConfigSummary(log.New(s2, "", 0), &srvmonitor.Config{})
	if !strings.Contains(s2.String(), "Anomaly monitor: DISABLED") {
		t.Fatalf("log = %q", s2.String())
	}
}
