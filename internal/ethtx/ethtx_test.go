package ethtx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// testKeyHex is a throwaway signer key used by the session tests.
const testKeyHex = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

// testKeyAddr is the address of testKeyHex.
var testKeyAddr = mustKeyAddr(testKeyHex)

func mustKeyAddr(hexKey string) common.Address {
	key, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		panic(err)
	}
	return crypto.PubkeyToAddress(key.PublicKey)
}

// session builds a Session against a fresh fake chain with a funded account.
func session(t *testing.T, verbose bool) (*testchain.Chain, *Session, *bytes.Buffer) {
	t.Helper()
	chain := testchain.New(t)
	chain.EnsureBlock(10)
	chain.SetBalance(testKeyAddr, big.NewInt(1_000_000_000_000_000_000)) // 1 ETH
	chain.SetNonce(testKeyAddr, 3)

	var out bytes.Buffer
	s, err := New(context.Background(), Options{
		RPCURL:         chain.URL(),
		PrivateKeyHex:  testKeyHex,
		Verbose:        verbose,
		Out:            &out,
		ReceiptTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return chain, s, &out
}

func TestNewSessionFetchesNetworkAndAccountState(t *testing.T) {
	chain, s, out := session(t, true)

	if s.Net.ChainID.Cmp(chain.ChainID()) != 0 {
		t.Errorf("chain id = %s, want %s", s.Net.ChainID, chain.ChainID())
	}
	if s.Net.BlockNum.Int64() != 10 {
		t.Errorf("latest block = %s, want 10", s.Net.BlockNum)
	}
	if s.Acc.Address != testKeyAddr {
		t.Errorf("account address = %s, want %s", s.Acc.Address, testKeyAddr)
	}
	if s.Acc.Nonce != 3 {
		t.Errorf("nonce = %d, want 3", s.Acc.Nonce)
	}
	if s.Acc.Balance.Cmp(big.NewInt(1_000_000_000_000_000_000)) != 0 {
		t.Errorf("balance = %s, want 1 ETH in wei", s.Acc.Balance)
	}

	text := out.String()
	for _, want := range []string{
		"NETWORK INFO",
		"Chain ID            = 1337",
		"ACCOUNT INFO",
		"Address             = " + testKeyAddr.String(),
		"Balance (ETH)       = 1.000000000000000000",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("verbose output missing %q\noutput:\n%s", want, text)
		}
	}
}

func TestNewSessionQuietPrintsNothing(t *testing.T) {
	_, _, out := session(t, false)
	if out.Len() != 0 {
		t.Errorf("quiet session printed %q, want empty", out.String())
	}
}

func TestConnectRequiresEndpoint(t *testing.T) {
	if _, err := Connect(context.Background(), ""); err == nil {
		t.Error("Connect with empty URL succeeded")
	}
}

func TestConnectUnreachableEndpoint(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if _, err := Connect(ctx, "http://127.0.0.1:1"); err == nil {
		t.Error("Connect to unreachable endpoint succeeded")
	}
}

func TestPrepareAccountRejectsBadKey(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	net, err := Connect(context.Background(), chain.URL())
	if err != nil {
		t.Fatalf("Connect: %v", err)
	}
	if _, err := PrepareAccount(context.Background(), net, "not-hex"); err == nil {
		t.Error("bad private key accepted")
	}
	if _, err := PrepareAccount(context.Background(), net, "abcd"); err == nil {
		t.Error("short private key accepted")
	}
}

func TestNewSessionRejectsBadKey(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	_, err := New(context.Background(), Options{RPCURL: chain.URL(), PrivateKeyHex: "zz"})
	if err == nil || !strings.Contains(err.Error(), "account setup failed") {
		t.Errorf("err = %v, want account setup failure", err)
	}
}

func TestAdjustGasPriceDoubles(t *testing.T) {
	if got := AdjustGasPrice(nil); got.Sign() != 0 {
		t.Errorf("AdjustGasPrice(nil) = %s, want 0", got)
	}
	got := AdjustGasPrice(big.NewInt(1_500_000_000))
	if got.Cmp(big.NewInt(3_000_000_000)) != 0 {
		t.Errorf("AdjustGasPrice(1.5 gwei) = %s, want 3 gwei", got)
	}
}

// signViaOpts builds a legacy transaction and signs it through the session's
// TransactOpts signer, the exact path the abigen bindings use.
func signViaOpts(t *testing.T, s *Session, to common.Address, value *big.Int, gasLimit uint64) *types.Transaction {
	t.Helper()
	opts := s.TransactOpts(value, gasLimit)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    opts.Nonce.Uint64(),
		GasPrice: opts.GasPrice,
		Gas:      opts.GasLimit,
		To:       &to,
		Value:    opts.Value,
	})
	signed, err := opts.Signer(opts.From, tx)
	if err != nil {
		t.Fatalf("signing through TransactOpts: %v", err)
	}
	return signed
}

