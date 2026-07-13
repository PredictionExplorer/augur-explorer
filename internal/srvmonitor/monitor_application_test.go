package srvmonitor

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestApplicationMonitorReportsBlockAndLag(t *testing.T) {
	t.Parallel()
	conn := &fakeConn{rowQueue: []scriptedRow{
		{vals: []any{"1"}},        // chain_id
		{vals: []any{int64(42)}},  // last_evt_id
		{vals: []any{int64(900)}}, // block_num
	}}
	shared := NewSharedRPCState()
	shared.UpdateOfficial("1", 1000)

	m := NewApplicationMonitor([]DatabaseConfig{dbConfig("CG app")}, shared, testIntervals())
	m.connect = connector(conn, nil)

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	st := m.apps[0]
	if st.LastBlockNum != 900 || st.OfficialLagDiff != 100 || st.ErrStr != "" || st.ChainID != "1" {
		t.Fatalf("status = %+v, want block 900 lag 100", st)
	}
	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}

	row := disp.Row(25)
	if !strings.Contains(row, "900") || !strings.Contains(row, "100") || !strings.Contains(row, "CG app") {
		t.Fatalf("row = %q, want block, lag and title", row)
	}
	if header := disp.Row(24); !strings.Contains(header, "Last Block Numbers in Postgres") {
		t.Fatalf("header = %q", header)
	}
}

func TestApplicationMonitorTableSelectionByIndex(t *testing.T) {
	t.Parallel()
	dbs := []DatabaseConfig{dbConfig("cg1"), dbConfig("cg2"), dbConfig("rw1"), dbConfig("rw2")}
	m := NewApplicationMonitor(dbs, nil, testIntervals())

	wantTables := []string{"cg_proc_status", "cg_proc_status", "rw_proc_status", "rw_proc_status"}
	for i, want := range wantTables {
		if m.apps[i].TableName != want {
			t.Fatalf("app %d table = %q, want %q", i, m.apps[i].TableName, want)
		}
	}
}

func TestApplicationMonitorNoRowsFallsBackToZero(t *testing.T) {
	t.Parallel()
	for name, queue := range map[string][]scriptedRow{
		"no contract_addresses": {{err: pgx.ErrNoRows}},
		"no proc status":        {{vals: []any{"1"}}, {err: pgx.ErrNoRows}},
		"no evt_log row":        {{vals: []any{"1"}}, {vals: []any{int64(5)}}, {err: pgx.ErrNoRows}},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			conn := &fakeConn{rowQueue: queue}
			m := NewApplicationMonitor([]DatabaseConfig{dbConfig("app")}, nil, testIntervals())
			m.connect = connector(conn, nil)

			errCh := make(chan string, 10)
			m.check(context.Background(), newFakeDisplay(), errCh)

			st := m.apps[0]
			if st.LastBlockNum != 0 || st.ErrStr != "" {
				t.Fatalf("status = %+v, want silent zero block", st)
			}
			if msgs := drain(errCh); len(msgs) != 0 {
				t.Fatalf("no-rows must not raise errors, got %v", msgs)
			}
		})
	}
}

func TestApplicationMonitorErrors(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		connectErr error
		queue      []scriptedRow
		wantErr    string
	}{
		"connect": {connectErr: errors.New("refused"), wantErr: "refused"},
		"chain query": {
			queue:   []scriptedRow{{err: errors.New("boom")}},
			wantErr: "Error boom",
		},
		"status query": {
			queue:   []scriptedRow{{vals: []any{"1"}}, {err: errors.New("bad table")}},
			wantErr: "Error bad table",
		},
		"block query": {
			queue:   []scriptedRow{{vals: []any{"1"}}, {vals: []any{int64(5)}}, {err: errors.New("io")}},
			wantErr: "Error io",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var conn DBConn
			if tc.connectErr == nil {
				conn = &fakeConn{rowQueue: tc.queue}
			}
			m := NewApplicationMonitor([]DatabaseConfig{dbConfig("app")}, nil, testIntervals())
			m.connect = connector(conn, tc.connectErr)

			errCh := make(chan string, 10)
			m.check(context.Background(), newFakeDisplay(), errCh)

			if m.apps[0].ErrStr != tc.wantErr {
				t.Fatalf("ErrStr = %q, want %q", m.apps[0].ErrStr, tc.wantErr)
			}
			msgs := drain(errCh)
			if len(msgs) != 1 || msgs[0] != tc.wantErr {
				t.Fatalf("errors = %v, want [%q]", msgs, tc.wantErr)
			}
		})
	}
}

func TestApplicationMonitorLagOmittedWithoutOfficialBlock(t *testing.T) {
	t.Parallel()
	conn := &fakeConn{rowQueue: []scriptedRow{
		{vals: []any{"42161"}},
		{vals: []any{int64(1)}},
		{vals: []any{int64(50)}},
	}}
	m := NewApplicationMonitor([]DatabaseConfig{dbConfig("app")}, NewSharedRPCState(), testIntervals())
	m.connect = connector(conn, nil)

	disp := newFakeDisplay()
	m.check(context.Background(), disp, make(chan string, 10))

	if got := m.apps[0].OfficialLagDiff; got != int64(^uint64(0)>>1) {
		t.Fatalf("lag = %d, want MaxInt64 sentinel", got)
	}
	if row := disp.Row(25); !strings.Contains(row, "------") {
		t.Fatalf("row = %q, want ------ lag placeholder", row)
	}
}

func TestApplicationMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	m := NewApplicationMonitor([]DatabaseConfig{dbConfig("app")}, nil, testIntervals())
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
