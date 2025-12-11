package dbs

import (
	"fmt"
	"os"
)

func (ss *SQLStorage) Get_last_block_timestamp() int64 {
	var query string
	query = "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
		"FROM block,last_block WHERE last_block.block_num=block.block_num"
	row := ss.db.QueryRow(query)
	var ts int64
	var err error
	err = row.Scan(&ts)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Error in Get_last_block_timestamp(): %v, q=%v", err, query))
		os.Exit(1)
	}
	return ts
}

func (ss *SQLStorage) Get_first_block_timestamp() int64 {
	var query string
	query = "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
		"FROM block ORDER BY block_num LIMIT 1"
	row := ss.db.QueryRow(query)
	var ts int64
	var err error
	err = row.Scan(&ts)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Error in Get_first_block_timestamp(): %v, q=%v", err, query))
		os.Exit(1)
	}
	return ts
}
