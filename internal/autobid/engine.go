// Package autobid is the automated CosmicGame bidding engine behind
// `cgctl autobid`. The pure decision rules live in decide.go; the Engine
// owns the market-refresh loop, transaction submission and receipt
// tracking, reconnection, round-change detection and session statistics.
//
// Every external dependency is injected through Config: the RPC dialer, the
// sleep function (so tests run without real delays) and the log writer.
// Cancellation of the Run context replaces the legacy signal-channel
// shutdown.
package autobid

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"sync"
	"sync/atomic"
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
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// rwalkMintEventTopic is the topic0 of the RandomWalk mint event.
var rwalkMintEventTopic = common.HexToHash("0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec")

// Engine timing and safety constants (not operator-configurable).
const (
	loopDelay            = 1 * time.Second        // between event-loop iterations
	errorDelay           = 500 * time.Millisecond // back-off after an error
	afterTxDelay         = 2 * time.Second        // after submitting a transaction
	idleDelay            = 5 * time.Second        // when there is nothing to do
	reconnectDelay       = 5 * time.Second        // before a reconnection attempt
	statsRPCTimeout      = 5 * time.Second        // bound the optional shutdown balance read
	maxReceiptRetries    = 3                      // abandon a pending-tx wait after this many polls
	bidGasLimit          = 1000000
	claimGasLimit        = 5000000
	maxConsecutiveErrors = 5  // trigger reconnect after this many consecutive errors
	maxReconnectAttempts = 10 // max reconnection attempts before giving up
	maxInitialBids       = 20 // safety limit for the initial-bidding loop
)

// Config configures an Engine. RPCURL, PrivateKeyHex and GameAddr are
// required; Limits fields must all be set (cmd/cgctl fills defaults).
type Config struct {
	// RPCURL is the Ethereum JSON-RPC endpoint.
	RPCURL string
	// PrivateKeyHex is the signer key: 64 hex characters, no 0x prefix.
	PrivateKeyHex string
	// GameAddr is the CosmicGame proxy contract address.
	GameAddr common.Address
	// Limits are the operator bidding limits.
	Limits Limits
	// InitialBidPrice, when set, makes the bot open each new round by
	// bidding until the price reaches this level.
	InitialBidPrice *big.Int
	// GasPriceMultiplier scales the node-suggested gas price on every
	// transaction; zero applies ethtx.DefaultGasPriceMultiplier.
	GasPriceMultiplier float64
	// Out receives the bot's log lines; defaults to os.Stdout.
	Out io.Writer
	// Sleep pauses between loop iterations; it must respect ctx
	// cancellation. Defaults to a context-aware timer sleep. Tests inject a
	// no-op to run scripted rounds instantly.
	Sleep func(ctx context.Context, d time.Duration) error
	// Dial opens the RPC connection; defaults to ethrpc.DialContext.
	// Reconnection re-dials through the same function.
	Dial func(ctx context.Context, url string) (*ethrpc.Client, error)
}

// sessionStats tracks statistics for one Run.
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

// pendingTxKind labels which transaction the engine is waiting on.
type pendingTxKind int

const (
	pendingNone pendingTxKind = iota
	pendingCSTBid
	pendingETHBid
	pendingRWalkMint
	pendingRWalkBid
	pendingClaim
)

// String names the pending-transaction kind for logs.
func (k pendingTxKind) String() string {
	switch k {
	case pendingNone:
		return "None"
	case pendingCSTBid:
		return "CSTBid"
	case pendingETHBid:
		return "ETHBid"
	case pendingRWalkMint:
		return "RWalkMint"
	case pendingRWalkBid:
		return "RWalkBid"
	case pendingClaim:
		return "Claim"
	default:
		return fmt.Sprintf("Unknown(%d)", int(k))
	}
}

// Engine is the automated bidding bot. Create with New, drive with Run.
type Engine struct {
	cfg Config

	rpcClient *ethrpc.Client
	ethClient *ethclient.Client

	gameContract  *cgcontracts.CosmicSignatureGame
	rwalkContract *rwcontracts.RWalk
	prizesWallet  *cgcontracts.PrizesWallet

	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    *big.Int

	roundNumPlayed     int64
	lastBidderNotified bool
	// startInitialBidding is set by checkRoundChange when a new round should
	// open with the initial-bidding loop.
	startInitialBidding bool

	// Pending transaction tracking.
	pendingKind    pendingTxKind
	pendingTxHash  common.Hash
	pendingValue   *big.Int // ETH attached to the pending bid (stats)
	retriesCounter int

	// RWalk token tracking. Atomics: the token search runs in a background
	// goroutine (the legacy bot raced on plain int64 fields here).
	nextRWalkTokenID atomic.Int64
	prevRWalkTokenID atomic.Int64
	rwalkSearching   atomic.Bool

	// Current loop variables (refreshed each iteration).
	market Market

	stats             sessionStats
	consecutiveErrors int
	reconnectCount    int

	// outMu serializes writes to cfg.Out: the RWalk token search logs from
	// its background goroutine.
	outMu     sync.Mutex
	closeOnce sync.Once
}

