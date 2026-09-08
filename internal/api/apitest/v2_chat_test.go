//go:build integration

package apitest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

const (
	v2ChatMessages = "/api/v2/cosmicgame/rounds/{round}/messages"
	v2ChatContext  = "/api/v2/cosmicgame/rounds/{round}/chat-context"
)

// The durable revision legitimately varies with shuffled moderation tests.
// Validate the real response first; normalize only this generated generation
// in golden output, including the corresponding opaque cursor field.
func pinChatGolden(t *testing.T, name string, recorder *httptest.ResponseRecorder) {
	t.Helper()
	body := canonicalJSON(t, recorder.Body.Bytes())
	if object, ok := body.(map[string]any); ok {
		if meta, ok := object["meta"].(map[string]any); ok {
			meta["revision"] = "1"
			for _, key := range []string{"nextCursor", "syncCursor"} {
				encoded, ok := meta[key].(string)
				if !ok {
					continue
				}
				payload, err := base64.RawURLEncoding.DecodeString(encoded)
				if err != nil {
					t.Fatal(err)
				}
				var cursor map[string]any
				if err := json.Unmarshal(payload, &cursor); err != nil {
					t.Fatal(err)
				}
				cursor["g"] = 1
				payload, err = json.Marshal(cursor)
				if err != nil {
					t.Fatal(err)
				}
				meta[key] = base64.RawURLEncoding.EncodeToString(payload)
			}
		}
	}
	compareV2Golden(t, name, response{Status: recorder.Code, ContentType: contentTypeOf(recorder), Body: body})
}

func TestAPIV2ChatMessagesAndContext(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	requestPage := func(name, target, template, round string) *httptest.ResponseRecorder {
		t.Helper()
		rr := h.do(t, request{path: target})
		validateV2Response(t, spec, v2GoldenCase{name: name, target: target, template: template, pathParams: map[string]string{"round": round}}, rr)
		pinChatGolden(t, name, rr)
		return rr
	}
	first := requestPage("chat_messages_first", "/api/v2/cosmicgame/rounds/0/messages?limit=1", v2ChatMessages, "0")
	var page apiv2.RoundChatMessagePage
	decodeV2JSON(t, first, &page)
	if len(page.Data) != 1 || page.Data[0].EventLogId != 5008 || page.Meta.NextCursor == nil || page.Meta.HasMore {
		t.Fatalf("first page=%+v", page)
	}
	if strings.Contains(first.Body.String(), "cstReward") || strings.Contains(first.Body.String(), "Donation") {
		t.Fatalf("bloated message DTO: %s", first.Body.String())
	}
	older := requestPage("chat_messages_older", "/api/v2/cosmicgame/rounds/0/messages?limit=1&cursor="+*page.Meta.NextCursor, v2ChatMessages, "0")
	var oldPage apiv2.RoundChatMessagePage
	decodeV2JSON(t, older, &oldPage)
	if len(oldPage.Data) != 1 || oldPage.Data[0].EventLogId != 5004 || oldPage.Meta.NextCursor != nil {
		t.Fatalf("older=%+v", oldPage)
	}
	requestPage("chat_messages_caught_up", "/api/v2/cosmicgame/rounds/0/messages?after="+page.Meta.SyncCursor, v2ChatMessages, "0")
	requestPage("chat_messages_empty", "/api/v2/cosmicgame/rounds/99999/messages", v2ChatMessages, "99999")
	contextResponse := requestPage("chat_context", "/api/v2/cosmicgame/rounds/0/chat-context", v2ChatContext, "0")
	var snapshot apiv2.RoundChatContext
	decodeV2JSON(t, contextResponse, &snapshot)
	if len(snapshot.Data) != 4 || snapshot.Meta.Revision != page.Meta.Revision {
		t.Fatalf("context=%+v", snapshot)
	}
	if strings.Contains(contextResponse.Body.String(), "message") || strings.Contains(contextResponse.Body.String(), "PriceWei") || strings.Contains(contextResponse.Body.String(), "Reward") {
		t.Fatalf("context leaked fields: %s", contextResponse.Body.String())
	}
	requestPage("chat_context_empty", "/api/v2/cosmicgame/rounds/99999/chat-context", v2ChatContext, "99999")
	requestPage("chat_messages_invalid_limit", "/api/v2/cosmicgame/rounds/0/messages?limit=201", v2ChatMessages, "0")
	requestPage("chat_messages_invalid_cursor", "/api/v2/cosmicgame/rounds/0/messages?cursor=invalid", v2ChatMessages, "0")
	requestPage("chat_messages_wrong_direction", "/api/v2/cosmicgame/rounds/0/messages?after="+*page.Meta.NextCursor, v2ChatMessages, "0")
	requestPage("chat_messages_wrong_round", "/api/v2/cosmicgame/rounds/1/messages?after="+page.Meta.SyncCursor, v2ChatMessages, "1")
	requestPage("chat_context_invalid_round", "/api/v2/cosmicgame/rounds/-1/chat-context", v2ChatContext, "-1")

	// An actual moderation write invalidates both kinds of historical cursor.
	if _, err := h.db.Exec(`INSERT INTO cg_banned_bids(bid_id,user_addr,created_at) VALUES(2001,'alice',1)`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := h.db.Exec(`DELETE FROM cg_banned_bids WHERE bid_id=2001`); err != nil {
			t.Error(err)
		}
	})
	stale := requestPage("chat_messages_reset", "/api/v2/cosmicgame/rounds/0/messages?after="+page.Meta.SyncCursor, v2ChatMessages, "0")
	if stale.Code != http.StatusConflict || !strings.Contains(stale.Body.String(), "feed-reset-required") {
		t.Fatalf("stale=%d %s", stale.Code, stale.Body.String())
	}

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	for _, target := range []string{"/api/v2/cosmicgame/rounds/0/messages", "/api/v2/cosmicgame/rounds/0/chat-context"} {
		rr := h.do(t, request{path: target, ctx: cancelled})
		if rr.Code != http.StatusInternalServerError {
			t.Fatalf("backend failure must not fall back: %d %s", rr.Code, rr.Body.String())
		}
	}
}
