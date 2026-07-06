package common

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// NFTImagePublicBase is the absolute URL prefix for NFT files served at GET /images/... (no trailing slash).
// When NFT_ASSETS_PUBLIC_BASE is unset, it is derived from the request (same scheme/host as this API).
// Set NFT_ASSETS_PUBLIC_BASE when the public asset URL must differ (CDN, external hostname, etc.).
func NFTImagePublicBase(c *gin.Context) string {
	if b := strings.TrimSpace(os.Getenv("NFT_ASSETS_PUBLIC_BASE")); b != "" {
		return NormalizeNFTAssetsPublicBase(b)
	}
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	if xf := strings.TrimSpace(c.Request.Header.Get("X-Forwarded-Proto")); xf != "" {
		scheme = strings.ToLower(xf)
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
	if strings.HasSuffix(b, "/randomwalk") {
		return strings.TrimSuffix(b, "/randomwalk") + "/images"
	}
	return b + "/images"
}

// NFTAssetsFlatPaths reports whether RandomWalk assets use /images/<file> (flat) instead of /images/randomwalk/<file>.
// Set NFT_ASSETS_FLAT_PATHS=1 on hosts like nfts.randomwalknft.com where files map directly under /images/.
func NFTAssetsFlatPaths() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("NFT_ASSETS_FLAT_PATHS")))
	return v == "1" || v == "true" || v == "yes"
}

// MetadataRandomWalkImagePublicBase is the absolute URL prefix (…/images, no trailing slash) for Random Walk
// ERC-721 metadata JSON fields `image` and `animation_url`.
//
// When NFT_ASSETS_PUBLIC_BASE is set, it is used (same as NFTImagePublicBase). When unset, this returns
// the public API image base: RandomWalk artwork is not served from the Vercel marketing site; it is
// served by websrv under GET /images/... (typically TLS on api.randomwalknft.com). Override with
// NFT_ASSETS_PUBLIC_BASE for other hosts (staging, IP-only deploys, etc.).
func MetadataRandomWalkImagePublicBase() string {
	if b := strings.TrimSpace(os.Getenv("NFT_ASSETS_PUBLIC_BASE")); b != "" {
		return NormalizeNFTAssetsPublicBase(b)
	}
	return "https://api.randomwalknft.com:1443/images"
}
