// The CosmicGame ETL: indexes every CosmicGame-family contract event into
// PostgreSQL. The shared process skeleton (typed configuration, process
// logger, RPC client, store, metrics, the indexing engine) lives in
// internal/etlapp; this binary contributes the CosmicGame module wiring:
// contract bootstrap, the handler set of internal/indexer/cosmicgame, the
// cg_proc_status watermark adapter, the contract-parameter drift audit,
// event-less V3 champion-duration recovery and the implementation /
// prizes-wallet recovery passes with their live hooks.
package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/ethcall"
	"github.com/PredictionExplorer/augur-explorer/internal/etlapp"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	cgindexer "github.com/PredictionExplorer/augur-explorer/internal/indexer/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// implementationCommitWait bounds how long the implementation backfill waits
// for the block carrying an Upgraded event to commit before giving up (the
// startup recovery pass catches anything left behind). Tests shrink it to
// pin the timeout arm.
var implementationCommitWait = 10 * time.Minute

// blockWatermark is the committed-watermark read the backfills poll; the
// production value is cgProgress, tests substitute failures and fixed
// positions.
type blockWatermark interface {
	LastBlock(ctx context.Context) (int64, error)
}

// backfillImplementationAfterCommit waits until the engine has committed
// blockNum (the Upgraded event's block; the hook fires inside that block's
// transaction) and then backfills the new implementation's constructor
// Initialized event. Failures are logged, not fatal: the startup recovery
// pass (RecoverImplementationEvents) repeats the attempt on the next run.
func backfillImplementationAfterCommit(
	ctx context.Context,
	logger *slog.Logger,
	progress blockWatermark,
	handlers *cgindexer.Handlers,
	engine *indexer.Engine,
	process indexer.ProcessFunc,
	impl ethcommon.Address,
	blockNum int64,
) {
	deadline := time.Now().Add(implementationCommitWait)
	for {
		last, err := progress.LastBlock(ctx)
		if err == nil && last >= blockNum {
			break
		}
		if ctx.Err() != nil {
			return
		}
		if time.Now().After(deadline) {
			logger.Warn("implementation backfill skipped: Upgraded block not committed in time",
				"implementation", impl.Hex(), "block", blockNum)
			return
		}
		time.Sleep(time.Second)
	}
	if blockNum < 0 {
		return
	}
	if _, err := handlers.BackfillImplementationEvents(ctx, engine, process, impl, uint64(blockNum)); err != nil {
		logger.Error("implementation backfill failed; the startup recovery pass will retry",
			"implementation", impl.Hex(), "block", blockNum, "err", err)
	}
}

// backfillPrizesWalletAfterCommit waits until the engine has committed
// blockNum (the PrizesWalletAddressChanged block; the hook fires inside that
// block's transaction) and then backfills the wallet's events from the
// announcement up to the committed watermark. The announcement and the
// wallet's first events can share one fetch batch whose FilterLogs call
// predates the wallet joining the address set; this closes that window.
// Failures are logged, not fatal: RecoverPrizesWalletEvents repeats the
// attempt on the next run.
func backfillPrizesWalletAfterCommit(
	ctx context.Context,
	logger *slog.Logger,
	progress blockWatermark,
	handlers *cgindexer.Handlers,
	engine *indexer.Engine,
	process indexer.ProcessFunc,
	wallet ethcommon.Address,
	blockNum int64,
) {
	if blockNum < 0 {
		return
	}
	deadline := time.Now().Add(implementationCommitWait)
	var last int64
	for {
		var err error
		last, err = progress.LastBlock(ctx)
		if err == nil && last >= blockNum {
			break
		}
		if ctx.Err() != nil {
			return
		}
		if time.Now().After(deadline) {
			logger.Warn("prizes wallet backfill skipped: announcement block not committed in time",
				"prizes_wallet", wallet.Hex(), "block", blockNum)
			return
		}
		time.Sleep(time.Second)
	}
	if _, err := handlers.BackfillPrizesWalletEvents(ctx, engine, process, wallet, uint64(blockNum), uint64(last)); err != nil {
		logger.Error("prizes wallet backfill failed; the startup recovery pass will retry",
			"prizes_wallet", wallet.Hex(), "block", blockNum, "err", err)
	}
}

