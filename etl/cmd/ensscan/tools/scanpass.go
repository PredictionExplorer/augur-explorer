// Scans 34 million password database and generates the label hash for each token (word)
package main

import (
	"os"
	"fmt"
	"encoding/hex"
	"bufio"

	 . "github.com/PredictionExplorer/augur-explorer/primitives"
)

func main() {

	names:=make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		password := scanner.Text()
		hash,err := LabelHash(password)
		if err != nil {
			continue
		}
		hash_str := hex.EncodeToString(hash[:])
		_,exists := names[hash_str]
		if exists {
			continue
		}
		fmt.Printf("%v\t%v\n",password,hash_str)
		names[hash_str]=struct{}{}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error : %v",err)
		os.Exit(1)
	}

}
