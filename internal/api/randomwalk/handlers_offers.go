package randomwalk

import (
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Current offers (API)
func (a *API) handleCurrentOffers(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addr, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addr.RandomWalkAid
	marketAid := addr.MarketPlaceAid
	pOrderBy := c.Param("order_by")
	var orderBy int64
	if len(pOrderBy) > 0 {
		var success bool
		orderBy, success = common.ParseIntFromRemoteOrError(c, JSON, &pOrderBy)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'order_by' parameter is not set")
		return
	}
	offers, err := a.repo.ActiveOffers(c.Request.Context(), rwalkAid, marketAid, int(orderBy))
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":    1,
		"error":     "",
		"Offers":    offers,
		"RWalkAid":  rwalkAid,
		"MarketAid": marketAid,
	})
}

// Floor price (API)
func (a *API) handleFloorPrice(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addr, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addr.RandomWalkAid
	marketAid := addr.MarketPlaceAid
	pRwalkAddr := addr.RandomWalk
	pMarketAddr := addr.MarketPlace
	noOffers, floorPrice, _, _, err := a.repo.FloorPrice(c.Request.Context(), rwalkAid, marketAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	// The legacy layer surfaced the driver's no-rows error text in DBError
	// when the order book was empty; clients may key off it.
	var dbErr string
	if noOffers {
		dbErr = legacyNoRowsText
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"FloorPrice": floorPrice,
		"DBError":    dbErr,
		"MarketAddr": pMarketAddr,
		"RWalkAddr":  pRwalkAddr,
		"RWalkAid":   rwalkAid,
		"MarketAid":  marketAid,
	})
}
