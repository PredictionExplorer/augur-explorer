/// Javascript building code for charts in UI
package main
import (
	"fmt"
	"html/template"
	"strings"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func build_js_data_obj(mdepth *MarketDepth) (template.JS,template.JS) {
	var asks_str string = "["
	var bids_str string = "["

	var last_price float64 = 0.0
	for i:=0 ; i < len(mdepth.Asks) ; i++ {
		if len(asks_str) > 1 {
			asks_str = asks_str + ","
		}
		var entry string
		entry = mkt_depth_entry_to_js_obj(&mdepth.Asks[i])
		asks_str= asks_str + entry
		last_price = mdepth.Asks[i].Price
	}
/*	// Possibly replace this with a line indicating the spread, as another Serie
	if len(mdepth.Asks) > 0 {
		if len(mdepth.Bids) > 0 {
			// add fake BID entry to fill the hole for the spread
			last_elt:=len(mdepth.Asks)-1
			fake_entry := mdepth.Asks[last_elt]
			fake_entry.Price = mdepth.Bids[0].Price*10
			bids_str = "[" + mkt_depth_entry_to_js_obj(&fake_entry)
		}
	}
*/
	for i:=0 ; i < len(mdepth.Bids) ; i++ {
		if len(bids_str) > 1 {
			bids_str = bids_str + ","
		}
		var entry string
		entry = mkt_depth_entry_to_js_obj(&mdepth.Bids[i])
		bids_str= bids_str + entry
		last_price = mdepth.Bids[i].Price
	}

	asks_str = asks_str + "]"
	bids_str = bids_str + "]"
	_ = last_price
	return template.JS(bids_str),template.JS(asks_str)
}
func build_js_price_history(orders *[]OrderInfo) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*orders) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*orders)[i];
		var entry string
		entry = "{" +
//				"x:" + fmt.Sprintf("\"%v\"",e.Date)  + "," +
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.Price) + "," +
				"price: " + fmt.Sprintf("%v",e.Price) + "," +
				"volume: " + fmt.Sprintf("%v",e.Amount) + "," +
				"click: function() {load_order_data(\"" +
					e.CreatorAddr +"\",\"" +
					e.FillerAddr+ "\"," +
					fmt.Sprintf("%v,%v,%v,\"%v\"",e.MktAid,e.Price,e.Amount,e.Date) +
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	fmt.Printf("JS price history string: %v\n",data_str)
	return template.JS(data_str)
}
func build_js_profit_loss_history(entries *[]PLEntry) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		outcome_escaped := strings.ReplaceAll(e.OutcomeStr,"\"","\\\"")
		descr_escaped := strings.ReplaceAll(e.MktDescr,"\"","\\\"")
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.AccumPl) + "," +
				"pl: " + fmt.Sprintf("%v",e.ImmediateProfit) + "," +
				"pl_accum: " + fmt.Sprintf("%v",e.AccumPl) + "," +
				"date: \"" + fmt.Sprintf("%v",e.Date) + "\"," +
				"click: function() {load_pl_data(" +
					fmt.Sprintf("%v,%v,%v,%v,%v,%v,\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",%v",
							e.Id,e.ClaimStatus,e.NetPosition,e.AvgPrice,e.ImmediateProfit,e.AccumPl,e.MktAddr,e.MktAddrSh,outcome_escaped,
							descr_escaped,e.Date,e.CounterPAddr,e.CounterPAddrSh,e.OrderHash,e.BlockNum) +
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_open_positions(entries *[]PLEntry) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		outcome_escaped := strings.ReplaceAll(e.OutcomeStr,"\"","\\\"")
		descr_escaped := strings.ReplaceAll(e.MktDescr,"\"","\\\"")
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.AccumFrozen) + "," +
				"frozen: " + fmt.Sprintf("%v",e.FrozenFunds) + "," +
				"frozen_accum: " + fmt.Sprintf("%v",e.AccumFrozen) + "," +
				"date: \"" + fmt.Sprintf("%v",e.Date) + "\"," +
				"click: function() {load_open_pos_data(" +
					fmt.Sprintf("%v,%v,%v,\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",%v",
							e.AvgPrice,e.FrozenFunds,e.NetPosition,e.MktAddr,e.MktAddrSh,outcome_escaped,
							descr_escaped,e.Date,e.CounterPAddr,e.CounterPAddrSh,e.OrderHash,e.BlockNum) +
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_cash_flow_data(entries *[]BlockCash) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		var entry string
		entry = "{" +
				//"x:" + fmt.Sprintf("%v",i)  + "," +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.Ts)  + "," +
				"y:"  + fmt.Sprintf("%.2f",e.AccumCashFlow) + "," +
				"block_num: " + fmt.Sprintf("%v",e.BlockNum) + "," +
				"cash: " + fmt.Sprintf("%v",e.AccumCashFlow) + "" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_uniq_addrs(entries *[]UniqueAddrEntry) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		var entry string
		entry = "{" +
