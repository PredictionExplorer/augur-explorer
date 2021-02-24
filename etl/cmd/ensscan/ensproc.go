package main

import (
	"os"
	"encoding/hex"
	"sort"
	"time"
	"context"
	"math/big"
	"bytes"
	"fmt"
	"strings"

	"github.com/wealdtech/go-ens/v3"
	"github.com/wealdtech/go-ens/v3/contracts/resolver"
	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/common"
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
func get_event_ids(from_evt_id,to_evt_id int64) []int64 {
	output := make([]int64 ,0,1024)
	for _,e := range inspected_events {
		event_list := storage.Get_evtlogs_by_signature_only_in_range(
			e.Signature,from_evt_id,to_evt_id,
		)
		/*Info.Printf("selected events for signature %v:\n",e.Signature)
		for _,evt_id := range event_list {
			Info.Printf("\tEvtId:\t%9v\n",evt_id)
		}*/
		output = append(output,event_list...)
	}
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	num_elts:=Remove_duplicates_int64(output)
	return output[0:num_elts]
}
func proc_name_registered1(log *types.Log,evt_id int64,tx_id int64) {
	var evt ENS_Name1
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TxHash = log.TxHash.String()
	var eth_event NameRegistered_v1
	err := ens_abi.Unpack(&eth_event,"NameRegistered1",log.Data)
	if err != nil {
		Error.Printf("Error upacking NameRegistered1: %v\n",err)
		Info.Printf("Error upacking NameRegistered1: %v\n",err)
		return
	}
	ctx := context.Background()
	bnum := big.NewInt(int64(log.BlockNumber))
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}
	evt.TimeStamp = int64(block_hdr.Time)
	eth_event.Label = log.Topics[1]
	eth_event.Owner = common.BytesToAddress(log.Topics[2][12:])
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	eth_event.Dump(Info)
	evt.TxHash = log.TxHash.String()
	evt.Label = hex.EncodeToString(eth_event.Label[:])
	owner_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.Owner = owner_addr.String()
	evt.Name = eth_event.Name
	evt.Cost = eth_event.Cost.String()
	evt.Expires = eth_event.Expires.Int64()
	evt.Contract = log.Address.String()
//		Info.Printf("label = %v, name=%v\n",hex.EncodeToString(eth_event.Label[:]),eth_event.Name)
//		Info.Printf("log data hex=%v\n",hex.EncodeToString(log.Data[:]))
//		Info.Printf("NameRegistered1: label=%v, Owner=%v, cost=%v, block %v tx %v\n",evt.Label,evt.Owner,log.BlockNumber,eth_event.Cost.String(),log.TxHash.String())
	storage.Insert_name_registered1(&evt)
}
func proc_name_registered2(log *types.Log,evt_id,tx_id int64) {
	var evt ENS_Name2
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	ctx := context.Background()
	bnum := big.NewInt(int64(log.BlockNumber))
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}
	evt.TimeStamp = int64(block_hdr.Time)
	evt.NameId = hex.EncodeToString(log.Topics[1][:])
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TxHash = log.TxHash.String()
	owner_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.Owner = owner_addr.String()
	expires := big.NewInt(0)
	expires.SetBytes(log.Data[:])
	evt.Expires = expires.Int64()
	evt.Contract = log.Address.String()
	Info.Printf("ENS_Name2 {\n")
	Info.Printf("\tOwner: %v\n",evt.Owner)
	Info.Printf("\tNameId: %v\n",evt.NameId)
	Info.Printf("\tExpires: %v\n",evt.Expires)
	Info.Printf("}")
