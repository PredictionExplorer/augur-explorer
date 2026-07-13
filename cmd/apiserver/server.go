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

type RWCGServer struct {
	store *store.Store
}

func create_rwcg_server() *RWCGServer {

	log_dir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), defaultLogDir)
	os.MkdirAll(log_dir, os.ModePerm)
	web_db_log_file := fmt.Sprintf("%v/%v", log_dir, "webserver-db.log")

	fname := fmt.Sprintf("%v/webserver_info.log", log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // #nosec G302 -- operational log, world-readable by design
	if err != nil {
		fmt.Printf("Can't start: %v\n", err)
		os.Exit(1)
	}
	Info = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	fname = fmt.Sprintf("%v/webserver_error.log", log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // #nosec G302 -- operational log, world-readable by design
	if err != nil {
		fmt.Printf("Can't start: %v\n", err)
		os.Exit(1)
	}
	Error = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Database log output (failed and slow queries) goes through the pgx
	// slog tracer into the file the legacy Init_log wrote to.
	dbLogHandle, err := os.OpenFile(web_db_log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // #nosec G302 -- operational log, world-readable by design
	if err != nil {
		fmt.Printf("Can't start: %v\n", err)
		os.Exit(1)
	}
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.New(slog.NewTextHandler(dbLogHandle, nil))
	st, err := store.New(context.Background(), cfg)
	if err != nil {
		fmt.Printf("Can't connect to PostgreSQL database.\nConnection error: %v\n%s", err, store.ConnectHint(err))
		os.Exit(1)
	}
	srv := new(RWCGServer)
	srv.store = st

	return srv
}
