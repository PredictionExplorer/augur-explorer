package tweets

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// twitterStub stands in for both Twitter endpoints (status update + media
// upload). It records every request and lets each test script the media
// endpoint's behavior per upload command.
type twitterStub struct {
	t *testing.T

	statusUpdate http.HandlerFunc
	media        http.HandlerFunc

	requests []stubRequest
}

type stubRequest struct {
	path string
	form map[string]string
	auth string
}

// installTwitterStub starts the stub server and points the package's endpoint
// variables at it for the duration of the test. Tests mutating package state
// must not run in parallel; none of these do.
func installTwitterStub(t *testing.T) *twitterStub {
	t.Helper()
	stub := &twitterStub{t: t}

	mux := http.NewServeMux()
	mux.HandleFunc("/1.1/statuses/update.json", func(w http.ResponseWriter, r *http.Request) {
		stub.record(r)
		if stub.statusUpdate != nil {
			stub.statusUpdate(w, r)
			return
		}
		_, _ = w.Write([]byte(`{"id":12345,"id_str":"12345"}`))
	})
	mux.HandleFunc("/1.1/media/upload.json", func(w http.ResponseWriter, r *http.Request) {
		stub.record(r)
		if stub.media != nil {
			stub.media(w, r)
			return
		}
		_, _ = w.Write([]byte(`{"media_id":777,"media_id_string":"777"}`))
	})
	srv := httptest.NewServer(mux)

	prevStatus, prevMedia, prevPoll := statusUpdateURL, mediaUploadURL, videoStatusPollInterval
	statusUpdateURL = srv.URL + "/1.1/statuses/update.json"
	mediaUploadURL = srv.URL + "/1.1/media/upload.json"
	videoStatusPollInterval = time.Millisecond
	t.Cleanup(func() {
		statusUpdateURL, mediaUploadURL, videoStatusPollInterval = prevStatus, prevMedia, prevPoll
		srv.Close()
	})
	return stub
}

func (s *twitterStub) record(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.t.Errorf("stub could not parse form: %v", err)
	}
	form := make(map[string]string, len(r.Form))
	for k := range r.Form {
		form[k] = r.Form.Get(k)
	}
	s.requests = append(s.requests, stubRequest{
		path: r.URL.Path,
		form: form,
		auth: r.Header.Get("Authorization"),
	})
}

func (s *twitterStub) request(i int) stubRequest {
	s.t.Helper()
	if i >= len(s.requests) {
		s.t.Fatalf("only %d requests recorded, want index %d", len(s.requests), i)
	}
	return s.requests[i]
}

func TestSendTweet(t *testing.T) {
	stub := installTwitterStub(t)

	status, body, err := SendTweet("api-key", "api-secret", "access-token", "token-secret", "hello", 7)
	if err != nil {
		t.Fatalf("SendTweet: %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("status = %d, want 200", status)
	}
	var resp StatusUpdateResponse
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		t.Fatalf("body %q is not the stubbed JSON: %v", body, err)
	}
	if resp.IdStr != "12345" {
		t.Errorf("IdStr = %q, want 12345", resp.IdStr)
	}

	req := stub.request(0)
	if req.form["status"] != "hello" {
		t.Errorf("status form field = %q", req.form["status"])
	}
	// The session nonce is formatted as lowercase hex into the OAuth header.
	if !strings.Contains(req.auth, fmt.Sprintf("oauth_nonce=%q", formatNonce(7))) {
		t.Errorf("Authorization %q missing session nonce", req.auth)
	}
	if !strings.Contains(req.auth, `oauth_token="access-token"`) {
		t.Errorf("Authorization %q missing access token", req.auth)
	}
}

