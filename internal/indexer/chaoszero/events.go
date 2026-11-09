// Decode/store pairs for the dispatched ChaosZero events.

package chaoszero

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	czmodel "github.com/PredictionExplorer/augur-explorer/internal/model/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// errRoundIDOverflow rejects a roundId topic that does not fit the BIGINT
// column; decode must fail the batch, never wrap.
var errRoundIDOverflow = errors.New("roundId exceeds int64")

// errQuestionOverflow rejects a question topic outside the uint8 range the
// contract emits (the eight markets use 0..7).
var errQuestionOverflow = errors.New("question exceeds uint8")

// eventBase fills the columns every cz_ table shares.
func eventBase(lg *types.Log, elog *store.EthereumEventLog) czmodel.EventBase {
	return czmodel.EventBase{
		EvtId:     elog.EvtID,
		BlockNum:  elog.BlockNum,
		TxId:      elog.TxID,
		Contract:  lg.Address.String(),
		TimeStamp: elog.TimeStamp,
	}
}

// roundIDFromTopic converts an indexed uint256 roundId topic to int64.
func roundIDFromTopic(topic ethcommon.Hash) (int64, error) {
	v := topic.Big()
	if !v.IsInt64() {
		return 0, errRoundIDOverflow
	}
	return v.Int64(), nil
}

// questionFromTopic converts an indexed uint8 question topic to int16.
func questionFromTopic(topic ethcommon.Hash) (int16, error) {
	v := topic.Big()
	if !v.IsUint64() || v.Uint64() > 255 {
		return 0, errQuestionOverflow
	}
	return int16(v.Uint64()), nil
}

// topicAddress extracts an indexed address topic.
func topicAddress(topic ethcommon.Hash) string {
	return ethcommon.BytesToAddress(topic[12:]).String()
}

// bigArrayStrings converts an unpacked uint256[8] into decimal strings for
// the NUMERIC[] columns.
func bigArrayStrings(vals [8]*big.Int) []string {
	out := make([]string, len(vals))
	for i, v := range vals {
		out[i] = v.String()
	}
	return out
}

// maskBools unpacks a uint8 outcome bitmask into its eight booleans (bit q
// is question q).
func maskBools(mask uint8) []bool {
	out := make([]bool, 8)
	for q := range out {
		out[q] = mask&(1<<q) != 0
	}
	return out
}

// =====================================================================
// IndependentOutcomeMarket
// =====================================================================

func (h *Handlers) decodeRoundInitialized(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.RoundInitialized, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	return &czmodel.RoundInitialized{
		EventBase: eventBase(lg, elog),
		RoundId:   roundID,
	}, nil
}

func (h *Handlers) storeRoundInitialized(ctx context.Context, evt *czmodel.RoundInitialized) error {
	h.log.Info("RoundInitialized", "evt_id", evt.EvtId, "round_id", evt.RoundId)

	return h.repo.InsertRoundInitialized(ctx, evt)
}

func (h *Handlers) decodeRoundFunded(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.RoundFunded, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt czmodel.ERoundFunded
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "RoundFunded", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	return &czmodel.RoundFunded{
		EventBase:     eventBase(lg, elog),
		RoundId:       roundID,
		Caller:        topicAddress(lg.Topics[2]),
		Donated:       ethEvt.Donated.String(),
		TransfersUsed: ethEvt.TransfersUsed.String(),
		CarryUsed:     ethEvt.CarryUsed.String(),
		Allocations:   bigArrayStrings(ethEvt.Allocations),
	}, nil
}

func (h *Handlers) storeRoundFunded(ctx context.Context, evt *czmodel.RoundFunded) error {
	h.log.Info("RoundFunded",
		"evt_id", evt.EvtId, "round_id", evt.RoundId, "caller", evt.Caller,
		"donated", evt.Donated, "transfers_used", evt.TransfersUsed, "carry_used", evt.CarryUsed)

	return h.repo.InsertRoundFunded(ctx, evt)
}

func (h *Handlers) decodeTrade(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.Trade, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt czmodel.ETrade
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "Trade", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	question, err := questionFromTopic(lg.Topics[3])
	if err != nil {
		return nil, err
	}
	return &czmodel.Trade{
		EventBase: eventBase(lg, elog),
		RoundId:   roundID,
		User:      topicAddress(lg.Topics[2]),
		Question:  question,
		Yes:       ethEvt.Yes,
		Buy:       ethEvt.Buy,
		Cash:      ethEvt.Cash.String(),
		Units:     ethEvt.Units.String(),
	}, nil
}

func (h *Handlers) storeTrade(ctx context.Context, evt *czmodel.Trade) error {
	h.log.Info("Trade",
		"evt_id", evt.EvtId, "round_id", evt.RoundId, "user", evt.User,
		"question", evt.Question, "yes", evt.Yes, "buy", evt.Buy,
		"cash", evt.Cash, "units", evt.Units)

	return h.repo.InsertTrade(ctx, evt)
}

