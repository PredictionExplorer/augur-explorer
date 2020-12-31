// Scans 34 million password database and generates the label hash for each token (word)
package main

import (
	"os"
	"fmt"
	"encoding/hex"
	"bufio"
	"strings"
	"unicode/utf8"

	 . "github.com/PredictionExplorer/augur-explorer/primitives"
)

func process_token(token string) bool {
	// Note: passwords are all unique, so we don't have unique index check here
	if len(token)==0 {
		return false
	}
	token = strings.ReplaceAll(token,` `,``)
	token = strings.ReplaceAll(token,`'`,``)
	token = strings.ReplaceAll(token,"\t",``)
	token = strings.ReplaceAll(token,"\r",``)
	token = strings.ReplaceAll(token,`\`,``)
	if len(token)==0 {
		return false
	}
	if !utf8.ValidString(token) {
		return false
	}
	hash,err := LabelHash(token)
	if err != nil {
		return false
	}
	hash_str := hex.EncodeToString(hash[:])
	fmt.Printf("%v\t%v\n",token,hash_str)
	return true
}
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		process_token(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error : %v",err)
		os.Exit(1)
	}

}
