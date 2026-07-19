package main

import (
	"bytes"
	"context"
	"math/big"
	"net/http"
	"sync/atomic"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/assets"
	"github.com/PredictionExplorer/augur-explorer/internal/ops/smoketest"
	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

type commandOutput struct {
	stdout string
	stderr string
	err    error
}

func executeCommand(command *cobra.Command, args ...string) commandOutput {
	var stdout, stderr bytes.Buffer
	command.SilenceErrors = true
	command.SilenceUsage = true
	command.SetOut(&stdout)
	command.SetErr(&stderr)
	command.SetArgs(args)
	err := command.Execute()
	return commandOutput{stdout: stdout.String(), stderr: stderr.String(), err: err}
}

// newCommandTestDB returns an opsDB fake whose Close call is observable.
func newCommandTestDB(t *testing.T) *testutil.ScriptedPgx {
	t.Helper()
	db := testutil.NewScriptedPgx()
	t.Cleanup(db.Close)
	return db
}

func assertCommandDBClosed(t *testing.T, db *testutil.ScriptedPgx) {
	t.Helper()
	if !db.Closed() {
		t.Fatal("database was not closed")
	}
}

type fakeOpsRPC struct {
	head    uint64
	headErr error
	closed  atomic.Bool
}

func (f *fakeOpsRPC) Close() { f.closed.Store(true) }

func (f *fakeOpsRPC) BlockNumber(context.Context) (uint64, error) {
	return f.head, f.headErr
}

func (*fakeOpsRPC) FilterLogs(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

func (*fakeOpsRPC) TransactionByHash(context.Context, common.Hash) (*types.Transaction, bool, error) {
	return nil, false, nil
}

func (*fakeOpsRPC) TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error) {
	return nil, nil
}

func (*fakeOpsRPC) BlockByNumber(context.Context, *big.Int) (*types.Block, error) {
	return nil, nil
}

type tokenSourceFunc func(context.Context, string) ([]assets.Token, error)

func (f tokenSourceFunc) TokenSeeds(ctx context.Context, schema string) ([]assets.Token, error) {
	return f(ctx, schema)
}

type tokenCountSourceFunc func(context.Context) (int64, error)

func (f tokenCountSourceFunc) MintedTokenCount(ctx context.Context) (int64, error) {
	return f(ctx)
}

type parameterSourceFunc func(context.Context) (smoketest.Params, error)

func (f parameterSourceFunc) Parameters(ctx context.Context) (smoketest.Params, error) {
	return f(ctx)
}

type httpClientFunc func(*http.Request) (*http.Response, error)

func (f httpClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

type commandRunnerFunc func(context.Context, string, ...string) ([]byte, error)

func (f commandRunnerFunc) CombinedOutput(ctx context.Context, name string, args ...string) ([]byte, error) {
	return f(ctx, name, args...)
}
