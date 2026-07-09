package randomwalk

import (
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// GET /api/randomwalk/contracts — marketplace + RandomWalk NFT contract addresses from rw_contracts (same source as ETL).
func apiRwalkContracts(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
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
