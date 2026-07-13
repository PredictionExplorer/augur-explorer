package randomwalk

import (
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// GET /api/randomwalk/contracts — marketplace + RandomWalk NFT contract addresses from rw_contracts (same source as ETL).
func (a *API) handleContracts(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":           1,
		"error":            "",
		"marketplace_addr": addrs.MarketPlace,
		"randomwalk_addr":  addrs.RandomWalk,
		"marketplace_aid":  addrs.MarketPlaceAid,
		"randomwalk_aid":   addrs.RandomWalkAid,
	})
}
