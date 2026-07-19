package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nsf/termbox-go"

	"github.com/PredictionExplorer/augur-explorer/internal/srvmonitor/termboxui"
)

// fakeUI embeds the recordingDisplay fake and adds a scriptable Init so
// runMain's display lifecycle runs without a real terminal.
type fakeUI struct {
	recordingDisplay

	initErr   error
	initPanic bool
}

func (d *fakeUI) Init() error {
	if d.initPanic {
		panic("scripted display panic")
	}
	return d.initErr
}

// minimalMonitorEnv satisfies Config.Validate (one RPC node, one Layer1
// database) with endpoints that fail fast, and points TMPDIR at a scratch
// directory for the monitor log.
func minimalMonitorEnv(t *testing.T) {
	t.Helper()
	t.Setenv("TMPDIR", t.TempDir())
	t.Setenv("RPC0_NAME", "test-rpc")
	t.Setenv("RPC0_URL", "http://127.0.0.1:1")
	t.Setenv("DB_L1_NAME_SRV1", "test-db")
	t.Setenv("DB_L1_HOST_SRV1", "127.0.0.1")
}

// stubUI routes runMain's display construction to the fake and restores the
// production wiring afterwards.
func stubUI(t *testing.T, ui *fakeUI) {
	t.Helper()
	newDisplay = func() display { return ui }
	pollEvent = func() termbox.Event {
		return termbox.Event{Type: termbox.EventKey, Ch: 'q'}
	}
	t.Cleanup(func() {
		newDisplay = func() display { return termboxui.New() }
		pollEvent = termbox.PollEvent
	})
}

// TestRunMainSetupError pins the exit-path contract for the pre-terminal
// phase: a failing setup (unwritable TMPDIR) surfaces as an error before the
// display is ever touched.
func TestRunMainSetupError(t *testing.T) {
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "does-not-exist"))
	err := runMain()
	if err == nil || !strings.Contains(err.Error(), "error opening log file") {
		t.Fatalf("runMain() = %v, want log-file open error", err)
	}
}

// TestRunMainDisplayInitError proves a display that cannot initialize turns
// into an error return (after which the deferred log rotation still runs).
func TestRunMainDisplayInitError(t *testing.T) {
	minimalMonitorEnv(t)
	stubUI(t, &fakeUI{initErr: errors.New("no tty")})

	err := runMain()
	if err == nil || !strings.Contains(err.Error(), "failed to initialize display") {
		t.Fatalf("runMain() = %v, want display init error", err)
	}
}

// TestRunMainRecoversPanic proves the deferred recovery converts a panic
// anywhere in the terminal phase into an ordinary error for main.
func TestRunMainRecoversPanic(t *testing.T) {
	minimalMonitorEnv(t)
	stubUI(t, &fakeUI{initPanic: true})

	err := runMain()
	if err == nil || !strings.Contains(err.Error(), "panic: scripted display panic") {
		t.Fatalf("runMain() = %v, want recovered panic error", err)
	}
}

// TestRunMainFullDashboardRun drives the whole wiring — setup, display init,
// event goroutine, signal registration, dashboard — over the fake display;
// the scripted 'q' key exits the event loop and runMain returns nil. (Draw
// content is asserted by the runDashboard/monitor tests; the 'q' exit races
// the first monitor paint by design.)
func TestRunMainFullDashboardRun(t *testing.T) {
	minimalMonitorEnv(t)
	ui := &fakeUI{}
	stubUI(t, ui)

	if err := runMain(); err != nil {
		t.Fatalf("runMain() = %v, want nil after 'q' exit", err)
	}
}

// TestMainExitsNonZeroOnError drives the real entrypoint through its failure
// arm via the stubbed exit seam.
func TestMainExitsNonZeroOnError(t *testing.T) {
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "does-not-exist"))
	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs; osExit = os.Exit })
	os.Args = []string{"srvmonitor"}
	code := 0
	osExit = func(c int) { code = c }

	main()

	if code != 1 {
		t.Fatalf("main() exit code = %d, want 1", code)
	}
}
