// Prizes-wallet tracking. The PrizesWallet is not upgradeable: a fix means
// deploying a new contract and pointing the game at it with setPrizesWallet
// (PrizesWalletAddressChanged). Prizes deposited into the old wallet remain
// there until their winners withdraw them, so the ETL must keep watching
// every wallet the game has ever used. This file keeps that set current:
// seeded at startup from cg_contracts plus the indexed
// PrizesWalletAddressChanged history (Contracts.PrizesWallets), grown at
// runtime by storePrizesWalletAddressChanged.
//
// A wallet never emits events before the PrizesWalletAddressChanged that
// announces it, but the announcement and the wallet's first events can land
// in the SAME engine fetch batch: the batch's FilterLogs call ran before the
// wallet joined the address set, so events between the announcement block
// and the batch's end are silently dropped. BackfillPrizesWalletEvents
// closes that window once the batch commits (same pattern as the
// implementation-constructor backfill), and RecoverPrizesWalletEvents
// repeats it at startup in case the process died in between.

package cosmicgame

import (
	"context"
	"fmt"
	"slices"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
)

// PrizesWalletBackfillLookaheadBlocks bounds the startup recovery window
// after a wallet's announcement block. The live after-commit backfill can
// only miss events inside the single engine batch that carried the
// announcement, so any value no smaller than the engine's largest fetch
// batch is safe; everything later was fetched with the wallet already in
// the FilterLogs set.
const PrizesWalletBackfillLookaheadBlocks uint64 = 50_000

// PrizesWallets returns the live set of prizes-wallet contracts.
func (h *Handlers) PrizesWallets() []ethcommon.Address {
	h.pwMu.RLock()
	defer h.pwMu.RUnlock()
	return slices.Clone(h.prizesWallets)
}

// addPrizesWallet records wallet; it reports whether the address was new.
func (h *Handlers) addPrizesWallet(wallet ethcommon.Address) bool {
	if wallet == (ethcommon.Address{}) {
		return false
	}
	h.pwMu.Lock()
	defer h.pwMu.Unlock()
	if slices.Contains(h.prizesWallets, wallet) {
		return false
	}
	h.prizesWallets = append(h.prizesWallets, wallet)
	return true
}

// SetOnNewPrizesWallet installs the callback storePrizesWalletAddressChanged
// invokes for a wallet address it has not seen before. cmd/cg-etl uses it to
// add the address to the engine's FilterLogs set. The callback runs inside
// the block's database transaction and must not write to the store itself.
func (h *Handlers) SetOnNewPrizesWallet(fn func(ctx context.Context, wallet ethcommon.Address, blockNum int64)) {
	h.pwMu.Lock()
	defer h.pwMu.Unlock()
	h.onNewPrizesWallet = fn
}

// notePrizesWallet is called by storePrizesWalletAddressChanged with the
// newly configured wallet.
func (h *Handlers) notePrizesWallet(ctx context.Context, wallet ethcommon.Address, blockNum int64) {
	if !h.addPrizesWallet(wallet) {
		return
	}
	h.log.Info("new prizes wallet", "prizes_wallet", wallet.Hex(), "block", blockNum)
	h.pwMu.RLock()
	fn := h.onNewPrizesWallet
	h.pwMu.RUnlock()
	if fn != nil {
		fn(ctx, wallet, blockNum)
	}
}

// prizesWalletSources lists every prizes wallet known so far; the dynamic
// source set of the PrizesWallet-emitted event handlers.
func (h *Handlers) prizesWalletSources() []ethcommon.Address {
	return h.PrizesWallets()
}

// BackfillPrizesWalletEvents fetches the logs wallet emitted in
// [fromBlock..toBlock], stores the missing evt_log rows and runs process
// over them. It closes the window between a wallet's announcement and the
// end of the engine batch that carried it (see the file header). It is
// idempotent: rows already present are skipped.
func (h *Handlers) BackfillPrizesWalletEvents(
	ctx context.Context,
	backfiller ImplementationBackfiller,
	process indexer.ProcessFunc,
	wallet ethcommon.Address,
	fromBlock, toBlock uint64,
) (indexer.BackfillStats, error) {
	h.addPrizesWallet(wallet)
	if toBlock < fromBlock {
		return indexer.BackfillStats{}, nil
	}
	stats, err := backfiller.BackfillContractEvtLogs(ctx, []ethcommon.Address{wallet}, fromBlock, toBlock, PrizesWalletBackfillLookaheadBlocks)
	if err != nil {
		return stats, fmt.Errorf("backfill prizes wallet %s logs [%d..%d]: %w", wallet.Hex(), fromBlock, toBlock, err)
	}
	for _, evtID := range stats.InsertedIDs {
		if err := process(ctx, evtID); err != nil {
			return stats, fmt.Errorf("processing backfilled prizes wallet event %d: %w", evtID, err)
		}
	}
	h.log.Info("prizes wallet events backfilled",
		"prizes_wallet", wallet.Hex(), "from_block", fromBlock, "to_block", toBlock,
		"logs_seen", stats.LogsSeen, "inserted", stats.Inserted, "skipped", stats.Skipped)
	return stats, nil
}

// RecoverPrizesWalletEvents is the startup pass: every wallet the indexed
// PrizesWalletAddressChanged history introduced gets its announcement window
// re-backfilled (announcement block up to the lookahead bound, capped at the
// committed watermark). The pass is idempotent, so it runs unconditionally;
// it repairs a crash between a wallet switch landing and its after-commit
// backfill finishing.
func (h *Handlers) RecoverPrizesWalletEvents(
	ctx context.Context,
	backfiller ImplementationBackfiller,
	process indexer.ProcessFunc,
	lastBlock int64,
) error {
	announcements, err := h.repo.PrizesWalletAnnouncements(ctx)
	if err != nil {
		return err
	}
	for _, a := range announcements {
		if a.BlockNum < 0 || lastBlock < a.BlockNum {
			continue
		}
		from := uint64(a.BlockNum)
		to := min(from+PrizesWalletBackfillLookaheadBlocks, uint64(lastBlock))
		if _, err := h.BackfillPrizesWalletEvents(ctx, backfiller, process,
			ethcommon.HexToAddress(a.Addr), from, to); err != nil {
			return err
		}
	}
	return nil
}
