package cosmicgame

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func scanCosmicTokenHolder(rows pgx.Rows, rec *p.CGCosmicTokenHolderRec) error {
	return rows.Scan(
		&rec.OwnerAid,
		&rec.OwnerAddr,
		&rec.Balance,
		&rec.BalanceFloat,
	)
}

// CosmicTokenHolders returns every CST (ERC-20) balance row ordered by
// balance descending.
func (r *Repo) CosmicTokenHolders(ctx context.Context) ([]p.CGCosmicTokenHolderRec, error) {
	query := `SELECT
			o.owner_aid,
			oa.addr,
			o.cur_balance,
			o.cur_balance/1e18
		FROM cg_costok_owner o
			LEFT JOIN address oa ON o.owner_aid=oa.address_id
		ORDER BY o.cur_balance DESC`
	return queryList(ctx, r, "cosmic token holders", 16, query, scanCosmicTokenHolder)
}

// CosmicTokenStatistics aggregates the CST (ERC-20) view of the game: total
// supply, how tokens enter (bidding rewards, marketing, prizes) and leave
// (CST bids), transfer counts, and the top-10 holders with their share of
// supply.
func (r *Repo) CosmicTokenStatistics(ctx context.Context) (p.CGCosmicTokenStats, error) {
	const op = "cosmic token statistics"
	var stats p.CGCosmicTokenStats

	err := r.pool().QueryRow(ctx, `SELECT
			COUNT(*) as holder_count,
			COALESCE(SUM(cur_balance), 0) as total_supply,
			COALESCE(SUM(cur_balance)/1e18, 0) as total_supply_eth
		FROM cg_costok_owner
		WHERE cur_balance > 0`).
		Scan(&stats.TotalHolders, &stats.TotalSupply, &stats.TotalSupplyEth)
	if err != nil {
		return stats, store.WrapError(op+": total supply", err)
	}

	// Bidding rewards (CST ERC20 given for bidding).
	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_reward), 0), COALESCE(SUM(cst_reward)/1e18, 0) FROM cg_bid WHERE cst_reward > 0").
		Scan(&stats.EarnedFromBidding, &stats.EarnedFromBiddingEth)
	if err != nil {
		return stats, store.WrapError(op+": bidding rewards", err)
	}

	// Marketing rewards (CST ERC20).
	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(amount), 0), COALESCE(SUM(amount)/1e18, 0) FROM cg_mkt_reward").
		Scan(&stats.DistributedToMarketers, &stats.DistributedToMarketersEth)
	if err != nil {
		return stats, store.WrapError(op+": marketing rewards", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM cg_prize_claim WHERE cst_amount > 0").
		Scan(&stats.GivenAsMainPrizes, &stats.GivenAsMainPrizesEth)
	if err != nil {
		return stats, store.WrapError(op+": main prizes", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM cg_raffle_nft_prize WHERE cst_amount > 0").
		Scan(&stats.GivenAsRafflePrizes, &stats.GivenAsRafflePrizesEth)
	if err != nil {
		return stats, store.WrapError(op+": raffle prizes", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM cg_chrono_warrior_prize WHERE cst_amount > 0").
		Scan(&stats.GivenAsChronoWarriorPrizes, &stats.GivenAsChronoWarriorPrizesEth)
	if err != nil {
		return stats, store.WrapError(op+": chrono warrior prizes", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(total_cst_consumed, 0), COALESCE(total_cst_consumed/1e18, 0) FROM cg_glob_stats LIMIT 1").
		Scan(&stats.ConsumedInBids, &stats.ConsumedInBidsEth)
	if err != nil {
		return stats, store.WrapError(op+": consumed in bids", err)
	}

	err = r.pool().QueryRow(ctx, `SELECT
			COUNT(CASE WHEN otype = 1 THEN 1 END) as mints,
			COUNT(CASE WHEN otype = 2 THEN 1 END) as burns,
			COUNT(CASE WHEN otype = 0 THEN 1 END) as transfers
		FROM cg_erc20_transfer`).
		Scan(&stats.TotalMints, &stats.TotalBurns, &stats.TotalTransfers)
	if err != nil {
		return stats, store.WrapError(op+": transfer counts", err)
	}

	topHolders, err := queryList(ctx, r, op+": top holders", 10, `SELECT
			o.owner_aid,
			a.addr,
			o.cur_balance,
			o.cur_balance/1e18
		FROM cg_costok_owner o
			LEFT JOIN address a ON o.owner_aid = a.address_id
		WHERE o.cur_balance > 0
		ORDER BY o.cur_balance DESC
		LIMIT 10`, scanCosmicTokenHolder)
	if err != nil {
		return stats, err
	}
	for i := range topHolders {
		if stats.TotalSupplyEth > 0 {
			topHolders[i].PercentOfSupply = (topHolders[i].BalanceFloat / stats.TotalSupplyEth) * 100.0
		}
	}
	stats.TopHolders = topHolders

	return stats, nil
}

// UserCosmicTokenSummary aggregates one user's CST (ERC-20) position: current
// balance, earnings by source, amounts consumed in CST bids and transfer
// activity counts. A user without a balance row reports a zero balance.
func (r *Repo) UserCosmicTokenSummary(ctx context.Context, userAid int64) (p.CGUserCosmicTokenSummary, error) {
	const op = "user cosmic token summary"
	var summary p.CGUserCosmicTokenSummary
	summary.UserAid = userAid

	err := r.pool().QueryRow(ctx,
		"SELECT COALESCE(cur_balance, 0), COALESCE(cur_balance/1e18, 0) FROM cg_costok_owner WHERE owner_aid=$1", userAid).
		Scan(&summary.CurrentBalance, &summary.CurrentBalanceEth)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return summary, store.WrapError(op+": balance", err)
		}
		summary.CurrentBalance = "0"
		summary.CurrentBalanceEth = 0
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_reward), 0), COALESCE(SUM(cst_reward)/1e18, 0) FROM cg_bid WHERE bidder_aid=$1 AND cst_reward > 0", userAid).
		Scan(&summary.EarnedFromBidding, &summary.EarnedFromBiddingEth)
	if err != nil {
		return summary, store.WrapError(op+": bidding rewards", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM cg_prize_claim WHERE winner_aid=$1 AND cst_amount > 0", userAid).
		Scan(&summary.EarnedFromMainPrizes, &summary.EarnedFromMainPrizesEth)
	if err != nil {
		return summary, store.WrapError(op+": main prizes", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM cg_raffle_nft_prize WHERE winner_aid=$1 AND cst_amount > 0", userAid).
		Scan(&summary.EarnedFromRafflePrizes, &summary.EarnedFromRafflePrizesEth)
	if err != nil {
		return summary, store.WrapError(op+": raffle prizes", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM cg_chrono_warrior_prize WHERE winner_aid=$1 AND cst_amount > 0", userAid).
		Scan(&summary.EarnedFromChronoWarrior, &summary.EarnedFromChronoWarriorEth)
	if err != nil {
		return summary, store.WrapError(op+": chrono warrior prizes", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(amount), 0), COALESCE(SUM(amount)/1e18, 0) FROM cg_mkt_reward WHERE marketer_aid=$1", userAid).
		Scan(&summary.EarnedFromMarketing, &summary.EarnedFromMarketingEth)
	if err != nil {
		return summary, store.WrapError(op+": marketing rewards", err)
	}

	err = r.pool().QueryRow(ctx,
		"SELECT COALESCE(SUM(cst_price), 0), COALESCE(SUM(cst_price)/1e18, 0) FROM cg_bid WHERE bidder_aid=$1 AND cst_price > 0", userAid).
		Scan(&summary.ConsumedInBids, &summary.ConsumedInBidsEth)
	if err != nil {
		return summary, store.WrapError(op+": consumed in bids", err)
	}

	summary.TotalEarnedEth = summary.EarnedFromBiddingEth +
		summary.EarnedFromMainPrizesEth +
		summary.EarnedFromRafflePrizesEth +
		summary.EarnedFromChronoWarriorEth +
		summary.EarnedFromMarketingEth
	summary.NetCSTFlowEth = summary.TotalEarnedEth - summary.ConsumedInBidsEth

	err = r.pool().QueryRow(ctx, `SELECT
			COUNT(CASE WHEN otype = 1 THEN 1 END) as mints,
			COUNT(CASE WHEN otype = 2 THEN 1 END) as burns,
			COUNT(*) as total_transfers
		FROM cg_erc20_transfer
		WHERE from_aid=$1 OR to_aid=$1`, userAid).
		Scan(&summary.NumMints, &summary.NumBurns, &summary.NumTransfers)
	if err != nil {
		return summary, store.WrapError(op+": activity counts", err)
	}

	return summary, nil
}

// CosmicTokenSupplyHistoryByBid returns one row per bid with the net CST
// supply change (cst_reward mint minus cst_price burn on CST bids) and
// running totals computed in SQL.
func (r *Repo) CosmicTokenSupplyHistoryByBid(ctx context.Context) ([]p.CGTotalSupplyHistoryRec, error) {
	query := `SELECT
		b.evtlog_id, b.bid_type, COALESCE(ba.addr, ''), b.block_num, COALESCE(t.id, 0), COALESCE(t.tx_hash, ''),
		EXTRACT(EPOCH FROM b.time_stamp)::BIGINT, b.time_stamp,
		GREATEST(COALESCE(b.cst_reward, 0), 0)::text,
		GREATEST(COALESCE(b.cst_reward, 0), 0)/1e18,
		(CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)::text,
		(CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)/1e18,
		(GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)::text,
		(GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)/1e18,
		SUM(GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)
		OVER (ORDER BY b.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)::text,
		SUM((GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)/1e18)
		OVER (ORDER BY b.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)
		FROM cg_bid b
		LEFT JOIN address ba ON b.bidder_aid = ba.address_id
		LEFT JOIN transaction t ON t.id = b.tx_id
		ORDER BY b.id`
	scan := func(rows pgx.Rows, rec *p.CGTotalSupplyHistoryRec) error {
		err := rows.Scan(
			&rec.BidInfoId,
			&rec.BidType,
			&rec.BidderAddr,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.MintAmount,
			&rec.MintAmountEth,
			&rec.BurnAmount,
			&rec.BurnAmountEth,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TotalSupply,
			&rec.TotalSupplyEth,
		)
		if err != nil {
			return err
		}
		rec.Tx.EvtLogId = rec.BidInfoId
		return nil
	}
	return queryList(ctx, r, "cosmic token supply history by bid", 256, query, scan)
}

// CosmicTokenSupplyHistoryByDate returns daily aggregates of CST supply
// change between fromDate and toDate (inclusive, YYYYMMDD), with running
// totals over all history up to each day.
func (r *Repo) CosmicTokenSupplyHistoryByDate(ctx context.Context, fromDate, toDate string) ([]p.CGTotalSupplyHistoryByDateRec, error) {
	query := `WITH daily AS (
		SELECT
		DATE(b.time_stamp) AS bid_date,
		COUNT(*)::bigint AS num_bids,
		SUM(GREATEST(COALESCE(b.cst_reward, 0), 0)) AS mint_amt,
		SUM(CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END) AS burn_amt,
		SUM(GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END) AS net_amt
		FROM cg_bid b
		GROUP BY DATE(b.time_stamp)
		), with_totals AS (
		SELECT
		d.bid_date,
		d.num_bids,
		d.mint_amt, d.burn_amt, d.net_amt,
		SUM(d.net_amt) OVER (ORDER BY d.bid_date ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS total_supply
		FROM daily d
		)
		SELECT
		TO_CHAR(w.bid_date, 'YYYYMMDD'),
		EXTRACT(EPOCH FROM w.bid_date)::BIGINT,
		w.bid_date::text,
		w.num_bids,
		w.mint_amt::text, w.mint_amt/1e18,
		w.burn_amt::text, w.burn_amt/1e18,
		w.net_amt::text, w.net_amt/1e18,
		w.total_supply::text, w.total_supply/1e18
		FROM with_totals w
		WHERE w.bid_date >= TO_DATE($1, 'YYYYMMDD') AND w.bid_date <= TO_DATE($2, 'YYYYMMDD')
		ORDER BY w.bid_date`
	scan := func(rows pgx.Rows, rec *p.CGTotalSupplyHistoryByDateRec) error {
		return rows.Scan(
			&rec.Date,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.NumBids,
			&rec.MintAmount,
			&rec.MintAmountEth,
			&rec.BurnAmount,
			&rec.BurnAmountEth,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TotalSupply,
			&rec.TotalSupplyEth,
		)
	}
	return queryList(ctx, r, "cosmic token supply history by date", 64, query, scan, fromDate, toDate)
}
