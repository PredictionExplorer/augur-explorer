package cosmicgame

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// RoundsMissingChampionDurations returns claimed rounds whose event-less V3
// championDurations snapshot has not been stored. Pre-V3 rounds legitimately
// stay at zero and are harmlessly rechecked during startup recovery.
func (r *Repo) RoundsMissingChampionDurations(ctx context.Context) ([]int64, error) {
	const query = `SELECT pc.round_num
		FROM cg_prize_claim pc
		LEFT JOIN cg_round_stats rs ON rs.round_num=pc.round_num
		WHERE COALESCE(rs.endurance_champion_duration,0)=0
			AND COALESCE(rs.chrono_warrior_duration,0)=0
		ORDER BY pc.round_num`
	scan := func(rows pgx.Rows, round *int64) error { return rows.Scan(round) }
	return queryList(ctx, r, "rounds missing champion durations", 16, query, scan)
}

// UpdateRoundChampionDurations stores championDurations(roundNum), creating
// the aggregate row when a claim arrived without any earlier round activity.
func (r *Repo) UpdateRoundChampionDurations(
	ctx context.Context,
	roundNum, enduranceDuration, chronoDuration int64,
) error {
	if roundNum < 0 || enduranceDuration < 0 || chronoDuration < 0 {
		return errors.New("update champion durations: values must be non-negative")
	}
	_, err := r.q(ctx).Exec(ctx, `INSERT INTO cg_round_stats(
			round_num,endurance_champion_duration,chrono_warrior_duration
		) VALUES($1,$2,$3)
		ON CONFLICT (round_num) DO UPDATE SET
			endurance_champion_duration=EXCLUDED.endurance_champion_duration,
			chrono_warrior_duration=EXCLUDED.chrono_warrior_duration`,
		roundNum, enduranceDuration, chronoDuration)
	return store.WrapError("update champion durations", err)
}

// RoundChampionDurations reads the persisted V3 duration pair.
func (r *Repo) RoundChampionDurations(ctx context.Context, roundNum int64) (endurance, chrono int64, err error) {
	err = r.q(ctx).QueryRow(ctx, `SELECT endurance_champion_duration,chrono_warrior_duration
		FROM cg_round_stats WHERE round_num=$1`, roundNum).Scan(&endurance, &chrono)
	if err != nil {
		return 0, 0, store.WrapError("round champion durations", err)
	}
	return endurance, chrono, nil
}

// InsertLiveStateUpdateIfChanged appends an observation only when it differs
// from the latest value for the same variable and round.
func (r *Repo) InsertLiveStateUpdateIfChanged(
	ctx context.Context,
	variableName string,
	contractAid, roundNum, blockNum, timeStamp int64,
	newValue string,
) (bool, error) {
	if variableName == "" || contractAid < 1 || roundNum < -1 || blockNum < 0 || timeStamp < 0 {
		return false, errors.New("insert live state update: invalid identity")
	}
	var id int64
	err := r.q(ctx).QueryRow(ctx, `WITH latest AS (
			SELECT new_value
			FROM cg_live_state_updates
			WHERE variable_name=$1 AND round_num=$3
			ORDER BY id DESC
			LIMIT 1
		)
		INSERT INTO cg_live_state_updates(
			variable_name,contract_aid,round_num,block_num,time_stamp,new_value
		)
		SELECT $1,$2,$3,$4,TO_TIMESTAMP($6),$5
		WHERE NOT EXISTS (SELECT 1 FROM latest)
			OR (SELECT new_value FROM latest) <> $5
		RETURNING id`,
		variableName, contractAid, roundNum, blockNum, newValue, timeStamp).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, store.WrapError("insert live state update", err)
	}
	return true, nil
}

// LatestLiveStateUpdate returns the newest audited value for one variable
// and round, or store.ErrNotFound when it has never been observed.
func (r *Repo) LatestLiveStateUpdate(
	ctx context.Context,
	variableName string,
	roundNum int64,
) (cgmodel.CGLiveStateUpdate, error) {
	var rec cgmodel.CGLiveStateUpdate
	err := r.q(ctx).QueryRow(ctx, `SELECT
			id,variable_name,contract_aid,round_num,block_num,
			EXTRACT(EPOCH FROM time_stamp)::BIGINT,new_value
		FROM cg_live_state_updates
		WHERE variable_name=$1 AND round_num=$2
		ORDER BY id DESC LIMIT 1`,
		variableName, roundNum).Scan(
		&rec.ID,
		&rec.VariableName,
		&rec.ContractAid,
		&rec.RoundNum,
		&rec.BlockNum,
		&rec.TimeStamp,
		&rec.NewValue,
	)
	if err != nil {
		return cgmodel.CGLiveStateUpdate{}, store.WrapError("latest live state update", err)
	}
	return rec, nil
}

// BidRewardsByBidID returns the normalized current/previous recipient split
// in reward-type order.
func (r *Repo) BidRewardsByBidID(ctx context.Context, bidID int64) ([]cgmodel.CGBidReward, error) {
	if bidID < 1 {
		return nil, fmt.Errorf("bid rewards: invalid bid id %d", bidID)
	}
	const query = `SELECT r.id,r.evtlog_id,r.bid_id,r.round_num,r.recipient_aid,a.addr,r.reward_type,r.amount
		FROM cg_bid_reward r
		JOIN address a ON a.address_id=r.recipient_aid
		WHERE r.bid_id=$1
		ORDER BY r.reward_type`
	scan := func(rows pgx.Rows, rec *cgmodel.CGBidReward) error {
		return rows.Scan(
			&rec.ID,
			&rec.EvtLogID,
			&rec.BidID,
			&rec.RoundNum,
			&rec.RecipientAid,
			&rec.RecipientAddr,
			&rec.RewardType,
			&rec.Amount,
		)
	}
	return queryList(ctx, r, "bid rewards", 2, query, scan, bidID)
}
