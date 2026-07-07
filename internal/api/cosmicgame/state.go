// state.go holds the cached CosmicGame contract/database state shared by the
// JSON API handlers. The values are loaded at startup by cosmic_game_init()
// and periodically refreshed by the reload_*_goroutine background goroutines.

package cosmicgame

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)
const (
	CONTRACT_CONSTANTS_REFRESH_TIME		= 5*60	// seconds
	CONTRACT_VARIABLES_REFRESH_TIME		= 5	// seconds
)
var (
	cosmic_game_addr				ethcommon.Address
	cosmic_signature_addr		ethcommon.Address
	cosmic_token_addr			ethcommon.Address
	charity_wallet_addr			ethcommon.Address
	marketing_wallet_addr		ethcommon.Address

	// contract constants (variables not frequently modified, and only by the owner)
	price_increase				string
	charity_addr				ethcommon.Address
	charity_percentage			int64
	token_reward				string
	prize_percentage			int64
	raffle_percentage			int64
	chrono_percentage			int64
	staking_percentage			int64
	time_increase				string
	initial_seconds				int64
	timeout_claim				int64
	roundstart_auclen			int64
	cst_dutch_auction_duration_change_divisor int64
	raffle_eth_winners_bidding	int64		// numRaffleETHWinnersBidding
	raffle_nft_winners_bidding	int64		// numRaffleNFTWinnersBidding
	raffle_nft_winners_staking_rwalk	int64	// numRaffleNFTWinnersStakingRWalk

	// contract variables (variables usually modified by a bid())
	bid_price					string
	bid_price_eth				float64
	prize_claim_date			int64
	prize_amount				string
	prize_amount_eth			float64
	raffle_amount				string
	raffle_amount_eth			float64
	staking_amount				string
	staking_amount_eth			float64
	round_num					int64
	round_activation_time_ts		int64	// RoundActivationTime() from contract (Unix sec), for dashboard when DB has none
	mainprize_microseconds_inc	string
	last_bidder					ethcommon.Address
	last_bidder_bid_time		int64
	charity_balance				string
	charity_balance_eth			float64
	round_start_ts				int64
	endurance_champ_addr		string
	endurance_duration			int64
	chrono_warrior_addr			string
	chrono_warrior_duration		int64
	lastcst_bidder_addr			string

	// contract counters	(collected via DB)
	bw_stats					CGStatistics

	arb_storagew				SQLStorageWrapper

	// arbRepo carries the queries already converted to the context-first,
	// error-returning Repo (Phase 1); arb_storagew keeps the rest until the
	// conversion completes.
	arbRepo						*Repo
)

// backgroundRefreshDisabled suppresses the reload_* refresh goroutines while
// keeping the synchronous initial loads in cosmic_game_init(). Only the API
// parity test harness sets it (via DisableBackgroundRefresh) so snapshots
// stay deterministic; production always runs the refresh loops. The whole
// mechanism goes away in Phase 2 when this package state becomes an
// injected ContractState component with a Run(ctx) lifecycle.
var backgroundRefreshDisabled bool

// DisableBackgroundRefresh prevents cosmic_game_init from starting the
// periodic contract/database refresh goroutines. Test-only; call before Init.
func DisableBackgroundRefresh() {
	backgroundRefreshDisabled = true
}

type rpcBlock struct {
    Timestamp         string      `json:"timestamp"`
}

type liveSpecialWinnersState struct {
	EnduranceChampionAddress        string
	EnduranceChampionDuration       int64
	EnduranceChampionStartTimeStamp int64
	PrevEnduranceChampionDuration   int64
	ChronoWarriorAddress            string
	ChronoWarriorDuration           int64
	ChronoWarriorIsLive             bool
	LastBidderAddress               string
	LastBidderLastBidTime           int64
	LastCstBidderAddress            string
	LastCstBidderLastBidTime        int64
	LastCstBidEventLogId            int64
	HasLastCstBidderLastBidTime     bool
	HasLastCstBidEventLogId         bool
	RoundNum                        int64
	SourceBlockNumber               uint64
	SourceBlockTimeStamp            int64
	Err                             error
}

