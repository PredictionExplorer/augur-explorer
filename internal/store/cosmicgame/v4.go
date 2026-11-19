// V4 (in development) event-row writes and address-set reads: the game
// proxy's art-auction wiring and ChaosZero cycle-funding events, plus the
// full event set of the non-upgradeable CosmicSignatureArtAuction contract.
// Same contracts as inserts.go/deletes.go: addresses resolve through the
// lookup-or-create cache, one INSERT per event, delete-before-insert
// idempotency via evtlog_id.

package cosmicgame

import (
	"context"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Game events.

// InsertArtAuctionConfigured records an ArtAuctionConfigured event.
func (r *Repo) InsertArtAuctionConfigured(ctx context.Context, evt *cgmodel.CGArtAuctionConfiguredEvent) error {
	const op = "insert into cg_art_auction_configured"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	auctionAid, err := r.addrID(ctx, evt.ArtAuctionAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_configured(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid,art_auction_aid" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid, auctionAid)
	return store.WrapError(op, err)
}

// DeleteArtAuctionConfigured removes an ArtAuctionConfigured row.
func (r *Repo) DeleteArtAuctionConfigured(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_configured", evtlogID)
}

// InsertArtAuctionOpeningBidChanged records an ArtAuctionOpeningBidChanged event.
func (r *Repo) InsertArtAuctionOpeningBidChanged(ctx context.Context, evt *cgmodel.CGArtAuctionOpeningBidChangedEvent) error {
	const op = "insert into cg_adm_art_auction_opening_bid"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	return r.insertAdminValue(ctx, "cg_adm_art_auction_opening_bid", "opening_bid",
		evt.EvtId, evt.BlockNum, evt.TxId, evt.TimeStamp, contractAid, evt.OpeningBid)
}

// DeleteArtAuctionOpeningBidChanged removes an ArtAuctionOpeningBidChanged row.
func (r *Repo) DeleteArtAuctionOpeningBidChanged(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_adm_art_auction_opening_bid", evtlogID)
}

// InsertArtAuctionNftsMinted records an ArtAuctionNftsMinted event.
func (r *Repo) InsertArtAuctionNftsMinted(ctx context.Context, evt *cgmodel.CGArtAuctionNftsMintedEvent) error {
	const op = "insert into cg_art_auction_nfts_minted"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	auctionAid, err := r.addrID(ctx, evt.ArtAuctionAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_nfts_minted(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"round_num,art_auction_aid,first_nft_id,num_nfts" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		evt.RoundNum, auctionAid, evt.FirstNftId, evt.NumNfts)
	return store.WrapError(op, err)
}

// DeleteArtAuctionNftsMinted removes an ArtAuctionNftsMinted row.
func (r *Repo) DeleteArtAuctionNftsMinted(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_nfts_minted", evtlogID)
}

// InsertChaosZeroCycleFunded records a ChaosZeroCycleFunded event.
func (r *Repo) InsertChaosZeroCycleFunded(ctx context.Context, evt *cgmodel.CGChaosZeroCycleFundedEvent) error {
	const op = "insert into cg_chaos_zero_cycle_funded"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	marketAid, err := r.addrID(ctx, evt.MarketAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_chaos_zero_cycle_funded(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"market_aid,cycle,amount,cumulative_amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		marketAid, evt.Cycle, evt.Amount, evt.CumulativeAmount)
	return store.WrapError(op, err)
}

// DeleteChaosZeroCycleFunded removes a ChaosZeroCycleFunded row.
func (r *Repo) DeleteChaosZeroCycleFunded(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_chaos_zero_cycle_funded", evtlogID)
}

// Art-auction contract events.

// InsertArtAuctionBatchRegistered records an AuctionBatchRegistered event.
func (r *Repo) InsertArtAuctionBatchRegistered(ctx context.Context, evt *cgmodel.CGArtAuctionBatchRegisteredEvent) error {
	const op = "insert into cg_art_auction_batch"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_batch(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"round_num,first_nft_id,num_nfts,eligible_at,opening_bid" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,TO_TIMESTAMP($9),$10)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		evt.RoundNum, evt.FirstNftId, evt.NumNfts, evt.EligibleAt, evt.OpeningBid)
	return store.WrapError(op, err)
}

// DeleteArtAuctionBatchRegistered removes an AuctionBatchRegistered row.
func (r *Repo) DeleteArtAuctionBatchRegistered(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_batch", evtlogID)
}

// InsertArtAuctionBid records an AuctionBidPlaced event.
func (r *Repo) InsertArtAuctionBid(ctx context.Context, evt *cgmodel.CGArtAuctionBidEvent) error {
	const op = "insert into cg_art_auction_bid"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	bidderAid, err := r.addrID(ctx, evt.BidderAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_bid(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"nft_id,round_num,bidder_aid,amount,eth_paid,end_at" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9,$10,TO_TIMESTAMP($11))"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		evt.NftId, evt.RoundNum, bidderAid, evt.Amount, evt.EthPaid, evt.EndAt)
	return store.WrapError(op, err)
}

// DeleteArtAuctionBid removes an AuctionBidPlaced row.
func (r *Repo) DeleteArtAuctionBid(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_bid", evtlogID)
}

// InsertArtAuctionRefundCredited records an AuctionRefundCredited event.
func (r *Repo) InsertArtAuctionRefundCredited(ctx context.Context, evt *cgmodel.CGArtAuctionRefundCreditedEvent) error {
	const op = "insert into cg_art_auction_refund_credited"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	bidderAid, err := r.addrID(ctx, evt.BidderAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_refund_credited(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"bidder_aid,nft_id,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		bidderAid, evt.NftId, evt.Amount)
	return store.WrapError(op, err)
}

// DeleteArtAuctionRefundCredited removes an AuctionRefundCredited row.
func (r *Repo) DeleteArtAuctionRefundCredited(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_refund_credited", evtlogID)
}

// InsertArtAuctionRefundWithdrawn records an AuctionRefundWithdrawn event.
func (r *Repo) InsertArtAuctionRefundWithdrawn(ctx context.Context, evt *cgmodel.CGArtAuctionRefundWithdrawnEvent) error {
	const op = "insert into cg_art_auction_refund_withdrawn"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	bidderAid, err := r.addrID(ctx, evt.BidderAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	recipientAid, err := r.addrID(ctx, evt.RecipientAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_refund_withdrawn(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"bidder_aid,recipient_aid,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		bidderAid, recipientAid, evt.Amount)
	return store.WrapError(op, err)
}

// DeleteArtAuctionRefundWithdrawn removes an AuctionRefundWithdrawn row.
func (r *Repo) DeleteArtAuctionRefundWithdrawn(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_refund_withdrawn", evtlogID)
}

// InsertArtAuctionSettled records an AuctionSettled event.
func (r *Repo) InsertArtAuctionSettled(ctx context.Context, evt *cgmodel.CGArtAuctionSettledEvent) error {
	const op = "insert into cg_art_auction_settled"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_settled(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"nft_id,round_num,winner_aid,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8,$9)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		evt.NftId, evt.RoundNum, winnerAid, evt.Amount)
	return store.WrapError(op, err)
}

// DeleteArtAuctionSettled removes an AuctionSettled row.
func (r *Repo) DeleteArtAuctionSettled(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_settled", evtlogID)
}

// InsertArtAuctionNftDelivered records an AuctionNftDelivered event.
func (r *Repo) InsertArtAuctionNftDelivered(ctx context.Context, evt *cgmodel.CGArtAuctionNftDeliveredEvent) error {
	const op = "insert into cg_art_auction_nft_delivered"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	winnerAid, err := r.addrID(ctx, evt.WinnerAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	recipientAid, err := r.addrID(ctx, evt.RecipientAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_nft_delivered(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"nft_id,winner_aid,recipient_aid" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		evt.NftId, winnerAid, recipientAid)
	return store.WrapError(op, err)
}

// DeleteArtAuctionNftDelivered removes an AuctionNftDelivered row.
func (r *Repo) DeleteArtAuctionNftDelivered(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_nft_delivered", evtlogID)
}

// InsertArtAuctionProceedsForwarded records an AuctionProceedsForwarded event.
func (r *Repo) InsertArtAuctionProceedsForwarded(ctx context.Context, evt *cgmodel.CGArtAuctionProceedsForwardedEvent) error {
	const op = "insert into cg_art_auction_proceeds"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_proceeds(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid," +
		"nft_id,round_num,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid,
		evt.NftId, evt.RoundNum, evt.Amount)
	return store.WrapError(op, err)
}

// DeleteArtAuctionProceedsForwarded removes an AuctionProceedsForwarded row.
func (r *Repo) DeleteArtAuctionProceedsForwarded(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_proceeds", evtlogID)
}

// InsertArtAuctionSurplusForwarded records an AuctionSurplusForwarded event.
func (r *Repo) InsertArtAuctionSurplusForwarded(ctx context.Context, evt *cgmodel.CGArtAuctionSurplusForwardedEvent) error {
	const op = "insert into cg_art_auction_surplus"
	contractAid, err := r.addrID(ctx, evt.ContractAddr, 0, 0)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := "INSERT INTO cg_art_auction_surplus(" +
		"evtlog_id,block_num,time_stamp,tx_id,contract_aid,amount" +
		") VALUES($1,$2,TO_TIMESTAMP($3),$4,$5,$6)"
	_, err = r.q(ctx).Exec(ctx, query,
		evt.EvtId, evt.BlockNum, evt.TimeStamp, evt.TxId, contractAid, evt.Amount)
	return store.WrapError(op, err)
}

// DeleteArtAuctionSurplusForwarded removes an AuctionSurplusForwarded row.
func (r *Repo) DeleteArtAuctionSurplusForwarded(ctx context.Context, evtlogID int64) error {
	return r.deleteByEvtlogID(ctx, "cg_art_auction_surplus", evtlogID)
}

// Address-set reads (mirroring HistoricPrizesWallets / PrizesWalletAnnouncements).

// HistoricArtAuctions returns every art-auction address an indexed
// ArtAuctionConfigured event announced, oldest first. Unlike the prizes
// wallet there is no deployment-time default: before the V4 upgrade the set
// is empty.
func (r *Repo) HistoricArtAuctions(ctx context.Context) ([]string, error) {
	rows, err := r.q(ctx).Query(ctx,
		"SELECT a.addr FROM cg_art_auction_configured c JOIN address a ON a.address_id = c.art_auction_aid "+
			"GROUP BY a.addr ORDER BY MIN(c.block_num), a.addr")
	if err != nil {
		return nil, store.WrapError("historic art auctions", err)
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			return nil, store.WrapError("historic art auctions", err)
		}
		out = append(out, addr)
	}
	return out, store.WrapError("historic art auctions", rows.Err())
}

// ArtAuctionAnnouncements returns every art-auction contract the indexed
// ArtAuctionConfigured history introduced, with the block of its first
// announcement — the point from which the contract may start emitting
// events. cg-etl uses it for the startup backfill recovery pass.
func (r *Repo) ArtAuctionAnnouncements(ctx context.Context) ([]ImplementationRef, error) {
	rows, err := r.q(ctx).Query(ctx,
		"SELECT a.addr, MIN(c.block_num) FROM cg_art_auction_configured c "+
			"JOIN address a ON a.address_id = c.art_auction_aid "+
			"GROUP BY a.addr ORDER BY MIN(c.block_num), a.addr")
	if err != nil {
		return nil, store.WrapError("art auction announcements", err)
	}
	defer rows.Close()
	var out []ImplementationRef
	for rows.Next() {
		var ref ImplementationRef
		if err := rows.Scan(&ref.Addr, &ref.BlockNum); err != nil {
			return nil, store.WrapError("art auction announcements", err)
		}
		out = append(out, ref)
	}
	return out, store.WrapError("art auction announcements", rows.Err())
}
