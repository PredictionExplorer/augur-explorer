// RandomWalk event processing: event-signature registry, marketplace and NFT
// event decoders (offers, buys, withdrawals, token names, mints, transfers) and dispatch.
package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
)

const (
	NEW_OFFER      = "55076e90b6b34a2569ffb2e1e34ee0da92d30ca423f0d6cfb317d252ade9a56a"
	ITEM_BOUGHT    = "caacc56f18ca259dc5175dae29eb0ca81407703a4819958c6885acbb7d4f3af3"
	OFFER_CANCELED = "0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951"
	WITHDRAWAL_EVT = "a11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7"
	TOKEN_NAME_EVT = "8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12"
	MINT_EVENT     = "ad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec"
	TRANSFER_EVT   = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

var (
	evt_new_offer, _      = hex.DecodeString(NEW_OFFER)
	evt_item_bought, _    = hex.DecodeString(ITEM_BOUGHT)
	evt_offer_canceled, _ = hex.DecodeString(OFFER_CANCELED)
	evt_withdrawal, _     = hex.DecodeString(WITHDRAWAL_EVT)
	evt_token_name, _     = hex.DecodeString(TOKEN_NAME_EVT)
	evt_transfer, _       = hex.DecodeString(TRANSFER_EVT)
	evt_mint_event, _     = hex.DecodeString(MINT_EVENT)
)

func proc_new_offer(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt rwp.NewOffer
	var eth_evt rwp.ENewOffer

	Info.Printf("Processing NewOffer event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	err := marketplace_abi.UnpackIntoInterface(&eth_evt, "NewOffer", log.Data)
	if err != nil {
		return fmt.Errorf("event NewOffer decode: %w", err)
	}

	if !bytes.Equal(log.Address.Bytes(), market_addr.Bytes()) {
		Info.Printf("Event doesn't belong to know address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Seller = eth_evt.Seller.String()
	evt.Buyer = eth_evt.Buyer.String()
	evt.Price = eth_evt.Price.String()
	evt.RWalkAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.OfferId = log.Topics[2].Big().Int64()
	evt.TokenId = log.Topics[3].Big().Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("NewOffer {\n")
	Info.Printf("\tNFT addr: %v\n", evt.RWalkAddr)
	Info.Printf("\tOfferId: %v\n", evt.OfferId)
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tSeller: %v\n", evt.Seller)
	Info.Printf("\tBuyer: %v\n", evt.Buyer)
	Info.Printf("}\n")

	return rwRepo.InsertNewOffer(ctx, &evt)
}

func proc_item_bought(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt rwp.ItemBought

	Info.Printf("Processing ItemBought id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), market_addr.Bytes()) {
		Info.Printf("Event doesn't belong to know address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OfferId = log.Topics[1].Big().Int64()
	evt.SellerAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.BuyerAddr = common.BytesToAddress(log.Topics[3][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("ItemBought {\n")
	Info.Printf("\tOfferId: %v\n", evt.OfferId)
	Info.Printf("\tSeller: %v\n", evt.SellerAddr)
	Info.Printf("\tBuyer: %v\n", evt.BuyerAddr)
	Info.Printf("}\n")

	return rwRepo.InsertItemBought(ctx, &evt)
}

func proc_offer_cancelled(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt rwp.OfferCanceled

	if !bytes.Equal(log.Address.Bytes(), market_addr.Bytes()) {
		Info.Printf("Event doesn't belong to know address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.OfferId = log.Topics[1].Big().Int64()

	Info.Printf("Processing OfferCanceled id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	exists, err := rwRepo.OfferExists(ctx, log.Address.String(), evt.OfferId)
	if err != nil {
		return fmt.Errorf("offer exists check for OfferCanceled: %w", err)
	}
	if !exists {
		Info.Printf(
			"Skipping OfferCanceled : offer %v for contract %v does not exist, skipping\n",
			evt.OfferId, log.Address.String(),
		)
		return nil
	}

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("OfferCanceled {\n")
	Info.Printf("\tOfferId: %v\n", evt.OfferId)
	Info.Printf("}\n")

	return rwRepo.InsertOfferCanceled(ctx, &evt)
}

func proc_withdrawal(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt rwp.Withdrawal
	var eth_evt rwp.EWithdrawalEvent

	Info.Printf("Processing WithdrawalEvent id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(), rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to know address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := randomwalk_abi.UnpackIntoInterface(&eth_evt, "WithdrawalEvent", log.Data)
	if err != nil {
		return fmt.Errorf("event WithdrawalEvent decode: %w", err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = log.Topics[1].Big().Int64()
	evt.Destination = eth_evt.Destination.String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("Withdrawal {\n")
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tDestination: %v\n", evt.Destination)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("}\n")

	return rwRepo.InsertWithdrawal(ctx, &evt)
}

func proc_token_name(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt rwp.TokenName
	var eth_evt rwp.ETokenNameEvent

	Info.Printf("Processing TokenName id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	if !bytes.Equal(log.Address.Bytes(), rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := randomwalk_abi.UnpackIntoInterface(&eth_evt, "TokenNameEvent", log.Data)
	if err != nil {
		return fmt.Errorf("event TokenName decode: %w", err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = eth_evt.TokenId.Int64()
	evt.NewName = eth_evt.NewName

	exists, err := rwRepo.TokenExists(ctx, log.Address.String(), evt.TokenId)
	if err != nil {
		return fmt.Errorf("token exists check for TokenName: %w", err)
	}
	if !exists {
		Info.Printf("Token name event skipped, token contract %v is not registered\n", log.Address.String())
		return nil
	}

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("TokenName {\n")
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tNewName: %v\n", evt.NewName)
	Info.Printf("}\n")

	return rwRepo.InsertTokenName(ctx, &evt)
}

func proc_mint_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt rwp.MintEvent
	var eth_evt rwp.EMintEvent

	Info.Printf("Processing MintEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), rwalk_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := randomwalk_abi.UnpackIntoInterface(&eth_evt, "MintEvent", log.Data)
	if err != nil {
		return fmt.Errorf("event MintEvent decode: %w", err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = log.Topics[1].Big().Int64()
	evt.Owner = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Seed = hex.EncodeToString(eth_evt.Seed[:])
	evt.SeedNum = common.BytesToHash(eth_evt.Seed[:]).Big().String()
	evt.Price = eth_evt.Price.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("MintEvent {\n")
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tOwner %v\n", evt.Owner)
	Info.Printf("\tSeed: %v\n", evt.Seed)
	Info.Printf("\tSeed Numeric: %v\n", evt.SeedNum)
	Info.Printf("\tPrice: %v\n", evt.Price)
	Info.Printf("}\n")

	return rwRepo.InsertMint(ctx, &evt)
}

func proc_transfer_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt rwp.Transfer

	if !bytes.Equal(log.Address.Bytes(), rwalk_addr.Bytes()) {
		Info.Printf("Skipping another instance of RandomWalk contract %v\n", log.Address.String())
		return nil
	}
	Info.Printf("Processing Token Transfer event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.To = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenId = log.Topics[3].Big().Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("Transfer {\n")
	Info.Printf("\tFrom: %v\n", evt.From)
	Info.Printf("\tTo: %v\n", evt.To)
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("}\n")

	return rwRepo.InsertTransfer(ctx, &evt)
}

// eventDispatchEntry pairs a decoded topic-0 signature with its handler.
type eventDispatchEntry struct {
	topic0  []byte
	handler func(context.Context, *types.Log, *EthereumEventLog) error
}

// eventDispatchTable returns the topic → handler registry. Built per call
// because the evt_* topic variables and the handlers' address guards are
// package globals initialized by main()/the test harness.
func eventDispatchTable() []eventDispatchEntry {
	return []eventDispatchEntry{
		{evt_new_offer, proc_new_offer},
		{evt_item_bought, proc_item_bought},
		{evt_offer_canceled, proc_offer_cancelled},
		{evt_withdrawal, proc_withdrawal},
		{evt_token_name, proc_token_name},
		{evt_transfer, proc_transfer_event},
		{evt_mint_event, proc_mint_event},
	}
}

// select_event_and_process dispatches the log to every matching event
// handler. Any handler error (decode failure or DB write failure) stops the
// dispatch and is returned to the polling loop, which leaves the batch
// unacknowledged for re-processing.
func select_event_and_process(ctx context.Context, log *types.Log, evtlog *EthereumEventLog) error {
	Info.Printf("processing event with sig = %v\n", log.Topics[0].String())
	topic0 := log.Topics[0].Bytes()
	for _, entry := range eventDispatchTable() {
		if !bytes.Equal(topic0, entry.topic0) {
			continue
		}
		if err := entry.handler(ctx, log, evtlog); err != nil {
			return err
		}
	}
	return nil
}

// process_single_event loads the stored evt_log row, reconstructs the
// Ethereum log from its RLP and dispatches it. All failures — a missing row,
// a corrupt RLP payload (previously a panic) or a handler error — are
// returned to the caller.
func process_single_event(ctx context.Context, evt_id int64) error {

	evtlog, err := dbStore.EventLog(ctx, evt_id)
	if err != nil {
		return fmt.Errorf("process_single_event(%v): %w", evt_id, err)
	}
	var log types.Log
	err = rlp.DecodeBytes(evtlog.RlpLog, &log)
	if err != nil {
		return fmt.Errorf("process_single_event(%v): RLP decode: %w", evt_id, err)
	}
	log.BlockNumber = uint64(evtlog.BlockNum)
	log.TxHash.SetBytes(common.HexToHash(evtlog.TxHash).Bytes())
	log.Address.SetBytes(common.HexToHash(evtlog.ContractAddress).Bytes())
	num_topics := len(log.Topics)
	if num_topics > 0 {
		return select_event_and_process(ctx, &log, &evtlog)
	}
	return nil
}
