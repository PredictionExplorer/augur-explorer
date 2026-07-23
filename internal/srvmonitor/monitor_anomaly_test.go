package srvmonitor

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func anomalyConfig() AnomalyConfig {
	return AnomalyConfig{User: "ops", Host: "websrv", RemoteFile: "/var/log/anomalies.log"}
}

func anomalyPayload(generatedAt time.Time, lines ...string) []byte {
	all := make([]string, 0, len(lines)+1)
	all = append(all, anomalyHeartbeat+strconv.FormatInt(generatedAt.Unix(), 10))
	all = append(all, lines...)
	return []byte(strings.Join(all, "\n") + "\n")
}

func TestAnomalyMonitorShowsLastLines(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, dir, testIntervals())
	now := time.Unix(1_720_000_000, 0)
	m.now = func() time.Time { return now }

	var gotArgs []string
	m.run = func(_ context.Context, name string, args ...string) ([]byte, error) {
		gotArgs = append([]string{name}, args...)
		// Simulate scp by writing the fetched file.
		content := anomalyPayload(now.Add(-5*time.Minute), "line1", "", "line2", "line3", "line4")
		return nil, os.WriteFile(m.localFile, content, 0o600)
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
	if header := disp.Row(40); !strings.Contains(header, "updated 5m ago") {
		t.Fatalf("header = %q, want freshness age", header)
	}
	if !m.hasHeartbeat || m.generatedAt.Unix() != now.Add(-5*time.Minute).Unix() || m.freshnessAge != 5*time.Minute || m.stale {
		t.Fatalf("freshness state = has:%v generated:%v age:%v stale:%v",
			m.hasHeartbeat, m.generatedAt, m.freshnessAge, m.stale)
	}
}

func TestAnomalyMonitorEmptyFileShowsNone(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, dir, testIntervals())
	now := time.Unix(1_720_000_000, 0)
	m.now = func() time.Time { return now }
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		return nil, os.WriteFile(m.localFile, anomalyPayload(now), 0o600)
	}

	disp := newFakeDisplay()
	m.check(context.Background(), disp, make(chan string, 10))

	if row := disp.Row(41); row != " None" {
		t.Fatalf("row = %q, want None", row)
	}
	if got := disp.FgAt(1, 41); got != ColorGreen {
		t.Fatalf("None color = %v, want green", got)
	}
	if header := disp.Row(40); !strings.Contains(header, "updated 0s ago") {
		t.Fatalf("header = %q", header)
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
	m.lines = []string{"old anomaly"}
	m.generatedAt = time.Unix(123, 0)
	m.freshnessAge = time.Hour
	m.hasHeartbeat = true
	m.stale = true
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
	if len(m.lines) != 0 || m.hasHeartbeat || m.stale || !m.generatedAt.IsZero() || m.freshnessAge != 0 {
		t.Fatalf("fetch failure leaked old state: %+v", m)
	}
	if header := disp.Row(40); strings.Contains(header, "STALE") || strings.Contains(header, "updated") {
		t.Fatalf("fetch failure was mislabeled as freshness state: %q", header)
	}
}

func TestAnomalyMonitorCannotRemovePreviousFetch(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, t.TempDir(), testIntervals())
	m.localFile = filepath.Join(t.TempDir(), "non-empty-directory")
	if err := os.Mkdir(m.localFile, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(m.localFile, "child"), []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	run := false
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		run = true
		return nil, nil
	}

	errCh := make(chan string, 1)
	m.check(context.Background(), newFakeDisplay(), errCh)
	if run {
		t.Fatal("scp ran after the previous destination could not be removed")
	}
	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "remove previous anomaly file") {
		t.Fatalf("errors = %v", msgs)
	}
}

