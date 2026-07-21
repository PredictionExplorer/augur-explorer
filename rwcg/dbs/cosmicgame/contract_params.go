package cosmicgame

import (
	"database/sql"
	"fmt"
	"math/big"
)

// Get_glob_stats_cst_reward_for_bidding returns cg_glob_stats.cst_reward_for_bidding as a decimal string.
func (sw *SQLStorageWrapper) Get_glob_stats_cst_reward_for_bidding() (string, error) {
	var reward string
	err := sw.S.Db().QueryRow(
		"SELECT cst_reward_for_bidding FROM "+sw.S.SchemaName()+".cg_glob_stats LIMIT 1",
	).Scan(&reward)
	if err != nil {
		return "", err
	}
	return reward, nil
}

// Get_latest_decimal_param returns the latest value from an admin/history table column.
func (sw *SQLStorageWrapper) Get_latest_decimal_param(table, column string) (string, bool, error) {
	query := fmt.Sprintf(
		"SELECT %s FROM %s.%s ORDER BY id DESC LIMIT 1",
		column, sw.S.SchemaName(), table,
	)
	var val sql.NullString
	err := sw.S.Db().QueryRow(query).Scan(&val)
	if err == sql.ErrNoRows {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	if !val.Valid {
		return "", false, nil
	}
	return val.String, true, nil
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

// Insert_live_state_update_if_changed records an event-less state variable value in
// cg_live_state_updates. The row is skipped when the latest recorded value for
// (variable_name, round_num) already equals newValue, so startup backfills are
// idempotent. Returns true when a row was inserted.
//
// This table must only receive variables that do NOT emit a *Changed event on-chain;
// evented variables are tracked via evt_log and their cg_adm_* history tables.
func (sw *SQLStorageWrapper) Insert_live_state_update_if_changed(variableName string, contractAid, roundNum, blockNum int64, newValue string) (bool, error) {
	var latest sql.NullString
	query := "SELECT new_value FROM " + sw.S.SchemaName() + ".cg_live_state_updates " +
		"WHERE variable_name = $1 AND round_num = $2 ORDER BY id DESC LIMIT 1"
	err := sw.S.Db().QueryRow(query, variableName, roundNum).Scan(&latest)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == nil && latest.Valid && DecimalStringsEqual(latest.String, newValue) {
		return false, nil
	}
	insertQuery := "INSERT INTO " + sw.S.SchemaName() + ".cg_live_state_updates" +
		"(variable_name, contract_aid, round_num, block_num, new_value) VALUES($1,$2,$3,$4,$5)"
	if _, err := sw.S.Db().Exec(insertQuery, variableName, contractAid, roundNum, blockNum, newValue); err != nil {
		return false, err
	}
	return true, nil
}
