package main

import (
	"fmt"
	"log"
	"os"
    "os/signal"
    "syscall"
	"sync"
	"path/filepath"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DELETER_LOG          = "deleter.log"
	NUM_BLOCKS_NO_DELETE	= 1000000	// leaves this number of blocks after head of the chain
	NUM_BLOCKS_TO_SCAN	= 10000
	NUM_WORKERS				= 10
)
var (
	Info                    *log.Logger
	Error                   *log.Logger
	storage 				*SQLStorage
	wg 						sync.WaitGroup
	in_statement			string
)
func process_single_block(bnum int64) {
	defer wg.Done()
	num_tx_to_our_contracts := storage.Get_deleter_count_non_deleteable_transactions_by_tx_to(bnum,in_statement)
	if (num_tx_to_our_contracts > 0) {
		Info.Printf("Block %v can't be deleted (tx.To used)\n",bnum)
		return
	}
	num_events_our_contracts := storage.Get_deleter_count_non_deleteable_transactions_by_events_emitted(bnum,in_statement)
	if (num_events_our_contracts > 0) {
		Info.Printf("Block %v can't be deleted (events used)\n",bnum)
		return
	}
	// block must be deleted (it contains none of our data)
	storage.Deleter_do_delete_block_transactions(bnum)
}
func main() {
	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	log_file_name := filepath.Join(log_dir, "deleter.log")
	logfile, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Err at log file creation: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	storage = Connect_to_storage(Info)
	storage.Log_msg("Log initialized\n")

	contracts:=storage.Get_deleter_contracts()
	first_block_num := storage.Get_first_block_num()
	last_block_num,_ := storage.Get_last_block_num()
	c := make(chan os.Signal,1)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}() 
	if len(contracts) == 0 {
		fmt.Printf("No contracts were registered in 'd_status' table, exiting\n")
		os.Exit(1)
	}
	in_statement = fmt.Sprintf("%d",contracts[0].ContractAid)
	for i:=1; i<len(contracts); i++ {
		in_statement = fmt.Sprintf("%s,%s",in_statement,fmt.Sprintf("%d",contracts[i].ContractAid))
	}
	cur_block_num := storage.Get_deleter_status()
	if cur_block_num < first_block_num { cur_block_num = first_block_num; }
	Info.Printf("first_block_bum=%v cur_block_num=%v, limit block num =%v\n",first_block_num,cur_block_num,(last_block_num-NUM_BLOCKS_NO_DELETE))

	for {
		if (cur_block_num  > (last_block_num - NUM_BLOCKS_NO_DELETE)) {
			Info.Printf("Skipping delete due to NUM_BLOCKS_NO_DELETE condition\n")
			os.Exit(1)
		}
		select {
			case exit_flag := <-exit_chan:
			if exit_flag {
				Info.Println("Exiting by user request.")
				os.Exit(0)
			}
			default:
		}
		for i:=int64(0);i<NUM_WORKERS;i++ {
			wg.Add(1)
			go process_single_block(cur_block_num + i)
		}
		wg.Wait()
		cur_block_num = cur_block_num + NUM_WORKERS
		storage.Update_deleter_status(cur_block_num)
	}
}
