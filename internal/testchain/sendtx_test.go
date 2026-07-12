package testchain

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// signedTx builds a legacy transaction signed with a throwaway key for the
// chain's chain id.
func signedTx(t *testing.T, chain *Chain, nonce uint64, to common.Address) (*types.Transaction, common.Address) {
	t.Helper()
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generating key: %v", err)
	}
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(2_000_000_000),
		Gas:      100_000,
		To:       &to,
		Value:    big.NewInt(7),
	})
	signed, err := types.SignTx(tx, types.LatestSignerForChainID(chain.ChainID()), key)
	if err != nil {
		t.Fatalf("signing tx: %v", err)
	}
	return signed, crypto.PubkeyToAddress(key.PublicKey)
}

func TestSuggestGasPriceDefaultAndOverride(t *testing.T) {
	chain, ec := client(t)

	got, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatalf("SuggestGasPrice: %v", err)
	}
	if got.Cmp(defaultGasPrice) != 0 {
		t.Errorf("default gas price = %s, want %s", got, defaultGasPrice)
	}

	chain.SetGasPrice(big.NewInt(42_000_000_000))
	got, err = ec.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatalf("SuggestGasPrice after override: %v", err)
	}
	if got.Cmp(big.NewInt(42_000_000_000)) != 0 {
		t.Errorf("overridden gas price = %s, want 42 gwei", got)
	}
}

func TestBalanceDefaultZeroAndSettable(t *testing.T) {
	chain, ec := client(t)
	addr := common.HexToAddress("0x1100000000000000000000000000000000000011")

	got, err := ec.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		t.Fatalf("BalanceAt: %v", err)
	}
	if got.Sign() != 0 {
		t.Errorf("default balance = %s, want 0", got)
	}

	want := new(big.Int)
	want.SetString("123456789012345678901", 10)
	chain.SetBalance(addr, want)
	got, err = ec.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		t.Fatalf("BalanceAt after SetBalance: %v", err)
	}
	if got.Cmp(want) != 0 {
		t.Errorf("balance = %s, want %s", got, want)
	}
}

func TestNonceSettableAndAdvancedBySend(t *testing.T) {
	chain, ec := client(t)
	to := common.HexToAddress("0x2200000000000000000000000000000000000022")

	tx, sender := signedTx(t, chain, 5, to)
	chain.SetNonce(sender, 5)

	nonce, err := ec.PendingNonceAt(context.Background(), sender)
	if err != nil {
		t.Fatalf("PendingNonceAt: %v", err)
	}
	if nonce != 5 {
		t.Errorf("set nonce = %d, want 5", nonce)
	}

	if err := ec.SendTransaction(context.Background(), tx); err != nil {
		t.Fatalf("SendTransaction: %v", err)
	}
	nonce, err = ec.PendingNonceAt(context.Background(), sender)
	if err != nil {
		t.Fatalf("PendingNonceAt after send: %v", err)
	}
	if nonce != 6 {
		t.Errorf("nonce after send = %d, want 6", nonce)
	}
}

func TestSendRawTransactionMinesAndServesReceipt(t *testing.T) {
	chain, ec := client(t)
	chain.EnsureBlock(50)
	to := common.HexToAddress("0x2200000000000000000000000000000000000022")
	tx, _ := signedTx(t, chain, 0, to)

	if err := ec.SendTransaction(context.Background(), tx); err != nil {
		t.Fatalf("SendTransaction: %v", err)
	}

	receipt, err := ec.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("TransactionReceipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		t.Errorf("status = %d, want successful", receipt.Status)
	}
	if receipt.BlockNumber.Int64() != 51 {
		t.Errorf("mined block = %s, want 51 (tip+1)", receipt.BlockNumber)
	}

	got, pending, err := ec.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("TransactionByHash: %v", err)
	}
	if pending {
		t.Error("submitted tx reported pending")
	}
	if got.Hash() != tx.Hash() {
		t.Errorf("round-tripped hash = %s, want %s", got.Hash(), tx.Hash())
	}
}

func TestMarkNextTxRevertedProducesFailedReceipt(t *testing.T) {
	chain, ec := client(t)
	to := common.HexToAddress("0x2200000000000000000000000000000000000022")
	tx, _ := signedTx(t, chain, 0, to)

	chain.MarkNextTxReverted()
	if err := ec.SendTransaction(context.Background(), tx); err != nil {
		t.Fatalf("SendTransaction: %v", err)
	}
	receipt, err := ec.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("TransactionReceipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusFailed {
		t.Errorf("status = %d, want failed (0)", receipt.Status)
	}

	// One-shot: the next transaction succeeds again.
	tx2, _ := signedTx(t, chain, 1, to)
	if err := ec.SendTransaction(context.Background(), tx2); err != nil {
		t.Fatalf("SendTransaction 2: %v", err)
	}
	receipt2, err := ec.TransactionReceipt(context.Background(), tx2.Hash())
	if err != nil {
		t.Fatalf("TransactionReceipt 2: %v", err)
	}
	if receipt2.Status != types.ReceiptStatusSuccessful {
		t.Errorf("second tx status = %d, want successful", receipt2.Status)
	}
}

func TestMarkNextTxPendingWithholdsReceiptUntilRelease(t *testing.T) {
	chain, ec := client(t)
	to := common.HexToAddress("0x2200000000000000000000000000000000000022")
	tx, _ := signedTx(t, chain, 0, to)

	chain.MarkNextTxPending()
	if err := ec.SendTransaction(context.Background(), tx); err != nil {
		t.Fatalf("SendTransaction: %v", err)
	}
	if _, err := ec.TransactionReceipt(context.Background(), tx.Hash()); err == nil {
		t.Fatal("receipt served while tx held pending; want not-found")
	}

	chain.ReleasePendingTxs()
	receipt, err := ec.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("TransactionReceipt after release: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		t.Errorf("released receipt status = %d, want successful", receipt.Status)
	}
}

func TestRejectNextSendWithSurfacesNodeError(t *testing.T) {
	chain, ec := client(t)
	to := common.HexToAddress("0x2200000000000000000000000000000000000022")
	tx, _ := signedTx(t, chain, 0, to)

	chain.RejectNextSendWith("insufficient funds for gas * price + value")
	err := ec.SendTransaction(context.Background(), tx)
	if err == nil || err.Error() != "insufficient funds for gas * price + value" {
		t.Fatalf("send error = %v, want the injected node rejection", err)
	}

	// One-shot: the retry succeeds.
	if err := ec.SendTransaction(context.Background(), tx); err != nil {
		t.Fatalf("SendTransaction retry: %v", err)
	}
}

func TestSendRawTransactionRejectsGarbage(t *testing.T) {
	_, ec := client(t)
	rpcClient := ec.Client()
	var result string
	if err := rpcClient.CallContext(context.Background(), &result, "eth_sendRawTransaction", "0xzz"); err == nil {
		t.Error("bad hex accepted")
	}
	if err := rpcClient.CallContext(context.Background(), &result, "eth_sendRawTransaction", "0xdeadbeef"); err == nil {
		t.Error("undecodable tx accepted")
	}
}
