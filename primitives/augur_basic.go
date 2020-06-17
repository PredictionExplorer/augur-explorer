package primitives

import (
	//"os"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/hex"
	"encoding/json"

	//"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
)
/* DISCONTINUED , to be delted
func dev_init_daicash(addr *common.Address) {
	if addr != nil {
		*addr = common.HexToAddress("5f3341EA5989aD3129E325027b8d908b63709A00")
	}
}
func prod_init_daicash(addr *common.Address) {
	if addr != nil {
		*addr = common.HexToAddress("6B175474E89094C44Da98b954EedeAC495271d0F")
	}
}
func dev_init_reputation_token(addr *common.Address) {
	if addr != nil {
		*addr = common.HexToAddress("0fF6ee01f88145298761a29A0372Ed24E16E73B1")
		//*addr = common.HexToAddress("B4D7f6747CEFbDcA11bDDd92a16134dc95B0DD9B") LegacyReputationToken
		//*addr = common.HexToAddress("B78B2B637d3861E601E54C00c054972c18A5e991")
	}
}
func prod_init_reputation_token(addr *common.Address) {
	if addr != nil {
		*addr = common.HexToAddress("B4D7f6747CEFbDcA11bDDd92a16134dc95B0DD9B")
	}
}
func dev_init_zerox(addr *common.Address) {
	if addr != nil {
		*addr = common.HexToAddress("6749E370e7B1955FFa924F4f75f5F12653C7512C")	// ZeroXTrade
	}
}
func prod_init_zerox(addr *common.Address) {
	if addr != nil {
		*addr = common.HexToAddress("6749E370e7B1955FFa924F4f75f5F12653C7512C")
	}
}
func dev_init_addresses(config_addr *ContractAddresses) {
	dev_init_zerox(config_addr.Zerox_addr)
	dev_init_daicash(config_addr.Dai_addr)
	dev_init_reputation_token(config_addr.Reputation_addr)
}
func prod_init_addresses(config_addr *ContractAddresses) {
	prod_init_zerox(config_addr.Zerox_addr)
	prod_init_daicash(config_addr.Dai_addr)
	prod_init_reputation_token(config_addr.Reputation_addr)
}
func Init_contract_addresses(addresses *ContractAddresses) {

	augur_prod := os.Getenv("AUGUR_PROD")
	if len(augur_prod) > 0 {
		prod_init_addresses(addresses)
	} else {
		dev_init_addresses(addresses)
	}
}
*/
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
