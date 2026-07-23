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

// adminTableIdent guards the table/column names interpolated into the
// admin-parameter SQL below. Callers pass compile-time literals from the
// cg-etl drift registry; the check turns any future slip into a loud error
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
	err := r.q(ctx).QueryRow(ctx,
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
	err = r.q(ctx).QueryRow(ctx, query).Scan(&val)
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
