package v2

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapUserRaffleEthDeposit(record cgstore.UserRaffleEthDepositRecord) (UserRaffleEthDeposit, error) {
	if record.RoundNum < 0 || record.WinnerIndex < 0 {
		return UserRaffleEthDeposit{}, errors.New("invalid deposit identity")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserRaffleEthDeposit{}, fmt.Errorf("deposit transaction: %w", err)
	}
	amount, err := requiredAmount(record.EthAmountWei)
	if err != nil {
		return UserRaffleEthDeposit{}, fmt.Errorf("deposit amount: %w", err)
	}

	result := UserRaffleEthDeposit{
		BlockNumber:     transaction.BlockNumber,
		Claimed:         record.Claimed,
		EthAmountWei:    amount,
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		Source:          Raffle,
		TransactionHash: transaction.TransactionHash,
		WinnerIndex:     record.WinnerIndex,
	}
	if record.IsChronoWarrior {
		result.Source = ChronoWarrior
	}

	if record.Withdrawal != nil {
		if !record.Claimed {
			return UserRaffleEthDeposit{}, errors.New("unclaimed deposit carries a withdrawal")
		}
		withdrawal, err := mapUserDepositWithdrawal(*record.Withdrawal)
		if err != nil {
			return UserRaffleEthDeposit{}, err
		}
		result.Withdrawal = &withdrawal
	}
	return result, nil
}

func mapUserDepositWithdrawal(record cgstore.UserDepositWithdrawalRecord) (UserDepositWithdrawal, error) {
	if record.EventLogID < 1 {
		return UserDepositWithdrawal{}, errors.New("invalid withdrawal event id")
	}
	if !isTransactionHash(record.TxHash) {
		return UserDepositWithdrawal{}, errors.New("invalid withdrawal transaction hash")
	}
	occurredAt, err := time.Parse(time.RFC3339Nano, record.DateTime)
	if err != nil {
		return UserDepositWithdrawal{}, fmt.Errorf("parse withdrawal timestamp: %w", err)
	}
	if !ethcommon.IsHexAddress(record.BeneficiaryAddr) {
		return UserDepositWithdrawal{}, errors.New("invalid withdrawal beneficiary address")
	}
	return UserDepositWithdrawal{
		BeneficiaryAddress: ethcommon.HexToAddress(record.BeneficiaryAddr).Hex(),
		EventLogId:         record.EventLogID,
		OccurredAt:         occurredAt.UTC(),
		TransactionHash:    strings.ToLower(record.TxHash),
	}, nil
}

func mapUserRaffleNftWin(record cgstore.UserRaffleNftWinRecord) (UserRaffleNftWin, error) {
	if record.RoundNum < 0 || record.WinnerIndex < 0 || record.TokenID < 0 {
		return UserRaffleNftWin{}, errors.New("invalid raffle NFT win identity")
	}
	if record.IsRWalk && !record.IsStaker {
		return UserRaffleNftWin{}, errors.New("RandomWalk raffle win without the staker flag")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserRaffleNftWin{}, fmt.Errorf("raffle NFT win transaction: %w", err)
	}
	cstAmount, err := requiredAmount(record.CstAmountWei)
	if err != nil {
		return UserRaffleNftWin{}, fmt.Errorf("raffle NFT win CST amount: %w", err)
	}
	return UserRaffleNftWin{
		BlockNumber:     transaction.BlockNumber,
		CstAmountWei:    cstAmount,
		EventLogId:      transaction.EventLogId,
		IsRandomWalk:    record.IsRWalk,
		IsStaker:        record.IsStaker,
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		TransactionHash: transaction.TransactionHash,
		WinnerIndex:     record.WinnerIndex,
	}, nil
}

func mapUserDonatedNft(record cgstore.UserDonatedNftRecord) (UserDonatedNft, error) {
	if record.RoundNum < 0 || record.TokenID < 0 || record.DonationIndex < 0 {
		return UserDonatedNft{}, errors.New("invalid donated NFT identity")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserDonatedNft{}, fmt.Errorf("donated NFT transaction: %w", err)
	}
	if !ethcommon.IsHexAddress(record.DonorAddr) {
		return UserDonatedNft{}, errors.New("invalid donated NFT donor address")
	}
	if !ethcommon.IsHexAddress(record.TokenAddr) {
		return UserDonatedNft{}, errors.New("invalid donated NFT token address")
	}
	if record.Claimed != (record.Claim != nil) {
		return UserDonatedNft{}, errors.New("donated NFT claim state is inconsistent")
	}

	result := UserDonatedNft{
		BlockNumber:     transaction.BlockNumber,
		Claimed:         record.Claimed,
		DonationIndex:   record.DonationIndex,
		DonorAddress:    ethcommon.HexToAddress(record.DonorAddr).Hex(),
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		TokenAddress:    ethcommon.HexToAddress(record.TokenAddr).Hex(),
		TokenId:         record.TokenID,
		TokenUri:        record.TokenURI,
		TransactionHash: transaction.TransactionHash,
	}
	if record.Claim != nil {
		claim, err := mapUserDonatedNftClaim(*record.Claim)
		if err != nil {
			return UserDonatedNft{}, err
		}
		result.Claim = &claim
	}
	return result, nil
}

