package randomwalk

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
)

// Topic-hash constants for every RandomWalk event the ETL dispatches;
// TestRegistryConstantsMatchABIEventIDs pins each constant to its
// ABI-derived event ID.
const (
	TopicNewOffer      = "55076e90b6b34a2569ffb2e1e34ee0da92d30ca423f0d6cfb317d252ade9a56a"
	TopicItemBought    = "caacc56f18ca259dc5175dae29eb0ca81407703a4819958c6885acbb7d4f3af3"
	TopicOfferCanceled = "0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951"
	TopicWithdrawalEvt = "a11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7"
	TopicTokenNameEvt  = "8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12" // #nosec G101 -- event signature hash, not a credential
	TopicMintEvent     = "ad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec"
	TopicTransferEvt   = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

// topicHash converts one of the topic-hash constants above to a common.Hash
// for handler registration.
func topicHash(hexConst string) ethcommon.Hash {
	return ethcommon.HexToHash("0x" + hexConst)
}
