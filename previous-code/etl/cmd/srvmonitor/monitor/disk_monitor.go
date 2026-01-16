package monitor

import (
	"context"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"
	
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
)

const (
	UpdateIntervalDisk = 600 // 10 minutes in seconds
)

// DiskStatus holds status for disk usage
type DiskStatus struct {
	Config types.DiskConfig
	Output string
	ErrStr string
	X, Y   int
}

// DiskMonitor monitors disk usage via SSH
type DiskMonitor struct {
	disks    []types.DiskConfig
	statuses []*DiskStatus
	position types.Position
	logger   *log.Logger
}

// NewDiskMonitor creates a new disk monitor
func NewDiskMonitor(disks []types.DiskConfig, logger *log.Logger) *DiskMonitor {
	m := &DiskMonitor{
		disks:    disks,
		statuses: make([]*DiskStatus, len(disks)),
		position: types.Position{X: 1, Y: 17},
		logger:   logger,
	}
	
	// Initialize statuses with different X positions
	xPositions := []int{1, 25, 50}
	for i, disk := range disks {
		x := 1
		if i < len(xPositions) {
			x = xPositions[i]
		}
		
		m.statuses[i] = &DiskStatus{
			Config: disk,
			X:      x,
			Y:      19, // Changed from 18 to make room for header
		}
	}
	
	return m
}

// Name returns the monitor name
func (m *DiskMonitor) Name() string {
	return "Disk Usage Monitor"
}

// GetDisplayPosition returns the display position
func (m *DiskMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *DiskMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
	// Do initial check immediately
	m.check(disp, errorChan)
	
	// Create ticker for periodic checks
	ticker := time.NewTicker(UpdateIntervalDisk * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			m.check(disp, errorChan)
		}
	}
}

// check performs a check cycle
func (m *DiskMonitor) check(disp display.Display, errorChan chan<- string) {
	if m.logger != nil {
		m.logger.Printf("Disk monitor: Starting check cycle for %d disk(s)", len(m.statuses))
	}
	
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))
	
	for _, status := range m.statuses {
		go m.checkDisk(status, &wg, errorChan)
	}
	
	wg.Wait()
	
	if m.logger != nil {
		m.logger.Printf("Disk monitor: Check cycle complete, updating display")
	}
	
	m.display(disp)
}

// checkDisk checks disk usage on a server
func (m *DiskMonitor) checkDisk(status *DiskStatus, wg *sync.WaitGroup, errorChan chan<- string) {
	defer wg.Done()
	
	cmd := exec.Command("/usr/bin/ssh",
		"-o", "StrictHostKeyChecking=accept-new",
		"-l", status.Config.User,
		status.Config.IP,
		"df --output=target,pcent", status.Config.DeviceList)
	
	output, err := cmd.Output()
	if err != nil {
		status.ErrStr = err.Error()
		errorChan <- status.ErrStr
		return
	}
	
	status.Output = string(output)
	status.ErrStr = ""
}

// display renders the disk usage
func (m *DiskMonitor) display(disp display.Display) {
	// Draw header (like other monitors)
	disp.DrawText(types.Position{X: 0, Y: 17},
		"----------- Disk Usage (df) ----------",
		types.ColorWhite, types.ColorDefault)
	
	for _, status := range m.statuses {
		// Title
		disp.DrawText(types.Position{X: status.X + 3, Y: status.Y}, status.Config.Title,
			types.ColorYellow, types.ColorDefault)
		
		// Output lines or error
		if status.ErrStr != "" {
			disp.DrawText(types.Position{X: status.X, Y: status.Y + 1}, "ERROR: "+status.ErrStr,
				types.ColorRed, types.ColorDefault)
		} else if status.Output != "" {
			lines := strings.Split(status.Output, "\n")
			for i := 1; i < len(lines)-1; i++ {
				line := lines[i]
				disp.DrawText(types.Position{X: status.X, Y: status.Y + i}, line,
					types.ColorWhite, types.ColorDefault)
			}
		} else {
			disp.DrawText(types.Position{X: status.X, Y: status.Y + 1}, "Checking...",
				types.ColorYellow, types.ColorDefault)
		}
	}
	
	disp.Flush()
}


