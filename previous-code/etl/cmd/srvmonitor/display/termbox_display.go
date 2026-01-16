package display

import (
	"sync"
	
	"github.com/nsf/termbox-go"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
)

// TermboxDisplay implements Display using termbox
type TermboxDisplay struct {
	mutex sync.Mutex
}

// NewTermboxDisplay creates a new termbox display
func NewTermboxDisplay() *TermboxDisplay {
	return &TermboxDisplay{}
}

// Init initializes termbox
func (d *TermboxDisplay) Init() error {
	return termbox.Init()
}

// Close closes termbox
func (d *TermboxDisplay) Close() error {
	termbox.Close()
	return nil
}

// Clear clears the screen
func (d *TermboxDisplay) Clear() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

// DrawLine draws a line with multiple segments
func (d *TermboxDisplay) DrawLine(line types.DisplayLine) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	
	x := line.Position.X
	y := line.Position.Y
	
	for _, segment := range line.Segments {
		for _, r := range segment.Text {
			termbox.SetCell(x, y, r, termbox.Attribute(segment.Foreground), termbox.Attribute(segment.Background))
			x++
		}
	}
}

// DrawText draws simple text at a position
func (d *TermboxDisplay) DrawText(pos types.Position, text string, fg, bg types.Color) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	
	x := pos.X
	for _, r := range text {
		termbox.SetCell(x, pos.Y, r, termbox.Attribute(fg), termbox.Attribute(bg))
		x++
	}
}

// Flush flushes changes to screen
func (d *TermboxDisplay) Flush() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	termbox.Flush()
}

// Size returns screen dimensions
func (d *TermboxDisplay) Size() (width, height int) {
	return termbox.Size()
}

// Lock locks the display mutex (for complex operations)
func (d *TermboxDisplay) Lock() {
	d.mutex.Lock()
}

// Unlock unlocks the display mutex
func (d *TermboxDisplay) Unlock() {
	d.mutex.Unlock()
}






