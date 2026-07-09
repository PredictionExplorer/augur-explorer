package randomwalk

import (
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Current offers (API)
func apiRwalkCurrentOffers(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addr, ok := rwContractAddrs(c)
	if !ok {
		return
	}
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
	offers, err := rwRepo.ActiveOffers(c.Request.Context(), rwalk_aid, market_aid, int(order_by))
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":    1,
		"error":     "",
		"Offers":    offers,
		"RWalkAid":  rwalk_aid,
		"MarketAid": market_aid,
	})
}

// Floor price (API)
func apiRwalkFloorPrice(c *httpx.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addr, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addr.RandomWalkAid
	market_aid := addr.MarketPlaceAid
	p_rwalk_addr := addr.RandomWalk
	p_market_addr := addr.MarketPlace
	no_offers, floor_price, _, _, err := rwRepo.FloorPrice(c.Request.Context(), rwalk_aid, market_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	// The legacy layer surfaced the driver's no-rows error text in DBError
	// when the order book was empty; clients may key off it.
	var db_err string
	if no_offers {
		db_err = legacyNoRowsText
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"FloorPrice": floor_price,
		"DBError":    db_err,
		"MarketAddr": p_market_addr,
		"RWalkAddr":  p_rwalk_addr,
		"RWalkAid":   rwalk_aid,
		"MarketAid":  market_aid,
	})
}
