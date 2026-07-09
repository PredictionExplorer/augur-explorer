// ETL-facing writes (offers, buys, mints, transfers, withdrawals, token
// names), the processing-status watermark, ranking accumulators and the
// notification read surface of the RandomWalk domain.

package randomwalk

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"

	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// =============================================================================
// PROCESSING STATUS
// =============================================================================

// ProcessingStatus returns the ETL watermark (last processed event id and
// block number), lazily creating the singleton rw_proc_status row on a fresh
// database.
func (r *Repo) ProcessingStatus(ctx context.Context) (rwp.ProcStatus, error) {
	const op = "randomwalk processing status"
	var output rwp.ProcStatus
	var lastID, lastBlock *int64
	err := r.pool().QueryRow(ctx, "SELECT last_evt_id,last_block FROM rw_proc_status").Scan(&lastID, &lastBlock)
	if err != nil {
		wrapped := store.WrapError(op, err)
		if !errors.Is(wrapped, store.ErrNotFound) {
			return output, wrapped
		}
		// Fresh database: create the singleton row and report the zero
		// watermark it holds.
		if _, err := r.pool().Exec(ctx, "INSERT INTO rw_proc_status DEFAULT VALUES"); err != nil {
			return output, store.WrapError(op+": insert default row", err)
		}
		if err := r.pool().QueryRow(ctx, "SELECT last_evt_id,last_block FROM rw_proc_status").Scan(&lastID, &lastBlock); err != nil {
			return output, store.WrapError(op, err)
		}
	}
	if lastID != nil {
		output.LastIdProcessed = *lastID
	}
	if lastBlock != nil {
		output.LastBlockNum = *lastBlock
	}
	return output, nil
}

// UpdateProcessingStatus persists the ETL watermark.
func (r *Repo) UpdateProcessingStatus(ctx context.Context, status *rwp.ProcStatus) error {
	_, err := r.pool().Exec(ctx, "UPDATE rw_proc_status SET last_evt_id = $1,last_block=$2",
		status.LastIdProcessed, status.LastBlockNum)
	return store.WrapError("update randomwalk processing status", err)
}

// =============================================================================
// CONTRACT CONFIGURATION
// =============================================================================

// ContractAddrs returns the marketplace and RandomWalk contract addresses
// with their address ids from the rw_contracts registry (one row). A
// database without the registry row (or with unregistered addresses) yields
// store.ErrNotFound.
func (r *Repo) ContractAddrs(ctx context.Context) (rwp.ContractAddresses, error) {
	var output rwp.ContractAddresses
	query := `SELECT
			marketplace_addr,randomwalk_addr,
			mp_a.address_id,rw_a.address_id
		FROM rw_contracts rw
			JOIN address mp_a ON rw.marketplace_addr=mp_a.addr
			JOIN address rw_a ON rw.randomwalk_addr=rw_a.addr`
	err := r.pool().QueryRow(ctx, query).Scan(
		&output.MarketPlace,
		&output.RandomWalk,
		&output.MarketPlaceAid,
		&output.RandomWalkAid,
	)
	if err != nil {
		return output, store.WrapError("randomwalk contract addrs", err)
	}
	return output, nil
}

// RawContractAddrs returns the marketplace and RandomWalk addresses straight
// from rw_contracts without joining the address table — the ETL uses it at
// startup to register both addresses before ContractAddrs can resolve them.
func (r *Repo) RawContractAddrs(ctx context.Context) (marketplace, randomwalk string, err error) {
	err = r.pool().QueryRow(ctx,
		"SELECT marketplace_addr, randomwalk_addr FROM rw_contracts LIMIT 1").Scan(&marketplace, &randomwalk)
	if err != nil {
		return "", "", store.WrapError("raw randomwalk contract addrs", err)
	}
	return marketplace, randomwalk, nil
}

// =============================================================================
// OFFER OPERATIONS
// =============================================================================

