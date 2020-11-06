package main

import (
	"os"
	"os/signal"
	"syscall"
	"bytes"
	"strings"
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
//	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
//	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	ERC20_TRANSFER				= "ddf252ad"
	BALANCER_NEW_POOL			= "8ccec77b"
	BALANCER_JOIN				= "63982df1"
	BALANCER_EXIT				= "e74c9155"
	BALANCER_SWAP				= "908fb5ee"
	AUGUR_FOUNDRY_WRAPPER_CREATED = "7dcd5c80"
	SIG_BALANCER_SET_SWAP_FEE	= "34e19907"
	SIG_BALANCER_SET_CONTROLLER = "92eefe9b"
	SIG_BALANCER_SET_PUBLIC_SWAP = "49b59552"
	SIG_BALANCER_FINALIZE		= "4bb278f3"
	SIG_BALANCER_BIND			= "e4e1e538"
	SIG_BALANCER_REBIND			= "3fdddaa2"
	SIG_BALANCER_UNBIND			= "cf5e7bd3"
	SIG_BALANCER_GULP			= "8c28cbe8"
)
var (
	evt_balancer_new_pool,_ = hex.DecodeString(BALANCER_NEW_POOL)
	evt_balancer_join,_ = hex.DecodeString(BALANCER_JOIN)
	evt_balancer_exit,_ = hex.DecodeString(BALANCER_EXIT)
	evt_balancer_swap,_ = hex.DecodeString(BALANCER_SWAP)
	b_balancer_set_swap_fee,_ = hex.DecodeString(SIG_BALANCER_SET_SWAP_FEE)
	b_balancer_set_controller,_ = hex.DecodeString(SIG_BALANCER_SET_CONTROLLER)
	b_balancer_set_public_swap,_ = hex.DecodeString(SIG_BALANCER_SET_PUBLIC_SWAP)
	b_balancer_finalize,_ = hex.DecodeString(SIG_BALANCER_FINALIZE)
	b_balancer_bind,_ = hex.DecodeString(SIG_BALANCER_BIND)
	b_balancer_rebind,_ = hex.DecodeString(SIG_BALANCER_REBIND)
	b_balancer_unbind,_ = hex.DecodeString(SIG_BALANCER_UNBIND)
	b_balancer_gulp,_ = hex.DecodeString(SIG_BALANCER_GULP)

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	Error   *log.Logger
	Info	*log.Logger
	market_order_id int64 = 0
	cash_abi *abi.ABI
	all_contracts map[string]interface{}
	caddrs *ContractAddresses
	inspected_events []InspectedEvent
	inspected_funcs []InspectedEvent
	bpool_abi abi.ABI

)
func build_list_of_inspected_events(pool,factory,exchange string) []InspectedEvent {

	// this is the list of all the events we read (not necesarilly insert into the DB, but check on them)
	inspected_events= make([]InspectedEvent,0,8)
	inspected_events = append(inspected_events,
		InspectedEvent {
			Signature:	hex.EncodeToString(evt_balancer_new_pool[:4]),
			ContractAid: storage.Lookup_or_create_address(factory,0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_balancer_join[:4]),
			ContractAid: storage.Lookup_or_create_address(pool,0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_balancer_exit[:4]),
			ContractAid: storage.Lookup_or_create_address(pool,0,0),
		},
		InspectedEvent {
			Signature: hex.EncodeToString(evt_balancer_swap[:4]),
			ContractAid: storage.Lookup_or_create_address(exchange,0,0),
		},
	)
	return inspected_events
}
func build_list_of_inspected_functions() []InspectedEvent {
	signatures := make([]InspectedEvent,0,8)
	signatures = append(signatures,
		InspectedEvent {
			Signature:  SIG_BALANCER_SET_SWAP_FEE,
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: SIG_BALANCER_SET_CONTROLLER,
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: SIG_BALANCER_SET_PUBLIC_SWAP,
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: SIG_BALANCER_FINALIZE,
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: SIG_BALANCER_BIND,
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: SIG_BALANCER_REBIND,
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: SIG_BALANCER_UNBIND,
			ContractAid: 0,
		},
		InspectedEvent {
			Signature: SIG_BALANCER_GULP,
			ContractAid: 0,
		},
	)
	return signatures
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
	for _,f := range inspected_funcs {
		event_list := storage.Get_LOG_CALL_evtlogs(
			f.Signature,from_evt_id,to_evt_id,
		)
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
func execute_event(e *EthereumEventLog,log *types.Log) {
	tx_hash,_,err := storage.Get_tx_hash_by_id(e.TxId)
	if err != nil {
		Error.Printf("Couldn't get tx record from DB: %v\n",err)
		os.Exit(1)
	}
	Info.Printf("execute_evt: block %v log.Address=%v tx_hash=%v\n",e.BlockNum,log.Address.String(),tx_hash)
	timestamp,err := storage.Get_block_timestamp(e.BlockNum)
	if err != nil {
		Error.Printf("Can't get block's timestamp (evtid=%v): %v\n",e.EvtId,err)
		os.Exit(1)
	}
	if len(log.Topics) == 0 {
		Error.Printf("Event id=%v has no topics. Undefined behaviour, revision required\n",e.EvtId)
		os.Exit(1)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],evt_balancer_new_pool) {
		if len(log.Topics)!=3 {
			Error.Printf("LOG_NEW_POOL event not compliant log.Topics!=3. evtid=%v\n",e.EvtId)
			os.Exit(1)
		}

		caller := common.BytesToAddress(log.Topics[1][12:])
		pool := common.BytesToAddress(log.Topics[2][12:])
		var evt BalancerNewPool
		evt.EvtId = e.EvtId
		evt.BlockNum = e.BlockNum
		evt.TxId = e.TxId
		evt.TimeStamp = timestamp
		evt.PoolAddr = pool.String()
		evt.CallerAddr = caller.String()
		Info.Printf("Insertint NEW_POOL with address %v\n",pool.String())
		storage.Delete_balancer_pool_created_evt(evt.EvtId)
		storage.Insert_balancer_pool_created_evt(&evt)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],evt_balancer_join) {
		if len(log.Topics)!=3 {
			Error.Printf("LOG_JOIN event not compliant log.Topics!=3. evtid=%v\n",e.EvtId)
			os.Exit(1)
		}
		caller := common.BytesToAddress(log.Topics[1][12:])
		token_in := common.BytesToAddress(log.Topics[2][12:])

		var joinevt ELOG_JOIN
		err := bpool_abi.Unpack(&joinevt,"LOG_JOIN",log.Data)
		if err != nil {
			Error.Printf("Event LOG_JOIN, decode error: %v",err)
			os.Exit(1)
		}
		var evt BalancerJoin
		evt.EvtId = e.EvtId
		evt.BlockNum = e.BlockNum
		evt.TxId = e.TxId
		evt.TimeStamp = timestamp
		evt.PoolAddr = log.Address.String()
		evt.CallerAddr = caller.String()
		evt.TokenInAddr = token_in.String()
		evt.AmountIn = joinevt.TokenAmountIn.String()
		Info.Printf("Inserting pool JOIN event of Holder %v to pool %v\n",caller.String(),evt.PoolAddr)
		storage.Delete_balancer_join_evt(evt.EvtId)
		storage.Insert_balancer_pool_join_evt(&evt)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],evt_balancer_exit) {
		if len(log.Topics)!=3 {
			Error.Printf("LOG_EXIT event not compliant log.Topics!=3. evtid=%v\n",e.EvtId)
			os.Exit(1)
		}
		caller := common.BytesToAddress(log.Topics[1][12:])
		token_out := common.BytesToAddress(log.Topics[2][12:])

		var exitevt ELOG_EXIT
		err := bpool_abi.Unpack(&exitevt,"LOG_EXIT",log.Data)
		if err != nil {
			Error.Printf("Event LOG_EXIT, decode error: %v",err)
			os.Exit(1)
		}
		var evt BalancerExit
		evt.EvtId = e.EvtId
		evt.BlockNum = e.BlockNum
		evt.TxId = e.TxId
		evt.TimeStamp = timestamp
		evt.PoolAddr = log.Address.String()
		evt.CallerAddr = caller.String()
		evt.TokenOutAddr = token_out.String()
		evt.AmountOut = exitevt.TokenAmountOut.String()
		Info.Printf("Inserting pool EXIT event of Holder %v from pool %v\n",caller.String(),evt.PoolAddr)
		storage.Delete_balancer_exit_evt(evt.EvtId)
		storage.Insert_balancer_pool_exit_evt(&evt)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],evt_balancer_swap) {
		if len(log.Topics)!=4 {
			Error.Printf("LOG_SWAP event not compliant log.Topics!=4. evtid=%v\n",e.EvtId)
			os.Exit(1)
		}

		var swapevt ELOG_SWAP
		err := bpool_abi.Unpack(&swapevt,"LOG_SWAP",log.Data)
		if err != nil {
			Error.Printf("Event LOG_SWAP, decode error: %v",err)
			os.Exit(1)
		}

		caller := common.BytesToAddress(log.Topics[1][12:])
		token_in := common.BytesToAddress(log.Topics[2][12:])
		token_out := common.BytesToAddress(log.Topics[3][12:])
		var evt BalancerSwap
		evt.EvtId = e.EvtId
		evt.BlockNum = e.BlockNum
		evt.TxId = e.TxId
		evt.TimeStamp = timestamp
		evt.PoolAddr = log.Address.String()
		evt.CallerAddr = caller.String()
		evt.TokenInAddr = token_in.String()
		evt.TokenOutAddr = token_out.String()
		evt.AmountIn = swapevt.TokenAmountIn.String()
		evt.AmountOut = swapevt.TokenAmountOut.String()
		Info.Printf("Inserting pool SWAP event of caller %v from pool %v\n",caller.String(),evt.PoolAddr)
		storage.Delete_balancer_swap_evt(evt.EvtId)
		storage.Insert_balancer_swap_evt(&evt)
	}
	// following events are Balancer's LOG_CALL events, but they are anonymous, so function signature in
	//		Topic[0] isn't present. Instead, Topic[0] is 'Sig' , Topic[1] is 'Caller'
	//		and log.Data[0] is 'data'. Since Topic[0] is the signature, we use it for bytes.Equal() instead
	//		of using tx.Input[0:4]
	if bytes.Equal(log.Topics[0].Bytes()[:4],b_balancer_set_swap_fee) {
		if !storage.Is_address_balancer_pool(log.Address.String()) {
			return
		}
		//Info.Printf("setfee: Data= %v\n",hex.EncodeToString(log.Data))
		var offset int = 32+32+4// first 32 - big.Int size; second 32 - length of Input; 4 - signature size
		fee := big.NewInt(0)
		fee.SetBytes(log.Data[offset+0:offset+32])
		//Info.Printf("Set fee with caller %v for pool %v\n",caller.String(),log.Address.String())
		var f SetSwapFee
		f.EvtId = e.EvtId
		f.BlockNum = e.BlockNum
		f.TxId = e.TxId
		f.TimeStamp = timestamp
		f.PoolAddr = log.Address.String()
		f.FeeStr = fee.String()
		storage.Delete_set_swap_fee(e.EvtId)
		storage.Insert_set_swap_fee(&f)
	}
	Info.Printf("before set_controller log.Address=%v topicsig=%v\n",log.Address.String(),hex.EncodeToString(log.Topics[0].Bytes()))
	if bytes.Equal(log.Topics[0].Bytes()[:4],b_balancer_set_controller) {
		if !storage.Is_address_balancer_pool(log.Address.String()) {
			return
		}
		//Info.Printf("controller: Data= %v\n",hex.EncodeToString(log.Data))
		var offset int = 32+32+4// first 32 - big.Int size; second 32 - length of Input; 4 - signature size
		controller := common.BytesToAddress(log.Data[offset+12:offset+32])
		var c SetController
		c.EvtId = e.EvtId
		c.BlockNum = e.BlockNum
		c.TxId = e.TxId
		c.TimeStamp = timestamp
		c.PoolAddr = log.Address.String()
		c.ControllerAddr = controller.String()
		//Info.Printf("set_controller for %v to %v\n",log.Address.String(),c.ControllerAddr)
		storage.Delete_set_controller(e.EvtId)
		storage.Insert_set_controller(&c)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],b_balancer_set_public_swap) {
		
		if !storage.Is_address_balancer_pool(log.Address.String()) {
			return
		}
		//Info.Printf("public: Data= %v\n",hex.EncodeToString(log.Data))
		var offset int = 32+32+4// first 32 - big.Int size; second 32 - length of Input; 4 - signature size
		var p SetPublic
		p.EvtId = e.EvtId
		p.BlockNum = e.BlockNum
		p.TxId = e.TxId
		p.TimeStamp = timestamp
		p.PoolAddr = log.Address.String()
		if log.Data[offset+31] == 1 {
			p.Public = true
		} else {
			p.Public = false
		}
		storage.Delete_set_public(e.EvtId)
		storage.Insert_set_public(&p)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],b_balancer_finalize) {
		if !storage.Is_address_balancer_pool(log.Address.String()) {
			return
		}
		//Info.Printf("finalize: Data= %v\n",hex.EncodeToString(log.Data))
		var f Finalize 
		f.EvtId = e.EvtId
		f.BlockNum = e.BlockNum
		f.TxId = e.TxId
		f.TimeStamp = timestamp
		f.PoolAddr = log.Address.String()
		storage.Delete_pool_finalize(e.EvtId)
		storage.Insert_pool_finalize(&f)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],b_balancer_bind) {
		if !storage.Is_address_balancer_pool(log.Address.String()) {
			return
		}
		Info.Printf("bind: Data= %v\n",hex.EncodeToString(log.Data))
		var offset int = 32+32+4// first 32 - big.Int size; second 32 - length of Input; 4 - signature size
		token := common.BytesToAddress(log.Data[offset+12:offset+32])
		balance := big.NewInt(0)
		balance.SetBytes(log.Data[offset+32:offset+64])
		denorm := big.NewInt(0)
		denorm.SetBytes(log.Data[offset+64:offset+96])
		var b PoolBind
		b.EvtId = e.EvtId
		b.BlockNum = e.BlockNum
		b.TxId = e.TxId
		b.TimeStamp = timestamp
		b.PoolAddr = log.Address.String()
		b.TokenAddr = token.String()
		b.Balance = balance.String()
		b.Denorm = denorm.String()
		storage.Delete_pool_bind(e.EvtId)
		storage.Insert_pool_bind(&b)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],b_balancer_unbind) {
		if !storage.Is_address_balancer_pool(log.Address.String()) {
			return
		}
		//Info.Printf("unbind: Data= %v\n",hex.EncodeToString(log.Data))
		var offset int = 32+32+4// first 32 - big.Int size; second 32 - length of Input; 4 - signature size
		token := common.BytesToAddress(log.Data[offset+12:offset+32])
		var u PoolUnBind
		u.EvtId = e.EvtId
		u.BlockNum = e.BlockNum
		u.TxId = e.TxId
		u.TimeStamp = timestamp
		u.PoolAddr = log.Address.String()
		u.TokenAddr = token.String()
		storage.Delete_pool_unbind(e.EvtId)
		storage.Insert_pool_unbind(&u)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],b_balancer_rebind) {
		
		if !storage.Is_address_balancer_pool(log.Address.String()) {
			return
		}
		//Info.Printf("rebind: Data= %v\n",hex.EncodeToString(log.Data))
		var offset int = 32+32+4// first 32 - big.Int size; second 32 - length of Input; 4 - signature size
		token := common.BytesToAddress(log.Data[offset+12:offset+32])
		balance := big.NewInt(0)
		balance.SetBytes(log.Data[offset+32:offset+64])
		denorm := big.NewInt(0)
		denorm.SetBytes(log.Data[offset+64:offset+96])
		/*type ReBindInput struct {
			Token			common.Address
			Balance			*big.Int
			Denorm			*big.Int
		}
		var input_data ReBindInput
		method, err := bpool_abi.MethodById(b_balancer_rebind)
		if err != nil {
			Info.Printf("Can't find method for rebind: %v\n",err)
			os.Exit(1)
		}
		err = method.Inputs.Unpack(&input_data, log.Data[64+4:])
		if err != nil {
			Error.Printf("Input decode error: %v",err)
			os.Exit(1)
		}*/
		var r PoolReBind
		r.EvtId = e.EvtId
		r.BlockNum = e.BlockNum
		r.TxId = e.TxId
		r.TimeStamp = timestamp
		r.PoolAddr = log.Address.String()
		r.TokenAddr = token.String()
		r.Balance = balance.String()
		r.Denorm = denorm.String()
		storage.Delete_pool_rebind(e.EvtId)
		storage.Insert_pool_rebind(&r)
	}
	if bytes.Equal(log.Topics[0].Bytes()[:4],b_balancer_gulp) {
		if !storage.Is_address_balancer_pool(log.Address.String()) {
			return
		}

		//Info.Printf("gulp: Data= %v\n",hex.EncodeToString(log.Data))
		var offset int = 32+32+4// first 32 - big.Int size; second 32 - length of Input; 4 - signature size
		token := common.BytesToAddress(log.Data[offset+12:offset+32])
		var g PoolGulp
		g.EvtId = e.EvtId
		g.BlockNum = e.BlockNum
		g.TxId = e.TxId
		g.TimeStamp = timestamp
		g.PoolAddr = log.Address.String()
		g.TokenAddr = token.String()
		storage.Delete_pool_gulp(e.EvtId)
		storage.Insert_pool_gulp(&g)
	}
}
func process_balancer_event(evt_id int64) {

	Info.Printf("Processing event id=%v\n",evt_id)
	evtlog := storage.Get_event_log(evt_id)
	var log types.Log
	rlp.DecodeBytes(evtlog.RlpLog,&log)
	execute_event(&evtlog,&log)
}
func process_balancer(exit_chan chan bool) {

	status := storage.Get_balancer_status()
	var max_batch_size int64 = 1024*100
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
			process_balancer_event(evt_id)
			status.LastEvtId=evt_id
			storage.Update_balancer_status(&status)
		}
		if len(events) == 0 {
			status.LastEvtId = id_upper_limit
			storage.Update_balancer_status(&status)
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/balancer_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/balancer_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/balancer_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

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

	abi_parsed := strings.NewReader(BPoolABI)
	bpool_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse Balancer Pool ABI: %v\n",err)
		os.Exit(1)
	}

	pool_addr,factory_addr,xchg_addr := storage.Get_balancer_contracts()
	inspected_events = build_list_of_inspected_events(pool_addr,factory_addr,xchg_addr)
	inspected_funcs = build_list_of_inspected_functions()
	process_balancer(exit_chan)
}
