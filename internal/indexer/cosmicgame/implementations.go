// Game implementation tracking. The CosmicSignatureGame proxy is upgraded in
// place; each implementation contract's constructor emits OpenZeppelin
// Initialized(type(uint64).max) from _disableInitializers(), in the deploy
// transaction that precedes the Upgraded event. Watching only the
// implementation registered in cg_contracts therefore captured one
// implementation's event and dropped the others (the live dev database had
// three implementations and one Initialized(-1) row). This file keeps the set
// of implementations current and backfills the constructor event of any
// implementation whose event the ETL never saw.

package cosmicgame

import (
	"context"
	"fmt"
	"slices"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
)

// ImplementationLookbackBlocks bounds how far before an Upgraded event the
// constructor Initialized of the new implementation is searched for. The
// deploy transaction normally lands seconds before the upgrade; the window is
// generous for a multi-step manual upgrade.
const ImplementationLookbackBlocks uint64 = 200_000

// Implementations returns the live set of game implementation contracts.
func (h *Handlers) Implementations() []ethcommon.Address {
	h.implMu.RLock()
	defer h.implMu.RUnlock()
	return slices.Clone(h.implementations)
}

// addImplementation records impl; it reports whether the address was new.
func (h *Handlers) addImplementation(impl ethcommon.Address) bool {
	if impl == (ethcommon.Address{}) {
		return false
	}
	h.implMu.Lock()
	defer h.implMu.Unlock()
	if slices.Contains(h.implementations, impl) {
		return false
	}
	h.implementations = append(h.implementations, impl)
	return true
}

// SetOnNewImplementation installs the callback storeUpgraded invokes for an
// implementation address it has not seen before. cmd/cg-etl uses it to add
// the address to the engine's FilterLogs set and to backfill the
// constructor event once the block carrying the Upgraded event has
// committed. The callback runs inside the block's database transaction and
// must not write to the store itself.
func (h *Handlers) SetOnNewImplementation(fn func(ctx context.Context, impl ethcommon.Address, blockNum int64)) {
	h.implMu.Lock()
	defer h.implMu.Unlock()
	h.onNewImplementation = fn
}

// noteImplementation is called by storeUpgraded with the new implementation.
func (h *Handlers) noteImplementation(ctx context.Context, impl ethcommon.Address, blockNum int64) {
	if !h.addImplementation(impl) {
		return
	}
	h.log.Info("new game implementation", "implementation", impl.Hex(), "block", blockNum)
	h.implMu.RLock()
	fn := h.onNewImplementation
	h.implMu.RUnlock()
	if fn != nil {
		fn(ctx, impl, blockNum)
	}
}

// initializedSources lists the contracts that may emit OpenZeppelin
// Initializable:Initialized: the ownership set plus every game
// implementation known so far.
func (h *Handlers) initializedSources() []ethcommon.Address {
	return append(h.ownershipSources(), h.Implementations()...)
}

// ImplementationBackfiller is the slice of indexer.Engine the recovery needs.
type ImplementationBackfiller interface {
	BackfillContractEvtLogs(ctx context.Context, contracts []ethcommon.Address, fromBlock, toBlock, batchSize uint64) (indexer.BackfillStats, error)
}

// BackfillImplementationEvents fetches the logs impl emitted in the
// ImplementationLookbackBlocks blocks up to and including upgradeBlock
// (its constructor Initialized), stores the missing evt_log rows and runs
// process over them so cg_adm_initialized receives the row. It is idempotent:
// rows already present are skipped.
func (h *Handlers) BackfillImplementationEvents(
	ctx context.Context,
	backfiller ImplementationBackfiller,
	process indexer.ProcessFunc,
	impl ethcommon.Address,
	upgradeBlock uint64,
) (indexer.BackfillStats, error) {
	h.addImplementation(impl)
	from := uint64(0)
	if upgradeBlock > ImplementationLookbackBlocks {
		from = upgradeBlock - ImplementationLookbackBlocks
	}
	stats, err := backfiller.BackfillContractEvtLogs(ctx, []ethcommon.Address{impl}, from, upgradeBlock, ImplementationLookbackBlocks)
	if err != nil {
		return stats, fmt.Errorf("backfill implementation %s logs [%d..%d]: %w", impl.Hex(), from, upgradeBlock, err)
	}
	for _, evtID := range stats.InsertedIDs {
		if err := process(ctx, evtID); err != nil {
			return stats, fmt.Errorf("processing backfilled implementation event %d: %w", evtID, err)
		}
	}
	h.log.Info("implementation events backfilled",
		"implementation", impl.Hex(), "from_block", from, "to_block", upgradeBlock,
		"logs_seen", stats.LogsSeen, "inserted", stats.Inserted, "skipped", stats.Skipped)
	return stats, nil
}

// RecoverImplementationEvents is the startup pass: every implementation an
// indexed Upgraded event introduced but for which no Initialized row exists
// gets BackfillImplementationEvents. A failure is returned to the caller;
// cg-etl treats it as fatal like the other startup recoveries.
func (h *Handlers) RecoverImplementationEvents(
	ctx context.Context,
	backfiller ImplementationBackfiller,
	process indexer.ProcessFunc,
) error {
	missing, err := h.repo.ImplementationsMissingInitialized(ctx)
	if err != nil {
		return err
	}
	for _, m := range missing {
		if m.BlockNum < 0 {
			return fmt.Errorf("implementation %s: negative upgrade block %d", m.Addr, m.BlockNum)
		}
		if _, err := h.BackfillImplementationEvents(ctx, backfiller, process,
			ethcommon.HexToAddress(m.Addr), uint64(m.BlockNum)); err != nil {
			return err
		}
	}
	return nil
}
