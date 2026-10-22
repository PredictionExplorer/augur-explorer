package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// The Bids page lists one round at a time, picked with ?round=N (default:
// the current round, taken from the dashboard call). ?round=all lists every
// bid of every round through the legacy bid/list/all endpoint.

const (
	bidsAllAPI = "/api/cosmicgame/bid/list/all/0/100000"
	// bidsByRoundAPI lists the bids of one round newest first (sort=1).
	bidsByRoundAPI = "/api/cosmicgame/bid/list/by_round/%d/1/0/100000"
	bidsPage       = "cosmicsignature/cg_bids.html"
)

// annotateBidGaps adds SinceLastBid (seconds between a bid and the bid before
// it) to every row of a bid list that can be paired with its predecessor.
// Rows are consecutive bids in either order; the row with the oldest
// timestamp has no predecessor in the list and gets no key, which the
// templates render as "—".
func annotateBidGaps(rows []any) {
	ts := func(i int) (int64, bool) {
		m, ok := rows[i].(map[string]any)
		if !ok {
			return 0, false
		}
		tx, ok := m["Tx"].(map[string]any)
		if !ok {
			return 0, false
		}
		v, ok := tx["TimeStamp"]
		if !ok {
			return 0, false
		}
		return int64(asFloat(v)), true
	}
	for i := range rows {
		cur, ok := ts(i)
		if !ok {
			continue
		}
		// The predecessor is the neighbour with the smaller timestamp.
		var prev int64 = -1
		for _, j := range []int{i + 1, i - 1} {
			if j < 0 || j >= len(rows) {
				continue
			}
			if t, ok := ts(j); ok && t <= cur && (prev < 0 || t > prev) {
				prev = t
			}
		}
		if prev >= 0 {
			rows[i].(map[string]any)["SinceLastBid"] = jsonNum(cur - prev)
		}
	}
}

// cstSupplyTotals sums the CST minted and burned by a list of bids.
type cstSupplyTotals struct {
	Minted, Burned float64
}

// Net is the total-supply change: minted minus burned.
func (t cstSupplyTotals) Net() float64 { return t.Minted - t.Burned }

// annotateCstSupply adds CstMinted, CstBurned and CstSupplyDelta (all in
// CST) to every bid row and returns the totals. Per the contracts, a bid
// mints BidCstRewardAmount to the previous bidder (ETH, RandomWalk and CST
// bids alike) and a CST bid additionally burns the CST price paid. V1 rows
// carry "-1" in BidCstRewardAmount; their reward is CSTRewardEth. Round-end
// prize mints are not bids and are not included.
func annotateCstSupply(rows []any) cstSupplyTotals {
	var tot cstSupplyTotals
	for _, r := range rows {
		m, ok := r.(map[string]any)
		if !ok {
			continue
		}
		minted := asFloat(m["BidCstRewardAmountEth"])
		if minted < 0 {
			minted = asFloat(m["CSTRewardEth"])
		}
		if minted < 0 {
			minted = 0
		}
		var burned float64
		if asFloat(m["BidType"]) == 2 {
			if p := asFloat(m["CstPriceEth"]); p > 0 {
				burned = p
			}
		}
		m["CstMinted"] = jsonNum(minted)
		m["CstBurned"] = jsonNum(burned)
		m["CstSupplyDelta"] = jsonNum(minted - burned)
		tot.Minted += minted
		tot.Burned += burned
	}
	return tot
}

// signedCst formats a CST amount with an explicit sign, 4 decimals. The
// result is sign and digits only, so it is safe to return unescaped (the
// template escaper would otherwise turn "+" into "&#43;").
func signedCst(v any) template.HTML {
	return template.HTML(fmt.Sprintf("%+.4f", asFloat(v)))
}

// bidsRound picks the round for the Bids page: the raw query value when it
// is a valid round number (0..cur), else the current round.
func bidsRound(raw string, cur int) int {
	if n, err := strconv.Atoi(raw); err == nil && n >= 0 && n <= cur {
		return n
	}
	return cur
}

// roundChoices lists the rounds for the picker, newest first.
func roundChoices(cur int) []int {
	out := make([]int, 0, cur+1)
	for r := cur; r >= 0; r-- {
		out = append(out, r)
	}
	return out
}

// currentRound reads CurRoundNum from the dashboard call.
func (s *server) currentRound(ctx context.Context) (int, error) {
	dash, err := s.fetch(ctx, dashboardAPI)
	if err != nil {
		return 0, err
	}
	return int(asFloat(dash["CurRoundNum"])), nil
}

func (s *server) bidsHandler(w http.ResponseWriter, r *http.Request) {
	meta := pageMeta{Project: "cosmicsignature", Path: r.URL.Path}
	cur, err := s.currentRound(r.Context())
	if err != nil {
		s.log.Error("API fetch failed", "api_path", dashboardAPI, "err", err)
		s.renderError(w, http.StatusBadGateway, err.Error())
		return
	}
	sel := r.URL.Query().Get("round")
	all := sel == "all"
	round := bidsRound(sel, cur)
	api, rows := fmt.Sprintf(bidsByRoundAPI, round), "BidsByRound"
	if all {
		api, rows = bidsAllAPI, "Bids"
	}
	meta.API = api
	data, err := s.fetch(r.Context(), api)
	if err != nil {
		s.log.Error("API fetch failed", "api_path", api, "err", err)
		s.renderError(w, http.StatusBadGateway, err.Error())
		return
	}
	if r.URL.Query().Get("view") == "fields" {
		s.renderFields(w, data, meta)
		return
	}
	if list, ok := data[rows].([]any); ok {
		annotateBidGaps(list)
		data["CstTotals"] = annotateCstSupply(list)
	}
	data["Bids"] = data[rows]
	data["CurRoundNum"] = cur
	data["RoundChoices"] = roundChoices(cur)
	data["AllRounds"] = all
	data["SelectedRound"] = round // ignored by the template when AllRounds
	s.renderPage(w, http.StatusOK, bidsPage, data, meta)
}
