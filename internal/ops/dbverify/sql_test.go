package dbverify_test

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

	"github.com/PredictionExplorer/augur-explorer/internal/ops/dbverify"
)

func TestSQLLoaderRequiresDatabase(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	var nilLoader *dbverify.SQLLoader

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "event records",
			call: func() error {
				_, err := nilLoader.LoadEventRecords(ctx, nil)
				return err
			},
		},
		{
			name: "transaction hashes",
			call: func() error {
				_, err := nilLoader.TransactionHashesFromEvents(ctx, nil)
				return err
			},
		},
		{
			name: "transactions",
			call: func() error {
				_, err := nilLoader.LoadTransactions(ctx, nil)
				return err
			},
		},
		{
			name: "block numbers",
			call: func() error {
				_, err := nilLoader.BlockNumbersFromEvents(ctx, nil)
				return err
			},
		},
		{
			name: "blocks",
			call: func() error {
				_, err := nilLoader.LoadBlocks(ctx, nil)
				return err
			},
		},
		{
			name: "event count",
			call: func() error {
				_, err := nilLoader.CountEventLogs(ctx, nil)
				return err
			},
		},
		{
			name: "detailed events",
			call: func() error {
				_, err := nilLoader.LoadDetailedEventLogs(ctx, nil, 0)
				return err
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if err := test.call(); err == nil || !strings.Contains(err.Error(), "database is required") {
				t.Fatalf("error = %v, want required-database error", err)
			}
		})
	}

	emptyLoader := &dbverify.SQLLoader{}
	if _, err := emptyLoader.LoadTransactions(ctx, []string{}); err == nil {
		t.Fatal("empty-filter load with a nil DB unexpectedly succeeded")
	}
	if _, err := dbverify.LoadRandomWalkContractAddressIDs(ctx, nil); err == nil {
		t.Fatal("LoadRandomWalkContractAddressIDs(nil) unexpectedly succeeded")
	}
}

func TestLoadRandomWalkContractAddressIDsBranches(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		db, script := newScriptedDB(t, queryExpectation{
			queryContains: "JOIN rw_contracts",
			columns:       []string{"address_id"},
			rows:          [][]driver.Value{{int64(4)}, {int64(9)}},
		})

		got, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)

		if err != nil {
			t.Fatalf("LoadRandomWalkContractAddressIDs() error = %v", err)
		}
		if want := []int64{4, 9}; !reflect.DeepEqual(got, want) {
			t.Errorf("ids = %v, want %v", got, want)
		}
		script.assertDone(t)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		db, script := newScriptedDB(t, queryExpectation{
			queryContains: "JOIN rw_contracts",
			columns:       []string{"address_id"},
		})

		_, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)

		if err == nil || !strings.Contains(err.Error(), "no contract addresses") {
			t.Fatalf("error = %v, want no-addresses error", err)
		}
		script.assertDone(t)
	})

	t.Run("query failure", func(t *testing.T) {
		t.Parallel()
		errQuery := errors.New("query failed")
		db, script := newScriptedDB(t, queryExpectation{
			queryContains: "JOIN rw_contracts",
			err:           errQuery,
		})

		_, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)

		if !errors.Is(err, errQuery) || !strings.Contains(err.Error(), "contract aids") {
			t.Fatalf("error = %v, want wrapped query failure", err)
		}
		script.assertDone(t)
	})

	t.Run("scan failure", func(t *testing.T) {
		t.Parallel()
		db, script := newScriptedDB(t, queryExpectation{
			queryContains: "JOIN rw_contracts",
			columns:       []string{"address_id"},
			rows:          [][]driver.Value{{"not-an-integer"}},
		})

		_, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)

		if err == nil || !strings.Contains(err.Error(), "scan contract aid") {
			t.Fatalf("error = %v, want scan failure", err)
		}
		script.assertDone(t)
	})

	t.Run("rows failure", func(t *testing.T) {
		t.Parallel()
		errRows := errors.New("rows failed")
		db, script := newScriptedDB(t, queryExpectation{
			queryContains: "JOIN rw_contracts",
			columns:       []string{"address_id"},
			rowsErr:       errRows,
		})

		_, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)

		if !errors.Is(err, errRows) || !strings.Contains(err.Error(), "contract aids") {
			t.Fatalf("error = %v, want wrapped rows failure", err)
		}
		script.assertDone(t)
	})
}

