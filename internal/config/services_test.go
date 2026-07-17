package config

import (
	"errors"
	"strings"
	"testing"
	"time"
)

// minimalAPIServerEnv is the smallest environment LoadAPIServer accepts.
func minimalAPIServerEnv() map[string]string {
	return map[string]string{
		"RPC_URL":   "https://rpc.example.com/v2/KEY",
		"HTTP_PORT": "8080",
	}
}

func TestLoadAPIServerMinimal(t *testing.T) {
	t.Parallel()
	cfg, err := LoadAPIServer(mapEnv(minimalAPIServerEnv()))
	if err != nil {
		t.Fatalf("LoadAPIServer: %v", err)
	}
	if !cfg.EnableCosmicGame || !cfg.EnableRandomWalk || !cfg.EnableFAQ {
		t.Errorf("module enables should default to true: %+v", cfg)
	}
	if cfg.ExploreMaxTokenID != 3766 {
		t.Errorf("ExploreMaxTokenID = %d, want the 3766 default", cfg.ExploreMaxTokenID)
	}
	if cfg.Log.Format != "text" || cfg.Log.Level != "info" {
		t.Errorf("Log defaults = %+v", cfg.Log)
	}
}

func TestLoadAPIServerFull(t *testing.T) {
	t.Parallel()
	env := minimalAPIServerEnv()
	env["ENABLE_ROUTES_FAQ"] = "false"
	env["RANKING_VOTE_CHAIN_IDS"] = "1, 42161"
	env["RANDOMWALK_EXPLORE_MAX_TOKEN_ID"] = "100"
	env["NFT_ASSETS_FLAT_PATHS"] = "1"
	env["WEBSRV_IMAGE_NO_CACHE"] = "true" // normalized: any accepted bool spelling works now
	env["DATABASE_URL"] = "postgres://u:p@h/d"
	env["LOG_FORMAT"] = "json"
	env["AI_BOT_BACKEND_URL"] = "http://faq:8000"
	env["V1_SUNSET_AT"] = "2027-01-01T00:00:00Z"
	cfg, err := LoadAPIServer(mapEnv(env))
	if err != nil {
		t.Fatalf("LoadAPIServer: %v", err)
	}
	if cfg.EnableFAQ {
		t.Error("ENABLE_ROUTES_FAQ=false ignored")
	}
	if len(cfg.RankingVoteChainIDs) != 2 || cfg.RankingVoteChainIDs[1] != 42161 {
		t.Errorf("RankingVoteChainIDs = %v", cfg.RankingVoteChainIDs)
	}
	if cfg.ExploreMaxTokenID != 100 || !cfg.NFTAssetsFlatPaths || !cfg.ImageNoCache {
		t.Errorf("cfg = %+v", cfg)
	}
	if cfg.DB.URL != "postgres://u:p@h/d" { //nolint:gosec // G101: fake test credentials
		t.Errorf("DB.URL = %q", cfg.DB.URL)
	}
	if cfg.FAQUpstream() != "http://faq:8000" {
		t.Errorf("FAQUpstream = %q", cfg.FAQUpstream())
	}
	if want := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC); !cfg.V1SunsetAt.Equal(want) {
		t.Errorf("V1SunsetAt = %v, want %v", cfg.V1SunsetAt, want)
	}
}

func TestLoadAPIServerAggregatesProblems(t *testing.T) {
	t.Parallel()
	_, err := LoadAPIServer(mapEnv(map[string]string{
		"ENABLE_ROUTES_COSMICGAME":        "maybe",
		"RANDOMWALK_EXPLORE_MAX_TOKEN_ID": "-5",
		"LOG_FORMAT":                      "xml",
	}))
	if err == nil {
		t.Fatal("bad config accepted")
	}
	for _, want := range []string{
		"RPC_URL: required but not set",
		"ENABLE_ROUTES_COSMICGAME: cannot parse",
		"HTTP_PORT: no listeners configured",
		"RANDOMWALK_EXPLORE_MAX_TOKEN_ID: must be a positive integer",
		"LOG_FORMAT: unknown format",
	} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("missing %q in:\n%s", want, err)
		}
	}
}

func TestLoadAPIServerMalformedMaxTokenIDReportsOnce(t *testing.T) {
	t.Parallel()
	env := minimalAPIServerEnv()
	env["RANDOMWALK_EXPLORE_MAX_TOKEN_ID"] = "many"
	_, err := LoadAPIServer(mapEnv(env))
	if err == nil {
		t.Fatal("malformed max token id accepted")
	}
	if n := strings.Count(err.Error(), "RANDOMWALK_EXPLORE_MAX_TOKEN_ID"); n != 1 {
		t.Errorf("variable reported %d times, want once:\n%s", n, err)
	}
}

func TestLoadAPIServerHTTPSOnlyListener(t *testing.T) {
	t.Parallel()
	cfg, err := LoadAPIServer(mapEnv(map[string]string{
		"RPC_URL":        "https://rpc.example.com",
		"HTTPS_HOSTNAME": ":443",
	}))
	if err != nil {
		t.Fatalf("HTTPS-only config rejected: %v", err)
	}
	if cfg.HTTPSHostname != ":443" {
		t.Errorf("HTTPSHostname = %q", cfg.HTTPSHostname)
	}
}

