package cosmicgame

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// TokenMetadataHandler is the exported entry point for the bare ERC-721
// tokenURI route (GET /metadata/:token_id), dispatched by host in the main
// router. On the Cosmic Signature host it serves Cosmic Signature metadata.
func TokenMetadataHandler(c *httpx.Context) {
	api_cosmic_game_cst_metadata(c)
}

// GET /api/cosmicgame/cst/metadata/:token_id — OpenSea-compatible metadata JSON (image hosted under /images/...).
// Uses the same token row as /cst/info. Image base defaults to this API's origin + /images; optional NFT_ASSETS_PUBLIC_BASE overrides.
func api_cosmic_game_cst_metadata(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !Enabled || !dbInitialized() {
		common.RespondErrorJSON(c, "CosmicGame module or database not available")
		return
	}
	base := common.NFTImagePublicBase(c)
	if base == "" {
		c.JSON(http.StatusInternalServerError, httpx.H{
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
	tokenInfo, err := arbRepo.CosmicSignatureTokenInfo(c.Request.Context(), tokenID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			c.JSON(http.StatusNotFound, httpx.H{"error": "record not found"})
			return
		}
		respondStoreError(c, err)
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

	// Immutable, marketplace-filterable traits. "Imprinted" uses the mint
	// timestamp (unix seconds) and is omitted when unavailable.
	attributes := []httpx.H{
		{"trait_type": "Round", "display_type": "number", "value": tokenInfo.RoundNum},
	}
	if tokenInfo.Tx.TimeStamp > 0 {
		attributes = append(attributes, httpx.H{
			"trait_type":   "Imprinted",
			"display_type": "date",
			"value":        tokenInfo.Tx.TimeStamp,
		})
	}
	attributes = append(attributes, httpx.H{"trait_type": "seed", "value": tokenInfo.Seed})

	meta := httpx.H{
		"name":          name,
		"description":   desc,
		"image":         image,
		"animation_url": animationURL,
		// Art is rendered on pure black; matches the marketplace frame.
		"background_color": "000000",
		"external_url":     fmt.Sprintf("https://www.cosmicsignature.com/detail/%d", tokenID),
		"attributes":       attributes,
		"properties": httpx.H{
			"seed":      tokenInfo.Seed,
			"token_id":  tokenID,
			"owner":     tokenInfo.CurOwnerAddr,
			"round_num": tokenInfo.RoundNum,
		},
	}
	c.JSON(http.StatusOK, meta)
}
