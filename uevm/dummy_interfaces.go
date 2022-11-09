package uevm				// EVM for Uniswap (v3)

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/consensus"
)
/* the following interfaces are dummy, then do not return any meaningful data, we only
	need them to make a dummy vm.Call() to simulated instance of Uniswap v3 contract to reproduce
	the Swap (also Join/Mint) math and obtain the same values as on MainNet (while using our own stateDB)
	Since Uniswap contract doesn't rely on global EVM variables we can use dummy values
*/

type DummyChainContext struct {
	dummy_engine		*DummyEngine
}
func NewDummyChainContext() *DummyChainContext {

	output := new(DummyChainContext)
	return output
}
func (this *DummyChainContext) Engine() consensus.Engine {
	return this.dummy_engine
}
type DummyEngine struct {

}
func NewDummyEngine() *DummyEngine {

	return new(DummyEngine)
}
func (this *DummyEngine) Author(header *types.Header) (common.Address, error) {

	return common.Address{},nil
}
func (this *DummyEngine) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {

	return nil
}
func (this *DummyEngine) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {

	panic("VerifyHeaders() is not implemented")
	return nil,nil
}
func (this *DummyEngine) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {

	return nil
}
func (this *DummyEngine) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {

	return nil
}
func (this *DummyEngine) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,uncles []*types.Header) {

}
func (this *DummyEngine) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	return nil,nil
}
func (this *DummyEngine) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	return nil
}
func (this *DummyEngine)  SealHash(header *types.Header) common.Hash {
	return common.Hash{}
}
func (this *DummyEngine) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return big.NewInt(1)
}
func (this *DummyEngine) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return nil
}
func (this *DummyEngine) Close() error {
	return nil
} 
func CustomGetHashFn(ref *types.Header, chain core.ChainContext) func(n uint64) common.Hash {
	return func(n uint64) common.Hash {
		return common.Hash{}
	}
}

func NewDummyBlockContext(bnum,t *big.Int) *vm.BlockContext {

	output := new(vm.BlockContext)
	output.BlockNumber = big.NewInt(0).Set(bnum)
	output.Time = big.NewInt(0).Set(t)

	output.GasLimit = 10000000000
	output.Difficulty = big.NewInt(1)
	output.BaseFee = big.NewInt(1)
	//output.Random = new(common.Hash)

	output.CanTransfer = func(vm.StateDB, common.Address, *big.Int) bool { return true }
	output.Transfer = func(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
			db.SubBalance(sender, amount)
			db.AddBalance(recipient, amount)
			negative := big.NewInt(0).Cmp(db.GetBalance(sender)) > 0
			if negative {
				panic(fmt.Sprintf("failed transferring amount %v from %v to %v , result is negative balance %v",amount.String(),sender.String(),recipient.String(),db.GetBalance(sender).String()))
			}
	}
	output.GetHash = func (n uint64) common.Hash { return common.Hash{}}
	return output
}
