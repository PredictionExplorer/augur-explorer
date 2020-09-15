package main
import (
	"math/big"
	"encoding/hex"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/0xProject/0x-mesh/zeroex"
	ztypes "github.com/0xProject/0x-mesh/common/types"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func convert_db_event_to_mesh_order(evt *MeshEvent) *ztypes.OrderInfo {
	//Note we are using this convertion routine because we want to keep the other
	//		functions in dbs package dependant on 0x Mesh types, and not our custom types
	//		for a possible reuse in case of design change
	oinfo := new(ztypes.OrderInfo)
	oinfo.OrderHash = common.HexToHash(evt.OrderHash)
	oinfo.FillableTakerAssetAmount = new(big.Int)
	oinfo.FillableTakerAssetAmount.SetString(evt.FillableAmount,10)

	oinfo.SignedOrder.ChainID = new(big.Int)
	oinfo.SignedOrder.ChainID.SetInt64(int64(evt.ChainId))
	oinfo.SignedOrder.ExchangeAddress = common.HexToAddress(evt.ExchangeAddress)
	oinfo.SignedOrder.MakerAddress = common.HexToAddress(evt.MakerAddress)
	oinfo.SignedOrder.MakerAssetData,_ = hex.DecodeString(evt.MakerAssetData)
	oinfo.SignedOrder.MakerFeeAssetData,_ = hex.DecodeString(evt.MakerFeeAssetData)
	oinfo.SignedOrder.MakerAssetAmount = new(big.Int)
	oinfo.SignedOrder.MakerAssetAmount.SetString(evt.MakerAssetAmount,10)
	oinfo.SignedOrder.MakerFee = new(big.Int)
	oinfo.SignedOrder.MakerFee.SetString(evt.MakerFee,10)
	oinfo.SignedOrder.TakerAddress = common.HexToAddress(evt.TakerAddress)
	oinfo.SignedOrder.TakerAssetData,_ = hex.DecodeString(evt.TakerAssetData)
	oinfo.SignedOrder.TakerFeeAssetData,_ = hex.DecodeString(evt.TakerFeeAssetData)
	oinfo.SignedOrder.TakerAssetAmount = new(big.Int)
	oinfo.SignedOrder.TakerAssetAmount.SetString(evt.TakerAssetAmount,10)
	oinfo.SignedOrder.TakerFee = new(big.Int)
	oinfo.SignedOrder.TakerFee.SetString(evt.TakerFee,10)
	oinfo.SignedOrder.SenderAddress = common.HexToAddress(evt.SenderAddress)
	oinfo.SignedOrder.FeeRecipientAddress = common.HexToAddress(evt.FeeRecipientAddress)
	oinfo.SignedOrder.ExpirationTimeSeconds = new(big.Int)
	oinfo.SignedOrder.ExpirationTimeSeconds.SetInt64(evt.ExpirationTime)
	oinfo.SignedOrder.Salt = new(big.Int)
	oinfo.SignedOrder.Salt.SetString(evt.Salt,10)

	return oinfo
}
func oo_insert(order_hash *string,order *zeroex.SignedOrder,fillable_amount *big.Int,timestamp int64) error {

	ctx := context.Background()
	var copts = new(bind.CallOpts)
	adata,err := ctrct_zerox_trade.DecodeAssetData(copts,order.MakerAssetData)
	if err!=nil {
		Error.Printf("couldn't decode asset data for order %v : %v\n",*order_hash,err)
		return err
	}
	unpacked_id,err := ctrct_zerox_trade.UnpackTokenId(copts,adata.TokenIds[0])
	if err!=nil {
		Error.Printf("Unpack token id failed for order %v: %v\n",*order_hash,err)
		return err
	}
	num:=big.NewInt(int64(owner_fld_offset))
	key:=common.BigToHash(num)
	eoa,err := eclient.StorageAt(ctx,order.MakerAddress,key,nil)
	Info.Printf("oo_insert: order_hash=%v\n",*order_hash)
	Info.Printf(
		"oo insert: maker=%v eoa=%v; err=%v\n",
		order.MakerAddress.String(),hex.EncodeToString(eoa[:]),err,
	)
	var eoa_addr_str string
	if err == nil {
		eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
	} else {
		Info.Printf(
			"ethclient::StorageAt() failed for order %v, maker addr %v: %v. " +
			"Order will be inserted without EOA link. (ETH_STORAGE_FAIL)",
			order.MakerAddress.String(),*order_hash,err,
		)
		return err
	}
	err = storage.Insert_open_order(
		order_hash,order,fillable_amount,eoa_addr_str,&unpacked_id,OOOpCodeCreated,timestamp,
	)
	return err
}
func proc_open_orders() {

	status := storage.Get_mesh_proc_status()
	db_events := storage.Get_mesh_events_from_id(status.LastIdProcessed)
	for _,db_evt := range db_events {

		zorder := convert_db_event_to_mesh_order(&db_evt)
		evt_code := MeshEvtCode(db_evt.EvtCode)
		switch evt_code {
		case MeshEvtAdded:
			err:=oo_insert(
				&db_evt.OrderHash,
				zorder.SignedOrder,
				zorder.FillableTakerAssetAmount,
				db_evt.Timestamp,
			)
			if err!=nil {
				Info.Printf("Error inserting order %v: %v\n",db_evt.OrderHash,err)
			}
		case MeshEvtExpired:
			storage.Delete_open_0x_order(db_evt.OrderHash,OOOpCodeExpired)
		case MeshEvtCancelled:
			storage.Delete_open_0x_order(db_evt.OrderHash,OOOpCodeCancelledByUser)
		case MeshEvtFullyFilled:
			// FULLY FILLED event: quantity of the order matches filling quantity
			storage.Delete_open_0x_order(db_evt.OrderHash,OOOpCodeNone)
		case MeshEvtFilled:
			// FILLED event: partial order fill
			storage.Update_0x_order_on_partial_fill(zorder)
			// the following are rare events, so we don't implement them, just do a resync
		}
		storage.Set_mesh_proc_status(db_evt.Id)
	}
}
