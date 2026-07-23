// Package cosmicgame includes a startup contract-parameter drift audit for
// evented live configuration. Every value inspected here
// emits a Changed event, so the indexed cg_adm_* history is authoritative.
// The audit compares live, block-pinned reads with the latest indexed values
// and reports drift; it never creates evt_log, transaction, or correction
// rows. Event-less state belongs in cg_live_state_updates instead.
package cosmicgame

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethcall"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const (
	contractMechanicsUnknown int64 = 0
	contractMechanicsV1      int64 = 1
	contractMechanicsV2      int64 = 2
	contractMechanicsV3      int64 = 3
)

type contractParamSync struct {
	name   string
	table  string
	column string
	read   func(
		v1 *cgc.CosmicSignatureGame,
		v2 *cgc.CosmicSignatureGameV2,
		v3 *cgc.CosmicSignatureGameV3,
		opts *bind.CallOpts,
	) (string, error)
}

// CheckContractParamsDrift compares evented live configuration with indexed
// history at one chain head. The returned count is informational; read or DB
// errors are returned, while individual unsupported getters are logged and
// skipped.
func CheckContractParamsDrift(
	ctx context.Context,
	repo *cgdb.Repo,
	client *ethclient.Client,
	gameAddr, prizesWalletAddr string,
	logger *slog.Logger,
) (int, error) {
	if repo == nil {
		return 0, errors.New("contract drift audit: repo is nil")
	}
	if client == nil {
		return 0, errors.New("contract drift audit: eth client is nil")
	}
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}

	headerCtx, cancel := context.WithTimeout(ctx, ethcall.DefaultTimeout)
	defer cancel()
	header, err := client.HeaderByNumber(headerCtx, nil)
	if err != nil {
		return 0, fmt.Errorf("contract drift audit: latest header: %w", err)
	}

	backend := ethcall.NewBoundedReadBackend(client, ethcall.DefaultTimeout)
	game := ethcommon.HexToAddress(gameAddr)
	v1, _ := cgc.NewCosmicSignatureGame(game, backend)
	v2, _ := cgc.NewCosmicSignatureGameV2(game, backend)
	v3, _ := cgc.NewCosmicSignatureGameV3(game, backend)
	opts := &bind.CallOpts{Context: ctx, BlockNumber: new(big.Int).Set(header.Number)}
	mechanics := probeContractMechanics(v1, v2, v3, opts)

	drifted := 0
	compare := func(name, table, column, chainValue string) error {
		dbValue, hasRow, err := repo.LatestDecimalParam(ctx, table, column)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		if !hasRow {
			logger.Info("contract drift audit: no indexed history",
				"parameter", name,
				"chain_value", chainValue)
			return nil
		}
		if !cgdb.DecimalStringsEqual(dbValue, chainValue) {
			drifted++
			logger.Error("contract parameter drift",
				"parameter", name,
				"db_value", dbValue,
				"chain_value", chainValue,
				"block", header.Number.String())
		}
		return nil
	}

	for _, parameter := range buildContractParamSyncList(mechanics) {
		chainValue, err := parameter.read(v1, v2, v3, opts)
		if err != nil {
			logger.Warn("contract drift audit: parameter skipped",
				"parameter", parameter.name, "err", err)
			continue
		}
		if err := compare(parameter.name, parameter.table, parameter.column, chainValue); err != nil {
			return drifted, err
		}
		if parameter.name == "cst_reward_for_bidding" {
			globReward, err := repo.GlobStatsCstRewardForBidding(ctx)
			if err != nil {
				return drifted, fmt.Errorf("cst_reward_for_bidding (cg_glob_stats): %w", err)
			}
			if !cgdb.DecimalStringsEqual(globReward, chainValue) {
				drifted++
				logger.Error("contract parameter drift",
					"parameter", "cg_glob_stats.cst_reward_for_bidding",
					"db_value", globReward,
					"chain_value", chainValue,
					"block", header.Number.String())
			}
		}
	}

	if value, err := readDelayDuration(v1, v2, v3, opts); err != nil {
		logger.Warn("contract drift audit: parameter skipped",
			"parameter", "delay_before_round_activation", "err", err)
	} else if err := compare(
		"delay_before_round_activation",
		"cg_delay_duration",
		"new_value",
		value.String(),
	); err != nil {
		return drifted, err
	}

	if prizesWalletAddr != "" {
		wallet, _ := cgc.NewPrizesWallet(ethcommon.HexToAddress(prizesWalletAddr), backend)
		value, callErr := wallet.TimeoutDurationToWithdrawPrizes(opts)
		if callErr != nil {
			logger.Warn("contract drift audit: parameter skipped",
				"parameter", "timeout_withdraw_prizes", "err", callErr)
		} else if err := compare(
			"timeout_withdraw_prizes",
			"cg_adm_timeout_withdraw",
			"new_timeout",
			value.String(),
		); err != nil {
			return drifted, err
		}
	}

	if mechanics >= contractMechanicsV2 {
		if value, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err != nil {
			logger.Warn("contract drift audit: parameter skipped",
				"parameter", "cst_dutch_auction_duration_change_divisor", "err", err)
		} else if err := compare(
			"cst_dutch_auction_duration_change_divisor",
			"cg_adm_cst_auclen_chg_div",
			"new_len",
			value.String(),
		); err != nil {
			return drifted, err
		}
	}

	logger.Info("contract drift audit complete",
		"drifted", drifted,
		"mechanics", mechanics,
		"block", header.Number.String())
	return drifted, nil
}

