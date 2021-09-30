package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

)
func wrapped_token_info(c *gin.Context) {

	p_address := c.Param("address")
	wrapper_addr,valid:=is_address_valid(c,false,p_address)
	if !valid {
		return
	}
	wrapper_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(wrapper_addr)
	if err == nil {
		respond_error(c,fmt.Sprintf("Address %v not found",p_address))
		return
	}
	winfo,err := augur_srv.db_augur.Get_wrapped_token_info(wrapper_aid)
	if err != nil {
		respond_error(c,fmt.Sprintf("ShareToken wrapper with address %v not found",p_address))
		return
	}
	c.HTML(http.StatusOK, "wrapped_sharetokens/token_info.html", gin.H{
		"WrapperInfo" : winfo,
	})
}
func user_wrapped_token_transfers(c *gin.Context) {

	p_user:= c.Param("user")
	user_addr,valid := is_address_valid(c,false,p_user)
	if !valid {
		return
	}
	p_wrapper:= c.Param("wrapper")
	wrapper_addr,valid := is_address_valid(c,false,p_wrapper)
	if !valid {
		return
	}
	user_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(user_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",user_addr),
		})
		return
	}
	wrapper_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(wrapper_addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",wrapper_addr),
		})
		return
	}
	wrapper_info,_ := augur_srv.db_augur.Get_wrapped_token_info(wrapper_aid)
	market_info,err := augur_srv.db_augur.Get_market_info(wrapper_info.MktAddr,wrapper_info.OutcomeIdx,true)
	total_rows,transfers:= augur_srv.db_augur.Get_user_wrapped_shtoken_transfers(user_aid,wrapper_aid,0,10000)
	c.HTML(http.StatusOK, "user_wrapped_transfers.html", gin.H{
			"UserAddr" : user_addr,
			"MarketInfo" : market_info,
			"TokenInfo" : wrapper_info,
			"Transfers" : transfers,
			"TotalRows" : total_rows,
	})
}
func wrapped_tokens(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid := is_address_valid(c,false,market)
	if !valid {
		return
	}
	market_info,err := augur_srv.db_augur.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	wrappers := augur_srv.db_augur.Get_wrapped_tokens_for_market(market_info.MktAid)
	c.HTML(http.StatusOK, "wrapper_contracts.html", gin.H{
			"WrapperContracts" : wrappers,
			"Market": market_info,
	})
}
func wrapped_token_transfers(c *gin.Context) {

	address:= c.Param("address")
	addr,valid := is_address_valid(c,false,address)
	if !valid {
		return
	}
	aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(addr)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Such address wasn't found: %v",address),
		})
		return
	}
	wrapper_info,_ := augur_srv.db_augur.Get_wrapped_token_info(aid)
	market_info,err := augur_srv.db_augur.Get_market_info(wrapper_info.MktAddr,wrapper_info.OutcomeIdx,true)
	transfers,total_rows := augur_srv.db_augur.Get_wrapped_token_transfers(aid,0,500)
	c.HTML(http.StatusOK, "wrapped_sharetokens/transfers.html", gin.H{
			"MarketInfo" : market_info,
			"TokenInfo" : wrapper_info,
			"TotalRows" : total_rows,
			"WrappedTransfers" : transfers,
	})
}
func show_augur_foundry_contracts(c *gin.Context) {

	wrappers:= augur_srv.db_augur.Get_augur_foundry_wrapper_list()
	c.HTML(http.StatusOK, "augur_foundry_wrappers.html", gin.H{
		"ERC20MarketOutcomeWrappers" : wrappers,
	})
}
