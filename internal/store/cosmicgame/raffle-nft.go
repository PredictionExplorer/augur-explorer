package cosmicgame

import (
	"context"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
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

func scanRaffleNFTWinner(rows pgx.Rows, rec *p.CGRaffleNFTWinnerRec) error {
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
func (r *Repo) RaffleNFTWinnersByRound(ctx context.Context, roundNum int64, isStaker bool) ([]p.CGRaffleNFTWinnerRec, error) {
	query := "SELECT " + raffleNFTWinnerColumns + `
		WHERE p.round_num=$1 AND p.is_staker=$2
		ORDER BY p.id DESC`
	return queryList(ctx, r, "raffle nft winners by round", 256, query, scanRaffleNFTWinner, roundNum, isStaker)
}

// RaffleNFTWinners returns raffle NFT winners across all rounds, newest
// first, offset/limit paginated (limit 0 means unbounded).
func (r *Repo) RaffleNFTWinners(ctx context.Context, offset, limit int) ([]p.CGRaffleNFTWinnerRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := "SELECT " + raffleNFTWinnerColumns + `
		ORDER BY p.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "raffle nft winners", 256, query, scanRaffleNFTWinner, offset, limit)
}
