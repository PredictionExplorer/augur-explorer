package primitives
import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
// We have them both, Uniswap and Balancer
type ELOG_NEW_POOL struct { //0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}
type ELOG_JOIN struct {	//0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a
	Caller        common.Address
	TokenIn       common.Address
	TokenAmountIn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}
type ELOG_EXIT struct {	//0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed
	Caller         common.Address
	TokenOut       common.Address
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type ELOG_SWAP struct { // 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378
	Caller         common.Address
	TokenIn        common.Address
	TokenOut       common.Address
	TokenAmountIn  *big.Int
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}
type UPairCreated struct {
	Token0			common.Address
	Token1			common.Address
	Pair			common.Address
	PairSeq			*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}
type UPairSwap struct {
	Decimals0				int
	Decimals1				int
	Sender					common.Address
	Amount0In				*big.Int
	Amount1In				*big.Int
	Amount0Out				*big.Int
	Amount1Out				*big.Int
	To						common.Address
	Raw						types.Log // Blockchain specific contextual infos
}
