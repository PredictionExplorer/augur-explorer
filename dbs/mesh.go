// 0x Mesh events
//	A layer separating Augur trading from 0x Mesh orders
//		Augur trading events are built in 'oohist' table
package dbs

import (
	"fmt"
	"os"
	"math/big"
	"errors"
	"encoding/hex"
	"database/sql"
	pq  "github.com/lib/pq"

	//ztypes "github.com/0xProject/0x-mesh/common/types"
	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Try_insert_0x_mesh_order_event(aid int64,timestamp int64,oi *p.OrderInfo0x,ospec *p.ZxMeshOrderSpec,amount_fill *big.Int,event_code p.MeshEvtCode) bool {

	_,err := ss.Lookup_market_by_addr_str(ospec.Market.String())
	if err != nil {
		ss.Info.Printf(
			"Try_insert_0x_mesh_order_event() fails: market %v isn't registered\n",
			ospec.Market.String(),
		)
		return false
	}
	ss.Insert_0x_mesh_order_event(0,aid,timestamp,oi,ospec,amount_fill,event_code)
	return true
}
func (ss *SQLStorage) Is_ADD_event_missing(order_hash string) bool {

	var query string
	query = "SELECT id FROM mesh_evt WHERE order_hash=$1 AND evt_code=2"
	var null_id sql.NullInt64
	err := ss.db.QueryRow(query,order_hash).Scan(&null_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			return true
		} else {
			ss.Log_msg(fmt.Sprintf("DB error : %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	return false
}
func (ss *SQLStorage) Insert_0x_mesh_order_event(mktord_id int64,aid int64,timestamp int64,oi *p.OrderInfo0x,ospec *p.ZxMeshOrderSpec,amount_fill *big.Int,event_code p.MeshEvtCode) bool {

	ss.Info.Printf("Inserting 0x mesh event, order %v\n",oi.OrderHash.String())
	if ospec == nil {
		ss.Log_msg(
			fmt.Sprintf(
				"Null ospec parameter in Insert_0x_mesh_order_event(h=%v)",
				oi.OrderHash.String(),
			),
		)
		os.Exit(1)
	}
	var query string
	// Prevent insertion of event without ADD event. If ADD event is missing then simulate it
	if event_code != p.MeshEvtAdded {
		query = "SELECT id FROM mesh_evt WHERE order_hash=$1 AND evt_code=2"
		var null_id sql.NullInt64
		err := ss.db.QueryRow(query,oi.OrderHash.String()).Scan(&null_id);
		if (err!=nil) {
			if (err==sql.ErrNoRows) {
				// ADD event is missing, INSERT
				ts := int64(oi.SignedOrder.Salt.Int64()/1000)
				ss.do_insert_0x_mesh_order_event(0,aid,ts,oi,ospec,nil,p.MeshEvtAdded)
			} else {
				ss.Log_msg(fmt.Sprintf("DB error : %v, q=%v",err,query))
				os.Exit(1)
			}
		}
	}
	market_aid,err := ss.do_insert_0x_mesh_order_event(mktord_id,aid,timestamp,oi,ospec,amount_fill,event_code)
	if err != nil {
		ss.Info.Printf("not updating price estimate due to error: %v\n",err)
		return false
	}
	// now we need to update all posterior records because future price estimate values
	// depend on past values
	ss.Update_future_price_estimates(market_aid,int(ospec.Outcome),timestamp)
	if ( event_code == p.MeshEvtFilled) || (event_code == p.MeshEvtFullyFilled) {
		ss.Update_oo_fillable_amount(oi.OrderHash.String(),oi.SignedOrder)
	}
	return true
}
func (ss *SQLStorage) Update_future_price_estimates(market_aid int64,outcome_idx int,timestamp int64) {

	var query string
	query = "SELECT "+
			"id,market_aid,outcome_idx,FLOOR(EXTRACT(EPOCH FROM time_stamp))::BIGINT as ts " +
			"FROM mesh_evt " +
			"WHERE market_aid=$1 AND outcome_idx=$2 AND time_stamp > TO_TIMESTAMP($3) " +
			"ORDER BY time_stamp"

	rows,err := ss.db.Query(query,market_aid,outcome_idx,timestamp)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var meshevt_id,market_aid,ts int64
		var outcome_idx int
		err=rows.Scan(&meshevt_id,&market_aid,&outcome_idx,&ts)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Scan failed: %v, q=%v",err,query))
			os.Exit(1)
		}
		_,err = ss.db.Exec(
			"SELECT update_price_estimate($1,$2,$3,TO_TIMESTAMP($4))",
			meshevt_id,market_aid,outcome_idx,ts,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Update future price estimate failed: %v",err))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) do_insert_0x_mesh_order_event(mktord_id int64,aid int64,timestamp int64,oi *p.OrderInfo0x,ospec *p.ZxMeshOrderSpec,amount_fill *big.Int,event_code p.MeshEvtCode) (int64,error) {

	market_aid := ss.Lookup_or_create_address(ospec.Market.String(),0,0)// block and tx_id will be updated on market creation event
	amount_fill_str := "0"
	if amount_fill != nil {
		amount_fill_str = amount_fill.String()
	}
	var mktord_field,mktord_value string
	if mktord_id != 0 { // non-nul value is inserted to related mesh_evt to mktord table
		mktord_field = "mktord_id,"
		mktord_value = fmt.Sprintf("%v,",mktord_id)
	}
	var query string
	var err error
	query = "INSERT INTO mesh_evt (" +
				"aid," +
				"time_stamp,fillable_amount,evt_code," +
				mktord_field + "market_aid,outcome_idx,otype,price," +
				"order_hash,chain_id,exchange_addr," +
				"maker_addr," +
				"maker_asset_data," +
				"maker_fee_asset_data," +
				"maker_asset_amount," +
				"maker_fee," +
				"taker_address," +
				"taker_asset_data," +
				"taker_fee_asset_data," +
				"taker_asset_amount," +
				"taker_fee," +
				"sender_address," +
				"fee_recipient_address," +
				"expiration_time," +
				"salt," +
				"signature," +
				"amount_fill" +
			") VALUES (" +
					"$1," +
					"TO_TIMESTAMP($2),($3::decimal/1e+18),$4,"+
					mktord_value + "$5,$6,$7,$8," +
					"$9,$10,$11," +
					"$12,$13,$14," +
					"($15::decimal/1e+18),($16::decimal/1e+18),"+
					"$17,$18,$19,"+
					"($20::decimal/1e+18),($21::decimal/1e+18),"+
					"$22,$23,TO_TIMESTAMP($24::BIGINT),$25,$26,($27::decimal/1e+18)"+
			") ON CONFLICT DO NOTHING"

	d_query := fmt.Sprintf("INSERT INTO mesh_evt (" +
				"aid," +
				"time_stamp,fillable_amount,evt_code," +
				mktord_field + "market_aid,outcome_idx,otype,price," +
				"order_hash,chain_id,exchange_addr," +
				"maker_addr,maker_asset_data,maker_fee_asset_data," +
				"maker_asset_amount,maker_fee," +
				"taker_address,taker_asset_data,taker_fee_asset_data," +
				"taker_asset_amount,taker_fee," +
				"sender_address," +
				"fee_recipient_address," +
				"expiration_time," +
				"salt," +
				"signature," +
				"amount_fill"+
			") VALUES (" +
					"%v," +
					"TO_TIMESTAMP(%v),(%v::decimal/1e+18),%v,"+
					"%v %v,%v,%v,%v," +
					"'%v',%v,'%v'," +
					"'%v','%v','%v'," +
					"(%v::decimal/1e+18),(%v::decimal/1e+18),"+
					"'%v','%v','%v',"+
					"(%v::decimal/1e+18),(%v::decimal/1e+18),"+
					"'%v','%v',TO_TIMESTAMP(%v::BIGINT),%v,'%v',(%v::decimal/1e+18)"+
			") ON CONFLICT DO NOTHING",
			aid,
			timestamp,oi.FillableTakerAssetAmount.String(),event_code,
			mktord_value,market_aid,ospec.Outcome,ospec.Type,ospec.Price.String(),
			oi.OrderHash.String(),oi.SignedOrder.ChainID.Int64(),oi.SignedOrder.ExchangeAddress.String(),
			oi.SignedOrder.MakerAddress.String(),hex.EncodeToString(oi.SignedOrder.MakerAssetData),hex.EncodeToString(oi.SignedOrder.MakerFeeAssetData),
			oi.SignedOrder.MakerAssetAmount.String(),oi.SignedOrder.MakerFee.String(),
			oi.SignedOrder.TakerAddress.String(),hex.EncodeToString(oi.SignedOrder.TakerAssetData),hex.EncodeToString(oi.SignedOrder.TakerFeeAssetData),
			oi.SignedOrder.TakerAssetAmount.String(),oi.SignedOrder.TakerFee.String(),
			oi.SignedOrder.SenderAddress.String(),oi.SignedOrder.FeeRecipientAddress.String(),oi.SignedOrder.ExpirationTimeSeconds.Int64(),oi.SignedOrder.Salt.String(),hex.EncodeToString(oi.SignedOrder.Signature),amount_fill_str,
	)
	ss.Info.Printf("q=%v\n",d_query)
	_,err = ss.db.Exec(query,
		aid,
		timestamp,oi.FillableTakerAssetAmount.String(),event_code,
		market_aid,ospec.Outcome,ospec.Type,ospec.Price.String(),
		oi.OrderHash.String(),oi.SignedOrder.ChainID.Int64(),oi.SignedOrder.ExchangeAddress.String(),
		oi.SignedOrder.MakerAddress.String(),hex.EncodeToString(oi.SignedOrder.MakerAssetData),hex.EncodeToString(oi.SignedOrder.MakerFeeAssetData),
		oi.SignedOrder.MakerAssetAmount.String(),oi.SignedOrder.MakerFee.String(),
		oi.SignedOrder.TakerAddress.String(),hex.EncodeToString(oi.SignedOrder.TakerAssetData),hex.EncodeToString(oi.SignedOrder.TakerFeeAssetData),
		oi.SignedOrder.TakerAssetAmount.String(),oi.SignedOrder.TakerFee.String(),
		oi.SignedOrder.SenderAddress.String(),oi.SignedOrder.FeeRecipientAddress.String(),oi.SignedOrder.ExpirationTimeSeconds.Int64(),oi.SignedOrder.Salt.String(),hex.EncodeToString(oi.SignedOrder.Signature),amount_fill_str,
	)
	ss.Info.Printf("err=%v\n",err)
	if (err!=nil) {
		switch err.(type) {
			default:
				ss.Log_msg(fmt.Sprintf("do_insert_0x_mesh_order_event() failed: %v ; q=%v",err,query))
			case *pq.Error:
				pq_err,_ := err.(*pq.Error)
				ss.Log_msg(
					fmt.Sprintf(
						"do_insert_0x_mesh_order_event() failed: %v %v %v %v; q=%v",
						pq_err,pq_err.Routine,pq_err.Position,pq_err.Where,
						query,
					),
				)
		}
		return 0,err
	}
	return market_aid,nil
}
func (ss *SQLStorage) Get_mesh_proc_status() p.MeshProcStatus {

	var status p.MeshProcStatus
	var query string
	var null_id sql.NullInt64

	for {
		query = "SELECT last_id_processed FROM mesh_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO mesh_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		status.LastIdProcessed = null_id.Int64
	}
	return status
}
func (ss *SQLStorage) Set_mesh_proc_status(last_id int64) {

	var query string = "UPDATE mesh_status SET last_id_processed=$1"
	res,err:=ss.db.Exec(query,last_id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("Set_mesh_proc_status() failed: %v",err))
		os.Exit(1)
	}
	affected_rows,err:=res.RowsAffected()
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("Error getting RowsAffected in Set_mesh_proc_status(): %v",err))
		os.Exit(1)
	}
	if affected_rows>0 {
		// break
	} else {
		query = "INSERT INTO mesh_status VALUES($1)"
		_,err := ss.db.Exec(query,last_id)
		if (err!=nil) {
			ss.Log_msg(fmt.Sprintf("Set_mesh_proc_status() failed on INSERT: %v",err));
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) Get_mesh_events_from_id(id int64) []p.MeshEvent {

	records := make([]p.MeshEvent,0,16)

	var query string
	query = "SELECT " +
				"id," +
				"FLOOR(EXTRACT(EPOCH FROM time_stamp))::BIGINT as time_stamp," +
				"ROUND(fillable_amount*1e+18) AS fillable_amount,evt_code,order_hash," +
				"chain_id,exchange_addr,maker_addr,maker_asset_data,maker_fee_asset_data," +
				"ROUND(maker_asset_amount*1e+18) AS maker_asset_amount,"+
				"ROUND(maker_fee*1e+18) AS maker_fee,taker_address," +
				"taker_asset_data,taker_fee_asset_data," +
				"ROUND(taker_asset_amount*1e+18) AS taker_asset_amount," +
				"ROUND(taker_fee*1e+18) AS taker_fee,sender_address,fee_recipient_address," +
				"FLOOR(EXTRACT(EPOCH FROM expiration_time))::BIGINT as expiration_ts," +
				"salt,signature " +
			"FROM mesh_evt " +
			"WHERE id > $1 " +
			"ORDER BY ID"
//	ss.Info.Println("(id=%v) q=%v\n",id,query)
	rows,err := ss.db.Query(query,id)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var rec p.MeshEvent

	defer rows.Close()
	for rows.Next() {
		err=rows.Scan(
			&rec.Id,
			&rec.Timestamp,
			&rec.FillableAmount,
			&rec.EvtCode,
			&rec.OrderHash,
			&rec.ChainId,
			&rec.ExchangeAddress,
			&rec.MakerAddress,
			&rec.MakerAssetData,
			&rec.MakerFeeAssetData,
			&rec.MakerAssetAmount,
			&rec.MakerFee,
			&rec.TakerAddress,
			&rec.TakerAssetData,
			&rec.TakerFeeAssetData,
			&rec.TakerAssetAmount,
			&rec.TakerFee,
			&rec.SenderAddress,
			&rec.FeeRecipientAddress,
			&rec.ExpirationTime,
			&rec.Salt,
			&rec.Signature,
		)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (ss *SQLStorage) Get_mesh_event_by_order_hash(order_hash string) (p.MeshEvent,error) {

	var order p.MeshEvent
	order.OrderHash = order_hash
	var query string
	query = "SELECT " +
				"FLOOR(EXTRACT(EPOCH FROM time_stamp) AS time_stamp," +
				"fillable_amount," +
				"evt_code," +
				"chain_id," +
				"exchange_addr," +
				"maker_addr," +
				"maker_asset_data," +
				"maker_fee_asset_data," +
				"maker_asset_amount," +
				"maker_fee," +
				"taker_address," +
				"taker_asset_data," +
				"taker_asset_amount," +
				"taker_fee," +
				"taker_fee_asset_data," +
				"sender_address," +
				"fee_recipient_address," +
				"FLOOR(EXTRACT(EPOCH FROM expiration_time))::BIGINT AS expiration_time," +
				"salt," +
				"signature" +
			"FROM mesh_evt " +
			"WHERE order_hash = $1 " +
			"LIMIT 1"
	row := ss.db.QueryRow(query,order_hash)
	err := row.Scan(
		&order.Timestamp,
		&order.FillableAmount,
		&order.EvtCode,
		&order.ChainId,
		&order.ExchangeAddress,
		&order.MakerAddress,
		&order.MakerAssetData,
		&order.MakerFeeAssetData,
		&order.MakerAssetAmount,
		&order.MakerFee,
		&order.TakerAddress,
		&order.TakerAssetData,
		&order.TakerFeeAssetData,
		&order.TakerAssetAmount,
		&order.TakerFee,
		&order.SenderAddress,
		&order.FeeRecipientAddress,
		&order.ExpirationTime,
		&order.Salt,
		&order.Signature,
	)
	if err!=nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
		os.Exit(1)
	}
	return order,nil
}
func (ss *SQLStorage) Lookup_maker_eoa_wallet_ids(possible_eoa_addr *common.Address,maker_addr *common.Address) (int64,int64,error) {

	// Note: the MakerAddress can be either EOA of the User or Wallet contract of the User
	//			we need to figure out which one we have been given
	//	This function returns error if the addres is unregistered (by 'unregistered' we mean that it wasnt
	//	activated for Augur trading (no Approvals or no Wallet Contract created)
	var wallet_aid int64 = 0
	var eoa_aid int64 = 0
	var err error
	zero_addr := common.Address{}
	if zero_addr.String()==possible_eoa_addr.String() {	// MakerAddress must be an EOA address
		eoa_aid,err = ss.Nonfatal_lookup_address_id(maker_addr.String())
		if err != nil {
			// Note: we can't INSERT an address from here because we need Transaction Hash (for reference)
			ss.Info.Printf(
				"MakerAddress %v is unregistered in the DB as EOA ",
				maker_addr.String(),
			)
			return eoa_aid,wallet_aid,errors.New(
				fmt.Sprintf("MakerAddress %v is unregistered in the DB as EOA.",maker_addr.String()),
			)
		}
		// Now we need to validate this this address is indeed an EOA
		wallet_aid,err = ss.Lookup_wallet_aid(eoa_aid)
		if err != nil {
			ss.Info.Printf(
				"MakerAddress %v doesn't have an associated Wallet contract",
				maker_addr.String(),
			)
			return eoa_aid,wallet_aid,errors.New(
				"EOA address provided from Mesh listener is zero address and un-registered.",
			)
		}
	} else { // MakerAddress has an EOA address, means MakerAddress is a Wallet contract
		wallet_aid,err = ss.Nonfatal_lookup_address_id(maker_addr.String())
		if err != nil {
			// Maker is not Wallet Contract, then it must be EOA
			ss.Info.Printf(
				"MakerAddress %v is unregistered in the DB as wallet : %v",
				maker_addr.String(),err,
			)
			return eoa_aid,wallet_aid,errors.New(
				fmt.Sprintf("MakerAddress %v is unregistered in the DB.",maker_addr.String()),
			)
		}
		eoa_aid,err = ss.Lookup_eoa_aid(wallet_aid)
		if err != nil {
			ss.Info.Printf(
				"MakerAddress %v is a Wallet contract that doesn't have an associated EOA address: %v",
				maker_addr.String(),err,
			)
			return eoa_aid,wallet_aid,errors.New(
				fmt.Sprintf("MakerAddress %v doesn't have associated EOA",maker_addr.String()),
			)
		}
	}
	return eoa_aid,wallet_aid,nil
}
func (ss *SQLStorage) Get_price_estimate_update_market_list() []p.UpdatePriceEst {

	var query string
	query = "SELECT " +
				"market_aid," +
				"FLOOR(EXTRACT(EPOCH FROM create_timestamp))::BIGINT as created_ts, " +
				"market_type," +
				"outcomes " +
			"FROM market "
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.UpdatePriceEst,0,256)

	defer rows.Close()
	for rows.Next() {
		var rec p.UpdatePriceEst
		var market_aid int64
		var timestamp int64
		var market_type int
		var outcomes string
		err=rows.Scan(&market_aid,&timestamp,&market_type,&outcomes)
		if err != nil {
			ss.Info.Printf("DB error: %v\n",err)
			os.Exit(1)
		}
		rec.MktAid = market_aid
		rec.TimeStamp = timestamp
		rec.MktType = market_type
		rec.Outcomes = outcomes
		records = append(records,rec)
	}
	return records
}
