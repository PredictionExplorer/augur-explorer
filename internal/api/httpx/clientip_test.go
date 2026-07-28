package httpx

import (
	"net/http"
	"net/http/httptest"
	"net/netip"
	"testing"
)

// setProxies installs a trusted-proxy set for one test and restores the
// empty default afterwards (the set is process-global).
func setProxies(t *testing.T, cidrs ...string) {
	t.Helper()
	prefixes := make([]netip.Prefix, 0, len(cidrs))
	for _, c := range cidrs {
		prefixes = append(prefixes, netip.MustParsePrefix(c))
	}
	SetTrustedProxies(prefixes)
	t.Cleanup(func() { SetTrustedProxies(nil) })
}

func TestClientIPDirectConnection(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		want       string
	}{
		{"ipv4 with port", "187.246.6.100:51234", "187.246.6.100"},
		{"ipv6 with port", "[2001:db8::1]:443", "2001:db8::1"},
		{"bare ip", "10.1.2.3", "10.1.2.3"},
		{"garbage", "not-an-address", ""},
		{"empty", "", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/x", nil)
			r.RemoteAddr = test.remoteAddr
			if got := ClientIP(r); got != test.want {
				t.Errorf("ClientIP = %q, want %q", got, test.want)
			}
		})
	}
}

func TestClientIPIgnoresHeadersWithoutTrustedProxies(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/x", nil)
	r.RemoteAddr = "203.0.113.7:1000"
	r.Header.Set("X-Forwarded-For", "1.2.3.4")
	r.Header.Set("X-Real-Ip", "5.6.7.8")
	if got := ClientIP(r); got != "203.0.113.7" {
		t.Errorf("ClientIP = %q, want the spoof-proof peer address", got)
	}
}

func TestClientIPIgnoresHeadersFromUntrustedPeer(t *testing.T) {
	setProxies(t, "10.0.0.0/8")
	r := httptest.NewRequest(http.MethodGet, "/x", nil)
	r.RemoteAddr = "203.0.113.7:1000" // not in 10.0.0.0/8
	r.Header.Set("X-Forwarded-For", "1.2.3.4")
	if got := ClientIP(r); got != "203.0.113.7" {
		t.Errorf("ClientIP = %q, want peer address", got)
	}
}

func TestClientIPBehindTrustedProxy(t *testing.T) {
	setProxies(t, "127.0.0.1/32", "69.10.55.2/32")

	tests := []struct {
		name string
		xff  []string
		xrip string
		want string
	}{
		{name: "single hop", xff: []string{"187.246.6.100"}, want: "187.246.6.100"},
		{name: "client spoof plus real", xff: []string{"1.2.3.4, 187.246.6.100"}, want: "187.246.6.100"},
		{name: "chained trusted proxies", xff: []string{"187.246.6.100, 69.10.55.2"}, want: "187.246.6.100"},
		{name: "two header lines", xff: []string{"187.246.6.100", "69.10.55.2"}, want: "187.246.6.100"},
		{name: "all hops trusted falls back to real-ip", xff: []string{"69.10.55.2"}, xrip: "187.246.6.100", want: "187.246.6.100"},
		{name: "malformed chain falls back to real-ip", xff: []string{"unknown"}, xrip: "187.246.6.100", want: "187.246.6.100"},
		{name: "malformed chain and no real-ip falls back to peer", xff: []string{"unknown"}, want: "127.0.0.1"},
		{name: "real-ip only", xrip: "187.246.6.100", want: "187.246.6.100"},
		{name: "no headers", want: "127.0.0.1"},
		{name: "mapped ipv4 client", xff: []string{"::ffff:187.246.6.100"}, want: "187.246.6.100"},
		{name: "entry with port", xff: []string{"187.246.6.100:4242"}, want: "187.246.6.100"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/x", nil)
			r.RemoteAddr = "127.0.0.1:33000"
			for _, v := range test.xff {
				r.Header.Add("X-Forwarded-For", v)
			}
			if test.xrip != "" {
				r.Header.Set("X-Real-Ip", test.xrip)
			}
			if got := ClientIP(r); got != test.want {
				t.Errorf("ClientIP = %q, want %q", got, test.want)
			}
		})
	}
}

func TestClientIPTrustedProxyMappedPeer(t *testing.T) {
	// Go listeners often report IPv4 peers as 4-in-6; the trusted-set
	// match must unmap before comparing.
	setProxies(t, "69.10.55.2/32")
	r := httptest.NewRequest(http.MethodGet, "/x", nil)
	r.RemoteAddr = "[::ffff:69.10.55.2]:1000"
	r.Header.Set("X-Forwarded-For", "187.246.6.100")
	if got := ClientIP(r); got != "187.246.6.100" {
		t.Errorf("ClientIP = %q, want forwarded client", got)
	}
}
