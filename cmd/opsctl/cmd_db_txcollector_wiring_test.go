package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/dbverify"
	"github.com/PredictionExplorer/augur-explorer/internal/ops/txcollector"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

func TestDBVerifyCommandWiring(t *testing.T) {
	primary := newCommandTestDB(t)
	secondary := newCommandTestDB(t)
	var openCalls []string
	deps := defaultDBVerifyDeps()
	deps.openDB = func(driverName, conn string) (*sql.DB, error) {
		if driverName != "postgres" {
			t.Fatalf("driver = %q", driverName)
		}
		openCalls = append(openCalls, conn)
		if len(openCalls) == 1 {
			return primary, nil
		}
		return secondary, nil
	}
	deps.loadIDs = func(ctx context.Context, db *sql.DB) ([]int64, error) {
		if err := ctx.Err(); err != nil || db != primary {
			t.Fatalf("load ids context/db = %v/%p", err, db)
		}
		return []int64{8, 13}, nil
	}
	deps.verify = func(
		ctx context.Context,
		primaryLoader, secondaryLoader dbverify.Loader,
		ids []int64,
		reportLimit int,
	) (dbverify.VerifyReport, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ids, []int64{8, 13}) ||
			reportLimit != dbverify.DefaultVerifyReportLimit {
			t.Fatalf("ids/limit = %v/%d", ids, reportLimit)
		}
		gotPrimary, ok := primaryLoader.(*dbverify.SQLLoader)
		if !ok || gotPrimary.DB != primary {
			t.Fatalf("primary loader = %#v", primaryLoader)
		}
		gotSecondary, ok := secondaryLoader.(*dbverify.SQLLoader)
		if !ok || gotSecondary.DB != secondary {
			t.Fatalf("secondary loader = %#v", secondaryLoader)
		}
		return dbverify.VerifyReport{}, nil
	}
	result := executeCommand(
		newDBVerifyCmdWithDeps(deps),
		"--primary", "primary-dsn",
		"--secondary", "secondary-dsn",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	if !reflect.DeepEqual(openCalls, []string{"primary-dsn", "secondary-dsn"}) {
		t.Fatalf("open calls = %v", openCalls)
	}
	for _, want := range []string{
		"Connected to primary database",
		"Connected to secondary database",
		"Found 2 contract address IDs",
		"All tables match perfectly",
	} {
		if !strings.Contains(result.stderr, want) {
			t.Fatalf("stderr missing %q: %q", want, result.stderr)
		}
	}
	assertCommandDBClosed(t, primary)
	assertCommandDBClosed(t, secondary)
}

