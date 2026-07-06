package randomwalk

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Current offers (API)
func apiRwalkCurrentOffers(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addr := rwContractAddrs()
	rwalk_aid := addr.RandomWalkAid
	market_aid := addr.MarketPlaceAid
	p_order_by := c.Param("order_by")
	var order_by int64
	if len(p_order_by) > 0 {
		var success bool
		order_by, success = common.ParseIntFromRemoteOrError(c, JSON, &p_order_by)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'order_by' parameter is not set")
		return
	}
	offers := rw_storagew.Get_active_offers(rwalk_aid, market_aid, int(order_by))
	c.JSON(http.StatusOK, gin.H{
		"status":   1,
		"error":    "",
		"Offers":   offers,
		"RWalkAid": rwalk_aid,
		"MarketAid": market_aid,
	})
}

// Floor price (API)
func apiRwalkFloorPrice(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addr := rwContractAddrs()
	rwalk_aid := addr.RandomWalkAid
	market_aid := addr.MarketPlaceAid
	p_rwalk_addr := addr.RandomWalk
	p_market_addr := addr.MarketPlace
	_, floor_price, _, _, err := rw_storagew.Get_floor_price(rwalk_aid, market_aid)
	var db_err string
	if err != nil {
		db_err = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    1,
		"error":     "",
		"FloorPrice": floor_price,
		"DBError":    db_err,
		"MarketAddr": p_market_addr,
		"RWalkAddr":  p_rwalk_addr,
		"RWalkAid":   rwalk_aid,
		"MarketAid":  market_aid,
	})
}
