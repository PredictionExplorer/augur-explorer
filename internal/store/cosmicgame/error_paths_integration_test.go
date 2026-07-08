//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

// TestErrorPathsConvertedFiles extends TestErrorPaths to the files converted
// in the second Phase 1 batch: one representative method per file proves the
// context-first plumbing (cancellation aborts with context.Canceled in the
// chain, a closed pool errors instead of panicking or exiting). The legacy
// layer called os.Exit(1) on all of these paths.
func TestErrorPathsConvertedFiles(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cancelled, cancel := context.WithCancel(ctx)
	cancel()

	cancelledCalls := map[string]func() error{
		"BidFrequencyByPeriod (bidding_analytics.go)": func() error {
			_, err := r.BidFrequencyByPeriod(cancelled, fixtureStartTs, fixtureEndTs, 900)
			return err
		},
		"TopBidderActivePeriods (bidding_analytics.go)": func() error {
			_, _, err := r.TopBidderActivePeriods(cancelled, 3, fixtureStartTs, fixtureEndTs, 1, 1)
			return err
		},
		"BidTimeBounds (bidding_analytics.go)": func() error {
			_, _, err := r.BidTimeBounds(cancelled)
			return err
		},
		"UnclaimedPrizeEthDeposits (raffle-eth.go)": func() error {
			_, err := r.UnclaimedPrizeEthDeposits(cancelled, aidCarol, 0, 10)
			return err
		},
		"NFTDonations (nft-donations.go)": func() error {
			_, err := r.NFTDonations(cancelled, 0, 10)
			return err
		},
		"ERC20DonationsByRoundSummarized (erc20-donations.go)": func() error {
			_, err := r.ERC20DonationsByRoundSummarized(cancelled, 0)
			return err
		},
		"CosmicSignatureTokenInfo (tokens-erc721.go)": func() error {
			_, err := r.CosmicSignatureTokenInfo(cancelled, 1)
			return err
		},
		"LatestDecimalParam (contract_params.go)": func() error {
			_, _, err := r.LatestDecimalParam(cancelled, "cg_adm_charity_pcent", "percentage")
			return err
		},
		"BidsByRound (bidding.go)": func() error {
			_, _, err := r.BidsByRound(cancelled, 0, 0, 0, 10)
			return err
		},
		"EthDonations (eth-donations.go)": func() error {
			_, err := r.EthDonations(cancelled)
			return err
		},
		"ResolveAdminEventValues (admin_events_resolve.go)": func() error {
			events := []p.CGAdminEvent{{RecordType: 18, EvtLogId: 6000, IntegerValue: 10100}}
			return r.ResolveAdminEventValues(cancelled, events)
		},
		"StakeActionCstInfo (staking.go)": func() error {
			_, err := r.StakeActionCstInfo(cancelled, 1)
			return err
		},
		"GlobalStakingRwalkHistory (staking.go)": func() error {
			_, err := r.GlobalStakingRwalkHistory(cancelled, 0, 10)
			return err
		},
		"UserInfo (user-specific.go)": func() error {
			_, err := r.UserInfo(cancelled, aidAlice)
			return err
		},
		"UserNotifRedBoxRewards (user-specific.go)": func() error {
			_, err := r.UserNotifRedBoxRewards(cancelled, aidAlice)
			return err
		},
		"CosmicGameStatistics (statistics.go)": func() error {
			_, err := r.CosmicGameStatistics(cancelled)
			return err
		},
		"ClaimsByRound (statistics.go)": func() error {
			_, err := r.ClaimsByRound(cancelled)
			return err
		},
	}
	for name, call := range cancelledCalls {
		if err := call(); !errors.Is(err, context.Canceled) {
			t.Errorf("%s with cancelled ctx = %v, want context.Canceled in chain", name, err)
		}
	}

	// A closed pool yields an error, not a panic or exit.
	spare, err := spareStore(ctx)
	if err != nil {
		t.Fatalf("connecting spare store: %v", err)
	}
	spareRepo := NewRepo(spare)
	spare.Close()

	closedPoolCalls := map[string]func() error{
		"Bids": func() error {
			_, err := spareRepo.Bids(ctx, 0, 1)
			return err
		},
		"CosmicGameRoundStatistics": func() error {
			_, err := spareRepo.CosmicGameRoundStatistics(ctx, 0)
			return err
		},
		"StakingCstUserDepositRewards": func() error {
			_, err := spareRepo.StakingCstUserDepositRewards(ctx, aidAlice)
			return err
		},
		"PrizeClaimsByUser": func() error {
			_, err := spareRepo.PrizeClaimsByUser(ctx, aidAlice)
			return err
		},
		"ClaimDetailByRound": func() error {
			_, err := spareRepo.ClaimDetailByRound(ctx, 0)
			return err
		},
		"CharityDonations": func() error {
			_, err := spareRepo.CharityDonations(ctx, aidCosmicGame)
			return err
		},
		"CosmicSignatureTokenCount": func() error {
			_, err := spareRepo.CosmicSignatureTokenCount(ctx)
			return err
		},
		"InsertAdminCorrectionDecimal": func() error {
			meta := &AdminCorrectionMeta{EvtId: 5001, BlockNum: 1, TxId: 1001, TimeStamp: 1, ContractAid: aidCosmicGame}
			return spareRepo.InsertAdminCorrectionDecimal(ctx, "cg_adm_charity_pcent", "percentage", "1", meta, 0)
		},
	}
	for name, call := range closedPoolCalls {
		if err := call(); err == nil {
			t.Errorf("%s on closed pool succeeded, want error", name)
		}
	}
}