// InsertNewOffer records a NewOffer marketplace event. Offers whose seller
// is the zero address are BUY offers (otype 0), everything else SELL
// (otype 1).
func (r *Repo) InsertNewOffer(ctx context.Context, evt *rwp.NewOffer) error {
	const op = "insert into rw_new_offer"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	rwalkAid, err := r.addrID(ctx, evt.RWalkAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	buyerAid, err := r.addrID(ctx, evt.Buyer, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	sellerAid, err := r.addrID(ctx, evt.Seller, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	otype := 1
	if evt.Seller == "0x0000000000000000000000000000000000000000" {
		otype = 0
	}
	query := `INSERT INTO rw_new_offer(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			rwalk_aid,offer_id,otype,token_id,buyer_aid,seller_aid,active,price
		) VALUES (
			$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11,$12,$13
		)`
	_, err = r.pool().Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		rwalkAid,
		evt.OfferId,
		otype,
		evt.TokenId,
		buyerAid,
		sellerAid,
		true,
		evt.Price,
	)
	return store.WrapError(op, err)
}

// InsertItemBought records an ItemBought marketplace event.
func (r *Repo) InsertItemBought(ctx context.Context, evt *rwp.ItemBought) error {
	const op = "insert into rw_item_bought"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	sellerAid, err := r.addrID(ctx, evt.SellerAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	buyerAid, err := r.addrID(ctx, evt.BuyerAddr, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO rw_item_bought(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			offer_id,seller_aid,buyer_aid
		) VALUES (
			$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8
		)`
	_, err = r.pool().Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.OfferId,
		sellerAid,
		buyerAid,
	)
	return store.WrapError(op, err)
}

// InsertOfferCanceled records an OfferCanceled marketplace event.
func (r *Repo) InsertOfferCanceled(ctx context.Context, evt *rwp.OfferCanceled) error {
	const op = "insert into rw_offer_canceled"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO rw_offer_canceled(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			offer_id
		) VALUES (
			$1,$2,$3,TO_TIMESTAMP($4),$5,$6
		)`
	_, err = r.pool().Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.OfferId,
	)
	return store.WrapError(op, err)
}

// =============================================================================
// TOKEN OPERATIONS
// =============================================================================

// InsertWithdrawal records a WithdrawalEvent of the RandomWalk contract.
func (r *Repo) InsertWithdrawal(ctx context.Context, evt *rwp.Withdrawal) error {
	const op = "insert into rw_withdrawal"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	aid, err := r.addrID(ctx, evt.Destination, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO rw_withdrawal(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			aid,token_id,amount
		) VALUES (
			$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8
		)`
	_, err = r.pool().Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		aid,
		evt.TokenId,
		evt.Amount,
	)
	return store.WrapError(op, err)
}

// InsertTokenName records a TokenNameEvent (token renamed).
func (r *Repo) InsertTokenName(ctx context.Context, evt *rwp.TokenName) error {
	const op = "insert into rw_token_name"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO rw_token_name(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			token_id,new_name
		) VALUES (
			$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7
		)`
	_, err = r.pool().Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.TokenId,
		evt.NewName,
	)
	return store.WrapError(op, err)
}

// InsertMint records a MintEvent of the RandomWalk contract.
func (r *Repo) InsertMint(ctx context.Context, evt *rwp.MintEvent) error {
	const op = "insert into rw_mint_evt"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	ownerAid, err := r.addrID(ctx, evt.Owner, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	query := `INSERT INTO rw_mint_evt(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			token_id,owner_aid,seed,seed_num,price
		) VALUES (
			$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10
		)`
	_, err = r.pool().Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.TokenId,
		ownerAid,
		evt.Seed,
		evt.SeedNum,
		evt.Price,
	)
	return store.WrapError(op, err)
}

// InsertTransfer records an ERC-721 Transfer of a RandomWalk token. Mints
// (from the zero address) are otype 1, burns (to the zero address) otype 2,
// wallet-to-wallet transfers otype 0.
func (r *Repo) InsertTransfer(ctx context.Context, evt *rwp.Transfer) error {
	const op = "insert into rw_transfer"
	contractAid, err := r.addrID(ctx, evt.Contract, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	fromAid, err := r.addrID(ctx, evt.From, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	toAid, err := r.addrID(ctx, evt.To, evt.BlockNum, evt.TxId)
	if err != nil {
		return store.WrapError(op, err)
	}
	var zero common.Address
	otype := 0
	if common.HexToAddress(evt.From) == zero {
		otype = 1
	}
	if common.HexToAddress(evt.To) == zero {
		otype = 2
	}
	query := `INSERT INTO rw_transfer(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			token_id,from_aid,to_aid,otype
		) VALUES (
			$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9
		)`
	_, err = r.pool().Exec(ctx, query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contractAid,
		evt.TokenId,
		fromAid,
		toAid,
		otype,
	)
	return store.WrapError(op, err)
}

// OfferExists reports whether an offer with offerID exists for the contract.
// An unknown contract address yields false. (The legacy layer also returned
// false when the address lookup hit a real DB failure, silently skipping the
// event; such errors propagate now.)
func (r *Repo) OfferExists(ctx context.Context, contractAddr string, offerID int64) (bool, error) {
	contractAid, err := r.store.LookupAddressID(ctx, contractAddr)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return false, nil
		}
		return false, err
	}
	var id int64
	err = r.pool().QueryRow(ctx,
		"SELECT id FROM rw_new_offer WHERE contract_aid=$1 AND offer_id=$2",
		contractAid, offerID).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, store.WrapError("offer exists check", err)
	}
	return true, nil
}

// TokenExists reports whether a mint event exists for tokenID on the given
// contract. An unknown contract address yields false.
func (r *Repo) TokenExists(ctx context.Context, contractAddr string, tokenID int64) (bool, error) {
	contractAid, err := r.store.LookupAddressID(ctx, contractAddr)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return false, nil
		}
		return false, err
	}
	var id int64
	err = r.pool().QueryRow(ctx,
		"SELECT id FROM rw_mint_evt WHERE contract_aid=$1 AND token_id=$2",
		contractAid, tokenID).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, store.WrapError("token exists check", err)
	}
	return true, nil
}

