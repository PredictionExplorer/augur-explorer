package main

import (
	"fmt"
	"math/big"
	"bytes"
	//"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)
func bigint_ptr_slice_to_str(data *[]*big.Int,separator string) string {
	var output bytes.Buffer
	length := len(*data)
	for i:=0 ; i< length ; i++  {
		if i>0 {
			output.WriteString(separator)
		}
		output.WriteString((*data)[i].String())
	}
	return output.String()
}
func outcomes_to_str(outcomes *[][32]byte,separator string) string {
	var output bytes.Buffer
	length := len(*outcomes)
	for i:=0 ; i<length ; i++ {
		if i>0 {
			output.WriteString(separator)
		}
		var zero_pos int = 0
		for ; zero_pos < 32 ; zero_pos++ {
			if (*outcomes)[i][zero_pos] == 0 {
				break
			}
		}
		s := string((*outcomes)[i][:zero_pos])
		fmt.Printf("outcome_idx %v : zero_pos = %v str=%v\n",i,zero_pos,s)
		//s := hex.EncodeToString((*outcomes)[i][:])
		output.WriteString(s)
	}
	return output.String()
}
func addresses_to_str(addresses *[]common.Address,separator string) string {
	var output bytes.Buffer
	length := len(*addresses)
	for i:=0 ; i<length ; i++ {
		if i>0 {
			output.WriteString(separator)
		}
		output.WriteString((*addresses)[i].String())
	}
	return output.String()
}
