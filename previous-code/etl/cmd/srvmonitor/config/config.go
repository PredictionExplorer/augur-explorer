package config

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/PredictionExplorer/augur-explorer/previous-code/etl/cmd/srvmonitor/types"
)

// Config holds all application configuration
type Config struct {
	// RPC Nodes
	RPCNodes []types.RPCConfig
	
	// Databases
	Layer1DBs      []types.DatabaseConfig
	ApplicationDBs []types.DatabaseConfig
	
	// Image Monitoring
	RWalkImage types.ImageServerConfig
	RWalkDB    types.DatabaseConfig
	
	// Web APIs
	WebAPIs []types.WebAPIConfig
	
	// Disk Monitoring
	DiskMonitors []types.DiskConfig
	
	// Official RPC identifiers
	OfficialRPCMainnet    string
	OfficialRPCArbitrum   string
	OfficialRPCSepolia    string
	OfficialRPCSepoliaArb string
	
	// Intervals
	UpdateInterval         int
	UpdateIntervalDiskCmd  int
	ImageCheckInterval     int
	RPCBlockWait           int
	DBBlockWait            int
	
	// Mobile Notifications
	MobileNotif bool
}

// LoadFromEnv loads configuration from environment variables
func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		UpdateInterval:        60,
		UpdateIntervalDiskCmd: 600,
		ImageCheckInterval:    900,
		RPCBlockWait:          60,
		DBBlockWait:           60,
	}
	
	// Load official RPC identifiers
	cfg.OfficialRPCMainnet = os.Getenv("OFFICIAL_RPC_MAINNET")
	cfg.OfficialRPCArbitrum = os.Getenv("OFFICIAL_RPC_ARBITRUM")
	cfg.OfficialRPCSepolia = os.Getenv("OFFICIAL_RPC_SEPOLIA")
	cfg.OfficialRPCSepoliaArb = os.Getenv("OFFICIAL_RPC_SEPOLIA_ARB")
	
	// Load RPC nodes (RPC0 through RPC9)
	for i := 0; i < 10; i++ {
		name := os.Getenv(fmt.Sprintf("RPC%d_NAME", i))
		url := os.Getenv(fmt.Sprintf("RPC%d_URL", i))
		chainID := os.Getenv(fmt.Sprintf("RPC%d_CHAINID", i))
		
		if name != "" && url != "" {
			isOfficial := name == cfg.OfficialRPCMainnet ||
				name == cfg.OfficialRPCArbitrum ||
				name == cfg.OfficialRPCSepolia ||
				name == cfg.OfficialRPCSepoliaArb
			
			cfg.RPCNodes = append(cfg.RPCNodes, types.RPCConfig{
				Name:       name,
				URL:        url,
				ChainID:    chainID,
				IsOfficial: isOfficial,
			})
		}
	}
	
	// Load Layer 1 databases (DB1 through DB4)
	for i := 1; i <= 4; i++ {
		name := os.Getenv(fmt.Sprintf("DB_L1_NAME_SRV%d", i))
		host := os.Getenv(fmt.Sprintf("DB_L1_HOST_SRV%d", i))
		dbname := os.Getenv(fmt.Sprintf("DB_L1_DBNAME_SRV%d", i))
		user := os.Getenv(fmt.Sprintf("DB_L1_USER_SRV%d", i))
		pass := os.Getenv(fmt.Sprintf("DB_L1_PASS_SRV%d", i))
		
		if name != "" && host != "" {
			cfg.Layer1DBs = append(cfg.Layer1DBs, types.DatabaseConfig{
				Name:   name,
				Host:   host,
				DBName: dbname,
				User:   user,
				Pass:   pass,
			})
		}
	}
	
	// Load Application Layer databases
	for i := 1; i <= 4; i++ {
		title := os.Getenv(fmt.Sprintf("APP_STATUS_SRV%d_TITLE", i))
		host := os.Getenv(fmt.Sprintf("APP_STATUS_SRV%d_HOST", i))
		dbname := os.Getenv(fmt.Sprintf("APP_STATUS_SRV%d_DBNAME", i))
		user := os.Getenv(fmt.Sprintf("APP_STATUS_SRV%d_USER", i))
		pass := os.Getenv(fmt.Sprintf("APP_STATUS_SRV%d_PASS", i))
		
		if title != "" && host != "" {
			cfg.ApplicationDBs = append(cfg.ApplicationDBs, types.DatabaseConfig{
				Name:   title,
				Host:   host,
				DBName: dbname,
				User:   user,
				Pass:   pass,
			})
		}
	}
	
	// Load Web APIs
	for i := 1; i <= 4; i++ {
		title := os.Getenv(fmt.Sprintf("SRV%d_WEB_API_NAME", i))
		host := os.Getenv(fmt.Sprintf("SRV%d_WEB_API_HOST", i))
		port := os.Getenv(fmt.Sprintf("SRV%d_WEB_API_PORT", i))
		uri := os.Getenv(fmt.Sprintf("SRV%d_WEB_API_URI", i))
		
		if title != "" && host != "" {
			cfg.WebAPIs = append(cfg.WebAPIs, types.WebAPIConfig{
				Title: title,
				Host:  host,
				Port:  port,
				URI:   uri,
			})
		}
	}
	
	// Load Disk Monitors
	for i := 1; i <= 3; i++ {
		title := os.Getenv(fmt.Sprintf("SSH_CMD_DF_SRV%d_NAME", i))
		user := os.Getenv(fmt.Sprintf("SSH_CMD_DF_SRV%d_USER", i))
		ip := os.Getenv(fmt.Sprintf("SSH_CMD_DF_SRV%d_IP", i))
		devices := os.Getenv(fmt.Sprintf("SSH_CMD_DF_SRV%d_DEVICES", i))
		
		if title != "" && user != "" {
			cfg.DiskMonitors = append(cfg.DiskMonitors, types.DiskConfig{
				Title:      title,
				User:       user,
				IP:         ip,
				DeviceList: devices,
			})
		}
	}
	
	// Load RWalk Image Monitoring
	cfg.RWalkDB = types.DatabaseConfig{
		Name:   os.Getenv("DB_RWLK_NAME_SRV"),
		Host:   os.Getenv("DB_RWLK_HOST_SRV"),
		DBName: os.Getenv("DB_RWLK_DBNAME_SRV"),
		User:   os.Getenv("DB_RWLK_USER_SRV"),
		Pass:   os.Getenv("DB_RWLK_PASS_SRV"),
	}
	
	cfg.RWalkImage = types.ImageServerConfig{
		Name:         "RWalk Thumbnails",
		URL:          "https://nfts.cosmicsignature.com/images/randomwalk",
		ContractAddr: os.Getenv("RWALK_CONTRACT_ADDR"),
		RPCURL:       os.Getenv("RPC1_URL"),
	}
	
	// Load Mobile Notification setting
	mobileNotif := strings.ToLower(os.Getenv("MOBILE_NOTIF"))
	cfg.MobileNotif = (mobileNotif == "yes" || mobileNotif == "true" || mobileNotif == "1")
	
	// Validate critical configuration
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	
	return cfg, nil
}

// Validate checks if required configuration is present
func (c *Config) Validate() error {
	var missing []string
	
	if len(c.RPCNodes) == 0 {
		missing = append(missing, "at least one RPC node (RPC0_NAME, RPC0_URL)")
	}
	
	if len(c.Layer1DBs) == 0 {
		missing = append(missing, "at least one Layer1 database")
	}
	
	if len(missing) > 0 {
		return fmt.Errorf("missing required configuration: %s", strings.Join(missing, ", "))
	}
	
	return nil
}

// GetOfficialRPCName returns the official RPC name for a given chain ID
func (c *Config) GetOfficialRPCName(chainID string) string {
	switch chainID {
	case "1":
		return c.OfficialRPCMainnet
	case "42161":
		return c.OfficialRPCArbitrum
	case "11155111":
		return c.OfficialRPCSepolia
	case "421614":
		return c.OfficialRPCSepoliaArb
	default:
		return ""
	}
}


