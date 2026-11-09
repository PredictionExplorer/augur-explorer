// Package chaoszero implements the ChaosZero event handlers of the indexing
// engine: one decode/store pair per dispatched event type of the
// IndependentOutcomeMarket (round lifecycle, trades, splits/merges,
// settlement, claims and the inherited ERC-1155 transfers/approvals) and the
// OutcomeSettlementVault (escrow and payouts), built over injected
// dependencies and registered in an indexer.Registry. cmd/cz-etl is reduced
// to wiring these handlers into an indexer.Engine.
//
// Like the RandomWalk handlers, the store steps insert without a preceding
// delete-by-evtlog-id: re-processing the same evt_log row violates the
// tables' UNIQUE(evtlog_id) constraints. Re-processing only happens through
// the reorg path, which cascade-deletes the rows first.
package chaoszero

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	czc "github.com/PredictionExplorer/augur-explorer/contracts/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	czmodel "github.com/PredictionExplorer/augur-explorer/internal/model/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	czstore "github.com/PredictionExplorer/augur-explorer/internal/store/chaoszero"
)

// Contracts carries the resolved market and settlement-vault addresses (the
// two event sources; the outcome source emits no events).
type Contracts struct {
	Market ethcommon.Address
	Vault  ethcommon.Address
}

// All returns the contracts the FilterLogs subscription watches.
func (c Contracts) All() []ethcommon.Address {
	return []ethcommon.Address{c.Market, c.Vault}
}

// BootstrapContracts reads the contract registry (cz_contracts), registers
// the market, outcome-source and vault addresses in the address table (the
// fresh-database bootstrap) and resolves them. The source is registered for
// address joins only; it never emits events.
func BootstrapContracts(ctx context.Context, repo *czstore.Repo, st *store.Store) (Contracts, czmodel.ContractAddresses, error) {
	rawMarket, rawSource, rawVault, err := repo.RawContractAddrs(ctx)
	if err != nil {
		return Contracts{}, czmodel.ContractAddresses{}, fmt.Errorf("reading cz_contracts: %w", err)
	}
	for _, contractAddr := range []string{rawMarket, rawSource, rawVault} {
		if _, err := st.LookupOrCreateAddress(ctx, contractAddr, 0, 0); err != nil {
			return Contracts{}, czmodel.ContractAddresses{}, fmt.Errorf("registering contract address %v: %w", contractAddr, err)
		}
	}
	addrs, err := repo.ContractAddrs(ctx)
	if err != nil {
		return Contracts{}, addrs, fmt.Errorf("resolving contract addresses: %w", err)
	}
	return Contracts{
		Market: ethcommon.HexToAddress(addrs.Market),
		Vault:  ethcommon.HexToAddress(addrs.Vault),
	}, addrs, nil
}

// Config assembles the handler set's dependencies.
type Config struct {
	// Repo runs every ChaosZero domain write.
	Repo *czstore.Repo
	// Contracts are the resolved contract addresses (BootstrapContracts).
	Contracts Contracts
	// Logger receives one structured record per processed event; nil
	// discards them.
	Logger *slog.Logger
}

// Handlers is the ChaosZero event-handler set: dependencies, parsed ABIs
// and the registry built over them.
type Handlers struct {
	repo *czstore.Repo
	c    Contracts
	log  *slog.Logger

	marketABI *abi.ABI
	vaultABI  *abi.ABI

	registry *indexer.Registry
}

// New parses the contract ABIs, validates the dependencies and builds the
// handler registry.
func New(cfg Config) (*Handlers, error) {
	if cfg.Repo == nil {
		return nil, errors.New("chaoszero handlers: Config.Repo is required")
	}
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}

	h := &Handlers{repo: cfg.Repo, c: cfg.Contracts, log: logger}

	marketABI, err := abi.JSON(strings.NewReader(czc.MarketABI))
	if err != nil {
		return nil, fmt.Errorf("parsing IndependentOutcomeMarket ABI: %w", err)
	}
	h.marketABI = &marketABI
	vaultABI, err := abi.JSON(strings.NewReader(czc.VaultABI))
	if err != nil {
		return nil, fmt.Errorf("parsing OutcomeSettlementVault ABI: %w", err)
	}
	h.vaultABI = &vaultABI

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

// eventHandlers returns every ChaosZero event handler in registration order.
func (h *Handlers) eventHandlers() []indexer.EventHandler {
	market := []ethcommon.Address{h.c.Market}
	vault := []ethcommon.Address{h.c.Vault}
	return []indexer.EventHandler{
		indexer.NewHandler(topicHash(TopicRoundInitialized), "RoundInitialized", market, h.decodeRoundInitialized, h.storeRoundInitialized),
		indexer.NewHandler(topicHash(TopicRoundFunded), "RoundFunded", market, h.decodeRoundFunded, h.storeRoundFunded),
		indexer.NewHandler(topicHash(TopicTrade), "Trade", market, h.decodeTrade, h.storeTrade),
		indexer.NewHandler(topicHash(TopicPositionSplit), "PositionSplit", market, h.decodePositionSplit, h.storePositionSplit),
		indexer.NewHandler(topicHash(TopicPositionsMerged), "PositionsMerged", market, h.decodePositionsMerged, h.storePositionsMerged),
		indexer.NewHandler(topicHash(TopicRoundSettled), "RoundSettled", market, h.decodeRoundSettled, h.storeRoundSettled),
		indexer.NewHandler(topicHash(TopicClaimed), "Claimed", market, h.decodeClaimed, h.storeClaimed),
		indexer.NewHandler(topicHash(TopicAnswersObserved), "AnswersObserved", market, h.decodeAnswersObserved, h.storeAnswersObserved),
		indexer.NewHandler(topicHash(TopicCycleEscrowed), "CycleEscrowed", vault, h.decodeCycleEscrowed, h.storeCycleEscrowed),
		indexer.NewHandler(topicHash(TopicPaid), "Paid", vault, h.decodePaid, h.storePaid),
		// OpenZeppelin ERC-1155 events the market inherits.
		indexer.NewHandler(topicHash(TopicTransferSingle), "TransferSingle", market, h.decodeTransferSingle, h.storePositionTransfer),
		indexer.NewHandler(topicHash(TopicTransferBatch), "TransferBatch", market, h.decodeTransferBatch, h.storePositionTransfer),
		indexer.NewHandler(topicHash(TopicApprovalForAll), "ApprovalForAll", market, h.decodeApprovalForAll, h.storeApprovalForAll),
	}
}

// requireTopics guards the indexed-topic accesses of a decode function.
func requireTopics(lg *types.Log, n int) error {
	if len(lg.Topics) < n {
		return fmt.Errorf("%d topics, want %d", len(lg.Topics), n)
	}
	return nil
}