func fetchLiveSpecialWinnersState() liveSpecialWinnersState {
	var state liveSpecialWinnersState
	state.RoundNum = round_num

	if EthClient == nil {
		state.Err = fmt.Errorf("ethereum client is not configured")
		return state
	}

	ctx := context.Background()
	header, err := EthClient.HeaderByNumber(ctx, nil)
	if err != nil {
		state.Err = fmt.Errorf("failed to fetch latest block header: %w", err)
		return state
	}
	state.SourceBlockNumber = header.Number.Uint64()
	state.SourceBlockTimeStamp = int64(header.Time)

	contract, err := NewCosmicSignatureGame(cosmic_game_addr, EthClient)
	if err != nil {
		state.Err = fmt.Errorf("failed to instantiate CosmicGame contract: %w", err)
		return state
	}

	copts := bind.CallOpts{Context: ctx, BlockNumber: header.Number}

	var (
		champs struct {
			EnduranceChampionAddress  ethcommon.Address
			EnduranceChampionDuration *big.Int
			ChronoWarriorAddress      ethcommon.Address
			ChronoWarriorDuration     *big.Int
		}
		enduranceStartTs            *big.Int
		prevEnduranceDuration       *big.Int
		storedEnduranceChampionAddr ethcommon.Address
		storedEnduranceChampionDur  *big.Int
		lastBidder                  ethcommon.Address
		lastCstBidder               ethcommon.Address
		storedChronoWarriorDur      *big.Int
		lastBidderLastBidTime       *big.Int
		lastCstBidderLastBidTime    *big.Int
	)

	var wg sync.WaitGroup
	var mu sync.Mutex
	recordErr := func(label string, err error) {
		if err == nil {
			return
		}
		mu.Lock()
		defer mu.Unlock()
		if state.Err == nil {
			state.Err = fmt.Errorf("%s: %w", label, err)
		}
	}

	wg.Add(8)
	go func() {
		defer wg.Done()
		result, err := contract.TryGetCurrentChampions(&copts)
		if err != nil {
			recordErr("TryGetCurrentChampions", err)
			return
		}
		mu.Lock()
		champs = result
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.EnduranceChampionStartTimeStamp(&copts)
		if err != nil {
			recordErr("EnduranceChampionStartTimeStamp", err)
			return
		}
		mu.Lock()
		enduranceStartTs = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.PrevEnduranceChampionDuration(&copts)
		if err != nil {
			recordErr("PrevEnduranceChampionDuration", err)
			return
		}
		mu.Lock()
		prevEnduranceDuration = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.LastBidderAddress(&copts)
		if err != nil {
			recordErr("LastBidderAddress", err)
			return
		}
		mu.Lock()
		lastBidder = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.LastCstBidderAddress(&copts)
		if err != nil {
			recordErr("LastCstBidderAddress", err)
			return
		}
		mu.Lock()
		lastCstBidder = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.ChronoWarriorDuration(&copts)
		if err != nil {
			recordErr("ChronoWarriorDuration", err)
			return
		}
		mu.Lock()
		storedChronoWarriorDur = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.EnduranceChampionAddress(&copts)
		if err != nil {
			recordErr("EnduranceChampionAddress", err)
			return
		}
		mu.Lock()
		storedEnduranceChampionAddr = val
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val, err := contract.EnduranceChampionDuration(&copts)
		if err != nil {
			recordErr("EnduranceChampionDuration", err)
			return
		}
		mu.Lock()
		storedEnduranceChampionDur = val
		mu.Unlock()
	}()
	wg.Wait()

	if state.Err != nil {
		return state
	}

	state.EnduranceChampionAddress = champs.EnduranceChampionAddress.String()
	state.EnduranceChampionDuration = champs.EnduranceChampionDuration.Int64()
	state.ChronoWarriorAddress = champs.ChronoWarriorAddress.String()
	state.ChronoWarriorDuration = champs.ChronoWarriorDuration.Int64()
	state.LastBidderAddress = lastBidder.String()
	state.LastCstBidderAddress = lastCstBidder.String()
	// The chrono-segment anchor (EnduranceChampionStartTimeStamp / PrevEnduranceChampionDuration)
	// and ChronoWarriorIsLive are computed below, after we have the last bidder's lastBidTimeStamp,
	// so the anchor stays consistent with the LIVE champion returned by tryGetCurrentChampions().

	if state.RoundNum >= 0 {
		roundBig := big.NewInt(state.RoundNum)
		wg.Add(2)
		go func() {
			defer wg.Done()
			info, err := contract.BiddersInfo(&copts, roundBig, lastBidder)
			if err != nil {
				recordErr("BiddersInfo(lastBidder)", err)
				return
			}
			mu.Lock()
			lastBidderLastBidTime = info.LastBidTimeStamp
			mu.Unlock()
		}()
		go func() {
			defer wg.Done()
			if lastCstBidder == (ethcommon.Address{}) {
				return
			}
			info, err := contract.BiddersInfo(&copts, roundBig, lastCstBidder)
			if err != nil {
				recordErr("BiddersInfo(lastCstBidder)", err)
				return
			}
			mu.Lock()
			lastCstBidderLastBidTime = info.LastBidTimeStamp
			mu.Unlock()
		}()
		wg.Wait()

		if state.Err != nil {
			return state
		}

		if lastBidderLastBidTime != nil {
			state.LastBidderLastBidTime = lastBidderLastBidTime.Int64()
		}
		if lastCstBidderLastBidTime != nil {
			state.LastCstBidderLastBidTime = lastCstBidderLastBidTime.Int64()
			state.HasLastCstBidderLastBidTime = true
		}

		// Derive the LIVE endurance-champion anchor consistently with tryGetCurrentChampions().
		// When the current last bidder has overtaken the stored endurance record, the contract
		// recomputes the champion in-memory, but the enduranceChampionStartTimeStamp() and
		// prevEnduranceChampionDuration() storage getters still return the OLD (stale) anchor.
		// Mixing the live champion (from tryGetCurrentChampions) with that stale anchor produced
		// a wrong Chrono-Warrior "record-growing segment" and "is live" status.
		liveEnduranceStart := enduranceStartTs.Int64()
		livePrevDuration := prevEnduranceDuration.Int64()
		if lastBidder != (ethcommon.Address{}) && lastBidderLastBidTime != nil {
			lastBidDuration := state.SourceBlockTimeStamp - lastBidderLastBidTime.Int64()
			if storedEnduranceChampionAddr == (ethcommon.Address{}) {
				// No champion recorded yet: the last bidder is the (live) endurance champion.
				liveEnduranceStart = lastBidderLastBidTime.Int64()
			} else if storedEnduranceChampionDur != nil && lastBidDuration > storedEnduranceChampionDur.Int64() {
				// Last bidder overtook the stored record: champion start/prev are recomputed live.
				livePrevDuration = storedEnduranceChampionDur.Int64()
				liveEnduranceStart = lastBidderLastBidTime.Int64()
			}
		}
		state.EnduranceChampionStartTimeStamp = liveEnduranceStart
		state.PrevEnduranceChampionDuration = livePrevDuration

		chronoSegmentStart := liveEnduranceStart + livePrevDuration
		currentChronoSegmentDuration := state.SourceBlockTimeStamp - chronoSegmentStart
		state.ChronoWarriorIsLive = currentChronoSegmentDuration > storedChronoWarriorDur.Int64()

		if lastCstBidder != (ethcommon.Address{}) {
			if evtlogId, ok := arb_storagew.Get_last_cst_bid_evtlog_for_bidder(state.RoundNum, lastCstBidder.String()); ok {
				state.LastCstBidEventLogId = evtlogId
				state.HasLastCstBidEventLogId = true
			}
		}
	}

	return state
}

