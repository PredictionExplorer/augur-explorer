package cosmicgame

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCGBidRecJSONIncludesV2BidPlacedFields(t *testing.T) {
	rec := CGBidRec{
		BidCstRewardAmount:         "18757985052718138840",
		BidCstRewardAmountEth:      18.757985052718137,
		CstDutchAuctionDuration:      "1750",
		CstDutchAuctionDurationInt: 1750,
	}
	raw, err := json.Marshal(rec)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	body := string(raw)
	for _, key := range []string{
		"BidCstRewardAmount",
		"BidCstRewardAmountEth",
		"CstDutchAuctionDuration",
		"CstDutchAuctionDurationInt",
	} {
		if !strings.Contains(body, key) {
			t.Fatalf("JSON missing %q: %s", key, body)
		}
	}
}
