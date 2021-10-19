package primitives

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

)
type RW_API_Offer struct {
	Id				int64
	EvtLogId		int64
	BlockNum		int64
	TxId			int64
	TimeStamp		string
	ContractAid		int64
	OfferId			int64
	SellerAid		int64
	SellerAddr		string
	BuyerAid		int64
	BuyerAddr		string
	TokenId			int64
	Active			bool
	Price			float64
}