// New connects to the RPC endpoint, loads the contracts and prepares the
// signing account state.
func New(ctx context.Context, cfg Config) (*Engine, error) {
	if cfg.Out == nil {
		cfg.Out = os.Stdout
	}
	if cfg.Sleep == nil {
		cfg.Sleep = sleepContext
	}
	if cfg.Dial == nil {
		cfg.Dial = ethrpc.DialContext
	}
	if cfg.GasPriceMultiplier <= 0 {
		cfg.GasPriceMultiplier = ethtx.DefaultGasPriceMultiplier
	}

	e := &Engine{
		cfg:            cfg,
		roundNumPlayed: -1,
		stats: sessionStats{
			startTime:     time.Now(),
			totalEthSpent: big.NewInt(0),
		},
	}
	e.nextRWalkTokenID.Store(-1)
	e.prevRWalkTokenID.Store(-1)
	ready := false
	defer func() {
		if !ready {
			e.Close()
		}
	}()

	var err error
	e.rpcClient, err = cfg.Dial(ctx, cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to ETH RPC: %w", err)
	}
	e.ethClient = ethclient.NewClient(e.rpcClient)

	e.chainID, err = e.ethClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting chain id: %w", err)
	}

	e.privateKey, err = crypto.HexToECDSA(cfg.PrivateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %w", err)
	}
	publicKey, ok := e.privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("couldn't derive public key")
	}
	e.address = crypto.PubkeyToAddress(*publicKey)

	if err := e.bindContracts(ctx); err != nil {
		return nil, err
	}

	roundNum, err := e.gameContract.RoundNum(callOpts(ctx))
	if err != nil {
		return nil, fmt.Errorf("getting roundNum: %w", err)
	}
	e.roundNumPlayed = roundNum.Int64()

	e.stats.startBalance, err = e.ethClient.BalanceAt(ctx, e.address, nil)
	if err != nil {
		return nil, fmt.Errorf("getting initial balance: %w", err)
	}

	ready = true
	return e, nil
}

// Close releases the Engine's Ethereum RPC client. It is idempotent and must
// be called after Run returns; an Engine is intended for one Run lifecycle.
func (e *Engine) Close() {
	if e == nil {
		return
	}
	e.closeOnce.Do(func() {
		if e.rpcClient != nil {
			e.rpcClient.Close()
		}
	})
}

func callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

// sleepContext sleeps for d unless ctx is cancelled first.
func sleepContext(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return ctx.Err()
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// bindContracts (re)instantiates the contract bindings against the current
// client: the game proxy plus the RandomWalk and PrizesWallet addresses it
// publishes.
func (e *Engine) bindContracts(ctx context.Context) error {
	var err error
	e.gameContract, err = cgcontracts.NewCosmicSignatureGame(e.cfg.GameAddr, e.ethClient)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicGame contract: %w", err)
	}
	rwalkAddr, err := e.gameContract.RandomWalkNft(callOpts(ctx))
	if err != nil {
		return fmt.Errorf("getting RWalk addr: %w", err)
	}
	e.rwalkContract, err = rwcontracts.NewRWalk(rwalkAddr, e.ethClient)
	if err != nil {
		return fmt.Errorf("creating RWalk instance: %w", err)
	}
	prizesAddr, err := e.gameContract.PrizesWallet(callOpts(ctx))
	if err != nil {
		return fmt.Errorf("fetching PrizesWallet address: %w", err)
	}
	e.prizesWallet, err = cgcontracts.NewPrizesWallet(prizesAddr, e.ethClient)
	if err != nil {
		return fmt.Errorf("creating PrizesWallet instance: %w", err)
	}
	return nil
}

// logf prints one timestamped log line to the configured writer.
func (e *Engine) logf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	e.outMu.Lock()
	defer e.outMu.Unlock()
	fmt.Fprintf(e.cfg.Out, "[%s] %s\n", time.Now().Format("15:04:05"), msg)
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

// isMyAddress reports whether addr is the bot's signing address.
func (e *Engine) isMyAddress(addr common.Address) bool {
	return addr == e.address
}

