package toolutil

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

// CollectorConfig is read from the transaction-collector JSON config file.
type CollectorConfig struct {
	RPCURL     string `json:"rpc_url"`
	OutputDir  string `json:"output_dir"`
	StartBlock uint64 `json:"start_block"`

	ContractAddresses []string       `json:"contract_addresses"`
	Contracts         *ContractAddrs `json:"contracts"`
}

// ContractAddrs mirrors cg_contracts in game-mechanics.sql.
type ContractAddrs struct {
	CosmicGameAddr         string `json:"cosmic_game_addr"`
	CosmicSignatureAddr    string `json:"cosmic_signature_addr"`
	CosmicTokenAddr        string `json:"cosmic_token_addr"`
	CosmicDaoAddr          string `json:"cosmic_dao_addr"`
	CharityWalletAddr      string `json:"charity_wallet_addr"`
	PrizesWalletAddr       string `json:"prizes_wallet_addr"`
	RandomWalkAddr         string `json:"random_walk_addr"`
	StakingWalletCSTAddr   string `json:"staking_wallet_cst_addr"`
	StakingWalletRWalkAddr string `json:"staking_wallet_rwalk_addr"`
	MarketingWalletAddr    string `json:"marketing_wallet_addr"`
	ImplementationAddr     string `json:"implementation_addr"`
}

// LoadCollectorConfig reads and validates a collector JSON config file.
func LoadCollectorConfig(path string) (*CollectorConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg CollectorConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	if cfg.OutputDir == "" {
		return nil, fmt.Errorf("output_dir is required")
	}
	return &cfg, nil
}

// ResolveContractAddresses returns deduplicated checksummed hex addresses from config.
func (c *CollectorConfig) ResolveContractAddresses() ([]string, error) {
	seen := make(map[string]struct{})
	var out []string
	add := func(raw string) error {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			return nil
		}
		if !ethcommon.IsHexAddress(raw) {
			return fmt.Errorf("invalid address: %q", raw)
		}
		addr := ethcommon.HexToAddress(raw).Hex()
		if _, ok := seen[addr]; ok {
			return nil
		}
		seen[addr] = struct{}{}
		out = append(out, addr)
		return nil
	}

	for _, a := range c.ContractAddresses {
		if err := add(a); err != nil {
			return nil, err
		}
	}
	if c.Contracts != nil {
		ca := c.Contracts
		for _, a := range []string{
			ca.CosmicGameAddr,
			ca.CosmicSignatureAddr,
			ca.CosmicTokenAddr,
			ca.CosmicDaoAddr,
			ca.CharityWalletAddr,
			ca.PrizesWalletAddr,
			ca.RandomWalkAddr,
			ca.StakingWalletCSTAddr,
			ca.StakingWalletRWalkAddr,
			ca.MarketingWalletAddr,
			ca.ImplementationAddr,
		} {
			if err := add(a); err != nil {
				return nil, err
			}
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no contract addresses: set contract_addresses and/or contracts in config")
	}
	return out, nil
}
