package main
import (
	"fmt"
	"time"
	"os"
	"strconv"
	"encoding/csv"
	"net/http"
	"github.com/gin-gonic/gin"

)
func rwalk_index_page(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	caddrs := augur_srv.db_arbitrum.Get_randomwalk_contract_addresses()
	top5tokens := augur_srv.db_arbitrum.Get_top5_traded_tokens()
	c.HTML(http.StatusOK, "rw_index.html", gin.H{
		"ContractAddresses":caddrs,
		"Top5Tokens":top5tokens,
	})
}
func rwalk_current_offers(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		rwalk_aid = -1
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		market_aid = -1
	}

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
	offers := augur_srv.db_arbitrum.Get_active_offers(rwalk_aid,market_aid,int(order_by))

	c.HTML(http.StatusOK, "rw_current_offers.html", gin.H{
		"Offers" : offers,
		"RWalkAid": rwalk_aid,
		"RWalkAddr" : p_rwalk_addr,
		"MarketAid": market_aid,
		"MarketAddr" : p_market_addr,
	})
}
func rwalk_floor_price(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		rwalk_aid = -1
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		market_aid = -1
	}

	_,floor_price,_,_,err := augur_srv.db_arbitrum.Get_floor_price(rwalk_aid,market_aid)
	var db_err string
	if err != nil { db_err = err.Error() }
	c.HTML(http.StatusOK, "rw_floor_price.html", gin.H{
		"FloorPrice" : floor_price,
		"DBError": db_err,
		"MarketAddr":p_market_addr,
		"RWalkAddr":p_rwalk_addr,
		"RWalkAid": rwalk_aid,
		"MarketAid": market_aid,
	})
}
func rwalk_token_list_seq(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"NTF address wasn't found in the 'address' table")
		return
	}
	tokens:= augur_srv.db_arbitrum.Get_minted_tokens_sequentially(rwalk_aid,0 ,10000000000)

	fin_ts := int(time.Now().Unix())
	interval := int(2 * 24 * 60* 60)
	init_ts := fin_ts - interval
	c.HTML(http.StatusOK, "rw_tokens_minted.html", gin.H{
		"MintedTokens" : tokens,
		"RWalkAddr" : p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
		"InitTs":init_ts,
		"FinTs":fin_ts,
		"Interval":interval,
	})
}
func rwalk_token_list_period(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"NTF address wasn't found in the 'address' table")
		return
	}
	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}
	tokens := augur_srv.db_arbitrum.Get_minted_tokens_by_period(rwalk_aid,ini,fin)

	c.HTML(http.StatusOK, "rw_tokens_minted_period.html", gin.H{
		"MintedTokens" : tokens,
		"RWalkAddr" : p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
		"InitTs": ini,
		"FinTs":fin,
		"InitDate" : time.Unix(int64(ini),0).String(),
		"FinDate":time.Unix(int64(fin),0).String(),
	})
}
func rwalk_trading_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	var market_aid int64 = 0
	if p_market_addr != "0x0000000000000000000000000000000000000000" {
		var err error
		market_aid,err = augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
		if err != nil {
			respond_error(c,"Market address doesn't exist in the database")
			return
		}
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	sales := augur_srv.db_arbitrum.Get_trading_history(market_aid,offset,limit)

	c.HTML(http.StatusOK, "rw_trading_history.html", gin.H{
		"Trading" : sales,
		"MarketPlaceAddr": p_market_addr,
		"MarketPlaceAid" : market_aid,
	})
}
func rwalk_sale_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	var market_aid int64 = 0
	if p_market_addr != "0x0000000000000000000000000000000000000000" {
		var err error
		market_aid,err = augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
		if err != nil {
			respond_error(c,"Market address doesn't exist in the database")
			return
		}
	}
	offset:=int(0)
	limit:= int(100000)
	sales := augur_srv.db_arbitrum.Get_sale_history(market_aid,offset,limit)

	c.HTML(http.StatusOK, "rw_sale_history.html", gin.H{
		"Trading" : sales,
		"MarketPlaceAddr": p_market_addr,
		"MarketPlaceAid" : market_aid,
	})
}
func rwalk_token_stats(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of NFT token address in the Db has failed")
		return
	}
	stats := augur_srv.db_arbitrum.Get_random_walk_stats(rwalk_aid)

	c.HTML(http.StatusOK, "rw_token_stats.html", gin.H{
		"TokenStats" : stats,
		"RWalkAid": rwalk_aid,
		"RWalkAddr" : p_rwalk_addr,
	})
}
func rwalk_market_stats(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		respond_error(c,"Market address doesn't exist in the database")
		return
	}
	stats := augur_srv.db_arbitrum.Get_market_stats(market_aid)

	c.HTML(http.StatusOK, "rw_market_stats.html", gin.H{
		"MarketStats" : stats,
		"MarketAid": market_aid,
		"MarketAddr" : p_market_addr,
	})
}
func rwalk_token_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,HTTP,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'token_id' parameter is not set")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of 'rwalk_addr' failed, address doesn't exist")
	}
	offset := int(0) ; limit:= int(100000)
	/*success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}*/
	history := augur_srv.db_arbitrum.Get_token_full_history(rwalk_aid,token_id,offset,limit)
	token_info,err := augur_srv.db_arbitrum.Get_rwalk_token_info(rwalk_aid,token_id)
	if err != nil {
		fmt.Printf("Error getting token info for token_id=%v, rwalk_aid=%v : %v\n",token_id,rwalk_aid,err)
	}

	c.HTML(http.StatusOK, "rw_token_history.html", gin.H{
		"TokenId" : token_id,
		"TokenHistory" : history,
		"RWalkAddr" : p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
		"TokenInfo" : token_info,
	})
}
func rwalk_trading_volume_by_period(c *gin.Context) {

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		respond_error(c,"Market address doesn't exist in the database")
		return
	}

	vol_hist := augur_srv.db_arbitrum.Get_market_trading_volume_by_period(market_aid,init_ts,fin_ts,interval_secs)
	volume_data := build_js_randomwalk_volume_history(&vol_hist)
	c.HTML(http.StatusOK, "rw_volume_history.html", gin.H{
		"VolumeHistory" : vol_hist,
		"VolumeData" : volume_data,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
		"MarketAddr" : p_market_addr,
		"MarketAid" : market_aid,
	})
}
func rwalk_token_name_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,HTTP,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'token_id' parameter is not set")
		return
	}
	name_changes := augur_srv.db_arbitrum.Get_name_changes_for_token(token_id)

	c.HTML(http.StatusOK, "rw_token_names.html", gin.H{
		"TokenNameChanges" : name_changes,
	})
}
func rwalk_tokens_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var err error
		user_aid, err = strconv.ParseInt(p_user_aid,10,64)
		if err != nil {
			if (len(p_user_aid) != 40) && (len(p_user_aid)!=42) {
				respond_error(c,"Can't resolve user identifier to valid address ID or address hex")
				return
			} else {
				user_aid,err = augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_user_aid)
			}
		}

	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}
	user_addr,err := augur_srv.db_arbitrum.Lookup_address(user_aid)
	if err != nil {
		respond_error(c,"Address lookup on user_aid failed")
		return
	}
	user_tokens := augur_srv.db_arbitrum.Get_random_walk_tokens_by_user(user_aid)

	c.HTML(http.StatusOK, "rw_tokens_by_user.html", gin.H{
		"UserTokens" : user_tokens,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
	})
}
func rwalk_trading_history_by_user(c *gin.Context) {

	if !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var err error
		user_aid, err = strconv.ParseInt(p_user_aid,10,64)
		if err != nil {
			if (len(p_user_aid) != 40) && (len(p_user_aid)!=42) {
				respond_error(c,"Can't resolve user identifier to valid address ID or address hex")
				return
			} else {
				user_aid,err = augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_user_aid)
				if err != nil {
					respond_error(c,"Cant find provided user")
					return
				}
			}
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}

	user_addr,err := augur_srv.db_arbitrum.Lookup_address(user_aid)
	if err != nil {
		respond_error(c,fmt.Sprintf("Address lookup on user_aid %v failed: %v",user_aid,err))
		return
	}
	user_trading := augur_srv.db_arbitrum.Get_trading_history_by_user(user_aid)

	c.HTML(http.StatusOK, "rw_trading_by_user.html", gin.H{
		"UserTrading" : user_trading,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
	})
}
func rwalk_user_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of NFT token failed")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,HTTP,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}
	user_addr,err := augur_srv.db_arbitrum.Lookup_address(user_aid)
	if err != nil {
		respond_error(c,"Address lookup on user_aid failed")
		return
	}
	user_info,dberr := augur_srv.db_arbitrum.Get_rwalk_user_info(user_aid,rwalk_aid)
	var dberr_string string
	if dberr != nil {
		dberr_string = dberr.Error()
	}
	c.HTML(http.StatusOK, "rw_user_info.html", gin.H{
		"UserInfo" : user_info,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
		"DBError" : dberr_string,
	})
}
func rwalk_top_users(c *gin.Context) {

	top_profit_makers := augur_srv.db_arbitrum.Get_randomwalk_top_profit_makers()
	top_trade_makers := augur_srv.db_arbitrum.Get_randomwalk_top_trade_makers()
	top_volume_makers := augur_srv.db_arbitrum.Get_randomwalk_top_volume_makers()
	c.HTML(http.StatusOK, "rw_top_users.html", gin.H{
			"title": "Top 100 Users of RandomWalk Token",
			"ProfitMakers" : top_profit_makers,
			"TradeMakers" : top_trade_makers,
			"VolumeMakers" : top_volume_makers,
	})
}
func rwalk_mint_intervals(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of NFT token failed")
		return
	}
	mint_intervals := augur_srv.db_arbitrum.Get_rwalk_mint_intervals(rwalk_aid)
	mint_data := build_js_randomwalk_mint_intervals(&mint_intervals)

	c.HTML(http.StatusOK, "rw_mint_intervals.html", gin.H{
			"MintIntervals" : mint_intervals,
			"MintIntervalData" : mint_data,
			"RWalkAid" : rwalk_aid,
			"RWalkAddr" : p_rwalk_addr,
	})
}
func rwalk_withdrawal_chart(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of NFT token failed")
		return
	}
	withdrawal_entries := augur_srv.db_arbitrum.Get_rwalk_withdrawal_chart(rwalk_aid)
	withdrawal_data := build_js_randomwalk_withdrawal_chart(&withdrawal_entries)
	rwalk_stats := augur_srv.db_arbitrum.Get_random_walk_stats(rwalk_aid)

	c.HTML(http.StatusOK, "rw_withdrawal_chart.html", gin.H{
			"WithdrawalEntries" : withdrawal_entries,
			"WithdrawalData" : withdrawal_data,
			"ContractStatistics": rwalk_stats,
			"RWalkAid" : rwalk_aid,
			"RWalkAddr" : p_rwalk_addr,
	})
}
func rwalk_floor_price_over_time(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of NFT token failed")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		respond_error(c,"Market address doesn't exist in the database")
		return
	}
	success,ini,fin,interval := parse_timeframe_params(c)
	if !success {
		return
	}
	if ini == 0 {
		ini = 1636676049
	}
	if fin == 0 {
		fin = int(time.Now().Unix()) 
	}
	if interval == 0 || interval == 2147483647 {
		interval = 24*60*60
	}
	price_entries := augur_srv.db_arbitrum.Get_rwalk_floor_price_for_periods(rwalk_aid,market_aid,ini,fin,interval)
	price_data := build_js_floor_price_data(&price_entries)
	rwalk_stats := augur_srv.db_arbitrum.Get_random_walk_stats(rwalk_aid)

	c.HTML(http.StatusOK, "rw_floor_price_over_time.html", gin.H{
			"PriceEntries" : price_entries,
			"PriceData" : price_data,
			"ContractStatistics": rwalk_stats,
			"RWalkAid" : rwalk_aid,
			"RWalkAddr" : p_rwalk_addr,
			"MarketAid":market_aid,
			"MarketAddr":p_market_addr,
			"InitTs" : ini,
			"FinTs": fin,
			"Interval": interval,
	})
}
func rwalk_token_csv_export(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	c.Writer.Header().Set("Cache-Control", "must-revalidate")
	c.Writer.Header().Set("Pragma","must-revalidate")
	c.Writer.Header().Set("Content-type","application/vnd.ms-excel")
	c.Writer.Header().Set("Content-disposition","attachment; filename=mints.csv")
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"NTF address wasn't found in the 'address' table")
		return
	}
	data := augur_srv.db_arbitrum.Get_minted_tokens_for_CSV(rwalk_aid)

	// convert the data to golang-way to store csv
	header:= []string {	// header
			"BlockNum",
			"TimeStamp",
			"DateTime",
			"ContractAddr",
			"TokenId",
			"MinterAddr",
			"SeedHex",
			"SeedNum",
			"PriceMinted",
			"TxHash",
			"NumTrades",
			"TotalVolume",
			"LastPrice",
			"TokenName",
			"CurOwner",
		}
	fname := "/tmp/mints.csv"
	f, err := os.Create(fname)
	if err != nil {
		respond_error(c,fmt.Sprintf("Cant create file: %v\n",err.Error()))
		return
	}
	w := csv.NewWriter(f)
	err = w.Write(header);
	if err != nil {
		respond_error(c,fmt.Sprintf("Error during header write to csv: %v\n",err.Error()))
		return
	}
	for i:=0; i<len(data); i++ {
		rec := &data[i]
		row := []string {
			fmt.Sprintf("%d",rec.BlockNum),
			fmt.Sprintf("%d",rec.TimeStamp),
			rec.DateTime,
			rec.ContractAddr,
			fmt.Sprintf("%v",rec.TokenId),
			rec.MinterAddr,
			rec.Seed,
			fmt.Sprintf("%s",rec.SeedNum),
			fmt.Sprintf("%f",rec.Price),
			rec.TxHash,
			fmt.Sprintf("%d",rec.NumTrades),
			fmt.Sprintf("%f",rec.TotalVolume),
			fmt.Sprintf("%f",rec.LastPrice),
			rec.LastName,
			rec.LastOwner,
		}
		err = w.Write(row);
		if err != nil {
			respond_error(c,fmt.Sprintf("Error during write to csv: %v\n",err.Error()))
			return
		}
	}
	w.Flush()
	f.Close()
	c.File(fname)
}
func rwalk_token_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of NFT token address in the Db has failed")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,HTTP,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'token_id' parameter is not set")
		return
	}
	token_info,err := augur_srv.db_arbitrum.Get_rwalk_token_info(rwalk_aid,token_id)
	if err != nil {
		respond_error(c,fmt.Sprintf("Error during query execution: %v",err))
		return
	}

	c.HTML(http.StatusOK, "rw_token_info.html", gin.H{
		"TokenInfo" : token_info,
		"RWalkAddr": p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
	})
}
