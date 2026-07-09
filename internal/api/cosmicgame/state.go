// state.go wires the injected contractstate.State component that replaced
// the legacy package-level contract/database state globals and their
// unkillable refresh goroutines (Phase 2, docs/MODERNIZATION.md §6.2).

package cosmicgame

import (
	"context"
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

var (
	// contractState caches the contract/database state the handlers read
	// via Snapshot(). Set once by Init; handler-level injection arrives with
	// the Phase 2 v2 Server.
	contractState *contractstate.State

	// arbStore exposes the shared base store for address lookups.
	arbStore *store.Store

	// arbRepo carries the context-first CosmicGame queries (Phase 1 Repo).
	arbRepo *cgdb.Repo
)

// initContractState loads the contract registry, builds the state component
// and performs the synchronous initial loads. The refresh loops are started
// separately via StartBackgroundRefresh, so tests get deterministic
// snapshots by simply not starting them.
func initContractState(ctx context.Context) error {
	if !dbInitialized() {
		return errors.New("cosmicgame: database link wasn't configured")
	}
	arbStore = common.Ctx.Store
	arbRepo = cgdb.NewRepo(common.Ctx.Store)

	caddrs, err := arbRepo.ContractAddrs(ctx)
	if err != nil {
		return fmt.Errorf("cosmicgame: reading contract addresses: %w", err)
	}
	st, err := contractstate.New(contractstate.Config{
		EthClient: EthClient,
		DB:        arbRepo,
		Addrs: contractstate.Addresses{
			CosmicGame:      ethcommon.HexToAddress(caddrs.CosmicGameAddr),
			CosmicSignature: ethcommon.HexToAddress(caddrs.CosmicSignatureAddr),
			CosmicToken:     ethcommon.HexToAddress(caddrs.CosmicTokenAddr),
			CharityWallet:   ethcommon.HexToAddress(caddrs.CharityWalletAddr),
			MarketingWallet: ethcommon.HexToAddress(caddrs.MarketingWalletAddr),
		},
		Info:  Info,
		Error: Error,
	})
	if err != nil {
		return fmt.Errorf("cosmicgame: building contract state: %w", err)
	}
	contractState = st
	contractState.LoadInitial(ctx)
	return nil
}

// StartBackgroundRefresh launches the periodic contract/database state
// refresh loops and returns immediately; cancelling ctx stops them. Call it
// after a successful Init. The API parity test harness deliberately never
// calls it, keeping snapshots deterministic.
func StartBackgroundRefresh(ctx context.Context) {
	if contractState == nil {
		return
	}
	go contractState.Run(ctx)
}
