package srvmonitor

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	// idracCrashMarker prefixes the script's crash line ("CRASH <detail>").
	idracCrashMarker = "CRASH"
	// idracOverlayInterval is how often the full-screen crash box is
	// repainted so other monitors' section redraws cannot erase it.
	idracOverlayInterval = time.Second
	// idracMaxLineWidth truncates section lines to fit the terminal.
	idracMaxLineWidth = 100
)

// envRunner executes an external command with an explicit environment and
// returns its combined output. The iDRAC monitor passes credentials through
// the environment (never argv, which is world-readable in /proc); tests
// inject fakes.
type envRunner func(ctx context.Context, env []string, name string, args ...string) ([]byte, error)

// runWithEnv is the production envRunner.
func runWithEnv(ctx context.Context, env []string, name string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args...) //nolint:gosec // G204: command comes from operator-provided monitor config
	cmd.Env = env
	return cmd.CombinedOutput()
}

// DefaultIDRACScriptPath returns the idrac_check.sh path next to the running
// binary, mirroring how the script is shipped in cmd/srvmonitor.
func DefaultIDRACScriptPath() string {
	exe, err := os.Executable()
	if err != nil {
		return "idrac_check.sh"
	}
	return filepath.Join(filepath.Dir(exe), "idrac_check.sh")
}

// IDRACStatus holds the last check result for one iDRAC.
type IDRACStatus struct {
	Config  IDRACConfig
	Checked bool   // at least one check completed
	Crashed bool   // a fatal SEL entry is currently active
	Detail  string // crash detail from the script (timestamp + message)
	ErrStr  string // check failure (unreachable iDRAC, bad credentials, ...)
}

// title returns the display label for the entry.
func (s *IDRACStatus) title() string {
	if s.Config.Name != "" {
		return s.Config.Name
	}
	return s.Config.Host
}

// IDRACMonitor periodically runs the read-only idrac_check.sh script against
// every configured iDRAC. A detected crash paints a large centered red box
// ("SERVER <host> CRASHED") over the dashboard; the box is repainted every
// second while the crash is active so it stays visible no matter what the
// other monitors redraw. The monitor never issues any power action.
type IDRACMonitor struct {
	mu       sync.Mutex
	statuses []*IDRACStatus

	scriptPath      string
	position        Position
	logger          *slog.Logger
	interval        time.Duration
	overlayInterval time.Duration
	run             envRunner
	environ         func() []string
}

// NewIDRACMonitor creates an iDRAC crash monitor rooted at baseY. An empty
// scriptPath selects idrac_check.sh next to the running binary.
func NewIDRACMonitor(cfgs []IDRACConfig, scriptPath string, baseY int, logger *slog.Logger, iv Intervals) *IDRACMonitor {
	if scriptPath == "" {
		scriptPath = DefaultIDRACScriptPath()
	}
	statuses := make([]*IDRACStatus, len(cfgs))
	for i, cfg := range cfgs {
		statuses[i] = &IDRACStatus{Config: cfg}
	}
	return &IDRACMonitor{
		statuses:        statuses,
		scriptPath:      scriptPath,
		position:        Position{X: 0, Y: baseY},
		logger:          logger,
		interval:        iv.IDRAC,
		overlayInterval: idracOverlayInterval,
		run:             runWithEnv,
		environ:         os.Environ,
	}
}

// Name returns the monitor name.
func (m *IDRACMonitor) Name() string {
	return "iDRAC Crash Monitor"
}

