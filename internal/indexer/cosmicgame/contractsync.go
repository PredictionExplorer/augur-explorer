// The startup contract-parameter sync: reads live monetary/timing settings
// from the chain and inserts admin-correction rows when the stored values
// disagree (check-then-correct; a clean run writes nothing).

package cosmicgame

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethcall"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const (
	chainSyncTxHash             = "0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
	chainSyncLogIndexFloor uint = 990000
)

const (
	contractMechanicsUnknown int64 = 0
	contractMechanicsV1      int64 = 1
	contractMechanicsV2      int64 = 2
)

type contractParamSync struct {
	name         string
	table        string
	column       string
	contractAddr string
	read         func(v1 *cgc.CosmicSignatureGame, v2 *cgc.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error)
}

func allocChainSyncEvtlog(ctx context.Context, st *store.Store, contractAddr string, client *ethclient.Client) (*cgdb.AdminCorrectionMeta, error) {
	headerCtx, cancel := context.WithTimeout(ctx, ethcall.DefaultTimeout)
	defer cancel()
	header, err := client.HeaderByNumber(headerCtx, nil)
	if err != nil {
		return nil, fmt.Errorf("HeaderByNumber: %w", err)
	}
	if header.Time > math.MaxInt64 {
		return nil, fmt.Errorf("header timestamp %d overflows int64", header.Time)
	}
	if err := st.InsertBlock(ctx, header); err != nil {
		return nil, fmt.Errorf("InsertBlock: %w", err)
	}

	blockNum := header.Number.Int64()
	txID, err := st.InsertMinimalTransaction(ctx, chainSyncTxHash, blockNum)
	if err != nil {
		return nil, fmt.Errorf("InsertMinimalTransaction: %w", err)
	}
	logIndex, err := st.NextEventLogIndex(ctx, blockNum, chainSyncLogIndexFloor)
	if err != nil {
		return nil, fmt.Errorf("NextEventLogIndex: %w", err)
	}

	contractAid, err := st.LookupOrCreateAddress(ctx, contractAddr, blockNum, txID)
	if err != nil {
		return nil, fmt.Errorf("LookupOrCreateAddress: %w", err)
	}
	syncLog := types.Log{
		Address:     ethcommon.HexToAddress(contractAddr),
		BlockNumber: header.Number.Uint64(),
		TxHash:      ethcommon.HexToHash(chainSyncTxHash),
		Index:       logIndex,
	}

	evtID, err := st.InsertEventLog(ctx, syncLog, txID, contractAid)
	if err != nil {
		return nil, fmt.Errorf("InsertEventLog: %w", err)
	}

	return &cgdb.AdminCorrectionMeta{
		EvtID:       evtID,
		BlockNum:    blockNum,
		TxID:        txID,
		TimeStamp:   int64(header.Time), // #nosec G115 -- checked above
		ContractAid: contractAid,
	}, nil
}

// paramSyncer applies check-then-correct sync policy on top of the Repo
// primitives: it compares the latest stored admin value with the chain value
// and inserts a correction row only on mismatch. The contract-address
// resolution stays lazy (only on an actual insert) so a clean sync run
// leaves the address table untouched, exactly like the legacy layer.
type paramSyncer struct {
	ctx   context.Context
	repo  *cgdb.Repo
	store *store.Store
	meta  *cgdb.AdminCorrectionMeta
}

func (s *paramSyncer) contractAid(contractAddr string) (int64, error) {
	if contractAddr == "" {
		return 0, nil // Repo substitutes meta.ContractAid
	}
	aid, err := s.store.LookupOrCreateAddress(s.ctx, contractAddr, s.meta.BlockNum, s.meta.TxID)
	if err != nil {
		return 0, fmt.Errorf("correction contract address %v: %w", contractAddr, err)
	}
	return aid, nil
}

