package main

// Sidebar navigation: every explorer route grouped by subject so any table
// is one click away from any page. Groups with a Param render an input in
// the sidebar; their item links contain a "{param}" placeholder that the
// page script substitutes with the input value.

type navItem struct {
	Label string
	Href  string
}

type navGroup struct {
	Title string
	// Param is the placeholder name used in the group's item links
	// ("round_num", "user_addr", ...). Empty for plain link groups.
	Param string
	// Hint is the input placeholder shown for Param groups.
	Hint  string
	Items []navItem
}

// projectNav returns the sidebar groups for a project ("cosmicsignature" or
// "randomwalk"); nil for the explorer index.
func projectNav(project string) []navGroup {
	switch project {
	case "cosmicsignature":
		return cgNav
	case "randomwalk":
		return rwNav
	}
	return nil
}

var cgNav = []navGroup{
	{Title: "Game", Items: []navItem{
		{"Dashboard", "/cosmicsignature/"},
		{"Bids", "/cosmicsignature/bid/list/all"},
		{"Rounds", "/cosmicsignature/rounds/list"},
		{"Prize history", "/cosmicsignature/prizes/history/global"},
		{"Claims by round", "/cosmicsignature/statistics/claims/by_round"},
		{"Special winners (live)", "/cosmicsignature/bid/current_special_winners"},
		{"Bid type ratio", "/cosmicsignature/bid/bid_type_ratio"},
		{"Used RWalk NFTs", "/cosmicsignature/bid/used_randomwalk_nfts"},
		{"CST bid price (live)", "/cosmicsignature/bid/cst_price"},
		{"ETH bid price (live)", "/cosmicsignature/bid/eth_price"},
		{"Time until prize", "/cosmicsignature/time/until_prize"},
	}},
	{Title: "Prizes", Items: []navItem{
		{"ETH deposits (all)", "/cosmicsignature/prizes/eth/all"},
		{"Raffle ETH", "/cosmicsignature/prizes/eth/raffle"},
		{"Chrono warrior ETH", "/cosmicsignature/prizes/eth/chronowarrior"},
		{"Raffle NFT wins", "/cosmicsignature/raffle/nft/all/list"},
		{"ROI leaderboard", "/cosmicsignature/statistics/leaderboard/roi"},
		{"Unique winners", "/cosmicsignature/statistics/unique/winners"},
		{"Unique bidders", "/cosmicsignature/statistics/unique/bidders"},
	}},
	{Title: "Tokens", Items: []navItem{
		{"CS NFT mints", "/cosmicsignature/cst/list/all"},
		{"CS NFT named only", "/cosmicsignature/cst/names/named_only"},
		{"CS NFT distribution", "/cosmicsignature/cst/distribution"},
		{"CosmicToken balances", "/cosmicsignature/ct/balances"},
		{"CosmicToken statistics", "/cosmicsignature/ct/statistics"},
	}},
	{Title: "Donations", Items: []navItem{
		{"ETH simple", "/cosmicsignature/donations/eth/simple/list"},
		{"ETH with info", "/cosmicsignature/donations/eth/with_info/list"},
		{"ETH both (all)", "/cosmicsignature/donations/eth/both/all"},
		{"ERC20 global", "/cosmicsignature/donations/erc20/global"},
		{"ERC20 claims", "/cosmicsignature/donations/erc20/claims"},
		{"NFT list", "/cosmicsignature/donations/nft/list"},
		{"NFT claims", "/cosmicsignature/donations/nft/claims"},
		{"NFT statistics", "/cosmicsignature/donations/nft/statistics"},
		{"Unique donors", "/cosmicsignature/statistics/unique/donors"},
		{"Charity deposits", "/cosmicsignature/donations/charity/deposits"},
		{"Charity CG deposits", "/cosmicsignature/donations/charity/cg_deposits"},
		{"Charity voluntary", "/cosmicsignature/donations/charity/voluntary"},
		{"Charity withdrawals", "/cosmicsignature/donations/charity/withdrawals"},
	}},
	{Title: "Staking CST", Items: []navItem{
		{"Actions", "/cosmicsignature/staking/cst/actions/global"},
		{"Staked tokens", "/cosmicsignature/staking/cst/staked_tokens/all"},
		{"Rewards (deposits)", "/cosmicsignature/staking/cst/rewards/global"},
		{"Raffle mints", "/cosmicsignature/staking/cst/mints/global"},
		{"Unique stakers", "/cosmicsignature/statistics/unique/stakers/cst"},
	}},
	{Title: "Staking RWalk", Items: []navItem{
		{"Actions", "/cosmicsignature/staking/randomwalk/actions/global"},
		{"Staked tokens", "/cosmicsignature/staking/randomwalk/staked_tokens/all"},
		{"Raffle mints", "/cosmicsignature/staking/randomwalk/mints/global"},
		{"Unique stakers", "/cosmicsignature/statistics/unique/stakers/randomwalk"},
		{"Unique stakers (both)", "/cosmicsignature/statistics/unique/stakers/both"},
	}},
	{Title: "System", Items: []navItem{
		{"Mode changes", "/cosmicsignature/system/modelist"},
		{"Marketing rewards", "/cosmicsignature/marketing/rewards/global"},
		{"Marketing config", "/cosmicsignature/marketing/config/current"},
		{"Server time", "/cosmicsignature/time/current"},
		{"API docs", "/cosmicsignature/api_docs"},
	}},
	{Title: "By round", Param: "round_num", Hint: "round #", Items: []navItem{
		{"Round info", "/cosmicsignature/round/info/{round_num}"},
		{"Bids", "/cosmicsignature/bid/list/all?round={round_num}"},
		{"Claim detail", "/cosmicsignature/statistics/claims/detail/{round_num}"},
		{"ETH prizes", "/cosmicsignature/prizes/eth/round/{round_num}"},
		{"Raffle NFT winners", "/cosmicsignature/raffle/nft/by_round/{round_num}"},
		{"Staking CST deposits", "/cosmicsignature/staking/cst/rewards/by_round/{round_num}"},
		{"Donations ETH simple", "/cosmicsignature/donations/eth/simple/by_round/{round_num}"},
		{"Donations ETH w/ info", "/cosmicsignature/donations/eth/with_info/by_round/{round_num}"},
		{"Donations ETH both", "/cosmicsignature/donations/eth/both/by_round/{round_num}"},
		{"Donations ERC20 detailed", "/cosmicsignature/donations/erc20/by_round/detailed/{round_num}"},
		{"Donations ERC20 all", "/cosmicsignature/donations/erc20/by_round/all/{round_num}"},
		{"Donations ERC20 summarized", "/cosmicsignature/donations/erc20/by_round/summarized/{round_num}"},
		{"Donated NFTs", "/cosmicsignature/donations/nft/by_round/{round_num}"},
		{"Donated NFTs unclaimed", "/cosmicsignature/donations/nft/unclaimed/by_round/{round_num}"},
	}},
	{Title: "By user", Param: "user_addr", Hint: "0x address", Items: []navItem{
		{"User info", "/cosmicsignature/user/info/{user_addr}"},
		{"Balances", "/cosmicsignature/user/balances/{user_addr}"},
		{"Notification box", "/cosmicsignature/user/notif_red_box/{user_addr}"},
		{"Prize history", "/cosmicsignature/prizes/history/by_user/{user_addr}/0/100000"},
		{"ETH prizes", "/cosmicsignature/prizes/eth/user/{user_addr}"},
		{"ETH prizes raffle", "/cosmicsignature/prizes/eth/user_raffle/{user_addr}"},
		{"ETH prizes chrono", "/cosmicsignature/prizes/eth/user_chronowarrior/{user_addr}"},
		{"ETH prizes unclaimed", "/cosmicsignature/prizes/eth/user_unclaimed/{user_addr}"},
		{"Raffle deposits", "/cosmicsignature/prizes/deposits/raffle/by_user/{user_addr}"},
		{"Chrono deposits", "/cosmicsignature/prizes/deposits/chrono_warrior/by_user/{user_addr}"},
		{"Raffle NFT wins", "/cosmicsignature/raffle/nft/by_user/{user_addr}"},
		{"CS NFTs owned", "/cosmicsignature/cst/list/by_user/{user_addr}"},
		{"CS NFT transfers", "/cosmicsignature/cst/transfers/by_user/{user_addr}"},
		{"CosmicToken summary", "/cosmicsignature/ct/summary/by_user/{user_addr}"},
		{"CosmicToken transfers", "/cosmicsignature/ct/transfers/by_user/{user_addr}"},
		{"Donations ETH", "/cosmicsignature/donations/eth/by_user/{user_addr}"},
		{"Donations ERC20 made", "/cosmicsignature/donations/erc20/donated/by_user/{user_addr}"},
		{"Donations ERC20 received", "/cosmicsignature/donations/erc20/by_user/{user_addr}"},
		{"Donations ERC20 claims", "/cosmicsignature/donations/erc20/claims/by_user/{user_addr}"},
		{"Donated NFTs made", "/cosmicsignature/donations/nft/by_user/{user_addr}"},
		{"Donated NFT claims", "/cosmicsignature/donations/nft/claims/by_user/{user_addr}"},
		{"Donated NFTs unclaimed", "/cosmicsignature/donations/nft/unclaimed/by_user/{user_addr}"},
		{"Staking CST tokens", "/cosmicsignature/staking/cst/staked_tokens/by_user/{user_addr}"},
		{"Staking CST actions", "/cosmicsignature/staking/cst/actions/by_user/{user_addr}"},
		{"Staking CST rewards to claim", "/cosmicsignature/staking/cst/rewards/to_claim/by_user/{user_addr}"},
		{"Staking CST rewards collected", "/cosmicsignature/staking/cst/rewards/collected/by_user/{user_addr}"},
		{"Staking CST rewards by token", "/cosmicsignature/staking/cst/rewards/by_user/by_token/summary/{user_addr}"},
		{"Staking CST rewards by deposit", "/cosmicsignature/staking/cst/rewards/by_user/by_deposit/{user_addr}"},
		{"Staking CST mints", "/cosmicsignature/staking/cst/mints/by_user/{user_addr}"},
		{"Staking RWalk tokens", "/cosmicsignature/staking/randomwalk/staked_tokens/by_user/{user_addr}"},
		{"Staking RWalk actions", "/cosmicsignature/staking/randomwalk/actions/by_user/{user_addr}"},
		{"Staking RWalk mints", "/cosmicsignature/staking/randomwalk/mints/by_user/{user_addr}"},
		{"Marketing rewards", "/cosmicsignature/marketing/rewards/by_user/{user_addr}"},
	}},
	{Title: "By CS token", Param: "token_id", Hint: "token id", Items: []navItem{
		{"Token info", "/cosmicsignature/cst/info/{token_id}"},
		{"Name history", "/cosmicsignature/cst/names/history/{token_id}"},
		{"Ownership transfers", "/cosmicsignature/cst/transfers/all/{token_id}"},
	}},
	{Title: "By record id", Param: "record_id", Hint: "id", Items: []navItem{
		{"Bid (evtlog id)", "/cosmicsignature/bid/info/{record_id}"},
		{"ETH donation w/ info", "/cosmicsignature/donations/eth/with_info/info/{record_id}"},
		{"NFT donation", "/cosmicsignature/donations/nft/info/{record_id}"},
		{"ERC20 donation", "/cosmicsignature/donations/erc20/info/{record_id}"},
		{"Stake action CST", "/cosmicsignature/staking/cst/actions/info/{record_id}"},
		{"Stake action RWalk", "/cosmicsignature/staking/randomwalk/actions/info/{record_id}"},
		{"Admin events from evtlog", "/cosmicsignature/system/admin_events/{record_id}/9223372036854775807"},
	}},
}

