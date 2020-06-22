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