// reconnect re-establishes the RPC connection and re-binds the contracts.
func (e *Engine) reconnect(ctx context.Context) error {
	e.reconnectCount++
	e.stats.reconnectCount++
	e.logf("[RECONNECT] Attempting reconnection (%d/%d)...", e.reconnectCount, maxReconnectAttempts)

	if e.reconnectCount > maxReconnectAttempts {
		return fmt.Errorf("max reconnection attempts (%d) exceeded", maxReconnectAttempts)
	}

	if e.rpcClient != nil {
		e.rpcClient.Close()
	}

	if err := e.cfg.Sleep(ctx, reconnectDelay); err != nil {
		return err
	}

	rpcClient, err := e.cfg.Dial(ctx, e.cfg.RPCURL)
	if err != nil {
		return fmt.Errorf("reconnect failed - can't connect to ETH RPC: %w", err)
	}
	e.rpcClient = rpcClient
	e.ethClient = ethclient.NewClient(rpcClient)

	chainID, err := e.ethClient.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("reconnect failed - can't get chain ID: %w", err)
	}
	if chainID.Cmp(e.chainID) != 0 {
		return fmt.Errorf("chain ID mismatch after reconnect: expected %v, got %v", e.chainID, chainID)
	}

	if err := e.bindContracts(ctx); err != nil {
		return fmt.Errorf("reconnect failed: %w", err)
	}

	e.consecutiveErrors = 0
	e.logf("[RECONNECT] Successfully reconnected to RPC")
	return nil
}

// createTransactOpts builds signing transaction options with a fresh nonce
// and gas price. The configured gas-price multiplier is applied (the legacy
// bot hardcoded 2x here, ignoring GAS_PRICE_MULTIPLIER).
func (e *Engine) createTransactOpts(ctx context.Context, value *big.Int, gasLimit uint64) (*bind.TransactOpts, error) {
	nonce, err := e.ethClient.PendingNonceAt(ctx, e.address)
	if err != nil {
		return nil, fmt.Errorf("getting nonce: %w", err)
	}
	gasPrice, err := e.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting gas price: %w", err)
	}

	txopts := &bind.TransactOpts{
		Context:  ctx,
		From:     e.address,
		Nonce:    new(big.Int).SetUint64(nonce),
		GasPrice: ethtx.AdjustGasPriceBy(gasPrice, e.cfg.GasPriceMultiplier),
		GasLimit: gasLimit,
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.NewEIP155Signer(e.chainID)
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), e.privateKey)
			if err != nil {
				return nil, fmt.Errorf("signing: %w", err)
			}
			return tx.WithSignature(signer, signature)
		},
	}
	if value != nil {
		txopts.Value = new(big.Int).Set(value)
	}
	return txopts, nil
}

// refreshMarket fetches the current market snapshot from the contracts.
func (e *Engine) refreshMarket(ctx context.Context) error {
	var m Market
	var err error
	opts := callOpts(ctx)

	m.CstPrice, err = e.gameContract.GetNextCstBidPrice(opts)
	if err != nil {
		return fmt.Errorf("getting CST bid price: %w", err)
	}
	m.EthBidPrice, err = e.gameContract.GetNextEthBidPrice(opts)
	if err != nil {
		return fmt.Errorf("getting ETH bid price: %w", err)
	}
	m.TimeUntilPrize, err = e.gameContract.GetDurationUntilMainPrize(opts)
	if err != nil {
		return fmt.Errorf("getting time until prize: %w", err)
	}
	tokenAddr, err := e.gameContract.Token(opts)
	if err != nil {
		return fmt.Errorf("getting token address: %w", err)
	}
	tokenContract, err := cgcontracts.NewERC20(tokenAddr, e.ethClient)
	if err != nil {
		return fmt.Errorf("instantiating token contract: %w", err)
	}
	m.CstBalance, err = tokenContract.BalanceOf(opts, e.address)
	if err != nil {
		return fmt.Errorf("getting CST balance: %w", err)
	}
	m.LastBidder, err = e.gameContract.LastBidderAddress(opts)
	if err != nil {
		return fmt.Errorf("getting last bidder: %w", err)
	}
	m.EthBalance, err = e.ethClient.BalanceAt(ctx, e.address, nil)
	if err != nil {
		return fmt.Errorf("getting ETH balance: %w", err)
	}
	m.RWalkMintPrice, err = e.rwalkContract.GetMintPrice(opts)
	if err != nil {
		return fmt.Errorf("getting RWalk mint price: %w", err)
	}

	e.market = m
	return nil
}

