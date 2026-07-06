// Package common - Javascript chart-data builders embedded in JSON API responses
package common

import (
	"fmt"
	"html/template"
	"time"

	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
)

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
