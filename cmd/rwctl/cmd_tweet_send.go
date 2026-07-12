package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// tweetSendDefaultMessage is the test message the legacy twsend tool posted.
const tweetSendDefaultMessage = "finally works https://api.randomwalknft.com:1443/images/randomwalk/003246_black.png"

// newTweetSendCmd builds the tweet-send subcommand (legacy twsend tool).
func newTweetSendCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tweet-send [api_key] [api_secret] [access_token] [token_secret] [nonce] [message]",
		Short: "Send a tweet on behalf of a user",
		Long: "Sends a tweet on behalf of a user, using raw OAuth 1.0a credentials.\n" +
			"The nonce is a hex number; it is incremented before use. When message is\n" +
			"omitted, the legacy twsend test message is posted.",
		Args: cobra.RangeArgs(5, 6),
		RunE: func(cmd *cobra.Command, args []string) error {
			apiKey := args[0]
			apiSecret := args[1]
			accessToken := args[2]
			tokenSecret := args[3]
			sessionNonce, err := strconv.ParseUint(args[4], 16, 64)
			if err != nil {
				return fmt.Errorf("parsing nonce error: %w", err)
			}
			sessionNonce++
			message := tweetSendDefaultMessage
			if len(args) == 6 {
				message = args[5]
			}

			statusCode, body, err := tweets.SendTweet(apiKey, apiSecret, accessToken, tokenSecret, message, sessionNonce)
			if err != nil {
				return fmt.Errorf("error sending tweet (status %v; body = %v): %w", statusCode, body, err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Successfully sent\n")
			fmt.Fprintf(cmd.OutOrStdout(), "dump body:\n")
			fmt.Fprintln(cmd.OutOrStdout(), body)
			return nil
		},
	}
}

func init() { register(newTweetSendCmd()) }