func TestDBVerifyCommandMismatchAndErrors(t *testing.T) {
	t.Run("mismatch", func(t *testing.T) {
		primary := newCommandTestDB(t)
		secondary := newCommandTestDB(t)
		calls := 0
		deps := defaultDBVerifyDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			calls++
			if calls == 1 {
				return primary, nil
			}
			return secondary, nil
		}
		deps.loadIDs = func(context.Context, *sql.DB) ([]int64, error) { return []int64{1}, nil }
		deps.verify = func(
			context.Context,
			dbverify.Loader,
			dbverify.Loader,
			[]int64,
			int,
		) (dbverify.VerifyReport, error) {
			return dbverify.VerifyReport{
				Events: dbverify.Comparison{MissingTotal: 1},
			}, nil
		}
		result := executeCommand(
			newDBVerifyCmdWithDeps(deps),
			"--primary", "primary",
			"--secondary", "secondary",
		)
		if result.err == nil || result.err.Error() != "verification FAILED" {
			t.Fatalf("error = %v", result.err)
		}
		if !strings.Contains(result.stderr, "evt_log: MISMATCH") {
			t.Fatalf("stderr = %q", result.stderr)
		}
	})

	t.Run("primary open", func(t *testing.T) {
		want := errors.New("primary unavailable")
		deps := defaultDBVerifyDeps()
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(
			newDBVerifyCmdWithDeps(deps),
			"--primary", "primary",
			"--secondary", "secondary",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect to primary") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("secondary open closes primary", func(t *testing.T) {
		primary := newCommandTestDB(t)
		want := errors.New("secondary unavailable")
		calls := 0
		deps := defaultDBVerifyDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			calls++
			if calls == 1 {
				return primary, nil
			}
			return nil, want
		}
		result := executeCommand(
			newDBVerifyCmdWithDeps(deps),
			"--primary", "primary",
			"--secondary", "secondary",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect to secondary") {
			t.Fatalf("error = %v", result.err)
		}
		assertCommandDBClosed(t, primary)
	})

	t.Run("contract ids", func(t *testing.T) {
		primary := newCommandTestDB(t)
		secondary := newCommandTestDB(t)
		want := errors.New("contracts failed")
		calls := 0
		deps := defaultDBVerifyDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			calls++
			if calls == 1 {
				return primary, nil
			}
			return secondary, nil
		}
		deps.loadIDs = func(context.Context, *sql.DB) ([]int64, error) { return nil, want }
		result := executeCommand(
			newDBVerifyCmdWithDeps(deps),
			"--primary", "primary",
			"--secondary", "secondary",
		)
		if !errors.Is(result.err, want) {
			t.Fatalf("error = %v", result.err)
		}
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "verification", err: errors.New("comparison failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			primary := newCommandTestDB(t)
			secondary := newCommandTestDB(t)
			calls := 0
			deps := defaultDBVerifyDeps()
			deps.openDB = func(string, string) (*sql.DB, error) {
				calls++
				if calls == 1 {
					return primary, nil
				}
				return secondary, nil
			}
			deps.loadIDs = func(context.Context, *sql.DB) ([]int64, error) { return []int64{1}, nil }
			deps.verify = func(
				context.Context,
				dbverify.Loader,
				dbverify.Loader,
				[]int64,
				int,
			) (dbverify.VerifyReport, error) {
				return dbverify.VerifyReport{}, test.err
			}
			result := executeCommand(
				newDBVerifyCmdWithDeps(deps),
				"--primary", "primary",
				"--secondary", "secondary",
			)
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
		})
	}
}

