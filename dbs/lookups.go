package dbs

import (
	"fmt"
	"os"
	"runtime/debug"
	"database/sql"
	_  "github.com/lib/pq"

)
func (ss *SQLStorage) lookup_universe_id(addr string) (int64,error) {

	return ss.Nonfatal_lookup_address_id(addr)
}
func (ss *SQLStorage) Lookup_eoa_aid(wallet_aid int64) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT eoa_aid FROM eoa_wallet WHERE wallet_aid=$1"
	err:=ss.db.QueryRow(query,wallet_aid).Scan(&addr_id);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("lookup_eoa_aid(wallet_aid=%v) sql error=%v\n",wallet_aid,err))
			os.Exit(1)
		}
		return 0,err
	}
	return addr_id,nil
}
func (ss *SQLStorage) Lookup_wallet_aid(eoa_aid int64) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT wallet_aid FROM eoa_wallet WHERE eoa_aid=$1"
	err:=ss.db.QueryRow(query,eoa_aid).Scan(&addr_id);
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		} else {
		}
		return 0,err
	}

	return addr_id,nil
}
func (ss *SQLStorage) Nonfatal_lookup_address_id(addr string) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT address_id FROM address WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v ,q=%v",query))
			os.Exit(1)
		}
		return 0,err
	}

	return addr_id,nil
}
func (ss *SQLStorage) Lookup_market_by_addr_str(addr string) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT address_id FROM address AS a,market AS m WHERE m.market_aid=a.address_id AND a.addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
	if (err!=nil) {
		if err!=sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v ,q=%v",query))
			os.Exit(1)
		}
		return 0,err
	}

	return addr_id,nil
}
func (ss *SQLStorage) Lookup_address(aid int64) (string,error) {

	var addr string;
	var query string
	query="SELECT addr FROM address WHERE address_id=$1"
	err:=ss.db.QueryRow(query,aid).Scan(&addr);
	return addr,err
}
func (ss *SQLStorage) Lookup_address_id(addr string) int64 {

	var addr_id int64 = 0;
	var query string
	query="SELECT address_id FROM address WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			ss.Log_msg(fmt.Sprintf("Forced address lookup failed for %v : addr not found",addr))
			debug.PrintStack()
			os.Exit(1)
		} else {
			ss.Log_msg(fmt.Sprintf("DB error upon address lookup: %v ; q=%v",err,query))
			os.Exit(1)
		}
	}

	return addr_id
}
func (ss *SQLStorage) lookup_market_id(addr string) (int64,int) {
	//Return values: market_aid & market_type
	var addr_id int64;
	var mkt_type int;
	var query string
	query=	"SELECT m.market_type,a.address_id FROM address a,market m " +
			"WHERE m.market_aid=a.address_id AND a.addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&mkt_type,&addr_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error upon address lookup: %v",err))
			os.Exit(1)
		}
	}

	return addr_id,mkt_type
}
func (ss *SQLStorage) Lookup_or_create_address(addr string,block_num int64,tx_id int64) int64 {

	var addr_id int64;
	var query string
	query="SELECT address_id FROM address WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			addr_id = ss.create_address(addr,block_num,tx_id)
			return addr_id
		} else {
			ss.Log_msg(fmt.Sprintf("DB error in getting address id : %v",err))
		}
	}

	return addr_id
}
func (ss *SQLStorage) create_address(addr string,block_num int64,tx_id int64) int64 {

	var addr_id int64;
	var query string
	if len(addr) == 0 {
		ss.Log_msg(fmt.Sprintf("Attempt to insert address with len=0, block %v, tx_id=%v",block_num,tx_id))
		os.Exit(1)
	}
	query = "INSERT INTO address(addr,block_num,tx_id) VALUES($1,$2,$3) RETURNING address_id"
	row:=ss.db.QueryRow(query,addr,block_num,tx_id);
	err:=row.Scan(&addr_id)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error in address insertion: %v : %v",query,err))
		os.Exit(1)
	}
	if addr_id==0 {
		ss.Log_msg(fmt.Sprintf("DB error, addr_id after INSERT is 0"))
		os.Exit(1)
	}
	ss.Info.Printf("create_address: %v    aid=%v\n",addr,addr_id)

	return addr_id
}
func (ss *SQLStorage) lookup_or_create_category(categories string) int64 {

	var cat_id int64
	var query string

	query="SELECT cat_id FROM category WHERE category=$1"
	err:=ss.db.QueryRow(query,categories).Scan(&cat_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			query = "INSERT INTO category(category) VALUES($1) RETURNING cat_id"
			row:=ss.db.QueryRow(query,categories);
			err:=row.Scan(&cat_id)
			if err!=nil {
				ss.Log_msg(fmt.Sprintf("DB error in category insertion: %v : %v",query,err))
				os.Exit(1)
			}
			if cat_id==0 {
				ss.Log_msg(fmt.Sprintf("DB error, cat_id after INSERT is 0"))
				os.Exit(1)
			}
			return cat_id
		} else {
			ss.Log_msg(fmt.Sprintf("DB error, cat_id after INSERT is 0"))
			os.Exit(1)
		}
	}

	return cat_id
}
func (ss *SQLStorage) Ustats_entry_exists(aid int64) (bool,error) {

	var addr_id int64;
	var query string
	query="SELECT aid FROM ustats WHERE aid=$1"
	err:=ss.db.QueryRow(query,aid).Scan(&addr_id);
	if (err!=nil) {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Ustats_entry_exists(aid=%v) sql error=%v\n",aid,err))
			os.Exit(1)
		}
		return false,err
	}
	return true,nil
}
