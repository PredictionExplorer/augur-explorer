package main
import (
	"fmt"
	"time"
	"bytes"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
func market_uniswap_pairs(c *gin.Context) {

	market := c.Param("market")
	market_addr,valid:=is_address_valid(c,false,market)
	if !valid {
		return
	}
	minfo,err := augur_srv.db_augur.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	pairs := augur_srv.db_augur.Get_market_uniswap_pairs(minfo.MktAid)
	c.HTML(http.StatusOK, "market_upairs.html", gin.H{
			"Market" : minfo,
			"MarketUniswapPairs": pairs,
	})

}
func uniswap_swaps(c *gin.Context) {

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
	now_ts := time.Now().Unix()
	past_ts := now_ts - 100 * 3600 * 24
	pair_info,_:= augur_srv.db_augur.Get_uniswap_pair_info(aid)
	swaps := augur_srv.db_augur.Get_uniswap_swaps(aid,0,200)
	c.HTML(http.StatusOK, "uniswap_swaps.html", gin.H{
			"PairInfo" : pair_info,
			"PairSwaps" : swaps,
			"SampleFinTs" : now_ts,
			"SampleInitTs" : past_ts,
	})
}
func show_upair_swap_prices(c *gin.Context) {

	p_pair_aid := c.Param("pair_aid")
	var pair_aid int64
	if len(p_pair_aid) > 0 {
		var success bool
		pair_aid,success = parse_int_from_remote_or_error(c,false,&p_pair_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"Pair ID is not set")
		return
	}
	p_inverse := c.Param("inverse")
	var inverse int64
	if len(p_inverse) > 0 {
		var success bool
		inverse,success = parse_int_from_remote_or_error(c,false,&p_inverse)
		if !success {
			return
		}
	} else {
		respond_error(c,"'inverse' parameter is not set")
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

	bool_inverse := false
	if inverse > 0 {
		bool_inverse = true
	}
	pair_info,_:= augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	prices := augur_srv.db_augur.Get_uniswap_token_prices(pair_aid,bool_inverse,init_ts,fin_ts,interval_secs)
	js_prices := build_js_upair_swap_prices(&prices)
	c.HTML(http.StatusOK, "upair_prices.html", gin.H{
			"PairInfo" : pair_info,
			"Prices" : prices,
			"JSPriceData" :js_prices,
			"InitTimeStamp": init_ts,
			"FinTimeSTamp": fin_ts,
	})
}
func show_single_uniswap_swap(c *gin.Context) {

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

	swap,err := augur_srv.db_augur.Get_uniswap_swap_by_id(id)
	if err!=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Error getting swap with id=%v : %v",id,err),
		})
		return
	}

	c.HTML(http.StatusOK, "single_uniswap_swap.html", gin.H{
			"UniswapSwap" : swap,
			"Id": id,
	})
}
func uniswap_correct_for_difference_in_decimals(value *big.Float,decimals1,decimals2 int) {
	if decimals1 != decimals2 {
		var dec_diff int = 0
		if decimals1 < decimals2 {
			dec_diff = decimals2 - decimals1;
			divisor_str := fmt.Sprintf("1%0*d",dec_diff, 0)
			divisor_big := big.NewFloat(0.0)
			divisor_big.SetString(divisor_str)
			value.Quo(value,divisor_big)
		} else {
			dec_diff = decimals1 - decimals2;
			multiplier_str := fmt.Sprintf("1%0*d",dec_diff, 0)
			multiplier_big := big.NewFloat(0.0)
			multiplier_big.SetString(multiplier_str)
			value.Mul(value,multiplier_big)
		}
	}
}
func uniswap_calc_slippage(pair_addr_str string,token_str string,amount_str string) (*big.Float,*big.Int,error) {
	// note: we are receiving decimals as parameter because the fetch porcess to get decimals from the
	//		contract is more complicated than just calling Decimals() on the contract. The code to
	//		fetch all token info is at primitives/augur_utils.go 
	//		the decimals should be provided from the DB

	addr := common.HexToAddress(pair_addr_str)
	qtoken := common.HexToAddress(token_str)

	ctrct_pair,err := NewUniswapV2Pair(addr,rpcclient)
	if err != nil {
		return nil,nil,err
	}
	var copts = new(bind.CallOpts)
	reserves,err := ctrct_pair.GetReserves(copts)
	if err != nil {
		return nil,nil,err
	}
	token0,err := ctrct_pair.Token0(copts)
	if err != nil {
		return nil,nil,err
	}
	token1,err := ctrct_pair.Token1(copts)
	if err != nil {
		return nil,nil,err
	}
	_=token1
	var r1,r2 *big.Int
	if bytes.Equal(qtoken.Bytes(),token0.Bytes()) {
		r1=reserves.Reserve0
		r2=reserves.Reserve1
	} else {
		r1=reserves.Reserve1
		r2=reserves.Reserve0
	}
	_,_,router02_addr_str := augur_srv.db_augur.Get_uniswap_contracts()
	router02_addr := common.HexToAddress(router02_addr_str)
	ctrct_router,err := NewUniswapV2Router(router02_addr,rpcclient)
	amount := big.NewInt(0)
	amount.SetString(amount_str,10)
	token_amount_out,err := ctrct_router.GetAmountOut(copts,amount,r1,r2)

	// calculate spot price before swap
	spot_price_before := big.NewFloat(0.0)
	r1_float := big.NewFloat(0.0)
	r1_float.SetString(r1.String())
	r2_float := big.NewFloat(0.0)
	r2_float.SetString(r2.String())
	spot_price_before.Quo(r1_float,r2_float)

	r1big := big.NewInt(0)
	r2big := big.NewInt(0)
	r1big.Set(r1)
	r1big.Add(r1big,amount)
	r2big.Sub(r2,token_amount_out)
	spot_price_after := big.NewFloat(0.0)
	r1_float = big.NewFloat(0.0)
	r1_float.SetString(r1.String())
	amount_float := big.NewFloat(0.0)
	amount_float.SetString(amount.String())
	r1_float.Add(r1_float,amount_float)
	r2_float = big.NewFloat(0.0)
	r2_float.SetString(r2.String())
	token_out_float := big.NewFloat(0.0)
	r2_float.Sub(r2_float,token_out_float)
	spot_price_after.Quo(r1_float,r2_float)

	slippage_float:= big.NewFloat(0.0)
	slippage_float.Sub(spot_price_after,spot_price_before)
	return slippage_float,token_amount_out,nil
}
func produce_uniswap_slippages(pi *MarketUPair,amount_str string) ([]TokenSlippage,error) {

	output := make([]TokenSlippage,0,2)
	{
		var ts TokenSlippage
		ts.Token1Addr = pi.Token0Addr
		ts.Token2Addr = pi.Token1Addr
		ts.Token1Symbol = pi.Token0Symbol
		ts.Token2Symbol = pi.Token1Symbol
		ts.Decimals1 = pi.Token0Decimals
		ts.Decimals2 = pi.Token1Decimals
		ts.PoolAddr = pi.PairAddr
		ts.NumSwaps = pi.TotalSwaps
		in_float := big.NewFloat(0.0)
		in_float.SetString(amount_str)
		ts.AmountIn,_ = in_float.Float64()

		amount := fmt.Sprintf("%v%0*d",amount_str,ts.Decimals1,0)
		slippage,token_amount_out,err := uniswap_calc_slippage(pi.PairAddr,ts.Token1Addr,amount)
		if err != nil {
			return output,err
		}
		uniswap_correct_for_difference_in_decimals(slippage,ts.Decimals2,ts.Decimals1)
		ts.Slippage,_ = slippage.Float64()

		famount := big.NewFloat(0.0)
		famount.SetString(token_amount_out.String())
		divisor := fmt.Sprintf("%v%0*d",1,ts.Decimals2,0)
		fdivisor := big.NewFloat(0.0)
		fdivisor.SetString(divisor)
		famount.Quo(famount,fdivisor)
		ts.AmountOut,_ = famount.Float64()

		output = append(output,ts)
	}
	{
		var ts TokenSlippage
		ts.Token1Addr = pi.Token1Addr
		ts.Token2Addr = pi.Token0Addr
		ts.Token1Symbol = pi.Token1Symbol
		ts.Token2Symbol = pi.Token0Symbol
		ts.Decimals1 = pi.Token1Decimals
		ts.Decimals2 = pi.Token0Decimals
		ts.PoolAddr = pi.PairAddr
		ts.NumSwaps = pi.TotalSwaps
		in_float := big.NewFloat(0.0)
		in_float.SetString(amount_str)
		ts.AmountIn,_ = in_float.Float64()

		amount := fmt.Sprintf("%v%0*d",amount_str,ts.Decimals1,0)
		slippage,token_amount_out,err := uniswap_calc_slippage(pi.PairAddr,ts.Token1Addr,amount)
		if err != nil {
			return output,err
		}
		uniswap_correct_for_difference_in_decimals(slippage,ts.Decimals2,ts.Decimals1)
		ts.Slippage,_ = slippage.Float64()

		famount := big.NewFloat(0.0)
		famount.SetString(token_amount_out.String())
		divisor := fmt.Sprintf("%v%0*d",1,ts.Decimals2,0)
		fdivisor := big.NewFloat(0.0)
		fdivisor.SetString(divisor)
		famount.Quo(famount,fdivisor)
		ts.AmountOut,_ = famount.Float64()

		output = append(output,ts)
	}
	return output,nil
}
func show_uniswap_slippage(c *gin.Context) {

	p_pair:= c.Param("pair")
	pair_addr,valid:=is_address_valid(c,false,p_pair)
	if !valid {
		return
	}
	pair_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(pair_addr)
	if err != nil {
		respond_error(c,fmt.Sprintf("Address %v not found",))
		return
	}
	pair_info,err := augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	if err != nil {
		respond_error(c,err.Error())
		return
	}
	//slippages,err := produce_uniswap_slippages(&pair_info,amount_to_trade)
	slippages := augur_srv.db_augur.Get_uniswap_latest_slippages(pair_aid)
	//amount_to_trade := "100";
	var amount_to_trade float64 = 0.0
	if len(slippages) > 0 {
		amount_to_trade = slippages[0].AmountIn
	}
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "uniswap/slippages.html", gin.H{
		"PairInfo" : pair_info,
		"TokenSlippage" : slippages,
		"AmountToTrade" : amount_to_trade,
	})
}
func rt_show_uniswap_slippage(c *gin.Context) {

	p_pair:= c.Param("pair")
	pair_addr,valid:=is_address_valid(c,false,p_pair)
	if !valid {
		return
	}
	pair_aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(pair_addr)
	if err != nil {
		respond_error(c,fmt.Sprintf("Address %v not found",))
		return
	}
	pair_info,err := augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	if err != nil {
		respond_error(c,err.Error())
		return
	}
	amount_to_trade := "100";
	slippages,err := produce_uniswap_slippages(&pair_info,amount_to_trade)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "uniswap_slippages.html", gin.H{
		"PairInfo" : pair_info,
		"TokenSlippage" : slippages,
		"AmountToTrade" : amount_to_trade,
	})
}
func user_uniswap_swaps(c *gin.Context) {

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
	swaps,total_rows := augur_srv.db_augur.Get_user_uniswap_swaps(user_aid,0,200)
	c.HTML(http.StatusOK, "user_uniswap_swaps.html", gin.H{
		"UserInfo" : user_info,
		"UserSwaps" : swaps,
		"TotalRows" : total_rows,
	})
}
