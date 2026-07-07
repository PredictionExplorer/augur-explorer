package main

import (
	"context"
	"fmt"
	"log"
	"os"

	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

const (
	DEFAULT_DB_LOG_FILE_NAME = "/var/tmp/backend-db.log"
)

type RWCGServer struct {
	store *store.Store
	db    *store.SQLStorage
}

func create_rwcg_server() *RWCGServer {

	log_dir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), DEFAULT_LOG_DIR)
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

	st, err := store.New(context.Background(), store.ConfigFromEnv())
	if err != nil {
		fmt.Printf("Can't connect to PostgreSQL database.\nConnection error: %v\n%s", err, store.ConnectHint(err))
		os.Exit(1)
	}
	srv := new(RWCGServer)
	srv.store = st
	srv.db = store.NewSQLStorageFromDB(st.DB(), Info)
	if err := srv.db.Init_log(web_db_log_file); err != nil {
		fmt.Printf("Can't start: %v\n", err)
		os.Exit(1)
	}

	return srv
}
