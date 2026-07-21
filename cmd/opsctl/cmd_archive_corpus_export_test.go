package main

import (
	"context"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
)

func TestArchiveCorpusExportCommandWiring(t *testing.T) {
	db := newCommandTestDB(t)
	var (
		gotConn    string
		gotOptions archive.CorpusExportOptions
	)
	deps := defaultArchiveCorpusExportDeps()
	deps.openDB = func(_ context.Context, conn string) (opsDB, error) {
		gotConn = conn
		return db, nil
	}
	deps.export = func(
		ctx context.Context,
		gotDB archive.Querier,
		options archive.CorpusExportOptions,
		w io.Writer,
	) (archive.CorpusExportStats, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if gotDB != db {
			t.Fatal("wrong archive database")
		}
		gotOptions = options
		if _, err := io.WriteString(w, "{\"corpus\":true}\n"); err != nil {
			t.Fatal(err)
		}
		return archive.CorpusExportStats{Transactions: 2, EventLogs: 3}, nil
	}

	result := executeCommand(
		newArchiveCorpusExportCmdWithDeps(deps),
		"--db", "archive-dsn",
		"--project", "cosmicgame",
		"--tx-hash", corpusCommandTxA,
		"--tx-hash", corpusCommandTxB,
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	if gotConn != "archive-dsn" {
		t.Fatalf("connection = %q", gotConn)
	}
	if !reflect.DeepEqual(gotOptions, archive.CorpusExportOptions{
		Project:  "cosmicgame",
		TxHashes: []string{corpusCommandTxA, corpusCommandTxB},
	}) {
		t.Fatalf("options = %+v", gotOptions)
	}
	if result.stdout != "{\"corpus\":true}\n" {
		t.Fatalf("stdout = %q", result.stdout)
	}
	if !strings.Contains(result.stderr, "exported 3 event logs from 2 complete transactions") {
		t.Fatalf("stderr = %q", result.stderr)
	}
	assertCommandDBClosed(t, db)
}

func TestArchiveCorpusExportCommandFailures(t *testing.T) {
	t.Run("database open", func(t *testing.T) {
		sentinel := errors.New("archive unavailable")
		deps := defaultArchiveCorpusExportDeps()
		deps.openDB = func(context.Context, string) (opsDB, error) {
			return nil, sentinel
		}
		result := executeCommand(
			newArchiveCorpusExportCmdWithDeps(deps),
			"--db", "archive-dsn",
			"--project", "randomwalk",
			"--tx-hash", corpusCommandTxA,
		)
		if !errors.Is(result.err, sentinel) ||
			!strings.Contains(result.err.Error(), "connect to archive database") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("export", func(t *testing.T) {
		db := newCommandTestDB(t)
		sentinel := errors.New("bad archive row")
		deps := defaultArchiveCorpusExportDeps()
		deps.openDB = func(context.Context, string) (opsDB, error) { return db, nil }
		deps.export = func(
			context.Context,
			archive.Querier,
			archive.CorpusExportOptions,
			io.Writer,
		) (archive.CorpusExportStats, error) {
			return archive.CorpusExportStats{}, sentinel
		}
		result := executeCommand(
			newArchiveCorpusExportCmdWithDeps(deps),
			"--db", "archive-dsn",
			"--project", "randomwalk",
			"--tx-hash", corpusCommandTxA,
		)
		if !errors.Is(result.err, sentinel) {
			t.Fatalf("error = %v", result.err)
		}
		if result.stdout != "" || result.stderr != "" {
			t.Fatalf("output on failure: stdout=%q stderr=%q", result.stdout, result.stderr)
		}
		assertCommandDBClosed(t, db)
	})
}

const (
	corpusCommandTxA = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	corpusCommandTxB = "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
)
