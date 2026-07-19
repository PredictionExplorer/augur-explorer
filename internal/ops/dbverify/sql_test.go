package dbverify_test

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/dbverify"
	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
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
			rows:          [][]any{{int64(4)}, {int64(9)}},
		})

		got, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)
		if err != nil {
			t.Fatalf("LoadRandomWalkContractAddressIDs() error = %v", err)
		}
		if want := []int64{4, 9}; !reflect.DeepEqual(got, want) {
			t.Errorf("ids = %v, want %v", got, want)
		}
		script.AssertDone(t)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		db, script := newScriptedDB(t, queryExpectation{
			queryContains: "JOIN rw_contracts",
		})

		_, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)

		if err == nil || !strings.Contains(err.Error(), "no contract addresses") {
			t.Fatalf("error = %v, want no-addresses error", err)
		}
		script.AssertDone(t)
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
		script.AssertDone(t)
	})

	t.Run("scan failure", func(t *testing.T) {
		t.Parallel()
		db, script := newScriptedDB(t, queryExpectation{
			queryContains: "JOIN rw_contracts",
			rows:          [][]any{{"not-an-integer"}},
		})

		_, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)

		if err == nil || !strings.Contains(err.Error(), "scan contract aid") {
			t.Fatalf("error = %v, want scan failure", err)
		}
		script.AssertDone(t)
	})

	t.Run("rows failure", func(t *testing.T) {
		t.Parallel()
		errRows := errors.New("rows failed")
		db, script := newScriptedDB(t, queryExpectation{
			queryContains: "JOIN rw_contracts",
			rowsErr:       errRows,
		})

		_, err := dbverify.LoadRandomWalkContractAddressIDs(context.Background(), db)

		if !errors.Is(err, errRows) || !strings.Contains(err.Error(), "contract aids") {
			t.Fatalf("error = %v, want wrapped rows failure", err)
		}
		script.AssertDone(t)
	})
}

func TestSQLLoaderLoadEventRecordsBranches(t *testing.T) {
	t.Parallel()
	t.Run("all rows preserve duplicate multiplicity", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "ORDER BY e.log_rlp, e.block_num, t.tx_hash, e.id",
			argCount:      new(0),
			rows: [][]any{
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
		script.AssertDone(t)
	})

	for name, filter := range map[string][]int64{
		"empty filter": {},
		"value filter": {7, 8},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			loader, script := newScriptedLoader(t, queryExpectation{
				queryContains: "WHERE e.contract_aid = ANY($1)",
				argCount:      new(1),
			})

			got, err := loader.LoadEventRecords(context.Background(), filter)
			if err != nil {
				t.Fatalf("LoadEventRecords() error = %v", err)
			}
			if len(got) != 0 {
				t.Errorf("records = %+v, want empty", got)
			}
			script.AssertDone(t)
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
		script.AssertDone(t)
	})

	t.Run("scan failure", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "FROM evt_log e",
			rows:          [][]any{{"bad-block", "tx", []byte{0x01}}},
		})

		_, err := loader.LoadEventRecords(context.Background(), nil)

		if err == nil || !strings.Contains(err.Error(), "scan event") {
			t.Fatalf("error = %v, want event scan failure", err)
		}
		script.AssertDone(t)
	})

	t.Run("rows failure", func(t *testing.T) {
		t.Parallel()
		errRows := errors.New("rows failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "FROM evt_log e",
			rowsErr:       errRows,
		})

		_, err := loader.LoadEventRecords(context.Background(), nil)

		if !errors.Is(err, errRows) {
			t.Fatalf("error = %v, want rows failure", err)
		}
		script.AssertDone(t)
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
				argCount:      new(1),
				rows:          [][]any{{"a"}, {"b"}},
			})

			got, err := loader.TransactionHashesFromEvents(context.Background(), filter)
			if err != nil {
				t.Fatalf("TransactionHashesFromEvents() error = %v", err)
			}
			if want := []string{"a", "b"}; !reflect.DeepEqual(got, want) {
				t.Errorf("hashes = %v, want %v", got, want)
			}
			script.AssertDone(t)
		})
	}

	runSingleColumnLoaderErrors(
		t,
		"transaction hashes",
		"SELECT DISTINCT t.tx_hash",
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
		script.AssertDone(t)
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
				argCount:      new(test.argCount),
				rows:          [][]any{{int64(1), "tx", int64(2), int64(3)}},
			})

			got, err := loader.LoadTransactions(context.Background(), test.filter)
			if err != nil {
				t.Fatalf("LoadTransactions() error = %v", err)
			}
			want := dbverify.TransactionRecord{BlockNum: 1, TxHash: "tx", GasUsed: 2, NumLogs: 3}
			if got["tx"] != want {
				t.Errorf("transaction = %+v, want %+v", got["tx"], want)
			}
			script.AssertDone(t)
		})
	}

	runRecordLoaderErrors(
		t,
		"transactions",
		"FROM transaction",
		[]any{"bad-block", "tx", int64(2), int64(3)},
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
				argCount:      new(1),
				rows:          [][]any{{int64(1)}, {int64(2)}},
			})

			got, err := loader.BlockNumbersFromEvents(context.Background(), filter)
			if err != nil {
				t.Fatalf("BlockNumbersFromEvents() error = %v", err)
			}
			if want := []int64{1, 2}; !reflect.DeepEqual(got, want) {
				t.Errorf("numbers = %v, want %v", got, want)
			}
			script.AssertDone(t)
		})
	}

	runSingleColumnLoaderErrors(
		t,
		"block numbers",
		"SELECT DISTINCT t.block_num",
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
		script.AssertDone(t)
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
				argCount:      new(test.argCount),
				rows:          [][]any{{int64(1), "block", "parent", int64(2)}},
			})

			got, err := loader.LoadBlocks(context.Background(), test.filter)
			if err != nil {
				t.Fatalf("LoadBlocks() error = %v", err)
			}
			want := dbverify.BlockRecord{BlockNum: 1, BlockHash: "block", ParentHash: "parent", NumTx: 2}
			if got["block"] != want {
				t.Errorf("block = %+v, want %+v", got["block"], want)
			}
			script.AssertDone(t)
		})
	}

	runRecordLoaderErrors(
		t,
		"blocks",
		"FROM block",
		[]any{"bad-block", "block", "parent", int64(2)},
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
				argCount:      new(test.argCount),
				rows:          [][]any{{int64(12)}},
			})

			got, err := loader.CountEventLogs(context.Background(), test.filter)

			if err != nil || got != 12 {
				t.Fatalf("CountEventLogs() = %d, %v; want 12, nil", got, err)
			}
			script.AssertDone(t)
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
		script.AssertDone(t)
	})

	t.Run("scan failure", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: "SELECT COUNT(*)",
			rows:          [][]any{{"bad-count"}},
		})

		_, err := loader.CountEventLogs(context.Background(), nil)

		if err == nil {
			t.Fatal("CountEventLogs() scan unexpectedly succeeded")
		}
		script.AssertDone(t)
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
				argCount:      new(test.argCount),
				rows: [][]any{{
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
			script.AssertDone(t)
		})
	}

	runRecordLoaderErrors(
		t,
		"detailed events",
		"JOIN address a",
		[]any{"bad-block", "tx", "contract", "topic", []byte{0xaa}},
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
	badValue any,
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
		script.AssertDone(t)
	})
	t.Run(name+" scan failure", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			rows:          [][]any{{badValue}},
		})

		err := call(loader)

		if err == nil || !strings.Contains(err.Error(), scanMessage) {
			t.Fatalf("error = %v, want %q scan failure", err, scanMessage)
		}
		script.AssertDone(t)
	})
	t.Run(name+" rows failure", func(t *testing.T) {
		t.Parallel()
		errRows := errors.New("rows failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			rowsErr:       errRows,
		})

		err := call(loader)

		if !errors.Is(err, errRows) {
			t.Fatalf("error = %v, want rows failure", err)
		}
		script.AssertDone(t)
	})
}

