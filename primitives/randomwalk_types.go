package primitives

import (

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

)
type RandomWalkProcStatus struct {
	LastIdProcessed			int64
	LastBlockNum			int64
}
type RW_ContractAddresses struct {
	MarketPlace				string
	RandomWalk				string
	MarketPlaceAid			int64
	RandomWalkAid			int64
}
