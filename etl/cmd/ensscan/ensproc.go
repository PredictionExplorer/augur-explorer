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
	"errors"
	"strings"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
//	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/wealdtech/go-ens/v3"
	"github.com/wealdtech/go-ens/v3/contracts/resolver"
	"github.com/wealdtech/go-ens/v3/contracts/baseregistrar"
	//"github.com/wealdtech/go-ens/v3/contracts/auctionregistrar"
	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
type std_proc_func func(log *types.Log,evt_id,tx_id,timestamp int64)
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
			Signature: hex.EncodeToString(evt_registrar_transfer[:4]),
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
func calculate_name_hash(node,label []byte) []byte {

	Info.Printf("calculate_name_hash(): node: %v\n",hex.EncodeToString(node[:]))
	Info.Printf("calculate_name_hash(): label: %v\n",hex.EncodeToString(label[:]))
	data :=make([]byte,32,64)
	copy(data,node[:]) // copying Node (bytes)
	Info.Printf("data: %v\n",hex.EncodeToString(data[:]))
	data = append(data[:],label[:]...)
	Info.Printf("data: %v\n",hex.EncodeToString(data[:]))
	keccak:= sha3.NewLegacyKeccak256()
	if _, err := keccak.Write(data[:]); err != nil {
		Error.Printf("cant calculate name hash: %v\n",err)
		os.Exit(1)
	}
	khash := keccak.Sum(nil)
	Info.Printf("hash: %v\n",hex.EncodeToString(khash[:]))
	return khash
}
func get_node_hash_via_new_owner_event(tx_id int64,tx_hash *common.Hash,value[]byte,is_label bool) ([]byte,[]byte,error) {
    // NameRegistered event doesn't provide the node hash in the event itself,
    // but there is a common pattern: when ENS name is registered a NewOwner event
    // is emitted and it contains Node hash. This function extracts node hash value
    // from linked NewOwner event
	possible_events,err := storage.Get_events_by_sig_and_tx_id(tx_id,"ce0457fe")
	if err != nil {
		Error.Printf("Error getting events by sig and tx_id: %v\n",err)
		os.Exit(1)
	}
	if len(possible_events) == 0 {
		// try from Receipts
		ctx := context.Background()
		receipt,err := eclient.TransactionReceipt(ctx,*tx_hash)
		if err != nil {
			Error.Printf("Receipt call failed for %v : %v)\n",tx_hash,err)
			os.Exit(1)
		}
		for i:=0; i<len(receipt.Logs); i++ {
			log := receipt.Logs[i]
			if !bytes.Equal(log.Topics[0].Bytes(),evt_newowner[:]) {
				continue
			}
			if is_label {
				if bytes.Equal(log.Topics[2][:],value[:]) && (len(log.Data)>0) {
					return log.Topics[1][:],value[:],nil
				}
			} else {
				// its the FQDN (aka 'id' or 'name hash')
				namehash := calculate_name_hash(log.Topics[1][:],log.Topics[2][:])
				Info.Printf(
					"namehash: %v , node: %v , label %v \n",
					hex.EncodeToString(namehash[:]),
					hex.EncodeToString(log.Topics[1][:]),
					hex.EncodeToString(log.Topics[2][:]),
				)
				if bytes.Equal(value,namehash[:]) {
					return log.Topics[1][:],log.Topics[2][:],nil
				}
			}
		}
		Error.Printf("No NewOwner events found after scanning transaction receipt logs, tx=%v\n",tx_hash.String())
		return nil,nil,errors.New("No NewOwner events found")
	}
	// this code decodes events from DB
	for i:=0; i<len(possible_events); i++ {
		evtlog := possible_events[i]
		var log types.Log
		err := rlp.DecodeBytes(evtlog.RlpLog,&log)
		if err != nil {
			Error.Printf("Error decoding RLP of event id=%v: %v\n",evtlog.EvtId)
			os.Exit(1)
		}
		if len(log.Data)<32 {
			Error.Printf("Found NewOwner event but the log.Data size is less than 32: %v\n",len(log.Data))
			os.Exit(1)
		}
		if !bytes.Equal(log.Topics[0].Bytes(),evt_newowner[:]) {
			continue
		}
		if is_label {
			if bytes.Equal(log.Topics[2][:],value[:]) {
				return log.Topics[1][:],value[:],nil
			}
		} else {
			// its the FQDN (aka 'id' or 'name hash')
			namehash := calculate_name_hash(log.Topics[1][:],log.Topics[2][:])
			if bytes.Equal(value,namehash[:]) {
				return log.Topics[1][:],log.Topics[2][:],nil
			}
		}

	}
	Error.Printf("Couldn't find NewOwner event linked with NameRegistered event, tx_id=%v\n",tx_id)
	return nil,nil,errors.New("No NewOwner events linked with NameRegistered events found")
}

