package main

import (
	"os"
	"os/signal"
	"syscall"
	"sort"
	"time"
	"fmt"
	"context"
	"log"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG				= "db.log"

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
	EXECUTE_WALLET_TRANSACTION = "78dc0eed"
	TRADE = "2f562016"
	CLAIM_PROCEEDS = "db754422"
	VALIDITY_BOND_CHANGED = "69af68e366a0570364e3a086f3b5ac79f08ecc3f93eaccbfcf3864809b12b5d8"
	NOSHOW_BOND_CHANGED = "d1fc3f2cb1387e602db0e6f8f22649df65df5246eeff281cf6d1ef62feda4ece"
	DISPUTE_CROWDSOURCER_CREATED = "f9a0b30bcf861874bf36630742f0d56b22648898d7cdd0cd785d74acd17e0d44"
	DISPUTE_WINDOW_CREATED = "97f8b399e255f30d56b759b645c86652624ee258937579ff4a747abaeae857c4"
	DESIGNATED_REPORT_STAKE_CHANGED = "9c75a088fcb0527d67a80a7d0a5006bbabe02f4b23984234ae68b2b146f001bc"
	INITIAL_REPORTER_REDEEMED = "3ffffb51f92f91faf4ba8c906f5a0180d1033be93b1e227cd92c872dc234fdf0"
	COMPLETE_SETS_PURCHASED = "fe06587917de7df83a446bcbb889cee699d7fc35b7b53e263282c2acb5a16499"
	COMPLETE_SETS_SOLD = "dd7dcfa6708112395eb94e9b1889295fb19af21ef290e918256838c979b2dfbd"
)
var (
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
	exec_wtx_sig ,_ = hex.DecodeString(EXECUTE_WALLET_TRANSACTION)
	trade_sig,_ = hex.DecodeString(TRADE)
	claim_proceeds_sig,_ = hex.DecodeString(CLAIM_PROCEEDS)
	evt_validity_bond_changed,_ = hex.DecodeString(VALIDITY_BOND_CHANGED)
	evt_noshow_bond_changed,_ = hex.DecodeString(NOSHOW_BOND_CHANGED)
	evt_dispute_crowdsourcer_created,_ = hex.DecodeString(DISPUTE_CROWDSOURCER_CREATED)
	evt_dispute_window_created,_ = hex.DecodeString(DISPUTE_WINDOW_CREATED)
	evt_designated_report_stake_changed,_ = hex.DecodeString(DESIGNATED_REPORT_STAKE_CHANGED)
	evt_complete_sets_purchased,_ = hex.DecodeString(COMPLETE_SETS_PURCHASED)
	evt_complete_sets_sold,_ = hex.DecodeString(COMPLETE_SETS_SOLD)
	evt_initial_reporter_redeemed,_ = hex.DecodeString(INITIAL_REPORTER_REDEEMED)

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	Error   *log.Logger
	Info	*log.Logger
	market_order_id int64 = 0
	all_contracts map[string]interface{}
	caddrs *ContractAddresses
	inspected_events []InspectedEvent

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

	eclient *ethclient.Client
	rpcclient *rpc.Client

	owner_fld_offset int64 = int64(OWNER_FIELD_OFFSET)	// offset to AugurContract::owner field obtained with eth_getStorage()
)
func get_event_ids(from_tx_id,to_tx_id int64) []int64 {
	output := make([]int64 ,0,1024)
	for _,e := range inspected_events {
		event_list := storage.Get_tx_ids_from_evt_logs_by_signature(
			e.Signature,e.ContractAid,from_tx_id,to_tx_id,
		)
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

func process_augur_trading_events(exit_chan chan bool,caddrs *ContractAddresses) {

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
		status := storage.Get_augur_process_status()
		id_upper_limit := status.LastTxId + max_batch_size
		last_tx_id,err := storage.Get_last_tx_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly transaction table is empty")
			os.Exit(1)
		}
		if id_upper_limit > last_tx_id {
			id_upper_limit = last_tx_id
		}
		events := get_event_ids(status.LastTxId,id_upper_limit)
		for _,tx_id := range events {
			err := process_transaction(tx_id)
			if err != nil {
				Info.Printf("Error processing tx with tx_id=%v : %v\n",tx_id,err)
				os.Exit(1)
			}
			status.LastTxId=tx_id
			storage.Update_augur_process_status(&status)
		}
		if len(events) == 0 {
			status.LastTxId = id_upper_limit
			storage.Update_augur_process_status(&status)
		}
		proc_open_orders()
		time.Sleep(1 * time.Second)
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/augur_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/augur_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/augur_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
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

	caddrs_obj,err := storage.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs = &caddrs_obj

	all_contracts = Load_all_artifacts("./abis/augur-artifacts-abi.json")

	augur_init(caddrs,&all_contracts)

	inspected_events = build_list_of_inspected_events()
	process_augur_trading_events(exit_chan,&caddrs_obj)
}
