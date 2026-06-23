// API smoke test for the cosmicgame websrv.
//
// Hits every /api/cosmicgame/... endpoint and reports non-200 responses (and
// in-body error/status:0) in RED as FAILED. A 400/500 or an "error" body
// usually means a broken SQL query in the handler.
//
// Real URL parameters (addresses, ids, rounds, token ids, timestamps) are
// fetched from the production database so the queries are exercised with live
// values. Missing values fall back to type-valid defaults.
//
// Build (single-file, like the other tools here):
//   go build api_smoke_test.go
//
// Run (env from your cg-prod.env):
//   source ~/configs/cg-prod.env
//   ./api_smoke_test
//   API_BASE=http://127.0.0.1:9090 ./api_smoke_test   # override base explicitly
//
// Env used:
//   PGSQL_HOST (host:port), PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD  (DB for param values)
//   HTTP_PORT  (websrv port; API base defaults to http://127.0.0.1:$HTTP_PORT)
//   API_BASE   (optional; overrides the base URL)
package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	colReset  = "\033[0m"
	colRed    = "\033[1;31m"
	colGreen  = "\033[32m"
	colYellow = "\033[33m"
)

type params struct {
	userAddr     string
	stakerCst    string
	roundNum     string
	bidEvtlogID  string
	bidRound     string
	bidPosition  string
	tokenID      string
	tokenName    string
	ethWiID      string
	nftDonID     string
	erc20DonID   string
	cstActionID  string
	rwalkActID   string
	depositID    string
	nftTokenAddr string
	tsMin        string
	tsMax        string
	fromDate     string
	toDate       string
}

func main() {
	base := apiBase()
	db := mustConnectDB()
	defer db.Close()

	p := fetchParams(db)
	fmt.Printf("API base : %s\n", base)
	fmt.Printf("Params   : userAddr=%s round=%s bidEvtlog=%s tokenId=%s nftDonId=%s erc20DonId=%s cstAction=%s rwalkAction=%s deposit=%s\n\n",
		p.userAddr, p.roundNum, p.bidEvtlogID, p.tokenID, p.nftDonID, p.erc20DonID, p.cstActionID, p.rwalkActID, p.depositID)

	endpoints := buildEndpoints(p)
	client := &http.Client{Timeout: 60 * time.Second}

	var failures []string
	okCount := 0
	for _, ep := range endpoints {
		full := base + ep
		status, bad, reason := doGet(client, full)
		if bad {
			fmt.Printf("%sFAILED%s [%s] %s  %s\n", colRed, colReset, statusStr(status), ep, reason)
			failures = append(failures, fmt.Sprintf("[%s] %s  %s", statusStr(status), ep, reason))
		} else {
			okCount++
			fmt.Printf("%sOK%s     [200] %s\n", colGreen, colReset, ep)
		}
	}

	fmt.Printf("\n==================== SUMMARY ====================\n")
	fmt.Printf("Total: %d   %sOK: %d%s   %sFAILED: %d%s\n", len(endpoints), colGreen, okCount, colReset, colRed, len(failures), colReset)
	if len(failures) > 0 {
		fmt.Printf("\n%sFailures:%s\n", colRed, colReset)
		for _, f := range failures {
			fmt.Printf("  %s%s%s\n", colRed, f, colReset)
		}
		os.Exit(1)
	}
	fmt.Printf("%sAll endpoints returned 200 with no error body.%s\n", colGreen, colReset)
}

func apiBase() string {
	if b := strings.TrimSpace(os.Getenv("API_BASE")); b != "" {
		return strings.TrimRight(b, "/")
	}
	port := strings.TrimSpace(os.Getenv("HTTP_PORT"))
	if port == "" {
		port = "9090"
	}
	return "http://127.0.0.1:" + port
}

func mustConnectDB() *sql.DB {
	host := os.Getenv("PGSQL_HOST")
	user := os.Getenv("PGSQL_USERNAME")
	dbName := os.Getenv("PGSQL_DATABASE")
	pass := os.Getenv("PGSQL_PASSWORD")
	if host == "" || user == "" || dbName == "" {
		fmt.Fprintf(os.Stderr, "%sPGSQL_HOST / PGSQL_USERNAME / PGSQL_DATABASE are required%s\n", colRed, colReset)
		os.Exit(2)
	}
	dsn := (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, pass),
		Host:     host, // host:port
		Path:     "/" + dbName,
		RawQuery: "sslmode=disable",
	}).String()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sDB open failed: %v%s\n", colRed, err, colReset)
		os.Exit(2)
	}
	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "%sDB ping failed: %v%s\n", colRed, err, colReset)
		os.Exit(2)
	}
	return db
}

