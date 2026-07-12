package rwbot

// TwitterNotifier adapts the internal/notify/tweets senders to the Tweeter
// seam: it owns the OAuth nonce sequence, enforces the HTTP 200 contract and
// extracts the tweet id used to thread video replies.

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// sendFunc matches the tweets.SendTweetWith{Image,Video} signatures.
type sendFunc func(apiKey, apiSecret, accessToken, tokenSecret, message string, sessionNonce uint64, media []byte, replyTweet string) (int, string, error)

// TwitterNotifier posts tweets through the vendored Twitter client.
type TwitterNotifier struct {
	keys  tweets.TwitterKeys
	nonce atomic.Uint64

	// Injected in tests; production uses the tweets package senders.
	sendImage sendFunc
	sendVideo sendFunc
}

// NewTwitterNotifier builds a Tweeter for the given account keys. The OAuth
// nonce sequence is seeded from the wall clock, like the legacy bots.
func NewTwitterNotifier(keys tweets.TwitterKeys) *TwitterNotifier {
	t := &TwitterNotifier{
		keys:      keys,
		sendImage: tweets.SendTweetWithImage,
		sendVideo: tweets.SendTweetWithVideo,
	}
	t.nonce.Store(uint64(time.Now().UnixNano())) // wall-clock nonce seed, same as the legacy bots
	return t
}

// TweetWithImage posts msg with the image attached and returns the created
// tweet's id.
func (t *TwitterNotifier) TweetWithImage(_ context.Context, msg string, image []byte) (string, error) {
	status, body, err := t.sendImage(
		t.keys.ApiKey, t.keys.ApiSecret, t.keys.TokenKey, t.keys.TokenSecret,
		msg, t.nonce.Add(1), image, "",
	)
	if err != nil {
		return "", fmt.Errorf("sending tweet (status %d): %w", status, err)
	}
	if status != 200 {
		return "", fmt.Errorf("sending tweet: HTTP status %d: %s", status, body)
	}
	return parseTweetID(body)
}

// TweetWithVideo posts msg with the video attached, threaded under
// replyToID when non-empty.
func (t *TwitterNotifier) TweetWithVideo(_ context.Context, msg string, video []byte, replyToID string) error {
	status, body, err := t.sendVideo(
		t.keys.ApiKey, t.keys.ApiSecret, t.keys.TokenKey, t.keys.TokenSecret,
		msg, t.nonce.Add(1), video, replyToID,
	)
	if err != nil {
		return fmt.Errorf("sending video tweet (status %d): %w", status, err)
	}
	if status != 200 {
		return fmt.Errorf("sending video tweet: HTTP status %d: %s", status, body)
	}
	return nil
}

// parseTweetID extracts the tweet id from a status-update response body.
func parseTweetID(body string) (string, error) {
	var resp tweets.StatusUpdateResponse
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		return "", fmt.Errorf("decoding status update response: %w", err)
	}
	return resp.IdStr, nil
}
