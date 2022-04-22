package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	//. "github.com/PredictionExplorer/augur-explorer/dbs/balancerv2"
)
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
		return
	}
	pool_info := storagew.Get_pool_info(pool_id)
	pool_tokens := storagew.Get_pool_registered_tokens(pool_aid)
	c.HTML(http.StatusOK, "poolinfo.html", gin.H{
			"title": "Balancer v2 Pool Info",
			"PoolInfo" : pool_info,
			"Tokens" : pool_tokens,
	})
}
func bal_v2_pool_token_history(c *gin.Context) {

	pool_aid,success := parse_integer_param_or_error(c,"pool_aid",FMT_HTML)
	if !success { return }
	pool_id,err := storagew.Lookup_pool_id_by_addr_id(pool_aid)
	if err != nil {
		respond_error_html(c,"Can't find pool id for provided pool address id")
		return
	}
	tok_aid,success := parse_integer_param_or_error(c,"tok_aid",FMT_HTML)
	if !success { return }
	tok_addr,err := storagew.S.Layer1_lookup_address(tok_aid)
	if err!=nil {
		fmt.Printf("err=%v\n",err)
		respond_error_html(c,"Token address ID not found")
		return
	}
	balances:=storagew.Get_pool_token_balance_history(pool_aid,tok_aid)
	c.HTML(http.StatusOK, "tokbalhistory.html", gin.H{
			"title": "Balancer v2 Pool Info",
			"PoolAid" : pool_aid,
			"PoolId" : pool_id,
			"Balances" : balances,
			"TokenAid" : tok_aid,
			"TokenAddr" : tok_addr,
	})
}
