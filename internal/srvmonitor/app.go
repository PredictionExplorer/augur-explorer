package srvmonitor

import (
	"log/slog"
)

// Layout holds the computed screen positions of every display section. The
// left column stacks RPC (Y=0), SQL DB (Y=11), disk usage (Y=17), Postgres
// block numbers (Y=24, application then event-table entries), web APIs,
// images and optionally anomalies; SSL certificates form a right-hand column
// aligned with the Postgres section.
type Layout struct {
	// EventTableBaseY is the first row of the event-table entries, directly
	// below the application entries in the "Last Block Numbers" section.
	EventTableBaseY int
	// WebAPIBaseY is the header row of the Web API section.
	WebAPIBaseY int
	// ImageBaseY is the header row of the image section.
	ImageBaseY int
	// AnomalyBaseY is the header row of the anomaly section (used only when
	// anomaly monitoring is enabled).
	AnomalyBaseY int
	// SSLBaseX/SSLBaseY position the SSL certificate column.
	SSLBaseX, SSLBaseY int
	// ErrorX/ErrorY position the shared two-line error area.
	ErrorX, ErrorY int
}

// ComputeLayout derives every section position from the configured monitor
// counts.
func ComputeLayout(cfg *Config) Layout {
	l := Layout{}

	// Application Monitor: header at Y=24, entries at Y=25+i.
	// Event Table entries continue directly below the application entries.
	l.EventTableBaseY = 25 + len(cfg.ApplicationDBs)

	// Web API starts after Application DBs + Event Table DBs + 1 line gap.
	l.WebAPIBaseY = 25 + len(cfg.ApplicationDBs) + len(cfg.EventTableDBs) + 1

	// Image starts after Web API header + entries + 1 line gap.
	l.ImageBaseY = l.WebAPIBaseY + 1 + len(cfg.WebAPIs) + 1

	// WebSrv Anomalies section sits directly below the Image section
	// (Image occupies its header + 2 content lines).
	l.AnomalyBaseY = l.ImageBaseY + 3

	// Bottom of the left-hand column (Postgres / Web API / Image [/ Anomaly] stack).
	leftColumnBottomY := l.ImageBaseY + 3
	if cfg.Anomaly.Enabled() {
		// Anomaly header at AnomalyBaseY, then AnomalyDisplayCount content lines.
		leftColumnBottomY = l.AnomalyBaseY + AnomalyDisplayCount
	}

	// SSL Certificates are placed in a right-hand column, aligned with the
	// top of the "Last Block Numbers in Postgres" section (Y=24). This runs
	// the SSL list alongside the Postgres/WebAPI stack instead of stacking it
	// below the Image section, reclaiming that vertical space. The X offset
	// clears the widest left-column content (Web API host:port ends near
	// column 60).
	l.SSLBaseX = 62
	l.SSLBaseY = 24
	sslColumnBottomY := l.SSLBaseY
	if len(cfg.SSLCerts) > 0 {
		sslColumnBottomY = l.SSLBaseY + 1 + len(cfg.SSLCerts)
	}

	// Error area: by default at the bottom of the left column. When SSL
	// certs are shown, place it in the right column directly beneath the SSL
	// list, so the (mostly cert-related) warnings sit next to the
	// certificate list and the bottom-left space stays free.
	l.ErrorX = 1
	l.ErrorY = leftColumnBottomY + 1
	if len(cfg.SSLCerts) > 0 {
		l.ErrorX = l.SSLBaseX
		l.ErrorY = sslColumnBottomY + 1
	}

	return l
}

// BuildManager creates the alarm tracker and manager and registers every
// monitor the config enables, mirroring the gating rules of the legacy main:
// RPC always; databases, event tables, applications, web APIs, disks and SSL
// certificates when configured; the image monitor when both its database and
// server URL are set; the anomaly monitor when fully configured.
//
// anomalyDir is where the fetched anomaly file is stored (empty selects the
// system temp directory).
func BuildManager(cfg *Config, disp Display, logger *slog.Logger, anomalyDir string) (*Manager, Layout) {
	layout := ComputeLayout(cfg)
	iv := cfg.Intervals

	alarmTracker := NewAlarmTracker(cfg.MobileNotif, logger)
	if cfg.MobileNotif {
		logger.Info("Mobile notifications enabled")
	}
	mgr := NewManager(disp, logger, alarmTracker, layout.ErrorX, layout.ErrorY)

	// Shared RPC state for lag calculations
	sharedRPCState := NewSharedRPCState()

	officialNames := map[string]string{
		"mainnet":     cfg.OfficialRPCMainnet,
		"arbitrum":    cfg.OfficialRPCArbitrum,
		"sepolia":     cfg.OfficialRPCSepolia,
		"sepolia_arb": cfg.OfficialRPCSepoliaArb,
	}
	mgr.Register(NewRPCMonitor(cfg.RPCNodes, officialNames, sharedRPCState, iv))

	if len(cfg.Layer1DBs) > 0 {
		mgr.Register(NewDatabaseMonitor(cfg.Layer1DBs, iv))
	}

	if len(cfg.ApplicationDBs) > 0 {
		mgr.Register(NewApplicationMonitor(cfg.ApplicationDBs, sharedRPCState, iv))
	}

	if len(cfg.EventTableDBs) > 0 {
		mgr.Register(NewEventTableMonitor(cfg.EventTableDBs, layout.EventTableBaseY, iv))
	}

	if len(cfg.WebAPIs) > 0 {
		mgr.Register(NewWebAPIMonitor(cfg.WebAPIs, layout.WebAPIBaseY, iv))
	}

	if len(cfg.DiskMonitors) > 0 {
		mgr.Register(NewDiskMonitor(cfg.DiskMonitors, logger, iv))
	}

	if cfg.RWalkDB.Host != "" && cfg.RWalkImage.URL != "" {
		mgr.Register(NewImageMonitor(cfg.RWalkImage, cfg.RWalkDB, layout.ImageBaseY, iv))
	}

	if cfg.Anomaly.Enabled() {
		mgr.Register(NewAnomalyMonitor(cfg.Anomaly, layout.AnomalyBaseY, logger, anomalyDir, iv))
	}

	if len(cfg.SSLCerts) > 0 {
		mgr.Register(NewSSLMonitor(cfg.SSLCerts, layout.SSLBaseX, layout.SSLBaseY, iv))
	}

	for _, name := range mgr.MonitorNames() {
		logger.Info("Registered: " + name)
	}

	return mgr, layout
}
