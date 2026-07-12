package v2

import (
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const userMappingAddress = "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

func TestMapUserProfile(t *testing.T) {
	t.Parallel()
	maxBid := "000120000000000000"
	record := validUserProfileRecord()
	record.MaxETHBidWei = &maxBid

	got, err := mapUserProfile(record)
	if err != nil {
		t.Fatalf("mapUserProfile: %v", err)
	}
	if got.Address != ethcommon.HexToAddress(userMappingAddress).Hex() {
		t.Fatalf("address = %q", got.Address)
	}
	if got.Bidding.BidCount != 5 || got.Bidding.MaxEthBidWei == nil ||
		*got.Bidding.MaxEthBidWei != "120000000000000" ||
		got.Bidding.TotalEthSpentWei != "430000000000000000" ||
		got.Bidding.TotalCstSpentWei != "999999999999999999999999999999999999" {
		t.Fatalf("bidding = %+v", got.Bidding)
	}
	if got.Prizes.PrizeCount != 7 || got.Prizes.CstPrizeCount != 2 ||
		got.Prizes.NftPrizeCount != 2 || got.Prizes.TotalEthWonWei != "640000000000000000" {
		t.Fatalf("prizes = %+v", got.Prizes)
	}
	if got.Raffles.EthPrizeCount != 1 || got.Raffles.NftPrizeCount != 2 ||
		got.Raffles.TotalCstWonWei != "200000000000000000000" {
		t.Fatalf("raffles = %+v", got.Raffles)
	}
	if got.EthDonations.DonationCount != 3 ||
		got.EthDonations.TotalDonatedWei != "90000000000000000" {
		t.Fatalf("donations = %+v", got.EthDonations)
	}
	if got.Transfers.CosmicTokenTransferCount != 4 ||
		got.Transfers.CosmicSignatureTransferCount != 5 {
		t.Fatalf("transfers = %+v", got.Transfers)
	}
	if got.CstStaking.StakedTokenCount != 2 ||
		got.CstStaking.TotalRewardWei != "1000000000000000000" ||
		got.RandomWalkStaking.MintedTokenCount != 4 {
		t.Fatalf("staking = %+v / %+v", got.CstStaking, got.RandomWalkStaking)
	}
}

func TestMapUserProfileWithoutETHBid(t *testing.T) {
	t.Parallel()
	record := validUserProfileRecord()
	record.MaxETHBidWei = nil
	got, err := mapUserProfile(record)
	if err != nil {
		t.Fatal(err)
	}
	if got.Bidding.MaxEthBidWei != nil {
		t.Fatalf("max ETH bid = %v, want omitted", got.Bidding.MaxEthBidWei)
	}
}

func TestZeroUserProfile(t *testing.T) {
	t.Parallel()
	address := ethcommon.HexToAddress(userMappingAddress).Hex()
	got := zeroUserProfile(address)
	if got.Address != address || got.Bidding.BidCount != 0 ||
		got.Bidding.MaxEthBidWei != nil ||
		got.Bidding.TotalEthSpentWei != "0" ||
		got.Prizes.TotalEthWonWei != "0" ||
		got.CstStaking.TotalRewardWei != "0" ||
		got.RandomWalkStaking.MintedTokenCount != 0 {
		t.Fatalf("zero profile = %+v", got)
	}
}

func TestMapUserProfileRejectsMalformedRecords(t *testing.T) {
	t.Parallel()
	tests := map[string]func(*cgstore.UserProfileRecord){
		"address": func(record *cgstore.UserProfileRecord) {
			record.Address = "bad"
		},
		"bid count": func(record *cgstore.UserProfileRecord) {
			record.BidCount = -1
		},
		"prize count": func(record *cgstore.UserProfileRecord) {
			record.PrizeCount = -1
		},
		"raffle count": func(record *cgstore.UserProfileRecord) {
			record.RaffleNFTPrizeCount = -1
		},
		"donation count": func(record *cgstore.UserProfileRecord) {
			record.ETHDonationCount = -1
		},
		"transfer count": func(record *cgstore.UserProfileRecord) {
			record.CosmicTokenTransferCount = -1
		},
		"bid amount": func(record *cgstore.UserProfileRecord) {
			record.TotalETHSpentWei = "-1"
		},
		"maximum bid": func(record *cgstore.UserProfileRecord) {
			value := "bad"
			record.MaxETHBidWei = &value
		},
		"prize amount": func(record *cgstore.UserProfileRecord) {
			record.TotalETHWonWei = "1.5"
		},
		"raffle amount": func(record *cgstore.UserProfileRecord) {
			record.RaffleCSTTotalWei = ""
		},
		"donation amount": func(record *cgstore.UserProfileRecord) {
			record.ETHDonatedWei = "bad"
		},
		"CST staking count": func(record *cgstore.UserProfileRecord) {
			record.CSTStakeActionCount = -1
		},
		"CST staking amount": func(record *cgstore.UserProfileRecord) {
			record.CSTTotalRewardWei = "-1"
		},
		"RandomWalk staking count": func(record *cgstore.UserProfileRecord) {
			record.RandomWalkMintedTokenCount = -1
		},
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validUserProfileRecord()
			mutate(&record)
			if _, err := mapUserProfile(record); err == nil {
				t.Fatal("malformed profile was accepted")
			}
		})
	}
}

func validUserProfileRecord() cgstore.UserProfileRecord {
	return cgstore.UserProfileRecord{
		Address:                      userMappingAddress,
		BidCount:                     5,
		TotalETHSpentWei:             "430000000000000000",
		TotalCSTSpentWei:             "999999999999999999999999999999999999",
		PrizeCount:                   7,
		MaxMainPrizeETHWei:           "500000000000000000",
		TotalETHWonWei:               "640000000000000000",
		CSTPrizeCount:                2,
		NFTPrizeCount:                2,
		UnclaimedNFTCount:            1,
		RaffleETHPrizeCount:          1,
		RaffleETHTotalWei:            "140000000000000000",
		RaffleNFTPrizeCount:          2,
		RaffleCSTTotalWei:            "200000000000000000000",
		ETHDonationCount:             3,
		ETHDonatedWei:                "90000000000000000",
		CosmicTokenTransferCount:     4,
		CosmicSignatureTransferCount: 5,
		CSTStakedTokenCount:          2,
		CSTStakeActionCount:          3,
		CSTUnstakeActionCount:        1,
		CSTTotalRewardWei:            "1000000000000000000",
		CSTUnclaimedRewardWei:        "250000000000000000",
		RandomWalkStakedTokenCount:   3,
		RandomWalkStakeActionCount:   4,
		RandomWalkUnstakeActionCount: 1,
		RandomWalkMintedTokenCount:   4,
	}
}
