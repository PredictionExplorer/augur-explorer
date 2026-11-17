//go:build integration

package chaoszero

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestStatistics(t *testing.T) {
	r := repo(t)
	golden(t, "statistics", func() any {
		stats, err := r.Statistics(context.Background())
		if err != nil {
			t.Fatalf("Statistics: %v", err)
		}
		return stats
	})
}

func TestRoundsList(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "rounds_list", func() any {
		rounds, err := r.RoundsList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("RoundsList: %v", err)
		}
		return rounds
	})
	golden(t, "rounds_list_paged", func() any {
		rounds, err := r.RoundsList(ctx, 1, 1)
		if err != nil {
			t.Fatalf("RoundsList(paged): %v", err)
		}
		return rounds
	})
}

func TestRoundDetail(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// Round 0 is settled: every sub-list is populated.
	golden(t, "round_detail_0", func() any {
		detail, err := r.RoundDetail(ctx, 0)
		if err != nil {
			t.Fatalf("RoundDetail(0): %v", err)
		}
		return detail
	})
	// Round 1 is open: no settlement, escrow or claims.
	golden(t, "round_detail_1", func() any {
		detail, err := r.RoundDetail(ctx, 1)
		if err != nil {
			t.Fatalf("RoundDetail(1): %v", err)
		}
		return detail
	})
	if _, err := r.RoundDetail(ctx, 999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("RoundDetail(999) = %v, want store.ErrNotFound", err)
	}
}

func TestTradesList(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "trades_list", func() any {
		trades, err := r.TradesList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("TradesList: %v", err)
		}
		return trades
	})
	golden(t, "trades_by_round_0", func() any {
		trades, err := r.TradesByRound(ctx, 0, 0, 100)
		if err != nil {
			t.Fatalf("TradesByRound: %v", err)
		}
		return trades
	})
	golden(t, "trades_by_user_alice", func() any {
		trades, err := r.TradesByUser(ctx, aidAlice, 0, 100)
		if err != nil {
			t.Fatalf("TradesByUser: %v", err)
		}
		return trades
	})
	empty, err := r.TradesByUser(ctx, 9999, 0, 100)
	if err != nil {
		t.Fatalf("TradesByUser(unknown): %v", err)
	}
	if len(empty) != 0 {
		t.Errorf("TradesByUser(unknown) = %d rows, want 0", len(empty))
	}
}

func TestTradersList(t *testing.T) {
	r := repo(t)
	// Realized-profit order: alice (+5), dave (-4), bob (-5), carol (-8).
	golden(t, "traders_list", func() any {
		traders, err := r.TradersList(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("TradersList: %v", err)
		}
		return traders
	})
}

func TestFundingList(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "funding_list", func() any {
		fundings, err := r.FundingList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("FundingList: %v", err)
		}
		return fundings
	})
	golden(t, "funding_by_round_0", func() any {
		fundings, err := r.FundingByRound(ctx, 0)
		if err != nil {
			t.Fatalf("FundingByRound: %v", err)
		}
		return fundings
	})
}

func TestSplitsAndMerges(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "splits_list", func() any {
		splits, err := r.SplitsList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("SplitsList: %v", err)
		}
		return splits
	})
	golden(t, "merges_list", func() any {
		merges, err := r.MergesList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("MergesList: %v", err)
		}
		return merges
	})
}

func TestSettlementsList(t *testing.T) {
	r := repo(t)
	golden(t, "settlements_list", func() any {
		settlements, err := r.SettlementsList(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("SettlementsList: %v", err)
		}
		return settlements
	})
}

func TestClaims(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "claims_list", func() any {
		claims, err := r.ClaimsList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("ClaimsList: %v", err)
		}
		return claims
	})
	golden(t, "claims_by_user_alice", func() any {
		claims, err := r.ClaimsByUser(ctx, aidAlice, 0, 100)
		if err != nil {
			t.Fatalf("ClaimsByUser: %v", err)
		}
		return claims
	})
}

func TestAnswersList(t *testing.T) {
	r := repo(t)
	golden(t, "answers_list", func() any {
		answers, err := r.AnswersList(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("AnswersList: %v", err)
		}
		return answers
	})
}

func TestEscrowsAndPayouts(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "escrows_list", func() any {
		escrows, err := r.EscrowsList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("EscrowsList: %v", err)
		}
		return escrows
	})
	golden(t, "payouts_list", func() any {
		payouts, err := r.PayoutsList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("PayoutsList: %v", err)
		}
		return payouts
	})
}

func TestTransfers(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "transfers_list", func() any {
		transfers, err := r.TransfersList(ctx, 0, 100)
		if err != nil {
			t.Fatalf("TransfersList: %v", err)
		}
		return transfers
	})
	// Dave both received a single transfer and two batch rows.
	golden(t, "transfers_by_user_dave", func() any {
		transfers, err := r.TransfersByUser(ctx, aidDave, 0, 100)
		if err != nil {
			t.Fatalf("TransfersByUser: %v", err)
		}
		return transfers
	})
}

func TestApprovalsList(t *testing.T) {
	r := repo(t)
	golden(t, "approvals_list", func() any {
		approvals, err := r.ApprovalsList(context.Background(), 0, 100)
		if err != nil {
			t.Fatalf("ApprovalsList: %v", err)
		}
		return approvals
	})
}

