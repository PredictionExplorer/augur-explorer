package rwbot

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

func TestContractWithdrawalReaderConvertsWeiToEth(t *testing.T) {
	chain := testchain.New(t)
	addr := common.HexToAddress("0x8950000000000000000000000000000000000001")
	wei, _ := new(big.Int).SetString("2500000000000000000", 10) // 2.5 ETH
	chain.RegisterCall(addr, testchain.MustContractStub(rwcontracts.RWalkMetaData.ABI).
		Return("withdrawalAmount", wei).Handler())

	ec, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(ec.Close)
	ctrct, err := rwcontracts.NewRWalk(addr, ec)
	if err != nil {
		t.Fatalf("NewRWalk: %v", err)
	}

	got, err := ContractWithdrawalReader{Contract: ctrct}.WithdrawalAmountEth(context.Background())
	if err != nil {
		t.Fatalf("WithdrawalAmountEth: %v", err)
	}
	if got != 2.5 {
		t.Errorf("amount = %v, want 2.5", got)
	}
}

func TestContractWithdrawalReaderPropagatesRevert(t *testing.T) {
	chain := testchain.New(t)
	addr := common.HexToAddress("0x8950000000000000000000000000000000000002")
	// No handler registered: the call reverts like a missing contract.
	ec, err := ethclient.Dial(chain.URL())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(ec.Close)
	ctrct, err := rwcontracts.NewRWalk(addr, ec)
	if err != nil {
		t.Fatalf("NewRWalk: %v", err)
	}

	if _, err := (ContractWithdrawalReader{Contract: ctrct}).WithdrawalAmountEth(context.Background()); err == nil ||
		!strings.Contains(err.Error(), "no call handler") {
		t.Errorf("err = %v, want the node's revert", err)
	}
}
