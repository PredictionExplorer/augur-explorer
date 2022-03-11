package dbs

import (
	"fmt"
	"os"
	"strings"
	"database/sql"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/balancer_v2"
)

func (ss *SQLStorage) Insert_pool_created(evt *p.BalV2PoolCreated) {

	var query string
	query =  "INSERT INTO pool_created(block_num,time_stamp,tx_index,log_index,pool_aid) " +
				"VALUES($1,TO_TIMESTAMP($2),$3,$4,$5)"
	
	_,err := ss.db.Exec(query,
		evt.BlockNum,
		evt.TimeStamp,
		evt.TxIndex,
		evt.LogIndex,
		pool_aid,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into pool_created table: %v\n",err))
		os.Exit(1)
	}

}
