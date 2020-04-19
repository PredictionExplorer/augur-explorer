///
/// Dumps hashes of the signatures of Contract's events using provided ABI file (on the commandline)

package main
import (
	"os"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"fmt"
	"bytes"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func dump_event_signatures(file_name string) {

	//abi_data, err := ioutil.ReadFile("/home/niko/eth/abis/Augur.abi")
	abi_data, err := ioutil.ReadFile(file_name)
	check(err)
	abi_rdr := bytes.NewReader(abi_data)
	contract_abi,err:=abi.JSON(abi_rdr)

	for evt:=range contract_abi.Events {
		fmt.Println(contract_abi.Events[evt].ID().String(),"\t",evt)
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %v [contracts_abi_filename.abi]\n",os.Args[0])
		os.Exit(1)
	}
	dump_event_signatures(os.Args[1])
}
