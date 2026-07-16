package v2

import (
	"errors"
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapCosmicSignatureMintType(source cgstore.CosmicSignatureMintSource) (CosmicSignatureMintType, error) {
	switch source {
	case cgstore.MintSourceMainPrize:
		return CosmicSignatureMintTypeMainPrize, nil
	case cgstore.MintSourceBidderRaffle:
		return CosmicSignatureMintTypeBidderRaffle, nil
	case cgstore.MintSourceRandomWalkStaker:
		return CosmicSignatureMintTypeRandomWalkStakerRaffle, nil
	case cgstore.MintSourceCosmicSigStaker:
		return CosmicSignatureMintTypeCosmicSignatureStakerRaffle, nil
	case cgstore.MintSourceEnduranceChampion:
		return CosmicSignatureMintTypeEnduranceChampion, nil
	case cgstore.MintSourceLastCstBidder:
		return CosmicSignatureMintTypeLastCstBidder, nil
	case cgstore.MintSourceChronoWarriorPrize:
		return CosmicSignatureMintTypeChronoWarrior, nil
	default:
		return "", fmt.Errorf("unknown mint source %q", source)
	}
}

func mapUserOwnedToken(record cgstore.UserOwnedTokenRecord) (UserCosmicSignatureToken, error) {
	if record.TokenID < 0 || record.MintRound < 0 {
		return UserCosmicSignatureToken{}, errors.New("invalid owned token identity")
	}
	if record.Seed == "" {
		return UserCosmicSignatureToken{}, errors.New("owned token misses its mint seed")
	}
	if !ethcommon.IsHexAddress(record.WinnerAddr) {
		return UserCosmicSignatureToken{}, errors.New("invalid owned token winner address")
	}
	mintType, err := mapCosmicSignatureMintType(record.MintSource)
	if err != nil {
		return UserCosmicSignatureToken{}, err
	}
	transaction, err := mapClaimTransaction(record.MintTx)
	if err != nil {
		return UserCosmicSignatureToken{}, fmt.Errorf("mint transaction: %w", err)
	}
	result := UserCosmicSignatureToken{
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		MintRound:       record.MintRound,
		MintType:        mintType,
		MintedAt:        transaction.OccurredAt,
		NftTokenId:      record.TokenID,
		Seed:            record.Seed,
		Staked:          record.Staked,
		TransactionHash: transaction.TransactionHash,
		WinnerAddress:   ethcommon.HexToAddress(record.WinnerAddr).Hex(),
	}
	if record.TokenName != "" {
		name := record.TokenName
		result.TokenName = &name
	}
	return result, nil
}

func mapTokenTransferType(otype int64) (TokenTransferType, error) {
	switch otype {
	case 0:
		return TokenTransferTypeTransfer, nil
	case 1:
		return TokenTransferTypeMint, nil
	case 2:
		return TokenTransferTypeBurn, nil
	default:
		return "", fmt.Errorf("unknown transfer type %d", otype)
	}
}

func mapTokenTransferDirection(direction cgstore.UserTransferDirection) (TokenTransferDirection, error) {
	switch direction {
	case cgstore.UserTransferIn:
		return In, nil
	case cgstore.UserTransferOut:
		return Out, nil
	case cgstore.UserTransferSelf:
		return Self, nil
	default:
		return "", fmt.Errorf("unknown transfer direction %q", direction)
	}
}

func mapUserCosmicSignatureTransfer(
	record cgstore.UserCosmicSignatureTransferRecord,
) (UserCosmicSignatureTransfer, error) {
	if record.TokenID < 0 {
		return UserCosmicSignatureTransfer{}, errors.New("invalid transfer token identity")
	}
	if !ethcommon.IsHexAddress(record.FromAddr) || !ethcommon.IsHexAddress(record.ToAddr) {
		return UserCosmicSignatureTransfer{}, errors.New("invalid transfer counterparty address")
	}
	transferType, err := mapTokenTransferType(record.TransferType)
	if err != nil {
		return UserCosmicSignatureTransfer{}, err
	}
	direction, err := mapTokenTransferDirection(record.Direction)
	if err != nil {
		return UserCosmicSignatureTransfer{}, err
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserCosmicSignatureTransfer{}, fmt.Errorf("transfer transaction: %w", err)
	}
	return UserCosmicSignatureTransfer{
		BlockNumber:     transaction.BlockNumber,
		Direction:       direction,
		EventLogId:      transaction.EventLogId,
		FromAddress:     ethcommon.HexToAddress(record.FromAddr).Hex(),
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		ToAddress:       ethcommon.HexToAddress(record.ToAddr).Hex(),
		TransactionHash: transaction.TransactionHash,
		TransferType:    transferType,
	}, nil
}

func mapUserCosmicTokenTransfer(
	record cgstore.UserCosmicTokenTransferRecord,
) (UserCosmicTokenTransfer, error) {
	if !ethcommon.IsHexAddress(record.FromAddr) || !ethcommon.IsHexAddress(record.ToAddr) {
		return UserCosmicTokenTransfer{}, errors.New("invalid transfer counterparty address")
	}
	amount, err := requiredAmount(record.AmountWei)
	if err != nil {
		return UserCosmicTokenTransfer{}, fmt.Errorf("transfer amount: %w", err)
	}
	transferType, err := mapTokenTransferType(record.TransferType)
	if err != nil {
		return UserCosmicTokenTransfer{}, err
	}
	direction, err := mapTokenTransferDirection(record.Direction)
	if err != nil {
		return UserCosmicTokenTransfer{}, err
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserCosmicTokenTransfer{}, fmt.Errorf("transfer transaction: %w", err)
	}
	return UserCosmicTokenTransfer{
		AmountWei:       amount,
		BlockNumber:     transaction.BlockNumber,
		Direction:       direction,
		EventLogId:      transaction.EventLogId,
		FromAddress:     ethcommon.HexToAddress(record.FromAddr).Hex(),
		OccurredAt:      transaction.OccurredAt,
		ToAddress:       ethcommon.HexToAddress(record.ToAddr).Hex(),
		TransactionHash: transaction.TransactionHash,
		TransferType:    transferType,
	}, nil
}

func mapUserMarketingReward(record cgstore.UserMarketingRewardRecord) (UserMarketingReward, error) {
	amount, err := requiredAmount(record.AmountWei)
	if err != nil {
		return UserMarketingReward{}, fmt.Errorf("marketing reward amount: %w", err)
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return UserMarketingReward{}, fmt.Errorf("marketing reward transaction: %w", err)
	}
	return UserMarketingReward{
		AmountWei:       amount,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		TransactionHash: transaction.TransactionHash,
	}, nil
}

// signedAmount validates an exact signed decimal string (netWei may be
// negative when a wallet consumed more CST in bids than it earned).
func signedAmount(value string) (string, error) {
	if value == "" {
		return "", errors.New("amount is empty")
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return "", fmt.Errorf("invalid decimal %q", value)
	}
	return amount.String(), nil
}

func mapUserCosmicTokenSummary(
	address string,
	record cgstore.UserCosmicTokenSummaryRecord,
) (CosmicGameUserCosmicTokenSummary, error) {
	if record.TransferCount < 0 || record.MintCount < 0 || record.BurnCount < 0 ||
		record.MintCount+record.BurnCount > record.TransferCount {
		return CosmicGameUserCosmicTokenSummary{}, errors.New("transfer counters are inconsistent")
	}
	balance, err := requiredAmount(record.BalanceWei)
	if err != nil {
		return CosmicGameUserCosmicTokenSummary{}, fmt.Errorf("balance: %w", err)
	}
	earnings := UserCosmicTokenEarnings{}
	for _, field := range []struct {
		name   string
		source string
		target *string
	}{
		{"bidding rewards", record.BiddingRewardsWei, &earnings.BiddingRewardsWei},
		{"main prizes", record.MainPrizesWei, &earnings.MainPrizesWei},
		{"raffle prizes", record.RafflePrizesWei, &earnings.RafflePrizesWei},
		{"chrono warrior prizes", record.ChronoWarriorPrizesWei, &earnings.ChronoWarriorPrizesWei},
		{"endurance champion prizes", record.EnduranceChampionPrizesWei, &earnings.EnduranceChampionPrizesWei},
		{"last CST bidder prizes", record.LastCstBidderPrizesWei, &earnings.LastCstBidderPrizesWei},
		{"marketing rewards", record.MarketingRewardsWei, &earnings.MarketingRewardsWei},
		{"total earned", record.TotalEarnedWei, &earnings.TotalWei},
	} {
		amount, err := requiredAmount(field.source)
		if err != nil {
			return CosmicGameUserCosmicTokenSummary{}, fmt.Errorf("%s: %w", field.name, err)
		}
		*field.target = amount
	}
	sourceSum := new(big.Int)
	for _, source := range []string{
		earnings.BiddingRewardsWei,
		earnings.MainPrizesWei,
		earnings.RafflePrizesWei,
		earnings.ChronoWarriorPrizesWei,
		earnings.EnduranceChampionPrizesWei,
		earnings.LastCstBidderPrizesWei,
		earnings.MarketingRewardsWei,
	} {
		part, _ := new(big.Int).SetString(source, 10)
		sourceSum.Add(sourceSum, part)
	}
	if sourceSum.String() != earnings.TotalWei {
		return CosmicGameUserCosmicTokenSummary{}, errors.New("earning sources do not add up")
	}
	consumed, err := requiredAmount(record.ConsumedInBidsWei)
	if err != nil {
		return CosmicGameUserCosmicTokenSummary{}, fmt.Errorf("consumed in bids: %w", err)
	}
	net, err := signedAmount(record.NetWei)
	if err != nil {
		return CosmicGameUserCosmicTokenSummary{}, fmt.Errorf("net flow: %w", err)
	}
	consumedInt, _ := new(big.Int).SetString(consumed, 10)
	if new(big.Int).Sub(sourceSum, consumedInt).String() != net {
		return CosmicGameUserCosmicTokenSummary{}, errors.New("net flow diverges from earnings minus consumption")
	}
	return CosmicGameUserCosmicTokenSummary{
		Address:           address,
		BalanceWei:        balance,
		ConsumedInBidsWei: consumed,
		Earned:            earnings,
		NetWei:            net,
		Transfers: UserCosmicTokenTransferCounts{
			BurnCount:  record.BurnCount,
			MintCount:  record.MintCount,
			TotalCount: record.TransferCount,
		},
	}, nil
}

// zeroUserCosmicTokenSummary is the exact shape a valid but unindexed
// wallet receives.
func zeroUserCosmicTokenSummary(address string) CosmicGameUserCosmicTokenSummary {
	return CosmicGameUserCosmicTokenSummary{
		Address:           address,
		BalanceWei:        "0",
		ConsumedInBidsWei: "0",
		Earned: UserCosmicTokenEarnings{
			BiddingRewardsWei:          "0",
			MainPrizesWei:              "0",
			RafflePrizesWei:            "0",
			ChronoWarriorPrizesWei:     "0",
			EnduranceChampionPrizesWei: "0",
			LastCstBidderPrizesWei:     "0",
			MarketingRewardsWei:        "0",
			TotalWei:                   "0",
		},
		NetWei:    "0",
		Transfers: UserCosmicTokenTransferCounts{},
	}
}

func mapUserPendingWinnings(
	address string,
	record cgstore.UserPendingWinningsRecord,
) (CosmicGameUserPendingWinnings, error) {
	if record.DonatedNftCount < 0 || record.DonatedErc20TokenCount < 0 {
		return CosmicGameUserPendingWinnings{}, errors.New("pending winning counters are negative")
	}
	raffle, err := requiredAmount(record.RaffleEthWei)
	if err != nil {
		return CosmicGameUserPendingWinnings{}, fmt.Errorf("raffle ETH: %w", err)
	}
	chrono, err := requiredAmount(record.ChronoWarriorEthWei)
	if err != nil {
		return CosmicGameUserPendingWinnings{}, fmt.Errorf("chrono warrior ETH: %w", err)
	}
	staking, err := requiredAmount(record.StakingRewardWei)
	if err != nil {
		return CosmicGameUserPendingWinnings{}, fmt.Errorf("staking reward: %w", err)
	}
	return CosmicGameUserPendingWinnings{
		Address:                address,
		ChronoWarriorEthWei:    chrono,
		DonatedErc20TokenCount: record.DonatedErc20TokenCount,
		DonatedNftCount:        record.DonatedNftCount,
		RaffleEthWei:           raffle,
		StakingRewardWei:       staking,
	}, nil
}

// zeroUserPendingWinnings is the exact shape a valid but unindexed wallet
// receives.
func zeroUserPendingWinnings(address string) CosmicGameUserPendingWinnings {
	return CosmicGameUserPendingWinnings{
		Address:             address,
		ChronoWarriorEthWei: "0",
		RaffleEthWei:        "0",
		StakingRewardWei:    "0",
	}
}