func TestTransactOptsSignsEIP155ForNetworkChain(t *testing.T) {
	_, s, _ := session(t, false)
	to := common.HexToAddress("0x2200000000000000000000000000000000000022")

	signed := signViaOpts(t, s, to, big.NewInt(5), GasLimitHighComplexity)

	if signed.ChainId().Cmp(s.Net.ChainID) != 0 {
		t.Errorf("signed chain id = %s, want %s", signed.ChainId(), s.Net.ChainID)
	}
	sender, err := types.Sender(types.NewEIP155Signer(s.Net.ChainID), signed)
	if err != nil {
		t.Fatalf("recovering sender: %v", err)
	}
	if sender != s.Acc.Address {
		t.Errorf("recovered sender = %s, want %s", sender, s.Acc.Address)
	}
	// The 2.0x gas policy is applied to the network's suggested price.
	want := AdjustGasPrice(s.Net.GasPrice)
	if signed.GasPrice().Cmp(want) != 0 {
		t.Errorf("gas price = %s, want doubled %s", signed.GasPrice(), want)
	}
}

func TestFinishTxSuccessQuietAndVerbose(t *testing.T) {
	for _, verbose := range []bool{false, true} {
		_, s, out := session(t, verbose)
		to := common.HexToAddress("0x2200000000000000000000000000000000000022")
		signed := signViaOpts(t, s, to, big.NewInt(0), GasLimitContractCall)
		if err := s.Net.Client.SendTransaction(context.Background(), signed); err != nil {
			t.Fatalf("SendTransaction: %v", err)
		}

		if err := s.FinishTx(context.Background(), signed, nil); err != nil {
			t.Fatalf("FinishTx (verbose=%v): %v", verbose, err)
		}
		text := out.String()
		if verbose {
			if !strings.Contains(text, "Status              = SUCCESS") || !strings.Contains(text, "Gas Used") {
				t.Errorf("verbose success output missing sections:\n%s", text)
			}
		} else if !strings.Contains(text, "Success. Tx hash = "+signed.Hash().String()) {
			t.Errorf("quiet success output = %q", text)
		}
	}
}

func TestFinishTxReportsSendError(t *testing.T) {
	_, s, out := session(t, false)
	err := s.FinishTx(context.Background(), nil, errors.New("insufficient funds"))
	if err == nil {
		t.Fatal("FinishTx with send error returned nil")
	}
	if !strings.Contains(out.String(), "insufficient funds") {
		t.Errorf("output = %q, want the send error", out.String())
	}
}

func TestFinishTxReportsNilTransaction(t *testing.T) {
	_, s, out := session(t, false)
	if err := s.FinishTx(context.Background(), nil, nil); err == nil {
		t.Fatal("FinishTx with nil tx returned nil")
	}
	if !strings.Contains(out.String(), "transaction is nil") {
		t.Errorf("output = %q, want nil-transaction report", out.String())
	}
}

func TestFinishTxReportsRevert(t *testing.T) {
	for _, verbose := range []bool{false, true} {
		chain, s, out := session(t, verbose)
		to := common.HexToAddress("0x2200000000000000000000000000000000000022")
		signed := signViaOpts(t, s, to, big.NewInt(0), GasLimitContractCall)
		chain.MarkNextTxReverted()
		if err := s.Net.Client.SendTransaction(context.Background(), signed); err != nil {
			t.Fatalf("SendTransaction: %v", err)
		}

		if err := s.FinishTx(context.Background(), signed, nil); err == nil {
			t.Fatal("FinishTx for reverted tx returned nil")
		}
		text := out.String()
		if verbose {
			if !strings.Contains(text, "Status              = REVERTED") {
				t.Errorf("verbose revert output:\n%s", text)
			}
		} else if !strings.Contains(text, "Transaction reverted on-chain") {
			t.Errorf("quiet revert output = %q", text)
		}
	}
}

