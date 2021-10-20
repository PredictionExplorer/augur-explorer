/// API v1
package main
import (
	//"fmt"
	//"strconv"

	"net/http"
	"github.com/gin-gonic/gin"

)
func api_rwalk_current_offers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	offers := augur_srv.db_arbitrum.Get_active_offers()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offers" : offers,
	})
}

