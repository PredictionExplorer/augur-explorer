package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func api_cosmic_game_user_unique_stakers_both(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	unique_stakers := arb_storagew.Get_unique_stakers_both()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UniqueStakersBoth" : unique_stakers,
	})
}
