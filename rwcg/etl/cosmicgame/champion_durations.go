package main

// V3 champion durations (championDurations[roundNum]).
//
// CosmicSignatureGameV3 saves the final endurance-champion and chrono-warrior durations
// into a public mapping at claimMainPrize() time, WITHOUT emitting an event. The ETL
// therefore fetches them via eth_call:
//   - when processing a V3 MainPrizeClaimed event (non-fatal: the claim is registered
//     even if the RPC call fails), and
//   - at ETL startup, backfilling every claimed round whose values are still unset
//     (which is also the recovery path for claim-time RPC failures: restart the ETL).
//
// Pre-V3 rounds legitimately return (0, 0) from the mapping and keep zeros in the DB.
// On a V2 contract the championDurations selector does not exist, so the first call
// reverts and the backfill is skipped entirely.

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
)

const CHAMPION_DURATIONS_CALL_TIMEOUT = 30 * time.Second

var cosmic_game_v3_caller *CosmicSignatureGameV3Caller

func init_champion_durations_caller() {
	var err error
	cosmic_game_v3_caller, err = NewCosmicSignatureGameV3Caller(cosmic_game_addr, eclient)
	if err != nil {
		// ABI parse failure; champion durations simply won't be collected.
		Info.Printf("Can't instantiate CosmicSignatureGameV3 caller (champion durations disabled): %v\n", err)
		Error.Printf("Can't instantiate CosmicSignatureGameV3 caller (champion durations disabled): %v\n", err)
		cosmic_game_v3_caller = nil
	}
}

// fetch_champion_durations reads championDurations[round_num] from the contract with a
// 30 second timeout. Returns an error on RPC failure/timeout, or if the contract is not
// V3 yet (the call reverts because the selector does not exist on V1/V2).
func fetch_champion_durations(round_num int64) (endurance int64, chrono int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), CHAMPION_DURATIONS_CALL_TIMEOUT)
	defer cancel()
	copts := &bind.CallOpts{Context: ctx}
	out, err := cosmic_game_v3_caller.ChampionDurations(copts, big.NewInt(round_num))
	if err != nil {
		return 0, 0, err
	}
	return out.EnduranceChampion.Int64(), out.ChronoWarrior.Int64(), nil
}

// store_champion_durations_for_round fetches and stores the durations for one claimed round.
// RPC failures are non-fatal by design: the claim record must survive even when the RPC
// server is down; a later ETL restart backfills the missing values.
func store_champion_durations_for_round(round_num int64) bool {
	if cosmic_game_v3_caller == nil {
		return false
	}
	endurance, chrono, err := fetch_champion_durations(round_num)
	if err != nil {
		Info.Printf("championDurations(%v) eth_call failed (claim registered without durations; restart ETL to backfill): %v\n", round_num, err)
		Error.Printf("championDurations(%v) eth_call failed: %v\n", round_num, err)
		return false
	}
	if endurance == 0 && chrono == 0 {
		// Pre-V3 round (mapping empty on-chain) or nothing to record; leave the zeros in place.
		return true
	}
	storagew.Update_round_champion_durations(round_num, endurance, chrono)
	Info.Printf("championDurations(%v): endurance=%v chrono=%v (stored in cg_round_stats)\n", round_num, endurance, chrono)
	return true
}

// backfill_champion_durations runs at ETL startup and populates champion durations for
// every claimed round that has none yet. It aborts on the first RPC error because a
// failure is systemic, not per-round: either the RPC server is unreachable or the
// contract is still V1/V2 (selector reverts for every round). Never fatal.
func backfill_champion_durations() {
	if cosmic_game_v3_caller == nil {
		return
	}
	rounds := storagew.Get_rounds_missing_champion_durations()
	if len(rounds) == 0 {
		Info.Printf("Champion durations backfill: nothing to do\n")
		return
	}
	Info.Printf("Champion durations backfill: checking %v round(s) with unset values\n", len(rounds))
	for _, round_num := range rounds {
		endurance, chrono, err := fetch_champion_durations(round_num)
		if err != nil {
			Info.Printf("Champion durations backfill stopped at round %v (contract not V3 yet, or RPC failure): %v\n", round_num, err)
			return
		}
		if endurance == 0 && chrono == 0 {
			continue // pre-V3 round; keep zeros
		}
		storagew.Update_round_champion_durations(round_num, endurance, chrono)
		Info.Printf("Champion durations backfill: round %v endurance=%v chrono=%v\n", round_num, endurance, chrono)
	}
	Info.Printf("Champion durations backfill: done\n")
}
