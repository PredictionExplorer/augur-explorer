// Scans 300million email addresses file and generates the label hash for each token (word)
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
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		idx := strings.Index(line, "@")
		if idx < 0 {
			continue
		}
		parsed := strings.Split(line,"@")
		if len(parsed) > 1 {
			username := parsed[0]
			domain_name := parsed[1]
			words_user := strings.Split(username,".")
			length := len(words_user)
			for i:=0 ; i<length ; i++ {
				hash,err := LabelHash(words_user[i])
				if err != nil {
					continue
				}
				hash_str := hex.EncodeToString(hash[:])
				_,exists := names[hash_str]
				if exists {
					continue
				}
				fmt.Printf("%v\t%v\n",words_user[i],hash_str)
				names[hash_str]=struct{}{}
			}
			words_domain := strings.Split(domain_name,".")
			length = len(words_domain)
			for i:=0 ; i<length ; i++ {
				hash,err := LabelHash(words_domain[i])
				if err != nil {
					continue
				}
				hash_str := hex.EncodeToString(hash[:])
				_,exists := names[hash_str]
				if exists {
					continue
				}
				fmt.Printf("%v\t%v\n",words_domain[i],hash_str)
				names[hash_str]=struct{}{}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error : %v",err)
		os.Exit(1)
	}

}
