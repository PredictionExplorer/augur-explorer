package archive

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

type recordingLogger struct {
	mu    sync.Mutex
	lines []string
}

func (l *recordingLogger) Printf(format string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lines = append(l.lines, fmt.Sprintf(format, args...))
}

func (l *recordingLogger) Println(args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lines = append(l.lines, fmt.Sprintln(args...))
}

func (l *recordingLogger) contains(text string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, line := range l.lines {
		if strings.Contains(line, text) {
			return true
		}
	}
	return false
}

func TestExportValidationAndUtilities(t *testing.T) {
	if _, err := ExportProjects(context.Background(), nil, nil); err == nil {
		t.Fatal("ExportProjects accepted a nil exporter")
	}
	for name, exporter := range map[string]*SQLExporter{
		"nil receiver": nil,
		"nil source":   {Destination: &sql.DB{}},
		"nil dest":     {Source: &sql.DB{}},
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := exporter.ExportProject(context.Background(), ProjectRandomWalk); err == nil {
				t.Fatal("ExportProject accepted missing databases")
			}
		})
	}

	logger := &recordingLogger{}
	exporter := &SQLExporter{Logger: logger}
	exporter.printf("value=%d", 7)
	exporter.println("line")
	if !logger.contains("value=7") || !logger.contains("line") {
		t.Fatalf("logger lines = %v", logger.lines)
	}
	if got := exporter.batchSize(); got != DefaultExportBatchSize {
		t.Errorf("default batch = %d", got)
	}
	exporter.BatchSize = 3
	if got := exporter.batchSize(); got != 3 {
		t.Errorf("configured batch = %d", got)
	}
	if got := perSecond(10, time.Now().Add(time.Hour)); got != 0 {
		t.Errorf("future-start rate = %f", got)
	}
	if got := perSecond(10, time.Now().Add(-time.Second)); got <= 0 {
		t.Errorf("normal rate = %f", got)
	}
	if got := percentage(4, 0); got != 0 {
		t.Errorf("zero-total percentage = %f", got)
	}
	if got := percentage(1, 4); got != 25 {
		t.Errorf("percentage = %f", got)
	}
}

func contractSourceOps(after ...scriptOp) []scriptOp {
	ops := []scriptOp{
		queryOp("JOIN rw_contracts", []string{"address_id"}, []driver.Value{int64(8)}),
		queryOp("SELECT addr FROM address", []string{"addr"}, []driver.Value{"0x08"}),
	}
	return append(ops, after...)
}

func emptyEventSourceOps(after ...scriptOp) []scriptOp {
	ops := []scriptOp{
		queryOp("SELECT COUNT(*) FROM evt_log", []string{"count"}, []driver.Value{int64(0)}),
		queryOp("FROM evt_log e JOIN transaction", []string{
			"block_num", "id", "tx_id", "log_index", "tx_hash", "addr", "topic0_sig", "log_rlp",
		}),
	}
	return append(ops, after...)
}

func eventDestinationOps(after ...scriptOp) []scriptOp {
	ops := []scriptOp{
		queryOp("COALESCE(MAX(evt_id)", []string{"max"}, []driver.Value{int64(0)}),
		prepareOp("INSERT INTO arch_evtlog"),
	}
	return append(ops, after...)
}

