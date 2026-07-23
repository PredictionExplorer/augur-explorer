// Bounded chain RPC: every call the engine makes carries a deadline. The
// engine's failure machinery (backoff, adaptive batch shrinking, the
// consecutive-failure circuit breaker) only works on calls that RETURN — a
// black-holed JSON-RPC endpoint used to hang the ETL forever, invisibly,
// because a call that never completes never counts as a failure. Wrapping
// the Client seam bounds every present and future call site in one place,
// including the in-transaction lookups that run on the shutdown-immune
// WithoutCancel context (finish-the-batch semantics stay, but "the batch"
// can no longer take forever).

package indexer

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DefaultRPCCallTimeout bounds one chain RPC call (D22). Point reads finish
// in milliseconds; the slowest legitimate call is an eth_getLogs over an
// adaptive range of up to a million blocks, which providers answer well
// under a minute or reject outright — and a timed-out fetch shrinks the
// batch and retries, so an occasionally slow provider self-heals while a
// dead one trips the breaker.
const DefaultRPCCallTimeout = time.Minute

// boundedClient decorates a Client with a per-call deadline. A caller
// context that already expires sooner keeps its earlier deadline
// (context.WithTimeout never extends one).
type boundedClient struct {
	inner   Client
	timeout time.Duration
}

// newBoundedClient wraps inner with the per-call deadline; timeout must be
// positive (New substitutes DefaultRPCCallTimeout for zero).
func newBoundedClient(inner Client, timeout time.Duration) Client {
	return boundedClient{inner: inner, timeout: timeout}
}

// BlockNumber returns the current chain head.
func (c boundedClient) BlockNumber(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	return c.inner.BlockNumber(ctx)
}

// HeaderByNumber returns the header of the given block (nil = latest).
func (c boundedClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	return c.inner.HeaderByNumber(ctx, number)
}

// FilterLogs executes an eth_getLogs query.
func (c boundedClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	return c.inner.FilterLogs(ctx, q)
}

// TransactionByHash returns the transaction and whether it is pending.
func (c boundedClient) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	return c.inner.TransactionByHash(ctx, hash)
}

// TransactionReceipt returns the receipt of a mined transaction.
func (c boundedClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	return c.inner.TransactionReceipt(ctx, txHash)
}
