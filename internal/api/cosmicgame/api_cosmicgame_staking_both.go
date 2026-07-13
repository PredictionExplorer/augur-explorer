package cosmicgame

import (
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

func (a *API) handleUserUniqueStakersBoth(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	uniqueStakers, err := a.repo.UniqueStakersBoth(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"UniqueStakersBoth": uniqueStakers,
	})
}
