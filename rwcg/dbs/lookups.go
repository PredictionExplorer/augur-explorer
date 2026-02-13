// Package dbs - address lookups and cache used by rwcg ETL and APIs
package dbs

import (
	"database/sql"
	"fmt"
	"os"
	"runtime/debug"
)

var (
	amap map[string]int64 = make(map[string]int64)
)

func (ss *SQLStorage) Nonfatal_lookup_address_id(addr string) (int64, error) {
	aid, exists := amap[addr]
	if exists {
		return aid, nil
	}
	query := "SELECT address_id FROM address WHERE addr=$1"
	err := ss.db.QueryRow(query, addr).Scan(&aid)
	if err != nil {
		if err != sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("DB error: %v ,q=%v", query, err))
			os.Exit(1)
		}
		return 0, err
	}
	amap[addr] = aid
	return aid, nil
}

func (ss *SQLStorage) Lookup_address(aid int64) (string, error) {
	var addr string
	query := "SELECT addr FROM address WHERE address_id=$1"
	err := ss.db.QueryRow(query, aid).Scan(&addr)
	return addr, err
}

func (ss *SQLStorage) Lookup_address_id(addr string) int64 {
	var addr_id int64 = 0
	query := "SELECT address_id FROM address WHERE addr=$1"
	err := ss.db.QueryRow(query, addr).Scan(&addr_id)
	if err != nil {
		if err == sql.ErrNoRows {
			ss.Log_msg(fmt.Sprintf("Forced address lookup failed for %v : addr not found", addr))
			ss.Log_msg(fmt.Sprintf("Printing stack trace to help locating the actual function with error"))
			debug.PrintStack()
			os.Exit(1)
		} else {
			ss.Log_msg(fmt.Sprintf("DB error upon address lookup: %v ; q=%v", err, query))
			os.Exit(1)
		}
	}
	return addr_id
}

func (ss *SQLStorage) Lookup_or_create_address(addr string, block_num int64, tx_id int64) int64 {
	aid, exists := amap[addr]
	if exists {
		return aid
	}
	query := "SELECT address_id FROM address WHERE addr=$1"
	err := ss.db.QueryRow(query, addr).Scan(&aid)
	if err != nil {
		if err == sql.ErrNoRows {
			aid = ss.create_address(addr, block_num, tx_id)
			amap[addr] = aid
			return aid
		} else {
			ss.Log_msg(fmt.Sprintf("DB error in getting address id : %v", err))
		}
	}
	amap[addr] = aid
	return aid
}

func (ss *SQLStorage) create_address(addr string, block_num int64, tx_id int64) int64 {
	var addr_id int64
	if len(addr) == 0 {
		ss.Log_msg(fmt.Sprintf("Attempt to insert address with len=0, block %v, tx_id=%v", block_num, tx_id))
		os.Exit(1)
	}
	query := "INSERT INTO address(addr,block_num,tx_id) VALUES($1,$2,$3) RETURNING address_id"
	row := ss.db.QueryRow(query, addr, block_num, tx_id)
	err := row.Scan(&addr_id)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error in address insertion: %v : %v", query, err))
		os.Exit(1)
	}
	if addr_id == 0 {
		ss.Log_msg(fmt.Sprintf("DB error, addr_id after INSERT is 0"))
		os.Exit(1)
	}
	return addr_id
}
