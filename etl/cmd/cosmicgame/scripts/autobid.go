// Cosmic Signature Game - Automated Bidding Bot
// Uses event-based programming model with a state machine
package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)

// =============================================================================
// Constants
// =============================================================================

// Internal constants (not user-configurable)
const (
	RWALK_MINT_EVENT    = "ad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec"
	TIME_DELAY_SEC      = 1
	TIME_DELAY_ON_ERROR = 500 // milliseconds
	DELAY_AFTER_TX      = 2   // seconds
	DELAY_NO_ACTION     = 5   // seconds
	MAX_RETRIES         = 3   // if something doesn't work for 3 times, then we abandon the task
	BID_GAS_LIMIT       = 1000000
	CLAIM_GAS_LIMIT     = 5000000
)

// Default values for user-configurable parameters
const (
	DEFAULT_MAX_ETH_BID_ETHER      = 5     // in ETH (will be converted to wei)
	DEFAULT_MAX_CST_BID_AMOUNT     = 9     // in CST tokens (will be converted to wei)
	DEFAULT_RWALK_BID_START_PRICE  = 0.1   // in ETH - only use RWALK when bid price above this
	DEFAULT_TIME_UNTIL_PRIZE_LIMIT = 15    // seconds before prize to start bidding
	DEFAULT_CST_BID_ANYWAY         = true  // keep bidding with CST even when last bidder
)

// FlowState represents the current state in the event-based state machine
type FlowState int

const (
	FlowUninitialized FlowState = iota
	FlowNotLastBidder
	FlowIAmLastBidder
	FlowNeedToBidWithCST
	FlowWaitingForCSTBidTx
	FlowNeedToBidWithETH
	FlowWaitingForETHBidTx
	FlowNeedToBidWithRWalk
	FlowNeedToClaimPrize
	FlowWaitingForClaimPrizeTx
	FlowWaitingForRWalkBidTx
	FlowWaitingForRWalkMint
	FlowNeedToSendRWalkBidTx
	FlowInitialBidding
)

// String returns the name of the flow state for logging
func (f FlowState) String() string {
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

// =============================================================================
// Configuration
// =============================================================================

// Config holds user-configurable parameters
type Config struct {
	MaxEthBid          *big.Int // maximum ETH to spend on bidding (wei)
	MaxCstBid          *big.Int // max CST price for bidding (wei)
	RWalkMinPrice      *big.Int // only use RWALK when bid price above this (wei)
	TimeBeforePrize    int64    // seconds before prize to start bidding
	CstBidAnyway       bool     // keep bidding with CST even when last bidder
	InitialBidPrice    *big.Int // initial bid price level (optional)
	RpcURL             string
	PrivateKeyHex      string
	GameContractAddr   string
}

// =============================================================================
// BiddingBot - Main struct encapsulating all state
// =============================================================================

// BiddingBot encapsulates all state and logic for automated bidding
type BiddingBot struct {
	// Configuration
	config Config

	// Ethereum clients
	rpcClient *ethrpc.Client
	ethClient *ethclient.Client
	callOpts  bind.CallOpts

	// Contracts
	gameContract   *CosmicSignatureGame
	rwalkContract  *RWalk
	prizesWallet   *PrizesWallet

	// Account info
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    *big.Int

	// State machine
	flowState        FlowState
	roundNumPlayed   int64
	lastBidderNotified bool

	// Pending transaction tracking
	ethBidTxHash     common.Hash
	cstBidTxHash     common.Hash
	rwalkMintTxHash  common.Hash
	rwalkBidTxHash   common.Hash
	claimPrizeTxHash common.Hash
	retriesCounter   int

	// RWalk token tracking
	nextRWalkTokenID int64
	prevRWalkTokenID int64

	// Current loop variables (refreshed each iteration)
	gasPrice       *big.Int
	cstPrice       *big.Int
	bidPrice       *big.Int
	timeUntilPrize *big.Int
	cstBalance     *big.Int
	ethBalance     *big.Int
	lastBidder     common.Address
	rwalkMintPrice *big.Int
}

// =============================================================================
// Bot Initialization
// =============================================================================

// NewBiddingBot creates and initializes a new BiddingBot
func NewBiddingBot(cfg Config) (*BiddingBot, error) {
	bot := &BiddingBot{
		config:           cfg,
		flowState:        FlowUninitialized,
		roundNumPlayed:   -1,
		nextRWalkTokenID: -1,
		prevRWalkTokenID: -1,
	}

	// Connect to RPC
	var err error
	bot.rpcClient, err = ethrpc.DialContext(context.Background(), cfg.RpcURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to ETH RPC: %v", err)
	}
	bot.ethClient = ethclient.NewClient(bot.rpcClient)

	// Get chain ID
	bot.chainID, err = bot.ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting chain id: %v", err)
	}

	// Setup private key and address
	bot.privateKey, err = crypto.HexToECDSA(cfg.PrivateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}
	publicKey := bot.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("couldn't derive public key")
	}
	bot.address = crypto.PubkeyToAddress(*publicKeyECDSA)

	// Setup game contract
	gameAddr := common.HexToAddress(cfg.GameContractAddr)
	bot.gameContract, err = NewCosmicSignatureGame(gameAddr, bot.ethClient)
		if err != nil {
		return nil, fmt.Errorf("failed to instantiate CosmicGame contract: %v", err)
	}

	// Setup RWalk contract
	rwalkAddr, err := bot.gameContract.RandomWalkNft(&bot.callOpts)
	if err != nil {
		return nil, fmt.Errorf("error getting RWalk addr: %v", err)
	}
	bot.rwalkContract, err = NewRWalk(rwalkAddr, bot.ethClient)
	if err != nil {
		return nil, fmt.Errorf("error creating RWalk instance: %v", err)
	}

	// Setup prizes wallet contract
	prizesAddr, err := bot.gameContract.PrizesWallet(&bot.callOpts)
	if err != nil {
		return nil, fmt.Errorf("error fetching PrizesWallet address: %v", err)
	}
	bot.prizesWallet, err = NewPrizesWallet(prizesAddr, bot.ethClient)
	if err != nil {
		return nil, fmt.Errorf("error creating PrizesWallet instance: %v", err)
	}

	// Get initial round number
	roundNum, err := bot.gameContract.RoundNum(&bot.callOpts)
	if err != nil {
		return nil, fmt.Errorf("error getting roundNum: %v", err)
	}
	bot.roundNumPlayed = roundNum.Int64()

	// Get initial gas price
	bot.gasPrice, err = bot.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting suggested gas price: %v", err)
	}

	return bot, nil
}

