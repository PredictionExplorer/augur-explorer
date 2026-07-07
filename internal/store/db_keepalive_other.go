//go:build !linux

package store

import "syscall"

// Non-Linux platforms: rely on net.Dialer's built-in keepalive (SO_KEEPALIVE
// with this idle/interval). TCP_KEEPIDLE / TCP_KEEPCNT / TCP_USER_TIMEOUT are
// not portable, so they're skipped here.
var dialerKeepAlive = dbKeepaliveIdle

func tcpKeepaliveControl(_, _ string, _ syscall.RawConn) error {
	return nil
}