// =============================================================================
// RANKING OPERATIONS
// =============================================================================

// updateRank UPDATEs one rw_uranks rank/value column pair, inserting the row
// when the user has no rank entry yet (the rwctl top-rated cron owns this
// table; there is no trigger maintaining it).
func (r *Repo) updateRank(ctx context.Context, op, rankColumn, valueColumn string, aid int64, rank float64, value any) error {
	res, err := r.pool().Exec(ctx,
		"UPDATE rw_uranks SET "+rankColumn+" = $2,"+valueColumn+"=$3 WHERE aid = $1",
		aid, rank, value)
	if err != nil {
		return store.WrapError(op, err)
	}
	if res.RowsAffected() == 0 {
		_, err := r.pool().Exec(ctx,
			"INSERT INTO rw_uranks(aid,"+rankColumn+","+valueColumn+") VALUES($1,$2,$3)",
			aid, rank, value)
		if err != nil {
			return store.WrapError(op+": insert", err)
		}
	}
	return nil
}

// UpdateTopProfitRank upserts a user's profit-leaderboard percentile and
// profit value.
func (r *Repo) UpdateTopProfitRank(ctx context.Context, aid int64, rank, profit float64) error {
	return r.updateRank(ctx, "update top profit rank", "top_profit", "profit", aid, rank, profit)
}

// UpdateTopTotalTradesRank upserts a user's trades-leaderboard percentile
// and trade count.
func (r *Repo) UpdateTopTotalTradesRank(ctx context.Context, aid int64, rank float64, totalTrades int64) error {
	return r.updateRank(ctx, "update top total trades rank", "top_trades", "total_trades", aid, rank, totalTrades)
}

// UpdateTopVolumeRank upserts a user's volume-leaderboard percentile and
// traded volume.
func (r *Repo) UpdateTopVolumeRank(ctx context.Context, aid int64, rank, volume float64) error {
	return r.updateRank(ctx, "update top volume rank", "top_volume", "volume", aid, rank, volume)
}

