package termboxui

import (
	"errors"
	"testing"

	"github.com/nsf/termbox-go"

	"github.com/PredictionExplorer/augur-explorer/internal/srvmonitor"
)

// recorded is one SetCell invocation.
type recorded struct {
	x, y   int
	ch     rune
	fg, bg termbox.Attribute
}

// fakeBackend captures every delegated termbox call.
type fakeBackend struct {
	initErr  error
	inits    int
	closes   int
	clears   int
	clearFg  termbox.Attribute
	clearBg  termbox.Attribute
	cells    []recorded
	flushes  int
	flushErr error
	w, h     int
}

func newDisplay(b *fakeBackend) *Display {
	return &Display{
		init:  func() error { b.inits++; return b.initErr },
		close: func() { b.closes++ },
		clear: func(fg, bg termbox.Attribute) error {
			b.clears++
			b.clearFg, b.clearBg = fg, bg
			return nil
		},
		setCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
			b.cells = append(b.cells, recorded{x: x, y: y, ch: ch, fg: fg, bg: bg})
		},
		flush: func() error { b.flushes++; return b.flushErr },
		size:  func() (int, int) { return b.w, b.h },
	}
}

func TestNewWiresTermboxFunctions(t *testing.T) {
	t.Parallel()
	d := New()
	if d.init == nil || d.close == nil || d.clear == nil || d.setCell == nil || d.flush == nil || d.size == nil {
		t.Fatal("New left a termbox delegate nil")
	}
}

func TestInitCloseDelegation(t *testing.T) {
	t.Parallel()
	b := &fakeBackend{}
	d := newDisplay(b)

	if err := d.Init(); err != nil {
		t.Fatal(err)
	}
	if err := d.Close(); err != nil {
		t.Fatal(err)
	}
	if b.inits != 1 || b.closes != 1 {
		t.Fatalf("backend = %+v", b)
	}

	// Init errors (no TTY) pass through.
	b.initErr = errors.New("no tty")
	if err := d.Init(); err == nil {
		t.Fatal("init error must propagate")
	}
}

func TestDrawTextLaysOutRunes(t *testing.T) {
	t.Parallel()
	b := &fakeBackend{}
	d := newDisplay(b)

	d.DrawText(srvmonitor.Position{X: 5, Y: 2}, "ab", srvmonitor.ColorGreen, srvmonitor.ColorDefault)

	if len(b.cells) != 2 {
		t.Fatalf("cells = %+v", b.cells)
	}
	first, second := b.cells[0], b.cells[1]
	if first.x != 5 || first.y != 2 || first.ch != 'a' {
		t.Fatalf("first = %+v", first)
	}
	// The column advances one cell per rune; colors map 1:1 onto termbox
	// attributes.
	if second.x != 6 || second.y != 2 || second.ch != 'b' {
		t.Fatalf("second = %+v", second)
	}
	if first.fg != termbox.ColorGreen || first.bg != termbox.ColorDefault {
		t.Fatalf("colors = %+v", first)
	}
}

func TestDrawTextMultibyteRunes(t *testing.T) {
	t.Parallel()
	b := &fakeBackend{}
	d := newDisplay(b)

	d.DrawText(srvmonitor.Position{X: 0, Y: 0}, "é⚠", srvmonitor.ColorRed, srvmonitor.ColorDefault)

	if len(b.cells) != 2 {
		t.Fatalf("cells = %+v", b.cells)
	}
	if b.cells[0].ch != 'é' || b.cells[1].ch != '⚠' || b.cells[1].x != 1 {
		t.Fatalf("cells = %+v", b.cells)
	}
}

func TestClearFlushSize(t *testing.T) {
	t.Parallel()
	b := &fakeBackend{w: 80, h: 24}
	d := newDisplay(b)

	d.Clear()
	if b.clears != 1 || b.clearFg != termbox.ColorDefault || b.clearBg != termbox.ColorDefault {
		t.Fatalf("backend = %+v", b)
	}

	d.Flush()
	b.flushErr = errors.New("broken pipe")
	d.Flush() // flush errors are deliberately swallowed
	if b.flushes != 2 {
		t.Fatalf("flushes = %d", b.flushes)
	}

	if w, h := d.Size(); w != 80 || h != 24 {
		t.Fatalf("size = %dx%d", w, h)
	}
}

func TestConcurrentDrawsAreSerialized(t *testing.T) {
	t.Parallel()
	b := &fakeBackend{}
	d := newDisplay(b)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 100; i++ {
			d.DrawText(srvmonitor.Position{X: 0, Y: 1}, "xy", srvmonitor.ColorWhite, srvmonitor.ColorDefault)
		}
	}()
	for i := 0; i < 100; i++ {
		d.DrawText(srvmonitor.Position{X: 0, Y: 2}, "ab", srvmonitor.ColorWhite, srvmonitor.ColorDefault)
		d.Clear()
		d.Flush()
	}
	<-done

	// The mutex kept every two-rune write atomic: cells arrive in pairs.
	for i := 0; i+1 < len(b.cells); i += 2 {
		if b.cells[i].y != b.cells[i+1].y {
			t.Fatalf("interleaved write at %d: %+v %+v", i, b.cells[i], b.cells[i+1])
		}
	}
}
