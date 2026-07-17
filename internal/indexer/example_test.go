package indexer_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// tableProgress adapts a domain watermark to the engine interface the way
// the ETL binaries adapt their cg_proc_status / rw_proc_status rows.
type tableProgress struct{ lastBlock int64 }

func (p *tableProgress) LastBlock(context.Context) (int64, error)      { return p.lastBlock, nil }
func (p *tableProgress) SetLastBlock(_ context.Context, b int64) error { p.lastBlock = b; return nil }

// Example shows how cmd/cg-etl and cmd/rw-etl configure the shared engine:
// inject the store, an Ethereum client, the watermark adapter, the contract
// filter set and a per-event processor, then hand control to Run until the
// context is cancelled. There is no Output directive because the example
// needs an Ethereum node and a database; it is compiled but not executed
// (the engine's Run loop is covered by unit and integration tests).
func Example() {
	ctx := context.Background()

	st, err := store.New(ctx, store.ConfigFromEnv())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer st.Close()

	client, err := ethclient.DialContext(ctx, os.Getenv("RPC_URL"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	processEvent := func(ctx context.Context, evtID int64) error {
		// Decode the stored log and apply the domain writes; a returned
		// error fails the batch, which the engine retries with backoff.
		// Production wires internal/indexer.LogProcessor over a Registry
		// of typed EventHandlers here.
		_ = evtID
		return nil
	}

	engine, err := indexer.New(indexer.Config{
		Store:     st,
		Client:    client,
		Progress:  &tableProgress{},
		Process:   processEvent,
		Contracts: []common.Address{common.HexToAddress("0x000000000000000000000000000000000000dEaD")},
		Logger:    slog.Default(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run polls for logs, persists blocks/transactions/events, dispatches
	// each stored event to processEvent and advances the watermark. It
	// returns nil on context cancellation (graceful shutdown) and an error
	// only after the consecutive-failure circuit breaker trips.
	if err := engine.Run(ctx); err != nil {
		fmt.Println(err)
	}
}
