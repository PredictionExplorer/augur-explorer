package v2

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	openapi_types "github.com/oapi-codegen/runtime/types"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapGlobalToken(record cgstore.GlobalTokenRecord) (CosmicSignatureToken, error) {
	if record.TokenID < 0 || record.MintRound < 0 {
		return CosmicSignatureToken{}, errors.New("invalid token identity")
	}
	if record.Seed == "" {
		return CosmicSignatureToken{}, errors.New("token misses its mint seed")
	}
	if !ethcommon.IsHexAddress(record.WinnerAddr) || !ethcommon.IsHexAddress(record.CurOwnerAddr) {
		return CosmicSignatureToken{}, errors.New("invalid token owner address")
	}
	mintType, err := mapCosmicSignatureMintType(record.MintSource)
	if err != nil {
		return CosmicSignatureToken{}, err
	}
	transaction, err := mapClaimTransaction(record.MintTx)
	if err != nil {
		return CosmicSignatureToken{}, fmt.Errorf("mint transaction: %w", err)
	}
	result := CosmicSignatureToken{
		BlockNumber:         transaction.BlockNumber,
		CurrentOwnerAddress: ethcommon.HexToAddress(record.CurOwnerAddr).Hex(),
		EventLogId:          transaction.EventLogId,
		MintRound:           record.MintRound,
		MintType:            mintType,
		MintedAt:            transaction.OccurredAt,
		NftTokenId:          record.TokenID,
		Seed:                record.Seed,
		Staked:              record.Staked,
		TransactionHash:     transaction.TransactionHash,
		WinnerAddress:       ethcommon.HexToAddress(record.WinnerAddr).Hex(),
	}
	if record.TokenName != "" {
		name := record.TokenName
		result.TokenName = &name
	}
	return result, nil
}

func mapGlobalTokenDetail(
	record cgstore.GlobalTokenDetailRecord,
) (CosmicGameCosmicSignatureTokenDetail, error) {
	token, err := mapGlobalToken(record.GlobalTokenRecord)
	if err != nil {
		return CosmicGameCosmicSignatureTokenDetail{}, err
	}
	result := CosmicGameCosmicSignatureTokenDetail{
		BlockNumber:         token.BlockNumber,
		CurrentOwnerAddress: token.CurrentOwnerAddress,
		EventLogId:          token.EventLogId,
		MintRound:           token.MintRound,
		MintType:            token.MintType,
		MintedAt:            token.MintedAt,
		NftTokenId:          token.NftTokenId,
		Seed:                token.Seed,
		Staked:              token.Staked,
		TokenName:           token.TokenName,
		TransactionHash:     token.TransactionHash,
		WinnerAddress:       token.WinnerAddress,
	}
	if record.Staked != (record.CurrentStake != nil) {
		return CosmicGameCosmicSignatureTokenDetail{},
			errors.New("staked flag disagrees with the stake action")
	}
	if record.CurrentStake != nil {
		stake := record.CurrentStake
		if stake.StakeActionID < 0 {
			return CosmicGameCosmicSignatureTokenDetail{}, errors.New("invalid stake action id")
		}
		if !ethcommon.IsHexAddress(stake.StakerAddr) {
			return CosmicGameCosmicSignatureTokenDetail{}, errors.New("invalid staker address")
		}
		stakedAt, err := time.Parse(time.RFC3339Nano, stake.StakedAtText)
		if err != nil {
			return CosmicGameCosmicSignatureTokenDetail{},
				fmt.Errorf("parse stake timestamp: %w", err)
		}
		result.CurrentStake = &CosmicSignatureTokenStake{
			StakeActionId: stake.StakeActionID,
			StakedAt:      stakedAt.UTC(),
			StakerAddress: ethcommon.HexToAddress(stake.StakerAddr).Hex(),
		}
	}
	return result, nil
}

