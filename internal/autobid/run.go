package autobid

// The engine's event loop: refresh the market snapshot, let the pure
// decision core pick an action, execute it, and track the pending
// transaction until its receipt arrives.

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// Run starts the main event loop. It returns nil when the round ends or the
// context is cancelled, and an error on fatal conditions (blockchain reset,
// reconnect exhaustion).
func (e *Engine) Run(ctx context.Context) error {
	e.logf("Playing round %v", e.roundNumPlayed)
	e.printConfig()

	var (
		searchCancel context.CancelFunc
		searchDone   chan struct{}
	)
	cancelSearch := func() {
		if searchDone == nil {
			return
		}
		searchCancel()
	}
	waitSearch := func() {
		if searchDone == nil {
			return
		}
		<-searchDone
		searchCancel = nil
		searchDone = nil
	}
	shutdown := func() {
		cancelSearch()
		e.Close()
		waitSearch()
	}
	defer shutdown()

	for {
		if searchDone != nil {
			select {
			case <-searchDone:
				searchCancel()
				searchCancel = nil
				searchDone = nil
			default:
			}
		}
		if ctx.Err() != nil {
			shutdown()
			e.logf("Shutting down gracefully...")
			e.printStats(ctx)
			return nil
		}

		// Check connection health and reconnect if needed.
		if e.consecutiveErrors >= maxConsecutiveErrors {
			cancelSearch()
			if e.rpcClient != nil {
				// The search may be blocked in an RPC transport that does not
				// return on request-context cancellation until the client is
				// closed. This client is about to be replaced by reconnect.
				e.rpcClient.Close()
			}
			waitSearch()
			e.logf("[WARNING] %d consecutive errors detected, connection may be unstable", e.consecutiveErrors)
			if err := e.reconnect(ctx); err != nil {
				if ctx.Err() != nil {
					e.logf("Shutting down gracefully...")
					e.printStats(ctx)
					return nil
				}
				e.logf("[ERROR] Reconnection failed: %v", err)
				if e.reconnectCount >= maxReconnectAttempts {
					e.logf("[FATAL] Max reconnection attempts reached, exiting...")
					e.printStats(ctx)
					return fmt.Errorf("max reconnection attempts reached: %w", err)
				}
				if err := e.cfg.Sleep(ctx, reconnectDelay); err != nil {
					continue
				}
				continue
			}
		}

		e.logf("=== Event loop (pending=%s, rwalkNext=%v) ===", e.pendingKind, e.nextRWalkTokenID.Load())

		// Search for RWalk tokens in the background.
		if e.nextRWalkTokenID.Load() == -1 && searchDone == nil {
			searchCtx, cancel := context.WithCancel(ctx)
			done := make(chan struct{})
			searchCancel = cancel
			searchDone = done
			go func() {
				defer close(done)
				e.findRWalkTokenID(searchCtx)
			}()
		}

		// Refresh market data.
		if err := e.refreshMarket(ctx); err != nil {
			e.logf("Error refreshing data: %v", err)
			e.consecutiveErrors++
			_ = e.cfg.Sleep(ctx, errorDelay)
			continue
		}
		e.consecutiveErrors = 0
		e.reconnectCount = 0

		// A submitted transaction is resolved before anything else — in
		// particular before the round-change exit, so a claim that ends the
		// round is still confirmed and counted (the legacy loop could exit
		// with the claim receipt unread).
		if e.pendingKind != pendingNone {
			e.waitPendingTx(ctx)
		} else {
			// Check for a round change.
			if stop, err := e.checkRoundChange(ctx); stop {
				return err
			}
			e.step(ctx)
		}

		if err := e.cfg.Sleep(ctx, loopDelay); err != nil {
			continue
		}
	}
}

// step runs one decision iteration against the fresh market snapshot: the
// initial-bidding opener, the timeout-claim opportunity, or the regular
// decision core.
func (e *Engine) step(ctx context.Context) {
	if e.startInitialBidding {
		e.startInitialBidding = false
		e.runInitialBidding(ctx)
		return
	}

	// The timeout-claim opportunity beats the regular decision.
	if e.checkTimeoutClaim(ctx) {
		e.claimPrize(ctx)
		return
	}

	action, reason := Decide(e.market, e.cfg.Limits, e.address)
	e.logf("Decision: %s (%s; timeUntilPrize=%v)", action, reason, e.market.TimeUntilPrize.Int64())

	// Track the last-bidder role edge for the legacy log line.
	if e.isMyAddress(e.market.LastBidder) {
		if !e.lastBidderNotified {
			e.logf("I am last bidder")
			e.lastBidderNotified = true
		}
	} else {
		e.lastBidderNotified = false
	}

	switch action {
	case ActionBidCST:
		e.bidCST(ctx)
	case ActionBidETH:
		e.bidETH(ctx)
	case ActionBidRWalk:
		e.bidRWalk(ctx)
	case ActionClaimPrize:
		e.claimPrize(ctx)
	case ActionWait:
		// Nothing to do this iteration.
	}
}

