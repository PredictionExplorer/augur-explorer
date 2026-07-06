package cosmicgame

import (
	"database/sql"
	"fmt"
	"math/big"
)

// AdminCorrectionMeta holds evt_log/block metadata for synthetic admin correction inserts.
type AdminCorrectionMeta struct {
	EvtId       int64
	BlockNum    int64
	TxId        int64
	TimeStamp   int64
	ContractAid int64
}

// Get_glob_stats_cst_reward_for_bidding returns cg_glob_stats.cst_reward_for_bidding as a decimal string.
func (sw *SQLStorageWrapper) Get_glob_stats_cst_reward_for_bidding() (string, error) {
	var reward string
	err := sw.S.Db().QueryRow(
		"SELECT cst_reward_for_bidding FROM " + sw.S.SchemaName() + ".cg_glob_stats LIMIT 1",
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

func (sw *SQLStorageWrapper) contractAidForCorrection(meta *AdminCorrectionMeta, contractAddr string, blockNum, txId int64) (int64, error) {
	if contractAddr == "" {
		return meta.ContractAid, nil
	}
	aid, err := sw.S.Lookup_or_create_address(contractAddr, blockNum, txId)
	if err != nil {
		return 0, fmt.Errorf("correction contract address %v: %w", contractAddr, err)
	}
	return aid, nil
}

// Insert_admin_correction_decimal_row inserts a correction row into an admin/history table.
func (sw *SQLStorageWrapper) Insert_admin_correction_decimal_row(table, column, value string, meta *AdminCorrectionMeta) error {
	return sw.Insert_admin_correction_decimal_row_for_contract(table, column, value, meta, "")
}

// Insert_admin_correction_decimal_row_for_contract inserts a correction row for a peripheral contract address.
func (sw *SQLStorageWrapper) Insert_admin_correction_decimal_row_for_contract(table, column, value string, meta *AdminCorrectionMeta, contractAddr string) error {
	contractAid, err := sw.contractAidForCorrection(meta, contractAddr, meta.BlockNum, meta.TxId)
	if err != nil {
		return err
	}
	query := fmt.Sprintf(
		"INSERT INTO %s.%s (evtlog_id, block_num, tx_id, time_stamp, contract_aid, %s) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6)",
		sw.S.SchemaName(), table, column,
	)
	_, err = sw.S.Db().Exec(query,
		meta.EvtId,
		meta.BlockNum,
		meta.TxId,
		meta.TimeStamp,
		contractAid,
		value,
	)
	return err
}

// SyncAdminDecimalParamIfMismatch compares latest SQL value with wantValue; inserts on mismatch.
func (sw *SQLStorageWrapper) SyncAdminDecimalParamIfMismatch(name, table, column, wantValue string, meta *AdminCorrectionMeta) (bool, error) {
	return sw.SyncAdminDecimalParamIfMismatchForContract(name, table, column, wantValue, "", meta)
}

// SyncAdminDecimalParamIfMismatchForContract syncs a param stored against an optional non-game contract address.
func (sw *SQLStorageWrapper) SyncAdminDecimalParamIfMismatchForContract(name, table, column, wantValue, contractAddr string, meta *AdminCorrectionMeta) (bool, error) {
	dbValue, hasRow, err := sw.Get_latest_decimal_param(table, column)
	if err != nil {
		return false, err
	}
	if hasRow && DecimalStringsEqual(dbValue, wantValue) {
		return false, nil
	}
	if err := sw.Insert_admin_correction_decimal_row_for_contract(table, column, wantValue, meta, contractAddr); err != nil {
		return false, fmt.Errorf("%s: %w", name, err)
	}
	return true, nil
}

// SyncAdminInt64ParamIfMismatch syncs integer admin params stored as DECIMAL/BIGINT columns.
func (sw *SQLStorageWrapper) SyncAdminInt64ParamIfMismatch(name, table, column string, wantValue int64, meta *AdminCorrectionMeta) (bool, error) {
	return sw.SyncAdminInt64ParamIfMismatchForContract(name, table, column, wantValue, meta, "")
}

// SyncAdminInt64ParamIfMismatchForContract syncs integer params for an optional peripheral contract.
func (sw *SQLStorageWrapper) SyncAdminInt64ParamIfMismatchForContract(name, table, column string, wantValue int64, meta *AdminCorrectionMeta, contractAddr string) (bool, error) {
	dbStr, hasRow, err := sw.Get_latest_decimal_param(table, column)
	if err != nil {
		return false, err
	}
	if hasRow {
		dbInt, ok := new(big.Int).SetString(dbStr, 10)
		if ok && dbInt.Int64() == wantValue {
			return false, nil
		}
	}
	valStr := fmt.Sprintf("%d", wantValue)
	if err := sw.Insert_admin_correction_decimal_row_for_contract(table, column, valStr, meta, contractAddr); err != nil {
		return false, fmt.Errorf("%s: %w", name, err)
	}
	return true, nil
}

// Insert_admin_correction_erc20_reward inserts cg_adm_erc20_reward (trigger updates cg_glob_stats).
func (sw *SQLStorageWrapper) Insert_admin_correction_erc20_reward(reward string, meta *AdminCorrectionMeta) error {
	query := "INSERT INTO " + sw.S.SchemaName() + ".cg_adm_erc20_reward(" +
		"evtlog_id, block_num, tx_id, time_stamp, contract_aid, new_reward" +
		") VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6)"
	_, err := sw.S.Db().Exec(query,
		meta.EvtId,
		meta.BlockNum,
		meta.TxId,
		meta.TimeStamp,
		meta.ContractAid,
		reward,
	)
	return err
}

// SyncCstRewardIfMismatch aligns cg_adm_erc20_reward and cg_glob_stats with wantValue.
func (sw *SQLStorageWrapper) SyncCstRewardIfMismatch(wantValue string, meta *AdminCorrectionMeta) (bool, error) {
	globReward, err := sw.Get_glob_stats_cst_reward_for_bidding()
	if err != nil {
		return false, err
	}
	dbLatest, hasRow, err := sw.Get_latest_decimal_param("cg_adm_erc20_reward", "new_reward")
	if err != nil {
		return false, err
	}
	if hasRow && DecimalStringsEqual(dbLatest, wantValue) && DecimalStringsEqual(globReward, wantValue) {
		return false, nil
	}
	if err := sw.Insert_admin_correction_erc20_reward(wantValue, meta); err != nil {
		return false, fmt.Errorf("cst_reward_for_bidding: %w", err)
	}
	return true, nil
}