func TestSQLExporterProjectStageErrors(t *testing.T) {
	sentinel := errors.New("stage failed")
	tests := []struct {
		name    string
		source  []scriptOp
		dest    []scriptOp
		want    string
		project string
	}{
		{
			name:    "contract loading",
			source:  []scriptOp{queryErrorOp("JOIN rw_contracts", sentinel)},
			want:    "contract aids",
			project: ProjectRandomWalk,
		},
		{
			name: "event export",
			source: contractSourceOps(
				queryErrorOp("SELECT COUNT(*) FROM evt_log", sentinel),
			),
			want:    "count events",
			project: ProjectRandomWalk,
		},
		{
			name:   "transaction export",
			source: contractSourceOps(emptyEventSourceOps()...),
			dest: eventDestinationOps(
				queryErrorOp("SELECT DISTINCT e.tx_hash", sentinel),
			),
			want:    "find missing transactions",
			project: ProjectRandomWalk,
		},
		{
			name:   "block export",
			source: contractSourceOps(emptyEventSourceOps()...),
			dest: eventDestinationOps(
				queryOp("SELECT DISTINCT e.tx_hash", []string{"tx_hash"}),
				queryErrorOp("SELECT DISTINCT tx.block_num", sentinel),
			),
			want:    "find missing blocks",
			project: ProjectRandomWalk,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exporter := &SQLExporter{
				Source:      openScriptDB(t, test.source...),
				Destination: openScriptDB(t, test.dest...),
			}
			_, err := exporter.ExportProject(context.Background(), test.project)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestEventLogResumeFloorErrorAndEmpty(t *testing.T) {
	floor, positions, err := EventLogResumeFloor(context.Background(), openScriptDB(t), nil)
	if err != nil || floor != 0 || len(positions) != 0 {
		t.Fatalf("empty resume = %d, %v, %v", floor, positions, err)
	}
	sentinel := errors.New("resume failed")
	_, positions, err = EventLogResumeFloor(
		context.Background(),
		openScriptDB(t,
			queryOp("COALESCE(MAX(evt_id)", []string{"max"}, []driver.Value{int64(10)}),
			queryErrorOp("COALESCE(MAX(evt_id)", sentinel),
		),
		[]string{"a", "b"},
	)
	if !errors.Is(err, sentinel) || len(positions) != 1 {
		t.Fatalf("partial resume positions = %v, error = %v", positions, err)
	}
}

func TestExportEventLogsSetupFailures(t *testing.T) {
	sentinel := errors.New("event setup failed")
	contracts := Contracts{AddressIDs: []int64{8}, Addresses: []string{"0x08"}}
	tests := []struct {
		name   string
		source []scriptOp
		dest   []scriptOp
		want   string
	}{
		{
			name:   "count",
			source: []scriptOp{queryErrorOp("SELECT COUNT(*) FROM evt_log", sentinel)},
			want:   "count events",
		},
		{
			name: "resume",
			source: []scriptOp{
				queryOp("SELECT COUNT(*) FROM evt_log", []string{"count"}, []driver.Value{int64(1)}),
			},
			dest: []scriptOp{queryErrorOp("COALESCE(MAX(evt_id)", sentinel)},
			want: "read resume position",
		},
		{
			name: "prepare",
			source: []scriptOp{
				queryOp("SELECT COUNT(*) FROM evt_log", []string{"count"}, []driver.Value{int64(1)}),
			},
			dest: []scriptOp{
				queryOp("COALESCE(MAX(evt_id)", []string{"max"}, []driver.Value{int64(0)}),
				prepareErrorOp("INSERT INTO arch_evtlog", sentinel),
			},
			want: "prepare arch_evtlog",
		},
		{
			name: "batch",
			source: []scriptOp{
				queryOp("SELECT COUNT(*) FROM evt_log", []string{"count"}, []driver.Value{int64(1)}),
				queryErrorOp("FROM evt_log e", sentinel),
			},
			dest: []scriptOp{
				queryOp("COALESCE(MAX(evt_id)", []string{"max"}, []driver.Value{int64(0)}),
				prepareOp("INSERT INTO arch_evtlog"),
			},
			want: "query evt_log batch",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exporter := &SQLExporter{
				Source:      openScriptDB(t, test.source...),
				Destination: openScriptDB(t, test.dest...),
			}
			_, err := exporter.exportEventLogs(context.Background(), contracts)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestExportLoopsHonorCancellation(t *testing.T) {
	t.Run("events", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		prepare := prepareOp("INSERT INTO arch_evtlog")
		prepare.onCall = cancel
		exporter := &SQLExporter{
			Source: openScriptDB(t,
				queryOp("SELECT COUNT(*) FROM evt_log", []string{"count"}, []driver.Value{int64(1)}),
			),
			Destination: openScriptDB(t,
				queryOp("COALESCE(MAX(evt_id)", []string{"max"}, []driver.Value{int64(0)}),
				prepare,
			),
		}
		_, err := exporter.exportEventLogs(ctx, Contracts{
			AddressIDs: []int64{8},
			Addresses:  []string{"0x08"},
		})
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("transactions", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		prepare := prepareOp("INSERT INTO arch_tx")
		prepare.onCall = cancel
		exporter := &SQLExporter{
			Source: openScriptDB(t),
			Destination: openScriptDB(t,
				queryOp("SELECT DISTINCT e.tx_hash", []string{"tx_hash"}, []driver.Value{"tx"}),
				prepare,
			),
		}
		_, err := exporter.exportTransactions(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("blocks", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		prepare := prepareOp("INSERT INTO arch_block")
		prepare.onCall = cancel
		exporter := &SQLExporter{
			Source: openScriptDB(t),
			Destination: openScriptDB(t,
				queryOp("SELECT DISTINCT tx.block_num", []string{"block_num"}, []driver.Value{int64(10)}),
				prepare,
			),
		}
		_, err := exporter.exportBlocks(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v", err)
		}
	})
}

func prepareStatement(t *testing.T, ops ...scriptOp) *sql.Stmt {
	t.Helper()
	db := openScriptDB(t, ops...)
	statement, err := db.PrepareContext(context.Background(), "INSERT test statement")
	if err != nil {
		t.Fatalf("preparing scripted statement: %v", err)
	}
	t.Cleanup(func() { _ = statement.Close() })
	return statement
}

func eventRow() []driver.Value {
	return []driver.Value{
		int64(10), int64(20), int64(30), int64(1),
		"0xtx", "0xcontract", "deadbeef", []byte{0x01},
	}
}

func TestExportEventLogBatchFailures(t *testing.T) {
	sentinel := errors.New("event batch failed")
	tests := []struct {
		name   string
		source scriptOp
		stmt   []scriptOp
		want   string
	}{
		{
			name:   "query",
			source: queryErrorOp("FROM evt_log e", sentinel),
			stmt:   []scriptOp{prepareOp("INSERT test statement")},
			want:   "query evt_log batch",
		},
		{
			name: "scan",
			source: queryOp("FROM evt_log e", []string{
				"block_num", "id", "tx_id", "log_index", "tx_hash", "addr", "topic0_sig", "log_rlp",
			}, []driver.Value{"bad", int64(2), int64(3), int64(0), "tx", "addr", "topic", []byte{1}}),
			stmt: []scriptOp{prepareOp("INSERT test statement")},
			want: "scan evt_log row",
		},
		{
			name: "insert",
			source: queryOp("FROM evt_log e", []string{
				"block_num", "id", "tx_id", "log_index", "tx_hash", "addr", "topic0_sig", "log_rlp",
			}, eventRow()),
			stmt: []scriptOp{
				prepareOp("INSERT test statement"),
				execErrorOp("INSERT test statement", sentinel),
			},
			want: "insert arch_evtlog",
		},
		{
			name: "iteration",
			source: scriptOp{
				kind:      "query",
				contains:  "FROM evt_log e",
				columns:   []string{"block_num", "id", "tx_id", "log_index", "tx_hash", "addr", "topic0_sig", "log_rlp"},
				nextErrAt: 0,
				nextErr:   sentinel,
			},
			stmt: []scriptOp{prepareOp("INSERT test statement")},
			want: "iterate evt_log batch",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			statement := prepareStatement(t, test.stmt...)
			exporter := &SQLExporter{Source: openScriptDB(t, test.source)}
			_, _, err := exporter.exportEventLogBatch(
				context.Background(),
				statement,
				[]int64{8},
				0,
				map[int64]struct{}{},
			)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestExportTransactionsFailures(t *testing.T) {
	sentinel := errors.New("transaction export failed")
	tests := []struct {
		name   string
		source []scriptOp
		dest   []scriptOp
		want   string
	}{
		{
			name: "missing query",
			dest: []scriptOp{queryErrorOp("SELECT DISTINCT e.tx_hash", sentinel)},
			want: "find missing transactions",
		},
		{
			name: "prepare",
			dest: []scriptOp{
				queryOp("SELECT DISTINCT e.tx_hash", []string{"tx_hash"}, []driver.Value{"tx"}),
				prepareErrorOp("INSERT INTO arch_tx", sentinel),
			},
			want: "prepare arch_tx",
		},
		{
			name:   "batch",
			source: []scriptOp{queryErrorOp("FROM transaction", sentinel)},
			dest: []scriptOp{
				queryOp("SELECT DISTINCT e.tx_hash", []string{"tx_hash"}, []driver.Value{"tx"}),
				prepareOp("INSERT INTO arch_tx"),
			},
			want: "query transaction batch",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exporter := &SQLExporter{
				Source:      openScriptDB(t, test.source...),
				Destination: openScriptDB(t, test.dest...),
				BatchSize:   1,
			}
			_, err := exporter.exportTransactions(context.Background())
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func transactionRow(input any) []driver.Value {
	return []driver.Value{
		int64(10), int64(1), int64(2), int64(21_000), int64(3), int64(4),
		false, "1", "2", "0xtx", input,
	}
}

func TestExportTransactionBatchBranches(t *testing.T) {
	sentinel := errors.New("transaction batch failed")
	columns := []string{
		"block_num", "from_aid", "to_aid", "gas_used", "tx_index", "num_logs",
		"ctrct_create", "value", "gas_price", "tx_hash", "input_sig",
	}
	tests := []struct {
		name   string
		source scriptOp
		stmt   []scriptOp
		want   string
		ok     bool
	}{
		{
			name:   "query",
			source: queryErrorOp("FROM transaction", sentinel),
			stmt:   []scriptOp{prepareOp("INSERT test statement")},
			want:   "query transaction batch",
		},
		{
			name:   "scan",
			source: queryOp("FROM transaction", columns, transactionRow(nil)[:3]),
			stmt:   []scriptOp{prepareOp("INSERT test statement")},
			want:   "scan transaction row",
		},
		{
			name:   "insert",
			source: queryOp("FROM transaction", columns, transactionRow("0x12345678")),
			stmt: []scriptOp{
				prepareOp("INSERT test statement"),
				execErrorOp("INSERT test statement", sentinel),
			},
			want: "insert arch_tx",
		},
		{
			name: "iteration",
			source: scriptOp{
				kind:      "query",
				contains:  "FROM transaction",
				columns:   columns,
				nextErrAt: 0,
				nextErr:   sentinel,
			},
			stmt: []scriptOp{prepareOp("INSERT test statement")},
			want: "iterate transaction batch",
		},
		{
			name:   "nullable input signature",
			source: queryOp("FROM transaction", columns, transactionRow(nil)),
			stmt: []scriptOp{
				prepareOp("INSERT test statement"),
				execOp("INSERT test statement", 1),
			},
			ok: true,
		},
		{
			name:   "valid input signature",
			source: queryOp("FROM transaction", columns, transactionRow("0x12345678")),
			stmt: []scriptOp{
				prepareOp("INSERT test statement"),
				execOp("INSERT test statement", 1),
			},
			ok: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			statement := prepareStatement(t, test.stmt...)
			exporter := &SQLExporter{Source: openScriptDB(t, test.source)}
			count, err := exporter.exportTransactionBatch(
				context.Background(),
				statement,
				[]string{"tx"},
				map[int64]struct{}{},
			)
			if test.ok {
				if err != nil || count != 1 {
					t.Fatalf("count/error = %d / %v", count, err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestExportBlocksFailures(t *testing.T) {
	sentinel := errors.New("block export failed")
	tests := []struct {
		name   string
		source []scriptOp
		dest   []scriptOp
		want   string
	}{
		{
			name: "missing query",
			dest: []scriptOp{queryErrorOp("SELECT DISTINCT tx.block_num", sentinel)},
			want: "find missing blocks",
		},
		{
			name: "prepare",
			dest: []scriptOp{
				queryOp("SELECT DISTINCT tx.block_num", []string{"block_num"}, []driver.Value{int64(10)}),
				prepareErrorOp("INSERT INTO arch_block", sentinel),
			},
			want: "prepare arch_block",
		},
		{
			name:   "batch",
			source: []scriptOp{queryErrorOp("FROM block", sentinel)},
			dest: []scriptOp{
				queryOp("SELECT DISTINCT tx.block_num", []string{"block_num"}, []driver.Value{int64(10)}),
				prepareOp("INSERT INTO arch_block"),
			},
			want: "query block batch",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exporter := &SQLExporter{
				Source:      openScriptDB(t, test.source...),
				Destination: openScriptDB(t, test.dest...),
				BatchSize:   1,
			}
			_, err := exporter.exportBlocks(context.Background())
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func blockRow() []driver.Value {
	return []driver.Value{
		int64(10), int64(2), time.Unix(1_700_000_000, 0), "3.5", "hash", "parent",
	}
}

func TestExportBlockBatchFailures(t *testing.T) {
	sentinel := errors.New("block batch failed")
	columns := []string{"block_num", "num_tx", "ts", "cash_flow", "block_hash", "parent_hash"}
	tests := []struct {
		name   string
		source scriptOp
		stmt   []scriptOp
		want   string
	}{
		{
			name:   "query",
			source: queryErrorOp("FROM block", sentinel),
			stmt:   []scriptOp{prepareOp("INSERT test statement")},
			want:   "query block batch",
		},
		{
			name:   "scan",
			source: queryOp("FROM block", columns, []driver.Value{"bad", int64(1), time.Now(), "0", "h", "p"}),
			stmt:   []scriptOp{prepareOp("INSERT test statement")},
			want:   "scan block row",
		},
		{
			name:   "insert",
			source: queryOp("FROM block", columns, blockRow()),
			stmt: []scriptOp{
				prepareOp("INSERT test statement"),
				execErrorOp("INSERT test statement", sentinel),
			},
			want: "insert arch_block",
		},
		{
			name: "iteration",
			source: scriptOp{
				kind:      "query",
				contains:  "FROM block",
				columns:   columns,
				nextErrAt: 0,
				nextErr:   sentinel,
			},
			stmt: []scriptOp{prepareOp("INSERT test statement")},
			want: "iterate block batch",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			statement := prepareStatement(t, test.stmt...)
			exporter := &SQLExporter{Source: openScriptDB(t, test.source)}
			_, err := exporter.exportBlockBatch(context.Background(), statement, []int64{10})
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestColumnQueryFailures(t *testing.T) {
	sentinel := errors.New("column query failed")
	for _, test := range []struct {
		name string
		run  func(*sql.DB) error
		op   scriptOp
	}{
		{
			name: "string query",
			run: func(db *sql.DB) error {
				_, err := queryStringColumn(context.Background(), db, "SELECT strings")
				return err
			},
			op: queryErrorOp("SELECT strings", sentinel),
		},
		{
			name: "string scan",
			run: func(db *sql.DB) error {
				_, err := queryStringColumn(context.Background(), db, "SELECT strings")
				return err
			},
			op: queryOp("SELECT strings", []string{"value"}, []driver.Value{nil}),
		},
		{
			name: "string iteration",
			run: func(db *sql.DB) error {
				_, err := queryStringColumn(context.Background(), db, "SELECT strings")
				return err
			},
			op: scriptOp{kind: "query", contains: "SELECT strings", columns: []string{"value"}, nextErrAt: 0, nextErr: sentinel},
		},
		{
			name: "int query",
			run: func(db *sql.DB) error {
				_, err := queryInt64Column(context.Background(), db, "SELECT ints")
				return err
			},
			op: queryErrorOp("SELECT ints", sentinel),
		},
		{
			name: "int scan",
			run: func(db *sql.DB) error {
				_, err := queryInt64Column(context.Background(), db, "SELECT ints")
				return err
			},
			op: queryOp("SELECT ints", []string{"value"}, []driver.Value{"bad"}),
		},
		{
			name: "int iteration",
			run: func(db *sql.DB) error {
				_, err := queryInt64Column(context.Background(), db, "SELECT ints")
				return err
			},
			op: scriptOp{kind: "query", contains: "SELECT ints", columns: []string{"value"}, nextErrAt: 0, nextErr: sentinel},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := test.run(openScriptDB(t, test.op)); err == nil {
				t.Fatal("expected query helper failure")
			}
		})
	}
}