// submitTx records a submitted transaction for receipt tracking.
func (e *Engine) submitTx(kind pendingTxKind, hash [32]byte, value *big.Int) {
	e.pendingKind = kind
	e.pendingTxHash = hash
	e.pendingValue = value
	e.retriesCounter = 0
}

func (e *Engine) bidCST(ctx context.Context) {
	txopts, err := e.createTransactOpts(ctx, nil, bidGasLimit)
	if err != nil {
		e.logf("Error creating tx opts: %v", err)
		return
	}
	tx, err := e.gameContract.BidWithCst(txopts, e.market.CstPrice, "")
	if err != nil {
		e.logf("BidWithCST error: %v", err)
		return
	}
	e.logf("CST bid tx: %v", tx.Hash().Hex())
	e.submitTx(pendingCSTBid, tx.Hash(), nil)
	_ = e.cfg.Sleep(ctx, afterTxDelay)
}

func (e *Engine) bidETH(ctx context.Context) {
	txopts, err := e.createTransactOpts(ctx, e.market.EthBidPrice, bidGasLimit)
	if err != nil {
		e.logf("Error creating tx opts: %v", err)
		return
	}
	tx, err := e.gameContract.BidWithEth(txopts, big.NewInt(-1), "")
	if err != nil {
		e.logf("BidWithEth error: %v", err)
		return
	}
	e.logf("ETH bid tx (%v ETH): %v", fmtEth(e.market.EthBidPrice), tx.Hash().Hex())
	e.submitTx(pendingETHBid, tx.Hash(), new(big.Int).Set(e.market.EthBidPrice))
	_ = e.cfg.Sleep(ctx, afterTxDelay)
}

// bidRWalk bids with an owned unused RandomWalk token, minting one first
// when none is available.
func (e *Engine) bidRWalk(ctx context.Context) {
	if e.nextRWalkTokenID.Load() > -1 {
		e.logf("Using pre-minted RWalk token %v", e.nextRWalkTokenID.Load())
		e.sendRWalkBid(ctx)
		return
	}

	e.logf("Need to mint RWalk token first")
	txopts, err := e.createTransactOpts(ctx, e.market.RWalkMintPrice, bidGasLimit)
	if err != nil {
		e.logf("Error creating tx opts: %v", err)
		return
	}
	tx, err := e.rwalkContract.Mint(txopts)
	if err != nil {
		e.logf("RWalk mint error: %v", err)
		return
	}
	e.logf("RWalk mint tx: %v", tx.Hash().Hex())
	e.submitTx(pendingRWalkMint, tx.Hash(), nil)
	_ = e.cfg.Sleep(ctx, afterTxDelay)
}

// sendRWalkBid submits the half-price ETH bid that consumes the RWalk token.
func (e *Engine) sendRWalkBid(ctx context.Context) {
	tokenID := e.nextRWalkTokenID.Load()
	txopts, err := e.createTransactOpts(ctx, e.market.EthBidPrice, bidGasLimit)
	if err != nil {
		e.logf("Error creating tx opts: %v", err)
		return
	}
	tx, err := e.gameContract.BidWithEth(txopts, big.NewInt(tokenID), "")
	if err != nil {
		e.logf("BidWithEth (RWalk) error: %v", err)
		return
	}
	e.logf("RWalk bid tx: %v (token %v)", tx.Hash().Hex(), tokenID)
	// RWalk bids cost half the ETH price.
	e.submitTx(pendingRWalkBid, tx.Hash(), new(big.Int).Div(e.market.EthBidPrice, big.NewInt(2)))
	_ = e.cfg.Sleep(ctx, afterTxDelay)
}

func (e *Engine) claimPrize(ctx context.Context) {
	txopts, err := e.createTransactOpts(ctx, nil, claimGasLimit)
	if err != nil {
		e.logf("Error creating tx opts: %v", err)
		return
	}
	tx, err := e.gameContract.ClaimMainPrize(txopts)
	if err != nil {
		e.logf("ClaimMainPrize error: %v", err)
		return
	}
	e.logf("ClaimPrize tx: %v", tx.Hash().Hex())
	e.submitTx(pendingClaim, tx.Hash(), nil)
	_ = e.cfg.Sleep(ctx, afterTxDelay)
}

// waitPendingTx polls the pending transaction's receipt once, applying the
// bounded retry policy: after maxReceiptRetries missing receipts the
// transaction is abandoned (the next loop iteration re-decides from fresh
// market data).
func (e *Engine) waitPendingTx(ctx context.Context) {
	kind := e.pendingKind
	receipt, err := e.ethClient.TransactionReceipt(ctx, e.pendingTxHash)
	if err != nil {
		e.retriesCounter++
		if e.retriesCounter >= maxReceiptRetries {
			e.logf("Max retries reached for %s tx", kind)
			e.pendingKind = pendingNone
		}
		_ = e.cfg.Sleep(ctx, errorDelay)
		return
	}

	e.pendingKind = pendingNone
	if receipt.Status != types.ReceiptStatusSuccessful {
		e.logf("%s tx failed", kind)
		e.stats.failedTxCount++
		return
	}

	switch kind {
	case pendingCSTBid:
		e.stats.cstBidCount++
		e.logAfterBid(ctx, "CST")
	case pendingETHBid:
		e.stats.ethBidCount++
		e.stats.totalEthSpent.Add(e.stats.totalEthSpent, e.pendingValue)
		e.logAfterBid(ctx, "ETH")
	case pendingRWalkBid:
		e.stats.rwalkBidCount++
		e.stats.totalEthSpent.Add(e.stats.totalEthSpent, e.pendingValue)
		e.nextRWalkTokenID.Store(-1)
		e.logAfterBid(ctx, "RWalk")
	case pendingRWalkMint:
		e.handleMintReceipt(ctx, receipt)
	case pendingClaim:
		e.logf("Prize claimed successfully!")
		e.stats.prizesClaimed++
		e.stats.roundsWon++
	case pendingNone:
		// Unreachable: kind is captured before the pendingNone reset.
	}
}