func TestSQLLoaderLoadEventRecordsBranches(t *testing.T) {
	t.Parallel()
	t.Run("all rows preserve duplicate multiplicity", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "ORDER BY e.log_rlp, e.block_num, t.tx_hash, e.id",
			argCount:      intPointer(0),
			columns:       []string{"block_num", "tx_hash", "log_rlp"},
			rows: [][]driver.Value{
				{int64(1), "tx-a", []byte{0xaa}},
				{int64(2), "tx-b", []byte{0xaa}},
			},
		})

		got, err := loader.LoadEventRecords(context.Background(), nil)

		if err != nil {
			t.Fatalf("LoadEventRecords() error = %v", err)
		}
		if len(got) != 2 || got["aa"].TxHash != "tx-a" || got["aa#2"].TxHash != "tx-b" {
			t.Errorf("records = %+v, want both duplicate occurrences", got)
		}
		script.assertDone(t)
	})

	for name, filter := range map[string][]int64{
		"empty filter": {},
		"value filter": {7, 8},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			loader, script := newScriptedLoader(t, queryExpectation{
				queryContains: "WHERE e.contract_aid = ANY($1)",
				argCount:      intPointer(1),
				columns:       []string{"block_num", "tx_hash", "log_rlp"},
			})

			got, err := loader.LoadEventRecords(context.Background(), filter)

			if err != nil {
				t.Fatalf("LoadEventRecords() error = %v", err)
			}
			if len(got) != 0 {
				t.Errorf("records = %+v, want empty", got)
			}
			script.assertDone(t)
		})
	}

	t.Run("query failure", func(t *testing.T) {
		t.Parallel()
		errQuery := errors.New("query failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "FROM evt_log e",
			err:           errQuery,
		})

		_, err := loader.LoadEventRecords(context.Background(), nil)

		if !errors.Is(err, errQuery) {
			t.Fatalf("error = %v, want query failure", err)
		}
		script.assertDone(t)
	})

	t.Run("scan failure", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "FROM evt_log e",
			columns:       []string{"block_num", "tx_hash", "log_rlp"},
			rows:          [][]driver.Value{{"bad-block", "tx", []byte{0x01}}},
		})

		_, err := loader.LoadEventRecords(context.Background(), nil)

		if err == nil || !strings.Contains(err.Error(), "scan event") {
			t.Fatalf("error = %v, want event scan failure", err)
		}
		script.assertDone(t)
	})

	t.Run("rows failure", func(t *testing.T) {
		t.Parallel()
		errRows := errors.New("rows failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "FROM evt_log e",
			columns:       []string{"block_num", "tx_hash", "log_rlp"},
			rowsErr:       errRows,
		})

		_, err := loader.LoadEventRecords(context.Background(), nil)

		if !errors.Is(err, errRows) {
			t.Fatalf("error = %v, want rows failure", err)
		}
		script.assertDone(t)
	})
}

func TestSQLLoaderTransactionHashesBranches(t *testing.T) {
	t.Parallel()
	for name, filter := range map[string][]int64{
		"nil filter":   nil,
		"empty filter": {},
		"value filter": {7},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			loader, script := newScriptedLoader(t, queryExpectation{
				queryContains: "SELECT DISTINCT t.tx_hash",
				argCount:      intPointer(1),
				columns:       []string{"tx_hash"},
				rows:          [][]driver.Value{{"a"}, {"b"}},
			})

			got, err := loader.TransactionHashesFromEvents(context.Background(), filter)

			if err != nil {
				t.Fatalf("TransactionHashesFromEvents() error = %v", err)
			}
			if want := []string{"a", "b"}; !reflect.DeepEqual(got, want) {
				t.Errorf("hashes = %v, want %v", got, want)
			}
			script.assertDone(t)
		})
	}

	runSingleColumnLoaderErrors(
		t,
		"transaction hashes",
		"SELECT DISTINCT t.tx_hash",
		"tx_hash",
		nil,
		func(loader *dbverify.SQLLoader) error {
			_, err := loader.TransactionHashesFromEvents(context.Background(), []int64{1})
			return err
		},
		"scan tx_hash",
	)
}

