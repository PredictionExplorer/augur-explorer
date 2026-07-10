// Package randomwalk implements the RandomWalk event handlers of the
// indexing engine: one decode/store pair per dispatched event type
// (marketplace offers/sales/cancellations, NFT mints, transfers, name
// changes and withdrawals), built over injected dependencies and registered
// in an indexer.Registry. cmd/rw-etl is reduced to wiring these handlers
// into an indexer.Engine.
//
// Unlike the CosmicGame handlers, the RandomWalk store steps insert without
// a preceding delete-by-evtlog-id: re-processing the same evt_log row
// violates the tables' UNIQUE(evtlog_id) constraints. Re-processing only
// happens through the reorg path, which cascade-deletes the rows first.
package randomwalk

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	rwc "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// Contracts carries the resolved marketplace and RandomWalk NFT addresses.
type Contracts struct {
	Market     ethcommon.Address
	RandomWalk ethcommon.Address
}

// All returns the contracts the FilterLogs subscription watches.
func (c Contracts) All() []ethcommon.Address {
	return []ethcommon.Address{c.RandomWalk, c.Market}
}

// BootstrapContracts reads the contract registry (rw_contracts), registers
// both addresses in the address table (the fresh-database bootstrap) and
// resolves them.
func BootstrapContracts(ctx context.Context, repo *rwstore.Repo, st *store.Store) (Contracts, rwp.ContractAddresses, error) {
	rawMarketplace, rawRandomwalk, err := repo.RawContractAddrs(ctx)
	if err != nil {
		return Contracts{}, rwp.ContractAddresses{}, fmt.Errorf("reading rw_contracts: %w", err)
	}
	for _, contractAddr := range []string{rawMarketplace, rawRandomwalk} {
		if _, err := st.LookupOrCreateAddress(ctx, contractAddr, 0, 0); err != nil {
			return Contracts{}, rwp.ContractAddresses{}, fmt.Errorf("registering contract address %v: %w", contractAddr, err)
		}
	}
	addrs, err := repo.ContractAddrs(ctx)
	if err != nil {
		return Contracts{}, addrs, fmt.Errorf("resolving contract addresses: %w", err)
	}
	return Contracts{
		Market:     ethcommon.HexToAddress(addrs.MarketPlace),
		RandomWalk: ethcommon.HexToAddress(addrs.RandomWalk),
	}, addrs, nil
}

// Config assembles the handler set's dependencies.
type Config struct {
	// Repo runs every RandomWalk domain write and existence guard.
	Repo *rwstore.Repo
	// Contracts are the resolved contract addresses (BootstrapContracts).
	Contracts Contracts
	// Logger receives one structured record per processed event; nil
	// discards them.
	Logger *slog.Logger
}

// Handlers is the RandomWalk event-handler set: dependencies, parsed ABIs
// and the registry built over them.
type Handlers struct {
	repo *rwstore.Repo
	c    Contracts
	log  *slog.Logger

	marketABI *abi.ABI
	rwalkABI  *abi.ABI

	registry *indexer.Registry
}

// New parses the contract ABIs, validates the dependencies and builds the
// handler registry.
func New(cfg Config) (*Handlers, error) {
	if cfg.Repo == nil {
		return nil, fmt.Errorf("randomwalk handlers: Config.Repo is required")
	}
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}

	h := &Handlers{repo: cfg.Repo, c: cfg.Contracts, log: logger}

	marketABI, err := abi.JSON(strings.NewReader(rwc.RWMarketABI))
	if err != nil {
		return nil, fmt.Errorf("parsing Marketplace ABI: %w", err)
	}
	h.marketABI = &marketABI
	rwalkABI, err := abi.JSON(strings.NewReader(rwc.RWalkABI))
	if err != nil {
		return nil, fmt.Errorf("parsing RandomWalk ABI: %w", err)
	}
	h.rwalkABI = &rwalkABI

	registry, err := indexer.NewRegistry(h.eventHandlers()...)
	if err != nil {
		return nil, err
	}
	h.registry = registry
	return h, nil
}

// Registry returns the handler registry; it plugs into indexer.LogProcessor
// and Config.TopicName.
func (h *Handlers) Registry() *indexer.Registry { return h.registry }

// eventHandlers returns every RandomWalk event handler in registration order.
func (h *Handlers) eventHandlers() []indexer.EventHandler {
	market := []ethcommon.Address{h.c.Market}
	rwalk := []ethcommon.Address{h.c.RandomWalk}
	return []indexer.EventHandler{
		indexer.NewHandler(topicHash(NEW_OFFER), "NewOffer", market, h.decodeNewOffer, h.storeNewOffer),
		indexer.NewHandler(topicHash(ITEM_BOUGHT), "ItemBought", market, h.decodeItemBought, h.storeItemBought),
		indexer.NewHandler(topicHash(OFFER_CANCELED), "OfferCanceled", market, h.decodeOfferCanceled, h.storeOfferCanceled),
		indexer.NewHandler(topicHash(WITHDRAWAL_EVT), "WithdrawalEvent", rwalk, h.decodeWithdrawal, h.storeWithdrawal),
		indexer.NewHandler(topicHash(TOKEN_NAME_EVT), "TokenNameEvent", rwalk, h.decodeTokenName, h.storeTokenName),
		indexer.NewHandler(topicHash(TRANSFER_EVT), "Transfer", rwalk, h.decodeTransfer, h.storeTransfer),
		indexer.NewHandler(topicHash(MINT_EVENT), "MintEvent", rwalk, h.decodeMintEvent, h.storeMintEvent),
	}
}

// requireTopics guards the indexed-topic accesses of a decode function.
func requireTopics(lg *types.Log, n int) error {
	if len(lg.Topics) < n {
		return fmt.Errorf("%d topics, want %d", len(lg.Topics), n)
	}
	return nil
}
