package cstscan

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var (
	cstDriverSequence atomic.Uint64
	errCSTUnsupported = errors.New("cstscan test driver: unsupported operation")
)

type cstQueryFunc func(context.Context, string, []driver.NamedValue) (driver.Rows, error)

type cstTestDriver struct {
	query cstQueryFunc
}

func (d cstTestDriver) Open(string) (driver.Conn, error) {
	return &cstTestConn{query: d.query}, nil
}

type cstTestConn struct {
	query cstQueryFunc
}

func (c *cstTestConn) Prepare(string) (driver.Stmt, error) {
	return nil, errCSTUnsupported
}

func (c *cstTestConn) Close() error {
	return nil
}

func (c *cstTestConn) Begin() (driver.Tx, error) {
	return nil, errCSTUnsupported
}

func (c *cstTestConn) QueryContext(
	ctx context.Context,
	query string,
	args []driver.NamedValue,
) (driver.Rows, error) {
	return c.query(ctx, query, args)
}

type cstTestRows struct {
	mu          sync.Mutex
	columns     []string
	values      [][]driver.Value
	next        int
	terminalErr error
	onNext      func()
	closed      bool
}

func (r *cstTestRows) Columns() []string {
	return append([]string(nil), r.columns...)
}

func (r *cstTestRows) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closed = true
	return nil
}

func (r *cstTestRows) Next(dest []driver.Value) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.next < len(r.values) {
		copy(dest, r.values[r.next])
		r.next++
		if r.onNext != nil {
			r.onNext()
		}
		return nil
	}
	if r.terminalErr != nil {
		err := r.terminalErr
		r.terminalErr = nil
		return err
	}
	return io.EOF
}

func (r *cstTestRows) isClosed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closed
}

func openCSTTestDB(t *testing.T, query cstQueryFunc) *sql.DB {
	t.Helper()
	name := fmt.Sprintf("cstscan-test-%d", cstDriverSequence.Add(1))
	sql.Register(name, cstTestDriver{query: query})
	db, err := sql.Open(name, "")
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("closing test database: %v", err)
		}
	})
	return db
}

func TestPostgresKeySourceSuccess(t *testing.T) {
	hash := common.HexToHash("0x1234")
	rows := &cstTestRows{
		columns: []string{"tx_hash", "log_index"},
		values:  [][]driver.Value{{"0x" + strings.ToUpper(hash.Hex()[2:]), int64(7)}},
	}
	db := openCSTTestDB(t, func(_ context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
		if !strings.Contains(query, "FROM cg_adm_cst_auclen_chg_div") {
			t.Fatalf("query = %q", query)
		}
		if len(args) != 0 {
			t.Fatalf("args = %v, want none", args)
		}
		return rows, nil
	})

	keys, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background())
	if err != nil {
		t.Fatalf("LoadKeys: %v", err)
	}
	want := map[EventKey]struct{}{{TxHash: hash, LogIndex: 7}: {}}
	if !reflect.DeepEqual(keys, want) {
		t.Fatalf("keys = %#v, want %#v", keys, want)
	}
	if !rows.isClosed() {
		t.Fatal("query rows were not closed")
	}
}

func TestPostgresKeySourceErrors(t *testing.T) {
	t.Run("nil database", func(t *testing.T) {
		if _, err := (PostgresKeySource{}).LoadKeys(context.Background()); err == nil {
			t.Fatal("LoadKeys succeeded without a database")
		}
	})

	t.Run("query", func(t *testing.T) {
		boom := errors.New("query failed")
		db := openCSTTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return nil, boom
		})
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background()); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want query error", err)
		}
	})

	t.Run("scan", func(t *testing.T) {
		db := openCSTTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return &cstTestRows{
				columns: []string{"tx_hash", "log_index"},
				values:  [][]driver.Value{{common.HexToHash("0x01").Hex(), "not-an-index"}},
			}, nil
		})
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background()); err == nil ||
			!strings.Contains(err.Error(), "db scan") {
			t.Fatalf("error = %v, want scan error", err)
		}
	})

	t.Run("negative log index", func(t *testing.T) {
		db := openCSTTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return &cstTestRows{
				columns: []string{"tx_hash", "log_index"},
				values:  [][]driver.Value{{common.HexToHash("0x01").Hex(), int64(-1)}},
			}, nil
		})
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background()); err == nil ||
			!strings.Contains(err.Error(), "negative log index") {
			t.Fatalf("error = %v, want negative-index error", err)
		}
	})

	t.Run("rows", func(t *testing.T) {
		boom := errors.New("rows failed")
		db := openCSTTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return &cstTestRows{
				columns:     []string{"tx_hash", "log_index"},
				terminalErr: boom,
			}, nil
		})
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background()); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want rows error", err)
		}
	})

	t.Run("cancellation while iterating", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		db := openCSTTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return &cstTestRows{
				columns: []string{"tx_hash", "log_index"},
				values:  [][]driver.Value{{common.HexToHash("0x01").Hex(), int64(1)}},
				onNext:  cancel,
			}, nil
		})
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(ctx); !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
	})
}
