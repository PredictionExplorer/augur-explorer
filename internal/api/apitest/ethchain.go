//go:build integration

package apitest

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cg "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// Fixture participant addresses (see internal/testfixtures/seed/01_layer1.sql);
// shared by the chain stubs below and the parity cases.
const (
	addrAlice = "0x2100000000000000000000000000000000000021" // round-0 main + chrono winner, CST staker
	addrBob   = "0x2200000000000000000000000000000000000022" // rwalk bidder, endurance winner, banned bid
	addrCarol = "0x2300000000000000000000000000000000000023" // CST bidder, lastcst + raffle winner, rwalk staker
	addrDave  = "0x2400000000000000000000000000000000000024" // eth donor, raffle NFT winner, round-1 winner
	addrEmma  = "0x2500000000000000000000000000000000000025" // donor with info, marketer, round-2 winner
)

// Chain-side fixture state. The values are chosen to cohere with the seeded
// database (internal/testfixtures/seed): round 3 is open, alice placed the
// round's last bid in block 129, carol is the most recent CST bidder, and
// the deterministic testchain timestamp formula pins every derived time.
const (
	// chainTipBlock is the fake chain's head; its testchain timestamp
	// (BaseTime + tip*100 = 1767239800) is what time-derived responses pin.
	chainTipBlock = 142

	// chainRoundNum matches the open fixture round.
	chainRoundNum = 3

	// chainLastBidTime is alice's round-3 bid timestamp (fixture block 129).
	chainLastBidTime = 1767228500
)

// chainStubs bundles the re-stubbable contract handlers the failure-matrix
// tests manipulate per case.
type chainStubs struct {
	game      *testchain.ContractStub
	token     *testchain.ContractStub
	marketing *testchain.ContractStub
}

