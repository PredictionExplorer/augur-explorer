//go:build integration

package cosmicgame

import (
	"context"
	"math"
	"testing"
)

func TestLiveStateUpdateIdempotence(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	const round = int64(9999)
	t.Cleanup(func() {
		_, _ = r.pool().Exec(context.Background(),
			"DELETE FROM cg_live_state_updates WHERE round_num=$1", round)
	})

	inserted, err := r.InsertLiveStateUpdateIfChanged(
		ctx, "endurance_champion_duration", aidCosmicGame, round, 500, 1000, "700",
	)
	if err != nil || !inserted {
		t.Fatalf("first live update = %v, %v", inserted, err)
	}
	inserted, err = r.InsertLiveStateUpdateIfChanged(
		ctx, "endurance_champion_duration", aidCosmicGame, round, 501, 1001, "0700",
	)
	if err != nil || inserted {
		t.Fatalf("equivalent live update = %v, %v; want skipped", inserted, err)
	}
	inserted, err = r.InsertLiveStateUpdateIfChanged(
		ctx, "endurance_champion_duration", aidCosmicGame, round, 502, 1002, "701",
	)
	if err != nil || !inserted {
		t.Fatalf("changed live update = %v, %v", inserted, err)
	}
	latest, err := r.LatestLiveStateUpdate(ctx, "endurance_champion_duration", round)
	if err != nil || latest.NewValue != "701" || latest.BlockNum != 502 {
		t.Fatalf("latest live update = %+v, %v", latest, err)
	}
}

func TestRoundChampionDurationsRoundTrip(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	const round = int64(9998)
	t.Cleanup(func() {
		_, _ = r.pool().Exec(context.Background(),
			"DELETE FROM cg_round_stats WHERE round_num=$1", round)
	})
	if err := r.UpdateRoundChampionDurations(ctx, round, 700, 900); err != nil {
		t.Fatal(err)
	}
	endurance, chrono, err := r.RoundChampionDurations(ctx, round)
	if err != nil || endurance != 700 || chrono != 900 {
		t.Fatalf("champion durations = %d,%d,%v", endurance, chrono, err)
	}
	if err := r.UpdateRoundChampionDurations(ctx, round, 701, 901); err != nil {
		t.Fatal(err)
	}
	endurance, chrono, err = r.RoundChampionDurations(ctx, round)
	if err != nil || endurance != 701 || chrono != 901 {
		t.Fatalf("updated champion durations = %d,%d,%v", endurance, chrono, err)
	}
}

func TestBidRewardsAreQueryable(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	var bidID int64
	if err := r.q(ctx).QueryRow(ctx, "SELECT id FROM cg_bid ORDER BY id LIMIT 1").Scan(&bidID); err != nil {
		t.Fatal(err)
	}
	if _, err := r.q(ctx).Exec(ctx, `INSERT INTO cg_bid_reward(
			evtlog_id,bid_id,round_num,recipient_aid,reward_type,amount
		)
		SELECT evtlog_id,id,round_num,bidder_aid,0,cst_reward
		FROM cg_bid WHERE id=$1`, bidID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, _ = r.pool().Exec(context.Background(),
			"DELETE FROM cg_bid_reward WHERE bid_id=$1", bidID)
	})
	rewards, err := r.BidRewardsByBidID(ctx, bidID)
	if err != nil {
		t.Fatal(err)
	}
	if len(rewards) != 1 || rewards[0].RewardType != 0 || rewards[0].RecipientAddr == "" {
		t.Fatalf("bid rewards = %+v", rewards)
	}
}

func TestRoundInfoRejectsOverflowingV3TokenRange(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	if _, err := r.q(ctx).Exec(ctx, `UPDATE cg_prize_claim
		SET token_id=$1,num_cs_nfts=2
		WHERE round_num=2`,
		int64(math.MaxInt64)); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		// Fixture round 2 uses first token 8 and one main-prize NFT.
		_, _ = r.pool().Exec(context.Background(), `UPDATE cg_prize_claim
			SET token_id=8,num_cs_nfts=1 WHERE round_num=2`)
	})
	if _, err := r.RoundInfo(ctx, 2); err == nil {
		t.Fatal("overflowing V3 token range was accepted")
	}
}
