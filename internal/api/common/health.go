package common

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// draining reports whether the process is shutting down. While draining,
// /readyz answers 503 so load balancers stop routing new traffic to this
// instance; /healthz stays 200 because the process is still alive and
// finishing its in-flight requests.
var draining atomic.Bool

// SetDraining marks the process as shutting down. It is one-way: readiness
// stays false until the process exits.
func SetDraining() { draining.Store(true) }

// SetServing starts a new server lifecycle in the ready-to-serve state. The
// production process calls it once before opening listeners; tests that invoke
// the process runner repeatedly use the same lifecycle boundary.
func SetServing() { draining.Store(false) }

// RegisterHealthRoutes adds liveness and readiness probes to the public router.
// st may be nil, in which case /readyz always reports unready.
func RegisterHealthRoutes(r *httpx.Router, st *store.Store) {
	var ping func(context.Context) error
	if st != nil {
		ping = func(ctx context.Context) error {
			return st.Pool().Ping(ctx)
		}
	}
	registerHealthRoutes(r, ping)
}

// registerHealthRoutes keeps readiness behavior deterministic in tests without
// requiring a live PostgreSQL pool.
func registerHealthRoutes(r *httpx.Router, ping func(context.Context) error) {
	// Liveness: the process is up and serving.
	r.GET("/healthz", func(c *httpx.Context) {
		c.String(http.StatusOK, "ok")
	})
	// Readiness: not draining and dependencies are reachable (database ping).
	r.GET("/readyz", func(c *httpx.Context) {
		if draining.Load() {
			c.JSON(http.StatusServiceUnavailable, httpx.H{"status": "draining"})
			return
		}
		if ping == nil {
			c.JSON(http.StatusServiceUnavailable, httpx.H{"status": "unready", "reason": "database not configured"})
			return
		}
		if err := ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, httpx.H{"status": "unready", "reason": err.Error()})
			return
		}
		c.JSON(http.StatusOK, httpx.H{"status": "ready"})
	})
}
