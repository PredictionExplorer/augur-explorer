package main

// The automated bidding bot behind the autobid subcommand. It is an
// event-based state machine, ported faithfully from the legacy autobid
// script: each loop iteration refreshes market data and processes flow
// states until the round ends or the operator interrupts the bot.

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// Internal bot constants (not user-configurable).
const (
	// rwalkMintEventTopic is the topic0 of the RandomWalk mint event.
	rwalkMintEventTopic = "ad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec"

	timeDelaySec         = 1   // seconds between event loop iterations
	timeDelayOnErrorMs   = 500 // milliseconds to back off after an error
	delayAfterTxSec      = 2   // seconds to wait after submitting a transaction
	delayNoActionSec     = 5   // seconds to wait when there is nothing to do
	maxRetries           = 3   // abandon a pending-tx wait after this many retries
	bidGasLimit          = 1000000
	claimGasLimit        = 5000000
	maxConsecutiveErrors = 5  // trigger reconnect after this many consecutive errors
	reconnectDelaySec    = 5  // delay before a reconnection attempt
	maxReconnectAttempts = 10 // max reconnection attempts before giving up
)

// botFlowState represents the current state in the bot's state machine.
type botFlowState int

const (
	flowUninitialized botFlowState = iota
	flowNotLastBidder
	flowIAmLastBidder
	flowNeedToBidWithCST
	flowWaitingForCSTBidTx
	flowNeedToBidWithETH
	flowWaitingForETHBidTx
	flowNeedToBidWithRWalk
	flowNeedToClaimPrize
	flowWaitingForClaimPrizeTx
	flowWaitingForRWalkBidTx
	flowWaitingForRWalkMint
	flowNeedToSendRWalkBidTx
	flowInitialBidding
)

// String returns the name of the flow state for logging.
func (f botFlowState) String() string {
	names := []string{
		"Uninitialized", "NotLastBidder", "IAmLastBidder", "NeedToBidWithCST",
		"WaitingForCSTBidTx", "NeedToBidWithETH", "WaitingForETHBidTx",
		"NeedToBidWithRWalk", "NeedToClaimPrize", "WaitingForClaimPrizeTx",
		"WaitingForRWalkBidTx", "WaitingForRWalkMint", "NeedToSendRWalkBidTx",
		"InitialBidding",
	}
	if int(f) < len(names) {
		return names[f]
	}
	return fmt.Sprintf("Unknown(%d)", f)
}

// sessionStats tracks statistics for the current bot session.
type sessionStats struct {
	startTime      time.Time
	startBalance   *big.Int
	ethBidCount    int
	cstBidCount    int
	rwalkBidCount  int
	totalEthSpent  *big.Int
	failedTxCount  int
	roundsWon      int
	prizesClaimed  int
	reconnectCount int
}

// biddingBot encapsulates all state and logic for automated bidding.
type biddingBot struct {
	config autobidConfig

	// Ethereum clients.
	rpcClient *ethrpc.Client
	ethClient *ethclient.Client
	callOpts  bind.CallOpts

	// Contracts.
	gameContract  *cgcontracts.CosmicSignatureGame
	rwalkContract *rwcontracts.RWalk
	prizesWallet  *cgcontracts.PrizesWallet

	// Account info.
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    *big.Int

	// State machine.
	flowState          botFlowState
	roundNumPlayed     int64
	lastBidderNotified bool

	// Pending transaction tracking.
	ethBidTxHash     common.Hash
	cstBidTxHash     common.Hash
	rwalkMintTxHash  common.Hash
	rwalkBidTxHash   common.Hash
	claimPrizeTxHash common.Hash
	retriesCounter   int

	// RWalk token tracking.
	nextRWalkTokenID int64
	prevRWalkTokenID int64

	// Current loop variables (refreshed each iteration).
	gasPrice       *big.Int
	cstPrice       *big.Int
	bidPrice       *big.Int
	timeUntilPrize *big.Int
	cstBalance     *big.Int
	ethBalance     *big.Int
	lastBidder     common.Address
	rwalkMintPrice *big.Int

	// Session statistics and shutdown.
	stats        sessionStats
	shutdownChan chan os.Signal

	// Connection health tracking.
	consecutiveErrors int
	reconnectCount    int
}

