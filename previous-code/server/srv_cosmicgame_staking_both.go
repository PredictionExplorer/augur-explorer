package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func cosmic_game_unique_stakers_both(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	unique_stakers := arb_storagew.Get_unique_stakers_both()
	c.HTML(http.StatusOK, "cg_unique_stakers_both.html", gin.H{
		"UniqueStakersBoth" : unique_stakers,
	})
}
