package main

import (
	"fmt"
	"time"
	"html/template"

	. "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"

)
type ChartJSData struct {
	Labels			template.JS
	DataSet1		template.JS
}
func build_js_fee_returns(report_type int64,records* []BalV2FeeReturns) ChartJSData {
	var data_str string = "["
	var label_str string = "["
 
	for i:=0 ; i < len(*records) ; i++ {
		if len(data_str) > 1 {
			data_str = data_str + ","
		}
		if len(label_str) > 1 {
			label_str = label_str + ","
		}
		var e = &(*records)[i];
		data_str = data_str + fmt.Sprintf("%v",e.FeeReturnsUSD)

		ts := time.Unix(e.TimeStamp,0)
		var date_formatted string
		switch report_type {
			case 0:	// hourly
				date_formatted = fmt.Sprintf("\"Hour %v\"",ts.Format("15"))
			case 1: // daily
				date_formatted = fmt.Sprintf("\"%v\"",ts.Format("Jan 2"))
			case 2: // weekly
				year,week_num := ts.ISOWeek()
				date_formatted = fmt.Sprintf("\"Week %v of %v\"",week_num,year)
		}
		label_str = label_str + date_formatted
	}
	data_str = data_str + "]"
	label_str = label_str + "]"
	var output ChartJSData
	output.Labels = template.JS(label_str)
	output.DataSet1 = template.JS(data_str)
	return output
}
