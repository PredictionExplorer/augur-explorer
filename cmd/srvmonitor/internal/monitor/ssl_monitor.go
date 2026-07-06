package monitor

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/display"
	"github.com/PredictionExplorer/augur-explorer/cmd/srvmonitor/internal/types"
)

const (
	SSLCheckInterval   = 3600 // 1 hour
	SSLWarningDays     = 10
	defaultSSLCertPort = "443"
)

// SSLCertStatus holds expiry status for a single certificate
type SSLCertStatus struct {
	Config types.SSLCertConfig
	Days   int
	ErrStr string
}

// SSLMonitor monitors SSL certificate expiry via remote TLS connections
type SSLMonitor struct {
	certs    []types.SSLCertConfig
	statuses []SSLCertStatus
	position types.Position
}

// NewSSLMonitor creates a new SSL certificate monitor. baseX/baseY position
// the top-left of the section, allowing it to be placed in a right-hand column.
func NewSSLMonitor(certs []types.SSLCertConfig, baseX, baseY int) *SSLMonitor {
	return &SSLMonitor{
		certs:    certs,
		statuses: make([]SSLCertStatus, len(certs)),
		position: types.Position{X: baseX, Y: baseY},
	}
}

// Name returns the monitor name
func (m *SSLMonitor) Name() string {
	return "SSL Monitor"
}

// GetDisplayPosition returns the display position
func (m *SSLMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *SSLMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.check(disp, errorChan)
			time.Sleep(SSLCheckInterval * time.Second)
		}
	}
}

func (m *SSLMonitor) check(disp display.Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.certs))

	for i := range m.certs {
		go func(idx int) {
			defer wg.Done()
			cfg := m.certs[idx]
			label := sslCertLabel(cfg)
			m.statuses[idx].Config = cfg
			days, err := daysUntilCertExpiry(cfg)
			if err != nil {
				m.statuses[idx].Days = -1
				m.statuses[idx].ErrStr = err.Error()
				errorChan <- fmt.Sprintf("SSL cert %s: %v", label, err)
				return
			}
			m.statuses[idx].Days = days
			m.statuses[idx].ErrStr = ""
			if days <= SSLWarningDays {
				errorChan <- fmt.Sprintf("SSL cert %s expires in %d days", label, days)
			}
		}(i)
	}

	wg.Wait()
	m.display(disp)
}

func sslCertPort(cfg types.SSLCertConfig) string {
	if cfg.Port != "" {
		return cfg.Port
	}
	return defaultSSLCertPort
}

func sslCertServerName(cfg types.SSLCertConfig) string {
	if cfg.ServerName != "" {
		return cfg.ServerName
	}
	return cfg.Host
}

func sslCertLabel(cfg types.SSLCertConfig) string {
	if cfg.Name != "" {
		return cfg.Name
	}
	port := sslCertPort(cfg)
	if port == defaultSSLCertPort {
		return cfg.Host
	}
	return net.JoinHostPort(cfg.Host, port)
}

func daysUntilCertExpiry(cfg types.SSLCertConfig) (int, error) {
	addr := net.JoinHostPort(cfg.Host, sslCertPort(cfg))
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: 10 * time.Second},
		"tcp",
		addr,
		&tls.Config{
			ServerName: sslCertServerName(cfg),
			MinVersion: tls.VersionTLS12,
		},
	)
	if err != nil {
		return 0, err
	}
	defer func() { _ = conn.Close() }() // best-effort close; connection is read-only

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return 0, fmt.Errorf("no certificate presented by %s", sslCertLabel(cfg))
	}

	days := int(time.Until(certs[0].NotAfter).Hours() / 24)
	return days, nil
}

func (m *SSLMonitor) display(disp display.Display) {
	x := m.position.X
	y := m.position.Y

	disp.DrawText(types.Position{X: x, Y: y},
		"--- SSL Certificates ---",
		types.ColorWhite, types.ColorDefault)

	for i, status := range m.statuses {
		lineY := y + 1 + i
		label := fmt.Sprintf(" %s: ", sslCertLabel(status.Config))

		disp.DrawText(types.Position{X: x, Y: lineY}, label, types.ColorWhite, types.ColorDefault)

		var valueStr string
		var valueColor types.Color
		if status.ErrStr != "" {
			valueStr = "Error"
			valueColor = types.ColorRed
		} else if status.Days < 0 {
			valueStr = "Expired"
			valueColor = types.ColorRed
		} else if status.Days == 0 {
			valueStr = "Expires today"
			valueColor = types.ColorRed
		} else if status.Days == 1 {
			valueStr = "1 day left"
			valueColor = colorForDaysLeft(status.Days)
		} else {
			valueStr = fmt.Sprintf("%d days left", status.Days)
			valueColor = colorForDaysLeft(status.Days)
		}

		disp.DrawText(types.Position{X: x + len(label), Y: lineY}, valueStr, valueColor, types.ColorDefault)
	}

	disp.Flush()
}

func colorForDaysLeft(days int) types.Color {
	if days <= SSLWarningDays {
		return types.ColorRed
	}
	return types.ColorGreen
}