func mapTokenNameChange(record cgstore.TokenNameChangeRecord) (TokenNameChange, error) {
	if record.TokenID < 0 {
		return TokenNameChange{}, errors.New("invalid rename token identity")
	}
	if !ethcommon.IsHexAddress(record.ChangedBy) {
		return TokenNameChange{}, errors.New("invalid rename author address")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return TokenNameChange{}, fmt.Errorf("rename transaction: %w", err)
	}
	return TokenNameChange{
		BlockNumber:      transaction.BlockNumber,
		ChangedByAddress: ethcommon.HexToAddress(record.ChangedBy).Hex(),
		EventLogId:       transaction.EventLogId,
		NftTokenId:       record.TokenID,
		OccurredAt:       transaction.OccurredAt,
		TokenName:        record.NewName,
		TransactionHash:  transaction.TransactionHash,
	}, nil
}

func mapTokenTransfer(record cgstore.TokenTransferRecord) (CosmicSignatureTokenTransfer, error) {
	if record.TokenID < 0 {
		return CosmicSignatureTokenTransfer{}, errors.New("invalid transfer token identity")
	}
	if !ethcommon.IsHexAddress(record.FromAddr) || !ethcommon.IsHexAddress(record.ToAddr) {
		return CosmicSignatureTokenTransfer{}, errors.New("invalid transfer counterparty address")
	}
	transferType, err := mapTokenTransferType(record.TransferType)
	if err != nil {
		return CosmicSignatureTokenTransfer{}, err
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return CosmicSignatureTokenTransfer{}, fmt.Errorf("transfer transaction: %w", err)
	}
	return CosmicSignatureTokenTransfer{
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		FromAddress:     ethcommon.HexToAddress(record.FromAddr).Hex(),
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		ToAddress:       ethcommon.HexToAddress(record.ToAddr).Hex(),
		TransactionHash: transaction.TransactionHash,
		TransferType:    transferType,
	}, nil
}

func mapCosmicSignatureHolder(
	record cgstore.CosmicSignatureHolderRecord,
) (CosmicSignatureHolder, error) {
	if record.TokenCount < 1 {
		return CosmicSignatureHolder{}, errors.New("holder without tokens")
	}
	if !ethcommon.IsHexAddress(record.Address) {
		return CosmicSignatureHolder{}, errors.New("invalid holder address")
	}
	return CosmicSignatureHolder{
		OwnerAddress: ethcommon.HexToAddress(record.Address).Hex(),
		TokenCount:   record.TokenCount,
	}, nil
}

func mapCosmicTokenHolder(record cgstore.CosmicTokenHolderRecord) (CosmicTokenHolder, error) {
	if !ethcommon.IsHexAddress(record.Address) {
		return CosmicTokenHolder{}, errors.New("invalid holder address")
	}
	balance, err := requiredAmount(record.BalanceWei)
	if err != nil {
		return CosmicTokenHolder{}, fmt.Errorf("holder balance: %w", err)
	}
	if balance == "0" {
		return CosmicTokenHolder{}, errors.New("holder without balance")
	}
	return CosmicTokenHolder{
		OwnerAddress: ethcommon.HexToAddress(record.Address).Hex(),
		BalanceWei:   balance,
	}, nil
}

// supplySharePercentage computes balance/supply as a canonical two-decimal
// percentage string from exact integers.
func supplySharePercentage(balanceWei, supplyWei string) (string, error) {
	balance, ok := new(big.Int).SetString(balanceWei, 10)
	if !ok || balance.Sign() < 0 {
		return "", fmt.Errorf("invalid balance %q", balanceWei)
	}
	supply, ok := new(big.Int).SetString(supplyWei, 10)
	if !ok || supply.Sign() <= 0 {
		return "", fmt.Errorf("invalid supply %q", supplyWei)
	}
	if balance.Cmp(supply) > 0 {
		return "", errors.New("holder balance exceeds total supply")
	}
	value := new(big.Rat).SetFrac(balance, supply)
	value.Mul(value, big.NewRat(100, 1))
	return canonicalDecimal(value.FloatString(2), false)
}