func TestSQLLoaderLoadTransactionsBranches(t *testing.T) {
	t.Parallel()
	t.Run("empty filter returns without querying", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t)

		got, err := loader.LoadTransactions(context.Background(), []string{})

		if err != nil || len(got) != 0 {
			t.Fatalf("LoadTransactions(empty) = %+v, %v", got, err)
		}
		script.assertDone(t)
	})

	for name, test := range map[string]struct {
		filter        []string
		queryContains string
		argCount      int
	}{
		"all": {
			filter:        nil,
			queryContains: "FROM transaction ORDER BY tx_hash",
			argCount:      0,
		},
		"filtered": {
			filter:        []string{"tx"},
			queryContains: "WHERE tx_hash = ANY($1)",
			argCount:      1,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			loader, script := newScriptedLoader(t, queryExpectation{
				queryContains: test.queryContains,
				argCount:      intPointer(test.argCount),
				columns:       []string{"block_num", "tx_hash", "gas_used", "num_logs"},
				rows:          [][]driver.Value{{int64(1), "tx", int64(2), int64(3)}},
			})

			got, err := loader.LoadTransactions(context.Background(), test.filter)

			if err != nil {
				t.Fatalf("LoadTransactions() error = %v", err)
			}
			want := dbverify.TransactionRecord{BlockNum: 1, TxHash: "tx", GasUsed: 2, NumLogs: 3}
			if got["tx"] != want {
				t.Errorf("transaction = %+v, want %+v", got["tx"], want)
			}
			script.assertDone(t)
		})
	}

	runRecordLoaderErrors(
		t,
		"transactions",
		"FROM transaction",
		[]string{"block_num", "tx_hash", "gas_used", "num_logs"},
		[]driver.Value{"bad-block", "tx", int64(2), int64(3)},
		func(loader *dbverify.SQLLoader) error {
			_, err := loader.LoadTransactions(context.Background(), nil)
			return err
		},
		"scan transaction",
	)
}

func TestSQLLoaderBlockNumbersBranches(t *testing.T) {
	t.Parallel()
	for name, filter := range map[string][]int64{
		"nil filter":   nil,
		"empty filter": {},
		"value filter": {7},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			loader, script := newScriptedLoader(t, queryExpectation{
				queryContains: "SELECT DISTINCT t.block_num",
				argCount:      intPointer(1),
				columns:       []string{"block_num"},
				rows:          [][]driver.Value{{int64(1)}, {int64(2)}},
			})

			got, err := loader.BlockNumbersFromEvents(context.Background(), filter)

			if err != nil {
				t.Fatalf("BlockNumbersFromEvents() error = %v", err)
			}
			if want := []int64{1, 2}; !reflect.DeepEqual(got, want) {
				t.Errorf("numbers = %v, want %v", got, want)
			}
			script.assertDone(t)
		})
	}

	runSingleColumnLoaderErrors(
		t,
		"block numbers",
		"SELECT DISTINCT t.block_num",
		"block_num",
		"bad-integer",
		func(loader *dbverify.SQLLoader) error {
			_, err := loader.BlockNumbersFromEvents(context.Background(), []int64{1})
			return err
		},
		"scan block_num",
	)
}

