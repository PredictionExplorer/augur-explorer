// Command notibot announces RandomWalk NFT events (mints, marketplace offers
// and purchases, floor-price changes) on Twitter and Discord, and keeps the
// Discord statistics channels (mint count, price, last mint, last reward)
// up to date.
//
// The monitoring loop lives in internal/notify/rwbot; this binary is wiring:
// flags, credential files, the Ethereum/PostgreSQL connections and the
// Discord/ffmpeg adapters.
//
// Configuration comes from environment variables:
//
//	RPC_URL            Ethereum/Arbitrum JSON-RPC endpoint (required)
//	PGSQL_*            PostgreSQL connection (required)
//	TWITTER_KEYS_FILE  JSON credentials file under $HOME/configs (--twitter)
//	DISCORD_KEYS_FILE  JSON credentials file under $HOME/configs (--discord)
//
// Discord bot permission required to update statistics channels:
//
//	Manage Channel, Connect
//
// Permissions for other users, to keep them out of the statistical channels:
//
//	View Channel yes; Manage Channel no; Connect no
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/andersfylling/disgord"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/notify/rwbot"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// defaultLogDir is the legacy operational log directory under $HOME.
const defaultLogDir = "ae_logs"

func main() {
	flagTwitter := flag.Bool("twitter", false, "Send messages to Twitter")
	flagDiscord := flag.Bool("discord", false, "Send messages to Discord")
	flag.Parse()
	if !*flagTwitter && !*flagDiscord {
		fmt.Fprintf(os.Stderr, "Please use --twitter or --discord flag\n")
		os.Exit(1)
	}
	ctx, stopSignals := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()
	if err := run(ctx, *flagTwitter, *flagDiscord); err != nil {
		fmt.Fprintf(os.Stderr, "notibot: %v\n", err)
		os.Exit(1)
	}
}

// run wires the notification engine and blocks until ctx is cancelled
// (SIGINT/SIGTERM in production; clean nil) or a fatal database failure
// (error; systemd restarts the bot and it resumes from the persisted
// watermark).
func run(ctx context.Context, twitterOn, discordOn bool) error {
	logDir := filepath.Join(os.Getenv("HOME"), defaultLogDir)
	if err := os.MkdirAll(logDir, 0o750); err != nil { //nolint:gosec // G703: operator-owned $HOME log directory, same as the legacy bot
		return fmt.Errorf("creating log directory %s: %w", logDir, err)
	}
	infoLog, err := openLogFile(filepath.Join(logDir, "notibot_info.log"))
	if err != nil {
		return err
	}
	defer infoLog.Close() //nolint:errcheck // append-only log handle at exit
	errorLog, err := openLogFile(filepath.Join(logDir, "notibot_error.log"))
	if err != nil {
		return err
	}
	defer errorLog.Close() //nolint:errcheck // append-only log handle at exit
	dbLog, err := openLogFile(filepath.Join(logDir, "notibot_db.log"))
	if err != nil {
		return err
	}
	defer dbLog.Close() //nolint:errcheck // append-only log handle at exit
	logger := slog.New(indexer.NewDualLogHandler(infoLog, errorLog))

	rpcURL := os.Getenv("RPC_URL")
	eclient, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return fmt.Errorf("connecting to ETH node at %q: %w", rpcURL, err)
	}
	logger.Info("connected to ETH node", "rpc_url", rpcURL)

	storeCfg := store.ConfigFromEnv()
	storeCfg.Logger = slog.New(slog.NewTextHandler(dbLog, nil))
	st, err := store.New(ctx, storeCfg)
	if err != nil {
		return fmt.Errorf("connecting to storage: %w\n%s", err, store.ConnectHint(err))
	}
	defer st.Close()
	repo := rwstore.NewRepo(st)

	addrs, err := repo.ContractAddrs(ctx)
	if err != nil {
		return fmt.Errorf("resolving RandomWalk contract addresses: %w", err)
	}
	rwalkAddr := common.HexToAddress(addrs.RandomWalk)
	logger.Info("resolved contracts", "randomwalk", rwalkAddr.String(), "marketplace", addrs.MarketPlace)

	rwalkCtrct, err := rwcontracts.NewRWalk(rwalkAddr, eclient)
	if err != nil {
		return fmt.Errorf("instantiating RandomWalk contract %s: %w", rwalkAddr.String(), err)
	}

	cfg := rwbot.Config{
		Data:       repo,
		RWalkAid:   addrs.RandomWalkAid,
		MarketAid:  addrs.MarketPlaceAid,
		Media:      rwbot.HTTPFetcher{},
		Resample:   resampleVideo,
		Withdrawal: rwbot.ContractWithdrawalReader{Contract: rwalkCtrct},
		SendVideos: true,
		Logger:     logger,
	}

	if twitterOn {
		keys, err := readTwitterKeys()
		if err != nil {
			return err
		}
		cfg.Twitter = rwbot.NewTwitterNotifier(keys)
		logger.Info("loaded twitter keys")
	}
	if discordOn {
		keys, err := readDiscordKeys()
		if err != nil {
			return err
		}
		client := newDiscordClient(disgord.Config{BotToken: keys.TokenKey})
		cfg.Discord = newDiscordSink(client, keys, logger)
		logger.Info("loaded discord keys",
			"main_channel", keys.MainChannelId,
			"mint_stats_channel", keys.MintStatsChanId,
			"price_stats_channel", keys.PriceStatsChanId,
			"date_stats_channel", keys.DateStatsChanId,
			"reward_stats_channel", keys.RewardStatsChanId,
		)
	}

	engine, err := rwbot.New(cfg)
	if err != nil {
		return err
	}
	if err := engine.Run(ctx); err != nil {
		logger.Error("notification engine failed", "err", err)
		return err
	}
	logger.Info("exiting by user request")
	return nil
}

// newDiscordClient builds the disgord client; a package variable so the
// integration tests can point it at a stub Discord REST server (disgord.New
// probes the bot identity at construction).
var newDiscordClient = func(cfg disgord.Config) *disgord.Client {
	return disgord.New(cfg)
}

// openLogFile opens an append-only log file, creating it if needed.
func openLogFile(path string) (*os.File, error) {
	f, err := os.OpenFile(filepath.Clean(path), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666) //nolint:gosec // operator log path under $HOME, world-readable like the legacy logs
	if err != nil {
		return nil, fmt.Errorf("opening log file %s: %w", path, err)
	}
	return f, nil
}
