package randomwalk

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// Current offers (API + HTML)
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

func rwalk_current_offers(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	addr := rwContractAddrs()
	rwalk_aid := addr.RandomWalkAid
	market_aid := addr.MarketPlaceAid
	p_rwalk_addr := addr.RandomWalk
	p_market_addr := addr.MarketPlace
	p_order_by := c.Param("order_by")
	var order_by int64
	if len(p_order_by) > 0 {
		var success bool
		order_by, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_order_by)
		if !success {
			return
		}
	} else {
		common.RespondError(c, "'order_by' parameter is not set")
		return
	}
	offers := rw_storagew.Get_active_offers(rwalk_aid, market_aid, int(order_by))
	c.HTML(http.StatusOK, "rw_current_offers.html", gin.H{
		"Offers":     offers,
		"RWalkAid":   rwalk_aid,
		"RWalkAddr":  p_rwalk_addr,
		"MarketAid":  market_aid,
		"MarketAddr": p_market_addr,
	})
}

// Floor price (API + HTML)
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

func rwalk_floor_price(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
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
	c.HTML(http.StatusOK, "rw_floor_price.html", gin.H{
		"FloorPrice": floor_price,
		"DBError":    db_err,
		"MarketAddr": p_market_addr,
		"RWalkAddr":  p_rwalk_addr,
		"RWalkAid":   rwalk_aid,
		"MarketAid":  market_aid,
	})
}
