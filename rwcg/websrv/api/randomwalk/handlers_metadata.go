package randomwalk

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// GET /api/randomwalk/metadata/:token_id — ERC-721-style JSON for marketplaces (image + animation_url).
// Requires NFT_ASSETS_PUBLIC_BASE (e.g. https://example.com/images or http://localhost:8080/images).
func apiRandomwalkTokenMetadata(c *gin.Context) {
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	base := strings.TrimRight(strings.TrimSpace(os.Getenv("NFT_ASSETS_PUBLIC_BASE")), "/")
	if base == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "NFT_ASSETS_PUBLIC_BASE is not set (public base URL for /images, no trailing slash)",
		})
		return
	}
	p := c.Param("token_id")
	var tokenID int64
	if _, err := fmt.Sscanf(p, "%d", &tokenID); err != nil || tokenID <= 0 {
		common.RespondErrorJSON(c, "invalid token_id")
		return
	}
	addrs := rw_storagew.Get_randomwalk_contract_addresses()
	info, err := rw_storagew.Get_rwalk_token_info(addrs.RandomWalkAid, tokenID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "token not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pad := fmt.Sprintf("%06d", tokenID)
	name := strings.TrimSpace(info.CurName)
	if name == "" {
		name = fmt.Sprintf("Random Walk #%s", pad)
	}
	meta := gin.H{
		"name":            name,
		"description":     "Random Walk NFT",
		"image":           fmt.Sprintf("%s/randomwalk/%s_black.png", base, pad),
		"animation_url":   fmt.Sprintf("%s/randomwalk/%s_black_single.mp4", base, pad),
		"external_url":    fmt.Sprintf("https://randomwalknft.com/detail/%d", tokenID),
		"attributes":      []gin.H{{"trait_type": "seed", "value": info.SeedHex}},
		"properties":      gin.H{"seed": info.SeedHex},
	}
	c.JSON(http.StatusOK, meta)
}
