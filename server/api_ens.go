/// API v1
package main
import (
	"fmt"

	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common"
	ens "github.com/wealdtech/go-ens/v3"
)
func a1_ens_name_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	fqdn:= c.Param("fqdn")

	ens_info,err := augur_srv.db_augur.Get_ens_record_info(fqdn)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("%v",err),
		})
		return
	}

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error" : err_str,
			"ENSInfo" : ens_info,
	})
}
func a1_ens_name_lookup(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}

	names,total_names := augur_srv.db_augur.Lookup_ens_names(eoa_aid)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error" : err_str,
			"Names" : names,
			"TotalRows" : total_names,
	})
}
func a1_user_ens_names(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	user_info,err := augur_srv.db_augur.Get_user_info(user_aid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("Error getting UserInfo: %v",err.Error()),
		})
	}


	active_names,active_total_rows := augur_srv.db_augur.Get_user_ens_names_active(user_aid,offset,limit)
	inactive_names,inactive_total_rows := augur_srv.db_augur.Get_user_ens_names_inactive(user_aid,offset,limit)
	addr_changes,achanges_total_rows := augur_srv.db_augur.Get_user_address_change_history(user_aid,offset,limit)
	ownership_changes,own_changes_total_rows := augur_srv.db_augur.Get_user_ownership_change_history(user_aid,offset,limit)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error" : err_str,
			"UserInfo" : user_info,
			"ENS_Names_Active" : active_names,
			"ENS_Names_Inactive" : inactive_names,
			"ENS_OwnershipChanges" : ownership_changes,
			"ENS_AddrChanges" : addr_changes,
			"TotalRowsActive" : active_total_rows,
			"TotalRowsInactive" :inactive_total_rows,
			"TotalRowsAddrChanges" : achanges_total_rows,
			"TotalRowsOwnershipChanges" : own_changes_total_rows,
	})
}
func a1_node_text_key_value_pairs(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	node := c.Param("node")
	fqdn,key_value_pairs:= augur_srv.db_augur.Get_node_text_key_values(node)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"error" : err_str,
		"Node" : node,
		"FullName" : fqdn,
		"KeyValuePairs" : key_value_pairs,
	})
}
func a1_ens_reverse_lookup(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_address:= c.Param("address")
/*	address_str,aid,success := json_validate_and_lookup_address_or_aid(c,&p_address)
	if !success {
		return
	}*/
	addr := common.HexToAddress(p_address)
	name, err := ens.ReverseResolve(rpcclient, addr)
	Info.Printf("reverse lookup of %v, name=%v\n",addr.String(),name)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"Addr" : p_address,
			"Name" : name,
			"status": status,
			"error": err_str,
	})
}