// RankingDataForAllUsers aggregates per-user trade count, profit and volume
// across contracts (input of the rwctl top-rated rank computation).
func (r *Repo) RankingDataForAllUsers(ctx context.Context) ([]rwp.RankStats, error) {
	query := `SELECT
			user_aid,
			SUM(total_num_trades) AS tot_trades,
			SUM(total_profit) AS profit,
			SUM(total_vol) AS tot_volume
		FROM rw_user_stats
		GROUP BY user_aid`
	return queryList(ctx, r, "ranking data for all users", 8, query, func(rows pgx.Rows, rec *rwp.RankStats) error {
		return rows.Scan(&rec.Aid, &rec.TotalTrades, &rec.ProfitLoss, &rec.VolumeTraded)
	})
}

// TopProfitMakers returns the profit leaderboard (best percentile first,
// top 100).
func (r *Repo) TopProfitMakers(ctx context.Context) ([]rwp.ProfitMaker, error) {
	query := `SELECT
			a.addr,
			r.top_profit,
			r.profit/1e+18
		FROM rw_uranks AS r
		LEFT JOIN address AS a ON r.aid = a.address_id
		ORDER BY r.top_profit ASC,r.profit DESC LIMIT 100`
	return queryList(ctx, r, "top profit makers", 101, query, func(rows pgx.Rows, rec *rwp.ProfitMaker) error {
		return rows.Scan(&rec.Addr, &rec.Percentage, &rec.ProfitLoss)
	})
}

// TopTradeMakers returns the trade-count leaderboard (top 100).
func (r *Repo) TopTradeMakers(ctx context.Context) ([]rwp.TradeMaker, error) {
	query := `SELECT
			a.addr,
			r.top_trades,
			r.total_trades
		FROM rw_uranks AS r
		LEFT JOIN address AS a ON r.aid = a.address_id
		ORDER BY r.top_trades ASC,r.total_trades DESC LIMIT 100`
	return queryList(ctx, r, "top trade makers", 101, query, func(rows pgx.Rows, rec *rwp.TradeMaker) error {
		return rows.Scan(&rec.Addr, &rec.Percentage, &rec.TotalTrades)
	})
}

// TopVolumeMakers returns the traded-volume leaderboard (top 100).
func (r *Repo) TopVolumeMakers(ctx context.Context) ([]rwp.VolumeMaker, error) {
	query := `SELECT
			a.addr,
			r.top_volume,
			r.volume/1e+18
		FROM rw_uranks AS r
		LEFT JOIN address AS a ON r.aid = a.address_id
		ORDER BY r.top_volume ASC,r.volume DESC LIMIT 100`
	return queryList(ctx, r, "top volume makers", 101, query, func(rows pgx.Rows, rec *rwp.VolumeMaker) error {
		return rows.Scan(&rec.Addr, &rec.Percentage, &rec.Volume)
	})
}

// =============================================================================
// NOTIFICATION & MESSAGING
// =============================================================================

// MintEventsForNotification returns mints after startTs (timestamp, token,
// price, seed) for notification bots.
func (r *Repo) MintEventsForNotification(ctx context.Context, rwalkAid, startTs int64) ([]rwp.NotificationEvent, error) {
	query := `SELECT
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,
			token_id,
			price/1e+18 AS price,
			seed
		FROM rw_mint_evt m
		WHERE (contract_aid=$1) AND (time_stamp > TO_TIMESTAMP($2))`
	return queryList(ctx, r, "mint events for notification", 101, query, func(rows pgx.Rows, rec *rwp.NotificationEvent) error {
		return rows.Scan(
			&rec.TimeStampMinted,
			&rec.TokenId,
			&rec.Price,
			&rec.SeedHex,
		)
	}, rwalkAid, startTs)
}

