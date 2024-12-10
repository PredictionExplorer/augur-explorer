package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

)
func cosmic_game_donations_cg_simple_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_simple_list(0, 10000)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_simple_list.html", gin.H{
		"CosmicGameDonations" : donations,
	})
}
func cosmic_game_donations_cg_simple_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_simple_by_round(round_num)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_simple_by_round.html", gin.H{
		"CosmicGameDonations" : donations,
		"RoundNum": round_num,
	})
}
func cosmic_game_donations_cg_with_info_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_with_info_simple_list(0, 10000)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_with_info_list.html", gin.H{
		"CosmicGameDonations" : donations,
	})
}
func cosmic_game_donations_cg_with_info_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_with_info_by_round(round_num)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_with_info_by_round.html", gin.H{
		"CosmicGameDonations" : donations,
		"RoundNum": round_num,
	})
}
func cosmic_game_donations_cg_with_info_record_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_record_id := c.Param("record_id")
	var record_id int64
	if len(p_record_id) > 0 {
		var success bool
		record_id,success = parse_int_from_remote_or_error(c,HTTP,&p_record_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'record_id' parameter is not set")
		return
	}
	record_info := arb_storagew.Get_donation_with_info_record_info(record_id)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_with_info_record_info.html", gin.H{
		"ETHDonation" : record_info,
		"RecordId": record_id,
	})
}
func cosmic_game_donations_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_by_user(user_aid)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_by_user.html", gin.H{
		"CombinedDonationRecords" : donations,
		"UserAddr": p_user_addr,
		"UserAid": user_aid,
	})
}
func cosmic_game_donations_cg_both_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_both_by_round(round_num)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_both_by_round.html", gin.H{
		"CosmicGameDonations" : donations,
		"RoundNum": round_num,
	})
}
func cosmic_game_donations_cg_both_all(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_both_all()
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_both_all.html", gin.H{
		"CosmicGameDonations" : donations,
	})
}
func cosmic_game_donations_erc20_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	donations := arb_storagew.Get_erc20_donations_by_round(round_num)
	c.HTML(http.StatusOK, "cg_donations_erc20_by_round.html", gin.H{
		"DonationsERC20" : donations,
		"RoundNum": round_num,
	})
}
func cosmic_game_donations_erc20_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	donations := arb_storagew.Get_erc20_donations_by_user(user_aid)
	c.HTML(http.StatusOK, "cg_donations_erc20_by_user.html", gin.H{
		"DonationsERC20ByUser" : donations,
		"UserAddr": p_user_addr,
		"UserAid": user_aid,
	})
}