func mapUserDonatedNftClaim(record cgstore.UserDonatedNftClaimRecord) (UserDonatedNftClaim, error) {
	if record.EventLogID < 1 {
		return UserDonatedNftClaim{}, errors.New("invalid donated NFT claim event id")
	}
	if !isTransactionHash(record.TxHash) {
		return UserDonatedNftClaim{}, errors.New("invalid donated NFT claim transaction hash")
	}
	occurredAt, err := time.Parse(time.RFC3339Nano, record.DateTime)
	if err != nil {
		return UserDonatedNftClaim{}, fmt.Errorf("parse donated NFT claim timestamp: %w", err)
	}
	if !ethcommon.IsHexAddress(record.ClaimerAddr) {
		return UserDonatedNftClaim{}, errors.New("invalid donated NFT claimer address")
	}
	return UserDonatedNftClaim{
		ClaimerAddress:  ethcommon.HexToAddress(record.ClaimerAddr).Hex(),
		EventLogId:      record.EventLogID,
		OccurredAt:      occurredAt.UTC(),
		TransactionHash: strings.ToLower(record.TxHash),
	}, nil
}

func mapUserDonatedErc20(record cgstore.UserDonatedErc20Record) (UserDonatedErc20, error) {
	if record.RoundNum < 0 {
		return UserDonatedErc20{}, errors.New("invalid donated ERC-20 round")
	}
	if !ethcommon.IsHexAddress(record.TokenAddr) {
		return UserDonatedErc20{}, errors.New("invalid donated ERC-20 token address")
	}
	donated, err := requiredAmount(record.DonatedBaseUnits)
	if err != nil {
		return UserDonatedErc20{}, fmt.Errorf("donated ERC-20 total: %w", err)
	}
	claimed, err := requiredAmount(record.ClaimedBaseUnits)
	if err != nil {
		return UserDonatedErc20{}, fmt.Errorf("donated ERC-20 claimed total: %w", err)
	}
	remaining, err := requiredAmount(record.RemainingBaseUnits)
	if err != nil {
		return UserDonatedErc20{}, fmt.Errorf("donated ERC-20 remaining total: %w", err)
	}
	if !donatedErc20TotalsConsistent(donated, claimed, remaining) {
		return UserDonatedErc20{}, errors.New("donated ERC-20 totals are inconsistent")
	}

	result := UserDonatedErc20{
		ClaimedBaseUnits:   claimed,
		DonatedBaseUnits:   donated,
		RemainingBaseUnits: remaining,
		Round:              record.RoundNum,
		TokenAddress:       ethcommon.HexToAddress(record.TokenAddr).Hex(),
	}
	if record.LastClaim != nil {
		claim, err := mapUserDonatedErc20Claim(*record.LastClaim)
		if err != nil {
			return UserDonatedErc20{}, err
		}
		result.LastClaim = &claim
	}
	return result, nil
}

// donatedErc20TotalsConsistent proves donated = claimed + remaining over the
// exact integer values, catching aggregation drift before it reaches a
// client.
func donatedErc20TotalsConsistent(donated, claimed, remaining string) bool {
	donatedInt, ok := new(big.Int).SetString(donated, 10)
	if !ok {
		return false
	}
	claimedInt, ok := new(big.Int).SetString(claimed, 10)
	if !ok {
		return false
	}
	remainingInt, ok := new(big.Int).SetString(remaining, 10)
	if !ok {
		return false
	}
	return donatedInt.Cmp(new(big.Int).Add(claimedInt, remainingInt)) == 0
}

func mapUserDonatedErc20Claim(record cgstore.UserDonatedErc20ClaimRecord) (UserDonatedErc20Claim, error) {
	if record.EventLogID < 1 {
		return UserDonatedErc20Claim{}, errors.New("invalid donated ERC-20 claim event id")
	}
	if !isTransactionHash(record.TxHash) {
		return UserDonatedErc20Claim{}, errors.New("invalid donated ERC-20 claim transaction hash")
	}
	occurredAt, err := time.Parse(time.RFC3339Nano, record.DateTime)
	if err != nil {
		return UserDonatedErc20Claim{}, fmt.Errorf("parse donated ERC-20 claim timestamp: %w", err)
	}
	if !ethcommon.IsHexAddress(record.ClaimerAddr) {
		return UserDonatedErc20Claim{}, errors.New("invalid donated ERC-20 claimer address")
	}
	amount, err := requiredAmount(record.AmountBaseUnits)
	if err != nil {
		return UserDonatedErc20Claim{}, fmt.Errorf("donated ERC-20 claim amount: %w", err)
	}
	return UserDonatedErc20Claim{
		AmountBaseUnits: amount,
		ClaimerAddress:  ethcommon.HexToAddress(record.ClaimerAddr).Hex(),
		EventLogId:      record.EventLogID,
		OccurredAt:      occurredAt.UTC(),
		TransactionHash: strings.ToLower(record.TxHash),
	}, nil
}
