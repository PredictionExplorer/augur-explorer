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

const (
	URL       string = "https://api.twitter.com/1.1/statuses/update.json"
	MEDIA_URL string = "https://upload.twitter.com/1.1/media/upload.json"
)

type TwitterKeys struct {
	ApiKey      string
	ApiSecret   string
	TokenKey    string
	TokenSecret string
}
type MediaImage struct {
	Image_type string `json:"image_type"`
	W          int64  `json:"w"`
	H          int64  `json:"h"`
}
type ImageResponse struct {
	Media_id           int64      `json:"media_id"`
	Media_id_string    string     `json:"media_id_string"`
	Media_key          string     `json:"media_key"`
	Size               int64      `json:"size"`
	Expires_after_secs int64      `json:"expires_after_secs"`
	Image              MediaImage `json:"image"`
}
type StatusUpdateResponse struct {
	Id    int64  `json:"id"`
	IdStr string `json:"id_str"`
}
type ProcessingInfo struct {
	State            string `json:"in_progress"`
	Check_after_secs int64  `json:"check_after_secs"`
	Progress_percent int64  `json:"progress_percent"`
}
type VideoUploadStatus struct {
	Media_id           int64          `json:"media_id"`
	Media_id_string    string         `json:"media_id_string"`
	Expires_after_secs int64          `json:"expires_after_secs"`
	Processing_info    ProcessingInfo `json:"processing_info"`
}

func decode_response(resp *http.Response, data interface{}) error {
	p, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	// note: we did ReadAll() because we need to have an option to dump the received body
	body_stream := bytes.NewReader(p)
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusAccepted) && (resp.StatusCode != http.StatusNoContent) {
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, string(p))
	}
	return json.NewDecoder(body_stream).Decode(data)
}
func format_nonce(nonce_int uint64) string {
	return strconv.FormatUint(nonce_int, 16)
}