func TestAnomalyMonitorReadFailure(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, t.TempDir(), testIntervals())
	m.lines = []string{"old anomaly"}
	m.generatedAt = time.Unix(123, 0)
	m.freshnessAge = time.Hour
	m.hasHeartbeat = true
	m.stale = true
	if err := os.WriteFile(m.localFile, anomalyPayload(time.Unix(123, 0), "old anomaly"), 0o600); err != nil {
		t.Fatal(err)
	}
	// scp "succeeds" but never writes the local file.
	m.run = func(context.Context, string, ...string) ([]byte, error) { return nil, nil }

	errCh := make(chan string, 10)
	disp := newFakeDisplay()
	m.check(context.Background(), disp, errCh)

	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "Anomaly read:") {
		t.Fatalf("errors = %v", msgs)
	}
	if len(m.lines) != 0 || m.hasHeartbeat || m.stale || !m.generatedAt.IsZero() || m.freshnessAge != 0 {
		t.Fatalf("read failure leaked old state: %+v", m)
	}
	if header := disp.Row(40); strings.Contains(header, "STALE") || strings.Contains(header, "updated") {
		t.Fatalf("read failure was mislabeled as freshness state: %q", header)
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

func TestAnomalyMonitorFreshnessBoundariesAndClockSkew(t *testing.T) {
	t.Parallel()
	now := time.Unix(1_720_000_000, 0)
	cases := []struct {
		name       string
		generated  time.Time
		wantAge    time.Duration
		wantStale  bool
		wantStatus string
	}{
		{
			name:       "fresh",
			generated:  now.Add(-9*time.Minute - 59*time.Second),
			wantAge:    9*time.Minute + 59*time.Second,
			wantStatus: "updated 9m ago",
		},
		{
			name:       "exact boundary remains fresh",
			generated:  now.Add(-10 * time.Minute),
			wantAge:    10 * time.Minute,
			wantStatus: "updated 10m ago",
		},
		{
			name:       "past boundary is stale",
			generated:  now.Add(-10*time.Minute - time.Second),
			wantAge:    10*time.Minute + time.Second,
			wantStale:  true,
			wantStatus: "STALE 10m",
		},
		{
			name:       "future heartbeat clamps clock skew",
			generated:  now.Add(5 * time.Minute),
			wantStatus: "updated 0s ago",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			cfg := anomalyConfig()
			cfg.StaleAfter = 10 * time.Minute
			logger, _ := newTestLogger()
			m := NewAnomalyMonitor(cfg, 40, logger, t.TempDir(), testIntervals())
			m.now = func() time.Time { return now }
			m.run = func(context.Context, string, ...string) ([]byte, error) {
				return nil, os.WriteFile(m.localFile, anomalyPayload(tc.generated, "boom"), 0o600)
			}

			disp := newFakeDisplay()
			errCh := make(chan string, 2)
			m.check(context.Background(), disp, errCh)

			if m.freshnessAge != tc.wantAge || m.stale != tc.wantStale {
				t.Fatalf("age/stale = %v/%v, want %v/%v",
					m.freshnessAge, m.stale, tc.wantAge, tc.wantStale)
			}
			header := disp.Row(40)
			if !strings.Contains(header, tc.wantStatus) {
				t.Fatalf("header = %q, want status %q", header, tc.wantStatus)
			}
			statusX := strings.Index(header, strings.Fields(tc.wantStatus)[0])
			wantColor := ColorGreen
			if tc.wantStale {
				wantColor = ColorRed
			}
			if statusX < 0 || disp.FgAt(statusX, 40) != wantColor {
				t.Fatalf("status color at %d = %v, want %v", statusX, disp.FgAt(statusX, 40), wantColor)
			}
			msgs := drain(errCh)
			if !tc.wantStale && len(msgs) != 0 {
				t.Fatalf("fresh heartbeat emitted errors: %v", msgs)
			}
			if tc.wantStale {
				if len(msgs) != 1 || !strings.Contains(msgs[0], "STALE") ||
					!strings.Contains(msgs[0], "within 10m") {
					t.Fatalf("stale errors = %v", msgs)
				}
			}
		})
	}
}

func TestAnomalyMonitorLegacyMarkersNeverBecomeStale(t *testing.T) {
	t.Parallel()
	for name, content := range map[string]string{
		"missing":   "legacy anomaly\n",
		"malformed": "#TS=broken\nlegacy anomaly\n",
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			cfg := anomalyConfig()
			cfg.StaleAfter = time.Second
			logger, _ := newTestLogger()
			m := NewAnomalyMonitor(cfg, 40, logger, t.TempDir(), testIntervals())
			m.now = func() time.Time { return time.Unix(1_720_000_000, 0) }
			m.run = func(context.Context, string, ...string) ([]byte, error) {
				return nil, os.WriteFile(m.localFile, []byte(content), 0o600)
			}

			disp := newFakeDisplay()
			errCh := make(chan string, 1)
			m.check(context.Background(), disp, errCh)

			if m.hasHeartbeat || m.stale || m.freshnessAge != 0 {
				t.Fatalf("legacy state = has:%v stale:%v age:%v", m.hasHeartbeat, m.stale, m.freshnessAge)
			}
			if msgs := drain(errCh); len(msgs) != 0 {
				t.Fatalf("legacy file emitted stale alarm: %v", msgs)
			}
			header := disp.Row(40)
			if strings.Contains(header, "updated") || strings.Contains(header, "STALE") {
				t.Fatalf("legacy header gained freshness label: %q", header)
			}
			if row := disp.Row(41); !strings.Contains(row, "legacy anomaly") {
				t.Fatalf("anomaly row = %q", row)
			}
		})
	}
}

