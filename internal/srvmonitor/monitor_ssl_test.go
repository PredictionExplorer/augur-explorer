package srvmonitor

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net"
	"strings"
	"testing"
	"time"
)

// startTLSServer serves TLS on a random local port with a self-signed
// certificate expiring at notAfter. It returns host, port and a pool
// trusting the certificate.
func startTLSServer(t *testing.T, serverName string, notAfter time.Time) (string, string, *x509.CertPool) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: serverName},
		DNSNames:              []string{serverName},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	der, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	cert := tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}

	parsed, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AddCert(parsed)

	listener, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = listener.Close() })

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				// Drive the handshake so the client sees the certificate,
				// then close.
				if tc, ok := c.(*tls.Conn); ok {
					_ = tc.Handshake()
				}
				_ = c.Close()
			}(conn)
		}
	}()

	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	return host, port, pool
}

func TestSSLMonitorDaysLeft(t *testing.T) {
	t.Parallel()
	host, port, pool := startTLSServer(t, "example.test", time.Now().Add(90*24*time.Hour))

	certs := []SSLCertConfig{{Name: "prod cert", Host: host, Port: port, ServerName: "example.test"}}
	m := NewSSLMonitor(certs, 62, 24, testIntervals())
	m.rootCAs = pool

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	st := m.statuses[0]
	if st.ErrStr != "" || st.Days < 88 || st.Days > 90 {
		t.Fatalf("status = %+v, want ~89 days", st)
	}
	row := disp.Row(25)
	if !strings.Contains(row, "prod cert:") || !strings.Contains(row, "days left") {
		t.Fatalf("row = %q", row)
	}
	if header := disp.Row(24); !strings.Contains(header, "SSL Certificates") {
		t.Fatalf("header = %q", header)
	}
	// A comfortably distant expiry renders green.
	labelWidth := len(" prod cert: ")
	if got := disp.FgAt(62+labelWidth, 25); got != ColorGreen {
		t.Fatalf("value color = %v, want green", got)
	}
}

func TestSSLMonitorExpiringSoonRaisesAlarm(t *testing.T) {
	t.Parallel()
	host, port, pool := startTLSServer(t, "soon.test", time.Now().Add(5*24*time.Hour))

	certs := []SSLCertConfig{{Host: host, Port: port, ServerName: "soon.test"}}
	m := NewSSLMonitor(certs, 0, 24, testIntervals())
	m.rootCAs = pool

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "expires in") {
		t.Fatalf("errors = %v, want expiry warning", msgs)
	}
	if row := disp.Row(25); !strings.Contains(row, "days left") {
		t.Fatalf("row = %q", row)
	}
}

func TestSSLMonitorUnreachableAndBadName(t *testing.T) {
	t.Parallel()

	t.Run("unreachable", func(t *testing.T) {
		t.Parallel()
		m := NewSSLMonitor([]SSLCertConfig{{Host: "127.0.0.1", Port: "1"}}, 0, 24, testIntervals())
		errCh := make(chan string, 10)
		disp := newFakeDisplay()
		m.check(context.Background(), disp, errCh)

		st := m.statuses[0]
		if st.ErrStr == "" || st.Days != -1 {
			t.Fatalf("status = %+v, want dial error", st)
		}
		msgs := drain(errCh)
		if len(msgs) != 1 || !strings.Contains(msgs[0], "SSL cert 127.0.0.1:1:") {
			t.Fatalf("errors = %v", msgs)
		}
		if row := disp.Row(25); !strings.Contains(row, "Error") {
			t.Fatalf("row = %q", row)
		}
	})

	t.Run("wrong server name", func(t *testing.T) {
		t.Parallel()
		host, port, pool := startTLSServer(t, "right.test", time.Now().Add(30*24*time.Hour))
		// The client requires a matching SNI/SAN; "wrong.test" fails
		// verification even against a trusted root.
		m := NewSSLMonitor([]SSLCertConfig{{Host: host, Port: port, ServerName: "wrong.test"}}, 0, 24, testIntervals())
		m.rootCAs = pool
		errCh := make(chan string, 10)
		m.check(context.Background(), newFakeDisplay(), errCh)

		if m.statuses[0].ErrStr == "" {
			t.Fatal("wrong server name must fail verification")
		}
	})

	t.Run("untrusted issuer", func(t *testing.T) {
		t.Parallel()
		host, port, _ := startTLSServer(t, "self.test", time.Now().Add(30*24*time.Hour))
		// Without the pool the self-signed certificate fails chain
		// verification against the system roots.
		m := NewSSLMonitor([]SSLCertConfig{{Host: host, Port: port, ServerName: "self.test"}}, 0, 24, testIntervals())
		errCh := make(chan string, 10)
		m.check(context.Background(), newFakeDisplay(), errCh)

		if m.statuses[0].ErrStr == "" {
			t.Fatal("untrusted issuer must fail verification")
		}
	})
}

