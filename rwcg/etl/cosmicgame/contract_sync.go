package main

// Startup drift check for contract parameters.
//
// Every parameter listed here emits a *Changed event, and the ETL decodes all of them,
// so the DB history tables are the source of truth. This check reads the live values
// via eth_call and compares them against the latest DB values: a mismatch means an
// event was missed or not decoded — a bug to investigate — and is logged as an error.
// NOTHING is written to the database. In particular, no synthetic evt_log rows, no
// sentinel transactions and no correction rows are created (that legacy mechanism was
// removed; see cg_live_state_updates for state variables that have no events).
//
// A parameter with no DB history at all is reported as info: it simply has never been
// changed on-chain (its value was set in the constructor/initializer without an event),
// so there is nothing to compare against.

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cgdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/cosmicgame"
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
	read         func(v1 *cosmicgame.CosmicSignatureGame, v2 *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error)
}

// checkContractParamsDrift compares live chain values against the latest DB history
// values and returns the number of drifted parameters. Read-only: logs findings and
// writes nothing.
func checkContractParamsDrift(sw *cgdb.SQLStorageWrapper, client *ethclient.Client, gameAddr, prizesWalletAddr string, info, errLog *log.Logger) (int, error) {
	if client == nil {
		return 0, fmt.Errorf("eth client is nil")
	}
	game := ethcommon.HexToAddress(gameAddr)
	v1, _ := cosmicgame.NewCosmicSignatureGame(game, client)
	v2, _ := cosmicgame.NewCosmicSignatureGameV2(game, client)
	if v1 == nil && v2 == nil {
		return 0, fmt.Errorf("can't bind CosmicGame at %s", gameAddr)
	}

	var copts bind.CallOpts
	mechanics := probeContractMechanics(v1, v2, &copts)

	var drifted int
	compare := func(name, table, column, chainValue string) error {
		dbValue, hasRow, err := sw.Get_latest_decimal_param(table, column)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		if !hasRow {
			info.Printf("drift check: %s has no DB history (never changed on-chain); chain value: %s", name, chainValue)
			return nil
		}
		if !cgdb.DecimalStringsEqual(dbValue, chainValue) {
			drifted++
			errLog.Printf("drift check: DRIFT %s: db=%s chain=%s (missed or undecoded event — investigate)", name, dbValue, chainValue)
			info.Printf("drift check: DRIFT %s: db=%s chain=%s", name, dbValue, chainValue)
		}
		return nil
	}

	for _, p := range buildContractParamSyncList(mechanics) {
		chainValue, err := p.read(v1, v2, &copts)
		if err != nil {
			errLog.Printf("drift check: skip %s: %v", p.name, err)
			continue
		}
		if err := compare(p.name, p.table, p.column, chainValue); err != nil {
			return drifted, err
		}
		if p.name == "cst_reward_for_bidding" {
			// cg_glob_stats holds a trigger-maintained copy; verify it as well.
			globReward, err := sw.Get_glob_stats_cst_reward_for_bidding()
			if err != nil {
				return drifted, fmt.Errorf("cst_reward_for_bidding (cg_glob_stats): %w", err)
			}
			if !cgdb.DecimalStringsEqual(globReward, chainValue) {
				drifted++
				errLog.Printf("drift check: DRIFT cg_glob_stats.cst_reward_for_bidding: db=%s chain=%s", globReward, chainValue)
			}
		}
	}

	if delay, err := readDelayDuration(v1, v2, &copts); err == nil {
		if err := compare("delay_before_round_activation", "cg_delay_duration", "new_value", fmt.Sprintf("%d", delay)); err != nil {
			return drifted, err
		}
	} else {
		errLog.Printf("drift check: skip delay_before_round_activation: %v", err)
	}

	if prizesWalletAddr != "" {
		pw, err := cosmicgame.NewPrizesWallet(ethcommon.HexToAddress(prizesWalletAddr), client)
		if err != nil {
			errLog.Printf("drift check: skip timeout_withdraw_prizes: bind PrizesWallet: %v", err)
		} else if val, err := pw.TimeoutDurationToWithdrawPrizes(&copts); err != nil {
			errLog.Printf("drift check: skip timeout_withdraw_prizes: %v", err)
		} else if err := compare("timeout_withdraw_prizes", "cg_adm_timeout_withdraw", "new_timeout", val.String()); err != nil {
			return drifted, err
		}
	}

	if cstAucChangeDiv := readCSTAuctionDurationChangeDivisor(v2, &copts, mechanics); cstAucChangeDiv >= 0 {
		if err := compare("cst_dutch_auction_duration_change_divisor", "cg_adm_cst_auclen_chg_div", "new_len", fmt.Sprintf("%d", cstAucChangeDiv)); err != nil {
			return drifted, err
		}
	}

	if drifted == 0 {
		info.Printf("drift check complete: all parameters match chain (mechanics v%d)", mechanics)
	} else {
		info.Printf("drift check complete: %d parameter(s) DRIFTED from chain (mechanics v%d) — see error log", drifted, mechanics)
	}
	return drifted, nil
}

