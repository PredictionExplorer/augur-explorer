package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func rwalk_current_offers(c *gin.Context) {

	offers := augur_srv.db_arbitrum.Get_active_offers()

	c.HTML(http.StatusOK, "current_offers.html", gin.H{
		"Offers" : offers,
	})
}

