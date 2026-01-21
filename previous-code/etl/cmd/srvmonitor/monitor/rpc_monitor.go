package monitor

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
	
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/display"
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/types"
)

const (
	RPCBlockWait     = 60 // seconds to wait before second getBlock() call
	UpdateIntervalRPC = 60 // seconds between checks
)

// RPCStatus holds status for a single RPC node
type RPCStatus struct {
	Config          types.RPCConfig
	LastBlockNum    int64
	Alive           bool
	OfficialLagDiff int64
	ErrStr          string
	X, Y            int
}

// RPCMonitor monitors RPC nodes
type RPCMonitor struct {
	nodes              []types.RPCConfig
	statuses           []*RPCStatus
	officialMainnet    *RPCStatus
	officialArbitrum   *RPCStatus
	officialSepolia    *RPCStatus
	officialSepoliaArb *RPCStatus
	position           types.Position
	sharedState        *SharedRPCState
}

// NewRPCMonitor creates a new RPC monitor
func NewRPCMonitor(nodes []types.RPCConfig, officialNames map[string]string, sharedState *SharedRPCState) *RPCMonitor {
	m := &RPCMonitor{
		nodes:       nodes,
		statuses:    make([]*RPCStatus, len(nodes)),
		position:    types.Position{X: 0, Y: 0},
		sharedState: sharedState,
	}
	
	// Initialize statuses
	for i, node := range nodes {
		m.statuses[i] = &RPCStatus{
			Config:          node,
			OfficialLagDiff: math.MaxInt64,
			X:               1,
			Y:               i + 1,
		}
		
		// Track official nodes
		if node.IsOfficial {
			switch node.ChainID {
			case "1":
				if node.Name == officialNames["mainnet"] {
					m.officialMainnet = m.statuses[i]
				}
			case "11155111":
				if node.Name == officialNames["sepolia"] {
					m.officialSepolia = m.statuses[i]
				}
			case "42161":
				if node.Name == officialNames["arbitrum"] {
					m.officialArbitrum = m.statuses[i]
				}
			case "421614":
				if node.Name == officialNames["sepolia_arb"] {
					m.officialSepoliaArb = m.statuses[i]
				}
			}
		}
	}
	
	return m
}

// Name returns the monitor name
func (m *RPCMonitor) Name() string {
	return "RPC Monitor"
}

// GetDisplayPosition returns the display position
func (m *RPCMonitor) GetDisplayPosition() types.Position {
	return m.position
}

// Start begins monitoring
func (m *RPCMonitor) Start(ctx context.Context, disp display.Display, errorChan chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.check(disp, errorChan)
			time.Sleep(UpdateIntervalRPC * time.Second)
		}
	}
}

// check performs a check cycle
func (m *RPCMonitor) check(disp display.Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))
	
	for _, status := range m.statuses {
		go m.checkNode(status, &wg, errorChan)
	}
	
	wg.Wait()
	m.display(disp)
}

