package v2

import (
	"testing"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

func TestMapBid(t *testing.T) {
	t.Parallel()

	record := validBidRecord()
	got, err := mapBid(record)
	if err != nil {
		t.Fatalf("mapBid: %v", err)
	}

	if got.EventLogId != 5008 || got.Round != 2 || got.Position != 3 {
		t.Errorf("identity fields = (%d,%d,%d), want (5008,2,3)", got.EventLogId, got.Round, got.Position)
	}
	if got.BidType != Cst {
		t.Errorf("BidType = %q, want %q", got.BidType, Cst)
	}
	if got.OccurredAt.Location().String() != "UTC" || got.PrizeAt.Location().String() != "UTC" {
		t.Errorf("timestamps are not normalized to UTC: %s / %s", got.OccurredAt.Location(), got.PrizeAt.Location())
	}
	if got.EthPriceWei != nil {
		t.Errorf("legacy -1 ETH price should be omitted, got %q", *got.EthPriceWei)
	}
	if got.CstPriceWei == nil || *got.CstPriceWei != "200000000000000000000" {
		t.Errorf("CstPriceWei = %v", got.CstPriceWei)
	}
	if got.BidCstRewardAmountWei == nil || *got.BidCstRewardAmountWei != "98000000000000000000" {
		t.Errorf("BidCstRewardAmountWei = %v", got.BidCstRewardAmountWei)
	}
	if got.CstDutchAuctionDurationSeconds == nil || *got.CstDutchAuctionDurationSeconds != 1800 {
		t.Errorf("CstDutchAuctionDurationSeconds = %v", got.CstDutchAuctionDurationSeconds)
	}
	if got.RandomWalkTokenId == nil || *got.RandomWalkTokenId != 13 {
		t.Errorf("RandomWalkTokenId = %v", got.RandomWalkTokenId)
	}
	if got.Message == nil || *got.Message != "hello" {
		t.Errorf("Message = %v", got.Message)
	}
	if got.NftDonation == nil || got.NftDonation.TokenId != 777 {
		t.Errorf("NftDonation = %+v", got.NftDonation)
	}
	if got.Erc20Donation == nil || got.Erc20Donation.AmountWei != "500000000000000000000" {
		t.Errorf("Erc20Donation = %+v", got.Erc20Donation)
	}
}

func TestMapBidOmitsAbsentOptionalFields(t *testing.T) {
	t.Parallel()

	record := validBidRecord()
	record.EthPrice = "10"
	record.CstPrice = "-1"
	record.RWalkNFTId = -1
	record.Message = ""
	record.BidCstRewardAmount = "-1"
	record.CstDutchAuctionDurationInt = -1
	record.NFTDonationTokenId = -1
	record.NFTDonationTokenAddr = ""
	record.NFTTokenURI = ""
	record.DonatedERC20TokenAddr = ""
	record.DonatedERC20TokenAmount = ""

	got, err := mapBid(record)
	if err != nil {
		t.Fatalf("mapBid: %v", err)
	}
	if got.EthPriceWei == nil || *got.EthPriceWei != "10" {
		t.Errorf("EthPriceWei = %v", got.EthPriceWei)
	}
	if got.CstPriceWei != nil ||
		got.RandomWalkTokenId != nil ||
		got.Message != nil ||
		got.BidCstRewardAmountWei != nil ||
		got.CstDutchAuctionDurationSeconds != nil ||
		got.NftDonation != nil ||
		got.Erc20Donation != nil {
		t.Fatalf("optional sentinel fields were not omitted: %+v", got)
	}
}

func TestMapBidRejectsInvalidStoreData(t *testing.T) {
	t.Parallel()

	tests := map[string]func(*cgprimitives.CGBidRec){
		"identity":           func(r *cgprimitives.CGBidRec) { r.BidPosition = 0 },
		"occurred timestamp": func(r *cgprimitives.CGBidRec) { r.Tx.DateTime = "not-a-time" },
		"prize timestamp":    func(r *cgprimitives.CGBidRec) { r.PrizeTimeDate = "not-a-time" },
		"bidder address":     func(r *cgprimitives.CGBidRec) { r.BidderAddr = "bad" },
		"transaction hash":   func(r *cgprimitives.CGBidRec) { r.Tx.TxHash = "0x01" },
		"cst reward":         func(r *cgprimitives.CGBidRec) { r.CSTReward = "-1" },
		"eth price":          func(r *cgprimitives.CGBidRec) { r.EthPrice = "1.2" },
		"cst price":          func(r *cgprimitives.CGBidRec) { r.CstPrice = "" },
		"bid reward":         func(r *cgprimitives.CGBidRec) { r.BidCstRewardAmount = "-2" },
		"nft address":        func(r *cgprimitives.CGBidRec) { r.NFTDonationTokenAddr = "bad" },
		"erc20 address":      func(r *cgprimitives.CGBidRec) { r.DonatedERC20TokenAddr = "bad" },
		"erc20 amount":       func(r *cgprimitives.CGBidRec) { r.DonatedERC20TokenAmount = "wat" },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			record := validBidRecord()
			mutate(&record)
			if _, err := mapBid(record); err == nil {
				t.Fatal("mapBid accepted invalid store data")
			}
		})
	}
}

func TestMapBidType(t *testing.T) {
	t.Parallel()

	tests := map[int64]BidType{0: Eth, 1: RandomWalk, 2: Cst, 3: Unknown, -1: Unknown}
	for input, want := range tests {
		if got := mapBidType(input); got != want {
			t.Errorf("mapBidType(%d) = %q, want %q", input, got, want)
		}
	}
}

func validBidRecord() cgprimitives.CGBidRec {
	// #nosec G101 -- deterministic chain amounts and addresses, not credentials.
	return cgprimitives.CGBidRec{
		Tx: cgprimitives.Transaction{
			EvtLogId: 5008,
			BlockNum: 103,
			TxHash:   "0xf000000000000000000000000000000000000000000000000000000000001004",
			DateTime: "2026-01-01T00:05:00-05:00",
		},
		BidderAddr:                 "0x2300000000000000000000000000000000000023",
		EthPrice:                   "-1",
		CstPrice:                   "200000000000000000000",
		RWalkNFTId:                 13,
		RoundNum:                   2,
		BidType:                    2,
		BidPosition:                3,
		PrizeTimeDate:              "2026-01-01T01:05:00-05:00",
		CSTReward:                  "100000000000000000000",
		BidCstRewardAmount:         "98000000000000000000",
		CstDutchAuctionDurationInt: 1800,
		NFTDonationTokenId:         777,
		NFTDonationTokenAddr:       "0x2700000000000000000000000000000000000027",
		NFTTokenURI:                "https://nft.example/777",
		Message:                    "hello",
		DonatedERC20TokenAddr:      "0x2600000000000000000000000000000000000026",
		DonatedERC20TokenAmount:    "500000000000000000000",
		CstDutchAuctionDuration:    "1800",
		DonatedERC20TokenAmountEth: 500,
		BidCstRewardAmountEth:      98,
		CSTRewardEth:               100,
		CstPriceEth:                200,
		EthPriceEth:                -1,
	}
}
