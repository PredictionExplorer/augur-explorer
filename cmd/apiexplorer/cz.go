// ChaosZero explorer pages: the route table, the sidebar navigation and the
// question/mask helpers. ChaosZero runs eight independent binary markets per
// Cosmic Signature cycle; the question list mirrors the frontend's
// QUESTIONS table (frontend/src/shared/questions.ts).
package main

import "fmt"

// czPages maps the ChaosZero explorer routes to templates and API paths
// (API server module internal/api/chaoszero).
var czPages = []page{
	{"GET /chaoszero/{$}", "chaoszero/home.html", "/api/chaoszero/statistics"},
	{"GET /chaoszero/contracts", "chaoszero/cz_contracts.html", "/api/chaoszero/contracts"},
	{"GET /chaoszero/rounds/list", "chaoszero/cz_rounds.html", "/api/chaoszero/rounds/list/0/100000"},
	{"GET /chaoszero/rounds/list/{offset}/{limit}", "chaoszero/cz_rounds.html", "/api/chaoszero/rounds/list/{offset}/{limit}"},
	{"GET /chaoszero/round/info/{round_num}", "chaoszero/cz_round_info.html", "/api/chaoszero/rounds/info/{round_num}"},
	{"GET /chaoszero/rounds/info/{round_num}", "chaoszero/cz_round_info.html", "/api/chaoszero/rounds/info/{round_num}"},
	{"GET /chaoszero/trades/list", "chaoszero/cz_trades.html", "/api/chaoszero/trades/list/0/1000"},
	{"GET /chaoszero/trades/list/{offset}/{limit}", "chaoszero/cz_trades.html", "/api/chaoszero/trades/list/{offset}/{limit}"},
	{"GET /chaoszero/trades/by_round/{round_num}", "chaoszero/cz_trades.html", "/api/chaoszero/trades/by_round/{round_num}/0/1000"},
	{"GET /chaoszero/trades/by_round/{round_num}/{offset}/{limit}", "chaoszero/cz_trades.html", "/api/chaoszero/trades/by_round/{round_num}/{offset}/{limit}"},
	{"GET /chaoszero/trades/by_user/{user_addr}", "chaoszero/cz_trades.html", "/api/chaoszero/trades/by_user/{user_addr}/0/1000"},
	{"GET /chaoszero/trades/by_user/{user_addr}/{offset}/{limit}", "chaoszero/cz_trades.html", "/api/chaoszero/trades/by_user/{user_addr}/{offset}/{limit}"},
	{"GET /chaoszero/traders/list", "chaoszero/cz_traders.html", "/api/chaoszero/traders/list/0/1000"},
	{"GET /chaoszero/traders/list/{offset}/{limit}", "chaoszero/cz_traders.html", "/api/chaoszero/traders/list/{offset}/{limit}"},
	{"GET /chaoszero/funding/list", "chaoszero/cz_funding.html", "/api/chaoszero/funding/list/0/1000"},
	{"GET /chaoszero/funding/list/{offset}/{limit}", "chaoszero/cz_funding.html", "/api/chaoszero/funding/list/{offset}/{limit}"},
	{"GET /chaoszero/splits/list", "chaoszero/cz_splits.html", "/api/chaoszero/splits/list/0/1000"},
	{"GET /chaoszero/splits/list/{offset}/{limit}", "chaoszero/cz_splits.html", "/api/chaoszero/splits/list/{offset}/{limit}"},
	{"GET /chaoszero/merges/list", "chaoszero/cz_merges.html", "/api/chaoszero/merges/list/0/1000"},
	{"GET /chaoszero/merges/list/{offset}/{limit}", "chaoszero/cz_merges.html", "/api/chaoszero/merges/list/{offset}/{limit}"},
	{"GET /chaoszero/settlements/list", "chaoszero/cz_settlements.html", "/api/chaoszero/settlements/list/0/1000"},
	{"GET /chaoszero/settlements/list/{offset}/{limit}", "chaoszero/cz_settlements.html", "/api/chaoszero/settlements/list/{offset}/{limit}"},
	{"GET /chaoszero/claims/list", "chaoszero/cz_claims.html", "/api/chaoszero/claims/list/0/1000"},
	{"GET /chaoszero/claims/list/{offset}/{limit}", "chaoszero/cz_claims.html", "/api/chaoszero/claims/list/{offset}/{limit}"},
	{"GET /chaoszero/claims/by_user/{user_addr}", "chaoszero/cz_claims.html", "/api/chaoszero/claims/by_user/{user_addr}/0/1000"},
	{"GET /chaoszero/claims/by_user/{user_addr}/{offset}/{limit}", "chaoszero/cz_claims.html", "/api/chaoszero/claims/by_user/{user_addr}/{offset}/{limit}"},
	{"GET /chaoszero/answers/list", "chaoszero/cz_answers.html", "/api/chaoszero/answers/list/0/1000"},
	{"GET /chaoszero/answers/list/{offset}/{limit}", "chaoszero/cz_answers.html", "/api/chaoszero/answers/list/{offset}/{limit}"},
	{"GET /chaoszero/escrows/list", "chaoszero/cz_escrows.html", "/api/chaoszero/escrows/list/0/1000"},
	{"GET /chaoszero/escrows/list/{offset}/{limit}", "chaoszero/cz_escrows.html", "/api/chaoszero/escrows/list/{offset}/{limit}"},
	{"GET /chaoszero/payouts/list", "chaoszero/cz_payouts.html", "/api/chaoszero/payouts/list/0/1000"},
	{"GET /chaoszero/payouts/list/{offset}/{limit}", "chaoszero/cz_payouts.html", "/api/chaoszero/payouts/list/{offset}/{limit}"},
	{"GET /chaoszero/transfers/list", "chaoszero/cz_transfers.html", "/api/chaoszero/transfers/list/0/1000"},
	{"GET /chaoszero/transfers/list/{offset}/{limit}", "chaoszero/cz_transfers.html", "/api/chaoszero/transfers/list/{offset}/{limit}"},
	{"GET /chaoszero/transfers/by_user/{user_addr}", "chaoszero/cz_transfers.html", "/api/chaoszero/transfers/by_user/{user_addr}/0/1000"},
	{"GET /chaoszero/transfers/by_user/{user_addr}/{offset}/{limit}", "chaoszero/cz_transfers.html", "/api/chaoszero/transfers/by_user/{user_addr}/{offset}/{limit}"},
	{"GET /chaoszero/approvals/list", "chaoszero/cz_approvals.html", "/api/chaoszero/approvals/list/0/1000"},
	{"GET /chaoszero/approvals/list/{offset}/{limit}", "chaoszero/cz_approvals.html", "/api/chaoszero/approvals/list/{offset}/{limit}"},
}

