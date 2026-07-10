package v2

import (
	"testing"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const participantTestAddress = "0x2100000000000000000000000000000000000021"

func TestParticipantMappings(t *testing.T) {
	t.Parallel()
	bidder, err := mapBidderParticipant(cgstore.BidderParticipantRecord{
		BidderAid: 21, Address: participantTestAddress, BidCount: 3, MaxBidWei: "-1",
	})
	if err != nil || bidder.BidderAddress != participantTestAddress ||
		bidder.BidCount != 3 || bidder.MaxBidWei != nil {
		t.Fatalf("bidder = %+v, %v", bidder, err)
	}

	winner, err := mapWinnerParticipant(cgstore.WinnerParticipantRecord{
		WinnerAid: 21, Address: participantTestAddress, PrizeCount: 7,
		MaxMainPrizeETHWei: "0500", TotalETHWonWei: "640", CSTPrizeCount: 2,
		NFTPrizeCount: 2, UnclaimedNFTCount: 1, TotalETHSpentWei: "430",
	})
	if err != nil || winner.MaxMainPrizeEthWei != "500" || winner.TotalEthWonWei != "640" ||
		winner.TotalEthSpentWei != "430" || winner.UnclaimedNftCount != 1 {
		t.Fatalf("winner = %+v, %v", winner, err)
	}

	donor, err := mapDonorParticipant(cgstore.DonorParticipantRecord{
		DonorAid: 21, Address: participantTestAddress,
		DonationCount: 2, TotalDonatedWei: "999999999999999999999999999999999999",
	})
	if err != nil || donor.DonationCount != 2 ||
		donor.TotalDonatedWei != "999999999999999999999999999999999999" {
		t.Fatalf("donor = %+v, %v", donor, err)
	}

	cst, err := mapCSTStakerParticipant(cgstore.CSTStakerParticipantRecord{
		StakerAid: 21, Address: participantTestAddress, StakedTokenCount: 2,
		StakeActionCount: 3, UnstakeActionCount: 1,
		TotalRewardWei: "1000", UnclaimedRewardWei: "500",
	})
	if err != nil || cst.Staking.StakedTokenCount != 2 ||
		cst.Staking.TotalRewardWei != "1000" {
		t.Fatalf("CST staker = %+v, %v", cst, err)
	}

	randomWalk, err := mapRandomWalkStakerParticipant(cgstore.RandomWalkStakerParticipantRecord{
		StakerAid: 21, Address: participantTestAddress, StakedTokenCount: 2,
		StakeActionCount: 3, UnstakeActionCount: 1, MintedTokenCount: 4,
	})
	if err != nil || randomWalk.Staking.StakedTokenCount != 2 ||
		randomWalk.Staking.MintedTokenCount != 4 {
		t.Fatalf("RandomWalk staker = %+v, %v", randomWalk, err)
	}

	dual, err := mapDualStakerParticipant(cgstore.DualStakerParticipantRecord{
		StakerAid: 21, Address: participantTestAddress, TotalStakedTokenCount: 5,
		CSTStakedTokenCount: 2, CSTStakeActionCount: 3, CSTUnstakeActionCount: 1,
		CSTTotalRewardWei: "1000", CSTUnclaimedRewardWei: "500",
		RandomWalkStakedTokenCount: 3, RandomWalkStakeActionCount: 3,
		RandomWalkUnstakeActionCount: 0, RandomWalkMintedTokenCount: 1,
	})
	if err != nil || dual.TotalStakedTokenCount != 5 ||
		dual.Cst.StakedTokenCount != 2 || dual.RandomWalk.StakedTokenCount != 3 {
		t.Fatalf("dual staker = %+v, %v", dual, err)
	}
}

func TestParticipantMappingsRejectMalformedRecords(t *testing.T) {
	t.Parallel()
	tests := map[string]func() error{
		"bidder ID": func() error {
			_, err := mapBidderParticipant(cgstore.BidderParticipantRecord{
				Address: participantTestAddress, BidCount: 1, MaxBidWei: "1",
			})
			return err
		},
		"bidder address": func() error {
			_, err := mapBidderParticipant(cgstore.BidderParticipantRecord{
				BidderAid: 1, Address: "bad", BidCount: 1, MaxBidWei: "1",
			})
			return err
		},
		"bidder count": func() error {
			_, err := mapBidderParticipant(cgstore.BidderParticipantRecord{
				BidderAid: 1, Address: participantTestAddress, BidCount: 0, MaxBidWei: "1",
			})
			return err
		},
		"bidder amount": func() error {
			_, err := mapBidderParticipant(cgstore.BidderParticipantRecord{
				BidderAid: 1, Address: participantTestAddress, BidCount: 1, MaxBidWei: "bad",
			})
			return err
		},
		"winner count": func() error {
			record := validWinnerParticipantRecord()
			record.PrizeCount = 0
			_, err := mapWinnerParticipant(record)
			return err
		},
		"winner amount": func() error {
			record := validWinnerParticipantRecord()
			record.TotalETHWonWei = "-1"
			_, err := mapWinnerParticipant(record)
			return err
		},
		"donor count": func() error {
			_, err := mapDonorParticipant(cgstore.DonorParticipantRecord{
				DonorAid: 1, Address: participantTestAddress,
				DonationCount: 0, TotalDonatedWei: "1",
			})
			return err
		},
		"CST reward": func() error {
			_, err := mapCSTStakerParticipant(cgstore.CSTStakerParticipantRecord{
				StakerAid: 1, Address: participantTestAddress,
				StakeActionCount: 1, TotalRewardWei: "-1", UnclaimedRewardWei: "0",
			})
			return err
		},
		"CST action": func() error {
			_, err := mapCSTStakerParticipant(cgstore.CSTStakerParticipantRecord{
				StakerAid: 1, Address: participantTestAddress,
				TotalRewardWei: "0", UnclaimedRewardWei: "0",
			})
			return err
		},
		"RandomWalk count": func() error {
			_, err := mapRandomWalkStakerParticipant(cgstore.RandomWalkStakerParticipantRecord{
				StakerAid: 1, Address: participantTestAddress, StakeActionCount: 1,
				MintedTokenCount: -1,
			})
			return err
		},
		"RandomWalk action": func() error {
			_, err := mapRandomWalkStakerParticipant(cgstore.RandomWalkStakerParticipantRecord{
				StakerAid: 1, Address: participantTestAddress,
			})
			return err
		},
		"dual total": func() error {
			record := validDualStakerParticipantRecord()
			record.CSTStakedTokenCount = 0
			record.TotalStakedTokenCount = 1
			_, err := mapDualStakerParticipant(record)
			return err
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := test(); err == nil {
				t.Fatal("malformed record was accepted")
			}
		})
	}
}

func validWinnerParticipantRecord() cgstore.WinnerParticipantRecord {
	return cgstore.WinnerParticipantRecord{
		WinnerAid: 1, Address: participantTestAddress, PrizeCount: 1,
		MaxMainPrizeETHWei: "1", TotalETHWonWei: "1", TotalETHSpentWei: "0",
	}
}

func validDualStakerParticipantRecord() cgstore.DualStakerParticipantRecord {
	return cgstore.DualStakerParticipantRecord{
		StakerAid: 1, Address: participantTestAddress, TotalStakedTokenCount: 2,
		CSTStakedTokenCount: 1, CSTTotalRewardWei: "1", CSTUnclaimedRewardWei: "0",
		RandomWalkStakedTokenCount: 1,
	}
}