//		Info.Printf("label = %v, name=%v\n",hex.EncodeToString(eth_event.Label[:]),eth_event.Name)
//		Info.Printf("log data hex=%v\n",hex.EncodeToString(log.Data[:]))
//		Info.Printf("NameRegistered1: label=%v, Owner=%v, cost=%v, block %v tx %v\n",evt.Label,evt.Owner,log.BlockNumber,eth_event.Cost.String(),log.TxHash.String())
//		storage.Insert_name_registered1(&evt)
}
func proc_newowner(log *types.Log,evt_id,tx_id int64) {

///		Info.Printf("%v: log = %+v\n",i,log)
	var evt ENS_NewOwner
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	bnum := big.NewInt(int64(log.BlockNumber))
	ctx := context.Background()
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}
	evt.TimeStamp = int64(block_hdr.Time)
	evt.Label = hex.EncodeToString(log.Topics[2][:])
	evt.Node = hex.EncodeToString(log.Topics[1][:])
	evt.Owner = common.BytesToAddress(log.Data[12:]).String()
	var new_node_hash [32]byte
	data :=make([]byte,32,64)
	copy(data,log.Topics[1][:]) // copying Node (bytes)
	data = append(data[:],log.Topics[2].Bytes()...)
	sha := sha3.NewLegacyKeccak256()
	if _, err := sha.Write(data[:]); err != nil {
		Error.Printf("cant calculate hash of new node: %v\n",err)
		os.Exit(1)
	}
	sha.Sum(new_node_hash[:0])
	evt.FQDN = hex.EncodeToString(new_node_hash[:])
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	Info.Printf("NewOwner %v for %v (node: %v, label: %v)\n",evt.Owner,evt.FQDN,evt.Node,evt.Label)
	evt.TxHash = log.TxHash.String()
	evt.Contract = log.Address.String()
	storage.Insert_new_owner(&evt)

}
func proc_name_registered3(log *types.Log,evt_id,tx_id int64) {

	var evt ENS_Name3
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	ctx := context.Background()
	bnum := big.NewInt(int64(log.BlockNumber))
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}

	var eth_event NameRegistered_v3
	err = ens_abi.Unpack(&eth_event,"NameRegistered3",log.Data)
	if err != nil {
		Error.Printf("Error upacking NameRegistered3: %v\n",err)
		Info.Printf("Error upacking NameRegistered3: %v\n",err)
		return
	}
	evt.TimeStamp = int64(block_hdr.Time)
	eth_event.Caller= common.BytesToAddress(log.Topics[1][12:])
	eth_event.Beneficiary = common.BytesToAddress(log.Topics[2][12:])
	eth_event.Label = log.Topics[3]
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TxHash = log.TxHash.String()
	evt.Label = hex.EncodeToString(eth_event.Label[:])
	evt.Subdomain = eth_event.Subdomain
	evt.CreatedDate = eth_event.CreatedDate.Int64()
	evt.Contract = log.Address.String()

	Info.Printf("ENS_Name3 {\n")
	Info.Printf("\tCaller: %v\n",eth_event.Caller.String())
	Info.Printf("\tBeneficiary: %v\n",eth_event.Beneficiary.String())
	Info.Printf("\tLabel: %v\n",eth_event.Label)
	Info.Printf("\tSubdomain: %v\n",eth_event.Subdomain)
	Info.Printf("\tCreatedDate: %v\n",evt.CreatedDate)
	Info.Printf("}")
}
func proc_hash_invalidated(log *types.Log,evt_id,tx_id int64) {

	var evt ENS_HashInvalidated
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	ctx := context.Background()
	bnum := big.NewInt(int64(log.BlockNumber))
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}

	var eth_event HashInvalidated
	err = ens_abi.Unpack(&eth_event,"HashInvalidated",log.Data)
	if err != nil {
		Error.Printf("Error upacking HashInvalidated: %v\n",err)
		Info.Printf("Error upacking HashInvalidated: %v\n",err)
		return
	}
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TimeStamp = int64(block_hdr.Time)
	copy(eth_event.Hash[:],log.Topics[1].Bytes())
	//eth_event.Name = Bytes32_to_string(log.Topics[2].Bytes())
	eth_event.Name = hex.EncodeToString(log.Topics[2][:])
	evt.TxHash = log.TxHash.String()
	evt.Hash = hex.EncodeToString(eth_event.Hash[:])
	evt.Name = eth_event.Name
	evt.RegistrationDate=eth_event.RegistrationDate.Int64()
	evt.Value = eth_event.Value.String()
	evt.Contract = log.Address.String()

	Info.Printf("HashInvalidated {\n")
	Info.Printf("\tHash: %v\n",hex.EncodeToString(eth_event.Hash[:]))
	Info.Printf("\tName: %v\n",eth_event.Name)
	Info.Printf("\tValue: %v\n",eth_event.Value.String())
	Info.Printf("\tRegDate: %v\n",eth_event.RegistrationDate.String())
	Info.Printf("}")
	storage.Insert_hash_invalidated(&evt)
}
func proc_new_resolver(log *types.Log,evt_id,tx_id int64) {
	var evt ENS_NewResolver
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	ctx := context.Background()
	bnum := big.NewInt(int64(log.BlockNumber))
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}

	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TimeStamp = int64(block_hdr.Time)
	evt.Node = hex.EncodeToString(log.Topics[1][:])
	addr := common.BytesToAddress(log.Data[12:])
	evt.Address = addr.String()
	evt.TxHash = log.TxHash.String()
	evt.Contract = log.Address.String()
	Info.Printf("NewResolver {\n")
	Info.Printf("\tNode: %v\n",evt.Node)
	Info.Printf("\tAddress: %v\n",evt.Address)
	Info.Printf("}")
	storage.Insert_new_resolver(&evt)
}
func proc_registry_transfer(log *types.Log,evt_id,tx_id int64) {

	var evt ENS_RegistryTransfer
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	ctx := context.Background()
	bnum := big.NewInt(int64(log.BlockNumber))
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}

	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TimeStamp = int64(block_hdr.Time)
	evt.Node = hex.EncodeToString(log.Topics[1][:])
	addr := common.BytesToAddress(log.Data[12:])
	evt.Owner = addr.String()
	evt.TxHash = log.TxHash.String()
	evt.Contract = log.Address.String()
	Info.Printf("(Registry) Transfer {\n")
	Info.Printf("\tNode: %v\n",evt.Node)
	Info.Printf("\tAddress: %v\n",evt.Owner)
	Info.Printf("}")
	storage.Insert_registry_transfer(&evt)
}
func proc_text_changed(log *types.Log,evt_id,tx_id int64) {

	var evt ENS_TextChanged
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	ctx := context.Background()
	bnum := big.NewInt(int64(log.BlockNumber))
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}

	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TimeStamp = int64(block_hdr.Time)
	evt.Node = hex.EncodeToString(log.Topics[1][:])
	if len(log.Data) < 64 {
		Error.Printf("Got event with log.Data of length lower than 64 bytes: %v bytes, skipping\n",len(log.Data))
		Info.Printf("Got event with log.Data of length lower than 64 bytes: %v bytes, skipping\n",len(log.Data))
		return
	}
	key_data := log.Data[64:]
	length := bytes.Index(key_data,[]byte{0})
	if length == -1 {
		length = 0
	}
	evt.Key = string(key_data[:length])
	evt.TxHash = log.TxHash.String()
	evt.Contract = log.Address.String()

	// the event doesn't provide the value, so we have to make a call to the contract
	var node [32]byte
	copy(node[:],log.Topics[1][:])
	registry,err := ens.NewRegistry(eclient)
	if err != nil {
		err_str := fmt.Sprintf(
			"Can't instantiate Registry contract for node %v (tx %v): %v",
			evt.Node,evt.TxHash,err,
		)
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	resolver_addr,err := registry.Contract.Resolver(nil,node)
	if err != nil {
		err_str := fmt.Sprintf(
			"Can't fetch Resolver contract addr for node %v (tx %v): %v",
			evt.Node,evt.TxHash,err,
		)
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	resolver_ctrct,err := resolver.NewContract(resolver_addr, eclient)
	if err != nil {
		err_str := fmt.Sprintf(
			"Can't instantiate Resolver contract for node %v (tx %v, ctrct addr %v): %v",
			evt.Node,evt.TxHash,resolver_addr.String(),err,
		)
		Info.Print(err_str)
		Error.Print(err_str)
		os.Exit(1)
	}
	text,err := resolver_ctrct.Text(nil,node,evt.Key)
	if err != nil {
		err_str := fmt.Sprintf(
			"Can't call Text() method for node %v (tx %v ctrct addrt %v): %v",
			evt.Node,evt.TxHash,resolver_addr.String(),err,
		)
		Info.Print(err_str)
		Error.Print(err_str)
		return
	}
	textbytes := []byte(evt.Value)
	Info.Printf("key bytes = %v\n",hex.EncodeToString([]byte(evt.Key)))
	Info.Printf("text bytes = %v\n",hex.EncodeToString(textbytes))
	evt.Value = hex.EncodeToString([]byte(strings.ReplaceAll(text,"\x00", "")))
	Info.Printf("text bytes after replace= %v\n",hex.EncodeToString(textbytes))
	Info.Printf("TextChanged {\n")
	Info.Printf("\tNode: %v\n",evt.Node)
	Info.Printf("\tKey: %v\n",evt.Key)
	Info.Printf("\tValue: %v\n",evt.Value)
	Info.Printf("}")
	storage.Insert_text_changed(&evt)
}
func proc_hash_registered(log *types.Log,evt_id,tx_id int64) {

	var evt ENS_HashRegistered
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	ctx := context.Background()
	bnum := big.NewInt(int64(log.BlockNumber))
	block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
	if err != nil {
		Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
		Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
	}

	var eth_event HashRegistered
	err = ens_abi.Unpack(&eth_event,"HashRegistered",log.Data)
	if err != nil {
		Error.Printf("Error upacking HashRegistered: %v\n",err)
		Info.Printf("Error upacking HashRegistered: %v\n",err)
		return
	}
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TimeStamp = int64(block_hdr.Time)
	copy(eth_event.Hash[:],log.Topics[1].Bytes())
	//eth_event.Name = Bytes32_to_string(log.Topics[2].Bytes())
	eth_event.Owner = common.BytesToAddress(log.Topics[2][12:])
	evt.Owner = eth_event.Owner.String()
	evt.TxHash = log.TxHash.String()
	evt.Hash = hex.EncodeToString(eth_event.Hash[:])
	evt.RegistrationDate=eth_event.RegistrationDate.Int64()
	evt.Value = eth_event.Value.String()
	evt.Contract = log.Address.String()

	Info.Printf("HashRegistered {\n")
	Info.Printf("\tHash: %v\n",hex.EncodeToString(eth_event.Hash[:]))
	Info.Printf("\tOwner: %v\n",evt.Owner)
	Info.Printf("\tValue: %v\n",eth_event.Value.String())
	Info.Printf("\tRegDate: %v\n",eth_event.RegistrationDate.String())
	Info.Printf("}")
	storage.Insert_hash_registered(&evt)
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
	log.BlockNumber=uint64(evtlog.BlockNum)
	log.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
	log.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
	num_topics := len(log.Topics)
	if num_topics > 0 {
		Info.Printf("found event with sig = %v\n",log.Topics[0].String())
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_newowner) {
			proc_newowner(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_name_registered1) {
			proc_name_registered1(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_name_registered2) {
			proc_name_registered2(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_name_registered3) {
			proc_name_registered3(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_hash_invalidated) {
			proc_hash_invalidated(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_hash_registered) {
			proc_hash_registered(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_new_resolver) {
			proc_new_resolver(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_registry_transfer) {
			proc_registry_transfer(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_text_changed) {
			proc_text_changed(&log,evtlog.EvtId,evtlog.TxId)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_name_bought) {
			//PENDING
			//proc_(&log)
		}

	}
	return nil
}
