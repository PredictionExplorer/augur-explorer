// Package routes assembles the shared API router: global middleware, the
// frozen v1 table, and an optional generated v2 server. cmd/apiserver and the
// API integration suites build through this package, so production and test
// wiring cannot drift.
package routes

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/policy"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	v2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// Options carries the injected API modules and production-only extras. The
// zero value builds a router with the standard middleware and health probes
// but no module routes; a module is registered exactly when it is non-nil
// (replacing the legacy package-level enable flags).
type Options struct {
	// AccessLog enables the per-request structured log line (production).
	AccessLog *slog.Logger

	// PanicLog receives recovered handler panics; nil uses slog.Default().
	PanicLog *slog.Logger

	// CosmicGame registers the frozen v1 CosmicGame surface.
	CosmicGame *cosmicgame.API

	// RandomWalk registers the frozen v1 RandomWalk surface.
	RandomWalk *randomwalk.API

	// FAQ registers the FAQ bot proxy routes.
	FAQ *faq.Proxy

	// V2 registers the generated /api/v2 surface. Nil builds v1 only, which
	// is useful for the frozen v1 route-drift test.
	V2 *v2.Server

	// V1SunsetAt, when non-zero, announces the earliest v1 removal moment
	// through the RFC 8594 Sunset header on every deprecated v1 response
	// (config V1_SUNSET_AT). Zero — the default until the D6 sunset gates
	// are met — omits the header; the Deprecation header and migration
	// Link are emitted regardless.
	V1SunsetAt time.Time

	// Extra is appended to the global middleware chain after the standard
	// stack (production adds Prometheus request metrics here).
	Extra []httpx.Middleware

	// OnRequestTimeout observes every request whose processing deadline
	// expired and was answered with the 503 timeout rendering (production
	// counts them in rwcg_http_request_timeouts_total). Nil ignores them.
	OnRequestTimeout func(*http.Request)

	// RegisterExtra registers additional routes and their middleware
	// (production adds the env-gated static asset routes here).
	RegisterExtra func(*httpx.Router)
}

// New builds the standard middleware chain — CORS, panic recovery, optional
// access log, gzip compression, per-IP rate limiting, extras, the request
// processing deadline, the request body cap, conditional requests — and
// registers every injected module, the v2 surface, health probes, the
// version endpoint, and host-dispatched metadata.
// st may be nil (bare route-table construction for drift tests).
func New(st *store.Store, opts Options) *httpx.Router {
	r := httpx.NewRouter()

	// Outermost: every response under a deprecated v1 path — including
	// preflights, rate-limit rejections and 404s — announces the
	// deprecation (RFC 9745) and, once configured, the sunset date.
	r.Use(common.DeprecationHeaders(common.DeprecationPolicy{
		Match:        func(req *http.Request) bool { return policy.V1Deprecated(req.URL.Path) },
		DeprecatedAt: policy.V1DeprecatedAt,
		LinkURL:      policy.V1MigrationGuideURL,
		SunsetAt:     opts.V1SunsetAt,
	}))
	r.Use(common.CORS(), common.Recovery(opts.PanicLog))
	if opts.AccessLog != nil {
		r.Use(common.AccessLog(opts.AccessLog))
	}
	// Compression sits inside the access log (so logged bytes are wire
	// bytes) and outside everything that produces bodies.
	r.Use(common.Compress())
	// Baseline abuse protection: per-IP token bucket across the whole API.
	// Generous enough for legitimate frontends; mutating endpoints add their
	// own stricter limits at registration.
	r.Use(common.RateLimit(50, 100))
	r.Use(opts.Extra...)
	// Bounded request time: one deadline on the request context (inside the
	// metrics extras, so timeout 503s are observable like every response);
	// a post-deadline internal error renders as the per-family 503 timeout.
	// The FAQ proxy is policy-exempt — see RequestDeadlineFor.
	r.Use(common.RequestDeadline(requestDeadlinePolicy, opts.OnRequestTimeout))
	// Bounded request bodies: declared oversized bodies answer 413 here
	// (inside the metrics extras, so rejections are observable); undeclared
	// ones fail at first read past the cap inside the consuming handler.
	r.Use(common.MaxRequestBody(common.MaxRequestBodyBytes))
	// Innermost global: hashes identity bodies before compression, answers
	// If-None-Match revalidations with 304 and applies the default
	// Cache-Control policy where no layer chose one.
	r.Use(common.ConditionalETag())

	// Liveness/readiness probes and build identity.
	common.RegisterHealthRoutes(r, st)
	r.GET("/version", func(c *httpx.Context) {
		c.JSON(http.StatusOK, version.Get())
	})

	// Production extras (static assets) register before the API modules so
	// their subtree middleware lands in the documented chain position.
	if opts.RegisterExtra != nil {
		opts.RegisterExtra(r)
	}

	if opts.RandomWalk != nil {
		opts.RandomWalk.RegisterRoutes(r)
	}
	if opts.CosmicGame != nil {
		opts.CosmicGame.RegisterRoutes(r)
	}
	if opts.FAQ != nil {
		opts.FAQ.RegisterRoutes(r)
	}
	if opts.V2 != nil {
		opts.V2.RegisterRoutes(r)
	}

	// Bare ERC-721 tokenURI route. Both projects' on-chain baseURI is
	// https://<host>/metadata/, and this single webserv serves both the
	// RandomWalk and Cosmic Signature hosts, so dispatch by request Host:
	// a Cosmic Signature host serves Cosmic Signature metadata, anything
	// else (RandomWalk hosts) serves RandomWalk metadata. A disabled module
	// answers its legacy "module not available" envelope.
	r.GET("/metadata/{token_id}", func(c *httpx.Context) {
		if common.MetadataHostServesCosmicSignature(c.Request.Host, c.Request.Header.Get("X-Forwarded-Host")) {
			if opts.CosmicGame != nil {
				opts.CosmicGame.TokenMetadata(c)
				return
			}
			common.RespondErrorJSON(c, "CosmicGame module or database not available")
			return
		}
		if opts.RandomWalk != nil {
			opts.RandomWalk.TokenMetadata(c)
			return
		}
		common.RespondErrorJSON(c, "Database link wasn't configured")
	})

	return r
}
