package srvmonitor

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// EventTableStatus holds status for a single event table database.
type EventTableStatus struct {
	Config    EventTableConfig
	LastEvtID int64
	Alive     bool
	ErrStr    string
	X, Y      int
}

// EventTableMonitor checks that a status column (e.g. cg_proc_status'
// last_evt_id) increases between two reads one wait apart, proving the ETL
// behind it is making progress.
type EventTableMonitor struct {
	statuses []*EventTableStatus
	position Position
	interval time.Duration
	connect  Connector
	// wait pauses between the two column reads of one check. Tests replace
	// it to advance the column instead of sleeping.
	wait func(ctx context.Context)
}

// NewEventTableMonitor creates a new event table monitor.
func NewEventTableMonitor(databases []EventTableConfig, baseY int, iv Intervals) *EventTableMonitor {
	m := &EventTableMonitor{
		statuses: make([]*EventTableStatus, len(databases)),
		position: Position{X: 0, Y: baseY},
		interval: iv.EventTable,
		connect:  ConnectPostgres,
	}
	m.wait = func(ctx context.Context) { sleepCtx(ctx, iv.EventTableWait) }

	// Initialize statuses - entries start directly at baseY (no header)
	for i, db := range databases {
		m.statuses[i] = &EventTableStatus{
			Config: db,
			X:      1,
			Y:      baseY + i,
		}
	}

	return m
}

// Name returns the monitor name.
func (m *EventTableMonitor) Name() string {
	return "Event Table Monitor"
}

// Start begins monitoring.
func (m *EventTableMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check performs a check cycle.
func (m *EventTableMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))

	for _, status := range m.statuses {
		go func(status *EventTableStatus) {
			defer wg.Done()
			m.checkDatabase(ctx, status, errorChan)
		}(status)
	}

	wg.Wait()
	m.display(disp)
}

// checkDatabase checks a single event table database.
func (m *EventTableMonitor) checkDatabase(ctx context.Context, status *EventTableStatus, errorChan chan<- string) {
	status.ErrStr = ""
	status.Alive = false

	dbobj, err := m.connect(ctx, status.Config.DatabaseConfig)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}
	defer func() { _ = dbobj.Close(ctx) }() // best-effort close of per-check connection

	// Column/table names come from operator-provided monitor config.
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT 1", status.Config.ColumnName, status.Config.TableName)

	var val1 int64
	err = dbobj.QueryRow(ctx, query).Scan(&val1)
	if err != nil {
		if isNoRows(err) {
			return
		}
		status.ErrStr = fmt.Sprintf("Error %v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	m.wait(ctx)

	var val2 int64
	err = dbobj.QueryRow(ctx, query).Scan(&val2)
	if err != nil {
		if isNoRows(err) {
			return
		}
		status.ErrStr = fmt.Sprintf("Error %v", err)
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	diff := val2 - val1
	if diff <= 0 {
		status.ErrStr = fmt.Sprintf("%s not increasing (%s = %v)", status.Config.ColumnName, status.Config.ColumnName, val2)
		sendErr(ctx, errorChan, status.ErrStr)
	} else {
		status.Alive = true
	}
	status.LastEvtID = val2
}

// display renders the event table status.
func (m *EventTableMonitor) display(disp Display) {
	// No separate header - entries appear in the same section as Application monitors

	// Each database (matching Application monitor format)
	for _, status := range m.statuses {
		y := status.Y

		// Event ID (first column - same position as block number in Application monitor)
		disp.DrawText(Position{X: status.X, Y: y}, fmt.Sprintf("%9d", status.LastEvtID),
			ColorBlue, ColorDefault)

		// Status indicator (second column - same position as lag in Application monitor)
		statusStr := " Alive"
		color := ColorGreen
		if !status.Alive {
			statusStr = " DOWN "
			color = ColorRed
		}
		disp.DrawText(Position{X: status.X + 10, Y: y}, statusStr, color, ColorDefault)

		// Name (third column - same position as name in Application monitor)
		disp.DrawText(Position{X: status.X + 20, Y: y}, status.Config.Name, ColorWhite, ColorDefault)
	}

	disp.Flush()
}
