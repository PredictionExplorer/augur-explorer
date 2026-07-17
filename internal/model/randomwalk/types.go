// Package randomwalk defines the RandomWalk NFT data model: the event and
// row types shared by the ETL, storage, API and notification layers.
package randomwalk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// =====================================================================
// Processing Status
// =====================================================================

// ProcStatus is the rw-etl progress watermark (rw_proc_status row): the last
// processed event-log ID and block number.
type ProcStatus struct {
	LastIdProcessed int64
	LastBlockNum    int64
}

// =====================================================================
// Contract Configuration
// =====================================================================

// ContractAddresses carries the deployed RandomWalk NFT and marketplace
// contract addresses with their resolved address IDs.
type ContractAddresses struct {
	MarketPlace    string
	RandomWalk     string
	MarketPlaceAid int64
	RandomWalkAid  int64
}

// =====================================================================
// Ethereum Event Types (parsed from RLP)
// =====================================================================

// ENewOffer is the decoded marketplace NewOffer contract event.
type ENewOffer struct {
	// signature: 55076e90b6b34a2569ffb2e1e34ee0da92d30ca423f0d6cfb317d252ade9a56a
	NftAddress common.Address
	OfferId    *big.Int
	TokenId    *big.Int
	Seller     common.Address
	Buyer      common.Address
	Price      *big.Int
	Raw        types.Log
}

// EItemBought is the decoded marketplace ItemBought contract event.
type EItemBought struct {
	// signature: 0xcaacc56f18ca259dc5175dae29eb0ca81407703a4819958c6885acbb7d4f3af3
	OfferId *big.Int
	Seller  common.Address
	Buyer   common.Address
	Raw     types.Log
}

// EOfferCanceled is the decoded marketplace OfferCanceled contract event.
type EOfferCanceled struct {
	// signature: 0x0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951
	OfferId *big.Int
	Raw     types.Log
}

// EWithdrawalEvent is the decoded RandomWalk WithdrawalEvent contract event
// (half of the accumulated mint funds paid to the last minter).
type EWithdrawalEvent struct {
	// signature: 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7
	TokenId     *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log
}

// ETokenNameEvent is the decoded RandomWalk TokenNameEvent contract event
// (a token rename).
type ETokenNameEvent struct {
	// signature: 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12
	TokenId *big.Int
	NewName string
	Raw     types.Log
}

// EMintEvent is the decoded RandomWalk MintEvent contract event with the
// token's generation seed.
type EMintEvent struct {
	// signature: 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec
	TokenId *big.Int
	Owner   common.Address
	Seed    [32]byte
	Price   *big.Int
	Raw     types.Log
}

// ETransfer is the decoded ERC-721 Transfer event of the RandomWalk
// contract.
type ETransfer struct {
	// signature: 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log
}

// =====================================================================
// Database Record Types
// =====================================================================

// NewOffer is the rw_new_offer row: one marketplace offer creation.
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

// ItemBought is the rw_item_bought row: one marketplace sale settling an
// offer.
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

// OfferCanceled is the rw_offer_canceled row: one offer cancellation.
type OfferCanceled struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	OfferId   int64
}

// Withdrawal is the rw_withdrawal row: one mint-funds withdrawal by the
// last minter.
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

// TokenName is the rw_token_name row: one token rename event.
type TokenName struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	TimeStamp int64
	Contract  string
	TokenId   int64
	NewName   string
}

// Transfer is the rw_transfer row: one ERC-721 transfer of a RandomWalk
// token (mints and burns use the zero address).
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

// MintEvent is the rw_mint row: one RandomWalk token mint with its seed and
// price.
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

// NotificationEvent is one queued bot notification about a mint or
// marketplace event, discriminated by EvtType.
type NotificationEvent struct {
	TokenId         int64
	TimeStampMinted int64
	Price           float64
	SeedHex         string
	EvtType         int64 // 0-undefined,1-Mint,2-NewOffer Sell, 3-ItemBought, 4 -FloorPriceChanged, 5 - NewOffer Buy
}

// NotificationEvent2 is the extended notification record carrying the
// event-log ID used by the notibot watermark.
type NotificationEvent2 struct {
	TokenId         int64
	TxId            int64 // currently unused, but considered for future use
	EvtLogId        int64
	TimeStampMinted int64
	Price           float64
	SeedHex         string
	EvtType         int64 // 0-undefined,1-Mint,2-NewOffer Sell, 3-ItemBought, 4 -FloorPriceChanged, 5 - NewOffer Buy
}

// TransferEntry is a minimal transfer notification (from, to, token).
type TransferEntry struct {
	From    string
	To      string
	TokenId int64
}

// MsgStatus is the notibot messaging watermark (rw_messaging_status row):
// the last event already announced to social channels.
type MsgStatus struct {
	TxId      int64
	EvtLogId  int64
	BlockNum  int64
	TimeStamp int64
}