func TestSQLLoaderLoadBlocksBranches(t *testing.T) {
	t.Parallel()
	t.Run("empty filter returns without querying", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t)

		got, err := loader.LoadBlocks(context.Background(), []int64{})

		if err != nil || len(got) != 0 {
			t.Fatalf("LoadBlocks(empty) = %+v, %v", got, err)
		}
		script.assertDone(t)
	})

	for name, test := range map[string]struct {
		filter        []int64
		queryContains string
		argCount      int
	}{
		"all": {
			filter:        nil,
			queryContains: "FROM block ORDER BY block_num",
			argCount:      0,
		},
		"filtered": {
			filter:        []int64{1},
			queryContains: "WHERE block_num = ANY($1)",
			argCount:      1,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			loader, script := newScriptedLoader(t, queryExpectation{
				queryContains: test.queryContains,
				argCount:      intPointer(test.argCount),
				columns:       []string{"block_num", "block_hash", "parent_hash", "num_tx"},
				rows:          [][]driver.Value{{int64(1), "block", "parent", int64(2)}},
			})

			got, err := loader.LoadBlocks(context.Background(), test.filter)

			if err != nil {
				t.Fatalf("LoadBlocks() error = %v", err)
			}
			want := dbverify.BlockRecord{BlockNum: 1, BlockHash: "block", ParentHash: "parent", NumTx: 2}
			if got["block"] != want {
				t.Errorf("block = %+v, want %+v", got["block"], want)
			}
			script.assertDone(t)
		})
	}

	runRecordLoaderErrors(
		t,
		"blocks",
		"FROM block",
		[]string{"block_num", "block_hash", "parent_hash", "num_tx"},
		[]driver.Value{"bad-block", "block", "parent", int64(2)},
		func(loader *dbverify.SQLLoader) error {
			_, err := loader.LoadBlocks(context.Background(), nil)
			return err
		},
		"scan block",
	)
}

func TestSQLLoaderCountEventLogsBranches(t *testing.T) {
	t.Parallel()
	for name, test := range map[string]struct {
		filter        []int64
		queryContains string
		argCount      int
	}{
		"all": {
			filter:        nil,
			queryContains: "SELECT COUNT(*) FROM evt_log",
			argCount:      0,
		},
		"empty filter": {
			filter:        []int64{},
			queryContains: "WHERE contract_aid = ANY($1)",
			argCount:      1,
		},
		"value filter": {
			filter:        []int64{7},
			queryContains: "WHERE contract_aid = ANY($1)",
			argCount:      1,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			loader, script := newScriptedLoader(t, queryExpectation{
				queryContains: test.queryContains,
				argCount:      intPointer(test.argCount),
				columns:       []string{"count"},
				rows:          [][]driver.Value{{int64(12)}},
			})

			got, err := loader.CountEventLogs(context.Background(), test.filter)

			if err != nil || got != 12 {
				t.Fatalf("CountEventLogs() = %d, %v; want 12, nil", got, err)
			}
			script.assertDone(t)
		})
	}

	t.Run("query failure", func(t *testing.T) {
		t.Parallel()
		errQuery := errors.New("query failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "SELECT COUNT(*)",
			err:           errQuery,
		})

		_, err := loader.CountEventLogs(context.Background(), nil)

		if !errors.Is(err, errQuery) {
			t.Fatalf("error = %v, want query failure", err)
		}
		script.assertDone(t)
	})

	t.Run("scan failure", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "SELECT COUNT(*)",
			columns:       []string{"count"},
			rows:          [][]driver.Value{{"bad-count"}},
		})

		_, err := loader.CountEventLogs(context.Background(), nil)

		if err == nil {
			t.Fatal("CountEventLogs() scan unexpectedly succeeded")
		}
		script.assertDone(t)
	})
}

func TestSQLLoaderDetailedEventLogsBranches(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		filter        []int64
		limit         int
		queryContains []string
		argCount      int
	}{
		{
			name:          "all without limit",
			filter:        nil,
			limit:         0,
			queryContains: []string{"ORDER BY e.block_num, e.id"},
			argCount:      0,
		},
		{
			name:          "empty filter without limit",
			filter:        []int64{},
			limit:         -1,
			queryContains: []string{"WHERE e.contract_aid = ANY($1)"},
			argCount:      1,
		},
		{
			name:          "all with limit",
			filter:        nil,
			limit:         3,
			queryContains: []string{"LIMIT $1"},
			argCount:      1,
		},
		{
			name:          "filtered with limit",
			filter:        []int64{7},
			limit:         3,
			queryContains: []string{"WHERE e.contract_aid = ANY($1)", "LIMIT $2"},
			argCount:      2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			loader, script := newScriptedLoader(t, queryExpectation{
				queryContains: test.queryContains[0],
				queryParts:    test.queryContains,
				argCount:      intPointer(test.argCount),
				columns:       []string{"block_num", "tx_hash", "addr", "topic0_sig", "log_rlp"},
				rows: [][]driver.Value{{
					int64(1),
					"tx",
					"contract",
					"topic",
					[]byte{0xaa},
				}},
			})

			got, err := loader.LoadDetailedEventLogs(context.Background(), test.filter, test.limit)

			if err != nil {
				t.Fatalf("LoadDetailedEventLogs() error = %v", err)
			}
			want := []dbverify.EventLogRecord{{
				BlockNum:        1,
				TxHash:          "tx",
				ContractAddress: "contract",
				Topic0Sig:       "topic",
				LogRLP:          []byte{0xaa},
			}}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("events = %+v, want %+v", got, want)
			}
			script.assertDone(t)
		})
	}

	runRecordLoaderErrors(
		t,
		"detailed events",
		"JOIN address a",
		[]string{"block_num", "tx_hash", "addr", "topic0_sig", "log_rlp"},
		[]driver.Value{"bad-block", "tx", "contract", "topic", []byte{0xaa}},
		func(loader *dbverify.SQLLoader) error {
			_, err := loader.LoadDetailedEventLogs(context.Background(), nil, 0)
			return err
		},
		"scan event",
	)
}