func TestAnomalyMonitorRepeatedStaleAlarmStopsOnCancel(t *testing.T) {
	t.Parallel()
	now := time.Unix(1_720_000_000, 0)
	cfg := anomalyConfig()
	cfg.StaleAfter = 2 * time.Second
	logger, _ := newTestLogger()
	iv := testIntervals()
	iv.Anomaly = time.Millisecond
	m := NewAnomalyMonitor(cfg, 40, logger, t.TempDir(), iv)
	m.now = func() time.Time { return now }
	m.run = func(context.Context, string, ...string) ([]byte, error) {
		return nil, os.WriteFile(m.localFile, anomalyPayload(now.Add(-3*time.Second)), 0o600)
	}

	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan string, AlarmThreshold+10)
	done := make(chan struct{})
	go func() {
		m.Start(ctx, newFakeDisplay(), errCh)
		close(done)
	}()

	msgs := make([]string, 0, AlarmThreshold)
	for len(msgs) < AlarmThreshold {
		select {
		case msg := <-errCh:
			msgs = append(msgs, msg)
		case <-time.After(5 * time.Second):
			t.Fatal("timed out waiting for repeated stale alarms")
		}
	}
	cancel()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("stale monitor did not stop after cancellation")
	}

	for i := 1; i < len(msgs); i++ {
		if msgs[i] != msgs[0] {
			t.Fatalf("alarm key changed across cycles: %q != %q", msgs[i], msgs[0])
		}
	}
	if !strings.Contains(msgs[0], "within 2s") {
		t.Fatalf("alarm = %q", msgs[0])
	}

	// The stable message reaches the existing exact-string alarm threshold.
	tracker, runner, _, _ := newAlarmTest(true)
	for _, msg := range msgs {
		tracker.RecordAlarm(context.Background(), msg)
	}
	if names := runner.commandNames(); len(names) != 1 || names[0] != "termux-notification" {
		t.Fatalf("notification commands = %v, want one threshold notification", names)
	}
}

func TestAnomalyMonitorStaleThresholdDefaults(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, t.TempDir(), testIntervals())
	if m.staleAfter != DefaultAnomalyStaleAfter {
		t.Fatalf("staleAfter = %v, want %v", m.staleAfter, DefaultAnomalyStaleAfter)
	}
	cfg := anomalyConfig()
	cfg.StaleAfter = 7 * time.Minute
	m = NewAnomalyMonitor(cfg, 40, logger, t.TempDir(), testIntervals())
	if m.staleAfter != 7*time.Minute {
		t.Fatalf("custom staleAfter = %v", m.staleAfter)
	}
}

