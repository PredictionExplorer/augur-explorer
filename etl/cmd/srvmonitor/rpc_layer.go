package main

import (
	"time"
	"fmt"
	"sync"
	"context"
	"github.com/nsf/termbox-go"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
		update_global_errors(status.ErrStr)
		wg.Done()
		return
	}
	eclient := ethclient.NewClient(rpc_obj)
	latestBlock1, err := eclient.HeaderByNumber(context.Background(), nil)
    if err != nil {
		status.ErrStr = err.Error()
		update_global_errors(status.ErrStr)
		wg.Done()
		return
    }
	time.Sleep(WAIT_RPC_BLOCK_NUM*time.Second)
	latestBlock2, err := eclient.HeaderByNumber(context.Background(), nil)
    if err != nil {
		status.ErrStr = err.Error()
		update_global_errors(status.ErrStr)
		wg.Done()
		return
    }
	diff := latestBlock2.Number.Int64() - latestBlock1.Number.Int64()
	if diff == 0 {
		status.ErrStr=fmt.Sprintf("Block difference is zero (last block = %v)",latestBlock2.Number.Int64())
		update_global_errors(status.ErrStr)
	} else {
		status.Alive = true
	}
	status.LastBlockNum = latestBlock2.Number.Int64()
	wg.Done()

}
func print_rpc_status_line(status *RPCStatus) {

	printAtPosition(status.X,status.Y,status.RPCName,termbox.ColorWhite,termbox.ColorDefault)
	printAtPosition(status.X+25,status.Y,status.RPCUrl,termbox.ColorWhite,termbox.ColorDefault)
	alive_str := string("Alive")
	if !status.Alive  {
		alive_str = "DOWN "
		printAtPosition(status.X+60,status.Y,alive_str,termbox.ColorRed,termbox.ColorDefault)
	} else {
		printAtPosition(status.X+60,status.Y,alive_str,termbox.ColorGreen,termbox.ColorDefault)
	}
	printAtPosition(status.X+70,status.Y,fmt.Sprintf("%v",status.LastBlockNum),termbox.ColorBlue,termbox.ColorDefault)
	if len(status.ErrStr) > 0 {
		update_global_errors(status.ErrStr)
	}
}
func print_current_rpc_status() {
	printAtPosition(0, 0, "--------------------- RPC Nodes ------------------------------",termbox.ColorWhite,termbox.ColorDefault)
	print_rpc_status_line(&rpc0)
	print_rpc_status_line(&rpc1)
	print_rpc_status_line(&rpc2)
	print_rpc_status_line(&rpc3)
	print_rpc_status_line(&rpc4)
	print_rpc_status_line(&rpc5)
	print_rpc_status_line(&rpc6)
	print_rpc_status_line(&rpc7)
	print_rpc_status_line(&rpc8)
	termbox.Flush()
}
