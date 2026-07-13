package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
)

type archiveExporterFunc func(context.Context, string) (archive.ExportStats, error)

func (f archiveExporterFunc) ExportProject(ctx context.Context, project string) (archive.ExportStats, error) {
	return f(ctx, project)
}

type archiveVerifierFunc func(context.Context, string, archive.VerifyOptions) (archive.ProjectVerification, error)

func (f archiveVerifierFunc) VerifyProject(
	ctx context.Context,
	project string,
	options archive.VerifyOptions,
) (archive.ProjectVerification, error) {
	return f(ctx, project, options)
}

func TestArchiveExportCommandWiring(t *testing.T) {
	srcDB := newCommandTestDB(t)
	dstDB := newCommandTestDB(t)
	marker := archiveExporterFunc(func(context.Context, string) (archive.ExportStats, error) {
		return archive.ExportStats{}, nil
	})
	var (
		openCalls []string
		projects  []string
	)
	deps := defaultArchiveExportDeps()
	deps.resolveProjects = func(project string) ([]string, error) {
		if project != "both" {
			t.Fatalf("project selector = %q", project)
		}
		return []string{archive.ProjectCosmicGame, archive.ProjectRandomWalk}, nil
	}
	deps.openDB = func(driverName, conn string) (*sql.DB, error) {
		if driverName != "postgres" {
			t.Fatalf("driver = %q", driverName)
		}
		openCalls = append(openCalls, conn)
		if len(openCalls) == 1 {
			return srcDB, nil
		}
		return dstDB, nil
	}
	deps.newExporter = func(source, destination *sql.DB, logger archive.Logger) archive.ProjectExporter {
		if source != srcDB || destination != dstDB {
			t.Fatalf("exporter databases = %p, %p", source, destination)
		}
		logger.Println("export factory ready")
		return marker
	}
	deps.exportProjects = func(
		ctx context.Context,
		gotProjects []string,
		exporter archive.ProjectExporter,
	) ([]archive.ExportResult, error) {
		if err := ctx.Err(); err != nil {
			t.Fatalf("unexpected context error: %v", err)
		}
		if exporter == nil {
			t.Fatal("exporter was not passed through")
		}
		projects = append([]string(nil), gotProjects...)
		return []archive.ExportResult{{Project: archive.ProjectCosmicGame}}, nil
	}

	result := executeCommand(
		newArchiveExportCmdWithDeps(deps),
		"--src", "source-dsn",
		"--dst", "destination-dsn",
		"--project", "both",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	if !reflect.DeepEqual(openCalls, []string{"source-dsn", "destination-dsn"}) {
		t.Fatalf("open calls = %v", openCalls)
	}
	if !reflect.DeepEqual(projects, []string{archive.ProjectCosmicGame, archive.ProjectRandomWalk}) {
		t.Fatalf("projects = %v", projects)
	}
	if srcDB.Stats().MaxOpenConnections != 2 || dstDB.Stats().MaxOpenConnections != 2 {
		t.Fatalf("max connections = source %d destination %d",
			srcDB.Stats().MaxOpenConnections, dstDB.Stats().MaxOpenConnections)
	}
	if !strings.Contains(result.stderr, "export factory ready") ||
		!strings.Contains(result.stderr, "=== All exports complete ===") {
		t.Fatalf("stderr = %q", result.stderr)
	}
	assertCommandDBClosed(t, srcDB)
	assertCommandDBClosed(t, dstDB)
}

func TestArchiveExportCommandFailuresAndCancellation(t *testing.T) {
	t.Run("project validation", func(t *testing.T) {
		deps := defaultArchiveExportDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			t.Fatal("database opened after invalid project")
			return nil, nil
		}
		result := executeCommand(
			newArchiveExportCmdWithDeps(deps),
			"--src", "source",
			"--dst", "destination",
			"--project", "invalid",
		)
		if result.err == nil || !strings.Contains(result.err.Error(), "invalid project") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("source open", func(t *testing.T) {
		want := errors.New("source unavailable")
		deps := defaultArchiveExportDeps()
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(
			newArchiveExportCmdWithDeps(deps),
			"--src", "source",
			"--dst", "destination",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect to source") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("destination open closes source", func(t *testing.T) {
		srcDB := newCommandTestDB(t)
		want := errors.New("destination unavailable")
		calls := 0
		deps := defaultArchiveExportDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			calls++
			if calls == 1 {
				return srcDB, nil
			}
			return nil, want
		}
		result := executeCommand(
			newArchiveExportCmdWithDeps(deps),
			"--src", "source",
			"--dst", "destination",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect to destination") {
			t.Fatalf("error = %v", result.err)
		}
		assertCommandDBClosed(t, srcDB)
	})

	t.Run("export error", func(t *testing.T) {
		srcDB := newCommandTestDB(t)
		dstDB := newCommandTestDB(t)
		want := errors.New("copy failed")
		calls := 0
		deps := defaultArchiveExportDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			calls++
			if calls == 1 {
				return srcDB, nil
			}
			return dstDB, nil
		}
		deps.newExporter = func(*sql.DB, *sql.DB, archive.Logger) archive.ProjectExporter {
			return archiveExporterFunc(func(context.Context, string) (archive.ExportStats, error) {
				return archive.ExportStats{}, nil
			})
		}
		deps.exportProjects = func(
			context.Context,
			[]string,
			archive.ProjectExporter,
		) ([]archive.ExportResult, error) {
			return nil, want
		}
		result := executeCommand(
			newArchiveExportCmdWithDeps(deps),
			"--src", "source",
			"--dst", "destination",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, want) {
			t.Fatalf("error = %v", result.err)
		}
		assertCommandDBClosed(t, srcDB)
		assertCommandDBClosed(t, dstDB)
	})

	t.Run("canceled context", func(t *testing.T) {
		srcDB := newCommandTestDB(t)
		dstDB := newCommandTestDB(t)
		calls := 0
		deps := defaultArchiveExportDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			calls++
			if calls == 1 {
				return srcDB, nil
			}
			return dstDB, nil
		}
		deps.newExporter = func(*sql.DB, *sql.DB, archive.Logger) archive.ProjectExporter { return nil }
		deps.exportProjects = func(
			ctx context.Context,
			_ []string,
			_ archive.ProjectExporter,
		) ([]archive.ExportResult, error) {
			return nil, ctx.Err()
		}
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		command := newArchiveExportCmdWithDeps(deps)
		command.SetContext(ctx)
		result := executeCommand(
			command,
			"--src", "source",
			"--dst", "destination",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, context.Canceled) {
			t.Fatalf("error = %v", result.err)
		}
	})
}

