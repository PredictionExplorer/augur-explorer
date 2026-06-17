package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
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

// resolveNFTStaticMount chooses how GET /images/... maps to disk.
// Nested (default): /images/randomwalk/<file> — Cosmic CDN and legacy URLs.
// Flat (NFT_ASSETS_FLAT_PATHS=1): /images/<file> — nfts.randomwalknft.com layout; fs root is the randomwalk folder.
func resolveNFTStaticMount(abs string) (mount string, fsRoot string) {
	flat := common.NFTAssetsFlatPaths()

	rwNested := filepath.Join(abs, "randomwalk")
	thumbsNested, _ := filepath.Glob(filepath.Join(rwNested, "*_black_thumb.jpg"))
	if len(thumbsNested) > 0 {
		if flat {
			log.Printf("NFT assets: found RandomWalk thumbs under %s (flat layout -> URL /images/<file>)", rwNested)
			return "/images", rwNested
		}
		log.Printf("NFT assets: found RandomWalk thumbs under %s (standard layout)", rwNested)
		return "/images", abs
	}
	thumbsRoot, _ := filepath.Glob(filepath.Join(abs, "*_black_thumb.jpg"))
	if len(thumbsRoot) > 0 {
		if flat {
			log.Printf("NFT assets: found RandomWalk thumbs in %s (flat layout -> URL /images/<file>)", abs)
			return "/images", abs
		}
		log.Printf("NFT assets: found RandomWalk thumbs in %s (compact layout -> URL /images/randomwalk/)", abs)
		return "/images/randomwalk", abs
	}
	if st, err := os.Stat(rwNested); err == nil && st.IsDir() {
		if flat {
			return "/images", rwNested
		}
		return "/images", abs
	}
	if filepath.Base(abs) == "randomwalk" {
		if flat {
			return "/images", abs
		}
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

// cosmicAssetPrefix scopes the per-seed package fallback to Cosmic Signature
// assets. webserv serves both RandomWalk and Cosmic Signature files from the
// same /images mount, and only Cosmic Signature uses the package layout, so the
// fallback must never apply to RandomWalk (or any other) paths.
const cosmicAssetPrefix = "new/cosmicsignature/"

// resolveAssetFile maps a request to a file on disk, translating the legacy /
// flat Cosmic Signature URLs onto whatever on-disk layout is present.
//
// The frontend and metadata only ever request these flat URLs:
//
//	new/cosmicsignature/0x<seed>.png              (full image)
//	new/cosmicsignature/0x<seed>.mp4              (video)
//	new/cosmicsignature/0x<seed>/thumb_card.webp  (thumbnail; also thumb_micro)
//
// Cosmic Signature assets have shipped in three layouts over time. We try the
// real requested path first (so anything addressing a nested file directly
// still works), then the candidates below in priority order:
//
//	flat files:        new/cosmicsignature/0x<seed>.png   / .mp4
//	current packages:  new/cosmicsignature/0x<seed>/images/web/full.webp,
//	                   .../images/web/preview.webp, .../images/source/master.png,
//	                   .../videos/web/main.mp4, .../videos/hq/main.mp4
//	legacy packages:   new/cosmicsignature/0x<seed>/image.png / video.mp4
//
// Some frontend callers omit the 0x prefix (e.g. the home banner), so the seed
// segment is normalized to 0x<seed> when it looks like bare hex. RandomWalk and
// every other subtree are unaffected: a missing file 404s as before.
func resolveAssetFile(rootAbs, urlRel string) (full string, ok bool) {
	if full, ok = safeFileUnderRoot(rootAbs, urlRel); ok {
		return full, true
	}

	if !strings.HasPrefix(urlRel, cosmicAssetPrefix) {
		return "", false
	}

	rest := strings.TrimPrefix(urlRel, cosmicAssetPrefix) // "0x<seed>.png" or "0x<seed>/thumb_card.webp"
	ext := strings.ToLower(filepath.Ext(rest))

	var candidates []string
	switch {
	case ext == ".png" && !strings.Contains(rest, "/"):
		seed := normalizeSeedSegment(strings.TrimSuffix(rest, filepath.Ext(rest)))
		candidates = []string{
			seed + "/images/web/full.webp",
			seed + "/images/source/master.png",
			seed + "/image.png", // legacy package layout
		}
	case ext == ".mp4" && !strings.Contains(rest, "/"):
		seed := normalizeSeedSegment(strings.TrimSuffix(rest, filepath.Ext(rest)))
		candidates = []string{
			seed + "/videos/web/main.mp4",
			seed + "/videos/hq/main.mp4",
			seed + "/video.mp4", // legacy package layout
		}
	case ext == ".webp" && strings.Contains(rest, "/thumb_"):
		seed := normalizeSeedSegment(rest[:strings.IndexByte(rest, '/')])
		candidates = []string{
			seed + "/images/web/preview.webp",
			seed + "/images/web/full.webp",
		}
	default:
		return "", false
	}

	for _, c := range candidates {
		if full, ok = safeFileUnderRoot(rootAbs, cosmicAssetPrefix+c); ok {
			return full, true
		}
	}
	return "", false
}

// normalizeSeedSegment prepends "0x" when seg looks like a bare hex seed, so
// callers that omit the prefix still match the on-disk 0x<seed> directory.
func normalizeSeedSegment(seg string) string {
	if strings.HasPrefix(seg, "0x") || strings.HasPrefix(seg, "0X") {
		return seg
	}
	if seg != "" && isHex(seg) {
		return "0x" + seg
	}
	return seg
}

func isHex(s string) bool {
	for _, r := range s {
		if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
			return false
		}
	}
	return true
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
			full, ok := resolveAssetFile(fsRoot, rel)
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
		} else if fsRoot != abs {
			log.Printf("Serving RandomWalk NFT files from %q at %s/ (flat URL layout, files under randomwalk/)", fsRoot, mount)
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