func runRecordLoaderErrors(
	t *testing.T,
	name string,
	queryContains string,
	badRow []any,
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
		script.AssertDone(t)
	})
	t.Run(name+" scan failure", func(t *testing.T) {
		t.Parallel()
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			rows:          [][]any{badRow},
		})

		err := call(loader)

		if err == nil || !strings.Contains(err.Error(), scanMessage) {
			t.Fatalf("error = %v, want %q scan failure", err, scanMessage)
		}
		script.AssertDone(t)
	})
	t.Run(name+" rows failure", func(t *testing.T) {
		t.Parallel()
		errRows := errors.New("rows failed")
		loader, script := newScriptedLoader(t, queryExpectation{
			queryContains: queryContains,
			rowsErr:       errRows,
		})

		err := call(loader)

		if !errors.Is(err, errRows) {
			t.Fatalf("error = %v, want rows failure", err)
		}
		script.AssertDone(t)
	})
}

type queryExpectation struct {
	queryContains string
	queryParts    []string
	argCount      *int
	rows          [][]any
	rowsErr       error
	err           error
}

func (e queryExpectation) op() testutil.PgxOp {
	return testutil.PgxOp{
		Kind:     "query",
		Contains: e.queryContains,
		Parts:    e.queryParts,
		ArgCount: e.argCount,
		Rows:     e.rows,
		RowsErr:  e.rowsErr,
		Err:      e.err,
	}
}

func newScriptedLoader(t *testing.T, expectations ...queryExpectation) (*dbverify.SQLLoader, *testutil.ScriptedPgx) {
	t.Helper()
	db, script := newScriptedDB(t, expectations...)
	return &dbverify.SQLLoader{DB: db}, script
}

func newScriptedDB(t *testing.T, expectations ...queryExpectation) (dbverify.Querier, *testutil.ScriptedPgx) {
	t.Helper()
	ops := make([]testutil.PgxOp, 0, len(expectations))
	for _, expectation := range expectations {
		ops = append(ops, expectation.op())
	}
	script := testutil.NewScriptedPgx(ops...)
	return script, script
}
