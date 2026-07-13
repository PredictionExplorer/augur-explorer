// Package apitest hosts the frozen v1 parity suite and the v2 contract/golden
// suite. Both boot the real router against the same seeded testcontainers
// database (docs/MODERNIZATION.md §§4.1 and 6.2).
//
// The package has two halves:
//
//   - Unit half (no build tag): the v1 OpenAPI route-drift test, which builds
//     the shared route table (internal/api/routes) with no database.
//   - Integration half (//go:build integration): the shared harness, v1
//     parity goldens, and v2 kin-openapi-validated goldens.
package apitest

import (
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/faq"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
)

// buildBareRouter registers every route without touching a database.
// Handlers on the returned router must not be invoked; it exists so tests
// can enumerate the served route table.
func buildBareRouter() *httpx.Router {
	return routes.New(nil, routes.Options{
		CosmicGame: cosmicgame.NewBare(),
		RandomWalk: randomwalk.NewBare(),
		FAQ:        faq.New(faq.Options{UpstreamURL: "http://127.0.0.1:8000"}),
	})
}