// =============================================================================
// Helper Methods
// =============================================================================

// fmtEth formats wei as ETH string
func fmtEth(wei *big.Int) string {
	if wei == nil {
		return "0"
	}
	ether := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return ethValue.Text('f', 18)
}

// isMyAddress checks if the given address matches the bot's address
func (bot *BiddingBot) isMyAddress(addr common.Address) bool {
	return bytes.Equal(addr.Bytes(), bot.address.Bytes())
}

// isZeroAddress checks if the given address is zero
func isZeroAddress(addr common.Address) bool {
	var zero common.Address
	return bytes.Equal(addr.Bytes(), zero.Bytes())
}

// createTransactOpts creates transaction options with signing
// Refreshes gas price for every transaction to ensure accurate pricing
func (bot *BiddingBot) createTransactOpts(value *big.Int, gasLimit uint64) (*bind.TransactOpts, error) {
	nonce, err := bot.ethClient.PendingNonceAt(context.Background(), bot.address)
	if err != nil {
		return nil, fmt.Errorf("error getting nonce: %v", err)
	}

	// Refresh gas price for this transaction
	gasPrice, err := bot.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting gas price: %v", err)
	}
	bot.gasPrice = gasPrice // Update cached value too

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
			return nil, fmt.Errorf("error signing: %v", err)
		}
		return tx.WithSignature(signer, signature)
	}

	return txopts, nil
}

// refreshMarketData fetches current market data from contracts
func (bot *BiddingBot) refreshMarketData() error {
	var err error

	bot.cstPrice, err = bot.gameContract.GetNextCstBidPrice(&bot.callOpts)
		if err != nil {
		return fmt.Errorf("error getting CST bid price: %v", err)
		}

	bot.bidPrice, err = bot.gameContract.GetNextEthBidPrice(&bot.callOpts)
		if err != nil {
		return fmt.Errorf("error getting ETH bid price: %v", err)
		}

	bot.timeUntilPrize, err = bot.gameContract.GetDurationUntilMainPrize(&bot.callOpts)
		if err != nil {
		return fmt.Errorf("error getting time until prize: %v", err)
		}

	tokenAddr, err := bot.gameContract.Token(&bot.callOpts)
		if err != nil {
		return fmt.Errorf("error getting token address: %v", err)
		}
	tokenContract, err := NewERC20(tokenAddr, bot.ethClient)
		if err != nil {
		return fmt.Errorf("error instantiating token contract: %v", err)
		}
	bot.cstBalance, err = tokenContract.BalanceOf(&bot.callOpts, bot.address)
		if err != nil {
		return fmt.Errorf("error getting CST balance: %v", err)
		}

	bot.lastBidder, err = bot.gameContract.LastBidderAddress(&bot.callOpts)
		if err != nil {
		return fmt.Errorf("error getting last bidder: %v", err)
		}

	bot.ethBalance, err = bot.ethClient.BalanceAt(context.Background(), bot.address, nil)
		if err != nil {
		return fmt.Errorf("error getting ETH balance: %v", err)
		}

	bot.rwalkMintPrice, err = bot.rwalkContract.GetMintPrice(&bot.callOpts)
		if err != nil {
		return fmt.Errorf("error getting RWalk mint price: %v", err)
	}

	return nil
}

