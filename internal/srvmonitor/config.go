package srvmonitor

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/policy"
)

// Intervals bundles every polling period and in-check wait the monitors use.
// LoadFromEnv fills it with the production defaults; tests shrink it to
// milliseconds.
type Intervals struct {
	// RPC is the pause between RPC monitor check cycles.
	RPC time.Duration
	// RPCBlockWait separates the two head-block reads of one RPC check.
	RPCBlockWait time.Duration
	// DB is the pause between database/application/web-API check cycles.
	DB time.Duration
	// DBBlockWait separates the two block-number reads of one database check.
	DBBlockWait time.Duration
	// EventTable is the pause between event-table check cycles.
	EventTable time.Duration
	// EventTableWait separates the two column reads of one event-table check.
	EventTableWait time.Duration
	// Disk is the pause between disk-usage check cycles.
	Disk time.Duration
	// Image is the pause between image-server check cycles.
	Image time.Duration
	// SSL is the pause between certificate expiry check cycles.
	SSL time.Duration
	// Anomaly is the pause between anomaly-file fetches.
	Anomaly time.Duration
}

// DefaultIntervals returns the production polling periods.
func DefaultIntervals() Intervals {
	return Intervals{
		RPC:            60 * time.Second,
		RPCBlockWait:   60 * time.Second,
		DB:             60 * time.Second,
		DBBlockWait:    60 * time.Second,
		EventTable:     60 * time.Second,
		EventTableWait: 120 * time.Second,
		Disk:           600 * time.Second,
		Image:          900 * time.Second,
		SSL:            3600 * time.Second,
		Anomaly:        300 * time.Second,
	}
}

// Config holds all application configuration.
type Config struct {
	// RPC Nodes
	RPCNodes []RPCConfig

	// Databases
	Layer1DBs      []DatabaseConfig
	ApplicationDBs []DatabaseConfig
	EventTableDBs  []EventTableConfig

	// Image Monitoring
	RWalkImage ImageServerConfig
	RWalkDB    DatabaseConfig

	// Web APIs
	WebAPIs []WebAPIConfig

	// Disk Monitoring
	DiskMonitors []DiskConfig

	// SSL Certificates
	SSLCerts []SSLCertConfig

	// WebSrv anomaly monitoring (error-log anomalies fetched via scp)
	Anomaly AnomalyConfig

	// Official RPC identifiers
	OfficialRPCMainnet    string
	OfficialRPCArbitrum   string
	OfficialRPCSepolia    string
	OfficialRPCSepoliaArb string

	// Intervals holds every polling period; monitors receive these values
	// instead of hardcoding their own.
	Intervals Intervals

	// Mobile Notifications
	MobileNotif bool
}

