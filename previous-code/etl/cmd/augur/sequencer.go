package main

import (
	"bytes"

	"github.com/ethereum/go-ethereum/core/types"
)
type EventSequencer struct {	// determines the order for contained events
	unordered_list		[]*types.Log
}
const (
	AgTxType_Unclassified        = iota
	AgTxType_MarketOrder
	AgTxType_MarketFinalized
)
func (sequencer *EventSequencer) append_event(new_log *types.Log) {

	sequencer.unordered_list = append(sequencer.unordered_list,new_log)
}
func (sequencer *EventSequencer) get_ordered_event_list() []*types.Log {
	// determines the correct event sequence for different event combinations
	// this variation returns events in inverted order
	return sequencer.unordered_list
}
func (sequencer *EventSequencer) get_inverse_ordered_event_list() []*types.Log {
	// determines the correct event sequence for different event combinations
	// this variation returns events in inverted order
	output := make([]*types.Log,0,8)
	for i := len(sequencer.unordered_list) - 1; i >= 0; i-- {
		output = append(output,sequencer.unordered_list[i])
	}
	return output
}
func (sequencer *EventSequencer) get_events_for_market_finalized_case() []*types.Log {
	// we must move TradingProceedsClaimed events to the end, so they execut before ProfitLoss events

	output := make([]*types.Log,0,8)
	var market_finalized *types.Log
	profit_loss_events := make([]*types.Log,0,8)
	proceed_events := make([]*types.Log,0,8)
	other_events := make([]*types.Log,0,8)

	for i := 0 ; i < len(sequencer.unordered_list) ; i++ {
		if len(sequencer.unordered_list[i].Topics) == 0 {
			other_events = append(other_events,sequencer.unordered_list[i])
			continue
		}
		if 0 == bytes.Compare(sequencer.unordered_list[i].Topics[0].Bytes(),evt_profit_loss_changed) {
			profit_loss_events = append(profit_loss_events,sequencer.unordered_list[i])
			continue
		}
		if 0 == bytes.Compare(sequencer.unordered_list[i].Topics[0].Bytes(),evt_trading_proceeds_claimed) {
			proceed_events = append(proceed_events,sequencer.unordered_list[i])
			continue
		}
		if 0 == bytes.Compare(sequencer.unordered_list[i].Topics[0].Bytes(),evt_market_finalized) {
			market_finalized = sequencer.unordered_list[i]
			continue
		}
		other_events = append(other_events,sequencer.unordered_list[i])
	}
	output = append(output,market_finalized)
	for i := 0; i < len(profit_loss_events) ; i++ {
		output = append(output,profit_loss_events[i])
	}
	for i := 0; i < len(proceed_events); i++ {
		output = append(output,proceed_events[i])
	}
	for i := 0; i < len(other_events); i++ {
		output = append(output,other_events[i])
	}
	return output
}
func (sequencer *EventSequencer) get_events_for_market_order_case() []*types.Log {
	// single Ethereum transaction can pack multiple Market Orders, so we need to put
	// all OderEvent events at the beginning, and after that all the remaining events
	// (we need to put order events at the beginning because the other events are linked
	//		to 'mktord' table by ID (like ProfitLoss event for example)
	// Note: the event's can be processed in inverted order because MarketVolumeChanged events
	//	contain 'Volume' field with depends on the order because it is a sum operation

	output := make([]*types.Log,0,8)
	var fill_event *types.Log
	marketorder_events := make([]*types.Log,0,8)
	other_events := make([]*types.Log,0,8)

	for i := 0 ; i < len(sequencer.unordered_list) ; i++ {
		if len(sequencer.unordered_list[i].Topics) == 0 {
			other_events = append(other_events,sequencer.unordered_list[i])
			continue
		}
		if 0 == bytes.Compare(sequencer.unordered_list[i].Topics[0].Bytes(),evt_exchange_fill) {
			fill_event = sequencer.unordered_list[i]
			continue
		}
		if 0 == bytes.Compare(sequencer.unordered_list[i].Topics[0].Bytes(),evt_market_order) {
			marketorder_events = append(marketorder_events,sequencer.unordered_list[i])
			continue
		}
		other_events = append(other_events,sequencer.unordered_list[i])
	}
	output = append(output,fill_event)	// Fill events goes first because we need to extract initial_amount
	for i := 0; i < len(marketorder_events) ; i++ {
		output = append(output,marketorder_events[i])
	}
	for i := 0; i < len(other_events); i++ {
		output = append(output,other_events[i])
	}
	return output
}
