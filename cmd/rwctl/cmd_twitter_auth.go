package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// readOAuthAppCredentials loads the application (consumer) credentials from a
// JSON config file of the form {"Credentials":{"Token":...,"Secret":...}}.
func readOAuthAppCredentials(path string) (tweets.Credentials, error) {
	var cfg struct {
		Credentials tweets.Credentials
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return cfg.Credentials, fmt.Errorf("can't read oauth config at %v: %w", path, err)
	}
	if err := json.Unmarshal(b, &cfg); err != nil {
		return cfg.Credentials, fmt.Errorf("can't parse oauth config at %v: %w", path, err)
	}
	return cfg.Credentials, nil
}

// newTwitterAuthCmd builds the twitter-auth subcommand. It runs the OAuth
// 1.0a out-of-band (PIN) authorization flow used to provision the token
// credentials stored in the TWITTER_KEYS_FILE config (legacy oauth.go /
// twitteroob tools).
func newTwitterAuthCmd() *cobra.Command {
	var credPath string
	c := &cobra.Command{
		Use:   "twitter-auth",
		Short: "Authorize the Twitter application and obtain token credentials",
		Long: "Runs the OAuth 1.0a out-of-band (PIN based) authorization flow: requests temporary\n" +
			"credentials, prints the authorization URL, reads the verification code from stdin and\n" +
			"prints the token credentials to store in the TWITTER_KEYS_FILE config. Finally sends a\n" +
			"test tweet using the new credentials.",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			appCreds, err := readOAuthAppCredentials(credPath)
			if err != nil {
				return err
			}
			client := tweets.Client{
				TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
				ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
				TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
				Credentials:                   appCreds,
				APIKey:                        appCreds.Token,
			}

			tempCred, err := client.RequestTemporaryCredentials(nil, "oob", nil)
			if err != nil {
				return fmt.Errorf("RequestTemporaryCredentials: %w", err)
			}
			fmt.Printf("Credentials:\n")
			fmt.Printf("\tToken: %v\n", tempCred.Token)
			fmt.Printf("\tSecret: %v\n", tempCred.Secret)
			u := client.AuthorizationURL(tempCred, nil)

			fmt.Printf("1. Go to %s\n2. Authorize the application\n3. Enter verification code:\n", u)

			var code string
			fmt.Scanln(&code)

			tokenCred, _, err := client.RequestToken(nil, tempCred, code)
			if err != nil {
				return fmt.Errorf("RequestToken: %w", err)
			}
			fmt.Printf("Token credentials:\n")
			fmt.Printf("\tToken: %v\n", tokenCred.Token)
			fmt.Printf("\tToken Secret: %v\n", tokenCred.Secret)

			form := url.Values{"status": {"got authorization from account owner"}}
			resp, err := client.Post(nil, tokenCred,
				"https://api.twitter.com/1.1/statuses/update.json", form)
			if err != nil {
				return fmt.Errorf("test tweet failed: %w", err)
			}
			defer resp.Body.Close()
			if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
				return fmt.Errorf("error reading response: %w", err)
			}
			return nil
		},
	}
	c.Flags().StringVar(&credPath, "config", "config.json",
		"path to configuration file containing the application's credentials")
	return c
}

func init() { register(newTwitterAuthCmd()) }
