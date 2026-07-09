// Package routes assembles the v1 API router: the shared global middleware
// chain and the complete route table. cmd/apiserver and the API parity suite
// (internal/api/apitest) build their routers through this package, so the
// surface that is served in production is exactly the surface the golden
// tests pin — the two can no longer drift apart.
package routes

import (
	"log/slog"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Options carries the production-only extras. The zero value builds the
// test-harness router: full middleware chain and route table, but no access
// log, no metrics and no static assets.
type Options struct {
	// AccessLog enables the per-request structured log line (production).
	AccessLog *slog.Logger

	// PanicLog receives recovered handler panics; nil uses slog.Default().
	PanicLog *slog.Logger

	// Extra is appended to the global middleware chain after the standard
	// stack (production adds Prometheus request metrics here).
	Extra []httpx.Middleware

	// RegisterExtra registers additional routes and their middleware
	// (production adds the env-gated static asset routes here).
	RegisterExtra func(*httpx.Router)
}

// New builds the standard middleware chain — CORS, panic recovery, optional
// access log, per-IP rate limiting, extras — and registers every v1 route:
// health probes, the three API modules, and the host-dispatched bare
// /metadata route. Module routes honor their enable flags (module Init).
// st may be nil (bare route-table construction for drift tests).
func New(st *store.Store, opts Options) *httpx.Router {
	r := httpx.NewRouter()

	r.Use(common.CORS(), common.Recovery(opts.PanicLog))
	if opts.AccessLog != nil {
		r.Use(common.AccessLog(opts.AccessLog))
	}
	// Baseline abuse protection: per-IP token bucket across the whole API.
	// Generous enough for legitimate frontends; mutating endpoints add their
	// own stricter limits at registration.
	r.Use(common.RateLimit(50, 100))
	r.Use(opts.Extra...)

	// Liveness/readiness probes.
	common.RegisterHealthRoutes(r, st)

	// Production extras (static assets) register before the API modules so
	// their subtree middleware lands in the documented chain position.
	if opts.RegisterExtra != nil {
		opts.RegisterExtra(r)
	}

	randomwalk.RegisterAPIRoutes(r)
	cosmicgame.RegisterAPIRoutes(r)
	faq.RegisterAPIRoutes(r)

	// Bare ERC-721 tokenURI route. Both projects' on-chain baseURI is
	// https://<host>/metadata/, and this single webserv serves both the
	// RandomWalk and Cosmic Signature hosts, so dispatch by request Host:
	// a Cosmic Signature host serves Cosmic Signature metadata, anything
	// else (RandomWalk hosts) serves RandomWalk metadata.
	r.GET("/metadata/{token_id}", func(c *httpx.Context) {
		if common.MetadataHostServesCosmicSignature(c.Request.Host, c.Request.Header.Get("X-Forwarded-Host")) {
			cosmicgame.TokenMetadataHandler(c)
			return
		}
		randomwalk.TokenMetadataHandler(c)
	})

	return r
}
