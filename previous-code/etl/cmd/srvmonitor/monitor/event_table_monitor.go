package monitor

import (
	"context"
	"fmt"
	"sync"
	"time"
	
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/types"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/utils"
)

const (
	EventTableWait     = 120 // seconds to wait before second query (2 minutes)
	UpdateIntervalEvt  = 60  // seconds between check cycles
)

// EventTableStatus holds status for a single event table database
type EventTableStatus struct {
	Config     types.EventTableConfig
	LastEvtID  int64
	Alive      bool
	ErrStr     string
	X, Y       int
}

// EventTableMonitor monitors event table databases for last_evt_id progression
type EventTableMonitor struct {
	databases []types.EventTableConfig
	statuses  []*EventTableStatus
	position  types.Position
	baseY     int
}

// NewEventTableMonitor creates a new event table monitor
func NewEventTableMonitor(databases []types.EventTableConfig, baseY int) *EventTableMonitor {
	m := &EventTableMonitor{
		databases: databases,
		statuses:  make([]*EventTableStatus, len(databases)),
		position:  types.Position{X: 0, Y: baseY},
		baseY:     baseY,
	}
	
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

// Name returns the monitor name
func (m *EventTableMonitor) Name() string {
	return "Event Table Monitor"
}

// GetDisplayPosition returns the display position
func (m *EventTableMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *EventTableMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.check(disp, errorChan)
			time.Sleep(UpdateIntervalEvt * time.Second)
		}
	}
}

// check performs a check cycle
func (m *EventTableMonitor) check(disp display.Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))
	
	for _, status := range m.statuses {
		go m.checkDatabase(status, &wg, errorChan)
	}
	
	wg.Wait()
	m.display(disp)
}

// checkDatabase checks a single event table database
func (m *EventTableMonitor) checkDatabase(status *EventTableStatus, wg *sync.WaitGroup, errorChan chan<- string) {
	defer wg.Done()
	
	status.ErrStr = ""
	status.Alive = false
	
	dbobj, err := utils.ConnectPostgresEventTable(status.Config)
	if err != nil {
		status.ErrStr = fmt.Sprintf("%v", err)
		errorChan <- status.ErrStr
		return
	}
	defer dbobj.Close()
	
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT 1", status.Config.ColumnName, status.Config.TableName)
	
	var val1 int64
	err = dbobj.QueryRow(query).Scan(&val1)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v", err)
		errorChan <- status.ErrStr
		return
	}
	
	time.Sleep(EventTableWait * time.Second)
	
	var val2 int64
	err = dbobj.QueryRow(query).Scan(&val2)
	if err != nil {
		status.ErrStr = fmt.Sprintf("Error %v", err)
		errorChan <- status.ErrStr
		return
	}
	
	diff := val2 - val1
	if diff <= 0 {
		status.ErrStr = fmt.Sprintf("%s not increasing (%s = %v)", status.Config.ColumnName, status.Config.ColumnName, val2)
		errorChan <- status.ErrStr
	} else {
		status.Alive = true
	}
	status.LastEvtID = val2
}

// display renders the event table status
func (m *EventTableMonitor) display(disp display.Display) {
	// No separate header - entries appear in the same section as Application monitors
	
	// Each database (matching Application monitor format)
	for _, status := range m.statuses {
		y := status.Y
		
		// Event ID (first column - same position as block number in Application monitor)
		disp.DrawText(types.Position{X: status.X, Y: y}, fmt.Sprintf("%9d", status.LastEvtID),
			types.ColorBlue, types.ColorDefault)
		
		// Status indicator (second column - same position as lag in Application monitor)
		statusStr := " Alive"
		color := types.ColorGreen
		if !status.Alive {
			statusStr = " DOWN "
			color = types.ColorRed
		}
		disp.DrawText(types.Position{X: status.X + 10, Y: y}, statusStr, color, types.ColorDefault)
		
		// Name (third column - same position as name in Application monitor)
		disp.DrawText(types.Position{X: status.X + 20, Y: y}, status.Config.Name, types.ColorWhite, types.ColorDefault)
	}
	
	disp.Flush()
}
