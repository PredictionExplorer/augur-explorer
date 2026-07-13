// Package termboxui implements srvmonitor's Display interface on top of
// termbox-go. It is the only production package touching the terminal, so
// the monitoring engine stays testable with an in-memory display.
package termboxui

import (
	"sync"

	"github.com/nsf/termbox-go"

	"github.com/PredictionExplorer/augur-explorer/internal/srvmonitor"
)

// Display renders to the terminal via termbox. The mutex serializes cell
// writes from concurrently running monitors. The termbox package functions
// are held as fields so the adapter's own logic (locking, rune layout,
// color mapping) is testable without a TTY.
type Display struct {
	mutex sync.Mutex

	init    func() error
	close   func()
	clear   func(fg, bg termbox.Attribute) error
	setCell func(x, y int, ch rune, fg, bg termbox.Attribute)
	flush   func() error
	size    func() (int, int)
}

// New creates a termbox display.
func New() *Display {
	return &Display{
		init:    termbox.Init,
		close:   termbox.Close,
		clear:   termbox.Clear,
		setCell: termbox.SetCell,
		flush:   termbox.Flush,
		size:    termbox.Size,
	}
}

// Init initializes termbox.
func (d *Display) Init() error {
	return d.init()
}

// Close closes termbox.
func (d *Display) Close() error {
	d.close()
	return nil
}

// Clear clears the screen.
func (d *Display) Clear() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_ = d.clear(termbox.ColorDefault, termbox.ColorDefault) // termbox.Clear never fails after successful Init
}

// DrawText draws simple text at a position, one rune per column.
func (d *Display) DrawText(pos srvmonitor.Position, text string, fg, bg srvmonitor.Color) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	x := pos.X
	for _, r := range text {
		d.setCell(x, pos.Y, r, termbox.Attribute(fg), termbox.Attribute(bg))
		x++
	}
}

// Flush flushes changes to screen.
func (d *Display) Flush() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_ = d.flush() // best-effort: a failed flush only skips one screen refresh
}

// Size returns screen dimensions.
func (d *Display) Size() (width, height int) {
	return d.size()
}