// readBodyAndStatusError reads the response body and builds the (status code,
// body, error) triple returned by the Send* helpers. A non-nil error is
// returned when the response status differs from http.StatusOK.
func readBodyAndStatusError(resp *http.Response) (int, string, error) {
	var err_out error
	if resp.StatusCode != http.StatusOK {
		err_out = fmt.Errorf("error at Twitter occurred: status = %v", resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil && err_out == nil {
		err_out = fmt.Errorf("error reading response body: %w", err)
	}
	return resp.StatusCode, string(b), err_out
}
func newTwitterClient(api_key, api_secret, access_token string, session_nonce uint64) Client {
	var client Client
	client.Credentials.Token = api_key     // user-generated token
	client.Credentials.Secret = api_secret // app secret
	client.APIKey = api_key
	client.ClientToken = access_token
	client.Nonce = format_nonce(session_nonce)
	return client
}
func SendTweet(api_key, api_secret, access_token, token_secret, message string, session_nonce uint64) (int, string, error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
	client := newTwitterClient(api_key, api_secret, access_token, session_nonce)

	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	form := url.Values{"status": {message}}
	resp, err := client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path

	return readBodyAndStatusError(resp)
}
func SendTweetWithImage(api_key, api_secret, access_token, token_secret, message string, session_nonce uint64, image_data []byte, reply_tweet string) (int, string, error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
	client := newTwitterClient(api_key, api_secret, access_token, session_nonce)

	encoded_image_data := base64.StdEncoding.EncodeToString(image_data)
	form := url.Values{
		"media_data": {encoded_image_data},
	}
	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	resp, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	var image_response ImageResponse
	err = decode_response(resp, &image_response)
	_ = resp.Body.Close() // best-effort: body was fully read by decode_response
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("error decoding media upload response: %w", err)
	}

	form = url.Values{
		"status":                {message},
		"media_ids":             {image_response.Media_id_string},
		"in_reply_to_status_id": {reply_tweet},
	}
	resp, err = client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path

	return readBodyAndStatusError(resp)
}
func SendTweetWithAttachment(api_key, api_secret, access_token, token_secret, message string, session_nonce uint64, reply_tweet_id string, attachment_url string) (int, string, error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
	client := newTwitterClient(api_key, api_secret, access_token, session_nonce)

	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	form := url.Values{
		"status":                {message},
		"in_reply_to_status_id": {reply_tweet_id},
		"attachment_url":        {attachment_url},
	}
	resp, err := client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path

	return readBodyAndStatusError(resp)
}
func SendTweetWithMedia(api_key, api_secret, access_token, token_secret, message string, session_nonce uint64, media_type string, image_data []byte, reply_tweet string) (int, string, error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
	client := newTwitterClient(api_key, api_secret, access_token, session_nonce)
	encoded_image_data := base64.StdEncoding.EncodeToString(image_data)
	form := url.Values{
		"media_data":     {encoded_image_data},
		"media_category": {"tweet_video"},
	}
	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	resp, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	var image_response ImageResponse
	err = decode_response(resp, &image_response)
	_ = resp.Body.Close() // best-effort: body was fully read by decode_response
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("error decoding media upload response: %w", err)
	}

	form = url.Values{
		"status":                {message},
		"media_ids":             {image_response.Media_id_string},
		"in_reply_to_status_id": {reply_tweet},
	}
	resp, err = client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp.Body.Close() }() // best-effort close on read path

	return readBodyAndStatusError(resp)
}
func SendTweetWithVideo(api_key, api_secret, access_token, token_secret, message string, session_nonce uint64, video_data []byte, reply_tweet string) (int, string, error) {
	// Return values:
	//		Status code
	//		Body (converted to string)
	//		Error from net/http, if any
	client := newTwitterClient(api_key, api_secret, access_token, session_nonce)
	encoded_video_data := base64.StdEncoding.EncodeToString(video_data)
	form := url.Values{
		"command":        {"INIT"},
		"total_bytes":    {fmt.Sprintf("%v", len(video_data))},
		"media_category": {"tweet_video"},
		"media_type":     {"video/mp4"},
	}
	var token_credentials Credentials
	token_credentials.Token = access_token
	token_credentials.Secret = token_secret

	//---------- INIT
	resp_init, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp_init.Body.Close() }() // best-effort close on read path
	if resp_init.StatusCode != http.StatusAccepted {
		b, read_err := io.ReadAll(resp_init.Body)
		if read_err != nil {
			return resp_init.StatusCode, "", fmt.Errorf("error reading INIT response body: %w", read_err)
		}
		return resp_init.StatusCode, string(b), fmt.Errorf("twitter INIT request returned status %v", resp_init.Status)
	}
	var img_response_init ImageResponse
	err = decode_response(resp_init, &img_response_init)
	if err != nil {
		return 0, "", err
	}
	upload_id := img_response_init.Media_id_string

	//------------ APPEND
	form = url.Values{
		"command":       {"APPEND"},
		"media_id":      {upload_id},
		"media_data":    {encoded_video_data + "\r\n"},
		"segment_index": {fmt.Sprintf("%v", 0)},
	}
	client.Header = make(map[string][]string)
	client.Header.Set("Content-Type", "application/octet-stream")
	client.Header.Set("Content-Transfer-Encoding", "base64")
	resp_append, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp_append.Body.Close() }() // best-effort close on read path
	if (resp_append.StatusCode != http.StatusOK) && (resp_append.StatusCode != http.StatusNoContent) {
		b, read_err := io.ReadAll(resp_append.Body)
		if read_err != nil {
			return resp_append.StatusCode, "", fmt.Errorf("error reading APPEND response body: %w", read_err)
		}
		return resp_append.StatusCode, string(b), fmt.Errorf("twitter APPEND request returned status %v", resp_append.Status)
	}
	if resp_append.StatusCode != http.StatusNoContent {
		var img_response_append ImageResponse
		err = decode_response(resp_append, &img_response_append)
		if err != nil {
			return 0, "", err
		}
	}
	client.Header.Del("Content-Type")
	client.Header.Del("Content-Transfer-Encoding")
	//----------- FINALIZE
	form = url.Values{
		"command":  {"FINALIZE"},
		"media_id": {upload_id},
	}
	resp_finalize, err := client.Post(nil, &token_credentials, MEDIA_URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp_finalize.Body.Close() }() // best-effort close on read path
	if (resp_finalize.StatusCode != http.StatusOK) && (resp_finalize.StatusCode != http.StatusNoContent) {
		b, read_err := io.ReadAll(resp_finalize.Body)
		if read_err != nil {
			return resp_finalize.StatusCode, "", fmt.Errorf("error reading FINALIZE response body: %w", read_err)
		}
		return resp_finalize.StatusCode, string(b), fmt.Errorf("twitter FINALIZE request returned status %v", resp_finalize.Status)
	}
	if resp_finalize.StatusCode != http.StatusNoContent {
		var img_response_finalize ImageResponse
		err = decode_response(resp_finalize, &img_response_finalize)
		if err != nil {
			return 0, "", err
		}
	}

	// Now check if Twitter have processed the video and finished on its end
	form = url.Values{
		"command":  {"STATUS"},
		"media_id": {upload_id},
	}
	const MAX_LOOP_LIMIT int = 8
	counter := 0
	for {
		resp_media_status, err := client.Get(nil, &token_credentials, MEDIA_URL, form)
		if err != nil {
			return 0, "", fmt.Errorf("get error: %w", err)
		}
		var vid_upload_status VideoUploadStatus
		err = decode_response(resp_media_status, &vid_upload_status)
		_ = resp_media_status.Body.Close() // best-effort: body was fully read by decode_response
		if err != nil {
			return 0, "", err
		}
		if vid_upload_status.Processing_info.Progress_percent == 100 {
			break
		}
		time.Sleep(2 * time.Second)
		counter++
		if counter >= MAX_LOOP_LIMIT {
			return 0, "", errors.New("aborted due to infinite loop condition check")
		}
	}

	// Send STATUS UPDATE request
	form = url.Values{
		"status":                {message},
		"media_ids":             {upload_id},
		"in_reply_to_status_id": {reply_tweet},
	}
	resp_update_status, err := client.Post(nil, &token_credentials, URL, form)
	if err != nil {
		return 0, "", fmt.Errorf("post error: %w", err)
	}
	defer func() { _ = resp_update_status.Body.Close() }() // best-effort close on read path

	return readBodyAndStatusError(resp_update_status)
}