// Start begins monitoring. Beside the usual check loop it runs the overlay
// painter, which keeps the red crash box on top of the dashboard.
func (m *IDRACMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	go m.runOverlay(ctx, disp)
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check runs one script invocation per iDRAC in parallel.
func (m *IDRACMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	m.logger.Info(fmt.Sprintf("iDRAC monitor: starting check cycle for %d interface(s)", len(m.statuses)))

	var wg sync.WaitGroup
	wg.Add(len(m.statuses))
	for _, status := range m.statuses {
		go func(status *IDRACStatus) {
			defer wg.Done()
			m.checkOne(ctx, status, errorChan)
		}(status)
	}
	wg.Wait()

	m.display(disp)
}

// checkOne runs the script for one iDRAC and interprets its result:
// "OK" means healthy, "CRASH <detail>" means a fatal SEL entry is active,
// anything else is a failure of the check itself.
func (m *IDRACMonitor) checkOne(ctx context.Context, status *IDRACStatus, errorChan chan<- string) {
	env := append(m.environ(),
		"IDRAC="+status.Config.Host,
		"IDRAC_USER="+status.Config.User,
		"IDRAC_PASS="+status.Config.Pass,
	)
	out, err := m.run(ctx, env, m.scriptPath)
	line := firstNonEmptyLine(string(out))

	m.mu.Lock()
	defer m.mu.Unlock()
	status.Checked = true

	switch {
	case strings.HasPrefix(line, idracCrashMarker):
		// The script exits 2 on a crash, so err is non-nil here by design;
		// the CRASH marker is authoritative.
		status.Crashed = true
		status.Detail = strings.TrimSpace(strings.TrimPrefix(line, idracCrashMarker))
		status.ErrStr = ""
		alarm := fmt.Sprintf("Server %s CRASHED: %s", status.title(), status.Detail)
		m.logger.Info("iDRAC monitor: " + alarm)
		sendErr(ctx, errorChan, alarm)
	case err != nil:
		status.Crashed = false
		status.Detail = ""
		status.ErrStr = err.Error()
		if line != "" {
			status.ErrStr = line
		}
		sendErr(ctx, errorChan, fmt.Sprintf("iDRAC %s check failed: %s", status.title(), status.ErrStr))
	default:
		status.Crashed = false
		status.Detail = ""
		status.ErrStr = ""
	}
}

// display renders the iDRAC section.
func (m *IDRACMonitor) display(disp Display) {
	m.mu.Lock()
	defer m.mu.Unlock()

	y := m.position.Y
	disp.DrawText(Position{X: 0, Y: y},
		"----------------- iDRAC Health (crash detection): ",
		ColorWhite, ColorDefault)

	blank := strings.Repeat(" ", idracMaxLineWidth)
	for i, status := range m.statuses {
		row := y + 1 + i
		disp.DrawText(Position{X: 1, Y: row}, blank, ColorDefault, ColorDefault)

		label := status.title() + ": "
		switch {
		case !status.Checked:
			disp.DrawText(Position{X: 1, Y: row}, label+"Checking...", ColorYellow, ColorDefault)
		case status.Crashed:
			disp.DrawText(Position{X: 1, Y: row},
				truncateLine(label+"CRASHED "+status.Detail, idracMaxLineWidth),
				ColorWhite, ColorRed)
		case status.ErrStr != "":
			disp.DrawText(Position{X: 1, Y: row},
				truncateLine(label+"ERROR: "+status.ErrStr, idracMaxLineWidth),
				ColorRed, ColorDefault)
		default:
			disp.DrawText(Position{X: 1, Y: row}, label+"OK", ColorGreen, ColorDefault)
		}
	}

	disp.Flush()
}

// runOverlay repaints the crash box every overlayInterval while any iDRAC
// reports a crash. When the last crash clears it wipes the screen once so
// the stale red box disappears; the section monitors repaint themselves on
// their next cycles.
func (m *IDRACMonitor) runOverlay(ctx context.Context, disp Display) {
	ticker := time.NewTicker(m.overlayInterval)
	defer ticker.Stop()

	active := false
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			crashed := m.crashedTitles()
			switch {
			case len(crashed) > 0:
				m.drawOverlay(disp, crashed)
				active = true
			case active:
				disp.Clear()
				disp.Flush()
				active = false
			}
		}
	}
}

// crashedTitles returns the display labels of every crashed server.
func (m *IDRACMonitor) crashedTitles() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	var crashed []string
	for _, status := range m.statuses {
		if status.Crashed {
			crashed = append(crashed, status.title())
		}
	}
	return crashed
}

// drawOverlay paints a large centered solid-red box announcing every crashed
// server in white-on-red text.
func (m *IDRACMonitor) drawOverlay(disp Display, crashed []string) {
	width, height := disp.Size()

	lines := make([]string, 0, len(crashed)+2)
	lines = append(lines, "!!!  CRITICAL SYSTEM FAILURE  !!!", "")
	for _, title := range crashed {
		lines = append(lines, fmt.Sprintf("SERVER  %s  CRASHED", title))
	}

	boxW := 50
	for _, line := range lines {
		if len(line)+12 > boxW {
			boxW = len(line) + 12
		}
	}
	if boxW > width {
		boxW = width
	}
	boxH := len(lines) + 4
	x0 := max((width-boxW)/2, 0)
	y0 := max((height-boxH)/2, 0)

	blank := strings.Repeat(" ", boxW)
	for row := range boxH {
		disp.DrawText(Position{X: x0, Y: y0 + row}, blank, ColorWhite, ColorRed)
	}
	for i, line := range lines {
		x := max(x0+(boxW-len(line))/2, x0)
		disp.DrawText(Position{X: x, Y: y0 + 2 + i}, truncateLine(line, boxW), ColorWhite, ColorRed)
	}

	disp.Flush()
}

// firstNonEmptyLine returns the first non-blank line of s, trimmed.
func firstNonEmptyLine(s string) string {
	for line := range strings.Lines(s) {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			return trimmed
		}
	}
	return ""
}
