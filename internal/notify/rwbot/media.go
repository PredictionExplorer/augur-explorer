package rwbot

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// maxMediaBytes bounds a single media download (the generated videos are a
// few MB; 256MB leaves ample headroom while preventing an unbounded read).
const maxMediaBytes = 256 << 20

// HTTPFetcher is the production Fetcher: a plain GET whose body is read into
// memory (the legacy temp-file round trip is gone; only the ffmpeg adapter
// still needs files, and it owns them).
type HTTPFetcher struct {
	// Client defaults to a client with a 5-minute timeout (video files).
	Client *http.Client
}

// Fetch downloads url and returns the HTTP status plus the body for 200
// responses.
func (f HTTPFetcher) Fetch(ctx context.Context, url string) (int, []byte, error) {
	client := f.Client
	if client == nil {
		client = &http.Client{Timeout: 5 * time.Minute}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, fmt.Errorf("building request for %s: %w", url, err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("fetching %s: %w", url, err)
	}
	defer resp.Body.Close() //nolint:errcheck // read-only body
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil, nil
	}
	data, err := io.ReadAll(io.LimitReader(resp.Body, maxMediaBytes))
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("reading body of %s: %w", url, err)
	}
	return resp.StatusCode, data, nil
}
