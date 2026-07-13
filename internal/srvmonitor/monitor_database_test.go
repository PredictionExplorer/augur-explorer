package srvmonitor

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

func dbConfig(name string) DatabaseConfig {
	return DatabaseConfig{Name: name, Host: "db1.example:5432", DBName: "layer1", User: "monitor", Pass: "pw"}
}

func TestDatabaseMonitorAdvancingBlocks(t *testing.T) {
	t.Parallel()
	conn := &fakeConn{rowQueue: []scriptedRow{
		{vals: []any{int64(100)}},
		{vals: []any{int64(105)}},
	}}
	m := NewDatabaseMonitor([]DatabaseConfig{dbConfig("L1 main")}, testIntervals())
	m.connect = connector(conn, nil)

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	st := m.statuses[0]
	if !st.Alive || st.LastBlockNum != 105 || st.ErrStr != "" {
		t.Fatalf("status = %+v, want alive at block 105", st)
	}
	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	if !conn.closed {
		t.Fatal("per-check connection was not closed")
	}

	row := disp.Row(13)
	if !strings.HasPrefix(row, " Alive") || !strings.Contains(row, "105") || !strings.Contains(row, "L1 main") {
		t.Fatalf("row 13 = %q, want alive row with block and name", row)
	}
	if got := disp.FgAt(1, 13); got != ColorGreen {
		t.Fatalf("alive color = %v, want green", got)
	}
	if header := disp.Row(11); !strings.Contains(header, "SQL DB") {
		t.Fatalf("header row = %q", header)
	}
}

func TestDatabaseMonitorStalledBlocks(t *testing.T) {
	t.Parallel()
	conn := &fakeConn{rowQueue: []scriptedRow{
		{vals: []any{int64(100)}},
		{vals: []any{int64(100)}},
	}}
	m := NewDatabaseMonitor([]DatabaseConfig{dbConfig("L1 main")}, testIntervals())
	m.connect = connector(conn, nil)

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	st := m.statuses[0]
	if st.Alive {
		t.Fatal("stalled database must not be alive")
	}
	if want := "Block difference is zero (last block = 100)"; st.ErrStr != want {
		t.Fatalf("ErrStr = %q, want %q", st.ErrStr, want)
	}
	msgs := drain(errCh)
	if len(msgs) != 1 || msgs[0] != st.ErrStr {
		t.Fatalf("errors = %v, want the stall message", msgs)
	}
	if got := disp.FgAt(1, 13); got != ColorRed {
		t.Fatalf("down color = %v, want red", got)
	}
	if row := disp.Row(13); !strings.HasPrefix(row, " DOWN") {
		t.Fatalf("row = %q, want DOWN", row)
	}
}

func TestDatabaseMonitorConnectFailure(t *testing.T) {
	t.Parallel()
	m := NewDatabaseMonitor([]DatabaseConfig{dbConfig("L1 main")}, testIntervals())
	m.connect = connector(nil, errors.New("connection refused"))

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	st := m.statuses[0]
	if st.Alive || st.ErrStr != "connection refused" {
		t.Fatalf("status = %+v, want connect error", st)
	}
	if msgs := drain(errCh); len(msgs) != 1 || msgs[0] != "connection refused" {
		t.Fatalf("errors = %v", msgs)
	}
}

func TestDatabaseMonitorQueryFailures(t *testing.T) {
	t.Parallel()
	// First read fails.
	conn := &fakeConn{rowQueue: []scriptedRow{{err: errors.New("relation missing")}}}
	m := NewDatabaseMonitor([]DatabaseConfig{dbConfig("L1")}, testIntervals())
	m.connect = connector(conn, nil)
	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)
	if want := "Error relation missing"; m.statuses[0].ErrStr != want {
		t.Fatalf("ErrStr = %q, want %q", m.statuses[0].ErrStr, want)
	}
	if !conn.closed {
		t.Fatal("connection leaked after first-read failure")
	}

	// Second read fails.
	conn2 := &fakeConn{rowQueue: []scriptedRow{
		{vals: []any{int64(7)}},
		{err: errors.New("timeout")},
	}}
	m2 := NewDatabaseMonitor([]DatabaseConfig{dbConfig("L1")}, testIntervals())
	m2.connect = connector(conn2, nil)
	m2.check(context.Background(), newFakeDisplay(), errCh)
	if want := "Error timeout"; m2.statuses[0].ErrStr != want {
		t.Fatalf("ErrStr = %q, want %q", m2.statuses[0].ErrStr, want)
	}
}

func TestDatabaseMonitorEmptyBlockTable(t *testing.T) {
	t.Parallel()
	for name, queue := range map[string][]scriptedRow{
		"first read":  {{err: pgx.ErrNoRows}},
		"second read": {{vals: []any{int64(5)}}, {err: pgx.ErrNoRows}},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			conn := &fakeConn{rowQueue: queue}
			m := NewDatabaseMonitor([]DatabaseConfig{dbConfig("L1")}, testIntervals())
			m.connect = connector(conn, nil)

			errCh := make(chan string, 10)
			m.check(context.Background(), newFakeDisplay(), errCh)

			st := m.statuses[0]
			if st.Alive || st.ErrStr != "" {
				t.Fatalf("status = %+v, want silent not-alive on empty table", st)
			}
			if msgs := drain(errCh); len(msgs) != 0 {
				t.Fatalf("empty table must not raise errors, got %v", msgs)
			}
		})
	}
}

func TestDatabaseMonitorBlockWaitHookRunsBetweenReads(t *testing.T) {
	t.Parallel()
	conn := &fakeConn{rowQueue: []scriptedRow{
		{vals: []any{int64(1)}},
		{vals: []any{int64(2)}},
	}}
	m := NewDatabaseMonitor([]DatabaseConfig{dbConfig("L1")}, testIntervals())
	m.connect = connector(conn, nil)
	var queriesAtWait int
	m.blockWait = func(context.Context) {
		conn.mu.Lock()
		queriesAtWait = len(conn.queries)
		conn.mu.Unlock()
	}

	m.check(context.Background(), newFakeDisplay(), make(chan string, 10))

	if queriesAtWait != 1 {
		t.Fatalf("blockWait ran after %d queries, want 1 (between the two reads)", queriesAtWait)
	}
}

func TestDatabaseMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	conn := &fakeConn{}
	m := NewDatabaseMonitor([]DatabaseConfig{dbConfig("L1")}, testIntervals())
	cycles := make(chan struct{}, 100)
	m.connect = func(context.Context, DatabaseConfig) (DBConn, error) {
		select {
		case cycles <- struct{}{}:
		default:
		}
		return conn, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		m.Start(ctx, newFakeDisplay(), make(chan string, 100))
		close(done)
	}()

	<-cycles
	<-cycles // at least two loop iterations ran
	cancel()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Start did not stop on cancellation")
	}
}
