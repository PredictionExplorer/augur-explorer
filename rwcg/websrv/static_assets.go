package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// warnNFTAssetsLayout logs when we cannot find RandomWalk thumbs in the expected places.
func warnNFTAssetsLayout(abs string, mount string) {
	rwNested := filepath.Join(abs, "randomwalk")
	thumbsNested, _ := filepath.Glob(filepath.Join(rwNested, "*_black_thumb.jpg"))
	thumbsRoot, _ := filepath.Glob(filepath.Join(abs, "*_black_thumb.jpg"))
	if len(thumbsNested) > 0 || len(thumbsRoot) > 0 {
		return
	}
	log.Printf("NFT_ASSETS_ROOT: no *_black_thumb.jpg under %q or %q (detail page images will 404 until assets exist).", rwNested, abs)
	log.Printf("Current mount: %s -> filesystem root %q", mount, abs)
}

// resolveNFTStaticMount chooses how /images/randomwalk/<file> maps to disk.
// 1) Thumbs under .../randomwalk/*.jpg  -> mount /images, root = parent (Cosmic /images/new/ still works).
// 2) Thumbs directly in NFT_ASSETS_ROOT -> mount /images/randomwalk, root = that dir (inner randomwalk folder).
// 3) Heuristics when no thumbs yet (empty DB / fresh checkout).
func resolveNFTStaticMount(abs string) (mount string, fsRoot string) {
	rwNested := filepath.Join(abs, "randomwalk")
	thumbsNested, _ := filepath.Glob(filepath.Join(rwNested, "*_black_thumb.jpg"))
	if len(thumbsNested) > 0 {
		log.Printf("NFT assets: found RandomWalk thumbs under %s (standard layout)", rwNested)
		return "/images", abs
	}
	thumbsRoot, _ := filepath.Glob(filepath.Join(abs, "*_black_thumb.jpg"))
	if len(thumbsRoot) > 0 {
		log.Printf("NFT assets: found RandomWalk thumbs in %s (compact layout -> URL /images/randomwalk/)", abs)
		return "/images/randomwalk", abs
	}
	if st, err := os.Stat(rwNested); err == nil && st.IsDir() {
		return "/images", abs
	}
	if filepath.Base(abs) == "randomwalk" {
		return "/images/randomwalk", abs
	}
	return "/images", abs
}

func safeFileUnderRoot(rootAbs, urlRel string) (full string, ok bool) {
	if strings.Contains(urlRel, "..") {
		return "", false
	}
	rel := filepath.FromSlash(urlRel)
	full = filepath.Join(rootAbs, rel)
	rootClean := filepath.Clean(rootAbs)
	fileClean := filepath.Clean(full)
	if rootClean == fileClean {
		return "", false
	}
	sep := string(os.PathSeparator)
	if !strings.HasPrefix(fileClean+sep, rootClean+sep) {
		return "", false
	}
	st, err := os.Stat(fileClean)
	if err != nil || st.IsDir() {
		return "", false
	}
	return fileClean, true
}

// registerStaticAssetRoutes serves nft-assets mirror at GET /images/*.
// NFT_ASSETS_ROOT: parent dir containing randomwalk/, or the randomwalk dir itself (see resolveNFTStaticMount).
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
			path := c.Request.URL.Path
			isImg := strings.HasPrefix(path, "/images/")
			if isImg {
				// Browsers cache aggressively on max-age; reload often skips the network → no line in websrv.
				// Set WEBSRV_IMAGE_NO_CACHE=1 in dev to force revalidation / no-store so each reload hits the API.
				if strings.TrimSpace(os.Getenv("WEBSRV_IMAGE_NO_CACHE")) == "1" {
					c.Header("Cache-Control", "no-store")
				} else {
					c.Header("Cache-Control", "public, max-age=3600, must-revalidate")
				}
			}
			c.Next()
			if !isImg || c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
				return
			}
			st := c.Writer.Status()
			if strings.TrimSpace(os.Getenv("WEBSRV_LOG_IMAGE_REQUESTS")) == "1" || st >= 400 {
				log.Printf("[images] %d %s", st, path)
			}
		})

		mount, fsRoot := resolveNFTStaticMount(abs)
		warnNFTAssetsLayout(abs, mount)

		handler := func(c *gin.Context) {
			rel := strings.TrimPrefix(c.Param("filepath"), "/")
			if rel == "" {
				c.Status(http.StatusNotFound)
				return
			}
			full, ok := safeFileUnderRoot(fsRoot, rel)
			if !ok {
				c.Status(http.StatusNotFound)
				return
			}
			c.File(full)
		}
		r.GET(mount+"/*filepath", handler)
		r.HEAD(mount+"/*filepath", handler)

		if mount == "/images/randomwalk" {
			log.Printf("Serving RandomWalk NFT files from %q at %s/ (token assets live directly in NFT_ASSETS_ROOT)", fsRoot, mount)
			log.Printf("Note: /images/new/... needs NFT_ASSETS_ROOT set to the parent of randomwalk/ if you use Cosmic assets too.")
		} else {
			log.Printf("Serving NFT assets from %q at %s/ (expect randomwalk/ and optional new/ under that root)", fsRoot, mount)
		}
		if strings.TrimSpace(os.Getenv("WEBSRV_IMAGE_NO_CACHE")) == "1" {
			log.Printf("WEBSRV_IMAGE_NO_CACHE=1: /images responses use Cache-Control: no-store (each reload should hit websrv).")
		} else {
			log.Printf("Image GETs may not appear on reload: browsers cache /images for max-age=3600 (no TCP = no log). Use DevTools \"Disable cache\", or WEBSRV_IMAGE_NO_CACHE=1, or WEBSRV_LOG_IMAGE_REQUESTS=1 for [images] lines when requests do occur.")
		}
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
