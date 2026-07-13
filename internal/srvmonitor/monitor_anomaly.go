package srvmonitor

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	// AnomalyDisplayCount is the number of most-recent anomalies to show.
	AnomalyDisplayCount = 3
	// anomalyMaxLineWidth truncates long anomaly lines to fit the terminal.
	anomalyMaxLineWidth = 100
)

// AnomalyMonitor fetches the websrv anomalies file (produced by the
// loganomaly tool on the production host) via scp and displays the most
// recent entries.
type AnomalyMonitor struct {
	config    AnomalyConfig
	position  Position
	localFile string
	logger    *log.Logger
	lines     []string
	errStr    string
	interval  time.Duration
	run       CommandRunner
}

// NewAnomalyMonitor creates a new anomaly monitor rooted at baseY. The
// fetched file is stored under localDir (empty selects the system temp
// directory).
func NewAnomalyMonitor(cfg AnomalyConfig, baseY int, logger *log.Logger, localDir string, iv Intervals) *AnomalyMonitor {
	if localDir == "" {
		localDir = os.TempDir()
	}
	return &AnomalyMonitor{
		config:    cfg,
		position:  Position{X: 0, Y: baseY},
		localFile: filepath.Join(localDir, "srvmonitor_anomalies.log"),
		logger:    logger,
		interval:  iv.Anomaly,
		run:       runCombinedOutput,
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
		m.errStr = err.Error()
		sendErr(ctx, errorChan, fmt.Sprintf("Anomaly fetch: %s", m.errStr))
		m.display(disp)
		return
	}

	lines, err := readLastLines(m.localFile, AnomalyDisplayCount)
	if err != nil {
		m.errStr = err.Error()
		sendErr(ctx, errorChan, fmt.Sprintf("Anomaly read: %s", m.errStr))
		m.display(disp)
		return
	}

	m.errStr = ""
	m.lines = lines
	m.display(disp)
}

// fetch copies the remote anomalies file to the local temp path via scp.
func (m *AnomalyMonitor) fetch(ctx context.Context) error {
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
	disp.DrawText(Position{X: 0, Y: y},
		"----------------- "+title+" ----------------",
		ColorWhite, ColorDefault)

	// Clear the content rows so shorter new lines fully overwrite older ones.
	blank := strings.Repeat(" ", anomalyMaxLineWidth)
	for i := 0; i < AnomalyDisplayCount; i++ {
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

// readLastLines returns the last n non-empty lines of a file.
func readLastLines(path string, n int) ([]string, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	raw := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	lines := make([]string, 0, len(raw))
	for _, l := range raw {
		if strings.TrimSpace(l) != "" {
			lines = append(lines, l)
		}
	}
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}
	return lines, nil
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