// syncDecimal aligns one admin/history column with wantValue, attributing
// the correction row to contractAddr when non-empty.
func (s *paramSyncer) syncDecimal(name, table, column, wantValue, contractAddr string) (bool, error) {
	dbValue, hasRow, err := s.repo.LatestDecimalParam(s.ctx, table, column)
	if err != nil {
		return false, err
	}
	if hasRow && cgdb.DecimalStringsEqual(dbValue, wantValue) {
		return false, nil
	}
	aid, err := s.contractAid(contractAddr)
	if err != nil {
		return false, err
	}
	if err := s.repo.InsertAdminCorrectionDecimal(s.ctx, table, column, wantValue, s.meta, aid); err != nil {
		return false, fmt.Errorf("%s: %w", name, err)
	}
	return true, nil
}

// syncInt64 is syncDecimal for integer admin params stored as DECIMAL/BIGINT columns.
func (s *paramSyncer) syncInt64(name, table, column string, wantValue int64, contractAddr string) (bool, error) {
	dbStr, hasRow, err := s.repo.LatestDecimalParam(s.ctx, table, column)
	if err != nil {
		return false, err
	}
	if hasRow {
		dbInt, ok := new(big.Int).SetString(dbStr, 10)
		if ok && dbInt.Int64() == wantValue {
			return false, nil
		}
	}
	aid, err := s.contractAid(contractAddr)
	if err != nil {
		return false, err
	}
	valStr := strconv.FormatInt(wantValue, 10)
	if err := s.repo.InsertAdminCorrectionDecimal(s.ctx, table, column, valStr, s.meta, aid); err != nil {
		return false, fmt.Errorf("%s: %w", name, err)
	}
	return true, nil
}

// syncCstReward aligns cg_adm_erc20_reward and cg_glob_stats with wantValue.
func (s *paramSyncer) syncCstReward(wantValue string) (bool, error) {
	globReward, err := s.repo.GlobStatsCstRewardForBidding(s.ctx)
	if err != nil {
		return false, err
	}
	dbLatest, hasRow, err := s.repo.LatestDecimalParam(s.ctx, "cg_adm_erc20_reward", "new_reward")
	if err != nil {
		return false, err
	}
	if hasRow && cgdb.DecimalStringsEqual(dbLatest, wantValue) && cgdb.DecimalStringsEqual(globReward, wantValue) {
		return false, nil
	}
	if err := s.repo.InsertAdminCorrectionERC20Reward(s.ctx, wantValue, s.meta); err != nil {
		return false, fmt.Errorf("cst_reward_for_bidding: %w", err)
	}
	return true, nil
}

