package main
import (
	"net/http"
	"github.com/gin-gonic/gin"


)
func main_page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Index Page",
	})
}
func bal_v2_poolinfo(c *gin.Context) {

	pool_id := c.Query("pool_id")
	c.HTML(http.StatusOK, "poolinfo.html", gin.H{
			"title": "Balancer v2 Pool Info",
			"PoolId" : pool_id,
	})
}
