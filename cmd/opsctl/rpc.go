package main

import (
	"context"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// rpcRequestTimeout bounds every HTTP JSON-RPC exchange of the opsctl chain
// tools (D22: node-fill, transaction collection and the CST auction scan run
// unattended for hours, so a black-holed endpoint must fail a call, not park
// the run). It sits above logscan's per-call fetch bound (one minute), so
// the scanner's own deadline stays the effective FilterLogs limit.
const rpcRequestTimeout = 90 * time.Second

// dialBoundedEthClient dials rpcURL with an injected HTTP client whose
// per-request timeout bounds every call made over the connection. WebSocket
// endpoints ignore the HTTP client and keep their caller-context bounds.
func dialBoundedEthClient(ctx context.Context, rpcURL string) (*ethclient.Client, error) {
	rpcClient, err := rpc.DialOptions(ctx, rpcURL, rpc.WithHTTPClient(&http.Client{Timeout: rpcRequestTimeout}))
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(rpcClient), nil
}
