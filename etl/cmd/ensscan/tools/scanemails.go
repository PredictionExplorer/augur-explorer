// Scans 300million email addresses file and generates the label hash for each token (word)
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
var (
	some_names map[string]struct{} = make(map[string]struct{})
)
func process_token(token string) bool {

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
	if len(token) <= 10 {
		_,exists := some_names[token]
		if exists {
			return true
		}
	}
	if !utf8.ValidString(token) {
		return false
	}
	hash,err := LabelHash(token)
	if err != nil {
		return false
	}
	if len(token) <= 10 {
		some_names[token] = struct{}{}
	}
	hash_str := hex.EncodeToString(hash[:])
	fmt.Printf("%v\t%v\n",token,hash_str)
	return true
}
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf,32*1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		idx := strings.Index(line, "@")
		if idx < 0 {
			continue
		}
		var bad_line bool = false
		parsed := strings.Split(line,"@")
		for j:=0 ; j < len(parsed) ; j ++ {
			email_addr := parsed[j]
			tokens := strings.Split(email_addr,".")
			for k:=0 ; k<len(tokens) ; k++ {
				tok := tokens[k]
				result := process_token(tok)
				if result == false {
					bad_line = true
				}
			}
		}
		if bad_line {
			os.Stderr.WriteString(line+"\n")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error : %v",err)
		os.Exit(1)
	}
}
