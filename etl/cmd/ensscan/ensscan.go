package main

import (
	"os"
	"os/signal"
	"syscall"
	"context"
	"math/big"
	"fmt"
	"log"
	"encoding/hex"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	ENS_NEWOWNER				= "ce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82"

	ENS_NAME_REGISTERED1		= "ca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f"
	//	NameRegistered (
	//		string name,
	//		index_topic_1 bytes32 label,
	//		index_topic_2 address owner,
	//		uint256 cost,
	//		uint256 expires
	//	)

	ENS_NAME_REGISTERED2		= "b3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9"

	ENS_NAME_REGISTERED3		= "570313dae523ecb48b1176a4b60272e5ea7ec637f5b2d09983cbc4bf25e7e9e3"
	//	NameRegistered (
	//		index_topic_1 address _caller,
	//		index_topic_2 address _beneficiary,
	//		index_topic_3 bytes32 _labelHash,
	//		string _subdomain,
	//		uint256 _createdDate
	//	)

	ENS_ADDR_CHANGED			= "52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2"
	ENS_ADDRESS_CHANGED			= "65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752"
	NEW_RESOLVER				= "335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0"

)
var (
	evt_newowner,_ = hex.DecodeString(ENS_NEWOWNER)

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
	Error   *log.Logger
	Info	*log.Logger

	market_order_id int64 = 0

)
func initial_load_
func initial_load(block_num_limit int64) {

}
func process_name_registered_events() {

}
func main() {


	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/ensscan_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/ensscan_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/ensscan_error.log",log_dir)
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
/*
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
*/
	filter := ethereum.FilterQuery{}
	//filter.FromBlock = big.NewInt(11000000)
	filter.FromBlock = big.NewInt(11470000)
	filter.ToBlock = nil // latest
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_newowner)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	filter.Addresses = nil
	Info.Printf("Submitting filter logs query with signature %v\n",hex.EncodeToString(signature.Bytes()))
	Info.Printf("filter query = %+v\n",filter)
	logs,err := eclient.FilterLogs(context.Background(),filter)
	if err!= nil {
		Error.Printf("Error: %v\n",err)
		Info.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	Info.Printf("call error=%v, len logs=%v\n",err,len(logs))
	for i,log := range logs {
		Info.Printf("%v: log = %+v\n",i,log)
	}
}
