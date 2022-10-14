package main

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/hex"
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
	//fmt.Printf("%+v\n",d)
	var head_slot_str string
	var ok bool
	head_slot_str,ok = d.Data["head_slot"].(string)
	if !ok {
		Error.Printf("Can't extract value of 'head_slot' field")
		os.Exit(1)
	}
	var head_slot int64
	ead_slot,_ = strconv.ParseInt(head_slot_str,10,64)
	return head_slot
}
func get_randao_reveal(req_body []byte) (int64,[]byte) {

	var d GenericRequestType
	err := json.Unmarshal(req_body,&d)
	if err != nil {
		Error.Printf("Can't unmarshal JSON: %v\n",err)
		os.Exit(1)
	}
	var ok bool
	var message_iface map[string]interface{} 
	message_iface,ok = d.Data["message"].(map[string]interface{})
	var body map[string]interface{}
	body,ok = message_iface["body"].(map[string]interface{})
	if !ok {
		Error.Printf("Can't cast 'body' field to map\n")
		os.Exit(1)
	}
	var randao_reveal_str string
	randao_reveal_str,ok = body["randao_reveal"].(string)
	if !ok {
		Error.Printf("Can't extract value of 'randao_reveal' field")
		os.Exit(1)
	}
	randao_reveal,err := hex.DecodeString(randao_reveal_str[2:])
	if err != nil {
		fmt.Printf("Error at hex.DecodeString(): %v\n",err)
		os.Exit(1)
	}
	var proposer_index_str string
	proposer_index_str,ok = message_iface["proposer_index"].(string)
	if !ok {
		Error.Printf("Can't extract value of 'proposer_index' field")
		os.Exit(1)
	}
	proposer_index,_ := strconv.ParseInt(proposer_index_str,10,64)

	return proposer_index,randao_reveal
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
	//fmt.Printf("response  %+v\n",string(body))
	head_slot := get_head_slot(body)
	fmt.Printf("head slot=%v\n",head_slot)
	for {
		url := fmt.Sprintf("%v/eth/v2/beacon/blocks/%v",API_URL,head_slot)
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
		//fmt.Printf("%+v\n",string(body))
		var req GenericRequestType
		err = json.Unmarshal(body,&req)
		if err != nil {
			Error.Printf("Can't unmarshal JSON: %v\n",err)
			os.Exit(1)
		}
		head_slot = head_slot + int64(BLOCK_INCREMENT)
		proposer_index,randao_reveal := get_randao_reveal(body)
		fmt.Printf("Proposer %v\t%v\n",proposer_index,hex.EncodeToString(randao_reveal))
		time.Sleep(time.Duration(REQUEST_DELAY) * time.Second)
	}
}