// SyncContractParams reads live monetary/timed settings from RPC and inserts SQL
// correction rows when the latest admin/history value differs from chain.
// cmd/cg-etl runs it once at startup, before the polling loop.
func SyncContractParams(ctx context.Context, repo *cgdb.Repo, st *store.Store, client *ethclient.Client, gameAddr, prizesWalletAddr string, logger *slog.Logger) error {
	if client == nil {
		return errors.New("eth client is nil")
	}
	// Every chain read is bounded: the bindings run over a read-bounded
	// backend (per-call deadline) and the CallOpts carry the sync context,
	// so a wedged RPC node fails the startup sync loudly instead of
	// wedging cg-etl before its polling loop ever starts.
	backend := ethcall.NewBoundedReadBackend(client, ethcall.DefaultTimeout)
	game := ethcommon.HexToAddress(gameAddr)
	v1, _ := cgc.NewCosmicSignatureGame(game, backend)
	v2, _ := cgc.NewCosmicSignatureGameV2(game, backend)
	if v1 == nil && v2 == nil {
		return fmt.Errorf("can't bind CosmicGame at %s", gameAddr)
	}

	copts := bind.CallOpts{Context: ctx}
	mechanics := probeContractMechanics(v1, v2, &copts)

	meta, err := allocChainSyncEvtlog(ctx, st, gameAddr, client)
	if err != nil {
		return fmt.Errorf("alloc chain sync evtlog: %w", err)
	}
	syncer := &paramSyncer{ctx: ctx, repo: repo, store: st, meta: meta}

	var updated int
	for _, p := range buildContractParamSyncList(mechanics) {
		chainValue, err := p.read(v1, v2, &copts)
		if err != nil {
			logger.Warn("chain sync: parameter skipped", "param", p.name, "err", err)
			continue
		}
		var changed bool
		if p.name == "cst_reward_for_bidding" {
			changed, err = syncer.syncCstReward(chainValue)
		} else {
			changed, err = syncer.syncDecimal(p.name, p.table, p.column, chainValue, p.contractAddr)
		}
		if err != nil {
			return err
		}
		if changed {
			updated++
			logger.Info("chain sync: parameter corrected", "param", p.name, "value", chainValue)
		}
	}

	if delay, err := readDelayDuration(v1, v2, &copts); err == nil {
		changed, err := syncer.syncInt64("delay_before_round_activation", "cg_delay_duration", "new_value", delay, "")
		if err != nil {
			return err
		}
		if changed {
			updated++
			logger.Info("chain sync: parameter corrected", "param", "delay_before_round_activation", "value", delay)
		}
	} else {
		logger.Warn("chain sync: parameter skipped", "param", "delay_before_round_activation", "err", err)
	}

	if prizesWalletAddr != "" {
		pw, err := cgc.NewPrizesWallet(ethcommon.HexToAddress(prizesWalletAddr), backend)
		if err != nil {
			logger.Warn("chain sync: parameter skipped", "param", "timeout_withdraw_prizes", "err", fmt.Errorf("bind PrizesWallet: %w", err))
		} else if val, err := pw.TimeoutDurationToWithdrawPrizes(&copts); err != nil {
			logger.Warn("chain sync: parameter skipped", "param", "timeout_withdraw_prizes", "err", err)
		} else {
			changed, err := syncer.syncInt64(
				"timeout_withdraw_prizes", "cg_adm_timeout_withdraw", "new_timeout",
				val.Int64(), prizesWalletAddr,
			)
			if err != nil {
				return err
			}
			if changed {
				updated++
				logger.Info("chain sync: parameter corrected", "param", "timeout_withdraw_prizes", "value", val.String())
			}
		}
	}

	if cstAucChangeDiv := readCSTAuctionDurationChangeDivisor(v2, &copts, mechanics); cstAucChangeDiv >= 0 {
		valStr := strconv.FormatInt(cstAucChangeDiv, 10)
		changed, err := syncer.syncDecimal("cst_dutch_auction_duration_change_divisor", "cg_adm_cst_auclen_chg_div", "new_len", valStr, "")
		if err != nil {
			return err
		}
		if changed {
			updated++
			logger.Info("chain sync: parameter corrected", "param", "cst_dutch_auction_duration_change_divisor", "value", valStr)
		}
	}

	logger.Info("chain sync complete", "corrected", updated, "mechanics", mechanics)
	return nil
}

func probeContractMechanics(v1 *cgc.CosmicSignatureGame, v2 *cgc.CosmicSignatureGameV2, opts *bind.CallOpts) int64 {
	if v2 != nil {
		if _, err := v2.CstDutchAuctionDuration(opts); err == nil {
			return contractMechanicsV2
		}
	}
	if v1 != nil {
		if _, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			return contractMechanicsV1
		}
	}
	if v2 != nil {
		if _, err := v2.BidCstRewardAmountMultiplier(opts); err == nil {
			return contractMechanicsV2
		}
	}
	return contractMechanicsUnknown
}

func readDelayDuration(v1 *cgc.CosmicSignatureGame, v2 *cgc.CosmicSignatureGameV2, opts *bind.CallOpts) (int64, error) {
	if v1 != nil {
		if val, err := v1.DelayDurationBeforeRoundActivation(opts); err == nil {
			return val.Int64(), nil
		}
	}
	if v2 != nil {
		if val, err := v2.DelayDurationBeforeRoundActivation(opts); err == nil {
			return val.Int64(), nil
		}
	}
	return 0, errors.New("can't read delayDurationBeforeRoundActivation")
}

func readCstReward(v1 *cgc.CosmicSignatureGame, v2 *cgc.CosmicSignatureGameV2, opts *bind.CallOpts, mechanics int64) (string, error) {
	if mechanics == contractMechanicsV2 && v2 != nil {
		if val, err := v2.BidCstRewardAmountMultiplier(opts); err == nil {
			return val.String(), nil
		}
	}
	if v1 != nil {
		if val, err := v1.CstRewardAmountForBidding(opts); err == nil {
			return val.String(), nil
		}
	}
	if v2 != nil {
		if val, err := v2.BidCstRewardAmountMultiplier(opts); err == nil {
			return val.String(), nil
		}
	}
	return "", errors.New("can't read CST bid reward")
}