// newBiddingBot creates and initializes a bidding bot from the configuration.
func newBiddingBot(cfg autobidConfig) (*biddingBot, error) {
	bot := &biddingBot{
		config:           cfg,
		flowState:        flowUninitialized,
		roundNumPlayed:   -1,
		nextRWalkTokenID: -1,
		prevRWalkTokenID: -1,
		stats: sessionStats{
			startTime:     time.Now(),
			totalEthSpent: big.NewInt(0),
		},
		shutdownChan: make(chan os.Signal, 1),
	}

	// Setup signal handling for graceful shutdown.
	signal.Notify(bot.shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	var err error
	bot.rpcClient, err = ethrpc.DialContext(context.Background(), cfg.rpcURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to ETH RPC: %w", err)
	}
	bot.ethClient = ethclient.NewClient(bot.rpcClient)

	bot.chainID, err = bot.ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("getting chain id: %w", err)
	}

	bot.privateKey, err = crypto.HexToECDSA(cfg.privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %w", err)
	}
	publicKey, ok := bot.privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("couldn't derive public key")
	}
	bot.address = crypto.PubkeyToAddress(*publicKey)

	gameAddr := common.HexToAddress(cfg.gameContractAddr)
	bot.gameContract, err = cgcontracts.NewCosmicSignatureGame(gameAddr, bot.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate CosmicGame contract: %w", err)
	}

	rwalkAddr, err := bot.gameContract.RandomWalkNft(&bot.callOpts)
	if err != nil {
		return nil, fmt.Errorf("getting RWalk addr: %w", err)
	}
	bot.rwalkContract, err = rwcontracts.NewRWalk(rwalkAddr, bot.ethClient)
	if err != nil {
		return nil, fmt.Errorf("creating RWalk instance: %w", err)
	}

	prizesAddr, err := bot.gameContract.PrizesWallet(&bot.callOpts)
	if err != nil {
		return nil, fmt.Errorf("fetching PrizesWallet address: %w", err)
	}
	bot.prizesWallet, err = cgcontracts.NewPrizesWallet(prizesAddr, bot.ethClient)
	if err != nil {
		return nil, fmt.Errorf("creating PrizesWallet instance: %w", err)
	}

	roundNum, err := bot.gameContract.RoundNum(&bot.callOpts)
	if err != nil {
		return nil, fmt.Errorf("getting roundNum: %w", err)
	}
	bot.roundNumPlayed = roundNum.Int64()

	bot.gasPrice, err = bot.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("getting suggested gas price: %w", err)
	}

	bot.stats.startBalance, err = bot.ethClient.BalanceAt(context.Background(), bot.address, nil)
	if err != nil {
		return nil, fmt.Errorf("getting initial balance: %w", err)
	}

	return bot, nil
}

// reconnect attempts to re-establish the RPC connection.
func (bot *biddingBot) reconnect() error {
	bot.reconnectCount++
	bot.stats.reconnectCount++
	botLog("[RECONNECT] Attempting reconnection (%d/%d)...", bot.reconnectCount, maxReconnectAttempts)

	if bot.reconnectCount > maxReconnectAttempts {
		return fmt.Errorf("max reconnection attempts (%d) exceeded", maxReconnectAttempts)
	}

	if bot.rpcClient != nil {
		bot.rpcClient.Close()
	}

	time.Sleep(reconnectDelaySec * time.Second)

	var err error
	bot.rpcClient, err = ethrpc.DialContext(context.Background(), bot.config.rpcURL)
	if err != nil {
		return fmt.Errorf("reconnect failed - can't connect to ETH RPC: %w", err)
	}
	bot.ethClient = ethclient.NewClient(bot.rpcClient)

	chainID, err := bot.ethClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("reconnect failed - can't get chain ID: %w", err)
	}
	if chainID.Cmp(bot.chainID) != 0 {
		return fmt.Errorf("chain ID mismatch after reconnect: expected %v, got %v", bot.chainID, chainID)
	}

	// Re-instantiate contracts with the new client.
	gameAddr := common.HexToAddress(bot.config.gameContractAddr)
	bot.gameContract, err = cgcontracts.NewCosmicSignatureGame(gameAddr, bot.ethClient)
	if err != nil {
		return fmt.Errorf("reconnect failed - can't instantiate game contract: %w", err)
	}

	rwalkAddr, err := bot.gameContract.RandomWalkNft(&bot.callOpts)
	if err != nil {
		return fmt.Errorf("reconnect failed - can't get RWalk address: %w", err)
	}
	bot.rwalkContract, err = rwcontracts.NewRWalk(rwalkAddr, bot.ethClient)
	if err != nil {
		return fmt.Errorf("reconnect failed - can't instantiate RWalk contract: %w", err)
	}

	prizesAddr, err := bot.gameContract.PrizesWallet(&bot.callOpts)
	if err != nil {
		return fmt.Errorf("reconnect failed - can't get PrizesWallet address: %w", err)
	}
	bot.prizesWallet, err = cgcontracts.NewPrizesWallet(prizesAddr, bot.ethClient)
	if err != nil {
		return fmt.Errorf("reconnect failed - can't instantiate PrizesWallet contract: %w", err)
	}

	bot.consecutiveErrors = 0
	botLog("[RECONNECT] Successfully reconnected to RPC")
	return nil
}

// checkConnectionHealth reports whether the connection is healthy enough to
// keep going without a reconnect.
func (bot *biddingBot) checkConnectionHealth() bool {
	if bot.consecutiveErrors >= maxConsecutiveErrors {
		botLog("[WARNING] %d consecutive errors detected, connection may be unstable", bot.consecutiveErrors)
		return false
	}
	return true
}

// fmtEth formats wei as an ETH string.
func fmtEth(wei *big.Int) string {
	if wei == nil {
		return "0"
	}
	ether := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return ethValue.Text('f', 18)
}

// botLog prints a message with a timestamp.
func botLog(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), msg)
}

// isMyAddress checks whether the given address is the bot's address.
func (bot *biddingBot) isMyAddress(addr common.Address) bool {
	return bytes.Equal(addr.Bytes(), bot.address.Bytes())
}

// isZeroAddress checks whether the given address is the zero address.
func isZeroAddress(addr common.Address) bool {
	var zero common.Address
	return bytes.Equal(addr.Bytes(), zero.Bytes())
}