// q1 runs a single-value query, returning fallback on any error or no row.
func q1(db *sql.DB, query, fallback string) string {
	var v sql.NullString
	if err := db.QueryRow(query).Scan(&v); err != nil || !v.Valid || v.String == "" {
		return fallback
	}
	return v.String
}

func fetchParams(db *sql.DB) params {
	p := params{}
	p.userAddr = q1(db, "SELECT a.addr FROM cg_bid b JOIN address a ON a.address_id=b.bidder_aid LIMIT 1",
		"0x0000000000000000000000000000000000000000")
	p.stakerCst = q1(db, "SELECT a.addr FROM cg_nft_staked_cst s JOIN address a ON a.address_id=s.staker_aid LIMIT 1", p.userAddr)
	// Prefer a round that actually has a prize claim; else the latest bid round.
	p.roundNum = q1(db, "SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1",
		q1(db, "SELECT round_num FROM cg_bid ORDER BY round_num DESC LIMIT 1", "0"))
	p.bidEvtlogID = q1(db, "SELECT evtlog_id FROM cg_bid ORDER BY id DESC LIMIT 1", "1")
	p.bidRound = q1(db, "SELECT round_num FROM cg_bid ORDER BY id DESC LIMIT 1", p.roundNum)
	p.bidPosition = q1(db, "SELECT bid_position FROM cg_bid ORDER BY id DESC LIMIT 1", "1")
	p.tokenID = q1(db, "SELECT token_id FROM cg_mint_event ORDER BY token_id DESC LIMIT 1", "0")
	p.tokenName = q1(db, "SELECT name FROM cg_token_name WHERE COALESCE(name,'')<>'' LIMIT 1", "a")
	p.ethWiID = q1(db, "SELECT id FROM cg_eth_donated_wi ORDER BY id DESC LIMIT 1", "1")
	p.nftDonID = q1(db, "SELECT id FROM cg_nft_donation ORDER BY id DESC LIMIT 1", "1")
	p.erc20DonID = q1(db, "SELECT id FROM cg_erc20_donation ORDER BY id DESC LIMIT 1", "1")
	p.cstActionID = q1(db, "SELECT action_id FROM cg_nft_staked_cst ORDER BY action_id DESC LIMIT 1", "0")
	p.rwalkActID = q1(db, "SELECT action_id FROM cg_nft_staked_rwalk ORDER BY action_id DESC LIMIT 1", "0")
	p.depositID = q1(db, "SELECT deposit_id FROM cg_staking_eth_deposit ORDER BY deposit_id DESC LIMIT 1", "0")
	p.nftTokenAddr = q1(db, "SELECT a.addr FROM cg_nft_donation d JOIN address a ON a.address_id=d.token_aid LIMIT 1",
		"0x0000000000000000000000000000000000000000")
	p.tsMin = q1(db, "SELECT EXTRACT(EPOCH FROM MIN(time_stamp))::bigint::text FROM cg_bid", "0")
	p.tsMax = q1(db, "SELECT EXTRACT(EPOCH FROM MAX(time_stamp))::bigint::text FROM cg_bid", "2000000000")
	// This endpoint expects YYYYMMDD (8 digits, no dashes).
	p.fromDate = q1(db, "SELECT to_char(MIN(time_stamp),'YYYYMMDD') FROM cg_bid", "20230101")
	p.toDate = q1(db, "SELECT to_char(MAX(time_stamp) + interval '1 day','YYYYMMDD') FROM cg_bid", "20300101")
	return p
}

