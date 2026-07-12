package main

// Command-level tests: every transaction/read subcommand executes with
// cobra-parsed args against a fake chain (internal/testchain) through the
// production wiring — env-var configuration, internal/ethtx session, abigen
// bindings, EIP-155 signing, receipt waiting and the legacy output format.
// Each test builds a fresh command instance so persistent flag state cannot
// leak between runs.

import (
	"bytes"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
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
	testRWalkAddr  = common.HexToAddress("0x895a000000000000000000000000000000000001")
	testMarketAddr = common.HexToAddress("0x47ef000000000000000000000000000000000002")
)

// startFundedChain boots a fake chain with a block, a funded signer account
// and the tx environment variables set.
func startFundedChain(t *testing.T) *testchain.Chain {
	t.Helper()
	chain := testchain.New(t)
	chain.EnsureBlock(100)
	oneHundredEth, _ := new(big.Int).SetString("100000000000000000000", 10)
	chain.SetBalance(testSignerAddr, oneHundredEth)
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("PKEY_HEX", testSignerKey)
	return chain
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

func TestMintCommandSubmitsTransaction(t *testing.T) {
	startFundedChain(t)
	out, err := executeCmd(t, newMintCmd(), testRWalkAddr.Hex(), "1000000000000000000")
	if err != nil {
		t.Fatalf("mint: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "Success. Tx hash = 0x") {
		t.Errorf("mint output = %q, want quiet success line", out)
	}
}

func TestMintCommandVerboseSections(t *testing.T) {
	startFundedChain(t)
	out, err := executeCmd(t, newMintCmd(), "-i", testRWalkAddr.Hex(), "1000000000000000000")
	if err != nil {
		t.Fatalf("mint -i: %v\noutput: %s", err, out)
	}
	for _, want := range []string{"NETWORK INFO", "ACCOUNT INFO", "MINT INFO", "SUBMITTING TRANSACTION", "Status              = SUCCESS"} {
		if !strings.Contains(out, want) {
			t.Errorf("verbose mint output missing %q\noutput:\n%s", want, out)
		}
	}
}

func TestMintCommandInsufficientBalance(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(100)
	// No SetBalance: the account has zero wei.
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("PKEY_HEX", testSignerKey)

	_, err := executeCmd(t, newMintCmd(), testRWalkAddr.Hex(), "1000000000000000000")
	if err == nil || !strings.Contains(err.Error(), "insufficient balance") {
		t.Errorf("mint with empty account = %v, want insufficient balance", err)
	}
}

func TestMintCommandRevertedTransaction(t *testing.T) {
	chain := startFundedChain(t)
	chain.MarkNextTxReverted()
	out, err := executeCmd(t, newMintCmd(), testRWalkAddr.Hex(), "1000")
	if err == nil || !strings.Contains(err.Error(), "transaction did not succeed") {
		t.Fatalf("mint of reverted tx = %v, want failure", err)
	}
	if !strings.Contains(out, "Transaction reverted on-chain") {
		t.Errorf("output = %q, want revert report", out)
	}
}

func TestMintCommandInvalidAmount(t *testing.T) {
	_, err := executeCmd(t, newMintCmd(), testRWalkAddr.Hex(), "not-a-number")
	if err == nil || !strings.Contains(err.Error(), "amount_wei") {
		t.Errorf("invalid amount error = %v", err)
	}
}

func TestMintCommandMissingEnv(t *testing.T) {
	t.Setenv("RPC_URL", "")
	t.Setenv("PKEY_HEX", "")
	_, err := executeCmd(t, newMintCmd(), testRWalkAddr.Hex(), "1000")
	if err == nil || !strings.Contains(err.Error(), "RPC_URL") {
		t.Errorf("missing env error = %v, want RPC_URL mention", err)
	}
}

func TestTransferApproveSetNameCommands(t *testing.T) {
	cases := []struct {
		name string
		cmd  func() *cobra.Command
		args []string
	}{
		{"transfer", newTransferCmd, []string{testRWalkAddr.Hex(), "42", "0x1111111111111111111111111111111111111111"}},
		{"approve", newApproveCmd, []string{testRWalkAddr.Hex(), "0x2222222222222222222222222222222222222222"}},
		{"set-name", newSetNameCmd, []string{testRWalkAddr.Hex(), "7", "my token"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			startFundedChain(t)
			out, err := executeCmd(t, tc.cmd(), tc.args...)
			if err != nil {
				t.Fatalf("%s: %v\noutput: %s", tc.name, err, out)
			}
			if !strings.Contains(out, "Success. Tx hash = 0x") {
				t.Errorf("%s output = %q", tc.name, out)
			}
		})
	}
}

func TestNewOfferCommand(t *testing.T) {
	t.Run("SELL", func(t *testing.T) {
		startFundedChain(t)
		out, err := executeCmd(t, newNewOfferCmd(),
			"SELL", testMarketAddr.Hex(), testRWalkAddr.Hex(), "7", "2500000000000000000")
		if err != nil {
			t.Fatalf("new-offer SELL: %v\noutput: %s", err, out)
		}
		if !strings.Contains(out, "Success. Tx hash = 0x") {
			t.Errorf("output = %q", out)
		}
	})
	t.Run("BUY sends the offer value", func(t *testing.T) {
		startFundedChain(t)
		out, err := executeCmd(t, newNewOfferCmd(),
			"-i", "BUY", testMarketAddr.Hex(), testRWalkAddr.Hex(), "7", "2500000000000000000")
		if err != nil {
			t.Fatalf("new-offer BUY: %v\noutput: %s", err, out)
		}
		if !strings.Contains(out, "Value (ETH)         = 2.500000000000000000") {
			t.Errorf("BUY offer output missing the attached value:\n%s", out)
		}
	})
	t.Run("invalid type", func(t *testing.T) {
		_, err := executeCmd(t, newNewOfferCmd(),
			"LEND", testMarketAddr.Hex(), testRWalkAddr.Hex(), "7", "1")
		if err == nil || !strings.Contains(err.Error(), "BUY or SELL") {
			t.Errorf("invalid offer type error = %v", err)
		}
	})
}

// stubMarketOffer registers an offers(id) answer on the market contract:
// a SELL offer (zero buyer) when sell is true, otherwise a BUY offer.
func stubMarketOffer(t *testing.T, chain *testchain.Chain, sell bool) {
	t.Helper()
	buyer := common.HexToAddress("0x3333333333333333333333333333333333333333")
	seller := common.HexToAddress("0x4444444444444444444444444444444444444444")
	if sell {
		buyer = common.Address{}
	}
	stub := testchain.MustContractStub(rwcontracts.RWMarketMetaData.ABI).
		Return("offers", testRWalkAddr, big.NewInt(7), big.NewInt(1_000_000_000_000_000_000), seller, buyer, true)
	chain.RegisterCall(testMarketAddr, stub.Handler())
}

func TestAcceptOfferCommand(t *testing.T) {
	for _, sell := range []bool{true, false} {
		name := "BUY"
		if sell {
			name = "SELL"
		}
		t.Run(name, func(t *testing.T) {
			chain := startFundedChain(t)
			stubMarketOffer(t, chain, sell)
			out, err := executeCmd(t, newAcceptOfferCmd(), "-i", testMarketAddr.Hex(), "5")
			if err != nil {
				t.Fatalf("accept-offer: %v\noutput: %s", err, out)
			}
			if !strings.Contains(out, "Type                        = "+name) {
				t.Errorf("accept-offer output missing offer type %s:\n%s", name, out)
			}
			if !strings.Contains(out, "Status              = SUCCESS") {
				t.Errorf("accept-offer output missing success:\n%s", out)
			}
		})
	}
}

func TestCancelOfferCommand(t *testing.T) {
	chain := startFundedChain(t)
	stubMarketOffer(t, chain, true)
	out, err := executeCmd(t, newCancelOfferCmd(), testMarketAddr.Hex(), "5")
	if err != nil {
		t.Fatalf("cancel-offer: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "Success. Tx hash = 0x") {
		t.Errorf("output = %q", out)
	}
}

func TestAcceptOfferUnreadableOffer(t *testing.T) {
	startFundedChain(t) // no offers stub registered: the read reverts
	_, err := executeCmd(t, newAcceptOfferCmd(), testMarketAddr.Hex(), "5")
	if err == nil || !strings.Contains(err.Error(), "offers(offer_id=5)") {
		t.Errorf("unreadable offer error = %v", err)
	}
}

// stubRWalkReads registers the read-only RandomWalk contract surface the
// read commands consume.
func stubRWalkReads(t *testing.T, chain *testchain.Chain) {
	t.Helper()
	stub := testchain.MustContractStub(rwcontracts.RWalkMetaData.ABI).
		Return("getMintPrice", big.NewInt(1_250_000_000_000_000)).
		Return("nextTokenId", big.NewInt(4068)).
		Return("timeUntilWithdrawal", big.NewInt(86400)).
		Return("withdrawalAmount", big.NewInt(2_000_000_000_000_000_000)).
		Return("numWithdrawals", big.NewInt(9)).
		Return("lastMinter", testSignerAddr).
		Return("tokenURI", "https://example/token/0").
		Return("ownerOf", testSignerAddr)
	chain.RegisterCall(testRWalkAddr, stub.Handler())
}

func TestReadCommands(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	stubRWalkReads(t, chain)
	marketStub := testchain.MustContractStub(rwcontracts.RWMarketMetaData.ABI).
		Return("numOffers", big.NewInt(17))
	chain.RegisterCall(testMarketAddr, marketStub.Handler())
	t.Setenv("RPC_URL", chain.URL())

	cases := []struct {
		name string
		cmd  func() *cobra.Command
		args []string
		want []string
	}{
		{"price", newPriceCmd, []string{testRWalkAddr.Hex()}, []string{"Mint price = 1250000000000000"}},
		{"withdrawal", newWithdrawalCmd, []string{testRWalkAddr.Hex()}, []string{"Withdrawal amount = 2000000000000000000"}},
		{"owner-of", newOwnerOfCmd, []string{testRWalkAddr.Hex(), "3"}, []string{"Owner  = " + testSignerAddr.Hex()}},
		{"token-uri", newTokenURICmd, []string{testRWalkAddr.Hex(), "0"}, []string{"https://example/token/0"}},
		{"status-market", newStatusMarketCmd, []string{testMarketAddr.Hex()}, []string{"NumOffers = 17"}},
		{"status", newStatusCmd, []string{testRWalkAddr.Hex()}, []string{
			"Next token ID = 4068",
			"Time remaining: 86400",
			"Withdrawal amount: 2.000000000000000000",
			"Num withdrawals: 9",
			"Last minter: " + testSignerAddr.Hex(),
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := executeCmd(t, tc.cmd(), tc.args...)
			if err != nil {
				t.Fatalf("%s: %v\noutput: %s", tc.name, err, out)
			}
			for _, want := range tc.want {
				if !strings.Contains(out, want) {
					t.Errorf("%s output missing %q\noutput:\n%s", tc.name, want, out)
				}
			}
		})
	}
}

func TestReadCommandContractRevert(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	t.Setenv("RPC_URL", chain.URL())
	// No stub registered: every eth_call reverts.
	_, err := executeCmd(t, newPriceCmd(), testRWalkAddr.Hex())
	if err == nil || !strings.Contains(err.Error(), "GetMintPrice") {
		t.Errorf("price against dead contract = %v", err)
	}
}

func TestRootCommandRegistersAllSubcommands(t *testing.T) {
	root := newRootCmd()
	want := []string{
		"mint", "transfer", "approve", "set-name",
		"new-offer", "accept-offer", "cancel-offer",
		"price", "withdrawal", "status", "status-market", "owner-of", "token-uri",
		"scan-mints", "scan-transfers", "verify-owner", "verify-erc20-transfers",
		"top-rated", "notify-bot", "tweet-mints", "tweet-send", "tweet-reply-image", "twitter-auth",
	}
	registered := map[string]bool{}
	for _, c := range root.Commands() {
		registered[c.Name()] = true
	}
	for _, name := range want {
		if !registered[name] {
			t.Errorf("subcommand %q not registered", name)
		}
	}
}

func TestScanTransfersCommand(t *testing.T) {
	chain := testchain.New(t)
	rwalkAddr := testRWalkAddr
	tx1 := chain.AddTx(5, rwalkAddr, nil)
	makeTransfer := func(token int64) *types.Log {
		return &types.Log{
			Address: rwalkAddr,
			Topics: []common.Hash{
				transferEventTopic,
				common.HexToHash("0x1111111111111111111111111111111111111111"),
				common.HexToHash("0x2222222222222222222222222222222222222222"),
				common.BigToHash(big.NewInt(token)),
			},
			BlockNumber: 5,
			TxHash:      tx1.Hash(),
		}
	}
	chain.AttachLogs(tx1.Hash(), []*types.Log{makeTransfer(3601), makeTransfer(42)})
	t.Setenv("RPC_URL", chain.URL())

	out, err := executeCmd(t, newScanTransfersCmd(), rwalkAddr.Hex())
	if err != nil {
		t.Fatalf("scan-transfers: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "(tok 3601)") {
		t.Errorf("output missing the default-token transfer:\n%s", out)
	}
	if strings.Contains(out, "(tok 42)") {
		t.Errorf("output includes a filtered-out token:\n%s", out)
	}

	out, err = executeCmd(t, newScanTransfersCmd(), "--token-id", "42", rwalkAddr.Hex())
	if err != nil {
		t.Fatalf("scan-transfers --token-id: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "(tok 42)") {
		t.Errorf("output missing the selected token:\n%s", out)
	}
}

func TestTweetReplyImageMissingMediaFile(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("TWITTER_KEYS_FILE", "tw.json")
	if err := os.MkdirAll(filepath.Join(home, "configs"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(home, "configs", "tw.json"),
		[]byte(`{"ApiKey":"a","ApiSecret":"b","TokenKey":"c","TokenSecret":"d"}`), 0o600); err != nil {
		t.Fatal(err)
	}

	out, err := executeCmd(t, newTweetReplyImageCmd(), "123", filepath.Join(home, "missing.mp4"), "msg")
	if err == nil || !strings.Contains(err.Error(), "can't read media data") {
		t.Errorf("tweet-reply-image = %v\noutput: %s", err, out)
	}
}

func TestTweetSendRejectsBadNonce(t *testing.T) {
	_, err := executeCmd(t, newTweetSendCmd(), "k", "s", "at", "ts", "not-hex")
	if err == nil || !strings.Contains(err.Error(), "parsing nonce") {
		t.Errorf("tweet-send with bad nonce = %v", err)
	}
}

func TestNotifyBotDialFailure(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("TWITTER_KEYS_FILE", "tw.json")
	t.Setenv("RPC_URL", "")
	if err := os.MkdirAll(filepath.Join(home, "configs"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(home, "configs", "tw.json"),
		[]byte(`{"ApiKey":"a","ApiSecret":"b","TokenKey":"c","TokenSecret":"d"}`), 0o600); err != nil {
		t.Fatal(err)
	}

	_, err := executeCmd(t, newNotifyBotCmd())
	if err == nil || !strings.Contains(err.Error(), "can't connect to ETH RPC") {
		t.Errorf("notify-bot without RPC = %v", err)
	}
}

func TestTwitterAuthCommandMissingConfig(t *testing.T) {
	_, err := executeCmd(t, newTwitterAuthCmd(), "--config", filepath.Join(t.TempDir(), "nope.json"))
	if err == nil || !strings.Contains(err.Error(), "can't read oauth config") {
		t.Errorf("twitter-auth without config = %v", err)
	}
}

func TestMintSessionMissingPrivateKey(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("PKEY_HEX", "")
	_, err := executeCmd(t, newMintCmd(), testRWalkAddr.Hex(), "1000")
	if err == nil || !strings.Contains(err.Error(), "PKEY_HEX") {
		t.Errorf("mint without key = %v", err)
	}
}

func TestRankValue(t *testing.T) {
	if got := rankValue(0, 100); got != 1.0 {
		t.Errorf("rankValue(0,100) = %v, want 1.0 (top)", got)
	}
	if got := rankValue(50, 100); got != 51.0 {
		t.Errorf("rankValue(50,100) = %v, want 51.0", got)
	}
}
