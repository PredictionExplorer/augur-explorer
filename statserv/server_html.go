package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

	//. "github.com/PredictionExplorer/augur-explorer/dbs/balancerv2"
)
func respond_error_html(c *gin.Context,error_text string) {

	c.HTML(http.StatusBadRequest, "error.html", gin.H{
		"title": "Augur Markets: Error",
		"ErrDescr": error_text,
	})
}
func respond_error_json(c *gin.Context,error_text string) {

	c.JSON(http.StatusBadRequest, gin.H{
		"status": 0,
		"error": error_text,
	})
}
func main_page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Index Page",
	})
}
func bal_v2_poolinfo(c *gin.Context) {

	pool_id := c.Query("pool_id")
	pool_aid,err := storagew.Lookup_pool_address_id(pool_id)
	if err != nil {
		respond_error_html(c,"No outcome provided")
	}
	pool_tokens := storagew.Get_pool_registered_tokens(pool_aid)
	c.HTML(http.StatusOK, "poolinfo.html", gin.H{
			"title": "Balancer v2 Pool Info",
			"PoolId" : pool_id,
			"Tokens" : pool_tokens,
	})
}
