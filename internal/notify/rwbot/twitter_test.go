package rwbot

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// recordedSend captures one adapter call to a tweets send function.
type recordedSend struct {
	msg   string
	nonce uint64
	media []byte
	reply string
}

func stubbedNotifier(imageStatus int, imageBody string, imageErr error, videoStatus int, videoBody string, videoErr error) (*TwitterNotifier, *[]recordedSend) {
	var calls []recordedSend
	n := NewTwitterNotifier(tweets.TwitterKeys{APIKey: "k", APISecret: "s", TokenKey: "tk", TokenSecret: "ts"})
	n.nonce.Store(100)
	n.sendImage = func(_, _, _, _, message string, nonce uint64, media []byte, reply string) (int, string, error) {
		calls = append(calls, recordedSend{msg: message, nonce: nonce, media: media, reply: reply})
		return imageStatus, imageBody, imageErr
	}
	n.sendVideo = func(_, _, _, _, message string, nonce uint64, media []byte, reply string) (int, string, error) {
		calls = append(calls, recordedSend{msg: message, nonce: nonce, media: media, reply: reply})
		return videoStatus, videoBody, videoErr
	}
	return n, &calls
}

func TestTweetWithImageReturnsTweetID(t *testing.T) {
	n, calls := stubbedNotifier(200, `{"id":123,"id_str":"123"}`, nil, 200, "", nil)
	id, err := n.TweetWithImage(context.Background(), "hello", []byte("img"))
	if err != nil {
		t.Fatalf("TweetWithImage: %v", err)
	}
	if id != "123" {
		t.Errorf("tweet id = %q, want 123", id)
	}
	if len(*calls) != 1 {
		t.Fatalf("calls = %d, want 1", len(*calls))
	}
	call := (*calls)[0]
	if call.msg != "hello" || string(call.media) != "img" || call.reply != "" {
		t.Errorf("recorded call = %+v", call)
	}
	if call.nonce != 101 {
		t.Errorf("nonce = %d, want 101 (incremented from seed 100)", call.nonce)
	}
}

func TestTweetWithImageErrorPaths(t *testing.T) {
	n, _ := stubbedNotifier(0, "", errors.New("network"), 200, "", nil)
	if _, err := n.TweetWithImage(context.Background(), "m", nil); err == nil || !strings.Contains(err.Error(), "network") {
		t.Errorf("transport error = %v", err)
	}

	n, _ = stubbedNotifier(403, "forbidden", nil, 200, "", nil)
	if _, err := n.TweetWithImage(context.Background(), "m", nil); err == nil || !strings.Contains(err.Error(), "status 403") {
		t.Errorf("non-200 error = %v", err)
	}

	n, _ = stubbedNotifier(200, "{not json", nil, 200, "", nil)
	if _, err := n.TweetWithImage(context.Background(), "m", nil); err == nil || !strings.Contains(err.Error(), "decoding") {
		t.Errorf("decode error = %v", err)
	}
}

func TestTweetWithVideoThreadsReplyAndChecksStatus(t *testing.T) {
	n, calls := stubbedNotifier(200, `{"id_str":"1"}`, nil, 200, "ok", nil)
	if err := n.TweetWithVideo(context.Background(), "vid msg", []byte("vv"), "999"); err != nil {
		t.Fatalf("TweetWithVideo: %v", err)
	}
	call := (*calls)[0]
	if call.msg != "vid msg" || string(call.media) != "vv" || call.reply != "999" {
		t.Errorf("recorded call = %+v", call)
	}

	n, _ = stubbedNotifier(200, "", nil, 500, "server error", nil)
	if err := n.TweetWithVideo(context.Background(), "m", nil, ""); err == nil || !strings.Contains(err.Error(), "status 500") {
		t.Errorf("non-200 video error = %v", err)
	}

	n, _ = stubbedNotifier(200, "", nil, 0, "", errors.New("boom"))
	if err := n.TweetWithVideo(context.Background(), "m", nil, ""); err == nil || !strings.Contains(err.Error(), "boom") {
		t.Errorf("video transport error = %v", err)
	}
}

func TestNonceSequenceAdvancesAcrossCalls(t *testing.T) {
	n, calls := stubbedNotifier(200, `{"id_str":"1"}`, nil, 200, "", nil)
	_, _ = n.TweetWithImage(context.Background(), "a", nil)
	_ = n.TweetWithVideo(context.Background(), "b", nil, "")
	_, _ = n.TweetWithImage(context.Background(), "c", nil)
	if len(*calls) != 3 {
		t.Fatalf("calls = %d, want 3", len(*calls))
	}
	for i, want := range []uint64{101, 102, 103} {
		if (*calls)[i].nonce != want {
			t.Errorf("call %d nonce = %d, want %d (strictly increasing)", i, (*calls)[i].nonce, want)
		}
	}
}

func TestNewTwitterNotifierSeedsNonceAndProductionSenders(t *testing.T) {
	n := NewTwitterNotifier(tweets.TwitterKeys{})
	if n.nonce.Load() == 0 {
		t.Error("nonce seed is zero, want wall-clock seed")
	}
	if n.sendImage == nil || n.sendVideo == nil {
		t.Error("production senders not wired")
	}
}
