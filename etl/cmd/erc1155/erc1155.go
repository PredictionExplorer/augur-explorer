package main

import (
	"os"
	"os/signal"
	"syscall"
	"errors"
	"time"
	"fmt"
	"log"
	"strings"
	"context"
	"sort"
	"bytes"
	"encoding/hex"
	"unicode/utf8"
	//"encoding/json"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	ERC1155_TRANSFER_SINGLE = "c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	ERC1155_TRANSFER_BATCH = "4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
	ERC1155_URI = "6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b"
)
var (
	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
	Error   *log.Logger
	Info	*log.Logger
	BalancesLog	*log.Logger
	market_order_id int64 = 0
	erc1155_abi abi.ABI
	all_contracts map[string]interface{}

	err_invalid_erc1155_format error = errors.New("Invalid ERC1155 event structure)")
	inspected_events []InspectedEvent
)
var (
	evt_erc1155_transfer_single,_ = hex.DecodeString(ERC1155_TRANSFER_SINGLE)
	evt_erc1155_transfer_batch,_ = hex.DecodeString(ERC1155_TRANSFER_BATCH)
	evt_erc1155_uri,_ = hex.DecodeString(ERC1155_URI)
)
func build_list_of_inspected_events() []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)

	inspected_events= make([]InspectedEvent,0,32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_erc1155_transfer_single[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_erc1155_transfer_batch[:4]),
			ContractAid: 0,
		},
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_erc1155_uri[:4]),
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
		output = append(output,event_list...)
	}
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	num_elts:=Remove_duplicates_int64(output)
	return output[0:num_elts]
}
func proc_erc1155_transfer_single(evtlog *EthereumEventLog) error {
	var log types.Log
	err := rlp.DecodeBytes(evtlog.RlpLog,&log)
	if err!= nil {
		panic(fmt.Sprintf("RLP Decode error: %v",err))
	}
	log.BlockNumber=uint64(evtlog.BlockNum)
	log.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
	log.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
	var mevt Evt_ERC1155TransferSingle
	if len(log.Topics) != 4  {
		Info.Printf("ERC1155 transfer single event is not compliant log.Topics != 4. Tx hash=%v\n",log.TxHash.String())
		return err_invalid_erc1155_format
	}
	mevt.Operator = common.BytesToAddress(log.Topics[1][12:])
	mevt.From = common.BytesToAddress(log.Topics[2][12:])
	mevt.To= common.BytesToAddress(log.Topics[3][12:])
	err = erc1155_abi.Unpack(&mevt,"TransferSingle",log.Data)
	if err != nil {
		Error.Printf("signature=%v\n",log.Topics[0].String())
		Error.Printf("address=%v\n",log.Address.String())
		Error.Printf("tx hash = %v\n",log.TxHash.String())
		Error.Printf("log.Data=%v, data len=%v\n",hex.EncodeToString(log.Data[:]),len(log.Data[:]))
		Error.Printf("Event ERC1155_Transfer, decode error: %v",err)
	} else {
		Info.Printf("ERC1155_TransferSingle event, contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		Info.Printf("ERC1155:TransferSingle {\n")
		Info.Printf("\tOperator: %v\n",mevt.Operator.String())
		Info.Printf("\tTokenId: %v\n",hex.EncodeToString(common.BigToHash(mevt.Id).Bytes()))
		Info.Printf("\tFrom: %v\n",mevt.From.String())
		Info.Printf("\tTo: %v\n",mevt.To.String())
		Info.Printf("\tValue: %v\n",mevt.Value.String())
		Info.Printf("}\n")
		storage.Insert_ERC1155_transfer_single(&mevt,log.Address.String(),evtlog.BlockNum,evtlog.TxId,evtlog.EvtId,evtlog.TimeStamp)
	}
	return nil
}
func proc_erc1155_transfer_batch(evtlog *EthereumEventLog) error {
	var log types.Log
	err := rlp.DecodeBytes(evtlog.RlpLog,&log)
	if err!= nil {
		panic(fmt.Sprintf("RLP Decode error: %v",err))
	}
	log.BlockNumber=uint64(evtlog.BlockNum)
	log.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
	log.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
	var mevt Evt_ERC1155TransferBatch
	if len(log.Topics) != 4  {
		Info.Printf("ERC1155 transfer batch event is not compliant log.Topics != 4. Tx hash=%v\n",log.TxHash.String())
		return err_invalid_erc1155_format
	}
	mevt.Operator = common.BytesToAddress(log.Topics[1][12:])
	mevt.From = common.BytesToAddress(log.Topics[2][12:])
	mevt.To= common.BytesToAddress(log.Topics[3][12:])
	err = erc1155_abi.Unpack(&mevt,"TransferBatch",log.Data)
	if err != nil {
		Error.Printf("signature=%v\n",log.Topics[0].String())
		Error.Printf("address=%v\n",log.Address.String())
		Error.Printf("tx hash = %v\n",log.TxHash.String())
		Error.Printf("log.Data=%v, data len=%v\n",hex.EncodeToString(log.Data[:]),len(log.Data[:]))
		Error.Printf("Event ERC1155_TransferBatch, decode error: %v",err)
		return err
	}
	var token_ids string
	for i:=0; i<len(mevt.Ids); i++ {
		token_id_hex := hex.EncodeToString(common.BigToHash(mevt.Ids[i]).Bytes())
		if len(token_ids) > 0 {
			token_ids = token_ids + ","
		}
		token_ids = token_ids + token_id_hex
	}
	if len(mevt.Values) != len(mevt.Ids) {
		err_str := fmt.Sprintf(
			"Invalid ERC1155 TransferBatch Event for evt_id=%v, tx %v , "+
													"values and token ids of different length",
			evtlog.EvtId,evtlog.TxHash,
		)
		Error.Printf("%v",err_str)
		Info.Printf("%v",err_str)
		return errors.New(err_str)
	}
	amounts := Bigint_ptr_slice_to_str(&mevt.Values,",")
	Info.Printf("ERC1155_TransferBatch event, contract %v (block=%v) :\n",
								log.Address.String(),log.BlockNumber)
	Info.Printf("ERC1155:TransferBatch {\n")
	Info.Printf("\tOperator: %v\n",mevt.Operator.String())
	Info.Printf("\tTokenIds: %v\n",token_ids)
	Info.Printf("\tFrom: %v\n",mevt.From.String())
	Info.Printf("\tTo: %v\n",mevt.To.String())
	Info.Printf("\tAmounts: %v\n",amounts)
	Info.Printf("}\n")
	storage.Insert_ERC1155_transfer_batch(&mevt,log.Address.String(),evtlog.BlockNum,evtlog.TxId,evtlog.EvtId,evtlog.TimeStamp)
	return nil
}
func proc_erc1155_uri(evtlog *EthereumEventLog) {

	var log types.Log
	err := rlp.DecodeBytes(evtlog.RlpLog,&log)
	if err!= nil {
		panic(fmt.Sprintf("RLP Decode error: %v",err))
	}
	var eth_evt Evt_ERC1155URI

	Info.Printf("Processing URI event id=%v, txhash %v\n",evtlog.EvtId,evtlog.TxHash)

	if len(log.Topics) < 2 {
		Error.Printf("URI in tx hash %v doesn't have indexed 'ID' field\n",evtlog.TxHash)
		return
	}
	eth_evt.Id = common.BytesToHash(log.Topics[1][:]).Big()

	err = erc1155_abi.Unpack(&eth_evt,"URI",log.Data)
	if err != nil {
		Error.Printf("Event URI decode error: %v",err)
		os.Exit(1)
	}
	var empty_array [32]byte
	//Info.Printf("Valu=%v\n",hex.EncodeToString([]byte(eth_evt.Value)))
	if bytes.Equal([]byte(eth_evt.Value),empty_array[:]) {
		eth_evt.Value = ""
	}
	if !utf8.ValidString(eth_evt.Value) {
		Error.Printf(
			"Invalid 'Value' field for Uri event: Value=%v, setting to '', tx_hash=%v\n",
			eth_evt.Value,evtlog.TxHash,
		)
		eth_evt.Value = ""
	}
	Info.Printf("Contract: %v\n",log.Address.String())
	Info.Printf("URI {\n")
	Info.Printf("\tId: %v\n",eth_evt.Id)
	Info.Printf("\tValue: %v\n",eth_evt.Value)
	Info.Printf("}\n")

	storage.Insert_ERC1155_URI(&eth_evt,log.Address.String(),evtlog.BlockNum,evtlog.TxId,evtlog.EvtId,evtlog.TimeStamp)
}

func process_erc1155_tokens(exit_chan chan bool) {

	var max_batch_size int64 = 1024*32
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}

		status := storage.Get_erc1155_process_status()
		id_upper_limit := status.LastEvtId + max_batch_size
		last_evt_id,err := storage.Get_last_evtlog_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_evt_id {
			id_upper_limit = last_evt_id
		}

		tok_events := get_event_ids(status.LastEvtId,id_upper_limit)
		for _,evt_id := range tok_events {
			evtlog := storage.Get_event_log(evt_id)
			sig_decoded,_ := hex.DecodeString(evtlog.Topic0_Sig)
			if bytes.Equal(sig_decoded,evt_erc1155_transfer_single[:4]) {
				proc_erc1155_transfer_single(&evtlog)
			}
			if bytes.Equal(sig_decoded,evt_erc1155_transfer_batch[:4]) {
				proc_erc1155_transfer_batch(&evtlog)
			}
			if bytes.Equal(sig_decoded,evt_erc1155_uri[:4]) {
				proc_erc1155_uri(&evtlog)
			}
			status.LastEvtId = evt_id
			storage.Update_erc1155_process_status(&status)
		}
		Info.Printf("processed %v events\n",len(tok_events))
		if len(tok_events) == 0 {
			last_evt_log_id_on_chain,err := storage.Get_last_evtlog_id()
			if err != nil {
				Info.Printf("Error getting last event log id: %v\n",err)
				os.Exit(1)
			}
			if last_evt_log_id_on_chain > id_upper_limit {
				// only advance upper range if events within the range have filled id value space
				status.LastEvtId = id_upper_limit
				storage.Update_erc1155_process_status(&status)
			}
			time.Sleep(1 * time.Second) // sleep only if there is no data
			return
		}
	}
}
func process_tokens(exit_chan chan bool) {
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}
		process_erc1155_tokens(exit_chan)
		time.Sleep(1 * time.Second)
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/erc1155_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/erc1155_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/erc1155_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	network_chain_id,err :=eclient.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("Error getting network id: %v\n",err)
		os.Exit(1)
	}
	stored_chain_id := storage.Get_stored_chain_id()
	if stored_chain_id != network_chain_id.Int64() {
		if stored_chain_id == 0 {
			fmt.Printf("Database not initialized, chain_id=0\n")
			os.Exit(1)
		} else {
			fmt.Printf(
				"Network chain_id = %v , my chain_id = %v. Mismatch, exiting",
				network_chain_id.Int64(),stored_chain_id,
			)
		}
	}

	abi_parsed := strings.NewReader(ERC1155ABI)
	erc1155_abi,err = abi.JSON(abi_parsed)
	if err != nil {
		Info.Printf("Can't parse ERC1155 token ABI")
		os.Exit(1)
	}

	build_list_of_inspected_events()
	process_tokens(exit_chan)
}
