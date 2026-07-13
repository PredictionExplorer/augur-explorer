package cosmicgame

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"regexp"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// AdminCorrectionMeta holds evt_log/block metadata for synthetic admin correction inserts.
type AdminCorrectionMeta struct {
	EvtID       int64
	BlockNum    int64
	TxID        int64
	TimeStamp   int64
	ContractAid int64
}

// adminTableIdent guards the table/column names interpolated into the
// admin-parameter SQL below. Callers pass compile-time literals from the
// cg-etl sync registry; the check turns any future slip into a loud error
// instead of an injection vector.
var adminTableIdent = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

func checkAdminIdent(kind, name string) error {
	if !adminTableIdent.MatchString(name) {
		return fmt.Errorf("invalid %s identifier %q", kind, name)
	}
	return nil
}

// GlobStatsCstRewardForBidding returns cg_glob_stats.cst_reward_for_bidding
// as a decimal string.
func (r *Repo) GlobStatsCstRewardForBidding(ctx context.Context) (string, error) {
	var reward string
	err := r.pool().QueryRow(ctx,
		"SELECT cst_reward_for_bidding FROM cg_glob_stats LIMIT 1",
	).Scan(&reward)
	if err != nil {
		return "", store.WrapError("glob stats cst reward for bidding", err)
	}
	return reward, nil
}

// LatestDecimalParam returns the latest value of an admin/history table
// column. hasRow is false when the table is empty or the latest value is
// NULL.
func (r *Repo) LatestDecimalParam(ctx context.Context, table, column string) (value string, hasRow bool, err error) {
	if err := checkAdminIdent("table", table); err != nil {
		return "", false, err
	}
	if err := checkAdminIdent("column", column); err != nil {
		return "", false, err
	}
	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT 1", column, table)
	var val *string
	err = r.pool().QueryRow(ctx, query).Scan(&val)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, store.WrapError("latest decimal param "+table+"."+column, err)
	}
	if val == nil {
		return "", false, nil
	}
	return *val, true, nil
}

// DecimalStringsEqual compares two base-10 integer strings (wei or seconds).
func DecimalStringsEqual(a, b string) bool {
	av, okA := new(big.Int).SetString(a, 10)
	bv, okB := new(big.Int).SetString(b, 10)
	if !okA || !okB {
		return a == b
	}
	return av.Cmp(bv) == 0
}

// InsertAdminCorrectionDecimal inserts a correction row into an
// admin/history table. contractAid 0 means the row belongs to the game
// contract of meta; pass a non-zero aid to attribute it to a peripheral
// contract instead.
func (r *Repo) InsertAdminCorrectionDecimal(ctx context.Context, table, column, value string, meta *AdminCorrectionMeta, contractAid int64) error {
	if err := checkAdminIdent("table", table); err != nil {
		return err
	}
	if err := checkAdminIdent("column", column); err != nil {
		return err
	}
	if contractAid == 0 {
		contractAid = meta.ContractAid
	}
	query := fmt.Sprintf(
		"INSERT INTO %s (evtlog_id, block_num, tx_id, time_stamp, contract_aid, %s) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6)",
		table, column,
	)
	_, err := r.pool().Exec(ctx, query,
		meta.EvtID,
		meta.BlockNum,
		meta.TxID,
		meta.TimeStamp,
		contractAid,
		value,
	)
	if err != nil {
		return store.WrapError("insert admin correction "+table+"."+column, err)
	}
	return nil
}

// InsertAdminCorrectionERC20Reward inserts a cg_adm_erc20_reward row (the
// insert trigger updates cg_glob_stats.cst_reward_for_bidding).
func (r *Repo) InsertAdminCorrectionERC20Reward(ctx context.Context, reward string, meta *AdminCorrectionMeta) error {
	query := "INSERT INTO cg_adm_erc20_reward(" +
		"evtlog_id, block_num, tx_id, time_stamp, contract_aid, new_reward" +
		") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6)"
	_, err := r.pool().Exec(ctx, query,
		meta.EvtID,
		meta.BlockNum,
		meta.TxID,
		meta.TimeStamp,
		meta.ContractAid,
		reward,
	)
	if err != nil {
		return store.WrapError("insert admin correction erc20 reward", err)
	}
	return nil
}
