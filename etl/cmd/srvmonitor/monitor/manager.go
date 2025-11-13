package monitor

import (
	"context"
	"log"
	
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
)

// Manager coordinates all monitors
type Manager struct {
	monitors   []Monitor
	display    display.Display
	errorChan  chan string
	globalErrs [2]string
	logger     *log.Logger
}

// NewManager creates a new monitor manager
func NewManager(display display.Display, logger *log.Logger) *Manager {
	return &Manager{
		monitors:  make([]Monitor, 0),
		display:   display,
		errorChan: make(chan string, 100),
		logger:    logger,
	}
}

// Register adds a monitor to the manager
func (m *Manager) Register(monitor Monitor) {
	m.monitors = append(m.monitors, monitor)
}

// Start begins all monitors
func (m *Manager) Start(ctx context.Context) {
	// Start error handler
	go m.handleErrors(ctx)
	
	// Start all monitors
	for _, monitor := range m.monitors {
		m.logger.Printf("Starting monitor: %s", monitor.Name())
		go monitor.Start(ctx, m.display, m.errorChan)
	}
}

// handleErrors processes errors from monitors
func (m *Manager) handleErrors(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-m.errorChan:
			if err == "" {
				continue
			}
			
			m.logger.Printf("Monitor error: %v", err)
			
			// Update global error display
			if m.globalErrs[0] == "" {
				m.globalErrs[0] = err
				m.display.DrawText(types.Position{X: 1, Y: 35}, err, types.ColorYellow, types.ColorDefault)
				m.display.Flush()
			} else if m.globalErrs[1] == "" {
				m.globalErrs[1] = err
				m.display.DrawText(types.Position{X: 1, Y: 36}, err, types.ColorYellow, types.ColorDefault)
				m.display.Flush()
			}
		}
	}
}


