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
	"encoding/hex"
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
	ERC20_TRANSFER = "ddf252ad"
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
	erc20_abi abi.ABI
	all_contracts map[string]interface{}
	caddrs *ContractAddresses
	bad_erc20_token map[common.Address]struct{}
	bad_for_decode map[common.Address]struct{}
	info_checked map[common.Address]struct{}

	err_invalid_erc20_format error = errors.New("Invalid ERC20 event structure (num topics != 3)")
	inspected_events []InspectedEvent
)
var (
	evt_erc20_transfer,_ = hex.DecodeString(ERC20_TRANSFER)

)
func fetch_and_store_erc20_info(token_addr common.Address) (int,error) {
	// note: this func is called as goroutine for speed. however duplicate calls can occur,
	//      which are prevented with DO NOTHING on conflict in the INSERT query
	_,already_checked := info_checked[token_addr]
	if already_checked {
		return 0,nil
	}
	info_checked[token_addr] = struct{}{}
	_,is_bad := bad_erc20_token[token_addr]
	if is_bad {
		return 0,nil
	}
	found,info := storage.Get_ERC20Info_v2(token_addr.String())
	if found {
		return info.Decimals,nil
	}
	erc20_info,err := Fetch_erc20_info(eclient,&token_addr)
	if err != nil {
		Error.Printf("Couldn't fetch ERC20 token info for addr %v : %v\n",token_addr.String(),err)
		bad_erc20_token[token_addr]=struct{}{}
		return 0,err
	}
	erc20_info.Address = token_addr.String()
	storage.Update_ERC20Info_v2(&erc20_info)
	return erc20_info.Decimals,nil
}

func build_list_of_inspected_events() []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)

	inspected_events= make([]InspectedEvent,0,32)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_erc20_transfer[:4]),
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
func proc_erc20_transfer(evt_id int64) error {
	evtlog := storage.Get_event_log(evt_id)
	var log types.Log
	err := rlp.DecodeBytes(evtlog.RlpLog,&log)
	if err!= nil {
		panic(fmt.Sprintf("RLP Decode error: %v",err))
	}
	_,bad_token := bad_for_decode[log.Address]
	if bad_token {
		return nil
	}
	log.BlockNumber=uint64(evtlog.BlockNum)
	log.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
	log.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
	var mevt ETransfer
	if len(log.Topics) < 3 {
		Info.Printf("ERC20 transfer event is not compliant log.Topics < 3. Tx hash=%v\n",log.TxHash.String())
		return err_invalid_erc20_format
	}
	mevt.From= common.BytesToAddress(log.Topics[1][12:])
	mevt.To= common.BytesToAddress(log.Topics[2][12:])
	err = erc20_abi.Unpack(&mevt,"Transfer",log.Data)
	if err != nil {

		Error.Printf("signature=%v\n",log.Topics[0].String())
		Error.Printf("address=%v\n",log.Address.String())
		Error.Printf("tx hash = %v\n",log.TxHash.String())
		Error.Printf("log.Data=%v, data len=%v\n",hex.EncodeToString(log.Data[:]),len(log.Data[:]))
		Error.Printf("Event ERC20_Transfer, decode error: %v",err)
		bad_for_decode[log.Address]=struct{}{}
	} else {
		Info.Printf("ERC20_Transfer event, contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump(Info)
		storage.Insert_ERC20_token_transfer(log.Address.String(),&mevt,evtlog.BlockNum,evtlog.TxId,evt_id,evtlog.TimeStamp)

		fetch_and_store_erc20_info(log.Address)
	}
	return nil
}
func process_erc20_tokens(exit_chan chan bool) {

	var max_batch_size int64 = 1024*8
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}

		status := storage.Get_erc20_process_status()
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
			err := proc_erc20_transfer(evt_id)
			if err != nil {
				if err != err_invalid_erc20_format {
					fmt.Printf("Exiting on error: %v\n",err)
					os.Exit(1)
				}
			}
			status.LastEvtId = evt_id
			storage.Update_erc20_process_status(&status)
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
				storage.Update_erc20_process_status(&status)
			}
			time.Sleep(1 * time.Second) // sleep only if there is no data
			return
		}
	}
}
func process_tokens(exit_chan chan bool,caddrs *ContractAddresses) {
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}
		process_erc20_tokens(exit_chan)
		time.Sleep(1 * time.Second)
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/erc20_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/erc20_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/erc20_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/erc20_balance_update.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	BalancesLog = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

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

	caddrs_obj,err := storage.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs = &caddrs_obj

	abi_parsed := strings.NewReader(OwnedERC20ABI)
	erc20_abi,err = abi.JSON(abi_parsed)
	if err != nil {
		Info.Printf("Can't parse ERC20 token ABI")
		os.Exit(1)
	}

	bad_erc20_token = make(map[common.Address]struct{})
	bad_for_decode = make(map[common.Address]struct{})
	info_checked = make(map[common.Address]struct{})
	build_list_of_inspected_events()
	//go balance_updater(BalancesLog)
	process_tokens(exit_chan,&caddrs_obj)
}