func TestDBEvtlogDiffCommandWiring(t *testing.T) {
	primary := newCommandTestDB(t)
	secondary := newCommandTestDB(t)
	calls := 0
	deps := defaultDBEvtlogDiffDeps()
	deps.openDB = func(driverName, conn string) (*sql.DB, error) {
		if driverName != "postgres" {
			t.Fatalf("driver = %q", driverName)
		}
		calls++
		if calls == 1 {
			if conn != "primary-dsn" {
				t.Fatalf("primary conn = %q", conn)
			}
			return primary, nil
		}
		if conn != "secondary-dsn" {
			t.Fatalf("secondary conn = %q", conn)
		}
		return secondary, nil
	}
	deps.loadIDs = func(context.Context, *sql.DB) ([]int64, error) { return []int64{21}, nil }
	deps.diff = func(
		_ context.Context,
		primaryLoader, secondaryLoader dbverify.Loader,
		ids []int64,
		loadLimit, reportLimit int,
	) (dbverify.EventLogDiffReport, error) {
		if !reflect.DeepEqual(ids, []int64{21}) || loadLimit != 12 ||
			reportLimit != dbverify.DefaultDiffReportLimit {
			t.Fatalf("ids/limits = %v/%d/%d", ids, loadLimit, reportLimit)
		}
		if primaryLoader.(*dbverify.SQLLoader).DB != primary ||
			secondaryLoader.(*dbverify.SQLLoader).DB != secondary {
			t.Fatal("loaders reference wrong databases")
		}
		return dbverify.EventLogDiffReport{
			PrimaryTotal:    4,
			SecondaryTotal:  5,
			PrimaryLoaded:   4,
			SecondaryLoaded: 5,
		}, nil
	}
	result := executeCommand(
		newDBEvtlogDiffCmdWithDeps(deps),
		"--primary", "primary-dsn",
		"--secondary", "secondary-dsn",
		"--limit", "12",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	for _, want := range []string{
		"Connected to primary database",
		"Connected to secondary database",
		"Found 1 contract address IDs",
		"Primary has 4 records",
		"Loaded 5 events from secondary",
	} {
		if !strings.Contains(result.stderr, want) {
			t.Fatalf("stderr missing %q: %q", want, result.stderr)
		}
	}
	assertCommandDBClosed(t, primary)
	assertCommandDBClosed(t, secondary)
}

func TestDBEvtlogDiffCommandErrors(t *testing.T) {
	t.Run("primary open", func(t *testing.T) {
		want := errors.New("primary unavailable")
		deps := defaultDBEvtlogDiffDeps()
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(
			newDBEvtlogDiffCmdWithDeps(deps),
			"--primary", "primary",
			"--secondary", "secondary",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect to primary") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("secondary open", func(t *testing.T) {
		primary := newCommandTestDB(t)
		want := errors.New("secondary unavailable")
		calls := 0
		deps := defaultDBEvtlogDiffDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			calls++
			if calls == 1 {
				return primary, nil
			}
			return nil, want
		}
		result := executeCommand(
			newDBEvtlogDiffCmdWithDeps(deps),
			"--primary", "primary",
			"--secondary", "secondary",
		)
		if !errors.Is(result.err, want) {
			t.Fatalf("error = %v", result.err)
		}
		assertCommandDBClosed(t, primary)
	})

	t.Run("contract ids", func(t *testing.T) {
		primary := newCommandTestDB(t)
		secondary := newCommandTestDB(t)
		want := errors.New("contracts failed")
		calls := 0
		deps := defaultDBEvtlogDiffDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			calls++
			if calls == 1 {
				return primary, nil
			}
			return secondary, nil
		}
		deps.loadIDs = func(context.Context, *sql.DB) ([]int64, error) { return nil, want }
		result := executeCommand(
			newDBEvtlogDiffCmdWithDeps(deps),
			"--primary", "primary",
			"--secondary", "secondary",
		)
		if !errors.Is(result.err, want) {
			t.Fatalf("error = %v", result.err)
		}
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "diff", err: errors.New("diff failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			primary := newCommandTestDB(t)
			secondary := newCommandTestDB(t)
			calls := 0
			deps := defaultDBEvtlogDiffDeps()
			deps.openDB = func(string, string) (*sql.DB, error) {
				calls++
				if calls == 1 {
					return primary, nil
				}
				return secondary, nil
			}
			deps.loadIDs = func(context.Context, *sql.DB) ([]int64, error) { return []int64{1}, nil }
			deps.diff = func(
				context.Context,
				dbverify.Loader,
				dbverify.Loader,
				[]int64,
				int,
				int,
			) (dbverify.EventLogDiffReport, error) {
				return dbverify.EventLogDiffReport{}, test.err
			}
			result := executeCommand(
				newDBEvtlogDiffCmdWithDeps(deps),
				"--primary", "primary",
				"--secondary", "secondary",
			)
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
		})
	}
}

func collectorConfig() *toolutil.CollectorConfig {
	return &toolutil.CollectorConfig{
		RPCURL:            "http://rpc.example",
		OutputDir:         "/tmp/collector-output",
		StartBlock:        10,
		ContractAddresses: []string{"0x0000000000000000000000000000000000000001"},
	}
}