// checkTimeoutClaim switches to claiming when the winner missed the claim
// window (then anyone may claim the prize).
func (e *Engine) checkTimeoutClaim(ctx context.Context) bool {
	if isZeroAddress(e.market.LastBidder) {
		return false
	}
	opts := callOpts(ctx)
	timeoutDuration, err := e.gameContract.TimeoutDurationToClaimMainPrize(opts)
	if err != nil {
		return false
	}
	prizeTime, err := e.gameContract.MainPrizeTime(opts)
	if err != nil {
		return false
	}
	var result map[string]any
	if err := e.rpcClient.CallContext(ctx, &result, "eth_getBlockByNumber", "pending", false); err != nil {
		return false
	}
	timestampHex, ok := result["timestamp"].(string)
	if !ok {
		return false
	}
	blockTimestamp, err := hexutil.DecodeBig(timestampHex)
	if err != nil {
		return false
	}
	if timeoutClaimExpired(prizeTime, timeoutDuration, blockTimestamp) {
		e.logf("Winner didn't claim prize during claim window, I am going to claim it")
		return true
	}
	return false
}

// ErrChainReset is returned by Run when the round number decreases, meaning
// the chain the bot was playing on was reset (a dev-net redeploy). The bot
// exits to prevent unintended spending.
var ErrChainReset = errors.New("blockchain was reset")

// checkRoundChange detects the end of the round played. It returns stop=true
// when the engine must terminate: with a nil error when the round ended
// normally, or wrapping ErrChainReset when the round number went backwards.
func (e *Engine) checkRoundChange(ctx context.Context) (stop bool, err error) {
	opts := callOpts(ctx)
	rnum, err := e.gameContract.RoundNum(opts)
	if err != nil {
		e.logf("Error getting roundNum: %v", err)
		return false, nil
	}
	if rnum.Int64() == e.roundNumPlayed {
		return false, nil
	}
	e.logf("Round changed (was %v, now %v)", e.roundNumPlayed, rnum)

	// SAFETY: detect a blockchain reset (round went backwards).
	if rnum.Int64() < e.roundNumPlayed {
		e.logf("ERROR: Round number decreased (%v -> %v) - blockchain was reset!", e.roundNumPlayed, rnum.Int64())
		e.logf("Exiting to prevent unintended spending. Restart bot manually.")
		e.printStats(ctx)
		return true, fmt.Errorf("%w (round %v -> %v)", ErrChainReset, e.roundNumPlayed, rnum.Int64())
	}

	winner, err := e.prizesWallet.MainPrizeBeneficiaryAddresses(opts, big.NewInt(e.roundNumPlayed))
	if err == nil {
		if e.isMyAddress(winner) {
			e.logf("I am the winner of round %v!", e.roundNumPlayed)
			e.stats.roundsWon++
		} else {
			e.logf("I am not the winner of round %v", e.roundNumPlayed)
		}
	}

	if e.cfg.InitialBidPrice != nil {
		e.logf("Playing new round with initial bids")
		e.roundNumPlayed = rnum.Int64()
		e.startInitialBidding = true
		return false, nil
	}
	e.logf("Round ended, exiting...")
	e.printStats(ctx)
	return true, nil
}

// findRWalkTokenID searches for an unused RWalk token owned by the bot. It
// runs in a background goroutine; at most one search is active at a time.
func (e *Engine) findRWalkTokenID(ctx context.Context) {
	if !e.rwalkSearching.CompareAndSwap(false, true) {
		return
	}
	defer e.rwalkSearching.Store(false)
	opts := callOpts(ctx)

	if next := e.nextRWalkTokenID.Load(); next > -1 {
		wasUsed, err := e.gameContract.UsedRandomWalkNfts(opts, big.NewInt(next))
		if err == nil && wasUsed.Cmp(big.NewInt(1)) == 0 {
			e.logf("Resetting nextRWalkTokenID (%v) - already used", next)
			e.nextRWalkTokenID.Store(-1)
		} else if err == nil {
			return // Already have a valid token.
		}
	}

	targetID := e.prevRWalkTokenID.Load() + 1
	lastTokenID, err := e.rwalkContract.NextTokenId(opts)
	if err != nil {
		if ctx.Err() == nil {
			e.logf("Error calling NextTokenId(): %v", err)
		}
		return
	}

	for targetID < lastTokenID.Int64() {
		if ctx.Err() != nil {
			return
		}
		owner, err := e.rwalkContract.OwnerOf(opts, big.NewInt(targetID))
		if err != nil {
			return
		}
		e.prevRWalkTokenID.Store(targetID)

		if e.isMyAddress(owner) {
			wasUsed, err := e.gameContract.UsedRandomWalkNfts(opts, big.NewInt(targetID))
			if err == nil && wasUsed.Sign() == 0 {
				e.nextRWalkTokenID.Store(targetID)
				e.logf("Found RWalk token %v for bidding", targetID)
				return
			}
		}
		targetID++
	}
}

