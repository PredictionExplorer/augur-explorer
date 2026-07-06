package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// newTweetReplyImageCmd builds the tweet-reply-image subcommand (legacy
// twsend_img_reply tool).
func newTweetReplyImageCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tweet-reply-image [reply_to_id] [media_file] [message]",
		Short: "Send a tweet with attached media as a reply",
		Long: "Sends a tweet with media_file (image or video) attached as reply-to.\n\n" +
			"Environment variables:\n  TWITTER_KEYS_FILE  name of the JSON credentials file under $HOME/configs (required)",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			keys, err := readTwitterKeys()
			if err != nil {
				return err
			}

			replyToID := args[0]
			mediaFilename := args[1]
			message := args[2]
			fmt.Printf("Media file: %v\n", mediaFilename)
			fmt.Printf("Reply to id: %v\n", replyToID)
			mediaData, err := os.ReadFile(mediaFilename)
			if err != nil {
				return fmt.Errorf("can't read media data at %v: %w", mediaFilename, err)
			}

			nonce := uint64(time.Now().UnixNano()) + 1
			statusCode, body, err := tweets.SendTweetWithVideo(
				keys.ApiKey,
				keys.ApiSecret,
				keys.TokenKey,
				keys.TokenSecret,
				message,
				nonce,
				mediaData,
				replyToID,
			)
			if err != nil {
				return fmt.Errorf("error sending tweet with media (status %v; body = %v): %w", statusCode, body, err)
			}
			fmt.Printf("body after send: %v\n", body)
			var statusResp tweets.StatusUpdateResponse
			if err := json.NewDecoder(strings.NewReader(body)).Decode(&statusResp); err != nil {
				return fmt.Errorf("error at decode response: %w", err)
			}
			fmt.Printf("status_resp.Id=%v\n", statusResp.Id)
			fmt.Printf("status_resp.IdStr=%v\n", statusResp.IdStr)
			return nil
		},
	}
}

func init() { register(newTweetReplyImageCmd()) }
