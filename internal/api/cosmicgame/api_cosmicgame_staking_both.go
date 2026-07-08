package cosmicgame

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

func api_cosmic_game_user_unique_stakers_both(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	unique_stakers, err := arbRepo.UniqueStakersBoth(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"UniqueStakersBoth": unique_stakers,
	})
}
