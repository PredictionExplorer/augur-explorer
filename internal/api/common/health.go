package common

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// RegisterHealthRoutes adds liveness and readiness probes to the public router.
// st may be nil, in which case /readyz always reports unready.
func RegisterHealthRoutes(r *gin.Engine, st *store.Store) {
	// Liveness: the process is up and serving.
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	// Readiness: dependencies are reachable (database ping).
	r.GET("/readyz", func(c *gin.Context) {
		if st == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unready", "reason": "database not configured"})
			return
		}
		if err := st.Pool().Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unready", "reason": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})
}
