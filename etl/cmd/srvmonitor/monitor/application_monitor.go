package monitor

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
	
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/utils"
)

// ApplicationStatus holds status for application layer
type ApplicationStatus struct {
	Config          types.DatabaseConfig
	TableName       string
	LastBlockNum    int64
	OfficialLagDiff int64
	ErrStr          string
	ChainID         string
	X, Y            int
}

// ApplicationMonitor monitors application layer databases
type ApplicationMonitor struct {
	apps        []ApplicationStatus
	sharedState *SharedRPCState
	position    types.Position
}

// NewApplicationMonitor creates a new application monitor
func NewApplicationMonitor(dbs []types.DatabaseConfig, sharedState *SharedRPCState) *ApplicationMonitor {
	m := &ApplicationMonitor{
		apps:        make([]ApplicationStatus, 0),
		sharedState: sharedState,
		position:    types.Position{X: 1, Y: 24},
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

// Name returns the monitor name
func (m *ApplicationMonitor) Name() string {
	return "Application Layer Monitor"
}

// GetDisplayPosition returns the display position
func (m *ApplicationMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *ApplicationMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
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
func (m *ApplicationMonitor) check(disp display.Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.apps))
	
	for i := range m.apps {
		go m.checkApp(&m.apps[i], &wg, errorChan)
	}
	
	wg.Wait()
	m.display(disp)
}

// checkApp checks a single application database
func (m *ApplicationMonitor) checkApp(status *ApplicationStatus, wg *sync.WaitGroup, errorChan chan<- string) {
	defer wg.Done()
	
	status.ErrStr = ""
	
	dbobj, err := utils.ConnectPostgres(status.Config)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v", err)
		errorChan <- status.ErrStr
		return
	}
	defer dbobj.Close()
	
	var chainIDStr string
	err = dbobj.QueryRow("SELECT chain_id FROM contract_addresses").Scan(&chainIDStr)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v", err)
		errorChan <- status.ErrStr
		return
	}
	status.ChainID = chainIDStr
	
	var lastEvtID int64
	err = dbobj.QueryRow("SELECT last_evt_id FROM " + status.TableName).Scan(&lastEvtID)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v", err)
		errorChan <- status.ErrStr
		return
	}
	
	var bnum int64
	err = dbobj.QueryRow("SELECT block_num FROM evt_log WHERE id=$1", lastEvtID).Scan(&bnum)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v", err)
		errorChan <- status.ErrStr
		return
	}
	
	status.LastBlockNum = bnum
	
	// Calculate lag using shared state
	if m.sharedState != nil {
		switch chainIDStr {
		case "1":
			officialBlock := m.sharedState.GetOfficialMainnet()
			if officialBlock != 0 {
				status.OfficialLagDiff = officialBlock - status.LastBlockNum
			}
		case "11155111":
			officialBlock := m.sharedState.GetOfficialSepolia()
			if officialBlock != 0 {
				status.OfficialLagDiff = officialBlock - status.LastBlockNum
			}
		case "42161":
			officialBlock := m.sharedState.GetOfficialArbitrum()
			if officialBlock != 0 {
				status.OfficialLagDiff = officialBlock - status.LastBlockNum
			}
		case "421614":
			officialBlock := m.sharedState.GetOfficialSepoliaArb()
			if officialBlock != 0 {
				status.OfficialLagDiff = officialBlock - status.LastBlockNum
			}
		}
	}
}

// display renders the application status
func (m *ApplicationMonitor) display(disp display.Display) {
	// Header
	disp.DrawText(types.Position{X: 1, Y: 24},
		"----------- Last Block Numbers in Postgres ----------",
		types.ColorWhite, types.ColorDefault)
	
	// Each application
	for _, status := range m.apps {
		y := status.Y
		
		// Title
		disp.DrawText(types.Position{X: status.X, Y: y}, status.Config.Name, types.ColorWhite, types.ColorDefault)
		
		// Block number
		disp.DrawText(types.Position{X: status.X + 35, Y: y}, fmt.Sprintf("%9d", status.LastBlockNum),
			types.ColorBlue, types.ColorDefault)
		
		// Official lag
		officialDiff := "------"
		if status.OfficialLagDiff != math.MaxInt64 {
			officialDiff = fmt.Sprintf("%6v", status.OfficialLagDiff)
		}
		disp.DrawText(types.Position{X: status.X + 45, Y: y}, officialDiff, types.ColorBlue, types.ColorDefault)
	}
	
	disp.Flush()
}

