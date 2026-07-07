// Package dbs - address lookups and cache used by rwcg ETL and APIs
package dbs

import (
	"database/sql"
	"fmt"
	"sync"
)

var (
	amapMu sync.RWMutex
	amap   map[string]int64 = make(map[string]int64)
)

// ResetAddressCacheForTests clears the process-wide address-id cache. Test
// harnesses that truncate and re-seed the address table between cases must
// call it, otherwise cached ids from a previous seeding would leak into the
// next one. (Phase 1 replaces this package state with a per-Store cache.)
func ResetAddressCacheForTests() {
	amapMu.Lock()
	amap = make(map[string]int64)
	amapMu.Unlock()
}

// Nonfatal_lookup_address_id returns the address_id for addr.
// A missing address yields sql.ErrNoRows; any other failure is returned as a
// wrapped DB error.
func (ss *SQLStorage) Nonfatal_lookup_address_id(addr string) (int64, error) {
	amapMu.RLock()
	aid, exists := amap[addr]
	amapMu.RUnlock()
	if exists {
		return aid, nil
	}
	query := "SELECT address_id FROM address WHERE addr=$1"
	err := ss.db.QueryRow(query, addr).Scan(&aid)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, fmt.Errorf("address id lookup for %v: %w", addr, err)
		}
		return 0, err
	}
	amapMu.Lock()
	if cached, ok := amap[addr]; ok {
		amapMu.Unlock()
		return cached, nil
	}
	amap[addr] = aid
	amapMu.Unlock()
	return aid, nil
}

func (ss *SQLStorage) Lookup_address(aid int64) (string, error) {
	var addr string
	query := "SELECT addr FROM address WHERE address_id=$1"
	err := ss.db.QueryRow(query, aid).Scan(&addr)
	return addr, err
}

// Lookup_address_id returns the address_id for addr. Unlike
// Nonfatal_lookup_address_id, a missing address is reported as an error too
// (callers use it for addresses that must already exist).
func (ss *SQLStorage) Lookup_address_id(addr string) (int64, error) {
	var addr_id int64 = 0
	query := "SELECT address_id FROM address WHERE addr=$1"
	err := ss.db.QueryRow(query, addr).Scan(&addr_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("forced address lookup for %v: addr not found: %w", addr, err)
		}
		return 0, fmt.Errorf("address lookup for %v: %w", addr, err)
	}
	return addr_id, nil
}

// Lookup_or_create_address returns the address_id for addr, inserting a new
// address row (recorded against block_num/tx_id) when it doesn't exist yet.
func (ss *SQLStorage) Lookup_or_create_address(addr string, block_num int64, tx_id int64) (int64, error) {
	amapMu.RLock()
	aid, exists := amap[addr]
	amapMu.RUnlock()
	if exists {
		return aid, nil
	}
	query := "SELECT address_id FROM address WHERE addr=$1"
	err := ss.db.QueryRow(query, addr).Scan(&aid)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, fmt.Errorf("address id lookup for %v: %w", addr, err)
		}
		aid, err = ss.create_address(addr, block_num, tx_id)
		if err != nil {
			return 0, err
		}
	}
	amapMu.Lock()
	if cached, ok := amap[addr]; ok {
		amapMu.Unlock()
		return cached, nil
	}
	amap[addr] = aid
	amapMu.Unlock()
	return aid, nil
}

func (ss *SQLStorage) create_address(addr string, block_num int64, tx_id int64) (int64, error) {
	var addr_id int64
	if len(addr) == 0 {
		return 0, fmt.Errorf("attempt to insert address with len=0, block %v, tx_id=%v", block_num, tx_id)
	}
	query := "INSERT INTO address(addr,block_num,tx_id) VALUES($1,$2,$3) RETURNING address_id"
	row := ss.db.QueryRow(query, addr, block_num, tx_id)
	err := row.Scan(&addr_id)
	if err != nil {
		return 0, fmt.Errorf("address insertion for %v: %w", addr, err)
	}
	if addr_id == 0 {
		return 0, fmt.Errorf("address insertion for %v: addr_id after INSERT is 0", addr)
	}
	return addr_id, nil
}
