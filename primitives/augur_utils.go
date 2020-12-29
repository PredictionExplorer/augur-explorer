package primitives

import (
	"math/big"
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)
func Eth_addr_is_zero(addr_ptr *common.Address) bool {
	if bytes.Equal(addr_ptr.Bytes(), common.Address{}.Bytes()) {
		return true
	}
	return false
}
func Bigint_ptr_slice_to_str(data *[]*big.Int,separator string) string {
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
func Outcomes_to_str(outcomes *[][32]byte,separator string) string {
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
func Fetch_erc20_info(client *ethclient.Client,contract_address *common.Address) (ERC20Info,error) {

	var erc20Info ERC20Info
	var copts = new(bind.CallOpts)

	contract,err := NewERC20Wrapper(*contract_address,client)
	if err != nil {
		return erc20Info,errors.New(fmt.Sprintf("NewERC20Wrapper error: %v",err))
	}

	total_supply,err := contract.TotalSupply(copts)
	if err != nil {
		total_supply = big.NewInt(0)
	}
	erc20Info.TotalSupply = total_supply.String()

	decimals,err := contract.Decimals(copts)
	if err != nil {
		erc20Info.Decimals = 18
	} else {
		erc20Info.Decimals = int(decimals)
	}

	symbol,err := contract.Symbol(copts)
	if err != nil {
		old_contract,err := NewOldERC20Token(*contract_address,client)
		if err != nil {
			return erc20Info,errors.New(fmt.Sprintf("OldERC20Token instantiation error: %v",err))
		}
		byte_symbol,err := old_contract.Symbol(copts)
		if err != nil {
			//return erc20Info,errors.New(fmt.Sprintf("Symbol() old version error: %v",err))
		} else {
			length := bytes.Index(byte_symbol[:], []byte{0})
			if length == -1 {
				length = 32
			}
			erc20Info.Symbol = string(byte_symbol[:length])
		}
	} else {
		erc20Info.Symbol = symbol
	}

	name,err := contract.Name(copts)
	if err != nil {
		old_contract,err := NewOldERC20Token(*contract_address,client)
		if err != nil {
			return erc20Info,errors.New(fmt.Sprintf("OldERC20Token instantiation error: %v",err))
		}
		byte_name,err := old_contract.Name(copts)
		if err != nil {
			//return erc20Info,errors.New(fmt.Sprintf("Name() old version error: %v",err))
		} else {
			length := bytes.Index(byte_name[:], []byte{0})
			if length == -1 {
				length = 32
			}
			erc20Info.Name = string(byte_name[:length])
		}
	} else {
		erc20Info.Name = name
	}

	return erc20Info,nil
}
func Bytes32_to_string(data []byte) string {

	length := bytes.Index(data, []byte{0})
	if length == -1 {
		length = 32
	}
	return string(data[:length])
}