// handleMintReceipt extracts the minted token id from the mint event and
// immediately submits the RWalk bid with it.
func (e *Engine) handleMintReceipt(ctx context.Context, receipt *types.Receipt) {
	for _, elog := range receipt.Logs {
		if len(elog.Topics) > 1 && elog.Topics[0] == rwalkMintEventTopic {
			tokenID := elog.Topics[1].Big().Int64()
			e.nextRWalkTokenID.Store(tokenID)
			e.logf("Minted RWalk token %v", tokenID)
			e.sendRWalkBid(ctx)
			return
		}
	}
	e.logf("RWalk mint event not found")
}

// logAfterBid reports whether the bot holds the last-bidder position after a
// confirmed bid.
func (e *Engine) logAfterBid(ctx context.Context, bidType string) {
	lastBidder, err := e.gameContract.LastBidderAddress(callOpts(ctx))
	if err != nil {
		e.logf("Error checking last bidder: %v", err)
		return
	}
	if e.isMyAddress(lastBidder) {
		e.logf("I am last bidder after %s bid", bidType)
		e.lastBidderNotified = true
		return
	}
	e.logf("Not last bidder after %s bid", bidType)
	e.lastBidderNotified = false
}

// runInitialBidding bids repeatedly at round start until the price reaches
// the configured level, the balance runs out, or the safety cap is hit.
func (e *Engine) runInitialBidding(ctx context.Context) {
	e.logf("Initial bidding: price=%v, limit=%v", fmtEth(e.market.EthBidPrice), fmtEth(e.cfg.InitialBidPrice))

	bidCount := 0
	totalSpent := big.NewInt(0)
	ethBalance := e.market.EthBalance
	opts := callOpts(ctx)

	for failures := 0; failures <= 5; {
		if ctx.Err() != nil {
			break
		}
		if bidCount >= maxInitialBids {
			e.logf("SAFETY: Reached max initial bids (%d), stopping", maxInitialBids)
			break
		}

		bidPrice, err := e.gameContract.GetNextEthBidPrice(opts)
		if err != nil {
			e.logf("Error getting bid price: %v", err)
			failures++
			_ = e.cfg.Sleep(ctx, errorDelay)
			continue
		}
		if e.cfg.InitialBidPrice.Cmp(bidPrice) < 0 {
			e.logf("Reached bid price limit, stopping initial bidding")
			break
		}
		// Safety check: don't exceed the ETH balance.
		if bidPrice.Cmp(ethBalance) >= 0 {
			e.logf("Insufficient ETH balance for initial bid, stopping")
			break
		}

		txopts, err := e.createTransactOpts(ctx, bidPrice, bidGasLimit)
		if err != nil {
			failures++
			_ = e.cfg.Sleep(ctx, errorDelay)
			continue
		}
		tx, err := e.gameContract.BidWithEth(txopts, big.NewInt(-1), "")
		if err != nil {
			e.logf("Bid error: %v", err)
			failures++
			_ = e.cfg.Sleep(ctx, errorDelay)
			continue
		}

		bidCount++
		totalSpent.Add(totalSpent, bidPrice)
		e.logf("Initial bid #%d tx (%v ETH): %v", bidCount, fmtEth(bidPrice), tx.Hash().Hex())
		_ = e.cfg.Sleep(ctx, afterTxDelay)

		// Wait for the receipt.
		for range 5 {
			_ = e.cfg.Sleep(ctx, afterTxDelay)
			receipt, err := e.ethClient.TransactionReceipt(ctx, tx.Hash())
			if err == nil {
				if receipt.Status == types.ReceiptStatusSuccessful {
					e.logf("Initial bid successful")
					e.stats.ethBidCount++
					e.stats.totalEthSpent.Add(e.stats.totalEthSpent, bidPrice)
				} else {
					e.logf("Initial bid tx failed, stopping")
					e.stats.failedTxCount++
					failures = 6 // Force exit.
				}
				break
			}
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return
			}
		}
		if failures <= 5 {
			failures = 0
		}

		// Refresh the ETH balance after each bid.
		if bal, err := e.ethClient.BalanceAt(ctx, e.address, nil); err == nil {
			ethBalance = bal
		}
	}

	e.logf("Initial bidding finished: %d bids, %v ETH spent", bidCount, fmtEth(totalSpent))
}