// createTransactOpts creates transaction options with signing. The gas price
// is refreshed for every transaction to ensure accurate pricing.
func (bot *biddingBot) createTransactOpts(value *big.Int, gasLimit uint64) (*bind.TransactOpts, error) {
	nonce, err := bot.ethClient.PendingNonceAt(context.Background(), bot.address)
	if err != nil {
		return nil, fmt.Errorf("getting nonce: %w", err)
	}

	gasPrice, err := bot.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("getting gas price: %w", err)
	}
	bot.gasPrice = gasPrice

	txopts := bind.NewKeyedTransactor(bot.privateKey)
	txopts.Nonce = big.NewInt(int64(nonce))
	txopts.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2))
	txopts.GasLimit = gasLimit
	if value != nil {
		txopts.Value = new(big.Int).Set(value)
	}

	txopts.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(bot.chainID)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), bot.privateKey)
		if err != nil {
			return nil, fmt.Errorf("signing: %w", err)
		}
		return tx.WithSignature(signer, signature)
	}

	return txopts, nil
}

// refreshMarketData fetches current market data from the contracts.
func (bot *biddingBot) refreshMarketData() error {
	var err error

	bot.cstPrice, err = bot.gameContract.GetNextCstBidPrice(&bot.callOpts)
	if err != nil {
		return fmt.Errorf("getting CST bid price: %w", err)
	}

	bot.bidPrice, err = bot.gameContract.GetNextEthBidPrice(&bot.callOpts)
	if err != nil {
		return fmt.Errorf("getting ETH bid price: %w", err)
	}

	bot.timeUntilPrize, err = bot.gameContract.GetDurationUntilMainPrize(&bot.callOpts)
	if err != nil {
		return fmt.Errorf("getting time until prize: %w", err)
	}

	tokenAddr, err := bot.gameContract.Token(&bot.callOpts)
	if err != nil {
		return fmt.Errorf("getting token address: %w", err)
	}
	tokenContract, err := cgcontracts.NewERC20(tokenAddr, bot.ethClient)
	if err != nil {
		return fmt.Errorf("instantiating token contract: %w", err)
	}
	bot.cstBalance, err = tokenContract.BalanceOf(&bot.callOpts, bot.address)
	if err != nil {
		return fmt.Errorf("getting CST balance: %w", err)
	}

	bot.lastBidder, err = bot.gameContract.LastBidderAddress(&bot.callOpts)
	if err != nil {
		return fmt.Errorf("getting last bidder: %w", err)
	}

	bot.ethBalance, err = bot.ethClient.BalanceAt(context.Background(), bot.address, nil)
	if err != nil {
		return fmt.Errorf("getting ETH balance: %w", err)
	}

	bot.rwalkMintPrice, err = bot.rwalkContract.GetMintPrice(&bot.callOpts)
	if err != nil {
		return fmt.Errorf("getting RWalk mint price: %w", err)
	}

	return nil
}

// checkTimeoutClaim checks if the bot can claim the prize because the winner
// missed the claim window.
func (bot *biddingBot) checkTimeoutClaim() {
	if isZeroAddress(bot.lastBidder) {
		return
	}

	timeoutDuration, err := bot.gameContract.TimeoutDurationToClaimMainPrize(&bot.callOpts)
	if err != nil {
		return
	}
	prizeTime, err := bot.gameContract.MainPrizeTime(&bot.callOpts)
	if err != nil {
		return
	}

	var result map[string]interface{}
	err = bot.rpcClient.CallContext(context.Background(), &result, "eth_getBlockByNumber", "pending", false)
	if err != nil {
		return
	}

	if timestampHex, ok := result["timestamp"].(string); ok {
		blockTimestamp, err := hexutil.DecodeBig(timestampHex)
		if err != nil {
			return
		}
		timeoutExpired := new(big.Int).Add(prizeTime, timeoutDuration)
		if timeoutExpired.Cmp(blockTimestamp) < 0 {
			botLog("Winner didn't claim prize during claim window, I am going to claim it")
			bot.flowState = flowNeedToClaimPrize
		}
	}
}

// checkRoundChange checks if the round has changed. It returns stop=true when
// the bot must terminate: with a nil error when the round ended normally, or
// with an error when the blockchain was reset (round number decreased).
func (bot *biddingBot) checkRoundChange() (stop bool, err error) {
	rnum, err := bot.gameContract.RoundNum(&bot.callOpts)
	if err != nil {
		botLog("Error getting roundNum: %v", err)
		return false, nil
	}

	if rnum.Int64() != bot.roundNumPlayed {
		botLog("Round changed (was %v, now %v)", bot.roundNumPlayed, rnum)

		// SAFETY: detect a blockchain reset (round went backwards).
		if rnum.Int64() < bot.roundNumPlayed {
			botLog("ERROR: Round number decreased (%v -> %v) - blockchain was reset!",
				bot.roundNumPlayed, rnum.Int64())
			botLog("Exiting to prevent unintended spending. Restart bot manually.")
			bot.printStats()
			return true, fmt.Errorf("blockchain was reset (round %v -> %v)", bot.roundNumPlayed, rnum.Int64())
		}

		winner, err := bot.prizesWallet.MainPrizeBeneficiaryAddresses(&bot.callOpts, big.NewInt(bot.roundNumPlayed))
		if err == nil {
			if bot.isMyAddress(winner) {
				botLog("I am the winner of round %v!", bot.roundNumPlayed)
				bot.stats.roundsWon++
			} else {
				botLog("I am not the winner of round %v", bot.roundNumPlayed)
			}
		}

		if bot.config.initialBidPrice != nil {
			botLog("Playing new round with initial bids")
			bot.roundNumPlayed = rnum.Int64()
			bot.flowState = flowInitialBidding
			return false, nil
		}
		botLog("Round ended, exiting...")
		bot.printStats()
		return true, nil
	}
	return false, nil
}

