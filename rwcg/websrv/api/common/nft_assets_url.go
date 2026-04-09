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
		return strings.TrimRight(b, "/")
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

// MetadataRandomWalkImagePublicBase is the absolute URL prefix (…/images, no trailing slash) for Random Walk
// ERC-721 metadata JSON fields `image` and `animation_url`.
//
// When NFT_ASSETS_PUBLIC_BASE is set, it is used (same as NFTImagePublicBase). When unset, this returns
// https://randomwalknft.com/images so metadata served from dev/staging/local hosts does not embed
// localhost in URLs that wallets and marketplaces cache on-chain.
func MetadataRandomWalkImagePublicBase() string {
	if b := strings.TrimSpace(os.Getenv("NFT_ASSETS_PUBLIC_BASE")); b != "" {
		return strings.TrimRight(b, "/")
	}
	return "https://randomwalknft.com/images"
}