func (h *Handlers) decodePositionSplit(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.PositionSplit, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt czmodel.EPositionSplit
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "PositionSplit", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	question, err := questionFromTopic(lg.Topics[3])
	if err != nil {
		return nil, err
	}
	return &czmodel.PositionSplit{
		EventBase: eventBase(lg, elog),
		RoundId:   roundID,
		User:      topicAddress(lg.Topics[2]),
		Question:  question,
		Amount:    ethEvt.Amount.String(),
	}, nil
}

func (h *Handlers) storePositionSplit(ctx context.Context, evt *czmodel.PositionSplit) error {
	h.log.Info("PositionSplit",
		"evt_id", evt.EvtId, "round_id", evt.RoundId, "user", evt.User,
		"question", evt.Question, "amount", evt.Amount)

	return h.repo.InsertPositionSplit(ctx, evt)
}

func (h *Handlers) decodePositionsMerged(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.PositionsMerged, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt czmodel.EPositionsMerged
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "PositionsMerged", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	return &czmodel.PositionsMerged{
		EventBase:     eventBase(lg, elog),
		RoundId:       roundID,
		User:          topicAddress(lg.Topics[2]),
		Amounts:       bigArrayStrings(ethEvt.Amounts),
		CollateralOut: ethEvt.CollateralOut.String(),
	}, nil
}

func (h *Handlers) storePositionsMerged(ctx context.Context, evt *czmodel.PositionsMerged) error {
	h.log.Info("PositionsMerged",
		"evt_id", evt.EvtId, "round_id", evt.RoundId, "user", evt.User,
		"collateral_out", evt.CollateralOut)

	return h.repo.InsertPositionsMerged(ctx, evt)
}

func (h *Handlers) decodeRoundSettled(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.RoundSettled, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt czmodel.ERoundSettled
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "RoundSettled", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	return &czmodel.RoundSettled{
		EventBase:  eventBase(lg, elog),
		RoundId:    roundID,
		ValuesMask: int16(ethEvt.Values),
		Answers:    maskBools(ethEvt.Values),
		Escrowed:   bigArrayStrings(ethEvt.Escrowed),
		Carried:    ethEvt.Carried.String(),
	}, nil
}

func (h *Handlers) storeRoundSettled(ctx context.Context, evt *czmodel.RoundSettled) error {
	h.log.Info("RoundSettled",
		"evt_id", evt.EvtId, "round_id", evt.RoundId,
		"values_mask", evt.ValuesMask, "carried", evt.Carried)

	return h.repo.InsertRoundSettled(ctx, evt)
}

func (h *Handlers) decodeClaimed(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.Claim, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt czmodel.EClaimed
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "Claimed", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	return &czmodel.Claim{
		EventBase: eventBase(lg, elog),
		RoundId:   roundID,
		User:      topicAddress(lg.Topics[2]),
		Amounts:   bigArrayStrings(ethEvt.Amounts),
		Total:     ethEvt.Total.String(),
	}, nil
}

func (h *Handlers) storeClaimed(ctx context.Context, evt *czmodel.Claim) error {
	h.log.Info("Claimed",
		"evt_id", evt.EvtId, "round_id", evt.RoundId, "user", evt.User, "total", evt.Total)

	return h.repo.InsertClaim(ctx, evt)
}

func (h *Handlers) decodeAnswersObserved(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.AnswersObserved, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt czmodel.EAnswersObserved
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "AnswersObserved", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	return &czmodel.AnswersObserved{
		EventBase:  eventBase(lg, elog),
		RoundId:    roundID,
		KnownMask:  int16(ethEvt.KnownMask),
		ValuesMask: int16(ethEvt.Values),
	}, nil
}

func (h *Handlers) storeAnswersObserved(ctx context.Context, evt *czmodel.AnswersObserved) error {
	h.log.Info("AnswersObserved",
		"evt_id", evt.EvtId, "round_id", evt.RoundId,
		"known_mask", evt.KnownMask, "values_mask", evt.ValuesMask)

	return h.repo.InsertAnswersObserved(ctx, evt)
}

// =====================================================================
// OutcomeSettlementVault
// =====================================================================

