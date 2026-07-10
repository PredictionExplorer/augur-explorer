package types

import (
	"time"
	"github.com/nsf/termbox-go"
)

// Position represents a screen position
type Position struct {
	X int
	Y int
}

// Color represents a display color
type Color termbox.Attribute

// Common colors
const (
	ColorDefault Color = Color(termbox.ColorDefault)
	ColorWhite   Color = Color(termbox.ColorWhite)
	ColorRed     Color = Color(termbox.ColorRed)
	ColorGreen   Color = Color(termbox.ColorGreen)
	ColorYellow  Color = Color(termbox.ColorYellow)
	ColorBlue    Color = Color(termbox.ColorBlue)
	ColorCyan    Color = Color(termbox.ColorCyan)
)

// TextSegment represents a piece of colored text
type TextSegment struct {
	Text       string
	Foreground Color
	Background Color
}

// DisplayLine represents a line to display
type DisplayLine struct {
	Position Position
	Segments []TextSegment
}

// MonitorStatus represents the status of a monitor
type MonitorStatus string

const (
	StatusOK      MonitorStatus = "Ok"
	StatusFail    MonitorStatus = "Fail"
	StatusUnknown MonitorStatus = "Unknown"
)

// MonitorResult contains the result of a monitor check
type MonitorResult struct {
	Status    MonitorStatus
	Message   string
	Error     error
	Timestamp time.Time
}

// DatabaseConfig holds database connection information
type DatabaseConfig struct {
	Name   string
	Host   string
	DBName string
	User   string
	Pass   string
}

// RPCConfig holds RPC node information
type RPCConfig struct {
	Name       string
	URL        string
	ChainID    string
	IsOfficial bool
}

// WebAPIConfig holds web API information
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

// DiskConfig holds disk monitoring configuration
type DiskConfig struct {
	Title      string
	User       string
	IP         string
	DeviceList string
}

// ImageServerConfig holds image server monitoring configuration
type ImageServerConfig struct {
	Name         string
	URL          string
	ContractAddr string
	RPCURL       string
}

// SSLCertConfig holds SSL certificate monitoring configuration
type SSLCertConfig struct {
	Name       string // optional display label; defaults to Host or Host:Port
	Host       string // address to connect to (hostname or IP)
	Port       string // optional; defaults to 443
	ServerName string // optional TLS SNI; defaults to Host (set when Host is an IP)
}

// AnomalyConfig holds configuration for fetching the websrv anomalies file
// (produced by the loganomaly tool on the production host) via scp.
type AnomalyConfig struct {
	Title       string // optional display label
	User        string // ssh user, e.g. "cgprod"
	Host        string // ssh host/alias, e.g. "cosmic1"
	RemoteFile  string // path to the anomalies file on the remote host
	StaleSecond int    // flag the feed STALE if not regenerated within this many seconds (0 = default)
}

// EventTableConfig holds event table monitoring configuration
type EventTableConfig struct {
	Name       string
	Host       string
	DBName     string
	User       string
	Pass       string
	TableName  string // e.g., "cg_proc_status" or "rw_proc_status"
	ColumnName string // e.g., "last_evt_id" or "last_block"
}






