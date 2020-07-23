// Data Base Storage
package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) lookup_universe_id(addr string) (int64,error) {

	var query string
	query="SELECT universe_id FROM universe WHERE universe_addr=$1"
	var universe_id int64 = 0
	err:=ss.db.QueryRow(query,addr).Scan(&universe_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			return 0,err
		} else {
			ss.Log_msg(fmt.Sprintf("DB error looking up for Universe record: %v",err))
			os.Exit(1)
		}
	}
	return universe_id,nil
}
func (ss *SQLStorage) Lookup_eoa_aid(wallet_aid int64) (int64,error) {

	var addr_id int64;
	var query string
	query="SELECT eoa_aid FROM ustats WHERE wallet_aid=$1"
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
	query="SELECT wallet_aid FROM ustats WHERE eoa_aid=$1"
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
func (ss *SQLStorage) lookup_market_by_addr_str(addr string) (int64,error) {

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
func (ss *SQLStorage) Lookup_address(eoa_aid int64) (string,error) {

	var addr string;
	var query string
	query="SELECT addr FROM address WHERE address_id=$1"
	err:=ss.db.QueryRow(query,eoa_aid).Scan(&addr);
	return addr,err
}
func (ss *SQLStorage) lookup_address_id(addr string) int64 {

	var addr_id int64;
	var query string
	query="SELECT address_id FROM address WHERE addr=$1"
	err:=ss.db.QueryRow(query,addr).Scan(&addr_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
		} else {
			ss.Log_msg(fmt.Sprintf("DB error upon address lookup: %v",err))
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
func (ss *SQLStorage) Lookup_or_create_address(addr string,block_num p.BlockNumber,tx_id int64) int64 {

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
func (ss *SQLStorage) create_address(addr string,block_num p.BlockNumber,tx_id int64) int64 {

	var addr_id int64;
	var query string

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