func cosmic_game_init() {

	if  !dbInitialized() {
		err_str := fmt.Sprintf("cosmic_game_init(): Database link wasn't configured")
		Info.Print(err_str)
		Error.Print(err_str)
	}
	arb_storagew.S=common.Ctx.Db
	arb_storagew.S.Db_set_schema_name("public")
	arbRepo = NewRepo(common.Ctx.Store)
	bw_caddrs, err := arbRepo.ContractAddrs(context.Background())
	if err != nil {
		// Startup cannot proceed without the contract registry; this keeps
		// the legacy fatal behavior (the exit used to live inside the store).
		err_str := fmt.Sprintf("cosmic_game_init(): reading contract addresses: %v", err)
		Error.Print(err_str)
		Info.Print(err_str)
		fmt.Printf("\nFATAL: %s\n", err_str)
		fmt.Printf("HINT: If you don't need CosmicGame, set ENABLE_ROUTES_COSMICGAME=false in websrv .env\n\n")
		os.Exit(1)
	}
	cosmic_game_addr = ethcommon.HexToAddress(bw_caddrs.CosmicGameAddr)
	cosmic_signature_addr = ethcommon.HexToAddress(bw_caddrs.CosmicSignatureAddr)
	cosmic_token_addr = ethcommon.HexToAddress(bw_caddrs.CosmicTokenAddr)
	charity_wallet_addr = ethcommon.HexToAddress(bw_caddrs.CharityWalletAddr)
	marketing_wallet_addr = ethcommon.HexToAddress(bw_caddrs.MarketingWalletAddr)
	do_reload_contract_variables()
	do_reload_database_variables()
	if backgroundRefreshDisabled {
		// Run the constants load that reload_constants_goroutine would have
		// performed immediately, then stop: no periodic refresh in tests.
		do_reload_contract_constants()
		return
	}
	go reload_database_variables_goroutine()
	go reload_constants_goroutine()
	go reload_variables_goroutine()
}
func get_cosmic_game_contract_balance() float64 {

	cg_eth_bal,err := EthClient.BalanceAt(context.Background(),cosmic_game_addr,nil)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceAt() call for cosmic game: %v\n",err)
		Info.Print(err_str)
		Error.Print(err_str)
		return math.NaN()
	}
	f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
	f_balance := big.NewFloat(0.0).SetInt(cg_eth_bal)
	f_quo := big.NewFloat(0.0).Quo(f_balance,f_divisor)
	f_out,_ := f_quo.Float64()
	return f_out
}
func do_reload_contract_constants() {
	var copts bind.CallOpts
	code,err := EthClient.CodeAt(context.Background(), cosmic_game_addr, nil)
	if (err != nil) {
		err_str := fmt.Sprintf("Can't instantiate Cosmic gane contract: %v\n",err)
		Error.Print(err_str)
		Info.Print(err_str)
		fmt.Print(err_str)
	}
	if len(code) == 0 {
		err_str := fmt.Sprintf("Can't instantiate Cosmic gane contract: no code at given address\n")
		Error.Print(err_str)
		Info.Print(err_str)
		fmt.Print(err_str)
	}
	v1contract, v2contract := bindCosmicGameLiveReaders(cosmic_game_addr, EthClient)
	if v1contract == nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract at %v . Contract constants won't be fetched\n", cosmic_game_addr)
		Error.Print(err_str)
		Info.Print(err_str)
	} else {
		bwcontract := v1contract
		var err error
		var tmp_val *big.Int
		tmp_val,err = bwcontract.EthBidPriceIncreaseDivisor(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PriceIncrease() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			price_increase = "error"
		} else { price_increase=tmp_val.String() }
		charity_addr,err = bwcontract.CharityAddress(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
		}
		tmp_val,err = bwcontract.CharityEthDonationAmountPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			charity_percentage = 0
		} else { charity_percentage = tmp_val.Int64() }
		rewardStr, err := readTokenReward(bwcontract, v2contract, &copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TokenReward() call: %v\n", err)
			Error.Print(err_str)
			Info.Print(err_str)
			token_reward = "error"
		} else {
			token_reward = rewardStr
		}
		tmp_val,err = bwcontract.MainEthPrizeAmountPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizePercentage() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			prize_percentage = -1
		} else { prize_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.RaffleTotalEthPrizeAmountForBiddersPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at RafflePercentage() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			raffle_percentage = -1
		} else { raffle_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.ChronoWarriorEthPrizeAmountPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at ChronoWarriorEthPrizeAmountPercentage(() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			chrono_percentage = -1
		} else { chrono_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at StakingPercentage() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			staking_percentage = -1
		} else { staking_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.MainPrizeTimeIncrementIncreaseDivisor(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TimeIncrease() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			time_increase = "error"
		} else { time_increase = tmp_val.String() }
		tmp_val,err = bwcontract.NumRaffleEthPrizesForBidders(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumRaffleETHWinnersBidding() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			raffle_eth_winners_bidding = -1 
		} else { raffle_eth_winners_bidding = tmp_val.Int64()}
		tmp_val,err = bwcontract.NumRaffleCosmicSignatureNftsForBidders(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumRaffleNFTWinnersBidding() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			raffle_nft_winners_bidding = -1
		} else { raffle_nft_winners_bidding = tmp_val.Int64() }
		tmp_val,err = bwcontract.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumRaffleNFTWinnersStakingRWalk() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			raffle_nft_winners_staking_rwalk = -1
		} else { raffle_nft_winners_staking_rwalk = tmp_val.Int64() }
		cst_dutch_auction_duration_change_divisor = readCSTAuctionDurationChangeDivisor(bwcontract, v2contract, &copts)
	}
}
func do_reload_contract_variables() {
	var copts bind.CallOpts
	v1contract, v2contract := bindCosmicGameLiveReaders(cosmic_game_addr, EthClient)
	if v1contract == nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract at %v . Contract constants won't be fetched\n", cosmic_game_addr)
		Error.Print(err_str)
		Info.Print(err_str)
	} else {
		bwcontract := v1contract
		var err error
		var tmp_val *big.Int
		f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
		tmp_val,err = bwcontract.GetNextEthBidPrice(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at GetBidPrice() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			bid_price = "error"
		} else {
			bid_price = tmp_val.String()
			f_bid_price := big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_bid_price,f_divisor)
			bid_price_eth,_ = f_quo.Float64()
		}
		tmp_val,err = bwcontract.GetDurationUntilMainPrize(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeTime() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			prize_claim_date = -1
		} else { prize_claim_date = tmp_val.Int64() }
		tmp_val , err = bwcontract.GetMainEthPrizeAmount(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeAmount() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			prize_amount = "error"
		} else {
			prize_amount = tmp_val.String()
			f_prize_amount:= big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_prize_amount,f_divisor)
			prize_amount_eth,_ = f_quo.Float64()
		}
		tmp_val , err = bwcontract.GetCosmicSignatureNftStakingTotalEthRewardAmount(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at GetCosmicSignatureNftStakingTotalEthRewardAmount() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			staking_amount = "error"
		} else {
			staking_amount = tmp_val.String()
			f_staking_amount:= big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_staking_amount,f_divisor)
			staking_amount_eth,_ = f_quo.Float64()
		}
		tmp_val , err = bwcontract.GetRaffleTotalEthPrizeAmountForBidders(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at RaffleAmount() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			raffle_amount = "error"
		} else {
			raffle_amount = tmp_val.String()
			f_raffle_amount:= big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_raffle_amount,f_divisor)
			raffle_amount_eth,_ = f_quo.Float64()
		}
		tmp_val , err = bwcontract.RoundNum(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumPrizes() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			round_num = -1
		} else { round_num = tmp_val.Int64() }
		tmp_val,err = bwcontract.MainPrizeTimeIncrementInMicroSeconds(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at MainPrizeTimeIncrementInMicroseconds() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			mainprize_microseconds_inc = "error"
		} else { mainprize_microseconds_inc = tmp_val.String() }
		last_bidder,err = bwcontract.LastBidderAddress(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at LastBidder() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
		}
		if round_num > -1 {
			tmp_bidder_info,err := bwcontract.BiddersInfo(&copts,big.NewInt(round_num),last_bidder)
			if err != nil {
				err_str := fmt.Sprintf("Error at BiddersInfo() call: %v\n",err)
				Error.Print(err_str)
				Info.Print(err_str)
				last_bidder_bid_time = -1
			} else { last_bidder_bid_time = tmp_bidder_info.LastBidTimeStamp.Int64() }
		}
		tmp_val,err = bwcontract.InitialDurationUntilMainPrizeDivisor(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at ImitialDurationUntilMainPrizeDivisor() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			initial_seconds = -1
		} else { initial_seconds = tmp_val.Int64() }
		tmp_val,err = bwcontract.TimeoutDurationToClaimMainPrize(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TimeoutClaimPrize() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			timeout_claim = -1
		} else { timeout_claim = tmp_val.Int64() }
		roundstart_auclen = readRoundStartCSTAuctionSetting(bwcontract, v2contract, &copts)
		if roundstart_auclen == -1 {
			err_str := "Error reading CST round-start auction setting (V1 divisor / V2 duration)\n"
			Error.Print(err_str)
			Info.Print(err_str)
		}
		tmp_val,err = EthClient.BalanceAt(context.Background(),charity_addr,nil)
		if err != nil {
			err_str := fmt.Sprintf("Error at BalanceAt() call for charity addr: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
			charity_balance = "error"
		} else {
			charity_balance = tmp_val.String()
			f_charity_balance := big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_charity_balance,f_divisor)
			charity_balance_eth,_ = f_quo.Float64()
		}
		//tmp_addr1, tmp_duration1,tmp_addr2,tmp_duration2, err := bwcontract.TryGetCurrentChampions(&copts);
		champs,err := bwcontract.TryGetCurrentChampions(&copts);
		if err != nil {
			err_str := fmt.Sprintf("Error at TryGetCurrentChampions() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
		} else {
			endurance_champ_addr = champs.EnduranceChampionAddress.String()
			endurance_duration = champs.EnduranceChampionDuration.Int64()
			chrono_warrior_addr = champs.ChronoWarriorAddress.String()
			chrono_warrior_duration = champs.ChronoWarriorDuration.Int64()
		}
		tmp_addr, err := bwcontract.LastCstBidderAddress(&copts);
		if err != nil {
			err_str := fmt.Sprintf("Error at lastCstBidderAddress() call: %v\n",err)
			Error.Print(err_str)
			Info.Print(err_str)
		} else {
			lastcst_bidder_addr = tmp_addr.String()
		}
	}
}
func do_reload_database_variables() {
	bw_stats = arb_storagew.Get_cosmic_game_statistics()
	round_start_ts = arb_storagew.Get_round_start_timestamp(bw_stats.TotalPrizes)
}
func reload_constants_goroutine() {
	// we will load contract constants up web requests but to avoid having to restart
	// the server every time a constant changes we will have a refresh of the constants
	// every few minutes
	for {
		do_reload_contract_constants()
		time.Sleep(CONTRACT_CONSTANTS_REFRESH_TIME * time.Second)
	}
}
func reload_variables_goroutine() {
	for {
		do_reload_contract_variables()
		time.Sleep(CONTRACT_VARIABLES_REFRESH_TIME * time.Second)
	}
}
func reload_database_variables_goroutine() {
	for {
		do_reload_database_variables()
		time.Sleep(CONTRACT_VARIABLES_REFRESH_TIME * time.Second)
	}
}
