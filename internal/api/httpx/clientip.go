package httpx

import (
	"net"
	"net/http"
	"net/netip"
	"strings"
	"sync/atomic"
)

// trustedProxies is the set of reverse-proxy networks whose forwarding
// headers ClientIP trusts. Empty (the default) means no proxy is trusted
// and the TCP peer address is always the client, matching the legacy
// no-proxy deployments.
var trustedProxies atomic.Pointer[[]netip.Prefix]

// SetTrustedProxies installs the trusted reverse-proxy set (TRUSTED_PROXIES).
// Call it once at startup, before serving; nil or empty disables proxy-header
// resolution.
func SetTrustedProxies(prefixes []netip.Prefix) {
	trustedProxies.Store(&prefixes)
}

func isTrustedProxy(addr netip.Addr) bool {
	prefixes := trustedProxies.Load()
	if prefixes == nil {
		return false
	}
	for _, p := range *prefixes {
		if p.Contains(addr) {
			return true
		}
	}
	return false
}

// ClientIP returns the IP of the client that issued r, or "" when it cannot
// be determined.
//
// When the TCP peer is one of the configured trusted proxies, the client is
// resolved from the forwarding headers that proxies like nginx set: the
// X-Forwarded-For chain is walked right to left and the first address that
// is not itself a trusted proxy wins (entries further left are supplied by
// the client and spoofable); X-Real-IP is the fallback. In every other case
// — no trusted proxies configured, or a direct connection — the TCP peer
// address is returned, matching the legacy behavior.
func ClientIP(r *http.Request) string {
	peer := parseAddr(r.RemoteAddr)
	if !peer.IsValid() {
		return ""
	}
	if !isTrustedProxy(peer) {
		return peer.String()
	}
	if ip := forwardedClient(r.Header.Values("X-Forwarded-For")); ip.IsValid() {
		return ip.String()
	}
	if ip := parseAddr(r.Header.Get("X-Real-Ip")); ip.IsValid() {
		return ip.String()
	}
	return peer.String()
}

// forwardedClient resolves the client from an X-Forwarded-For chain: the
// rightmost address that is not a trusted proxy. A malformed entry aborts
// the walk (nothing further left can be trusted).
func forwardedClient(headerValues []string) netip.Addr {
	chain := make([]string, 0, len(headerValues))
	for _, v := range headerValues {
		chain = append(chain, strings.Split(v, ",")...)
	}
	for i := len(chain) - 1; i >= 0; i-- {
		addr := parseAddr(chain[i])
		if !addr.IsValid() {
			return netip.Addr{}
		}
		if !isTrustedProxy(addr) {
			return addr
		}
	}
	return netip.Addr{}
}

// parseAddr parses an IP with an optional port ("1.2.3.4", "1.2.3.4:56",
// "[::1]:56"), unmapping 4-in-6 addresses so 127.0.0.1 and ::ffff:127.0.0.1
// compare equal against the trusted set.
func parseAddr(s string) netip.Addr {
	s = strings.TrimSpace(s)
	if host, _, err := net.SplitHostPort(s); err == nil {
		s = host
	}
	addr, err := netip.ParseAddr(s)
	if err != nil {
		return netip.Addr{}
	}
	return addr.Unmap()
}
