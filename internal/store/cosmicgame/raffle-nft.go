package cosmicgame

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

const raffleNFTWinnerColumns = `
	p.evtlog_id,
	p.block_num,
	t.id,
	t.tx_hash,
	EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
	p.time_stamp,
	p.winner_aid,
	wa.addr,
	p.round_num,
	p.token_id,
	p.cst_amount,
	p.cst_amount/1e18 cst_amount_eth,
	p.winner_idx,
	p.is_rwalk,
	p.is_staker
FROM cg_raffle_nft_prize p
	LEFT JOIN transaction t ON t.id=p.tx_id
	LEFT JOIN address wa ON p.winner_aid=wa.address_id`

func scanRaffleNFTWinner(rows pgx.Rows, rec *cgmodel.CGRaffleNFTWinnerRec) error {
	return rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&rec.RoundNum,
		&rec.TokenId,
		&rec.CstAmount,
		&rec.CstAmountEth,
		&rec.WinnerIndex,
		&rec.IsRWalk,
		&rec.IsStaker,
	)
}

// RaffleNFTWinnersByRound returns the raffle NFT winners of one round.
// isStaker selects the staker raffle (true) or the bidder raffle (false);
// the legacy version baked the flag into the SQL text, it is a bound
// parameter now.
func (r *Repo) RaffleNFTWinnersByRound(ctx context.Context, roundNum int64, isStaker bool) ([]cgmodel.CGRaffleNFTWinnerRec, error) {
	query := "SELECT " + raffleNFTWinnerColumns + `
		WHERE p.round_num=$1 AND p.is_staker=$2
		ORDER BY p.id DESC`
	return queryList(ctx, r, "raffle nft winners by round", 256, query, scanRaffleNFTWinner, roundNum, isStaker)
}

// RaffleNFTWinnerPageCursor identifies the last winner returned by
// RaffleNFTWinnersByRoundPage.
type RaffleNFTWinnerPageCursor struct {
	WinnerIndex int64
	EventLogID  int64
}

// RaffleNFTWinnersByRoundPage returns at most limit winners from one raffle
// pool after the supplied ascending keyset cursor. A nil cursor starts at the
// first winner index.
func (r *Repo) RaffleNFTWinnersByRoundPage(
	ctx context.Context,
	roundNum int64,
	isStaker bool,
	after *RaffleNFTWinnerPageCursor,
	limit int,
) (records []cgmodel.CGRaffleNFTWinnerRec, hasMore bool, err error) {
	const op = "raffle nft winners by round page"
	if roundNum < 0 {
		return nil, false, fmt.Errorf("%s: round must be non-negative", op)
	}
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: limit must be positive", op)
	}

	query := "SELECT " + raffleNFTWinnerColumns + `
		WHERE p.round_num=$1 AND p.is_staker=$2
		ORDER BY p.winner_idx, p.evtlog_id
		LIMIT $3`
	args := []any{roundNum, isStaker, limit + 1}
	if after != nil {
		if after.WinnerIndex < 0 || after.EventLogID < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		query = "SELECT " + raffleNFTWinnerColumns + `
			WHERE p.round_num=$1
				AND p.is_staker=$2
				AND (p.winner_idx, p.evtlog_id) > ($3, $4)
			ORDER BY p.winner_idx, p.evtlog_id
			LIMIT $5`
		args = []any{roundNum, isStaker, after.WinnerIndex, after.EventLogID, limit + 1}
	}

	records, err = queryList(ctx, r, op, limit+1, query, scanRaffleNFTWinner, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}

// RaffleNFTWinners returns raffle NFT winners across all rounds, newest
// first, offset/limit paginated (limit 0 means unbounded).
func (r *Repo) RaffleNFTWinners(ctx context.Context, offset, limit int) ([]cgmodel.CGRaffleNFTWinnerRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := "SELECT " + raffleNFTWinnerColumns + `
		ORDER BY p.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "raffle nft winners", 256, query, scanRaffleNFTWinner, offset, limit)
}