//				"x:" + fmt.Sprintf("\"%v\"",e.Day)  + "," +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.Ts)  + "," +
				"y:"  + fmt.Sprintf("%v",e.NumAddrsAccum) + "," +
				"num_addrs: " + fmt.Sprintf("%v",e.NumAddrs) + "," +
				"num_addrs_accum: " + fmt.Sprintf("%v",e.NumAddrsAccum) + "," +
				"date_str: " + fmt.Sprintf("\"%v\"",e.Day) + "" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_global_gas_usage_data(entries *[]GasSpent,field int) template.JS {
	var data_str string = "["
	Info.Printf("dumping entries for field=%v\n",field)
	for i:=0 ; i < len(*entries) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*entries)[i];
		var datum string
		switch field {
			case 0: datum=e.Trading
			case 1: datum=e.Reporting
			case 2: datum=e.Markets
			case 3: datum=e.Total
			case 4: datum=e.EthTrading
			case 5: datum=e.EthReporting
			case 6: datum=e.EthMarkets
			case 7: datum=e.EthTotal
		}
		var entry string
		entry = "{" +
				//"x:" + fmt.Sprintf("%v",i)  + "," +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.Ts)  + "," +
				"y:"  + fmt.Sprintf("%v",datum) + "," +
				"day: " + fmt.Sprintf("new Date(%v * 1000)",e.Ts)+ "" +
				"}"
		data_str= data_str + entry
		Info.Printf("datum=%v\n",entry)
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_price_estimate_history(prices *[]PriceEstimate) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*prices) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*prices)[i];
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.TimeStamp)  + "," +
				"y:"  + fmt.Sprintf("%v",e.PriceEst) + "," +
				"price: " + fmt.Sprintf("%v",e.PriceEst) + "," +
				"spread: " + fmt.Sprintf("%v",e.Spread) + "," +
				"date:" + fmt.Sprintf("\"%v\"",e.Date) + "," +
				"click: function() {load_pest_data(" +
					fmt.Sprintf("new Date(%v * 1000)",e.TimeStamp)+"," +
					fmt.Sprintf("%v,%v,%v,%v,",e.PriceEst,e.Spread,e.MaxBid,e.MinAsk) +
					fmt.Sprintf("%v,%v,%v,",e.WeightedPriceEst,e.WMaxBid,e.WMinAsk) +
					fmt.Sprintf("%v",e.EvtCode) +
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_weighted_price_history(prices *[]PriceEstimate) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*prices) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*prices)[i];
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.TimeStamp)  + "," +
				"y:"  + fmt.Sprintf("%v",e.WeightedPriceEst) + "," +
				"price: " + fmt.Sprintf("%v",e.WeightedPriceEst) + "," +
				"spread: " + fmt.Sprintf("%v",e.Spread) + "," +
				"wmaxbid:" + fmt.Sprintf("\"%v\"",e.WMaxBid) + "," +
				"wminask:" + fmt.Sprintf("\"%v\"",e.WMinAsk) + "," +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
