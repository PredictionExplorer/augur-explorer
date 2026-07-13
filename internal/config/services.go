package config

import (
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// DB selects the PostgreSQL connection, shared by every database-backed
// service. DATABASE_URL (12-factor style) wins over the legacy PGSQL_*
// variables when both are set.
type DB struct {
	// URL is a postgres:// connection URL (or libpq DSN). It embeds the
	// password, so it is secret as a whole.
	URL string `env:"DATABASE_URL" secret:"true"`
	// User, Password, Database and Host are the legacy PGSQL_* settings;
	// an empty Host selects the local Unix socket with peer auth.
	User     string `env:"PGSQL_USERNAME"`
	Password string `env:"PGSQL_PASSWORD" secret:"true"`
	Database string `env:"PGSQL_DATABASE"`
	Host     string `env:"PGSQL_HOST"`
}

// StoreConfig bridges the loaded DB settings into the store's Config.
// Attach a Logger before calling store.New to enable query tracing.
func (d DB) StoreConfig() store.Config {
	return store.Config{
		URL:      d.URL,
		User:     d.User,
		Password: d.Password,
		Database: d.Database,
		Host:     d.Host,
	}
}

// APIServer is the configuration of cmd/apiserver.
type APIServer struct {
	// RPCURL is the Ethereum JSON-RPC endpoint for live contract reads.
	RPCURL string `env:"RPC_URL" required:"true" secret:"url"`

	// HTTPPort serves plain HTTP on :port when set. At least one of
	// HTTPPort and HTTPSHostname must be configured.
	HTTPPort string `env:"HTTP_PORT"`
	// HTTPSHostname is the primary TLS bind address (e.g. :443).
	HTTPSHostname string `env:"HTTPS_HOSTNAME"`
	// HTTPSExtraListenAddr starts a second TLS listener with the same
	// routes and certificates (e.g. :1443).
	HTTPSExtraListenAddr string `env:"HTTPS_EXTRA_LISTEN_ADDR"`

	// TLS certificate pairs; empty paths fall back to
	// $HOME/configs/server.{crt,key}. The optional second pair serves
	// another domain via SNI.
	TLSCertFile  string `env:"TLS_CERT_FILE"`
	TLSKeyFile   string `env:"TLS_KEY_FILE"`
	TLSCertFile2 string `env:"TLS_CERT_FILE_2"`
	TLSKeyFile2  string `env:"TLS_KEY_FILE_2"`

	// Module switches: a disabled module registers no routes and the
	// shared /metadata dispatch answers its legacy unavailable envelope.
	EnableCosmicGame bool `env:"ENABLE_ROUTES_COSMICGAME" default:"true"`
	EnableRandomWalk bool `env:"ENABLE_ROUTES_RANDOMWALK" default:"true"`
	EnableFAQ        bool `env:"ENABLE_ROUTES_FAQ" default:"true"`

	// MetricsAddr serves /metrics and pprof on a private listener; empty
	// disables it. Never expose it publicly.
	MetricsAddr string `env:"METRICS_ADDR"`

	// AdminAPIKey guards the moderation endpoints (fail-closed: unset
	// answers 503). RankingAdminKey guards the direct Elo-match route and
	// falls back to AdminAPIKey.
	AdminAPIKey     string `env:"ADMIN_API_KEY" secret:"true"`
	RankingAdminKey string `env:"RANKING_ADMIN_KEY" secret:"true"`

	// RankingVoteChainIDs allowlists the wallet-signed beauty-vote chain
	// ids (comma-separated). Empty allows any chain id.
	RankingVoteChainIDs []int64 `env:"RANKING_VOTE_CHAIN_IDS"`
	// ExploreMaxTokenID bounds the RandomWalk explore/random token id
	// range.
	ExploreMaxTokenID int64 `env:"RANDOMWALK_EXPLORE_MAX_TOKEN_ID" default:"3766"`

	// Static NFT asset serving (routes register only when the roots are
	// set; see cmd/apiserver/static_assets.go).
	NFTAssetsRoot       string `env:"NFT_ASSETS_ROOT"`
	NFTAssetsPublicBase string `env:"NFT_ASSETS_PUBLIC_BASE"`
	NFTAssetsFlatPaths  bool   `env:"NFT_ASSETS_FLAT_PATHS"`
	StaticABIDir        string `env:"STATIC_ABI_DIR"`
	ImageNoCache        bool   `env:"WEBSRV_IMAGE_NO_CACHE"`
	LogImageRequests    bool   `env:"WEBSRV_LOG_IMAGE_REQUESTS"`

	// FAQ bot upstream: FAQUpstreamURL wins, then the legacy alias, then
	// the faq package's http://127.0.0.1:8000 default.
	FAQUpstreamURL       string `env:"AI_BOT_BACKEND_URL"`
	FAQUpstreamURLLegacy string `env:"FAQ_BOT_UPSTREAM_URL"`

	DB  DB
	Log Log
}

// FAQUpstream resolves the configured FAQ upstream URL, preferring
// AI_BOT_BACKEND_URL over the legacy FAQ_BOT_UPSTREAM_URL alias. Empty means
// "use the faq package default".
func (c *APIServer) FAQUpstream() string {
	if c.FAQUpstreamURL != "" {
		return c.FAQUpstreamURL
	}
	return c.FAQUpstreamURLLegacy
}

// LoadAPIServer loads and validates the API server configuration.
func LoadAPIServer(getenv func(string) string) (*APIServer, error) {
	cfg := &APIServer{}
	err := Load(cfg, getenv)
	err = appendProblems(err, cfg.Log.validate()...)
	if cfg.HTTPPort == "" && cfg.HTTPSHostname == "" {
		err = appendProblems(err, "HTTP_PORT: no listeners configured — set HTTP_PORT and/or HTTPS_HOSTNAME")
	}
	if cfg.ExploreMaxTokenID <= 0 && !hasProblem(err, "RANDOMWALK_EXPLORE_MAX_TOKEN_ID") {
		err = appendProblems(err, "RANDOMWALK_EXPLORE_MAX_TOKEN_ID: must be a positive integer")
	}
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// ETL is the configuration shared by cmd/cg-etl and cmd/rw-etl.
type ETL struct {
	// RPCURL is the Ethereum JSON-RPC endpoint the indexer polls.
	RPCURL string `env:"RPC_URL" required:"true" secret:"url"`
	// MetricsAddr serves /metrics and pprof on a private listener; empty
	// disables it. Use a different port per process on shared hosts.
	MetricsAddr string `env:"METRICS_ADDR"`

	DB  DB
	Log Log
}

// LoadETL loads and validates an ETL configuration.
func LoadETL(getenv func(string) string) (*ETL, error) {
	cfg := &ETL{}
	err := Load(cfg, getenv)
	err = appendProblems(err, cfg.Log.validate()...)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Notibot is the configuration of cmd/notibot.
type Notibot struct {
	// RPCURL is the Ethereum JSON-RPC endpoint for contract reads.
	RPCURL string `env:"RPC_URL" required:"true" secret:"url"`
	// TwitterKeysFile and DiscordKeysFile name credential files under
	// $HOME/configs (required by --twitter / --discord respectively).
	TwitterKeysFile string `env:"TWITTER_KEYS_FILE"`
	DiscordKeysFile string `env:"DISCORD_KEYS_FILE"`

	DB  DB
	Log Log
}

// LoadNotibot loads and validates the notification bot configuration.
func LoadNotibot(getenv func(string) string) (*Notibot, error) {
	cfg := &Notibot{}
	err := Load(cfg, getenv)
	err = appendProblems(err, cfg.Log.validate()...)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// ImggenMonitor is the configuration of cmd/imggen-monitor.
type ImggenMonitor struct {
	// RequestURL is the POST endpoint of the artifact generator service;
	// ImageURL/VideoURL are the bases where artifacts are served
	// (<base><id>.png / <base><id>.mp4).
	RequestURL string `env:"IM_REQUEST_URL" required:"true"`
	ImageURL   string `env:"IM_IMAGE_URL" required:"true"`
	VideoURL   string `env:"IM_VIDEO_URL" required:"true"`

	DB  DB
	Log Log
}

// LoadImggenMonitor loads and validates the artifact monitor configuration.
func LoadImggenMonitor(getenv func(string) string) (*ImggenMonitor, error) {
	cfg := &ImggenMonitor{}
	err := Load(cfg, getenv)
	err = appendProblems(err, cfg.Log.validate()...)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// RwalkAlarm is the configuration of cmd/rwalk-alarm.
type RwalkAlarm struct {
	// PhoneList is the comma-separated person:phone list to notify.
	PhoneList string `env:"PHONE_LIST" required:"true"`
	// WhatsAppToken and WhatsAppPhoneID are the WhatsApp Cloud API
	// credentials.
	WhatsAppToken   string `env:"WHATSAPP_TOKEN" required:"true" secret:"true"`
	WhatsAppPhoneID string `env:"WHATSAPP_PHONE_ID" required:"true"`

	Log Log
}

// LoadRwalkAlarm loads and validates the URL watchdog configuration.
func LoadRwalkAlarm(getenv func(string) string) (*RwalkAlarm, error) {
	cfg := &RwalkAlarm{}
	err := Load(cfg, getenv)
	err = appendProblems(err, cfg.Log.validate()...)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// ServiceConfigs returns a zero value of every service configuration struct
// keyed by service name. The .env.example verification test walks these to
// prove the documented environment matches the code.
func ServiceConfigs() map[string]any {
	return map[string]any{
		"apiserver":      &APIServer{},
		"etl":            &ETL{},
		"notibot":        &Notibot{},
		"imggen-monitor": &ImggenMonitor{},
		"rwalk-alarm":    &RwalkAlarm{},
	}
}

// hasProblem reports whether err already carries a problem line for the
// named variable, so follow-up validation does not double-report it.
func hasProblem(err error, varName string) bool {
	le, ok := err.(*LoadError) //nolint:errorlint // Load always returns *LoadError; this is internal plumbing
	if !ok {
		return false
	}
	for _, p := range le.Problems {
		if strings.HasPrefix(p, varName+":") {
			return true
		}
	}
	return false
}

// appendProblems merges extra problem lines into an existing Load error
// (creating one when err is nil and problems exist).
func appendProblems(err error, problems ...string) error {
	if len(problems) == 0 {
		return err
	}
	if err == nil {
		return &LoadError{Problems: problems}
	}
	le, ok := err.(*LoadError) //nolint:errorlint // Load always returns *LoadError; this is internal plumbing
	if !ok {
		return err
	}
	le.Problems = append(le.Problems, problems...)
	return le
}
