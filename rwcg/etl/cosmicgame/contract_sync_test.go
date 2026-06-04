package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TestProbeContractMechanicsUnknownWithoutBindings(t *testing.T) {
	got := probeContractMechanics(nil, nil, &bind.CallOpts{})
	if got != contractMechanicsUnknown {
		t.Fatalf("expected unknown mechanics, got %d", got)
	}
}

func TestReadCSTAuctionDurationChangeDivisorV1(t *testing.T) {
	got := readCSTAuctionDurationChangeDivisor(nil, &bind.CallOpts{}, contractMechanicsV1)
	if got != -1 {
		t.Fatalf("expected -1 on V1, got %d", got)
	}
}
