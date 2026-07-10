package v2

import (
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapRoundRaffleEthDeposit(
	record cgstore.RaffleEthDepositRecord,
) (RoundRaffleEthDeposit, error) {
	if record.RoundNum < 0 || record.WinnerIndex < 0 {
		return RoundRaffleEthDeposit{}, errors.New("invalid raffle ETH deposit identity")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return RoundRaffleEthDeposit{}, fmt.Errorf("raffle ETH deposit transaction: %w", err)
	}
	if !ethcommon.IsHexAddress(record.WinnerAddr) {
		return RoundRaffleEthDeposit{}, errors.New("invalid raffle ETH deposit winner address")
	}
	amount, err := requiredAmount(record.EthAmountWei)
	if err != nil {
		return RoundRaffleEthDeposit{}, fmt.Errorf("raffle ETH deposit amount: %w", err)
	}
	return RoundRaffleEthDeposit{
		BlockNumber:     transaction.BlockNumber,
		Claimed:         record.Claimed,
		EthAmountWei:    amount,
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		TransactionHash: transaction.TransactionHash,
		WinnerAddress:   ethcommon.HexToAddress(record.WinnerAddr).Hex(),
		WinnerIndex:     record.WinnerIndex,
	}, nil
}
