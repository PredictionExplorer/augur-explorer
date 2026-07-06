package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const chainSyncTxHash = "0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"

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

func chainSyncLogIndex() uint {
	return 990000 + uint(time.Now().UnixNano()%10000)
}

func allocChainSyncEvtlog(sw *cgdb.SQLStorageWrapper, contractAddr string, client *ethclient.Client) (*cgdb.AdminCorrectionMeta, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("HeaderByNumber: %w", err)
	}
	if err := sw.S.Insert_block(header); err != nil {
		return nil, fmt.Errorf("Insert_block: %w", err)
	}

	blockNum := header.Number.Int64()
	txId, err := sw.S.Insert_minimal_transaction(chainSyncTxHash, blockNum)
	if err != nil {
		return nil, fmt.Errorf("Insert_minimal_transaction: %w", err)
	}

	contractAid, err := sw.S.Lookup_or_create_address(contractAddr, blockNum, txId)
	if err != nil {
		return nil, fmt.Errorf("Lookup_or_create_address: %w", err)
	}
	syncLog := types.Log{
		Address:     ethcommon.HexToAddress(contractAddr),
		BlockNumber: header.Number.Uint64(),
		TxHash:      ethcommon.HexToHash(chainSyncTxHash),
		Index:       chainSyncLogIndex(),
	}

	evtId, err := sw.S.Insert_event_log(syncLog, txId, contractAid)
	if err != nil {
		return nil, fmt.Errorf("Insert_event_log: %w", err)
	}

	return &cgdb.AdminCorrectionMeta{
		EvtId:       evtId,
		BlockNum:    blockNum,
		TxId:        txId,
		TimeStamp:   int64(header.Time),
		ContractAid: contractAid,
	}, nil
}

// syncContractParamsFromChain reads live monetary/timed settings from RPC and inserts SQL
// correction rows when the latest admin/history value differs from chain.
func syncContractParamsFromChain(sw *cgdb.SQLStorageWrapper, client *ethclient.Client, gameAddr, prizesWalletAddr string, info, errLog *log.Logger) error {
	if client == nil {
		return fmt.Errorf("eth client is nil")
	}
	game := ethcommon.HexToAddress(gameAddr)
	v1, _ := cosmicgame.NewCosmicSignatureGame(game, client)
	v2, _ := cosmicgame.NewCosmicSignatureGameV2(game, client)
	if v1 == nil && v2 == nil {
		return fmt.Errorf("can't bind CosmicGame at %s", gameAddr)
	}

	var copts bind.CallOpts
	mechanics := probeContractMechanics(v1, v2, &copts)

	meta, err := allocChainSyncEvtlog(sw, gameAddr, client)
	if err != nil {
		return fmt.Errorf("alloc chain sync evtlog: %w", err)
	}

	var updated int
	for _, p := range buildContractParamSyncList(mechanics) {
		chainValue, err := p.read(v1, v2, &copts)
		if err != nil {
			errLog.Printf("chain sync: skip %s: %v", p.name, err)
			continue
		}
		var changed bool
		if p.name == "cst_reward_for_bidding" {
			changed, err = sw.SyncCstRewardIfMismatch(chainValue, meta)
		} else if p.contractAddr != "" {
			changed, err = sw.SyncAdminDecimalParamIfMismatchForContract(p.name, p.table, p.column, chainValue, p.contractAddr, meta)
		} else {
			changed, err = sw.SyncAdminDecimalParamIfMismatch(p.name, p.table, p.column, chainValue, meta)
		}
		if err != nil {
			return err
		}
		if changed {
			updated++
			info.Printf("chain sync: updated %s => %s", p.name, chainValue)
		}
	}

	if delay, err := readDelayDuration(v1, v2, &copts); err == nil {
		changed, err := sw.SyncAdminInt64ParamIfMismatch("delay_before_round_activation", "cg_delay_duration", "new_value", delay, meta)
		if err != nil {
			return err
		}
		if changed {
			updated++
			info.Printf("chain sync: updated delay_before_round_activation => %d", delay)
		}
	} else {
		errLog.Printf("chain sync: skip delay_before_round_activation: %v", err)
	}

	if prizesWalletAddr != "" {
		pw, err := cosmicgame.NewPrizesWallet(ethcommon.HexToAddress(prizesWalletAddr), client)
		if err != nil {
			errLog.Printf("chain sync: skip timeout_withdraw_prizes: bind PrizesWallet: %v", err)
		} else if val, err := pw.TimeoutDurationToWithdrawPrizes(&copts); err != nil {
			errLog.Printf("chain sync: skip timeout_withdraw_prizes: %v", err)
		} else {
			changed, err := sw.SyncAdminInt64ParamIfMismatchForContract(
				"timeout_withdraw_prizes", "cg_adm_timeout_withdraw", "new_timeout",
				val.Int64(), meta, prizesWalletAddr,
			)
			if err != nil {
				return err
			}
			if changed {
				updated++
				info.Printf("chain sync: updated timeout_withdraw_prizes => %s", val.String())
			}
		}
	}

	if cstAucChangeDiv := readCSTAuctionDurationChangeDivisor(v2, &copts, mechanics); cstAucChangeDiv >= 0 {
		valStr := fmt.Sprintf("%d", cstAucChangeDiv)
		changed, err := sw.SyncAdminDecimalParamIfMismatch("cst_dutch_auction_duration_change_divisor", "cg_adm_cst_auclen_chg_div", "new_len", valStr, meta)
		if err != nil {
			return err
		}
		if changed {
			updated++
			info.Printf("chain sync: updated cst_dutch_auction_duration_change_divisor => %s", valStr)
		}
	}

	info.Printf("chain sync complete: %d parameter(s) corrected from live RPC (mechanics v%d)", updated, mechanics)
	return nil
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