func TestAnomalyMonitorClockFallsBackToWallTime(t *testing.T) {
	t.Parallel()
	logger, _ := newTestLogger()
	m := NewAnomalyMonitor(anomalyConfig(), 40, logger, t.TempDir(), testIntervals())
	m.now = nil
	before := time.Now()
	got := m.clock()
	after := time.Now()
	if got.Before(before) || got.After(after) {
		t.Fatalf("fallback clock = %v, want between %v and %v", got, before, after)
	}
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

func TestHeartbeatAgeAndShortDuration(t *testing.T) {
	t.Parallel()
	now := time.Unix(100, 0)
	if got := heartbeatAge(now, now.Add(time.Second)); got != 0 {
		t.Fatalf("future heartbeat age = %v, want 0", got)
	}
	if got := heartbeatAge(now, now.Add(-time.Minute)); got != time.Minute {
		t.Fatalf("past heartbeat age = %v, want 1m", got)
	}
	for input, want := range map[time.Duration]string{
		-time.Second:                    "0s",
		45 * time.Second:                "45s",
		12*time.Minute + 30*time.Second: "12m",
		3*time.Hour + 5*time.Minute:     "3h5m",
	} {
		if got := shortDuration(input); got != want {
			t.Errorf("shortDuration(%v) = %q, want %q", input, got, want)
		}
	}
}

func TestParseAnomaliesHeartbeatCompatibility(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name          string
		input         string
		wantHeartbeat bool
		wantUnix      int64
		wantLines     []string
	}{
		{
			name:          "valid marker is stripped",
			input:         "#TS=1720000000\r\na\n\nb\r\nc\nd\n",
			wantHeartbeat: true,
			wantUnix:      1_720_000_000,
			wantLines:     []string{"b", "c", "d"},
		},
		{
			name:      "malformed marker is legacy",
			input:     "#TS=not-a-time\nfirst\nsecond\n",
			wantLines: []string{"first", "second"},
		},
		{
			name:      "missing marker is legacy",
			input:     "first\nsecond\n",
			wantLines: []string{"first", "second"},
		},
		{
			name:          "last valid marker wins and all markers are stripped",
			input:         "#TS=100\nfirst\n#TS=bad\n#TS=200\nsecond\n",
			wantHeartbeat: true,
			wantUnix:      200,
			wantLines:     []string{"first", "second"},
		},
		{
			name:      "indented lookalike remains an anomaly",
			input:     " #TS=123\n",
			wantLines: []string{" #TS=123"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := parseAnomalies(strings.NewReader(tc.input), 3)
			if err != nil {
				t.Fatal(err)
			}
			if got.hasHeartbeat != tc.wantHeartbeat {
				t.Fatalf("hasHeartbeat = %v, want %v", got.hasHeartbeat, tc.wantHeartbeat)
			}
			if got.hasHeartbeat && got.generatedAt.Unix() != tc.wantUnix {
				t.Fatalf("generatedAt = %v, want unix %d", got.generatedAt, tc.wantUnix)
			}
			if strings.Join(got.lines, "\x00") != strings.Join(tc.wantLines, "\x00") {
				t.Fatalf("lines = %q, want %q", got.lines, tc.wantLines)
			}
		})
	}
}

func TestReadAnomalies(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "f.log")
	if err := os.WriteFile(path, anomalyPayload(time.Unix(123, 0), "a", "b"), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := readAnomalies(path, 1)
	if err != nil {
		t.Fatal(err)
	}
	if !got.hasHeartbeat || got.generatedAt.Unix() != 123 || len(got.lines) != 1 || got.lines[0] != "b" {
		t.Fatalf("snapshot = %+v", got)
	}

	if _, err := readAnomalies(filepath.Join(t.TempDir(), "missing"), 3); err == nil {
		t.Fatal("missing file must error")
	}
}

func TestParseAnomaliesBoundsInput(t *testing.T) {
	t.Parallel()
	t.Run("line", func(t *testing.T) {
		t.Parallel()
		_, err := parseAnomaliesBounded(strings.NewReader("12345\n"), 3, 100, 4)
		if err == nil || !strings.Contains(err.Error(), "scan anomaly file") {
			t.Fatalf("error = %v, want line-size failure", err)
		}
	})
	t.Run("file", func(t *testing.T) {
		t.Parallel()
		_, err := parseAnomaliesBounded(strings.NewReader("a\nb\nc\n"), 3, 4, 100)
		if err == nil || !strings.Contains(err.Error(), "exceeds 4 bytes") {
			t.Fatalf("error = %v, want file-size failure", err)
		}
	})
	t.Run("zero retained rows still parses heartbeat", func(t *testing.T) {
		t.Parallel()
		got, err := parseAnomalies(strings.NewReader("#TS=123\na\n"), 0)
		if err != nil {
			t.Fatal(err)
		}
		if !got.hasHeartbeat || len(got.lines) != 0 {
			t.Fatalf("snapshot = %+v", got)
		}
	})
}

func FuzzParseAnomalies(f *testing.F) {
	f.Add([]byte("#TS=1720000000\n500 | GET /boom\n"))
	f.Add([]byte("#TS=bad\r\nlegacy\n"))
	f.Add([]byte("#TS=1\n#TS=2\n"))
	f.Add([]byte{0, '\n', 0xff, '\r', '\n'})
	f.Fuzz(func(t *testing.T, input []byte) {
		got, err := parseAnomalies(bytes.NewReader(input), AnomalyDisplayCount)
		if err != nil {
			return
		}
		if len(got.lines) > AnomalyDisplayCount {
			t.Fatalf("retained %d lines, limit %d", len(got.lines), AnomalyDisplayCount)
		}
		for _, line := range got.lines {
			if strings.HasPrefix(line, anomalyHeartbeat) {
				t.Fatalf("heartbeat marker leaked into anomaly rows: %q", line)
			}
			if strings.TrimSpace(line) == "" {
				t.Fatalf("blank anomaly row retained: %q", line)
			}
		}
	})
}
