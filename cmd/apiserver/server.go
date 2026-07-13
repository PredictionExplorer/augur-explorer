package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// defaultLogDir is the legacy operational log directory under $HOME.
const defaultLogDir = "ae_logs"

// serverDeps bundles the process-wide dependencies run assembles at boot:
// the shared store pool and the legacy info/error file loggers.
type serverDeps struct {
	store  *store.Store
	info   *log.Logger
	errlog *log.Logger
}

// newServerDeps opens the legacy log files under $HOME/ae_logs, connects the
// shared store pool (database log routed through the pgx slog tracer into
// webserver-db.log) and returns the bundle. Errors are fatal to the caller:
// the API server cannot run without its logs or database.
func newServerDeps(ctx context.Context, getenv func(string) string) (*serverDeps, error) {
	logDir := fmt.Sprintf("%v/%v", getenv("HOME"), defaultLogDir)
	_ = os.MkdirAll(logDir, os.ModePerm) // #nosec G301 G703 -- legacy log dir under $HOME; openAppendLog fails loudly if unusable

	infoFile, err := openAppendLog(fmt.Sprintf("%v/webserver_info.log", logDir))
	if err != nil {
		return nil, err
	}
	info := log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	errFile, err := openAppendLog(fmt.Sprintf("%v/webserver_error.log", logDir))
	if err != nil {
		return nil, err
	}
	errlog := log.New(errFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Database log output (failed and slow queries) goes through the pgx
	// slog tracer into the file the legacy Init_log wrote to.
	dbLogHandle, err := openAppendLog(fmt.Sprintf("%v/webserver-db.log", logDir))
	if err != nil {
		return nil, err
	}
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.New(slog.NewTextHandler(dbLogHandle, nil))
	st, err := store.New(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("can't connect to PostgreSQL database: %w\n%s", err, store.ConnectHint(err))
	}

	return &serverDeps{store: st, info: info, errlog: errlog}, nil
}

// openAppendLog opens (creating if needed) one of the server's append-only
// log files.
func openAppendLog(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666) // #nosec G302 G304 -- operational log under $HOME/ae_logs, world-readable by design
	if err != nil {
		return nil, fmt.Errorf("can't open log file: %w", err)
	}
	return f, nil
}
