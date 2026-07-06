package monitor

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/display"
	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/types"
)

// WebAPIStatus holds status for a web API
type WebAPIStatus struct {
	Config       types.WebAPIConfig
	Alive        bool
	ErrStr       string
	PublicAlive  bool
	PublicErrStr string
	X, Y         int
}

// WebAPIMonitor monitors web APIs
type WebAPIMonitor struct {
	apis     []types.WebAPIConfig
	statuses []*WebAPIStatus
	position types.Position
}

// NewWebAPIMonitor creates a new web API monitor
func NewWebAPIMonitor(apis []types.WebAPIConfig, baseY int) *WebAPIMonitor {
	m := &WebAPIMonitor{
		apis:     apis,
		statuses: make([]*WebAPIStatus, len(apis)),
		position: types.Position{X: 0, Y: baseY},
	}

	// Initialize statuses
	for i, api := range apis {
		m.statuses[i] = &WebAPIStatus{
			Config: api,
			X:      1,
			Y:      baseY + 1 + i,
		}
	}

	return m
}

// Name returns the monitor name
func (m *WebAPIMonitor) Name() string {
	return "Web API Monitor"
}

// GetDisplayPosition returns the display position
func (m *WebAPIMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *WebAPIMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.check(disp, errorChan)
			time.Sleep(UpdateIntervalDB * time.Second)
		}
	}
}

// check performs a check cycle
func (m *WebAPIMonitor) check(disp display.Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))

	for _, status := range m.statuses {
		go m.checkAPI(status, &wg, errorChan)
	}

	wg.Wait()
	m.display(disp)
}

// checkAPI checks a single web API. It always probes the internal Host:Port
// endpoint over plain HTTP, and additionally probes the public URL (through
// TLS/stunnel/nginx) when one is configured. Both probes run in parallel and
// write to independent fields so there is no data race.
func (m *WebAPIMonitor) checkAPI(status *WebAPIStatus, wg *sync.WaitGroup, errorChan chan<- string) {
	defer wg.Done()

	var innerWg sync.WaitGroup

	// Internal probe: confirms the backend process + DB are answering.
	// Keeps the lenient 2xx-4xx "alive" rule for backward compatibility.
	innerWg.Add(1)
	go func() {
		defer innerWg.Done()
		url := fmt.Sprintf("http://%s:%s%v", status.Config.Host, status.Config.Port, status.Config.URI)
		alive, err := probeInternal(url)
		if err != nil {
			status.Alive = false
			status.ErrStr = err.Error()
			errorChan <- fmt.Sprintf("%s (internal): %s", status.Config.Title, status.ErrStr)
			return
		}
		status.Alive = alive
		status.ErrStr = ""
		if !alive {
			errorChan <- fmt.Sprintf("%s (internal): unexpected HTTP status", status.Config.Title)
		}
	}()

	// Public probe: exercises the full user-facing path (DNS, TLS, proxy).
	// Only runs when a public URL is configured; requires a strict 200 + body.
	if status.Config.PublicURL != "" {
		innerWg.Add(1)
		go func() {
			defer innerWg.Done()
			if err := probePublic(status.Config.PublicURL); err != nil {
				status.PublicAlive = false
				status.PublicErrStr = err.Error()
				errorChan <- fmt.Sprintf("%s (public %s): %s", status.Config.Title, status.Config.PublicURL, status.PublicErrStr)
				return
			}
			status.PublicAlive = true
			status.PublicErrStr = ""
		}()
	}

	innerWg.Wait()
}

// probeInternal performs the lenient internal reachability check.
func probeInternal(url string) (bool, error) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if resp != nil {
		defer func() { _ = resp.Body.Close() }() // best-effort close on read path
	}
	if err != nil {
		return false, err
	}
	return (resp.StatusCode >= http.StatusOK) && (resp.StatusCode < http.StatusInternalServerError), nil
}

// probePublic performs a strict end-to-end check: exactly HTTP 200 and a
// non-empty body, so a proxy/error page does not read as healthy.
func probePublic(url string) error {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if resp != nil {
		defer func() { _ = resp.Body.Close() }() // best-effort close on read path
	}
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
	if len(body) == 0 {
		return fmt.Errorf("HTTP 200 but empty body")
	}
	return nil
}

// display renders the web API status
func (m *WebAPIMonitor) display(disp display.Display) {
	// Header (Int = internal Host:Port probe, Pub = public TLS/proxy probe)
	disp.DrawText(types.Position{X: 0, Y: m.position.Y},
		"------------------ Web API ( Int / Pub ) -------------------",
		types.ColorWhite, types.ColorDefault)

	// Each API
	for _, status := range m.statuses {
		y := status.Y

		// Internal status (first column)
		intStr := "Alive"
		intColor := types.ColorGreen
		if !status.Alive {
			intStr = "DOWN "
			intColor = types.ColorRed
		}
		disp.DrawText(types.Position{X: status.X, Y: y}, intStr, intColor, types.ColorDefault)

		// Public status (second column) - only when a public URL is configured
		if status.Config.PublicURL != "" {
			pubStr := "Alive"
			pubColor := types.ColorGreen
			if !status.PublicAlive {
				pubStr = "DOWN "
				pubColor = types.ColorRed
			}
			disp.DrawText(types.Position{X: status.X + 8, Y: y}, pubStr, pubColor, types.ColorDefault)
		} else {
			disp.DrawText(types.Position{X: status.X + 8, Y: y}, "  -  ", types.ColorWhite, types.ColorDefault)
		}

		// Title (third column - shifted right)
		disp.DrawText(types.Position{X: status.X + 16, Y: y}, status.Config.Title, types.ColorWhite, types.ColorDefault)

		// Host:Port (fourth column - shifted right)
		hostPort := status.Config.Host + ":" + status.Config.Port
		disp.DrawText(types.Position{X: status.X + 42, Y: y}, hostPort, types.ColorWhite, types.ColorDefault)
	}

	disp.Flush()
}
