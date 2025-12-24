// Package randomwalk contains types specific to the RandomWalk NFT project
package randomwalk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// =====================================================================
// Processing Status
// =====================================================================

type ProcStatus struct {
	LastIdProcessed int64
	LastBlockNum    int64
}

// =====================================================================
// Contract Configuration
// =====================================================================

type ContractAddresses struct {
	MarketPlace    string
	RandomWalk     string
	MarketPlaceAid int64
	RandomWalkAid  int64
}

// =====================================================================
// Ethereum Event Types (parsed from RLP)
// =====================================================================

type ENewOffer struct {
	//signature: 55076e90b6b34a2569ffb2e1e34ee0da92d30ca423f0d6cfb317d252ade9a56a
	NftAddress common.Address
	OfferId    *big.Int
	TokenId    *big.Int
	Seller     common.Address
	Buyer      common.Address
	Price      *big.Int
	Raw        types.Log
}

type EItemBought struct {
	//signature: 0xcaacc56f18ca259dc5175dae29eb0ca81407703a4819958c6885acbb7d4f3af3
	OfferId *big.Int
	Seller  common.Address
	Buyer   common.Address
	Raw     types.Log
}

type EOfferCanceled struct {
	//signature: 0x0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951
	OfferId *big.Int
	Raw     types.Log
}

type EWithdrawalEvent struct {
	//signature: 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7
	TokenId     *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log
}

type ETokenNameEvent struct {
	//signature: 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12
	TokenId *big.Int
	NewName string
	Raw     types.Log
}

type EMintEvent struct {
	//signature: 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec
	TokenId *big.Int
	Owner   common.Address
	Seed    [32]byte
	Price   *big.Int
	Raw     types.Log
}

type ETransfer struct {
	//signature: 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log
}

// =====================================================================
// Database Record Types
// =====================================================================

type NewOffer struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	OfferId   int64
	TokenId   int64
	Contract  string
	RWalkAddr string
	Buyer     string
	Seller    string
	Price     string
}

type ItemBought struct {
	EvtId      int64
	BlockNum   int64
	TxId       int64
	TimeStamp  int64
	Contract   string
	OfferId    int64
	SellerAddr string
	BuyerAddr  string
}

type OfferCanceled struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	OfferId   int64
}

type Withdrawal struct {
	EvtId       int64
	BlockNum    int64
	TxId        int64
	TimeStamp   int64
	Contract    string
	TokenId     int64
	Destination string
	Amount      string
}

type TokenName struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	TokenId   int64
	NewName   string
}

type Transfer struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	From      string
	To        string
	TokenId   int64
}

type MintEvent struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	TokenId   int64
	Owner     string
	Seed      string
	SeedNum   string
	Price     string
}

// =====================================================================
// Notification Types
// =====================================================================

type NotificationEvent struct {
	TokenId         int64
	TimeStampMinted int64
	Price           float64
	SeedHex         string
	EvtType         int64 //0-undefined,1-Mint,2-NewOffer Sell, 3-ItemBought, 4 -FloorPriceChanged, 5 - NewOffer Buy
}

type NotificationEvent2 struct {
	TokenId         int64
	TxId            int64 // currently unused, but considered for future use
	EvtLogId        int64
	TimeStampMinted int64
	Price           float64
	SeedHex         string
	EvtType         int64 //0-undefined,1-Mint,2-NewOffer Sell, 3-ItemBought, 4 -FloorPriceChanged, 5 - NewOffer Buy
}

type TransferEntry struct {
	From    string
	To      string
	TokenId int64
}

type MsgStatus struct {
	TxId      int64
	EvtLogId  int64
	BlockNum  int64
	TimeStamp int64
}

