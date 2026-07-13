package srvmonitor

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// webAPIProbeTimeout bounds one HTTP probe.
const webAPIProbeTimeout = 10 * time.Second

// WebAPIStatus holds status for a web API.
type WebAPIStatus struct {
	Config       WebAPIConfig
	Alive        bool
	ErrStr       string
	PublicAlive  bool
	PublicErrStr string
	X, Y         int
}

// WebAPIMonitor probes each API's internal Host:Port endpoint and, when
// configured, its public URL through the TLS/proxy path.
type WebAPIMonitor struct {
	statuses []*WebAPIStatus
	position Position
	interval time.Duration
	client   *http.Client
}

// NewWebAPIMonitor creates a new web API monitor.
func NewWebAPIMonitor(apis []WebAPIConfig, baseY int, iv Intervals) *WebAPIMonitor {
	m := &WebAPIMonitor{
		statuses: make([]*WebAPIStatus, len(apis)),
		position: Position{X: 0, Y: baseY},
		interval: iv.DB,
		client:   &http.Client{Timeout: webAPIProbeTimeout},
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

// Name returns the monitor name.
func (m *WebAPIMonitor) Name() string {
	return "Web API Monitor"
}

// Start begins monitoring.
func (m *WebAPIMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check performs a check cycle.
func (m *WebAPIMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))

	for _, status := range m.statuses {
		go func(status *WebAPIStatus) {
			defer wg.Done()
			m.checkAPI(ctx, status, errorChan)
		}(status)
	}

	wg.Wait()
	m.display(disp)
}

// checkAPI checks a single web API. It always probes the internal Host:Port
// endpoint over plain HTTP, and additionally probes the public URL (through
// TLS/stunnel/nginx) when one is configured. Both probes run in parallel and
// write to independent fields so there is no data race.
func (m *WebAPIMonitor) checkAPI(ctx context.Context, status *WebAPIStatus, errorChan chan<- string) {
	var innerWg sync.WaitGroup

	// Internal probe: confirms the backend process + DB are answering.
	// Keeps the lenient 2xx-4xx "alive" rule for backward compatibility.
	innerWg.Add(1)
	go func() {
		defer innerWg.Done()
		url := fmt.Sprintf("http://%s:%s%v", status.Config.Host, status.Config.Port, status.Config.URI)
		alive, err := m.probeInternal(ctx, url)
		if err != nil {
			status.Alive = false
			status.ErrStr = err.Error()
			sendErr(ctx, errorChan, fmt.Sprintf("%s (internal): %s", status.Config.Title, status.ErrStr))
			return
		}
		status.Alive = alive
		status.ErrStr = ""
		if !alive {
			sendErr(ctx, errorChan, fmt.Sprintf("%s (internal): unexpected HTTP status", status.Config.Title))
		}
	}()

	// Public probe: exercises the full user-facing path (DNS, TLS, proxy).
	// Only runs when a public URL is configured; requires a strict 200 + body.
	if status.Config.PublicURL != "" {
		innerWg.Add(1)
		go func() {
			defer innerWg.Done()
			if err := m.probePublic(ctx, status.Config.PublicURL); err != nil {
				status.PublicAlive = false
				status.PublicErrStr = err.Error()
				sendErr(ctx, errorChan, fmt.Sprintf("%s (public %s): %s", status.Config.Title, status.Config.PublicURL, status.PublicErrStr))
				return
			}
			status.PublicAlive = true
			status.PublicErrStr = ""
		}()
	}

	innerWg.Wait()
}

// probeInternal performs the lenient internal reachability check.
func (m *WebAPIMonitor) probeInternal(ctx context.Context, url string) (bool, error) {
	resp, err := m.get(ctx, url)
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
func (m *WebAPIMonitor) probePublic(ctx context.Context, url string) error {
	resp, err := m.get(ctx, url)
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

// get issues a context-aware GET through the monitor's client.
func (m *WebAPIMonitor) get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return m.client.Do(req)
}

// display renders the web API status.
func (m *WebAPIMonitor) display(disp Display) {
	// Header (Int = internal Host:Port probe, Pub = public TLS/proxy probe)
	disp.DrawText(Position{X: 0, Y: m.position.Y},
		"------------------ Web API ( Int / Pub ) -------------------",
		ColorWhite, ColorDefault)

	// Each API
	for _, status := range m.statuses {
		y := status.Y

		// Internal status (first column)
		intStr := "Alive"
		intColor := ColorGreen
		if !status.Alive {
			intStr = "DOWN "
			intColor = ColorRed
		}
		disp.DrawText(Position{X: status.X, Y: y}, intStr, intColor, ColorDefault)

		// Public status (second column) - only when a public URL is configured
		if status.Config.PublicURL != "" {
			pubStr := "Alive"
			pubColor := ColorGreen
			if !status.PublicAlive {
				pubStr = "DOWN "
				pubColor = ColorRed
			}
			disp.DrawText(Position{X: status.X + 8, Y: y}, pubStr, pubColor, ColorDefault)
		} else {
			disp.DrawText(Position{X: status.X + 8, Y: y}, "  -  ", ColorWhite, ColorDefault)
		}

		// Title (third column - shifted right)
		disp.DrawText(Position{X: status.X + 16, Y: y}, status.Config.Title, ColorWhite, ColorDefault)

		// Host:Port (fourth column - shifted right)
		hostPort := status.Config.Host + ":" + status.Config.Port
		disp.DrawText(Position{X: status.X + 42, Y: y}, hostPort, ColorWhite, ColorDefault)
	}

	disp.Flush()
}
