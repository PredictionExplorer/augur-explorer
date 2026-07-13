package srvmonitor

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func anomalyConfig() AnomalyConfig {
	return AnomalyConfig{User: "ops", Host: "websrv", RemoteFile: "/var/log/anomalies.log"}
}

func TestAnomalyMonitorShowsLastLines(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, dir, testIntervals())

	var gotArgs []string
	m.run = func(_ context.Context, name string, args ...string) ([]byte, error) {
		gotArgs = append([]string{name}, args...)
		// Simulate scp by writing the fetched file.
		content := "line1\n\nline2\r\nline3\nline4\n"
		return nil, os.WriteFile(m.localFile, []byte(content), 0o600)
	}

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	// scp target: user@host:remote -> local file under dir.
	if want := "ops@websrv:/var/log/anomalies.log"; gotArgs[len(gotArgs)-2] != want {
		t.Fatalf("scp source = %q, want %q", gotArgs[len(gotArgs)-2], want)
	}
	if gotArgs[len(gotArgs)-1] != filepath.Join(dir, "srvmonitor_anomalies.log") {
		t.Fatalf("scp destination = %q", gotArgs[len(gotArgs)-1])
	}

	// Only the last 3 non-empty lines render, all in red.
	if row := disp.Row(41); !strings.Contains(row, "line2") {
		t.Fatalf("row 41 = %q", row)
	}
	if row := disp.Row(42); !strings.Contains(row, "line3") {
		t.Fatalf("row 42 = %q", row)
	}
	if row := disp.Row(43); !strings.Contains(row, "line4") {
		t.Fatalf("row 43 = %q", row)
	}
	if got := disp.FgAt(1, 43); got != ColorRed {
		t.Fatalf("anomaly color = %v, want red", got)
	}
	if header := disp.Row(40); !strings.Contains(header, "WebSrv Anomalies (last 3)") {
		t.Fatalf("header = %q", header)
	}
}

func TestAnomalyMonitorEmptyFileShowsNone(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, dir, testIntervals())
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		return nil, os.WriteFile(m.localFile, []byte("\n\n"), 0o600)
	}

	disp := newFakeDisplay()
	m.check(context.Background(), disp, make(chan string, 10))

	if row := disp.Row(41); row != " None" {
		t.Fatalf("row = %q, want None", row)
	}
	if got := disp.FgAt(1, 41); got != ColorGreen {
		t.Fatalf("None color = %v, want green", got)
	}
}

func TestAnomalyMonitorCustomTitle(t *testing.T) {
	t.Parallel()
	cfg := anomalyConfig()
	cfg.Title = "Prod Anomalies"
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(cfg, 40, logger, t.TempDir(), testIntervals())
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		return nil, os.WriteFile(m.localFile, nil, 0o600)
	}

	disp := newFakeDisplay()
	m.check(context.Background(), disp, make(chan string, 10))

	if header := disp.Row(40); !strings.Contains(header, "Prod Anomalies") {
		t.Fatalf("header = %q", header)
	}
}

func TestAnomalyMonitorFetchFailure(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, t.TempDir(), testIntervals())
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		return []byte("Connection timed out\n"), errors.New("exit status 1")
	}

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "Anomaly fetch: exit status 1: Connection timed out") {
		t.Fatalf("errors = %v", msgs)
	}
	if row := disp.Row(41); !strings.Contains(row, "ERROR: exit status 1") {
		t.Fatalf("row = %q", row)
	}
}

func TestAnomalyMonitorReadFailure(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, t.TempDir(), testIntervals())
	// scp "succeeds" but never writes the local file.
	m.run = func(context.Context, string, ...string) ([]byte, error) { return nil, nil }

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "Anomaly read:") {
		t.Fatalf("errors = %v", msgs)
	}
}

func TestAnomalyMonitorDefaultLocalDir(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, "", testIntervals())
	if m.localFile != filepath.Join(os.TempDir(), "srvmonitor_anomalies.log") {
		t.Fatalf("localFile = %q", m.localFile)
	}
}

func TestAnomalyMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, t.TempDir(), testIntervals())
	cycles := make(chan struct{}, 100)
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		select {
		case cycles <- struct{}{}:
		default:
		}
		return nil, os.WriteFile(m.localFile, []byte("a\n"), 0o600)
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		m.Start(ctx, newFakeDisplay(), make(chan string, 100))
		close(done)
	}()

	<-cycles
	<-cycles
	cancel()
	waitFor(t, "loop exit", func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	})
}

func TestAnomalyMonitorLongLinesTruncated(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, t.TempDir(), testIntervals())
	long := strings.Repeat("x", 150)
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		return nil, os.WriteFile(m.localFile, []byte(long+"\n"), 0o600)
	}

	disp := newFakeDisplay()
	m.check(context.Background(), disp, make(chan string, 10))

	row := disp.Row(41)
	if len(row) > anomalyMaxLineWidth+1 { // +1 for the X=1 offset
		t.Fatalf("row length = %d, want truncated to %d", len(row), anomalyMaxLineWidth)
	}
	if !strings.HasSuffix(row, "...") {
		t.Fatalf("row = %q, want ellipsis suffix", row)
	}
}

func TestTruncateLine(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in   string
		max  int
		want string
	}{
		{"short", 10, "short"},
		{"exactly10!", 10, "exactly10!"},
		{"this is far too long", 10, "this is..."},
		{"abcdef", 3, "abc"},
		{"abcdef", 2, "ab"},
	}
	for _, tc := range cases {
		if got := truncateLine(tc.in, tc.max); got != tc.want {
			t.Fatalf("truncateLine(%q, %d) = %q, want %q", tc.in, tc.max, got, tc.want)
		}
	}
}

func TestReadLastLines(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "f.log")
	if err := os.WriteFile(path, []byte("a\n\nb\r\nc\nd\n \n"), 0o600); err != nil {
		t.Fatal(err)
	}
	lines, err := readLastLines(path, 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 || lines[0] != "b" || lines[1] != "c" || lines[2] != "d" {
		t.Fatalf("lines = %v", lines)
	}

	if _, err := readLastLines(filepath.Join(t.TempDir(), "missing"), 3); err == nil {
		t.Fatal("missing file must error")
	}
}
