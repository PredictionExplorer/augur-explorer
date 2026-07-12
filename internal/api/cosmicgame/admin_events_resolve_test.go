package cosmicgame

import (
	"context"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

var adminResolveTestAddress = ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")

func TestResolveAdminEventFromContract(t *testing.T) {
	t.Parallel()
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI).
		Return("getInitialDurationUntilMainPrize", big.NewInt(600)).
		Return("cstDutchAuctionDuration", big.NewInt(1800)).
		Return("getEthDutchAuctionDurations", big.NewInt(3600), big.NewInt(300)).
		Return("ethDutchAuctionBeginningBidPrice", big.NewInt(1_000_000_000_000_000_000)).
		Return("getNextEthBidPrice", big.NewInt(2_000_000_000_000_000_000))
	v1, v2 := newAdminResolveBindings(t, stub)
	tests := []struct {
		name  string
		event *p.CGAdminEvent
		want  string
	}{
		{name: "nil"},
		{name: "negative block", event: &p.CGAdminEvent{BlockNum: -1, RecordType: 20, IntegerValue: 4}},
		{name: "percent", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 20, IntegerValue: 4}, want: "25.0000%"},
		{name: "microseconds", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 21, IntegerValue: 1_500_000}, want: "1.500000 sec"},
		{name: "initial duration", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 22, IntegerValue: 4}, want: "10m (25.0000%)"},
		{name: "CST direct duration", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 25, IntegerValue: 90}, want: "90 sec"},
		{name: "CST contract duration", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 25}, want: "30m"},
		{name: "ETH duration", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 36, IntegerValue: 1}, want: "1h"},
		{name: "ETH end price", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 37, IntegerValue: 10}, want: "0.10000000 ETH"},
		{name: "CST duration change", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 39, IntegerValue: 10}, want: "3m change per bid"},
		{name: "unknown", event: &p.CGAdminEvent{BlockNum: 1, RecordType: 999, IntegerValue: 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := resolveAdminEventFromContract(v1, v2, test.event); got != test.want {
				t.Fatalf("resolved = %q, want %q", got, test.want)
			}
		})
	}
}

func TestResolveAdminEventEndingPriceFallsBackToNextBid(t *testing.T) {
	t.Parallel()
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI).
		Return("getNextEthBidPrice", big.NewInt(2_000_000_000_000_000_000))
	v1, v2 := newAdminResolveBindings(t, stub)
	event := &p.CGAdminEvent{BlockNum: 1, RecordType: 37, IntegerValue: 20}
	if got := resolveAdminEventFromContract(v1, v2, event); got != "0.10000000 ETH (from next bid price)" {
		t.Fatalf("fallback = %q", got)
	}
}

func TestResolveAdminEventUnavailableContractValues(t *testing.T) {
	t.Parallel()
	v1, v2 := newAdminResolveBindings(t,
		testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI))
	for _, recordType := range []int64{22, 25, 36, 37, 39} {
		event := &p.CGAdminEvent{BlockNum: 1, RecordType: recordType}
		if recordType != 25 {
			event.IntegerValue = 10
		}
		if got := resolveAdminEventFromContract(v1, v2, event); got != "" {
			t.Errorf("record type %d = %q, want empty", recordType, got)
		}
	}
}

func TestAdminEventFormatters(t *testing.T) {
	t.Parallel()
	for input, want := range map[int64]string{
		0: "", -1: "", 4: "25.0000%",
	} {
		if got := formatPercentFromDivisor(input); got != want {
			t.Errorf("percent %d = %q, want %q", input, got, want)
		}
	}
	for input, want := range map[int64]string{
		0: "0 sec", 1_000_000: "1 sec", 1_500_001: "1.500001 sec",
	} {
		if got := formatMicrosecondsAsSeconds(input); got != want {
			t.Errorf("microseconds %d = %q, want %q", input, got, want)
		}
	}
	for input, want := range map[int64]string{
		0: "0 sec", 119: "119 sec", 120: "2m", 121: "2m 1s",
		3600: "1h", 3660: "1h 1m", 3661: "1h 1m 1s",
	} {
		if got := formatDurationSeconds(input); got != want {
			t.Errorf("duration %d = %q, want %q", input, got, want)
		}
	}
	for input, want := range map[int64]string{
		0: "0 ETH", 1_000_000_000_000_000_000: "1.00000000 ETH",
		10_000_000_000_000: "0.000010000000 ETH",
	} {
		if got := formatEthFromWei(input); got != want {
			t.Errorf("ETH %d = %q, want %q", input, got, want)
		}
	}
}

func newAdminResolveBindings(
	t *testing.T,
	stub *testchain.ContractStub,
) (*cgc.CosmicSignatureGame, *cgc.CosmicSignatureGameV2) {
	t.Helper()
	chain := testchain.New(t)
	chain.RegisterCall(adminResolveTestAddress, stub.Handler())
	rpcClient, err := rpc.DialContext(context.Background(), chain.URL())
	if err != nil {
		t.Fatal(err)
	}
	client := ethclient.NewClient(rpcClient)
	t.Cleanup(client.Close)
	v1, err := cgc.NewCosmicSignatureGame(adminResolveTestAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	v2, err := cgc.NewCosmicSignatureGameV2(adminResolveTestAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	return v1, v2
}
