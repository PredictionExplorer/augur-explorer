package main

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
)
type SyncingDataType struct {
	HeadSlot			int64		`json: "head_slot,omitempty"`
	SyncDistance		int64		`json: "sync_distance,omitempty"`
	IsSyncing			bool		`json: "is_syncing,omitempty"`
	IsOptimistic		bool		`json: "is_optimistic,omitempty"`
}
type GenericRequestType struct {
	Data		map[string]interface{} `json: "data"`
}

func get_head_slot(req_body []byte) int64 {

	var d GenericRequestType
	err := json.Unmarshal(req_body,&d)
	if err != nil {
		Error.Printf("Can't unmarshal JSON: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n",d)
	var head_slot_str string
	var ok bool
	head_slot_str,ok = d.Data["head_slot"].(string)
	if !ok {
		Error.Printf("Can't extract value of 'head_slot' field")
		os.Exit(1)
	}
	var head_slot int64
	head_slot,_ = strconv.ParseInt(head_slot_str,10,64)
	return head_slot
}
func main_event_loop() {
	url := fmt.Sprintf("%v/eth/v1/node/syncing",API_URL)
	resp,err := http.Get(url)
	if err != nil {
		Error.Printf("Couldn't successfuly GET at % : %v\n",url,err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Error.Printf("Error in ReadAll: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("response  %+v\n",string(body))
	head_slot := get_head_slot(body)
	fmt.Printf("head slot=%v\n",head_slot)
	for {
		url := fmt.Sprintf("%v/eth/v1/beacon/headers/%v",API_URL,head_slot)
		fmt.Printf("URL: %v\n",url)
		resp,err := http.Get(url)
		if err != nil {
			Error.Printf("Couldn't successfuly GET at % : %v\n",url,err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Error.Printf("Error in ReadAll: %v\n",err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n",string(body))
		var req GenericRequestType
		err = json.Unmarshal(body,&req)
		if err != nil {
			Error.Printf("Can't unmarshal JSON: %v\n",err)
			os.Exit(1)
		}
		head_slot = head_slot + int64(BLOCK_INCREMENT)
		time.Sleep(time.Duration(REQUEST_DELAY) * time.Second)
	}
}
