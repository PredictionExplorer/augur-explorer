package cosmicgame
import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"

)
func cosmic_game_unique_stakers_both(c *gin.Context) {

	if  !dbInitialized() {
		common.RespondError(c,"Database link wasn't configured")
		return
	}
	unique_stakers := arb_storagew.Get_unique_stakers_both()
	c.HTML(http.StatusOK, "cg_unique_stakers_both.html", gin.H{
		"UniqueStakersBoth" : unique_stakers,
	})
}