// checkTimeoutClaim checks if we can claim prize due to timeout
func (bot *BiddingBot) checkTimeoutClaim() {
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
			fmt.Println("Winner didn't claim prize during claim window, I am going to claim it")
			bot.flowState = FlowNeedToClaimPrize
		}
	}
}

// checkRoundChange checks if the round has changed
func (bot *BiddingBot) checkRoundChange() bool {
	rnum, err := bot.gameContract.RoundNum(&bot.callOpts)
	if err != nil {
		fmt.Printf("Error getting roundNum: %v\n", err)
		return false
	}

	if rnum.Int64() != bot.roundNumPlayed {
		fmt.Printf("Round changed (was %v, now %v)\n", bot.roundNumPlayed, rnum)
		
		// SAFETY: Detect blockchain reset (round went backwards)
		if rnum.Int64() < bot.roundNumPlayed {
			fmt.Printf("ERROR: Round number decreased (%v -> %v) - blockchain was reset!\n", 
				bot.roundNumPlayed, rnum.Int64())
			fmt.Println("Exiting to prevent unintended spending. Restart bot manually.")
			os.Exit(1)
		}
		
		winner, err := bot.prizesWallet.MainPrizeBeneficiaryAddresses(&bot.callOpts, big.NewInt(bot.roundNumPlayed))
		if err == nil {
			if bot.isMyAddress(winner) {
				fmt.Printf("I am the winner of round %v!\n", bot.roundNumPlayed)
			} else {
				fmt.Printf("I am not the winner of round %v\n", bot.roundNumPlayed)
			}
		}

		if bot.config.InitialBidPrice != nil {
			fmt.Printf("Playing new round with initial bids\n")
			bot.roundNumPlayed = rnum.Int64()
			bot.flowState = FlowInitialBidding
			return false
		}
				os.Exit(0)
			}
	return false
}

// findRWalkTokenID searches for an unused RWalk token owned by the bot
func (bot *BiddingBot) findRWalkTokenID() {
	if bot.nextRWalkTokenID > -1 {
		wasUsed, err := bot.gameContract.UsedRandomWalkNfts(&bot.callOpts, big.NewInt(bot.nextRWalkTokenID))
		if err == nil && wasUsed.Cmp(big.NewInt(1)) == 0 {
			fmt.Printf("Resetting nextRWalkTokenID (%v) - already used\n", bot.nextRWalkTokenID)
			bot.nextRWalkTokenID = -1
		} else if err == nil {
			return // Already have a valid token
		}
	}

	targetID := bot.prevRWalkTokenID + 1
	lastTokenID, err := bot.rwalkContract.NextTokenId(&bot.callOpts)
	if err != nil {
		fmt.Printf("Error calling NextTokenId(): %v\n", err)
		return
	}

	for targetID < lastTokenID.Int64() {
		owner, err := bot.rwalkContract.OwnerOf(&bot.callOpts, big.NewInt(targetID))
		if err != nil {
			time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			return
		}
		bot.prevRWalkTokenID = targetID

		if bot.isMyAddress(owner) {
			wasUsed, err := bot.gameContract.UsedRandomWalkNfts(&bot.callOpts, big.NewInt(targetID))
			if err == nil && wasUsed.Cmp(big.NewInt(0)) == 0 {
				bot.nextRWalkTokenID = targetID
				fmt.Printf("Found RWalk token %v for bidding\n", bot.nextRWalkTokenID)
				return
			}
		}
		targetID++
	}
}

// canBidWithCST checks if CST bidding conditions are met
func (bot *BiddingBot) canBidWithCST() bool {
	return bot.config.MaxCstBid.Cmp(bot.cstBalance) <= 0 &&
		bot.cstPrice.Cmp(bot.config.MaxCstBid) <= 0
}

// =============================================================================
// Flow Handlers
// =============================================================================

// handleUninitialized handles the initial state
func (bot *BiddingBot) handleUninitialized() FlowState {
	if bot.isMyAddress(bot.lastBidder) {
		if !bot.lastBidderNotified {
			fmt.Println("I am last bidder")
			bot.lastBidderNotified = true
		}
		return FlowIAmLastBidder
	}
	fmt.Printf("I am not the last bidder (time until prize = %v)\n", bot.timeUntilPrize)
	bot.lastBidderNotified = false
	return FlowNotLastBidder
}

// handleNotLastBidder handles logic when bot is not the last bidder
func (bot *BiddingBot) handleNotLastBidder() (FlowState, bool) {
	// Check if we should bid with CST anyway
	if bot.config.CstBidAnyway && bot.canBidWithCST() {
		fmt.Printf("CST price (%v) below limit, bidding with CST\n", fmtEth(bot.cstPrice))
		return FlowNeedToBidWithCST, true
	}

	// Re-verify last bidder status
	if bot.isMyAddress(bot.lastBidder) {
		if !bot.lastBidderNotified {
			fmt.Println("I am last bidder")
			bot.lastBidderNotified = true
		}
		return FlowIAmLastBidder, false
	}

	// Check if it's time to bid
	if bot.timeUntilPrize.Cmp(big.NewInt(bot.config.TimeBeforePrize)) <= 0 {
		fmt.Printf("%v sec before prize, time to bid\n", bot.timeUntilPrize.Int64())
		return bot.decideBidType()
	}

	fmt.Printf("Not my time to bid yet (timeUntilPrize = %v)\n", bot.timeUntilPrize.Int64())
	return FlowNotLastBidder, false
}

