// Package cosmicgame implements the CosmicGame event handlers of the
// indexing engine: one decode/store pair per dispatched event type, built
// over injected dependencies (repo, store, contract caller, logger) and
// registered in an indexer.Registry. cmd/cg-etl is reduced to wiring these
// handlers into an indexer.Engine.
package cosmicgame

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// Contracts carries the resolved addresses of every CosmicGame-family
// contract plus the address-table id of the CosmicToken ERC20 (the bid
// handlers' reward pre-query filters evt_log rows on it).
type Contracts struct {
	Game            ethcommon.Address
	Signature       ethcommon.Address
	Token           ethcommon.Address
	Dao             ethcommon.Address
	CharityWallet   ethcommon.Address
	PrizesWallet    ethcommon.Address
	StakingCST      ethcommon.Address
	StakingRWalk    ethcommon.Address
	MarketingWallet ethcommon.Address
	Implementation  ethcommon.Address

	CosmicTokenAid int64
}

// All returns every contract the FilterLogs subscription watches. The DAO is
// included although Governor events are stored in evt_log only (no cg_dao_*
// layer-2 tables); the registry ignores unknown topics.
func (c Contracts) All() []ethcommon.Address {
	return []ethcommon.Address{
		c.Game,
		c.Signature,
		c.Token,
		c.Dao,
		c.CharityWallet,
		c.PrizesWallet,
		c.StakingCST,
		c.StakingRWalk,
		c.MarketingWallet,
		c.Implementation,
	}
}

// BootstrapContracts reads the contract registry (cg_contracts), registers
// every address in the address table (the fresh-database bootstrap) and
// resolves the ids the handlers need. Registration order is fixed: on an
// empty database it determines the assigned address ids.
func BootstrapContracts(ctx context.Context, repo *cgstore.Repo, st *store.Store) (Contracts, cgmodel.ContractAddrs, error) {
	addrs, err := repo.ContractAddrs(ctx)
	if err != nil {
		return Contracts{}, addrs, fmt.Errorf("reading contract addresses: %w", err)
	}

	for _, contractAddr := range []string{
		addrs.CosmicGameAddr,
		addrs.CosmicSignatureAddr,
		addrs.CosmicTokenAddr,
		addrs.CosmicDaoAddr,
		addrs.CharityWalletAddr,
		addrs.PrizesWalletAddr,
		addrs.RandomWalkAddr,
		addrs.StakingWalletCSTAddr,
		addrs.StakingWalletRWalkAddr,
		addrs.MarketingWalletAddr,
		addrs.ImplementationAddr,
	} {
		if _, err := st.LookupOrCreateAddress(ctx, contractAddr, 0, 0); err != nil {
			return Contracts{}, addrs, fmt.Errorf("registering contract address %v: %w", contractAddr, err)
		}
	}

	tokenAid, err := st.LookupAddressID(ctx, addrs.CosmicTokenAddr)
	if err != nil {
		return Contracts{}, addrs, fmt.Errorf("looking up CosmicToken address id: %w", err)
	}

	return Contracts{
		Game:            ethcommon.HexToAddress(addrs.CosmicGameAddr),
		Signature:       ethcommon.HexToAddress(addrs.CosmicSignatureAddr),
		Token:           ethcommon.HexToAddress(addrs.CosmicTokenAddr),
		Dao:             ethcommon.HexToAddress(addrs.CosmicDaoAddr),
		CharityWallet:   ethcommon.HexToAddress(addrs.CharityWalletAddr),
		PrizesWallet:    ethcommon.HexToAddress(addrs.PrizesWalletAddr),
		StakingCST:      ethcommon.HexToAddress(addrs.StakingWalletCSTAddr),
		StakingRWalk:    ethcommon.HexToAddress(addrs.StakingWalletRWalkAddr),
		MarketingWallet: ethcommon.HexToAddress(addrs.MarketingWalletAddr),
		Implementation:  ethcommon.HexToAddress(addrs.ImplementationAddr),
		CosmicTokenAid:  tokenAid,
	}, addrs, nil
}

