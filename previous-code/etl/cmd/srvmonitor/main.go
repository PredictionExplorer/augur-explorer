// Server Monitoring Application - Refactored Architecture
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	
	"github.com/nsf/termbox-go"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/config"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/monitor"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/types"
)

func main() {
	// Setup panic recovery
	defer func() {
		if r := recover(); r != nil {
			log.Printf("PANIC: %v\n", r)
			termbox.Close()
			os.Exit(1)
		}
	}()
	
	// Setup logging
	tmpDir := os.Getenv("TMPDIR")
	if tmpDir == "" {
		tmpDir = "/tmp"
	}
	logFilePath := filepath.Join(tmpDir, "srvmonitor.log")
	oldLogFilePath := filepath.Join(tmpDir, "srvmonitor-old.log")
	
	logfile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		os.Exit(1)
	}
	logger := log.New(logfile, "INFO: ", log.Ltime|log.Lshortfile)
	defer os.Rename(logFilePath, oldLogFilePath)
	
	logger.Printf("=== Server Monitor Starting ===")
	
	// Load configuration from environment variables
	cfg, err := config.LoadFromEnv()
	if err != nil {
		logger.Printf("Configuration error: %v", err)
		fmt.Printf("Configuration error: %v\n", err)
		os.Exit(1)
	}
	
	logger.Printf("Configuration loaded successfully:")
	logger.Printf("  - %d RPC nodes", len(cfg.RPCNodes))
	logger.Printf("  - %d Layer1 databases", len(cfg.Layer1DBs))
	logger.Printf("  - %d Event Table databases", len(cfg.EventTableDBs))
	logger.Printf("  - %d Application databases", len(cfg.ApplicationDBs))
	logger.Printf("  - %d Web APIs", len(cfg.WebAPIs))
	logger.Printf("  - %d Disk monitors", len(cfg.DiskMonitors))
	logger.Printf("  - Mobile notifications: %v", cfg.MobileNotif)
	
	// Initialize display
	disp := display.NewTermboxDisplay()
	if err := disp.Init(); err != nil {
		logger.Printf("Failed to initialize display: %v", err)
		fmt.Printf("Failed to initialize display: %v\n", err)
		os.Exit(1)
	}
	defer disp.Close()
	
	logger.Printf("Display initialized")
	
	// Create shared RPC state for lag calculations
	sharedRPCState := monitor.NewSharedRPCState()
	
	// Calculate dynamic Y positions for each section
	// Layout:
	// - RPC Nodes: Y=0 header, Y=1+ entries
	// - SQL DB (Layer1): Y=11 header, Y=13+ entries
	// - Disk Usage: Y=17 header, Y=19+ content
	// - Last Block Numbers in Postgres (Application + Event Table): Y=24 header, Y=25+ entries
	// - Web API: dynamic based on Application + Event Table count
	// - Image: dynamic based on Web API count
	// - Errors: dynamic based on Image position
	
	// Application Monitor: header at Y=24, entries at Y=25+i
	// Event Table entries continue at Y=25+len(ApplicationDBs)
	eventTableBaseY := 25 + len(cfg.ApplicationDBs)
	
	// Web API starts after Application DBs + Event Table DBs + 1 line gap
	webAPIBaseY := 25 + len(cfg.ApplicationDBs) + len(cfg.EventTableDBs) + 1
	
	// Image starts after Web API header + entries + 1 line gap
	imageBaseY := webAPIBaseY + 1 + len(cfg.WebAPIs) + 1
	
	// Error display starts after Image section (header + 2 content lines + 1 gap)
	errorDisplayY := imageBaseY + 3 + 1
	
	logger.Printf("Layout positions: EventTable=%d, WebAPI=%d, Image=%d, Errors=%d",
		eventTableBaseY, webAPIBaseY, imageBaseY, errorDisplayY)
	
	// Create monitor manager
	mgr := monitor.NewManager(disp, logger, cfg.MobileNotif, errorDisplayY)
	
	// Register RPC monitor
	officialNames := map[string]string{
		"mainnet":     cfg.OfficialRPCMainnet,
		"arbitrum":    cfg.OfficialRPCArbitrum,
		"sepolia":     cfg.OfficialRPCSepolia,
		"sepolia_arb": cfg.OfficialRPCSepoliaArb,
	}
	rpcMon := monitor.NewRPCMonitor(cfg.RPCNodes, officialNames, sharedRPCState)
	mgr.Register(rpcMon)
	logger.Printf("Registered: %s", rpcMon.Name())
	
	// Register Database monitor
	if len(cfg.Layer1DBs) > 0 {
		dbMon := monitor.NewDatabaseMonitor(cfg.Layer1DBs)
		mgr.Register(dbMon)
		logger.Printf("Registered: %s", dbMon.Name())
	}
	
	// Register Application Layer monitor
	if len(cfg.ApplicationDBs) > 0 {
		appMon := monitor.NewApplicationMonitor(cfg.ApplicationDBs, sharedRPCState)
		mgr.Register(appMon)
		logger.Printf("Registered: %s", appMon.Name())
	}
	
	// Register Event Table monitor (after Application DBs in "Last Block Numbers in Postgres" section)
	if len(cfg.EventTableDBs) > 0 {
		evtMon := monitor.NewEventTableMonitor(cfg.EventTableDBs, eventTableBaseY)
		mgr.Register(evtMon)
		logger.Printf("Registered: %s", evtMon.Name())
	}
	
	// Register Web API monitor
	if len(cfg.WebAPIs) > 0 {
		webMon := monitor.NewWebAPIMonitor(cfg.WebAPIs, webAPIBaseY)
		mgr.Register(webMon)
		logger.Printf("Registered: %s", webMon.Name())
	}
	
	// Register Disk monitor
	if len(cfg.DiskMonitors) > 0 {
		diskMon := monitor.NewDiskMonitor(cfg.DiskMonitors, logger)
		mgr.Register(diskMon)
		logger.Printf("Registered: %s", diskMon.Name())
	}
	
	// Register Image monitor
	if cfg.RWalkDB.Host != "" && cfg.RWalkImage.URL != "" {
		imgMon := monitor.NewImageMonitor(cfg.RWalkImage, cfg.RWalkDB, imageBaseY)
		mgr.Register(imgMon)
		logger.Printf("Registered: %s", imgMon.Name())
	}
	
	// Setup signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGWINCH, syscall.SIGUSR1)
	
	go func() {
		for sig := range sigChan {
			switch sig {
			case syscall.SIGWINCH:
				// Window resize - termbox will handle via EventResize
				logger.Printf("Window resize signal received")
			case syscall.SIGUSR1:
				// Test notification request
				logger.Printf("SIGUSR1 signal received, sending test notification")
				mgr.SendTestNotification()
			case os.Interrupt, syscall.SIGTERM:
				logger.Printf("Termination signal received, exiting")
				cancel()
			}
		}
	}()
	
	// Start all monitors
	logger.Printf("Starting all monitors...")
	mgr.Start(ctx)
	
	logger.Printf("Application fully started, entering event loop")
	
	// Event loop
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			// Exit on Ctrl+C or 'q' key
			if ev.Key == termbox.KeyCtrlC || ev.Ch == 'q' {
				logger.Printf("Exit key pressed")
				cancel()
				return
			}
		case termbox.EventResize:
			// Handle window resize
			logger.Printf("Resize event: w=%d, h=%d", ev.Width, ev.Height)
			disp.Clear()
			w, h := disp.Size()
			msg := "Refreshing display..."
			x := (w - len(msg)) / 2
			y := h / 2
			if x < 0 {
				x = 0
			}
			if y < 0 {
				y = 0
			}
			disp.DrawText(types.Position{X: x, Y: y}, msg, types.ColorYellow, types.ColorDefault)
			disp.Flush()
			// Clear the temporary message immediately so monitors can redraw cleanly
			disp.Clear()
			disp.Flush()
		case termbox.EventError:
			logger.Printf("Termbox error event: %v", ev.Err)
		case termbox.EventInterrupt:
			logger.Printf("Interrupt event received")
			cancel()
			return
		}
		
		// Check if context is done
		select {
		case <-ctx.Done():
			logger.Printf("Context cancelled, exiting event loop")
			return
		default:
		}
	}
}
