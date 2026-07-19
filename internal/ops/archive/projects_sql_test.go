package archive

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestLoadProjectContractsCosmicGame(t *testing.T) {
	db := openScriptDB(t,
		queryOp("JOIN cg_contracts",
			[]any{int64(2)},
			[]any{int64(3)},
		),
		queryOp("SELECT addr FROM address",
			[]any{"0x02"},
			[]any{"0x03"},
		),
	)
	contracts, err := LoadProjectContracts(context.Background(), db, ProjectCosmicGame)
	if err != nil {
		t.Fatalf("LoadProjectContracts() error = %v", err)
	}
	if len(contracts.AddressIDs) != 2 || len(contracts.Addresses) != 2 {
		t.Fatalf("contracts = %+v", contracts)
	}
}

func TestLoadProjectContractsFailures(t *testing.T) {
	sentinel := errors.New("database failure")
	tests := []struct {
		name    string
		project string
		ops     []scriptOp
		want    string
	}{
		{name: "unknown project", project: "bad", want: "unknown project"},
		{
			name:    "contract query",
			project: ProjectRandomWalk,
			ops:     []scriptOp{queryErrorOp("JOIN rw_contracts", sentinel)},
			want:    "contract aids",
		},
		{
			name:    "contract scan",
			project: ProjectRandomWalk,
			ops: []scriptOp{queryOp("JOIN rw_contracts",
				[]any{"not-an-integer"},
			)},
			want: "scan contract aid",
		},
		{
			name:    "contract iteration",
			project: ProjectRandomWalk,
			ops:     []scriptOp{queryIterErrorOp("JOIN rw_contracts", sentinel)},
			want:    "contract aids",
		},
		{
			name:    "no contracts",
			project: ProjectRandomWalk,
			ops:     []scriptOp{queryOp("JOIN rw_contracts")},
			want:    "no contract addresses found",
		},
		{
			name:    "address query",
			project: ProjectRandomWalk,
			ops: []scriptOp{
				queryOp("JOIN rw_contracts", []any{int64(8)}),
				queryErrorOp("SELECT addr FROM address", sentinel),
			},
			want: "resolve addrs",
		},
		{
			name:    "address scan",
			project: ProjectRandomWalk,
			ops: []scriptOp{
				queryOp("JOIN rw_contracts", []any{int64(8)}),
				queryOp("SELECT addr FROM address", []any{nil}),
			},
			want: "scan contract address",
		},
		{
			name:    "address iteration",
			project: ProjectRandomWalk,
			ops: []scriptOp{
				queryOp("JOIN rw_contracts", []any{int64(8)}),
				queryIterErrorOp("SELECT addr FROM address", sentinel),
			},
			want: "resolve addrs",
		},
		{
			name:    "no resolved addresses",
			project: ProjectRandomWalk,
			ops: []scriptOp{
				queryOp("JOIN rw_contracts", []any{int64(8)}),
				queryOp("SELECT addr FROM address"),
			},
			want: "no contract addresses resolved",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := openScriptDB(t, test.ops...)
			_, err := LoadProjectContracts(context.Background(), db, test.project)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}
