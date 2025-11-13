// Server Monitoring Application - Refactored Architecture
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	
	"github.com/nsf/termbox-go"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/config"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/monitor"
	"github.com/PredictionExplorer/augur-explorer/etl/cmd/srvmonitor/types"
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
	logfile, err := os.OpenFile("/tmp/srvmonitor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		os.Exit(1)
	}
	logger := log.New(logfile, "INFO: ", log.Ltime|log.Lshortfile)
	defer os.Rename("/tmp/srvmonitor.log", "/tmp/srvmonitor-old.log")
	
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
	logger.Printf("  - %d Application databases", len(cfg.ApplicationDBs))
	logger.Printf("  - %d Web APIs", len(cfg.WebAPIs))
	logger.Printf("  - %d Disk monitors", len(cfg.DiskMonitors))
	
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
	
	// Create monitor manager
	mgr := monitor.NewManager(disp, logger)
	
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
	
	// Register Web API monitor
	if len(cfg.WebAPIs) > 0 {
		webMon := monitor.NewWebAPIMonitor(cfg.WebAPIs)
		mgr.Register(webMon)
		logger.Printf("Registered: %s", webMon.Name())
	}
	
	// Register Disk monitor
	if len(cfg.DiskMonitors) > 0 {
		diskMon := monitor.NewDiskMonitor(cfg.DiskMonitors)
		mgr.Register(diskMon)
		logger.Printf("Registered: %s", diskMon.Name())
	}
	
	// Register Image monitor
	if cfg.RWalkDB.Host != "" && cfg.RWalkImage.URL != "" {
		imgMon := monitor.NewImageMonitor(cfg.RWalkImage, cfg.RWalkDB)
		mgr.Register(imgMon)
		logger.Printf("Registered: %s", imgMon.Name())
	}
	
	// Setup signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGWINCH)
	
	go func() {
		for sig := range sigChan {
			switch sig {
			case syscall.SIGWINCH:
				// Window resize - termbox will handle via EventResize
				logger.Printf("Window resize signal received")
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