func TestContractAddrs(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	market, source, vault, err := r.RawContractAddrs(ctx)
	if err != nil {
		t.Fatalf("RawContractAddrs: %v", err)
	}
	if market != addrMarket || source != addrSource || vault != addrVault {
		t.Errorf("RawContractAddrs = (%s, %s, %s), want fixture addresses", market, source, vault)
	}

	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		t.Fatalf("ContractAddrs: %v", err)
	}
	if addrs.Market != addrMarket || addrs.Source != addrSource || addrs.Vault != addrVault {
		t.Errorf("ContractAddrs = %+v, want fixture addresses", addrs)
	}
	if addrs.MarketAid == 0 || addrs.SourceAid == 0 || addrs.VaultAid == 0 {
		t.Errorf("ContractAddrs ids = %+v, want non-zero address ids", addrs)
	}
}

// TestProcessingStatusRoundtrip exercises the watermark read/write pair and
// the lazy re-creation of a lost singleton row, on a spare store so the
// shared fixtures stay untouched.
func TestProcessingStatusRoundtrip(t *testing.T) {
	repo(t) // skip guard
	ctx := context.Background()
	st, err := spareStore(ctx)
	if err != nil {
		t.Fatalf("opening spare store: %v", err)
	}
	defer st.Close()
	r := NewRepo(st)

	// No restore needed: the spare store is a dedicated throwaway database.
	original, err := r.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus: %v", err)
	}

	status := original
	status.LastBlockNum = 4242
	status.LastIdProcessed = 17
	if err := r.UpdateProcessingStatus(ctx, &status); err != nil {
		t.Fatalf("UpdateProcessingStatus: %v", err)
	}
	got, err := r.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus after update: %v", err)
	}
	if got.LastBlockNum != 4242 || got.LastIdProcessed != 17 {
		t.Errorf("roundtripped status = %+v, want {17 4242}", got)
	}

	// A lost singleton row is recreated with the migration defaults (the
	// production deployment block).
	if _, err := st.Pool().Exec(ctx, "DELETE FROM cz_proc_status"); err != nil {
		t.Fatalf("deleting watermark row: %v", err)
	}
	recreated, err := r.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus after delete: %v", err)
	}
	if recreated.LastBlockNum != 517037316 || recreated.LastIdProcessed != 0 {
		t.Errorf("recreated status = %+v, want the migration defaults {0 517037316}", recreated)
	}
}

// TestReadsFailOnCancelledContext loops the read surface over an already
// cancelled context: every method must return an error, not a zero value.
func TestReadsFailOnCancelledContext(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	reads := map[string]func() error{
		"Statistics":       func() error { _, err := r.Statistics(ctx); return err },
		"RoundsList":       func() error { _, err := r.RoundsList(ctx, 0, 1); return err },
		"RoundDetail":      func() error { _, err := r.RoundDetail(ctx, 0); return err },
		"TradesList":       func() error { _, err := r.TradesList(ctx, 0, 1); return err },
		"TradesByRound":    func() error { _, err := r.TradesByRound(ctx, 0, 0, 1); return err },
		"TradesByUser":     func() error { _, err := r.TradesByUser(ctx, aidAlice, 0, 1); return err },
		"TradersList":      func() error { _, err := r.TradersList(ctx, 0, 1); return err },
		"FundingList":      func() error { _, err := r.FundingList(ctx, 0, 1); return err },
		"FundingByRound":   func() error { _, err := r.FundingByRound(ctx, 0); return err },
		"SplitsList":       func() error { _, err := r.SplitsList(ctx, 0, 1); return err },
		"MergesList":       func() error { _, err := r.MergesList(ctx, 0, 1); return err },
		"SettlementsList":  func() error { _, err := r.SettlementsList(ctx, 0, 1); return err },
		"ClaimsList":       func() error { _, err := r.ClaimsList(ctx, 0, 1); return err },
		"ClaimsByUser":     func() error { _, err := r.ClaimsByUser(ctx, aidAlice, 0, 1); return err },
		"AnswersList":      func() error { _, err := r.AnswersList(ctx, 0, 1); return err },
		"EscrowsList":      func() error { _, err := r.EscrowsList(ctx, 0, 1); return err },
		"PayoutsList":      func() error { _, err := r.PayoutsList(ctx, 0, 1); return err },
		"TransfersList":    func() error { _, err := r.TransfersList(ctx, 0, 1); return err },
		"TransfersByUser":  func() error { _, err := r.TransfersByUser(ctx, aidDave, 0, 1); return err },
		"ApprovalsList":    func() error { _, err := r.ApprovalsList(ctx, 0, 1); return err },
		"RawContractAddrs": func() error { _, _, _, err := r.RawContractAddrs(ctx); return err },
		"ContractAddrs":    func() error { _, err := r.ContractAddrs(ctx); return err },
		"ProcessingStatus": func() error { _, err := r.ProcessingStatus(ctx); return err },
	}
	for name, read := range reads {
		if err := read(); err == nil {
			t.Errorf("%s returned nil error on a cancelled context", name)
		}
	}
}
