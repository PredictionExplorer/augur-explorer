package chaoszero

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
)

// Topic-hash constants for every ChaosZero event the ETL dispatches;
// TestRegistryConstantsMatchABIEventIDs pins each constant to its
// ABI-derived event ID.
const (
	// IndependentOutcomeMarket.
	TopicRoundInitialized = "e2d996ef28ba133e7c64e9c5b2304b7c17e6d2887389032940016a63d882d971"
	TopicRoundFunded      = "81f6a77353d45ea7a238715966269be47f3904e088afd1ff67420e901a610a5c"
	TopicTrade            = "0543184a8a4fc3b557461bf0aaea980418c71c89c983e44e6422d4975300498b"
	TopicPositionSplit    = "f35af27f5c587462d5a0f029fb4c578bc53bc3c87db59b8db657010446b46455"
	TopicPositionsMerged  = "81ac7d05a625b7e14ab408012a5fbd6cbf06f0e9d28d2d5413d078e323a5dc2c"
	TopicRoundSettled     = "7bac5aa6107fc77cc2426c3c4f1491806f51281ee9dad946b5d3ad5dd6e7b469"
	TopicClaimed          = "d166975969e02ab94330b286dd849bfa3771eced0fed813cbf4492fd17e73790"
	TopicAnswersObserved  = "6eed9ea57baa5a42a4d4294006a1f18100e35851110898d45bf710aaaddbea3b"
	// OutcomeSettlementVault.
	TopicCycleEscrowed = "16cbc14eaf1c725d6bf8f56d127952bd60a5a107cc8f428d97ced68b4f5d589a"
	TopicPaid          = "9f777629478be3b0998eab401d2606e8b3ac41e303df6a1d13f5f1cdf48ce29f"
	// OpenZeppelin ERC-1155 events the market inherits.
	TopicTransferSingle = "c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	TopicTransferBatch  = "4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
	TopicApprovalForAll = "17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
)

// topicHash converts one of the topic-hash constants above to a common.Hash
// for handler registration.
func topicHash(hexConst string) ethcommon.Hash {
	return ethcommon.HexToHash("0x" + hexConst)
}