func TestArchiveVerifyCommandWiringAndFailureReport(t *testing.T) {
	t.Run("passing strict verification", func(t *testing.T) {
		db := newCommandTestDB(t)
		var gotOptions archive.VerifyOptions
		deps := defaultArchiveVerifyDeps()
		deps.resolveProjects = func(project string) ([]string, error) {
			return []string{project}, nil
		}
		deps.openDB = func(driverName, conn string) (*sql.DB, error) {
			if driverName != "postgres" || conn != "archive-db" {
				t.Fatalf("open = %q %q", driverName, conn)
			}
			return db, nil
		}
		deps.newVerifier = func(gotDB *sql.DB, logger archive.Logger) archive.ProjectVerifier {
			if gotDB != db {
				t.Fatal("wrong verifier database")
			}
			logger.Println("verifier factory ready")
			return archiveVerifierFunc(func(
				context.Context,
				string,
				archive.VerifyOptions,
			) (archive.ProjectVerification, error) {
				return archive.ProjectVerification{}, nil
			})
		}
		deps.verifyProjects = func(
			_ context.Context,
			projects []string,
			options archive.VerifyOptions,
			verifier archive.ProjectVerifier,
		) (archive.VerificationReport, error) {
			if !reflect.DeepEqual(projects, []string{"randomwalk"}) || verifier == nil {
				t.Fatalf("projects/verifier = %v %#v", projects, verifier)
			}
			gotOptions = options
			return archive.VerificationReport{Passed: true}, nil
		}
		result := executeCommand(
			newArchiveVerifyCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
			"--strict-arch-block-metadata",
			"--strict-arch-tx-num-logs",
		)
		if result.err != nil {
			t.Fatal(result.err)
		}
		if gotOptions != (archive.VerifyOptions{StrictBlockMetadata: true, StrictTxNumLogs: true}) {
			t.Fatalf("options = %#v", gotOptions)
		}
		if !strings.Contains(result.stderr, "=== SUMMARY ===") ||
			!strings.Contains(result.stderr, "OK — no blocking mismatches") {
			t.Fatalf("stderr = %q", result.stderr)
		}
		assertCommandDBClosed(t, db)
	})

	t.Run("blocking mismatch", func(t *testing.T) {
		db := newCommandTestDB(t)
		deps := defaultArchiveVerifyDeps()
		deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
		deps.newVerifier = func(*sql.DB, archive.Logger) archive.ProjectVerifier { return nil }
		deps.verifyProjects = func(
			context.Context,
			[]string,
			archive.VerifyOptions,
			archive.ProjectVerifier,
		) (archive.VerificationReport, error) {
			return archive.VerificationReport{Passed: false}, nil
		}
		result := executeCommand(
			newArchiveVerifyCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "cosmicgame",
		)
		if result.err == nil || result.err.Error() != "archive verification failed" {
			t.Fatalf("error = %v", result.err)
		}
		if !strings.Contains(result.stderr, "FAILED — see details above") {
			t.Fatalf("stderr = %q", result.stderr)
		}
	})
}

