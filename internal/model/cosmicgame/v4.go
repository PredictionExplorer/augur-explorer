// V4 (in development) event models: the game proxy's art-auction wiring and
// ChaosZero cycle-funding events, plus the full event set of the
// non-upgradeable CosmicSignatureArtAuction contract. Big numbers travel as
// decimal strings, addresses as hex strings (the store layer resolves them
// to aids), timestamps that the contract emits as unix seconds as int64.
package cosmicgame

// CGArtAuctionConfiguredEvent records an ArtAuctionConfigured game event:
// the art-auction contract selected during initializeV4 or a later upgrade.
type CGArtAuctionConfiguredEvent struct {
	EvtId          int64
	BlockNum       int64
	TimeStamp      int64
	TxId           int64
	ContractAddr   string
	ArtAuctionAddr string
}

// CGArtAuctionOpeningBidChangedEvent records an ArtAuctionOpeningBidChanged
// game event: the owner updated the opening price for future batches.
type CGArtAuctionOpeningBidChangedEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	OpeningBid   string
}

// CGArtAuctionNftsMintedEvent records an ArtAuctionNftsMinted game event:
// the batch of NFTs a main-prize claim minted into art-auction custody.
type CGArtAuctionNftsMintedEvent struct {
	EvtId          int64
	BlockNum       int64
	TimeStamp      int64
	TxId           int64
	ContractAddr   string
	RoundNum       int64
	ArtAuctionAddr string
	FirstNftId     int64
	NumNfts        int64
}

// CGChaosZeroCycleFundedEvent records a ChaosZeroCycleFunded game event:
// CST minted and transferred to the ChaosZero market for one cycle.
type CGChaosZeroCycleFundedEvent struct {
	EvtId            int64
	BlockNum         int64
	TimeStamp        int64
	TxId             int64
	ContractAddr     string
	MarketAddr       string
	Cycle            int64
	Amount           string
	CumulativeAmount string
}

// CGArtAuctionBatchRegisteredEvent records an AuctionBatchRegistered
// art-auction event: fixed terms for one round's batch of eight NFTs.
type CGArtAuctionBatchRegisteredEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	RoundNum     int64
	FirstNftId   int64
	NumNfts      int64
	EligibleAt   int64 // unix seconds; bidding opens at this time
	OpeningBid   string
}

// CGArtAuctionBidEvent records an AuctionBidPlaced art-auction event.
// Amount is the bidder's total standing bid, EthPaid the ETH sent with this
// transaction (lower when the leader raises their own bid), EndAt the
// auction deadline after any anti-sniping extension.
type CGArtAuctionBidEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	NftId        int64
	RoundNum     int64
	BidderAddr   string
	Amount       string
	EthPaid      string
	EndAt        int64 // unix seconds
}

// CGArtAuctionRefundCreditedEvent records an AuctionRefundCredited
// art-auction event: an outbid bidder's funds became a pull-based credit.
type CGArtAuctionRefundCreditedEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	BidderAddr   string
	NftId        int64
	Amount       string
}

// CGArtAuctionRefundWithdrawnEvent records an AuctionRefundWithdrawn
// art-auction event.
type CGArtAuctionRefundWithdrawnEvent struct {
	EvtId         int64
	BlockNum      int64
	TimeStamp     int64
	TxId          int64
	ContractAddr  string
	BidderAddr    string
	RecipientAddr string
	Amount        string
}

// CGArtAuctionSettledEvent records an AuctionSettled art-auction event.
type CGArtAuctionSettledEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	NftId        int64
	RoundNum     int64
	WinnerAddr   string
	Amount       string
}

// CGArtAuctionNftDeliveredEvent records an AuctionNftDelivered art-auction
// event: the winner collected the NFT to an address of their choice.
type CGArtAuctionNftDeliveredEvent struct {
	EvtId         int64
	BlockNum      int64
	TimeStamp     int64
	TxId          int64
	ContractAddr  string
	NftId         int64
	WinnerAddr    string
	RecipientAddr string
}

// CGArtAuctionProceedsForwardedEvent records an AuctionProceedsForwarded
// art-auction event: a settled lot's final price donated into the game.
type CGArtAuctionProceedsForwardedEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	NftId        int64
	RoundNum     int64
	Amount       string
}

// CGArtAuctionSurplusForwardedEvent records an AuctionSurplusForwarded
// art-auction event: unexpected ETH swept into the game as a donation.
type CGArtAuctionSurplusForwardedEvent struct {
	EvtId        int64
	BlockNum     int64
	TimeStamp    int64
	TxId         int64
	ContractAddr string
	Amount       string
}
