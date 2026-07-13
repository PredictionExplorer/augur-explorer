package srvmonitor

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// DatabaseStatus holds status for a single database.
type DatabaseStatus struct {
	Config       DatabaseConfig
	LastBlockNum int64
	Alive        bool
	ErrStr       string
	X, Y         int
}

// DatabaseMonitor checks that every Layer 1 database is reachable and its
// block table grows between two reads one blockWait apart.
type DatabaseMonitor struct {
	statuses []*DatabaseStatus
	position Position
	interval time.Duration
	connect  Connector
	// blockWait pauses between the two block-number reads of one check.
	// Tests replace it to insert rows instead of sleeping.
	blockWait func(ctx context.Context)
}

// NewDatabaseMonitor creates a new database monitor.
func NewDatabaseMonitor(databases []DatabaseConfig, iv Intervals) *DatabaseMonitor {
	m := &DatabaseMonitor{
		statuses: make([]*DatabaseStatus, len(databases)),
		position: Position{X: 0, Y: 11},
		interval: iv.DB,
		connect:  ConnectPostgres,
	}
	m.blockWait = func(ctx context.Context) { sleepCtx(ctx, iv.DBBlockWait) }

	for i, db := range databases {
		m.statuses[i] = &DatabaseStatus{
			Config: db,
			X:      1,
			Y:      13 + i,
		}
	}

	return m
}

// Name returns the monitor name.
func (m *DatabaseMonitor) Name() string {
	return "Database Monitor"
}

// Start begins monitoring.
func (m *DatabaseMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check performs a check cycle.
func (m *DatabaseMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))

	for _, status := range m.statuses {
		go func(status *DatabaseStatus) {
			defer wg.Done()
			m.checkDatabase(ctx, status, errorChan)
		}(status)
	}

	wg.Wait()
	m.display(disp)
}

// checkDatabase checks a single database.
func (m *DatabaseMonitor) checkDatabase(ctx context.Context, status *DatabaseStatus, errorChan chan<- string) {
	status.ErrStr = ""
	status.Alive = false

	dbobj, err := m.connect(ctx, status.Config)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}
	defer func() { _ = dbobj.Close(ctx) }() // best-effort close of per-check connection

	const blockQuery = "SELECT block_num FROM block ORDER BY block_num DESC LIMIT 1"

	var bnum1 int64
	err = dbobj.QueryRow(ctx, blockQuery).Scan(&bnum1)
	if err != nil {
		if isNoRows(err) {
			return
		}
		status.ErrStr = fmt.Sprintf("Error %v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	m.blockWait(ctx)

	var bnum2 int64
	err = dbobj.QueryRow(ctx, blockQuery).Scan(&bnum2)
	if err != nil {
		if isNoRows(err) {
			return
		}
		status.ErrStr = fmt.Sprintf("Error %v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	diff := bnum2 - bnum1
	if diff == 0 {
		status.ErrStr = fmt.Sprintf("Block difference is zero (last block = %v)", bnum2)
		sendErr(ctx, errorChan, status.ErrStr)
	} else {
		status.Alive = true
	}
	status.LastBlockNum = bnum2
}

// display renders the database status.
func (m *DatabaseMonitor) display(disp Display) {
	// Header
	disp.DrawText(Position{X: 0, Y: 11},
		"--------------------- SQL DB --------------------------------",
		ColorWhite, ColorDefault)

	// Each database
	for _, status := range m.statuses {
		y := status.Y

		// Status (first column)
		aliveStr := "Alive"
		color := ColorGreen
		if !status.Alive {
			aliveStr = "DOWN "
			color = ColorRed
		}
		disp.DrawText(Position{X: status.X, Y: y}, aliveStr, color, ColorDefault)

		// Block number (second column)
		disp.DrawText(Position{X: status.X + 10, Y: y}, fmt.Sprintf("%v", status.LastBlockNum),
			ColorBlue, ColorDefault)

		// Name (third column - shifted right)
		disp.DrawText(Position{X: status.X + 20, Y: y}, status.Config.Name, ColorWhite, ColorDefault)

		// Host (fourth column - shifted right)
		disp.DrawText(Position{X: status.X + 55, Y: y}, status.Config.Host, ColorWhite, ColorDefault)
	}

	disp.Flush()
}
