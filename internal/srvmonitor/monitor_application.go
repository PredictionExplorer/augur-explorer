package srvmonitor

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

// ApplicationStatus holds status for application layer.
type ApplicationStatus struct {
	Config          DatabaseConfig
	TableName       string
	LastBlockNum    int64
	OfficialLagDiff int64
	ErrStr          string
	ChainID         string
	X, Y            int
}

// ApplicationMonitor reports the last indexed block of each application
// database (the ETL watermark resolved to its block number) and its lag
// behind the official RPC node of the same chain.
type ApplicationMonitor struct {
	apps        []ApplicationStatus
	sharedState *SharedRPCState
	position    Position
	interval    time.Duration
	connect     Connector
}

// NewApplicationMonitor creates a new application monitor.
func NewApplicationMonitor(dbs []DatabaseConfig, sharedState *SharedRPCState, iv Intervals) *ApplicationMonitor {
	m := &ApplicationMonitor{
		apps:        make([]ApplicationStatus, 0, len(dbs)),
		sharedState: sharedState,
		position:    Position{X: 1, Y: 24},
		interval:    iv.DB,
		connect:     ConnectPostgres,
	}

	// Initialize app statuses
	// First 2 are cosmic, last 2 are rwalk
	for i, db := range dbs {
		tableName := "cg_proc_status"
		if i >= 2 {
			tableName = "rw_proc_status"
		}

		m.apps = append(m.apps, ApplicationStatus{
			Config:          db,
			TableName:       tableName,
			OfficialLagDiff: math.MaxInt64,
			X:               1,
			Y:               25 + i,
		})
	}

	return m
}

// Name returns the monitor name.
func (m *ApplicationMonitor) Name() string {
	return "Application Layer Monitor"
}

// Start begins monitoring.
func (m *ApplicationMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check performs a check cycle.
func (m *ApplicationMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.apps))

	for i := range m.apps {
		go func(status *ApplicationStatus) {
			defer wg.Done()
			m.checkApp(ctx, status, errorChan)
		}(&m.apps[i])
	}

	wg.Wait()
	m.display(disp)
}

// checkApp checks a single application database.
func (m *ApplicationMonitor) checkApp(ctx context.Context, status *ApplicationStatus, errorChan chan<- string) {
	status.ErrStr = ""

	dbobj, err := m.connect(ctx, status.Config)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}
	defer func() { _ = dbobj.Close(ctx) }() // best-effort close of per-check connection

	var chainIDStr string
	err = dbobj.QueryRow(ctx, "SELECT chain_id FROM contract_addresses").Scan(&chainIDStr)
	if err != nil {
		if isNoRows(err) {
			status.LastBlockNum = 0
			return
		}
		status.ErrStr = fmt.Sprintf("Error %v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}
	status.ChainID = chainIDStr

	var lastEvtID int64
	// The table name is one of two package-internal constants, never input.
	err = dbobj.QueryRow(ctx, "SELECT last_evt_id FROM "+status.TableName).Scan(&lastEvtID)
	if err != nil {
		if isNoRows(err) {
			status.LastBlockNum = 0
			return
		}
		status.ErrStr = fmt.Sprintf("Error %v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	var bnum int64
	err = dbobj.QueryRow(ctx, "SELECT block_num FROM evt_log WHERE id=$1", lastEvtID).Scan(&bnum)
	if err != nil {
		if isNoRows(err) {
			status.LastBlockNum = 0
			return
		}
		status.ErrStr = fmt.Sprintf("Error %v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	status.LastBlockNum = bnum

	// Calculate lag using shared state
	if m.sharedState != nil {
		if officialBlock := m.sharedState.Official(chainIDStr); officialBlock != 0 {
			status.OfficialLagDiff = officialBlock - status.LastBlockNum
		}
	}
}

// display renders the application status.
func (m *ApplicationMonitor) display(disp Display) {
	// Header
	disp.DrawText(Position{X: 1, Y: 24},
		"----------- Last Block Numbers in Postgres ----------",
		ColorWhite, ColorDefault)

	// Each application
	for _, status := range m.apps {
		y := status.Y

		// Block number (first column)
		disp.DrawText(Position{X: status.X, Y: y}, fmt.Sprintf("%9d", status.LastBlockNum),
			ColorBlue, ColorDefault)

		// Official lag (second column)
		officialDiff := "------"
		if status.OfficialLagDiff != math.MaxInt64 {
			officialDiff = fmt.Sprintf("%6v", status.OfficialLagDiff)
		}
		disp.DrawText(Position{X: status.X + 10, Y: y}, officialDiff, ColorBlue, ColorDefault)

		// Title (third column - shifted right)
		disp.DrawText(Position{X: status.X + 20, Y: y}, status.Config.Name, ColorWhite, ColorDefault)
	}

	disp.Flush()
}