func buildEndpoints(p params) []string {
	const (
		off   = "0"
		lim   = "1000000"
		sort  = "0"
		inter = "3600"
		n     = "10"
		eS    = "0"
		eE    = "9223372036854775807"
	)
	ua := p.userAddr
	r := p.roundNum
	var e []string
	add := func(s string) { e = append(e, s) }

	// Statistics
	add("/api/cosmicgame/statistics/dashboard")
	add("/api/cosmicgame/statistics/counters")
	add("/api/cosmicgame/statistics/unique/bidders")
	add("/api/cosmicgame/statistics/unique/winners")
	add("/api/cosmicgame/statistics/unique/donors")
	add("/api/cosmicgame/statistics/unique/stakers/cst")
	add("/api/cosmicgame/statistics/unique/stakers/randomwalk")
	add("/api/cosmicgame/statistics/unique/stakers/rwalk")
	add("/api/cosmicgame/statistics/unique/stakers/both")
	add("/api/cosmicgame/statistics/bidding/activity/" + p.tsMin + "/" + p.tsMax + "/" + inter)
	add("/api/cosmicgame/statistics/bidding/frequency/" + p.tsMin + "/" + p.tsMax + "/" + inter)
	add("/api/cosmicgame/statistics/bidding/top_active_periods/" + n + "/" + p.tsMin + "/" + p.tsMax)
	add("/api/cosmicgame/statistics/bidding/time_bounds")

	// Rounds
	add("/api/cosmicgame/rounds/list/" + off + "/" + lim)
	add("/api/cosmicgame/rounds/info/" + r)
	add("/api/cosmicgame/rounds/current/time")

	// Prizes
	add("/api/cosmicgame/prizes/history/global/" + off + "/" + lim)
	add("/api/cosmicgame/prizes/history/by_user/" + ua + "/" + off + "/" + lim)
	add("/api/cosmicgame/prizes/eth/all/global")
	add("/api/cosmicgame/prizes/eth/all/global/" + off + "/" + lim)
	add("/api/cosmicgame/prizes/eth/raffle/global")
	add("/api/cosmicgame/prizes/eth/raffle/global/" + off + "/" + lim)
	add("/api/cosmicgame/prizes/eth/chronowarrior/global")
	add("/api/cosmicgame/prizes/eth/chronowarrior/global/" + off + "/" + lim)
	add("/api/cosmicgame/prizes/eth/all/by_user/" + ua)
	add("/api/cosmicgame/prizes/eth/raffle/by_user/" + ua)
	add("/api/cosmicgame/prizes/eth/chronowarrior/by_user/" + ua)
	add("/api/cosmicgame/prizes/eth/unclaimed/by_user/" + ua + "/" + off + "/" + lim)
	add("/api/cosmicgame/prizes/deposits/raffle/by_user/" + ua)
	add("/api/cosmicgame/prizes/deposits/chrono_warrior/by_user/" + ua)
	add("/api/cosmicgame/prizes/deposits/unclaimed/by_user/" + ua + "/" + off + "/" + lim)

	// Bid
	add("/api/cosmicgame/bid/list/all/" + off + "/" + lim)
	add("/api/cosmicgame/bid/info/" + p.bidEvtlogID)
	add("/api/cosmicgame/bid/info_by_pos/" + p.bidRound + "/" + p.bidPosition)
	add("/api/cosmicgame/bid/with_message/by_round/" + r)
	add("/api/cosmicgame/bid/list/by_round/" + r + "/" + sort + "/" + off + "/" + lim)
	add("/api/cosmicgame/bid/used_randomwalk_nfts")
	add("/api/cosmicgame/bid/used_rwalk_nfts")
	add("/api/cosmicgame/bid/cst_price")
	add("/api/cosmicgame/bid/eth_price")
	add("/api/cosmicgame/bid/current_special_winners")
	add("/api/cosmicgame/get_banned_bids")

	// CST NFT
	add("/api/cosmicgame/cst/list/all/" + off + "/" + lim)
	add("/api/cosmicgame/cst/list/by_user/" + ua + "/" + off + "/" + lim)
	add("/api/cosmicgame/cst/info/" + p.tokenID)
	add("/api/cosmicgame/cst/metadata/" + p.tokenID)
	add("/api/cosmicgame/cst/names/history/" + p.tokenID)
	add("/api/cosmicgame/cst/names/search/" + url.PathEscape(p.tokenName))
	add("/api/cosmicgame/cst/names/named_only")
	add("/api/cosmicgame/cst/transfers/all/" + p.tokenID + "/" + off + "/" + lim)
	add("/api/cosmicgame/cst/transfers/by_user/" + ua + "/" + off + "/" + lim)
	add("/api/cosmicgame/cst/distribution")

	// CT (cosmic token)
	add("/api/cosmicgame/ct/balances")
	add("/api/cosmicgame/ct/statistics")
	add("/api/cosmicgame/ct/summary/by_user/" + ua)
	add("/api/cosmicgame/ct/transfers/by_user/" + ua + "/" + off + "/" + lim)
	add("/api/cosmicgame/ct/total_supply_history_by_bid")
	add("/api/cosmicgame/ct/total_supply_history_by_date/" + p.fromDate + "/" + p.toDate)

	// User
	add("/api/cosmicgame/user/info/" + ua)
	add("/api/cosmicgame/user/notif_red_box/" + ua)
	add("/api/cosmicgame/user/balances/" + ua)

	// Donations
	add("/api/cosmicgame/donations/eth/simple/list/" + off + "/" + lim)
	add("/api/cosmicgame/donations/eth/simple/by_round/" + r)
	add("/api/cosmicgame/donations/eth/with_info/list/" + off + "/" + lim)
	add("/api/cosmicgame/donations/eth/with_info/by_round/" + r)
	add("/api/cosmicgame/donations/eth/with_info/info/" + p.ethWiID)
	add("/api/cosmicgame/donations/eth/by_user/" + ua)
	add("/api/cosmicgame/donations/eth/both/by_round/" + r)
	add("/api/cosmicgame/donations/eth/both/all")
	add("/api/cosmicgame/donations/charity/deposits")
	add("/api/cosmicgame/donations/charity/cg_deposits")
	add("/api/cosmicgame/donations/charity/voluntary")
	add("/api/cosmicgame/donations/charity/withdrawals")
	add("/api/cosmicgame/donations/nft/list/" + off + "/" + lim)
	add("/api/cosmicgame/donations/nft/info/" + p.nftDonID)
	add("/api/cosmicgame/donations/nft/by_user/" + ua)
	add("/api/cosmicgame/donations/nft/claims")
	add("/api/cosmicgame/donations/nft/claims/" + off + "/" + lim)
	add("/api/cosmicgame/donations/nft/claims/by_user/" + ua)
	add("/api/cosmicgame/donations/nft/statistics")
	add("/api/cosmicgame/donations/nft/by_round/" + r)
	add("/api/cosmicgame/donations/nft/by_token/" + p.nftTokenAddr)
	add("/api/cosmicgame/donations/nft/unclaimed/by_round/" + r)
	add("/api/cosmicgame/donations/nft/unclaimed/by_user/" + ua)
	add("/api/cosmicgame/donations/erc20/by_round/detailed/" + r)
	add("/api/cosmicgame/donations/erc20/by_round/all/" + r)
	add("/api/cosmicgame/donations/erc20/by_round/summarized/" + r)
	add("/api/cosmicgame/donations/erc20/donated/by_user/" + ua)
	add("/api/cosmicgame/donations/erc20/by_user/" + ua)
	add("/api/cosmicgame/donations/erc20/global/" + off + "/" + lim)
	add("/api/cosmicgame/donations/erc20/info/" + p.erc20DonID)
	add("/api/cosmicgame/donations/erc20/claims")
	add("/api/cosmicgame/donations/erc20/claims/" + off + "/" + lim)
	add("/api/cosmicgame/donations/erc20/claims/by_user/" + ua)
	add("/api/cosmicgame/donations/erc20/claims/by_round/" + r)

	// Raffle / deposits
	add("/api/cosmicgame/raffle/deposits/list")
	add("/api/cosmicgame/raffle/deposits/list/" + off + "/" + lim)
	add("/api/cosmicgame/raffle/deposits/by_round/" + r)
	add("/api/cosmicgame/eth_deposits/all/list/" + off + "/" + lim)
	add("/api/cosmicgame/eth_deposits/raffle_eth/list/" + off + "/" + lim)
	add("/api/cosmicgame/eth_deposits/chronowarrior_eth/list/" + off + "/" + lim)
	add("/api/cosmicgame/raffle/nft/all/list")
	add("/api/cosmicgame/raffle/nft/all/list/" + off + "/" + lim)
	add("/api/cosmicgame/raffle/nft/by_round/" + r)
	add("/api/cosmicgame/raffle/nft/by_user/" + ua)

	// Staking CST
	add("/api/cosmicgame/staking/cst/staked_tokens/all")
	add("/api/cosmicgame/staking/cst/staked_tokens/by_user/" + ua)
	add("/api/cosmicgame/staking/cst/actions/global/" + off + "/" + lim)
	add("/api/cosmicgame/staking/cst/actions/by_user/" + ua + "/" + off + "/" + lim)
	add("/api/cosmicgame/staking/cst/actions/info/" + p.cstActionID)
	add("/api/cosmicgame/staking/cst/rewards/global")
	add("/api/cosmicgame/staking/cst/rewards/to_claim/by_user/" + ua)
	add("/api/cosmicgame/staking/cst/rewards/collected/by_user/" + ua + "/" + off + "/" + lim)
	add("/api/cosmicgame/staking/cst/rewards/action_ids_by_deposit/" + p.stakerCst + "/" + p.depositID)
	add("/api/cosmicgame/staking/cst/rewards/by_user/by_token/summary/" + ua)
	add("/api/cosmicgame/staking/cst/rewards/by_user/by_token/details/" + ua + "/" + p.tokenID)
	add("/api/cosmicgame/staking/cst/rewards/by_user/by_deposit/" + ua)
	add("/api/cosmicgame/staking/cst/rewards/by_round/" + r)
	add("/api/cosmicgame/staking/cst/mints/global/" + off + "/" + lim)
	add("/api/cosmicgame/staking/cst/mints/by_user/" + ua)

	// Staking RandomWalk (canonical + rwalk alias)
	for _, base := range []string{"randomwalk", "rwalk"} {
		add("/api/cosmicgame/staking/" + base + "/actions/info/" + p.rwalkActID)
		add("/api/cosmicgame/staking/" + base + "/actions/global/" + off + "/" + lim)
		add("/api/cosmicgame/staking/" + base + "/actions/by_user/" + ua + "/" + off + "/" + lim)
		add("/api/cosmicgame/staking/" + base + "/mints/global/" + off + "/" + lim)
		add("/api/cosmicgame/staking/" + base + "/mints/by_user/" + ua)
		add("/api/cosmicgame/staking/" + base + "/staked_tokens/all")
		add("/api/cosmicgame/staking/" + base + "/staked_tokens/by_user/" + ua)
	}

	// Marketing
	add("/api/cosmicgame/marketing/rewards/global/" + off + "/" + lim)
	add("/api/cosmicgame/marketing/rewards/by_user/" + ua + "/" + off + "/" + lim)
	add("/api/cosmicgame/marketing/config/current")

	// Time
	add("/api/cosmicgame/time/current")
	add("/api/cosmicgame/time/until_prize")

	// System
	add("/api/cosmicgame/system/modelist")
	add("/api/cosmicgame/system/modelist/" + off + "/" + lim)
	add("/api/cosmicgame/system/admin_events/" + eS + "/" + eE)

	return e
}

