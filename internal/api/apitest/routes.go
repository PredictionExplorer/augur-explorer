// Package apitest hosts the v1 API parity suite: it boots the real router
// against a seeded testcontainers database and pins every route's JSON
// response as a golden file (docs/MODERNIZATION.md §4.1).
//
// The package has two halves:
//
//   - Unit half (no build tag): the OpenAPI route-drift test, which builds
//     the shared route table (internal/api/routes) with no database.
//   - Integration half (//go:build integration): the parity harness, seed
//     fixtures and golden tests, run via `make test-integration`.
package apitest

import (
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
)

// buildBareRouter registers every route without initializing module state or
// touching a database. Handlers on the returned router must not be invoked;
// it exists so tests can enumerate the served route table.
func buildBareRouter() *httpx.Router {
	// Route registration is gated behind each module's enable flag; force
	// them on without running the DB/RPC-dependent Init functions.
	cosmicgame.Enabled = true
	randomwalk.Init(true)
	faq.Enabled = true

	return routes.New(nil, routes.Options{})
}
