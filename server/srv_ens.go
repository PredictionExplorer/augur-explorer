package main
import (
	"fmt"
	"net/http"
	"encoding/hex"
	"github.com/gin-gonic/gin"

	"github.com/wealdtech/go-ens/v3"
)
func user_ens_names(c *gin.Context) {

	user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,user)
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
	user_info,err := augur_srv.db_augur.Get_user_info(user_aid)
	active_names,active_total_rows := augur_srv.db_augur.Get_user_ens_names_active(user_aid,0,1000000)
	inactive_names,inactive_total_rows := augur_srv.db_augur.Get_user_ens_names_inactive(user_aid,0,1000000)
	addr_changes,achanges_total_rows := augur_srv.db_augur.Get_user_address_change_history(user_aid,0,1000000)
	ownership_changes,own_changes_total_rows := augur_srv.db_augur.Get_user_ownership_change_history(user_aid,0,1000000)
	c.HTML(http.StatusOK, "user_ens_names.html", gin.H{
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
func show_node_text_data(c *gin.Context) {

	node := c.Param("node")
	fqdn,key_value_pairs:= augur_srv.db_augur.Get_node_text_key_values(node)
	c.HTML(http.StatusOK, "user_text_kv_pairs.html", gin.H{
		"Node" : node,
		"FullName" : fqdn,
		"KeyValuePairs" : key_value_pairs,
	})
}
func ens_name_info(c *gin.Context) {

	fqdn := c.Param("fqdn")
	ens_info,err := augur_srv.db_augur.Get_ens_record_info(fqdn)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err),
		})
		return
	}
	if len(ens_info.ContentHash) > 0 {
		data,err := hex.DecodeString(ens_info.ContentHash)
		if err==nil {
			ens_info.ContentHash,err = ens.ContenthashToString(data[:])
			Error.Printf(
				"Content hash bianry string for node %v  has invalid bin fmt : %v\n",
				ens_info.FQDN,err,
			)
		} else {
			Error.Printf(
				"Content hash bianry string couldn't be decoded for node %v : %v\n",
				ens_info.FQDN,err,
			)
		}

	}
	c.HTML(http.StatusOK, "ens_info.html", gin.H{
		"ENSInfo" : ens_info,
	})
}