// findRWalkTokenID searches for an unused RWalk token owned by the bot.
func (bot *biddingBot) findRWalkTokenID() {
	if bot.nextRWalkTokenID > -1 {
		wasUsed, err := bot.gameContract.UsedRandomWalkNfts(&bot.callOpts, big.NewInt(bot.nextRWalkTokenID))
		if err == nil && wasUsed.Cmp(big.NewInt(1)) == 0 {
			botLog("Resetting nextRWalkTokenID (%v) - already used", bot.nextRWalkTokenID)
			bot.nextRWalkTokenID = -1
		} else if err == nil {
			return // Already have a valid token.
		}
	}

	targetID := bot.prevRWalkTokenID + 1
	lastTokenID, err := bot.rwalkContract.NextTokenId(&bot.callOpts)
	if err != nil {
		botLog("Error calling NextTokenId(): %v", err)
		return
	}

	for targetID < lastTokenID.Int64() {
		owner, err := bot.rwalkContract.OwnerOf(&bot.callOpts, big.NewInt(targetID))
		if err != nil {
			time.Sleep(timeDelayOnErrorMs * time.Millisecond)
			return
		}
		bot.prevRWalkTokenID = targetID

		if bot.isMyAddress(owner) {
			wasUsed, err := bot.gameContract.UsedRandomWalkNfts(&bot.callOpts, big.NewInt(targetID))
			if err == nil && wasUsed.Cmp(big.NewInt(0)) == 0 {
				bot.nextRWalkTokenID = targetID
				botLog("Found RWalk token %v for bidding", bot.nextRWalkTokenID)
				return
			}
		}
		targetID++
	}
}

// canBidWithCST checks if the CST bidding conditions are met.
func (bot *biddingBot) canBidWithCST() bool {
	return bot.config.maxCstBid.Cmp(bot.cstBalance) <= 0 &&
		bot.cstPrice.Cmp(bot.config.maxCstBid) <= 0
}

// handleUninitialized handles the initial state.
func (bot *biddingBot) handleUninitialized() botFlowState {
	if bot.isMyAddress(bot.lastBidder) {
		if !bot.lastBidderNotified {
			botLog("I am last bidder")
			bot.lastBidderNotified = true
		}
		return flowIAmLastBidder
	}
	botLog("I am not the last bidder (time until prize = %v)", bot.timeUntilPrize)
	bot.lastBidderNotified = false
	return flowNotLastBidder
}

// handleNotLastBidder handles logic when the bot is not the last bidder.
func (bot *biddingBot) handleNotLastBidder() (botFlowState, bool) {
	// Check if we should bid with CST anyway (but NOT if first bid of round - must be ETH).
	if bot.config.cstBidAnyway && bot.canBidWithCST() && !isZeroAddress(bot.lastBidder) {
		botLog("CST price (%v) below limit, bidding with CST", fmtEth(bot.cstPrice))
		return flowNeedToBidWithCST, true
	}

	// Re-verify last bidder status.
	if bot.isMyAddress(bot.lastBidder) {
		if !bot.lastBidderNotified {
			botLog("I am last bidder")
			bot.lastBidderNotified = true
		}
		return flowIAmLastBidder, false
	}

	// Check if it's time to bid.
	if bot.timeUntilPrize.Cmp(big.NewInt(bot.config.timeBeforePrize)) <= 0 {
		botLog("%v sec before prize, time to bid", bot.timeUntilPrize.Int64())
		return bot.decideBidType()
	}

	botLog("Not my time to bid yet (timeUntilPrize = %v)", bot.timeUntilPrize.Int64())
	return flowNotLastBidder, false
}

// decideBidType decides which bidding method to use.
func (bot *biddingBot) decideBidType() (botFlowState, bool) {
	// First bid of round must be ETH (contract requirement).
	if isZeroAddress(bot.lastBidder) {
		botLog("First bid of round - must use ETH")
		return bot.tryPlainEthBid()
	}

	// Try CST first.
	if bot.canBidWithCST() {
		botLog("CST price (%v) below limit, bidding with CST", fmtEth(bot.cstPrice))
		return flowNeedToBidWithCST, true
	}

	// Can't bid with CST, check ETH options.
	if bot.config.maxCstBid.Cmp(bot.cstBalance) > 0 {
		botLog("Not enough CST balance for bid")
	} else {
		botLog("CST price above limit")
	}

	// Check if RWalk bidding is allowed.
	if bot.config.rwalkMinPrice.Cmp(bot.bidPrice) < 0 {
		rwalkDiscountedPrice := new(big.Int).Quo(bot.bidPrice, big.NewInt(2))
		bidWithRwalkPrice := new(big.Int).Add(bot.rwalkMintPrice, rwalkDiscountedPrice)

		botLog("RWALK+ETH costs %v, pure ETH costs %v", fmtEth(bidWithRwalkPrice), fmtEth(bot.bidPrice))

		if bot.bidPrice.Cmp(bidWithRwalkPrice) <= 0 {
			// Plain ETH is cheaper.
			return bot.tryPlainEthBid()
		}
		// RWalk is cheaper.
		if bidWithRwalkPrice.Cmp(bot.config.maxEthBid) < 0 {
			botLog("Bidding with RWalk (cheaper)")
			return flowNeedToBidWithRWalk, true
		}
		botLog("Out of funds even with RWalk")
		return flowNotLastBidder, false
	}

	// Plain ETH only.
	return bot.tryPlainEthBid()
}

