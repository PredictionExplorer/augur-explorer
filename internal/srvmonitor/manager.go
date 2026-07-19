package srvmonitor

import (
	"context"
	"fmt"
	"log/slog"
)

// Monitor represents a monitoring component.
type Monitor interface {
	// Name returns the monitor's name.
	Name() string

	// Start begins monitoring; it blocks until ctx is cancelled, so the
	// manager runs it in a goroutine.
	Start(ctx context.Context, display Display, errorChan chan<- string)
}

// Manager coordinates all monitors: it fans their errors into one channel,
// logs them, records them with the alarm tracker and paints the first two on
// the shared error area of the display.
type Manager struct {
	monitors      []Monitor
	display       Display
	errorChan     chan string
	globalErrs    [2]string
	logger        *slog.Logger
	alarmTracker  *AlarmTracker
	errorDisplayX int
	errorDisplayY int
}

// NewManager creates a new monitor manager.
func NewManager(display Display, logger *slog.Logger, alarmTracker *AlarmTracker, errorDisplayX, errorDisplayY int) *Manager {
	return &Manager{
		monitors:      make([]Monitor, 0),
		display:       display,
		errorChan:     make(chan string, 100),
		logger:        logger,
		alarmTracker:  alarmTracker,
		errorDisplayX: errorDisplayX,
		errorDisplayY: errorDisplayY,
	}
}

// Register adds a monitor to the manager.
func (m *Manager) Register(monitor Monitor) {
	m.monitors = append(m.monitors, monitor)
}

// MonitorNames returns the names of all registered monitors in registration
// order.
func (m *Manager) MonitorNames() []string {
	names := make([]string, 0, len(m.monitors))
	for _, mon := range m.monitors {
		names = append(names, mon.Name())
	}
	return names
}

// Start begins all monitors, the error handler and the alarm-state cleanup.
// It returns immediately; everything stops when ctx is cancelled.
func (m *Manager) Start(ctx context.Context) {
	go m.handleErrors(ctx)
	go m.alarmTracker.RunCleanup(ctx)

	for _, monitor := range m.monitors {
		m.logger.Info("Starting monitor: " + monitor.Name())
		go monitor.Start(ctx, m.display, m.errorChan)
	}
}

// handleErrors processes errors from monitors.
func (m *Manager) handleErrors(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-m.errorChan:
			if err == "" {
				continue
			}

			m.logger.Info(fmt.Sprintf("Monitor error: %v", err))

			// Record alarm for mobile notification tracking
			m.alarmTracker.RecordAlarm(ctx, err)

			// Update global error display
			if m.globalErrs[0] == "" {
				m.globalErrs[0] = err
				m.display.DrawText(Position{X: m.errorDisplayX, Y: m.errorDisplayY}, err, ColorYellow, ColorDefault)
				m.display.Flush()
			} else if m.globalErrs[1] == "" {
				m.globalErrs[1] = err
				m.display.DrawText(Position{X: m.errorDisplayX, Y: m.errorDisplayY + 1}, err, ColorYellow, ColorDefault)
				m.display.Flush()
			}
		}
	}
}

// SendTestNotification sends a test notification (for the SIGUSR1 signal
// handler).
func (m *Manager) SendTestNotification(ctx context.Context) {
	m.logger.Info("SIGUSR1 received: Sending test notification")
	m.alarmTracker.SendTestNotification(ctx, "TEST: Server Monitor notification test (triggered by SIGUSR1)")
}
