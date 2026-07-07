package common

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dbs "github.com/PredictionExplorer/augur-explorer/internal/store"
)

// RegisterHealthRoutes adds liveness and readiness probes to the public router.
// db may be nil, in which case /readyz always reports unready.
func RegisterHealthRoutes(r *gin.Engine, db *dbs.SQLStorage) {
	// Liveness: the process is up and serving.
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	// Readiness: dependencies are reachable (database ping).
	r.GET("/readyz", func(c *gin.Context) {
		if db == nil || db.Db() == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unready", "reason": "database not configured"})
			return
		}
		if err := db.Db().PingContext(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unready", "reason": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})
}
