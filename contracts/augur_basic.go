package contracts

import (
	"os"
	"fmt"
	"bytes"
	"math/big"
	"io/ioutil"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	//ztypes "github.com/0xProject/0x-mesh/common/types"
	"github.com/0xProject/0x-mesh/zeroex"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func dump_abi_events(a *abi.ABI) {

	fmt.Printf("Events:\n")
	for evt:=range a.Events {
		fmt.Printf("\t%v\t%v\n",a.Events[evt].ID().String(),evt)
	}

}
func dump_abi_methods(a *abi.ABI) {
	fmt.Printf("Methods:\n")
	for meth := range a.Methods {
		fmt.Printf("\t%v\t%v\n",hex.EncodeToString(a.Methods[meth].ID()),meth)
	}
}
func dump_all_artifacts(contracts *map[string]interface{}) {

	for contract_name , _ := range (*contracts) {
		fmt.Printf("Contract: %v\n",contract_name)
		abi:=Abi_from_artifacts(contracts,contract_name)
		dump_abi_events(abi)
		dump_abi_methods(abi)
	}
}
func Load_all_artifacts(filename string) map[string]interface{} {

	abi_data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Can't load abi file %v : %v\n",filename,err)
		os.Exit(1)
	}
	all_abis_rdr := bytes.NewReader(abi_data)
	if err != nil {
		fmt.Printf("Can't create reader for ABI from %v : %v\n",filename,err)
		os.Exit(1)
	}
	byte_data, err := ioutil.ReadAll(all_abis_rdr)
	if err != nil {
		fmt.Printf("Can't read ABI data for %v : %v\n",filename,err)
		os.Exit(1)
	}
	var contracts map[string]interface{}
	json.Unmarshal([]byte(byte_data), &contracts)
	return contracts
}
func Abi_from_artifacts(contracts *map[string]interface{},contract string) *abi.ABI {

	contract_abi:=(*contracts)[contract]
	contract_bytes, _ := json.Marshal(contract_abi) // convert back to JSON so Ethereum package can work
	rdr := bytes.NewReader(contract_bytes)
	ctrct_abi,err := abi.JSON(rdr)
	if err != nil {
		fmt.Printf("Can't read ABI for contract %v : %v\n",contract,err)
		os.Exit(1)
	}
	return &ctrct_abi
}
func load_abi(fname string) *abi.ABI {

	abi_data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Failed to ReadFile %v : %v\n",fname,err)
		os.Exit(1)
	}
	abi_rdr := bytes.NewReader(abi_data)
	if err != nil {
		fmt.Printf("Failed to create new Reader for %v : %v\n",fname,err)
		os.Exit(1)
	}
	abi,err := abi.JSON(abi_rdr)
	if err != nil {
		fmt.Printf("Failed to read ABI data for %v : %v\n",fname,err)
		os.Exit(1)
	}
	return &abi
}
func Get_contract_addresses_from_net(augur_trading_address common.Address,eclient *ethclient.Client) (ContractAddresses,error) {

	var caddrs ContractAddresses
	var err error
	var copts = new(bind.CallOpts)
	var key common.Hash

	caddrs.AugurTrading = augur_trading_address

	var ctrct_augurtrading *AugurTrading
	ctrct_augurtrading,err = NewAugurTrading(caddrs.AugurTrading,eclient)
	if err != nil {
		newerr := errors.New(fmt.Sprintf("Couldn't create AugurTrading instance: %v",err.Error()))
		return caddrs,newerr
	}

	caddrs.Augur,err = ctrct_augurtrading.Augur(copts)
	if err != nil {
		newerr := errors.New(fmt.Sprintf("Call to AugurTrading.sol:Augur() failed: %v",err.Error()))
		return caddrs,newerr
	}

	ctrct_augur,err := NewAugur(caddrs.Augur,eclient)
	if err != nil {
		newerr := errors.New(fmt.Sprintf("Couldn't create Augur contract instance: %v",err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("Cash"))
	caddrs.Dai,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("ProfitLoss"))
	caddrs.PL,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("ZeroXTrade"))
	caddrs.ZeroxTrade,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("ZeroXExchange"))
	caddrs.ZeroxXchg,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	caddrs.GenesisUniverse,err = ctrct_augur.GenesisUniverse(copts)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Getting GenesisUniverse failed: %v",err.Error()))
		return caddrs,newerr
	}

	var ctrct_universe *Universe
	ctrct_universe,err = NewUniverse(caddrs.GenesisUniverse,eclient)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Can't create Universe contract: %v,",err.Error()))
		return caddrs,newerr
	}

	caddrs.Reputation,err = ctrct_universe.GetReputationToken(copts)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Cant' get Reputation Token v2: %v",err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("AugurWalletRegistry"))
	caddrs.WalletReg,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("AugurWalletRegistryV2"))
	caddrs.WalletReg2,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("FillOrder"))
	caddrs.FillOrder,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("UniswapV2Factory"))
	caddrs.UniswapV2Factory,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("UniswapV2Router02"))
	caddrs.UniswapV2Router02,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("WETH9"))
	caddrs.WETH9,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	var ctrct_uniswapv2factory *UniswapV2Factory
	ctrct_uniswapv2factory,err = NewUniswapV2Factory(caddrs.UniswapV2Factory,eclient)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Can't create instance of UniswavV2Factory: %v",err))
		return caddrs,newerr
	}

	caddrs.EthXchg,err = ctrct_uniswapv2factory.GetPair(copts,caddrs.WETH9,caddrs.Dai)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("CAll to UniswapV2Factory::GetPair failed: %v",err))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("ShareToken"))
	caddrs.ShareToken,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("CreateOrder"))
	caddrs.CreateOrder,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("LegacyReputationToken"))
	caddrs.LegacyReputationToken,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("BuyParticipationTokens"))
	caddrs.BuyParticipationTokens,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("RedeemStake"))
	caddrs.RedeemStake,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("WarpSync"))
	caddrs.WarpSync,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("HotLoading"))
	caddrs.HotLoading,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("Affiliates"))
	caddrs.Affiliates,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("AffiliateValidator"))
	caddrs.AffiliateValidator,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("Time"))
	caddrs.Time,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("CancelOrder"))
	caddrs.CancelOrder,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("Orders"))
	caddrs.Orders,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("SimulateTrade"))
	caddrs.SimulateTrade,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("Trade"))
	caddrs.Trade,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("OICash"))
	caddrs.OICash,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("AuditFunds"))
	caddrs.AuditFunds,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("USDC"))
	caddrs.USDC,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("USDT"))
	caddrs.USDT,err = ctrct_augur.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}
	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("RelayHubV2"))
	caddrs.RelayHubV2,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed: %v",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	return caddrs,nil
}
func Contract_addresses_match(ca1 *ContractAddresses,ca2 *ContractAddresses) (int,error) {
	var num_mismatches int = 0
	var all_errors string = ""
	if !bytes.Equal(ca1.Augur.Bytes(),ca2.Augur.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "Augur address isn't matching "
	}
	if !bytes.Equal(ca1.AugurTrading.Bytes(),ca2.AugurTrading.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "AugurTrading address isn't matching "
	}
	if !bytes.Equal(ca1.PL.Bytes(),ca2.PL.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "ProfitLoss address isn't matching "
	}
	if !bytes.Equal(ca1.ZeroxTrade.Bytes(),ca2.ZeroxTrade.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "ZeroXTrade address isn't matching "
	}
	if !bytes.Equal(ca1.ZeroxXchg.Bytes(),ca2.ZeroxXchg.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "Exchange (0x) address isn't matching "
	}
	if !bytes.Equal(ca1.Dai.Bytes(),ca2.Dai.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "Cash (DAI) address isn't matching "
	}
	if !bytes.Equal(ca1.Reputation.Bytes(),ca2.Reputation.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "REPv2 token address isn't matching "
	}
	if !bytes.Equal(ca1.WalletReg.Bytes(),ca2.WalletReg.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "AugurWalletRegistry (v1) address isn't matching "
	}
	if !bytes.Equal(ca1.WalletReg2.Bytes(),ca2.WalletReg2.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "AugurWalletRegistry2 address isn't matching "
	}
	if !bytes.Equal(ca1.FillOrder.Bytes(),ca2.FillOrder.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "FillOrder address isn't matching "
	}
	if !bytes.Equal(ca1.EthXchg.Bytes(),ca2.EthXchg.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "EthXchg (Uniswap pair WETH9 + DAI) address isn't matching "
	}
	if !bytes.Equal(ca1.ShareToken.Bytes(),ca2.ShareToken.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "ShareToken address isn't matching "
	}
	if !bytes.Equal(ca1.GenesisUniverse.Bytes(),ca2.GenesisUniverse.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "GenesisUniverse address isn't matching "
	}
	if !bytes.Equal(ca1.CreateOrder.Bytes(),ca2.CreateOrder.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "CreateOrder address isn't matching "
	}
	if !bytes.Equal(ca1.LegacyReputationToken.Bytes(),ca2.LegacyReputationToken.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "LegacyReputationToken address isn't matching "
	}
	if !bytes.Equal(ca1.BuyParticipationTokens.Bytes(),ca2.BuyParticipationTokens.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "BuyParticipationToken address isn't matching "
	}
	if !bytes.Equal(ca1.RedeemStake.Bytes(),ca2.RedeemStake.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "RedeemStake address isn't matching "
	}
	if !bytes.Equal(ca1.WarpSync.Bytes(),ca2.WarpSync.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "WarpSync address isn't matching "
	}
	if !bytes.Equal(ca1.HotLoading.Bytes(),ca2.HotLoading.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "HotLoading address isn't matching "
	}
	if !bytes.Equal(ca1.Affiliates.Bytes(),ca2.Affiliates.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "Affiliates address isn't matching "
	}
	if !bytes.Equal(ca1.AffiliateValidator.Bytes(),ca2.AffiliateValidator.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "AffiliateValidator address isn't matching "
	}
	if !bytes.Equal(ca1.Time.Bytes(),ca2.Time.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "Time address isn't matching "
	}
	if !bytes.Equal(ca1.CancelOrder.Bytes(),ca2.CancelOrder.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "CancelOrder address isn't matching "
	}
	if !bytes.Equal(ca1.Orders.Bytes(),ca2.Orders.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "Orders address isn't matching "
	}
	if !bytes.Equal(ca1.SimulateTrade.Bytes(),ca2.SimulateTrade.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "SimulateTrade address isn't matching "
	}
	if !bytes.Equal(ca1.Trade.Bytes(),ca2.Trade.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "Trade address isn't matching "
	}
	if !bytes.Equal(ca1.OICash.Bytes(),ca2.OICash.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "OICash address isn't matching "
	}
	if !bytes.Equal(ca1.UniswapV2Factory.Bytes(),ca2.UniswapV2Factory.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "UniswapV2Factory address isn't matching "
	}
	if !bytes.Equal(ca1.UniswapV2Router02.Bytes(),ca2.UniswapV2Router02.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "UniswapV2Router02 address isn't matching "
	}
	if !bytes.Equal(ca1.AuditFunds.Bytes(),ca2.AuditFunds.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "AuditFunds address isn't matching "
	}
	if !bytes.Equal(ca1.WETH9.Bytes(),ca2.WETH9.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "WETH9 address isn't matching "
	}
	if !bytes.Equal(ca1.USDC.Bytes(),ca2.USDC.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "USDC address isn't matching "
	}
	if !bytes.Equal(ca1.USDT.Bytes(),ca2.USDT.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "USDT address isn't matching "
	}
	if !bytes.Equal(ca1.RelayHubV2.Bytes(),ca2.RelayHubV2.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "RelayHubV2 address isn't matching "
	}
	if !bytes.Equal(ca1.AccountLoader.Bytes(),ca2.AccountLoader.Bytes()) {
		num_mismatches++
		all_errors = all_errors + "Account loader address isn't matching "
	}
	return num_mismatches,errors.New(all_errors)
}
func Copy_iexchange_order_to_zeroxorder(in *IExchangeOrder) zeroex.Order {
	// copies the data between compatible types
	var out zeroex.Order

	out.MakerAddress.SetBytes(in.MakerAddress.Bytes())
	out.TakerAddress.SetBytes(in.TakerAddress.Bytes())
	out.FeeRecipientAddress.SetBytes(in.FeeRecipientAddress.Bytes())
	out.SenderAddress.SetBytes(in.SenderAddress.Bytes())
	out.MakerAssetAmount=new(big.Int)
	out.MakerAssetAmount.Set(in.MakerAssetAmount)
	out.TakerAssetAmount=new(big.Int)
	out.TakerAssetAmount.Set(in.TakerAssetAmount)
	out.MakerFee=new(big.Int)
	out.MakerFee.Set(in.MakerFee)
	out.ExpirationTimeSeconds=new(big.Int)
	out.ExpirationTimeSeconds.Set(in.ExpirationTimeSeconds)
	out.Salt=new(big.Int)
	out.Salt.Set(in.Salt)


	out.MakerAssetData = make([]byte,len(in.MakerAssetData))
	copy(out.MakerAssetData,in.MakerAssetData)

	out.TakerAssetData = make([]byte,len(in.TakerAssetData))
	copy(out.TakerAssetData,in.TakerAssetData)


	out.MakerFeeAssetData = make([]byte,len(in.MakerFeeAssetData))
	copy(out.MakerFeeAssetData,in.MakerFeeAssetData)

	out.TakerFeeAssetData = make([]byte,len(in.TakerFeeAssetData))
	copy(out.TakerFeeAssetData,in.TakerFeeAssetData)

	return out
}
func Copy_zerox_order_to_iexchange_order(in *zeroex.Order ) IExchangeOrder {
	// copies the data between compatible types
	var out IExchangeOrder

	out.MakerAddress.SetBytes(in.MakerAddress.Bytes())
	out.TakerAddress.SetBytes(in.TakerAddress.Bytes())
	out.FeeRecipientAddress.SetBytes(in.FeeRecipientAddress.Bytes())
	out.SenderAddress.SetBytes(in.SenderAddress.Bytes())
	out.MakerAssetAmount=new(big.Int)
	out.MakerAssetAmount.Set(in.MakerAssetAmount)
	out.TakerAssetAmount=new(big.Int)
	out.TakerAssetAmount.Set(in.TakerAssetAmount)
	out.MakerFee=new(big.Int)
	out.MakerFee.Set(in.MakerFee)
	out.ExpirationTimeSeconds=new(big.Int)
	out.ExpirationTimeSeconds.Set(in.ExpirationTimeSeconds)
	out.Salt=new(big.Int)
	out.Salt.Set(in.Salt)


	out.MakerAssetData = make([]byte,len(in.MakerAssetData))
	copy(out.MakerAssetData,in.MakerAssetData)

	out.TakerAssetData = make([]byte,len(in.TakerAssetData))
	copy(out.TakerAssetData,in.TakerAssetData)


	out.MakerFeeAssetData = make([]byte,len(in.MakerFeeAssetData))
	copy(out.MakerFeeAssetData,in.MakerFeeAssetData)

	out.TakerFeeAssetData = make([]byte,len(in.TakerFeeAssetData))
	copy(out.TakerFeeAssetData,in.TakerFeeAssetData)

	return out
}
func Unpack_sharetoken_id(packed_token_id []byte) (common.Address,uint8) {

	addr := common.BytesToAddress(packed_token_id[0:20])
	outcome_idx := packed_token_id[20]
	return addr,outcome_idx
}
func Fetch_erc20_info(client *ethclient.Client,contract_address *common.Address) (ERC20Info,error) {

	var erc20Info ERC20Info
	var copts = new(bind.CallOpts)

	contract,err := NewERC20Wrapper(*contract_address,client)
	if err != nil {
		return erc20Info,errors.New(fmt.Sprintf("NewERC20Wrapper error: %v",err))
	}

	total_supply,err := contract.TotalSupply(copts)
	if err != nil {
		total_supply = big.NewInt(0)
	}
	erc20Info.TotalSupply = total_supply.String()

	decimals,err := contract.Decimals(copts)
	if err != nil {
		erc20Info.Decimals = 18
	} else {
		erc20Info.Decimals = int(decimals)
	}

	symbol,err := contract.Symbol(copts)
	if err != nil {
		old_contract,err := NewOldERC20Token(*contract_address,client)
		if err != nil {
			return erc20Info,errors.New(fmt.Sprintf("OldERC20Token instantiation error: %v",err))
		}
		byte_symbol,err := old_contract.Symbol(copts)
		if err != nil {
			//return erc20Info,errors.New(fmt.Sprintf("Symbol() old version error: %v",err))
		} else {
			length := bytes.Index(byte_symbol[:], []byte{0})
			if length == -1 {
				length = 32
			}
			erc20Info.Symbol = string(byte_symbol[:length])
		}
	} else {
		erc20Info.Symbol = symbol
	}

	name,err := contract.Name(copts)
	if err != nil {
		old_contract,err := NewOldERC20Token(*contract_address,client)
		if err != nil {
			return erc20Info,errors.New(fmt.Sprintf("OldERC20Token instantiation error: %v",err))
		}
		byte_name,err := old_contract.Name(copts)
		if err != nil {
			//return erc20Info,errors.New(fmt.Sprintf("Name() old version error: %v",err))
		} else {
			length := bytes.Index(byte_name[:], []byte{0})
			if length == -1 {
				length = 32
			}
			erc20Info.Name = string(byte_name[:length])
		}
	} else {
		erc20Info.Name = name
	}

	return erc20Info,nil
}
