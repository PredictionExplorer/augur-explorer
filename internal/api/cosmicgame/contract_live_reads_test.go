package cosmicgame

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TestContractMechanicsVersionConstants(t *testing.T) {
	if contractMechanicsUnknown != 0 || contractMechanicsV1 != 1 || contractMechanicsV2 != 2 {
		t.Fatalf("unexpected mechanics version constants")
	}
}

func TestReadRoundStartWithoutBindings(t *testing.T) {
	got := readRoundStartCSTAuctionSetting(nil, nil, &bind.CallOpts{})
	if got != -1 {
		t.Fatalf("expected -1 without bindings, got %d", got)
	}
}

func TestReadChangeDivisorWithoutBindings(t *testing.T) {
	got := readCSTAuctionDurationChangeDivisor(nil, nil, &bind.CallOpts{})
	if got != -1 {
		t.Fatalf("expected -1 without bindings, got %d", got)
	}
}

func TestReadTokenRewardWithoutBindings(t *testing.T) {
	_, err := readTokenReward(nil, nil, &bind.CallOpts{})
	if err == nil {
		t.Fatal("expected error without bindings")
	}
}
