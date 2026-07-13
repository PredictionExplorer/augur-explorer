package srvmonitor

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
)

func evtConfig(name string) EventTableConfig {
	return EventTableConfig{
		DatabaseConfig: dbConfig(name),
		TableName:      "cg_proc_status",
		ColumnName:     "last_evt_id",
	}
}

func TestEventTableMonitorProgressing(t *testing.T) {
	t.Parallel()
	conn := &fakeConn{rowQueue: []scriptedRow{
		{vals: []any{int64(10)}},
		{vals: []any{int64(11)}},
	}}
	m := NewEventTableMonitor([]EventTableConfig{evtConfig("cg evt")}, 27, testIntervals())
	m.connect = connector(conn, nil)

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	st := m.statuses[0]
	if !st.Alive || st.LastEvtID != 11 || st.ErrStr != "" {
		t.Fatalf("status = %+v, want alive at evt 11", st)
	}
	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	if !conn.closed {
		t.Fatal("per-check connection was not closed")
	}
	// The query must interpolate the configured column/table names.
	if q := conn.queries[0]; q != "SELECT last_evt_id FROM cg_proc_status LIMIT 1" {
		t.Fatalf("query = %q", q)
	}
	row := disp.Row(27)
	if !strings.Contains(row, "11") || !strings.Contains(row, "Alive") || !strings.Contains(row, "cg evt") {
		t.Fatalf("row = %q", row)
	}
}

func TestEventTableMonitorStalledAndDecreasing(t *testing.T) {
	t.Parallel()
	for name, second := range map[string]int64{"stalled": 10, "decreasing": 9} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			conn := &fakeConn{rowQueue: []scriptedRow{
				{vals: []any{int64(10)}},
				{vals: []any{second}},
			}}
			m := NewEventTableMonitor([]EventTableConfig{evtConfig("cg evt")}, 27, testIntervals())
			m.connect = connector(conn, nil)

			errCh := make(chan string, 10)
			m.check(context.Background(), newFakeDisplay(), errCh)

			st := m.statuses[0]
			if st.Alive {
				t.Fatal("non-increasing column must not be alive")
			}
			if !strings.Contains(st.ErrStr, "last_evt_id not increasing") {
				t.Fatalf("ErrStr = %q", st.ErrStr)
			}
			if msgs := drain(errCh); len(msgs) != 1 {
				t.Fatalf("errors = %v", msgs)
			}
		})
	}
}

func TestEventTableMonitorErrorsAndEmpty(t *testing.T) {
	t.Parallel()
	// Connect failure.
	m := NewEventTableMonitor([]EventTableConfig{evtConfig("e")}, 27, testIntervals())
	m.connect = connector(nil, errors.New("refused"))
	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)
	if m.statuses[0].ErrStr != "refused" || len(drain(errCh)) != 1 {
		t.Fatalf("status = %+v", m.statuses[0])
	}

	// First read fails with a real error.
	m0 := NewEventTableMonitor([]EventTableConfig{evtConfig("e")}, 27, testIntervals())
	m0.connect = connector(&fakeConn{rowQueue: []scriptedRow{{err: errors.New("bad table")}}}, nil)
	m0.check(context.Background(), newFakeDisplay(), errCh)
	if m0.statuses[0].ErrStr != "Error bad table" || len(drain(errCh)) != 1 {
		t.Fatalf("status = %+v", m0.statuses[0])
	}

	// Empty table: silent not-alive.
	m2 := NewEventTableMonitor([]EventTableConfig{evtConfig("e")}, 27, testIntervals())
	m2.connect = connector(&fakeConn{rowQueue: []scriptedRow{{err: pgx.ErrNoRows}}}, nil)
	m2.check(context.Background(), newFakeDisplay(), errCh)
	if m2.statuses[0].ErrStr != "" || m2.statuses[0].Alive || len(drain(errCh)) != 0 {
		t.Fatalf("status = %+v", m2.statuses[0])
	}

	// Second read fails.
	m3 := NewEventTableMonitor([]EventTableConfig{evtConfig("e")}, 27, testIntervals())
	m3.connect = connector(&fakeConn{rowQueue: []scriptedRow{
		{vals: []any{int64(3)}},
		{err: errors.New("io")},
	}}, nil)
	m3.check(context.Background(), newFakeDisplay(), errCh)
	if m3.statuses[0].ErrStr != "Error io" || len(drain(errCh)) != 1 {
		t.Fatalf("status = %+v", m3.statuses[0])
	}

	// Second read empty: silent.
	m4 := NewEventTableMonitor([]EventTableConfig{evtConfig("e")}, 27, testIntervals())
	m4.connect = connector(&fakeConn{rowQueue: []scriptedRow{
		{vals: []any{int64(3)}},
		{err: pgx.ErrNoRows},
	}}, nil)
	m4.check(context.Background(), newFakeDisplay(), errCh)
	if m4.statuses[0].ErrStr != "" || m4.statuses[0].Alive || len(drain(errCh)) != 0 {
		t.Fatalf("status = %+v", m4.statuses[0])
	}
}

func TestEventTableMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	m := NewEventTableMonitor([]EventTableConfig{evtConfig("e")}, 27, testIntervals())
	cycles := make(chan struct{}, 100)
	m.connect = func(context.Context, DatabaseConfig) (DBConn, error) {
		select {
		case cycles <- struct{}{}:
		default:
		}
		return nil, errors.New("down")
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		m.Start(ctx, newFakeDisplay(), make(chan string, 100))
		close(done)
	}()

	<-cycles
	<-cycles
	cancel()
	waitFor(t, "loop exit", func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	})
}

func TestEventTableMonitorWaitHookRunsBetweenReads(t *testing.T) {
	t.Parallel()
	conn := &fakeConn{rowQueue: []scriptedRow{
		{vals: []any{int64(1)}},
		{vals: []any{int64(2)}},
	}}
	m := NewEventTableMonitor([]EventTableConfig{evtConfig("e")}, 27, testIntervals())
	m.connect = connector(conn, nil)
	var queriesAtWait int
	m.wait = func(context.Context) {
		conn.mu.Lock()
		queriesAtWait = len(conn.queries)
		conn.mu.Unlock()
	}

	m.check(context.Background(), newFakeDisplay(), make(chan string, 10))

	if queriesAtWait != 1 {
		t.Fatalf("wait ran after %d queries, want 1", queriesAtWait)
	}
}
