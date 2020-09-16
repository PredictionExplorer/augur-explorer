// Augur ETL: Converts Augur Data to SQL database

package main

import (
	"os"
	"os/signal"
	"syscall"
	"io/ioutil"
	"strings"
	"time"
	"strconv"
	"fmt"
	"context"
	"log"
	"math/big"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	// ToDo: get these signatures from the abi files (after our code stabilizes, currently we will
	//	leave these constants visible to aid debugging processes)
	MARKET_CREATED = "ea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1"
	MARKET_OI_CHANGED = "213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268"
	MARKET_ORDER = "9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e"
	MARKET_FINALIZED = "6d39632c2dc10305bf5771cfff4af1851f07c03ea27b821cad382466bdf7a21f"
	INITIAL_REPORT_SUBMITTED = "c3ebb227c22e7644e9bef8822009f746a72c86f239760124d67fdc2c302b3115"
	MARKET_VOLUME_CHANGED_V1 = "e9f0af820300e73bae76c8e76943abe7fbb4224b49cb133e2dadc6f71acf6370"
	MARKET_VOLUME_CHANGED_V2 = "cc7cd5af4aead9d3a4a968c74d9063653dccf7110c5ced93fa86b8b03ef5ca24"
	DISPUTE_CROWDSOURCER_CONTRIBUTION = "e7f47639cdf56ec6c5451df334b73c9ca5cccd20da2c0f4e390e9bb71a6f672a"
	TOKENS_TRANSFERRED = "3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf"
	TOKEN_BALANCE_CHANGED = "63fd58f559b73fc4da5511c341ec8a7b31c5c48538ef83c6077712b6edf5f7cb"
	SHARE_TOKEN_BALANCE_CHANGED = "350ea32dc29530b9557420816d743c436f8397086f98c96292138edd69e01cb3"
	CANCEL_0X_ORDER = "be80e5687d7095071b7c4e7a56e0e67bfb9e8a39352f1690fdf74c1ee935c75e"
	TRANSFER_BATCH = "4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
	TRANSFER_SINGLE = "c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	PROFIT_LOSS_CHANGED = "59543b7f82735782aa5bdb97dff40ff288d4548a5865da513b40e4088e2ee77e"
	ERC20_TRANSFER = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	EXCHANGE_FILL = "6869791f0a34781b29882982cc39e882768cf2c96995c2a110c577c53bc932d5"
	TRADING_PROCEEDS_CLAIMED = "95366b7f64c6bb45149f9f7c522403fceebe5170ff76b8ffde2b0ab943ac11ce"
	ERC1155_APPROVAL_FOR_ALL = "17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
	ERC20_APPROVAL = "8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	EXEC_TX_STATUS = "ee9c28a7fe7177d351e891cb4ca5b7a4e4aba4974be67fb7665ba1ad0e703439"
	RELAYHUB_TRANSACTION_RELAYED = "ab74390d395916d9e0006298d47938a5def5d367054dcca78fa6ec84381f3f22"
	REGISTER_CONTRACT = "a037dd0e01f0488a530cb17065a6d2f284fae016004fc744ee2a41d5cacf85d5"
	UNIVERSE_CREATED = "e36b09d83f9cfa88c37f071fc2cfb5ff30b764cbd98088e70d965573c9ce5bbd"
	SET_REFERRER = "a18a7bfc"

	DEFAULT_WAIT_TIME = 5000	// 5 seconds
	DEFAULT_DB_LOG				= "db.log"
	//DEFAULT_LOG_DIR				= "ae_logs"
	MAX_APPROVAL_BASE10 string = "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	NUM_AUGUR_CONTRACTS int = 35
)
var (
	// these evt_ variables are here for speed to avoid calculation of Keccak256
	//		on each bytes.Compare() operation
	evt_market_created,_ = hex.DecodeString(MARKET_CREATED)
	evt_market_oi_changed,_ = hex.DecodeString(MARKET_OI_CHANGED)
	evt_market_order,_ = hex.DecodeString(MARKET_ORDER)
	evt_market_finalized,_ = hex.DecodeString(MARKET_FINALIZED)
	evt_initial_report_submitted,_ = hex.DecodeString(INITIAL_REPORT_SUBMITTED)
	evt_market_volume_changed_v1,_ = hex.DecodeString(MARKET_VOLUME_CHANGED_V1)
	evt_market_volume_changed_v2,_ = hex.DecodeString(MARKET_VOLUME_CHANGED_V2)
	evt_dispute_crowd_contrib,_ = hex.DecodeString(DISPUTE_CROWDSOURCER_CONTRIBUTION)
	evt_tokens_transferred,_ = hex.DecodeString(TOKENS_TRANSFERRED)
	evt_token_balance_changed,_ = hex.DecodeString(TOKEN_BALANCE_CHANGED)
	evt_share_token_balance_changed,_ = hex.DecodeString(SHARE_TOKEN_BALANCE_CHANGED)
	evt_cancel_0x_order,_ = hex.DecodeString(CANCEL_0X_ORDER)
	evt_transfer_batch,_ = hex.DecodeString(TRANSFER_BATCH)
	evt_transfer_single,_ = hex.DecodeString(TRANSFER_SINGLE)
	evt_profit_loss_changed,_ = hex.DecodeString(PROFIT_LOSS_CHANGED)
	evt_erc20_transfer,_ = hex.DecodeString(ERC20_TRANSFER)
	evt_exchange_fill,_ = hex.DecodeString(EXCHANGE_FILL)
	evt_trading_proceeds_claimed,_ = hex.DecodeString(TRADING_PROCEEDS_CLAIMED)
	evt_erc1155_approval_for_all,_ = hex.DecodeString(ERC1155_APPROVAL_FOR_ALL)
	evt_erc20_approval,_ = hex.DecodeString(ERC20_APPROVAL)
	evt_execute_tx_status,_ = hex.DecodeString(EXEC_TX_STATUS)
	evt_tx_relayed,_ = hex.DecodeString(RELAYHUB_TRANSACTION_RELAYED)
	evt_register_contract,_ = hex.DecodeString(REGISTER_CONTRACT)
	evt_universe_created,_ = hex.DecodeString(UNIVERSE_CREATED)
	sig_set_referrer,_ = hex.DecodeString(SET_REFERRER)

	exec_wtx_sig ,_ = hex.DecodeString("78dc0eed")

	storage *SQLStorage

	all_contracts map[string]interface{}
	inspected_events [][]byte

	augur_abi *abi.ABI
	trading_abi *abi.ABI
	zerox_trade_abi *abi.ABI
	cash_abi *abi.ABI
	exchange_abi *abi.ABI
	wallet_abi *abi.ABI

	ctrct_wallet_registry *AugurWalletRegistry
	ctrct_zerox_trade *ZeroX
	ctrct_dai_token *DAICash
	ctrct_pl *ProfitLoss

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	eclient *ethclient.Client
	rpcclient *rpc.Client

	// addresses of the contracts used in our code (for making eth.Call()s if needed)
	caddrs *ContractAddresses

	market_order_id int64 = 0
	owner_fld_offset int64 = int64(OWNER_FIELD_OFFSET)	// offset to AugurContract::owner field obtained with eth_getStorage()
	initial_amount *big.Int = nil	// Initial order amount extracted from MakerAssetData of Fill event
	
	set_back_block_num int64 = 0

	position_changes	[]*PosChg	// used to track changes in positions for debugging/verification

	Error   *log.Logger
	Info	*log.Logger

	//DISCONTINUED ErrChainSplit error = errors.New("Chainsplit detected")
	split_simulated bool = false

	max_approval *big.Int = big.NewInt(0)

)
type rpcBlockHash struct {
	Hash		string
}
func read_block_numbers(fname string)  []int64 {
	data,err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Can't open file %v containing comma-separated block numbers to be processed\n")
		os.Exit(1)
	}
	blocks_str := string(data)
	numbers := strings.Split(blocks_str,",")
	output := make([]int64,0,512)
	for i:=0 ; i<len(numbers); i++ {
		trimmed:=strings.ReplaceAll(numbers[i],"\n","")
		bnum,err:=strconv.Atoi(trimmed)
		if err!=nil {
			fmt.Printf("Can't convert block %v to number: %v . Aborting\n",numbers[i],err)
			os.Exit(1)
		}
		output = append(output,int64(bnum))
	}
	return output
}
func main() {
	//client, err := ethclient.Dial("http://:::8545")

	var block_numbers []int64
	stop_block := int(0)
	if len(os.Args) > 1 {
		var err error
		stop_block,err=strconv.Atoi(os.Args[1])
		if err != nil {
			// must be file number specifying block numbers to process
			block_numbers = read_block_numbers(os.Args[1])
		}
	}
	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
	}

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/etl_%v",log_dir,DEFAULT_DB_LOG)

	position_changes = make([]*PosChg,0,8)

	fname:=fmt.Sprintf("%v/etl_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/etl_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	max_approval.SetString(MAX_APPROVAL_BASE10,10)

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	ctx := context.Background()
	stored_chain_id := storage.Get_stored_chain_id()
	network_chain_id,err :=eclient.NetworkID(ctx)
	if err != nil {
		Fatalf("Can't get Network ID: %v\n",err)
	}
	if stored_chain_id != network_chain_id.Int64() {
		if stored_chain_id == 0 {
			// not initialized yet
			storage.Set_chain_id(network_chain_id.Int64())
		} else {
			Fatalf(
				"Network chain_id = %v , my chain_id = %v. Mismatch, exiting",
				network_chain_id.Int64(),stored_chain_id,
			)
		}
	}

	caddrs_obj,err := storage.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs=&caddrs_obj

	net_caddrs,err := Get_contract_addresses_from_net(caddrs_obj.AugurTrading,eclient)
	if err != nil {
		Fatalf("Can't get contract addresses from Ethereum Network: %v",err)
	}
	num_mismatches,match_errors := Contract_addresses_match(caddrs,&net_caddrs)
	if num_mismatches > 0 {
		if num_mismatches >= (NUM_AUGUR_CONTRACTS - 4) { // -1 for AugurTrading , -1 for AccountLoader
			Info.Printf("Empty contract addresses found, populating...")
			storage.Update_contract_addresses(&net_caddrs)
			Info.Printf("Please restart")
			os.Exit(0)
		} else {
			Error.Printf("%v contract addresses mismatch, errors: %v\n",num_mismatches,match_errors)
			Info.Printf("Exiting due to contract address mismatch.")
			os.Exit(1)
		}
	}
	augur_init(caddrs,&all_contracts)

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after block processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	//go balance_updater()	// updates DAI token balances very 10 seconds

	if len(block_numbers) > 0 {
		for i:=0 ; i<len(block_numbers); i++ {
			bnum := block_numbers[i]
			err := process_block(bnum,false,true)
			if err!=nil {
				fmt.Printf("Process failed: %v. Repeat again.\n",err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	}

  main_loop:
	latestBlock, err := eclient.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	}

	bnum,exists := storage.Get_last_block_num()
	if !exists {
		bnum = 0
	} else {
		bnum = bnum + 1
	}
	var bnum_high int64 = latestBlock.Number().Int64()
	if bnum_high < bnum {
		Info.Printf("Database has more blocks than the blockchain, aborting. Fix last_block table.\n")
		os.Exit(1)
	}
	if stop_block > 0 {
		Info.Printf("Will exit at block %v for debugging\n",stop_block)
		bnum_high = int64(stop_block)
	}
	for ; bnum<bnum_high; bnum++ {
		//block_hash:=common.HexToHash(block_hash_str)
		for {
			select {
				case exit_flag := <-exit_chan:
					if exit_flag {
						Info.Println("Exiting by user request.")
						os.Exit(0)
					}
				default:
			}
			proc_open_orders()
			err := process_block(bnum,true,false)
			if err==nil {
				break
			} else {
				// this is probably happening due to RPC unavailability, so we use a delay
				time.Sleep(1 * time.Second)
				if err == ErrChainSplit {
					bnum = set_back_block_num
					continue
				}
				Error.Printf("Block processing error: %v\n",err)
			}
		}
		//scan_profit_loss_data_for_debugging(bnum,&position_changes)
		position_changes=nil
	}// for block_num
	latestBlock, err = eclient.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	} else {
		if latestBlock.Number().Int64() >= bnum {
			time.Sleep(DEFAULT_WAIT_TIME * time.Millisecond)
		}
	}
	if stop_block == 0 {
		goto main_loop
	}
}