// decideBidType decides which bidding method to use
func (bot *BiddingBot) decideBidType() (FlowState, bool) {
	// Try CST first
	if bot.canBidWithCST() {
		fmt.Printf("CST price (%v) below limit, bidding with CST\n", fmtEth(bot.cstPrice))
		return FlowNeedToBidWithCST, true
	}

	// Can't bid with CST, check ETH options
	if bot.config.MaxCstBid.Cmp(bot.cstBalance) > 0 {
		fmt.Printf("Not enough CST balance for bid\n")
							} else { 
		fmt.Printf("CST price above limit\n")
	}

	// Check if RWalk bidding is allowed
	if bot.config.RWalkMinPrice.Cmp(bot.bidPrice) < 0 {
		rwalkDiscountedPrice := new(big.Int).Quo(bot.bidPrice, big.NewInt(2))
		bidWithRwalkPrice := new(big.Int).Add(bot.rwalkMintPrice, rwalkDiscountedPrice)

		fmt.Printf("RWALK+ETH costs %v, pure ETH costs %v\n", fmtEth(bidWithRwalkPrice), fmtEth(bot.bidPrice))

		if bot.bidPrice.Cmp(bidWithRwalkPrice) <= 0 {
			// Plain ETH is cheaper
			return bot.tryPlainEthBid()
								} else {
			// RWalk is cheaper
			if bidWithRwalkPrice.Cmp(bot.config.MaxEthBid) < 0 {
				fmt.Println("Bidding with RWalk (cheaper)")
				return FlowNeedToBidWithRWalk, true
			}
			fmt.Printf("Out of funds even with RWalk\n")
			return FlowNotLastBidder, false
		}
	}

	// Plain ETH only
	return bot.tryPlainEthBid()
}

// tryPlainEthBid attempts to bid with plain ETH
func (bot *BiddingBot) tryPlainEthBid() (FlowState, bool) {
	if bot.config.MaxEthBid.Cmp(bot.bidPrice) < 0 {
		fmt.Printf("ETH bid price (%v) above limit\n", fmtEth(bot.bidPrice))
		return FlowNotLastBidder, false
	}
	if bot.bidPrice.Cmp(bot.ethBalance) >= 0 {
		fmt.Printf("Insufficient ETH balance\n")
		return FlowNotLastBidder, false
	}
	fmt.Println("Bidding with plain ETH")
	return FlowNeedToBidWithETH, true
}

// handleIAmLastBidder handles logic when bot is the last bidder
func (bot *BiddingBot) handleIAmLastBidder() (FlowState, bool) {
	// Use cached lastBidder (already fetched in refreshMarketData)
	// Don't make duplicate contract calls that could cause inconsistency
	if !bot.isMyAddress(bot.lastBidder) {
		fmt.Printf("No longer last bidder\n")
		bot.lastBidderNotified = false
		return FlowNotLastBidder, false // Don't continue - need to refresh market data
	}

	// Check if we should keep bidding with CST
	if bot.config.CstBidAnyway && bot.canBidWithCST() {
		fmt.Printf("CST price low, bidding again\n")
		return FlowNeedToBidWithCST, true
	}

	// Check if we can claim prize
	if bot.timeUntilPrize.Cmp(big.NewInt(0)) == 0 {
		return FlowNeedToClaimPrize, true
	}

	return FlowIAmLastBidder, false
}

// handleBidWithCST handles CST bidding
func (bot *BiddingBot) handleBidWithCST() (FlowState, bool) {
	txopts, err := bot.createTransactOpts(nil, BID_GAS_LIMIT)
					if err != nil {
		fmt.Printf("Error creating tx opts: %v\n", err)
		return FlowUninitialized, false
	}

	tx, err := bot.gameContract.BidWithCst(txopts, bot.cstPrice, "")
					if err != nil {
		fmt.Printf("BidWithCST error: %v\n", err)
		return FlowUninitialized, false
	}

	fmt.Printf("CST bid tx: %v\n", tx.Hash().Hex())
	bot.cstBidTxHash = tx.Hash()
	bot.retriesCounter = 0
	return FlowWaitingForCSTBidTx, false
}

// handleWaitForCSTBidTx waits for CST bid transaction receipt
func (bot *BiddingBot) handleWaitForCSTBidTx() (FlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.cstBidTxHash)
					if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= MAX_RETRIES {
			fmt.Println("Max retries reached for CST bid tx")
			return FlowUninitialized, false
		}
		return FlowWaitingForCSTBidTx, false
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		fmt.Println("CST bid tx failed")
		return FlowUninitialized, false
	}

	return bot.checkLastBidderAfterBid("CST")
}

