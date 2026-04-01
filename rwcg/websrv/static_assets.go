package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// warnNFTAssetsLayout logs when NFT_ASSETS_ROOT is likely the inner `randomwalk` directory while
// the app serves URLs under /images/randomwalk/... (needs parent of `randomwalk` as root).
func warnNFTAssetsLayout(abs string) {
	nested := filepath.Join(abs, "randomwalk")
	if st, err := os.Stat(nested); err == nil && st.IsDir() {
		return
	}
	// Any *_{black,white}_thumb.jpg at root suggests files live directly in abs (wrong level).
	matches, _ := filepath.Glob(filepath.Join(abs, "*_black_thumb.jpg"))
	if len(matches) == 0 {
		matches, _ = filepath.Glob(filepath.Join(abs, "*_white_thumb.jpg"))
	}
	if len(matches) == 0 {
		return
	}
	parent := filepath.Dir(abs)
	log.Printf("NFT_ASSETS_ROOT layout: %q has token images at top level but no %q subdirectory.", abs, nested)
	log.Printf("Frontend uses /images/randomwalk/<file>. Set NFT_ASSETS_ROOT to the parent directory (e.g. %q) so files resolve under randomwalk/.", parent)
}

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
		warnNFTAssetsLayout(abs)
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