// tryPlainEthBid attempts to bid with plain ETH.
func (bot *biddingBot) tryPlainEthBid() (botFlowState, bool) {
	if bot.config.maxEthBid.Cmp(bot.bidPrice) < 0 {
		botLog("ETH bid price (%v) above limit", fmtEth(bot.bidPrice))
		return flowNotLastBidder, false
	}
	if bot.bidPrice.Cmp(bot.ethBalance) >= 0 {
		botLog("Insufficient ETH balance")
		return flowNotLastBidder, false
	}
	botLog("Bidding with plain ETH")
	return flowNeedToBidWithETH, true
}

// handleIAmLastBidder handles logic when the bot is the last bidder.
func (bot *biddingBot) handleIAmLastBidder() (botFlowState, bool) {
	// Use cached lastBidder (already fetched in refreshMarketData); don't make
	// duplicate contract calls that could cause inconsistency.
	if !bot.isMyAddress(bot.lastBidder) {
		botLog("No longer last bidder")
		bot.lastBidderNotified = false
		return flowNotLastBidder, false // Don't continue - need to refresh market data.
	}

	// Check if we should keep bidding with CST.
	if bot.config.cstBidAnyway && bot.canBidWithCST() {
		botLog("CST price low, bidding again")
		return flowNeedToBidWithCST, true
	}

	// Check if we can claim the prize.
	if bot.timeUntilPrize.Cmp(big.NewInt(0)) == 0 {
		return flowNeedToClaimPrize, true
	}

	return flowIAmLastBidder, false
}

// handleBidWithCST handles CST bidding.
func (bot *biddingBot) handleBidWithCST() (botFlowState, bool) {
	txopts, err := bot.createTransactOpts(nil, bidGasLimit)
	if err != nil {
		botLog("Error creating tx opts: %v", err)
		return flowUninitialized, false
	}

	tx, err := bot.gameContract.BidWithCst(txopts, bot.cstPrice, "")
	if err != nil {
		botLog("BidWithCST error: %v", err)
		return flowUninitialized, false
	}

	botLog("CST bid tx: %v", tx.Hash().Hex())
	bot.cstBidTxHash = tx.Hash()
	bot.retriesCounter = 0
	return flowWaitingForCSTBidTx, false
}

// handleWaitForCSTBidTx waits for the CST bid transaction receipt.
func (bot *biddingBot) handleWaitForCSTBidTx() (botFlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.cstBidTxHash)
	if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= maxRetries {
			botLog("Max retries reached for CST bid tx")
			return flowUninitialized, false
		}
		return flowWaitingForCSTBidTx, false
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		botLog("CST bid tx failed")
		bot.stats.failedTxCount++
		return flowUninitialized, false
	}

	bot.stats.cstBidCount++
	return bot.checkLastBidderAfterBid("CST")
}

// handleBidWithETH handles plain ETH bidding.
func (bot *biddingBot) handleBidWithETH() (botFlowState, bool) {
	txopts, err := bot.createTransactOpts(bot.bidPrice, bidGasLimit)
	if err != nil {
		botLog("Error creating tx opts: %v", err)
		return flowUninitialized, false
	}

	tx, err := bot.gameContract.BidWithEth(txopts, big.NewInt(-1), "")
	if err != nil {
		botLog("BidWithEth error: %v", err)
		return flowUninitialized, false
	}

	botLog("ETH bid tx (%v ETH): %v", fmtEth(bot.bidPrice), tx.Hash().Hex())
	bot.ethBidTxHash = tx.Hash()
	bot.retriesCounter = 0
	return flowWaitingForETHBidTx, false
}

// handleWaitForETHBidTx waits for the ETH bid transaction receipt.
func (bot *biddingBot) handleWaitForETHBidTx() (botFlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.ethBidTxHash)
	if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= maxRetries {
			botLog("Max retries reached for ETH bid tx")
			return flowUninitialized, false
		}
		return flowWaitingForETHBidTx, false
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		botLog("ETH bid tx failed")
		bot.stats.failedTxCount++
		return flowUninitialized, false
	}

	bot.stats.ethBidCount++
	bot.stats.totalEthSpent.Add(bot.stats.totalEthSpent, bot.bidPrice)
	return bot.checkLastBidderAfterBid("ETH")
}

// handleBidWithRWalk handles the RWalk bidding flow.
func (bot *biddingBot) handleBidWithRWalk() (botFlowState, bool) {
	if bot.nextRWalkTokenID > -1 {
		botLog("Using pre-minted RWalk token %v", bot.nextRWalkTokenID)
		return flowNeedToSendRWalkBidTx, true
	}

	botLog("Need to mint RWalk token first")
	txopts, err := bot.createTransactOpts(bot.rwalkMintPrice, bidGasLimit)
	if err != nil {
		botLog("Error creating tx opts: %v", err)
		return flowUninitialized, false
	}

	tx, err := bot.rwalkContract.Mint(txopts)
	if err != nil {
		botLog("RWalk mint error: %v", err)
		return flowUninitialized, false
	}

	botLog("RWalk mint tx: %v", tx.Hash().Hex())
	bot.rwalkMintTxHash = tx.Hash()
	bot.retriesCounter = 0
	return flowWaitingForRWalkMint, false
}

