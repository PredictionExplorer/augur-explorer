package gamever

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

var gameAddr = common.HexToAddress("0x1000000000000000000000000000000000000001")

func dialChain(t *testing.T, chain *testchain.Chain) *ethclient.Client {
	t.Helper()
	client, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dialing test chain: %v", err)
	}
	t.Cleanup(client.Close)
	return client
}

func newGameChain(t *testing.T) (*testchain.Chain, *testchain.ContractStub) {
	t.Helper()
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	stub := testchain.MustContractStub(
		cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI, cgc.CosmicSignatureGameV3ABI)
	chain.RegisterCall(gameAddr, stub.Handler())
	return chain, stub
}

func TestDetectV3(t *testing.T) {
	chain, stub := newGameChain(t)
	stub.Return("cstBidPriceDeclineMultiplier", big.NewInt(1))
	// A V3 contract also answers the V2 getter; V3 must win.
	stub.Return("cstDutchAuctionDuration", big.NewInt(1800))

	v, err := Detect(ethtx.CallOpts(), gameAddr, dialChain(t, chain))
	if err != nil || v != V3 {
		t.Fatalf("Detect = %v, %v; want V3, nil", v, err)
	}
	if !v.BidsTakeMinLimit() {
		t.Error("V3 must use the min-limit bid shape")
	}
}

func TestDetectV2(t *testing.T) {
	chain, stub := newGameChain(t)
	stub.Return("cstDutchAuctionDuration", big.NewInt(1800))

	v, err := Detect(ethtx.CallOpts(), gameAddr, dialChain(t, chain))
	if err != nil || v != V2 {
		t.Fatalf("Detect = %v, %v; want V2, nil", v, err)
	}
	if !v.BidsTakeMinLimit() {
		t.Error("V2 must use the min-limit bid shape")
	}
}

func TestDetectV1(t *testing.T) {
	chain, stub := newGameChain(t)
	stub.Return("lastBidderAddress", common.HexToAddress("0xcc"))

	v, err := Detect(ethtx.CallOpts(), gameAddr, dialChain(t, chain))
	if err != nil || v != V1 {
		t.Fatalf("Detect = %v, %v; want V1, nil", v, err)
	}
	if v.BidsTakeMinLimit() {
		t.Error("V1 must use the legacy bid shape")
	}
}

func TestDetectNotAGame(t *testing.T) {
	chain, _ := newGameChain(t)
	// Nothing stubbed: every probe reverts, including the roundNum sanity
	// check, so this is not a game contract (or a wrong address).
	_, err := Detect(ethtx.CallOpts(), gameAddr, dialChain(t, chain))
	if err == nil || !strings.Contains(err.Error(), "no CosmicSignatureGame") {
		t.Fatalf("Detect = %v, want no-game error", err)
	}
}
