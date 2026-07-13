package srvmonitor

// Display is the interface monitors render to. The production implementation
// is internal/srvmonitor/termboxui; tests use an in-memory fake.
type Display interface {
	// Init initializes the display.
	Init() error

	// Close closes the display.
	Close() error

	// Clear clears the entire screen.
	Clear()

	// DrawText draws text at a position with colors.
	DrawText(pos Position, text string, fg, bg Color)

	// Flush flushes changes to screen.
	Flush()

	// Size returns the screen dimensions.
	Size() (width, height int)
}
