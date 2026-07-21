package testchain

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// client dials the fake chain through the real go-ethereum client, which is
// exactly how production code will consume it.
func client(t *testing.T) (*Chain, *ethclient.Client) {
	t.Helper()
	chain := New(t)
	ec, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dialing fake chain: %v", err)
	}
	t.Cleanup(ec.Close)
	return chain, ec
}

func TestHeaderRoundTripPreservesHash(t *testing.T) {
	chain, ec := client(t)
	want := chain.EnsureBlock(42)

	got, err := ec.HeaderByNumber(context.Background(), big.NewInt(42))
	if err != nil {
		t.Fatalf("HeaderByNumber: %v", err)
	}
	if got.Hash() != want.Hash() {
		t.Errorf("hash mismatch after JSON round trip: server %s, client %s", want.Hash(), got.Hash())
	}
	if got.Time != BlockTime(42) {
		t.Errorf("timestamp = %d, want %d", got.Time, BlockTime(42))
	}
}

func TestHeaderChainsToParent(t *testing.T) {
	chain, _ := client(t)
	parent := chain.EnsureBlock(10)
	child := chain.EnsureBlock(11)
	if child.ParentHash != parent.Hash() {
		t.Errorf("block 11 parent hash = %s, want %s", child.ParentHash, parent.Hash())
	}
}

func TestMissingBlockIsNotFound(t *testing.T) {
	_, ec := client(t)
	_, err := ec.HeaderByNumber(context.Background(), big.NewInt(999))
	if !errors.Is(err, ethereum.NotFound) {
		t.Errorf("missing block error = %v, want ethereum.NotFound", err)
	}
}

// TestParseBlockTagRejectsOverflowingHeight pins the fake node's guard: a
// hex block tag beyond int64 answers an error instead of wrapping into a
// negative height.
func TestParseBlockTagRejectsOverflowingHeight(t *testing.T) {
	chain := New(t)
	_, err := chain.parseBlockTag(json.RawMessage(`"0xffffffffffffffff"`))
	if err == nil || !strings.Contains(err.Error(), "overflows int64") {
		t.Fatalf("parseBlockTag error = %v, want int64 overflow rejection", err)
	}
	height, err := chain.parseBlockTag(json.RawMessage(`"0x2a"`))
	if err != nil || height != 42 {
		t.Fatalf("parseBlockTag(0x2a) = %d, %v; want 42", height, err)
	}
}

func TestTransactionRoundTripRecoversSender(t *testing.T) {
	chain, ec := client(t)
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := chain.AddTx(100, to, []byte{0xde, 0xad, 0xbe, 0xef})

	got, pending, err := ec.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("TransactionByHash: %v", err)
	}
	if pending {
		t.Error("transaction reported as pending")
	}
	// The production insert path recovers the sender from the signature; it
	// must yield the chain's fixed signer.
	from, err := types.Sender(types.LatestSignerForChainID(got.ChainId()), got)
	if err != nil {
		t.Fatalf("recovering sender: %v", err)
	}
	if from != chain.Sender() {
		t.Errorf("recovered sender = %s, want %s", from, chain.Sender())
	}
	if *got.To() != to {
		t.Errorf("to = %s, want %s", got.To(), to)
	}
}

func TestMissingTransactionIsNotFound(t *testing.T) {
	_, ec := client(t)
	_, _, err := ec.TransactionByHash(context.Background(), common.HexToHash("0xdead"))
	if !errors.Is(err, ethereum.NotFound) {
		t.Errorf("missing tx error = %v, want ethereum.NotFound", err)
	}
}

func TestReceiptCarriesAttachedLogs(t *testing.T) {
	chain, ec := client(t)
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := chain.AddTx(100, to, nil)
	chain.AttachLogs(tx.Hash(), []*types.Log{{
		Address:     to,
		Topics:      []common.Hash{common.HexToHash("0x01")},
		Data:        []byte{0x02},
		BlockNumber: 100,
		TxHash:      tx.Hash(),
		Index:       0,
	}})

	receipt, err := ec.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("TransactionReceipt: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		t.Errorf("status = %d, want successful", receipt.Status)
	}
	if len(receipt.Logs) != 1 || receipt.Logs[0].Address != to {
		t.Errorf("receipt logs = %+v, want the attached log", receipt.Logs)
	}
}

func TestFilterLogsByRangeAndAddress(t *testing.T) {
	chain, ec := client(t)
	addrA := common.HexToAddress("0xaa00000000000000000000000000000000000001")
	addrB := common.HexToAddress("0xbb00000000000000000000000000000000000002")
	tx1 := chain.AddTx(10, addrA, nil)
	tx2 := chain.AddTx(20, addrB, nil)
	chain.AttachLogs(tx1.Hash(), []*types.Log{{Address: addrA, Topics: []common.Hash{{}}, BlockNumber: 10, TxHash: tx1.Hash()}})
	chain.AttachLogs(tx2.Hash(), []*types.Log{{Address: addrB, Topics: []common.Hash{{}}, BlockNumber: 20, TxHash: tx2.Hash()}})

	logs, err := ec.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(1),
		ToBlock:   big.NewInt(30),
		Addresses: []common.Address{addrA},
	})
	if err != nil {
		t.Fatalf("FilterLogs: %v", err)
	}
	if len(logs) != 1 || logs[0].Address != addrA {
		t.Errorf("filtered logs = %+v, want exactly the addrA log", logs)
	}

	logs, err = ec.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(15),
		ToBlock:   big.NewInt(30),
	})
	if err != nil {
		t.Fatalf("FilterLogs: %v", err)
	}
	if len(logs) != 1 || logs[0].BlockNumber != 20 {
		t.Errorf("range-filtered logs = %+v, want only block 20", logs)
	}
}

