// Donation events: ETH/ERC20/NFT donations, donated token/NFT claims and
// charity wallet traffic (donations received/sent, funds transferred to charity).
package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// bidIDByEvtlog adapts Repo.BidIDByEvtlog to the handler contract: -1 when
// the event log carries no bid, real DB errors propagate.
func bidIDByEvtlog(ctx context.Context, evtlogID int64) (int64, error) {
	id, err := cgRepo.BidIDByEvtlog(ctx, evtlogID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return -1, nil
		}
		return 0, fmt.Errorf("BidIDByEvtlog(%v): %w", evtlogID, err)
	}
	return id, nil
}

func proc_donation_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGDonationEvent
	var eth_evt CosmicSignatureGameEthDonated

	Info.Printf("Processing DonationEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "EthDonated", log.Data)
	if err != nil {
		return fmt.Errorf("EthDonated (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.Amount = eth_evt.Amount.String()
	evt.RoundNum = log.Topics[1].Big().Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("DonationEvent {\n")
	Info.Printf("\tDonor: %v\n", evt.DonorAddr)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("\tRound: %v\n", evt.RoundNum)
	Info.Printf("}\n")

	if err := cgRepo.DeleteEthDonation(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertEthDonation(ctx, &evt)
}
func get_donation_data(record_id int64) (string, error) {

	cosmic_game_ctrct, err := NewCosmicSignatureGame(cosmic_game_addr, eclient)
	if err != nil {
		return "", err
	}
	var copts bind.CallOpts
	dinfo_rec, err := cosmic_game_ctrct.EthDonationWithInfoRecords(&copts, big.NewInt(record_id))
	if err != nil {
		return "", err
	}
	return dinfo_rec.Data, err
}
func proc_donation_with_info_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGDonationWithInfoEvent
	var eth_evt CosmicSignatureGameEthDonatedWithInfo

	Info.Printf("Processing DonationWithInfoEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), cosmic_game_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "EthDonatedWithInfo", log.Data)
	if err != nil {
		return fmt.Errorf("EthDonatedWithInfo (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.RecordId = log.Topics[3].Big().Int64()
	evt.Amount = eth_evt.Amount.String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	data_json, err := get_donation_data(evt.RecordId)
	if err != nil {
		return fmt.Errorf("EthDonatedWithInfo (evt id %v): fetching donation info record: %w", elog.EvtId, err)
	}

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("DonationWithInfoEvent {\n")
	Info.Printf("\tDonor: %v\n", evt.DonorAddr)
	Info.Printf("\tRecordId: %v\n", evt.RecordId)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("\tData JSON: %v\n", data_json)
	Info.Printf("}\n")

	if err := cgRepo.DeleteEthDonationWithInfo(ctx, evt.EvtId); err != nil {
		return err
	}
	if err := cgRepo.InsertEthDonationWithInfo(ctx, &evt); err != nil {
		return err
	}
	return cgRepo.InsertDonationJSON(ctx, evt.RecordId, data_json)
}
func proc_donation_received_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGDonationReceivedEvent
	var eth_evt CharityWalletDonationReceived

	Info.Printf("Processing DonationReceivedEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), charity_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt, "DonationReceived", log.Data)
	if err != nil {
		return fmt.Errorf("DonationReceived (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()
	evt.RoundNum, err = find_prize_num(ctx, evt.TxId)
	if err != nil {
		return fmt.Errorf("DonationReceived (evt id %v): %w", elog.EvtId, err)
	}

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("DonationReceivedEvent {\n")
	Info.Printf("\tDonor: %v\n", evt.DonorAddr)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("}\n")

	if err := cgRepo.DeleteDonationReceived(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertDonationReceived(ctx, &evt)
}
func proc_donation_sent_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGDonationSentEvent
	var eth_evt CharityWalletFundsTransferredToCharity

	Info.Printf("Processing DonationSentEvent event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), charity_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := charity_wallet_abi.UnpackIntoInterface(&eth_evt, "FundsTransferredToCharity", log.Data)
	if err != nil {
		return fmt.Errorf("FundsTransferredToCharity (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("FundsTransferredToCharity{\n")
	Info.Printf("\tCharity: %v\n", evt.CharityAddr)
	Info.Printf("\tAmount%v\n", evt.Amount)
	Info.Printf("}\n")

	if err := cgRepo.DeleteDonationSent(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertDonationSent(ctx, &evt)
}

// get_token_uri fetches tokenURI for a donated NFT. A failing contract call
// yields "" (the donation is stored without a URI, exactly like the legacy
// path); only a broken contract binding is an error.
func get_token_uri(token_id int64, contract_addr common.Address) (string, error) {

	c, err := NewCosmicSignatureNft(contract_addr, eclient) // we use cosmicsignature because its ERC721
	if err != nil {
		return "", fmt.Errorf("get_token_uri(): contract creation: %w", err)
	}
	var copts bind.CallOpts
	tok_uri, err := c.TokenURI(&copts, big.NewInt(token_id))
	if err != nil {
		err_str := fmt.Sprintf("Error at get_token_uri() during GetTokenURI() call: %v", err)
		Info.Print(err_str)
		Error.Print(err_str)
		return "", nil
	}
	return tok_uri, nil
}
func proc_erc20_donated_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGERC20DonationEvent
	var eth_evt IPrizesWalletTokenDonated

	Info.Printf("Processing TokenDonated event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt, "TokenDonated", log.Data)
	if err != nil {
		return fmt.Errorf("TokenDonated (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.DonorAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.Amount = eth_evt.Amount.String()
	evt.BidId, err = bidIDByEvtlog(ctx, evt.EvtId-1)
	if err != nil {
		return err
	}
	if evt.BidId == -1 { // if BidId = -1 , it could be that EvtId - 1 falls on Approval event, so Bid event will be EvtId - 2
		evt.BidId, err = bidIDByEvtlog(ctx, evt.EvtId-2)
		if err != nil {
			return err
		}
	}

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("TokenDonated{\n")
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tDonor: %v\n", evt.DonorAddr)

	Info.Printf("\tTokenAddr: %v\n", evt.TokenAddr)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("\tLinking to bid id: %v\n", evt.BidId)
	Info.Printf("}\n")

	if err := cgRepo.DeleteERC20Donation(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertERC20Donation(ctx, &evt)
}
func proc_nft_donation_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGNFTDonationEvent
	var eth_evt IPrizesWalletNftDonated

	Info.Printf("Processing NftDonated event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt, "NftDonated", log.Data)
	if err != nil {
		return fmt.Errorf("NftDonated (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.TokenId = eth_evt.NftId.Int64()
	// BidRowIDByEvtlogID reports 0 when the previous event carries no bid
	// (e.g. a pure Donate() call); only real DB failures error.
	evt.BidId, err = cgRepo.BidRowIDByEvtlogID(ctx, evt.EvtId-1)
	if err != nil {
		return fmt.Errorf("NftDonated (evt id %v): %w", elog.EvtId, err)
	}
	evt.NFTTokenURI, err = get_token_uri(evt.TokenId, common.HexToAddress(evt.TokenAddr))
	if err != nil {
		return fmt.Errorf("NftDonated (evt id %v): %w", elog.EvtId, err)
	}
	evt.Index = eth_evt.Index.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("NFTDonationEvent {\n")
	Info.Printf("\tDonor: %v\n", evt.DonorAddr)
	Info.Printf("\tNFTAddress: %v\n", evt.TokenAddr)
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("\tNFTTokenURI: %v\n", evt.NFTTokenURI)
	Info.Printf("}\n")

	if err := cgRepo.DeleteNFTDonation(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertNFTDonation(ctx, &evt)
}
func proc_donated_token_claimed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGDonatedTokenClaimed
	var eth_evt PrizesWalletDonatedTokenClaimed

	Info.Printf("Processing DonatedTokenClaimed event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt, "DonatedTokenClaimed", log.Data)
	if err != nil {
		return fmt.Errorf("DonatedTokenClaimed (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenAddr = eth_evt.TokenAddress.String()
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.Amount = eth_evt.Amount.String()
	evt.BeneficiaryAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[3][12:]).String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("DonatedTokenClaimedEvent{\n")
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tBeneficiary: %v\n", evt.BeneficiaryAddr)

	Info.Printf("\tTokenAddr: %v\n", evt.TokenAddr)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("}\n")

	if err := cgRepo.DeleteDonatedTokenClaim(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertDonatedTokenClaim(ctx, &evt)
}
func proc_donated_nft_claimed_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGDonatedNFTClaimed
	var eth_evt PrizesWalletDonatedNftClaimed

	Info.Printf("Processing DonatedNFTClaimed event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)

	if !bytes.Equal(log.Address.Bytes(), prizes_wallet_addr.Bytes()) {
		Info.Printf("Event doesn't belong to known address set (addr=%v), skipping\n", log.Address.String())
		return nil
	}
	err := prizes_wallet_abi.UnpackIntoInterface(&eth_evt, "DonatedNftClaimed", log.Data)
	if err != nil {
		return fmt.Errorf("DonatedNftClaimed (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = log.Topics[1].Big().Int64()
	evt.BeneficiaryAddr = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenAddr = common.BytesToAddress(log.Topics[3][12:]).String()
	evt.TokenId = eth_evt.NftId.String()
	evt.Index = eth_evt.Index.Int64()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("DonatedNFTClaimedEvent{\n")
	Info.Printf("\tRoundNum: %v\n", evt.RoundNum)
	Info.Printf("\tIndex: %v\n", evt.Index)
	Info.Printf("\tBeneficiary: %v\n", evt.BeneficiaryAddr)

	Info.Printf("\tTokenAddr: %v\n", evt.TokenAddr)
	Info.Printf("\tTokenId: %v\n", evt.TokenId)
	Info.Printf("}\n")

	if err := cgRepo.DeleteDonatedNFTClaim(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertDonatedNFTClaim(ctx, &evt)
}
func proc_funds_transferred_to_charity_event(ctx context.Context, log *types.Log, elog *EthereumEventLog) error {

	var evt CGFundsToCharity
	var eth_evt CosmicSignatureGameFundsTransferredToCharity

	if !bytes.Equal(log.Address.Bytes(), marketing_wallet_addr.Bytes()) {
		return nil
	}
	Info.Printf("Processing FundsTransferredToCharity event id=%v, txhash %v\n", elog.EvtId, elog.TxHash)
	err := cosmic_game_abi.UnpackIntoInterface(&eth_evt, "FundsTransferredToCharity", log.Data)
	if err != nil {
		return fmt.Errorf("FundsTransferredToCharity (evt id %v): decode: %w", elog.EvtId, err)
	}

	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = log.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CharityAddr = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.Amount = eth_evt.Amount.String()

	Info.Printf("Contract: %v\n", log.Address.String())
	Info.Printf("FundsTransferredToCharity{\n")
	Info.Printf("\tCharityAddress: %v\n", evt.CharityAddr)
	Info.Printf("\tAmount: %v\n", evt.Amount)
	Info.Printf("}\n")

	if err := cgRepo.DeleteFundsToCharity(ctx, evt.EvtId); err != nil {
		return err
	}
	return cgRepo.InsertFundsToCharity(ctx, &evt)
}
