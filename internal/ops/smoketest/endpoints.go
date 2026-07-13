package smoketest

import "net/url"

// BuildEndpoints returns the complete stable CosmicGame API smoke-test list.
func BuildEndpoints(input Params) []string {
	params := WithDefaults(input)
	const (
		offset    = "0"
		limit     = "1000000"
		sortParam = "0"
		interval  = "3600"
		topN      = "10"
		eventMin  = "0"
		eventMax  = "9223372036854775807"
	)
	user := params.UserAddress
	round := params.RoundNumber
	endpoints := make([]string, 0, 140)
	add := func(endpoint string) { endpoints = append(endpoints, endpoint) }

	// Statistics.
	add("/api/cosmicgame/statistics/dashboard")
	add("/api/cosmicgame/statistics/counters")
	add("/api/cosmicgame/statistics/unique/bidders")
	add("/api/cosmicgame/statistics/unique/winners")
	add("/api/cosmicgame/statistics/unique/donors")
	add("/api/cosmicgame/statistics/unique/stakers/cst")
	add("/api/cosmicgame/statistics/unique/stakers/randomwalk")
	add("/api/cosmicgame/statistics/unique/stakers/rwalk")
	add("/api/cosmicgame/statistics/unique/stakers/both")
	add("/api/cosmicgame/statistics/bidding/activity/" + params.TimestampMin + "/" + params.TimestampMax + "/" + interval)
	add("/api/cosmicgame/statistics/bidding/frequency/" + params.TimestampMin + "/" + params.TimestampMax + "/" + interval)
	add("/api/cosmicgame/statistics/bidding/top_active_periods/" + topN + "/" + params.TimestampMin + "/" + params.TimestampMax)
	add("/api/cosmicgame/statistics/bidding/time_bounds")

	// Rounds.
	add("/api/cosmicgame/rounds/list/" + offset + "/" + limit)
	add("/api/cosmicgame/rounds/info/" + round)
	add("/api/cosmicgame/rounds/current/time")

	// Prizes.
	add("/api/cosmicgame/prizes/history/global/" + offset + "/" + limit)
	add("/api/cosmicgame/prizes/history/by_user/" + user + "/" + offset + "/" + limit)
	add("/api/cosmicgame/prizes/eth/all/global")
	add("/api/cosmicgame/prizes/eth/all/global/" + offset + "/" + limit)
	add("/api/cosmicgame/prizes/eth/raffle/global")
	add("/api/cosmicgame/prizes/eth/raffle/global/" + offset + "/" + limit)
	add("/api/cosmicgame/prizes/eth/chronowarrior/global")
	add("/api/cosmicgame/prizes/eth/chronowarrior/global/" + offset + "/" + limit)
	add("/api/cosmicgame/prizes/eth/all/by_user/" + user)
	add("/api/cosmicgame/prizes/eth/raffle/by_user/" + user)
	add("/api/cosmicgame/prizes/eth/chronowarrior/by_user/" + user)
	add("/api/cosmicgame/prizes/eth/unclaimed/by_user/" + user + "/" + offset + "/" + limit)
	add("/api/cosmicgame/prizes/deposits/raffle/by_user/" + user)
	add("/api/cosmicgame/prizes/deposits/chrono_warrior/by_user/" + user)
	add("/api/cosmicgame/prizes/deposits/unclaimed/by_user/" + user + "/" + offset + "/" + limit)

	// Bid.
	add("/api/cosmicgame/bid/list/all/" + offset + "/" + limit)
	add("/api/cosmicgame/bid/info/" + params.BidEventLogID)
	add("/api/cosmicgame/bid/info_by_pos/" + params.BidRound + "/" + params.BidPosition)
	add("/api/cosmicgame/bid/with_message/by_round/" + round)
	add("/api/cosmicgame/bid/list/by_round/" + round + "/" + sortParam + "/" + offset + "/" + limit)
	add("/api/cosmicgame/bid/used_randomwalk_nfts")
	add("/api/cosmicgame/bid/used_rwalk_nfts")
	add("/api/cosmicgame/bid/cst_price")
	add("/api/cosmicgame/bid/eth_price")
	add("/api/cosmicgame/bid/current_special_winners")
	add("/api/cosmicgame/get_banned_bids")

	// CST NFT.
	add("/api/cosmicgame/cst/list/all/" + offset + "/" + limit)
	add("/api/cosmicgame/cst/list/by_user/" + user + "/" + offset + "/" + limit)
	add("/api/cosmicgame/cst/info/" + params.TokenID)
	add("/api/cosmicgame/cst/metadata/" + params.TokenID)
	add("/api/cosmicgame/cst/names/history/" + params.TokenID)
	add("/api/cosmicgame/cst/names/search/" + url.PathEscape(params.TokenName))
	add("/api/cosmicgame/cst/names/named_only")
	add("/api/cosmicgame/cst/transfers/all/" + params.TokenID + "/" + offset + "/" + limit)
	add("/api/cosmicgame/cst/transfers/by_user/" + user + "/" + offset + "/" + limit)
	add("/api/cosmicgame/cst/distribution")

	// CT (cosmic token).
	add("/api/cosmicgame/ct/balances")
	add("/api/cosmicgame/ct/statistics")
	add("/api/cosmicgame/ct/summary/by_user/" + user)
	add("/api/cosmicgame/ct/transfers/by_user/" + user + "/" + offset + "/" + limit)
	add("/api/cosmicgame/ct/total_supply_history_by_bid")
	add("/api/cosmicgame/ct/total_supply_history_by_date/" + params.FromDate + "/" + params.ToDate)

	// User.
	add("/api/cosmicgame/user/info/" + user)
	add("/api/cosmicgame/user/notif_red_box/" + user)
	add("/api/cosmicgame/user/balances/" + user)

	// Donations.
	add("/api/cosmicgame/donations/eth/simple/list/" + offset + "/" + limit)
	add("/api/cosmicgame/donations/eth/simple/by_round/" + round)
	add("/api/cosmicgame/donations/eth/with_info/list/" + offset + "/" + limit)
	add("/api/cosmicgame/donations/eth/with_info/by_round/" + round)
	add("/api/cosmicgame/donations/eth/with_info/info/" + params.ETHDonationID)
	add("/api/cosmicgame/donations/eth/by_user/" + user)
	add("/api/cosmicgame/donations/eth/both/by_round/" + round)
	add("/api/cosmicgame/donations/eth/both/all")
	add("/api/cosmicgame/donations/charity/deposits")
	add("/api/cosmicgame/donations/charity/cg_deposits")
	add("/api/cosmicgame/donations/charity/voluntary")
	add("/api/cosmicgame/donations/charity/withdrawals")
	add("/api/cosmicgame/donations/nft/list/" + offset + "/" + limit)
	add("/api/cosmicgame/donations/nft/info/" + params.NFTDonationID)
	add("/api/cosmicgame/donations/nft/by_user/" + user)
	add("/api/cosmicgame/donations/nft/claims")
	add("/api/cosmicgame/donations/nft/claims/" + offset + "/" + limit)
	add("/api/cosmicgame/donations/nft/claims/by_user/" + user)
	add("/api/cosmicgame/donations/nft/statistics")
	add("/api/cosmicgame/donations/nft/by_round/" + round)
	add("/api/cosmicgame/donations/nft/by_token/" + params.NFTTokenAddress)
	add("/api/cosmicgame/donations/nft/unclaimed/by_round/" + round)
	add("/api/cosmicgame/donations/nft/unclaimed/by_user/" + user)
	add("/api/cosmicgame/donations/erc20/by_round/detailed/" + round)
	add("/api/cosmicgame/donations/erc20/by_round/all/" + round)
	add("/api/cosmicgame/donations/erc20/by_round/summarized/" + round)
	add("/api/cosmicgame/donations/erc20/donated/by_user/" + user)
	add("/api/cosmicgame/donations/erc20/by_user/" + user)
	add("/api/cosmicgame/donations/erc20/global/" + offset + "/" + limit)
	add("/api/cosmicgame/donations/erc20/info/" + params.ERC20DonationID)
	add("/api/cosmicgame/donations/erc20/claims")
	add("/api/cosmicgame/donations/erc20/claims/" + offset + "/" + limit)
	add("/api/cosmicgame/donations/erc20/claims/by_user/" + user)
	add("/api/cosmicgame/donations/erc20/claims/by_round/" + round)

	// Raffle / deposits.
	add("/api/cosmicgame/raffle/deposits/list")
	add("/api/cosmicgame/raffle/deposits/list/" + offset + "/" + limit)
	add("/api/cosmicgame/raffle/deposits/by_round/" + round)
	add("/api/cosmicgame/eth_deposits/all/list/" + offset + "/" + limit)
	add("/api/cosmicgame/eth_deposits/raffle_eth/list/" + offset + "/" + limit)
	add("/api/cosmicgame/eth_deposits/chronowarrior_eth/list/" + offset + "/" + limit)
	add("/api/cosmicgame/raffle/nft/all/list")
	add("/api/cosmicgame/raffle/nft/all/list/" + offset + "/" + limit)
	add("/api/cosmicgame/raffle/nft/by_round/" + round)
	add("/api/cosmicgame/raffle/nft/by_user/" + user)

	// Staking CST.
	add("/api/cosmicgame/staking/cst/staked_tokens/all")
	add("/api/cosmicgame/staking/cst/staked_tokens/by_user/" + user)
	add("/api/cosmicgame/staking/cst/actions/global/" + offset + "/" + limit)
	add("/api/cosmicgame/staking/cst/actions/by_user/" + user + "/" + offset + "/" + limit)
	add("/api/cosmicgame/staking/cst/actions/info/" + params.CSTActionID)
	add("/api/cosmicgame/staking/cst/rewards/global")
	add("/api/cosmicgame/staking/cst/rewards/to_claim/by_user/" + user)
	add("/api/cosmicgame/staking/cst/rewards/collected/by_user/" + user + "/" + offset + "/" + limit)
	add("/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/" + params.CSTStakerAddress + "/" + params.DepositID)
	add("/api/cosmicgame/staking/cst/rewards/by_user/by_token/summary/" + user)
	add("/api/cosmicgame/staking/cst/rewards/by_user/by_token/details/" + user + "/" + params.TokenID)
	add("/api/cosmicgame/staking/cst/rewards/by_user/by_deposit/" + user)
	add("/api/cosmicgame/staking/cst/rewards/by_round/" + round)
	add("/api/cosmicgame/staking/cst/mints/global/" + offset + "/" + limit)
	add("/api/cosmicgame/staking/cst/mints/by_user/" + user)

	// Staking RandomWalk (canonical + rwalk alias).
	for _, resource := range []string{"randomwalk", "rwalk"} {
		add("/api/cosmicgame/staking/" + resource + "/actions/info/" + params.RandomWalkActionID)
		add("/api/cosmicgame/staking/" + resource + "/actions/global/" + offset + "/" + limit)
		add("/api/cosmicgame/staking/" + resource + "/actions/by_user/" + user + "/" + offset + "/" + limit)
		add("/api/cosmicgame/staking/" + resource + "/mints/global/" + offset + "/" + limit)
		add("/api/cosmicgame/staking/" + resource + "/mints/by_user/" + user)
		add("/api/cosmicgame/staking/" + resource + "/staked_tokens/all")
		add("/api/cosmicgame/staking/" + resource + "/staked_tokens/by_user/" + user)
	}

	// Marketing.
	add("/api/cosmicgame/marketing/rewards/global/" + offset + "/" + limit)
	add("/api/cosmicgame/marketing/rewards/by_user/" + user + "/" + offset + "/" + limit)
	add("/api/cosmicgame/marketing/config/current")

	// Time.
	add("/api/cosmicgame/time/current")
	add("/api/cosmicgame/time/until_prize")

	// System.
	add("/api/cosmicgame/system/modelist")
	add("/api/cosmicgame/system/modelist/" + offset + "/" + limit)
	add("/api/cosmicgame/system/admin_events/" + eventMin + "/" + eventMax)

	return endpoints
}
