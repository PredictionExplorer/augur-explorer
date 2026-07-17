package httpx_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// ExampleRouter shows the two middleware positions: Use installs global
// middleware around the whole route table (the production chain composes
// CORS, panic recovery, access logging, compression, rate limiting and
// conditional requests this way — see internal/api/routes), while a
// per-route middleware argument wraps one handler only.
func ExampleRouter() {
	r := httpx.NewRouter()

	// Global middleware: first registered runs outermost.
	tagResponse := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("X-Service", "rwcg")
			next.ServeHTTP(w, req)
		})
	}
	r.Use(tagResponse)

	// Per-route middleware: guards this handler only.
	requireKey := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.Header.Get("X-API-Key") == "" {
				http.Error(w, "missing key", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
	r.GET("/rounds/{round}", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"round": c.Param("round")})
	}, requireKey)

	req := httptest.NewRequest(http.MethodGet, "/rounds/7", nil)
	req.Header.Set("X-API-Key", "secret")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	fmt.Println(rec.Code, rec.Header().Get("X-Service"))
	fmt.Println(rec.Body.String())

	denied := httptest.NewRequest(http.MethodGet, "/rounds/7", nil)
	deniedRec := httptest.NewRecorder()
	r.ServeHTTP(deniedRec, denied)
	fmt.Println(deniedRec.Code, deniedRec.Header().Get("X-Service"))
	// Output:
	// 200 rwcg
	// {"round":"7"}
	// 401 rwcg
}
