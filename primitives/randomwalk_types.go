package primitives

import (
	"math/big"

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
type ERandomWalk_NewOffer struct {
	//signature: 0x8b4d06c200b17b9c1150172953ceb6fa3e7ace7623f6f933707badfa52c354cf
	OfferId *big.Int
	TokenId *big.Int
	Seller  common.Address
	Buyer   common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
type ERandomWalk_ItemBought struct {
	//signature: 0x7f7e375fbeaef0d6ebfc53af15b7aeed1d704e3424f34ef67e914b1f68f8c8ef
	OfferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
type ERandomWalk_OfferCanceled struct {
	//signature: 0x0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951
	OfferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
type ERandomWalk_WithdrawalEvent struct {
	//signature: 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7
	TokenId     *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}
type ERandomWalk_TokenNameEvent struct {
	//signature: 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12
	TokenId *big.Int
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}
type ERandomWalkTransfer struct {
	//signature: 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

type RW_NewOffer struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	OfferId					int64
	TokenId					int64
	Contract                string
	Buyer					string
	Seller					string
	Price					string
}
type RW_ItemBought struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	OfferId					int64
}
type RW_OfferCanceled struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	OfferId					int64
}
type RW_Withdrawal struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	TokenId					int64
	Destination				string
	Amount					string
}
type RW_TokenName struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	TokenId					int64
	NewName					string
}
type RW_Transfer struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	From					string
	To						string
	TokenId					int64
}
