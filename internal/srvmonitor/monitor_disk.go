package srvmonitor

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"
)

// DiskStatus holds status for disk usage.
type DiskStatus struct {
	Config DiskConfig
	Output string
	ErrStr string
	X, Y   int
}

// DiskMonitor reports remote disk usage by running df over ssh.
type DiskMonitor struct {
	statuses []*DiskStatus
	position Position
	logger   *slog.Logger
	interval time.Duration
	run      CommandRunner
}

// NewDiskMonitor creates a new disk monitor.
func NewDiskMonitor(disks []DiskConfig, logger *slog.Logger, iv Intervals) *DiskMonitor {
	m := &DiskMonitor{
		statuses: make([]*DiskStatus, len(disks)),
		position: Position{X: 1, Y: 17},
		logger:   logger,
		interval: iv.Disk,
		// df output is parsed line by line; use the stdout-only runner so
		// ssh banners on stderr cannot corrupt it.
		run: runStdout,
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
			Y:      19, // Below the section header at Y=17
		}
	}

	return m
}

// Name returns the monitor name.
func (m *DiskMonitor) Name() string {
	return "Disk Usage Monitor"
}

// Start begins monitoring.
func (m *DiskMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check performs a check cycle.
func (m *DiskMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	m.logger.Info(fmt.Sprintf("Disk monitor: Starting check cycle for %d disk(s)", len(m.statuses)))

	var wg sync.WaitGroup
	wg.Add(len(m.statuses))

	for _, status := range m.statuses {
		go func(status *DiskStatus) {
			defer wg.Done()
			m.checkDisk(ctx, status, errorChan)
		}(status)
	}

	wg.Wait()

	m.logger.Info("Disk monitor: Check cycle complete, updating display")

	m.display(disp)
}

// checkDisk checks disk usage on a server.
func (m *DiskMonitor) checkDisk(ctx context.Context, status *DiskStatus, errorChan chan<- string) {
	output, err := m.run(ctx, "/usr/bin/ssh",
		"-o", "StrictHostKeyChecking=accept-new",
		"-l", status.Config.User,
		status.Config.IP,
		"df --output=target,pcent", status.Config.DeviceList)
	if err != nil {
		status.ErrStr = err.Error()
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	status.Output = string(output)
	status.ErrStr = ""
}

// display renders the disk usage.
func (m *DiskMonitor) display(disp Display) {
	// Draw header (like other monitors)
	disp.DrawText(Position{X: 0, Y: 17},
		"----------- Disk Usage (df) ----------",
		ColorWhite, ColorDefault)

	for _, status := range m.statuses {
		// Title
		disp.DrawText(Position{X: status.X + 3, Y: status.Y}, status.Config.Title,
			ColorYellow, ColorDefault)

		// Output lines or error
		if status.ErrStr != "" {
			disp.DrawText(Position{X: status.X, Y: status.Y + 1}, "ERROR: "+status.ErrStr,
				ColorRed, ColorDefault)
		} else if status.Output != "" {
			lines := strings.Split(status.Output, "\n")
			for i := 1; i < len(lines)-1; i++ {
				line := lines[i]
				disp.DrawText(Position{X: status.X, Y: status.Y + i}, line,
					ColorWhite, ColorDefault)
			}
		} else {
			disp.DrawText(Position{X: status.X, Y: status.Y + 1}, "Checking...",
				ColorYellow, ColorDefault)
		}
	}

	disp.Flush()
}
