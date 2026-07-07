// Package apitest hosts the v1 API parity suite: it boots the real gin
// router against a seeded testcontainers database and pins every route's
// JSON response as a golden file (docs/MODERNIZATION.md §4.1).
//
// The package has two halves:
//
//   - Unit half (no build tag): route registration shared with the real
//     server and the OpenAPI route-drift test, which needs no database.
//   - Integration half (//go:build integration): the parity harness, seed
//     fixtures and golden tests, run via `make test-integration`.
package apitest

import (
	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// registerAllRoutes mirrors the route surface of cmd/apiserver/main.go: the
// three module route sets, health probes, and the host-dispatched bare
// /metadata route. Static-asset routes (env-gated, excluded from the OpenAPI
// spec) are intentionally not registered.
//
// Keep in sync with cmd/apiserver/main.go until Phase 2 extracts a shared
// router constructor.
func registerAllRoutes(r *gin.Engine, storage *store.SQLStorage) {
	common.RegisterHealthRoutes(r, storage)

	randomwalk.RegisterAPIRoutes(r)
	cosmicgame.RegisterAPIRoutes(r)
	faq.RegisterAPIRoutes(r)

	r.GET("/metadata/:token_id", func(c *gin.Context) {
		if common.MetadataHostServesCosmicSignature(c.Request.Host, c.Request.Header.Get("X-Forwarded-Host")) {
			cosmicgame.TokenMetadataHandler(c)
			return
		}
		randomwalk.TokenMetadataHandler(c)
	})
}

// buildBareRouter registers every route without initializing module state or
// touching a database. Handlers on the returned router must not be invoked;
// it exists so tests can enumerate the served route table.
func buildBareRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	// Route registration is gated behind each module's enable flag; force
	// them on without running the DB/RPC-dependent Init functions.
	cosmicgame.Enabled = true
	randomwalk.Init(true)
	faq.Enabled = true

	r := gin.New()
	registerAllRoutes(r, nil)
	return r
}
