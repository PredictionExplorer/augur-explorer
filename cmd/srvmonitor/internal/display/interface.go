// Package display abstracts terminal rendering for srvmonitor.
package display

import (
	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/types"
)

// Display is the interface for rendering to screen
type Display interface {
	// Initialize the display
	Init() error

	// Close the display
	Close() error

	// Clear the entire screen
	Clear()

	// Draw a line at the specified position
	DrawLine(line types.DisplayLine)

	// Draw text at a position with colors
	DrawText(pos types.Position, text string, fg, bg types.Color)

	// Flush changes to screen
	Flush()

	// Get screen dimensions
	Size() (width, height int)
}
