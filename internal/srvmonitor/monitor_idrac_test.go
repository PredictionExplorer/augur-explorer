package srvmonitor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"
)

// scriptedIDRACRunner serves per-host script results keyed on the IDRAC env
// value and records the environments it received.
type scriptedIDRACRunner struct {
	mu      sync.Mutex
	results map[string]struct {
		out string
		err error
	}
	envs  [][]string
	names []string
}

func newScriptedIDRACRunner() *scriptedIDRACRunner {
	return &scriptedIDRACRunner{results: make(map[string]struct {
		out string
		err error
	})}
}

func (r *scriptedIDRACRunner) set(host, out string, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.results[host] = struct {
		out string
		err error
	}{out, err}
}

func (r *scriptedIDRACRunner) run(_ context.Context, env []string, name string, _ ...string) ([]byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.envs = append(r.envs, env)
	r.names = append(r.names, name)
	for _, kv := range env {
		if host, ok := strings.CutPrefix(kv, "IDRAC="); ok {
			res, exists := r.results[host]
			if !exists {
				return nil, fmt.Errorf("unexpected host %q", host)
			}
			return []byte(res.out), res.err
		}
	}
	return nil, errors.New("no IDRAC in env")
}

func newIDRACMonitorForTest(cfgs []IDRACConfig, runner *scriptedIDRACRunner) *IDRACMonitor {
	logger := slog.New(slog.DiscardHandler)
	m := NewIDRACMonitor(cfgs, "/opt/monitor/idrac_check.sh", 40, logger, testIntervals())
	m.overlayInterval = time.Millisecond
	m.run = runner.run
	m.environ = func() []string { return []string{"PATH=/usr/bin"} }
	return m
}

func TestIDRACMonitorHealthy(t *testing.T) {
	t.Parallel()
	runner := newScriptedIDRACRunner()
	runner.set("10.0.0.1", "OK\n", nil)
	m := newIDRACMonitorForTest([]IDRACConfig{
		{Name: "prod r640", Host: "10.0.0.1", User: "mon", Pass: "pw"},
	}, runner)
	disp := newFakeDisplay()
	errCh := make(chan string, 10)

	m.check(context.Background(), disp, errCh)

	if got := disp.Row(41); !strings.Contains(got, "prod r640: OK") {
		t.Fatalf("row = %q", got)
	}
	if disp.FgAt(1, 41) != ColorGreen {
		t.Fatalf("fg = %v, want green", disp.FgAt(1, 41))
	}
	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}

	// Credentials travel via the environment, and the configured script is
	// what runs.
	runner.mu.Lock()
	defer runner.mu.Unlock()
	if runner.names[0] != "/opt/monitor/idrac_check.sh" {
		t.Fatalf("script = %q", runner.names[0])
	}
	joined := strings.Join(runner.envs[0], "\n")
	for _, want := range []string{"IDRAC=10.0.0.1", "IDRAC_USER=mon", "IDRAC_PASS=pw", "PATH=/usr/bin"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("env %q missing %q", joined, want)
		}
	}
}

func TestIDRACMonitorCrashAlarmsAndSection(t *testing.T) {
	t.Parallel()
	runner := newScriptedIDRACRunner()
	// The script exits 2 on a crash, so the runner returns an error too.
	runner.set("10.0.0.1",
		"CRASH 2026-07-30T04:44:41-05:00 A bus fatal error was detected on a component at slot 2.\n",
		errors.New("exit status 2"))
	m := newIDRACMonitorForTest([]IDRACConfig{
		{Host: "10.0.0.1", User: "mon", Pass: "pw"},
	}, runner)
	disp := newFakeDisplay()
	errCh := make(chan string, 10)

	m.check(context.Background(), disp, errCh)

	row := disp.Row(41)
	if !strings.Contains(row, "10.0.0.1: CRASHED") || !strings.Contains(row, "bus fatal error") {
		t.Fatalf("row = %q", row)
	}
	if disp.BgAt(1, 41) != ColorRed || disp.FgAt(1, 41) != ColorWhite {
		t.Fatalf("colors = fg %v bg %v, want white on red", disp.FgAt(1, 41), disp.BgAt(1, 41))
	}
	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "Server 10.0.0.1 CRASHED") {
		t.Fatalf("alarms = %v", msgs)
	}
}

