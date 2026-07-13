package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/config"
)

// staticAssets carries the static-file serving configuration and the process
// logger (previously scattered os.Getenv reads).
type staticAssets struct {
	// nftAssetsRoot enables GET /images/* when non-empty (NFT_ASSETS_ROOT).
	nftAssetsRoot string
	// flatPaths selects the flat /images/<file> layout (NFT_ASSETS_FLAT_PATHS).
	flatPaths bool
	// abiDir enables GET /static/* when non-empty (STATIC_ABI_DIR).
	abiDir string
	// noCache serves /images with Cache-Control: no-store (WEBSRV_IMAGE_NO_CACHE).
	noCache bool
	// logRequests logs successful image requests too (WEBSRV_LOG_IMAGE_REQUESTS).
	logRequests bool
	logger      *slog.Logger
}

// staticAssetsConfig maps the loaded service configuration onto the static
// asset options.
func staticAssetsConfig(cfg *config.APIServer, logger *slog.Logger) staticAssets {
	return staticAssets{
		nftAssetsRoot: strings.TrimSpace(cfg.NFTAssetsRoot),
		flatPaths:     cfg.NFTAssetsFlatPaths,
		abiDir:        strings.TrimSpace(cfg.StaticABIDir),
		noCache:       cfg.ImageNoCache,
		logRequests:   cfg.LogImageRequests,
		logger:        logger,
	}
}

// warnNFTAssetsLayout logs when we cannot find RandomWalk thumbs in the expected places.
func (sa staticAssets) warnNFTAssetsLayout(abs string, mount string) {
	rwNested := filepath.Join(abs, "randomwalk")
	thumbsNested, _ := filepath.Glob(filepath.Join(rwNested, "*_black_thumb.jpg"))
	thumbsRoot, _ := filepath.Glob(filepath.Join(abs, "*_black_thumb.jpg"))
	if len(thumbsNested) > 0 || len(thumbsRoot) > 0 {
		return
	}
	sa.logger.Warn(fmt.Sprintf("NFT_ASSETS_ROOT: no *_black_thumb.jpg under %q or %q (detail page images will 404 until assets exist).", rwNested, abs))
	sa.logger.Warn(fmt.Sprintf("Current mount: %s -> filesystem root %q", mount, abs))
}

// resolveNFTStaticMount chooses how GET /images/... maps to disk.
// Nested (default): /images/randomwalk/<file> — Cosmic CDN and legacy URLs.
// Flat (NFT_ASSETS_FLAT_PATHS): /images/<file> — nfts.randomwalknft.com layout; fs root is the randomwalk folder.
func (sa staticAssets) resolveNFTStaticMount(abs string) (mount string, fsRoot string) {
	flat := sa.flatPaths

	rwNested := filepath.Join(abs, "randomwalk")
	thumbsNested, _ := filepath.Glob(filepath.Join(rwNested, "*_black_thumb.jpg"))
	if len(thumbsNested) > 0 {
		if flat {
			sa.logger.Info(fmt.Sprintf("NFT assets: found RandomWalk thumbs under %s (flat layout -> URL /images/<file>)", rwNested))
			return "/images", rwNested
		}
		sa.logger.Info(fmt.Sprintf("NFT assets: found RandomWalk thumbs under %s (standard layout)", rwNested))
		return "/images", abs
	}
	thumbsRoot, _ := filepath.Glob(filepath.Join(abs, "*_black_thumb.jpg"))
	if len(thumbsRoot) > 0 {
		if flat {
			sa.logger.Info(fmt.Sprintf("NFT assets: found RandomWalk thumbs in %s (flat layout -> URL /images/<file>)", abs))
			return "/images", abs
		}
		sa.logger.Info(fmt.Sprintf("NFT assets: found RandomWalk thumbs in %s (compact layout -> URL /images/randomwalk/)", abs))
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
		isDigit := r >= '0' && r <= '9'
		isLower := r >= 'a' && r <= 'f'
		isUpper := r >= 'A' && r <= 'F'
		if !isDigit && !isLower && !isUpper {
			return false
		}
	}
	return true
}