// handleBidWithETH handles plain ETH bidding
func (bot *BiddingBot) handleBidWithETH() (FlowState, bool) {
	txopts, err := bot.createTransactOpts(bot.bidPrice, BID_GAS_LIMIT)
						if err != nil {
		fmt.Printf("Error creating tx opts: %v\n", err)
		return FlowUninitialized, false
	}

	tx, err := bot.gameContract.BidWithEth(txopts, big.NewInt(-1), "")
					if err != nil {
		fmt.Printf("BidWithEth error: %v\n", err)
		return FlowUninitialized, false
	}

	fmt.Printf("ETH bid tx (%v ETH): %v\n", fmtEth(bot.bidPrice), tx.Hash().Hex())
	bot.ethBidTxHash = tx.Hash()
	bot.retriesCounter = 0
	return FlowWaitingForETHBidTx, false
}

// handleWaitForETHBidTx waits for ETH bid transaction receipt
func (bot *BiddingBot) handleWaitForETHBidTx() (FlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.ethBidTxHash)
						if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= MAX_RETRIES {
			fmt.Println("Max retries reached for ETH bid tx")
			return FlowUninitialized, false
		}
		return FlowWaitingForETHBidTx, false
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		fmt.Println("ETH bid tx failed")
		return FlowUninitialized, false
	}

	return bot.checkLastBidderAfterBid("ETH")
}

// handleBidWithRWalk handles RWalk bidding flow
func (bot *BiddingBot) handleBidWithRWalk() (FlowState, bool) {
	if bot.nextRWalkTokenID > -1 {
		fmt.Printf("Using pre-minted RWalk token %v\n", bot.nextRWalkTokenID)
		return FlowNeedToSendRWalkBidTx, true
	}

	fmt.Println("Need to mint RWalk token first")
	txopts, err := bot.createTransactOpts(bot.rwalkMintPrice, BID_GAS_LIMIT)
					if err != nil {
		fmt.Printf("Error creating tx opts: %v\n", err)
		return FlowUninitialized, false
	}

	tx, err := bot.rwalkContract.Mint(txopts)
						if err != nil {
		fmt.Printf("RWalk mint error: %v\n", err)
		return FlowUninitialized, false
	}

	fmt.Printf("RWalk mint tx: %v\n", tx.Hash().Hex())
	bot.rwalkMintTxHash = tx.Hash()
	bot.retriesCounter = 0
	return FlowWaitingForRWalkMint, false
}

// handleWaitForRWalkMint waits for RWalk mint transaction
func (bot *BiddingBot) handleWaitForRWalkMint() (FlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.rwalkMintTxHash)
					if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= MAX_RETRIES {
			fmt.Println("Max retries reached for RWalk mint tx")
			return FlowUninitialized, false
		}
		return FlowWaitingForRWalkMint, false
	}

	evtMintEvent, _ := hex.DecodeString(RWALK_MINT_EVENT)
	for _, elog := range receipt.Logs {
		if len(elog.Topics) > 0 && bytes.Equal(elog.Topics[0].Bytes(), evtMintEvent) {
			bot.nextRWalkTokenID = elog.Topics[1].Big().Int64()
			fmt.Printf("Minted RWalk token %v\n", bot.nextRWalkTokenID)
			return FlowNeedToSendRWalkBidTx, true
		}
	}

	fmt.Println("RWalk mint event not found")
	return FlowUninitialized, false
}

// handleSendRWalkBidTx sends bid with RWalk token
func (bot *BiddingBot) handleSendRWalkBidTx() (FlowState, bool) {
	txopts, err := bot.createTransactOpts(bot.bidPrice, BID_GAS_LIMIT)
					if err != nil {
		fmt.Printf("Error creating tx opts: %v\n", err)
		return FlowUninitialized, false
	}

	tx, err := bot.gameContract.BidWithEth(txopts, big.NewInt(bot.nextRWalkTokenID), "")
						if err != nil {
		fmt.Printf("BidWithEth (RWalk) error: %v\n", err)
		return FlowUninitialized, false
	}

	fmt.Printf("RWalk bid tx: %v (token %v)\n", tx.Hash().Hex(), bot.nextRWalkTokenID)
	bot.rwalkBidTxHash = tx.Hash()
	bot.retriesCounter = 0
	return FlowWaitingForRWalkBidTx, false
}

// handleWaitForRWalkBidTx waits for RWalk bid transaction
func (bot *BiddingBot) handleWaitForRWalkBidTx() (FlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.rwalkBidTxHash)
					if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= MAX_RETRIES {
			fmt.Println("Max retries reached for RWalk bid tx")
			return FlowUninitialized, false
		}
		return FlowWaitingForRWalkBidTx, false
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		fmt.Println("RWalk bid tx failed")
		return FlowUninitialized, false
	}

	bot.nextRWalkTokenID = -1
	return bot.checkLastBidderAfterBid("RWalk")
}

