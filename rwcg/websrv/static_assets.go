package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// registerStaticAssetRoutes serves nft-assets mirror at GET /images/*.
// NFT_ASSETS_ROOT: absolute path to directory whose contents appear under /images/
// (e.g. .../nft-assets with new/, randomwalk/).
//
// STATIC_ABI_DIR: if set, directory whose files are served from /static/ (place abi/RandomWalkNFT.json etc. here).
func registerStaticAssetRoutes(r *gin.Engine) {
	if root := strings.TrimSpace(os.Getenv("NFT_ASSETS_ROOT")); root != "" {
		abs, err := filepath.Abs(root)
		if err != nil {
			log.Printf("NFT_ASSETS_ROOT invalid: %v", err)
			return
		}
		r.Use(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/images/") {
				c.Header("Cache-Control", "public, max-age=3600, must-revalidate")
			}
			c.Next()
		})
		r.Static("/images", abs)
		log.Printf("Serving NFT assets from %s at /images/", abs)
	}

	if abiRoot := strings.TrimSpace(os.Getenv("STATIC_ABI_DIR")); abiRoot != "" {
		abs, err := filepath.Abs(abiRoot)
		if err != nil {
			log.Printf("STATIC_ABI_DIR invalid: %v", err)
			return
		}
		r.Static("/static", abs)
		log.Printf("Serving /static from %s", abs)
	}
}