func runSingleColumnLoaderErrors(
	t *testing.T,
	name string,
	queryContains string,
	column string,
	badValue driver.Value,
	call func(*dbverify.SQLLoader) error,
	scanMessage string,
) {
	t.Helper()
	t.Run(name+" query failure", func(t *testing.T) {
		t.Parallel()
		errQuery := errors.New("query failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			err:           errQuery,
		})

		err := call(loader)

		if !errors.Is(err, errQuery) {
			t.Fatalf("error = %v, want query failure", err)
		}
		script.assertDone(t)
	})
	t.Run(name+" scan failure", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			columns:       []string{column},
			rows:          [][]driver.Value{{badValue}},
		})

		err := call(loader)

		if err == nil || !strings.Contains(err.Error(), scanMessage) {
			t.Fatalf("error = %v, want %q scan failure", err, scanMessage)
		}
		script.assertDone(t)
	})
	t.Run(name+" rows failure", func(t *testing.T) {
		t.Parallel()
		errRows := errors.New("rows failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			columns:       []string{column},
			rowsErr:       errRows,
		})

		err := call(loader)

		if !errors.Is(err, errRows) {
			t.Fatalf("error = %v, want rows failure", err)
		}
		script.assertDone(t)
	})
}

func runRecordLoaderErrors(
	t *testing.T,
	name string,
	queryContains string,
	columns []string,
	badRow []driver.Value,
	call func(*dbverify.SQLLoader) error,
	scanMessage string,
) {
	t.Helper()
	t.Run(name+" query failure", func(t *testing.T) {
		t.Parallel()
		errQuery := errors.New("query failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			err:           errQuery,
		})

		err := call(loader)

		if !errors.Is(err, errQuery) {
			t.Fatalf("error = %v, want query failure", err)
		}
		script.assertDone(t)
	})
	t.Run(name+" scan failure", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			columns:       columns,
			rows:          [][]driver.Value{badRow},
		})

		err := call(loader)

		if err == nil || !strings.Contains(err.Error(), scanMessage) {
			t.Fatalf("error = %v, want %q scan failure", err, scanMessage)
		}
		script.assertDone(t)
	})
	t.Run(name+" rows failure", func(t *testing.T) {
		t.Parallel()
		errRows := errors.New("rows failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			columns:       columns,
			rowsErr:       errRows,
		})

		err := call(loader)

		if !errors.Is(err, errRows) {
			t.Fatalf("error = %v, want rows failure", err)
		}
		script.assertDone(t)
	})
}

const scriptedDriverName = "dbverify-scripted"

var (
	scriptedDatabases  sync.Map
	scriptedDBSequence atomic.Uint64
)

func init() {
	sql.Register(scriptedDriverName, scriptedDriver{})
}

type queryExpectation struct {
	queryContains string
	queryParts    []string
	argCount      *int
	columns       []string
	rows          [][]driver.Value
	rowsErr       error
	err           error
}

type queryScript struct {
	mu           sync.Mutex
	expectations []queryExpectation
	validation   []string
}

