//go:build integration

package cosmicgame

import (
	"context"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

func TestV3ChampionDurationClaimFailureAndStartupRecovery(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	const (
		blockNum = int64(5300)
		roundNum = int64(11)
	)
	testChain.Reorg(blockNum)
	// An unconfigured V3 stub makes championDurations revert. The claim must
	// still commit so startup recovery can repair only the event-less fields.
	testChain.RegisterCall(addr(fxGameAddr),
		testchain.MustContractStub(cgc.CosmicSignatureGameV3ABI).Handler())
	ingestTx(t, blockNum, addr(fxGameAddr), 0, []*types.Log{
		buildLog(t, gameV3ABI, "MainPrizeClaimed", addr(fxGameAddr),
			[]any{bigInt(roundNum), addr(fxAlice), bigInt(200)},
			[]any{eth(4), eth(12), bigInt(3), bigInt(600)}),
	})

	var claims, endurance, chrono int64
	if err := testDB.SQL.QueryRow(
		`SELECT COUNT(*),MAX(rs.endurance_champion_duration),MAX(rs.chrono_warrior_duration)
		 FROM cg_prize_claim pc JOIN cg_round_stats rs ON rs.round_num=pc.round_num
		 WHERE pc.round_num=$1`, roundNum,
	).Scan(&claims, &endurance, &chrono); err != nil {
		t.Fatal(err)
	}
	if claims != 1 || endurance != 0 || chrono != 0 {
		t.Fatalf("claim after read failure = count:%d durations:%d/%d", claims, endurance, chrono)
	}

	registerCallHandlers()
	championDurationsMu.Lock()
	championDurations[roundNum] = [2]int64{701, 901}
	championDurationsMu.Unlock()
	if err := testHandlers.RecoverChampionDurations(
		context.Background(), blockNum+1, 1767231000,
	); err != nil {
		t.Fatalf("recover champion durations: %v", err)
	}
	endurance, chrono, err := cgRepo.RoundChampionDurations(context.Background(), roundNum)
	if err != nil || endurance != 701 || chrono != 901 {
		t.Fatalf("recovered durations = %d/%d, %v", endurance, chrono, err)
	}
	var auditRows int
	if err := testDB.SQL.QueryRow(
		"SELECT COUNT(*) FROM cg_live_state_updates WHERE round_num=$1", roundNum,
	).Scan(&auditRows); err != nil {
		t.Fatal(err)
	}
	if auditRows != 2 {
		t.Fatalf("live-state audit rows = %d, want 2", auditRows)
	}
	if err := testHandlers.RecoverChampionDurations(
		context.Background(), blockNum+2, 1767231100,
	); err != nil {
		t.Fatal(err)
	}
	var auditRowsAfter int
	if err := testDB.SQL.QueryRow(
		"SELECT COUNT(*) FROM cg_live_state_updates WHERE round_num=$1", roundNum,
	).Scan(&auditRowsAfter); err != nil {
		t.Fatal(err)
	}
	if auditRowsAfter != auditRows {
		t.Fatalf("idempotent recovery added rows: %d -> %d", auditRows, auditRowsAfter)
	}
}

func TestV3ClaimReorgRemovesLiveAuditAndReplayConverges(t *testing.T) {
	resetDB(t)
	const (
		blockNum = int64(5400)
		roundNum = int64(12)
	)
	testChain.Reorg(blockNum)
	championDurationsMu.Lock()
	championDurations[roundNum] = [2]int64{702, 902}
	championDurationsMu.Unlock()
	buildClaim := func() *types.Log {
		return buildLog(t, gameV3ABI, "MainPrizeClaimed", addr(fxGameAddr),
			[]any{bigInt(roundNum), addr(fxAlice), bigInt(300)},
			[]any{eth(5), eth(13), bigInt(3), bigInt(600)})
	}
	ingestTx(t, blockNum, addr(fxGameAddr), 0, []*types.Log{buildClaim()})
	full := snapshot(t)

	testChain.Reorg(blockNum)
	inserted, err := testIndexer.EnsureBlockExists(
		context.Background(), blockNum, testChain.BlockHash(blockNum).Hex(),
	)
	if err != nil || !inserted {
		t.Fatalf("detect V3 claim reorg = %v, %v", inserted, err)
	}
	var claims, updates int
	if err := testDB.SQL.QueryRow(
		"SELECT COUNT(*) FROM cg_prize_claim WHERE round_num=$1", roundNum,
	).Scan(&claims); err != nil {
		t.Fatal(err)
	}
	if err := testDB.SQL.QueryRow(
		"SELECT COUNT(*) FROM cg_live_state_updates WHERE round_num=$1", roundNum,
	).Scan(&updates); err != nil {
		t.Fatal(err)
	}
	if claims != 0 || updates != 0 {
		t.Fatalf("reorg residue = claims:%d updates:%d", claims, updates)
	}

	ingestTx(t, blockNum, addr(fxGameAddr), 0, []*types.Log{buildClaim()})
	requireNoDiff(t, full, snapshot(t), "V3 claim reorg replay")
}

func TestChampionDurationZeroAndOverflowReads(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	read, err := testHandlers.captureChampionDurations(
		context.Background(), 999, 5500, 1767232000,
	)
	if err != nil || !read {
		t.Fatalf("zero duration capture = %v, %v", read, err)
	}
	var rows int
	if err := testDB.SQL.QueryRow(
		"SELECT COUNT(*) FROM cg_live_state_updates WHERE round_num=999",
	).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows != 0 {
		t.Fatalf("zero duration capture wrote %d audit rows", rows)
	}

	overflow := new(big.Int).Lsh(big.NewInt(1), 80)
	testChain.RegisterCall(addr(fxGameAddr),
		testchain.MustContractStub(cgc.CosmicSignatureGameV3ABI).
			Return("championDurations", overflow, big.NewInt(1)).
			Handler())
	if _, _, err := testHandlers.readChampionDurations(context.Background(), 999); err == nil {
		t.Fatal("overflowing champion duration was accepted")
	}
}

func TestV3AdminStoreDeleteErrorsPropagate(t *testing.T) {
	tests := []struct {
		name  string
		table string
		store func(context.Context) error
	}{
		{
			name:  "duration divisor",
			table: "cg_adm_late_bid_dur_divisor",
			store: func(ctx context.Context) error {
				return testHandlers.storeRoundLateBidDurationDivisorChanged(
					ctx, &cgmodel.CGRoundLateBidDurationDivisorChanged{EvtId: 1},
				)
			},
		},
		{
			name:  "premium base",
			table: "cg_adm_late_bid_premium_base_mul",
			store: func(ctx context.Context) error {
				return testHandlers.storeRoundLateBidPremiumBaseMultiplierChanged(
					ctx, &cgmodel.CGRoundLateBidPricePremiumAmountBaseMultiplierChanged{EvtId: 1},
				)
			},
		},
		{
			name:  "premium exponent",
			table: "cg_adm_late_bid_premium_exponent",
			store: func(ctx context.Context) error {
				return testHandlers.storeRoundLateBidPremiumExponentChanged(
					ctx, &cgmodel.CGRoundLateBidPricePremiumAmountExponentChanged{EvtId: 1},
				)
			},
		},
		{
			name:  "reward percentage",
			table: "cg_adm_last_bidder_reward_pct",
			store: func(ctx context.Context) error {
				return testHandlers.storeLastBidderRewardPercentageChanged(
					ctx, &cgmodel.CGLastBidderBidCstRewardAmountPercentageChanged{EvtId: 1},
				)
			},
		},
		{
			name:  "main prize count",
			table: "cg_adm_main_prize_num_nfts",
			store: func(ctx context.Context) error {
				return testHandlers.storeMainPrizeNumNftsChanged(
					ctx, &cgmodel.CGMainPrizeNumCosmicSignatureNftsChanged{EvtId: 1},
				)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resetDB(t)
			restore := faultHarnessTable(t, test.table)
			err := test.store(context.Background())
			restore()
			if err == nil {
				t.Fatal("faulted V3 admin store succeeded")
			}
		})
	}
}

func TestV3RewardAndPrizeLookupErrorPaths(t *testing.T) {
	resetDB(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, _, err := testHandlers.cstBidRewards(cancelled, 1, 1, fxAlice); err == nil {
		t.Fatal("cancelled reward lookup succeeded")
	}
	if _, err := testHandlers.prizeRoundInTx(cancelled, 1); err == nil {
		t.Fatal("cancelled prize lookup succeeded")
	}

	ids := ingestTx(t, 5600, addr(fxGameAddr), 0, []*types.Log{
		buildLog(t, erc20ABI, "Transfer", addr(fxTokenAddr),
			[]any{addr("0x0000000000000000000000000000000000000000"), addr(fxAlice)},
			[]any{eth(10)}),
		buildLog(t, gameV3ABI, "BidPlaced", addr(fxGameAddr),
			[]any{bigInt(1), addr(fxAlice), bigInt(-1)},
			[]any{bigInt(1), bigInt(-1), "coverage", eth(10), bigInt(60), bigInt(1767233000)}),
	})
	var txID int64
	if err := testDB.SQL.QueryRow(
		"SELECT tx_id FROM evt_log WHERE id=$1", ids[1],
	).Scan(&txID); err != nil {
		t.Fatal(err)
	}
	if _, err := testDB.SQL.Exec(
		"UPDATE evt_log SET log_rlp=$1 WHERE id=$2", []byte{1}, ids[0],
	); err != nil {
		t.Fatal(err)
	}
	if _, _, _, err := testHandlers.cstBidRewards(
		context.Background(), ids[1], txID, fxAlice,
	); err == nil {
		t.Fatal("invalid reward RLP was accepted")
	}

	malformed := types.Log{
		Address: addr(fxTokenAddr),
		Topics: []ethcommon.Hash{
			topicHash(TopicTransferEvt),
			{},
			ethcommon.BytesToHash(addr(fxAlice).Bytes()),
		},
		Data: []byte{1},
	}
	raw, err := rlp.EncodeToBytes(&malformed)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := testDB.SQL.Exec(
		"UPDATE evt_log SET log_rlp=$1 WHERE id=$2", raw, ids[0],
	); err != nil {
		t.Fatal(err)
	}
	if _, _, _, err := testHandlers.cstBidRewards(
		context.Background(), ids[1], txID, fxAlice,
	); err == nil {
		t.Fatal("malformed reward Transfer was accepted")
	}
}
