package main

import (
	"fmt"
	"html"
	"html/template"
)

// Record-type labels shared by the dashboard tails and the fields view.

type badge struct{ Label, Class string }

// prizeTypes mirrors cg_prize.record_type (see cg_prizes_history_global.html).
var prizeTypes = map[int]badge{
	0:  {"Main Prize ETH", "badge-warning"},
	1:  {"Main Prize CST", "badge-info"},
	2:  {"Main Prize CS NFT", "badge-success"},
	3:  {"Last CST Bidder NFT", "badge-default"},
	4:  {"Last CST Bidder CST", "badge-info"},
	5:  {"Endurance Champion NFT", "badge-success"},
	6:  {"Endurance Champion CST", "badge-info"},
	7:  {"Chrono Warrior ETH", "badge-warning"},
	8:  {"Chrono Warrior CST", "badge-info"},
	9:  {"Chrono Warrior NFT", "badge-success"},
	10: {"Raffle ETH (Bidders)", "badge-warning"},
	11: {"Raffle CST (Bidders)", "badge-info"},
	12: {"Raffle NFT (Bidders)", "badge-success"},
	13: {"Raffle CST (RWalk)", "badge-info"},
	14: {"Raffle NFT (RWalk)", "badge-success"},
	15: {"Staking Deposit ETH", "badge-warning"},
	16: {"Donated NFT", "badge-default"},
	17: {"Donated ERC20", "badge-default"},
	18: {"Raffle ETH", "badge-default"},
}

// bidTypes mirrors cg_bid.bid_type.
var bidTypes = map[int]badge{
	0: {"ETH", "badge-info"},
	1: {"RandomWalk", "badge-success"},
	2: {"CST", "badge-warning"},
}

func renderBadge(table map[int]badge, v any) template.HTML {
	n := int(asFloat(v))
	b, ok := table[n]
	if !ok {
		return template.HTML(fmt.Sprintf(`<span class="badge badge-default">type %d</span>`, n)) // #nosec G203 -- integer
	}
	return template.HTML(fmt.Sprintf(`<span class="badge %s">%s</span>`, b.Class, html.EscapeString(b.Label))) // #nosec G203 -- escaped
}

// prizeTypeBadge renders a cg_prize record type.
func prizeTypeBadge(v any) template.HTML { return renderBadge(prizeTypes, v) }

// bidTypeBadge renders a bid type.
func bidTypeBadge(v any) template.HTML { return renderBadge(bidTypes, v) }

// mechanicsVersion names the contract mechanics generation.
func mechanicsVersion(v any) string {
	switch int(asFloat(v)) {
	case 1:
		return "V1"
	case 2:
		return "V2"
	case 3:
		return "V3.1"
	}
	return "unknown"
}
