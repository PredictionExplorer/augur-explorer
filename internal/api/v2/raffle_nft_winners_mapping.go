package v2

import (
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

func mapRoundRaffleNftWinner(
	record cgprimitives.CGRaffleNFTWinnerRec,
	expectedStaker bool,
) (RoundRaffleNftWinner, error) {
	if record.RoundNum < 0 || record.WinnerIndex < 0 || record.TokenId < 0 {
		return RoundRaffleNftWinner{}, errors.New("invalid raffle NFT winner identity")
	}
	if record.IsStaker != expectedStaker {
		return RoundRaffleNftWinner{}, errors.New("raffle NFT winner belongs to another pool")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return RoundRaffleNftWinner{}, fmt.Errorf("raffle NFT winner transaction: %w", err)
	}
	if !ethcommon.IsHexAddress(record.WinnerAddr) {
		return RoundRaffleNftWinner{}, errors.New("invalid raffle NFT winner address")
	}
	cstAmount, err := requiredAmount(record.CstAmount)
	if err != nil {
		return RoundRaffleNftWinner{}, fmt.Errorf("raffle NFT winner CST amount: %w", err)
	}
	return RoundRaffleNftWinner{
		BlockNumber:     transaction.BlockNumber,
		CstAmountWei:    cstAmount,
		EventLogId:      transaction.EventLogId,
		IsRandomWalk:    record.IsRWalk,
		NftTokenId:      record.TokenId,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		TransactionHash: transaction.TransactionHash,
		WinnerAddress:   ethcommon.HexToAddress(record.WinnerAddr).Hex(),
		WinnerIndex:     record.WinnerIndex,
	}, nil
}
