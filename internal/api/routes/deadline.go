package routes

import (
	"net/http"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// RequestDeadlineFor returns the processing deadline of one request path:
// common.DefaultRequestDeadline for everything except the FAQ proxy, which
// is exempt (zero). The FAQ upstream is an LLM service whose legitimate
// answers exceed the API budget; its time bound is the proxy's own HTTP
// client timeout, and its size bound the proxy's buffered-response cap.
// Exported so the timeout tests and the drift tests pin the same policy the
// router installs.
func RequestDeadlineFor(path string) time.Duration {
	if strings.HasPrefix(path, "/api/cosmicgame/faq/") {
		return 0
	}
	return common.DefaultRequestDeadline
}

// requestDeadlinePolicy adapts RequestDeadlineFor to the middleware's
// per-request signature.
func requestDeadlinePolicy(r *http.Request) time.Duration {
	return RequestDeadlineFor(r.URL.Path)
}
