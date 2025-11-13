package monitor

import (
	"context"
	"fmt"
	"sync"
	"time"
	
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/utils"
)

const (
	DBBlockWait     = 60 // seconds to wait before second query
	UpdateIntervalDB = 60 // seconds between checks
)

// DatabaseStatus holds status for a single database
type DatabaseStatus struct {
	Config       types.DatabaseConfig
	LastBlockNum int64
	Alive        bool
	ErrStr       string
	X, Y         int
}

// DatabaseMonitor monitors Layer 1 databases
type DatabaseMonitor struct {
	databases []types.DatabaseConfig
	statuses  []*DatabaseStatus
	position  types.Position
}

// NewDatabaseMonitor creates a new database monitor
func NewDatabaseMonitor(databases []types.DatabaseConfig) *DatabaseMonitor {
	m := &DatabaseMonitor{
		databases: databases,
		statuses:  make([]*DatabaseStatus, len(databases)),
		position:  types.Position{X: 0, Y: 11},
	}
	
	// Initialize statuses
	for i, db := range databases {
		m.statuses[i] = &DatabaseStatus{
			Config: db,
			X:      1,
			Y:      13 + i,
		}
	}
	
	return m
}

// Name returns the monitor name
func (m *DatabaseMonitor) Name() string {
	return "Database Monitor"
}

// GetDisplayPosition returns the display position
func (m *DatabaseMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *DatabaseMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.check(disp, errorChan)
			time.Sleep(UpdateIntervalDB * time.Second)
		}
	}
}

// check performs a check cycle
func (m *DatabaseMonitor) check(disp display.Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))
	
	for _, status := range m.statuses {
		go m.checkDatabase(status, &wg, errorChan)
	}
	
	wg.Wait()
	m.display(disp)
}

// checkDatabase checks a single database
func (m *DatabaseMonitor) checkDatabase(status *DatabaseStatus, wg *sync.WaitGroup, errorChan chan<- string) {
	defer wg.Done()
	
	status.ErrStr = ""
	status.Alive = false
	
	dbobj, err := utils.ConnectPostgres(status.Config)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v", err)
		errorChan <- status.ErrStr
		return
	}
	defer dbobj.Close()
	
	var bnum1 int64
	err = dbobj.QueryRow("SELECT block_num FROM block ORDER BY block_num DESC LIMIT 1").Scan(&bnum1)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v", err)
		errorChan <- status.ErrStr
		return
	}
	
	time.Sleep(DBBlockWait * time.Second)
	
	var bnum2 int64
	err = dbobj.QueryRow("SELECT block_num FROM block ORDER BY block_num DESC LIMIT 1").Scan(&bnum2)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v", err)
		errorChan <- status.ErrStr
		return
	}
	
	diff := bnum2 - bnum1
	if diff == 0 {
		status.ErrStr = fmt.Sprintf("Block difference is zero (last block = %v)", bnum2)
		errorChan <- status.ErrStr
	} else {
		status.Alive = true
	}
	status.LastBlockNum = bnum2
}

// display renders the database status
func (m *DatabaseMonitor) display(disp display.Display) {
	// Header
	disp.DrawText(types.Position{X: 0, Y: 11},
		"--------------------- SQL DB --------------------------------",
		types.ColorWhite, types.ColorDefault)
	
	// Each database
	for _, status := range m.statuses {
		y := status.Y
		
		// Name
		disp.DrawText(types.Position{X: status.X, Y: y}, status.Config.Name, types.ColorWhite, types.ColorDefault)
		
		// Host
		disp.DrawText(types.Position{X: status.X + 35, Y: y}, status.Config.Host, types.ColorWhite, types.ColorDefault)
		
		// Status
		aliveStr := "Alive"
		color := types.ColorGreen
		if !status.Alive {
			aliveStr = "DOWN "
			color = types.ColorRed
		}
		disp.DrawText(types.Position{X: status.X + 60, Y: y}, aliveStr, color, types.ColorDefault)
		
		// Block number
		disp.DrawText(types.Position{X: status.X + 70, Y: y}, fmt.Sprintf("%v", status.LastBlockNum),
			types.ColorBlue, types.ColorDefault)
	}
	
	disp.Flush()
}