func mapCosmicTokenStatistics(
	record cgstore.CosmicTokenStatisticsRecord,
) (CosmicGameCosmicTokenStatistics, error) {
	if record.HolderCount < 0 {
		return CosmicGameCosmicTokenStatistics{}, errors.New("holder count is negative")
	}
	if record.TransferCount < 0 || record.MintCount < 0 || record.BurnCount < 0 ||
		record.MintCount+record.BurnCount > record.TransferCount {
		return CosmicGameCosmicTokenStatistics{}, errors.New("transfer counters are inconsistent")
	}
	supply, err := requiredAmount(record.TotalSupplyWei)
	if err != nil {
		return CosmicGameCosmicTokenStatistics{}, fmt.Errorf("total supply: %w", err)
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
			return CosmicGameCosmicTokenStatistics{}, fmt.Errorf("%s: %w", field.name, err)
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
		return CosmicGameCosmicTokenStatistics{}, errors.New("earning sources do not add up")
	}
	consumed, err := requiredAmount(record.ConsumedInBidsWei)
	if err != nil {
		return CosmicGameCosmicTokenStatistics{}, fmt.Errorf("consumed in bids: %w", err)
	}
	net, err := signedAmount(record.NetWei)
	if err != nil {
		return CosmicGameCosmicTokenStatistics{}, fmt.Errorf("net issuance: %w", err)
	}
	consumedInt, _ := new(big.Int).SetString(consumed, 10)
	if new(big.Int).Sub(sourceSum, consumedInt).String() != net {
		return CosmicGameCosmicTokenStatistics{},
			errors.New("net issuance diverges from earnings minus consumption")
	}

	if len(record.TopHolders) > 10 {
		return CosmicGameCosmicTokenStatistics{}, errors.New("too many top holders")
	}
	if record.HolderCount < int64(len(record.TopHolders)) {
		return CosmicGameCosmicTokenStatistics{},
			errors.New("holder count is below the top-holder list")
	}
	topHolders := make([]CosmicTokenTopHolder, 0, len(record.TopHolders))
	var previousBalance *big.Int
	previousAid := int64(0)
	for _, holder := range record.TopHolders {
		if !ethcommon.IsHexAddress(holder.Address) {
			return CosmicGameCosmicTokenStatistics{}, errors.New("invalid top-holder address")
		}
		balance, err := requiredAmount(holder.BalanceWei)
		if err != nil {
			return CosmicGameCosmicTokenStatistics{}, fmt.Errorf("top-holder balance: %w", err)
		}
		balanceInt, _ := new(big.Int).SetString(balance, 10)
		if balanceInt.Sign() <= 0 {
			return CosmicGameCosmicTokenStatistics{}, errors.New("top holder without balance")
		}
		if previousBalance != nil {
			comparison := balanceInt.Cmp(previousBalance)
			if comparison > 0 || (comparison == 0 && holder.OwnerAid <= previousAid) {
				return CosmicGameCosmicTokenStatistics{}, errors.New("top holders are unordered")
			}
		}
		share, err := supplySharePercentage(balance, supply)
		if err != nil {
			return CosmicGameCosmicTokenStatistics{}, fmt.Errorf("top-holder share: %w", err)
		}
		topHolders = append(topHolders, CosmicTokenTopHolder{
			BalanceWei:    balance,
			OwnerAddress:  ethcommon.HexToAddress(holder.Address).Hex(),
			ShareOfSupply: share,
		})
		previousBalance = balanceInt
		previousAid = holder.OwnerAid
	}

	return CosmicGameCosmicTokenStatistics{
		ConsumedInBidsWei: consumed,
		Earned:            earnings,
		HolderCount:       record.HolderCount,
		NetWei:            net,
		TopHolders:        topHolders,
		TotalSupplyWei:    supply,
		Transfers: UserCosmicTokenTransferCounts{
			BurnCount:  record.BurnCount,
			MintCount:  record.MintCount,
			TotalCount: record.TransferCount,
		},
	}, nil
}