var czNav = []navGroup{
	{Title: "ChaosZero", Items: []navItem{
		{"Dashboard", "/chaoszero/"},
		{"Rounds (combined)", "/chaoszero/rounds/list"},
		{"Trades", "/chaoszero/trades/list"},
		{"Unique traders", "/chaoszero/traders/list"},
		{"Round funding", "/chaoszero/funding/list"},
		{"Settlements", "/chaoszero/settlements/list"},
		{"Claims", "/chaoszero/claims/list"},
		{"Contracts", "/chaoszero/contracts"},
	}},
	{Title: "Positions", Items: []navItem{
		{"Splits", "/chaoszero/splits/list"},
		{"Merges", "/chaoszero/merges/list"},
		{"Transfers (ERC-1155)", "/chaoszero/transfers/list"},
		{"Operator approvals", "/chaoszero/approvals/list"},
	}},
	{Title: "Settlement vault", Items: []navItem{
		{"Escrowed cycles", "/chaoszero/escrows/list"},
		{"Payouts", "/chaoszero/payouts/list"},
		{"Answers observed", "/chaoszero/answers/list"},
	}},
	{Title: "By round", Param: "round_num", Hint: "round #", Items: []navItem{
		{"Round info (combined)", "/chaoszero/round/info/{round_num}"},
		{"Trades", "/chaoszero/trades/by_round/{round_num}"},
	}},
	{Title: "By user", Param: "user_addr", Hint: "0x address", Items: []navItem{
		{"Trades", "/chaoszero/trades/by_user/{user_addr}"},
		{"Claims", "/chaoszero/claims/by_user/{user_addr}"},
		{"Position transfers", "/chaoszero/transfers/by_user/{user_addr}"},
	}},
}

// czQuestionTitles are the eight market questions, indexed by question id
// (mirrors frontend/src/shared/questions.ts).
var czQuestionTitles = [8]string{
	"More gestures?",
	"Longer Endurance?",
	"Longer Chrono?",
	"More ETH spent?",
	"More CST burned?",
	"Odd gesture count?",
	"One gesture, both titles?",
	"Higher CST peak?",
}

// oneMinus returns 1-v for complementary probabilities (P(NO) = 1-P(YES)).
func oneMinus(v any) float64 {
	return 1 - asFloat(v)
}

// asInt converts an API JSON number for use with the builtin index func.
func asInt(v any) int {
	return int(asFloat(v))
}

// czQuestion renders "N · title" for a question id from API JSON.
func czQuestion(v any) string {
	q := int(asFloat(v))
	if q < 0 || q >= len(czQuestionTitles) {
		return fmt.Sprint(v)
	}
	return fmt.Sprintf("%d · %s", q, czQuestionTitles[q])
}

// czQuestionTitle returns just the title for a question id (loop indices).
func czQuestionTitle(v any) string {
	q := int(asFloat(v))
	if q < 0 || q >= len(czQuestionTitles) {
		return ""
	}
	return czQuestionTitles[q]
}

// maskBits renders an 8-bit mask as individual Y/n glyphs, bit 0 (question
// 0) first — e.g. 0b00100101 renders "Y n Y n n Y n n".
func maskBits(v any) string {
	m := int(asFloat(v))
	out := make([]byte, 0, 16)
	for q := 0; q < 8; q++ {
		if q > 0 {
			out = append(out, ' ')
		}
		if m&(1<<q) != 0 {
			out = append(out, 'Y')
		} else {
			out = append(out, 'n')
		}
	}
	return string(out)
}
