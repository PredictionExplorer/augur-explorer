package main

import (
	"os"
	"log"
	"bytes"
	"encoding/hex"
	"math/big"
	"context"
	//"fmt"
	//"database/sql"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	ztypes "github.com/0xProject/0x-mesh/common/types"
	"github.com/0xProject/0x-mesh/zeroex"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (
	MARKET_ORDER = "9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e"
	CANCEL_0X_ORDER = "be80e5687d7095071b7c4e7a56e0e67bfb9e8a39352f1690fdf74c1ee935c75e"
)
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	storage *SQLStorage
	l1_storage *SQLStorage

	fill_order_id int64 = 0			// during event processing, holds id of record in mktord from Fill evt
	market_order_id int64 = 0

	Info    *log.Logger

	eclient *ethclient.Client
	rpcclient *rpc.Client

	all_contracts map[string]interface{}
	caddrs *ContractAddresses
	zerox_trade_abi *abi.ABI
	wallet_abi *abi.ABI
	ctrct_zerox_trade *ZeroX
	trading_abi *abi.ABI

	evt_market_order,_ = hex.DecodeString(MARKET_ORDER)
	evt_cancel_0x_order,_ = hex.DecodeString(CANCEL_0X_ORDER)
)
type ExecWalletTxInputStruct struct {
	To common.Address `abi:"_to"`
	Data []byte `abi:"_data"`
	Value *big.Int `abi:"_value"`
	Payment *big.Int `abi:"_payment"`
	ReferralAddress common.Address `abi:"_referralAddress"`
	Fingerprint [32]byte `abi:"_fingerprint"`
	DesiredSignerBalance *big.Int `abi:"_desiredSignerBalance"`
	MaxExchangeRateInDai *big.Int `abi:"_maxExchangeRateInDai"`
	RevertOnFailure bool `abi:"_revertOnFailure"`
}
type TradeInputStruct struct {
	RequestedFillAmount		*big.Int `abi:"_requestedFillAmount"`
	Fingerprint				[32]byte `abi:"_fingerprint"`
	TradeGroupId			[32]byte `abi:"_tradeGroupId"`
	MaxProtocolFeeDai		*big.Int `abi:"_maxProtocolFeeDai"`
	MaxTrades				*big.Int `abi:"_maxTrades"`
	Orders					[]IExchangeOrder `abi:"_orders"`
	Signatures				[][]byte `abi:"_signatures"`
}
type CancelPrdersInputStruct struct {
	Orders					[]IExchangeOrder `abi:"_orders"`
	Signatures				[][]byte `abi:"_signatures"`
	MaxProtocolFeeDai		*big.Int `abi:"_maxProtocolFeeDai"`

}
func get_order_data(o *IExchangeOrder) (zeroex.Order,common.Hash) {

	var zero_order zeroex.Order
	zero_order.ChainID=big.NewInt(caddrs.ChainId)
	zero_order.ExchangeAddress.SetBytes(caddrs.ZeroxXchg.Bytes())
	zero_order.MakerAddress.SetBytes(o.MakerAddress.Bytes())
	zero_order.MakerAssetData = make([]byte,len(o.MakerAssetData))
	copy(zero_order.MakerAssetData,o.MakerAssetData)
	zero_order.MakerFeeAssetData = make([]byte,len(o.MakerFeeAssetData))
	copy(zero_order.MakerFeeAssetData,o.MakerFeeAssetData)
	zero_order.MakerAssetAmount = new(big.Int)
	zero_order.MakerAssetAmount.Set(o.MakerAssetAmount)
	zero_order.MakerFee = new(big.Int)
	zero_order.MakerFee.Set(o.MakerFee)
	zero_order.TakerAddress.SetBytes(o.TakerAddress.Bytes())
	zero_order.TakerAssetData = make([]byte,len(o.TakerAssetData))
	copy(zero_order.TakerAssetData,o.TakerAssetData)
	zero_order.TakerFeeAssetData = make([]byte,len(o.TakerFeeAssetData))
	copy(zero_order.TakerFeeAssetData,o.TakerFeeAssetData)
	zero_order.TakerAssetAmount = new(big.Int)
	zero_order.TakerAssetAmount.Set(o.TakerAssetAmount)
	zero_order.TakerFee = new(big.Int)
	zero_order.TakerFee.Set(o.TakerFee)
	zero_order.SenderAddress.SetBytes(o.SenderAddress.Bytes())
	zero_order.FeeRecipientAddress.SetBytes(o.FeeRecipientAddress.Bytes())
	zero_order.ExpirationTimeSeconds = new(big.Int)
	zero_order.ExpirationTimeSeconds.Set(o.ExpirationTimeSeconds)
	zero_order.Salt = new(big.Int)
	zero_order.Salt.Set(o.Salt)
	hash,err:=zero_order.ComputeOrderHash()
	if err!=nil {
		Fatalf("can't compute ZeroX order hash: %v\n",err)
	}
	Info.Printf("get_order_hash() returning %v\n",hash.String())
	return zero_order,hash
}
func decode_0x_orders(input_data []byte,method_sig []byte) (map[string]*ztypes.OrderInfo,map[string]*ZxMeshOrderSpec) {

	output1 := make(map[string]*ztypes.OrderInfo,0)
	output2 := make(map[string]*ZxMeshOrderSpec,0)

	var trade_input_data_decoded TradeInputStruct
	var cancel_order_input_data_decoded CancelPrdersInputStruct
	var decoded_orders []IExchangeOrder
	var decoded_signatures [][]byte

	zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
	if 0 == bytes.Compare(method_sig,zeroex_trade_sig) {
		method, err := zerox_trade_abi.MethodById(method_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&trade_input_data_decoded, input_data)
		if err != nil {
			Fatalf("Couldn't decode input of tx: %v",err)
		}
		decoded_orders = trade_input_data_decoded.Orders
		decoded_signatures = trade_input_data_decoded.Signatures
	}
	zeroex_cancel_sig,_ := hex.DecodeString("4ea96c30")
	if 0 == bytes.Compare(method_sig,zeroex_cancel_sig) {
		method, err := zerox_trade_abi.MethodById(method_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&cancel_order_input_data_decoded, input_data)
		if err != nil {
			Fatalf("Couldn't decode input of tx: %v",err)
		}
		decoded_orders=cancel_order_input_data_decoded.Orders
		decoded_signatures=cancel_order_input_data_decoded.Signatures
	}
	if len(decoded_orders) > 0 {
		for i,order := range decoded_orders {
			ord,h := get_order_data(&order)
			hash_str := h.String()
			Info.Printf(
				"Order %v (%v), maker amount = %v, taker amount=%v\n",
				i,hash_str,order.MakerAssetAmount,order.TakerAssetAmount,
			)
			order_info := new(ztypes.OrderInfo)
			order_info.OrderHash.SetBytes(h.Bytes())
			order_info.SignedOrder=new(zeroex.SignedOrder)
			order_info.SignedOrder.Signature=make([]byte,len(decoded_signatures[i]))
			order_info.SignedOrder.Order = ord
			order_info.FillableTakerAssetAmount = big.NewInt(0) // this value is incorrect, but we don't have the correct one
			copy(order_info.SignedOrder.Signature,decoded_signatures[i])
			output1[hash_str]=order_info
			ospec := get_ospec(ord.MakerAssetData,&hash_str)
			output2[hash_str] = ospec
		}
	} else {
		Info.Printf("Undefined behavior: no orders detected on the input of ZeroXTrade::trade()")
		os.Exit(1)
	}

	return output1,output2
}
func extract_orders_from_input(tx_data []byte) (map[string]*ztypes.OrderInfo,map[string]*ZxMeshOrderSpec) {
	// returns orders in one map and initial amounts in another map
	if len(tx_data) < 32 {
		return make(map[string]*ztypes.OrderInfo,0),make(map[string]*ZxMeshOrderSpec,0)

	}
	input_sig := tx_data[:4]
	decoded_sig ,_ := hex.DecodeString("78dc0eed")
	if 0 == bytes.Compare(input_sig,decoded_sig) {
		input_data_raw:= tx_data[4:]
		var input_data ExecWalletTxInputStruct
		method, err := wallet_abi.MethodById(decoded_sig)
		if err != nil {
			Fatalf("Method not found")
		}
		err = method.Inputs.Unpack(&input_data, input_data_raw)
		if err != nil {
			Fatalf("Couldn't decode input of tx %v",err)
		}

		// check for internal transactions for the Wallet Registry contract
		if len(input_data.Data) >= 4 {
			input_sig := input_data.Data[:4]
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("augur_wallet_call: ZeroEx::trade()\n")
				return decode_0x_orders(input_data.Data[4:],zeroex_trade_sig)
			}
			zeroex_cancel_sig,_ := hex.DecodeString("4ea96c30")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				Info.Printf("augur_wallet_call: ZeroEx::cancelOrder()\n")
				return decode_0x_orders(input_data.Data[4:],zeroex_cancel_sig)
			}
		}
	} else {
		if len(input_sig) >= 4 {
			input_data_raw:= tx_data[4:]
			//Info.Printf("tx input= %v\n",hex.EncodeToString(input_data_raw))
			zeroex_trade_sig ,_ := hex.DecodeString("2f562016")
			if 0 == bytes.Compare(input_sig,zeroex_trade_sig) {
				return decode_0x_orders(input_data_raw,zeroex_trade_sig)
			}
			zeroex_cancel_sig,_ := hex.DecodeString("4ea96c30")
			if 0 == bytes.Compare(input_sig,zeroex_cancel_sig) {
				return decode_0x_orders(input_data_raw,zeroex_cancel_sig)
			}
		}
	}
	return make(map[string]*ztypes.OrderInfo,0),make(map[string]*ZxMeshOrderSpec,0)
}
func get_ospec(maker_asset_data []byte,order_hash *string) *ZxMeshOrderSpec {

	var copts = new(bind.CallOpts)
	adata,err := ctrct_zerox_trade.DecodeAssetData(copts,maker_asset_data)
	if err!=nil {
		Info.Printf("couldn't decode asset data for order %v : %v\n",*order_hash,err)
		os.Exit(1)
	}
	unpacked_id,err := ctrct_zerox_trade.UnpackTokenId(copts,adata.TokenIds[0])
	if err!=nil {
		Info.Printf("Unpack token id failed for order %v: %v\n",*order_hash,err)
		os.Exit(1)
	}
	return &unpacked_id
}
func create_mesh_history_table(ss *SQLStorage) {

	var query string

	query = "CREATE TABLE IF NOT EXISTS mesh_history ( " +
				"id                      BIGSERIAL PRIMARY KEY," +
			//-- Event fields:
				"time_stamp              TIMESTAMPTZ NOT NULL," +
			"fillable_amount         DECIMAL(32,18) NOT NULL," +
			"evt_code                SMALLINT NOT NULL," +
			//-- Augur fields:
			"market_aid              BIGINT NOT NULL," +
			"outcome_idx             SMALLINT NOT NULL," +
			"otype                   SMALLINT NOT NULL," +
			"price                   DECIMAL(32,18) NOT NULL," +
			//-- Fill fields:
			"amount_fill             DECIMAL(32,18) DEFAULT 0.0," +
			//-- `Order` struct follows:
			"order_hash              CHAR(66) NOT NULL," +
			"chain_id                INT NOT NULL," +
			"exchange_addr           CHAR(42) NOT NULL," +
			"maker_addr              CHAR(42) NOT NULL," +
			"maker_asset_data        TEXT NOT NULL," +
			"maker_fee_asset_data    TEXT NOT NULL, " +
			"maker_asset_amount      DECIMAL(32,18) NOT NULL," +
			"maker_fee               DECIMAL(32,18) NOT NULL," +
			"taker_address           CHAR(42) NOT NULL," +
			"taker_asset_data        TEXT NOT NULL," +
			"taker_fee_asset_data    TEXT NOT NULL," +
			"taker_asset_amount      DECIMAL(32,18) NOT NULL," +
			"taker_fee               DECIMAL(32,18) NOT NULL," +
			"sender_address          CHAR(42) NOT NULL," +
			"fee_recipient_address   CHAR(42) NOT NULL," +
			"expiration_time         TIMESTAMPTZ NOT NULL," +
			"salt                    TEXT NOT NULL,"+
			"signature               TEXT," +
			"UNIQUE(order_hash,evt_code) "+
		")";
	_,err := ss.Db().Exec(query)
	if err!=nil {
		Info.Printf("Can't create mesh historic table: %v\n",err)
		os.Exit(1)
	}
}
func get_all_fill_events() []string {
	// Return Value: array of transaction ids that processed Fill event

	records := make([]string,0,1024)

	var query string
	query = "SELECT tx.tx_hash FROM evt_log,transaction as tx " +
			"WHERE evt_log.tx_id=tx.id AND ((topic0_sig='6869791f') OR (topic0_sig='be80e568'))"

	rows,err := l1_storage.Db().Query(query)
	if err != nil {
		Info.Printf("Error in query %v: %v\n",query,err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var tx_hash string
		err=rows.Scan(&tx_hash)
		if err!=nil {
			Info.Printf("DB error: %v, q=%v",err,query)
			os.Exit(1)
		}
		records = append(records,tx_hash)
	}
	Info.Printf("query %v returned %v rows\n",query,len(records))
	return records
}
func process_onchain_fill_events(txs []string) {

	ctx := context.Background()
	for _,tx_hash := range(txs) {
		hash:=common.HexToHash(tx_hash)
		tx,_,err := eclient.TransactionByHash(ctx,hash)
		if err!= nil {
			Info.Printf("Error getting transaction %v: %v\n",tx_hash,err)
			os.Exit(1)
		}
		orders,ospecs := extract_orders_from_input(tx.Data())
		receipt,err := eclient.TransactionReceipt(ctx,hash)
		if err != nil {
			Info.Printf("Error getting receipt for %v: %v\n",tx_hash,err)
			os.Exit(1)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			Info.Printf("Failed transaction, skipping\n")
			continue
		}
		bnum := receipt.BlockNumber.Int64()
		timestamp,err := storage.Get_block_timestamp(bnum)
		if err != nil {
			Info.Printf("Error getting timestamp for block %v : %v\n",bnum,err)
			os.Exit(1)
		}
		//Info.Printf("Receipt has %v logs\n",len(receipt.Logs))
		for _,log := range receipt.Logs {
			if len(log.Topics) == 0 {
				continue
			}
			if bytes.Equal(log.Topics[0].Bytes(),evt_market_order) {
				Info.Printf("Processing tx_hash %v\n",tx_hash)
				var evt EOrderEvent
				err := trading_abi.Unpack(&evt,"OrderEvent",log.Data)
				if err != nil {
					Fatalf("Event OrderEvent decode error: %v",err)
					return
				}
				if !bytes.Equal(log.Address.Bytes(),caddrs.AugurTrading.Bytes()) {
					Info.Printf(
						"OrderEvent received and ignored (belongs to different contract: %v) at block %v (EVENT_IGNORE)",
						log.Address.String(),receipt.BlockNumber.Int64(),
					)
					continue
				}
				var order_hash_obj = common.BytesToHash(evt.OrderId[:])
				var order_hash = order_hash_obj.String()
				zorder := orders[order_hash]
				if zorder == nil {
					Info.Printf("Error, couldn't find order %v in tx input",order_hash)
					os.Exit(1)
				}
				initial_amount := zorder.SignedOrder.MakerAssetAmount
				amount_filled := evt.Uint256Data[6]
				mesh_evt_code := MeshEvtFullyFilled
				if 0 != initial_amount.Cmp(amount_filled) {
					mesh_evt_code = MeshEvtFilled
				}
				Info.Printf(
					"OrderEvent: 0x Mesh Evt inset: ohash: %v, ts: %v, fill_amount: %v, evt: %v\n",
					order_hash,timestamp,amount_filled.String,mesh_evt_code,
				)
				/*storage.Insert_0x_mesh_order_event(
					timestamp,zorder,ospecs[order_hash],amount_filled,mesh_evt_code,
				)
				market_aid:=storage.Lookup_address_id(ospec.Market.String())
				storage.Update_future_price_estimates(market_aid,int(ospec.Outcome),timestamp)
				*/
				Info.Printf(
					"\tOrder type: %v, market: %v, outcome %v, price: %v\n",
					ospecs[order_hash].Type,
					ospecs[order_hash].Market.String(),
					ospecs[order_hash].Outcome,
					ospecs[order_hash].Price.String(),
				)
			}
			if bytes.Equal(log.Topics[0].Bytes(),evt_cancel_0x_order) {
				Info.Printf("Processing tx_hash %v\n",tx_hash)

				var evt ECancelZeroXOrder
				err := trading_abi.Unpack(&evt,"CancelZeroXOrder",log.Data)
				if err != nil {
					Fatalf("Event CancelZeroXOrder decode error: %v",err)
					return
				}
				if !bytes.Equal(log.Address.Bytes(),caddrs.AugurTrading.Bytes()) {
					Info.Printf(
						"Cancel0xOrder event received and ignored (belongs to different contract: %v) at block %v (EVENT_IGNORE)",
						log.Address.String(),receipt.BlockNumber.Int64(),
					)
					continue
				}
				var order_hash_obj = common.BytesToHash(evt.OrderHash[:])
				var order_hash = order_hash_obj.String()
				zorder := orders[order_hash]
				if zorder == nil {
					Info.Printf("Error, couldn't find order %v in tx input",order_hash)
					os.Exit(1)
				}
				spec:=ospecs[order_hash]
				Info.Printf(
					"CancelOrder: 0x Mesh Evt insert: ohash: %v, ts: %v, evt: %v\n",order_hash,timestamp)
				/*
				storage.Insert_0x_mesh_order_event(timestamp,zorder,spec,nil,p.MeshEvtCancelled)
				market_aid:=storage.Lookup_address_id(ospec.Market.String())
				storage.Update_future_price_estimates(market_aid,int(ospec.Outcome),timestamp)
				*/
				Info.Printf(
					"\tOrder type: %v, market: %v, outcome %v, price: %v\n",
					spec.Type,spec.Market.String(),spec.Outcome,spec.Price.String(),
				)
			}
		}
	}
}
func main() {

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
	}
	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	var err error
	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)

	eclient = ethclient.NewClient(rpcclient)
	var dummy_int_var,market_id int64
	l1_storage = New_sql_storage(
		&dummy_int_var,Info,
		os.Getenv("L1_HOST"),
		os.Getenv("L1_DB"),
		os.Getenv("L1_USER"),
		os.Getenv("L1_PASSWD"),
	)
	storage = Connect_to_storage(&market_id,Info)

	caddrs_obj,err := storage.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs=&caddrs_obj

	all_contracts = Load_all_artifacts("../abis/augur-artifacts-abi.json")

	// Augur service involves 39 contracts in total. We only use a few of them
	zerox_trade_abi = Abi_from_artifacts(&all_contracts,"ZeroXTrade")
	ctrct_zerox_trade, err = NewZeroX(caddrs.ZeroxTrade,eclient)
	if err != nil {
		Fatalf("Failed to instantiate a ZeroX contract: %v", err)
	}
	wallet_abi = Abi_from_artifacts(&all_contracts,"AugurWalletRegistry")
	ctrct_zerox_trade, err = NewZeroX(caddrs.ZeroxTrade,eclient)
	if err != nil {
		Fatalf("Failed to instantiate a ZeroX contract: %v", err)
	}
	trading_abi = Abi_from_artifacts(&all_contracts,"AugurTrading")

	create_mesh_history_table(l1_storage)
	txs := get_all_fill_events()
	process_onchain_fill_events(txs)
}
