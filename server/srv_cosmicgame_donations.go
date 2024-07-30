package main
import (
	"net/http"
	`"github.com/gin-gonic/gin"

)
func cosmic_game_donations_eth_simple_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game()
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_simple_list.html", gin.H{
		"CosmicGameDonations" : donations,
	})
}