// handleWaitForRWalkMint waits for the RWalk mint transaction.
func (bot *biddingBot) handleWaitForRWalkMint() (botFlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.rwalkMintTxHash)
	if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= maxRetries {
			botLog("Max retries reached for RWalk mint tx")
			return flowUninitialized, false
		}
		return flowWaitingForRWalkMint, false
	}

	evtMintEvent, _ := hex.DecodeString(rwalkMintEventTopic)
	for _, elog := range receipt.Logs {
		if len(elog.Topics) > 0 && bytes.Equal(elog.Topics[0].Bytes(), evtMintEvent) {
			bot.nextRWalkTokenID = elog.Topics[1].Big().Int64()
			botLog("Minted RWalk token %v", bot.nextRWalkTokenID)
			return flowNeedToSendRWalkBidTx, true
		}
	}

	botLog("RWalk mint event not found")
	return flowUninitialized, false
}

// handleSendRWalkBidTx sends a bid using an RWalk token.
func (bot *biddingBot) handleSendRWalkBidTx() (botFlowState, bool) {
	txopts, err := bot.createTransactOpts(bot.bidPrice, bidGasLimit)
	if err != nil {
		botLog("Error creating tx opts: %v", err)
		return flowUninitialized, false
	}

	tx, err := bot.gameContract.BidWithEth(txopts, big.NewInt(bot.nextRWalkTokenID), "")
	if err != nil {
		botLog("BidWithEth (RWalk) error: %v", err)
		return flowUninitialized, false
	}

	botLog("RWalk bid tx: %v (token %v)", tx.Hash().Hex(), bot.nextRWalkTokenID)
	bot.rwalkBidTxHash = tx.Hash()
	bot.retriesCounter = 0
	return flowWaitingForRWalkBidTx, false
}

// handleWaitForRWalkBidTx waits for the RWalk bid transaction.
func (bot *biddingBot) handleWaitForRWalkBidTx() (botFlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.rwalkBidTxHash)
	if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= maxRetries {
			botLog("Max retries reached for RWalk bid tx")
			return flowUninitialized, false
		}
		return flowWaitingForRWalkBidTx, false
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		botLog("RWalk bid tx failed")
		bot.stats.failedTxCount++
		return flowUninitialized, false
	}

	bot.stats.rwalkBidCount++
	// RWalk bids cost half the ETH price.
	halfPrice := new(big.Int).Div(bot.bidPrice, big.NewInt(2))
	bot.stats.totalEthSpent.Add(bot.stats.totalEthSpent, halfPrice)
	bot.nextRWalkTokenID = -1
	return bot.checkLastBidderAfterBid("RWalk")
}

// handleClaimPrize handles claiming the prize.
func (bot *biddingBot) handleClaimPrize() (botFlowState, bool) {
	txopts, err := bot.createTransactOpts(nil, claimGasLimit)
	if err != nil {
		botLog("Error creating tx opts: %v", err)
		return flowUninitialized, false
	}

	tx, err := bot.gameContract.ClaimMainPrize(txopts)
	if err != nil {
		botLog("ClaimMainPrize error: %v", err)
		return flowUninitialized, false
	}

	botLog("ClaimPrize tx: %v", tx.Hash().Hex())
	bot.claimPrizeTxHash = tx.Hash()
	bot.retriesCounter = 0
	return flowWaitingForClaimPrizeTx, false
}

// handleWaitForClaimPrizeTx waits for the claim prize transaction.
func (bot *biddingBot) handleWaitForClaimPrizeTx() (botFlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.claimPrizeTxHash)
	if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= maxRetries {
			botLog("Max retries reached for claim prize tx")
			return flowUninitialized, false
		}
		return flowWaitingForClaimPrizeTx, false
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		botLog("Prize claimed successfully!")
		bot.stats.prizesClaimed++
		bot.stats.roundsWon++
	} else {
		botLog("Claim prize tx failed")
		bot.stats.failedTxCount++
	}
	return flowUninitialized, false
}

// handleInitialBidding handles the initial bidding phase at round start.
func (bot *biddingBot) handleInitialBidding() (botFlowState, bool) {
	botLog("Initial bidding: price=%v, limit=%v", fmtEth(bot.bidPrice), fmtEth(bot.config.initialBidPrice))

	const maxInitialBids = 20 // Safety limit to prevent runaway bidding.
	bidCount := 0
	totalSpent := big.NewInt(0)

	for failures := 0; failures <= 5; {
		if bidCount >= maxInitialBids {
			botLog("SAFETY: Reached max initial bids (%d), stopping", maxInitialBids)
			break
		}

		bidPrice, err := bot.gameContract.GetNextEthBidPrice(&bot.callOpts)
		if err != nil {
			botLog("Error getting bid price: %v", err)
			failures++
			time.Sleep(timeDelayOnErrorMs * time.Millisecond)
			continue
		}

		if bot.config.initialBidPrice.Cmp(bidPrice) < 0 {
			botLog("Reached bid price limit, stopping initial bidding")
			break
		}

		// Safety check: don't exceed the ETH balance.
		if bidPrice.Cmp(bot.ethBalance) >= 0 {
			botLog("Insufficient ETH balance for initial bid, stopping")
			break
		}

		txopts, err := bot.createTransactOpts(bidPrice, bidGasLimit)
		if err != nil {
			failures++
			time.Sleep(timeDelayOnErrorMs * time.Millisecond)
			continue
		}

		tx, err := bot.gameContract.BidWithEth(txopts, big.NewInt(-1), "")
		if err != nil {
			botLog("Bid error: %v", err)
			failures++
			time.Sleep(timeDelayOnErrorMs * time.Millisecond)
			continue
		}

		bidCount++
		totalSpent.Add(totalSpent, bidPrice)
		botLog("Initial bid #%d tx (%v ETH): %v", bidCount, fmtEth(bidPrice), tx.Hash().Hex())
		time.Sleep(delayAfterTxSec * time.Second)

		// Wait for the receipt.
		for i := 0; i < 5; i++ {
			time.Sleep(delayAfterTxSec * time.Second)
			receipt, err := bot.ethClient.TransactionReceipt(context.Background(), tx.Hash())
			if err == nil {
				if receipt.Status == types.ReceiptStatusSuccessful {
					botLog("Initial bid successful")
					bot.stats.ethBidCount++
					bot.stats.totalEthSpent.Add(bot.stats.totalEthSpent, bidPrice)
				} else {
					botLog("Initial bid tx failed, stopping")
					bot.stats.failedTxCount++
					failures = 6 // Force exit.
				}
				break
			}
		}
		failures = 0

		// Refresh the ETH balance after each bid.
		bot.ethBalance, _ = bot.ethClient.BalanceAt(context.Background(), bot.address, nil)
	}

	botLog("Initial bidding finished: %d bids, %v ETH spent", bidCount, fmtEth(totalSpent))
	return flowUninitialized, false
}