func TestSendTweetErrorStatusStillReturnsBody(t *testing.T) {
	stub := installTwitterStub(t)
	stub.statusUpdate = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"errors":[{"code":187}]}`))
	}

	status, body, err := SendTweet("k", "s", "at", "ts", "dup", 1)
	if err == nil {
		t.Fatal("expected error for 403 response")
	}
	if status != http.StatusForbidden {
		t.Errorf("status = %d, want 403", status)
	}
	if !strings.Contains(body, "187") {
		t.Errorf("body %q should carry the Twitter error payload", body)
	}
}

func TestSendTweetTransportError(t *testing.T) {
	prev := statusUpdateURL
	statusUpdateURL = "http://127.0.0.1:1/unreachable"
	t.Cleanup(func() { statusUpdateURL = prev })

	status, _, err := SendTweet("k", "s", "at", "ts", "msg", 1)
	if err == nil || status != 0 {
		t.Fatalf("want transport error and status 0, got status=%d err=%v", status, err)
	}
}

func TestSendTweetWithImage(t *testing.T) {
	stub := installTwitterStub(t)
	imageData := []byte{0x89, 'P', 'N', 'G'}

	status, _, err := SendTweetWithImage("k", "s", "at", "ts", "look at this", 3, imageData, "999")
	if err != nil {
		t.Fatalf("SendTweetWithImage: %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("status = %d, want 200", status)
	}

	upload := stub.request(0)
	if upload.path != "/1.1/media/upload.json" {
		t.Errorf("first request path = %q, want media upload", upload.path)
	}
	if got := upload.form["media_data"]; got != base64.StdEncoding.EncodeToString(imageData) {
		t.Errorf("media_data = %q, want base64 of image bytes", got)
	}

	update := stub.request(1)
	if update.path != "/1.1/statuses/update.json" {
		t.Errorf("second request path = %q, want status update", update.path)
	}
	if update.form["media_ids"] != "777" {
		t.Errorf("media_ids = %q, want the uploaded id 777", update.form["media_ids"])
	}
	if update.form["in_reply_to_status_id"] != "999" {
		t.Errorf("in_reply_to_status_id = %q, want 999", update.form["in_reply_to_status_id"])
	}
}

func TestSendTweetWithImageUploadRejected(t *testing.T) {
	stub := installTwitterStub(t)
	stub.media = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"errors":[{"message":"media type unrecognized"}]}`))
	}

	status, _, err := SendTweetWithImage("k", "s", "at", "ts", "msg", 1, []byte{1}, "")
	if err == nil {
		t.Fatal("expected error when the media upload is rejected")
	}
	if status != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", status)
	}
	if !strings.Contains(err.Error(), "media type unrecognized") {
		t.Errorf("error %q should carry the upload response", err)
	}
	if len(stub.requests) != 1 {
		t.Errorf("no status update must be attempted after a failed upload; got %d requests", len(stub.requests))
	}
}

// videoStub scripts the four-stage upload protocol. statusPolls counts STATUS
// requests so tests can serve "in progress" a fixed number of times.
type videoStub struct {
	stub        *twitterStub
	statusPolls int

	initStatus     int
	appendStatus   int
	finalizeStatus int
	pendingPolls   int
}

func installVideoStub(t *testing.T) *videoStub {
	t.Helper()
	vs := &videoStub{
		stub:           installTwitterStub(t),
		initStatus:     http.StatusAccepted,
		appendStatus:   http.StatusNoContent,
		finalizeStatus: http.StatusOK,
	}
	vs.stub.media = func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("video stub parse form: %v", err)
		}
		switch r.Form.Get("command") {
		case "INIT":
			w.WriteHeader(vs.initStatus)
			_, _ = w.Write([]byte(`{"media_id":42,"media_id_string":"42"}`))
		case "APPEND":
			w.WriteHeader(vs.appendStatus)
			if vs.appendStatus == http.StatusOK {
				_, _ = w.Write([]byte(`{"media_id":42,"media_id_string":"42"}`))
			}
		case "FINALIZE":
			w.WriteHeader(vs.finalizeStatus)
			if vs.finalizeStatus == http.StatusOK {
				_, _ = w.Write([]byte(`{"media_id":42,"media_id_string":"42"}`))
			}
		case "STATUS":
			vs.statusPolls++
			progress := int64(100)
			if vs.statusPolls <= vs.pendingPolls {
				progress = 50
			}
			_, _ = fmt.Fprintf(w, `{"media_id":42,"media_id_string":"42","processing_info":{"state":"in_progress","progress_percent":%d}}`, progress)
		default:
			t.Errorf("unexpected media command %q", r.Form.Get("command"))
		}
	}
	return vs
}

