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

// LatestAddressParam returns the address (0x-hex) the latest row of an
// address-change table points at through its aidColumn. hasRow is false
// when the table is empty.
func (r *Repo) LatestAddressParam(ctx context.Context, table, aidColumn string) (addr string, hasRow bool, err error) {
	if err := checkAdminIdent("table", table); err != nil {
		return "", false, err
	}
	if err := checkAdminIdent("column", aidColumn); err != nil {
		return "", false, err
	}
	query := fmt.Sprintf(
		"SELECT a.addr FROM %s r JOIN address a ON a.address_id = r.%s ORDER BY r.id DESC LIMIT 1",
		table, aidColumn)
	err = r.q(ctx).QueryRow(ctx, query).Scan(&addr)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, store.WrapError("latest address param "+table+"."+aidColumn, err)
	}
	return addr, true, nil
}

// UpgradedImplementations returns every implementation address an indexed
// ERC-1967 Upgraded event pointed the game proxy at, oldest first.
func (r *Repo) UpgradedImplementations(ctx context.Context) ([]string, error) {
	rows, err := r.q(ctx).Query(ctx,
		"SELECT a.addr FROM cg_adm_upgraded u JOIN address a ON a.address_id = u.implementation_aid "+
			"GROUP BY a.addr ORDER BY MIN(u.block_num), a.addr")
	if err != nil {
		return nil, store.WrapError("upgraded implementations", err)
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			return nil, store.WrapError("upgraded implementations", err)
		}
		out = append(out, addr)
	}
	return out, store.WrapError("upgraded implementations", rows.Err())
}

// HistoricPrizesWallets returns every prizes-wallet address an indexed
// PrizesWalletAddressChanged event pointed the game at, oldest first. The
// wallet configured at deployment never emitted that event; callers union
// this list with cg_contracts.prizes_wallet_addr.
func (r *Repo) HistoricPrizesWallets(ctx context.Context) ([]string, error) {
	rows, err := r.q(ctx).Query(ctx,
		"SELECT a.addr FROM cg_adm_prizes_wallet_addr w JOIN address a ON a.address_id = w.new_wallet_aid "+
			"GROUP BY a.addr ORDER BY MIN(w.block_num), a.addr")
	if err != nil {
		return nil, store.WrapError("historic prizes wallets", err)
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			return nil, store.WrapError("historic prizes wallets", err)
		}
		out = append(out, addr)
	}
	return out, store.WrapError("historic prizes wallets", rows.Err())
}

// ImplementationRef is an implementation address with the block of the
// first Upgraded event that introduced it.
type ImplementationRef struct {
	Addr     string
	BlockNum int64
}

// ImplementationsMissingInitialized returns the implementations introduced
// by an Upgraded event whose constructor Initialized event has no
// cg_adm_initialized row (the ETL was not watching the address when the
// deploy transaction was indexed).
// PrizesWalletAnnouncements returns every prizes wallet the indexed
// PrizesWalletAddressChanged history introduced, with the block of its first
// announcement — the point from which the wallet may start emitting events.
// cg-etl uses it for the startup backfill recovery pass.
func (r *Repo) PrizesWalletAnnouncements(ctx context.Context) ([]ImplementationRef, error) {
	rows, err := r.q(ctx).Query(ctx,
		"SELECT a.addr, MIN(w.block_num) FROM cg_adm_prizes_wallet_addr w "+
			"JOIN address a ON a.address_id = w.new_wallet_aid "+
			"GROUP BY a.addr ORDER BY MIN(w.block_num), a.addr")
	if err != nil {
		return nil, store.WrapError("prizes wallet announcements", err)
	}
	defer rows.Close()
	var out []ImplementationRef
	for rows.Next() {
		var ref ImplementationRef
		if err := rows.Scan(&ref.Addr, &ref.BlockNum); err != nil {
			return nil, store.WrapError("prizes wallet announcements", err)
		}
		out = append(out, ref)
	}
	return out, store.WrapError("prizes wallet announcements", rows.Err())
}

func (r *Repo) ImplementationsMissingInitialized(ctx context.Context) ([]ImplementationRef, error) {
	rows, err := r.q(ctx).Query(ctx,
		"SELECT a.addr, MIN(u.block_num) FROM cg_adm_upgraded u "+
			"JOIN address a ON a.address_id = u.implementation_aid "+
			"WHERE NOT EXISTS (SELECT 1 FROM cg_adm_initialized i WHERE i.contract_aid = u.implementation_aid) "+
			"GROUP BY a.addr ORDER BY MIN(u.block_num), a.addr")
	if err != nil {
		return nil, store.WrapError("implementations missing initialized", err)
	}
	defer rows.Close()
	var out []ImplementationRef
	for rows.Next() {
		var ref ImplementationRef
		if err := rows.Scan(&ref.Addr, &ref.BlockNum); err != nil {
			return nil, store.WrapError("implementations missing initialized", err)
		}
		out = append(out, ref)
	}
	return out, store.WrapError("implementations missing initialized", rows.Err())
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
