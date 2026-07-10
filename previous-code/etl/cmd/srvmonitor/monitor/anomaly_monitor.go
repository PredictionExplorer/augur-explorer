package monitor

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/types"
)

const (
	UpdateIntervalAnomaly     = 300  // seconds between scp fetches (5 minutes)
	AnomalyDisplayCount       = 3    // number of most-recent anomalies to show
	anomalyMaxLineWidth       = 100  // truncate long anomaly lines to fit the terminal
	anomalyStaleSecondsDefult = 1800 // default staleness threshold (30 min) if not configured
	anomalyTSMarker           = "#TS=" // first-line generation-time marker written by loganomaly
)

// AnomalyMonitor fetches the websrv anomalies file (produced by the loganomaly
// tool on the production host) via scp and displays the most recent entries.
// It also detects a stale feed (parser dead / never started) via the generation
// timestamp that loganomaly writes on every run.
type AnomalyMonitor struct {
	config    types.AnomalyConfig
	position  types.Position
	localFile string
	logger    *log.Logger
	staleSecs int

	lines   []string
	errStr  string
	genTime time.Time
	hasGen  bool
	stale   bool
}

// NewAnomalyMonitor creates a new anomaly monitor rooted at baseY.
func NewAnomalyMonitor(cfg types.AnomalyConfig, baseY int, logger *log.Logger) *AnomalyMonitor {
	tmpDir := os.Getenv("TMPDIR")
	if tmpDir == "" {
		tmpDir = "/tmp"
	}
	staleSecs := cfg.StaleSecond
	if staleSecs <= 0 {
		staleSecs = anomalyStaleSecondsDefult
	}
	return &AnomalyMonitor{
		config:    cfg,
		position:  types.Position{X: 0, Y: baseY},
		localFile: filepath.Join(tmpDir, "srvmonitor_anomalies.log"),
		logger:    logger,
		staleSecs: staleSecs,
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
		m.stale = false
		errorChan <- fmt.Sprintf("Anomaly fetch: %s", m.errStr)
		m.display(disp)
		return
	}

	genTime, hasGen, anomalies, err := readAnomalies(m.localFile)
	if err != nil {
		m.errStr = err.Error()
		m.stale = false
		errorChan <- fmt.Sprintf("Anomaly read: %s", m.errStr)
		m.display(disp)
		return
	}

	m.errStr = ""
	m.genTime = genTime
	m.hasGen = hasGen
	m.lines = lastN(anomalies, AnomalyDisplayCount)

	// A stale feed means loganomaly stopped regenerating the file (dead or never
	// started). It rewrites the timestamp every run even with zero anomalies, so
	// an old timestamp is a reliable "parser down" signal.
	m.stale = hasGen && time.Since(genTime) > time.Duration(m.staleSecs)*time.Second
	if m.stale {
		errorChan <- fmt.Sprintf("WebSrv anomaly feed STALE: no update for %s (loganomaly may be down on %s)",
			shortDur(time.Since(genTime)), m.config.Host)
	}

	m.display(disp)
}

// fetch copies the remote anomalies file to the local temp path via scp.
func (m *AnomalyMonitor) fetch() error {
	remote := fmt.Sprintf("%s@%s:%s", m.config.User, m.config.Host, m.config.RemoteFile)
	cmd := exec.Command("/usr/bin/scp",
		"-o", "StrictHostKeyChecking=accept-new",
		"-o", "ConnectTimeout=10",
		remote, m.localFile)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, strings.TrimSpace(string(out)))
	}
	return nil
}

// display renders the anomaly section
func (m *AnomalyMonitor) display(disp display.Display) {
	y := m.position.Y

	title := "WebSrv Anomalies (last 3): "
	if m.config.Title != "" {
		title = m.config.Title + ": "
	}
	header := "----------------- " + title
	disp.DrawText(types.Position{X: 0, Y: y}, header, types.ColorWhite, types.ColorDefault)

	// Freshness / staleness indicator drawn inline after the header title.
	fx := len(header)
	disp.DrawText(types.Position{X: fx, Y: y}, strings.Repeat(" ", 40), types.ColorDefault, types.ColorDefault)
	if m.errStr == "" && m.hasGen {
		age := shortDur(time.Since(m.genTime))
		if m.stale {
			disp.DrawText(types.Position{X: fx, Y: y},
				fmt.Sprintf("STALE %s - loganomaly down?", age), types.ColorRed, types.ColorDefault)
		} else {
			disp.DrawText(types.Position{X: fx, Y: y},
				fmt.Sprintf("updated %s ago", age), types.ColorGreen, types.ColorDefault)
		}
	}

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

// readAnomalies reads the anomalies file, separating the generation-time marker
// (if present) from the anomaly lines.
func readAnomalies(path string) (genTime time.Time, hasGen bool, lines []string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return time.Time{}, false, nil, err
	}

	for _, l := range strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n") {
		if strings.HasPrefix(l, anomalyTSMarker) {
			if secs, perr := strconv.ParseInt(strings.TrimSpace(l[len(anomalyTSMarker):]), 10, 64); perr == nil {
				genTime = time.Unix(secs, 0)
				hasGen = true
			}
			continue
		}
		if strings.TrimSpace(l) != "" {
			lines = append(lines, l)
		}
	}
	return genTime, hasGen, lines, nil
}

// lastN returns the last n elements of s (or all of s if shorter).
func lastN(s []string, n int) []string {
	if len(s) > n {
		return s[len(s)-n:]
	}
	return s
}

// shortDur formats a duration compactly, e.g. "45s", "12m", "3h5m".
func shortDur(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	}
	return fmt.Sprintf("%dh%dm", int(d.Hours()), int(d.Minutes())%60)
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
