package main

import (
	"os"
	"os/signal"
	"syscall"
	"bytes"
	"strings"
	"context"
	"time"
	"sort"
	"fmt"
	"log"
	"math/big"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	ERC20_TRANSFER				= "ddf252ad"
	PAIR_CREATED				= "0d3648bd"
	PAIR_SWAP					= "d78ad95f"
	WETH_DEPOSIT				= "e1fffcc4"
)
var (
	evt_pair_created,_ = hex.DecodeString(PAIR_CREATED)
	evt_pair_swap,_ = hex.DecodeString(PAIR_SWAP)

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
	Error   *log.Logger
	Info	*log.Logger
	inspected_events []InspectedEvent
	factory_abi abi.ABI
	pair_abi abi.ABI
	factory_addr_str string
)
func  update_uniswap_slippage_if_applicable(block_num int64,pair_addr string) {
	pair_aid,err := storage.Nonfatal_lookup_address_id(pair_addr)
	if err != nil {
		return
	}
	num_augur_tokens,err := storage.Get_uniswap_augur_tokens(pair_aid)
	if num_augur_tokens == 0 {
		return
	}
	pair_info,err := storage.Get_uniswap_pair_info(pair_aid)
	if err != nil {
		return
	}
	amount_to_trade := "100";
	_,_,router02_addr_str := storage.Get_uniswap_contracts()
	slippages,err := Produce_uniswap_slippages(eclient,router02_addr_str,&pair_info,amount_to_trade)
	storage.Update_uniswap_slippages(block_num,pair_aid,slippages)
}
func update_existing_pairs() {

	var copts = new(bind.CallOpts)
	bci := storage.Get_first_event_log()

	factory_addr := common.HexToAddress(factory_addr_str)
	factory_contract,err := NewUniswapV2Factory(factory_addr,eclient)
	if err != nil {
		Error.Printf("Error creating Uniswap Factory (%v) instance: %v\n",factory_addr_str,err)
		os.Exit(1)
	}
	pair_len,err := factory_contract.AllPairsLength(copts)
	if err != nil {
		Error.Printf("Error getting all pair length (f=%v): %v",factory_addr.String(),err)
		os.Exit(1)
	}
	ipair_len := int(pair_len.Int64())
	Info.Printf("beggining pair array scan len=%v\n",ipair_len)
	for i:=int(ipair_len); i > 0; i-- {	// scanning from the end for speed
		paddr,err := factory_contract.AllPairs(copts,big.NewInt(int64(i-1)))
		if err != nil {
			Error.Printf("Error scanning AllPairs array: %v\n",err)
			os.Exit(1)
		}
		exists := storage.Pair_exists(paddr.String())
		if !exists {
			pair_contract,err := NewUniswapV2Pair(paddr,eclient)
			if err != nil {
				Error.Printf("Error creating Uniswap Pair instance: %v\n",err)
				os.Exit(1)
			}
			token0,err := pair_contract.Token0(copts)
			if err != nil {
				Error.Printf("Error getting Token0 from Uniswap pair %v: %v\n",paddr.String(),err)
				os.Exit(1)
			}
			token1,err := pair_contract.Token1(copts)
			if err != nil {
				Error.Printf("Error getting Token1 from Uniswap pair %v: %v\n",paddr.String(),err)
				os.Exit(1)
			}
			var evt UPairCreated
			evt.Token0 = token0
			evt.Token1 = token1
			evt.Pair = paddr
			evt.PairSeq = big.NewInt(int64(i))
			Info.Printf("scan: inserting pair %v seq %v\n",evt.Pair.String(),i)
			storage.Insert_uniswap_pair_created_evt(&bci,&evt)
		}
	}
	Info.Printf("finished pair array scan\n")
}
/*DISCONTINUED
func fetch_and_store_erc20_info(token_addr common.Address) {
	// note: this func is called as goroutine for speed. however duplicate calls can occur,
	//		which are prevented with DO NOTHING on conflict in the INSERT query
	found,_ := storage.Get_ERC20Info(token_addr.String())
	if found {
		return
	}
	erc20_info,err := Fetch_erc20_info(eclient,&token_addr)
	if err != nil {
		Error.Printf("Couldn't fetch ERC20 token info for addr %v : %v\n",token_addr.String(),err)
		return
	}
	erc20_info.Address = token_addr.String()
	storage.Insert_ERC20Info(&erc20_info)
}*/
func fetch_and_store_erc20_info(token_addr common.Address) (int,error) {
	// note: this func is called as goroutine for speed. however duplicate calls can occur,
	//		which are prevented with DO NOTHING on conflict in the INSERT query
	found,info := storage.Get_ERC20Info(token_addr.String())
	if found {
		return info.Decimals,nil
	}
	erc20_info,err := Fetch_erc20_info(eclient,&token_addr)
	if err != nil {
		Error.Printf("Couldn't fetch ERC20 token info for addr %v : %v\n",token_addr.String(),err)
		return 0,err
	}
	erc20_info.Address = token_addr.String()
	storage.Insert_ERC20Info(&erc20_info)
	return erc20_info.Decimals,nil
}
func build_list_of_inspected_events(factory,router1,router2 string) []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0,8)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	PAIR_CREATED,
			ContractAid: storage.Lookup_or_create_address(factory,0,0),
		},
		InspectedEvent {
			Signature: PAIR_SWAP,
			ContractAid: 0,
		},
	)
	return inspected_events
}
func get_event_ids(from_evt_id,to_evt_id int64) []int64 {
	output := make([]int64 ,0,1024)
	for _,e := range inspected_events {
		var event_list []int64
		if e.ContractAid > 0 {
			event_list = storage.Get_evtlogs_by_signature_in_range(
				e.Signature,fmt.Sprintf("%v",e.ContractAid),from_evt_id,to_evt_id,
			)
		} else {
			event_list = storage.Get_evtlogs_by_signature_only_in_range(
				e.Signature,from_evt_id,to_evt_id,
			)
		}
		/*Info.Printf("selected events for signature %v:\n",e.Signature)
		for _,evt_id := range event_list {
			Info.Printf("\tEvtId:\t%9v\n",evt_id)
		}*/
		output = append(output,event_list...)
	}
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	num_elts:=remove_duplicates(output)
	return output[0:num_elts]
}
func remove_duplicates(nums []int64) int {
	if len(nums) == 0 {
		return 0
	}
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
func find_requester(tx_id int64,amount *big.Int) *common.Address {
	
	Info.Printf("finding requester for tx_id=%v, amount=%v\n",tx_id,amount.String())
	raw_logs := storage.Find_uniswap_transfer_events(tx_id)
	for _,rlog  := range raw_logs {
		var log types.Log
		rlp.DecodeBytes(rlog,&log)
		var transf ETransfer
		err := pair_abi.UnpackIntoInterface(&transf,"Transfer",log.Data)
		if err != nil {
			Error.Printf("Error unpacking Transfer event to find swap requester: %v\n",err)
			os.Exit(1)
		}
		transf.From= common.BytesToAddress(log.Topics[1][12:])
		transf.To= common.BytesToAddress(log.Topics[2][12:])
		Info.Printf("log.Address=%v\n",log.Address.String())
		transf.Dump(Info)
		if amount.Cmp(transf.Value) == 0 {
			addr := new(common.Address)
			addr.SetBytes(transf.From.Bytes())
			Info.Printf("found requester, addr=%v\n",addr.String())
			return addr
		} else {
			Info.Printf(
				"possible requester addr=%v and amount=%v don't match",
				transf.From.String(),transf.Value.String(),
			)
		}
	}
	Info.Printf("requester wasn't found, returning nil")
	return nil
}
func execute_event(e *EthereumEventLog,log *types.Log) error {
	tx_hash,_,timestamp,err := storage.Get_tx_hash_with_timestamp_by_id(e.TxId)
	if err != nil {
		Error.Printf("Couldn't get tx record from DB: %v\n",err)
		os.Exit(1)
	}
	Info.Printf("execute_evt: block %v log.Address= %v tx_hash= %v\n",e.BlockNum,log.Address.String(),tx_hash)
	if len(log.Topics) == 0 {
		Error.Printf("Event id=%v has no topics. Undefined behaviour, revision required\n",e.EvtId)
		os.Exit(1)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],evt_pair_created) {
		if len(log.Topics)!=3 {
			Error.Printf("PairCreated event not compliant log.Topics!=3. evtid=%v\n",e.EvtId)
			os.Exit(1)
		}
		var evt UPairCreated
		evt.Token0 = common.BytesToAddress(log.Topics[1][12:])
		evt.Token1 = common.BytesToAddress(log.Topics[2][12:])
		evt.Pair= common.BytesToAddress(log.Data[12:32])
		evt.PairSeq = big.NewInt(0)
		evt.PairSeq.SetBytes(log.Data[32:64])

		var bci BasicChainInfo
		bci.EvtId = e.EvtId
		bci.BlockNum = e.BlockNum
		bci.TxId = e.TxId
		bci.TimeStamp = timestamp
		if !storage.Pair_exists(evt.Pair.String()) {
			Info.Printf("Inserting PairCreated with address %v (pseq=%v)\n",evt.Pair.String(),evt.PairSeq.Int64())
			storage.Delete_uniswap_pair_created_evt(e.EvtId)
			storage.Insert_uniswap_pair_created_evt(&bci,&evt)
		}
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],evt_pair_swap) {
		if len(log.Topics)!=3 {
			Error.Printf("PairSwap event not compliant log.Topics!=3. evtid=%v\n",e.EvtId)
			os.Exit(1)
		}
		var evt UPairSwap
		evt.Sender = common.BytesToAddress(log.Topics[1][12:])
		evt.To = common.BytesToAddress(log.Topics[2][12:])
		err := pair_abi.UnpackIntoInterface(&evt,"Swap",log.Data)
		if err != nil {
			Error.Printf("Can't decode uniswap Swap event: %v\n",err)
			os.Exit(1)
		}
		var bci BasicChainInfo
		bci.EvtId = e.EvtId
		bci.BlockNum = e.BlockNum
		bci.TxId = e.TxId
		bci.TimeStamp = timestamp
		evt.Dump(Info)
		Info.Printf("Swap event on contract %v\n",log.Address.String())
		exists := storage.Pair_exists(log.Address.String())
		if !exists { // we are missing some PairCreated events because our data begins from 18 Jul 20
			Info.Printf(
				"Found Swap event that doesn't belongs to Unisvap V2 (contract %v )\n",
				log.Address.String(),
			)
		} else {
			Info.Printf("Inserting uniswap Swap event of Pair %v\n",log.Address.String())
			ptoks,err := storage.Get_pair_tokens(log.Address.String())
			if err != nil {
				Error.Printf("Error getting pair info for addr %v\n",log.Address.String())
				return err
			}
			evt.Decimals0,err = fetch_and_store_erc20_info(ptoks.Token0Addr)
			if err != nil {
				Error.Printf("Error getting decimals for addr %v\n",ptoks.Token0Addr.String())
				return err
			}
			evt.Decimals1,err = fetch_and_store_erc20_info(ptoks.Token1Addr)
			if err != nil {
				Error.Printf("Error getting decimals for addr %v\n",ptoks.Token1Addr.String())
				return err
			}
			storage.Delete_uniswap_pair_swap_evt(e.EvtId)
			storage.Insert_uniswap_pair_swap_evt(&log.Address,&bci,&evt)
			go update_uniswap_slippage_if_applicable(bci.BlockNum,log.Address.String())
		}
	}
	Info.Printf("\n\n")
	return nil
}
func process_uniswap_event(evt_id int64,exit_chan chan bool) error {

	select {
		case exit_flag := <-exit_chan:
			if exit_flag {
				Info.Println("Exiting by user request.")
				os.Exit(0)
			}
		default:
	}
	Info.Printf("Processing event id=%v\n",evt_id)
	evtlog := storage.Get_event_log(evt_id)
	var log types.Log
	rlp.DecodeBytes(evtlog.RlpLog,&log)
	return execute_event(&evtlog,&log)
}
func process_uniswap(exit_chan chan bool) {

	status := storage.Get_uniswap_status()
	var max_batch_size int64 = 1024*64
	for {
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
		last_chain_evt_id,err := storage.Get_last_evtlog_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_chain_evt_id {
			id_upper_limit = last_chain_evt_id
		}
		events := get_event_ids(status.LastEvtId,id_upper_limit)
		for _,evt_id := range events {
			err := process_uniswap_event(evt_id,exit_chan)
			if err != nil {
				Error.Printf(
					"Can't process swap event, cant get decimals : %v\n",err,
				)
				time.Sleep(1 * time.Second)
				break
			}
			status.LastEvtId=evt_id
			storage.Update_uniswap_status(&status)
		}
		if len(events) == 0 {
			status.LastEvtId = id_upper_limit
			storage.Update_uniswap_status(&status)
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {


	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/uniswap_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/uniswap_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/uniswap_error.log",log_dir)
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
	storage = Connect_to_storage(Info)
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

	abi_parsed := strings.NewReader(UniswapV2FactoryABI)
	factory_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse Uniswap Factory ABI: %v\n",err)
		os.Exit(1)
	}
	abi_parsed = strings.NewReader(UniswapV2PairABI)
	pair_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse Uniswap Pair ABI: %v\n",err)
		os.Exit(1)
	}

	var router1,router2 string
	factory_addr_str,router1,router2 = storage.Get_uniswap_contracts()
	inspected_events = build_list_of_inspected_events(factory_addr_str,router1,router2)
	update_existing_pairs()
	process_uniswap(exit_chan)
}
