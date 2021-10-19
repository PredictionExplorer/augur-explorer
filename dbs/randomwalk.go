package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_randomwalk_processing_status() p.RandomWalkProcStatus {

	var output p.RandomWalkProcStatus
	var null_id,null_block sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id,last_block FROM rw_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id,&null_block)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO rw_proc_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		output.LastIdProcessed = null_id.Int64
	}
	if null_block.Valid {
		output.LastBlockNum = null_block.Int64
	}
	return output
}
func (ss *SQLStorage) Update_randomwalk_process_status(status *p.RandomWalkProcStatus) {

	var query string
	query = "UPDATE rw_proc_status SET last_evt_id = $1,last_block=$2"

	_,err := ss.db.Exec(query,status.LastIdProcessed,status.LastBlockNum)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_randomwalk_contract_addresses() p.RW_ContractAddresses {

	var output p.RW_ContractAddresses
	var query string
	query = "SELECT "+
				"marketplace_addr,randomwalk_addr," +
				"mp_a.address_id,rw_a.address_id "+
			"FROM rw_contracts rw " +
				"JOIN address mp_a ON rw.marketplace_addr=mp_a.addr " +
				"JOIN address rw_a ON rw.randomwalk_addr=rw_a.addr "

	res := ss.db.QueryRow(query)
	err := res.Scan(
		&output.MarketPlace,
		&output.RandomWalk,
		&output.MarketPlaceAid,
		&output.RandomWalkAid,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output
		}
		ss.Log_msg(fmt.Sprintf("Get_randomwalk_contract_addresses() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	return output
}
func (ss *SQLStorage) Insert_new_offer(evt *p.RW_NewOffer) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	buyer_aid:=ss.Lookup_or_create_address(evt.Buyer,evt.BlockNum,evt.TxId)
	seller_aid:=ss.Lookup_or_create_address(evt.Seller,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_new_offer(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"offer_id,token_id,buyer_aid,seller_aid,active,price" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.OfferId,
		evt.TokenId,
		buyer_aid,
		seller_aid,
		true,
		evt.Price,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into new_offer table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_item_bought(evt *p.RW_ItemBought) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_item_bought(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"offer_id" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.OfferId,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into item_bought table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_offer_canceled(evt *p.RW_OfferCanceled) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_offer_canceled(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"offer_id" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.OfferId,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into offer_canceled table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_withdrawal(evt *p.RW_Withdrawal) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	aid:=ss.Lookup_or_create_address(evt.Destination,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_withdrawal(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"aid,token_id,amount" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		aid,
		evt.TokenId,
		evt.Amount,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into rw_withdrawal table: %v\n",err))
		os.Exit(1)
	}

}
func (ss *SQLStorage) Insert_token_name(evt *p.RW_TokenName) {

	contract_aid:=ss.Lookup_or_create_address(evt.Contract,evt.BlockNum,evt.TxId)
	var query string
	query = "INSERT INTO rw_token_name(" +
				"evtlog_id,block_num,tx_id,time_stamp,contract_aid, "+
				"token_id,new_name" +
			") VALUES (" +
				"$1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7"+
			")"
	_,err := ss.db.Exec(query,
		evt.EvtId,
		evt.BlockNum,
		evt.TxId,
		evt.TimeStamp,
		contract_aid,
		evt.TokenId,
		evt.NewName,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: can't insert into rw_token_name table: %v\n",err))
		os.Exit(1)
	}

}
