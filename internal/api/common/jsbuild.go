// Package common - Javascript chart-data builders embedded in JSON API responses
package common

import (
	"fmt"
	"html/template"
	"time"

	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
)

// BuildJSRandomwalkWithdrawalChart renders the withdrawal-chart rows as the
// legacy inline-JavaScript array the frontend charts consume. Every field is
// numeric or a Go-formatted timestamp, so the unescaped template.JS cannot
// carry attacker-controlled markup.
func BuildJSRandomwalkWithdrawalChart(intervals *[]rwmodel.WithdrawalChartEntry) template.JS {
	dataStr := "["

	for i := range len(*intervals) {
		if len(dataStr) > 1 {
			dataStr = dataStr + ","
		}
		e := &(*intervals)[i]
		var entry string
		ts := time.Unix(e.TimeStamp, 0)
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
	return template.JS(dataStr) // #nosec G203 -- numeric-only content built above; no user-controlled strings
}

// BuildJSFloorPriceData renders the floor-price rows as the legacy inline
// JavaScript array; like the withdrawal chart, the content is numeric-only.
func BuildJSFloorPriceData(intervals *[]rwmodel.FloorPrice) template.JS {
	dataStr := "["

	for i := range len(*intervals) {
		if len(dataStr) > 1 {
			dataStr = dataStr + ","
		}
		e := &(*intervals)[i]
		var entry string
		ts := time.Unix(e.TimeStamp, 0)
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
	return template.JS(dataStr) // #nosec G203 -- numeric-only content built above; no user-controlled strings
}
