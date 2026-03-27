package cosmicgame

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// GET /api/cosmicgame/cst/metadata/:token_id — OpenSea-compatible metadata JSON (image hosted under /images/...).
// Uses the same token row as /cst/info. Requires NFT_ASSETS_PUBLIC_BASE.
func api_cosmic_game_cst_metadata(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !Enabled || !dbInitialized() {
		common.RespondErrorJSON(c, "CosmicGame module or database not available")
		return
	}
	base := strings.TrimRight(strings.TrimSpace(os.Getenv("NFT_ASSETS_PUBLIC_BASE")), "/")
	if base == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "NFT_ASSETS_PUBLIC_BASE is not set (public base for /images, no trailing slash)",
		})
		return
	}
	p := c.Param("token_id")
	var tokenID int64
	if _, err := fmt.Sscanf(p, "%d", &tokenID); err != nil || tokenID <= 0 {
		common.RespondErrorJSON(c, "invalid token_id")
		return
	}
	ok, tokenInfo := arb_storagew.Get_cosmic_signature_token_info(tokenID)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	seedHex := strings.TrimSpace(tokenInfo.Seed)
	if !strings.HasPrefix(strings.ToLower(seedHex), "0x") {
		seedHex = "0x" + seedHex
	}
	image := fmt.Sprintf("%s/new/cosmicsignature/%s.png", base, seedHex)
	desc := fmt.Sprintf(
		"Discover the unique attributes and ownership history of Cosmic Signature Token #%d, an exclusive digital collectible from the Cosmic Signature game.",
		tokenID,
	)
	name := strings.TrimSpace(tokenInfo.TokenName)
	if name == "" {
		name = fmt.Sprintf("Cosmic Signature #%d", tokenID)
	}
	meta := gin.H{
		"name":         name,
		"description":  desc,
		"image":        image,
		"external_url": fmt.Sprintf("https://www.cosmicsignature.com/detail/%d", tokenID),
		"attributes": []gin.H{
			{"trait_type": "seed", "value": tokenInfo.Seed},
		},
		"properties": gin.H{
			"seed":      tokenInfo.Seed,
			"token_id":  tokenID,
			"owner":     tokenInfo.CurOwnerAddr,
			"round_num": tokenInfo.RoundNum,
		},
	}
	c.JSON(http.StatusOK, meta)
}
