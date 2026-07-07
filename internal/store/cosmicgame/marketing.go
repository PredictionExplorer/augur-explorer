package cosmicgame

import (
	"context"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

const marketingRewardColumns = `
	r.id,
	r.evtlog_id,
	r.block_num,
	tx.id,
	tx.tx_hash,
	EXTRACT(EPOCH FROM r.time_stamp)::BIGINT,
	r.time_stamp,
	r.amount,
	r.amount/1e18,
	r.marketer_aid,
	ma.addr
FROM cg_mkt_reward r
	LEFT JOIN transaction tx ON tx.id=r.tx_id
	LEFT JOIN address ma ON r.marketer_aid=ma.address_id`

func scanMarketingReward(rows pgx.Rows, rec *p.CGMarketingRewardRec) error {
	return rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.Amount,
		&rec.AmountEth,
		&rec.MarketerAid,
		&rec.MarketerAddr,
	)
}

// MarketingRewardHistoryGlobal returns marketing reward events across all
// marketers, newest first, paginated by offset/limit.
func (r *Repo) MarketingRewardHistoryGlobal(ctx context.Context, offset, limit int) ([]p.CGMarketingRewardRec, error) {
	query := "SELECT " + marketingRewardColumns + `
		ORDER BY r.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "marketing reward history global", 16, query, scanMarketingReward, offset, limit)
}

// MarketingRewardsByUser returns every marketing reward paid to one marketer,
// newest first.
func (r *Repo) MarketingRewardsByUser(ctx context.Context, marketerAid int64) ([]p.CGMarketingRewardRec, error) {
	query := "SELECT " + marketingRewardColumns + `
		WHERE r.marketer_aid=$1
		ORDER BY r.id DESC`
	return queryList(ctx, r, "marketing rewards by user", 16, query, scanMarketingReward, marketerAid)
}
