package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func rwalk_current_offers(c *gin.Context) {

	p_order_by := c.Param("order_by")
	var order_by int64
	if len(p_order_by) > 0 {
		var success bool
		order_by,success = parse_int_from_remote_or_error(c,HTTP,&p_order_by)
		if !success {
			return
		}
	} else {
		respond_error(c,"'order_by' parameter is not set")
		return
	}
	offers := augur_srv.db_arbitrum.Get_active_offers(int(order_by))

	c.HTML(http.StatusOK, "current_offers.html", gin.H{
		"Offers" : offers,
	})
}

