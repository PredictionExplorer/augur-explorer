package primitives

import (
	//"os"
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
)
const (
	DEFAULT_LOG_DIR	 = "ae_logs"
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

	abi_data, err := ioutil.ReadFile("./abis/augur-artifacts-abi.json")
	check(err)
	all_abis_rdr := bytes.NewReader(abi_data)
	check(err)
	byte_data, err := ioutil.ReadAll(all_abis_rdr)
	check(err)
	var contracts map[string]interface{}
	json.Unmarshal([]byte(byte_data), &contracts)
	return contracts
}
func Abi_from_artifacts(contracts *map[string]interface{},contract string) *abi.ABI {

	contract_abi:=(*contracts)[contract]
	contract_bytes, _ := json.Marshal(contract_abi) // convert back to JSON so Ethereum package can work
	rdr := bytes.NewReader(contract_bytes)
	ctrct_abi,err := abi.JSON(rdr)
	check(err)
	return &ctrct_abi
}
func load_abi(fname string) *abi.ABI {

	abi_data, err := ioutil.ReadFile(fname)
	check(err)
	abi_rdr := bytes.NewReader(abi_data)
	check(err)
	abi,err := abi.JSON(abi_rdr)
	check(err)
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
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("ZeroXTrade"))
	caddrs.ZeroxTrade,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("ZeroXExchange"))
	caddrs.ZeroxXchg,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	caddrs.Universe,err = ctrct_augur.GenesisUniverse(copts)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Getting GenesisUniverse failed: %v",err.Error()))
		return caddrs,newerr
	}

	var ctrct_universe *Universe
	ctrct_universe,err = NewUniverse(caddrs.Universe,eclient)
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
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("AugurWalletRegistryV2"))
	caddrs.WalletReg2,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	key = common.BigToHash(big.NewInt(0))
	copy(key[:],[]byte("FillOrder"))
	caddrs.FillOrder,err = ctrct_augurtrading.Lookup(copts,key)
	if err != nil {
		newerr:=errors.New(fmt.Sprintf("Lookup of %v failed",string(key[:]),err.Error()))
		return caddrs,newerr
	}

	return caddrs,nil
}
