package main

import (
	"os"
	"encoding/hex"
	"sort"
	"time"


	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (

	inspected_events []InspectedEvent
)
func build_list_of_inspected_events() []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)

	inspected_events= make([]InspectedEvent,0,32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_newowner[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_name_registered1[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_name_registered2[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_name_registered3[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: 	hex.EncodeToString(evt_hash_invalidated[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_hash_registered[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_new_resolver[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_registry_transfer[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_text_changed[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_name_bought[:4]),
			ContractAid: 0,
		},
	)
	return inspected_events
}
func get_event_ids(from_tx_id,to_tx_id int64) []int64 {
	output := make([]int64 ,0,1024)
	for _,e := range inspected_events {
		event_list := storage.Get_tx_ids_from_evt_logs_by_signature(
			e.Signature,e.ContractAid,from_tx_id,to_tx_id,
		)
		output = append(output,event_list...)
	}
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	num_elts:=Remove_duplicates_int64(output)
	return output[0:num_elts]
}
func init_ens_processing() {

	build_list_of_inspected_events()

}
func process_ens_events(exit_chan chan bool) {

	var max_batch_size int64 = 1024*100
	for {
		status := storage.Get_ens_proc_status()
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		Info.Printf("scanning event range from %v to %v\n",status.LastEvtId,status.LastEvtId+max_batch_size)
		id_upper_limit := status.LastEvtId + max_batch_size
		last_evt_id,err := storage.Get_last_evtlog_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_evt_id {
			id_upper_limit = last_evt_id
		}
		events := get_event_ids(status.LastEvtId,id_upper_limit)
		for _,evt_id := range events {
			err := process_ens_event(evt_id)
			if err != nil {
				Error.Printf("Pausing event processing loop for 5 sec due to error")
				time.Sleep(5 * time.Second)
				break
			}
			status.LastEvtId=evt_id
			storage.Update_ens_proc_status(&status)
		}
		if len(events) == 0 {
			last_evt_log_id_on_chain,err := storage.Get_last_evtlog_id()
			if err != nil {
				Info.Printf("Error getting last event log id: %v\n",err)
				os.Exit(1)
			}
			if last_evt_log_id_on_chain > id_upper_limit {
				// only advance upper range if events within the range have filled id value space
				status.LastEvtId = id_upper_limit
				storage.Update_ens_proc_status(&status)
			}
			time.Sleep(1 * time.Second) // sleep only if there is no data
		}
	}
}
func process_ens_event(evt_id int64) error {

	evtlog := storage.Get_event_log(evt_id)
	var log types.Log
	rlp.DecodeBytes(evtlog.RlpLog,&log)
	return nil
}