func newScriptedLoader(t *testing.T, expectations ...queryExpectation) (*dbverify.SQLLoader, *queryScript) {
	t.Helper()
	db, script := newScriptedDB(t, expectations...)
	return &dbverify.SQLLoader{DB: db}, script
}

func newScriptedDB(t *testing.T, expectations ...queryExpectation) (*sql.DB, *queryScript) {
	t.Helper()
	name := fmt.Sprintf("dbverify-script-%d", scriptedDBSequence.Add(1))
	script := &queryScript{expectations: append([]queryExpectation(nil), expectations...)}
	scriptedDatabases.Store(name, script)
	db, err := sql.Open(scriptedDriverName, name)
	if err != nil {
		t.Fatalf("opening scripted database: %v", err)
	}
	db.SetMaxOpenConns(1)
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("closing scripted database: %v", err)
		}
		scriptedDatabases.Delete(name)
	})
	return db, script
}

func (s *queryScript) next(query string, args []driver.NamedValue) (driver.Rows, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.expectations) == 0 {
		return nil, errors.New("unexpected query")
	}
	expectation := s.expectations[0]
	s.expectations = s.expectations[1:]
	normalized := strings.Join(strings.Fields(query), " ")
	if expectation.queryContains != "" && !strings.Contains(normalized, expectation.queryContains) {
		s.validation = append(
			s.validation,
			fmt.Sprintf("query %q does not contain %q", normalized, expectation.queryContains),
		)
	}
	for _, part := range expectation.queryParts {
		if !strings.Contains(normalized, part) {
			s.validation = append(s.validation, fmt.Sprintf("query %q does not contain %q", normalized, part))
		}
	}
	if expectation.argCount != nil && len(args) != *expectation.argCount {
		s.validation = append(
			s.validation,
			fmt.Sprintf("query %q has %d args, want %d", normalized, len(args), *expectation.argCount),
		)
	}
	if expectation.err != nil {
		return nil, expectation.err
	}
	return &scriptedRows{
		columns: append([]string(nil), expectation.columns...),
		rows:    cloneDriverRows(expectation.rows),
		rowsErr: expectation.rowsErr,
	}, nil
}

func (s *queryScript) assertDone(t *testing.T) {
	t.Helper()
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.expectations) != 0 {
		t.Errorf("%d scripted queries were not executed", len(s.expectations))
	}
	for _, message := range s.validation {
		t.Error(message)
	}
}

type scriptedDriver struct{}

func (scriptedDriver) Open(name string) (driver.Conn, error) {
	value, ok := scriptedDatabases.Load(name)
	if !ok {
		return nil, fmt.Errorf("scripted database %q is not registered", name)
	}
	script, ok := value.(*queryScript)
	if !ok {
		return nil, fmt.Errorf("scripted database %q has invalid state", name)
	}
	return &scriptedConn{script: script}, nil
}

type scriptedConn struct {
	script *queryScript
}

func (c *scriptedConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("scripted prepared statements are unsupported")
}

func (c *scriptedConn) Close() error {
	return nil
}

func (c *scriptedConn) Begin() (driver.Tx, error) {
	return nil, errors.New("scripted transactions are unsupported")
}

func (c *scriptedConn) QueryContext(
	_ context.Context,
	query string,
	args []driver.NamedValue,
) (driver.Rows, error) {
	return c.script.next(query, args)
}

type scriptedRows struct {
	columns []string
	rows    [][]driver.Value
	rowsErr error
	index   int
}

func (r *scriptedRows) Columns() []string {
	return r.columns
}

func (r *scriptedRows) Close() error {
	return nil
}

func (r *scriptedRows) Next(destination []driver.Value) error {
	if r.index < len(r.rows) {
		copy(destination, r.rows[r.index])
		r.index++
		return nil
	}
	if r.rowsErr != nil {
		err := r.rowsErr
		r.rowsErr = nil
		return err
	}
	return io.EOF
}

func cloneDriverRows(rows [][]driver.Value) [][]driver.Value {
	cloned := make([][]driver.Value, len(rows))
	for index, row := range rows {
		cloned[index] = append([]driver.Value(nil), row...)
	}
	return cloned
}

func intPointer(value int) *int {
	return &value
}
