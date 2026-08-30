package traits

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// maxBodyBytes caps what one package file may deliver. Trait files are a few
// kilobytes and manifests smaller still; the cap exists so a misconfigured
// or hostile asset host cannot make the ingester allocate without bound.
const maxBodyBytes = 1 << 20

// errBodyTooLarge reports a package file that exceeded maxBodyBytes.
var errBodyTooLarge = errors.New("response body exceeds the size limit")

// fetchStatus is the outcome the ingester branches on. Everything that is
// not one of these three is reported as an error instead.
type fetchStatus int

const (
	// fetchOK carries a fresh body to parse.
	fetchOK fetchStatus = iota
	// fetchNotModified means the stored validator is still current.
	fetchNotModified
	// fetchMissing means the generator has not uploaded this package yet.
	// New mints lag the pipeline by minutes to hours, so this is the
	// expected steady state for recent tokens, not a failure.
	fetchMissing
)

// fetchResult is one conditional GET against the asset host.
type fetchResult struct {
	Status fetchStatus
	Body   []byte
	ETag   string
}

// fetcher performs the conditional GETs the ingest loop needs.
type fetcher struct {
	client *http.Client
	base   string
}

// newFetcher validates the asset host base URL and returns a fetcher for it.
// The base is the host root (for example https://nfts.cosmicsignature.com),
// not the /images mount: the trait contract files are published at /traits
// and /asset-manifests alongside it.
func newFetcher(client *http.Client, base string) (*fetcher, error) {
	trimmed := strings.TrimRight(strings.TrimSpace(base), "/")
	if trimmed == "" {
		return nil, errors.New("empty asset host base URL")
	}
	parsed, err := url.Parse(trimmed)
	if err != nil {
		return nil, fmt.Errorf("parsing asset host base URL %q: %w", base, err)
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return nil, fmt.Errorf("asset host base URL %q must be http or https", base)
	}
	if parsed.Host == "" {
		return nil, fmt.Errorf("asset host base URL %q has no host", base)
	}
	return &fetcher{client: client, base: trimmed}, nil
}

// traitsURL is the published location of one seed's trait contract.
func (f *fetcher) traitsURL(seed string) string {
	return f.base + "/traits/" + seed + ".json"
}

// manifestURL is the published location of one seed's asset manifest.
func (f *fetcher) manifestURL(seed string) string {
	return f.base + "/asset-manifests/" + seed + ".json"
}

// get performs a conditional GET, returning the classified outcome. A
// non-nil error means the request could not be completed or answered an
// unexpected status; 404 and 304 are outcomes, not errors.
func (f *fetcher) get(ctx context.Context, target, etag string) (fetchResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return fetchResult{}, fmt.Errorf("building request for %s: %w", target, err)
	}
	req.Header.Set("Accept", "application/json")
	if etag != "" {
		req.Header.Set("If-None-Match", etag)
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return fetchResult{}, fmt.Errorf("fetching %s: %w", target, err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path

	switch resp.StatusCode {
	case http.StatusNotModified:
		// 304 carries no body; the caller keeps the row it already has.
		return fetchResult{Status: fetchNotModified, ETag: resp.Header.Get("ETag")}, nil
	case http.StatusNotFound, http.StatusGone:
		return fetchResult{Status: fetchMissing}, nil
	case http.StatusOK:
		body, err := readCapped(resp.Body)
		if err != nil {
			return fetchResult{}, fmt.Errorf("reading %s: %w", target, err)
		}
		return fetchResult{Status: fetchOK, Body: body, ETag: resp.Header.Get("ETag")}, nil
	default:
		return fetchResult{}, fmt.Errorf("fetching %s: unexpected status %s", target, resp.Status)
	}
}

// readCapped reads at most maxBodyBytes, reporting oversize bodies rather
// than silently truncating them into a parse error.
func readCapped(r io.Reader) ([]byte, error) {
	body, err := io.ReadAll(io.LimitReader(r, maxBodyBytes+1))
	if err != nil {
		return nil, err
	}
	if len(body) > maxBodyBytes {
		return nil, errBodyTooLarge
	}
	return body, nil
}
