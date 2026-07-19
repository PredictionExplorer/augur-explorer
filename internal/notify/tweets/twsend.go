package tweets

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// statusUpdateURL and mediaUploadURL are package variables (not constants) so
// the httptest suite can point the senders at a stub server; production never
// changes them.
var (
	statusUpdateURL = "https://api.twitter.com/1.1/statuses/update.json"
	mediaUploadURL  = "https://upload.twitter.com/1.1/media/upload.json"
)

// videoStatusPollInterval is the wait between STATUS polls while Twitter
// processes an uploaded video. A variable so tests do not sleep for real.
var videoStatusPollInterval = 2 * time.Second

// TwitterKeys carries the four credentials the send helpers need, as loaded
// from the operator's TWITTER_KEYS_FILE config. The tags pin the legacy key
// spellings operators already have on disk.
type TwitterKeys struct {
	APIKey      string `json:"ApiKey"`
	APISecret   string `json:"ApiSecret"`
	TokenKey    string
	TokenSecret string
}

// MediaImage is the image block of a media-upload response.
type MediaImage struct {
	ImageType string `json:"image_type"`
	W         int64  `json:"w"`
	H         int64  `json:"h"`
}

// ImageResponse is the media-upload endpoint's response (also used for the
// INIT/FINALIZE stages of video uploads, which share the same shape).
type ImageResponse struct {
	MediaID          int64      `json:"media_id"`
	MediaIDString    string     `json:"media_id_string"`
	MediaKey         string     `json:"media_key"`
	Size             int64      `json:"size"`
	ExpiresAfterSecs int64      `json:"expires_after_secs"`
	Image            MediaImage `json:"image"`
}

// StatusUpdateResponse is the statuses/update response subset the callers
// read back (the tweet id, used for reply threading).
type StatusUpdateResponse struct {
	ID    int64  `json:"id"`
	IDStr string `json:"id_str"`
}

// ProcessingInfo reports Twitter's server-side video processing progress.
type ProcessingInfo struct {
	State           string `json:"state"`
	CheckAfterSecs  int64  `json:"check_after_secs"`
	ProgressPercent int64  `json:"progress_percent"`
}

// VideoUploadStatus is the STATUS response polled after a video FINALIZE.
type VideoUploadStatus struct {
	MediaID          int64          `json:"media_id"`
	MediaIDString    string         `json:"media_id_string"`
	ExpiresAfterSecs int64          `json:"expires_after_secs"`
	ProcessingInfo   ProcessingInfo `json:"processing_info"`
}

func decodeResponse(resp *http.Response, data any) error {
	p, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	// note: we did ReadAll() because we need to have an option to dump the received body
	bodyStream := bytes.NewReader(p)
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusAccepted) && (resp.StatusCode != http.StatusNoContent) {
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, string(p))
	}
	return json.NewDecoder(bodyStream).Decode(data)
}

func formatNonce(nonce uint64) string {
	return strconv.FormatUint(nonce, 16)
}