func readRoundStartCSTAuction(v1 *cgc.CosmicSignatureGame, v2 *cgc.CosmicSignatureGameV2, opts *bind.CallOpts, mechanics int64) (string, error) {
	if mechanics == contractMechanicsV2 && v2 != nil {
		if val, err := v2.CstDutchAuctionDuration(opts); err == nil {
			return val.String(), nil
		}
	}
	if v1 != nil {
		if val, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			return val.String(), nil
		}
	}
	if v2 != nil {
		if val, err := v2.CstDutchAuctionDuration(opts); err == nil {
			return val.String(), nil
		}
	}
	return "", errors.New("can't read CST dutch auction duration")
}

func readCSTAuctionDurationChangeDivisor(v2 *cgc.CosmicSignatureGameV2, opts *bind.CallOpts, mechanics int64) int64 {
	if mechanics == contractMechanicsV1 || v2 == nil {
		return -1
	}
	if val, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
		return val.Int64()
	}
	return -1
}

func buildContractParamSyncList(mechanics int64) []contractParamSync {
	v1Big := func(read func(*cgc.CosmicSignatureGame, *bind.CallOpts) (*big.Int, error)) func(*cgc.CosmicSignatureGame, *cgc.CosmicSignatureGameV2, *bind.CallOpts) (string, error) {
		return func(v1 *cgc.CosmicSignatureGame, _ *cgc.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error) {
			if v1 == nil {
				return "", errors.New("v1 binding unavailable")
			}
			val, err := read(v1, opts)
			if err != nil {
				return "", err
			}
			return val.String(), nil
		}
	}

	return []contractParamSync{
		{
			name: "cst_reward_for_bidding", table: "cg_adm_erc20_reward", column: "new_reward",
			read: func(v1 *cgc.CosmicSignatureGame, v2 *cgc.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error) {
				return readCstReward(v1, v2, opts, mechanics)
			},
		},
		{
			name: "cst_dutch_auction_duration", table: "cg_adm_cst_auclen", column: "new_len",
			read: func(v1 *cgc.CosmicSignatureGame, v2 *cgc.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error) {
				return readRoundStartCSTAuction(v1, v2, opts, mechanics)
			},
		},
		{
			name: "timeout_claim_main_prize", table: "cg_adm_timeout_claimprize", column: "new_timeout",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.TimeoutDurationToClaimMainPrize(opts)
			}),
		},
		{
			name: "eth_bid_price_increase_divisor", table: "cg_adm_price_inc", column: "new_price_increase",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.EthBidPriceIncreaseDivisor(opts)
			}),
		},
		{
			name: "main_prize_time_increment_divisor", table: "cg_adm_time_inc", column: "new_time_inc",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.MainPrizeTimeIncrementIncreaseDivisor(opts)
			}),
		},
		{
			name: "main_prize_microseconds_increment", table: "cg_adm_prize_microsec", column: "new_microseconds",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.MainPrizeTimeIncrementInMicroSeconds(opts)
			}),
		},
		{
			name: "initial_duration_until_main_prize_divisor", table: "cg_adm_inisecprize", column: "new_inisec",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.InitialDurationUntilMainPrizeDivisor(opts)
			}),
		},
		{
			name: "eth_dutch_auction_duration_divisor", table: "cg_adm_eth_auclen", column: "new_len",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.EthDutchAuctionDurationDivisor(opts)
			}),
		},
		{
			name: "eth_dutch_auction_ending_bid_price_divisor", table: "cg_adm_eth_auc_endprice", column: "new_len",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.EthDutchAuctionEndingBidPriceDivisor(opts)
			}),
		},
		{
			name: "cst_dutch_auction_beginning_bid_price_min_limit", table: "cg_adm_cst_min_limit", column: "min_limit",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.CstDutchAuctionBeginningBidPriceMinLimit(opts)
			}),
		},
		{
			name: "marketing_wallet_cst_contribution", table: "cg_adm_mkt_reward", column: "new_reward",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.MarketingWalletCstContributionAmount(opts)
			}),
		},
	}
}