// backfillArtAuctionAfterCommit is backfillPrizesWalletAfterCommit for an
// art-auction contract announced by ArtAuctionConfigured (V4): wait for the
// announcement block to commit, then backfill the contract's events up to
// the committed watermark. Failures are logged, not fatal:
// RecoverArtAuctionEvents repeats the attempt on the next run.
func backfillArtAuctionAfterCommit(
	ctx context.Context,
	logger *slog.Logger,
	progress blockWatermark,
	handlers *cgindexer.Handlers,
	engine *indexer.Engine,
	process indexer.ProcessFunc,
	auction ethcommon.Address,
	blockNum int64,
) {
	if blockNum < 0 {
		return
	}
	deadline := time.Now().Add(implementationCommitWait)
	var last int64
	for {
		var err error
		last, err = progress.LastBlock(ctx)
		if err == nil && last >= blockNum {
			break
		}
		if ctx.Err() != nil {
			return
		}
		if time.Now().After(deadline) {
			logger.Warn("art auction backfill skipped: announcement block not committed in time",
				"art_auction", auction.Hex(), "block", blockNum)
			return
		}
		time.Sleep(time.Second)
	}
	if _, err := handlers.BackfillArtAuctionEvents(ctx, engine, process, auction, uint64(blockNum), uint64(last)); err != nil {
		logger.Error("art auction backfill failed; the startup recovery pass will retry",
			"art_auction", auction.Hex(), "block", blockNum, "err", err)
	}
}

// cgProgress adapts the cg_proc_status row to the engine's watermark
// interface, preserving last_evt_id across writes.
type cgProgress struct {
	repo *cgstore.Repo
}

func (p cgProgress) LastBlock(ctx context.Context) (int64, error) {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return 0, err
	}
	return status.LastBlockNum, nil
}

func (p cgProgress) SetLastBlock(ctx context.Context, block int64) error {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return err
	}
	status.LastBlockNum = block
	return p.repo.UpdateProcessingStatus(ctx, &status)
}

// osExit is stubbed by tests that drive main through its failure arm.
var osExit = os.Exit

func main() {
	if version.HandleFlag(os.Args[1:], os.Stdout) {
		return
	}
	if err := runMain(); err != nil {
		fmt.Fprintf(os.Stderr, "cg-etl: %v\n", err)
		osExit(1)
	}
}

// runMain owns the signal-scoped context so its deferred cleanup always runs
// before main decides the exit code (os.Exit skips deferred calls).
func runMain() error {
	// Graceful shutdown: on SIGINT/SIGTERM/SIGHUP finish the current event
	// batch, write status, and exit 0 cleanly. The engine checks ctx between
	// batches and during waits.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()
	return run(ctx, os.Getenv, os.Stdout, prometheus.DefaultRegisterer, prometheus.DefaultGatherer)
}

// run drives the shared ETL skeleton with the CosmicGame module wiring; the
// signature is the integration tests' seam.
func run(ctx context.Context, getenv func(string) string, logOut io.Writer, reg prometheus.Registerer, gatherer prometheus.Gatherer) error {
	return etlapp.Run(ctx, getenv, logOut, reg, gatherer, wireCosmicGame)
}

