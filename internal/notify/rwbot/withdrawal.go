package rwbot

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// ContractWithdrawalReader reads withdrawalAmount() from the RandomWalk
// contract.
type ContractWithdrawalReader struct {
	Contract *rwcontracts.RWalk
}

// WithdrawalAmountEth returns the current withdrawal amount converted from
// wei to ETH (the notification texts use the rounded ETH value).
func (r ContractWithdrawalReader) WithdrawalAmountEth(ctx context.Context) (float64, error) {
	amount, err := r.Contract.WithdrawalAmount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}
	f := new(big.Float).SetInt(amount)
	f = f.Quo(f, big.NewFloat(1e18))
	out, _ := f.Float64()
	return out, nil
}
