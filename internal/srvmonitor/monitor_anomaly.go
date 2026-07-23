package srvmonitor

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	// AnomalyDisplayCount is the number of most-recent anomalies to show.
	AnomalyDisplayCount = 3
	// anomalyMaxLineWidth truncates long anomaly lines to fit the terminal.
	anomalyMaxLineWidth = 100
	// anomalyMaxLineBytes and anomalyMaxFileBytes bound untrusted remote input.
	anomalyMaxLineBytes = 1 << 20
	anomalyMaxFileBytes = 64 << 20
	anomalyHeartbeat    = "#TS="
)

// AnomalyMonitor fetches the websrv anomalies file (produced by the
// loganomaly tool on the production host) via scp and displays the most
// recent entries.
type AnomalyMonitor struct {
	config     AnomalyConfig
	position   Position
	localFile  string
	logger     *slog.Logger
	interval   time.Duration
	staleAfter time.Duration
	run        CommandRunner
	now        func() time.Time

	lines        []string
	errStr       string
	generatedAt  time.Time
	freshnessAge time.Duration
	hasHeartbeat bool
	stale        bool
}

// NewAnomalyMonitor creates a new anomaly monitor rooted at baseY. The
// fetched file is stored under localDir (empty selects the system temp
// directory).
func NewAnomalyMonitor(cfg AnomalyConfig, baseY int, logger *slog.Logger, localDir string, iv Intervals) *AnomalyMonitor {
	if localDir == "" {
		localDir = os.TempDir()
	}
	staleAfter := cfg.StaleAfter
	if staleAfter <= 0 {
		staleAfter = DefaultAnomalyStaleAfter
	}
	return &AnomalyMonitor{
		config:     cfg,
		position:   Position{X: 0, Y: baseY},
		localFile:  filepath.Join(localDir, "srvmonitor_anomalies.log"),
		logger:     logger,
		interval:   iv.Anomaly,
		staleAfter: staleAfter,
		run:        runCombinedOutput,
		now:        time.Now,
	}
}

// Name returns the monitor name.
func (m *AnomalyMonitor) Name() string {
	return "WebSrv Anomaly Monitor"
}

// Start begins monitoring.
func (m *AnomalyMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check performs a fetch + parse cycle.
func (m *AnomalyMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	if err := m.fetch(ctx); err != nil {
		m.setFailure(err)
		sendErr(ctx, errorChan, "Anomaly fetch: "+m.errStr)
		m.display(disp)
		return
	}

	snapshot, err := readAnomalies(m.localFile, AnomalyDisplayCount)
	if err != nil {
		m.setFailure(err)
		sendErr(ctx, errorChan, "Anomaly read: "+m.errStr)
		m.display(disp)
		return
	}

	m.errStr = ""
	m.lines = snapshot.lines
	m.generatedAt = snapshot.generatedAt
	m.hasHeartbeat = snapshot.hasHeartbeat
	m.freshnessAge = 0
	m.stale = false
	if m.hasHeartbeat {
		m.freshnessAge = heartbeatAge(m.clock(), m.generatedAt)
		m.stale = m.freshnessAge > m.staleAfter
	}
	if m.stale {
		sendErr(ctx, errorChan, m.staleAlarm())
	}
	m.display(disp)
}

func (m *AnomalyMonitor) clock() time.Time {
	if m.now != nil {
		return m.now()
	}
	return time.Now()
}

func (m *AnomalyMonitor) setFailure(err error) {
	m.lines = nil
	m.errStr = err.Error()
	m.generatedAt = time.Time{}
	m.freshnessAge = 0
	m.hasHeartbeat = false
	m.stale = false
}

func (m *AnomalyMonitor) staleAlarm() string {
	return fmt.Sprintf(
		"WebSrv anomaly feed STALE: no heartbeat within %s (loganomaly may be down on %s)",
		shortDuration(m.staleAfter),
		m.config.Host,
	)
}

// fetch copies the remote anomalies file to the local temp path via scp.
func (m *AnomalyMonitor) fetch(ctx context.Context) error {
	// Never let a prior successful generation masquerade as the result of a
	// later fetch. scp recreates this destination on success.
	if err := os.Remove(m.localFile); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("remove previous anomaly file: %w", err)
	}
	remote := fmt.Sprintf("%s@%s:%s", m.config.User, m.config.Host, m.config.RemoteFile)
	out, err := m.run(ctx, "/usr/bin/scp",
		"-o", "StrictHostKeyChecking=accept-new",
		"-o", "ConnectTimeout=10",
		remote, m.localFile)
	if err != nil {
		return fmt.Errorf("%w: %s", err, strings.TrimSpace(string(out)))
	}
	return nil
}

