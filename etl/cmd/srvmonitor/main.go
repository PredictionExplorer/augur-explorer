// Used for monitoring our server installation

package main

import (
//	"net/http"
	"os"
	"fmt"
	"log"
	"time"
	//"errors"
//	"io/ioutil"
	"context"
	"sync"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nsf/termbox-go"

//	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
type RPCStatus struct {
	LastBlockNum		int64
	Alive				bool	// if there is block difference over last 60 seconds, node is alive
	RPCUrl				string
	RPCName				string
	ErrStr				string
	X					int
	Y					int
}
const (
	WAIT_RPC_BLOCK_NUM	= 60		// seconds to wait before second getBlock() call
)
var (
	Error   *log.Logger
	Info	*log.Logger
	storage *SQLStorage

	rpc0,rpc1,rpc2,rpc3,rpc4,rpc5,rpc6		RPCStatus
)
func printAtPosition(x, y int, text string, fg, bg termbox.Attribute) {
	for i, r := range text {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}
func init_rpc_status_struct(s *RPCStatus,name string,url string,x int,y int) {
	s.RPCName = name
	s.RPCUrl = url
	s.X = x
	s.Y = y
}
func check_rpc_status(status *RPCStatus, wg *sync.WaitGroup) {

	if len(status.RPCUrl) == 0 {
		status.RPCUrl = "*** not set ***"
		wg.Done()
		return
	}
	status.ErrStr = ""
	status.Alive = false
	rpc_obj, err:=rpc.DialContext(context.Background(), status.RPCUrl)
	if err != nil {
		status.ErrStr = err.Error()
	}
	eclient := ethclient.NewClient(rpc_obj)
	latestBlock1, err := eclient.HeaderByNumber(context.Background(), nil)
    if err != nil {
		status.ErrStr = err.Error()
    }
	time.Sleep(WAIT_RPC_BLOCK_NUM*time.Second)
	latestBlock2, err := eclient.HeaderByNumber(context.Background(), nil)
    if err != nil {
		status.ErrStr = err.Error()
    }
	diff := latestBlock2.Number.Int64() - latestBlock1.Number.Int64()
	fmt.Printf("diff for %v is %v\n",status.RPCName,diff)
	if diff == 0 {
		status.ErrStr=fmt.Sprintf("Block difference is zero (last block = %v)",latestBlock2.Number.Int64())
	} else {
		fmt.Printf("setting alive to true for %v\n",status.RPCName)
		status.Alive = true
	}
	status.LastBlockNum = latestBlock2.Number.Int64()
	wg.Done()

}
func print_rpc_status_line(status *RPCStatus) {

	printAtPosition(status.X,status.Y,status.RPCName,termbox.ColorGreen,termbox.ColorBlack)
	printAtPosition(status.X+20,status.Y,status.RPCUrl,termbox.ColorGreen,termbox.ColorBlack)
	alive_str := string("Alive")
	if !status.Alive  {
		alive_str = "DOWN"
	}
	printAtPosition(status.X+48,status.Y,alive_str,termbox.ColorGreen,termbox.ColorBlack)
	printAtPosition(status.X+60,status.Y,fmt.Sprintf("%v",status.LastBlockNum),termbox.ColorGreen,termbox.ColorBlack)
	printAtPosition(status.X+75,status.Y,fmt.Sprintf("%v",status.ErrStr),termbox.ColorGreen,termbox.ColorBlack)
}
func print_current_rpc_status() {
	print_rpc_status_line(&rpc0)
	print_rpc_status_line(&rpc1)
	termbox.Flush()
}
func check_rpc_services() {

	var wg_rpcs sync.WaitGroup
	wg_rpcs.Add(2);
	init_rpc_status_struct(&rpc0,os.Getenv("RPC0_NAME"),os.Getenv("RPC0_URL"),5,2)
	init_rpc_status_struct(&rpc1,os.Getenv("RPC1_NAME"),os.Getenv("RPC1_URL"),5,3)
	init_rpc_status_struct(&rpc2,os.Getenv("RPC2_NAME"),os.Getenv("RPC2_URL"),5,4)
	go check_rpc_status(&rpc0,&wg_rpcs); 
	go check_rpc_status(&rpc1,&wg_rpcs); 
	wg_rpcs.Wait() 
	print_current_rpc_status()
}
func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatalf("Failed to initialize termbox: %v", err)
	}
	defer termbox.Close()
	//Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	fmt.Printf("\n\n\n\n\n\n")

//	storage = Connect_to_storage(Info)
	check_rpc_services()
	printAtPosition(3,20,fmt.Sprintf("Press any key to exit"),termbox.ColorGreen,termbox.ColorBlack)
	termbox.PollEvent()
}
