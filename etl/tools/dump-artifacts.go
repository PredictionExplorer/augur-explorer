// Augur ETL: Converts Augur Data to SQL database

package main

import (
	"bytes"
	"sort"
	"fmt"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
)
const (
)
var (
	all_contracts map[string]interface{}
)
func check(e error) {
	if e != nil {
		panic(fmt.Sprintf("Exiting Augur extractor with error: %v",e))
	}
}
func dump_abi_events(contract_name string,a *abi.ABI) {

	fmt.Printf("Events:\n")
	for evt:=range a.Events {
		fmt.Printf("\teee::%v::%v::%v\n",contract_name,a.Events[evt].ID().String(),evt)
	}

}
func dump_abi_methods(contract_name string,a *abi.ABI) {
	fmt.Printf("Methods:\n")
	for meth := range a.Methods {
		fmt.Printf("\tmmm::%v\t%v::%v\n",contract_name,hex.EncodeToString(a.Methods[meth].ID()),meth)
	}
}
func dump_all_artifacts() {

	names := make([]string, 0, len(all_contracts))
	for k := range all_contracts {
		names = append(names, k)
	}
	sort.Strings(names)
	for i:=0; i<len(names); i++ {
		contract_name := names[i]
		fmt.Printf("Contract: %v\n",contract_name)
		abi:=abi_from_artifacts(contract_name)
		dump_abi_events(contract_name,abi)
		dump_abi_methods(contract_name,abi)
	}
}
func load_all_artifacts(filename string) map[string]interface{} {

	abi_data, err := ioutil.ReadFile(filename)
	check(err)
	all_abis_rdr := bytes.NewReader(abi_data)
	check(err)
	byte_data, err := ioutil.ReadAll(all_abis_rdr)
	check(err)
	var all_contracts map[string]interface{}
	json.Unmarshal([]byte(byte_data), &all_contracts)
	return all_contracts
}
func abi_from_artifacts(contract string) *abi.ABI {

	contract_abi:=all_contracts[contract]
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
func main() {

//	all_contracts = load_all_artifacts("./abis/augur-artifacts-abi.json")
	all_contracts = load_all_artifacts("./abis/augur-artifacts-abi-5jun.json")
	dump_all_artifacts()

}
