// Command srvmonitor is an interactive terminal dashboard monitoring the
// operational estate: RPC nodes, PostgreSQL databases, ETL progress, web
// APIs, disks, SSL certificates, image servers and web-server anomalies.
// All monitoring logic lives in internal/srvmonitor; this binary wires the
// environment configuration, the termbox UI and signal handling together.
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

	"github.com/PredictionExplorer/augur-explorer/internal/srvmonitor"
	"github.com/PredictionExplorer/augur-explorer/internal/srvmonitor/termboxui"
)

// setupResult carries everything main needs from the testable setup phase.
type setupResult struct {
	logger  *log.Logger
	cfg     *srvmonitor.Config
	tmpDir  string
	logPath string
	oldPath string
}

// setup opens the monitor log and loads the configuration. It is separated
// from main so the pre-terminal phase is testable.
func setup(getenv func(string) string) (*setupResult, error) {
	tmpDir := getenv("TMPDIR")
	if tmpDir == "" {
		tmpDir = "/tmp"
	}
	tmpDir = filepath.Clean(tmpDir)
	logFilePath := filepath.Join(tmpDir, "srvmonitor.log")
	oldLogFilePath := filepath.Join(tmpDir, "srvmonitor-old.log")

	logfile, err := os.OpenFile(filepath.Clean(logFilePath), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %w", err)
	}
	logger := log.New(logfile, "INFO: ", log.Ltime|log.Lshortfile)

	logger.Printf("=== Server Monitor Starting ===")

	// Load configuration from environment variables
	cfg, err := srvmonitor.LoadFromEnv(getenv)
	if err != nil {
		logger.Printf("Configuration error: %v", err)
		return nil, fmt.Errorf("configuration error: %w", err)
	}

	logConfigSummary(logger, cfg)

	return &setupResult{
		logger:  logger,
		cfg:     cfg,
		tmpDir:  tmpDir,
		logPath: logFilePath,
		oldPath: oldLogFilePath,
	}, nil
}

func main() {
	// Setup panic recovery
	defer func() {
		if r := recover(); r != nil {
			log.Printf("PANIC: %v\n", r)
			termbox.Close()
			os.Exit(1)
		}
	}()

	s, err := setup(os.Getenv)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	logger := s.logger
	defer func() { _ = os.Rename(s.logPath, s.oldPath) }() // best-effort log rotation on exit

	// Initialize display
	disp := termboxui.New()
	if err := disp.Init(); err != nil {
		logger.Printf("Failed to initialize display: %v", err)
		fmt.Printf("Failed to initialize display: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = disp.Close() }() // termbox Close never returns an error

	logger.Printf("Display initialized")
	cfg, tmpDir := s.cfg, s.tmpDir

	// Build the manager with every configured monitor registered
	mgr, layout := srvmonitor.BuildManager(cfg, disp, logger, tmpDir)
	logger.Printf("Layout positions: EventTable=%d, WebAPI=%d, Image=%d, Anomaly=%d, SSL=(%d,%d), Errors=(%d,%d)",
		layout.EventTableBaseY, layout.WebAPIBaseY, layout.ImageBaseY, layout.AnomalyBaseY,
		layout.SSLBaseX, layout.SSLBaseY, layout.ErrorX, layout.ErrorY)

	// Setup signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGWINCH, syscall.SIGUSR1)
	go handleSignals(ctx, sigChan, mgr, logger, cancel)

	// Start all monitors and the alarm-state cleanup
	logger.Printf("Starting all monitors...")
	mgr.Start(ctx)

	logger.Printf("Application fully started, entering event loop")

	// Feed termbox events to the loop; the goroutine unblocks (and the
	// process exits) once the loop returns.
	events := make(chan termbox.Event)
	go func() {
		for {
			ev := termbox.PollEvent()
			select {
			case events <- ev:
			case <-ctx.Done():
				return
			}
		}
	}()

	runEventLoop(ctx, events, disp, logger, cancel)
}

// logConfigSummary writes the loaded configuration counts to the log.
func logConfigSummary(logger *log.Logger, cfg *srvmonitor.Config) {
	logger.Printf("Configuration loaded successfully:")
	logger.Printf("  - %d RPC nodes", len(cfg.RPCNodes))
	logger.Printf("  - %d Layer1 databases", len(cfg.Layer1DBs))
	logger.Printf("  - %d Event Table databases", len(cfg.EventTableDBs))
	logger.Printf("  - %d Application databases", len(cfg.ApplicationDBs))
	logger.Printf("  - %d Web APIs", len(cfg.WebAPIs))
	logger.Printf("  - %d Disk monitors", len(cfg.DiskMonitors))
	logger.Printf("  - %d SSL certificates", len(cfg.SSLCerts))
	if cfg.Anomaly.Enabled() {
		logger.Printf("  - Anomaly monitor: %s@%s:%s", cfg.Anomaly.User, cfg.Anomaly.Host, cfg.Anomaly.RemoteFile)
	} else {
		logger.Printf("  - Anomaly monitor: DISABLED (set ANOMALY_SSH_USER, ANOMALY_SSH_HOST, ANOMALY_REMOTE_FILE)")
	}
	logger.Printf("  - Mobile notifications: %v", cfg.MobileNotif)
}

// notifier abstracts the manager for the signal handler.
type notifier interface {
	SendTestNotification(ctx context.Context)
}

// handleSignals reacts to process signals: SIGWINCH is informational (termbox
// delivers a resize event), SIGUSR1 triggers a test notification, and
// SIGINT/SIGTERM cancel the run.
func handleSignals(ctx context.Context, sigChan <-chan os.Signal, mgr notifier, logger *log.Logger, cancel context.CancelFunc) {
	for {
		select {
		case <-ctx.Done():
			return
		case sig := <-sigChan:
			switch sig {
			case syscall.SIGWINCH:
				logger.Printf("Window resize signal received")
			case syscall.SIGUSR1:
				logger.Printf("SIGUSR1 signal received, sending test notification")
				mgr.SendTestNotification(ctx)
			case os.Interrupt, syscall.SIGTERM:
				logger.Printf("Termination signal received, exiting")
				cancel()
			}
		}
	}
}

// runEventLoop processes termbox events until the context is cancelled or an
// exit key arrives: 'q'/Ctrl+C quit, resize repaints, error events are
// logged.
func runEventLoop(ctx context.Context, events <-chan termbox.Event, disp srvmonitor.Display, logger *log.Logger, cancel context.CancelFunc) {
	for {
		select {
		case <-ctx.Done():
			logger.Printf("Context cancelled, exiting event loop")
			return
		case ev := <-events:
			switch ev.Type {
			case termbox.EventKey:
				// Exit on Ctrl+C or 'q' key
				if ev.Key == termbox.KeyCtrlC || ev.Ch == 'q' {
					logger.Printf("Exit key pressed")
					cancel()
					return
				}
			case termbox.EventResize:
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
				disp.DrawText(srvmonitor.Position{X: x, Y: y}, msg, srvmonitor.ColorYellow, srvmonitor.ColorDefault)
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
		}
	}
}
