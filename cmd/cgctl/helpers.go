package main

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

// defaultLocalGameAddr is the first contract address deployed by a fresh
// Hardhat node; read-only commands fall back to it when no address is given.
const defaultLocalGameAddr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"

// parseAddress validates that s is a hex Ethereum address and returns it.
func parseAddress(name, s string) (common.Address, error) {
	if !common.IsHexAddress(s) {
		return common.Address{}, fmt.Errorf("invalid %s: %q is not a hex address", name, s)
	}
	return common.HexToAddress(s), nil
}

// parseBigInt parses a base-10 integer of arbitrary size.
func parseBigInt(name, s string) (*big.Int, error) {
	v := new(big.Int)
	if _, ok := v.SetString(s, 10); !ok {
		return nil, fmt.Errorf("invalid %s value provided: %s", name, s)
	}
	return v, nil
}

// parseInt64 parses a base-10 int64.
func parseInt64(name, s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing %s: %w", name, err)
	}
	return v, nil
}
