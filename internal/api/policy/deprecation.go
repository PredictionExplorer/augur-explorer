// Package policy owns API-surface classifications shared by serving and
// operational clients without pulling the full router into either side.
package policy

import (
	"strings"
	"time"
)

// V1DeprecatedAt is the moment the frozen v1 API was marked deprecated in
// docs/openapi.yaml. It is the source of the RFC 9745 Deprecation header.
var V1DeprecatedAt = time.Date(2026, time.July, 16, 0, 0, 0, 0, time.UTC)

// V1MigrationGuideURL points clients at the v1-to-v2 endpoint mapping.
const V1MigrationGuideURL = "https://github.com/PredictionExplorer/augur-explorer/blob/master/docs/api-v2-migration.md"

// V1Deprecated reports whether path belongs to the deprecated frozen v1 API.
//
// The FAQ proxy is a separate service without a v2 replacement. Contract-
// pinned metadata and operational endpoints are outside the deprecated
// prefixes by construction.
func V1Deprecated(path string) bool {
	if strings.HasPrefix(path, "/api/cosmicgame/") {
		return !strings.HasPrefix(path, "/api/cosmicgame/faq/")
	}
	return strings.HasPrefix(path, "/api/randomwalk/")
}