// checkLastBidderAfterBid checks if the bot is the last bidder after a bid.
func (bot *biddingBot) checkLastBidderAfterBid(bidType string) (botFlowState, bool) {
	lastBidder, err := bot.gameContract.LastBidderAddress(&bot.callOpts)
	if err != nil {
		botLog("Error checking last bidder: %v", err)
		return flowUninitialized, false
	}

	if bot.isMyAddress(lastBidder) {
		botLog("I am last bidder after %s bid", bidType)
		bot.lastBidderNotified = true
		return flowIAmLastBidder, false
	}

	botLog("Not last bidder after %s bid", bidType)
	bot.lastBidderNotified = false
	return flowNotLastBidder, false
}

// run starts the main event loop. It returns nil when the round ends or the
// bot is interrupted, and an error on fatal conditions (blockchain reset,
// reconnect exhaustion).
func (bot *biddingBot) run() error {
	botLog("Playing round %v", bot.roundNumPlayed)
	botLog("Press Ctrl+C to stop gracefully")
	bot.printConfig()

	for {
		// Check for a shutdown signal (non-blocking).
		select {
		case sig := <-bot.shutdownChan:
			botLog("Received signal: %v", sig)
			botLog("Shutting down gracefully...")
			bot.printStats()
			return nil
		default:
			// Continue normal operation.
		}

		// Check connection health and reconnect if needed.
		if !bot.checkConnectionHealth() {
			if err := bot.reconnect(); err != nil {
				botLog("[ERROR] Reconnection failed: %v", err)
				if bot.reconnectCount >= maxReconnectAttempts {
					botLog("[FATAL] Max reconnection attempts reached, exiting...")
					bot.printStats()
					return fmt.Errorf("max reconnection attempts reached: %w", err)
				}
				time.Sleep(reconnectDelaySec * time.Second)
				continue
			}
		}

		botLog("=== Event loop (flow=%s, rwalkNext=%v) ===",
			bot.flowState, bot.nextRWalkTokenID)

		// Search for RWalk tokens in the background.
		if bot.nextRWalkTokenID == -1 {
			go bot.findRWalkTokenID()
		}

		// Refresh market data.
		if err := bot.refreshMarketData(); err != nil {
			botLog("Error refreshing data: %v", err)
			bot.consecutiveErrors++
			time.Sleep(timeDelayOnErrorMs * time.Millisecond)
			continue
		}
		// Reset error counters on successful refresh.
		bot.consecutiveErrors = 0
		bot.reconnectCount = 0

		// Check for a timeout claim opportunity.
		bot.checkTimeoutClaim()

		// Check for a round change.
		if stop, err := bot.checkRoundChange(); stop {
			return err
		}

		// Process flow states.
		bot.processFlowStates()

		// Refresh the gas price.
		if gasPrice, err := bot.ethClient.SuggestGasPrice(context.Background()); err == nil {
			bot.gasPrice = gasPrice
		}

		time.Sleep(timeDelaySec * time.Second)
	}
}

