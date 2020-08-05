package dbs

import (
	"fmt"
	"os"
	"strings"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Check_main_stats() {

	var query string
	query="SELECT id FROM main_stats LIMIT 1";
	row := ss.db.QueryRow(query)
	var null_id sql.NullInt64
	var err error
	err=row.Scan(&null_id);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			query="INSERT INTO main_stats(universe_id) VALUES(1)";
			_,_ =ss.db.Exec(query)
		} else {
			ss.Log_msg(fmt.Sprintf("Error in check_main_stats(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_main_stats() p.MainStats {

	var query string
	query = "SELECT " +
				"markets_count," +
				"yesno_count," +
				"categ_count," +
				"scalar_count," +
				"active_count," +
				"money_at_stake," +
				"trades_count " +
			"FROM main_stats "

	row := ss.db.QueryRow(query)
	var err error
	var s p.MainStats
	err=row.Scan(
				&s.MarketsCount,
				&s.YesNoCount,
				&s.CategCount,
				&s.ScalarCount,
				&s.ActiveCount,
				&s.MoneyAtStake,
				&s.TradesCount,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v\n",err,query))
			os.Exit(1)
		}
	}
	s.FinalizedCount = (s.YesNoCount + s.CategCount + s.ScalarCount) - s.ActiveCount
	return s
}
func (ss *SQLStorage) Get_front_page_stats() p.FrontPageStats {

	var stats p.FrontPageStats
	var query string
	query = "SELECT markets_count,money_at_stake,trades_count " +
			"FROM main_stats WHERE universe_id=1"
	row := ss.db.QueryRow(query)
	err := row.Scan(
				&stats.MarketsCreated,
				&stats.MoneyAtStake,
				&stats.TradesCount,
	)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	return stats
}
func (ss *SQLStorage) Get_last_unique_addr_day() int64 {

	var query string
	query = "SELECT EXTRACT(EPOCH FROM day::TIMESTAMP)::BIGINT AS ts FROM unique_addrs ORDER BY day DESC LIMIT 1"
	row := ss.db.QueryRow(query)
	var ts int64
	var err error
	err=row.Scan(&ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_last_block_timestamp(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return ts
}
func (ss *SQLStorage) Update_unique_addresses_entry(ts int64,num_addrs int64) {
	var query string
	query = "UPDATE unique_addrs SET num_addrs = $2 WHERE day=to_timestamp($1)"
	res,err:=ss.db.Exec(query,ts,num_addrs)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Update_unique_addresses_entry() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Update_unique_addresses_entry(): %v",err))
		os.Exit(1)
	}
	if affected_rows == 0 {
		query = "INSERT INTO unique_addrs(day,num_addrs) VALUES(to_timestamp($1),$2)"
		_,err := ss.db.Exec(query,ts,num_addrs)
		if (err!=nil) {
			ss.Log_msg(
				fmt.Sprintf(
					"DB Error on INSERT in Update_unique_addresses_entry(): %v q=%v",
					err,query,
				),
			);
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_unique_addresses() []p.UniqueAddrEntry {

	records := make([]p.UniqueAddrEntry,0,365)
	var query string
	query = "SELECT " +
				"EXTRACT(EPOCH FROM day::TIMESTAMP)::BIGINT AS ts,"+
				"day," +
				"num_addrs "+
			"FROM unique_addrs ORDER BY day"
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var accumulator int64 = 0
	defer rows.Close()
	for rows.Next() {
		var rec p.UniqueAddrEntry
		err=rows.Scan(&rec.Ts,&rec.Day,&rec.NumAddrs)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		accumulator = accumulator + rec.NumAddrs
		rec.NumAddrsAccum = accumulator
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Calc_unique_addresses(ts_from int64,ts_to int64) (int64,bool) {

	var query string
	query = "SELECT count(*) FROM ( " +
				"SELECT DISTINCT u.eoa_aid FROM address a " +
				"JOIN ustats u ON u.eoa_aid=a.address_id " +
				"JOIN block b ON a.block_num=b.block_num " +
				"WHERE b.ts >= to_timestamp($1) AND b.ts < to_timestamp($2)" +
			") AS s"
	row := ss.db.QueryRow(query,ts_from,ts_to)
	var null_counter sql.NullInt64
	var err error
	err=row.Scan(&null_counter);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0,true	// this will never happen
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Calc_unique_addresses(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	query = "SELECT b.block_num FROM block b " +
			"WHERE b.ts >= to_timestamp($1) AND b.ts < to_timestamp($2)" +
			"LIMIT 1"
	row = ss.db.QueryRow(query,ts_from,ts_to)
	var no_rows bool = false
	var null_block_num sql.NullInt64
	err = row.Scan(&null_block_num)
	if err!=nil {
		if err == sql.ErrNoRows {
			no_rows=true
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Calc_unique_addresses(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}

	return null_counter.Int64,no_rows
}
func (ss *SQLStorage) Link_eoa_and_wallet_contract(eoa_aid, wallet_aid int64) {

	var query string
	query = "INSERT INTO ustats(eoa_aid,wallet_aid) " +
			"VALUES($1,$2) ON CONFLICT DO NOTHING"

	res,err:=ss.db.Exec(query,eoa_aid,wallet_aid)
	if (err!=nil) {
		if !strings.Contains(err.Error(),"duplicate key value") {
			ss.Log_msg(
				fmt.Sprintf(
					"Link_eoa_and_wallet_contract(%v,%v) failed: %v, q=%v",
					eoa_aid,wallet_aid,err,query,
				),
			)
			os.Exit(1)
		}
	} else {
		affected_rows,err:=res.RowsAffected()
		if err == nil {
			if affected_rows > 0 {
				ss.Info.Printf("eoa2wallet link success: eoa=%v wallet=%v\n",eoa_aid,wallet_aid)
			}
		}
	}
}
