// Package chaoszero defines the ChaosZero data model: the event and row
// types shared by the ETL and storage layers for the eight-market
// IndependentOutcomeMarket and its OutcomeSettlementVault.
package chaoszero

import (
	"math/big"
)

// =====================================================================
// Processing Status
// =====================================================================

// ProcStatus is the cz-etl progress watermark (cz_proc_status row): the last
// processed event-log ID and block number.
type ProcStatus struct {
	LastIdProcessed int64
	LastBlockNum    int64
}

// =====================================================================
// Contract Configuration
// =====================================================================

// ContractAddresses carries the deployed market, outcome source and
// settlement vault addresses with their resolved address IDs. The source
// emits no events; it is registered for address joins only.
type ContractAddresses struct {
	Market    string
	Source    string
	Vault     string
	MarketAid int64
	SourceAid int64
	VaultAid  int64
}

// =====================================================================
// Ethereum Event Types (ABI-unpacked non-indexed fields)
// =====================================================================

// ERoundFunded is the decoded non-indexed payload of RoundFunded.
type ERoundFunded struct {
	// signature: 81f6a77353d45ea7a238715966269be47f3904e088afd1ff67420e901a610a5c
	Donated       *big.Int
	TransfersUsed *big.Int
	CarryUsed     *big.Int
	Allocations   [8]*big.Int
}

// ETrade is the decoded non-indexed payload of Trade.
type ETrade struct {
	// signature: 0543184a8a4fc3b557461bf0aaea980418c71c89c983e44e6422d4975300498b
	Yes   bool
	Buy   bool
	Cash  *big.Int
	Units *big.Int
}

// EPositionSplit is the decoded non-indexed payload of PositionSplit.
type EPositionSplit struct {
	// signature: f35af27f5c587462d5a0f029fb4c578bc53bc3c87db59b8db657010446b46455
	Amount *big.Int
}

// EPositionsMerged is the decoded non-indexed payload of PositionsMerged.
type EPositionsMerged struct {
	// signature: 81ac7d05a625b7e14ab408012a5fbd6cbf06f0e9d28d2d5413d078e323a5dc2c
	Amounts       [8]*big.Int
	CollateralOut *big.Int
}

// ERoundSettled is the decoded non-indexed payload of RoundSettled.
type ERoundSettled struct {
	// signature: 7bac5aa6107fc77cc2426c3c4f1491806f51281ee9dad946b5d3ad5dd6e7b469
	Values   uint8
	Escrowed [8]*big.Int
	Carried  *big.Int
}

// EClaimed is the decoded non-indexed payload of Claimed.
type EClaimed struct {
	// signature: d166975969e02ab94330b286dd849bfa3771eced0fed813cbf4492fd17e73790
	Amounts [8]*big.Int
	Total   *big.Int
}

// EAnswersObserved is the decoded non-indexed payload of AnswersObserved.
type EAnswersObserved struct {
	// signature: 6eed9ea57baa5a42a4d4294006a1f18100e35851110898d45bf710aaaddbea3b
	KnownMask uint8
	Values    uint8
}

// ECycleEscrowed is the decoded non-indexed payload of CycleEscrowed.
type ECycleEscrowed struct {
	// signature: 16cbc14eaf1c725d6bf8f56d127952bd60a5a107cc8f428d97ced68b4f5d589a
	Amounts [8]*big.Int
}

// EPaid is the decoded non-indexed payload of Paid.
type EPaid struct {
	// signature: 9f777629478be3b0998eab401d2606e8b3ac41e303df6a1d13f5f1cdf48ce29f
	Amount *big.Int
}

// ETransferSingle is the decoded non-indexed payload of ERC-1155
// TransferSingle.
type ETransferSingle struct {
	// signature: c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62
	Id    *big.Int
	Value *big.Int
}

// ETransferBatch is the decoded non-indexed payload of ERC-1155
// TransferBatch.
type ETransferBatch struct {
	// signature: 4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb
	Ids    []*big.Int
	Values []*big.Int
}

// EApprovalForAll is the decoded non-indexed payload of ERC-1155
// ApprovalForAll.
type EApprovalForAll struct {
	// signature: 17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31
	Approved bool
}

// =====================================================================
// Row Types (decoded domain events, one struct per cz_ table)
// =====================================================================

// EventBase carries the columns every cz_ event table shares.
type EventBase struct {
	EvtId     int64
	BlockNum  int64
	TxId      int64
	Contract  string
	TimeStamp int64
}

// RoundInitialized is one cz_round_initialized row.
type RoundInitialized struct {
	EventBase
	RoundId int64
}

// RoundFunded is one cz_round_funded row. Decimal values are uint256
// strings; Allocations holds the 8 per-question splits.
type RoundFunded struct {
	EventBase
	RoundId       int64
	Caller        string
	Donated       string
	TransfersUsed string
	CarryUsed     string
	Allocations   []string
}

// Trade is one cz_trade row.
type Trade struct {
	EventBase
	RoundId  int64
	User     string
	Question int16
	Yes      bool
	Buy      bool
	Cash     string
	Units    string
}

// PositionSplit is one cz_position_split row.
type PositionSplit struct {
	EventBase
	RoundId  int64
	User     string
	Question int16
	Amount   string
}

// PositionsMerged is one cz_positions_merged row.
type PositionsMerged struct {
	EventBase
	RoundId       int64
	User          string
	Amounts       []string
	CollateralOut string
}

// RoundSettled is one cz_round_settled row. Answers is ValuesMask unpacked:
// Answers[q] is bit q.
type RoundSettled struct {
	EventBase
	RoundId    int64
	ValuesMask int16
	Answers    []bool
	Escrowed   []string
	Carried    string
}

// Claim is one cz_claim row.
type Claim struct {
	EventBase
	RoundId int64
	User    string
	Amounts []string
	Total   string
}

// AnswersObserved is one cz_answers_observed row.
type AnswersObserved struct {
	EventBase
	RoundId    int64
	KnownMask  int16
	ValuesMask int16
}

// CycleEscrowed is one cz_cycle_escrowed row (vault).
type CycleEscrowed struct {
	EventBase
	RoundId int64
	Amounts []string
}

// Paid is one cz_paid row (vault).
type Paid struct {
	EventBase
	RoundId   int64
	Question  int16
	Recipient string
	Amount    string
}

// PositionTransferItem is one (id, value) pair of an ERC-1155 transfer with
// the position id decoded ((roundId << 4) | (question << 1) | yes).
type PositionTransferItem struct {
	PositionId string
	RoundId    int64
	Question   int16
	Yes        bool
	Amount     string
}

// PositionTransfer is one decoded TransferSingle or TransferBatch: n items
// stored as n cz_position_transfer rows sharing the evtlog_id.
type PositionTransfer struct {
	EventBase
	Operator string
	From     string
	To       string
	Items    []PositionTransferItem
}

// ApprovalForAll is one cz_approval_for_all row.
type ApprovalForAll struct {
	EventBase
	Owner    string
	Operator string
	Approved bool
}
