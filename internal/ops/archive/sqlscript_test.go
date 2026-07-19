package archive

import (
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// scriptOp aliases the shared scripted pgx operation; the constructors below
// keep this package's historical test vocabulary.
type scriptOp = testutil.PgxOp

func queryOp(contains string, rows ...[]any) scriptOp {
	return scriptOp{Kind: "query", Contains: contains, Rows: rows}
}

func queryErrorOp(contains string, err error) scriptOp {
	return scriptOp{Kind: "query", Contains: contains, Err: err}
}

func queryIterErrorOp(contains string, err error) scriptOp {
	return scriptOp{Kind: "query", Contains: contains, RowsErr: err}
}

func execOp(contains string, affected int64) scriptOp {
	return scriptOp{Kind: "exec", Contains: contains, Affected: affected}
}

func execErrorOp(contains string, err error) scriptOp {
	return scriptOp{Kind: "exec", Contains: contains, Err: err}
}

// openScriptDB builds a scripted pgx querier serving ops in order and fails
// the test when operations remain unconsumed at cleanup.
func openScriptDB(t *testing.T, ops ...scriptOp) *testutil.ScriptedPgx {
	t.Helper()
	script := testutil.NewScriptedPgx(ops...)
	t.Cleanup(func() { script.AssertDone(t) })
	return script
}
