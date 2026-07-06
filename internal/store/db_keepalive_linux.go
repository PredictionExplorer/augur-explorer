//go:build linux

package dbs

import (
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

// On Linux we manage keepalive entirely through the socket options below, so
// disable net.Dialer's own (coarser) keepalive handling to keep it from
// overwriting our TCP_KEEPIDLE / TCP_KEEPINTVL values.
var dialerKeepAlive = time.Duration(-1)

// tcpKeepaliveControl enables SO_KEEPALIVE and tunes the per-connection probe
// timing plus TCP_USER_TIMEOUT, so a dead or flaky link is dropped in seconds
// rather than the kernel defaults (~2h idle, ~15min of unacked data). This is
// the socket-level equivalent of libpq's keepalives_idle / keepalives_interval /
// keepalives_count / tcp_user_timeout.
func tcpKeepaliveControl(network, _ string, c syscall.RawConn) error {
	switch network {
	case "tcp", "tcp4", "tcp6":
	default:
		return nil // unix socket: nothing to tune
	}
	var sockErr error
	ctrlErr := c.Control(func(fd uintptr) {
		f := int(fd)
		set := func(level, opt, val int) {
			if sockErr == nil {
				sockErr = unix.SetsockoptInt(f, level, opt, val)
			}
		}
		set(unix.SOL_SOCKET, unix.SO_KEEPALIVE, 1)
		set(unix.IPPROTO_TCP, unix.TCP_KEEPIDLE, int(dbKeepaliveIdle.Seconds()))
		set(unix.IPPROTO_TCP, unix.TCP_KEEPINTVL, int(dbKeepaliveInterval.Seconds()))
		set(unix.IPPROTO_TCP, unix.TCP_KEEPCNT, dbKeepaliveCount)
		set(unix.IPPROTO_TCP, unix.TCP_USER_TIMEOUT, int(dbTCPUserTimeout.Milliseconds()))
	})
	if ctrlErr != nil {
		return ctrlErr
	}
	return sockErr
}
