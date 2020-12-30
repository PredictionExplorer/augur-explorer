// Scans 61k proper name dictionary for English
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

	names:=make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := strings.Split(line,",")
		if len(parsed) > 1 {
			word := parsed[1]
			hash,err := LabelHash(word)
			if err != nil {
				continue
			}
			hash_str := hex.EncodeToString(hash[:])
			_,exists := names[hash_str]
			if exists {
				continue
			}
			fmt.Printf("%v\t%v\n",word,hash_str)
			names[hash_str]=struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error : %v",err)
		os.Exit(1)
	}

}
