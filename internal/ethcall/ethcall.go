// Package ethcall bounds read-only Ethereum contract calls in time. The
// generated abigen bindings inherit whatever context their CallOpts carry —
// frequently context.Background() or a long-lived loop context — so a
// black-holed JSON-RPC endpoint could park a caller forever. Wrapping the
// binding's caller (or full backend) here gives every eth_call/eth_getCode a
// deadline without touching the generated code or its call sites (D22).
package ethcall

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// DefaultTimeout bounds one read-only contract call. Contract view reads are
// point queries a healthy node answers in milliseconds; 15s tolerates a slow
// provider while guaranteeing a wedged one surfaces as an error the caller's
// existing failure handling can act on.
const DefaultTimeout = 15 * time.Second

// BoundedCaller decorates a bind.ContractCaller with a per-call deadline. A
// caller context that already expires sooner keeps its earlier deadline
// (context.WithTimeout never extends one).
type BoundedCaller struct {
	inner   bind.ContractCaller
	timeout time.Duration
}

// NewBoundedCaller wraps inner with a per-call deadline; a non-positive
// timeout applies DefaultTimeout.
func NewBoundedCaller(inner bind.ContractCaller, timeout time.Duration) BoundedCaller {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	return BoundedCaller{inner: inner, timeout: timeout}
}

// CodeAt returns the code of the given account, bounded by the deadline.
func (c BoundedCaller) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	return c.inner.CodeAt(ctx, contract, blockNumber)
}

// CallContract executes an eth_call, bounded by the deadline.
func (c BoundedCaller) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	return c.inner.CallContract(ctx, call, blockNumber)
}

// BoundedReadBackend decorates a full bind.ContractBackend, bounding its
// read (ContractCaller) surface while every other method passes through
// unchanged. It exists for bindings that are constructed over a full backend
// but used read-only (the ETL contract-parameter sync, the notification
// bots' contract views); code that transacts must bound its submissions
// itself — internal/ethtx does.
type BoundedReadBackend struct {
	bind.ContractBackend

	timeout time.Duration
}

// NewBoundedReadBackend wraps backend with a per-read-call deadline; a
// non-positive timeout applies DefaultTimeout.
func NewBoundedReadBackend(backend bind.ContractBackend, timeout time.Duration) BoundedReadBackend {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	return BoundedReadBackend{ContractBackend: backend, timeout: timeout}
}

// CodeAt returns the code of the given account, bounded by the deadline.
func (b BoundedReadBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, b.timeout)
	defer cancel()
	return b.ContractBackend.CodeAt(ctx, contract, blockNumber)
}

// CallContract executes an eth_call, bounded by the deadline.
func (b BoundedReadBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, b.timeout)
	defer cancel()
	return b.ContractBackend.CallContract(ctx, call, blockNumber)
}
