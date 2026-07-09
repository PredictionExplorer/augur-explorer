//go:build integration

package randomwalk

import (
	"context"
	"errors"
	"testing"
	"time"

	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
)

// TestErrorPathsConvertedFiles proves the context-first plumbing of the
// converted RandomWalk store: one representative method per file must abort
// with context.Canceled in the chain when the context is cancelled, and a
// closed pool must yield an error instead of panicking or exiting. The
// legacy layer called os.Exit(1) on all of these paths.
func TestErrorPathsConvertedFiles(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cancelled, cancel := context.WithCancel(ctx)
	cancel()

	cancelledCalls := map[string]func() error{
		"ProcessingStatus (randomwalk.go)": func() error {
			_, err := r.ProcessingStatus(cancelled)
			return err
		},
		"UpdateProcessingStatus (randomwalk.go)": func() error {
			return r.UpdateProcessingStatus(cancelled, &rwp.ProcStatus{})
		},
		"ContractAddrs (randomwalk.go)": func() error {
			_, err := r.ContractAddrs(cancelled)
			return err
		},
		"InsertNewOffer (randomwalk.go)": func() error {
			return r.InsertNewOffer(cancelled, &rwp.NewOffer{
				EvtId: -1, Contract: "0xEeee00000000000000000000000000000000Eeee",
				RWalkAddr: addrRandomWalk, Buyer: "0xEeee00000000000000000000000000000000Ffff",
				Seller: "0xEeee00000000000000000000000000000000Aaaa", Price: "0",
			})
		},
		"InsertMint (randomwalk.go)": func() error {
			return r.InsertMint(cancelled, &rwp.MintEvent{
				EvtId: -1, Contract: "0xEeee00000000000000000000000000000000Eeee",
				Owner: "0xEeee00000000000000000000000000000000Ffff",
				Seed:  "00", SeedNum: "0", Price: "0",
			})
		},
		"OfferExists (randomwalk.go)": func() error {
			_, err := r.OfferExists(cancelled, addrMarketplace, 1)
			return err
		},
		"UpdateTopProfitRank (randomwalk.go)": func() error {
			return r.UpdateTopProfitRank(cancelled, aidCarol, 1, 1)
		},
		"MessagingStatus (randomwalk.go)": func() error {
			_, err := r.MessagingStatus(cancelled)
			return err
		},
		"AllEventsForNotificationSinceEvtlog (randomwalk.go)": func() error {
			_, err := r.AllEventsForNotificationSinceEvtlog(cancelled, aidRandomWalk, 0)
			return err
		},
		"ActiveOffers (randomwalk_api.go)": func() error {
			_, err := r.ActiveOffers(cancelled, aidRandomWalk, aidMarketplace, 0)
			return err
		},
		"MintedTokensSequentially (randomwalk_api.go)": func() error {
			_, err := r.MintedTokensSequentially(cancelled, aidRandomWalk, 0, 10)
			return err
		},
		"TradingHistory (randomwalk_api.go)": func() error {
			_, err := r.TradingHistory(cancelled, aidMarketplace, 0, 10)
			return err
		},
		"RandomWalkStats (randomwalk_api.go)": func() error {
			_, err := r.RandomWalkStats(cancelled, aidRandomWalk)
			return err
		},
		"TokenFullHistory (randomwalk_api.go)": func() error {
			_, err := r.TokenFullHistory(cancelled, aidRandomWalk, 10, 0, 10)
			return err
		},
		"FloorPrice (randomwalk_api.go)": func() error {
			_, _, _, _, err := r.FloorPrice(cancelled, aidRandomWalk, aidMarketplace)
			return err
		},
		"UserInfo (randomwalk_api.go)": func() error {
			_, err := r.UserInfo(cancelled, aidCarol, aidRandomWalk)
			return err
		},
		"TokenInfo (randomwalk_api.go)": func() error {
			_, err := r.TokenInfo(cancelled, aidRandomWalk, 10)
			return err
		},
		"MintReport (randomwalk_api.go)": func() error {
			_, err := r.MintReport(cancelled)
			return err
		},
		"ExploreRandomTokenIDs (ranking.go)": func() error {
			_, err := r.ExploreRandomTokenIDs(cancelled, aidRandomWalk, 100, 2)
			return err
		},
		"CountRankingMatches (ranking.go)": func() error {
			_, err := r.CountRankingMatches(cancelled)
			return err
		},
		"RatingPair (ranking.go)": func() error {
			_, _, err := r.RatingPair(cancelled, 10, 11)
			return err
		},
		"InsertRankingVoteNonce (ranking.go)": func() error {
			return r.InsertRankingVoteNonce(cancelled, "error-path-nonce", time.Minute)
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
		"ContractAddrs": func() error {
			_, err := spareRepo.ContractAddrs(ctx)
			return err
		},
		"ActiveOffers": func() error {
			_, err := spareRepo.ActiveOffers(ctx, aidRandomWalk, aidMarketplace, 0)
			return err
		},
		"TokensByUser": func() error {
			_, err := spareRepo.TokensByUser(ctx, aidDave)
			return err
		},
		"MarketStats": func() error {
			_, err := spareRepo.MarketStats(ctx, aidMarketplace)
			return err
		},
		"TokenMinted": func() error {
			_, err := spareRepo.TokenMinted(ctx, 10)
			return err
		},
		"RatingOrder": func() error {
			_, err := spareRepo.RatingOrder(ctx, aidRandomWalk)
			return err
		},
		"HasRankingVoteForVoterPair": func() error {
			_, err := spareRepo.HasRankingVoteForVoterPair(ctx, aidAlice, 10, 11)
			return err
		},
		"UpdateMessagingStatus": func() error {
			return spareRepo.UpdateMessagingStatus(ctx, &rwp.MsgStatus{})
		},
		"InsertItemBought": func() error {
			return spareRepo.InsertItemBought(ctx, &rwp.ItemBought{
				EvtId: -1, Contract: "0xEeee00000000000000000000000000000000Eeee",
				SellerAddr: "0xEeee00000000000000000000000000000000Ffff",
				BuyerAddr:  "0xEeee00000000000000000000000000000000Aaaa",
			})
		},
		"LastMintTimestamp": func() error {
			_, err := spareRepo.LastMintTimestamp(ctx)
			return err
		},
	}
	for name, call := range closedPoolCalls {
		if err := call(); err == nil {
			t.Errorf("%s on a closed pool returned nil error", name)
		}
	}
}
