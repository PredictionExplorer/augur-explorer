package primitives
import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
type NameRegistered_v1 struct { //0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
	Name    string
	Label   [32]byte
	Owner   common.Address
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
type NameRegistered_v3 struct { //570313dae523ecb48b1176a4b60272e5ea7ec637f5b2d09983cbc4bf25e7e9e3
	Caller		common.Address
	Beneficiary common.Address
	Label		[32]byte
	Subdomain	string
	CreatedDate	*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
type HashInvalidated struct { //1f9c649fe47e58bb60f4e52f0d90e4c47a526c9f90c5113df842c025970b66ad
	Hash					[32]byte
	Name					string
	Value					*big.Int
	RegistrationDate		*big.Int
	Raw						types.Log // Blockchain specific contextual infos
}
type HashRegistered struct { //0f0c27adfd84b60b6f456b0e87cdccb1e5fb9603991588d87fa99f5b6b61e670
	Hash					[32]byte
	Owner					common.Address
	Value					*big.Int
	RegistrationDate		*big.Int
	Raw						types.Log // Blockchain specific contextual infos
}
type AddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

