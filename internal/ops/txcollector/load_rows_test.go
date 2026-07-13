package txcollector

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
	txDriverSequence atomic.Uint64
	errTXUnsupported = errors.New("txcollector test driver: unsupported operation")
)

type txQueryFunc func(context.Context, string, []driver.NamedValue) (driver.Rows, error)

type txTestDriver struct {
	query txQueryFunc
}

func (d txTestDriver) Open(string) (driver.Conn, error) {
	return &txTestConn{query: d.query}, nil
}

type txTestConn struct {
	query txQueryFunc
}

func (c *txTestConn) Prepare(string) (driver.Stmt, error) {
	return nil, errTXUnsupported
}

func (c *txTestConn) Close() error {
	return nil
}

func (c *txTestConn) Begin() (driver.Tx, error) {
	return nil, errTXUnsupported
}

func (c *txTestConn) QueryContext(
	ctx context.Context,
	query string,
	args []driver.NamedValue,
) (driver.Rows, error) {
	return c.query(ctx, query, args)
}

type txTestRows struct {
	mu          sync.Mutex
	columns     []string
	values      [][]driver.Value
	next        int
	terminalErr error
	onNext      func()
	closed      bool
}

func (r *txTestRows) Columns() []string {
	return append([]string(nil), r.columns...)
}

func (r *txTestRows) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closed = true
	return nil
}

func (r *txTestRows) Next(dest []driver.Value) error {
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

func (r *txTestRows) isClosed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closed
}

func openTXTestDB(t *testing.T, query txQueryFunc) *sql.DB {
	t.Helper()
	name := fmt.Sprintf("txcollector-test-%d", txDriverSequence.Add(1))
	sql.Register(name, txTestDriver{query: query})
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

func eventRowColumns() []string {
	return []string{"block_num", "log_index", "tx_hash", "addr", "topic0_sig", "log_rlp"}
}

func TestLoadEventRowsSuccessAndQueryArguments(t *testing.T) {
	txHash := common.HexToHash("0xabcd")
	address := common.HexToAddress("0xabcdef0000000000000000000000000000001234")
	rows := &txTestRows{
		columns: eventRowColumns(),
		values: [][]driver.Value{{
			int64(125),
			int64(7),
			"0x" + strings.ToUpper(txHash.Hex()[2:]),
			strings.ToLower(address.Hex()),
			"deadbeef",
			[]byte{1, 2, 3},
		}},
	}
	contracts := []string{address.Hex()}
	db := openTXTestDB(t, func(_ context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
		if !strings.Contains(query, "a.addr = ANY($1)") ||
			!strings.Contains(query, "e.block_num >= $2") ||
			!strings.Contains(query, "ORDER BY t.tx_hash, e.log_index") {
			t.Fatalf("query = %q", query)
		}
		if len(args) != 2 {
			t.Fatalf("args = %v, want two", args)
		}
		if got, ok := args[0].Value.(string); !ok || !strings.Contains(got, contracts[0]) {
			t.Fatalf("contract argument = %#v", args[0].Value)
		}
		if args[1].Value != int64(120) {
			t.Fatalf("from-block argument = %#v, want int64(120)", args[1].Value)
		}
		return rows, nil
	})

	got, err := LoadEventRows(context.Background(), db, contracts, 120)
	if err != nil {
		t.Fatalf("LoadEventRows: %v", err)
	}
	want := []EventRow{{
		BlockNum:     125,
		LogIndex:     7,
		TxHash:       txHash.Hex(),
		ContractAddr: address.Hex(),
		Topic0Sig:    "deadbeef",
		LogRLP:       []byte{1, 2, 3},
	}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("rows = %#v, want %#v", got, want)
	}
	if !rows.isClosed() {
		t.Fatal("query rows were not closed")
	}
}

func TestLoadEventRowsErrors(t *testing.T) {
	t.Run("nil database", func(t *testing.T) {
		if _, err := LoadEventRows(context.Background(), nil, []string{"0x1"}, 0); err == nil {
			t.Fatal("LoadEventRows succeeded without a database")
		}
	})

	t.Run("empty contracts", func(t *testing.T) {
		db := openTXTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			t.Fatal("query called for empty contracts")
			return nil, nil
		})
		if _, err := LoadEventRows(context.Background(), db, nil, 0); err == nil {
			t.Fatal("LoadEventRows succeeded without contracts")
		}
	})

	t.Run("query", func(t *testing.T) {
		boom := errors.New("query failed")
		db := openTXTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return nil, boom
		})
		if _, err := LoadEventRows(context.Background(), db, []string{"0x1"}, 0); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want query error", err)
		}
	})

	t.Run("scan", func(t *testing.T) {
		db := openTXTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return &txTestRows{
				columns: eventRowColumns(),
				values:  [][]driver.Value{{"not-a-block", int64(1), "0x1", "0x2", "topic", []byte{1}}},
			}, nil
		})
		if _, err := LoadEventRows(context.Background(), db, []string{"0x1"}, 0); err == nil {
			t.Fatal("LoadEventRows succeeded with an invalid row")
		}
	})

	t.Run("rows", func(t *testing.T) {
		boom := errors.New("rows failed")
		db := openTXTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return &txTestRows{columns: eventRowColumns(), terminalErr: boom}, nil
		})
		if _, err := LoadEventRows(context.Background(), db, []string{"0x1"}, 0); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want rows error", err)
		}
	})

	t.Run("cancellation while iterating", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		db := openTXTestDB(t, func(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
			return &txTestRows{
				columns: eventRowColumns(),
				values: [][]driver.Value{{
					int64(1), int64(1), "0x1", "0x2", "topic", []byte{1},
				}},
				onNext: cancel,
			}, nil
		})
		if _, err := LoadEventRows(ctx, db, []string{"0x1"}, 0); !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
	})
}
