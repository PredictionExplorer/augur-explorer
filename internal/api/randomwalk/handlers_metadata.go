package randomwalk

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
// router. On RandomWalk hosts it serves RandomWalk metadata.
func TokenMetadataHandler(c *httpx.Context) {
	apiRandomwalkTokenMetadata(c)
}

// GET /api/randomwalk/metadata/:token_id and GET /metadata/:token_id — ERC-721 metadata JSON.
// On-chain baseURI is often https://<api-host>/metadata/ (e.g. legacy randomwalknft-api.com or api1.randomwalknft.com).
// image/animation_url use MetadataRandomWalkImagePublicBase (default API /images/...; override with NFT_ASSETS_PUBLIC_BASE).
func apiRandomwalkTokenMetadata(c *httpx.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	base := common.MetadataRandomWalkImagePublicBase()
	if base == "" {
		c.JSON(http.StatusInternalServerError, httpx.H{
			"error": "cannot derive public /images base URL (set NFT_ASSETS_PUBLIC_BASE)",
		})
		return
	}
	p := c.Param("token_id")
	var tokenID int64
	n, err := fmt.Sscanf(p, "%d", &tokenID)
	if err != nil || n != 1 || tokenID < 0 {
		common.RespondErrorJSON(c, "invalid token_id")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	info, err := rwRepo.TokenInfo(c.Request.Context(), addrs.RandomWalkAid, tokenID)
	if errors.Is(err, store.ErrNotFound) {
		c.JSON(http.StatusNotFound, httpx.H{"error": "token not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpx.H{"error": err.Error()})
		return
	}
	pad := fmt.Sprintf("%06d", tokenID)
	name := strings.TrimSpace(info.CurName)
	if name == "" {
		name = fmt.Sprintf("Random Walk #%s", pad)
	}
	var imageURL, animationURL string
	if common.NFTAssetsFlatPaths() {
		imageURL = fmt.Sprintf("%s/%s_black.png", base, pad)
		animationURL = fmt.Sprintf("%s/%s_black_single.mp4", base, pad)
	} else {
		imageURL = fmt.Sprintf("%s/randomwalk/%s_black.png", base, pad)
		animationURL = fmt.Sprintf("%s/randomwalk/%s_black_single.mp4", base, pad)
	}
	meta := httpx.H{
		"name":          name,
		"description":   "Random Walk NFT",
		"image":         imageURL,
		"animation_url": animationURL,
		"external_url":  fmt.Sprintf("https://randomwalknft.com/detail/%d", tokenID),
		"attributes":    []httpx.H{{"trait_type": "seed", "value": info.SeedHex}},
		"properties":    httpx.H{"seed": info.SeedHex},
	}
	c.JSON(http.StatusOK, meta)
}
