package v2

import (
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

type roundPrizeAsset uint8

const (
	roundPrizeAssetEth roundPrizeAsset = iota + 1
	roundPrizeAssetCst
	roundPrizeAssetNft
)

func mapRoundPrize(record cgmodel.CGPrizeHistory) (RoundPrize, error) {
	if record.RoundNum < 0 || record.WinnerIndex < 0 {
		return RoundPrize{}, errors.New("invalid round-prize identity")
	}
	transaction, err := mapClaimTransaction(record.Tx)
	if err != nil {
		return RoundPrize{}, fmt.Errorf("prize transaction: %w", err)
	}
	prizeType, asset, err := mapRoundPrizeType(record.RecordType)
	if err != nil {
		return RoundPrize{}, err
	}

	result := RoundPrize{
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogId,
		OccurredAt:      transaction.OccurredAt,
		Round:           record.RoundNum,
		TransactionHash: transaction.TransactionHash,
		Type:            prizeType,
		WinnerIndex:     record.WinnerIndex,
	}

	if prizeType != CosmicSignatureStakingEth {
		if !ethcommon.IsHexAddress(record.WinnerAddr) {
			return RoundPrize{}, errors.New("invalid round-prize winner address")
		}
		winner := ethcommon.HexToAddress(record.WinnerAddr).Hex()
		result.WinnerAddress = &winner
	} else if record.WinnerAddr != "(All CS NFT Stakers)" {
		return RoundPrize{}, errors.New("round-wide staking prize has an invalid beneficiary")
	}

	switch asset {
	case roundPrizeAssetEth:
		amount, err := requiredAmount(record.Amount)
		if err != nil {
			return RoundPrize{}, fmt.Errorf("round-prize ETH amount: %w", err)
		}
		if record.TokenId != -1 {
			return RoundPrize{}, errors.New("ETH prize has an NFT token id")
		}
		result.EthAmountWei = &amount
	case roundPrizeAssetCst:
		amount, err := requiredAmount(record.Amount)
		if err != nil {
			return RoundPrize{}, fmt.Errorf("round-prize CST amount: %w", err)
		}
		if record.TokenId != -1 {
			return RoundPrize{}, errors.New("CST prize has an NFT token id")
		}
		result.CstAmountWei = &amount
	case roundPrizeAssetNft:
		if record.TokenId < 0 {
			return RoundPrize{}, errors.New("NFT prize has no token id")
		}
		amount, err := amountOrZero(record.Amount)
		if err != nil || amount != "0" {
			return RoundPrize{}, errors.New("NFT prize has a fungible amount")
		}
		tokenID := record.TokenId
		result.NftTokenId = &tokenID
	default:
		return RoundPrize{}, errors.New("unsupported round-prize asset")
	}
	return result, nil
}

func mapRoundPrizeType(recordType int64) (RoundPrizeType, roundPrizeAsset, error) {
	switch recordType {
	case 0:
		return MainPrizeEth, roundPrizeAssetEth, nil
	case 1:
		return MainPrizeCst, roundPrizeAssetCst, nil
	case 2:
		return MainPrizeNft, roundPrizeAssetNft, nil
	case 3:
		return LastCstBidderNft, roundPrizeAssetNft, nil
	case 4:
		return LastCstBidderCst, roundPrizeAssetCst, nil
	case 5:
		return EnduranceChampionNft, roundPrizeAssetNft, nil
	case 6:
		return EnduranceChampionCst, roundPrizeAssetCst, nil
	case 7:
		return ChronoWarriorEth, roundPrizeAssetEth, nil
	case 8:
		return ChronoWarriorCst, roundPrizeAssetCst, nil
	case 9:
		return ChronoWarriorNft, roundPrizeAssetNft, nil
	case 10:
		return BidderRaffleEth, roundPrizeAssetEth, nil
	case 11:
		return BidderRaffleCst, roundPrizeAssetCst, nil
	case 12:
		return BidderRaffleNft, roundPrizeAssetNft, nil
	case 13:
		return RandomWalkStakerRaffleCst, roundPrizeAssetCst, nil
	case 14:
		return RandomWalkStakerRaffleNft, roundPrizeAssetNft, nil
	case 15:
		return CosmicSignatureStakingEth, roundPrizeAssetEth, nil
	default:
		return "", 0, fmt.Errorf("unknown round-prize type %d", recordType)
	}
}
