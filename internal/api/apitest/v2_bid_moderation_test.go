//go:build integration

package apitest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

const (
	v2BannedBids     = "/api/v2/cosmicgame/moderation/banned-bids"
	v2BannedBidsItem = "/api/v2/cosmicgame/moderation/banned-bids/{bidId}"
)

func restoreV2BidModerationState(t *testing.T, h *harness) {
	t.Helper()
	if _, err := h.db.Exec(`DELETE FROM cg_banned_bids WHERE bid_id <> 2002`); err != nil {
		t.Fatalf("restoring bid-ban fixture: %v", err)
	}
}

func pinV2ExchangeGolden(
	t *testing.T,
	spec *openapi3.T,
	name, method, template, target, requestBody string,
	headers map[string]string,
	recorder *httptest.ResponseRecorder,
) {
	t.Helper()
	validateV2MutationResponse(t, spec, method, template, target, requestBody, headers, recorder)
	var body any
	if recorder.Body.Len() > 0 {
		body = canonicalJSON(t, recorder.Body.Bytes())
	}
	compareV2Golden(t, name, response{
		Status:      recorder.Code,
		ContentType: contentTypeOf(recorder),
		Body:        body,
	})
}

func TestAPIV2BidModerationReads(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	restoreV2BidModerationState(t, h)
	t.Cleanup(func() { restoreV2BidModerationState(t, h) })

	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{name: "bid_moderation_list", target: v2BannedBids, template: v2BannedBids, pathParams: map[string]string{}},
		{name: "bid_moderation_list_error_limit", target: v2BannedBids + "?limit=0", template: v2BannedBids, pathParams: map[string]string{}},
		{name: "bid_moderation_list_error_cursor", target: v2BannedBids + "?cursor=invalid", template: v2BannedBids, pathParams: map[string]string{}},
	})

	// Use fixed high row ids so the continuation cursor is deterministic even
	// when shuffled mutation tests have advanced the table sequence.
	if _, err := h.db.Exec(`INSERT INTO cg_banned_bids(id,bid_id,user_addr,created_at) VALUES
		(900001,2010,$1,1767230100),
		(900002,2011,$2,1767230200)`,
		addrAlice, addrBob); err != nil {
		t.Fatalf("seeding deterministic bid-ban page: %v", err)
	}
	first := h.get(t, v2BannedBids+"?limit=1")
	var page apiv2.BannedBidPage
	decodeV2JSON(t, first, &page)
	if len(page.Data) != 1 || page.Data[0].BidId != 2011 || page.Meta.NextCursor == nil {
		t.Fatalf("first moderation page = %+v", page)
	}
	wantCursor := base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"i":900002}`))
	if *page.Meta.NextCursor != wantCursor {
		t.Fatalf("next cursor = %q, want %q", *page.Meta.NextCursor, wantCursor)
	}
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{
			name:       "bid_moderation_list_first_page",
			target:     v2BannedBids + "?limit=1",
			template:   v2BannedBids,
			pathParams: map[string]string{},
		},
		{
			name:       "bid_moderation_list_next_page",
			target:     v2BannedBids + "?limit=1&cursor=" + wantCursor,
			template:   v2BannedBids,
			pathParams: map[string]string{},
		},
	})
}

func TestAPIV2BidModerationLifecycle(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	restoreV2BidModerationState(t, h)
	t.Cleanup(func() { restoreV2BidModerationState(t, h) })

	body := `{"bidId":2011}`
	adminHeaders := map[string]string{"X-Admin-Key": adminKey}

	unauthorized := postV2(t, h, v2BannedBids, body, nil)
	pinV2MutationGolden(t, spec, "bid_moderation_create_error_unauthorized",
		v2BannedBids, v2BannedBids, body, nil, unauthorized, nil)
	if unauthorized.Code != http.StatusUnauthorized {
		t.Fatalf("missing-key status = %d", unauthorized.Code)
	}
	wrong := postV2(t, h, v2BannedBids, body, map[string]string{"X-Admin-Key": "wrong"})
	if wrong.Code != http.StatusUnauthorized {
		t.Fatalf("wrong-key status = %d", wrong.Code)
	}

	created := postV2(t, h, v2BannedBids, body, adminHeaders)
	if created.Code != http.StatusCreated {
		t.Fatalf("create status = %d; body=%s", created.Code, created.Body.String())
	}
	var ban apiv2.BannedBid
	if err := json.Unmarshal(created.Body.Bytes(), &ban); err != nil {
		t.Fatalf("decode created ban: %v", err)
	}
	if ban.BidId != 2011 || ban.BidderAddress != addrBob ||
		!ban.BannedAt.Equal(time.Unix(1767230000, 0).UTC()) {
		t.Fatalf("created ban = %+v", ban)
	}
	pinV2MutationGolden(t, spec, "bid_moderation_created",
		v2BannedBids, v2BannedBids, body, adminHeaders, created, nil)

	duplicate := postV2(t, h, v2BannedBids, body, adminHeaders)
	if duplicate.Code != http.StatusConflict {
		t.Fatalf("duplicate status = %d; body=%s", duplicate.Code, duplicate.Body.String())
	}
	pinV2MutationGolden(t, spec, "bid_moderation_create_error_conflict",
		v2BannedBids, v2BannedBids, body, adminHeaders, duplicate, nil)

	deleted := h.do(t, request{
		method:  http.MethodDelete,
		path:    v2BannedBids + "/2011",
		headers: adminHeaders,
	})
	if deleted.Code != http.StatusNoContent || deleted.Body.Len() != 0 {
		t.Fatalf("delete status = %d; body=%s", deleted.Code, deleted.Body.String())
	}
	pinV2ExchangeGolden(t, spec, "bid_moderation_deleted",
		http.MethodDelete, v2BannedBidsItem, v2BannedBids+"/2011", "", adminHeaders, deleted)

	missing := h.do(t, request{
		method:  http.MethodDelete,
		path:    v2BannedBids + "/2011",
		headers: adminHeaders,
	})
	if missing.Code != http.StatusNotFound {
		t.Fatalf("missing-delete status = %d; body=%s", missing.Code, missing.Body.String())
	}
	pinV2ExchangeGolden(t, spec, "bid_moderation_delete_error_not_found",
		http.MethodDelete, v2BannedBidsItem, v2BannedBids+"/2011", "", adminHeaders, missing)

	var fixtureCount int
	if err := h.db.QueryRow(`SELECT COUNT(*) FROM cg_banned_bids WHERE bid_id=2002`).Scan(&fixtureCount); err != nil {
		t.Fatalf("checking fixture ban: %v", err)
	}
	if fixtureCount != 1 {
		t.Fatalf("fixture ban count = %d, want 1", fixtureCount)
	}
}

func TestAPIV2BidModerationProblems(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	restoreV2BidModerationState(t, h)
	t.Cleanup(func() { restoreV2BidModerationState(t, h) })
	adminHeaders := map[string]string{"X-Admin-Key": adminKey}

	for _, tc := range []struct {
		name string
		body string
		want int
	}{
		{"bid_moderation_create_error_body", `{"bidId":["not-a-number"]}`, http.StatusBadRequest},
		{"bid_moderation_create_error_invalid_bid", `{"bidId":0}`, http.StatusBadRequest},
		{"bid_moderation_create_error_missing_bid", `{"bidId":999999}`, http.StatusNotFound},
	} {
		response := postV2(t, h, v2BannedBids, tc.body, adminHeaders)
		if response.Code != tc.want {
			t.Fatalf("%s status = %d, want %d; body=%s", tc.name, response.Code, tc.want, response.Body.String())
		}
		pinV2MutationGolden(t, spec, tc.name,
			v2BannedBids, v2BannedBids, tc.body, adminHeaders, response, nil)
	}

	invalidDelete := h.do(t, request{
		method:  http.MethodDelete,
		path:    v2BannedBids + "/0",
		headers: adminHeaders,
	})
	if invalidDelete.Code != http.StatusBadRequest {
		t.Fatalf("invalid-delete status = %d; body=%s", invalidDelete.Code, invalidDelete.Body.String())
	}
	pinV2ExchangeGolden(t, spec, "bid_moderation_delete_error_invalid_bid",
		http.MethodDelete, v2BannedBidsItem, v2BannedBids+"/0", "", adminHeaders, invalidDelete)
}

func TestAPIV2BidModerationAdminFailsClosed(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	keyless, err := apiv2.NewServer(
		h.store,
		h.state,
		slog.New(slog.DiscardHandler),
		apiv2.WithClock(func() time.Time { return time.Unix(1767230000, 0) }),
		apiv2.WithAdmin(apiv2.AdminConfig{
			AdminKeys: []common.AdminKey{{Name: "ADMIN_API_KEY", Value: ""}},
		}),
	)
	if err != nil {
		t.Fatalf("building keyless v2 server: %v", err)
	}
	router := routes.New(h.store, routes.Options{V2: keyless})
	body := `{"bidId":2011}`
	request := httptest.NewRequest(http.MethodPost, v2BannedBids, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Admin-Key", adminKey)
	request.RemoteAddr = "10.99.99.2:4242"
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusServiceUnavailable {
		t.Fatalf("keyless status = %d; body=%s", recorder.Code, recorder.Body.String())
	}
	pinV2MutationGolden(t, spec, "bid_moderation_create_error_admin_disabled",
		v2BannedBids, v2BannedBids, body, nil, recorder, nil)
}

func TestAPIV2BidModerationWriteRateLimit(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	restoreV2BidModerationState(t, h)
	t.Cleanup(func() { restoreV2BidModerationState(t, h) })
	const limitedIP = "10.77.77.8:4242"
	headers := map[string]string{"X-Admin-Key": adminKey}
	body := `{"bidId":2011}`

	var limited *httptest.ResponseRecorder
	for range 8 {
		response := h.do(t, request{
			method:     http.MethodPost,
			path:       v2BannedBids,
			body:       strings.NewReader(body),
			headers:    headers,
			remoteAddr: limitedIP,
		})
		switch response.Code {
		case http.StatusCreated, http.StatusConflict:
		case http.StatusTooManyRequests:
			limited = response
		default:
			t.Fatalf("rate-limit probe status = %d; body=%s", response.Code, response.Body.String())
		}
		if limited != nil {
			break
		}
	}
	if limited == nil {
		t.Fatal("moderation create bucket was not exhausted")
	}
	if limited.Header().Get("Retry-After") != "1" {
		t.Fatalf("Retry-After = %q", limited.Header().Get("Retry-After"))
	}
	pinV2MutationGolden(t, spec, "bid_moderation_create_error_rate_limited",
		v2BannedBids, v2BannedBids, body, headers, limited, nil)
}

func TestAPIV2BidModerationCancellationIsOpaque(t *testing.T) {
	h := server(t)
	spec := v2RankingSpec(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	response := h.do(t, request{path: v2BannedBids, ctx: cancelled})
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("cancelled list status = %d; body=%s", response.Code, response.Body.String())
	}
	runV2GoldenCases(t, h, spec, []v2GoldenCase{{
		name:       "bid_moderation_list_error_internal",
		target:     v2BannedBids,
		template:   v2BannedBids,
		pathParams: map[string]string{},
		ctx:        cancelled,
	}})
}
