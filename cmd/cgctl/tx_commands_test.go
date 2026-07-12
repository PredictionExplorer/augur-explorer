package main

// Command-level tests: every transaction subcommand executes with
// cobra-parsed args against a fake chain (internal/testchain) through the
// production wiring — env-var configuration, the shared internal/ethtx
// session, abigen bindings, EIP-155 signing, receipt waiting and the legacy
// output format. Each test builds a fresh command instance so persistent
// flag state cannot leak between runs.

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// testSignerKey is the throwaway key the command tests sign with.
const testSignerKey = "8da4ef21b864d2cc526dbdb2a120bd2874c36c9d0a1fb7f8c63d7f7a8b41de8f"

var testSignerAddr = func() common.Address {
	key, err := crypto.HexToECDSA(testSignerKey)
	if err != nil {
		panic(err)
	}
	return crypto.PubkeyToAddress(key.PublicKey)
}()

var (
	testGameAddr  = common.HexToAddress("0x1000000000000000000000000000000000000001")
	testTokenAddr = common.HexToAddress("0x3000000000000000000000000000000000000003")
	testNFTAddr   = common.HexToAddress("0x5000000000000000000000000000000000000005")
	otherAddr     = common.HexToAddress("0x00000000000000000000000000000000000000cc")
)

// eth converts a float ETH amount to wei for test readability.
func eth(f float64) *big.Int {
	wei := new(big.Float).Mul(big.NewFloat(f), big.NewFloat(1e18))
	out, _ := wei.Int(nil)
	return out
}

// startFundedChain boots a fake chain with a block, a funded signer account
// and the tx environment variables set.
func startFundedChain(t *testing.T) *testchain.Chain {
	t.Helper()
	chain := testchain.New(t)
	chain.EnsureBlock(100)
	chain.SetBalance(testSignerAddr, eth(100))
	chain.SetGasPrice(big.NewInt(1_000_000_000))
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("PKEY_HEX", testSignerKey)
	t.Setenv("GAS_PRICE_MULTIPLIER", "")
	return chain
}

// gameStubFull returns a CosmicGame stub (V1+V2 ABIs) pre-loaded with the
// reads the transaction commands make.
func gameStubFull() *testchain.ContractStub {
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI)
	stub.Return("roundNum", big.NewInt(3))
	stub.Return("getNextEthBidPrice", eth(0.05))
	stub.Return("lastBidderAddress", otherAddr)
	stub.Return("getTotalNumBids", big.NewInt(7))
	stub.Return("getMainEthPrizeAmount", eth(2))
	stub.Return("getDurationUntilMainPrize", big.NewInt(0))
	stub.Return("delayDurationBeforeRoundActivation", big.NewInt(0))
	stub.Return("owner", testSignerAddr)
	stub.Return("roundActivationTime", big.NewInt(int64(testchain.BlockTime(100))+1000))
	stub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3600*1_000_000))
	stub.Return("initialDurationUntilMainPrizeDivisor", big.NewInt(200))
	return stub
}

// registerGame installs the stub at the game address.
func registerGame(t *testing.T, chain *testchain.Chain, stub *testchain.ContractStub) {
	t.Helper()
	chain.RegisterCall(testGameAddr, stub.Handler())
}

// executeCmd runs a freshly built command with args and returns its combined
// output.
func executeCmd(t *testing.T, cmd *cobra.Command, args ...string) (string, error) {
	t.Helper()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return out.String(), err
}

func TestBidCommandSubmitsAndWaits(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, gameStubFull())

	out, err := executeCmd(t, newBidCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("bid: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "Success. Tx hash = 0x") {
		t.Errorf("bid output = %q, want quiet success line", out)
	}
	if got := chain.SubmittedTxCount(); got != 1 {
		t.Errorf("submitted txs = %d, want 1", got)
	}
}

func TestBidCommandVerboseSections(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, gameStubFull())

	out, err := executeCmd(t, newBidCmd(), "-i", testGameAddr.Hex())
	if err != nil {
		t.Fatalf("bid -i: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"NETWORK INFO", "ACCOUNT INFO", "CONTRACT", "ROUND INFO",
		"Round Number", "SUBMITTING TRANSACTION", "Status              = SUCCESS",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("verbose bid output missing %q\noutput:\n%s", want, out)
		}
	}
}

func TestBidCommandRevertReported(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, gameStubFull())
	chain.MarkNextTxReverted()

	out, err := executeCmd(t, newBidCmd(), testGameAddr.Hex())
	if err == nil || !strings.Contains(err.Error(), "transaction did not succeed") {
		t.Fatalf("bid of reverted tx = %v, want failure", err)
	}
	if !strings.Contains(out, "Transaction reverted on-chain") {
		t.Errorf("output = %q, want revert report", out)
	}
}