var rwNav = []navGroup{
	{Title: "RandomWalk", Items: []navItem{
		{"Dashboard", "/randomwalk/"},
		{"Mint report", "/randomwalk/mint_report"},
		{"Token mints", "/randomwalk/tokens/list/sequential"},
		{"Trade history", "/randomwalk/trading/history"},
		{"Finalized sales", "/randomwalk/trading/sales"},
		{"Active offers", "/randomwalk/current_offers/0"},
		{"Floor price", "/randomwalk/floor_price"},
	}},
	{Title: "Statistics", Items: []navItem{
		{"Market stats", "/randomwalk/statistics/by_market"},
		{"Token stats", "/randomwalk/statistics/by_token"},
		{"Mint intervals", "/randomwalk/statistics/mint_intervals"},
		{"Withdrawal chart", "/randomwalk/statistics/withdrawal_chart"},
		{"Floor price chart", "/randomwalk/statistics/floor_price/0/0/0"},
		{"Volume chart", "/randomwalk/statistics/trading_volume/0/0/0"},
	}},
	{Title: "By token", Param: "token_id", Hint: "token id", Items: []navItem{
		{"Token info", "/randomwalk/tokens/info/{token_id}"},
		{"History", "/randomwalk/tokens/history/{token_id}"},
		{"Name changes", "/randomwalk/tokens/name_changes/{token_id}"},
	}},
	{Title: "By user (address id)", Param: "user_aid", Hint: "address id", Items: []navItem{
		{"User info", "/randomwalk/user/info/{user_aid}"},
		{"Tokens", "/randomwalk/tokens/by_user/{user_aid}"},
		{"Trading", "/randomwalk/trading/by_user/{user_aid}"},
	}},
}
