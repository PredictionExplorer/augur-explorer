package common

import "strings"

// MetadataHostServesCosmicSignature decides which project's ERC-721 metadata a
// bare /metadata/:token_id request is for. Both projects' on-chain baseURI is
// https://<host>/metadata/ and this single server serves both hosts, so the
// dispatch key is the effective request host: the Host header, overridden by
// the first entry of X-Forwarded-Host when a proxy supplies one. A host that
// mentions "cosmicsignature" (any case) gets Cosmic Signature metadata;
// anything else (the RandomWalk hosts) gets RandomWalk metadata.
func MetadataHostServesCosmicSignature(host, xForwardedHost string) bool {
	effective := strings.ToLower(host)
	if xForwardedHost != "" {
		if i := strings.IndexByte(xForwardedHost, ','); i >= 0 {
			xForwardedHost = xForwardedHost[:i]
		}
		effective = strings.ToLower(strings.TrimSpace(xForwardedHost))
	}
	return strings.Contains(effective, "cosmicsignature")
}