// wireCosmicGame registers the contract addresses (fresh-database
// bootstrap), builds the event-handler set over them, audits
// contract-parameter drift, recovers event-less V3 champion durations and
// defers the implementation/prizes-wallet recovery passes (which need the
// built engine) to the PostEngine hook.
func wireCosmicGame(ctx context.Context, deps etlapp.Deps) (etlapp.Module, error) {
	logger := deps.Logger
	eclient := deps.Client
	cgRepo := cgstore.NewRepo(deps.Store)

	contracts, cgAddrs, err := cgindexer.BootstrapContracts(ctx, cgRepo, deps.Store)
	if err != nil {
		return etlapp.Module{}, fmt.Errorf("contract address bootstrap failed: %w", err)
	}
	logger.Info("All contract addresses registered in address table")

	handlers, err := cgindexer.New(cgindexer.Config{
		Repo:      cgRepo,
		Store:     deps.Store,
		Caller:    eclient,
		Contracts: contracts,
		Logger:    logger,
	})
	// The error branch below is unreachable through run(): cgindexer.New only
	// rejects a nil Repo/Store/Caller, and this wiring always supplies them.
	// Kept as a guard against future constructor validation.
	if err != nil {
		return etlapp.Module{}, fmt.Errorf("can't build event handlers: %w", err)
	}

	if _, err := cgindexer.CheckContractParamsDrift(
		ctx,
		cgRepo,
		eclient,
		cgAddrs.CosmicGameAddr,
		cgAddrs.PrizesWalletAddr,
		logger,
	); err != nil {
		logger.Error("Contract parameter drift audit failed", "err", err)
		return etlapp.Module{}, fmt.Errorf("contract parameter drift audit failed: %w", err)
	}

	recoveryCtx, cancelRecovery := context.WithTimeout(ctx, ethcall.DefaultTimeout)
	header, headerErr := eclient.HeaderByNumber(recoveryCtx, nil)
	cancelRecovery()
	if headerErr != nil {
		logger.Warn("Champion duration recovery skipped: latest header unavailable", "err", headerErr)
	} else if header.Time > math.MaxInt64 {
		return etlapp.Module{}, fmt.Errorf("champion duration recovery: header timestamp %d overflows int64", header.Time)
	} else if err := handlers.RecoverChampionDurations(
		ctx,
		header.Number.Int64(),
		int64(header.Time), // #nosec G115 -- bounded above
	); err != nil {
		return etlapp.Module{}, fmt.Errorf("champion duration recovery failed: %w", err)
	}

	progress := cgProgress{repo: cgRepo}
	postEngine := func(ctx context.Context, engine *indexer.Engine, process indexer.ProcessFunc) error {
		// Game implementations: constructor events of implementations the
		// ETL was not watching when they were deployed are backfilled now,
		// and a new implementation reported by Upgraded joins the
		// FilterLogs set at once, with its constructor event backfilled
		// once the block commits.
		if err := handlers.RecoverImplementationEvents(ctx, engine, process); err != nil {
			return fmt.Errorf("implementation event recovery failed: %w", err)
		}
		handlers.SetOnNewImplementation(func(_ context.Context, impl ethcommon.Address, blockNum int64) {
			engine.AddContracts(impl)
			go backfillImplementationAfterCommit(ctx, logger, progress, handlers, engine, process, impl, blockNum)
		})

		// Prizes wallets: a new wallet (PrizesWalletAddressChanged) joins
		// the FilterLogs set at once, and the fetch batch that carried the
		// announcement — whose FilterLogs call predates the wallet — is
		// re-covered by a backfill once the batch commits. The startup pass
		// repairs a crash between the two.
		if lastBlock, err := progress.LastBlock(ctx); err != nil {
			return fmt.Errorf("prizes wallet event recovery: read watermark: %w", err)
		} else if err := handlers.RecoverPrizesWalletEvents(ctx, engine, process, lastBlock); err != nil {
			return fmt.Errorf("prizes wallet event recovery failed: %w", err)
		}
		handlers.SetOnNewPrizesWallet(func(_ context.Context, wallet ethcommon.Address, blockNum int64) {
			engine.AddContracts(wallet)
			go backfillPrizesWalletAfterCommit(ctx, logger, progress, handlers, engine, process, wallet, blockNum)
		})

		// Art auctions (V4): a new contract (ArtAuctionConfigured) joins the
		// FilterLogs set at once, with the announcement batch re-covered by
		// a backfill once it commits; the startup pass repairs a crash
		// between the two. Before the V4 upgrade both are no-ops.
		if lastBlock, err := progress.LastBlock(ctx); err != nil {
			return fmt.Errorf("art auction event recovery: read watermark: %w", err)
		} else if err := handlers.RecoverArtAuctionEvents(ctx, engine, process, lastBlock); err != nil {
			return fmt.Errorf("art auction event recovery failed: %w", err)
		}
		handlers.SetOnNewArtAuction(func(_ context.Context, auction ethcommon.Address, blockNum int64) {
			engine.AddContracts(auction)
			go backfillArtAuctionAfterCommit(ctx, logger, progress, handlers, engine, process, auction, blockNum)
		})
		return nil
	}

	return etlapp.Module{
		Progress:   progress,
		Registry:   handlers.Registry(),
		Contracts:  contracts.All(),
		PostEngine: postEngine,
	}, nil
}
