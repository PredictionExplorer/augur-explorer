package archive

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

type scriptOp struct {
	kind         string
	contains     string
	columns      []string
	rows         [][]driver.Value
	err          error
	nextErrAt    int
	nextErr      error
	closeErr     error
	rowsAffected int64
	affectedErr  error
	onCall       func()
}

func queryOp(contains string, columns []string, rows ...[]driver.Value) scriptOp {
	return scriptOp{kind: "query", contains: contains, columns: columns, rows: rows, nextErrAt: -1}
}

func queryErrorOp(contains string, err error) scriptOp {
	return scriptOp{kind: "query", contains: contains, err: err, nextErrAt: -1}
}

func prepareOp(contains string) scriptOp {
	return scriptOp{kind: "prepare", contains: contains, nextErrAt: -1}
}

func prepareErrorOp(contains string, err error) scriptOp {
	return scriptOp{kind: "prepare", contains: contains, err: err, nextErrAt: -1}
}

func execOp(contains string, affected int64) scriptOp {
	return scriptOp{kind: "exec", contains: contains, rowsAffected: affected, nextErrAt: -1}
}

func execErrorOp(contains string, err error) scriptOp {
	return scriptOp{kind: "exec", contains: contains, err: err, nextErrAt: -1}
}

type sqlScript struct {
	mu  sync.Mutex
	ops []scriptOp
}

func (s *sqlScript) next(kind, query string) (scriptOp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.ops) == 0 {
		return scriptOp{}, fmt.Errorf("unexpected %s: %s", kind, compactSQL(query))
	}
	op := s.ops[0]
	s.ops = s.ops[1:]
	if op.kind != kind {
		return scriptOp{}, fmt.Errorf("got %s %q, want %s containing %q", kind, compactSQL(query), op.kind, op.contains)
	}
	if op.contains != "" && !strings.Contains(compactSQL(query), compactSQL(op.contains)) {
		return scriptOp{}, fmt.Errorf("%s %q does not contain %q", kind, compactSQL(query), compactSQL(op.contains))
	}
	if op.onCall != nil {
		op.onCall()
	}
	return op, nil
}

func (s *sqlScript) remaining() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.ops)
}

func compactSQL(query string) string {
	return strings.Join(strings.Fields(query), " ")
}

var scriptDriverID atomic.Uint64

func openScriptDB(t *testing.T, ops ...scriptOp) *sql.DB {
	t.Helper()
	script := &sqlScript{ops: append([]scriptOp(nil), ops...)}
	name := fmt.Sprintf("archive-script-%d", scriptDriverID.Add(1))
	sql.Register(name, &scriptDriver{script: script})
	db, err := sql.Open(name, "")
	if err != nil {
		t.Fatalf("opening scripted database: %v", err)
	}
	db.SetMaxOpenConns(1)
	t.Cleanup(func() {
		_ = db.Close()
		if remaining := script.remaining(); remaining != 0 {
			t.Errorf("script has %d unconsumed operations", remaining)
		}
	})
	return db
}

type scriptDriver struct {
	script *sqlScript
}

func (d *scriptDriver) Open(string) (driver.Conn, error) {
	return &scriptConn{script: d.script}, nil
}

type scriptConn struct {
	script *sqlScript
}

func (c *scriptConn) Prepare(query string) (driver.Stmt, error) {
	return c.PrepareContext(context.Background(), query)
}

func (c *scriptConn) PrepareContext(_ context.Context, query string) (driver.Stmt, error) {
	op, err := c.script.next("prepare", query)
	if err != nil {
		return nil, err
	}
	if op.err != nil {
		return nil, op.err
	}
	return &scriptStmt{script: c.script, query: query, closeErr: op.closeErr}, nil
}

func (*scriptConn) Close() error { return nil }

func (*scriptConn) Begin() (driver.Tx, error) {
	return nil, errors.New("transactions are not supported by scripted database")
}

func (c *scriptConn) QueryContext(_ context.Context, query string, _ []driver.NamedValue) (driver.Rows, error) {
	op, err := c.script.next("query", query)
	if err != nil {
		return nil, err
	}
	return rowsForOp(op)
}

func (c *scriptConn) ExecContext(_ context.Context, query string, _ []driver.NamedValue) (driver.Result, error) {
	op, err := c.script.next("exec", query)
	if err != nil {
		return nil, err
	}
	return resultForOp(op)
}

type scriptStmt struct {
	script   *sqlScript
	query    string
	closeErr error
}

func (s *scriptStmt) Close() error { return s.closeErr }

func (*scriptStmt) NumInput() int { return -1 }

func (s *scriptStmt) Exec([]driver.Value) (driver.Result, error) {
	return s.ExecContext(context.Background(), nil)
}

func (s *scriptStmt) ExecContext(_ context.Context, _ []driver.NamedValue) (driver.Result, error) {
	op, err := s.script.next("exec", s.query)
	if err != nil {
		return nil, err
	}
	return resultForOp(op)
}

func (s *scriptStmt) Query([]driver.Value) (driver.Rows, error) {
	return s.QueryContext(context.Background(), nil)
}

func (s *scriptStmt) QueryContext(_ context.Context, _ []driver.NamedValue) (driver.Rows, error) {
	op, err := s.script.next("query", s.query)
	if err != nil {
		return nil, err
	}
	return rowsForOp(op)
}

func rowsForOp(op scriptOp) (driver.Rows, error) {
	if op.err != nil {
		return nil, op.err
	}
	return &scriptRows{
		columns:   op.columns,
		rows:      op.rows,
		nextErrAt: op.nextErrAt,
		nextErr:   op.nextErr,
		closeErr:  op.closeErr,
	}, nil
}

func resultForOp(op scriptOp) (driver.Result, error) {
	if op.err != nil {
		return nil, op.err
	}
	return scriptResult{affected: op.rowsAffected, affectedErr: op.affectedErr}, nil
}

type scriptRows struct {
	columns   []string
	rows      [][]driver.Value
	index     int
	nextErrAt int
	nextErr   error
	closeErr  error
}

func (r *scriptRows) Columns() []string { return r.columns }

func (r *scriptRows) Close() error { return r.closeErr }

func (r *scriptRows) Next(dest []driver.Value) error {
	if r.nextErr != nil && r.index == r.nextErrAt {
		r.index++
		return r.nextErr
	}
	if r.index >= len(r.rows) {
		return io.EOF
	}
	copy(dest, r.rows[r.index])
	r.index++
	return nil
}

type scriptResult struct {
	affected    int64
	affectedErr error
}

func (scriptResult) LastInsertId() (int64, error) { return 0, nil }

func (r scriptResult) RowsAffected() (int64, error) {
	return r.affected, r.affectedErr
}