func probeContractMechanics(v1 *cosmicgame.CosmicSignatureGame, v2 *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts) int64 {
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

func readDelayDuration(v1 *cosmicgame.CosmicSignatureGame, v2 *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts) (int64, error) {
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
	return 0, fmt.Errorf("can't read delayDurationBeforeRoundActivation")
}

func readCstReward(v1 *cosmicgame.CosmicSignatureGame, v2 *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts, mechanics int64) (string, error) {
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
	return "", fmt.Errorf("can't read CST bid reward")
}

func readRoundStartCSTAuction(v1 *cosmicgame.CosmicSignatureGame, v2 *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts, mechanics int64) (string, error) {
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
	return "", fmt.Errorf("can't read CST dutch auction duration")
}

func readCSTAuctionDurationChangeDivisor(v2 *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts, mechanics int64) int64 {
	if mechanics == contractMechanicsV1 || v2 == nil {
		return -1
	}
	if val, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
		return val.Int64()
	}
	return -1
}

func buildContractParamSyncList(mechanics int64) []contractParamSync {
	v1Big := func(read func(*cosmicgame.CosmicSignatureGame, *bind.CallOpts) (*big.Int, error)) func(*cosmicgame.CosmicSignatureGame, *cosmicgame.CosmicSignatureGameV2, *bind.CallOpts) (string, error) {
		return func(v1 *cosmicgame.CosmicSignatureGame, _ *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error) {
			if v1 == nil {
				return "", fmt.Errorf("v1 binding unavailable")
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
			read: func(v1 *cosmicgame.CosmicSignatureGame, v2 *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error) {
				return readCstReward(v1, v2, opts, mechanics)
			},
		},
		{
			name: "cst_dutch_auction_duration", table: "cg_adm_cst_auclen", column: "new_len",
			read: func(v1 *cosmicgame.CosmicSignatureGame, v2 *cosmicgame.CosmicSignatureGameV2, opts *bind.CallOpts) (string, error) {
				return readRoundStartCSTAuction(v1, v2, opts, mechanics)
			},
		},
		{
			name: "timeout_claim_main_prize", table: "cg_adm_timeout_claimprize", column: "new_timeout",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.TimeoutDurationToClaimMainPrize(opts)
			}),
		},
		{
			name: "eth_bid_price_increase_divisor", table: "cg_adm_price_inc", column: "new_price_increase",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.EthBidPriceIncreaseDivisor(opts)
			}),
		},
		{
			name: "main_prize_time_increment_divisor", table: "cg_adm_time_inc", column: "new_time_inc",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.MainPrizeTimeIncrementIncreaseDivisor(opts)
			}),
		},
		{
			name: "main_prize_microseconds_increment", table: "cg_adm_prize_microsec", column: "new_microseconds",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.MainPrizeTimeIncrementInMicroSeconds(opts)
			}),
		},
		{
			name: "initial_duration_until_main_prize_divisor", table: "cg_adm_inisecprize", column: "new_inisec",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.InitialDurationUntilMainPrizeDivisor(opts)
			}),
		},
		{
			name: "eth_dutch_auction_duration_divisor", table: "cg_adm_eth_auclen", column: "new_len",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.EthDutchAuctionDurationDivisor(opts)
			}),
		},
		{
			name: "eth_dutch_auction_ending_bid_price_divisor", table: "cg_adm_eth_auc_endprice", column: "new_len",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.EthDutchAuctionEndingBidPriceDivisor(opts)
			}),
		},
		{
			name: "cst_dutch_auction_beginning_bid_price_min_limit", table: "cg_adm_cst_min_limit", column: "min_limit",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.CstDutchAuctionBeginningBidPriceMinLimit(opts)
			}),
		},
		{
			name: "marketing_wallet_cst_contribution", table: "cg_adm_mkt_reward", column: "new_reward",
			read: v1Big(func(v1 *cosmicgame.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.MarketingWalletCstContributionAmount(opts)
			}),
		},
	}
}
