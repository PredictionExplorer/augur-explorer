package store

import (
	"context"
	"net"
	"time"
)

// Connection resilience tuning. These mirror libpq's connect_timeout +
// keepalives_idle / keepalives_interval / keepalives_count / tcp_user_timeout.
// pgx accepts connect_timeout as a connection-string key, but the keepalive
// tuning knobs are applied at the socket level via a custom dialer
// (see keepaliveDialer below).
const (
	dbConnectTimeout    = 10 * time.Second
	dbKeepaliveIdle     = 30 * time.Second
	dbKeepaliveInterval = 10 * time.Second
	dbKeepaliveCount    = 5
	dbTCPUserTimeout    = 15 * time.Second

	// Connect retry: a transient link failure (e.g. "no route to host" on a
	// flaky Wi-Fi link) at process startup should not immediately terminate
	// the service. New retries the initial connection several times with
	// capped backoff so a short blip is absorbed.
	dbConnectMaxAttempts  = 10
	dbConnectRetryDelay   = 1 * time.Second
	dbConnectRetryMaxWait = 5 * time.Second
)

// keepaliveDialer dials with a bounded connect timeout and TCP keepalive probing
// (plus TCP_USER_TIMEOUT on Linux) so a dead or flaky link is detected in
// seconds instead of the kernel default of minutes/hours.
type keepaliveDialer struct{ d net.Dialer }

func newKeepaliveDialer() keepaliveDialer {
	return keepaliveDialer{d: net.Dialer{
		Timeout:   dbConnectTimeout,
		KeepAlive: dialerKeepAlive,     // platform-specific (see db_keepalive_*.go)
		Control:   tcpKeepaliveControl, // platform-specific socket-option tuning
	}}
}

// DialContext satisfies pgconn.DialFunc; pgx uses it for both TCP and Unix
// socket connections (the keepalive Control is a no-op for Unix sockets).
func (k keepaliveDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return k.d.DialContext(ctx, network, address)
}
