package primitives
import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
type EAugurFoundryWrapperCreated struct {
	TokenId      *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}