// Config assembles the handler set's dependencies.
type Config struct {
	// Repo runs every CosmicGame domain write.
	Repo *cgstore.Repo
	// Store serves the base-layer reads (stored sibling event logs).
	Store *store.Store
	// Caller performs the contract reads two handlers need (donation-info
	// records, donated-NFT tokenURI). *ethclient.Client satisfies it.
	Caller bind.ContractCaller
	// Contracts are the resolved contract addresses (BootstrapContracts).
	Contracts Contracts
	// Logger receives one structured record per processed event; nil
	// discards them.
	Logger *slog.Logger
}

// Handlers is the CosmicGame event-handler set: dependencies, parsed ABIs
// and the registry built over them.
type Handlers struct {
	repo   *cgstore.Repo
	store  *store.Store
	caller bind.ContractCaller
	c      Contracts
	log    *slog.Logger

	gameABI            *abi.ABI
	gameV2ABI          *abi.ABI
	signatureABI       *abi.ABI
	charityWalletABI   *abi.ABI
	prizesWalletABI    *abi.ABI
	stakingCSTABI      *abi.ABI
	stakingRWalkABI    *abi.ABI
	marketingWalletABI *abi.ABI
	erc20ABI           *abi.ABI
	erc1967ABI         *abi.ABI

	registry *indexer.Registry
}

// New parses the contract ABIs, validates the dependencies and builds the
// handler registry.
func New(cfg Config) (*Handlers, error) {
	if cfg.Repo == nil {
		return nil, fmt.Errorf("cosmicgame handlers: Config.Repo is required")
	}
	if cfg.Store == nil {
		return nil, fmt.Errorf("cosmicgame handlers: Config.Store is required")
	}
	if cfg.Caller == nil {
		return nil, fmt.Errorf("cosmicgame handlers: Config.Caller is required")
	}
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}

	h := &Handlers{
		repo:   cfg.Repo,
		store:  cfg.Store,
		caller: cfg.Caller,
		c:      cfg.Contracts,
		log:    logger,
	}
	for _, a := range []struct {
		dst  **abi.ABI
		name string
		raw  string
	}{
		{&h.gameABI, "CosmicSignatureGame", cgc.CosmicSignatureGameABI},
		{&h.gameV2ABI, "CosmicSignatureGameV2", cgc.CosmicSignatureGameV2ABI},
		{&h.signatureABI, "CosmicSignatureNft", cgc.CosmicSignatureNftABI},
		{&h.charityWalletABI, "CharityWallet", cgc.CharityWalletABI},
		{&h.prizesWalletABI, "PrizesWallet", cgc.PrizesWalletABI},
		{&h.stakingCSTABI, "IStakingWalletCosmicSignatureNft", cgc.IStakingWalletCosmicSignatureNftABI},
		{&h.stakingRWalkABI, "IStakingWalletRandomWalkNft", cgc.IStakingWalletRandomWalkNftABI},
		{&h.marketingWalletABI, "MarketingWallet", cgc.MarketingWalletABI},
		{&h.erc20ABI, "ERC20", cgc.ERC20ABI},
		{&h.erc1967ABI, "IERC1967", cgc.IERC1967ABI},
	} {
		parsed, err := abi.JSON(strings.NewReader(a.raw))
		if err != nil {
			return nil, fmt.Errorf("parsing %s ABI: %w", a.name, err)
		}
		*a.dst = &parsed
	}

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

// requireTopics guards the indexed-topic accesses of a decode function:
// a matching topic0 with fewer indexed topics than the event declares means
// corrupt input, reported as an error instead of an index panic.
func requireTopics(lg *types.Log, n int) error {
	if len(lg.Topics) < n {
		return fmt.Errorf("%d topics, want %d", len(lg.Topics), n)
	}
	return nil
}
