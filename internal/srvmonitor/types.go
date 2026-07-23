// Package srvmonitor implements the terminal server-monitoring engine behind
// cmd/srvmonitor: periodic checks of RPC nodes, PostgreSQL databases, web
// APIs, disks, SSL certificates, image servers and error-log anomalies, each
// rendering its section of a shared character display and reporting failures
// to a central manager that tracks alarms and sends mobile notifications.
//
// Every external system is behind an injectable seam (database connector,
// command runner, HTTP client, poll intervals), so the whole engine is
// testable without real servers; cmd/srvmonitor wires the production
// implementations plus the termbox UI (internal/srvmonitor/termboxui).
package srvmonitor

import "time"

// Position represents a screen position.
type Position struct {
	X int
	Y int
}

// Color identifies a terminal color. The values mirror termbox-go's
// Attribute constants so the termbox display can cast directly.
type Color uint16

// Terminal colors used by the monitors.
const (
	ColorDefault Color = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

// DatabaseConfig holds database connection information.
type DatabaseConfig struct {
	Name   string
	Host   string
	DBName string
	User   string
	Pass   string
}

// RPCConfig holds RPC node information.
type RPCConfig struct {
	Name       string
	URL        string
	ChainID    string
	IsOfficial bool
}

// WebAPIConfig holds web API information.
type WebAPIConfig struct {
	Title string
	Host  string
	Port  string
	URI   string
	// PublicURL is an optional full URL (e.g. https://a1.cosmicsignature.com/api/...)
	// for the same service, reached through the public path (TLS/stunnel/nginx).
	// When set, the monitor probes it in addition to the internal Host:Port check.
	PublicURL string
}

// DiskConfig holds disk monitoring configuration.
type DiskConfig struct {
	Title      string
	User       string
	IP         string
	DeviceList string
}

// ImageServerConfig holds image server monitoring configuration.
type ImageServerConfig struct {
	Name         string
	URL          string
	ContractAddr string
	RPCURL       string
}

// SSLCertConfig holds SSL certificate monitoring configuration.
type SSLCertConfig struct {
	Name       string // optional display label; defaults to Host or Host:Port
	Host       string // address to connect to (hostname or IP)
	Port       string // optional; defaults to 443
	ServerName string // optional TLS SNI; defaults to Host (set when Host is an IP)
}

// AnomalyConfig holds configuration for fetching the websrv anomalies file
// (produced by the loganomaly tool on the production host) via scp.
type AnomalyConfig struct {
	Title      string        // optional display label
	User       string        // ssh user, e.g. "cgprod"
	Host       string        // ssh host/alias, e.g. "cosmic1"
	RemoteFile string        // path to the anomalies file on the remote host
	StaleAfter time.Duration // maximum heartbeat age; zero selects DefaultAnomalyStaleAfter
}

// DefaultAnomalyStaleAfter is the maximum heartbeat age used when the
// operator does not configure ANOMALY_STALE_SECS.
const DefaultAnomalyStaleAfter = 30 * time.Minute

// Enabled reports whether anomaly monitoring is configured.
func (c AnomalyConfig) Enabled() bool {
	return c.User != "" && c.Host != "" && c.RemoteFile != ""
}

// EventTableConfig holds event table monitoring configuration.
type EventTableConfig struct {
	DatabaseConfig

	TableName  string // e.g., "cg_proc_status" or "rw_proc_status"
	ColumnName string // e.g., "last_evt_id" or "last_block"
}