func TestBidCommandInsufficientBalance(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(100)
	// No SetBalance: the account has zero wei.
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("PKEY_HEX", testSignerKey)
	registerGame(t, chain, gameStubFull())

	_, err := executeCmd(t, newBidCmd(), testGameAddr.Hex())
	if err == nil || !strings.Contains(err.Error(), "insufficient balance") {
		t.Errorf("bid with empty account = %v, want insufficient balance", err)
	}
}

func TestBidCommandInvalidAddress(t *testing.T) {
	_, err := executeCmd(t, newBidCmd(), "not-an-address")
	if err == nil || !strings.Contains(err.Error(), "cosmicgame-addr") {
		t.Errorf("invalid address error = %v", err)
	}
}

func TestBidCommandMissingEnv(t *testing.T) {
	t.Setenv("RPC_URL", "")
	t.Setenv("PKEY_HEX", "")
	_, err := executeCmd(t, newBidCmd(), testGameAddr.Hex())
	if err == nil || !strings.Contains(err.Error(), "RPC_URL") {
		t.Errorf("missing env error = %v, want RPC_URL mention", err)
	}
}

func TestBidCommandHonorsGasMultiplier(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, gameStubFull())
	t.Setenv("GAS_PRICE_MULTIPLIER", "1.5")

	out, err := executeCmd(t, newBidCmd(), "-i", testGameAddr.Hex())
	if err != nil {
		t.Fatalf("bid: %v\noutput: %s", err, out)
	}
	// Suggested price is 1 gwei; the displayed submission price applies 1.5x.
	if !strings.Contains(out, "Gas Price (gwei)    = 1.5000") {
		t.Errorf("verbose output missing the 1.5x price:\n%s", out)
	}
}

func TestBidCommandRejectsBadGasMultiplier(t *testing.T) {
	startFundedChain(t)
	t.Setenv("GAS_PRICE_MULTIPLIER", "free")
	_, err := executeCmd(t, newBidCmd(), testGameAddr.Hex())
	if err == nil || !strings.Contains(err.Error(), "GAS_PRICE_MULTIPLIER") {
		t.Errorf("bad multiplier error = %v", err)
	}
}

func TestDonateCommand(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, gameStubFull())
	chain.SetBalance(testGameAddr, eth(10))

	out, err := executeCmd(t, newDonateCmd(), testGameAddr.Hex(), eth(1).String())
	if err != nil {
		t.Fatalf("donate: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "Success. Tx hash = 0x") {
		t.Errorf("donate output = %q", out)
	}

	// Insufficient balance path.
	_, err = executeCmd(t, newDonateCmd(), testGameAddr.Hex(), eth(1000).String())
	if err == nil || !strings.Contains(err.Error(), "insufficient balance") {
		t.Errorf("oversized donation = %v, want insufficient balance", err)
	}

	// Invalid amount path.
	_, err = executeCmd(t, newDonateCmd(), testGameAddr.Hex(), "one-eth")
	if err == nil || !strings.Contains(err.Error(), "amount") {
		t.Errorf("invalid amount = %v", err)
	}
}

func TestClaimPrizeCommand(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, gameStubFull())

	out, err := executeCmd(t, newClaimPrizeCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("claim-prize: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "Success. Tx hash = 0x") {
		t.Errorf("claim output = %q", out)
	}
}

