package main
import (
	"net/http"
	`"github.com/gin-gonic/gin"

)
func api_cosmic_game_donations_eth(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	donations := arb_storagew.Get_donations_to_cosmic_game()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CharityDonations" : donations,
	})
}
