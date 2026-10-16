package main

import "fmt"

// Help-balloon text for the admin-events page: what each admin event means
// and what values it accepts. Keyed by the record_type numbering of
// cgmodel.CGAdminEvent (the registry in
// internal/store/cosmicgame/admin_events.go is the source of truth for the
// numbering; 43 and 45 are retired there and deliberately absent here).
var adminEventHelpTexts = map[int]string{
	1: "Share of a round's ETH pot donated to charity when the main prize is claimed. " +
		"Accepts a percentage (whole number); all pot percentages together must stay below 100.",
	2: "Share of a round's ETH pot paid to the main prize winner (the last bidder). " +
		"Accepts a percentage (whole number); all pot percentages together must stay below 100.",
	3: "Share of a round's ETH pot split evenly among the raffle ETH winners drawn from that round's bidders. " +
		"Accepts a percentage (whole number).",
	4: "Share of a round's ETH pot paid to the Cosmic Signature NFT staking rewards pool. " +
		"Accepts a percentage (whole number).",
	5: "How many raffle ETH prizes are drawn among bidders at the end of each round. " +
		"Accepts a non-negative whole number of winners.",
	6: "How many raffle Cosmic Signature NFTs are minted for randomly chosen bidders each round. " +
		"Accepts a non-negative whole number of winners.",
	7: "Pause between a main prize claim and the next round becoming active (bids rejected meanwhile). " +
		"Accepts a duration in seconds.",
	8: "How many raffle Cosmic Signature NFTs are minted for randomly chosen Random Walk NFT stakers each round. " +
		"Accepts a non-negative whole number of winners.",
	9: "New charity wallet: the address that receives the charity share of each round's pot. " +
		"Accepts a non-zero address.",
	10: "New Random Walk NFT contract address, used to validate discounted (half-price) ETH bids made with a Random Walk NFT. " +
		"Accepts a non-zero contract address.",
	11: "New PrizesWallet contract: escrows raffle/special ETH prizes and donated tokens and NFTs until winners withdraw them. " +
		"Accepts a non-zero contract address.",
	12: "New staking wallet for Cosmic Signature NFTs (receives and distributes the staking share of each pot). " +
		"Accepts a non-zero contract address.",
	13: "New staking wallet for Random Walk NFTs (its stakers are eligible for NFT raffles; no ETH rewards). " +
		"Accepts a non-zero contract address.",
	14: "New marketing wallet: receives a CST contribution on each main prize claim to fund promotion. " +
		"Accepts a non-zero contract address.",
	15: "New Cosmic Signature Token (CST, ERC-20) contract address. " +
		"Accepts a non-zero contract address.",
	16: "New Cosmic Signature NFT (ERC-721) contract address. " +
		"Accepts a non-zero contract address.",
	17: "The game proxy was upgraded to a new implementation contract (ERC-1967). " +
		"The value is the new implementation's address; the game generation (V1/V2/V3) changes here.",
	18: "Divisor controlling how the per-bid main-prize time increment itself grows when a round ends " +
		"(the increment is raised by 1/divisor each round). Accepts a positive whole number.",
	19: "How long after the main prize becomes claimable it stays reserved for the last bidder; " +
		"after this timeout anyone may claim it. Accepts a duration in seconds.",
	20: "Divisor setting how much the next ETH bid price rises after each ETH bid " +
		"(next price = paid price × (1 + 1/divisor); 100 means +1%). Accepts a positive whole number.",
	21: "Base time added to the main prize deadline per bid. " +
		"Accepts a duration in microseconds (e.g. 3600000000 = 1 hour).",
	22: "Initial main-prize countdown armed by the first bid of a round " +
		"(V1 set it directly in seconds; V2+ derive it from the time increment via a divisor). " +
		"Accepts a positive whole number (seconds or divisor, by game generation).",
	23: "New treasurer: the account authorized to manage and pay out the marketing wallet's CST funds. " +
		"Accepts a non-zero address.",
	24: "When the current round activates and starts accepting bids. " +
		"Accepts a Unix timestamp in seconds; settable only while the round is inactive.",
	25: "Length of the CST Dutch auction (the declining-price window for CST bids). " +
		"V1 set the duration in seconds; V2+ set a divisor applied to the main-prize time increment. " +
		"Accepts a positive whole number.",
	26: "CST reward multiplier for ERC-20 payouts to a round's participants. " +
		"Accepts an amount in wei of CST (shown here in whole CST).",
	27: "Floor for the price a new CST Dutch auction may begin at. " +
		"Accepts an amount in wei of CST (shown in whole CST). " +
		"On the review-branch V3 contracts the setter is disabled (the value is frozen at its V2 setting).",
	28: "CST contribution minted to the marketing wallet on each main prize claim. " +
		"Accepts an amount in wei of CST (shown in whole CST).",
	29: "CST reward for placing a bid. V1/V2 minted a fixed amount to the bidder; " +
		"V3 mints to the previous (outbid) bidder, scaled by how long their bid held the lead via this multiplier. " +
		"Accepts an amount in wei of CST (shown in whole CST).",
	30: "Maximum length of the message a bidder may attach to a bid. " +
		"Accepts a non-negative whole number of bytes.",
	31: "URL of the script that generates Cosmic Signature NFT artwork from each NFT's on-chain seed. " +
		"Accepts any string (expected to be a URL).",
	32: "Base URI for Cosmic Signature NFT token metadata (tokenURI = base URI + token ID). " +
		"Accepts any string (expected to be a URL ending in '/').",
	33: "An OpenZeppelin initializer ran on the game proxy. " +
		"The value is the initialization version: 1 = initial deployment, 2 = V2 upgrade, 3 = V3 upgrade. " +
		"The indexer derives each bid's mechanics generation from these events.",
	34: "Ownership of the contract moved to a new owner (OpenZeppelin Ownable). " +
		"Values are the previous and new owner addresses; the new owner controls all admin setters.",
	35: "How long winners' prizes stay reserved in the PrizesWallet; " +
		"after this timeout anyone may withdraw unclaimed ETH prizes. Accepts a duration in seconds.",
	36: "Divisor deriving the ETH Dutch auction duration from the main-prize time increment " +
		"(the auction prices the first ETH bid of a round). Accepts a positive whole number.",
	37: "Divisor setting the ETH Dutch auction's ending price relative to its beginning price " +
		"(2 means the price halves over the auction). Accepts a positive whole number.",
	38: "Share of a round's ETH pot paid to the Chrono-Warrior (longest cumulative time as Endurance Champion). " +
		"Accepts a percentage (whole number).",
	39: "Divisor controlling the per-bid adjustment of the CST Dutch auction duration " +
		"(each bid nudges the duration by 1/divisor). Accepts a positive whole number.",
	40: "Divisor deriving the round's late-bid window from the main-prize time increment: " +
		"bids placed within that window before the main prize deadline pay a price premium. " +
		"Accepts a positive whole number (V3 only).",
	41: "Base multiplier of the late-bid price premium charged near the main prize deadline. " +
		"Accepts a positive whole number (V3 only).",
	42: "Exponent shaping how steeply the late-bid price premium grows as the main prize deadline approaches. " +
		"Accepts a non-negative whole number (V3 only).",
	44: "How many Cosmic Signature NFTs are minted to the main prize winner on a claim. " +
		"Accepts a positive whole number (V3 only).",
	46: "Threshold for refunding ETH bid overpayment, measured in gas units: " +
		"overpayments cheaper than this much gas are kept by the game instead of refunded. " +
		"Accepts a non-negative whole number of gas units (V3 only).",
	47: "DAO: minimum CST voting power required to create a governance proposal. " +
		"Accepts an amount in wei of CST (shown in whole CST).",
	48: "DAO: delay between a proposal's creation and the start of voting. " +
		"Accepts a duration (Governor clock units; this DAO uses seconds).",
	49: "DAO: how long voting on a proposal stays open. " +
		"Accepts a duration (Governor clock units; this DAO uses seconds).",
	50: "DAO: quorum required for a proposal to pass, as a numerator over 100 of total CST voting supply. " +
		"Accepts a whole number from 0 to 100.",
	51: "DAO: a new governance proposal was created. " +
		"Values show the proposer's address, the proposal ID and its description.",
	52: "DAO: a passed proposal was queued for execution. " +
		"The numeric value is the earliest execution time (ETA) as a Unix timestamp.",
	53: "DAO: a queued proposal was executed on-chain.",
	54: "DAO: a proposal was canceled before execution.",
	55: "DAO: a vote was cast on a proposal. " +
		"Values show the voter, their voting weight in CST, the choice (For / Against / Abstain) and an optional reason.",
	56: "Cosmic Token: a delegate's voting power changed (CST transfers or delegation). " +
		"The value is the delegate's new total voting power in CST.",
}

// roundRangeLabel renders the header's round label: "Round N" when the
// range lies in one round, "Rounds N to M" when it spans several, and ""
// when the API did not supply the round numbers (older apiserver builds),
// so the page still renders against them.
func roundRangeLabel(start, end any) string {
	toInt := func(v any) (int64, bool) {
		switch n := v.(type) {
		case jsonNum:
			return int64(n), true
		case float64:
			return int64(n), true
		case int:
			return int64(n), true
		default:
			return 0, false
		}
	}
	s, okS := toInt(start)
	e, okE := toInt(end)
	switch {
	case !okS || !okE:
		return ""
	case s == e:
		return fmt.Sprintf("Round %d", s)
	default:
		return fmt.Sprintf("Rounds %d to %d", s, e)
	}
}

// adminEventHelp returns the help-balloon text for an admin event record
// type, or "" when none is defined (the template then renders no balloon).
// Record types arrive as jsonNum because the page data is decoded JSON and
// run through convertNums.
func adminEventHelp(recordType any) string {
	switch v := recordType.(type) {
	case jsonNum:
		return adminEventHelpTexts[int(v)]
	case float64:
		return adminEventHelpTexts[int(v)]
	case int:
		return adminEventHelpTexts[v]
	default:
		return ""
	}
}
