/// Import mesh events
package main
import (
	"os"
	"log"
	"strconv"
	"fmt"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/0xProject/0x-mesh/zeroex"
	ztypes "github.com/0xProject/0x-mesh/common/types"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	storage *SQLStorage

	fill_order_id int64 = 0			// during event processing, holds id of record in mktord from Fill evt
	market_order_id int64 = 0

	Info    *log.Logger

	caddrs *ContractAddresses
	eclient *ethclient.Client
	//rpcclient *rpc.Client
	zerox_contract *ZeroX
	adecoder *zeroex.AssetDataDecoder
)
func read_mesh_import_file(fname string) []MeshEvent {

	var num_fields int = 23
	data,err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Can't open file %v containing tab-separated 0x mesh events\n")
		os.Exit(1)
	}
	records := string(data)
	lines := strings.Split(records,"\n")
	output := make([]MeshEvent,0,1024)
	for i:=0 ; i<len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		if lines[i] == "\n" {
			continue
		}
		fields := strings.Split(lines[i],"\t")
		if len(fields) != num_fields {
			fmt.Printf("Error at line %v, number of fields must be %v and we have %v\n",i+1,num_fields,len(fields))
			fmt.Printf("line:\n%v\n",lines[i])
			os.Exit(1)
		}
		var rec MeshEvent
		var intval int
		intval,err = strconv.Atoi(fields[0])
		if err != nil {
			fmt.Printf("can't convert timestamp %v at line %v to integer\n",fields[0],i)
			os.Exit(1)
		}
		rec.Timestamp = int64(intval)
		intval,err = strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("can't convert EvtCode %v at line %v to integer\n",fields[1],i)
			os.Exit(1)
		}
		rec.EvtCode = intval
		rec.OrderHash = fields[2]
		intval,err = strconv.Atoi(fields[3])
		if err != nil {
			fmt.Printf("can't convert ChainId %v at line %v to integer\n",fields[3],i)
			os.Exit(1)
		}
		rec.ChainId = intval
		rec.ExchangeAddress = fields[4]
		rec.MakerAddress = fields[5]
		rec.MakerAssetData = fields[6]
		rec.MakerFeeAssetData = fields[7]
		rec.MakerAssetAmount = fields[8]
		rec.MakerFee = fields[9]
		rec.TakerAddress = fields[10]
		rec.TakerAssetData = fields[11]
		rec.TakerFeeAssetData = fields[12]
		rec.TakerAssetAmount = fields[13]
		rec.TakerFee = fields[14]
		rec.SenderAddress = fields[15]
		rec.FeeRecipientAddress = fields[16]
		intval,err = strconv.Atoi(fields[17])
		if err != nil {
			fmt.Printf("can't convert ExpirationTime %v at line %v to integer\n",fields[17],i)
			os.Exit(1)
		}
		rec.ExpirationTime = int64(intval)
		rec.Salt = fields[18]
		rec.Signature = fields[19]
		rec.FillableAmount = fields[20]
		rec.MarketAddress = fields[21]
		intval,err = strconv.Atoi(fields[22])
		if err != nil {
			fmt.Printf("can't convert NumTicks %v at line %v to integer\n",fields[22],i)
			os.Exit(1)
		}
		rec.NumTicks = intval
		output = append(output,rec)
	}
	return output
}
func dump_0x_order(o *zeroex.Order) {
	fmt.Printf("0x Mesh Order {\n")
	fmt.Printf("\tChainId: %v\n",o.ChainID.String())
	fmt.Printf("\tExchangeAddress: %v\n",o.ExchangeAddress.String())
	fmt.Printf("\tMakerAddress: %v\n",o.MakerAddress.String())
	fmt.Printf("\tMakerAssetData: %v\n",hex.EncodeToString(o.MakerAssetData))
	fmt.Printf("\tMakerFeeAssetData: %v\n",hex.EncodeToString(o.MakerFeeAssetData))
	fmt.Printf("\tMakerAssetAmount: %v\n",o.MakerAssetAmount.String())
	fmt.Printf("\tMakerFee: %v\n",o.MakerFee.String())
	fmt.Printf("\tTakerAddress: %v\n",o.TakerAddress.String())
	fmt.Printf("\tTakerAssetData: %v\n",hex.EncodeToString(o.TakerAssetData))
	fmt.Printf("\tTakerFeeAssetData: %v\n",hex.EncodeToString(o.TakerFeeAssetData))
	fmt.Printf("\tTakerAssetAmount: %v\n",o.TakerAssetAmount.String())
	fmt.Printf("\tTakerFee: %v\n",o.TakerFee.String())
	fmt.Printf("\tSenderAddress: %v\n",o.SenderAddress.String())
	fmt.Printf("\tFeeRecipientAddress: %v\n",o.FeeRecipientAddress.String())
	fmt.Printf("\tExpirationTimeSeconds: %v\n",o.ExpirationTimeSeconds.String())
	fmt.Printf("\tSalt: %v\n",o.Salt.String())
	fmt.Printf("}\n")
}
func make_0x_order(evt *MeshEvent) (zeroex.Order,common.Hash) {

	var err error
	var zero_order zeroex.Order
	zero_order.ChainID=new(big.Int)
	zero_order.ChainID.SetInt64(int64(evt.ChainId))
	zero_order.ExchangeAddress = common.HexToAddress(evt.ExchangeAddress)
	zero_order.MakerAddress = common.HexToAddress(evt.MakerAddress)
	zero_order.MakerAssetData,err = hex.DecodeString(evt.MakerAssetData)
	if err != nil {
		Fatalf("Error decoding MakerAssetData : %v\n",err)
	}
	zero_order.MakerFeeAssetData,err = hex.DecodeString(evt.MakerFeeAssetData)
	if err != nil {
		Fatalf("Error decoding MakerFeeAssetData : %v\n",err)
	}
	zero_order.MakerAssetAmount = new(big.Int)
	zero_order.MakerAssetAmount.SetString(evt.MakerAssetAmount,10)
	zero_order.MakerFee = new(big.Int)
	zero_order.MakerFee.SetString(evt.MakerFee,10)
	zero_order.TakerAddress = common.HexToAddress(evt.TakerAddress)
	zero_order.TakerAssetData,err = hex.DecodeString(evt.TakerAssetData)
	if err != nil {
		Fatalf("Error decoding TakerAssetData : %v\n",err)
	}
	zero_order.TakerFeeAssetData,err = hex.DecodeString(evt.TakerFeeAssetData)
	if err != nil {
		Fatalf("Error decoding TakerFeeAssetData : %v\n",err)
	}
	zero_order.TakerAssetAmount = new(big.Int)
	zero_order.TakerAssetAmount.SetString(evt.TakerAssetAmount,10)
	zero_order.TakerFee = new(big.Int)
	zero_order.TakerFee.SetString(evt.TakerFee,10)
	zero_order.SenderAddress = common.HexToAddress(evt.SenderAddress)
	zero_order.FeeRecipientAddress = common.HexToAddress(evt.FeeRecipientAddress)
	zero_order.ExpirationTimeSeconds = new(big.Int)
	zero_order.ExpirationTimeSeconds.SetInt64(evt.ExpirationTime)
	zero_order.Salt = new(big.Int)
	zero_order.Salt.SetString(evt.Salt,10)
	hash,err:=zero_order.ComputeOrderHash()
	if err!=nil {
		Fatalf("can't compute ZeroX order hash: %v\n",err)
	}
	//dump_0x_order(&zero_order)
	return zero_order,hash
}
func get_ospec(order *zeroex.SignedOrder,order_hash *string) (ZxMeshOrderSpec,error) {

	var copts = new(bind.CallOpts)
	adata,err := zerox_contract.DecodeAssetData(copts,order.MakerAssetData)
	if err!=nil {
		fmt.Printf("couldn't decode asset data for order %v : %v\n",*order_hash,err)
		return ZxMeshOrderSpec{},err
	}
	unpacked_id,err := zerox_contract.UnpackTokenId(copts,adata.TokenIds[0])
	if err!=nil {
		fmt.Printf("Unpack token id failed for order %v: %v\n",*order_hash,err)
		return ZxMeshOrderSpec{},err
	}
	return unpacked_id,err
}
func main() {	// returns 0 - no changes, 2 - day was added

	if len(os.Args) < 2 {
		fmt.Printf("usage: %v [mesh_events_file]\n",os.Args[0])
		os.Exit(1)
	}
	events := read_mesh_import_file(os.Args[1])

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)

	caddrs_obj,err := storage.Get_contract_addresses()
	if err != nil {
		log.Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs = &caddrs_obj
	if caddrs.ChainId == 0 {
		log.Fatalf("ChainID = 0, db is not initialized")
	}
	adecoder = zeroex.NewAssetDataDecoder()
	eclient, err = ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("ZeroX contract = %v\n",caddrs.ZeroxTrade.String())
	zerox_contract, err = NewZeroX(
			common.HexToAddress(
				caddrs.ZeroxTrade.String(),
			),
			eclient,
	)
	for i,evt := range events {
		evt_code := MeshEvtCode(evt.EvtCode)
		switch evt_code {
		case MeshEvtFilled,MeshEvtFullyFilled,MeshEvtCancelled:
			fmt.Printf("Skipping order %v , evt_code=%v it's a blockchain event",evt.OrderHash,evt_code)
			continue
		}
		order,hash := make_0x_order(&evt)
		if hash.String() != evt.OrderHash {
			fmt.Printf("line %v: computed order hash %v is different than stored order hash %v. some field in the order has wrong value\n",i+1,hash.String(),evt.OrderHash)
			os.Exit(1)
		}
		fmt.Printf("line %v: imported order with hash %v\n",i+1,evt.OrderHash)
		var order_info ztypes.OrderInfo
		order_info.OrderHash.SetBytes(hash.Bytes())
		order_info.SignedOrder = new(zeroex.SignedOrder)
		order_info.SignedOrder.Order = order
		order_info.SignedOrder.Signature,err = hex.DecodeString(evt.Signature)
		if err!= nil {
			fmt.Printf("can't decode signature HEX for order %v at line %v\n",evt.OrderHash,i+1)
			os.Exit(1)
		}
		order_info.FillableTakerAssetAmount = new(big.Int)
		order_info.FillableTakerAssetAmount.SetString(evt.FillableAmount,10)
		ospec,err := get_ospec(order_info.SignedOrder,&evt.OrderHash)
		if err!=nil {
			fmt.Printf("Error decoding market data: %v\n",err)
			continue
		}
//		Dump_0x_mesh_order(Info,&order_info)
//		DumpOrderSpec(Info,&ospec)
		aid := storage.Lookup_or_create_address(evt.MakerAddress,0,0)
		storage.Insert_0x_mesh_order_event(
			0,
			aid,
			evt.Timestamp,
			&order_info,
			&ospec,
			nil,//amount_filled (only events from blockchain have this)
			evt_code,
		)
	}
}
