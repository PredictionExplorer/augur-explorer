// Augur ETL: Converts Augur Data to SQL database

package main

import (
	"bytes"
	"sort"
	"fmt"
	"log"
	"os"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
)
var (
	all_contracts map[string]interface{}
	storage *SQLStorage
	Info    *log.Logger
	market_order_id int64 = 0
)
func check(e error) {
	if e != nil {
		panic(fmt.Sprintf("Exiting Augur extractor with error: %v",e))
	}
}
func insert_abi_events(contract_name string,a *abi.ABI) {

	fmt.Printf("Events:\n")
	for evt:=range a.Events {
		sig:=a.Events[evt].ID().Bytes()
		fmt.Printf("\teee::%v::%v::%v\n",contract_name,hex.EncodeToString(sig[:4]),evt)
		storage.Insert_event_signature(hex.EncodeToString(sig[:4]),evt,contract_name)
	}

}
func insert_abi_methods(contract_name string,a *abi.ABI) {
	fmt.Printf("Methods:\n")
	for meth := range a.Methods {
		fmt.Printf("\tmmm::%v\t%v::%v\n",contract_name,hex.EncodeToString(a.Methods[meth].ID()),meth)
		storage.Insert_function_signature(hex.EncodeToString(a.Methods[meth].ID()),meth,contract_name)
	}
}
func insert_all_artifacts() {

	names := make([]string, 0, len(all_contracts))
	for k := range all_contracts {
		names = append(names, k)
	}
	sort.Strings(names)
	for i:=0; i<len(names); i++ {
		contract_name := names[i]
		fmt.Printf("Contract: %v\n",contract_name)
		abi:=abi_from_artifacts(contract_name)
		insert_abi_events(contract_name,abi)
		insert_abi_methods(contract_name,abi)
		//dump_abi_events(contract_name,abi)
		//dump_abi_methods(contract_name,abi)
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

	if len(os.Args) < 3 {
		fmt.Printf("usage: %v [abi_file] [contract_list]\n",os.Args[0])
		os.Exit(1)
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)

	contract_names := os.Args[2]
	abi := load_abi(os.Args[1])
	insert_abi_events(contract_names,abi)
	insert_abi_methods(contract_names,abi)

}
