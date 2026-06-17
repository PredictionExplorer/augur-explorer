package cosmicgame

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// TokenMetadataHandler is the exported entry point for the bare ERC-721
// tokenURI route (GET /metadata/:token_id), dispatched by host in the main
// router. On the Cosmic Signature host it serves Cosmic Signature metadata.
func TokenMetadataHandler(c *gin.Context) {
	api_cosmic_game_cst_metadata(c)
}

// GET /api/cosmicgame/cst/metadata/:token_id — OpenSea-compatible metadata JSON (image hosted under /images/...).
// Uses the same token row as /cst/info. Image base defaults to this API's origin + /images; optional NFT_ASSETS_PUBLIC_BASE overrides.
func api_cosmic_game_cst_metadata(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !Enabled || !dbInitialized() {
		common.RespondErrorJSON(c, "CosmicGame module or database not available")
		return
	}
	base := common.NFTImagePublicBase(c)
	if base == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot derive public /images base URL (set Host or NFT_ASSETS_PUBLIC_BASE)",
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
	// animation_url is the marketplace-standard field (OpenSea et al.) for the
	// animated/video representation; built like image but pointing to the .mp4.
	animationURL := fmt.Sprintf("%s/new/cosmicsignature/%s.mp4", base, seedHex)
	desc := fmt.Sprintf(
		"Discover the unique attributes and ownership history of Cosmic Signature Token #%d, an exclusive digital collectible from the Cosmic Signature game.",
		tokenID,
	)
	name := strings.TrimSpace(tokenInfo.TokenName)
	if name == "" {
		name = fmt.Sprintf("Cosmic Signature #%d", tokenID)
	}
	meta := gin.H{
		"name":          name,
		"description":   desc,
		"image":         image,
		"animation_url": animationURL,
		"external_url":  fmt.Sprintf("https://www.cosmicsignature.com/detail/%d", tokenID),
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
