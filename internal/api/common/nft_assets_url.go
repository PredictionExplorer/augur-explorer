package common

import (
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// NFTImagePublicBase is the absolute URL prefix for NFT files served at GET /images/... (no trailing slash).
// configuredBase is the NFT_ASSETS_PUBLIC_BASE configuration value; when empty, the base is derived from
// the request (same scheme/host as this API). Configure it when the public asset URL must differ
// (CDN, external hostname, etc.).
func NFTImagePublicBase(c *httpx.Context, configuredBase string) string {
	if b := strings.TrimSpace(configuredBase); b != "" {
		return NormalizeNFTAssetsPublicBase(b)
	}
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	if xf := strings.TrimSpace(c.Request.Header.Get("X-Forwarded-Proto")); xf != "" {
		if i := strings.IndexByte(xf, ','); i >= 0 {
			xf = strings.TrimSpace(xf[:i])
		}
		if xf != "" {
			scheme = strings.ToLower(xf)
		}
	}
	host := strings.TrimSpace(c.Request.Host)
	if xfh := strings.TrimSpace(c.Request.Header.Get("X-Forwarded-Host")); xfh != "" {
		if i := strings.IndexByte(xfh, ','); i >= 0 {
			xfh = strings.TrimSpace(xfh[:i])
		}
		host = xfh
	}
	if host == "" {
		return ""
	}
	return scheme + "://" + host + "/images"
}

// NormalizeNFTAssetsPublicBase ensures the public asset prefix ends with "/images" (no trailing slash).
// Accepts either "https://host" or "https://host/images"; fixes "…/randomwalk" misconfiguration.
func NormalizeNFTAssetsPublicBase(b string) string {
	b = strings.TrimRight(strings.TrimSpace(b), "/")
	if b == "" {
		return ""
	}
	if strings.HasSuffix(b, "/images") {
		return b
	}
	if before, ok := strings.CutSuffix(b, "/randomwalk"); ok {
		return before + "/images"
	}
	return b + "/images"
}

// MetadataRandomWalkImagePublicBase is the absolute URL prefix (…/images, no trailing slash) for Random Walk
// ERC-721 metadata JSON fields `image` and `animation_url`.
//
// configuredBase is the NFT_ASSETS_PUBLIC_BASE configuration value. When set, it is used (same as
// NFTImagePublicBase). When empty, this returns the public API image base: RandomWalk artwork is not served
// from the Vercel marketing site; it is served by websrv under GET /images/... (typically TLS on
// api.randomwalknft.com). Configure the base for other hosts (staging, IP-only deploys, etc.).
func MetadataRandomWalkImagePublicBase(configuredBase string) string {
	if b := strings.TrimSpace(configuredBase); b != "" {
		return NormalizeNFTAssetsPublicBase(b)
	}
	return "https://api.randomwalknft.com:1443/images"
}