// registerChainState populates the fake chain the parity harness dials:
// the head block and the eth_call surface of every contract the v1 API
// reads live (CosmicSignatureGame V1, CosmicToken, MarketingWallet).
func registerChainState(chain *testchain.Chain) chainStubs {
	chain.EnsureBlock(chainTipBlock)

	game := ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")
	token := ethcommon.HexToAddress("0x4000000000000000000000000000000000000004")
	charity := ethcommon.HexToAddress("0x6000000000000000000000000000000000000006")
	marketing := ethcommon.HexToAddress("0x1100000000000000000000000000000000000011")
	alice := ethcommon.HexToAddress(addrAlice)
	bob := ethcommon.HexToAddress(addrBob)
	carol := ethcommon.HexToAddress(addrCarol)
	emma := ethcommon.HexToAddress(addrEmma)

	// CosmicSignatureGame, V1 generation: the V1-only CST auction divisor
	// answers, so the mechanics detector pins ContractMechanicsVersion=1.
	gameStub := testchain.MustContractStub(
		cg.CosmicSignatureGameMetaData.ABI,
		cg.CosmicSignatureGameV2MetaData.ABI,
	)

	// Owner-tunable constants.
	gameStub.Return("ethBidPriceIncreaseDivisor", big.NewInt(100))
	gameStub.Return("charityAddress", charity)
	gameStub.Return("charityEthDonationAmountPercentage", big.NewInt(10))
	gameStub.Return("cstRewardAmountForBidding", wei("100000000000000000000")) // 100 CST
	gameStub.Return("mainEthPrizeAmountPercentage", big.NewInt(25))
	gameStub.Return("raffleTotalEthPrizeAmountForBiddersPercentage", big.NewInt(5))
	gameStub.Return("chronoWarriorEthPrizeAmountPercentage", big.NewInt(7))
	gameStub.Return("cosmicSignatureNftStakingTotalEthRewardAmountPercentage", big.NewInt(10))
	gameStub.Return("mainPrizeTimeIncrementIncreaseDivisor", big.NewInt(50))
	gameStub.Return("numRaffleEthPrizesForBidders", big.NewInt(3))
	gameStub.Return("numRaffleCosmicSignatureNftsForBidders", big.NewInt(5))
	gameStub.Return("numRaffleCosmicSignatureNftsForRandomWalkNftStakers", big.NewInt(4))
	gameStub.Return("cstDutchAuctionDurationDivisor", big.NewInt(11)) // V1 marker
	gameStub.Return("cstDutchAuctionDuration", big.NewInt(28800))
	gameStub.Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(33))
	gameStub.Return("getBidCstRewardAmount", wei("99000000000000000000"))
	gameStub.Return("bidCstRewardAmountMultiplier", big.NewInt(7))

	// Per-round live state.
	gameStub.Return("roundNum", big.NewInt(chainRoundNum))
	gameStub.Return("getNextEthBidPrice", wei("1010000000000000"))     // 0.00101 ETH
	gameStub.Return("getNextCstBidPrice", wei("55000000000000000000")) // 55 CST
	gameStub.Return("getCstDutchAuctionDurations", big.NewInt(28800), big.NewInt(3600))
	gameStub.Return("getEthDutchAuctionDurations", big.NewInt(86400), big.NewInt(7200))
	gameStub.Return("getDurationUntilMainPrize", big.NewInt(3600))
	gameStub.Return("getMainEthPrizeAmount", wei("2500000000000000000"))                           // 2.5 ETH
	gameStub.Return("getCosmicSignatureNftStakingTotalEthRewardAmount", wei("100000000000000000")) // 0.1 ETH
	gameStub.Return("getRaffleTotalEthPrizeAmountForBidders", wei("50000000000000000"))            // 0.05 ETH
	gameStub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3600000000))
	gameStub.Return("lastBidderAddress", alice)
	gameStub.Return("initialDurationUntilMainPrizeDivisor", big.NewInt(2))
	gameStub.Return("timeoutDurationToClaimMainPrize", big.NewInt(86400))
	mainPrizeTime := new(big.Int).SetUint64(testchain.BlockTime(chainTipBlock))
	gameStub.Return("mainPrizeTime", mainPrizeTime.Add(mainPrizeTime, big.NewInt(3600)))
	gameStub.Return("delayDurationBeforeRoundActivation", big.NewInt(600))
	gameStub.Return("roundActivationTime", big.NewInt(1767228200))

	// Live special-winners surface.
	gameStub.Return("tryGetCurrentChampions", bob, big.NewInt(600), carol, big.NewInt(800))
	gameStub.Return("enduranceChampionAddress", bob)
	gameStub.Return("enduranceChampionDuration", big.NewInt(600))
	gameStub.Return("enduranceChampionStartTimeStamp", big.NewInt(1767227700))
	gameStub.Return("prevEnduranceChampionDuration", big.NewInt(300))
	gameStub.Return("chronoWarriorDuration", big.NewInt(800))
	gameStub.Return("lastCstBidderAddress", carol)
	gameStub.Handle("biddersInfo", func(args []any) ([]any, error) {
		// alice: her round-3 last bid; carol: no round-3 CST bid, all zeros.
		if addr, ok := args[1].(ethcommon.Address); ok && addr == alice {
			return []any{wei("3030000000000000"), big.NewInt(0), big.NewInt(chainLastBidTime)}, nil
		}
		return []any{big.NewInt(0), big.NewInt(0), big.NewInt(0)}, nil
	})

	// Admin-event enrichment reads (historical, block-pinned).
	gameStub.Return("getInitialDurationUntilMainPrize", big.NewInt(1800))
	gameStub.Return("ethDutchAuctionBeginningBidPrice", wei("2020000000000000"))

	chain.RegisterCall(game, gameStub.Handler())

	// CosmicToken: the token's own metadata surface plus the generic ERC-20
	// balance read used by user/balances.
	tokenStub := testchain.MustContractStub(cg.CosmicSignatureTokenMetaData.ABI, cg.ERC20MetaData.ABI)
	tokenStub.Return("name", "CosmicToken")
	tokenStub.Return("symbol", "CST")
	tokenStub.Return("decimals", uint8(18))
	tokenStub.Return("game", game)
	tokenStub.Return("balanceOf", wei("150000000000000000000")) // 150 CST for every holder
	chain.RegisterCall(token, tokenStub.Handler())

	// MarketingWallet configuration surface.
	marketingStub := testchain.MustContractStub(cg.MarketingWalletMetaData.ABI)
	marketingStub.Return("treasurerAddress", emma)
	marketingStub.Return("token", token)
	marketingStub.Return("owner", alice)
	chain.RegisterCall(marketing, marketingStub.Handler())
	return chainStubs{game: gameStub, token: tokenStub, marketing: marketingStub}
}

// wei parses a decimal wei literal.
func wei(s string) *big.Int {
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("apitest: bad wei literal: " + s)
	}
	return v
}

// newFAQStub returns a fake FAQ bot upstream with fixed responses, so the
// /api/cosmicgame/faq/* proxy routes can be snapshotted deterministically.
func newFAQStub() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok","service":"faq-bot-stub"}`))
	})
	mux.HandleFunc("/api/query", func(w http.ResponseWriter, r *http.Request) {
		// Echo the request body and Accept header so the proxy tests can
		// assert body/header forwarding (the /health golden stays fixed).
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]string{
			"answer":   "stub answer",
			"received": string(body),
			"accept":   r.Header.Get("Accept"),
		}
		_ = json.NewEncoder(w).Encode(resp)
	})
	mux.HandleFunc("/api/reindex", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"reindexed"}`))
	})
	return httptest.NewServer(mux)
}
