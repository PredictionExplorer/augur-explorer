package srvmonitor

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// RPCStatus holds status for a single RPC node.
type RPCStatus struct {
	Config          RPCConfig
	LastBlockNum    int64
	Alive           bool
	OfficialLagDiff int64
	ErrStr          string
	X, Y            int

	// readOK reports whether both head reads succeeded this cycle, i.e.
	// LastBlockNum is fresh rather than left over from an earlier cycle.
	readOK bool
}

// RPCMonitor checks that every configured RPC node is reachable and its head
// block advances between two reads one blockWait apart. Official nodes feed
// the shared state; other nodes report their lag behind the official node of
// the same chain.
type RPCMonitor struct {
	statuses    []*RPCStatus
	official    map[string]*RPCStatus // chain id -> official node status
	position    Position
	sharedState *SharedRPCState
	interval    time.Duration
	// blockWait pauses between the two head reads of one check. Tests
	// replace it to advance the fake chain instead of sleeping.
	blockWait func(ctx context.Context)
}

// NewRPCMonitor creates a new RPC monitor. officialNames maps the chain key
// ("mainnet", "arbitrum", "sepolia", "sepolia_arb") to the configured
// official node name for that chain.
func NewRPCMonitor(nodes []RPCConfig, officialNames map[string]string, sharedState *SharedRPCState, iv Intervals) *RPCMonitor {
	m := &RPCMonitor{
		statuses:    make([]*RPCStatus, len(nodes)),
		official:    make(map[string]*RPCStatus),
		position:    Position{X: 0, Y: 0},
		sharedState: sharedState,
		interval:    iv.RPC,
	}
	m.blockWait = func(ctx context.Context) { sleepCtx(ctx, iv.RPCBlockWait) }

	officialNameByChain := map[string]string{
		chainIDMainnet:    officialNames["mainnet"],
		chainIDSepolia:    officialNames["sepolia"],
		chainIDArbitrum:   officialNames["arbitrum"],
		chainIDSepoliaArb: officialNames["sepolia_arb"],
	}

	for i, node := range nodes {
		m.statuses[i] = &RPCStatus{
			Config:          node,
			OfficialLagDiff: math.MaxInt64,
			X:               1,
			Y:               i + 1,
		}
		if node.IsOfficial && node.Name == officialNameByChain[node.ChainID] {
			m.official[node.ChainID] = m.statuses[i]
		}
	}

	return m
}

// Name returns the monitor name.
func (m *RPCMonitor) Name() string {
	return "RPC Monitor"
}

// Start begins monitoring.
func (m *RPCMonitor) Start(ctx context.Context, disp Display, errorChan chan<- string) {
	runLoop(ctx, m.interval, func(ctx context.Context) {
		m.check(ctx, disp, errorChan)
	})
}

// check performs a check cycle: all nodes are probed in parallel, then the
// official block numbers are published and lags computed sequentially. The
// legacy implementation computed lags inside the parallel probes, racing on
// the official nodes' LastBlockNum fields.
func (m *RPCMonitor) check(ctx context.Context, disp Display, errorChan chan<- string) {
	var wg sync.WaitGroup
	wg.Add(len(m.statuses))
	for _, status := range m.statuses {
		go func(status *RPCStatus) {
			defer wg.Done()
			m.checkNode(ctx, status, errorChan)
		}(status)
	}
	wg.Wait()

	for chainID, status := range m.official {
		if m.sharedState != nil && status.readOK {
			m.sharedState.UpdateOfficial(chainID, status.LastBlockNum)
		}
	}
	for _, status := range m.statuses {
		if status.Config.IsOfficial || !status.readOK {
			continue
		}
		official, ok := m.official[status.Config.ChainID]
		if ok && official.LastBlockNum != 0 {
			status.OfficialLagDiff = official.LastBlockNum - status.LastBlockNum
		}
	}

	m.display(disp)
}

// checkNode checks a single RPC node.
func (m *RPCMonitor) checkNode(ctx context.Context, status *RPCStatus, errorChan chan<- string) {
	if status.Config.URL == "" {
		status.Config.URL = "*** not set ***"
		return
	}

	status.ErrStr = ""
	status.Alive = false
	status.readOK = false

	rpcObj, err := rpc.DialContext(ctx, status.Config.URL)
	if err != nil {
		status.ErrStr = err.Error()
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}
	defer rpcObj.Close()

	eclient := ethclient.NewClient(rpcObj)
	latestBlock1, err := eclient.HeaderByNumber(ctx, nil)
	if err != nil {
		status.ErrStr = err.Error()
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	m.blockWait(ctx)

	latestBlock2, err := eclient.HeaderByNumber(ctx, nil)
	if err != nil {
		status.ErrStr = err.Error()
		sendErr(ctx, errorChan, status.ErrStr)
		return
	}

	status.LastBlockNum = latestBlock2.Number.Int64()
	status.readOK = true

	diff := latestBlock2.Number.Int64() - latestBlock1.Number.Int64()
	if diff == 0 {
		status.ErrStr = fmt.Sprintf("Block difference is zero (last block = %v)", latestBlock2.Number.Int64())
		sendErr(ctx, errorChan, status.ErrStr)
	} else {
		status.Alive = true
	}
}

// display renders the RPC status.
func (m *RPCMonitor) display(disp Display) {
	// Header
	disp.DrawText(Position{X: 0, Y: 0},
		"--------------------- RPC Nodes ------------------------------",
		ColorWhite, ColorDefault)

	// Each RPC node
	for _, status := range m.statuses {
		y := status.Y

		// Status (first column)
		aliveStr := "Alive"
		color := ColorGreen
		if !status.Alive {
			aliveStr = "DOWN "
			color = ColorRed
		}
		disp.DrawText(Position{X: status.X, Y: y}, aliveStr, color, ColorDefault)

		// Block number (second column)
		disp.DrawText(Position{X: status.X + 10, Y: y}, strconv.FormatInt(status.LastBlockNum, 10),
			ColorBlue, ColorDefault)

		// Official lag (third column)
		officialDiff := "------"
		if status.OfficialLagDiff != math.MaxInt64 {
			officialDiff = fmt.Sprintf("%6v", status.OfficialLagDiff)
		}
		if status.Config.IsOfficial {
			officialDiff = fmt.Sprintf("%6s", "N/A")
		}
		disp.DrawText(Position{X: status.X + 20, Y: y}, officialDiff, ColorBlue, ColorDefault)

		// Name (fourth column - shifted right)
		disp.DrawText(Position{X: status.X + 30, Y: y}, status.Config.Name, ColorWhite, ColorDefault)

		// URL (fifth column - shifted right)
		disp.DrawText(Position{X: status.X + 55, Y: y}, status.Config.URL, ColorWhite, ColorDefault)
	}

	disp.Flush()
}
