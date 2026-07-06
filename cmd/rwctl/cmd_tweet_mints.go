package main

import (
	"github.com/spf13/cobra"
)

// newTweetMintsCmd builds the tweet-mints subcommand (legacy tweet_mints
// tool). The legacy tweet_mints.go was an identical copy of notif_bot.go, so
// both subcommands share the same implementation (see cmd_notify_bot.go).
func newTweetMintsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tweet-mints",
		Short: "Monitor new mint events and announce them on Twitter",
		Long: "Monitors new mint events and makes announcements to the Twitter account.\n" +
			"Identical to notify-bot; both replace a pair of identical legacy tools.\n\n" +
			notifyBotEnvHelp,
		Args: cobra.NoArgs,
		RunE: runNotifyBot,
	}
}

func init() { register(newTweetMintsCmd()) }