// imageCacheAndLogMiddleware sets Cache-Control on /images/ responses and
// logs failed (or, when enabled, all) image requests. It self-filters on the
// path prefix, so it is registered globally like the legacy version was.
func (sa staticAssets) imageCacheAndLogMiddleware() httpx.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			isImg := strings.HasPrefix(path, "/images/")
			if isImg {
				// Browsers cache aggressively on max-age; reload often skips the network → no line in websrv.
				// Set WEBSRV_IMAGE_NO_CACHE in dev to force revalidation / no-store so each reload hits the API.
				if sa.noCache {
					w.Header().Set("Cache-Control", "no-store")
				} else {
					w.Header().Set("Cache-Control", "public, max-age=3600, must-revalidate")
				}
			}
			rw := httpx.WrapResponseWriter(w)
			next.ServeHTTP(rw, r)
			if !isImg || r.Method != http.MethodGet && r.Method != http.MethodHead {
				return
			}
			if st := rw.Status(); sa.logRequests || st >= 400 {
				sa.logger.Info(fmt.Sprintf("[images] %d %q", st, path))
			}
		})
	}
}

// registerStaticAssetRoutes returns the route registration hook serving the
// nft-assets mirror at GET /images/* and, when configured, /static/*.
//
// NFT_ASSETS_ROOT: parent dir containing randomwalk/, or the randomwalk dir
// itself (see resolveNFTStaticMount). STATIC_ABI_DIR: if set, directory whose
// files are served from /static/ (place abi/RandomWalkNFT.json etc. here).
func registerStaticAssetRoutes(sa staticAssets) func(*httpx.Router) {
	return func(r *httpx.Router) {
		if sa.nftAssetsRoot != "" {
			abs, err := filepath.Abs(sa.nftAssetsRoot)
			if err != nil {
				sa.logger.Error(fmt.Sprintf("NFT_ASSETS_ROOT invalid: %v", err))
				return
			}
			r.Use(sa.imageCacheAndLogMiddleware())

			mount, fsRoot := sa.resolveNFTStaticMount(abs)
			sa.warnNFTAssetsLayout(abs, mount)

			handler := func(c *httpx.Context) {
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
			r.GET(mount+"/{filepath...}", handler)
			r.HEAD(mount+"/{filepath...}", handler)

			switch {
			case mount == "/images/randomwalk":
				sa.logger.Info(fmt.Sprintf("Serving RandomWalk NFT files from %q at %s/ (token assets live directly in NFT_ASSETS_ROOT)", fsRoot, mount))
				sa.logger.Info("Note: /images/new/... needs NFT_ASSETS_ROOT set to the parent of randomwalk/ if you use Cosmic assets too.")
			case fsRoot != abs:
				sa.logger.Info(fmt.Sprintf("Serving RandomWalk NFT files from %q at %s/ (flat URL layout, files under randomwalk/)", fsRoot, mount))
			default:
				sa.logger.Info(fmt.Sprintf("Serving NFT assets from %q at %s/ (expect randomwalk/ and optional new/ under that root)", fsRoot, mount))
			}
			if sa.noCache {
				sa.logger.Info("WEBSRV_IMAGE_NO_CACHE: /images responses use Cache-Control: no-store (each reload should hit websrv).")
			} else {
				sa.logger.Info("Image GETs may not appear on reload: browsers cache /images for max-age=3600 (no TCP = no log). Use DevTools \"Disable cache\", or WEBSRV_IMAGE_NO_CACHE=1, or WEBSRV_LOG_IMAGE_REQUESTS=1 for [images] lines when requests do occur.")
			}
		}

		if sa.abiDir != "" {
			abs, err := filepath.Abs(sa.abiDir)
			if err != nil {
				sa.logger.Error(fmt.Sprintf("STATIC_ABI_DIR invalid: %v", err))
				return
			}
			// Files only: directory requests 404 (no listings on a public API).
			staticHandler := func(c *httpx.Context) {
				full, ok := safeFileUnderRoot(abs, c.Param("filepath"))
				if !ok {
					c.Status(http.StatusNotFound)
					return
				}
				c.File(full)
			}
			r.GET("/static/{filepath...}", staticHandler)
			r.HEAD("/static/{filepath...}", staticHandler)
			sa.logger.Info("serving /static", "dir", abs)
		}
	}
}