// processFlowStates processes the state machine until no handler requests an
// immediate continuation.
func (bot *biddingBot) processFlowStates() {
	for {
		prevState := bot.flowState
		botLog("Processing state: %s (timeUntilPrize=%v)", bot.flowState, bot.timeUntilPrize.Int64())

		continueProcessing := false
		var sleepDuration time.Duration

		switch bot.flowState {
		case flowUninitialized:
			bot.flowState = bot.handleUninitialized()
			if bot.flowState == flowIAmLastBidder {
				sleepDuration = delayNoActionSec * time.Second
			} else if bot.flowState != flowUninitialized {
				continueProcessing = true
			}

		case flowNotLastBidder:
			bot.flowState, continueProcessing = bot.handleNotLastBidder()
			if !continueProcessing && bot.flowState == flowNotLastBidder {
				sleepDuration = delayNoActionSec * time.Second
			}

		case flowIAmLastBidder:
			bot.flowState, continueProcessing = bot.handleIAmLastBidder()
			if !continueProcessing && bot.flowState == flowIAmLastBidder {
				if bot.timeUntilPrize.Cmp(big.NewInt(delayNoActionSec)) >= 0 {
					sleepDuration = delayNoActionSec * time.Second
				}
			}

		case flowNeedToBidWithCST:
			bot.flowState, continueProcessing = bot.handleBidWithCST()
			sleepDuration = delayAfterTxSec * time.Second

		case flowWaitingForCSTBidTx:
			bot.flowState, continueProcessing = bot.handleWaitForCSTBidTx()
			if bot.flowState == flowWaitingForCSTBidTx {
				sleepDuration = timeDelayOnErrorMs * time.Millisecond
			}

		case flowNeedToBidWithETH:
			bot.flowState, continueProcessing = bot.handleBidWithETH()
			sleepDuration = delayAfterTxSec * time.Second

		case flowWaitingForETHBidTx:
			bot.flowState, continueProcessing = bot.handleWaitForETHBidTx()
			if bot.flowState == flowWaitingForETHBidTx {
				sleepDuration = timeDelayOnErrorMs * time.Millisecond
			}

		case flowNeedToBidWithRWalk:
			bot.flowState, continueProcessing = bot.handleBidWithRWalk()
			if !continueProcessing {
				sleepDuration = delayAfterTxSec * time.Second
			}

		case flowWaitingForRWalkMint:
			bot.flowState, continueProcessing = bot.handleWaitForRWalkMint()
			if bot.flowState == flowWaitingForRWalkMint {
				sleepDuration = timeDelayOnErrorMs * time.Millisecond
			}

		case flowNeedToSendRWalkBidTx:
			bot.flowState, continueProcessing = bot.handleSendRWalkBidTx()
			sleepDuration = delayAfterTxSec * time.Second

		case flowWaitingForRWalkBidTx:
			bot.flowState, continueProcessing = bot.handleWaitForRWalkBidTx()
			if bot.flowState == flowWaitingForRWalkBidTx {
				sleepDuration = timeDelayOnErrorMs * time.Millisecond
			}

		case flowNeedToClaimPrize:
			bot.flowState, continueProcessing = bot.handleClaimPrize()
			sleepDuration = delayAfterTxSec * time.Second

		case flowWaitingForClaimPrizeTx:
			bot.flowState, continueProcessing = bot.handleWaitForClaimPrizeTx()
			if bot.flowState == flowWaitingForClaimPrizeTx {
				sleepDuration = timeDelayOnErrorMs * time.Millisecond
			}

		case flowInitialBidding:
			bot.flowState, continueProcessing = bot.handleInitialBidding()

		default:
			botLog("Unknown flow state: %v", bot.flowState)
			sleepDuration = delayNoActionSec * time.Second
		}

		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}

		// Break out of the inner loop to refresh market data unless a handler
		// explicitly requested continuation.
		if !continueProcessing {
			break
		}
		// Also break if the state didn't change, to prevent infinite loops.
		if bot.flowState == prevState {
			break
		}
	}
}

// printConfig prints the current configuration.
func (bot *biddingBot) printConfig() {
	fmt.Println("Config params:")
	fmt.Printf("  MAX_ETH_BID: %v ETH\n", fmtEth(bot.config.maxEthBid))
	fmt.Printf("  MAX_CST_BID: %v CST\n", fmtEth(bot.config.maxCstBid))
	fmt.Printf("  RWALK_MIN_PRICE: %v ETH\n", fmtEth(bot.config.rwalkMinPrice))
	fmt.Printf("  TIME_BEFORE_PRIZE: %v secs\n", bot.config.timeBeforePrize)
	fmt.Printf("  CST_BID_ANYWAY: %v\n", bot.config.cstBidAnyway)
	if bot.config.initialBidPrice != nil {
		fmt.Printf("  AT_STARTUP_BID_UP_TO_PRICE_LEVEL: %v ETH\n", fmtEth(bot.config.initialBidPrice))
	}
}

// printStats prints the session statistics summary.
func (bot *biddingBot) printStats() {
	duration := time.Since(bot.stats.startTime)

	// Get the current balance to calculate the net change.
	currentBalance, err := bot.ethClient.BalanceAt(context.Background(), bot.address, nil)
	if err != nil {
		currentBalance = big.NewInt(0)
	}

	balanceChange := new(big.Int).Sub(currentBalance, bot.stats.startBalance)

	fmt.Println("\n" + "=" + "==========================================================")
	fmt.Println("                    SESSION SUMMARY")
	fmt.Println("==========================================================")
	fmt.Printf("  Duration:          %v\n", duration.Round(time.Second))
	fmt.Printf("  Starting Balance:  %v ETH\n", fmtEth(bot.stats.startBalance))
	fmt.Printf("  Current Balance:   %v ETH\n", fmtEth(currentBalance))
	if balanceChange.Sign() >= 0 {
		fmt.Printf("  Balance Change:    +%v ETH\n", fmtEth(balanceChange))
	} else {
		fmt.Printf("  Balance Change:    %v ETH\n", fmtEth(balanceChange))
	}
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("  ETH Bids:          %d\n", bot.stats.ethBidCount)
	fmt.Printf("  CST Bids:          %d\n", bot.stats.cstBidCount)
	fmt.Printf("  RWalk Bids:        %d\n", bot.stats.rwalkBidCount)
	fmt.Printf("  Total Bids:        %d\n", bot.stats.ethBidCount+bot.stats.cstBidCount+bot.stats.rwalkBidCount)
	fmt.Printf("  ETH Spent on Bids: %v ETH\n", fmtEth(bot.stats.totalEthSpent))
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("  Rounds Won:        %d\n", bot.stats.roundsWon)
	fmt.Printf("  Prizes Claimed:    %d\n", bot.stats.prizesClaimed)
	fmt.Printf("  Failed Txs:        %d\n", bot.stats.failedTxCount)
	fmt.Printf("  Reconnections:     %d\n", bot.stats.reconnectCount)
	fmt.Println("==========================================================")
}
