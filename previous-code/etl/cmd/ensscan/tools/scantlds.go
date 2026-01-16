// Scans top-TLD domains file generates the label hash 
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
			words := strings.Split(parsed[1],".")
			length := len(words)
			for i:=(length -2) ; i >=0 ; i-- {
				_,exists := names[words[i]]
				if exists {
					continue
				}
				hash,err := LabelHash(words[i])
				if err != nil {
					continue
				}
				hash_str := hex.EncodeToString(hash[:])
				fmt.Printf("%v\t%v\n",words[i],hash_str)
				names[words[i]]=struct{}{}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error : %v",err)
		os.Exit(1)
	}

}
