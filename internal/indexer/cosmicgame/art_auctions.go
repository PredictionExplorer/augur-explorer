// Art-auction contract tracking (V4, in development). The
// CosmicSignatureArtAuction is not upgradeable: a replacement means
// deploying a new contract and announcing it through a game upgrade's
// initializer (ArtAuctionConfigured). Lots registered with an old auction
// contract keep living there until settled, delivered and forwarded, so the
// ETL must keep watching every auction contract the game has ever announced.
// This file keeps that set current: seeded at startup from the indexed
// ArtAuctionConfigured history (Contracts.ArtAuctions; before the V4
// upgrade the set is empty), grown at runtime by storeArtAuctionConfigured.
//
// The same announcement-window hazard as the prizes wallet applies: the
// ArtAuctionConfigured announcement and the contract's first events can land
// in the same engine fetch batch, whose FilterLogs call ran before the
// address joined the set. BackfillArtAuctionEvents closes that window after
// the batch commits, and RecoverArtAuctionEvents repeats it at startup.

package cosmicgame

import (
	"context"
	"fmt"
	"slices"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
)

// ArtAuctionBackfillLookaheadBlocks bounds the startup recovery window after
// an auction contract's announcement block; same contract as
// PrizesWalletBackfillLookaheadBlocks.
const ArtAuctionBackfillLookaheadBlocks uint64 = 50_000

// ArtAuctions returns the live set of art-auction contracts.
func (h *Handlers) ArtAuctions() []ethcommon.Address {
	h.aaMu.RLock()
	defer h.aaMu.RUnlock()
	return slices.Clone(h.artAuctions)
}

// addArtAuction records auction; it reports whether the address was new.
func (h *Handlers) addArtAuction(auction ethcommon.Address) bool {
	if auction == (ethcommon.Address{}) {
		return false
	}
	h.aaMu.Lock()
	defer h.aaMu.Unlock()
	if slices.Contains(h.artAuctions, auction) {
		return false
	}
	h.artAuctions = append(h.artAuctions, auction)
	return true
}

// SetOnNewArtAuction installs the callback storeArtAuctionConfigured invokes
// for an auction address it has not seen before. cmd/cg-etl uses it to add
// the address to the engine's FilterLogs set. The callback runs inside the
// block's database transaction and must not write to the store itself.
func (h *Handlers) SetOnNewArtAuction(fn func(ctx context.Context, auction ethcommon.Address, blockNum int64)) {
	h.aaMu.Lock()
	defer h.aaMu.Unlock()
	h.onNewArtAuction = fn
}

// noteArtAuction is called by storeArtAuctionConfigured with the newly
// announced auction contract.
func (h *Handlers) noteArtAuction(ctx context.Context, auction ethcommon.Address, blockNum int64) {
	if !h.addArtAuction(auction) {
		return
	}
	h.log.Info("new art auction contract", "art_auction", auction.Hex(), "block", blockNum)
	h.aaMu.RLock()
	fn := h.onNewArtAuction
	h.aaMu.RUnlock()
	if fn != nil {
		fn(ctx, auction, blockNum)
	}
}

// artAuctionSources lists every art-auction contract known so far; the
// dynamic source set of the CosmicSignatureArtAuction event handlers.
func (h *Handlers) artAuctionSources() []ethcommon.Address {
	return h.ArtAuctions()
}

// BackfillArtAuctionEvents fetches the logs auction emitted in
// [fromBlock..toBlock], stores the missing evt_log rows and runs process
// over them. It closes the window between an auction contract's announcement
// and the end of the engine batch that carried it (see the file header). It
// is idempotent: rows already present are skipped.
func (h *Handlers) BackfillArtAuctionEvents(
	ctx context.Context,
	backfiller ImplementationBackfiller,
	process indexer.ProcessFunc,
	auction ethcommon.Address,
	fromBlock, toBlock uint64,
) (indexer.BackfillStats, error) {
	h.addArtAuction(auction)
	if toBlock < fromBlock {
		return indexer.BackfillStats{}, nil
	}
	stats, err := backfiller.BackfillContractEvtLogs(ctx, []ethcommon.Address{auction}, fromBlock, toBlock, ArtAuctionBackfillLookaheadBlocks)
	if err != nil {
		return stats, fmt.Errorf("backfill art auction %s logs [%d..%d]: %w", auction.Hex(), fromBlock, toBlock, err)
	}
	for _, evtID := range stats.InsertedIDs {
		if err := process(ctx, evtID); err != nil {
			return stats, fmt.Errorf("processing backfilled art auction event %d: %w", evtID, err)
		}
	}
	h.log.Info("art auction events backfilled",
		"art_auction", auction.Hex(), "from_block", fromBlock, "to_block", toBlock,
		"logs_seen", stats.LogsSeen, "inserted", stats.Inserted, "skipped", stats.Skipped)
	return stats, nil
}

// RecoverArtAuctionEvents is the startup pass: every auction contract the
// indexed ArtAuctionConfigured history introduced gets its announcement
// window re-backfilled (announcement block up to the lookahead bound, capped
// at the committed watermark). The pass is idempotent, so it runs
// unconditionally; it repairs a crash between an announcement landing and
// its after-commit backfill finishing.
func (h *Handlers) RecoverArtAuctionEvents(
	ctx context.Context,
	backfiller ImplementationBackfiller,
	process indexer.ProcessFunc,
	lastBlock int64,
) error {
	announcements, err := h.repo.ArtAuctionAnnouncements(ctx)
	if err != nil {
		return err
	}
	for _, a := range announcements {
		if a.BlockNum < 0 || lastBlock < a.BlockNum {
			continue
		}
		from := uint64(a.BlockNum)
		to := min(from+ArtAuctionBackfillLookaheadBlocks, uint64(lastBlock))
		if _, err := h.BackfillArtAuctionEvents(ctx, backfiller, process,
			ethcommon.HexToAddress(a.Addr), from, to); err != nil {
			return err
		}
	}
	return nil
}
