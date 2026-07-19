package txcollector

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type txQueryFunc func(ctx context.Context, query string, args []any) (pgx.Rows, error)

// txQuerier adapts a per-query callback to the pgx Querier seam.
type txQuerier struct {
	query txQueryFunc
}

func (q txQuerier) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return q.query(ctx, query, args)
}

// txTestRows is a pgx.Rows fake with per-row hooks and close tracking.
type txTestRows struct {
	mu          sync.Mutex
	values      [][]any
	next        int
	current     []any
	terminalErr error
	onNext      func()
	closed      bool
	err         error
}

func (r *txTestRows) Next() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.next < len(r.values) {
		r.current = r.values[r.next]
		r.next++
		if r.onNext != nil {
			r.onNext()
		}
		return true
	}
	if r.terminalErr != nil {
		r.err = r.terminalErr
		r.terminalErr = nil
	}
	return false
}

func (r *txTestRows) Scan(dest ...any) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	values := r.current
	if len(dest) != len(values) {
		return errors.New("txcollector test rows: destination arity mismatch")
	}
	for i, value := range values {
		switch d := dest[i].(type) {
		case *int64:
			v, ok := value.(int64)
			if !ok {
				return errors.New("txcollector test rows: not an int64")
			}
			*d = v
		case *int:
			v, ok := value.(int64)
			if !ok {
				return errors.New("txcollector test rows: not an int")
			}
			*d = int(v)
		case *string:
			v, ok := value.(string)
			if !ok {
				return errors.New("txcollector test rows: not a string")
			}
			*d = v
		case *[]byte:
			v, ok := value.([]byte)
			if !ok {
				return errors.New("txcollector test rows: not bytes")
			}
			*d = v
		default:
			return errors.New("txcollector test rows: unsupported destination")
		}
	}
	return nil
}

func (r *txTestRows) Err() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.err
}

func (r *txTestRows) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closed = true
}

func (r *txTestRows) isClosed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closed
}

func (*txTestRows) CommandTag() pgconn.CommandTag                { return pgconn.CommandTag{} }
func (*txTestRows) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (*txTestRows) RawValues() [][]byte                          { return nil }
func (*txTestRows) Values() ([]any, error)                       { return nil, nil }
func (*txTestRows) Conn() *pgx.Conn                              { return nil }

func TestLoadEventRowsSuccessAndQueryArguments(t *testing.T) {
	txHash := common.HexToHash("0xabcd")
	address := common.HexToAddress("0xabcdef0000000000000000000000000000001234")
	rows := &txTestRows{
		values: [][]any{{
			int64(125),
			int64(7),
			"0x" + strings.ToUpper(txHash.Hex()[2:]),
			strings.ToLower(address.Hex()),
			"deadbeef",
			[]byte{1, 2, 3},
		}},
	}
	contracts := []string{address.Hex()}
	db := txQuerier{query: func(_ context.Context, query string, args []any) (pgx.Rows, error) {
		if !strings.Contains(query, "a.addr = ANY($1)") ||
			!strings.Contains(query, "e.block_num >= $2") ||
			!strings.Contains(query, "ORDER BY t.tx_hash, e.log_index") {
			t.Fatalf("query = %q", query)
		}
		if len(args) != 2 {
			t.Fatalf("args = %v, want two", args)
		}
		if got, ok := args[0].([]string); !ok || !reflect.DeepEqual(got, contracts) {
			t.Fatalf("contract argument = %#v", args[0])
		}
		if args[1] != uint64(120) {
			t.Fatalf("from-block argument = %#v, want uint64(120)", args[1])
		}
		return rows, nil
	}}

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
		db := txQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			t.Fatal("query called for empty contracts")
			return nil, nil
		}}
		if _, err := LoadEventRows(context.Background(), db, nil, 0); err == nil {
			t.Fatal("LoadEventRows succeeded without contracts")
		}
	})

	t.Run("query", func(t *testing.T) {
		boom := errors.New("query failed")
		db := txQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return nil, boom
		}}
		if _, err := LoadEventRows(context.Background(), db, []string{"0x1"}, 0); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want query error", err)
		}
	})

	t.Run("scan", func(t *testing.T) {
		db := txQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return &txTestRows{
				values: [][]any{{"not-a-block", int64(1), "0x1", "0x2", "topic", []byte{1}}},
			}, nil
		}}
		if _, err := LoadEventRows(context.Background(), db, []string{"0x1"}, 0); err == nil {
			t.Fatal("LoadEventRows succeeded with an invalid row")
		}
	})

	t.Run("rows", func(t *testing.T) {
		boom := errors.New("rows failed")
		db := txQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return &txTestRows{terminalErr: boom}, nil
		}}
		if _, err := LoadEventRows(context.Background(), db, []string{"0x1"}, 0); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want rows error", err)
		}
	})

	t.Run("cancellation while iterating", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		db := txQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return &txTestRows{
				values: [][]any{{
					int64(1), int64(1), "0x1", "0x2", "topic", []byte{1},
				}},
				onNext: cancel,
			}, nil
		}}
		if _, err := LoadEventRows(ctx, db, []string{"0x1"}, 0); !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
	})
}
