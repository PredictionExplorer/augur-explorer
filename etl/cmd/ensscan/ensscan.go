package main

import (
	"os"
	"os/signal"
	"syscall"
	"context"
	"math/big"
	"fmt"
	"log"
	"bytes"
	"strings"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/wealdtech/go-ens/v3"
	"github.com/wealdtech/go-ens/v3/contracts/resolver"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	GRACE_PERIOD			int = 90 * 24 * 60 * 60 // 90 days
	ENS_NEWOWNER				= "ce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82"

	ENS_NAME_REGISTERED1		= "ca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f"
	//	NameRegistered (	// this was introduced in 2019
	//		string name,
	//		index_topic_1 bytes32 label,
	//		index_topic_2 address owner,
	//		uint256 cost,
	//		uint256 expires
	//	)

	ENS_NAME_REGISTERED2		= "b3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9"
	// NameRegistered (
	//		index_topic_1 uint256 id,
	//		index_topic_2 address owner,
	//		uint256 expires
	//	)

	ENS_NAME_REGISTERED3		= "570313dae523ecb48b1176a4b60272e5ea7ec637f5b2d09983cbc4bf25e7e9e3"
	//	NameRegistered (
	//		index_topic_1 address _caller,
	//		index_topic_2 address _beneficiary,
	//		index_topic_3 bytes32 _labelHash,
	//		string _subdomain,
	//		uint256 _createdDate
	//	)

	ENS_ADDR_CHANGED			= "52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2"
	ENS_ADDRESS_CHANGED			= "65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752"
	NEW_RESOLVER				= "335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0"
	ENS_TRANSFER				= "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	ENS_REGISTRY_TRANSFER		= "d4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266"
	HASH_REGISTERED				= "0f0c27adfd84b60b6f456b0e87cdccb1e5fb9603991588d87fa99f5b6b61e670"
	//  HashRegistered (	Sample tx: 0x216d8103a59c3f3921210b3a4c6aca32c21724dac8451d32ed0d88103c20b802
	//		index_topic_1 bytes32 hash,
	//		index_topic_2 address owner,
	//		uint256 value,
	//		uint256 registrationDate
	//	)
	HASH_INVALIDATED			= "1f9c649fe47e58bb60f4e52f0d90e4c47a526c9f90c5113df842c025970b66ad"
	NEW_TTL						= "1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68."
	ENS_TEXT_CHANGED			= "d8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550"
	NAME_BOUGHT					= "0xb8c56202a5ae8b00edfcd57a54ec6c3fb8d2f6deb3067a7ba11408a7bd014a3e"

	ENS_V1_REGISTRY_ADDR		= "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"	// 10 Mar 2017
	ENS_V2_REGISTRY_ADDR		= "0x314159265dD8dbb310642f98f50C066173C1259b"	// 30 Jan 2020

	//[common.Address].addr.reverse
	ADDRIN_REVERSE_NODE			= "91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e2"
	REVERSE_REG_V2_ADDR			= "0x084b1c3C81545d370f3634392De611CaaBFf8148"
	REVERSE_REG_V1_ADDR			= "0x9062C0A6Dbd6108336BcBe4593a3D1cE05512069"
)
var (
	evt_newowner,_ = hex.DecodeString(ENS_NEWOWNER)
	evt_name_registered1,_ = hex.DecodeString(ENS_NAME_REGISTERED1)
	evt_name_registered2,_ = hex.DecodeString(ENS_NAME_REGISTERED2)
	evt_name_registered3,_ = hex.DecodeString(ENS_NAME_REGISTERED3)
	evt_hash_invalidated,_ = hex.DecodeString(HASH_INVALIDATED)
	evt_hash_registered,_ = hex.DecodeString(HASH_REGISTERED)
	evt_new_resolver,_ = hex.DecodeString(NEW_RESOLVER)
	evt_registry_transfer,_ = hex.DecodeString(ENS_REGISTRY_TRANSFER)
	evt_text_changed,_		= hex.DecodeString(ENS_TEXT_CHANGED)
	evt_name_bought,_ = hex.DecodeString(NAME_BOUGHT)

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
	Error   *log.Logger
	Info	*log.Logger

	market_order_id int64 = 0
	ens_abi abi.ABI

	ens1_addr			= common.HexToAddress(ENS_V1_REGISTRY_ADDR)
	ens2_addr			= common.HexToAddress(ENS_V2_REGISTRY_ADDR)
)
func do_initiial_load_name_registered1(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
//	filter.FromBlock = big.NewInt(0)
//	filter.ToBlock = nil
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_name_registered1)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("NameRegisterd1: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
///		Info.Printf("%v: log = %+v\n",i,log)
		var evt ENS_Name1
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
		var eth_event NameRegistered_v1
		err := ens_abi.Unpack(&eth_event,"NameRegistered1",log.Data)
		if err != nil {
			Error.Printf("Error upacking NameRegistered1: %v\n",err)
			Info.Printf("Error upacking NameRegistered1: %v\n",err)
			continue
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
}
func do_initiial_load_new_owner(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
//	filter.FromBlock = big.NewInt(0)
//	filter.ToBlock = nil
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_newowner)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("NewOwner: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		if	ens1,ens2 :=
				bytes.Equal(ens1_addr.Bytes(),log.Address.Bytes()),
				bytes.Equal(ens2_addr.Bytes(),log.Address.Bytes());
			!(ens1 || ens2) {
				continue
		}
///		Info.Printf("%v: log = %+v\n",i,log)
		var evt ENS_NewOwner
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
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
}
func do_initiial_load_name_registered2(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
//	filter.FromBlock = big.NewInt(0)
//	filter.ToBlock = nil
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_name_registered2)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("NameRegisterd2: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
///		Info.Printf("%v: log = %+v\n",i,log)
		var evt ENS_Name2
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
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
}
func range_initial_load_name_registered1(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 7691000 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_name_registered1(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_name_registered1(block_num,next_block_num)
			block_num = next_block_num + 1
		}
		storage.Expire_ens_names(Info)
	}
}
func range_initial_load_new_owner(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 2933000 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_new_owner(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_new_owner(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func range_initial_load_name_registered2(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 0 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_name_registered2(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_name_registered2(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func do_initiial_load_name_registered3(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_name_registered3)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("NameRegisterd3: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		var evt ENS_Name3
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
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
			continue
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
}
func range_initial_load_name_registered3(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 0 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_name_registered3(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_name_registered3(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func do_initiial_load_hash_invalidated(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_hash_invalidated)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("HashInvalidated : block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		var evt ENS_HashInvalidated
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
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
			continue
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
}
func range_initial_load_hash_invalidated(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 0 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_hash_invalidated(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_hash_invalidated(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func do_initiial_load_new_resolver(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_new_resolver)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("NewResolver: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		var evt ENS_NewResolver
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
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
}
func range_initial_load_new_resolver(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 0 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_new_resolver(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_new_resolver(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func do_initiial_load_registry_transfer(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_registry_transfer)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("Registry Transfer: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		var evt ENS_RegistryTransfer
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
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
}
func range_initial_load_registry_transfer(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 0 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_registry_transfer(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_registry_transfer(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func do_initiial_load_text_changed(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_text_changed)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("TextChanged: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		var evt ENS_TextChanged
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
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
			continue
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
			continue
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
}
func range_initial_load_text_changed(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 0 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_text_changed(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_text_changed(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func do_initiial_load_hash_registered(block_num_from,block_num_to int64) {

	filter := ethereum.FilterQuery{}
	filter.FromBlock = big.NewInt(block_num_from)
	filter.ToBlock = big.NewInt(block_num_to)
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_hash_registered)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	Info.Printf("HashRegistered: block range: %v - %v\n",block_num_from,block_num_to)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	for _,log := range logs {
		if log.Removed {
			continue
		}
		var evt ENS_HashRegistered
		evt.EvtId = 0
		evt.BlockNum = int64(log.BlockNumber)
		evt.TxId = 0
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
			continue
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
}
func range_initial_load_hash_registered(exit_chan chan bool,block_num_limit int64) {

	var block_num int64 = 0 // found empirically
	for ; block_num <= block_num_limit ; {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}

		next_block_num := block_num + 1000 - 1
		if next_block_num > block_num_limit {
			do_initiial_load_hash_registered(block_num,block_num_limit)
			break
		} else {
			do_initiial_load_hash_registered(block_num,next_block_num)
			block_num = next_block_num + 1
		}
	}
}
func initial_load(exit_chan chan bool,block_num_limit int64) {
	//range_initial_load_name_registered1(exit_chan,block_num_limit)
	//range_initial_load_new_owner(exit_chan,block_num_limit)
	//range_initial_load_name_registered2(exit_chan,block_num_limit)
	//range_initial_load_name_registered3(exit_chan,block_num_limit)
	//range_initial_load_hash_invalidated(exit_chan,block_num_limit)
	//range_initial_load_new_resolver(exit_chan,block_num_limit)
	//range_initial_load_registry_transfer(exit_chan,block_num_limit)
	//range_initial_load_text_changed(exit_chan,block_num_limit)
	range_initial_load_hash_registered(exit_chan,block_num_limit)
}
func check_initial_load_completness() bool {

	const correct_num_active_names int64 = 14040; // empirically found
	num_active_names := storage.Get_count_of_active_names()
	if num_active_names != correct_num_active_names {
		msg := fmt.Sprintf(
			"Number of active names is %v but should be %v\n",
			num_active_names,
			correct_num_active_names,
		)
		Error.Printf(msg)
		Info.Printf(msg)
		return false
	}
	return true
}
func main() {

	var do_initial_load bool = false
	if len(os.Args) > 1 {
		if os.Args[1] != "-i" {
			fmt.Printf("Unknown option %v\n",os.Args[1])
			fmt.Printf("Usage: %v [-i]\n\t -i = initial load\n",os.Args[0])
			os.Exit(1)
		}
		do_initial_load = true
	}
	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/ensscan_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/ensscan_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/ensscan_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	rpcclient, err := rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)
	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()
	abi_parsed := strings.NewReader(ENSABI)
	ens_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse ENS ABI: %v\n",err)
		os.Exit(1)
	}
	status := storage.Get_ens_processing_status()
	if do_initial_load {
		initial_load(exit_chan,status.IniLoadBlockNumLimit)
		fmt.Printf("Initial load finished.")
		return
	}
}
