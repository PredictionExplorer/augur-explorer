package v2

import (
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func mapUserProfile(record cgstore.UserProfileRecord) (CosmicGameUserProfile, error) {
	if !ethcommon.IsHexAddress(record.Address) {
		return CosmicGameUserProfile{}, errors.New("invalid user address")
	}
	if err := nonNegativeParticipantCounts(map[string]int64{
		"bid":                      record.BidCount,
		"prize":                    record.PrizeCount,
		"CST prize":                record.CSTPrizeCount,
		"NFT prize":                record.NFTPrizeCount,
		"unclaimed NFT":            record.UnclaimedNFTCount,
		"raffle ETH prize":         record.RaffleETHPrizeCount,
		"raffle NFT prize":         record.RaffleNFTPrizeCount,
		"ETH donation":             record.ETHDonationCount,
		"CosmicToken transfer":     record.CosmicTokenTransferCount,
		"CosmicSignature transfer": record.CosmicSignatureTransferCount,
	}); err != nil {
		return CosmicGameUserProfile{}, err
	}

	totalETHSpent, err := userProfileAmount("total ETH spent", record.TotalETHSpentWei)
	if err != nil {
		return CosmicGameUserProfile{}, err
	}
	totalCSTSpent, err := userProfileAmount("total CST spent", record.TotalCSTSpentWei)
	if err != nil {
		return CosmicGameUserProfile{}, err
	}
	maxMainPrize, err := userProfileAmount("maximum main prize", record.MaxMainPrizeETHWei)
	if err != nil {
		return CosmicGameUserProfile{}, err
	}
	totalETHWon, err := userProfileAmount("total ETH won", record.TotalETHWonWei)
	if err != nil {
		return CosmicGameUserProfile{}, err
	}
	raffleETH, err := userProfileAmount("raffle ETH won", record.RaffleETHTotalWei)
	if err != nil {
		return CosmicGameUserProfile{}, err
	}
	raffleCST, err := userProfileAmount("raffle CST won", record.RaffleCSTTotalWei)
	if err != nil {
		return CosmicGameUserProfile{}, err
	}
	donatedETH, err := userProfileAmount("ETH donated", record.ETHDonatedWei)
	if err != nil {
		return CosmicGameUserProfile{}, err
	}

	var maxBid *string
	if record.MaxETHBidWei != nil {
		value, err := userProfileAmount("maximum ETH bid", *record.MaxETHBidWei)
		if err != nil {
			return CosmicGameUserProfile{}, err
		}
		maxBid = &value
	}

	cstStaking, err := mapCSTParticipantStats(
		record.CSTStakedTokenCount,
		record.CSTStakeActionCount,
		record.CSTUnstakeActionCount,
		record.CSTTotalRewardWei,
		record.CSTUnclaimedRewardWei,
	)
	if err != nil {
		return CosmicGameUserProfile{}, fmt.Errorf("CST staking: %w", err)
	}
	randomWalkStaking, err := mapRandomWalkParticipantStats(
		record.RandomWalkStakedTokenCount,
		record.RandomWalkStakeActionCount,
		record.RandomWalkUnstakeActionCount,
		record.RandomWalkMintedTokenCount,
	)
	if err != nil {
		return CosmicGameUserProfile{}, fmt.Errorf("RandomWalk staking: %w", err)
	}

	return CosmicGameUserProfile{
		Address: ethcommon.HexToAddress(record.Address).Hex(),
		Bidding: UserBiddingStats{
			BidCount:         record.BidCount,
			MaxEthBidWei:     maxBid,
			TotalCstSpentWei: totalCSTSpent,
			TotalEthSpentWei: totalETHSpent,
		},
		CstStaking: cstStaking,
		EthDonations: UserEthDonationStats{
			DonationCount:   record.ETHDonationCount,
			TotalDonatedWei: donatedETH,
		},
		Prizes: UserPrizeStats{
			CstPrizeCount:      record.CSTPrizeCount,
			MaxMainPrizeEthWei: maxMainPrize,
			NftPrizeCount:      record.NFTPrizeCount,
			PrizeCount:         record.PrizeCount,
			TotalEthWonWei:     totalETHWon,
			UnclaimedNftCount:  record.UnclaimedNFTCount,
		},
		Raffles: UserRaffleStats{
			EthPrizeCount:  record.RaffleETHPrizeCount,
			NftPrizeCount:  record.RaffleNFTPrizeCount,
			TotalCstWonWei: raffleCST,
			TotalEthWonWei: raffleETH,
		},
		RandomWalkStaking: randomWalkStaking,
		Transfers: UserTransferStats{
			CosmicSignatureTransferCount: record.CosmicSignatureTransferCount,
			CosmicTokenTransferCount:     record.CosmicTokenTransferCount,
		},
	}, nil
}

func zeroUserProfile(address string) CosmicGameUserProfile {
	return CosmicGameUserProfile{
		Address: address,
		Bidding: UserBiddingStats{
			TotalCstSpentWei: "0",
			TotalEthSpentWei: "0",
		},
		CstStaking: CstStakingParticipantStats{
			TotalRewardWei:     "0",
			UnclaimedRewardWei: "0",
		},
		EthDonations: UserEthDonationStats{TotalDonatedWei: "0"},
		Prizes: UserPrizeStats{
			MaxMainPrizeEthWei: "0",
			TotalEthWonWei:     "0",
		},
		Raffles: UserRaffleStats{
			TotalCstWonWei: "0",
			TotalEthWonWei: "0",
		},
		RandomWalkStaking: RandomWalkStakingParticipantStats{},
		Transfers:         UserTransferStats{},
	}
}

func userProfileAmount(name, value string) (string, error) {
	amount, err := requiredAmount(value)
	if err != nil {
		return "", fmt.Errorf("%s: %w", name, err)
	}
	return amount, nil
}