func TestTxCollectorRunCommandWiring(t *testing.T) {
	rpc := &fakeOpsRPC{head: 30}
	deps := defaultTxCollectorRunDeps()
	deps.loadConfig = func(path string) (*toolutil.CollectorConfig, error) {
		if path != "collector.json" {
			t.Fatalf("config path = %q", path)
		}
		return collectorConfig(), nil
	}
	deps.dialRPC = func(ctx context.Context, rpcURL string) (txCollectorRunRPC, error) {
		if err := ctx.Err(); err != nil || rpcURL != "http://rpc.example" {
			t.Fatalf("dial context/URL = %v/%q", err, rpcURL)
		}
		return rpc, nil
	}
	deps.run = func(
		ctx context.Context,
		cfg txcollector.Config,
		options txcollector.RunOptions,
	) (txcollector.RunStats, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if cfg.Client != rpc || cfg.OutputDir != "/tmp/collector-output" ||
			len(cfg.Contracts) != 1 ||
			cfg.Contracts[0] != common.HexToAddress("0x0000000000000000000000000000000000000001") {
			t.Fatalf("config = %#v", cfg)
		}
		if options.FromBlock != 12 || options.ToBlock != 25 ||
			options.InitialBatch != 500 || options.MinBatch != 500 ||
			options.RetryDelay != txcollector.DefaultRetryDelay {
			t.Fatalf("options = %#v", options)
		}
		cfg.Logger.Printf("collector engine invoked")
		return txcollector.RunStats{
			BlocksScanned:  14,
			LogsSeen:       3,
			TxUnique:       2,
			TxWritten:      2,
			ReceiptWritten: 2,
		}, nil
	}
	result := executeCommand(
		newTxCollectorRunCmdWithDeps(deps),
		"--config", "collector.json",
		"--start-block", "12",
		"--batch", "500",
		"--to", "25",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	for _, want := range []string{
		"Scanning blocks 12 .. 25 (batch 500)",
		"collector engine invoked",
		"done: blocks_scanned=14 logs=3 unique_tx=2",
	} {
		if !strings.Contains(result.stderr, want) {
			t.Fatalf("stderr missing %q: %q", want, result.stderr)
		}
	}
	if !rpc.closed.Load() {
		t.Fatal("RPC client was not closed")
	}
}

func TestTxCollectorRunValidationSetupAndErrors(t *testing.T) {
	t.Run("config load", func(t *testing.T) {
		want := errors.New("read failed")
		deps := defaultTxCollectorRunDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return nil, want }
		result := executeCommand(newTxCollectorRunCmdWithDeps(deps), "--config", "bad.json")
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "config") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("missing RPC URL", func(t *testing.T) {
		deps := defaultTxCollectorRunDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) {
			cfg := collectorConfig()
			cfg.RPCURL = ""
			return cfg, nil
		}
		result := executeCommand(newTxCollectorRunCmdWithDeps(deps), "--config", "collector.json")
		if result.err == nil || !strings.Contains(result.err.Error(), "rpc_url is required") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("invalid contract", func(t *testing.T) {
		deps := defaultTxCollectorRunDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) {
			cfg := collectorConfig()
			cfg.ContractAddresses = []string{"not-an-address"}
			return cfg, nil
		}
		result := executeCommand(newTxCollectorRunCmdWithDeps(deps), "--config", "collector.json")
		if result.err == nil || !strings.Contains(result.err.Error(), "contracts") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("RPC dial", func(t *testing.T) {
		want := errors.New("dial failed")
		deps := defaultTxCollectorRunDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
		deps.dialRPC = func(context.Context, string) (txCollectorRunRPC, error) { return nil, want }
		result := executeCommand(newTxCollectorRunCmdWithDeps(deps), "--config", "collector.json")
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "rpc connect") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("chain head", func(t *testing.T) {
		want := errors.New("head failed")
		rpc := &fakeOpsRPC{headErr: want}
		deps := defaultTxCollectorRunDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
		deps.dialRPC = func(context.Context, string) (txCollectorRunRPC, error) { return rpc, nil }
		result := executeCommand(newTxCollectorRunCmdWithDeps(deps), "--config", "collector.json")
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "chain head") {
			t.Fatalf("error = %v", result.err)
		}
		if !rpc.closed.Load() {
			t.Fatal("RPC client was not closed")
		}
	})

	t.Run("empty range", func(t *testing.T) {
		rpc := &fakeOpsRPC{head: 9}
		deps := defaultTxCollectorRunDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
		deps.dialRPC = func(context.Context, string) (txCollectorRunRPC, error) { return rpc, nil }
		deps.run = func(context.Context, txcollector.Config, txcollector.RunOptions) (txcollector.RunStats, error) {
			t.Fatal("collector ran for empty range")
			return txcollector.RunStats{}, nil
		}
		result := executeCommand(newTxCollectorRunCmdWithDeps(deps), "--config", "collector.json")
		if result.err != nil || !strings.Contains(result.stderr, "nothing to scan") {
			t.Fatalf("error/output = %v/%q", result.err, result.stderr)
		}
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "run error", err: errors.New("collector failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			rpc := &fakeOpsRPC{head: 20}
			deps := defaultTxCollectorRunDeps()
			deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
			deps.dialRPC = func(context.Context, string) (txCollectorRunRPC, error) { return rpc, nil }
			deps.run = func(context.Context, txcollector.Config, txcollector.RunOptions) (txcollector.RunStats, error) {
				return txcollector.RunStats{}, test.err
			}
			result := executeCommand(newTxCollectorRunCmdWithDeps(deps), "--config", "collector.json")
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
		})
	}
}

func TestTxCollectorVerifyCommandWiring(t *testing.T) {
	db := newCommandTestDB(t)
	rows := []txcollector.EventRow{{BlockNum: 14, TxHash: "0x01"}}
	envCalls := 0
	deps := defaultTxCollectorVerifyDeps()
	deps.loadConfig = func(path string) (*toolutil.CollectorConfig, error) {
		if path != "collector.json" {
			t.Fatalf("path = %q", path)
		}
		return collectorConfig(), nil
	}
	deps.postgresConn = func() (string, error) {
		envCalls++
		return "env-dsn", nil
	}
	deps.openDB = func(driverName, conn string) (*sql.DB, error) {
		if driverName != "postgres" || conn != "explicit-dsn" {
			t.Fatalf("open = %q/%q", driverName, conn)
		}
		return db, nil
	}
	deps.loadRows = func(
		ctx context.Context,
		gotDB *sql.DB,
		addresses []string,
		fromBlock uint64,
	) ([]txcollector.EventRow, error) {
		if err := ctx.Err(); err != nil || gotDB != db || len(addresses) != 1 || fromBlock != 22 {
			t.Fatalf("load rows = %v/%p/%v/%d", err, gotDB, addresses, fromBlock)
		}
		return rows, nil
	}
	deps.verify = func(ctx context.Context, cfg txcollector.VerifyConfig) (txcollector.VerifyStats, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if cfg.OutputDir != "/tmp/collector-output" || !reflect.DeepEqual(cfg.Rows, rows) ||
			cfg.MaxReport != 7 {
			t.Fatalf("verify config = %#v", cfg)
		}
		cfg.Logger.Printf("verification engine invoked")
		return txcollector.VerifyStats{EvtRowsTotal: 1, LogVerifiedOK: 1}, nil
	}
	result := executeCommand(
		newTxCollectorVerifyCmdWithDeps(deps),
		"--config", "collector.json",
		"--db", "explicit-dsn",
		"--start-block", "22",
		"--max-report", "7",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	if envCalls != 0 {
		t.Fatalf("environment connection calls = %d", envCalls)
	}
	for _, want := range []string{
		"Output dir: /tmp/collector-output",
		"evt_log block_num >= 22",
		"Loaded 1 evt_log rows",
		"verification engine invoked",
	} {
		if !strings.Contains(result.stderr, want) {
			t.Fatalf("stderr missing %q: %q", want, result.stderr)
		}
	}
	if db.Stats().MaxOpenConnections != 4 {
		t.Fatalf("max connections = %d", db.Stats().MaxOpenConnections)
	}
	assertCommandDBClosed(t, db)
}

func TestTxCollectorVerifyFallbackAndErrors(t *testing.T) {
	t.Run("environment connection fallback", func(t *testing.T) {
		db := newCommandTestDB(t)
		deps := defaultTxCollectorVerifyDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
		deps.postgresConn = func() (string, error) { return "environment-dsn", nil }
		deps.openDB = func(_ string, conn string) (*sql.DB, error) {
			if conn != "environment-dsn" {
				t.Fatalf("conn = %q", conn)
			}
			return db, nil
		}
		deps.loadRows = func(context.Context, *sql.DB, []string, uint64) ([]txcollector.EventRow, error) {
			return nil, nil
		}
		deps.verify = func(context.Context, txcollector.VerifyConfig) (txcollector.VerifyStats, error) {
			return txcollector.VerifyStats{}, nil
		}
		result := executeCommand(newTxCollectorVerifyCmdWithDeps(deps), "--config", "collector.json")
		if result.err != nil {
			t.Fatal(result.err)
		}
	})

	t.Run("config", func(t *testing.T) {
		want := errors.New("config failed")
		deps := defaultTxCollectorVerifyDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return nil, want }
		result := executeCommand(newTxCollectorVerifyCmdWithDeps(deps), "--config", "collector.json")
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "config") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("contracts", func(t *testing.T) {
		deps := defaultTxCollectorVerifyDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) {
			cfg := collectorConfig()
			cfg.ContractAddresses = []string{"invalid"}
			return cfg, nil
		}
		result := executeCommand(newTxCollectorVerifyCmdWithDeps(deps), "--config", "collector.json")
		if result.err == nil || !strings.Contains(result.err.Error(), "contracts") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("environment", func(t *testing.T) {
		want := errors.New("environment missing")
		deps := defaultTxCollectorVerifyDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
		deps.postgresConn = func() (string, error) { return "", want }
		result := executeCommand(newTxCollectorVerifyCmdWithDeps(deps), "--config", "collector.json")
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "db") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("open", func(t *testing.T) {
		want := errors.New("database unavailable")
		deps := defaultTxCollectorVerifyDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(
			newTxCollectorVerifyCmdWithDeps(deps),
			"--config", "collector.json",
			"--db", "database",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("load rows", func(t *testing.T) {
		db := newCommandTestDB(t)
		want := errors.New("query failed")
		deps := defaultTxCollectorVerifyDeps()
		deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
		deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
		deps.loadRows = func(context.Context, *sql.DB, []string, uint64) ([]txcollector.EventRow, error) {
			return nil, want
		}
		result := executeCommand(
			newTxCollectorVerifyCmdWithDeps(deps),
			"--config", "collector.json",
			"--db", "database",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "load evt_log") {
			t.Fatalf("error = %v", result.err)
		}
		assertCommandDBClosed(t, db)
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "verify", err: errors.New("verification failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			db := newCommandTestDB(t)
			deps := defaultTxCollectorVerifyDeps()
			deps.loadConfig = func(string) (*toolutil.CollectorConfig, error) { return collectorConfig(), nil }
			deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
			deps.loadRows = func(context.Context, *sql.DB, []string, uint64) ([]txcollector.EventRow, error) {
				return nil, nil
			}
			deps.verify = func(context.Context, txcollector.VerifyConfig) (txcollector.VerifyStats, error) {
				return txcollector.VerifyStats{}, test.err
			}
			result := executeCommand(
				newTxCollectorVerifyCmdWithDeps(deps),
				"--config", "collector.json",
				"--db", "database",
			)
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
		})
	}
}

func TestDBAndCollectorDefaultDependencyAdapters(t *testing.T) {
	db := newCommandTestDB(t)
	if loader, ok := defaultDBVerifyDeps().newLoader(db).(*dbverify.SQLLoader); !ok || loader.DB != db {
		t.Fatalf("default verify loader = %#v", loader)
	}
	if loader, ok := defaultDBEvtlogDiffDeps().newLoader(db).(*dbverify.SQLLoader); !ok || loader.DB != db {
		t.Fatalf("default diff loader = %#v", loader)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	logger := log.New(&strings.Builder{}, "", 0)
	err := runEvtlogDiff(
		ctx,
		logger,
		"postgres://user:pass@127.0.0.1:1/primary?sslmode=disable",
		"postgres://user:pass@127.0.0.1:1/secondary?sslmode=disable",
		0,
	)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("default diff cancellation = %v", err)
	}

	if _, err := defaultTxCollectorRunDeps().dialRPC(context.Background(), "://bad-rpc-url"); err == nil {
		t.Fatal("default collector RPC dialer accepted malformed URL")
	}
}
