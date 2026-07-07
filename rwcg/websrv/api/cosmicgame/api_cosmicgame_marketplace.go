package cosmicgame

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// Marketplace endpoints for Cosmic Signature NFTs traded on the shared
// RandomWalk marketplace contract. The offer/sale rows live in rw_new_offer /
// rw_item_bought (populated by the RandomWalk ETL, which indexes every NFT
// collection on the marketplace). Cosmic Signature offers are simply the rows
// whose rwalk_aid is the Cosmic Signature NFT contract.

// resolveMarketplaceAids resolves the Cosmic Signature NFT and marketplace
// contract addresses to their address ids. address ids are assigned per
// database, so they are resolved at runtime here rather than hardcoded.
func resolveMarketplaceAids(c *gin.Context) (nft_aid int64, market_aid int64, ok bool) {
	market_addr, err := arb_storagew.Get_marketplace_addr()
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Cannot read marketplace address from rw_contracts: %v", err))
		return 0, 0, false
	}
	nft_aid, err = arb_storagew.S.Nonfatal_lookup_address_id(cosmic_signature_addr.String())
	if err != nil {
		common.RespondErrorJSON(c, "Cosmic Signature NFT address is not indexed yet")
		return 0, 0, false
	}
	market_aid, err = arb_storagew.S.Nonfatal_lookup_address_id(market_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Marketplace address is not indexed yet")
		return 0, 0, false
	}
	return nft_aid, market_aid, true
}

// GET /api/cosmicgame/marketplace/current_offers/:order_by
func api_cosmic_game_marketplace_current_offers(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	nft_aid, market_aid, ok := resolveMarketplaceAids(c)
	if !ok {
		return
	}

	var order_by int64
	p_order_by := c.Param("order_by")
	if len(p_order_by) > 0 {
		var success bool
		order_by, success = common.ParseIntFromRemoteOrError(c, JSON, &p_order_by)
		if !success {
			return
		}
	}

	offers := arb_storagew.Get_marketplace_active_offers(nft_aid, market_aid, int(order_by))
	c.JSON(http.StatusOK, gin.H{
		"status":    1,
		"error":     "",
		"Offers":    offers,
		"NftAid":    nft_aid,
		"MarketAid": market_aid,
	})
}

// GET /api/cosmicgame/marketplace/floor_price
func api_cosmic_game_marketplace_floor_price(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	nft_aid, market_aid, ok := resolveMarketplaceAids(c)
	if !ok {
		return
	}

	_, floor_price, _, _, err := arb_storagew.Get_marketplace_floor_price(nft_aid, market_aid)
	var db_err string
	if err != nil {
		db_err = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     1,
		"error":      "",
		"FloorPrice": floor_price,
		"DBError":    db_err,
		"NftAid":     nft_aid,
		"MarketAid":  market_aid,
	})
}

// GET /api/cosmicgame/marketplace/trading/sales/:offset/:limit
func api_cosmic_game_marketplace_sales(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	nft_aid, market_aid, ok := resolveMarketplaceAids(c)
	if !ok {
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	sales := arb_storagew.Get_marketplace_sale_history(nft_aid, market_aid, offset, limit)
	c.JSON(http.StatusOK, gin.H{
		"status":    1,
		"error":     "",
		"Trading":   sales,
		"NftAid":    nft_aid,
		"MarketAid": market_aid,
	})
}