// doGet returns (statusCode, bad, reason). bad is true on transport error,
// non-200 status, or a JSON body that carries a non-empty "error" or "status":0.
func doGet(client *http.Client, fullURL string) (int, bool, string) {
	resp, err := client.Get(fullURL)
	if err != nil {
		return 0, true, "request error: " + err.Error()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, true, "non-200; body: " + snippet(body)
	}
	if reason, ok := bodyError(body); ok {
		return resp.StatusCode, true, "200 but error body: " + reason
	}
	return resp.StatusCode, false, ""
}

// bodyError best-effort detects {"status":0,...} or {"error":"..."} payloads.
func bodyError(body []byte) (string, bool) {
	s := strings.TrimSpace(string(body))
	if s == "" || s[0] != '{' {
		return "", false // arrays / non-object payloads: treat as fine
	}
	low := strings.ToLower(s)
	// crude but robust scan for an error/status signal without a schema
	if i := strings.Index(low, "\"error\""); i >= 0 {
		// find the value after the colon
		rest := s[i+len("\"error\""):]
		if c := strings.Index(rest, ":"); c >= 0 {
			val := strings.TrimSpace(rest[c+1:])
			if strings.HasPrefix(val, "\"") {
				if end := strings.Index(val[1:], "\""); end > 0 {
					msg := val[1 : 1+end]
					if msg != "" {
						return "error=" + msg, true
					}
				}
			}
		}
	}
	if strings.Contains(strings.ReplaceAll(low, " ", ""), "\"status\":0") {
		return "status=0; " + snippet(body), true
	}
	return "", false
}

func snippet(b []byte) string {
	s := strings.TrimSpace(string(b))
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > 160 {
		s = s[:160] + "..."
	}
	return s
}

func statusStr(code int) string {
	if code == 0 {
		return "ERR"
	}
	return fmt.Sprintf("%d", code)
}