func TestFinishTxReportsReceiptTimeout(t *testing.T) {
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	chain.SetBalance(testKeyAddr, big.NewInt(1_000_000_000_000_000_000))
	var out bytes.Buffer
	s, err := New(context.Background(), Options{
		RPCURL:         chain.URL(),
		PrivateKeyHex:  testKeyHex,
		Out:            &out,
		ReceiptTimeout: 50 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	to := common.HexToAddress("0x2200000000000000000000000000000000000022")
	signed := signViaOpts(t, s, to, big.NewInt(0), GasLimitContractCall)
	chain.MarkNextTxPending()
	if err := s.Net.Client.SendTransaction(context.Background(), signed); err != nil {
		t.Fatalf("SendTransaction: %v", err)
	}

	if err := s.FinishTx(context.Background(), signed, nil); err == nil {
		t.Fatal("FinishTx without a receipt returned nil")
	}
	if !strings.Contains(out.String(), "receipt not received") {
		t.Errorf("output = %q, want receipt-timeout report", out.String())
	}
}

// failingRPC serves eth JSON-RPC answering every method except failMethod,
// which errors; it covers the per-call failure branches of Connect and
// PrepareAccount.
func failingRPC(t *testing.T, failMethod string) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     json.RawMessage   `json:"id"`
			Method string            `json:"method"`
			Params []json.RawMessage `json:"params"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		resp := map[string]any{"jsonrpc": "2.0", "id": req.ID}
		if req.Method == failMethod {
			resp["error"] = map[string]any{"code": -32000, "message": "injected failure"}
		} else {
			switch req.Method {
			case "eth_chainId", "eth_gasPrice", "eth_getTransactionCount", "eth_getBalance":
				resp["result"] = "0x1"
			case "eth_getBlockByNumber":
				resp["result"] = map[string]any{
					"number": "0x1", "parentHash": "0x" + strings.Repeat("00", 32),
					"sha3Uncles": "0x" + strings.Repeat("00", 32), "miner": "0x" + strings.Repeat("00", 20),
					"stateRoot": "0x" + strings.Repeat("00", 32), "transactionsRoot": "0x" + strings.Repeat("00", 32),
					"receiptsRoot": "0x" + strings.Repeat("00", 32), "logsBloom": "0x" + strings.Repeat("00", 256),
					"difficulty": "0x1", "gasLimit": "0x1", "gasUsed": "0x0", "timestamp": "0x1",
					"extraData": "0x", "mixHash": "0x" + strings.Repeat("00", 32), "nonce": "0x0000000000000000",
					"hash": "0x" + strings.Repeat("11", 32),
				}
			default:
				resp["error"] = map[string]any{"code": -32601, "message": "unsupported"}
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	t.Cleanup(srv.Close)
	return srv.URL
}

func TestConnectPerCallFailures(t *testing.T) {
	cases := []struct {
		fail string
		want string
	}{
		{"eth_chainId", "chain ID"},
		{"eth_gasPrice", "suggested gas price"},
		{"eth_getBlockByNumber", "latest block"},
	}
	for _, tc := range cases {
		t.Run(tc.fail, func(t *testing.T) {
			_, err := Connect(context.Background(), failingRPC(t, tc.fail))
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("Connect with failing %s = %v, want mention of %q", tc.fail, err, tc.want)
			}
		})
	}
}

func TestConnectRejectsMalformedURL(t *testing.T) {
	if _, err := Connect(context.Background(), "://not-a-url"); err == nil {
		t.Error("malformed URL accepted")
	}
}

func TestNewWrapsConnectFailure(t *testing.T) {
	_, err := New(context.Background(), Options{RPCURL: "", PrivateKeyHex: testKeyHex})
	if err == nil || !strings.Contains(err.Error(), "network connection failed") {
		t.Errorf("New without endpoint = %v", err)
	}
}

func TestPrepareAccountPerCallFailures(t *testing.T) {
	for _, tc := range []struct {
		fail string
		want string
	}{
		{"eth_getTransactionCount", "account nonce"},
		{"eth_getBalance", "account balance"},
	} {
		t.Run(tc.fail, func(t *testing.T) {
			net, err := Connect(context.Background(), failingRPC(t, tc.fail))
			if err != nil {
				t.Fatalf("Connect: %v", err)
			}
			if _, err := PrepareAccount(context.Background(), net, testKeyHex); err == nil ||
				!strings.Contains(err.Error(), tc.want) {
				t.Errorf("PrepareAccount with failing %s = %v", tc.fail, err)
			}
		})
	}
}

func TestFinishTxVerboseErrorReports(t *testing.T) {
	t.Run("send error", func(t *testing.T) {
		_, s, out := session(t, true)
		if err := s.FinishTx(context.Background(), nil, errors.New("nonce too low")); err == nil {
			t.Fatal("FinishTx returned nil")
		}
		if !strings.Contains(out.String(), "Status              = FAILED") ||
			!strings.Contains(out.String(), "nonce too low") {
			t.Errorf("verbose send-error output:\n%s", out.String())
		}
	})
	t.Run("nil transaction", func(t *testing.T) {
		_, s, out := session(t, true)
		if err := s.FinishTx(context.Background(), nil, nil); err == nil {
			t.Fatal("FinishTx returned nil")
		}
		if !strings.Contains(out.String(), "Error               = transaction is nil") {
			t.Errorf("verbose nil-tx output:\n%s", out.String())
		}
	})
	t.Run("receipt wait failure", func(t *testing.T) {
		chain := testchain.New(t)
		chain.EnsureBlock(1)
		chain.SetBalance(testKeyAddr, big.NewInt(1_000_000_000_000_000_000))
		var out bytes.Buffer
		s, err := New(context.Background(), Options{
			RPCURL: chain.URL(), PrivateKeyHex: testKeyHex, Verbose: true,
			Out: &out, ReceiptTimeout: 50 * time.Millisecond,
		})
		if err != nil {
			t.Fatalf("New: %v", err)
		}
		to := common.HexToAddress("0x2200000000000000000000000000000000000022")
		signed := signViaOpts(t, s, to, big.NewInt(0), GasLimitContractCall)
		chain.MarkNextTxPending()
		if err := s.Net.Client.SendTransaction(context.Background(), signed); err != nil {
			t.Fatalf("SendTransaction: %v", err)
		}
		if err := s.FinishTx(context.Background(), signed, nil); err == nil {
			t.Fatal("FinishTx returned nil without a receipt")
		}
		if !strings.Contains(out.String(), "SUBMITTED (receipt wait failed)") {
			t.Errorf("verbose receipt-wait output:\n%s", out.String())
		}
	})
}

func TestVerboseTxSubmittingSection(t *testing.T) {
	var out bytes.Buffer
	o := Output{Verbose: true, W: &out}
	o.TxSubmitting("Mint", big.NewInt(2_000_000_000_000_000_000), 5000000, big.NewInt(2_000_000_000))
	text := out.String()
	for _, want := range []string{
		"SUBMITTING TRANSACTION",
		"Action              = Mint",
		"Value (ETH)         = 2.000000000000000000",
		"Gas Limit           = 5000000",
		"Max Gas Cost (ETH)  = 0.010000",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("TxSubmitting output missing %q\noutput:\n%s", want, text)
		}
	}
}

func TestQuietHelpersPrintNothing(t *testing.T) {
	var out bytes.Buffer
	o := Output{Verbose: false, W: &out}
	o.Section("X")
	o.KeyValue("k", "v")
	o.KeyValueEth("k", big.NewInt(1))
	o.TxSubmitting("A", nil, 1, big.NewInt(1))
	o.NetworkInfo(&Network{ChainID: big.NewInt(1), GasPrice: big.NewInt(1), BlockNum: big.NewInt(1)})
	o.AccountInfo(&Account{Balance: big.NewInt(1)})
	if out.Len() != 0 {
		t.Errorf("quiet helpers printed %q", out.String())
	}
}

func TestWeiConversions(t *testing.T) {
	if got := WeiToEthText(nil); got != "0.000000000000000000" {
		t.Errorf("WeiToEthText(nil) = %q", got)
	}
	wei, _ := new(big.Int).SetString("1234500000000000000", 10)
	if got := WeiToEthText(wei); got != "1.234500000000000000" {
		t.Errorf("WeiToEthText = %q, want 1.2345 ETH", got)
	}
	if got := WeiToGwei(nil); got != 0 {
		t.Errorf("WeiToGwei(nil) = %v", got)
	}
	if got := WeiToGwei(big.NewInt(2_500_000_000)); got != 2.5 {
		t.Errorf("WeiToGwei(2.5 gwei) = %v", got)
	}
}
