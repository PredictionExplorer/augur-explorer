// Package common - Javascript building code for charts in UI
package common

import (
	"fmt"
	"html/template"
	"time"

	rwp "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/randomwalk"
)

func BuildJSRandomwalkVolumeHistory(prices *[]rwp.API_VolumeHistory) template.JS {
	var dataStr string = "["

	for i := 0; i < len(*prices); i++ {
		if len(dataStr) > 1 {
			dataStr = dataStr + ","
		}
		var e = &(*prices)[i]
		var entry string
		ts := time.Unix(int64(e.StartTs), 0)
		dateStr := fmt.Sprintf("%v", ts)
		entry = "{" +
			"x:" + fmt.Sprintf("new Date(%v * 1000)", e.StartTs) + "," +
			"y:" + fmt.Sprintf("%v", e.VolumeAccum) + "," +
			"volume: " + fmt.Sprintf("%.18f", e.VolumeAccum) + "," +
			"date_str: \"" + dateStr + "\"" +
			"}"
		dataStr = dataStr + entry
	}
	dataStr = dataStr + "]"
	return template.JS(dataStr)
}

func BuildJSRandomwalkMintIntervals(intervals *[]rwp.API_MintInterval) template.JS {
	var dataStr string = "["

	for i := 0; i < len(*intervals); i++ {
		if len(dataStr) > 1 {
			dataStr = dataStr + ","
		}
		var e = &(*intervals)[i]
		var entry string
		ts := time.Unix(int64(e.TimeStamp), 0)
		dateStr := fmt.Sprintf("%v", ts)
		entry = "{" +
			"x:" + fmt.Sprintf("%v", e.TokenId) + "," +
			"y:" + fmt.Sprintf("%v", e.Interval) + "," +
			"interval: " + fmt.Sprintf("%v", e.Interval) + "," +
			"tokenid: " + fmt.Sprintf("%v", e.TokenId) + "," +
			"date_str: \"" + dateStr + "\"," +
			"timestamp:" + fmt.Sprintf("%v", e.TimeStamp) + "" +
			"}"
		dataStr = dataStr + entry
	}
	dataStr = dataStr + "]"
	return template.JS(dataStr)
}

func BuildJSRandomwalkWithdrawalChart(intervals *[]rwp.API_WithdrawalChartEntry) template.JS {
	var dataStr string = "["

	for i := 0; i < len(*intervals); i++ {
		if len(dataStr) > 1 {
			dataStr = dataStr + ","
		}
		var e = &(*intervals)[i]
		var entry string
		ts := time.Unix(int64(e.TimeStamp), 0)
		dateStr := fmt.Sprintf("%v", ts)
		entry = "{" +
			"x:" + fmt.Sprintf("new Date(%v * 1000)", e.TimeStamp) + "," +
			"y:" + fmt.Sprintf("%.18f", e.WithdrawalAmount) + "," +
			"amount: " + fmt.Sprintf("%v", e.WithdrawalAmount) + "," +
			"date_str: \"" + dateStr + "\"," +
			"timestamp:" + fmt.Sprintf("%v", e.TimeStamp) + "" +
			"}"
		dataStr = dataStr + entry
	}
	dataStr = dataStr + "]"
	return template.JS(dataStr)
}

func BuildJSFloorPriceData(intervals *[]rwp.API_FloorPrice) template.JS {
	var dataStr string = "["

	for i := 0; i < len(*intervals); i++ {
		if len(dataStr) > 1 {
			dataStr = dataStr + ","
		}
		var e = &(*intervals)[i]
		var entry string
		ts := time.Unix(int64(e.TimeStamp), 0)
		dateStr := fmt.Sprintf("%v", ts)
		entry = "{" +
			"x:" + fmt.Sprintf("new Date(%v * 1000)", e.TimeStamp) + "," +
			"y:" + fmt.Sprintf("%.18f", e.Price) + "," +
			"price: " + fmt.Sprintf("%v", e.Price) + "," +
			"date_str: \"" + dateStr + "\"" +
			"}"
		dataStr = dataStr + entry
	}
	dataStr = dataStr + "]"
	return template.JS(dataStr)
}
