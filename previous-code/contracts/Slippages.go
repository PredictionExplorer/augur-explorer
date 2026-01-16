package contracts

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)

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
func uniswap_calc_slippage(eclient *ethclient.Client,router02_addr_str string,pair_addr_str string,token_str string,amount_str string) (*big.Float,*big.Int,error) {
	// note: we are receiving decimals as parameter because the fetch porcess to get decimals from the
	//		contract is more complicated than just calling Decimals() on the contract. The code to
	//		fetch all token info is at primitives/augur_utils.go 
	//		the decimals should be provided from the DB

	addr := common.HexToAddress(pair_addr_str)
	qtoken := common.HexToAddress(token_str)

	ctrct_pair,err := NewUniswapV2Pair(addr,eclient)
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
	/*token1,err := ctrct_pair.Token1(copts)
	if err != nil {
		return nil,nil,err
	}*/
	var r1,r2 *big.Int
	if bytes.Equal(qtoken.Bytes(),token0.Bytes()) {
		r1=reserves.Reserve0
		r2=reserves.Reserve1
	} else {
		r1=reserves.Reserve1
		r2=reserves.Reserve0
	}
	router02_addr := common.HexToAddress(router02_addr_str)
	ctrct_router,err := NewUniswapV2Router(router02_addr,eclient)
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
func Produce_uniswap_slippages(eclient *ethclient.Client,router02_addr_str string,pi *MarketUPair,amount_str string) ([]TokenSlippage,error) {

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
		slippage,token_amount_out,err := uniswap_calc_slippage(eclient,router02_addr_str,pi.PairAddr,ts.Token1Addr,amount)
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
		slippage,token_amount_out,err := uniswap_calc_slippage(eclient,router02_addr_str,pi.PairAddr,ts.Token1Addr,amount)
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
func balancer_calc_slippage(eclient *ethclient.Client, addr_str string,token_in_str string,token_out_str string,amount_str string) (*big.Int,*big.Int,error) {

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
	if err != nil {
		return nil,nil,err
	}
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
func Produce_pool_slippages(eclient *ethclient.Client,amount_to_trade string,tokens []TokenSlippage) {

	for i:=0; i < len(tokens) ; i++ {
		t := &tokens[i]
		amount := fmt.Sprintf("%v%0*d",amount_to_trade,t.Decimals1, 0)
		slippage,amount_token_out,_:= balancer_calc_slippage(
			eclient,
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
}
