package cosmicgame

import (
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	marketstore "github.com/PredictionExplorer/augur-explorer/internal/store/marketplace"
)

// cosmicSignatureMarketOffer is the exact upstream v1 marketplace row. Its
// exported field names intentionally define the legacy JSON keys.
type cosmicSignatureMarketOffer struct {
	OfferID    int64 `json:"OfferId"`
	OfferType  int
	TokenID    int64 `json:"TokenId"`
	SellerAddr string
	BuyerAddr  string
	Active     bool
	Price      float64
	BlockNum   int64
	TimeStamp  int64
	DateTime   string
}

func cosmicSignatureMarketOffers(
	records []marketstore.LegacyOffer,
) []cosmicSignatureMarketOffer {
	output := make([]cosmicSignatureMarketOffer, 0, len(records))
	for i := range records {
		record := records[i]
		output = append(output, cosmicSignatureMarketOffer{
			OfferID:    record.OfferID,
			OfferType:  record.OfferType,
			TokenID:    record.TokenID,
			SellerAddr: record.SellerAddress,
			BuyerAddr:  record.BuyerAddress,
			Active:     record.Active,
			Price:      record.PriceETH,
			BlockNum:   record.BlockNumber,
			TimeStamp:  record.TimeStamp,
			DateTime:   record.LegacyDateTime,
		})
	}
	return output
}

// handleMarketplaceCurrentOffers implements the upstream compatibility route
// GET /api/cosmicgame/marketplace/current_offers/{order_by}.
func (a *API) handleMarketplaceCurrentOffers(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
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
	records, scope, err := a.repo.CosmicSignatureMarketplaceActiveOffers(
		c.Request.Context(),
		int(orderBy),
	)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":    1,
		"error":     "",
		"Offers":    cosmicSignatureMarketOffers(records),
		"NftAid":    scope.CollectionAid,
		"MarketAid": scope.MarketplaceAid,
	})
}

// handleMarketplaceFloorPrice implements the upstream compatibility route
// GET /api/cosmicgame/marketplace/floor_price.
func (a *API) handleMarketplaceFloorPrice(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	_, floorPrice, _, _, scope, err := a.repo.CosmicSignatureMarketplaceFloorPriceETH(
		c.Request.Context(),
	)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"FloorPrice": floorPrice,
		"DBError":    "",
		"NftAid":     scope.CollectionAid,
		"MarketAid":  scope.MarketplaceAid,
	})
}

// handleMarketplaceSales implements the upstream compatibility route
// GET /api/cosmicgame/marketplace/trading/sales/{offset}/{limit}.
func (a *API) handleMarketplaceSales(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	records, scope, err := a.repo.CosmicSignatureMarketplaceSales(
		c.Request.Context(),
		offset,
		limit,
	)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":    1,
		"error":     "",
		"Trading":   cosmicSignatureMarketOffers(records),
		"NftAid":    scope.CollectionAid,
		"MarketAid": scope.MarketplaceAid,
	})
}