func TestSSLCertHelpers(t *testing.T) {
	t.Parallel()
	if got := sslCertPort(SSLCertConfig{}); got != "443" {
		t.Fatalf("default port = %q", got)
	}
	if got := sslCertPort(SSLCertConfig{Port: "8443"}); got != "8443" {
		t.Fatalf("port = %q", got)
	}
	if got := sslCertServerName(SSLCertConfig{Host: "h"}); got != "h" {
		t.Fatalf("server name = %q", got)
	}
	if got := sslCertServerName(SSLCertConfig{Host: "h", ServerName: "sni"}); got != "sni" {
		t.Fatalf("server name = %q", got)
	}
	if got := sslCertLabel(SSLCertConfig{Name: "n", Host: "h"}); got != "n" {
		t.Fatalf("label = %q", got)
	}
	if got := sslCertLabel(SSLCertConfig{Host: "h"}); got != "h" {
		t.Fatalf("label = %q", got)
	}
	if got := sslCertLabel(SSLCertConfig{Host: "h", Port: "8443"}); got != "h:8443" {
		t.Fatalf("label = %q", got)
	}
}

func TestColorForDaysLeft(t *testing.T) {
	t.Parallel()
	if got := colorForDaysLeft(10); got != ColorRed {
		t.Fatalf("10 days = %v, want red", got)
	}
	if got := colorForDaysLeft(11); got != ColorGreen {
		t.Fatalf("11 days = %v, want green", got)
	}
}

func TestSSLMonitorDisplayEdgeStates(t *testing.T) {
	t.Parallel()
	m := NewSSLMonitor([]SSLCertConfig{
		{Name: "expired", Host: "a"},
		{Name: "today", Host: "b"},
		{Name: "oneday", Host: "c"},
	}, 0, 24, testIntervals())
	m.statuses[0] = SSLCertStatus{Config: m.certs[0], Days: -1}
	m.statuses[1] = SSLCertStatus{Config: m.certs[1], Days: 0}
	m.statuses[2] = SSLCertStatus{Config: m.certs[2], Days: 1}

	disp := newFakeDisplay()
	m.display(disp)

	if row := disp.Row(25); !strings.Contains(row, "Expired") {
		t.Fatalf("row = %q", row)
	}
	if row := disp.Row(26); !strings.Contains(row, "Expires today") {
		t.Fatalf("row = %q", row)
	}
	if row := disp.Row(27); !strings.Contains(row, "1 day left") {
		t.Fatalf("row = %q", row)
	}
}

func TestSSLMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	host, port, pool := startTLSServer(t, "loop.test", time.Now().Add(30*24*time.Hour))
	m := NewSSLMonitor([]SSLCertConfig{{Host: host, Port: port, ServerName: "loop.test"}}, 0, 24, testIntervals())
	m.rootCAs = pool

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	disp := newFakeDisplay()
	go func() {
		m.Start(ctx, disp, make(chan string, 100))
		close(done)
	}()

	waitFor(t, "two check cycles", func() bool { return disp.Flushes() >= 2 })
	cancel()
	waitFor(t, "loop exit", func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	})
}
