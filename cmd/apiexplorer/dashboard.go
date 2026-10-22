package main

import (
	"context"
	"net/http"
	"sync"
)

// The CosmicSignature dashboard aggregates several API calls into one
// screen: the statistics/dashboard payload plus live tails of the bids,
// prize-history and system-mode tables. The dashboard call is mandatory;
// the tails are optional and render as "unavailable" when they fail, so a
// single slow endpoint does not blank the whole page.

const (
	dashboardAPI = "/api/cosmicgame/statistics/dashboard"
	// recentBidRows is the size of the bids tail; one extra row is fetched
	// so the oldest shown bid still has a predecessor for SinceLastBid.
	recentBidRows = 12
	recentBidsAPI = "/api/cosmicgame/bid/list/all/0/13"
)

// dashboardExtras maps template keys to the optional API calls and the
// response field that holds the rows.
var dashboardExtras = []struct {
	Key, API, Field string
}{
	{"RecentBids", recentBidsAPI, "Bids"},
	{"RecentPrizes", "/api/cosmicgame/prizes/history/global/0/10", "GlobalPrizeHistory"},
	{"RecentModes", "/api/cosmicgame/system/modelist/0/5", "SystemModeChanges"},
	{"UntilPrize", "/api/cosmicgame/time/until_prize", "TimeUntilPrize"},
}

// fetchDashboard runs the dashboard call, then the extras concurrently.
func (s *server) fetchDashboard(ctx context.Context) (map[string]any, error) {
	data, err := s.fetch(ctx, dashboardAPI)
	if err != nil {
		return nil, err
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, ex := range dashboardExtras {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := s.fetch(ctx, ex.API)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				s.log.Warn("dashboard extra failed", "api_path", ex.API, "err", err)
				data[ex.Key] = nil
				return
			}
			data[ex.Key] = resp[ex.Field]
		}()
	}
	wg.Wait()
	if bids, ok := data["RecentBids"].([]any); ok {
		annotateBidGaps(bids)
		if len(bids) > recentBidRows {
			bids = bids[:recentBidRows]
		}
		data["RecentBids"] = bids
	}
	return data, nil
}

func (s *server) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	meta := pageMeta{Project: "cosmicsignature", Path: r.URL.Path, API: dashboardAPI}
	if r.URL.Query().Get("view") == "fields" {
		data, err := s.fetch(r.Context(), dashboardAPI)
		if err != nil {
			s.renderError(w, http.StatusBadGateway, err.Error())
			return
		}
		s.renderFields(w, data, meta)
		return
	}
	data, err := s.fetchDashboard(r.Context())
	if err != nil {
		s.log.Error("API fetch failed", "api_path", dashboardAPI, "err", err)
		s.renderError(w, http.StatusBadGateway, err.Error())
		return
	}
	s.renderPage(w, http.StatusOK, "cosmicsignature/cg_index.html", data, meta)
}