func TestFAQUpstreamPrecedence(t *testing.T) {
	t.Parallel()
	cfg := &APIServer{FAQUpstreamURL: "a", FAQUpstreamURLLegacy: "b"}
	if cfg.FAQUpstream() != "a" {
		t.Errorf("AI_BOT_BACKEND_URL should win, got %q", cfg.FAQUpstream())
	}
	cfg.FAQUpstreamURL = ""
	if cfg.FAQUpstream() != "b" {
		t.Errorf("legacy alias should apply, got %q", cfg.FAQUpstream())
	}
	cfg.FAQUpstreamURLLegacy = ""
	if cfg.FAQUpstream() != "" {
		t.Errorf("unset upstream should be empty, got %q", cfg.FAQUpstream())
	}
}

func TestLoadETL(t *testing.T) {
	t.Parallel()
	if _, err := LoadETL(mapEnv(nil)); err == nil || !strings.Contains(err.Error(), "RPC_URL: required") {
		t.Errorf("missing RPC_URL not reported: %v", err)
	}
	cfg, err := LoadETL(mapEnv(map[string]string{
		"RPC_URL":      "https://rpc.example.com",
		"METRICS_ADDR": "127.0.0.1:9091",
		"LOG_LEVEL":    "debug",
	}))
	if err != nil {
		t.Fatalf("LoadETL: %v", err)
	}
	if cfg.MetricsAddr != "127.0.0.1:9091" || cfg.Log.Level != "debug" {
		t.Errorf("cfg = %+v", cfg)
	}
}

func TestLoadNotibot(t *testing.T) {
	t.Parallel()
	if _, err := LoadNotibot(mapEnv(nil)); err == nil || !strings.Contains(err.Error(), "RPC_URL: required") {
		t.Errorf("missing RPC_URL not reported: %v", err)
	}
	cfg, err := LoadNotibot(mapEnv(map[string]string{
		"RPC_URL":           "https://rpc.example.com",
		"TWITTER_KEYS_FILE": "tk.json",
		"DISCORD_KEYS_FILE": "dk.json",
	}))
	if err != nil {
		t.Fatalf("LoadNotibot: %v", err)
	}
	if cfg.TwitterKeysFile != "tk.json" || cfg.DiscordKeysFile != "dk.json" {
		t.Errorf("cfg = %+v", cfg)
	}
}

func TestLoadImggenMonitor(t *testing.T) {
	t.Parallel()
	_, err := LoadImggenMonitor(mapEnv(nil))
	if err == nil {
		t.Fatal("missing IM_* accepted")
	}
	for _, want := range []string{"IM_REQUEST_URL: required", "IM_IMAGE_URL: required", "IM_VIDEO_URL: required"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("missing %q in:\n%s", want, err)
		}
	}
	cfg, err := LoadImggenMonitor(mapEnv(map[string]string{
		"IM_REQUEST_URL": "http://gen/generate",
		"IM_IMAGE_URL":   "http://assets/images/",
		"IM_VIDEO_URL":   "http://assets/videos/",
	}))
	if err != nil {
		t.Fatalf("LoadImggenMonitor: %v", err)
	}
	if cfg.RequestURL != "http://gen/generate" {
		t.Errorf("cfg = %+v", cfg)
	}
}

func TestLoadRwalkAlarm(t *testing.T) {
	t.Parallel()
	_, err := LoadRwalkAlarm(mapEnv(nil))
	if err == nil {
		t.Fatal("missing WhatsApp config accepted")
	}
	for _, want := range []string{"PHONE_LIST: required", "WHATSAPP_TOKEN: required", "WHATSAPP_PHONE_ID: required"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("missing %q in:\n%s", want, err)
		}
	}
	cfg, err := LoadRwalkAlarm(mapEnv(map[string]string{
		"PHONE_LIST":        "alice:15550001111",
		"WHATSAPP_TOKEN":    "tok",
		"WHATSAPP_PHONE_ID": "123",
	}))
	if err != nil {
		t.Fatalf("LoadRwalkAlarm: %v", err)
	}
	if cfg.PhoneList != "alice:15550001111" {
		t.Errorf("cfg = %+v", cfg)
	}
}

func TestDBStoreConfig(t *testing.T) {
	t.Parallel()
	db := DB{URL: "postgres://u@h/d", User: "u", Password: "p", Database: "d", Host: "h:5432"}
	sc := db.StoreConfig()
	if sc.URL != db.URL || sc.User != db.User || sc.Password != db.Password ||
		sc.Database != db.Database || sc.Host != db.Host {
		t.Errorf("StoreConfig() = %+v, want a field-for-field copy of %+v", sc, db)
	}
}

func TestProblemHelpersToleratePlainErrors(t *testing.T) {
	t.Parallel()
	plain := errors.New("not a load error")
	if hasProblem(plain, "X") {
		t.Error("hasProblem matched a non-LoadError")
	}
	if got := appendProblems(plain, "X: bad"); got != plain { //nolint:errorlint // identity check on purpose
		t.Errorf("appendProblems rewrote a non-LoadError: %v", got)
	}
	if got := appendProblems(nil); got != nil {
		t.Errorf("appendProblems(nil) = %v, want nil", got)
	}
}

func TestServiceConfigsCoverEveryService(t *testing.T) {
	t.Parallel()
	svcs := ServiceConfigs()
	for _, name := range []string{"apiserver", "etl", "notibot", "imggen-monitor", "rwalk-alarm"} {
		cfg, ok := svcs[name]
		if !ok {
			t.Errorf("ServiceConfigs missing %q", name)
			continue
		}
		if len(Vars(cfg)) == 0 {
			t.Errorf("service %q declares no variables", name)
		}
	}
}
