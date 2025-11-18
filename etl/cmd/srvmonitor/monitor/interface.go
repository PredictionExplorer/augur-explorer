package monitor

import (
	"context"
	
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
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




