package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// readOAuthAppCredentials loads the application (consumer) credentials from a
// JSON config file of the form {"Credentials":{"Token":...,"Secret":...}}.
func readOAuthAppCredentials(path string) (tweets.Credentials, error) {
	var cfg struct {
		Credentials tweets.Credentials
	}
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return cfg.Credentials, fmt.Errorf("can't read oauth config at %v: %w", path, err)
	}
	if err := json.Unmarshal(b, &cfg); err != nil {
		return cfg.Credentials, fmt.Errorf("can't parse oauth config at %v: %w", path, err)
	}
	return cfg.Credentials, nil
}

// twitterAuthEndpoints carries the OAuth 1.0a endpoints of the authorization
// flow; tests point them at a stub server.
type twitterAuthEndpoints struct {
	requestTokenURI string
	authorizeURI    string
	accessTokenURI  string
	statusUpdateURI string
}

// productionTwitterEndpoints are the real Twitter API endpoints.
func productionTwitterEndpoints() twitterAuthEndpoints {
	return twitterAuthEndpoints{ //nolint:gosec // G101: public API endpoint URLs, not credentials
		requestTokenURI: "https://api.twitter.com/oauth/request_token",
		authorizeURI:    "https://api.twitter.com/oauth/authorize",
		accessTokenURI:  "https://api.twitter.com/oauth/access_token",
		statusUpdateURI: "https://api.twitter.com/1.1/statuses/update.json",
	}
}

// runTwitterAuthFlow executes the OAuth 1.0a out-of-band (PIN) authorization
// flow: request temporary credentials, print the authorization URL, read the
// verification code from in, exchange it for token credentials, print them,
// and post a test tweet with the new credentials.
func runTwitterAuthFlow(appCreds tweets.Credentials, endpoints twitterAuthEndpoints, httpClient *http.Client, in io.Reader, out io.Writer) error {
	client := tweets.Client{
		TemporaryCredentialRequestURI: endpoints.requestTokenURI,
		ResourceOwnerAuthorizationURI: endpoints.authorizeURI,
		TokenRequestURI:               endpoints.accessTokenURI,
		Credentials:                   appCreds,
		APIKey:                        appCreds.Token,
	}

	tempCred, err := client.RequestTemporaryCredentials(httpClient, "oob", nil)
	if err != nil {
		return fmt.Errorf("RequestTemporaryCredentials: %w", err)
	}
	fmt.Fprintf(out, "Credentials:\n")
	fmt.Fprintf(out, "\tToken: %v\n", tempCred.Token)
	fmt.Fprintf(out, "\tSecret: %v\n", tempCred.Secret)
	u := client.AuthorizationURL(tempCred, nil)

	fmt.Fprintf(out, "1. Go to %s\n2. Authorize the application\n3. Enter verification code:\n", u)

	var code string
	if _, err := fmt.Fscanln(in, &code); err != nil {
		return fmt.Errorf("reading verification code: %w", err)
	}

	tokenCred, _, err := client.RequestToken(httpClient, tempCred, code)
	if err != nil {
		return fmt.Errorf("RequestToken: %w", err)
	}
	fmt.Fprintf(out, "Token credentials:\n")
	fmt.Fprintf(out, "\tToken: %v\n", tokenCred.Token)
	fmt.Fprintf(out, "\tToken Secret: %v\n", tokenCred.Secret)

	form := url.Values{"status": {"got authorization from account owner"}}
	resp, err := client.Post(httpClient, tokenCred, endpoints.statusUpdateURI, form)
	if err != nil {
		return fmt.Errorf("test tweet failed: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck // read-only body
	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}
	return nil
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
		RunE: func(cmd *cobra.Command, _ []string) error {
			appCreds, err := readOAuthAppCredentials(credPath)
			if err != nil {
				return err
			}
			return runTwitterAuthFlow(appCreds, productionTwitterEndpoints(), nil, cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	c.Flags().StringVar(&credPath, "config", "config.json",
		"path to configuration file containing the application's credentials")
	return c
}

func init() { register(newTwitterAuthCmd()) }
