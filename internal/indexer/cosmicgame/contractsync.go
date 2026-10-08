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

	// Contract address settings: every *AddressChanged event has a table
	// keyed by address id; compare the live getter with the latest row.
	for _, parameter := range contractAddressParams() {
		chainAddr, err := parameter.read(v1, opts)
		if err != nil {
			logger.Warn("contract drift audit: parameter skipped",
				"parameter", parameter.name, "err", err)
			continue
		}
		dbAddr, hasRow, err := repo.LatestAddressParam(ctx, parameter.table, parameter.aidColumn)
		if err != nil {
			return drifted, fmt.Errorf("%s: %w", parameter.name, err)
		}
		if !hasRow {
			logger.Info("contract drift audit: no indexed history",
				"parameter", parameter.name,
				"chain_value", chainAddr.Hex())
			continue
		}
		if ethcommon.HexToAddress(dbAddr) != chainAddr {
			drifted++
			logger.Error("contract parameter drift",
				"parameter", parameter.name,
				"db_value", dbAddr,
				"chain_value", chainAddr.Hex(),
				"block", header.Number.String())
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
		{
			name: "eth_bid_refund_amount_in_gas_to_swallow_max_limit", table: "cg_adm_eth_bid_refund_gas_limit", column: "new_value",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.EthBidRefundAmountInGasToSwallowMaxLimit(opts)
			}),
		},
		// Prize split percentages.
		{
			name: "charity_eth_donation_amount_percentage", table: "cg_adm_charity_pcent", column: "percentage",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.CharityEthDonationAmountPercentage(opts)
			}),
		},
		{
			name: "main_eth_prize_amount_percentage", table: "cg_adm_main_prize_pcent", column: "percentage",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.MainEthPrizeAmountPercentage(opts)
			}),
		},
		{
			name: "raffle_total_eth_prize_amount_for_bidders_percentage", table: "cg_adm_raffle_pcent", column: "percentage",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.RaffleTotalEthPrizeAmountForBiddersPercentage(opts)
			}),
		},
		{
			name: "chrono_warrior_eth_prize_amount_percentage", table: "cg_adm_chrono_pcent", column: "percentage",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.ChronoWarriorEthPrizeAmountPercentage(opts)
			}),
		},
		{
			name: "cosmic_signature_nft_staking_total_eth_reward_amount_percentage", table: "cg_adm_stake_pcent", column: "percentage",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(opts)
			}),
		},
		// Raffle counts.
		{
			name: "num_raffle_eth_prizes_for_bidders", table: "cg_adm_raf_eth_bidding", column: "num_winners",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.NumRaffleEthPrizesForBidders(opts)
			}),
		},
		{
			name: "num_raffle_cosmic_signature_nfts_for_bidders", table: "cg_adm_raf_nft_bidding", column: "num_winners",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.NumRaffleCosmicSignatureNftsForBidders(opts)
			}),
		},
		{
			name: "num_raffle_cosmic_signature_nfts_for_random_walk_nft_stakers", table: "cg_adm_raf_nft_staking_rwalk", column: "num_winners",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(opts)
			}),
		},
		// Remaining scalar settings.
		{
			name: "bid_message_length_max_limit", table: "cg_adm_msg_len", column: "new_length",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.BidMessageLengthMaxLimit(opts)
			}),
		},
		{
			name: "cst_prize_amount", table: "cg_adm_erc_rwd_mul", column: "new_reward",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.CstPrizeAmount(opts)
			}),
		},
		// roundActivationTime changes on every main-prize claim, each time
		// through RoundActivationTimeChanged, so the latest event must match
		// the live value between claims.
		{
			name: "round_activation_time", table: "cg_adm_acttime", column: "new_atime",
			read: v1Big(func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (*big.Int, error) {
				return v1.RoundActivationTime(opts)
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
				"cst_bid_price_decline_multiplier",
				"cg_adm_cst_price_decline_mul",
				func(v3 *cgc.CosmicSignatureGameV3, opts *bind.CallOpts) (*big.Int, error) {
					return v3.CstBidPriceDeclineMultiplier(opts)
				},
			),
			v3ContractParam(
				"cst_bid_price_decline_multiplier_change_divisor",
				"cg_adm_cst_price_decline_mul_div",
				func(v3 *cgc.CosmicSignatureGameV3, opts *bind.CallOpts) (*big.Int, error) {
					return v3.CstBidPriceDeclineMultiplierChangeDivisor(opts)
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

// contractAddressParam pairs an address getter with the *AddressChanged
// history table that records it (address-id column).
type contractAddressParam struct {
	name      string
	table     string
	aidColumn string
	read      func(v1 *cgc.CosmicSignatureGame, opts *bind.CallOpts) (ethcommon.Address, error)
}

// contractAddressParams lists the game's evented address settings. The
// getters exist unchanged on every generation, so the V1 binding serves all.
func contractAddressParams() []contractAddressParam {
	return []contractAddressParam{
		{
			name: "charity_address", table: "cg_adm_charity_wallet", aidColumn: "new_charity_aid",
			read: (*cgc.CosmicSignatureGame).CharityAddress,
		},
		{
			name: "random_walk_nft_address", table: "cg_adm_rwalk_addr", aidColumn: "new_rwalk_aid",
			read: (*cgc.CosmicSignatureGame).RandomWalkNft,
		},
		{
			name: "cosmic_signature_nft_address", table: "cg_adm_cossig_addr", aidColumn: "new_cossig_aid",
			read: (*cgc.CosmicSignatureGame).Nft,
		},
		{
			name: "cosmic_signature_token_address", table: "cg_adm_costok_addr", aidColumn: "new_costok_aid",
			read: (*cgc.CosmicSignatureGame).Token,
		},
		{
			name: "prizes_wallet_address", table: "cg_adm_prizes_wallet_addr", aidColumn: "new_wallet_aid",
			read: (*cgc.CosmicSignatureGame).PrizesWallet,
		},
		{
			name: "staking_wallet_cosmic_signature_nft_address", table: "cg_adm_staking_cst_addr", aidColumn: "new_staking_aid",
			read: (*cgc.CosmicSignatureGame).StakingWalletCosmicSignatureNft,
		},
		{
			name: "staking_wallet_random_walk_nft_address", table: "cg_adm_staking_rwalk_addr", aidColumn: "new_staking_aid",
			read: (*cgc.CosmicSignatureGame).StakingWalletRandomWalkNft,
		},
		{
			name: "marketing_wallet_address", table: "cg_adm_marketing_addr", aidColumn: "new_marketing_aid",
			read: (*cgc.CosmicSignatureGame).MarketingWallet,
		},
	}
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
