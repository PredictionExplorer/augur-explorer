package main

// End-to-end test of the autobid command wiring: cobra args → env config →
// autobid.New → Run against a scripted world where the round ends
// immediately, plus the initialization-failure path.

import (
	"math/big"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	rwc "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

var (
	autobidRWalkAddr  = common.HexToAddress("0x7000000000000000000000000000000000000007")
	autobidPrizesAddr = common.HexToAddress("0x4000000000000000000000000000000000000004")
)

// setAutobidEnv points the autobid environment at the chain.
func setAutobidEnv(t *testing.T, chain *testchain.Chain) {
	t.Helper()
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("PKEY_HEX", testSignerKey)
	t.Setenv("CGAME_ADDR", strings.TrimPrefix(testGameAddr.Hex(), "0x"))
	for _, k := range []string{
		"MAX_ETH_BID", "MAX_CST_BID", "RWALK_MIN_PRICE", "TIME_BEFORE_PRIZE",
		"CST_BID_ANYWAY", "AT_STARTUP_BID_UP_TO_PRICE_LEVEL", "GAS_PRICE_MULTIPLIER",
	} {
		t.Setenv(k, "")
	}
}

func TestAutobidCommandRunsUntilRoundEnds(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(100)
	chain.SetBalance(testSignerAddr, eth(100))
	chain.SetGasPrice(big.NewInt(1_000_000_000))
	setAutobidEnv(t, chain)

	game := testchain.MustContractStub(cgc.CosmicSignatureGameABI)
	game.Return("randomWalkNft", autobidRWalkAddr)
	game.Return("prizesWallet", autobidPrizesAddr)
	game.Return("token", testTokenAddr)
	game.Return("getNextCstBidPrice", eth(20))
	game.Return("getNextEthBidPrice", eth(0.05))
	game.Return("getDurationUntilMainPrize", big.NewInt(1000))
	game.Return("lastBidderAddress", otherAddr)
	game.Return("timeoutDurationToClaimMainPrize", big.NewInt(3600))
	game.Return("mainPrizeTime", big.NewInt(1<<40))
	game.Return("usedRandomWalkNfts", big.NewInt(0))
	var roundCalls atomic.Int64
	game.Handle("roundNum", func([]any) ([]any, error) {
		// The engine's startup read sees round 0; the first loop check sees
		// round 1 and exits.
		if roundCalls.Add(1) == 1 {
			return []any{big.NewInt(0)}, nil
		}
		return []any{big.NewInt(1)}, nil
	})
	chain.RegisterCall(testGameAddr, game.Handler())

	rwalk := testchain.MustContractStub(rwc.RWalkABI)
	rwalk.Return("getMintPrice", eth(0.02))
	rwalk.Return("nextTokenId", big.NewInt(0))
	chain.RegisterCall(autobidRWalkAddr, rwalk.Handler())

	token := testchain.MustContractStub(cgc.ERC20ABI)
	token.Return("balanceOf", eth(100))
	chain.RegisterCall(testTokenAddr, token.Handler())

	prizes := testchain.MustContractStub(cgc.PrizesWalletABI)
	prizes.Return("mainPrizeBeneficiaryAddresses", otherAddr)
	chain.RegisterCall(autobidPrizesAddr, prizes.Handler())

	out, err := executeCmd(t, newAutobidCmd())
	if err != nil {
		t.Fatalf("autobid: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"Playing round 0",
		"Config params:",
		"Round changed (was 0, now 1)",
		"I am not the winner of round 0",
		"Round ended, exiting...",
		"SESSION SUMMARY",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("autobid output missing %q:\n%s", want, out)
		}
	}
}

func TestAutobidCommandInitFailure(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	setAutobidEnv(t, chain)
	// No game contract registered: the RandomWalkNft read fails during New.

	out, err := executeCmd(t, newAutobidCmd())
	if err == nil || !strings.Contains(err.Error(), "failed to initialize bot") {
		t.Errorf("autobid init = %v\noutput: %s", err, out)
	}
}

func TestAutobidCommandConfigFailure(t *testing.T) {
	t.Setenv("RPC_URL", "")
	t.Setenv("PKEY_HEX", "")
	t.Setenv("CGAME_ADDR", "")
	_, err := executeCmd(t, newAutobidCmd())
	if err == nil || !strings.Contains(err.Error(), "configuration errors") {
		t.Errorf("autobid without env = %v", err)
	}
}

func TestDonationRecordsDefaultAddress(t *testing.T) {
	chain := startReadChain(t)
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI)
	stub.Return("numEthDonationWithInfoRecords", big.NewInt(0))
	chain.RegisterCall(common.HexToAddress(defaultLocalGameAddr), stub.Handler())

	out, err := executeCmd(t, newDonationRecordsCmd())
	if err != nil {
		t.Fatalf("donation-records default: %v\noutput: %s", err, out)
	}
	for _, want := range []string{"DEFAULT ADDRESS", "Using default", defaultLocalGameAddr} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestRootCommandRegistersSubcommands(t *testing.T) {
	root := newRootCmd()
	names := make(map[string]bool)
	for _, c := range root.Commands() {
		names[c.Name()] = true
	}
	for _, want := range []string{
		"autobid", "backfill-dao-evtlog", "bid", "claim-and-set-time-increment",
		"claim-prize", "deploy-erc20", "donate", "donation-records", "erc20",
		"info", "nft", "owner", "set-activation-delay", "set-charity-percentage",
		"set-initial-duration-divisor", "set-main-prize-percentage",
		"set-num-nft-winners", "set-num-raffle-winners", "set-raffle-percentage",
		"set-round-activation", "set-staking-percentage", "set-time-increment",
		"token-seed", "total-tokens",
	} {
		if !names[want] {
			t.Errorf("root command missing subcommand %q", want)
		}
	}

	out, err := executeCmd(t, newRootCmd(), "--help")
	if err != nil {
		t.Fatalf("--help: %v", err)
	}
	if !strings.Contains(out, "Operator CLI for the CosmicGame contracts") {
		t.Errorf("help output missing description:\n%s", out)
	}

	out, err = executeCmd(t, newRootCmd(), "--version")
	if err != nil {
		t.Fatalf("--version: %v", err)
	}
	if !strings.HasPrefix(out, "cgctl version ") || !strings.Contains(out, "commit") {
		t.Errorf("--version output = %q, want the build identity line", out)
	}
}
