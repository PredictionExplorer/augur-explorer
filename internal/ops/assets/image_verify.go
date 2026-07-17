package assets

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// HTTPClient is the request surface required by image verification.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TokenCountSource returns the number of minted RandomWalk tokens to check.
type TokenCountSource interface {
	MintedTokenCount(ctx context.Context) (int64, error)
}

// TokenHTTPStatus records a non-200, non-403 response.
type TokenHTTPStatus struct {
	TokenID int64
	Status  int
}

// TokenRequestError records a token whose HTTP request could not be completed.
type TokenRequestError struct {
	TokenID int64
	Err     error
}

// ImageVerificationSummary contains deterministic RandomWalk image results.
type ImageVerificationSummary struct {
	Tokens        int64
	Checked       int64
	OK            int64
	Forbidden     []int64
	OtherStatuses []TokenHTTPStatus
	RequestErrors []TokenRequestError
}

// VerifyTokenImagesOptions configures VerifyTokenImages.
type VerifyTokenImagesOptions struct {
	Source  TokenCountSource
	Client  HTTPClient
	BaseURL string
	Logger  Logger
}

// VerifyTokenImages checks each token's black-background image status. It does
// not download response bodies or write temporary image files.
func VerifyTokenImages(ctx context.Context, opts VerifyTokenImagesOptions) (ImageVerificationSummary, error) {
	var summary ImageVerificationSummary
	if err := ctx.Err(); err != nil {
		return summary, err
	}
	if opts.Source == nil {
		return summary, fmt.Errorf("token count source is nil")
	}
	if opts.Client == nil {
		return summary, fmt.Errorf("HTTP client is nil")
	}
	if strings.TrimSpace(opts.BaseURL) == "" {
		return summary, fmt.Errorf("image base URL is empty")
	}
	logger := opts.Logger
	if logger == nil {
		logger = discardLogger{}
	}

	count, err := opts.Source.MintedTokenCount(ctx)
	if err != nil {
		return summary, fmt.Errorf("read RandomWalk token count: %w", err)
	}
	if count < 0 {
		return summary, fmt.Errorf("RandomWalk token count is negative: %d", count)
	}
	summary.Tokens = count
	logger.Printf("num_tokens = %v", count)

	for tokenID := range count {
		if err := ctx.Err(); err != nil {
			return summary, err
		}
		imageURL := TokenImageURL(opts.BaseURL, tokenID)
		status, err := CheckTokenImage(ctx, opts.Client, imageURL)
		if err != nil {
			if ctxErr := ctx.Err(); ctxErr != nil {
				return summary, ctxErr
			}
			logger.Printf("Can't fetch image %v : %v", imageURL, err)
			summary.RequestErrors = append(summary.RequestErrors, TokenRequestError{TokenID: tokenID, Err: err})
			continue
		}
		summary.Checked++
		if status != http.StatusOK {
			logger.Printf("HTTP response was not 'Ok' : %v (url=%v)", status, imageURL)
		}
		switch status {
		case http.StatusOK:
			summary.OK++
		case http.StatusForbidden:
			logger.Printf("Image server returns 403 code for token %v...", tokenID)
			logger.Printf("token %v: FAIL", tokenID)
			summary.Forbidden = append(summary.Forbidden, tokenID)
		default:
			summary.OtherStatuses = append(summary.OtherStatuses, TokenHTTPStatus{
				TokenID: tokenID,
				Status:  status,
			})
		}
	}
	logger.Printf("Process ended, failed tokens: %v", len(summary.Forbidden))
	return summary, nil
}

// TokenImageURL constructs the zero-padded black-background image URL.
func TokenImageURL(baseURL string, tokenID int64) string {
	return fmt.Sprintf("%s/%06d_black.png", strings.TrimRight(baseURL, "/"), tokenID)
}

// CheckTokenImage performs one context-bound status-only image request.
func CheckTokenImage(ctx context.Context, client HTTPClient, imageURL string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp == nil {
		return 0, fmt.Errorf("HTTP client returned a nil response")
	}
	if resp.Body == nil {
		resp.Body = http.NoBody
	}
	defer func() { _ = resp.Body.Close() }()
	return resp.StatusCode, nil
}
