package main
import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)

func pool_swaps(c *gin.Context) {

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
	pool_info,_ := augur_srv.db_augur.Get_pool_info(aid)
	swaps := augur_srv.db_augur.Get_pool_swaps(aid,0,200)
	c.HTML(http.StatusOK, "pool_swaps.html", gin.H{
			"PoolInfo" : pool_info,
			"PoolSwaps" : swaps,
	})
}
func show_pool_swap_prices(c *gin.Context) {

	p_pool_aid := c.Param("pool_aid")
	var pool_aid int64
	if len(p_pool_aid) > 0 {
		var success bool
		pool_aid,success = parse_int_from_remote_or_error(c,false,&p_pool_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Pool ID is not set")
		return
	}
	p_token1_aid := c.Param("token1_aid")
	var token1_aid int64
	if len(p_token1_aid) > 0 {
		var success bool
		token1_aid,success = parse_int_from_remote_or_error(c,false,&p_token1_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Token1 ID is not set")
		return
	}
	p_token2_aid := c.Param("token2_aid")
	var token2_aid int64
	if len(p_token2_aid) > 0 {
		var success bool
		token2_aid,success = parse_int_from_remote_or_error(c,false,&p_token2_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Token2 ID is not set")
		return
	}
	var err error
	p_init_ts := c.Param("init_ts")
	var init_ts int
	if len(p_init_ts) > 0 {
		init_ts, err = strconv.Atoi(p_init_ts)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse init_ts",
			})
			return
		}
	}
	p_fin_ts := c.Param("fin_ts")
	var fin_ts int
	if len(p_fin_ts) > 0 {
		fin_ts, err = strconv.Atoi(p_fin_ts)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse fin_ts",
			})
			return
		}
	}
	if fin_ts == 0 {
		fin_ts = 2147483647
	}
	p_interval_secs := c.Param("interval_secs")
	var interval_secs int = 0
	if len(p_interval_secs) > 0 {
		interval_secs, err = strconv.Atoi(p_interval_secs)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": "Can't parse 'interval_secs' param",
			})
			return
		}
	}
	if interval_secs == 0 {
		interval_secs = 60*60
	}

	pool_info,_ := augur_srv.db_augur.Get_pool_info(pool_aid)
	token1_info,_ := augur_srv.db_augur.Get_bpool_token_info(pool_aid,token1_aid)
	token2_info,_ := augur_srv.db_augur.Get_bpool_token_info(pool_aid,token2_aid)
	prices := augur_srv.db_augur.Get_balancer_token_prices(pool_aid,token1_aid,token2_aid,init_ts,fin_ts,interval_secs)
	js_prices := build_js_bpool_swap_prices(&prices)
	c.HTML(http.StatusOK, "bswap_prices.html", gin.H{
			"PoolInfo" : pool_info,
			"Token1Info" : token1_info,
			"Token2Info" : token2_info,
			"Prices" : prices,
			"JSPriceData" :js_prices,
			"InitTimeStamp": init_ts,
			"FinTimeSTamp": fin_ts,
			"IntervalSecs": interval_secs,
	})
}
func show_single_balancer_swap(c *gin.Context) {

	p_id:= c.Param("id")
	var id int64
	if len(p_id) > 0 {
		var success bool
		id,success = parse_int_from_remote_or_error(c,false,&p_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'id' parameter is not set")
		return
	}

	swap,err := augur_srv.db_augur.Get_balancer_swap_by_id(id)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Error getting swap with id=%v : %v",id,err),
		})
		return
	}

	c.HTML(http.StatusOK, "single_balancer_swap.html", gin.H{
			"BalancerSwap" : swap,
			"Id": id,
	})
}
func balancer_calc_slippage(addr_str string,token_in_str string,token_out_str string,amount_str string) (*big.Int,*big.Int,error) {

	addr := common.HexToAddress(addr_str)
	token_in := common.HexToAddress(token_in_str)
	token_out := common.HexToAddress(token_out_str)
	ctrct_bpool,err := NewBPool(addr,eclient)
	if err != nil {
		return nil,nil,err
	}
	var copts = new(bind.CallOpts)
	ten := big.NewInt(10)
	max_price := big.NewInt(0)


	token_in_balance,err := ctrct_bpool.GetBalance(copts,token_in)
	if err != nil {
		return nil,nil,err
	}
	token_out_balance,err := ctrct_bpool.GetBalance(copts,token_out)
	if err != nil {
		return nil,nil,err
	}
	token_in_weight,err := ctrct_bpool.GetDenormalizedWeight(copts,token_in)
	if err != nil {
		return nil,nil,err
	}
	token_out_weight,err := ctrct_bpool.GetDenormalizedWeight(copts,token_out)
	if err != nil {
		return nil,nil,err
	}
	swap_fee,err := ctrct_bpool.GetSwapFee(copts)
	if err != nil {
		return nil,nil,err
	}
	spot_price,err := ctrct_bpool.CalcSpotPrice(copts,token_in_balance,token_in_weight,token_out_balance,token_out_weight,swap_fee)
	max_price.Mul(spot_price,ten)

	amount := big.NewInt(0)
	amount.SetString(amount_str,10)
	token_amount_out,err := ctrct_bpool.CalcOutGivenIn(copts,token_in_balance,token_in_weight,token_out_balance,token_out_weight,amount,swap_fee)
	if err != nil {
		return nil,nil,err
	}
	new_in_balance := big.NewInt(0)
	new_in_balance.Set(token_in_balance)
	new_in_balance.Add(new_in_balance,amount)
	new_out_balance := big.NewInt(0)
	new_out_balance.Set(token_out_balance)
	new_out_balance.Add(new_out_balance,token_amount_out)
	spot_price_after,err := ctrct_bpool.CalcSpotPrice(copts,new_in_balance,token_in_weight,new_out_balance,token_out_weight,swap_fee)
	if err != nil {
		return nil,nil,err
	}
	slippage := big.NewInt(0)
	slippage.Sub(spot_price,spot_price_after)
	return slippage,token_amount_out,nil
}
func produce_pool_slippages(amount_to_trade string,pool_aid int64) []TokenSlippage {

	tokens := augur_srv.db_augur.Get_balancer_pool_tokens_for_slippage(pool_aid)
	for i:=0; i < len(tokens) ; i++ {
		t := &tokens[i]
		amount := fmt.Sprintf("%v%0*d",amount_to_trade,t.Decimals1, 0)
		slippage,amount_token_out,_:= balancer_calc_slippage(
			t.PoolAddr,
			t.Token1Addr,
			t.Token2Addr,
			amount,
		)
		if slippage != nil {
			fslippage := big.NewFloat(0.0)
			fslippage.SetString(slippage.String())
			divisor1_str := fmt.Sprintf("1%0*d", t.Decimals1, 0)
			divisor2_str := fmt.Sprintf("1%0*d", t.Decimals2, 0)
			divisor1 := big.NewFloat(0.0)
			divisor1.SetString(divisor1_str)
			divisor2 := big.NewFloat(0.0)
			divisor2.SetString(divisor2_str)
			quo := big.NewFloat(0.0)
			quo.Quo(fslippage,divisor1)
			resulting_slippage,_ := quo.Float64()
			t.Slippage = resulting_slippage
			famount := big.NewFloat(0.0)
			famount.SetString(amount)
			famount.Quo(famount,divisor1)
			t.AmountIn,_ = famount.Float64()
			famount.SetString(amount_token_out.String())
			famount.Quo(famount,divisor2)
			t.AmountOut,_ = famount.Float64()
		}
	}
	return tokens
}
func show_pool_slippage(c *gin.Context) {

	p_pool:= c.Param("pool")
	pool_addr,valid:=is_address_valid(c,false,p_pool)
	if !valid {
		return
	}
	pool_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(pool_addr)
	if err != nil {
		respond_error(c,fmt.Sprintf("Address %v not found",))
		return
	}
	pool_info,_ := augur_srv.db_augur.Get_pool_info(pool_aid)
	//amount_to_trade := "100";
	//tokens := produce_pool_slippages(amount_to_trade,pool_aid)
	tokens := augur_srv.db_augur.Get_balancer_latest_slippages(pool_aid)
	var amount_to_trade float64 = 0.0
	if len(tokens) > 0 {
		amount_to_trade = tokens[0].AmountIn
	}
	c.HTML(http.StatusOK, "pool_slippage.html", gin.H{
		"PoolInfo" : pool_info,
		"TokenSlippage" : tokens,
		"AmountToTrade" : amount_to_trade,
	})
}
func user_balancer_swaps(c *gin.Context) {

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
	swaps,total_rows := augur_srv.db_augur.Get_user_balancer_swaps(user_aid,0,200)
	c.HTML(http.StatusOK, "user_balancer_swaps.html", gin.H{
		"UserInfo" : user_info,
		"UserSwaps" : swaps,
		"TotalRows" : total_rows,
	})
}
func arbitrum_augur_pools(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return 
	}
	pools := augur_srv.db_matic.Get_arbitrum_augur_pools()
	c.HTML(http.StatusOK, "arbitrum_augur_pools.html", gin.H{
		"ArbitrumAugurPools" : pools,
	})
}
