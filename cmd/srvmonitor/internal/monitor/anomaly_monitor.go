package monitor

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/display"
	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/types"
)

const (
	UpdateIntervalAnomaly = 300 // seconds between scp fetches (5 minutes)
	AnomalyDisplayCount   = 3   // number of most-recent anomalies to show
	anomalyMaxLineWidth   = 100 // truncate long anomaly lines to fit the terminal
)

// AnomalyMonitor fetches the websrv anomalies file (produced by the loganomaly
// tool on the production host) via scp and displays the most recent entries.
type AnomalyMonitor struct {
	config    types.AnomalyConfig
	position  types.Position
	localFile string
	logger    *log.Logger
	lines     []string
	errStr    string
}

// NewAnomalyMonitor creates a new anomaly monitor rooted at baseY.
func NewAnomalyMonitor(cfg types.AnomalyConfig, baseY int, logger *log.Logger) *AnomalyMonitor {
	tmpDir := os.Getenv("TMPDIR")
	if tmpDir == "" {
		tmpDir = "/tmp"
	}
	return &AnomalyMonitor{
		config:    cfg,
		position:  types.Position{X: 0, Y: baseY},
		localFile: filepath.Join(tmpDir, "srvmonitor_anomalies.log"),
		logger:    logger,
	}
}

// Name returns the monitor name
func (m *AnomalyMonitor) Name() string {
	return "WebSrv Anomaly Monitor"
}

// GetDisplayPosition returns the display position
func (m *AnomalyMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *AnomalyMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
	m.check(disp, errorChan)

	ticker := time.NewTicker(UpdateIntervalAnomaly * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			m.check(disp, errorChan)
		}
	}
}

// check performs a fetch + parse cycle
func (m *AnomalyMonitor) check(disp display.Display, errorChan chan<- string) {
	if err := m.fetch(); err != nil {
		m.errStr = err.Error()
		errorChan <- fmt.Sprintf("Anomaly fetch: %s", m.errStr)
		m.display(disp)
		return
	}

	lines, err := readLastLines(m.localFile, AnomalyDisplayCount)
	if err != nil {
		m.errStr = err.Error()
		errorChan <- fmt.Sprintf("Anomaly read: %s", m.errStr)
		m.display(disp)
		return
	}

	m.errStr = ""
	m.lines = lines
	m.display(disp)
}

// fetch copies the remote anomalies file to the local temp path via scp.
func (m *AnomalyMonitor) fetch() error {
	remote := fmt.Sprintf("%s@%s:%s", m.config.User, m.config.Host, m.config.RemoteFile)
	cmd := exec.Command("/usr/bin/scp", //nolint:gosec // G204: scp target comes from operator-provided monitor config
		"-o", "StrictHostKeyChecking=accept-new",
		"-o", "ConnectTimeout=10",
		remote, m.localFile)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, strings.TrimSpace(string(out)))
	}
	return nil
}

// display renders the anomaly section
func (m *AnomalyMonitor) display(disp display.Display) {
	y := m.position.Y

	title := "WebSrv Anomalies (last 3)"
	if m.config.Title != "" {
		title = m.config.Title
	}
	disp.DrawText(types.Position{X: 0, Y: y},
		"----------------- "+title+" ----------------",
		types.ColorWhite, types.ColorDefault)

	// Clear the content rows so shorter new lines fully overwrite older ones.
	blank := strings.Repeat(" ", anomalyMaxLineWidth)
	for i := 0; i < AnomalyDisplayCount; i++ {
		disp.DrawText(types.Position{X: 1, Y: y + 1 + i}, blank, types.ColorDefault, types.ColorDefault)
	}

	switch {
	case m.errStr != "":
		disp.DrawText(types.Position{X: 1, Y: y + 1},
			truncateLine("ERROR: "+m.errStr, anomalyMaxLineWidth),
			types.ColorRed, types.ColorDefault)
	case len(m.lines) == 0:
		disp.DrawText(types.Position{X: 1, Y: y + 1}, "None", types.ColorGreen, types.ColorDefault)
	default:
		for i, line := range m.lines {
			disp.DrawText(types.Position{X: 1, Y: y + 1 + i},
				truncateLine(line, anomalyMaxLineWidth),
				types.ColorRed, types.ColorDefault)
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
