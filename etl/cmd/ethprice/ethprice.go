package main
import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"fmt"
	"math/big"
	"context"
	"encoding/hex"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG              = "db.log"
	SWAP_EVT=   "c42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"
	MAX_BLOCKS	int64 = 1024*1
	USDC1_ADDR	string = "0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8"
	USDC2_ADDR	string = "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"
	DAI_ADDR string		= "0xC2e9F25Be6257c210d7Adf0D4Cd6E3E881ba25f8"
	//UNISWAP_V3_CREATION_BLOCK int64 = 12369621
	UNISWAP_V3_CREATION_BLOCK int64 = 13000000
)
var (
	storage *SQLStorage
	Info    *log.Logger
	Error	*log.Logger
	RPC_URL string
	evt_swap,_   = hex.DecodeString(SWAP_EVT)
	market_order_id int64 = 0

	rpcclient *rpc.Client
	eclient *ethclient.Client

	uniswap_v3_abi *abi.ABI
)
func get_blocks_timestamp(block_hash common.Hash) int64 {

	hdr,err := eclient.HeaderByHash(context.Background(),block_hash)
	if err != nil {
		Info.Printf("Error in HeaderByHash(): %v\n",err)
		Error.Printf("Error in HeaderByHash(): %v\n",err)
		os.Exit(1)
	}
	return int64(hdr.Time)
}
func lookup_token_code(addr_str,tx_hash string) int8 {

	Info.Printf("looking up swap code for addr %v\n",addr_str)
	if USDC1_ADDR == addr_str { return 0 }
	if USDC2_ADDR == addr_str { return 0 }
	if DAI_ADDR == addr_str { return 1 }
	Info.Printf("Error: can't lookup token code, tx_hash = %v\n",tx_hash)
	Error.Printf("Error: can't lookup token code, tx_hash = %v\n",tx_hash)
	os.Exit(1)
	return -1
}
func process_log(log *types.Log) {

	Info.Printf("Tx %v , log #%v\n",log.TxHash.String(),log.Index)
	var evt UniswapV3PoolSwap
	evt.Sender = common.BytesToAddress(log.Topics[1][12:])
	evt.Recipient = common.BytesToAddress(log.Topics[2][12:])
	err := uniswap_v3_abi.UnpackIntoInterface(&evt,"Swap",log.Data)
	if err != nil {
		Info.Printf("Error, cant unpack log idx=%v for tx %v : %v\n",log.Index,log.TxHash.String(),err)
		Error.Printf("Error, cant unpack log idx=%v for tx %v : %v\n",log.Index,log.TxHash.String(),err)
		os.Exit(1)
	}
	var swap EthpriceSwap
	swap.TxHash = log.TxHash.String()
	swap.TimeStamp = get_blocks_timestamp(log.BlockHash)
	swap.BlockNum = int64(log.BlockNumber)
	swap.TxIdx = int32(log.TxIndex)
	swap.LogIdx = int32(log.Index)
	swap.TokenCode = lookup_token_code(log.Address.String(),swap.TxHash)
	swap.Sender = evt.Sender.String()
	swap.Recipient = evt.Recipient.String()
	swap.Amount0 = evt.Amount0.String()
	swap.Amount1 = evt.Amount1.String()
	swap.SqrtPrice = evt.SqrtPriceX96.String()
	swap.Liquidity = evt.Liquidity.String()
	swap.Tick = evt.Tick.String()
	Info.Printf("Uniswap V3 Swap {\n")
	Info.Printf("\tTxHash: %v\n",swap.TxHash)
	Info.Printf("\tTimestamp: %v\n",swap.TimeStamp)
	Info.Printf("\tBlockNum: %v\n",swap.BlockNum)
	Info.Printf("\tTxIndex: %v\n",swap.TxIdx)
	Info.Printf("\tLogIdx: %v\n",swap.LogIdx)
	Info.Printf("\tTokenCode: %v\n",swap.TokenCode)
	Info.Printf("\tSender: %v\n",swap.Sender)
	Info.Printf("\tRecipient: %v\n",swap.Recipient)
	Info.Printf("\tAmount0: %v\n",swap.Amount0)
	Info.Printf("\tAmount1: %v\n",swap.Amount1)
	Info.Printf("\tSqrtPrice: %v\n",swap.SqrtPrice)
	Info.Printf("\tLiquidity: %v\n",swap.SqrtPrice)
	Info.Printf("\tTick: %v\n",swap.Tick)
	Info.Printf("}\n")
	storage.Ethprice_insert_swap_event(&swap)
}
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v [schema_name]\n")
		os.Exit(1)
	}

	schema_name := os.Args[1]

	RPC_URL = os.Getenv("RPC_URL")

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/ethprice_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/ethprice_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime)

	fname=fmt.Sprintf("%v/ethprice_error.log",log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Error = log.New(logfile,"ERROR: ",log.Ltime)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")
	storage.Db_set_schema_name(schema_name)

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err!= nil {
		fmt.Printf("Error connecting to node: %v\n",err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)

	abi_parsed := strings.NewReader(UniswapV3PoolABI)
	abi,err := abi.JSON(abi_parsed)
	if err != nil {
		fmt.Printf("Can't parse RandomWalk ABI: %v\n",err)
		os.Exit(1)
	}
	uniswap_v3_abi = &abi

	ctx := context.Background()
	starting_block := storage.Ethprice_get_last_processed_block()

	if starting_block == 0 {
		starting_block = UNISWAP_V3_CREATION_BLOCK
	}

	latestBlock, err := eclient.BlockByNumber(ctx, nil)
	if err != nil {
		fmt.Printf("Error geting latest block: %v\n",err)
		os.Exit(1)
	}
	latest_bnum := latestBlock.Number().Int64()


	from_block := int64(starting_block)
	to_block := int64(0)
	filter := ethereum.FilterQuery{}
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_swap)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	addresses := make([]common.Address,0,4)
	addresses = append(addresses,common.HexToAddress(USDC1_ADDR))
	addresses = append(addresses,common.HexToAddress(USDC2_ADDR))
	addresses = append(addresses,common.HexToAddress(DAI_ADDR))
	filter.Addresses = addresses

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	for ;from_block<=latest_bnum;from_block=from_block+MAX_BLOCKS {
		to_block = from_block + MAX_BLOCKS
		filter.FromBlock = big.NewInt(from_block)
		filter.ToBlock = big.NewInt(to_block)
		logs,err := eclient.FilterLogs(context.Background(),filter)
		if err != nil {
			fmt.Printf("Error querying events: %v\n",err)
			os.Exit(1)
		}
		for _,log := range logs {
			if log.Removed {
				continue
			} else {
				process_log(&log)
			}
			select {
				case exit_flag := <-exit_chan:
					if exit_flag {
						Info.Println("Exiting by user request.")
						os.Exit(0)
					}
				default:
			}
		}
	}
}
