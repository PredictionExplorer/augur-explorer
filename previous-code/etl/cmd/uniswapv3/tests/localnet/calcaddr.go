package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [sender_addr] [nonce]\n\t\t"+
			"calculates contract address using sender's addr and nonce\")\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	from := common.HexToAddress(os.Args[1])
	nonce,err := strconv.ParseInt(os.Args[2],10,64)
	if err != nil {
		fmt.Printf("Error parsing nonce field: %v\n",err)
		os.Exit(1)
	}
	sender := vm.AccountRef(from)
	db := state.NewDatabase(rawdb.NewMemoryDatabase())
	state, _ := state.New(common.Hash{}, db, nil)
	state.SetNonce(from,uint64(nonce))
	fmt.Printf("nonce after set = %v\n",state.GetNonce(sender.Address()))
	contract_addr := crypto.CreateAddress(sender.Address(), state.GetNonce(sender.Address()))
	fmt.Printf("contract addr = %v\n",contract_addr.String())
}
