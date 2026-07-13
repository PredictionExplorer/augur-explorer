// Package imggen implements the cmd/imggen-monitor engine: it verifies that
// every minted Cosmic Signature NFT has its generated image and video
// artifacts available on the artifact server, and optionally asks the
// generator service to (re)create missing ones.
//
// The HTTP client, token source, pacing and output writer are injected, so
// the scan/generate/wait pipeline is testable against httptest servers; the
// binary wires the environment configuration and the production repository.
package imggen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Default pacing of the scan and wait loops.
const (
	// DefaultPollInterval is the pause between presence probes while
	// waiting for artifacts to appear.
	DefaultPollInterval = 10 * time.Second
	// DefaultWaitTimeout bounds one artifact wait. The legacy monitor
	// polled forever, so a wedged generator hung the scan permanently.
	DefaultWaitTimeout = 15 * time.Minute
	// DefaultTokenPacing is the pause between tokens during a scan.
	DefaultTokenPacing = time.Second
)

// HTTPClient is the request surface the client needs.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client checks and requests generation of NFT image/video artifacts.
type Client struct {
	// RequestURL is the POST endpoint of the artifact generator service.
	RequestURL string
	// ImageURL is the base URL where images are served (<base><id>.png).
	ImageURL string
	// VideoURL is the base URL where videos are served (<base><id>.mp4).
	VideoURL string
	// HTTPClient issues the requests.
	HTTPClient HTTPClient
	// PollInterval is the pause between presence probes; zero selects
	// DefaultPollInterval.
	PollInterval time.Duration
	// WaitTimeout bounds one WaitUntilPresent call; zero selects
	// DefaultWaitTimeout.
	WaitTimeout time.Duration
}

// Generate asks the generator service to create image/video artifacts for a
// token.
func (c *Client) Generate(ctx context.Context, tokenID int64, seed string) error {
	// A fixed-shape map of an int64 and a string cannot fail to encode.
	payload, _ := json.Marshal(map[string]interface{}{"token_id": tokenID, "seed": seed})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.RequestURL, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("building generation request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("submitting generation request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path
	if resp.StatusCode >= 300 {
		return fmt.Errorf("generator service returned %s", resp.Status)
	}
	return nil
}

// Exists reports whether both the image and the video artifact are reachable.
func (c *Client) Exists(ctx context.Context, tokenID int64) (bool, error) {
	for _, url := range []string{
		fmt.Sprintf("%v%06d.png", c.ImageURL, tokenID),
		fmt.Sprintf("%v%06d.mp4", c.VideoURL, tokenID),
	} {
		req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
		if err != nil {
			return false, fmt.Errorf("building HEAD %s: %w", url, err)
		}
		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			return false, fmt.Errorf("HEAD %s: %w", url, err)
		}
		_ = resp.Body.Close() // HEAD responses carry no body
		if resp.StatusCode != http.StatusOK {
			return false, nil
		}
	}
	return true, nil
}

// WaitUntilPresent polls until the token's artifacts appear, an error
// occurs, the wait times out or ctx is cancelled. Progress dots go to out.
func (c *Client) WaitUntilPresent(ctx context.Context, tokenID int64, out io.Writer) error {
	pollInterval := c.PollInterval
	if pollInterval <= 0 {
		pollInterval = DefaultPollInterval
	}
	waitTimeout := c.WaitTimeout
	if waitTimeout <= 0 {
		waitTimeout = DefaultWaitTimeout
	}
	deadline := time.Now().Add(waitTimeout)

	for {
		exists, err := c.Exists(ctx, tokenID)
		if err != nil {
			return err
		}
		if exists {
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("artifacts for token %d still missing after %v", tokenID, waitTimeout)
		}
		fmt.Fprint(out, ".")
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(pollInterval):
		}
	}
}

// Token is one minted token to check.
type Token struct {
	ID   int64
	Seed string
}

// TokenSource lists the minted tokens to scan.
type TokenSource interface {
	Tokens(ctx context.Context) ([]Token, error)
}

// ScanOptions configures Scan.
type ScanOptions struct {
	Source TokenSource
	Client *Client
	// Regenerate requests generation of missing artifacts and waits for
	// them to appear.
	Regenerate bool
	// Out receives the per-token report lines.
	Out io.Writer
	// TokenPacing is the pause between tokens; zero selects
	// DefaultTokenPacing.
	TokenPacing time.Duration
}

// Scan checks artifact presence for every token; with Regenerate it asks the
// generator to fill gaps. Per-token failures are reported and the scan moves
// on; only a token-list failure or cancellation aborts it.
func Scan(ctx context.Context, opts ScanOptions) error {
	if opts.Source == nil {
		return fmt.Errorf("token source is nil")
	}
	if opts.Client == nil {
		return fmt.Errorf("client is nil")
	}
	pacing := opts.TokenPacing
	if pacing <= 0 {
		pacing = DefaultTokenPacing
	}

	if opts.Regenerate {
		fmt.Fprintln(opts.Out, "Regenerating missing images/videos")
	} else {
		fmt.Fprintln(opts.Out, "Checking image/video presence")
	}

	tokens, err := opts.Source.Tokens(ctx)
	if err != nil {
		return fmt.Errorf("failed to list tokens: %w", err)
	}
	for _, tok := range tokens {
		if err := ctx.Err(); err != nil {
			return err
		}
		fmt.Fprintf(opts.Out, "token id = %v    ", tok.ID)
		exists, err := opts.Client.Exists(ctx, tok.ID)
		switch {
		case err != nil:
			fmt.Fprintf(opts.Out, "error: %v\n", err)
		case exists:
			fmt.Fprintln(opts.Out, "image/video present")
		case !opts.Regenerate:
			fmt.Fprintln(opts.Out, "doesn't exist")
		default:
			fmt.Fprint(opts.Out, " regenerating ...")
			if err := opts.Client.Generate(ctx, tok.ID, tok.Seed); err != nil {
				fmt.Fprintf(opts.Out, " request failed: %v\n", err)
				continue
			}
			if err := opts.Client.WaitUntilPresent(ctx, tok.ID, opts.Out); err != nil {
				fmt.Fprintf(opts.Out, " aborting due to error: %v\n", err)
				continue
			}
			fmt.Fprintln(opts.Out, " done.")
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(pacing):
		}
	}
	return nil
}
