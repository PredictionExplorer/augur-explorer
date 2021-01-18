package main

import (
	"os"
	"fmt"
	"encoding/hex"

	 . "github.com/PredictionExplorer/augur-explorer/primitives"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("usage:	%v [ens_name]\n",os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	hash,err := LabelHash(name)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Label hash: %v\n",hex.EncodeToString(hash[:]))
}