// handleClaimPrize handles claiming the prize
func (bot *BiddingBot) handleClaimPrize() (FlowState, bool) {
	txopts, err := bot.createTransactOpts(nil, CLAIM_GAS_LIMIT)
					if err != nil {
		fmt.Printf("Error creating tx opts: %v\n", err)
		return FlowUninitialized, false
	}

	tx, err := bot.gameContract.ClaimMainPrize(txopts)
						if err != nil {
		fmt.Printf("ClaimMainPrize error: %v\n", err)
		return FlowUninitialized, false
	}

	fmt.Printf("ClaimPrize tx: %v\n", tx.Hash().Hex())
	bot.claimPrizeTxHash = tx.Hash()
	bot.retriesCounter = 0
	return FlowWaitingForClaimPrizeTx, false
}

// handleWaitForClaimPrizeTx waits for claim prize transaction
func (bot *BiddingBot) handleWaitForClaimPrizeTx() (FlowState, bool) {
	receipt, err := bot.ethClient.TransactionReceipt(context.Background(), bot.claimPrizeTxHash)
					if err != nil {
		bot.retriesCounter++
		if bot.retriesCounter >= MAX_RETRIES {
			fmt.Println("Max retries reached for claim prize tx")
			return FlowUninitialized, false
		}
		return FlowWaitingForClaimPrizeTx, false
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("Prize claimed successfully!")
					} else {
		fmt.Println("Claim prize tx failed")
	}
	return FlowUninitialized, false
}

// handleInitialBidding handles the initial bidding phase
func (bot *BiddingBot) handleInitialBidding() (FlowState, bool) {
	fmt.Printf("Initial bidding: price=%v, limit=%v\n", fmtEth(bot.bidPrice), fmtEth(bot.config.InitialBidPrice))

	const maxInitialBids = 20 // Safety limit to prevent runaway bidding
	bidCount := 0
	totalSpent := big.NewInt(0)

	for failures := 0; failures <= 5; {
		// Safety check: max bid count
		if bidCount >= maxInitialBids {
			fmt.Printf("SAFETY: Reached max initial bids (%d), stopping\n", maxInitialBids)
			break
		}

		bidPrice, err := bot.gameContract.GetNextEthBidPrice(&bot.callOpts)
					if err != nil {
			fmt.Printf("Error getting bid price: %v\n", err)
			failures++
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}

		if bot.config.InitialBidPrice.Cmp(bidPrice) < 0 {
			fmt.Println("Reached bid price limit, stopping initial bidding")
			break
		}

		// Safety check: don't exceed ETH balance
		if bidPrice.Cmp(bot.ethBalance) >= 0 {
			fmt.Println("Insufficient ETH balance for initial bid, stopping")
			break
		}

		txopts, err := bot.createTransactOpts(bidPrice, BID_GAS_LIMIT)
		if err != nil {
			failures++
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}

		tx, err := bot.gameContract.BidWithEth(txopts, big.NewInt(-1), "")
		if err != nil {
			fmt.Printf("Bid error: %v\n", err)
			failures++
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
			continue
		}

		bidCount++
		totalSpent.Add(totalSpent, bidPrice)
		fmt.Printf("Initial bid #%d tx (%v ETH): %v\n", bidCount, fmtEth(bidPrice), tx.Hash().Hex())
						time.Sleep(DELAY_AFTER_TX * time.Second)

		// Wait for receipt
		for i := 0; i < 5; i++ {
			time.Sleep(DELAY_AFTER_TX * time.Second)
			receipt, err := bot.ethClient.TransactionReceipt(context.Background(), tx.Hash())
			if err == nil {
				if receipt.Status == types.ReceiptStatusSuccessful {
					fmt.Println("Initial bid successful")
				} else {
					fmt.Println("Initial bid tx failed, stopping")
					failures = 6 // Force exit
				}
				break
			}
		}
		failures = 0

		// Refresh ETH balance after each bid
		bot.ethBalance, _ = bot.ethClient.BalanceAt(context.Background(), bot.address, nil)
	}

	fmt.Printf("Initial bidding finished: %d bids, %v ETH spent\n", bidCount, fmtEth(totalSpent))
	return FlowUninitialized, false
}

// checkLastBidderAfterBid checks if we're last bidder after a bid
func (bot *BiddingBot) checkLastBidderAfterBid(bidType string) (FlowState, bool) {
	lastBidder, err := bot.gameContract.LastBidderAddress(&bot.callOpts)
					if err != nil {
		fmt.Printf("Error checking last bidder: %v\n", err)
		return FlowUninitialized, false
	}

	if bot.isMyAddress(lastBidder) {
		fmt.Printf("I am last bidder after %s bid\n", bidType)
		bot.lastBidderNotified = true
		return FlowIAmLastBidder, false
	}

	fmt.Printf("Not last bidder after %s bid\n", bidType)
	bot.lastBidderNotified = false
	return FlowNotLastBidder, false
}