// display renders the anomaly section.
func (m *AnomalyMonitor) display(disp Display) {
	y := m.position.Y

	title := "WebSrv Anomalies (last 3)"
	if m.config.Title != "" {
		title = m.config.Title
	}
	header := "----------------- " + title + ": "
	disp.DrawText(Position{X: 0, Y: y}, header, ColorWhite, ColorDefault)

	// Clear and repaint freshness inline with the title. Legacy files without
	// a valid heartbeat intentionally show no freshness state.
	freshnessX := utf8.RuneCountInString(header)
	disp.DrawText(Position{X: freshnessX, Y: y}, strings.Repeat(" ", 40), ColorDefault, ColorDefault)
	if m.errStr == "" && m.hasHeartbeat {
		status := "updated " + shortDuration(m.freshnessAge) + " ago"
		color := ColorGreen
		if m.stale {
			status = "STALE " + shortDuration(m.freshnessAge) + " - loganomaly down?"
			color = ColorRed
		}
		disp.DrawText(Position{X: freshnessX, Y: y}, status, color, ColorDefault)
	}

	// Clear the content rows so shorter new lines fully overwrite older ones.
	blank := strings.Repeat(" ", anomalyMaxLineWidth)
	for i := range AnomalyDisplayCount {
		disp.DrawText(Position{X: 1, Y: y + 1 + i}, blank, ColorDefault, ColorDefault)
	}

	switch {
	case m.errStr != "":
		disp.DrawText(Position{X: 1, Y: y + 1},
			truncateLine("ERROR: "+m.errStr, anomalyMaxLineWidth),
			ColorRed, ColorDefault)
	case len(m.lines) == 0:
		disp.DrawText(Position{X: 1, Y: y + 1}, "None", ColorGreen, ColorDefault)
	default:
		for i, line := range m.lines {
			disp.DrawText(Position{X: 1, Y: y + 1 + i},
				truncateLine(line, anomalyMaxLineWidth),
				ColorRed, ColorDefault)
		}
	}

	disp.Flush()
}

type anomalySnapshot struct {
	lines        []string
	generatedAt  time.Time
	hasHeartbeat bool
}

// readAnomalies streams a bounded anomaly file, strips every heartbeat marker,
// and retains only the requested number of recent non-empty anomaly rows.
func readAnomalies(path string, keep int) (anomalySnapshot, error) {
	file, err := os.Open(filepath.Clean(path)) // #nosec G304 -- path is the monitor's private local file.
	if err != nil {
		return anomalySnapshot{}, err
	}
	defer func() { _ = file.Close() }()
	return parseAnomalies(file, keep)
}

func parseAnomalies(reader io.Reader, keep int) (anomalySnapshot, error) {
	return parseAnomaliesBounded(reader, keep, anomalyMaxFileBytes, anomalyMaxLineBytes)
}

func parseAnomaliesBounded(reader io.Reader, keep int, maxFileBytes int64, maxLineBytes int) (anomalySnapshot, error) {
	limited := &io.LimitedReader{R: reader, N: maxFileBytes + 1}
	scanner := bufio.NewScanner(limited)
	scanner.Buffer(make([]byte, min(64*1024, maxLineBytes)), maxLineBytes)

	snapshot := anomalySnapshot{lines: make([]string, 0, max(keep, 0))}
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\r")
		if after, ok := strings.CutPrefix(line, anomalyHeartbeat); ok {
			seconds, err := strconv.ParseInt(strings.TrimSpace(after), 10, 64)
			if err == nil {
				snapshot.generatedAt = time.Unix(seconds, 0)
				snapshot.hasHeartbeat = true
			}
			continue
		}
		if strings.TrimSpace(line) == "" || keep <= 0 {
			continue
		}
		if len(snapshot.lines) == keep {
			copy(snapshot.lines, snapshot.lines[1:])
			snapshot.lines[len(snapshot.lines)-1] = line
		} else {
			snapshot.lines = append(snapshot.lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return anomalySnapshot{}, fmt.Errorf("scan anomaly file: %w", err)
	}
	if limited.N == 0 {
		return anomalySnapshot{}, fmt.Errorf("anomaly file exceeds %d bytes", maxFileBytes)
	}
	return snapshot, nil
}

func heartbeatAge(now, generatedAt time.Time) time.Duration {
	if generatedAt.After(now) {
		return 0
	}
	return now.Sub(generatedAt)
}

// shortDuration formats a duration compactly, e.g. "45s", "12m", "3h5m".
func shortDuration(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d/time.Second))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d/time.Minute))
	}
	return fmt.Sprintf("%dh%dm", int(d/time.Hour), int(d%time.Hour/time.Minute))
}

// truncateLine shortens s to at most max characters, adding an ellipsis.
func truncateLine(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}