func TestIDRACMonitorCheckFailure(t *testing.T) {
	t.Parallel()
	runner := newScriptedIDRACRunner()
	runner.set("10.0.0.1", "curl: (7) Failed to connect\n", errors.New("exit status 1"))
	m := newIDRACMonitorForTest([]IDRACConfig{
		{Host: "10.0.0.1", User: "mon", Pass: "pw"},
	}, runner)
	disp := newFakeDisplay()
	errCh := make(chan string, 10)

	m.check(context.Background(), disp, errCh)

	if row := disp.Row(41); !strings.Contains(row, "ERROR: curl: (7)") {
		t.Fatalf("row = %q", row)
	}
	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "iDRAC 10.0.0.1 check failed") {
		t.Fatalf("alarms = %v", msgs)
	}
}

func TestIDRACMonitorOverlayLifecycle(t *testing.T) {
	t.Parallel()
	runner := newScriptedIDRACRunner()
	runner.set("10.0.0.1", "CRASH 2026-07-30T04:44:41-05:00 A bus fatal error was detected.\n",
		errors.New("exit status 2"))
	m := newIDRACMonitorForTest([]IDRACConfig{
		{Host: "10.0.0.1", User: "mon", Pass: "pw"},
	}, runner)
	disp := newFakeDisplay()
	errCh := make(chan string, 100)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go m.Start(ctx, disp, errCh)

	// The overlay painter must put a big white-on-red banner mid-screen.
	width, height := disp.Size()
	waitFor(t, "crash overlay", func() bool {
		return strings.Contains(disp.Row(height/2), "SERVER  10.0.0.1  CRASHED") ||
			strings.Contains(disp.Row(height/2+1), "SERVER  10.0.0.1  CRASHED")
	})
	if disp.BgAt(width/2, height/2) != ColorRed {
		t.Fatalf("center bg = %v, want red", disp.BgAt(width/2, height/2))
	}

	// Recovery: the next cycles read OK, the overlay wipes the screen once.
	before := disp.Cleared()
	runner.set("10.0.0.1", "OK\n", nil)
	waitFor(t, "overlay cleared after recovery", func() bool {
		return disp.Cleared() > before
	})
}

func TestIDRACMonitorName(t *testing.T) {
	t.Parallel()
	m := newIDRACMonitorForTest(nil, newScriptedIDRACRunner())
	if m.Name() != "iDRAC Crash Monitor" {
		t.Fatalf("Name() = %q", m.Name())
	}
}

func TestDefaultIDRACScriptPath(t *testing.T) {
	t.Parallel()
	path := DefaultIDRACScriptPath()
	if !strings.HasSuffix(path, "idrac_check.sh") {
		t.Fatalf("path = %q", path)
	}
	// An empty configured path selects the default next to the binary.
	m := NewIDRACMonitor(nil, "", 0, slog.New(slog.DiscardHandler), testIntervals())
	if m.scriptPath != path {
		t.Fatalf("scriptPath = %q, want %q", m.scriptPath, path)
	}
}

func TestFirstNonEmptyLine(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"":                     "",
		"\n\n":                 "",
		"OK\n":                 "OK",
		"\n  CRASH x\nOK\n":    "CRASH x",
		"  padded  \nsecond\n": "padded",
		"no trailing newline":  "no trailing newline",
		"\r\nwindows line\r\n": "windows line",
	}
	for in, want := range cases {
		if got := firstNonEmptyLine(in); got != want {
			t.Errorf("firstNonEmptyLine(%q) = %q, want %q", in, got, want)
		}
	}
}