// LoadFromEnv loads configuration through getenv (pass os.Getenv in
// production). A getenv seam instead of direct os.Getenv reads keeps the
// loader testable without mutating process state.
func LoadFromEnv(getenv func(string) string) (*Config, error) {
	cfg := &Config{
		Intervals: DefaultIntervals(),
	}
	var loadProblems []string

	// Load official RPC identifiers
	cfg.OfficialRPCMainnet = getenv("OFFICIAL_RPC_MAINNET")
	cfg.OfficialRPCArbitrum = getenv("OFFICIAL_RPC_ARBITRUM")
	cfg.OfficialRPCSepolia = getenv("OFFICIAL_RPC_SEPOLIA")
	cfg.OfficialRPCSepoliaArb = getenv("OFFICIAL_RPC_SEPOLIA_ARB")

	// Load RPC nodes (RPC0 through RPC9)
	for i := range 10 {
		name := getenv(fmt.Sprintf("RPC%d_NAME", i))
		url := getenv(fmt.Sprintf("RPC%d_URL", i))
		chainID := getenv(fmt.Sprintf("RPC%d_CHAINID", i))

		if name != "" && url != "" {
			isOfficial := name == cfg.OfficialRPCMainnet ||
				name == cfg.OfficialRPCArbitrum ||
				name == cfg.OfficialRPCSepolia ||
				name == cfg.OfficialRPCSepoliaArb

			cfg.RPCNodes = append(cfg.RPCNodes, RPCConfig{
				Name:       name,
				URL:        url,
				ChainID:    chainID,
				IsOfficial: isOfficial,
			})
		}
	}

	// Load Layer 1 databases (DB1 through DB4)
	for i := 1; i <= 4; i++ {
		name := getenv(fmt.Sprintf("DB_L1_NAME_SRV%d", i))
		host := getenv(fmt.Sprintf("DB_L1_HOST_SRV%d", i))
		dbname := getenv(fmt.Sprintf("DB_L1_DBNAME_SRV%d", i))
		user := getenv(fmt.Sprintf("DB_L1_USER_SRV%d", i))
		pass := getenv(fmt.Sprintf("DB_L1_PASS_SRV%d", i))

		if name != "" && host != "" {
			cfg.Layer1DBs = append(cfg.Layer1DBs, DatabaseConfig{
				Name:   name,
				Host:   host,
				DBName: dbname,
				User:   user,
				Pass:   pass,
			})
		}
	}

	// Load Event Table databases (DB_L1EVT1 through DB_L1EVT6)
	for i := 1; i <= 6; i++ {
		name := getenv(fmt.Sprintf("DB_L1EVT%d_NAME", i))
		host := getenv(fmt.Sprintf("DB_L1EVT%d_HOST", i))
		dbname := getenv(fmt.Sprintf("DB_L1EVT%d_DBNAME", i))
		user := getenv(fmt.Sprintf("DB_L1EVT%d_USER", i))
		pass := getenv(fmt.Sprintf("DB_L1EVT%d_PASS", i))
		tableName := getenv(fmt.Sprintf("DB_L1EVT%d_TABLE", i))
		columnName := getenv(fmt.Sprintf("DB_L1EVT%d_COLUMN", i))

		// Default to cg_proc_status if not specified
		if tableName == "" {
			tableName = "cg_proc_status"
		}
		// Default to last_evt_id if not specified
		if columnName == "" {
			columnName = "last_evt_id"
		}

		if name != "" && host != "" {
			cfg.EventTableDBs = append(cfg.EventTableDBs, EventTableConfig{
				DatabaseConfig: DatabaseConfig{
					Name:   name,
					Host:   host,
					DBName: dbname,
					User:   user,
					Pass:   pass,
				},
				TableName:  tableName,
				ColumnName: columnName,
			})
		}
	}

	// Load Application Layer databases
	for i := 1; i <= 4; i++ {
		title := getenv(fmt.Sprintf("APP_STATUS_SRV%d_TITLE", i))
		host := getenv(fmt.Sprintf("APP_STATUS_SRV%d_HOST", i))
		dbname := getenv(fmt.Sprintf("APP_STATUS_SRV%d_DBNAME", i))
		user := getenv(fmt.Sprintf("APP_STATUS_SRV%d_USER", i))
		pass := getenv(fmt.Sprintf("APP_STATUS_SRV%d_PASS", i))

		if title != "" && host != "" {
			cfg.ApplicationDBs = append(cfg.ApplicationDBs, DatabaseConfig{
				Name:   title,
				Host:   host,
				DBName: dbname,
				User:   user,
				Pass:   pass,
			})
		}
	}

	// Load Web APIs
	for i := 1; i <= 6; i++ {
		title := getenv(fmt.Sprintf("SRV%d_WEB_API_NAME", i))
		host := getenv(fmt.Sprintf("SRV%d_WEB_API_HOST", i))
		port := getenv(fmt.Sprintf("SRV%d_WEB_API_PORT", i))
		uri := getenv(fmt.Sprintf("SRV%d_WEB_API_URI", i))
		publicURL := getenv(fmt.Sprintf("SRV%d_WEB_API_PUBLIC_URL", i))

		if title != "" && host != "" {
			cfg.WebAPIs = append(cfg.WebAPIs, WebAPIConfig{
				Title:     title,
				Host:      host,
				Port:      port,
				URI:       uri,
				PublicURL: publicURL,
			})
		}
	}

	// Load Disk Monitors
	for i := 1; i <= 3; i++ {
		title := getenv(fmt.Sprintf("SSH_CMD_DF_SRV%d_NAME", i))
		user := getenv(fmt.Sprintf("SSH_CMD_DF_SRV%d_USER", i))
		ip := getenv(fmt.Sprintf("SSH_CMD_DF_SRV%d_IP", i))
		devices := getenv(fmt.Sprintf("SSH_CMD_DF_SRV%d_DEVICES", i))

		if title != "" && user != "" {
			cfg.DiskMonitors = append(cfg.DiskMonitors, DiskConfig{
				Title:      title,
				User:       user,
				IP:         ip,
				DeviceList: devices,
			})
		}
	}

	// Load SSL certificate monitors (SSL_CERT1_HOST, SSL_CERT1_PORT, ...)
	for i := 1; i <= 12; i++ {
		host := getenv(fmt.Sprintf("SSL_CERT%d_HOST", i))
		if host == "" {
			continue
		}
		cfg.SSLCerts = append(cfg.SSLCerts, SSLCertConfig{
			Name:       getenv(fmt.Sprintf("SSL_CERT%d_NAME", i)),
			Host:       host,
			Port:       getenv(fmt.Sprintf("SSL_CERT%d_PORT", i)),
			ServerName: getenv(fmt.Sprintf("SSL_CERT%d_SERVERNAME", i)),
		})
	}

	// Load WebSrv anomaly monitoring (optional; enabled when user/host/file set)
	staleAfter, err := parseAnomalyStaleAfter(getenv("ANOMALY_STALE_SECS"))
	if err != nil {
		loadProblems = append(loadProblems, err.Error())
		staleAfter = DefaultAnomalyStaleAfter
	}
	cfg.Anomaly = AnomalyConfig{
		Title:      getenv("ANOMALY_TITLE"),
		User:       getenv("ANOMALY_SSH_USER"),
		Host:       getenv("ANOMALY_SSH_HOST"),
		RemoteFile: getenv("ANOMALY_REMOTE_FILE"),
		StaleAfter: staleAfter,
	}

	// Load RWalk Image Monitoring
	cfg.RWalkDB = DatabaseConfig{
		Name:   getenv("DB_RWLK_NAME_SRV"),
		Host:   getenv("DB_RWLK_HOST_SRV"),
		DBName: getenv("DB_RWLK_DBNAME_SRV"),
		User:   getenv("DB_RWLK_USER_SRV"),
		Pass:   getenv("DB_RWLK_PASS_SRV"),
	}

	cfg.RWalkImage = ImageServerConfig{
		Name:         "RWalk Thumbnails",
		URL:          "https://nfts.cosmicsignature.com/images/randomwalk",
		ContractAddr: getenv("RWALK_CONTRACT_ADDR"),
		RPCURL:       getenv("RPC1_URL"),
	}

	// Load Mobile Notification setting
	mobileNotif := strings.ToLower(getenv("MOBILE_NOTIF"))
	cfg.MobileNotif = (mobileNotif == "yes" || mobileNotif == "true" || mobileNotif == "1")

	// Validate critical configuration and report malformed typed values
	// together with structural problems.
	loadProblems = append(loadProblems, cfg.validationProblems()...)
	if len(loadProblems) > 0 {
		return nil, fmt.Errorf("invalid configuration: %s", strings.Join(loadProblems, ", "))
	}

	return cfg, nil
}