func probeContractMechanics(
	v1 *cgc.CosmicSignatureGame,
	v2 *cgc.CosmicSignatureGameV2,
	v3 *cgc.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) int64 {
	if v3 != nil {
		if _, err := v3.MainPrizeNumCosmicSignatureNfts(opts); err == nil {
			return contractMechanicsV3
		}
	}
	if v2 != nil {
		if _, err := v2.CstDutchAuctionDurationChangeDivisor(opts); err == nil {
			return contractMechanicsV2
		}
	}
	if v1 != nil {
		if _, err := v1.CstDutchAuctionDurationDivisor(opts); err == nil {
			return contractMechanicsV1
		}
	}
	return contractMechanicsUnknown
}

func readDelayDuration(
	v1 *cgc.CosmicSignatureGame,
	v2 *cgc.CosmicSignatureGameV2,
	v3 *cgc.CosmicSignatureGameV3,
	opts *bind.CallOpts,
) (*big.Int, error) {
	if v3 != nil {
		if value, err := v3.DelayDurationBeforeRoundActivation(opts); err == nil {
			return value, nil
		}
	}
	if v2 != nil {
		if value, err := v2.DelayDurationBeforeRoundActivation(opts); err == nil {
			return value, nil
		}
	}
	if v1 != nil {
		if value, err := v1.DelayDurationBeforeRoundActivation(opts); err == nil {
			return value, nil
		}
	}
	return nil, errors.New("cannot read delayDurationBeforeRoundActivation")
}

func readCstReward(
	v1 *cgc.CosmicSignatureGame,
	v2 *cgc.CosmicSignatureGameV2,
	v3 *cgc.CosmicSignatureGameV3,
	opts *bind.CallOpts,
	mechanics int64,
) (string, error) {
	if mechanics == contractMechanicsV3 && v3 != nil {
		value, err := v3.BidCstRewardAmountMultiplier(opts)
		if err == nil {
			return value.String(), nil
		}
	}
	if mechanics == contractMechanicsV2 && v2 != nil {
		value, err := v2.BidCstRewardAmountMultiplier(opts)
		if err == nil {
			return value.String(), nil
		}
	}
	if v1 != nil {
		value, err := v1.CstRewardAmountForBidding(opts)
		if err == nil {
			return value.String(), nil
		}
	}
	return "", errors.New("cannot read CST bid reward configuration")
}