// =============================================================================
// Main Event Loop
// =============================================================================

// Run starts the main event loop
func (bot *BiddingBot) Run() {
	fmt.Printf("Playing round %v\n", bot.roundNumPlayed)
	bot.printConfig()

	for {
		fmt.Printf("\n=== Event loop (flow=%s, rwalkNext=%v) ===\n", 
			bot.flowState, bot.nextRWalkTokenID)

		// Search for RWalk tokens in background
		if bot.nextRWalkTokenID == -1 {
			go bot.findRWalkTokenID()
		}

		// Refresh market data
		if err := bot.refreshMarketData(); err != nil {
			fmt.Printf("Error refreshing data: %v\n", err)
						time.Sleep(TIME_DELAY_ON_ERROR * time.Millisecond)
						continue
					}

		// Check for timeout claim opportunity
		bot.checkTimeoutClaim()

		// Check for round change
		if bot.checkRoundChange() {
								continue
							}

		// Process flow states
		bot.processFlowStates()

		// Refresh gas price
		if gasPrice, err := bot.ethClient.SuggestGasPrice(context.Background()); err == nil {
			bot.gasPrice = gasPrice
		}

		time.Sleep(TIME_DELAY_SEC * time.Second)
	}
}

// processFlowStates processes the state machine
func (bot *BiddingBot) processFlowStates() {
	for {
		prevState := bot.flowState
		fmt.Printf("Processing state: %s (timeUntilPrize=%v)\n", bot.flowState, bot.timeUntilPrize.Int64())

		var continueProcessing bool
		var sleepDuration time.Duration

		switch bot.flowState {
		case FlowUninitialized:
			bot.flowState = bot.handleUninitialized()
			if bot.flowState == FlowIAmLastBidder {
				sleepDuration = DELAY_NO_ACTION * time.Second
			}

		case FlowNotLastBidder:
			bot.flowState, continueProcessing = bot.handleNotLastBidder()
			if !continueProcessing && bot.flowState == FlowNotLastBidder {
				sleepDuration = DELAY_NO_ACTION * time.Second
			}

		case FlowIAmLastBidder:
			bot.flowState, continueProcessing = bot.handleIAmLastBidder()
			if !continueProcessing && bot.flowState == FlowIAmLastBidder {
				if bot.timeUntilPrize.Cmp(big.NewInt(DELAY_NO_ACTION)) >= 0 {
					sleepDuration = DELAY_NO_ACTION * time.Second
				}
			}

		case FlowNeedToBidWithCST:
			bot.flowState, _ = bot.handleBidWithCST()
			sleepDuration = DELAY_AFTER_TX * time.Second

		case FlowWaitingForCSTBidTx:
			bot.flowState, _ = bot.handleWaitForCSTBidTx()
			if bot.flowState == FlowWaitingForCSTBidTx {
				sleepDuration = TIME_DELAY_ON_ERROR * time.Millisecond
			}

		case FlowNeedToBidWithETH:
			bot.flowState, _ = bot.handleBidWithETH()
			sleepDuration = DELAY_AFTER_TX * time.Second

		case FlowWaitingForETHBidTx:
			bot.flowState, _ = bot.handleWaitForETHBidTx()
			if bot.flowState == FlowWaitingForETHBidTx {
				sleepDuration = TIME_DELAY_ON_ERROR * time.Millisecond
			}

		case FlowNeedToBidWithRWalk:
			bot.flowState, continueProcessing = bot.handleBidWithRWalk()
			if !continueProcessing {
				sleepDuration = DELAY_AFTER_TX * time.Second
			}

		case FlowWaitingForRWalkMint:
			bot.flowState, continueProcessing = bot.handleWaitForRWalkMint()
			if bot.flowState == FlowWaitingForRWalkMint {
				sleepDuration = TIME_DELAY_ON_ERROR * time.Millisecond
			}

		case FlowNeedToSendRWalkBidTx:
			bot.flowState, _ = bot.handleSendRWalkBidTx()
			sleepDuration = DELAY_AFTER_TX * time.Second

		case FlowWaitingForRWalkBidTx:
			bot.flowState, _ = bot.handleWaitForRWalkBidTx()
			if bot.flowState == FlowWaitingForRWalkBidTx {
				sleepDuration = TIME_DELAY_ON_ERROR * time.Millisecond
			}

		case FlowNeedToClaimPrize:
			bot.flowState, _ = bot.handleClaimPrize()
			sleepDuration = DELAY_AFTER_TX * time.Second

		case FlowWaitingForClaimPrizeTx:
			bot.flowState, _ = bot.handleWaitForClaimPrizeTx()
			if bot.flowState == FlowWaitingForClaimPrizeTx {
				sleepDuration = TIME_DELAY_ON_ERROR * time.Millisecond
			}

		case FlowInitialBidding:
			bot.flowState, _ = bot.handleInitialBidding()

		default:
			fmt.Printf("Unknown flow state: %v\n", bot.flowState)
			sleepDuration = DELAY_NO_ACTION * time.Second
		}

		// Sleep if needed
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}

		// Break if state didn't change (need to refresh data)
		if bot.flowState == prevState && !continueProcessing {
							break
						}
	}
}

