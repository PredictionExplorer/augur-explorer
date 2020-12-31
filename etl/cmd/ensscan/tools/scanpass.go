// Scans 34 million password database and generates the label hash for each token (word)
package main

import (
	"os"
	"fmt"
	"encoding/hex"
	"bufio"
	"strings"

	 . "github.com/PredictionExplorer/augur-explorer/primitives"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		password := strings.ReplaceAll(scanner.Text(),`'`,`''`)
		hash,err := LabelHash(password)
		if err != nil {
			continue
		}
		hash_str := hex.EncodeToString(hash[:])
		fmt.Printf("%v\t%v\n",password,hash_str)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error : %v",err)
		os.Exit(1)
	}

}