// MessagingStatus returns the notification watermark; a rowless table (not
// seeded yet) yields the zero status.
func (r *Repo) MessagingStatus(ctx context.Context) (rwp.MsgStatus, error) {
	var output rwp.MsgStatus
	err := r.pool().QueryRow(ctx,
		"SELECT last_tx_id,last_evtlog_id,last_block_num,last_timestamp FROM rw_messaging_status").Scan(
		&output.TxId,
		&output.EvtLogId,
		&output.BlockNum,
		&output.TimeStamp,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return output, store.WrapError("messaging status", err)
	}
	return output, nil
}

// UpdateMessagingStatus persists the notification watermark.
func (r *Repo) UpdateMessagingStatus(ctx context.Context, status *rwp.MsgStatus) error {
	query := `UPDATE rw_messaging_status SET
			last_tx_id = $1,
			last_evtlog_id = $2,
			last_block_num = $3,
			last_timestamp = $4`
	_, err := r.pool().Exec(ctx, query, status.TxId, status.EvtLogId, status.BlockNum, status.TimeStamp)
	return store.WrapError("update messaging status", err)
}

// AllEventsForNotification returns mints, offers and buys after startTs in
// timestamp order (evt_type 1 mint, 2 sell offer, 5 buy offer, 3 bought).
func (r *Repo) AllEventsForNotification(ctx context.Context, rwalkAid, startTs int64) ([]rwp.NotificationEvent, error) {
	query := `SELECT
			ts,
			token_id,
			price,
			evt_type
		FROM (
		(
		SELECT
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,
			token_id,
			price/1e+18 AS price,
			1 AS evt_type
		FROM rw_mint_evt m
		WHERE (contract_aid=$1) AND (time_stamp > TO_TIMESTAMP($2))
		) UNION ALL(
		SELECT
			EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,
			token_id,
			price/1e+18 AS price,
			2 AS evt_type
		FROM rw_new_offer o
		WHERE (rwalk_aid=$1) AND (time_stamp > TO_TIMESTAMP($2)) AND (otype=1)
		) UNION ALL(
		SELECT
			EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,
			token_id,
			price/1e+18 AS price,
			5 AS evt_type
		FROM rw_new_offer o
		WHERE (rwalk_aid=$1) AND (time_stamp > TO_TIMESTAMP($2)) AND (otype=0)
		) UNION ALL (
		SELECT
			EXTRACT(EPOCH FROM b.time_stamp)::BIGINT as ts,
			o.token_id,
			o.price/1e+18 AS price,
			3 AS evt_type
		FROM rw_item_bought b
		JOIN rw_new_offer o ON (b.contract_aid=o.contract_aid) AND (b.offer_id=o.offer_id)
		WHERE (o.rwalk_aid=$1) AND (b.time_stamp > TO_TIMESTAMP($2))
		)
		) data
		ORDER BY ts`
	return queryList(ctx, r, "all events for notification", 101, query, scanNotificationEvent, rwalkAid, startTs)
}

func scanNotificationEvent(rows pgx.Rows, rec *rwp.NotificationEvent) error {
	return rows.Scan(
		&rec.TimeStampMinted,
		&rec.TokenId,
		&rec.Price,
		&rec.EvtType,
	)
}