func proc_name_registered1(log *types.Log,evt_id,tx_id,timestamp int64) {
	var evt ENS_Name1
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}
	evt.TxHash = log.TxHash.String()
	var eth_event NameRegistered_v1
	err := ens_abi.Unpack(&eth_event,"NameRegistered1",log.Data)
	if err != nil {
		Error.Printf("Error upacking NameRegistered1: %v\n",err)
		Info.Printf("Error upacking NameRegistered1: %v\n",err)
		os.Exit(1)
	}
	owner_addr := common.BytesToAddress(log.Topics[2][12:])

	eth_event.Label = log.Topics[1]
	eth_event.Owner = owner_addr
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	eth_event.Dump(Info)
	evt.TxHash = log.TxHash.String()
	evt.Label = hex.EncodeToString(eth_event.Label[:])

	node_hash,_,err := get_node_hash_via_new_owner_event(evt.TxId,&log.TxHash,eth_event.Label[:],true)
	if err != nil {
		Error.Printf("Error getting node hash: %v\n",err)
		os.Exit(1)
	}

	fqdn_bytes := calculate_name_hash(node_hash,eth_event.Label[:])
	evt.FQDN = hex.EncodeToString(fqdn_bytes[:])
	Info.Printf("resulting fqdn: %v\n",hex.EncodeToString(fqdn_bytes[:]))

	evt.Node = hex.EncodeToString(node_hash[:])
	evt.Owner = owner_addr.String()
	evt.Name = eth_event.Name
	evt.Cost = eth_event.Cost.String()
	evt.Expires = eth_event.Expires.Int64()
	evt.Contract = log.Address.String()
	storage.Insert_name_registered1(&evt)
}
func proc_name_registered2(log *types.Log,evt_id,tx_id,timestamp int64) {
	var evt ENS_Name2
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp

	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TxHash = log.TxHash.String()
	owner_addr := common.BytesToAddress(log.Topics[2][12:])
	evt.Owner = owner_addr.String()
	expires := big.NewInt(0)
	expires.SetBytes(log.Data[:])
	evt.Expires = expires.Int64()
	evt.Contract = log.Address.String()

	base_registrar_ctrct,err := baseregistrar.NewContract(log.Address,eclient)
	if err != nil {
		Error.Printf("Error instantiating base registrar contract %v : %v\n",log.Address.String(),err)
		Info.Printf("Error instantiating base registrar contract %v : %v\n",log.Address.String(),err)
		os.Exit(1)
	}
	var node_hash []byte
	label_hash := log.Topics[1][:]
	var node_bytes[32]byte
	var copts = new(bind.CallOpts)
	node_bytes,err = base_registrar_ctrct.BaseNode(copts)
	if err == nil {
		node_hash = node_bytes[:]
	} else {
		node_hash,_,err = get_node_hash_via_new_owner_event(evt.TxId,&log.TxHash,label_hash,true)
		if err != nil {
			Error.Printf("Error getting node hash: %v\n",err)
			os.Exit(1)
		}
	}
	fqdn_bytes := calculate_name_hash(node_hash,label_hash)
	evt.FQDN = hex.EncodeToString(fqdn_bytes)
	evt.Node = hex.EncodeToString(node_hash[:])
	evt.Label = hex.EncodeToString(label_hash[:])

	Info.Printf("ENS_Name2 {\n")
	Info.Printf("\tOwner: %v\n",evt.Owner)
	Info.Printf("\tNameId: %v\n",evt.Label)
	Info.Printf("\tExpires: %v\n",evt.Expires)
	Info.Printf("}")
	Info.Printf("Node: %v , Label: %v , FQDN: %v\n",evt.Node,evt.Label,evt.FQDN)
	storage.Insert_name_registered2(&evt)
}
func proc_name_renewed(log *types.Log,evt_id,tx_id,timestamp int64) {
	var evt ENS_NameRenewed
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}
	evt.TxHash = log.TxHash.String()
	var eth_event NameRenewed
	err := ens_abi.Unpack(&eth_event,"NameRenewed",log.Data)
	if err != nil {
		Error.Printf("Error upacking NameRenewed: %v\n",err)
		Info.Printf("Error upacking NameRenewed: %v\n",err)
		os.Exit(1)
	}

	eth_event.Label = log.Topics[1]
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	eth_event.Dump(Info)
	evt.TxHash = log.TxHash.String()
	evt.Label = hex.EncodeToString(eth_event.Label[:])

	node_hash,_,err := get_node_hash_via_new_owner_event(evt.TxId,&log.TxHash,eth_event.Label[:],true)
	if err != nil {
		Error.Printf("Error getting node hash: %v\n",err)
		os.Exit(1)
	}

	fqdn_bytes := calculate_name_hash(node_hash,eth_event.Label[:])
	evt.FQDN = hex.EncodeToString(fqdn_bytes[:])
	Info.Printf("resulting fqdn: %v\n",hex.EncodeToString(fqdn_bytes[:]))

	evt.Node = hex.EncodeToString(node_hash[:])
	evt.Name = eth_event.Name
	evt.Cost = eth_event.Cost.String()
	evt.Expires = eth_event.Expires.Int64()
	evt.Contract = log.Address.String()
	storage.Insert_name_renewed(&evt)
}
func proc_newowner(log *types.Log,evt_id,tx_id,timestamp int64) {

///		Info.Printf("%v: log = %+v\n",i,log)
	var evt ENS_NewOwner
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		bnum := big.NewInt(int64(log.BlockNumber))
		ctx := context.Background()
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}
	evt.Label = hex.EncodeToString(log.Topics[2][:])
	evt.Node = hex.EncodeToString(log.Topics[1][:])
	evt.Owner = common.BytesToAddress(log.Data[12:]).String()

	fqdn_bytes := calculate_name_hash(log.Topics[1][:],log.Topics[2][:])
	evt.FQDN = hex.EncodeToString(fqdn_bytes)
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	Info.Printf("NewOwner %v for %v (node: %v, label: %v)\n",evt.Owner,evt.FQDN,evt.Node,evt.Label)
	evt.TxHash = log.TxHash.String()
	evt.Contract = log.Address.String()
	storage.Insert_new_owner(&evt)

}
func proc_addr_changed1(log *types.Log,evt_id,tx_id,timestamp int64) {

///		Info.Printf("%v: log = %+v\n",i,log)
	var evt ENS_AddrChanged
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TxHash = log.TxHash.String()
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		bnum := big.NewInt(int64(log.BlockNumber))
		ctx := context.Background()
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}
	if len(log.Data) < 32 {	// not our event
		return
	}
	fqdn_hash := log.Topics[1][:]
	evt.FQDN = hex.EncodeToString(fqdn_hash)
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	Info.Printf("AddrChanged (addr= %v ) (fqdn: %v ) \n",evt.Address,evt.FQDN)
	evt.Contract = log.Address.String()
	addr := common.BytesToAddress(log.Data[12:])
	evt.Address = addr.String()
	storage.Insert_address_changed1(&evt)

}
func proc_address_changed2(log *types.Log,evt_id,tx_id,timestamp int64) {

///		Info.Printf("%v: log = %+v\n",i,log)
	var evt ENS_AddressChanged
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TxHash = log.TxHash.String()
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		bnum := big.NewInt(int64(log.BlockNumber))
		ctx := context.Background()
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}
	if len(log.Data) < 32 {	// not our event
		return
	}
	fqdn_hash := log.Topics[1][:]
	evt.FQDN = hex.EncodeToString(fqdn_hash)
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	Info.Printf("AddressChanged(2) (addr= %v ) (node: %v ) (coin: %v) \n",evt.Address,evt.FQDN,evt.CoinType)
	evt.Contract = log.Address.String()
	var eth_event AddressChanged
	err := ens_abi.Unpack(&eth_event,"AddressChanged",log.Data)
	if err != nil {
		Error.Printf("Error upacking AddressChanged: %v\n",err)
		Info.Printf("Error upacking AddressChanged: %v\n",err)
		os.Exit(1)
	}
	evt.CoinType=int(eth_event.CoinType.Int64())
	new_addr:=common.Address{}
	new_addr.SetBytes(eth_event.NewAddress[:])
	evt.Address = new_addr.String()
	storage.Insert_address_changed2(&evt)

}
func proc_name_registered3(log *types.Log,evt_id,tx_id,timestamp int64) {

	var evt ENS_Name3
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	var eth_event NameRegistered_v3
	err := ens_abi.Unpack(&eth_event,"NameRegistered3",log.Data)
	if err != nil {
		Error.Printf("Error upacking NameRegistered3: %v\n",err)
		Info.Printf("Error upacking NameRegistered3: %v\n",err)
		os.Exit(1)
	}
	eth_event.Caller= common.BytesToAddress(log.Topics[1][12:])
	eth_event.Beneficiary = common.BytesToAddress(log.Topics[2][12:])
	eth_event.Label = log.Topics[3]
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TxHash = log.TxHash.String()
	evt.Label = hex.EncodeToString(eth_event.Label[:])
	evt.Subdomain = eth_event.Subdomain
	evt.CreatedDate = eth_event.CreatedDate.Int64()
	evt.Contract = log.Address.String()
	evt.Beneficiary = eth_event.Beneficiary.String()
	evt.Caller = eth_event.Caller.String()

	node_hash,_,err := get_node_hash_via_new_owner_event(evt.TxId,&log.TxHash,eth_event.Label[:],true)
	if err != nil {
		Error.Printf("Error getting node hash: %v\n",err)
		os.Exit(1)
	}

	fqdn_bytes := calculate_name_hash(node_hash,eth_event.Label[:])
	evt.FQDN = hex.EncodeToString(fqdn_bytes)
	evt.Node = hex.EncodeToString(node_hash[:])

	Info.Printf("ENS_Name3 {\n")
	Info.Printf("\tCaller: %v\n",eth_event.Caller.String())
	Info.Printf("\tBeneficiary: %v\n",eth_event.Beneficiary.String())
	Info.Printf("\tLabel: %v\n",eth_event.Label)
	Info.Printf("\tSubdomain: %v\n",eth_event.Subdomain)
	Info.Printf("\tCreatedDate: %v\n",evt.CreatedDate)
	Info.Printf("}")
	Info.Printf("(node: %v, fqdn: %v\n",evt.Node,evt.FQDN)

	storage.Insert_name_registered3(&evt)
}
func proc_name_changed(log *types.Log,evt_id,tx_id,timestamp int64) {

	var evt ENS_NameChanged
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TxHash = log.TxHash.String()
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		bnum := big.NewInt(int64(log.BlockNumber))
		ctx := context.Background()
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}
	if len(log.Data) < 32 {	// not our event
		return
	}
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	Info.Printf("NameChanged (name= %v ) (coin: %v) \n",evt.Name,evt.Node)
	evt.Contract = log.Address.String()
	var eth_event NameChanged
	err := ens_abi.Unpack(&eth_event,"NameChanged",log.Data)
	if err != nil {
		Error.Printf("Error upacking NameChanged: %v\n",err)
		Info.Printf("Error upacking NameChanged: %v\n",err)
		os.Exit(1)
	}
	evt.Node = hex.EncodeToString(log.Topics[1][:])
	evt.Name = eth_event.Name
	storage.Insert_name_changed(&evt)

}
func proc_hash_invalidated(log *types.Log,evt_id,tx_id,timestamp int64) {

	var evt ENS_HashInvalidated
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	var eth_event HashInvalidated
	err := ens_abi.Unpack(&eth_event,"HashInvalidated",log.Data)
	if err != nil {
		Error.Printf("Error upacking HashInvalidated: %v\n",err)
		Info.Printf("Error upacking HashInvalidated: %v\n",err)
		os.Exit(1)
	}
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
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
func proc_new_resolver(log *types.Log,evt_id,tx_id,timestamp int64) {
	var evt ENS_NewResolver
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
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
func proc_registry_transfer(log *types.Log,evt_id,tx_id,timestamp int64) {

	var evt ENS_RegistryTransfer
	evt.EvtId = evt_id
	evt.TimeStamp = timestamp
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
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
func proc_registrar_transfer(log *types.Log,evt_id,tx_id,timestamp int64) {
	var evt ENS_RegistrarTransfer
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp

	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	if len(log.Topics) !=4 {
		Info.Printf("log.Topics length invalid, skipping\n")
		return
	}

	evt.TxHash = log.TxHash.String()
	from_addr := common.BytesToAddress(log.Topics[1][12:])
	to_addr := common.BytesToAddress(log.Topics[2][12:])
	token_id := log.Topics[3][:]
	evt.TokenId = hex.EncodeToString(token_id[:])
	evt.From = from_addr.String()
	evt.To = to_addr.String()
	evt.Contract = log.Address.String()

	base_registrar_ctrct,err := baseregistrar.NewContract(log.Address,eclient)
	if err != nil {
		Error.Printf("Error instantiating base registrar contract %v : %v\n",log.Address.String(),err)
		Info.Printf("Error instantiating base registrar contract %v : %v\n",log.Address.String(),err)
		os.Exit(1)
	}
	var node_hash []byte
	label_hash := token_id
	var node_bytes[32]byte
	var copts = new(bind.CallOpts)
	node_bytes,err = base_registrar_ctrct.BaseNode(copts)
	if err == nil {
		node_hash = node_bytes[:]
	} else {
		node_hash,_,err = get_node_hash_via_new_owner_event(evt.TxId,&log.TxHash,label_hash,true)
		if err != nil {
			Error.Printf("Error getting node hash: %v\n",err)
			//os.Exit(1)
			return
		}
	}
	fqdn_bytes := calculate_name_hash(node_hash,label_hash)
	evt.FQDN = hex.EncodeToString(fqdn_bytes)
	evt.Node = hex.EncodeToString(node_hash[:])
	evt.Label = hex.EncodeToString(label_hash[:])

	Info.Printf("Registrar Transfer{\n")
	Info.Printf("\tFrom: %v\n",evt.From)
	Info.Printf("\tTo: %v\n",evt.To)
	Info.Printf("\tTokenID: %v\n",evt.TokenId)
	Info.Printf("}")
	Info.Printf("Node: %v , Label: %v , FQDN: %v\n",evt.Node,evt.Label,evt.FQDN)
	storage.Insert_registrar_transfer(&evt)
}
func proc_text_changed(log *types.Log,evt_id,tx_id,timestamp int64) {

	var evt ENS_TextChanged
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
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
func proc_hash_registered(log *types.Log,evt_id,tx_id,timestamp int64) {

	var evt ENS_HashRegistered
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	var eth_event HashRegistered
	err := ens_abi.Unpack(&eth_event,"HashRegistered",log.Data)
	if err != nil {
		Error.Printf("Error upacking HashRegistered: %v\n",err)
		Info.Printf("Error upacking HashRegistered: %v\n",err)
		os.Exit(1)
	}
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
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
func proc_owner_changed(log *types.Log,evt_id,tx_id,timestamp int64) {
/*
	var evt ENS_HashRegistered
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	var eth_event HashRegistered
	err := ens_abi.Unpack(&eth_event,"HashRegistered",log.Data)
	if err != nil {
		Error.Printf("Error upacking HashRegistered: %v\n",err)
		Info.Printf("Error upacking HashRegistered: %v\n",err)
		os.Exit(1)
	}
	copy(eth_event.Hash[:],log.Topics[1].Bytes())
	//eth_event.Name = Bytes32_to_string(log.Topics[2].Bytes())
	eth_event.Owner = common.BytesToAddress(log.Topics[2][12:])
	evt.Owner = eth_event.Owner.String()
	evt.TxHash = log.TxHash.String()
	evt.Hash = hex.EncodeToString(eth_event.Hash[:])
	evt.RegistrationDate=eth_event.RegistrationDate.Int64()
	evt.Value = eth_event.Value.String()
	evt.Contract = log.Address.String()

	*/
	Info.Printf("Processing block %v, tx %v\n",int64(log.BlockNumber),log.TxHash.String())
	Info.Printf("OwnerChanged {\n")
	Info.Printf("\tLabel: %v\n",hex.EncodeToString(log.Topics[1][:]))
	Info.Printf("\toldOwner: %v\n",hex.EncodeToString(log.Topics[2][:]))
	Info.Printf("\tnewOwner: %v\n",hex.EncodeToString(log.Topics[3][:]))
	Info.Printf("}")
}
func proc_pubkey_changed(log *types.Log,evt_id,tx_id,timestamp int64) {

	var evt ENS_PubkeyChanged
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	var eth_event PubkeyChanged
	err := ens_abi.Unpack(&eth_event,"PubkeyChanged",log.Data)
	if err != nil {
		Error.Printf("Error upacking PubkeyChanged: %v\n",err)
		Info.Printf("Error upacking PubkeyChanged: %v\n",err)
		os.Exit(1)
	}
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TxHash = log.TxHash.String()
	evt.Contract = log.Address.String()
	evt.Node = hex.EncodeToString(log.Topics[1][:])
	evt.X = hex.EncodeToString(eth_event.X[:])
	evt.Y = hex.EncodeToString(eth_event.Y[:])

	Info.Printf("PubkeyChanged {\n")
	Info.Printf("\tNode: %v\n",hex.EncodeToString(eth_event.Node[:]))
	Info.Printf("\tX: %v\n",hex.EncodeToString(eth_event.X[:]))
	Info.Printf("\tY: %v\n",hex.EncodeToString(eth_event.Y[:]))
	Info.Printf("}")
	xhash := common.BytesToHash(eth_event.X[:])
	yhash := common.BytesToHash(eth_event.Y[:])
	xb := xhash.Big()
	yb := yhash.Big()
	pkey := &ecdsa.PublicKey {
		Curve:	secp256k1.S256(),
		X:		xb,
		Y:		yb,
	}
	addr := crypto.PubkeyToAddress(*pkey)
	evt.DerivedAddr = addr.String()
//		Error.Printf("X/Y decode errors for tx %v : err1=%v, err2=%v\n",evt.TxHash,err1,err2)
//		Info.Printf("X/Y decode errors for tx %v : err1=%v, err2=%v\n",evt.TxHash,err1,err2)

	storage.Insert_pubkey_changed(&evt)
}
func proc_contenthash_changed(log *types.Log,evt_id,tx_id,timestamp int64) {

	var evt ENS_ContentHashChanged
	evt.EvtId = evt_id
	evt.BlockNum = int64(log.BlockNumber)
	evt.TxId = tx_id
	evt.TimeStamp = timestamp
	if evt.TimeStamp == 0 {
		ctx := context.Background()
		bnum := big.NewInt(int64(log.BlockNumber))
		block_hdr,err := eclient.HeaderByNumber(ctx,bnum)
		if err != nil {
			Error.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			Info.Printf("Error getting block header %v : %v\n",log.BlockNumber,err)
			os.Exit(1)
		}
		evt.TimeStamp = int64(block_hdr.Time)
	}

	var eth_event ContenthashChanged
	err := ens_abi.Unpack(&eth_event,"ContenthashChanged",log.Data)
	if err != nil {
		Error.Printf("Error upacking ContentHashChanged: %v\n",err)
		Info.Printf("Error upacking ContentHashChanged: %v\n",err)
		os.Exit(1)
	}
	Info.Printf("Processing block %v, tx %v\n",evt.BlockNum,log.TxHash.String())
	evt.TxHash = log.TxHash.String()
	evt.Contract = log.Address.String()
	evt.Node = hex.EncodeToString(log.Topics[1][:])
	evt.Hash = hex.EncodeToString(eth_event.Hash[:])

	Info.Printf("ContentHashChanged {\n")
	Info.Printf("\tNode: %v\n",hex.EncodeToString(eth_event.Node[:]))
	Info.Printf("\tHash: %v\n",hex.EncodeToString(eth_event.Hash[:]))
	Info.Printf("}")
	storage.Insert_contenthash_changed(&evt)
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
			storage.Expire_ens_names(Info)
		}
	}
}
func process_ens_event(evt_id int64) error {

	evtlog := storage.Get_event_log(evt_id)
	var log types.Log
	err := rlp.DecodeBytes(evtlog.RlpLog,&log)
	if err!= nil {
		panic(fmt.Sprintf("RLP Decode error: %v",err))
	}
	log.BlockNumber=uint64(evtlog.BlockNum)
	log.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
	log.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
	num_topics := len(log.Topics)
	if num_topics > 0 {
		Info.Printf("found event with sig = %v\n",log.Topics[0].String())
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_newowner) {
			proc_newowner(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_addrchanged1) {
			proc_addr_changed1(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_addresschanged2) {
			proc_address_changed2(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_name_registered1) {
			proc_name_registered1(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_name_registered2) {
			proc_name_registered2(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_name_registered3) {
			proc_name_registered3(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_hash_invalidated) {
			proc_hash_invalidated(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_hash_registered) {
			proc_hash_registered(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_new_resolver) {
			proc_new_resolver(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_registry_transfer) {
			proc_registry_transfer(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_registrar_transfer) {
			proc_registrar_transfer(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_text_changed) {
			proc_text_changed(&log,evtlog.EvtId,evtlog.TxId,evtlog.TimeStamp)
		}
		if 0 == bytes.Compare(log.Topics[0].Bytes(),evt_name_bought) {
			//PENDING
			//proc_(&log)
		}

	}
	return nil
}
	/*this code doesn't work, pending for removal
	// Fetch Node hash from Contract itself
	base_registrar_ctrct,err := baseregistrar.NewContract(log.Address,eclient)
	if err != nil {
		Error.Printf("Error instantiating base registrar contract %v : %v\n",log.Address.String(),err)
		Info.Printf("Error instantiating base registrar contract %v : %v\n",log.Address.String(),err)
		os.Exit(1)
	}
	var node_hash [32]byte
	ens_resolver,err := storage.Get_ens_resolver(log.Address.String())
	if err != nil {
		Info.Printf("No resolver %v registered in the DB, querying geth\n",log.Address.String())
		var copts = new(bind.CallOpts)
		node_hash,err = base_registrar_ctrct.BaseNode(copts)
		if err != nil {
			Error.Printf("Error calling baseNode() for ctrct %v : %v\n",log.Address.String(),err)
			Info.Printf("Error calling baseNode() for ctrct %v : %v\n",log.Address.String(),err)
			auction_registrar_ctrct,err := auctionregistrar.NewContract(log.Address,eclient)
			if err != nil {
				Error.Printf("Error instantiating auction registrar contract %v : %v\n",log.Address.String(),err)
				Info.Printf("Error instantiating auction registrar contract %v : %v\n",log.Address.String(),err)
				os.Exit(1)
			}
			node_hash,err = auction_registrar_ctrct.RootNode(copts)
			if err != nil {
				Error.Printf("Error calling rootNode() for ctrct %v : %v\n",log.Address.String(),err)
				Info.Printf("Error calling rootNode() for ctrct %v : %v\n",log.Address.String(),err)
				os.Exit(1)
			}
		}
	} else {
		Info.Printf("resolver.Node=%v\n",ens_resolver.Node)
		hash,err := hex.DecodeString(ens_resolver.Node)
		if err!=nil {
			Error.Printf("Error decoding node hash str for %v: %v\n",log.Address.String,err)
			Info.Printf("Error decoding node hash str for %v: %v\n",log.Address.String,err)
			os.Exit(1)
		}
		copy(node_hash[:],hash[:])
	}
*/
