package v2

import (
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapRoundEthDonation(record cgstore.RoundEthDonationRecord) (RoundEthDonation, error) {
	if record.RoundNum < 0 {
		return RoundEthDonation{}, errors.New("invalid ETH donation round")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return RoundEthDonation{}, fmt.Errorf("ETH donation transaction: %w", err)
	}
	if !ethcommon.IsHexAddress(record.DonorAddr) {
		return RoundEthDonation{}, errors.New("invalid ETH donation donor address")
	}
	amount, err := requiredAmount(record.EthAmountWei)
	if err != nil {
		return RoundEthDonation{}, fmt.Errorf("ETH donation amount: %w", err)
	}

	result := RoundEthDonation{
		BlockNumber:     transaction.BlockNumber,
		DonorAddress:    ethcommon.HexToAddress(record.DonorAddr).Hex(),
		EthAmountWei:    amount,
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		TransactionHash: transaction.TransactionHash,
	}
	switch record.Kind {
	case cgstore.RoundEthDonationPlain:
		if record.ContractRecordID != nil || record.Data != nil {
			return RoundEthDonation{}, errors.New("plain ETH donation carries with-info fields")
		}
		result.Kind = Plain
	case cgstore.RoundEthDonationWithInfo:
		if record.ContractRecordID == nil || *record.ContractRecordID < 0 || record.Data == nil {
			return RoundEthDonation{}, errors.New("with-info ETH donation is missing variant fields")
		}
		recordID := *record.ContractRecordID
		data := *record.Data
		result.Kind = WithInfo
		result.ContractRecordId = &recordID
		result.Data = &data
	default:
		return RoundEthDonation{}, fmt.Errorf("unknown ETH donation kind %q", record.Kind)
	}
	return result, nil
}

func mapRoundERC20Donation(record cgstore.RoundERC20DonationRecord) (RoundErc20Donation, error) {
	if record.RoundNum < 0 {
		return RoundErc20Donation{}, errors.New("invalid ERC-20 donation round")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return RoundErc20Donation{}, fmt.Errorf("ERC-20 donation transaction: %w", err)
	}
	if !ethcommon.IsHexAddress(record.DonorAddr) {
		return RoundErc20Donation{}, errors.New("invalid ERC-20 donation donor address")
	}
	if !ethcommon.IsHexAddress(record.TokenAddr) {
		return RoundErc20Donation{}, errors.New("invalid ERC-20 donation token address")
	}
	amount, err := requiredAmount(record.AmountBaseUnits)
	if err != nil {
		return RoundErc20Donation{}, fmt.Errorf("ERC-20 donation amount: %w", err)
	}
	return RoundErc20Donation{
		AmountBaseUnits: amount,
		BlockNumber:     transaction.BlockNumber,
		DonorAddress:    ethcommon.HexToAddress(record.DonorAddr).Hex(),
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		TokenAddress:    ethcommon.HexToAddress(record.TokenAddr).Hex(),
		TransactionHash: transaction.TransactionHash,
	}, nil
}

func mapRoundNFTDonation(record cgstore.RoundNFTDonationRecord) (RoundNftDonation, error) {
	if record.RoundNum < 0 || record.TokenID < 0 || record.DonationIndex < 0 {
		return RoundNftDonation{}, errors.New("invalid NFT donation identity")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return RoundNftDonation{}, fmt.Errorf("NFT donation transaction: %w", err)
	}
	if !ethcommon.IsHexAddress(record.DonorAddr) {
		return RoundNftDonation{}, errors.New("invalid NFT donation donor address")
	}
	if !ethcommon.IsHexAddress(record.TokenAddr) {
		return RoundNftDonation{}, errors.New("invalid NFT donation token address")
	}
	return RoundNftDonation{
		BlockNumber:     transaction.BlockNumber,
		DonationIndex:   record.DonationIndex,
		DonorAddress:    ethcommon.HexToAddress(record.DonorAddr).Hex(),
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		TokenAddress:    ethcommon.HexToAddress(record.TokenAddr).Hex(),
		TokenId:         record.TokenID,
		TokenUri:        record.TokenURI,
		TransactionHash: transaction.TransactionHash,
	}, nil
}