func TestClaimPrizeCommandWarnsWhenNotClaimable(t *testing.T) {
	chain := startFundedChain(t)
	stub := gameStubFull()
	stub.Return("lastBidderAddress", otherAddr)
	stub.Return("getDurationUntilMainPrize", big.NewInt(120))
	registerGame(t, chain, stub)

	out, err := executeCmd(t, newClaimPrizeCmd(), "-i", testGameAddr.Hex())
	if err != nil {
		t.Fatalf("claim-prize -i: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"You are NOT the last bidder. Claim may fail unless timeout has passed.",
		"Prize is NOT yet claimable",
		"Wait Time Remaining",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestClaimPrizeWithDelayWaitsForBothReceipts(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, gameStubFull())

	out, err := executeCmd(t, newClaimPrizeCmd(), "-i", "--delay=120", testGameAddr.Hex())
	if err != nil {
		t.Fatalf("claim-prize --delay: %v\noutput: %s", err, out)
	}
	if got := chain.SubmittedTxCount(); got != 2 {
		t.Errorf("submitted txs = %d, want 2 (set delay + claim)", got)
	}
	for _, want := range []string{
		"STEP 1: SET DELAY", "STEP 2: CLAIM PRIZE", "SUMMARY",
		"Prize claimed successfully",
		"New round will activate 120 seconds after the claim",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
	// The legacy blind two-second sleep is gone: no "Waiting 2 seconds".
	if strings.Contains(out, "Waiting 2 seconds") {
		t.Errorf("output still contains the legacy blind sleep:\n%s", out)
	}
}

func TestClaimPrizeWithDelayAbortsWhenDelayTxFails(t *testing.T) {
	chain := startFundedChain(t)
	registerGame(t, chain, gameStubFull())
	chain.MarkNextTxReverted()

	out, err := executeCmd(t, newClaimPrizeCmd(), "--delay", testGameAddr.Hex())
	if err == nil || !strings.Contains(err.Error(), "failed to set delay, aborting") {
		t.Fatalf("delay revert = %v, want abort", err)
	}
	if got := chain.SubmittedTxCount(); got != 1 {
		t.Errorf("submitted txs = %d, want 1 (claim must not fire after a failed delay)", got)
	}
	_ = out
}

func TestStandaloneSetterCommands(t *testing.T) {
	cases := []struct {
		name string
		cmd  func() *cobra.Command
		args []string
	}{
		{"set-activation-delay", newSetActivationDelayCmd, []string{testGameAddr.Hex(), "300"}},
		{"set-round-activation", newSetRoundActivationCmd, []string{testGameAddr.Hex(), "1900000000"}},
		{"set-time-increment", newSetTimeIncrementCmd, []string{testGameAddr.Hex(), "3600"}},
		{"set-initial-duration-divisor", newSetInitialDurationDivisorCmd, []string{testGameAddr.Hex(), "100"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			chain := startFundedChain(t)
			registerGame(t, chain, gameStubFull())
			out, err := executeCmd(t, tc.cmd(), tc.args...)
			if err != nil {
				t.Fatalf("%s: %v\noutput: %s", tc.name, err, out)
			}
			if !strings.Contains(out, "Success. Tx hash = 0x") {
				t.Errorf("%s output = %q", tc.name, out)
			}
			if got := chain.SubmittedTxCount(); got != 1 {
				t.Errorf("%s submitted %d txs, want 1", tc.name, got)
			}
		})
	}
}

func TestStandaloneSetterValidation(t *testing.T) {
	startFundedChain(t)
	if _, err := executeCmd(t, newSetTimeIncrementCmd(), testGameAddr.Hex(), "0"); err == nil ||
		!strings.Contains(err.Error(), "must be positive") {
		t.Errorf("zero increment = %v", err)
	}
	if _, err := executeCmd(t, newSetInitialDurationDivisorCmd(), "--", testGameAddr.Hex(), "-5"); err == nil ||
		!strings.Contains(err.Error(), "must be positive") {
		t.Errorf("negative divisor = %v", err)
	}
}

func TestSetterWarnsWhenNotOwner(t *testing.T) {
	chain := startFundedChain(t)
	stub := gameStubFull()
	stub.Return("owner", otherAddr)
	registerGame(t, chain, stub)

	out, err := executeCmd(t, newSetActivationDelayCmd(), "-i", testGameAddr.Hex(), "60")
	if err != nil {
		t.Fatalf("set-activation-delay: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "You are NOT the contract owner. Transaction will likely fail.") {
		t.Errorf("output missing owner warning:\n%s", out)
	}
}

func TestSpecSetterCommands(t *testing.T) {
	// Every data-driven gameSetterSpec command: read the current value, show
	// current/new, submit and wait.
	specs := []struct {
		name string
		spec gameSetterSpec
		read string // stubbed read method
	}{
		{"set-charity-percentage", charityPercentageSpec, "charityEthDonationAmountPercentage"},
		{"set-main-prize-percentage", mainPrizePercentageSpec, "mainEthPrizeAmountPercentage"},
		{"set-staking-percentage", stakingPercentageSpec, "cosmicSignatureNftStakingTotalEthRewardAmountPercentage"},
		{"set-raffle-percentage", rafflePercentageSpec, "raffleTotalEthPrizeAmountForBiddersPercentage"},
		{"set-num-raffle-winners", numRaffleWinnersSpec, "numRaffleEthPrizesForBidders"},
		{"set-num-nft-winners", numNftWinnersSpec, "numRaffleCosmicSignatureNftsForBidders"},
	}
	for _, tc := range specs {
		t.Run(tc.name, func(t *testing.T) {
			chain := startFundedChain(t)
			stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI)
			stub.Return(tc.read, big.NewInt(10))
			registerGame(t, chain, stub)

			out, err := executeCmd(t, newGameSetterCmd(tc.spec), "-i", testGameAddr.Hex(), "25")
			if err != nil {
				t.Fatalf("%s: %v\noutput: %s", tc.name, err, out)
			}
			for _, want := range []string{tc.spec.section, "Current Value", "New Value", "Status              = SUCCESS"} {
				if !strings.Contains(out, want) {
					t.Errorf("%s output missing %q:\n%s", tc.name, want, out)
				}
			}
			if got := chain.SubmittedTxCount(); got != 1 {
				t.Errorf("%s submitted %d txs, want 1", tc.name, got)
			}
		})
	}
}

func TestSpecSetterReadFailure(t *testing.T) {
	chain := startFundedChain(t)
	// Stub with no methods: the read reverts.
	registerGame(t, chain, testchain.MustContractStub(cgc.CosmicSignatureGameABI))

	_, err := executeCmd(t, newGameSetterCmd(charityPercentageSpec), testGameAddr.Hex(), "25")
	if err == nil || !strings.Contains(err.Error(), "getting current value") {
		t.Errorf("read failure = %v", err)
	}
}

func TestERC20ApproveAndRevoke(t *testing.T) {
	tokenStub := func() *testchain.ContractStub {
		stub := testchain.MustContractStub(cgc.ERC20ABI)
		stub.Return("symbol", "CST")
		stub.Return("allowance", big.NewInt(0))
		stub.Return("balanceOf", eth(9))
		return stub
	}

	t.Run("approve", func(t *testing.T) {
		chain := startFundedChain(t)
		chain.RegisterCall(testTokenAddr, tokenStub().Handler())
		out, err := executeCmd(t, newERC20ApproveCmd(), "-i", testTokenAddr.Hex(), otherAddr.Hex())
		if err != nil {
			t.Fatalf("approve: %v\noutput: %s", err, out)
		}
		for _, want := range []string{"APPROVAL INFO", "MAX_UINT256 (unlimited)", "Status              = SUCCESS"} {
			if !strings.Contains(out, want) {
				t.Errorf("approve output missing %q:\n%s", want, out)
			}
		}
	})

	t.Run("revoke", func(t *testing.T) {
		chain := startFundedChain(t)
		chain.RegisterCall(testTokenAddr, tokenStub().Handler())
		out, err := executeCmd(t, newERC20RevokeCmd(), testTokenAddr.Hex(), otherAddr.Hex())
		if err != nil {
			t.Fatalf("revoke: %v\noutput: %s", err, out)
		}
		if !strings.Contains(out, "Success. Tx hash = 0x") {
			t.Errorf("revoke output = %q", out)
		}
	})

	t.Run("approve invalid spender", func(t *testing.T) {
		startFundedChain(t)
		_, err := executeCmd(t, newERC20ApproveCmd(), testTokenAddr.Hex(), "nope")
		if err == nil || !strings.Contains(err.Error(), "spender-addr") {
			t.Errorf("invalid spender = %v", err)
		}
	})
}

func TestNFTSetNameCommand(t *testing.T) {
	chain := startFundedChain(t)
	chain.RegisterCall(testNFTAddr, testchain.MustContractStub(cgc.CosmicSignatureNftABI).Handler())

	out, err := executeCmd(t, newNFTSetNameCmd(), "-i", testNFTAddr.Hex(), "7", "my token")
	if err != nil {
		t.Fatalf("set-name: %v\noutput: %s", err, out)
	}
	for _, want := range []string{"NFT NAME CONFIG", "New Name", "my token", "Status              = SUCCESS"} {
		if !strings.Contains(out, want) {
			t.Errorf("set-name output missing %q:\n%s", want, out)
		}
	}

	// Empty name variant.
	out, err = executeCmd(t, newNFTSetNameCmd(), "-i", testNFTAddr.Hex(), "7")
	if err != nil {
		t.Fatalf("set-name empty: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "(empty)") {
		t.Errorf("empty-name output missing (empty):\n%s", out)
	}
}

func TestDeployERC20Command(t *testing.T) {
	chain := startFundedChain(t)

	out, err := executeCmd(t, newDeployERC20Cmd())
	if err != nil {
		t.Fatalf("deploy-erc20: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"Deploying SampToken contract...",
		"=== DEPLOYMENT SUCCESSFUL ===",
		"Contract Address: 0x",
		"Token Name: ERC20 Token Sample1",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("deploy output missing %q:\n%s", want, out)
		}
	}
	if got := chain.SubmittedTxCount(); got != 1 {
		t.Errorf("submitted txs = %d, want 1", got)
	}
}

func TestDeployERC20CommandFailedStatus(t *testing.T) {
	chain := startFundedChain(t)
	chain.MarkNextTxReverted()

	_, err := executeCmd(t, newDeployERC20Cmd())
	if err == nil || !strings.Contains(err.Error(), "transaction failed with status") {
		t.Errorf("reverted deploy = %v", err)
	}
}

// The receipt-timeout wait path is covered at the session level in
// internal/ethtx (TestFinishTxReportsReceiptTimeout); commands cannot
// configure the timeout, so no command-level variant exists.
