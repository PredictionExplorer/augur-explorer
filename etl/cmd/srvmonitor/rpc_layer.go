package main

import (
	"os"
	"time"
	"fmt"
	"sync"
	"math"
	"context"
	"github.com/nsf/termbox-go"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
)
func calculate_official_flag(name string) bool {

	if name == os.Getenv("OFFICIAL_RPC_ARBITRUM") {
		return true
	}
	if name == os.Getenv("OFFICIAL_RPC_MAINNET") {
		return true
	}
	if name == os.Getenv("OFFICIAL_RPC_SEPOLIA") {
		return true
	}
	if name == os.Getenv("OFFICIAL_RPC_SEPOLIA_ARB") {
		return true
	}
	return false
}
func init_rpc_status_struct(s *RPCStatus,name string,url string,chain_id string,x int,y int) {
	s.RPCName = name
	s.RPCUrl = url
	s.X = x
	s.Y = y
	s.IsOfficial = calculate_official_flag(name)
	s.ChainId=chain_id
	if s.IsOfficial {
		if os.Getenv("OFFICIAL_RPC_MAINNET") == name {
			Official_mainnet_ptr = s
		}
		if os.Getenv("OFFICIAL_RPC_ARBITRUM") == name {
			Official_arbitrum_ptr = s
		}
		if os.Getenv("OFFICIAL_RPC_SEPOLIA_ARB") == name {
			Official_sepolia_arb_ptr = s
		}
	}
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
	status.LastBlockNum = latestBlock2.Number.Int64()
	time.Sleep(2*time.Second)	// this sleep is required to sync all parallel RPC calls that are made by all go-routines, because we calculate difference of blocks against official RPCs
	diff := latestBlock2.Number.Int64() - latestBlock1.Number.Int64()
	if diff == 0 {
		status.ErrStr=fmt.Sprintf("Block difference is zero (last block = %v)",latestBlock2.Number.Int64())
		update_global_errors(status.ErrStr)
	} else {
		status.Alive = true
	}
	if status.ChainId == "42161" {
		if Official_arbitrum_ptr != nil {
			if Official_arbitrum_ptr.LastBlockNum != 0 {
				status.OfficialLagDiff = Official_arbitrum_ptr.LastBlockNum - status.LastBlockNum
			} else {
			}
		} else {
			status.OfficialLagDiff = math.MaxInt64
		}
	}
	if status.ChainId == "421614" {
		if Official_sepolia_arb_ptr != nil {
			if Official_sepolia_arb_ptr.LastBlockNum != 0 {
				status.OfficialLagDiff = Official_sepolia_arb_ptr.LastBlockNum - status.LastBlockNum
			} else {
			}
		} else {
			status.OfficialLagDiff = math.MaxInt64
		}
	}
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
	var official_diff string = "------"
	if status.OfficialLagDiff != math.MaxInt64 {
		official_diff = fmt.Sprintf("%6v",status.OfficialLagDiff)
	}
	if status.IsOfficial { // it is not official rpc itself
		official_diff = fmt.Sprintf("%6s","N/A")
	}
	printAtPosition(status.X+80,status.Y,official_diff,termbox.ColorBlue,termbox.ColorDefault)
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
