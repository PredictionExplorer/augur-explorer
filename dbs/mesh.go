// 0x Mesh events
//	A layer separating Augur trading from 0x Mesh orders
//		Augur trading events are built in 'oohist' table
package dbs

import (
	"fmt"
	"os"
	"math/big"
	"encoding/hex"
	"database/sql"
	pq  "github.com/lib/pq"

//	"github.com/0xProject/0x-mesh/zeroex"
	ztypes "github.com/0xProject/0x-mesh/common/types"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Try_insert_0x_mesh_order_event(timestamp int64,oi *ztypes.OrderInfo,ospec *p.ZxMeshOrderSpec,amount_fill *big.Int,event_code p.MeshEvtCode) bool {

	_,err := ss.Lookup_market_by_addr_str(ospec.Market.String())
	if err != nil {
		ss.Info.Printf(
			"Try_insert_0x_mesh_order_event() fails: market %v isn't registered\n",
			ospec.Market.String(),
		)
		return false
	}
	ss.Insert_0x_mesh_order_event(timestamp,oi,ospec,amount_fill,event_code)
	return true
}
func (ss *SQLStorage) Insert_0x_mesh_order_event(timestamp int64,oi *ztypes.OrderInfo,ospec *p.ZxMeshOrderSpec,amount_fill *big.Int,event_code p.MeshEvtCode) {

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

	// Prevent duplicate insertion and exit if this is the case
	query = "SELECT id FROM mesh_evt WHERE order_hash=$1 AND evt_code=$2"
	var null_id sql.NullInt64
	err := ss.db.QueryRow(query,oi.OrderHash.String(),event_code).Scan(&null_id);
	if (err!=nil) {
		if (err==sql.ErrNoRows) {
			// break
		} else {
			ss.Log_msg(fmt.Sprintf("DB error : %v, q=%v",err,query))
			os.Exit(1)
		}
	} else {
		// event with this code already exists
		return
	}

	// Prevent insertion of event without ADD event. If ADD event is missing then simulate it
	if event_code != p.MeshEvtAdded {
		query = "SELECT id FROM mesh_evt WHERE order_hash=$1 AND evt_code=2"
		err = ss.db.QueryRow(query,oi.OrderHash.String()).Scan(&null_id);
		if (err!=nil) {
			if (err==sql.ErrNoRows) {
				// ADD event is missing, INSERT
				ts := int64(oi.SignedOrder.Order.Salt.Int64()/1000)
				ss.do_insert_0x_mesh_order_event(ts,oi,ospec,nil,p.MeshEvtAdded)
			} else {
				ss.Log_msg(fmt.Sprintf("DB error : %v, q=%v",err,query))
				os.Exit(1)
			}
		}
	}
	market_aid := ss.do_insert_0x_mesh_order_event(timestamp,oi,ospec,amount_fill,event_code)
	// now we need to update all posterior records because future price estimate values
	// depend on past values
	ss.Update_future_price_estimates(market_aid,int(ospec.Outcome),timestamp)
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
		ss.Info.Printf(
			"Future price estimate update: meshevt_id=%v, market_aid=%v,outc=%v,ts=%v\n",
			meshevt_id,market_aid,outcome_idx,ts,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("Update future price estimate failed: %v",err))
			os.Exit(1)
		}
	}
}
func (ss *SQLStorage) do_insert_0x_mesh_order_event(timestamp int64,oi *ztypes.OrderInfo,ospec *p.ZxMeshOrderSpec,amount_fill *big.Int,event_code p.MeshEvtCode) int64 {

	market_aid := ss.Lookup_or_create_address(ospec.Market.String(),0,0)// block and tx_id will be updated on market creation event
	amount_fill_str := "0"
	if amount_fill != nil {
		amount_fill_str = amount_fill.String()
	}
	var query string
	var err error
	query = "INSERT INTO mesh_evt (" +
				"time_stamp,fillable_amount,evt_code," +
				"market_aid,outcome_idx,otype,price," +
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
					"TO_TIMESTAMP($1),($2::decimal/1e+18),$3,"+
					"$4,$5,$6,$7," +
					"$8,$9,$10," +
					"$11,$12,$13," +
					"($14::decimal/1e+18),($15::decimal/1e+18),"+
					"$16,$17,$18,"+
					"($19::decimal/1e+18),($20::decimal/1e+18),"+
					"$21,$22,TO_TIMESTAMP($23::BIGINT),$24,$25,($26::decimal/1e+18)"+
			") ON CONFLICT DO NOTHING"

	d_query := fmt.Sprintf("INSERT INTO mesh_evt (" +
				"time_stamp,fillable_amount,evt_code," +
				"market_aid,outcome_idx,otype,price," +
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
					"TO_TIMESTAMP(%v),(%v::decimal/1e+18),%v,"+
					"%v,%v,%v,%v," +
					"'%v',%v,'%v'," +
					"'%v','%v','%v'," +
					"(%v::decimal/1e+18),(%v::decimal/1e+18),"+
					"'%v','%v','%v',"+
					"(%v::decimal/1e+18),(%v::decimal/1e+18),"+
					"'%v','%v',TO_TIMESTAMP(%v::BIGINT),%v,'%v',(%v::decimal/1e+18)"+
			") ON CONFLICT DO NOTHING",
			timestamp,oi.FillableTakerAssetAmount.String(),event_code,
			market_aid,ospec.Outcome,ospec.Type,ospec.Price.String(),
			oi.OrderHash.String(),oi.SignedOrder.Order.ChainID.Int64(),oi.SignedOrder.Order.ExchangeAddress.String(),
			oi.SignedOrder.Order.MakerAddress.String(),hex.EncodeToString(oi.SignedOrder.Order.MakerAssetData),hex.EncodeToString(oi.SignedOrder.Order.MakerFeeAssetData),
			oi.SignedOrder.Order.MakerAssetAmount.String(),oi.SignedOrder.Order.MakerFee.String(),
			oi.SignedOrder.Order.TakerAddress.String(),hex.EncodeToString(oi.SignedOrder.Order.TakerAssetData),hex.EncodeToString(oi.SignedOrder.Order.TakerFeeAssetData),
			oi.SignedOrder.Order.TakerAssetAmount.String(),oi.SignedOrder.Order.TakerFee.String(),
			oi.SignedOrder.Order.SenderAddress.String(),oi.SignedOrder.Order.FeeRecipientAddress.String(),oi.SignedOrder.Order.ExpirationTimeSeconds.Int64(),oi.SignedOrder.Order.Salt.String(),hex.EncodeToString(oi.SignedOrder.Signature),amount_fill_str,
	)
	ss.Info.Printf("q=%v\n",d_query)
	_,err = ss.db.Exec(query,
		timestamp,oi.FillableTakerAssetAmount.String(),event_code,
		market_aid,ospec.Outcome,ospec.Type,ospec.Price.String(),
		oi.OrderHash.String(),oi.SignedOrder.Order.ChainID.Int64(),oi.SignedOrder.Order.ExchangeAddress.String(),
		oi.SignedOrder.Order.MakerAddress.String(),hex.EncodeToString(oi.SignedOrder.Order.MakerAssetData),hex.EncodeToString(oi.SignedOrder.Order.MakerFeeAssetData),
		oi.SignedOrder.Order.MakerAssetAmount.String(),oi.SignedOrder.Order.MakerFee.String(),
		oi.SignedOrder.Order.TakerAddress.String(),hex.EncodeToString(oi.SignedOrder.Order.TakerAssetData),hex.EncodeToString(oi.SignedOrder.Order.TakerFeeAssetData),
		oi.SignedOrder.Order.TakerAssetAmount.String(),oi.SignedOrder.Order.TakerFee.String(),
		oi.SignedOrder.Order.SenderAddress.String(),oi.SignedOrder.Order.FeeRecipientAddress.String(),oi.SignedOrder.Order.ExpirationTimeSeconds.Int64(),oi.SignedOrder.Order.Salt.String(),hex.EncodeToString(oi.SignedOrder.Signature),amount_fill_str,
	)
	if (err!=nil) {
		pq_err := err.(*pq.Error)
		ss.Log_msg(
			fmt.Sprintf(
				"do_insert_0x_mesh_order_event() failed: %v %v %v %v; q=%v",
				pq_err,pq_err.Routine,pq_err.Position,pq_err.Where,
				query,
			),
		)
		os.Exit(1)
	}
	return market_aid
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
