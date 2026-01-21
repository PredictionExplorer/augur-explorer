package monitor

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
	
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/types"
)

// WebAPIStatus holds status for a web API
type WebAPIStatus struct {
	Config types.WebAPIConfig
	Alive  bool
	ErrStr string
	X, Y   int
}

// WebAPIMonitor monitors web APIs
type WebAPIMonitor struct {
	apis     []types.WebAPIConfig
	statuses []*WebAPIStatus
	position types.Position
}

// NewWebAPIMonitor creates a new web API monitor
func NewWebAPIMonitor(apis []types.WebAPIConfig) *WebAPIMonitor {
	m := &WebAPIMonitor{
		apis:     apis,
		statuses: make([]*WebAPIStatus, len(apis)),
		position: types.Position{X: 0, Y: 30},
	}
	
	// Initialize statuses
	for i, api := range apis {
		m.statuses[i] = &WebAPIStatus{
			Config: api,
			X:      1,
			Y:      31 + i,
		}
	}
	
	return m
}

// Name returns the monitor name
func (m *WebAPIMonitor) Name() string {
	return "Web API Monitor"
}

// GetDisplayPosition returns the display position
func (m *WebAPIMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *WebAPIMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
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
func (m *WebAPIMonitor) check(disp display.Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))
	
	for _, status := range m.statuses {
		go m.checkAPI(status, &wg, errorChan)
	}
	
	wg.Wait()
	m.display(disp)
}

// checkAPI checks a single web API
func (m *WebAPIMonitor) checkAPI(status *WebAPIStatus, wg *sync.WaitGroup, errorChan chan<- string) {
	defer wg.Done()
	
	url := fmt.Sprintf("http://%s:%s%v", status.Config.Host, status.Config.Port, status.Config.URI)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	
	if err != nil {
		status.Alive = false
		status.ErrStr = err.Error()
		errorChan <- status.ErrStr
	} else {
		status.Alive = (resp.StatusCode >= 200) && (resp.StatusCode < 500)
		status.ErrStr = ""
	}
}

// display renders the web API status
func (m *WebAPIMonitor) display(disp display.Display) {
	// Header
	disp.DrawText(types.Position{X: 0, Y: 30},
		"--------------------- Web API ------------------------------",
		types.ColorWhite, types.ColorDefault)
	
	// Each API
	for _, status := range m.statuses {
		y := status.Y
		
		// Status (first column)
		aliveStr := "Alive"
		color := types.ColorGreen
		if !status.Alive {
			aliveStr = "DOWN "
			color = types.ColorRed
		}
		disp.DrawText(types.Position{X: status.X, Y: y}, aliveStr, color, types.ColorDefault)
		
		// Title (second column - shifted right)
		disp.DrawText(types.Position{X: status.X + 10, Y: y}, status.Config.Title, types.ColorWhite, types.ColorDefault)
		
		// Host:Port (third column - shifted right)
		hostPort := status.Config.Host + ":" + status.Config.Port
		disp.DrawText(types.Position{X: status.X + 35, Y: y}, hostPort, types.ColorWhite, types.ColorDefault)
	}
	
	disp.Flush()
}


