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
	head_slot,_ = strconv.ParseInt(head_slot_str,10,64)
	return head_slot
}
func get_randao_reveal(req_body []byte) (int64,int64,[]byte,string) {

	var ok bool
	var check_iface map[string]interface{} 
	err := json.Unmarshal(req_body,&check_iface)
	if err == nil {
		var code_f float64
		code_f,ok = check_iface["code"].(float64)
		code := int64(code_f)
		if ok {
			return code,0 ,nil,""	// error occured
		}
	}
	var d GenericRequestType
	err = json.Unmarshal(req_body,&d)
	if err != nil {
		Error.Printf("Can't unmarshal JSON: %v\n",err)
		os.Exit(1)
	}
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

	state_root,ok := message_iface["state_root"].(string)
	if !ok {
		Error.Printf("Can't extract value of 'proposer_index' field")
		os.Exit(1)
	}

	return 0, proposer_index,randao_reveal,state_root
}
func get_validator_info(req_body []byte) (int64,string,string) {
	// Returns request code,balance and publike key
	fmt.Printf("body=%+v\n",string(req_body))

	var ok bool
	var check_iface map[string]interface{} 
	err := json.Unmarshal(req_body,&check_iface)
	if err == nil {
		var code_f float64
		code_f,ok = check_iface["code"].(float64)
		code := int64(code_f)
		if ok {
			return code,"",""	// error occured
		}
	}
	
	var d GenericRequestType
	err = json.Unmarshal(req_body,&d)
	if err != nil {
		Error.Printf("Can't unmarshal JSON: %v\n",err)
		os.Exit(1)
	}
	var validator_iface map[string]interface{} 
	validator_iface,ok = d.Data["validator"].(map[string]interface{})
	var balance_str string
	balance_str,ok = d.Data["balance"].(string)
	if !ok {
		Error.Printf("Can't extract value of 'balance' field in get_validator_info()")
		os.Exit(1)
	}
	var pubkey_str string
	pubkey_str,ok = validator_iface["pubkey"].(string)
	if !ok {
		Error.Printf("Can't extract value of 'pubkey' field in get_validator_info()")
		os.Exit(1)
	}
	return 0 ,balance_str,pubkey_str
}
func main_event_loop() {
	fout,err := os.Create("output.csv")
	if err != nil {
		fmt.Printf("Error creating file: %v\n",err)
		os.Exit(1)
	}
	defer fout.Close()

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
		//time.Sleep(time.Duration(7 * time.Second))
		//fmt.Printf("%+v\n",string(body))
		var req GenericRequestType
		err = json.Unmarshal(body,&req)
		if err != nil {
			Error.Printf("Can't unmarshal JSON: %v\n",err)
			os.Exit(1)
		}
		head_slot = head_slot + int64(BLOCK_INCREMENT)
		req_code,proposer_index,randao_reveal,state_id := get_randao_reveal(body)
		if req_code !=0 {
			fmt.Printf("Block not found, skipping\n")
			continue	// block not found
		}
		fmt.Printf("Proposer %v\t%v\n",proposer_index,hex.EncodeToString(randao_reveal))
		url = fmt.Sprintf("%v/eth/v1/beacon/states/%v/validators/%v",API_URL,state_id,proposer_index)
		fmt.Printf("URL: %v\n",url)
	again:
		resp,err = http.Get(url)
		if err != nil {
			Error.Printf("Couldn't successfuly GET at % : %v\n",url,err)
			os.Exit(1)
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			Error.Printf("Error in ReadAll (for validator info): %v\n",err)
			os.Exit(1)
		}
		code,v_balance,v_pubkey := get_validator_info(body)
		if code != 0 {
			fmt.Printf("Error %v, sleeping for 1 sec\n",code)
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Printf("Retrying url %v\n",url)
			goto again
		}
		fmt.Printf("Proposer's balance: %v (pubkey %v)\n",v_balance,v_pubkey)
		fout.WriteString(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n",head_slot,proposer_index,hex.EncodeToString(randao_reveal),v_balance,v_pubkey))
		//time.Sleep(time.Duration(REQUEST_DELAY) * time.Second)
		//os.Exit(1)
	}
}
