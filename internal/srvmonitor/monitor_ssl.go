package srvmonitor

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	// SSLWarningDays is the expiry proximity that turns a certificate red
	// and raises an alarm.
	SSLWarningDays     = 10
	defaultSSLCertPort = "443"
	sslDialTimeout     = 10 * time.Second
)

// SSLCertStatus holds expiry status for a single certificate.
type SSLCertStatus struct {
	Config SSLCertConfig
	Days   int
	ErrStr string
}

// SSLMonitor monitors SSL certificate expiry via remote TLS connections.
type SSLMonitor struct {
	certs    []SSLCertConfig
	statuses []SSLCertStatus
	position Position
	interval time.Duration
	// rootCAs overrides the trust store; nil selects the system roots.
	// Tests point it at their own CA.
	rootCAs *x509.CertPool
}

// NewSSLMonitor creates a new SSL certificate monitor. baseX/baseY position
// the top-left of the section, allowing it to be placed in a right-hand column.
func NewSSLMonitor(certs []SSLCertConfig, baseX, baseY int, iv Intervals) *SSLMonitor {
	return &SSLMonitor{
		certs:    certs,
		statuses: make([]SSLCertStatus, len(certs)),
		position: Position{X: baseX, Y: baseY},
		interval: iv.SSL,
	}
}

// Name returns the monitor name.
func (m *SSLMonitor) Name() string {
	return "SSL Monitor"
}

// Start begins monitoring.
func (m *SSLMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

func (m *SSLMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.certs))

	for i := range m.certs {
		go func(idx int) {
			defer wg.Done()
			cfg := m.certs[idx]
			label := sslCertLabel(cfg)
			m.statuses[idx].Config = cfg
			days, err := m.daysUntilCertExpiry(ctx, cfg)
			if err != nil {
				m.statuses[idx].Days = -1
				m.statuses[idx].ErrStr = err.Error()
				sendErr(ctx, errorChan, fmt.Sprintf("SSL cert %s: %v", label, err))
				return
			}
			m.statuses[idx].Days = days
			m.statuses[idx].ErrStr = ""
			if days <= SSLWarningDays {
				sendErr(ctx, errorChan, fmt.Sprintf("SSL cert %s expires in %d days", label, days))
			}
		}(i)
	}

	wg.Wait()
	m.display(disp)
}

func sslCertPort(cfg SSLCertConfig) string {
	if cfg.Port != "" {
		return cfg.Port
	}
	return defaultSSLCertPort
}

func sslCertServerName(cfg SSLCertConfig) string {
	if cfg.ServerName != "" {
		return cfg.ServerName
	}
	return cfg.Host
}

func sslCertLabel(cfg SSLCertConfig) string {
	if cfg.Name != "" {
		return cfg.Name
	}
	port := sslCertPort(cfg)
	if port == defaultSSLCertPort {
		return cfg.Host
	}
	return net.JoinHostPort(cfg.Host, port)
}

func (m *SSLMonitor) daysUntilCertExpiry(ctx context.Context, cfg SSLCertConfig) (int, error) {
	addr := net.JoinHostPort(cfg.Host, sslCertPort(cfg))
	dialer := &tls.Dialer{
		NetDialer: &net.Dialer{Timeout: sslDialTimeout},
		Config: &tls.Config{
			ServerName: sslCertServerName(cfg),
			MinVersion: tls.VersionTLS12,
			RootCAs:    m.rootCAs,
		},
	}
	netConn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return 0, err
	}
	conn := netConn.(*tls.Conn)
	defer func() { _ = conn.Close() }() // best-effort close; connection is read-only

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return 0, fmt.Errorf("no certificate presented by %s", sslCertLabel(cfg))
	}

	days := int(time.Until(certs[0].NotAfter).Hours() / 24)
	return days, nil
}

func (m *SSLMonitor) display(disp Display) {
	x := m.position.X
	y := m.position.Y

	disp.DrawText(Position{X: x, Y: y},
		"--- SSL Certificates ---",
		ColorWhite, ColorDefault)

	for i, status := range m.statuses {
		lineY := y + 1 + i
		label := fmt.Sprintf(" %s: ", sslCertLabel(status.Config))

		disp.DrawText(Position{X: x, Y: lineY}, label, ColorWhite, ColorDefault)

		var valueStr string
		var valueColor Color
		switch {
		case status.ErrStr != "":
			valueStr = "Error"
			valueColor = ColorRed
		case status.Days < 0:
			valueStr = "Expired"
			valueColor = ColorRed
		case status.Days == 0:
			valueStr = "Expires today"
			valueColor = ColorRed
		case status.Days == 1:
			valueStr = "1 day left"
			valueColor = colorForDaysLeft(status.Days)
		default:
			valueStr = fmt.Sprintf("%d days left", status.Days)
			valueColor = colorForDaysLeft(status.Days)
		}

		disp.DrawText(Position{X: x + len(label), Y: lineY}, valueStr, valueColor, ColorDefault)
	}

	disp.Flush()
}

func colorForDaysLeft(days int) Color {
	if days <= SSLWarningDays {
		return ColorRed
	}
	return ColorGreen
}
