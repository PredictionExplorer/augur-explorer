package srvmonitor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const (
	// siteCheckDisplayErrors caps how many error lines the section shows.
	// When sites start falling they produce dozens of findings; the operator
	// only needs a red line to notice, not the full list.
	siteCheckDisplayErrors = 2
	// siteCheckDisplayLines is the vertical space the section occupies:
	// a header plus the capped error lines (or the single OK line).
	siteCheckDisplayLines = 1 + siteCheckDisplayErrors
	// siteCheckMaxLineWidth truncates section lines to fit the terminal.
	siteCheckMaxLineWidth = 100
	// siteCheckTimeout bounds one subprocess run. A full three-site check
	// takes ~40s; well past that the run is hung, and the 5-minute cycle
	// must never overlap itself.
	siteCheckTimeout = 4 * time.Minute
)

// siteCheckFinding is one authentic problem in the checker's JSON report.
type siteCheckFinding struct {
	Severity string `json:"severity"`
	Kind     string `json:"kind"`
	Site     string `json:"site"`
	Message  string `json:"message"`
}

// siteCheckReport mirrors the JSON emitted by check-sites.js with "--json -".
// Only the fields the monitor consumes are declared.
type siteCheckReport struct {
	Sites []struct {
		Name      string             `json:"name"`
		Authentic []siteCheckFinding `json:"authentic"`
	} `json:"sites"`
	Endpoints []struct {
		Name      string             `json:"name"`
		Authentic []siteCheckFinding `json:"authentic"`
	} `json:"endpoints"`
	SSL []struct {
		Host     string             `json:"host"`
		Findings []siteCheckFinding `json:"findings"`
	} `json:"ssl"`
	AuthenticCount int `json:"authenticCount"`
}

// SiteCheckMonitor periodically launches the Node headless-browser site
// checker in its own process and reads its JSON report back over a stdout
// pipe. The section below the iDRAC block shows at most
// siteCheckDisplayErrors problems in red (plus a "+N more" count) so a
// falling site is noticeable at a glance without flooding the screen.
type SiteCheckMonitor struct {
	mu sync.Mutex
	// checked is true once the first run finished.
	checked bool
	// errors holds every authentic problem of the last run, formatted.
	errors []string

	cfg      SiteCheckConfig
	position Position
	logger   *slog.Logger
	interval time.Duration
	// run returns stdout only: the JSON report arrives over the stdout
	// pipe while the script's human-readable log goes to stderr.
	run CommandRunner
}

// NewSiteCheckMonitor creates the monitor rooted at baseY.
func NewSiteCheckMonitor(cfg SiteCheckConfig, baseY int, logger *slog.Logger, iv Intervals) *SiteCheckMonitor {
	return &SiteCheckMonitor{
		cfg:      cfg,
		position: Position{X: 0, Y: baseY},
		logger:   logger,
		interval: iv.SiteCheck,
		run:      runStdout,
	}
}

// Name returns the monitor name.
func (m *SiteCheckMonitor) Name() string {
	return "Site Checker (headless browser)"
}

// title returns the display label for the section header.
func (m *SiteCheckMonitor) title() string {
	if m.cfg.Title != "" {
		return m.cfg.Title
	}
	return "Site Checker"
}

// Start begins monitoring.
func (m *SiteCheckMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check runs one subprocess cycle: exec the Node script, read the JSON
// report over the stdout pipe, extract the authentic problems.
func (m *SiteCheckMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	m.logger.Info("Site check: launching " + m.cfg.Script)
	runCtx, cancel := context.WithTimeout(ctx, siteCheckTimeout)
	defer cancel()

	node := m.cfg.Node
	if node == "" {
		node = "node"
	}
	args := []string{m.cfg.Script, "--json", "-"}
	if m.cfg.Config != "" {
		args = append(args, "--config", m.cfg.Config)
	}

	// The script exits 1 when it found problems, so err is expected then;
	// the JSON on stdout is authoritative whenever it parses.
	out, err := m.run(runCtx, node, args...)
	problems := parseSiteCheckOutput(out, err)

	m.mu.Lock()
	m.checked = true
	m.errors = problems
	m.mu.Unlock()

	if len(problems) > 0 {
		m.logger.Info(fmt.Sprintf("Site check: %d problem(s): %s", len(problems), strings.Join(problems, " | ")))
	}
	for i, msg := range problems {
		if i >= siteCheckDisplayErrors {
			break
		}
		sendErr(ctx, errorChan, "Site check: "+msg)
	}

	m.display(disp)
}

// parseSiteCheckOutput turns one subprocess result into the error list. An
// unparseable (or empty) report means the checker itself failed, which is
// reported as a single error so a broken cron never masks a broken site.
func parseSiteCheckOutput(out []byte, runErr error) []string {
	var report siteCheckReport
	if len(out) == 0 || json.Unmarshal(out, &report) != nil {
		detail := "no JSON report"
		if runErr != nil {
			detail = firstNonEmptyLine(stderrOf(runErr)) // script's own error line, if any
			if detail == "" {
				detail = runErr.Error()
			}
		}
		return []string{"checker failed: " + detail}
	}

	var problems []string
	for _, site := range report.Sites {
		for _, f := range site.Authentic {
			problems = append(problems, site.Name+": "+f.Message)
		}
	}
	for _, ep := range report.Endpoints {
		for _, f := range ep.Authentic {
			problems = append(problems, ep.Name+": "+f.Message)
		}
	}
	for _, ssl := range report.SSL {
		for _, f := range ssl.Findings {
			if f.Severity == "warning" {
				continue // expiry warnings are the SSL monitor's job
			}
			problems = append(problems, "SSL "+ssl.Host+": "+f.Message)
		}
	}
	return problems
}

// stderrOf extracts the captured stderr of a finished subprocess, if the
// error carries one.
func stderrOf(err error) string {
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return string(exitErr.Stderr)
	}
	return ""
}

// display renders the section: header, then either a green OK line or up to
// siteCheckDisplayErrors red problem lines, the last one carrying a "+N
// more" count when the run produced more than fits.
func (m *SiteCheckMonitor) display(disp Display) {
	m.mu.Lock()
	defer m.mu.Unlock()

	y := m.position.Y
	disp.DrawText(Position{X: 0, Y: y},
		"----------------- "+m.title()+" (headless browser): ",
		ColorWhite, ColorDefault)

	blank := strings.Repeat(" ", siteCheckMaxLineWidth)
	for i := 0; i < siteCheckDisplayErrors; i++ {
		disp.DrawText(Position{X: 1, Y: y + 1 + i}, blank, ColorDefault, ColorDefault)
	}

	switch {
	case !m.checked:
		disp.DrawText(Position{X: 1, Y: y + 1}, "Checking...", ColorYellow, ColorDefault)
	case len(m.errors) == 0:
		disp.DrawText(Position{X: 1, Y: y + 1}, "All sites OK", ColorGreen, ColorDefault)
	default:
		shown := min(len(m.errors), siteCheckDisplayErrors)
		for i := 0; i < shown; i++ {
			line := m.errors[i]
			if i == shown-1 && len(m.errors) > shown {
				line += fmt.Sprintf(" (+%d more)", len(m.errors)-shown)
			}
			disp.DrawText(Position{X: 1, Y: y + 1 + i},
				truncateLine(line, siteCheckMaxLineWidth),
				ColorRed, ColorDefault)
		}
	}

	disp.Flush()
}
