package primitives
import (
	"github.com/ethereum/go-ethereum/common"
)
type Deleter_status struct {
	LastBlockNum		int64
	ContractAid			int64
	ContractAddr		string
	Info				string
	ContractEthAddr		common.Address
}