// printConfig logs the operator configuration at startup.
func (e *Engine) printConfig() {
	e.outMu.Lock()
	defer e.outMu.Unlock()
	fmt.Fprintln(e.cfg.Out, "Config params:")
	fmt.Fprintf(e.cfg.Out, "  MAX_ETH_BID: %v ETH\n", fmtEth(e.cfg.Limits.MaxEthBid))
	fmt.Fprintf(e.cfg.Out, "  MAX_CST_BID: %v CST\n", fmtEth(e.cfg.Limits.MaxCstBid))
	fmt.Fprintf(e.cfg.Out, "  RWALK_MIN_PRICE: %v ETH\n", fmtEth(e.cfg.Limits.RWalkMinPrice))
	fmt.Fprintf(e.cfg.Out, "  TIME_BEFORE_PRIZE: %v secs\n", e.cfg.Limits.TimeBeforePrize)
	fmt.Fprintf(e.cfg.Out, "  CST_BID_ANYWAY: %v\n", e.cfg.Limits.CstBidAnyway)
	if e.cfg.InitialBidPrice != nil {
		fmt.Fprintf(e.cfg.Out, "  AT_STARTUP_BID_UP_TO_PRICE_LEVEL: %v ETH\n", fmtEth(e.cfg.InitialBidPrice))
	}
}

// printStats logs the session statistics summary.
func (e *Engine) printStats(ctx context.Context) {
	duration := time.Since(e.stats.startTime)

	currentBalance := new(big.Int).Set(e.stats.startBalance)
	if e.market.EthBalance != nil {
		currentBalance.Set(e.market.EthBalance)
	}
	if ctx.Err() == nil {
		statsCtx, cancel := context.WithTimeout(ctx, statsRPCTimeout)
		defer cancel()
		if balance, err := e.ethClient.BalanceAt(statsCtx, e.address, nil); err == nil {
			currentBalance.Set(balance)
		}
	}
	balanceChange := new(big.Int).Sub(currentBalance, e.stats.startBalance)

	var b bytes.Buffer
	fmt.Fprintln(&b, "\n===========================================================")
	fmt.Fprintln(&b, "                    SESSION SUMMARY")
	fmt.Fprintln(&b, "==========================================================")
	fmt.Fprintf(&b, "  Duration:          %v\n", duration.Round(time.Second))
	fmt.Fprintf(&b, "  Starting Balance:  %v ETH\n", fmtEth(e.stats.startBalance))
	fmt.Fprintf(&b, "  Current Balance:   %v ETH\n", fmtEth(currentBalance))
	if balanceChange.Sign() >= 0 {
		fmt.Fprintf(&b, "  Balance Change:    +%v ETH\n", fmtEth(balanceChange))
	} else {
		fmt.Fprintf(&b, "  Balance Change:    %v ETH\n", fmtEth(balanceChange))
	}
	fmt.Fprintln(&b, "----------------------------------------------------------")
	fmt.Fprintf(&b, "  ETH Bids:          %d\n", e.stats.ethBidCount)
	fmt.Fprintf(&b, "  CST Bids:          %d\n", e.stats.cstBidCount)
	fmt.Fprintf(&b, "  RWalk Bids:        %d\n", e.stats.rwalkBidCount)
	fmt.Fprintf(&b, "  Total Bids:        %d\n", e.stats.ethBidCount+e.stats.cstBidCount+e.stats.rwalkBidCount)
	fmt.Fprintf(&b, "  ETH Spent on Bids: %v ETH\n", fmtEth(e.stats.totalEthSpent))
	fmt.Fprintln(&b, "----------------------------------------------------------")
	fmt.Fprintf(&b, "  Rounds Won:        %d\n", e.stats.roundsWon)
	fmt.Fprintf(&b, "  Prizes Claimed:    %d\n", e.stats.prizesClaimed)
	fmt.Fprintf(&b, "  Failed Txs:        %d\n", e.stats.failedTxCount)
	fmt.Fprintf(&b, "  Reconnections:     %d\n", e.stats.reconnectCount)
	fmt.Fprintln(&b, "==========================================================")
	e.outMu.Lock()
	defer e.outMu.Unlock()
	_, _ = e.cfg.Out.Write(b.Bytes())
}
