package primitives
import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)
type Evt_ERC1155TransferBatch struct {
	// Signature: 4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type Evt_ERC1155TransferSingle struct {
	// Signature: c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
type Evt_ERC1155URI struct {
	// Signature: 6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

type ERC1155ProcStatus struct {
	LastEvtId			int64
}
type ERC1155_B struct {
	Id					int64
	Aid					int64
	ParentId			int64
	BlockNum			int64
	ContractAid			int64
	Processed			bool
	Address				string
	ContractAddr		string
	Amount				string
	Balance				string
	BlockHash			string
}