func TestArchiveVerifyCommandSetupAndOperationalErrors(t *testing.T) {
	t.Run("invalid project", func(t *testing.T) {
		deps := defaultArchiveVerifyDeps()
		deps.openDB = func(string, string) (*sql.DB, error) {
			t.Fatal("database opened")
			return nil, nil
		}
		result := executeCommand(
			newArchiveVerifyCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "invalid",
		)
		if result.err == nil || !strings.Contains(result.err.Error(), "invalid project") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("database open", func(t *testing.T) {
		want := errors.New("database unavailable")
		deps := defaultArchiveVerifyDeps()
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(
			newArchiveVerifyCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect") {
			t.Fatalf("error = %v", result.err)
		}
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "operational error", err: errors.New("verification query failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			db := newCommandTestDB(t)
			deps := defaultArchiveVerifyDeps()
			deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
			deps.newVerifier = func(*sql.DB, archive.Logger) archive.ProjectVerifier { return nil }
			deps.verifyProjects = func(
				context.Context,
				[]string,
				archive.VerifyOptions,
				archive.ProjectVerifier,
			) (archive.VerificationReport, error) {
				return archive.VerificationReport{}, test.err
			}
			result := executeCommand(
				newArchiveVerifyCmdWithDeps(deps),
				"--db", "archive-db",
				"--project", "randomwalk",
			)
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
			assertCommandDBClosed(t, db)
		})
	}
}

func successfulArchiveNodeFillDeps(
	t *testing.T,
	db *sql.DB,
	rpc *fakeOpsRPC,
) (archiveNodeFillDeps, *bool) {
	t.Helper()
	closed := false
	deps := defaultArchiveNodeFillDeps()
	deps.getenv = func(name string) string {
		if name != "RPC_URL" {
			t.Fatalf("environment key = %q", name)
		}
		return "http://rpc.example"
	}
	deps.resolveProjects = func(project string) ([]string, error) {
		if project == "both" {
			return []string{archive.ProjectCosmicGame, archive.ProjectRandomWalk}, nil
		}
		return []string{project}, nil
	}
	deps.openStorage = func(ctx context.Context, conn string) (archiveNodeFillStorage, error) {
		if err := ctx.Err(); err != nil {
			return archiveNodeFillStorage{}, err
		}
		if conn != "archive-db" {
			t.Fatalf("db conn = %q", conn)
		}
		return archiveNodeFillStorage{
			db: db,
			close: func() {
				closed = true
				_ = db.Close()
			},
		}, nil
	}
	deps.requireSchema = func(context.Context, *sql.DB) error { return nil }
	deps.dialRPC = func(ctx context.Context, rpcURL string) (archiveNodeFillRPC, error) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if rpcURL != "http://rpc.example" {
			t.Fatalf("RPC URL = %q", rpcURL)
		}
		return rpc, nil
	}
	deps.newRepository = func(gotDB *sql.DB) archive.NodeFillRepository {
		if gotDB != db {
			t.Fatal("wrong repository database")
		}
		return nil
	}
	return deps, &closed
}

func TestArchiveNodeFillCommandWiringAndAliases(t *testing.T) {
	db := newCommandTestDB(t)
	rpc := &fakeOpsRPC{head: 100}
	deps, storageClosed := successfulArchiveNodeFillDeps(t, db, rpc)
	var (
		projects []string
		options  []archive.NodeFillOptions
	)
	deps.runProject = func(
		ctx context.Context,
		filler *archive.NodeFiller,
		project string,
		got archive.NodeFillOptions,
	) (archive.FillStats, error) {
		if err := ctx.Err(); err != nil {
			return archive.FillStats{}, err
		}
		if filler.Client != rpc || filler.Logger == nil {
			t.Fatalf("filler was not wired: %#v", filler)
		}
		projects = append(projects, project)
		options = append(options, got)
		return archive.FillStats{BlocksScanned: 2, LogsInserted: 1}, nil
	}
	result := executeCommand(
		newArchiveNodeFillCmdWithDeps(deps),
		"--db", "archive-db",
		"--project", "both",
		"--from", "7",
		"--start-block", "7",
		"--to", "55",
		"--batch", "11",
		"--dry-run",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	if !reflect.DeepEqual(projects, []string{archive.ProjectCosmicGame, archive.ProjectRandomWalk}) {
		t.Fatalf("projects = %v", projects)
	}
	wantOptions := archive.NodeFillOptions{FromBlock: 7, EndBlock: 55, BatchSize: 11, DryRun: true}
	for _, got := range options {
		if got != wantOptions {
			t.Fatalf("options = %#v, want %#v", got, wantOptions)
		}
	}
	for _, want := range []string{
		"Chain head: 100, scanning through block 55",
		"DRY RUN",
		"project: cosmicgame",
		"project: randomwalk",
		"TOTAL",
	} {
		if !strings.Contains(result.stderr, want) {
			t.Fatalf("stderr missing %q: %q", want, result.stderr)
		}
	}
	if !*storageClosed || !rpc.closed.Load() {
		t.Fatalf("resources closed = storage %v rpc %v", *storageClosed, rpc.closed.Load())
	}
	assertCommandDBClosed(t, db)
}

func TestArchiveNodeFillCommandSetupFailures(t *testing.T) {
	t.Run("alias disagreement precedes setup", func(t *testing.T) {
		deps := defaultArchiveNodeFillDeps()
		deps.getenv = func(string) string {
			t.Fatal("environment read after invalid aliases")
			return ""
		}
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
			"--from", "1",
			"--start-block", "2",
		)
		if result.err == nil || !strings.Contains(result.err.Error(), "disagree") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("missing RPC URL", func(t *testing.T) {
		deps := defaultArchiveNodeFillDeps()
		deps.getenv = func(string) string { return "" }
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
		)
		if result.err == nil || !strings.Contains(result.err.Error(), "RPC_URL") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("invalid project", func(t *testing.T) {
		deps := defaultArchiveNodeFillDeps()
		deps.getenv = func(string) string { return "rpc" }
		deps.openStorage = func(context.Context, string) (archiveNodeFillStorage, error) {
			t.Fatal("storage opened")
			return archiveNodeFillStorage{}, nil
		}
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "invalid",
		)
		if result.err == nil || !strings.Contains(result.err.Error(), "invalid project") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("storage open", func(t *testing.T) {
		want := errors.New("pool unavailable")
		deps := defaultArchiveNodeFillDeps()
		deps.getenv = func(string) string { return "rpc" }
		deps.openStorage = func(context.Context, string) (archiveNodeFillStorage, error) {
			return archiveNodeFillStorage{}, want
		}
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, want) {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("schema check", func(t *testing.T) {
		db := newCommandTestDB(t)
		rpc := &fakeOpsRPC{head: 1}
		deps, closed := successfulArchiveNodeFillDeps(t, db, rpc)
		want := errors.New("missing log index")
		deps.requireSchema = func(context.Context, *sql.DB) error { return want }
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "schema check") {
			t.Fatalf("error = %v", result.err)
		}
		if !*closed || rpc.closed.Load() {
			t.Fatalf("storage/rpc close state = %v/%v", *closed, rpc.closed.Load())
		}
	})

	t.Run("RPC dial", func(t *testing.T) {
		db := newCommandTestDB(t)
		rpc := &fakeOpsRPC{head: 1}
		deps, _ := successfulArchiveNodeFillDeps(t, db, rpc)
		want := errors.New("dial failed")
		deps.dialRPC = func(context.Context, string) (archiveNodeFillRPC, error) {
			return nil, want
		}
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "rpc connect") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("chain head", func(t *testing.T) {
		db := newCommandTestDB(t)
		want := errors.New("head failed")
		rpc := &fakeOpsRPC{headErr: want}
		deps, _ := successfulArchiveNodeFillDeps(t, db, rpc)
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "chain head") {
			t.Fatalf("error = %v", result.err)
		}
		if !rpc.closed.Load() {
			t.Fatal("RPC client was not closed")
		}
	})
}

func TestArchiveNodeFillRunErrorsAndCancellation(t *testing.T) {
	t.Run("best-effort row errors make command fail", func(t *testing.T) {
		db := newCommandTestDB(t)
		rpc := &fakeOpsRPC{head: 90}
		deps, _ := successfulArchiveNodeFillDeps(t, db, rpc)
		deps.runProject = func(
			context.Context,
			*archive.NodeFiller,
			string,
			archive.NodeFillOptions,
		) (archive.FillStats, error) {
			return archive.FillStats{RPCErrors: 2, DBErrors: 1}, nil
		}
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "randomwalk",
		)
		if result.err == nil || !strings.Contains(result.err.Error(), "unresolved errors (rpc=2 db=1)") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("project error stops subsequent projects", func(t *testing.T) {
		db := newCommandTestDB(t)
		rpc := &fakeOpsRPC{head: 90}
		deps, _ := successfulArchiveNodeFillDeps(t, db, rpc)
		want := errors.New("fill failed")
		calls := 0
		deps.runProject = func(
			context.Context,
			*archive.NodeFiller,
			string,
			archive.NodeFillOptions,
		) (archive.FillStats, error) {
			calls++
			return archive.FillStats{}, want
		}
		result := executeCommand(
			newArchiveNodeFillCmdWithDeps(deps),
			"--db", "archive-db",
			"--project", "both",
			"--start-block", "8",
		)
		if !errors.Is(result.err, want) || calls != 1 {
			t.Fatalf("calls=%d error=%v", calls, result.err)
		}
	})

	t.Run("canceled context", func(t *testing.T) {
		db := newCommandTestDB(t)
		rpc := &fakeOpsRPC{head: 90}
		deps, _ := successfulArchiveNodeFillDeps(t, db, rpc)
		deps.openStorage = func(context.Context, string) (archiveNodeFillStorage, error) {
			return archiveNodeFillStorage{db: db, close: func() { _ = db.Close() }}, nil
		}
		deps.dialRPC = func(context.Context, string) (archiveNodeFillRPC, error) {
			return rpc, nil
		}
		deps.runProject = func(
			ctx context.Context,
			_ *archive.NodeFiller,
			_ string,
			_ archive.NodeFillOptions,
		) (archive.FillStats, error) {
			return archive.FillStats{}, ctx.Err()
		}
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		command := newArchiveNodeFillCmdWithDeps(deps)
		command.SetContext(ctx)
		result := executeCommand(
			command,
			"--db", "archive-db",
			"--project", "randomwalk",
		)
		if !errors.Is(result.err, context.Canceled) {
			t.Fatalf("error = %v", result.err)
		}
	})
}

func TestArchiveDefaultDependencyFactories(t *testing.T) {
	src := newCommandTestDB(t)
	dst := newCommandTestDB(t)
	logger := log.New(&strings.Builder{}, "", 0)

	exportDeps := defaultArchiveExportDeps()
	if _, ok := exportDeps.newExporter(src, dst, logger).(*archive.SQLExporter); !ok {
		t.Fatal("default exporter is not SQLExporter")
	}
	verifyDeps := defaultArchiveVerifyDeps()
	if _, ok := verifyDeps.newVerifier(src, logger).(*archive.SQLVerifier); !ok {
		t.Fatal("default verifier is not SQLVerifier")
	}

	nodeDeps := defaultArchiveNodeFillDeps()
	if _, ok := nodeDeps.newRepository(src).(*archive.SQLNodeFillRepository); !ok {
		t.Fatal("default repository is not SQLNodeFillRepository")
	}
	_, err := nodeDeps.runProject(
		context.Background(),
		&archive.NodeFiller{},
		archive.ProjectRandomWalk,
		archive.NodeFillOptions{},
	)
	if err == nil || !strings.Contains(err.Error(), "repository is required") {
		t.Fatalf("default RunProject adapter error = %v", err)
	}
	if _, err := nodeDeps.dialRPC(context.Background(), "://bad-rpc-url"); err == nil {
		t.Fatal("default RPC dialer accepted malformed URL")
	}

	if _, err := openArchiveNodeFillStorage(context.Background(), "postgres://[::1"); err == nil ||
		!strings.Contains(err.Error(), "db pool config") {
		t.Fatalf("invalid pool config error = %v", err)
	}
	storage, err := openArchiveNodeFillStorage(
		context.Background(),
		"postgres://user:pass@127.0.0.1:1/database?sslmode=disable",
	)
	if err != nil {
		t.Fatalf("lazy pool construction: %v", err)
	}
	if storage.db == nil || storage.addressStore == nil || storage.close == nil {
		t.Fatalf("storage = %#v", storage)
	}
	storage.close()
}
