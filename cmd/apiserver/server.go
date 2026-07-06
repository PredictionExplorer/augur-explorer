package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/store"
)

const (
	DEFAULT_DB_LOG_FILE_NAME = "/var/tmp/backend-db.log"
)

type RWCGServer struct {
	db *SQLStorage
}

func create_rwcg_server() *RWCGServer {

	log_dir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	web_db_log_file := fmt.Sprintf("%v/%v", log_dir, "webserver-db.log")

	fname := fmt.Sprintf("%v/webserver_info.log", log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Can't start: %v\n", err)
		os.Exit(1)
	}
	Info = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	fname = fmt.Sprintf("%v/webserver_error.log", log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Can't start: %v\n", err)
		os.Exit(1)
	}
	Error = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	srv := new(RWCGServer)
	srv.db = Connect_to_storage(Info)
	if srv.db == nil {
		os.Exit(1)
	}
	if err := srv.db.Init_log(web_db_log_file); err != nil {
		fmt.Printf("Can't start: %v\n", err)
		os.Exit(1)
	}

	return srv
}
