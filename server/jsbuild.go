/// Javascript building code for charts in UI
package main
import (
	"fmt"
	"time"
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
				"x:" + fmt.Sprintf("%v",i)  + "," +
				"y:"  + fmt.Sprintf("%v",e.Price) + "," +
				"price: " + fmt.Sprintf("%v",e.Price) + "," +
				"volume: " + fmt.Sprintf("%v",e.Amount) + "," +
				"click: function() {load_order_data(\"" +
					e.CreatorAddr +"\",\"" +e.FillerAddr+ "\"," +
					fmt.Sprintf("%v,%v,%v,\"%v\"",e.MktAid,e.Price,e.Amount,e.Date) +
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
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
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.Ts)  + "," +
				"y:"  + fmt.Sprintf("%v",datum) + "," +
				"day: " + fmt.Sprintf("new Date(%v * 1000)",e.Ts)+ "" +
				"}"
		data_str= data_str + entry
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
func build_js_bpool_swap_prices(prices* []BSwapPrice) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*prices) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*prices)[i];
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.TimeStamp)  + "," +
				"y:"  + fmt.Sprintf("%v",e.Price) + "," +
				"price: " + fmt.Sprintf("%v",e.Price) + "," +
				"num_recs: " + fmt.Sprintf("%v",e.NumRecords) + "," +
				"date_str: " + fmt.Sprintf("\"%v\"",e.Date) + "," +
				"click: function() {load_price_data(\"" +
					e.Date+"\"," +fmt.Sprintf("%v",e.Price)+","+fmt.Sprintf("%v",e.NumRecords)+
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_upair_swap_prices(prices* []UPairPrice) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*prices) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*prices)[i];
		var entry string
		entry = "{" +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.TimeStamp)  + "," +
				"y:"  + fmt.Sprintf("%v",e.Price) + "," +
				"price: " + fmt.Sprintf("%v",e.Price) + "," +
				"num_recs: " + fmt.Sprintf("%v",e.NumRecords) + "," +
				"date_str: " + fmt.Sprintf("\"%v\"",e.Date) + "," +
				"click: function() {load_price_data(\"" +
					e.Date+"\"," +fmt.Sprintf("%v",e.Price)+ ","+fmt.Sprintf("%v",e.NumRecords)+
				")}" +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_ethusd_price_history(prices* []EthUsdPrice) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*prices) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*prices)[i];
		var entry string
		ts := time.Unix(int64(e.TimeStamp),0)
		date_str := fmt.Sprintf("%v",ts)
		entry = "{" +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.TimeStamp)  + "," +
				"y:"  + fmt.Sprintf("%v",e.Price) + "," +
				"price: " + fmt.Sprintf("%v",e.Price) + "," +
				"date_str: \"" + date_str + "\"," +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_noshow_bond_price_history(prices* []NoShowBondPrice) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*prices) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*prices)[i];
		var entry string
		ts := time.Unix(int64(e.TimeStamp),0)
		date_str := fmt.Sprintf("%v",ts)
		entry = "{" +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.TimeStamp)  + "," +
				"y:"  + fmt.Sprintf("%v",e.Price) + "," +
				"price: " + fmt.Sprintf("%v",e.Price) + "," +
				"date_str: \"" + date_str + "\"," +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
func build_js_validity_bond_price_history(prices* []ValidityBondPrice) template.JS {
	var data_str string = "["

	for i:=0 ; i < len(*prices) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		var e = &(*prices)[i];
		var entry string
		ts := time.Unix(int64(e.TimeStamp),0)
		date_str := fmt.Sprintf("%v",ts)
		entry = "{" +
				"x:" + fmt.Sprintf("new Date(%v * 1000)",e.TimeStamp)  + "," +
				"y:"  + fmt.Sprintf("%v",e.Price) + "," +
				"price: " + fmt.Sprintf("%v",e.Price) + "," +
				"date_str: \"" + date_str + "\"," +
				"}"
		data_str= data_str + entry
	}
	data_str = data_str + "]"
	return template.JS(data_str)
}
