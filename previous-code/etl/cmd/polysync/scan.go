package main

import (
	"os"
	//"fmt"
	"strings"
	"time"
	"encoding/json"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (
	layoutUS = "January 2, 2006"
)
func process_single_market(rec *Pol_Market_API_Record_JSON_v2) {

	if storage.Market_exists(rec.MarketId) {
		return	// we alredy have this market
	}
	var complementary_rec Pol_Market_API_Record_Complementary
	t, _ := time.Parse(layoutUS,rec.EndDate)
	complementary_rec.EndDateTs = t.Unix()
	t,_ = time.Parse(layoutUS,rec.StartDate)
	complementary_rec.StartDateTs = t.Unix()
	t,_ = time.Parse(layoutUS,rec.LowerBoundDate)
	complementary_rec.LowerBoundDateTs = t.Unix()
	t,_ = time.Parse(layoutUS,rec.UpperBoundDate)
	complementary_rec.UpperBoundDateTs = t.Unix()
	t,_ = time.Parse("2006-01-02T15:04:05.000Z",rec.CreatedAtDate)
	complementary_rec.CreatedAtDateTs = t.Unix()
	t,_ = time.Parse("2006-01-02T15:04:05.000Z",rec.UpdatedAt)
	complementary_rec.UpdatedAtTs = t.Unix()
	t,_ = time.Parse("2006-01-02T15:04:05.000Z",rec.ClosedTimeDate)
	complementary_rec.ClosedTimeDateTs = t.Unix()
	if rec.MarketType == "normal" {
		complementary_rec.MarketTypeCode = 1
	} else {
		if rec.MarketType == "scalar" {
			complementary_rec.MarketTypeCode = 2
		} else {
			Error.Printf("Error: market type '%v' unknown\n",rec.MarketType)
			os.Exit(1)
		}
	}
	rec.LowerBound = strings.ReplaceAll(rec.LowerBound,",","")
	rec.UpperBound = strings.ReplaceAll(rec.UpperBound,",","")
	storage.Insert_polymarket_from_web_api(rec,&complementary_rec,0,0)
}
func scan_markets(req_body []byte) {

	//var records Pol_Market_API_Response_JSON
	records := make([]Pol_Market_API_Record_JSON_v2,0,1024)
	err := json.Unmarshal(req_body,&records)
	if err != nil {
		Error.Printf("Can't unmarshal JSON: %v\n",err)
		os.Exit(1)
	}
	for i:=0;i<len(records);i++ {
		//	fmt.Printf("record %v:\n%+vi\n\n",i+1,records[i])
		process_single_market(&records[i])
	}
}

