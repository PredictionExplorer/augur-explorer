package main

import (
	"os"
	"errors"
	"time"
	"fmt"
	"log"
	"encoding/hex"
	"flag"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/dbs/balancerv2"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	ERC20_TRANSFER = "ddf252ad"
)
var (
	storagew				SQLStorageWrapper
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
	Error   *log.Logger
	Info	*log.Logger
	BalancesLog	*log.Logger
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
func fetch_and_store_erc20_info(token_addr common.Address,token_aid int64) error {
	_,already_checked := info_checked[token_addr]
	if already_checked {
		return nil
	}
	info_checked[token_addr] = struct{}{}
	_,is_bad := bad_erc20_token[token_addr]
	if is_bad {
		storagew.Insert_bad_erc20_token_mark(token_aid)
		return nil
	}
	Info.Printf("Getting ERC20 info\n")
	found,_ := storagew.S.Get_ERC20Info_v3(token_addr.String())
	Info.Printf("found=%v\n",found)
	if found {
		Info.Printf("Aid %v already exists\n",token_aid)
		return nil
	}
	erc20_info,err := Fetch_erc20_info(eclient,&token_addr)
	if err != nil {
		Info.Printf("Couldn't fetch ERC20 token info for addr %v : %v\n",token_addr.String(),err)
		bad_erc20_token[token_addr]=struct{}{}
		return err
	}
	erc20_info.Address = token_addr.String()
	storagew.S.Insert_ERC20Info_v3("erc20_info",&erc20_info)
	return nil
}
func process_token(tok_aid int64) {

	addr_str,err := storagew.S.Layer1_lookup_address(tok_aid)
	if err != nil {
		Info.Printf("Can't lookup address with aid=%v : %v\n",tok_aid,err)
		os.Exit(1)
	}

	addr := common.HexToAddress(addr_str)
	fetch_and_store_erc20_info(addr,tok_aid)
}
func main() {

	usage_str := fmt.Sprintf("usage: %v --schema [schema_name]\n",os.Args[0])
	if len(os.Args)<3 {
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	schema_name := flag.String("schema", "", "Schema name")
	flag.Parse()
	if len(*schema_name) < 3 {
		fmt.Printf("Schema name must be larger than 2 characters\n")
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	var err error
	RPC_URL = os.Getenv("RPC_URL")
	eclient, err = ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storagew.S = Connect_to_storage_with_schema(Info,*schema_name)
	Info.Printf("Schema name: %v\n",*schema_name)

	bad_erc20_token = make(map[common.Address]struct{})
	bad_for_decode = make(map[common.Address]struct{})
	info_checked = make(map[common.Address]struct{})
  next:
	last_block := storagew.Get_erc20_info_status_last_block()
	Info.Printf("Last block number = %v, starting scan\n",last_block)
	tokens,next_block := storagew.Get_token_aids_from_swaps(last_block)
	Info.Printf("Tokens fetched: %v , next_block=%v\n",len(tokens),next_block)
	for i:=0;i<len(tokens);i++ {
		aid := tokens[i]
		Info.Printf("Processing token aid=%v\n",aid)
		process_token(aid)
	}
	if next_block !=0 {
		storagew.Update_erc20_status_last_block(next_block)
		last_block=next_block
		Info.Printf("Set next_block to %v\n",next_block)
	} else {
		Info.Printf("next block stays the same\n")
	}
	time.Sleep(1*time.Second)
	goto next
}