// parseAnomalyStaleAfter parses ANOMALY_STALE_SECS as a strictly positive
// integer count of seconds. Empty selects the production default.
func parseAnomalyStaleAfter(raw string) (time.Duration, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return DefaultAnomalyStaleAfter, nil
	}
	seconds, err := strconv.ParseInt(raw, 10, 64)
	const maxDurationSeconds = (1<<63 - 1) / int64(time.Second)
	if err != nil || seconds <= 0 || seconds > maxDurationSeconds {
		return 0, fmt.Errorf(
			"ANOMALY_STALE_SECS: %q is not a positive integer number of seconds",
			raw,
		)
	}
	return time.Duration(seconds) * time.Second, nil
}

// Validate checks if required configuration is present.
func (c *Config) Validate() error {
	problems := c.validationProblems()
	if len(problems) > 0 {
		return fmt.Errorf("invalid configuration: %s", strings.Join(problems, ", "))
	}
	return nil
}

func (c *Config) validationProblems() []string {
	var problems []string

	if len(c.RPCNodes) == 0 {
		problems = append(problems, "at least one RPC node (RPC0_NAME, RPC0_URL)")
	}

	if len(c.Layer1DBs) == 0 {
		problems = append(problems, "at least one Layer1 database")
	}

	if c.Anomaly.StaleAfter < 0 {
		problems = append(problems, "ANOMALY_STALE_SECS must be positive")
	}

	for _, api := range c.WebAPIs {
		title := api.Title
		if title == "" {
			title = api.Host
		}
		internalURI := api.URI
		if internalURI == "" {
			internalURI = "/"
		}
		parsedInternal, err := url.ParseRequestURI(internalURI)
		if err != nil || !strings.HasPrefix(parsedInternal.Path, "/") {
			problems = append(problems, fmt.Sprintf("web API %q has an invalid internal URI", title))
		} else if policy.V1Deprecated(parsedInternal.Path) {
			problems = append(problems, fmt.Sprintf(
				"web API %q internal URI %q is deprecated v1; use /readyz or /api/v2",
				title,
				parsedInternal.Path,
			))
		}

		if api.PublicURL == "" {
			continue
		}
		parsedPublic, err := url.ParseRequestURI(api.PublicURL)
		if err != nil || parsedPublic.Host == "" ||
			(parsedPublic.Scheme != "http" && parsedPublic.Scheme != "https") {
			problems = append(problems, fmt.Sprintf("web API %q has an invalid public URL", title))
		} else if policy.V1Deprecated(parsedPublic.Path) {
			problems = append(problems, fmt.Sprintf(
				"web API %q public path %q is deprecated v1; use /api/v2",
				title,
				parsedPublic.Path,
			))
		}
	}

	return problems
}
