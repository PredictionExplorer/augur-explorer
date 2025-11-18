package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	
	_ "github.com/lib/pq"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
)

// ConnectPostgres connects to a PostgreSQL database
func ConnectPostgres(cfg types.DatabaseConfig) (*sql.DB, error) {
	host, port, err := net.SplitHostPort(cfg.Host)
	if err != nil {
		host = cfg.Host
		port = "5432"
	}
	
	connStr := fmt.Sprintf("user='%s' dbname='%s' password='%s' host='%s' port='%s'",
		cfg.User, cfg.DBName, cfg.Pass, host, port)
	
	dbobj, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error connecting: %v", err))
	}
	
	_, err = dbobj.Exec("SET timezone TO 0") // Setting timezone to UTC
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DB Error: %v", err))
	}
	
	return dbobj, nil
}




