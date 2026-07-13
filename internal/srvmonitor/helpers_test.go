package srvmonitor

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// testIntervals returns polling periods small enough that a test observes
// several cycles without waiting.
func testIntervals() Intervals {
	return Intervals{
		RPC:            time.Millisecond,
		RPCBlockWait:   time.Millisecond,
		DB:             time.Millisecond,
		DBBlockWait:    time.Millisecond,
		EventTable:     time.Millisecond,
		EventTableWait: time.Millisecond,
		Disk:           time.Millisecond,
		Image:          time.Millisecond,
		SSL:            time.Millisecond,
		Anomaly:        time.Millisecond,
	}
}

// cell is one drawn character with its colors.
type cell struct {
	ch rune
	fg Color
	bg Color
}

// fakeDisplay records every drawn cell, mirroring how termbox lays text out
// one rune per column.
type fakeDisplay struct {
	mu      sync.Mutex
	cells   map[[2]int]cell
	flushes int
	cleared int
	width   int
	height  int
}

func newFakeDisplay() *fakeDisplay {
	return &fakeDisplay{cells: make(map[[2]int]cell), width: 120, height: 50}
}

func (d *fakeDisplay) Init() error { return nil }
func (d *fakeDisplay) Close() error {
	return nil
}

func (d *fakeDisplay) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cells = make(map[[2]int]cell)
	d.cleared++
}

func (d *fakeDisplay) DrawText(pos Position, text string, fg, bg Color) {
	d.mu.Lock()
	defer d.mu.Unlock()
	x := pos.X
	for _, r := range text {
		d.cells[[2]int{x, pos.Y}] = cell{ch: r, fg: fg, bg: bg}
		x++
	}
}

func (d *fakeDisplay) Flush() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.flushes++
}

func (d *fakeDisplay) Size() (int, int) {
	return d.width, d.height
}

// Row renders row y as a string with trailing blanks trimmed; cells never
// drawn read as spaces.
func (d *fakeDisplay) Row(y int) string {
	d.mu.Lock()
	defer d.mu.Unlock()
	maxX := -1
	for key := range d.cells {
		if key[1] == y && key[0] > maxX {
			maxX = key[0]
		}
	}
	if maxX < 0 {
		return ""
	}
	var b strings.Builder
	for x := 0; x <= maxX; x++ {
		if c, ok := d.cells[[2]int{x, y}]; ok {
			b.WriteRune(c.ch)
		} else {
			b.WriteRune(' ')
		}
	}
	return strings.TrimRight(b.String(), " ")
}

// FgAt returns the foreground color of the cell at (x, y).
func (d *fakeDisplay) FgAt(x, y int) Color {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.cells[[2]int{x, y}].fg
}

// Flushes returns how many times Flush was called.
func (d *fakeDisplay) Flushes() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.flushes
}

// captureLogger collects log output for assertions.
type captureLogger struct {
	mu  sync.Mutex
	buf strings.Builder
}

func (l *captureLogger) Write(p []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.buf.Write(p)
}

func (l *captureLogger) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.buf.String()
}

func newTestLogger() (*slog.Logger, *captureLogger) {
	sink := &captureLogger{}
	return slog.New(slog.NewTextHandler(sink, nil)), sink
}

// scriptedRow is one QueryRow result: values assigned in order, or an error.
type scriptedRow struct {
	vals []any
	err  error
}

func (r scriptedRow) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}
	if len(dest) != len(r.vals) {
		return fmt.Errorf("scripted row has %d values, Scan got %d destinations", len(r.vals), len(dest))
	}
	for i, v := range r.vals {
		switch d := dest[i].(type) {
		case *int64:
			d64, ok := v.(int64)
			if !ok {
				return fmt.Errorf("destination %d wants int64, scripted %T", i, v)
			}
			*d = d64
		case *string:
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("destination %d wants string, scripted %T", i, v)
			}
			*d = s
		default:
			return fmt.Errorf("unsupported Scan destination %T", dest[i])
		}
	}
	return nil
}

// fakeConn scripts a sequence of QueryRow results and one Query result.
type fakeConn struct {
	mu       sync.Mutex
	rowQueue []scriptedRow
	queries  []string

	queryRows   [][]int64 // rows returned by Query, one int64 column
	queryErr    error
	rowsErr     error // error reported by the returned rows' Err()
	rowsScanErr error // error returned by the rows' Scan

	closed bool
}

func (c *fakeConn) QueryRow(_ context.Context, sql string, _ ...any) pgx.Row {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queries = append(c.queries, sql)
	if len(c.rowQueue) == 0 {
		return scriptedRow{err: fmt.Errorf("unexpected query: %s", sql)}
	}
	row := c.rowQueue[0]
	c.rowQueue = c.rowQueue[1:]
	return row
}

func (c *fakeConn) Query(_ context.Context, sql string, _ ...any) (pgx.Rows, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queries = append(c.queries, sql)
	if c.queryErr != nil {
		return nil, c.queryErr
	}
	return &fakeRows{rows: c.queryRows, err: c.rowsErr, scanErr: c.rowsScanErr}, nil
}

func (c *fakeConn) Close(context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.closed = true
	return nil
}

// fakeRows serves single-int64-column rows through the pgx.Rows interface.
type fakeRows struct {
	rows    [][]int64
	current int
	err     error
	scanErr error
}

func (r *fakeRows) Close()                                       {}
func (r *fakeRows) Err() error                                   { return r.err }
func (r *fakeRows) CommandTag() pgconn.CommandTag                { return pgconn.CommandTag{} }
func (r *fakeRows) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (r *fakeRows) Values() ([]any, error)                       { return nil, nil }
func (r *fakeRows) RawValues() [][]byte                          { return nil }
func (r *fakeRows) Conn() *pgx.Conn                              { return nil }
func (r *fakeRows) Next() bool                                   { return r.current < len(r.rows) }
func (r *fakeRows) Scan(dest ...any) error {
	if r.scanErr != nil {
		return r.scanErr
	}
	row := r.rows[r.current]
	r.current++
	for i := range dest {
		d, ok := dest[i].(*int64)
		if !ok {
			return fmt.Errorf("unsupported Scan destination %T", dest[i])
		}
		*d = row[i]
	}
	return nil
}

// connector returns a Connector serving the given connection (or error).
func connector(conn DBConn, err error) Connector {
	return func(context.Context, DatabaseConfig) (DBConn, error) {
		return conn, err
	}
}

// waitFor polls cond until it holds or the deadline passes.
func waitFor(t *testing.T, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(2 * time.Millisecond)
	}
	t.Fatalf("timed out waiting for %s", what)
}

// drain consumes and returns everything currently buffered on ch.
func drain(ch chan string) []string {
	var out []string
	for {
		select {
		case msg := <-ch:
			out = append(out, msg)
		default:
			return out
		}
	}
}
