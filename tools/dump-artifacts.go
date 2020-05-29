// Augur ETL: Converts Augur Data to SQL database

package main

import (
	"bytes"
	"sort"
	"fmt"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
)
const (
)
var (
	/*
	// these evt_ variables are here for speed to avoid calculation of Keccak256
	//		on each bytes.Compare() operation
	evt_market_created,_ = hex.DecodeString(MARKET_CREATED)
	evt_market_oi_changed,_ = hex.DecodeString(MARKET_OI_CHANGED)
	evt_market_order,_ = hex.DecodeString(MARKET_ORDER)
	evt_market_finalized,_ = hex.DecodeString(MARKET_FINALIZED)
	evt_initial_report_submitted,_ = hex.DecodeString(INITIAL_REPORT_SUBMITTED)
	evt_market_volume_changed,_ = hex.DecodeString(MARKET_VOLUME_CHANGED)
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
	evt_zerox_approval_for_all,_ = hex.DecodeString(ZEROX_APPROVAL_FOR_ALL)
	evt_erc20_approval,_ = hex.DecodeString(ERC20_APPROVAL)

	storage *SQLStorage
*/
	all_contracts map[string]interface{}
//	inspected_events [][]byte
/*
	augur_abi *abi.ABI
	trading_abi *abi.ABI
	zerox_abi *abi.ABI
	cash_abi *abi.ABI
	exchange_abi *abi.ABI
	wallet_abi *abi.ABI

	ctrct_wallet_registry *AugurWalletRegistry
	ctrct_zerox *ZeroX

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	rpcclient *ethclient.Client

	// addresses of the contracts used in our code (for making eth.Call()s if needed)
	dai_addr common.Address
	zerox_addr common.Address


	market_order_id int64 = 0
	*/
)
func check(e error) {
	if e != nil {
		panic(fmt.Sprintf("Exiting Augur extractor with error: %v",e))
	}
}
func dump_abi_events(contract_name string,a *abi.ABI) {

	fmt.Printf("Events:\n")
	for evt:=range a.Events {
		fmt.Printf("\teee::%v::%v::%v\n",contract_name,a.Events[evt].ID().String(),evt)
	}

}
func dump_abi_methods(contract_name string,a *abi.ABI) {
	fmt.Printf("Methods:\n")
	for meth := range a.Methods {
		fmt.Printf("\tmmm::%v\t%v::%v\n",contract_name,hex.EncodeToString(a.Methods[meth].ID()),meth)
	}
}
func dump_all_artifacts() {

	names := make([]string, 0, len(all_contracts))
	for k := range all_contracts {
		names = append(names, k)
	}
	sort.Strings(names)
	for i:=0; i<len(names); i++ {
		contract_name := names[i]
		fmt.Printf("Contract: %v\n",contract_name)
		abi:=abi_from_artifacts(contract_name)
		dump_abi_events(contract_name,abi)
		dump_abi_methods(contract_name,abi)
	}
}
func load_all_artifacts(filename string) map[string]interface{} {

	abi_data, err := ioutil.ReadFile("./abis/augur-artifacts-abi.json")
	check(err)
	all_abis_rdr := bytes.NewReader(abi_data)
	check(err)
	byte_data, err := ioutil.ReadAll(all_abis_rdr)
	check(err)
	var all_contracts map[string]interface{}
	json.Unmarshal([]byte(byte_data), &all_contracts)
	return all_contracts
}
func abi_from_artifacts(contract string) *abi.ABI {

	contract_abi:=all_contracts[contract]
	contract_bytes, _ := json.Marshal(contract_abi) // convert back to JSON so Ethereum package can work
	rdr := bytes.NewReader(contract_bytes)
	ctrct_abi,err := abi.JSON(rdr)
	check(err)
	return &ctrct_abi
}
func load_abi(fname string) *abi.ABI {

	abi_data, err := ioutil.ReadFile(fname)
	check(err)
	abi_rdr := bytes.NewReader(abi_data)
	check(err)
	abi,err := abi.JSON(abi_rdr)
	check(err)
	return &abi
}
func main() {

	//all_contracts = load_all_artifacts("./abis/augur-artifacts-abi.json")
	all_contracts = load_all_artifacts("./artifacts-list-before-may26.txt")
	dump_all_artifacts()

}