func (h *Handlers) decodeCycleEscrowed(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.CycleEscrowed, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt czmodel.ECycleEscrowed
	if err := h.vaultABI.UnpackIntoInterface(&ethEvt, "CycleEscrowed", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	return &czmodel.CycleEscrowed{
		EventBase: eventBase(lg, elog),
		RoundId:   roundID,
		Amounts:   bigArrayStrings(ethEvt.Amounts),
	}, nil
}

func (h *Handlers) storeCycleEscrowed(ctx context.Context, evt *czmodel.CycleEscrowed) error {
	h.log.Info("CycleEscrowed", "evt_id", evt.EvtId, "round_id", evt.RoundId)

	return h.repo.InsertCycleEscrowed(ctx, evt)
}

func (h *Handlers) decodePaid(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.Paid, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt czmodel.EPaid
	if err := h.vaultABI.UnpackIntoInterface(&ethEvt, "Paid", lg.Data); err != nil {
		return nil, err
	}
	roundID, err := roundIDFromTopic(lg.Topics[1])
	if err != nil {
		return nil, err
	}
	question, err := questionFromTopic(lg.Topics[2])
	if err != nil {
		return nil, err
	}
	return &czmodel.Paid{
		EventBase: eventBase(lg, elog),
		RoundId:   roundID,
		Question:  question,
		Recipient: topicAddress(lg.Topics[3]),
		Amount:    ethEvt.Amount.String(),
	}, nil
}

func (h *Handlers) storePaid(ctx context.Context, evt *czmodel.Paid) error {
	h.log.Info("Paid",
		"evt_id", evt.EvtId, "round_id", evt.RoundId, "question", evt.Question,
		"recipient", evt.Recipient, "amount", evt.Amount)

	return h.repo.InsertPaid(ctx, evt)
}

// =====================================================================
// OpenZeppelin ERC-1155 events the market inherits
// =====================================================================

// positionItem decodes one ERC-1155 (id, value) pair. The market mints
// position ids as (roundId << 4) | (question << 1) | yes
// (IndependentOutcomeMarket.positionId), so the components are recovered by
// shifting; a round beyond int64 fails the batch like every other overflow.
func positionItem(id, value *big.Int) (czmodel.PositionTransferItem, error) {
	round := new(big.Int).Rsh(id, 4)
	if !round.IsInt64() {
		return czmodel.PositionTransferItem{}, fmt.Errorf("position id %s: %w", id.String(), errRoundIDOverflow)
	}
	low := id.Bit(0)
	question := int16(new(big.Int).Rsh(id, 1).Uint64() & 0x7) // #nosec G115 -- masked to 3 bits
	return czmodel.PositionTransferItem{
		PositionId: id.String(),
		RoundId:    round.Int64(),
		Question:   question,
		Yes:        low == 1,
		Amount:     value.String(),
	}, nil
}

func (h *Handlers) decodeTransferSingle(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.PositionTransfer, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt czmodel.ETransferSingle
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "TransferSingle", lg.Data); err != nil {
		return nil, err
	}
	item, err := positionItem(ethEvt.Id, ethEvt.Value)
	if err != nil {
		return nil, err
	}
	return &czmodel.PositionTransfer{
		EventBase: eventBase(lg, elog),
		Operator:  topicAddress(lg.Topics[1]),
		From:      topicAddress(lg.Topics[2]),
		To:        topicAddress(lg.Topics[3]),
		Items:     []czmodel.PositionTransferItem{item},
	}, nil
}

func (h *Handlers) decodeTransferBatch(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.PositionTransfer, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt czmodel.ETransferBatch
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "TransferBatch", lg.Data); err != nil {
		return nil, err
	}
	if len(ethEvt.Ids) != len(ethEvt.Values) {
		return nil, fmt.Errorf("TransferBatch ids/values length mismatch: %d vs %d", len(ethEvt.Ids), len(ethEvt.Values))
	}
	items := make([]czmodel.PositionTransferItem, 0, len(ethEvt.Ids))
	for i, id := range ethEvt.Ids {
		item, err := positionItem(id, ethEvt.Values[i])
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return &czmodel.PositionTransfer{
		EventBase: eventBase(lg, elog),
		Operator:  topicAddress(lg.Topics[1]),
		From:      topicAddress(lg.Topics[2]),
		To:        topicAddress(lg.Topics[3]),
		Items:     items,
	}, nil
}

func (h *Handlers) storePositionTransfer(ctx context.Context, evt *czmodel.PositionTransfer) error {
	h.log.Info("PositionTransfer",
		"evt_id", evt.EvtId, "operator", evt.Operator, "from", evt.From,
		"to", evt.To, "items", len(evt.Items))

	return h.repo.InsertPositionTransfer(ctx, evt)
}

func (h *Handlers) decodeApprovalForAll(lg *types.Log, elog *store.EthereumEventLog) (*czmodel.ApprovalForAll, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt czmodel.EApprovalForAll
	if err := h.marketABI.UnpackIntoInterface(&ethEvt, "ApprovalForAll", lg.Data); err != nil {
		return nil, err
	}
	return &czmodel.ApprovalForAll{
		EventBase: eventBase(lg, elog),
		Owner:     topicAddress(lg.Topics[1]),
		Operator:  topicAddress(lg.Topics[2]),
		Approved:  ethEvt.Approved,
	}, nil
}

func (h *Handlers) storeApprovalForAll(ctx context.Context, evt *czmodel.ApprovalForAll) error {
	h.log.Info("ApprovalForAll",
		"evt_id", evt.EvtId, "owner", evt.Owner, "operator", evt.Operator, "approved", evt.Approved)

	return h.repo.InsertApprovalForAll(ctx, evt)
}
