//go:build integration

// Integration coverage for the chain-sync paths the V1-mechanics suite in
// contractsync_integration_test.go does not reach: a V2-mechanics sync run
// (duration model, reward multiplier, the V2-only auction-duration-change
// divisor), and the failure modes — missing client, unusable chain and
// correction writes rejected by the database.
package cosmicgame

import (
	"context"
	"log/slog"
	"math/big"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// v2GameStub answers the V2 read model: the duration probe succeeds (so
// mechanics resolve to V2), rewards come from the multiplier and the
// V2-only change divisor is readable. V1-named shared parameters answer too.
func v2GameStub() *testchain.ContractStub {
	return testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI).
		Return("cstDutchAuctionDuration", big.NewInt(43200)). // mechanics probe -> V2
		Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(25)).
		Return("bidCstRewardAmountMultiplier", big.NewInt(10)).
		Return("timeoutDurationToClaimMainPrize", big.NewInt(86400)).
		Return("ethBidPriceIncreaseDivisor", big.NewInt(100)).
		Return("mainPrizeTimeIncrementIncreaseDivisor", big.NewInt(50)).
		Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3_600_000_000)).
		Return("initialDurationUntilMainPrizeDivisor", big.NewInt(200)).
		Return("ethDutchAuctionDurationDivisor", big.NewInt(40)).
		Return("ethDutchAuctionEndingBidPriceDivisor", big.NewInt(10)).
		Return("cstDutchAuctionBeginningBidPriceMinLimit", eth(200)).
		Return("marketingWalletCstContributionAmount", eth(300)).
		Return("delayDurationBeforeRoundActivation", big.NewInt(1800))
}

func TestSyncContractParamsV2Mechanics(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	testChain.EnsureBlock(500)

	testChain.RegisterCall(addr(fxGameAddr), v2GameStub().Handler())
	testChain.RegisterCall(addr(fxPrizesAddr), prizesWalletStub().Handler())

	ctx := context.Background()
	if err := SyncContractParams(ctx, cgRepo, dbStore, eclient, fxGameAddr, fxPrizesAddr, slog.New(slog.DiscardHandler)); err != nil {
		t.Fatalf("V2 sync run: %v", err)
	}

	// V2 model: the auction length is the duration itself, the reward the
	// multiplier, and the change divisor — skipped entirely under V1 — is
	// corrected too.
	requireLatestParam(t, "cg_adm_cst_auclen", "new_len", "43200")
	requireLatestParam(t, "cg_adm_erc20_reward", "new_reward", "10")
	requireLatestParam(t, "cg_adm_cst_auclen_chg_div", "new_len", "25")
	requireLatestParam(t, "cg_adm_timeout_withdraw", "new_timeout", "604800")

	// A clean V2 re-run corrects nothing further.
	watched := append([]string{"address", "cg_adm_cst_auclen_chg_div"}, syncedAdminTables...)
	before := tableCounts(t, watched)
	if err := SyncContractParams(ctx, cgRepo, dbStore, eclient, fxGameAddr, fxPrizesAddr, slog.New(slog.DiscardHandler)); err != nil {
		t.Fatalf("second V2 sync run: %v", err)
	}
	after := tableCounts(t, watched)
	for _, table := range watched {
		if before[table] != after[table] {
			t.Errorf("%s rows changed on a clean V2 re-run: %d -> %d", table, before[table], after[table])
		}
	}
}

func TestSyncContractParamsFailureModes(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	testChain.EnsureBlock(500)

	ctx := context.Background()
	logger := slog.New(slog.DiscardHandler)

	t.Run("nil client", func(t *testing.T) {
		err := SyncContractParams(ctx, cgRepo, dbStore, nil, fxGameAddr, fxPrizesAddr, logger)
		if err == nil || !strings.Contains(err.Error(), "eth client is nil") {
			t.Errorf("nil client error = %v", err)
		}
	})

	t.Run("cancelled context aborts the evtlog allocation", func(t *testing.T) {
		testChain.RegisterCall(addr(fxGameAddr), v1GameStub().Handler())
		cancelled, cancel := context.WithCancel(ctx)
		cancel()
		err := SyncContractParams(cancelled, cgRepo, dbStore, eclient, fxGameAddr, fxPrizesAddr, logger)
		if err == nil || !strings.Contains(err.Error(), "alloc chain sync evtlog") {
			t.Errorf("cancelled-context error = %v", err)
		}
	})

	t.Run("rejected correction writes propagate", func(t *testing.T) {
		// Drive the check-then-correct primitives over a read-only pool:
		// the comparison reads succeed, the correction inserts are rejected
		// by the database, and every sync* wrapper must surface that error
		// (a swallowed failure would leave the DB silently out of sync with
		// the chain).
		roCfg, err := pgxpool.ParseConfig(testDB.ConnString)
		if err != nil {
			t.Fatalf("parsing conn string: %v", err)
		}
		roCfg.ConnConfig.RuntimeParams["default_transaction_read_only"] = "on"
		roCfg.ConnConfig.RuntimeParams["timezone"] = "UTC"
		roCfg.ConnConfig.RuntimeParams["search_path"] = "public"
		roPool, err := pgxpool.NewWithConfig(ctx, roCfg)
		if err != nil {
			t.Fatalf("creating read-only pool: %v", err)
		}
		roStore := store.NewFromPool(roPool)
		defer roStore.Close()

		syncer := &paramSyncer{
			ctx:   ctx,
			repo:  cgstore.NewRepo(roStore),
			store: roStore,
			meta:  &cgstore.AdminCorrectionMeta{EvtId: -1, BlockNum: 500, TxId: 1, TimeStamp: 1, ContractAid: 1},
		}
		if _, err := syncer.syncDecimal("price_inc", "cg_adm_price_inc", "new_price_increase", "31337", ""); err == nil {
			t.Error("syncDecimal on a read-only database succeeded")
		}
		if _, err := syncer.syncInt64("delay", "cg_delay_duration", "new_value", 31337, ""); err == nil {
			t.Error("syncInt64 on a read-only database succeeded")
		}
		if _, err := syncer.syncCstReward("31337"); err == nil {
			t.Error("syncCstReward on a read-only database succeeded")
		}
		// A non-empty contract address must be registered before the insert;
		// on a read-only pool the registration itself fails.
		if _, err := syncer.syncDecimal("timeout", "cg_adm_timeout_withdraw", "new_timeout", "31337",
			"0x9999000000000000000000000000000000009999"); err == nil {
			t.Error("syncDecimal with an unregistered contract address on a read-only database succeeded")
		}
	})
}
