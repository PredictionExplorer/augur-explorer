package randomwalk

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rwp "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// rwContractAddrs returns marketplace + RandomWalk addresses and AIDs from rw_contracts (same as ETL).
func rwContractAddrs() rwp.ContractAddresses {
	return rw_storagew.Get_randomwalk_contract_addresses()
}

// GET /api/randomwalk/contracts — marketplace + RandomWalk NFT contract addresses from rw_contracts (same source as ETL).
func apiRwalkContracts(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs := rw_storagew.Get_randomwalk_contract_addresses()
	c.JSON(http.StatusOK, gin.H{
		"status":            1,
		"error":             "",
		"marketplace_addr":  addrs.MarketPlace,
		"randomwalk_addr":   addrs.RandomWalk,
		"marketplace_aid":   addrs.MarketPlaceAid,
		"randomwalk_aid":    addrs.RandomWalkAid,
	})
}