func buildContractParamSyncList(mechanics int64) []contractParamSync {
	v1Big := func(
		read func(*cgc.CosmicSignatureGame, *bind.CallOpts) (*big.Int, error),
	) func(
		*cgc.CosmicSignatureGame,
		*cgc.CosmicSignatureGameV2,
		*cgc.CosmicSignatureGameV3,
		*bind.CallOpts,
	) (string, error) {
		return func(
			v1 *cgc.CosmicSignatureGame,
			_ *cgc.CosmicSignatureGameV2,
			_ *cgc.CosmicSignatureGameV3,
			opts *bind.CallOpts,
		) (string, error) {
			if v1 == nil {
				return "", errors.New("V1-compatible binding unavailable")
			}
			value, err := read(v1, opts)
			if err != nil {
				return "", err
			}
			return value.String(), nil
		}
	}

	parameters := []contractParamSync{
		{
			name: "cst_reward_for_bidding", table: "cg_adm_erc20_reward", column: "new_reward",
			read: func(
				v1 *cgc.CosmicSignatureGame,
				v2 *cgc.CosmicSignatureGameV2,
				v3 *cgc.CosmicSignatureGameV3,
				opts *bind.CallOpts,
			) (string, error) {
				return readCstReward(v1, v2, v3, opts, mechanics)
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

	switch mechanics {
	case contractMechanicsV1:
		parameters = append(parameters, contractParamSync{
			name:  "cst_dutch_auction_duration_divisor",
			table: "cg_adm_cst_auclen", column: "new_len",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.CstDutchAuctionDurationDivisor(opts)
			}),
		})
	case contractMechanicsV2:
		parameters = append(parameters, contractParamSync{
			name:  "cst_dutch_auction_duration",
			table: "cg_adm_cst_auclen", column: "new_len",
			read: func(
				_ *cgc.CosmicSignatureGame,
				v2 *cgc.CosmicSignatureGameV2,
				_ *cgc.CosmicSignatureGameV3,
				opts *bind.CallOpts,
			) (string, error) {
				value, err := v2.CstDutchAuctionDuration(opts)
				if err != nil {
					return "", err
				}
				return value.String(), nil
			},
		})
	case contractMechanicsV3:
		parameters = append(parameters,
			v3ContractParam(
				"round_late_bid_duration_divisor",
				"cg_adm_late_bid_dur_divisor",
				func(v3 *cgc.CosmicSignatureGameV3, opts *bind.CallOpts) (*big.Int, error) {
					return v3.RoundLateBidDurationDivisor(opts)
				},
			),
			v3ContractParam(
				"round_late_bid_price_premium_base_multiplier",
				"cg_adm_late_bid_premium_base_mul",
				func(v3 *cgc.CosmicSignatureGameV3, opts *bind.CallOpts) (*big.Int, error) {
					return v3.RoundLateBidPricePremiumAmountBaseMultiplier(opts)
				},
			),
			v3ContractParam(
				"round_late_bid_price_premium_exponent",
				"cg_adm_late_bid_premium_exponent",
				func(v3 *cgc.CosmicSignatureGameV3, opts *bind.CallOpts) (*big.Int, error) {
					return v3.RoundLateBidPricePremiumAmountExponent(opts)
				},
			),
			v3ContractParam(
				"last_bidder_bid_cst_reward_amount_percentage",
				"cg_adm_last_bidder_reward_pct",
				func(v3 *cgc.CosmicSignatureGameV3, opts *bind.CallOpts) (*big.Int, error) {
					return v3.LastBidderBidCstRewardAmountPercentage(opts)
				},
			),
			v3ContractParam(
				"main_prize_num_cosmic_signature_nfts",
				"cg_adm_main_prize_num_nfts",
				func(v3 *cgc.CosmicSignatureGameV3, opts *bind.CallOpts) (*big.Int, error) {
					return v3.MainPrizeNumCosmicSignatureNfts(opts)
				},
			),
		)
	}
	return parameters
}

func v3ContractParam(
	name, table string,
	read func(*cgc.CosmicSignatureGameV3, *bind.CallOpts) (*big.Int, error),
) contractParamSync {
	return contractParamSync{
		name: name, table: table, column: "new_value",
		read: func(
			_ *cgc.CosmicSignatureGame,
			_ *cgc.CosmicSignatureGameV2,
			v3 *cgc.CosmicSignatureGameV3,
			opts *bind.CallOpts,
		) (string, error) {
			if v3 == nil {
				return "", errors.New("V3 binding unavailable")
			}
			value, err := read(v3, opts)
			if err != nil {
				return "", err
			}
			return value.String(), nil
		},
	}
}
