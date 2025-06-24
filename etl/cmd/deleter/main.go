package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"math/big"
    "os/signal"
    "syscall"
	"context"
	"path/filepath"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DELETER_LOG          = "deleter.log"
	NUM_BLOCKS_NO_DELETE	= 1000000
)
var (
	eclient 				*ethclient.Client
    rpcclient *rpc.Client
	Info                    *log.Logger
	Error                   *log.Logger
	storage 				*SQLStorage
	RPC_URL                  = os.Getenv("RPC_URL")
)
func main() {
	
	var err error

	// Set up logging
	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	log_file := filepath.Join(log_dir, "deleter.log")
	Info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	rpcclient, err = rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	// Initialize storage
	storage := Connect_to_storage(Info)
	storage.Init_log(log_file)
	storage.Log_msg("Log initialized\n")

	statuses:=storage.Get_deleter_status()

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}() 
	if len(statuses) == 0 {
		fmt.Printf("No contracts were registered in 'd_status' table, exiting\n")
		os.Exit(1)
	}

	for {
		select {
			case exit_flag := <-exit_chan:
			if exit_flag {
				Info.Println("Exiting by user request.")
				os.Exit(0)
			}
			default:
		}

		for i:=0; i<len(statuses); i++ {
			last_block_num := statuses[i].LastBlockNum
			filter :=  ethereum.FilterQuery{}
			filter.FromBlock = big.NewInt(last_block_num)
			filter.ToBlock = big.NewInt(last_block_num)
			filter.Addresses = []common.Address{statuses[i].ContractEthAddr}
			logs,err := eclient.FilterLogs(context.Background(),filter)
			if err != nil {
				Info.Printf("Error at ethclient.Filter(): %v\n",err)
				fmt.Printf("Error at ethclient.Filter(): %v\n",err)
				os.Exit(1)
			}
			for j:=0; j<len(logs); j++ {
				delete_to_block := int64(logs[i].BlockNumber)
				for k:=last_block_num; k<delete_to_block; k++ {	// we delete 1 block per query since it is a time consuming op
					fmt.Printf("Deleting block %v (delete_to_block=%v)\n",k,delete_to_block)
				}
			}
		}

		time.Sleep(1 * time.Second) 
	}
}