func mapSupplyChange(record cgstore.SupplyChangeRecord) (CosmicTokenSupplyChange, error) {
	if !ethcommon.IsHexAddress(record.BidderAddr) {
		return CosmicTokenSupplyChange{}, errors.New("invalid supply-change bidder address")
	}
	minted, err := requiredAmount(record.MintedWei)
	if err != nil {
		return CosmicTokenSupplyChange{}, fmt.Errorf("minted amount: %w", err)
	}
	burned, err := requiredAmount(record.BurnedWei)
	if err != nil {
		return CosmicTokenSupplyChange{}, fmt.Errorf("burned amount: %w", err)
	}
	net, err := signedAmount(record.NetWei)
	if err != nil {
		return CosmicTokenSupplyChange{}, fmt.Errorf("net amount: %w", err)
	}
	mintedInt, _ := new(big.Int).SetString(minted, 10)
	burnedInt, _ := new(big.Int).SetString(burned, 10)
	if new(big.Int).Sub(mintedInt, burnedInt).String() != net {
		return CosmicTokenSupplyChange{}, errors.New("net diverges from minted minus burned")
	}
	total, err := requiredAmount(record.TotalSupplyWei)
	if err != nil {
		return CosmicTokenSupplyChange{}, fmt.Errorf("running supply: %w", err)
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return CosmicTokenSupplyChange{}, fmt.Errorf("supply-change transaction: %w", err)
	}
	return CosmicTokenSupplyChange{
		BidType:         mapBidType(record.BidType),
		BidderAddress:   ethcommon.HexToAddress(record.BidderAddr).Hex(),
		BlockNumber:     transaction.BlockNumber,
		BurnedWei:       burned,
		EventLogId:      transaction.EventLogId,
		MintedWei:       minted,
		NetWei:          net,
		OccurredAt:      transaction.OccurredAt,
		TotalSupplyWei:  total,
		TransactionHash: transaction.TransactionHash,
	}, nil
}

func mapDailySupply(record cgstore.DailySupplyRecord) (CosmicTokenDailySupply, error) {
	if record.BidCount < 1 {
		return CosmicTokenDailySupply{}, errors.New("daily row without bids")
	}
	day, err := time.Parse("2006-01-02", record.Date)
	if err != nil {
		return CosmicTokenDailySupply{}, fmt.Errorf("parse day: %w", err)
	}
	minted, err := requiredAmount(record.MintedWei)
	if err != nil {
		return CosmicTokenDailySupply{}, fmt.Errorf("minted amount: %w", err)
	}
	burned, err := requiredAmount(record.BurnedWei)
	if err != nil {
		return CosmicTokenDailySupply{}, fmt.Errorf("burned amount: %w", err)
	}
	net, err := signedAmount(record.NetWei)
	if err != nil {
		return CosmicTokenDailySupply{}, fmt.Errorf("net amount: %w", err)
	}
	mintedInt, _ := new(big.Int).SetString(minted, 10)
	burnedInt, _ := new(big.Int).SetString(burned, 10)
	if new(big.Int).Sub(mintedInt, burnedInt).String() != net {
		return CosmicTokenDailySupply{}, errors.New("net diverges from minted minus burned")
	}
	total, err := requiredAmount(record.TotalSupplyWei)
	if err != nil {
		return CosmicTokenDailySupply{}, fmt.Errorf("running supply: %w", err)
	}
	return CosmicTokenDailySupply{
		BidCount:       record.BidCount,
		BurnedWei:      burned,
		Date:           openapi_types.Date{Time: day},
		MintedWei:      minted,
		NetWei:         net,
		TotalSupplyWei: total,
	}, nil
}

func mapMarketingReward(record cgstore.MarketingRewardRecord) (MarketingReward, error) {
	if !ethcommon.IsHexAddress(record.MarketerAddr) {
		return MarketingReward{}, errors.New("invalid marketer address")
	}
	amount, err := requiredAmount(record.AmountWei)
	if err != nil {
		return MarketingReward{}, fmt.Errorf("marketing reward amount: %w", err)
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return MarketingReward{}, fmt.Errorf("marketing reward transaction: %w", err)
	}
	return MarketingReward{
		AmountWei:       amount,
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		MarketerAddress: ethcommon.HexToAddress(record.MarketerAddr).Hex(),
		OccurredAt:      transaction.OccurredAt,
		TransactionHash: transaction.TransactionHash,
	}, nil
}