func TestSendTweetWithVideo(t *testing.T) {
	vs := installVideoStub(t)
	vs.pendingPolls = 2 // two "in progress" polls before completion

	status, _, err := SendTweetWithVideo("k", "s", "at", "ts", "video!", 5, []byte("mp4data"), "111")
	if err != nil {
		t.Fatalf("SendTweetWithVideo: %v", err)
	}
	if status != http.StatusOK {
		t.Errorf("status = %d, want 200", status)
	}
	if vs.statusPolls != 3 {
		t.Errorf("STATUS polled %d times, want 3 (two pending + one complete)", vs.statusPolls)
	}

	// INIT carries the byte count; APPEND the base64 payload; the final
	// status update references the upload id from INIT.
	reqs := vs.stub.requests
	last := reqs[len(reqs)-1]
	if last.path != "/1.1/statuses/update.json" {
		t.Fatalf("last request path = %q, want status update", last.path)
	}
	if last.form["media_ids"] != "42" {
		t.Errorf("media_ids = %q, want 42", last.form["media_ids"])
	}
	var sawInit, sawAppend bool
	for _, r := range reqs {
		switch r.form["command"] {
		case "INIT":
			sawInit = true
			if r.form["total_bytes"] != "7" {
				t.Errorf("INIT total_bytes = %q, want 7", r.form["total_bytes"])
			}
		case "APPEND":
			sawAppend = true
			want := base64.StdEncoding.EncodeToString([]byte("mp4data")) + "\r\n"
			if r.form["media_data"] != want {
				t.Errorf("APPEND media_data = %q, want %q", r.form["media_data"], want)
			}
		}
	}
	if !sawInit || !sawAppend {
		t.Errorf("INIT/APPEND requests missing (init=%v append=%v)", sawInit, sawAppend)
	}
}

func TestSendTweetWithVideoAppendOKBody(t *testing.T) {
	vs := installVideoStub(t)
	vs.appendStatus = http.StatusOK // APPEND may answer 200 with a body instead of 204

	if _, _, err := SendTweetWithVideo("k", "s", "at", "ts", "v", 1, []byte("x"), ""); err != nil {
		t.Fatalf("SendTweetWithVideo with 200 APPEND: %v", err)
	}
}

func TestSendTweetWithVideoInitRejected(t *testing.T) {
	vs := installVideoStub(t)
	vs.initStatus = http.StatusBadRequest

	status, body, err := SendTweetWithVideo("k", "s", "at", "ts", "v", 1, []byte("x"), "")
	if err == nil {
		t.Fatal("expected INIT failure")
	}
	if status != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", status)
	}
	if !strings.Contains(body, "media_id") {
		t.Errorf("body %q should be the INIT response payload", body)
	}
	if len(vs.stub.requests) != 1 {
		t.Errorf("upload must stop after INIT failure; got %d requests", len(vs.stub.requests))
	}
}

func TestSendTweetWithVideoAppendRejected(t *testing.T) {
	vs := installVideoStub(t)
	vs.appendStatus = http.StatusBadRequest

	if _, _, err := SendTweetWithVideo("k", "s", "at", "ts", "v", 1, []byte("x"), ""); err == nil ||
		!strings.Contains(err.Error(), "APPEND") {
		t.Fatalf("want APPEND failure, got %v", err)
	}
}

func TestSendTweetWithVideoFinalizeRejected(t *testing.T) {
	vs := installVideoStub(t)
	vs.finalizeStatus = http.StatusBadRequest

	if _, _, err := SendTweetWithVideo("k", "s", "at", "ts", "v", 1, []byte("x"), ""); err == nil ||
		!strings.Contains(err.Error(), "FINALIZE") {
		t.Fatalf("want FINALIZE failure, got %v", err)
	}
}

func TestSendTweetWithVideoProcessingStuckAborts(t *testing.T) {
	vs := installVideoStub(t)
	vs.pendingPolls = 1000 // never completes

	_, _, err := SendTweetWithVideo("k", "s", "at", "ts", "v", 1, []byte("x"), "")
	if err == nil || !strings.Contains(err.Error(), "infinite loop") {
		t.Fatalf("want stuck-processing abort, got %v", err)
	}
	if vs.statusPolls != 8 {
		t.Errorf("STATUS polled %d times before aborting, want the 8-attempt cap", vs.statusPolls)
	}
}

func TestDecodeResponseRejectsErrorStatus(t *testing.T) {
	stub := installTwitterStub(t)
	stub.media = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("upstream broke"))
	}
	_, _, err := SendTweetWithImage("k", "s", "at", "ts", "m", 1, []byte{1}, "")
	if err == nil || !strings.Contains(err.Error(), "upstream broke") {
		t.Fatalf("decodeResponse should reject 500 with the body included, got %v", err)
	}
}

func TestFormatNonce(t *testing.T) {
	if got := formatNonce(255); got != "ff" {
		t.Errorf("formatNonce(255) = %q, want ff", got)
	}
	if got := formatNonce(0); got != "0" {
		t.Errorf("formatNonce(0) = %q, want 0", got)
	}
}
