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
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/config"
	"github.com/PredictionExplorer/augur-explorer/internal/ethcall"
	"github.com/PredictionExplorer/augur-explorer/internal/notify/rwbot"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// Server-side database time bounds (D22 defense in depth): notibot's queries
// are small notification reads and watermark writes, so PostgreSQL aborting
// anything longer keeps a wedged statement from silencing notifications
// forever.
const (
	dbStatementTimeout = time.Minute
	dbIdleInTxTimeout  = 5 * time.Minute
)

// osExit is stubbed by tests that drive main through its failure arm.
var osExit = os.Exit

// Package-level so main stays re-runnable under tests (the global FlagSet
// rejects duplicate definitions).
var (
	flagTwitter = flag.Bool("twitter", false, "Send messages to Twitter")
	flagDiscord = flag.Bool("discord", false, "Send messages to Discord")
)

func main() {
	// Before flag.Parse: --version must win over flag validation.
	if version.HandleFlag(os.Args[1:], os.Stdout) {
		return
	}
	flag.Parse()
	if !*flagTwitter && !*flagDiscord {
		fmt.Fprintf(os.Stderr, "Please use --twitter or --discord flag\n")
		osExit(1)
		return
	}
	if err := runMain(*flagTwitter, *flagDiscord); err != nil {
		fmt.Fprintf(os.Stderr, "notibot: %v\n", err)
		osExit(1)
	}
}

// runMain owns the signal-scoped context so its deferred cleanup always runs
// before main decides the exit code (os.Exit skips deferred calls).
func runMain(twitterOn, discordOn bool) error {
	ctx, stopSignals := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()
	return run(ctx, os.Getenv, os.Stdout, twitterOn, discordOn)
}

// run wires the notification engine and blocks until ctx is cancelled
// (SIGINT/SIGTERM in production; clean nil) or a fatal database failure
// (error; systemd restarts the bot and it resumes from the persisted
// watermark). Environment access goes through getenv and structured logs go
// to logOut so tests can inject configuration and inspect records.
func run(ctx context.Context, getenv func(string) string, logOut io.Writer, twitterOn, discordOn bool) error {
	cfg, err := config.LoadNotibot(getenv)
	if err != nil {
		return err
	}
	// One structured logger on stdout; journald owns persistence (§8.3 —
	// the legacy $HOME/ae_logs dual-file layout is gone).
	logger := cfg.Log.NewLogger(logOut)
	logger.LogAttrs(ctx, slog.LevelInfo, "build info", version.LogAttrs()...)
	logger.LogAttrs(ctx, slog.LevelInfo, "effective configuration", config.Attrs(cfg)...)

	eclient, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return fmt.Errorf("connecting to ETH node: %w", err)
	}
	logger.Info("connected to ETH node", "rpc_url", config.RedactURL(cfg.RPCURL))

	storeCfg := cfg.DB.StoreConfig()
	storeCfg.Logger = logger.With("component", "db")
	storeCfg.StatementTimeout = dbStatementTimeout
	storeCfg.IdleInTxSessionTimeout = dbIdleInTxTimeout
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

	// The withdrawal-amount view is the bot's only contract read; the
	// bounded read backend gives it a per-call deadline so a wedged RPC
	// node degrades one notification instead of stalling the loop (D22).
	rwalkCtrct, err := rwcontracts.NewRWalk(rwalkAddr, ethcall.NewBoundedReadBackend(eclient, ethcall.DefaultTimeout))
	if err != nil {
		return fmt.Errorf("instantiating RandomWalk contract %s: %w", rwalkAddr.String(), err)
	}

	botCfg := rwbot.Config{
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
		keys, err := readTwitterKeys(getenv("HOME"), cfg.TwitterKeysFile)
		if err != nil {
			return err
		}
		botCfg.Twitter = rwbot.NewTwitterNotifier(keys)
		logger.Info("loaded twitter keys")
	}
	if discordOn {
		keys, err := readDiscordKeys(getenv("HOME"), cfg.DiscordKeysFile)
		if err != nil {
			return err
		}
		client := newDiscordClient(disgord.Config{BotToken: keys.TokenKey})
		botCfg.Discord = newDiscordSink(client, keys, logger)
		logger.Info("loaded discord keys",
			"main_channel", keys.MainChannelID,
			"mint_stats_channel", keys.MintStatsChanID,
			"price_stats_channel", keys.PriceStatsChanID,
			"date_stats_channel", keys.DateStatsChanID,
			"reward_stats_channel", keys.RewardStatsChanID,
		)
	}

	engine, err := rwbot.New(botCfg)
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
var newDiscordClient = disgord.New