// printConfig prints the current configuration
func (bot *BiddingBot) printConfig() {
	fmt.Println("Config params:")
	fmt.Printf("  MAX_ETH_BID: %v ETH\n", fmtEth(bot.config.MaxEthBid))
	fmt.Printf("  MAX_CST_BID: %v CST\n", fmtEth(bot.config.MaxCstBid))
	fmt.Printf("  RWALK_MIN_PRICE: %v ETH\n", fmtEth(bot.config.RWalkMinPrice))
	fmt.Printf("  TIME_BEFORE_PRIZE: %v secs\n", bot.config.TimeBeforePrize)
	fmt.Printf("  CST_BID_ANYWAY: %v\n", bot.config.CstBidAnyway)
	if bot.config.InitialBidPrice != nil {
		fmt.Printf("  INITIAL_BID_PRICE: %v ETH\n", fmtEth(bot.config.InitialBidPrice))
	}
}

// =============================================================================
// Configuration Loading
// =============================================================================

func loadConfig() Config {
	validateRequiredEnvVars()

	cfg := Config{
		RpcURL:           os.Getenv("RPC_URL"),
		PrivateKeyHex:    os.Getenv("PKEY_HEX"),
		GameContractAddr: os.Getenv("CGAME_ADDR"),
		MaxEthBid:        getEnvBigIntEth("MAX_ETH_BID", DEFAULT_MAX_ETH_BID_ETHER),
		MaxCstBid:        getEnvBigIntEth("MAX_CST_BID", DEFAULT_MAX_CST_BID_AMOUNT),
		RWalkMinPrice:    getEnvBigIntEth("RWALK_MIN_PRICE", DEFAULT_RWALK_BID_START_PRICE),
		TimeBeforePrize:  getEnvInt64("TIME_BEFORE_PRIZE", DEFAULT_TIME_UNTIL_PRIZE_LIMIT),
		CstBidAnyway:     getEnvBool("CST_BID_ANYWAY", DEFAULT_CST_BID_ANYWAY),
	}

	// Parse optional initial bid price
	if level := os.Getenv("INITIAL_BID_PRICE_LEVEL"); level != "" {
		if val, valid := new(big.Int).SetString(level, 10); valid {
			cfg.InitialBidPrice = new(big.Int).Mul(big.NewInt(1e15), val)
		}
	}

	return cfg
}

func validateRequiredEnvVars() {
	var errors []string

	if os.Getenv("RPC_URL") == "" {
		errors = append(errors, "RPC_URL is required")
	}
	
	pkey := os.Getenv("PKEY_HEX")
	if pkey == "" {
		errors = append(errors, "PKEY_HEX is required")
	} else if len(pkey) != 64 {
		errors = append(errors, fmt.Sprintf("PKEY_HEX must be 64 chars (got %d)", len(pkey)))
	}

	addr := os.Getenv("CGAME_ADDR")
	if addr == "" {
		errors = append(errors, "CGAME_ADDR is required")
	} else if len(addr) != 40 {
		errors = append(errors, fmt.Sprintf("CGAME_ADDR must be 40 chars (got %d)", len(addr)))
	}

	if len(errors) > 0 {
		fmt.Println("Configuration errors:")
		for _, e := range errors {
			fmt.Printf("  - %s\n", e)
		}
		fmt.Println("\nRequired: RPC_URL, PKEY_HEX, CGAME_ADDR")
		fmt.Println("Optional: MAX_ETH_BID, MAX_CST_BID, RWALK_MIN_PRICE, TIME_BEFORE_PRIZE, CST_BID_ANYWAY")
		os.Exit(1)
	}
}

func getEnvBigIntEth(key string, defaultVal float64) *big.Int {
	if val := os.Getenv(key); val != "" {
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			wei := new(big.Float).Mul(big.NewFloat(f), big.NewFloat(1e18))
			result, _ := wei.Int(nil)
			return result
		}
		fmt.Printf("Warning: invalid %s, using default\n", key)
	}
	wei := new(big.Float).Mul(big.NewFloat(defaultVal), big.NewFloat(1e18))
	result, _ := wei.Int(nil)
	return result
}

func getEnvInt64(key string, defaultVal int64) int64 {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			return i
		}
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return defaultVal
}

// =============================================================================
// Main Entry Point
// =============================================================================

func main() {
	cfg := loadConfig()

	bot, err := NewBiddingBot(cfg)
		if err != nil {
		fmt.Printf("Failed to initialize bot: %v\n", err)
		os.Exit(1)
		}

	bot.Run()
}
