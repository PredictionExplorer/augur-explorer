package routes

import (
	"strings"
	"time"
)

// V1DeprecatedAt is the moment the frozen v1 API was marked deprecated in
// docs/openapi.yaml (docs-and-deprecation sprint). It is the single source
// of the RFC 9745 Deprecation header value on v1 responses.
// The removal date is separate: it flows in through Options.V1SunsetAt once
// the D6 sunset gates (consumer migration, 30 zero-traffic days, announced
// not-before date) produce one.
var V1DeprecatedAt = time.Date(2026, time.July, 16, 0, 0, 0, 0, time.UTC)

// V1MigrationGuideURL is the v1→v2 endpoint mapping served in the Link
// header of every deprecated v1 response.
const V1MigrationGuideURL = "https://github.com/PredictionExplorer/augur-explorer/blob/master/docs/api-v2-migration.md"

// V1Deprecated reports whether the request path belongs to the deprecated
// frozen v1 surface. It is the single policy the router middleware and the
// OpenAPI drift test share, so the spec's deprecated flags and the runtime
// headers can never disagree.
//
// The deprecated surface is everything under /api/cosmicgame/ and
// /api/randomwalk/ — except the FAQ proxy (/api/cosmicgame/faq/*), which
// fronts a separate service and has no v2 replacement. Everything else is
// out of scope by construction: /api/v2/* is the replacement surface, the
// metadata routes (/metadata/{token_id}, /cg/metadata/{token_id}) are
// contract-pinned tokenURI targets that can never move (D12), and
// /healthz, /readyz and /version are operational endpoints.
func V1Deprecated(path string) bool {
	if strings.HasPrefix(path, "/api/cosmicgame/") {
		return !strings.HasPrefix(path, "/api/cosmicgame/faq/")
	}
	return strings.HasPrefix(path, "/api/randomwalk/")
}