// checkNode checks a single RPC node
func (m *RPCMonitor) checkNode(status *RPCStatus, wg *sync.WaitGroup, errorChan chan<- string) {
	defer wg.Done()
	
	if status.Config.URL == "" {
		status.Config.URL = "*** not set ***"
		return
	}
	
	status.ErrStr = ""
	status.Alive = false
	
	rpcObj, err := rpc.DialContext(context.Background(), status.Config.URL)
	if err != nil {
		status.ErrStr = err.Error()
		errorChan <- status.ErrStr
		return
	}
	
	eclient := ethclient.NewClient(rpcObj)
	latestBlock1, err := eclient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		status.ErrStr = err.Error()
		errorChan <- status.ErrStr
		return
	}
	
	time.Sleep(RPCBlockWait * time.Second)
	
	latestBlock2, err := eclient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		status.ErrStr = err.Error()
		errorChan <- status.ErrStr
		return
	}
	
	status.LastBlockNum = latestBlock2.Number.Int64()
	time.Sleep(2 * time.Second) // sync all parallel RPC calls
	
	diff := latestBlock2.Number.Int64() - latestBlock1.Number.Int64()
	if diff == 0 {
		status.ErrStr = fmt.Sprintf("Block difference is zero (last block = %v)", latestBlock2.Number.Int64())
		errorChan <- status.ErrStr
	} else {
		status.Alive = true
	}
	
	// Update shared state if this is an official node
	if status.Config.IsOfficial && m.sharedState != nil {
		switch status.Config.ChainID {
		case "1":
			m.sharedState.UpdateOfficialMainnet(status.LastBlockNum)
		case "11155111":
			m.sharedState.UpdateOfficialSepolia(status.LastBlockNum)
		case "42161":
			m.sharedState.UpdateOfficialArbitrum(status.LastBlockNum)
		case "421614":
			m.sharedState.UpdateOfficialSepoliaArb(status.LastBlockNum)
		}
	}
	
	// Calculate lag against official RPC
	if status.Config.ChainID == "1" && m.officialMainnet != nil && !status.Config.IsOfficial {
		if m.officialMainnet.LastBlockNum != 0 {
			status.OfficialLagDiff = m.officialMainnet.LastBlockNum - status.LastBlockNum
		}
	}
	if status.Config.ChainID == "11155111" && m.officialSepolia != nil && !status.Config.IsOfficial {
		if m.officialSepolia.LastBlockNum != 0 {
			status.OfficialLagDiff = m.officialSepolia.LastBlockNum - status.LastBlockNum
		}
	}
	if status.Config.ChainID == "42161" && m.officialArbitrum != nil && !status.Config.IsOfficial {
		if m.officialArbitrum.LastBlockNum != 0 {
			status.OfficialLagDiff = m.officialArbitrum.LastBlockNum - status.LastBlockNum
		}
	}
	if status.Config.ChainID == "421614" && m.officialSepoliaArb != nil && !status.Config.IsOfficial {
		if m.officialSepoliaArb.LastBlockNum != 0 {
			status.OfficialLagDiff = m.officialSepoliaArb.LastBlockNum - status.LastBlockNum
		}
	}
}

// display renders the RPC status
func (m *RPCMonitor) display(disp display.Display) {
	// Header
	disp.DrawText(types.Position{X: 0, Y: 0},
		"--------------------- RPC Nodes ------------------------------",
		types.ColorWhite, types.ColorDefault)
	
	// Each RPC node
	for _, status := range m.statuses {
		y := status.Y
		
		// Status (first column)
		aliveStr := "Alive"
		color := types.ColorGreen
		if !status.Alive {
			aliveStr = "DOWN "
			color = types.ColorRed
		}
		disp.DrawText(types.Position{X: status.X, Y: y}, aliveStr, color, types.ColorDefault)
		
		// Block number (second column)
		disp.DrawText(types.Position{X: status.X + 10, Y: y}, fmt.Sprintf("%v", status.LastBlockNum),
			types.ColorBlue, types.ColorDefault)
		
		// Official lag (third column)
		officialDiff := "------"
		if status.OfficialLagDiff != math.MaxInt64 {
			officialDiff = fmt.Sprintf("%6v", status.OfficialLagDiff)
		}
		if status.Config.IsOfficial {
			officialDiff = fmt.Sprintf("%6s", "N/A")
		}
		disp.DrawText(types.Position{X: status.X + 20, Y: y}, officialDiff, types.ColorBlue, types.ColorDefault)
		
		// Name (fourth column - shifted right)
		disp.DrawText(types.Position{X: status.X + 30, Y: y}, status.Config.Name, types.ColorWhite, types.ColorDefault)
		
		// URL (fifth column - shifted right)
		disp.DrawText(types.Position{X: status.X + 55, Y: y}, status.Config.URL, types.ColorWhite, types.ColorDefault)
	}
	
	disp.Flush()
}