func TestBlockNumberTracksTip(t *testing.T) {
	chain, ec := client(t)
	chain.EnsureBlock(77)
	tip, err := ec.BlockNumber(context.Background())
	if err != nil {
		t.Fatalf("BlockNumber: %v", err)
	}
	if tip != 77 {
		t.Errorf("tip = %d, want 77", tip)
	}
}

func TestEthCallDispatchesToHandler(t *testing.T) {
	chain, ec := client(t)
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	chain.RegisterCall(contract, func(input []byte) ([]byte, error) {
		return append([]byte{0x42}, input...), nil
	})

	out, err := ec.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: []byte{0x01, 0x02},
	}, nil)
	if err != nil {
		t.Fatalf("CallContract: %v", err)
	}
	if len(out) != 3 || out[0] != 0x42 || out[1] != 0x01 || out[2] != 0x02 {
		t.Errorf("call output = %x, want 420102", out)
	}
}

func TestEthCallWithoutHandlerErrors(t *testing.T) {
	_, ec := client(t)
	to := common.HexToAddress("0xcc00000000000000000000000000000000000003")
	if _, err := ec.CallContract(context.Background(), ethereum.CallMsg{To: &to}, nil); err == nil {
		t.Error("expected error for unregistered call handler")
	}
}

func TestReorgChangesHashesAndDropsState(t *testing.T) {
	chain, ec := client(t)
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	chain.EnsureBlock(5)
	tx := chain.AddTx(6, to, nil)
	chain.AttachLogs(tx.Hash(), []*types.Log{{Address: to, Topics: []common.Hash{{}}, BlockNumber: 6, TxHash: tx.Hash()}})

	hash5 := chain.BlockHash(5)
	hash6 := chain.BlockHash(6)

	chain.Reorg(6)

	if chain.BlockHash(5) != hash5 {
		t.Error("block 5 hash changed even though the reorg starts at 6")
	}
	if chain.BlockHash(6) == hash6 {
		t.Error("block 6 hash unchanged after reorg")
	}
	if _, _, err := ec.TransactionByHash(context.Background(), tx.Hash()); !errors.Is(err, ethereum.NotFound) {
		t.Errorf("tx in reorged block: err = %v, want ethereum.NotFound", err)
	}
	logs, err := ec.FilterLogs(context.Background(), ethereum.FilterQuery{FromBlock: big.NewInt(0), ToBlock: big.NewInt(100)})
	if err != nil {
		t.Fatalf("FilterLogs: %v", err)
	}
	if len(logs) != 0 {
		t.Errorf("logs from reorged range survived: %+v", logs)
	}

	// Deterministic tx hashes: re-adding the same tx spec reproduces state.
	tx2 := chain.AddTx(6, to, nil)
	if tx2.Hash() != tx.Hash() {
		t.Errorf("re-added tx hash %s differs from original %s (nonce-keyed determinism broken)", tx2.Hash(), tx.Hash())
	}
}

func TestClearTransactionsFromPreservesBlockIdentity(t *testing.T) {
	chain, ec := client(t)
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := chain.AddTx(6, to, nil)
	chain.AttachLogs(tx.Hash(), []*types.Log{{
		Address:     to,
		Topics:      []common.Hash{{}},
		BlockNumber: 6,
		TxHash:      tx.Hash(),
	}})
	blockHash := chain.BlockHash(6)

	chain.ClearTransactionsFrom(6)

	if chain.BlockHash(6) != blockHash {
		t.Fatal("fixture cleanup changed the block hash")
	}
	if _, _, err := ec.TransactionByHash(context.Background(), tx.Hash()); !errors.Is(err, ethereum.NotFound) {
		t.Errorf("cleared transaction: err = %v, want ethereum.NotFound", err)
	}
	logs, err := ec.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(6),
		ToBlock:   big.NewInt(6),
	})
	if err != nil {
		t.Fatalf("FilterLogs: %v", err)
	}
	if len(logs) != 0 {
		t.Fatalf("cleared logs survived: %+v", logs)
	}
	if repeated := chain.AddTx(6, to, nil); repeated.Hash() != tx.Hash() {
		t.Fatalf("re-added transaction hash = %s, want %s", repeated.Hash(), tx.Hash())
	}
}

func TestRepeatedEnsureBlockIsStable(t *testing.T) {
	chain, _ := client(t)
	h1 := chain.BlockHash(9)
	h2 := chain.BlockHash(9)
	if h1 != h2 {
		t.Errorf("EnsureBlock not idempotent: %s vs %s", h1, h2)
	}
}