// AllEventsForNotificationSinceEvtlog is AllEventsForNotification keyed on
// the evt_log id watermark instead of a timestamp (notibot's cursor), with
// tx and evtlog ids included in the result.
func (r *Repo) AllEventsForNotificationSinceEvtlog(ctx context.Context, rwalkAid, startEvtlogID int64) ([]rwp.NotificationEvent2, error) {
	query := `SELECT
			ts,
			tx_id,
			evtlog_id,
			token_id,
			price,
			evt_type
		FROM (
		(
		SELECT
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,
			m.tx_id,
			m.evtlog_id,
			token_id,
			price/1e+18 AS price,
			1 AS evt_type
		FROM rw_mint_evt m
		WHERE (contract_aid=$1) AND (m.evtlog_id>$2)
		) UNION ALL(
		SELECT
			EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,
			o.tx_id,
			o.evtlog_id,
			token_id,
			price/1e+18 AS price,
			2 AS evt_type
		FROM rw_new_offer o
		WHERE (rwalk_aid=$1) AND (o.evtlog_id>$2) AND (otype=1)
		) UNION ALL(
		SELECT
			EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,
			o.tx_id,
			o.evtlog_id,
			token_id,
			price/1e+18 AS price,
			5 AS evt_type
		FROM rw_new_offer o
		WHERE (rwalk_aid=$1) AND (o.evtlog_id>$2) AND (otype=0)
		) UNION ALL (
		SELECT
			EXTRACT(EPOCH FROM b.time_stamp)::BIGINT as ts,
			b.tx_id,
			b.evtlog_id,
			o.token_id,
			o.price/1e+18 AS price,
			3 AS evt_type
		FROM rw_item_bought b
		JOIN rw_new_offer o ON (b.contract_aid=o.contract_aid) AND (b.offer_id=o.offer_id)
		WHERE (o.rwalk_aid=$1) AND (b.evtlog_id>$2)
		)
		) data
		ORDER BY evtlog_id`
	return queryList(ctx, r, "all events for notification since evtlog", 101, query, func(rows pgx.Rows, rec *rwp.NotificationEvent2) error {
		return rows.Scan(
			&rec.TimeStampMinted,
			&rec.TxId,
			&rec.EvtLogId,
			&rec.TokenId,
			&rec.Price,
			&rec.EvtType,
		)
	}, rwalkAid, startEvtlogID)
}

// AllEventsForNotificationMintsOnly is the mint-only variant of
// AllEventsForNotification (the offer/buy branches are intentionally
// disabled; kept for notification dry runs).
func (r *Repo) AllEventsForNotificationMintsOnly(ctx context.Context, rwalkAid, startTs int64) ([]rwp.NotificationEvent, error) {
	query := `SELECT
			ts,
			token_id,
			price,
			evt_type
		FROM (
		(
		SELECT
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,
			token_id,
			price/1e+18 AS price,
			1 AS evt_type
		FROM rw_mint_evt m
		WHERE (contract_aid=$1) AND (time_stamp > TO_TIMESTAMP($2))
		)
		) data
		ORDER BY ts`
	return queryList(ctx, r, "all events for notification (mints only)", 101, query, scanNotificationEvent, rwalkAid, startTs)
}

// ServerTimestamp returns the database server's current Unix timestamp.
func (r *Repo) ServerTimestamp(ctx context.Context) (int64, error) {
	var ts sql.NullInt64
	err := r.pool().QueryRow(ctx, "SELECT EXTRACT(EPOCH FROM now())::BIGINT AS ts").Scan(&ts)
	if err != nil {
		return 0, store.WrapError("server timestamp", err)
	}
	return ts.Int64, nil
}

// LastMintTimestamp returns the timestamp of the most recent mint, or 0 when
// no mints exist.
func (r *Repo) LastMintTimestamp(ctx context.Context) (int64, error) {
	var ts sql.NullInt64
	err := r.pool().QueryRow(ctx,
		"SELECT EXTRACT(EPOCH FROM time_stamp)::BIGINT AS ts FROM rw_mint_evt ORDER BY id DESC LIMIT 1").Scan(&ts)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, store.WrapError("last mint timestamp", err)
	}
	return ts.Int64, nil
}

// TokenTransfersByTxHash returns the RandomWalk token transfers contained in
// one transaction (rwctl verify-erc20-transfers uses it to cross-check
// chain data).
func (r *Repo) TokenTransfersByTxHash(ctx context.Context, txHash string) ([]rwp.TransferEntry, error) {
	query := `SELECT
			fa.addr,ta.addr,token_id
		FROM rw_transfer tr
			JOIN address fa ON tr.from_aid=fa.address_id
			JOIN address ta ON tr.to_aid=ta.address_id
			JOIN transaction tx ON tx.id=tr.tx_id
		WHERE tx.tx_hash=$1`
	return queryList(ctx, r, "token transfers by tx hash", 8, query, func(rows pgx.Rows, rec *rwp.TransferEntry) error {
		return rows.Scan(&rec.From, &rec.To, &rec.TokenId)
	}, txHash)
}
