package main

import (
	"os"
	"os/signal"
	"syscall"
	"bytes"
	"time"
	"fmt"
	"log"
	"strings"
	"math/big"
	"context"
	"encoding/hex"
	//"encoding/json"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	ERC20_TRANSFER = "ddf252ad"
	BALANCER_SWAP = "908fb5ee"
	AUGUR_FOUNDRY_WRAPPER_CREATED = "7dcd5c80"
)
var (
	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
	Error   *log.Logger
	Info	*log.Logger
	market_order_id int64 = 0
	cash_abi *abi.ABI
	af_abi abi.ABI
	all_contracts map[string]interface{}
	caddrs *ContractAddresses
)
func proc_erc20_transfer(log *types.Log,agtx *AugurTx,evtlog_id int64) {
	var mevt ETransfer
	if len(log.Topics)!=3 {
		Info.Printf("ERC20 transfer event is not compliant log.Topics!=3. Tx hash=%v\n",log.TxHash.String())
		return
	}
	mevt.From= common.BytesToAddress(log.Topics[1][12:])
	mevt.To= common.BytesToAddress(log.Topics[2][12:])
	start := time.Now()
	err := cash_abi.Unpack(&mevt,"Transfer",log.Data)
	duration := time.Since(start)
	Info.Printf("BENCH cash_abi.Unpack() took %v micrsec\n",duration.Microseconds())
	if err != nil {
		Error.Printf("signature=%v\n",log.Topics[0].String())
		Error.Printf("address=%v\n",log.Address.String())
		Error.Printf("tx hash = %v\n",log.TxHash.String())
		Error.Printf("log.Data=%v, data len=%v\n",hex.EncodeToString(log.Data[:]),len(log.Data[:]))
		Error.Printf("Event ERC20_Transfer, decode error: %v",err)
	} else {
		Info.Printf("ERC20_Transfer event, contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump(Info)
		if bytes.Equal(caddrs.Dai.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.Delete_DAI_transfer_by_evtlog_id(evtlog_id)	// prevention against duplication
			storage.Process_DAI_token_transfer(&mevt,caddrs,agtx,evtlog_id)
		}
		if bytes.Equal(caddrs.Reputation.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.Delete_REP_transfer_by_evtlog_id(evtlog_id) // prevention against duplication
			storage.Process_REP_token_transfer(&mevt,agtx,evtlog_id)
		}
	}
}
func process_erc20_tokens(contract_aids string) {

	for {
		status := storage.Get_tok_process_status()
		start1 := time.Now()
		tok_events := storage.Get_evt_log_ids_by_signature(ERC20_TRANSFER,contract_aids,status.LastEvtId,256)
		duration1 := time.Since(start1)
		Info.Printf("BENCH Get_token_transfers_batch() took %v milliseconds\n",duration1.Milliseconds())
		for _,evt := range tok_events {
			Info.Printf("event = %+v\n",evt)
			evtlog := storage.Get_event_log(evt.EvtId)
			var log types.Log
			rlp.DecodeBytes(evtlog.RlpLog,&log)
			log.Address.SetBytes(caddrs.Dai.Bytes())
			agtx := storage.Get_augur_transaction(evt.TxId)
			proc_erc20_transfer(&log,agtx,evt.EvtId)
			status.LastEvtId = evt.EvtId
			storage.Update_tok_process_status(&status)
		}
		Info.Printf("processed %v events\n",len(tok_events))
		if len(tok_events) == 0 {
			return
		}
	}
}
func process_afoundry_wrapper_created_events() {
	
	var copts = new(bind.CallOpts)
	af_addr := storage.Get_augur_foundry_contract_addr()
	af_contract_aid,err := storage.Nonfatal_lookup_address_id(af_addr)
	if err != nil {
		return
	}
	status := storage.Get_augur_foundry_status()
	ids_str := fmt.Sprintf("%v",af_contract_aid)
	events := storage.Get_evt_log_ids_by_signature(AUGUR_FOUNDRY_WRAPPER_CREATED,ids_str,status.LastEvtId,256)
	for _,e := range events {
		Info.Printf("Augur foundry wrapper created: %v",e)
		evtlog := storage.Get_event_log(e.EvtId)
		var log types.Log
		rlp.DecodeBytes(evtlog.RlpLog,&log)
		var evt EAugurFoundryWrapperCreated
		evt.TokenId = big.NewInt(0)
		evt.TokenId.SetBytes(log.Topics[1][:])
		err := af_abi.Unpack(&evt,"WrapperCreated",log.Data)
		if err != nil {
			Error.Printf("Error decoding WrapperCreated event: %v\n",err)
			os.Exit(1)
		}
		agtx := storage.Get_augur_transaction(e.TxId)
		_,bnum,err := storage.Get_tx_hash_by_id(agtx.TxId)
		if err != nil {
			Error.Printf("Error getting tx object by id (tx_id=%v): %v\n",e.TxId,err)
			os.Exit(1)
		}
		timestamp,err := storage.Get_block_timestamp(bnum)
		if err != nil {
			Error.Printf("Errror getting block's timestamp: %v\n",err)
			os.Exit(1)
		}
		wctrct,err := NewERC20Wrapper(evt.TokenAddress,eclient)
		if err!=nil {
			Error.Printf("Cant create contract instance for ERC20 Wrapper %v: ",evt.TokenAddress.String(),err)
			os.Exit(1)
		}
		decimals,err := wctrct.Decimals(copts)
		if err != nil {
			Error.Printf("Decimals() for %v failed: %v\n",evt.TokenAddress.String(),err)
			os.Exit(1)
		}
		symbol,err := wctrct.Symbol(copts)
		if err != nil {
			Error.Printf("Symbol() for %v failed: %v\n",evt.TokenAddress.String(),err)
			os.Exit(1)
		}
		name,err := wctrct.Name(copts)
		if err != nil {
			Error.Printf("Name() for %v failed: %v\n",evt.TokenAddress.String(),err)
			os.Exit(1)
		}
		tokid_bytes := evt.TokenId.Bytes()
		market_addr,outcome_idx := Unpack_sharetoken_id(tokid_bytes[0:21])
		Info.Printf("market_addr=%v\n",market_addr.String())
		storage.Insert_wrapper_created_evt(e.EvtId,timestamp,agtx,&evt,name,symbol,int(decimals),&market_addr,outcome_idx)
		status.LastEvtId = e.EvtId
		storage.Update_augur_foundry_status(&status)
	}
}
func process_single_erc20wrapped_transfer_evt(wrapper_aid int64,decimals int,evt *ShortEvtLog) {

	evtlog := storage.Get_event_log(evt.EvtId)
	var log types.Log
	rlp.DecodeBytes(evtlog.RlpLog,&log)
	agtx := storage.Get_augur_transaction(evt.TxId)

	var transfer ETransfer
	if len(log.Topics)!=3 {
		Info.Printf("ERC20 transfer event is not compliant log.Topics!=3. EvtId=%v\n",evt.EvtId)
		return
	}
	transfer.From= common.BytesToAddress(log.Topics[1][12:])
	transfer.To= common.BytesToAddress(log.Topics[2][12:])
	err := cash_abi.Unpack(&transfer,"Transfer",log.Data)
	if err !=  nil {
		Error.Printf("Can't unpack Transfer event: %v\n",err)
		os.Exit(1)
	}

	var t WShTokTransfer
	t.EvtLogId = evt.EvtId
	t.WrapperAid = wrapper_aid
	t.BlockNum = agtx.BlockNum
	t.TxId = agtx.TxId
	t.From = transfer.From.String()
	t.To = transfer.To.String()
	t.AmountStr = transfer.Value.String()
	storage.Insert_augur_foundry_transfer_evt(&t,decimals)
}
func process_single_erc20wrapped_contract(c *ERC20ShTokContract) {
	Info.Printf("Querying ERC20 transfer events for contract_id=%v\n",c.WrapperAid)
	events := storage.Get_evt_log_ids_by_signature(
		ERC20_TRANSFER,fmt.Sprintf("%v",c.WrapperAid),c.LastEvtId,512,
	)
	Info.Printf("processing wrapper id=%v , %v events\n",c.WrapperAid,len(events))
	for _,evt := range events {
		process_single_erc20wrapped_transfer_evt(c.WrapperAid,c.Decimals,&evt)
		storage.Update_wrapped_token_event_id_status(c.WrapperAid,evt.EvtId)
	}
}
func process_erc20wrapped_sharetokens() {

	wrp_ctrcts := storage.Get_erc20wrapped_sharetoken_contracts()
	for _,c := range wrp_ctrcts {
		process_single_erc20wrapped_contract(&c)
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
		dai_contract_aid:= storage.Lookup_or_create_address(caddrs.Dai.String(),0,0)
		rep_contract_aid := storage.Lookup_or_create_address(caddrs.Reputation.String(),0,0)
		_=dai_contract_aid
		_=rep_contract_aid
	//	process_erc20_tokens(fmt.Sprintf("%v,%v",dai_contract_aid,rep_contract_aid))
		process_afoundry_wrapper_created_events()
		process_erc20wrapped_sharetokens()
		time.Sleep(1 * time.Second)
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/tokens_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/tokens_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/tokens_error.log",log_dir)
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

	caddrs_obj,err := storage.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs = &caddrs_obj

	all_contracts = Load_all_artifacts("./abis/augur-artifacts-abi.json")

	cash_abi = Abi_from_artifacts(&all_contracts,"Cash")

	abi_parsed := strings.NewReader(AugurFoundryABI)
	af_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse Augur Foundry ABI: %v\n",err)
		os.Exit(1)
	}

	process_tokens(exit_chan,&caddrs_obj)
}
