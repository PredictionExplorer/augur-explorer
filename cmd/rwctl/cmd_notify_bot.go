package main

// The notify-bot and tweet-mints subcommands share one implementation: the
// Twitter-only wiring of the internal/notify/rwbot engine (the same engine
// cmd/notibot runs with Discord and videos enabled). The legacy in-memory
// timestamp watermark is gone: the engine resumes from the persisted
// rw_messaging_status row exactly like notibot.

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/notify/rwbot"
	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// readTwitterKeys loads the Twitter API credentials from
// $HOME/configs/$TWITTER_KEYS_FILE.
func readTwitterKeys() (tweets.TwitterKeys, error) {
	var keys tweets.TwitterKeys
	fileName := filepath.Join(os.Getenv("HOME"), "configs", os.Getenv("TWITTER_KEYS_FILE"))
	b, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return keys, fmt.Errorf("can't read configuration file with twitter account keys in %v: %w", fileName, err)
	}
	if err := json.Unmarshal(b, &keys); err != nil {
		return keys, fmt.Errorf("can't parse twitter account keys in %v: %w", fileName, err)
	}
	return keys, nil
}

// runNotifyBot wires database, RPC and Twitter credentials into the shared
// notification engine and runs it until interrupted.
func runNotifyBot(cmd *cobra.Command, _ []string) error {
	ctx, stopSignals := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	keys, err := readTwitterKeys()
	if err != nil {
		return err
	}

	eclient, err := dialEthClient()
	if err != nil {
		return err
	}

	repo, _, err := connectRWStorage(newInfoLogger())
	if err != nil {
		return err
	}

	addrs, err := repo.ContractAddrs(ctx)
	if err != nil {
		return fmt.Errorf("resolving contract addresses: %w", err)
	}
	rwalkAddr := common.HexToAddress(addrs.RandomWalk)
	logger.Info("resolved contracts", "randomwalk", rwalkAddr.String(), "marketplace", addrs.MarketPlace)

	rwalkCtrct, err := rwcontracts.NewRWalk(rwalkAddr, eclient)
	if err != nil {
		return fmt.Errorf("can't instantiate RandomWalk contract %v: %w", rwalkAddr.String(), err)
	}

	engine, err := rwbot.New(rwbot.Config{
		Data:       repo,
		RWalkAid:   addrs.RandomWalkAid,
		MarketAid:  addrs.MarketPlaceAid,
		Twitter:    rwbot.NewTwitterNotifier(keys),
		Media:      rwbot.HTTPFetcher{},
		Withdrawal: rwbot.ContractWithdrawalReader{Contract: rwalkCtrct},
		SendVideos: false, // the rwctl bot has always been image-only
		Logger:     logger,
	})
	if err != nil {
		return err
	}
	return engine.Run(ctx)
}

// notifyBotEnvHelp documents the environment variables of the notification bot.
const notifyBotEnvHelp = "Environment variables:\n" +
	"  RPC_URL             Ethereum RPC endpoint (required)\n" +
	"  TWITTER_KEYS_FILE   name of the JSON credentials file under $HOME/configs (required)\n" +
	"  PGSQL_*             PostgreSQL connection (required)"

// newNotifyBotCmd builds the notify-bot subcommand (legacy notif_bot tool).
func newNotifyBotCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "notify-bot",
		Short: "Monitor RandomWalk events and announce them on Twitter",
		Long: "Monitors new mint/offer/purchase events and floor price changes and announces them to the Twitter account.\n" +
			"Resumes from the persisted rw_messaging_status watermark (shared with notibot).\n\n" +
			notifyBotEnvHelp,
		Args: cobra.NoArgs,
		RunE: runNotifyBot,
	}
}

func init() { register(newNotifyBotCmd()) }
