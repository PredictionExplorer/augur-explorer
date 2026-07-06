// Package monitor implements the individual monitoring components
// (RPC nodes, databases, disks, web APIs, SSL certificates, images).
package monitor

import (
	"context"

	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/display"
	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/types"
)

// Monitor represents a monitoring component
type Monitor interface {
	// Name returns the monitor's name
	Name() string

	// Start begins monitoring (runs in goroutine)
	Start(ctx context.Context, display display.Display, errorChan chan<- string)

	// GetDisplayPosition returns where this monitor should display
	GetDisplayPosition() types.Position
}
