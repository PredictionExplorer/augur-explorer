package primitives

import (
	"math"
	"math/big"
	"bytes"

	"github.com/ethereum/go-ethereum/common"
)
const (
	DEFAULT_LOG_DIR	 = "ae_logs"
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
func Bytes32_to_string(data []byte) string {

	length := bytes.Index(data, []byte{0})
	if length == -1 {
		length = 32
	}
	return string(data[:length])
}
func Augur_UI_price_adjustments(price *float64,amount *float64,mkt_type int,decimals int) {

	// Price and amount are fixed floating points of 18 precision
	// According to specs, the price of the outcom can range between 0 to [num_ticks]
	// however Augur multiplies quanty and divides the price to allow 0..1 price ranges
	multiplier := math.Pow(10,float64(decimals))
	if mkt_type == MktTypeScalar {
		if price != nil {
			//*price = *price / float64(CATEGORICAL_MULTIPLIER)
			*price = *price / multiplier
		}
		if amount != nil {
			//*amount = *amount * float64(CATEGORICAL_MULTIPLIER)
			*amount = *amount * multiplier
		}
	} else {
		if price != nil {
			*price = *price / float64(CATEGORICAL_MULTIPLIER)
		}
		if amount != nil {
			*amount = *amount * float64(CATEGORICAL_MULTIPLIER)
		}
	}
}
func (obj *GasSpent) Has_rows() bool {
	if (obj.Num_trading==0) && (obj.Num_reporting==0) && (obj.Num_markets==0) && (obj.Num_total==0) {
		return false
	}
	return true
}
