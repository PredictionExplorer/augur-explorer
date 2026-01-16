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
	
	// Configure connection pool to prevent connection leaks
	dbobj.SetMaxOpenConns(1)              // Only 1 connection needed per check
	dbobj.SetMaxIdleConns(0)              // Don't keep idle connections
	dbobj.SetConnMaxLifetime(0)           // Connections live forever (we close the pool after each use)
	dbobj.SetConnMaxIdleTime(0)           // No idle timeout needed (we set MaxIdleConns to 0)
	
	_, err = dbobj.Exec("SET timezone TO 0") // Setting timezone to UTC
	if err != nil {
		dbobj.Close() // Close if we fail to set timezone
		return nil, errors.New(fmt.Sprintf("DB Error: %v", err))
	}
	
	return dbobj, nil
}






