package cstscan

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

type cstQueryFunc func(ctx context.Context, query string, args []any) (pgx.Rows, error)

// cstQuerier adapts a per-query callback to the pgx Querier seam.
type cstQuerier struct {
	query cstQueryFunc
}

func (q cstQuerier) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return q.query(ctx, query, args)
}

// cstTestRows is a pgx.Rows fake with per-row hooks and close tracking.
type cstTestRows struct {
	mu          sync.Mutex
	values      [][]any
	next        int
	current     []any
	terminalErr error
	onNext      func()
	closed      bool
	err         error
}

func (r *cstTestRows) Next() bool {
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

func (r *cstTestRows) Scan(dest ...any) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(dest) != len(r.current) {
		return errors.New("cstscan test rows: destination arity mismatch")
	}
	for i, value := range r.current {
		switch d := dest[i].(type) {
		case *string:
			v, ok := value.(string)
			if !ok {
				return errors.New("cstscan test rows: not a string")
			}
			*d = v
		case *int64:
			v, ok := value.(int64)
			if !ok {
				return errors.New("cstscan test rows: not an int64")
			}
			*d = v
		default:
			return errors.New("cstscan test rows: unsupported destination")
		}
	}
	return nil
}

func (r *cstTestRows) Err() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.err
}

func (r *cstTestRows) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closed = true
}

func (r *cstTestRows) isClosed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closed
}

func (*cstTestRows) CommandTag() pgconn.CommandTag                { return pgconn.CommandTag{} }
func (*cstTestRows) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (*cstTestRows) RawValues() [][]byte                          { return nil }
func (*cstTestRows) Values() ([]any, error)                       { return nil, nil }
func (*cstTestRows) Conn() *pgx.Conn                              { return nil }

func TestPostgresKeySourceSuccess(t *testing.T) {
	hash := common.HexToHash("0x1234")
	rows := &cstTestRows{
		values: [][]any{{"0x" + strings.ToUpper(hash.Hex()[2:]), int64(7)}},
	}
	db := cstQuerier{query: func(_ context.Context, query string, args []any) (pgx.Rows, error) {
		if !strings.Contains(query, "FROM cg_adm_cst_auclen_chg_div") {
			t.Fatalf("query = %q", query)
		}
		if len(args) != 0 {
			t.Fatalf("args = %v, want none", args)
		}
		return rows, nil
	}}

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
		db := cstQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return nil, boom
		}}
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background()); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want query error", err)
		}
	})

	t.Run("scan", func(t *testing.T) {
		db := cstQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return &cstTestRows{
				values: [][]any{{common.HexToHash("0x01").Hex(), "not-an-index"}},
			}, nil
		}}
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background()); err == nil ||
			!strings.Contains(err.Error(), "db scan") {
			t.Fatalf("error = %v, want scan error", err)
		}
	})

	t.Run("negative log index", func(t *testing.T) {
		db := cstQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return &cstTestRows{
				values: [][]any{{common.HexToHash("0x01").Hex(), int64(-1)}},
			}, nil
		}}
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background()); err == nil ||
			!strings.Contains(err.Error(), "negative log index") {
			t.Fatalf("error = %v, want negative-index error", err)
		}
	})

	t.Run("rows", func(t *testing.T) {
		boom := errors.New("rows failed")
		db := cstQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return &cstTestRows{terminalErr: boom}, nil
		}}
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(context.Background()); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want rows error", err)
		}
	})

	t.Run("cancellation while iterating", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		db := cstQuerier{query: func(context.Context, string, []any) (pgx.Rows, error) {
			return &cstTestRows{
				values: [][]any{{common.HexToHash("0x01").Hex(), int64(1)}},
				onNext: cancel,
			}, nil
		}}
		if _, err := (PostgresKeySource{DB: db}).LoadKeys(ctx); !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
	})
}
