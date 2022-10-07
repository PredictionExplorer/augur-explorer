package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)
type Syncing_Data_Type struct {
	HeadSlot			int64		`json: "head_slot",omitempty`
	SyncDistance		int64		`json: "sync_distance",omitempty`
	IsSyncing			bool		`json: "is_syncing",omitempty`
	IsOptimistic		bool		`json: "is_optimistic,omitempty"`
}
type Syncing_Type struct {
	Data		Syncing_Data_Type	`json: "data"`
}
func get_head_slot(req_body []byte) int64 {

	var remote_obj *Syncing_Type
	remote_obj = new(Syncing_Type)
	err := json.Unmarshal(req_body,remote_obj)
	if err != nil {
		Error.Printf("Can't unmarshal JSON: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n",remote_obj)
	return remote_obj.Data.HeadSlot
}
func main_event_loop() {
	url := fmt.Sprintf("%v//eth/v1/node/syncing",API_URL)
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
}