// readBodyAndStatusError reads the response body and builds the (status code,
// body, error) triple returned by the Send* helpers. A non-nil error is
// returned when the response status differs from http.StatusOK.
func readBodyAndStatusError(resp *http.Response) (int, string, error) {
	var errOut error
	if resp.StatusCode != http.StatusOK {
		errOut = fmt.Errorf("error at Twitter occurred: status = %v", resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil && errOut == nil {
		errOut = fmt.Errorf("error reading response body: %w", err)
	}
	return resp.StatusCode, string(b), errOut
}

func newTwitterClient(apiKey, apiSecret, accessToken string, sessionNonce uint64) Client {
	var client Client
	client.Credentials.Token = apiKey     // user-generated token
	client.Credentials.Secret = apiSecret // app secret
	client.APIKey = apiKey
	client.ClientToken = accessToken
	client.Nonce = formatNonce(sessionNonce)
	return client
}

// SendTweet posts a plain status update and returns the HTTP status code,
// the raw response body and an error when the status is not 200.
func SendTweet(apiKey, apiSecret, accessToken, tokenSecret, message string, sessionNonce uint64) (int, string, error) {
	client := newTwitterClient(apiKey, apiSecret, accessToken, sessionNonce)

	tokenCredentials := Credentials{Token: accessToken, Secret: tokenSecret}

	form := url.Values{"status": {message}}
	resp, err := client.Post(nil, &tokenCredentials, statusUpdateURL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path

	return readBodyAndStatusError(resp)
}

// SendTweetWithImage uploads imageData to the media endpoint and posts a
// status update referencing it (optionally as a reply to replyTweet). Returns
// the HTTP status code, raw response body and an error when the final status
// is not 200.
func SendTweetWithImage(apiKey, apiSecret, accessToken, tokenSecret, message string, sessionNonce uint64, imageData []byte, replyTweet string) (int, string, error) {
	client := newTwitterClient(apiKey, apiSecret, accessToken, sessionNonce)

	encodedImageData := base64.StdEncoding.EncodeToString(imageData)
	form := url.Values{
		"media_data": {encodedImageData},
	}
	tokenCredentials := Credentials{Token: accessToken, Secret: tokenSecret}

	resp, err := client.Post(nil, &tokenCredentials, mediaUploadURL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	var imageResponse ImageResponse
	err = decodeResponse(resp, &imageResponse)
	_ = resp.Body.Close() // best-effort: body was fully read by decodeResponse
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("error decoding media upload response: %w", err)
	}

	form = url.Values{
		"status":                {message},
		"media_ids":             {imageResponse.MediaIDString},
		"in_reply_to_status_id": {replyTweet},
	}
	resp, err = client.Post(nil, &tokenCredentials, statusUpdateURL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path

	return readBodyAndStatusError(resp)
}

// SendTweetWithVideo runs Twitter's chunked video upload protocol
// (INIT/APPEND/FINALIZE, then STATUS polling until processing completes) and
// posts a status update referencing the uploaded video. Returns the HTTP
// status code, raw response body and an error when the final status is not
// 200; the poll loop aborts after eight attempts.
func SendTweetWithVideo(apiKey, apiSecret, accessToken, tokenSecret, message string, sessionNonce uint64, videoData []byte, replyTweet string) (int, string, error) {
	client := newTwitterClient(apiKey, apiSecret, accessToken, sessionNonce)
	encodedVideoData := base64.StdEncoding.EncodeToString(videoData)
	form := url.Values{
		"command":        {"INIT"},
		"total_bytes":    {strconv.Itoa(len(videoData))},
		"media_category": {"tweet_video"},
		"media_type":     {"video/mp4"},
	}
	tokenCredentials := Credentials{Token: accessToken, Secret: tokenSecret}

	//---------- INIT
	respInit, err := client.Post(nil, &tokenCredentials, mediaUploadURL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = respInit.Body.Close() }() // best-effort close on read path
	if respInit.StatusCode != http.StatusAccepted {
		b, readErr := io.ReadAll(respInit.Body)
		if readErr != nil {
			return respInit.StatusCode, "", fmt.Errorf("error reading INIT response body: %w", readErr)
		}
		return respInit.StatusCode, string(b), fmt.Errorf("twitter INIT request returned status %v", respInit.Status)
	}
	var imgResponseInit ImageResponse
	err = decodeResponse(respInit, &imgResponseInit)
	if err != nil {
		return 0, "", err
	}
	uploadID := imgResponseInit.MediaIDString

	//------------ APPEND
	form = url.Values{
		"command":       {"APPEND"},
		"media_id":      {uploadID},
		"media_data":    {encodedVideoData + "\r\n"},
		"segment_index": {"0"},
	}
	client.Header = make(map[string][]string)
	client.Header.Set("Content-Type", "application/octet-stream")
	client.Header.Set("Content-Transfer-Encoding", "base64")
	respAppend, err := client.Post(nil, &tokenCredentials, mediaUploadURL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = respAppend.Body.Close() }() // best-effort close on read path
	if (respAppend.StatusCode != http.StatusOK) && (respAppend.StatusCode != http.StatusNoContent) {
		b, readErr := io.ReadAll(respAppend.Body)
		if readErr != nil {
			return respAppend.StatusCode, "", fmt.Errorf("error reading APPEND response body: %w", readErr)
		}
		return respAppend.StatusCode, string(b), fmt.Errorf("twitter APPEND request returned status %v", respAppend.Status)
	}
	if respAppend.StatusCode != http.StatusNoContent {
		var imgResponseAppend ImageResponse
		err = decodeResponse(respAppend, &imgResponseAppend)
		if err != nil {
			return 0, "", err
		}
	}
	client.Header.Del("Content-Type")
	client.Header.Del("Content-Transfer-Encoding")
	//----------- FINALIZE
	form = url.Values{
		"command":  {"FINALIZE"},
		"media_id": {uploadID},
	}
	respFinalize, err := client.Post(nil, &tokenCredentials, mediaUploadURL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = respFinalize.Body.Close() }() // best-effort close on read path
	if (respFinalize.StatusCode != http.StatusOK) && (respFinalize.StatusCode != http.StatusNoContent) {
		b, readErr := io.ReadAll(respFinalize.Body)
		if readErr != nil {
			return respFinalize.StatusCode, "", fmt.Errorf("error reading FINALIZE response body: %w", readErr)
		}
		return respFinalize.StatusCode, string(b), fmt.Errorf("twitter FINALIZE request returned status %v", respFinalize.Status)
	}
	if respFinalize.StatusCode != http.StatusNoContent {
		var imgResponseFinalize ImageResponse
		err = decodeResponse(respFinalize, &imgResponseFinalize)
		if err != nil {
			return 0, "", err
		}
	}

	// Now check if Twitter have processed the video and finished on its end
	form = url.Values{
		"command":  {"STATUS"},
		"media_id": {uploadID},
	}
	const maxLoopLimit int = 8
	counter := 0
	for {
		respMediaStatus, err := client.Get(nil, &tokenCredentials, mediaUploadURL, form)
		if err != nil {
			return 0, "", fmt.Errorf("get error: %w", err)
		}
		var vidUploadStatus VideoUploadStatus
		err = decodeResponse(respMediaStatus, &vidUploadStatus)
		_ = respMediaStatus.Body.Close() // best-effort: body was fully read by decodeResponse
		if err != nil {
			return 0, "", err
		}
		if vidUploadStatus.ProcessingInfo.ProgressPercent == 100 {
			break
		}
		time.Sleep(videoStatusPollInterval)
		counter++
		if counter >= maxLoopLimit {
			return 0, "", errors.New("aborted due to infinite loop condition check")
		}
	}

	// Send STATUS UPDATE request
	form = url.Values{
		"status":                {message},
		"media_ids":             {uploadID},
		"in_reply_to_status_id": {replyTweet},
	}
	respUpdateStatus, err := client.Post(nil, &tokenCredentials, statusUpdateURL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = respUpdateStatus.Body.Close() }() // best-effort close on read path

	return readBodyAndStatusError(respUpdateStatus)
}
