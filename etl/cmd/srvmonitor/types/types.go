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


