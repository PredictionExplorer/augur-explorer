package cosmicgame

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCGBidRecJSONIncludesV2BidPlacedFields(t *testing.T) {
	rec := CGBidRec{
		PreviousBidderCstRewardAmount: "16882186547446325056",
		PreviousCstRewardAmountEth:    16.882186547446324,
		ThisBidderCstRewardAmount:     "1875798505271813784",
		ThisCstRewardAmountEth:        1.8757985052718137,
		CstDutchAuctionDuration:      "1750",
		CstDutchAuctionDurationInt: 1750,
	}
	raw, err := json.Marshal(rec)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	body := string(raw)
	for _, key := range []string{
		"PreviousBidderCstRewardAmount",
		"PreviousCstRewardAmountEth",
		"ThisBidderCstRewardAmount",
		"ThisCstRewardAmountEth",
		"CstDutchAuctionDuration",
		"CstDutchAuctionDurationInt",
	} {
		if !strings.Contains(body, key) {
			t.Fatalf("JSON missing %q: %s", key, body)
		}
	}
}
