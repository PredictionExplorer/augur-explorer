package v2

import (
	"context"
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const (
	bidModerationKey     = "moderation-secret"
	bidModerationAddress = "0x52908400098527886E0F7030069857D2E4169EE7"
)

var bidModerationNow = time.Unix(1767230000, 0).UTC()

func generousAdminLimits() *AdminWriteLimits {
	return &AdminWriteLimits{
		CreateBidBan: RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
		DeleteBidBan: RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
	}
}

func newBidModerationTestServer(t *testing.T, bids fakeBidReader, config AdminConfig) *Server {
	t.Helper()
	if config.WriteLimits == nil {
		config.WriteLimits = generousAdminLimits()
	}
	server := newTestServer(t, bids)
	server.adminConfig = config
	server.now = func() time.Time { return bidModerationNow }
	return server
}

func moderationAdminConfig() AdminConfig {
	return AdminConfig{
		AdminKeys:   []common.AdminKey{{Name: "ADMIN_API_KEY", Value: bidModerationKey}},
		WriteLimits: generousAdminLimits(),
	}
}

func moderationHeaders() map[string]string {
	return map[string]string{"X-Admin-Key": bidModerationKey}
}

func serveDeleteBidBan(
	t *testing.T,
	server *Server,
	bidID string,
	headers map[string]string,
) *httptest.ResponseRecorder {
	t.Helper()
	router := httpx.NewRouter()
	server.RegisterRoutes(router)
	request := httptest.NewRequest(http.MethodDelete, bannedBidsInstance+"/"+bidID, nil)
	for name, value := range headers {
		request.Header.Set(name, value)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func TestMapBannedBid(t *testing.T) {
	t.Parallel()
	record := cgmodel.CGBannedBidRec{
		Id:        7,
		BidId:     42,
		UserAddr:  strings.ToLower(bidModerationAddress),
		CreatedAt: bidModerationNow.Unix(),
	}
	got, err := mapBannedBid(record)
	if err != nil {
		t.Fatalf("mapBannedBid: %v", err)
	}
	if got.BidId != 42 || got.BidderAddress != bidModerationAddress || !got.BannedAt.Equal(bidModerationNow) {
		t.Fatalf("mapped ban = %+v", got)
	}

	for name, mutate := range map[string]func(*cgmodel.CGBannedBidRec){
		"row id":    func(rec *cgmodel.CGBannedBidRec) { rec.Id = 0 },
		"bid id":    func(rec *cgmodel.CGBannedBidRec) { rec.BidId = 0 },
		"address":   func(rec *cgmodel.CGBannedBidRec) { rec.UserAddr = "not-an-address" },
		"pre epoch": func(rec *cgmodel.CGBannedBidRec) { rec.CreatedAt = -1 },
		"timestamp": func(rec *cgmodel.CGBannedBidRec) { rec.CreatedAt = math.MaxInt64 },
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bad := record
			mutate(&bad)
			if _, err := mapBannedBid(bad); err == nil {
				t.Fatalf("mapBannedBid accepted %+v", bad)
			}
		})
	}
}

func TestListCosmicGameBannedBids(t *testing.T) {
	t.Parallel()
	server := newBidModerationTestServer(t, fakeBidReader{
		bannedPage: func(_ context.Context, after *cgstore.BannedBidPageCursor, limit int) ([]cgmodel.CGBannedBidRec, bool, error) {
			if after != nil || limit != 2 {
				t.Errorf("BannedBidsPage(after=%+v, limit=%d)", after, limit)
			}
			return []cgmodel.CGBannedBidRec{
				{Id: 9, BidId: 2002, UserAddr: bidModerationAddress, CreatedAt: bidModerationNow.Unix()},
				{Id: 8, BidId: 2001, UserAddr: bidModerationAddress, CreatedAt: bidModerationNow.Add(-time.Second).Unix()},
			}, true, nil
		},
	}, AdminConfig{})
	response := serve(t, server, bannedBidsInstance+"?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
	}
	var page BannedBidPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.Limit != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeBidBanCursor(*page.Meta.NextCursor)
	if err != nil || cursor.ID != 8 {
		t.Fatalf("next cursor = %+v, %v", cursor, err)
	}

	t.Run("continuation", func(t *testing.T) {
		t.Parallel()
		encoded, err := encodeBidBanCursor(bidBanCursor{Version: bidBanCursorVersion, ID: 8})
		if err != nil {
			t.Fatal(err)
		}
		continued := newBidModerationTestServer(t, fakeBidReader{
			bannedPage: func(_ context.Context, after *cgstore.BannedBidPageCursor, limit int) ([]cgmodel.CGBannedBidRec, bool, error) {
				if after == nil || after.ID != 8 || limit != 50 {
					t.Errorf("BannedBidsPage(after=%+v, limit=%d)", after, limit)
				}
				return []cgmodel.CGBannedBidRec{}, false, nil
			},
		}, AdminConfig{})
		got := serve(t, continued, bannedBidsInstance+"?cursor="+encoded)
		if got.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", got.Code, got.Body.String())
		}
		var empty BannedBidPage
		decodeResponse(t, got, &empty)
		if empty.Data == nil || len(empty.Data) != 0 || empty.Meta.NextCursor != nil {
			t.Fatalf("empty page = %+v", empty)
		}
	})
}

func TestListCosmicGameBannedBidsFailures(t *testing.T) {
	t.Parallel()
	for name, target := range map[string]string{
		"zero limit":      bannedBidsInstance + "?limit=0",
		"oversized limit": bannedBidsInstance + "?limit=201",
		"invalid cursor":  bannedBidsInstance + "?cursor=not-a-cursor",
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newBidModerationTestServer(t, fakeBidReader{}, AdminConfig{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}

	cases := []struct {
		name string
		page func(context.Context, *cgstore.BannedBidPageCursor, int) ([]cgmodel.CGBannedBidRec, bool, error)
	}{
		{"store error", func(context.Context, *cgstore.BannedBidPageCursor, int) ([]cgmodel.CGBannedBidRec, bool, error) {
			return nil, false, errors.New("database password leaked")
		}},
		{"over cardinality", func(context.Context, *cgstore.BannedBidPageCursor, int) ([]cgmodel.CGBannedBidRec, bool, error) {
			rows := make([]cgmodel.CGBannedBidRec, 51)
			return rows, false, nil
		}},
		{"unordered", func(context.Context, *cgstore.BannedBidPageCursor, int) ([]cgmodel.CGBannedBidRec, bool, error) {
			return []cgmodel.CGBannedBidRec{
				{Id: 2, BidId: 2, UserAddr: bidModerationAddress, CreatedAt: 1},
				{Id: 3, BidId: 3, UserAddr: bidModerationAddress, CreatedAt: 1},
			}, false, nil
		}},
		{"invalid record", func(context.Context, *cgstore.BannedBidPageCursor, int) ([]cgmodel.CGBannedBidRec, bool, error) {
			return []cgmodel.CGBannedBidRec{{Id: 2, BidId: 2, UserAddr: "bad", CreatedAt: 1}}, false, nil
		}},
		{"more without row", func(context.Context, *cgstore.BannedBidPageCursor, int) ([]cgmodel.CGBannedBidRec, bool, error) {
			return []cgmodel.CGBannedBidRec{}, true, nil
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newBidModerationTestServer(t, fakeBidReader{bannedPage: tc.page}, AdminConfig{}), bannedBidsInstance)
			assertProblemKind(t, response, http.StatusInternalServerError, "internal")
			if strings.Contains(response.Body.String(), "password") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}
}

func TestCreateCosmicGameBidBan(t *testing.T) {
	t.Parallel()
	var created bool
	server := newBidModerationTestServer(t, fakeBidReader{
		bidderAddress: func(_ context.Context, bidID int64) (string, error) {
			if bidID != 2001 {
				t.Errorf("bid id = %d", bidID)
			}
			return strings.ToLower(bidModerationAddress), nil
		},
		createBan: func(_ context.Context, bidID int64, address string, bannedAt time.Time) (cgmodel.CGBannedBidRec, error) {
			created = true
			if bidID != 2001 || address != bidModerationAddress || !bannedAt.Equal(bidModerationNow) {
				t.Errorf("CreateBannedBid(%d, %q, %s)", bidID, address, bannedAt)
			}
			return cgmodel.CGBannedBidRec{
				Id: 7, BidId: bidID, UserAddr: address, CreatedAt: bannedAt.Unix(),
			}, nil
		},
	}, moderationAdminConfig())
	response := servePost(t, server, bannedBidsInstance, `{"bidId":2001}`, moderationHeaders())
	if response.Code != http.StatusCreated {
		t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
	}
	if !created {
		t.Fatal("bid ban was not created")
	}
	var ban BannedBid
	decodeResponse(t, response, &ban)
	if ban.BidId != 2001 || ban.BidderAddress != bidModerationAddress ||
		!ban.BannedAt.Equal(bidModerationNow) {
		t.Fatalf("ban = %+v", ban)
	}
}

func TestCreateCosmicGameBidBanFailures(t *testing.T) {
	t.Parallel()
	config := moderationAdminConfig()
	t.Run("request validation", func(t *testing.T) {
		t.Parallel()
		for name, body := range map[string]string{
			"malformed": `{"bidId":"secret-value"}`,
			"missing":   `{}`,
			"zero":      `{"bidId":0}`,
		} {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				response := servePost(t, newBidModerationTestServer(t, fakeBidReader{}, config),
					bannedBidsInstance, body, moderationHeaders())
				assertProblem(t, response, http.StatusBadRequest)
				if strings.Contains(response.Body.String(), "secret-value") {
					t.Fatalf("request value leaked: %s", response.Body.String())
				}
			})
		}
	})

	cases := []struct {
		name   string
		bids   fakeBidReader
		status int
		kind   string
	}{
		{"missing bid", fakeBidReader{
			bidderAddress: func(context.Context, int64) (string, error) { return "", store.ErrNotFound },
		}, http.StatusNotFound, "bid-not-found"},
		{"lookup error", fakeBidReader{
			bidderAddress: func(context.Context, int64) (string, error) { return "", errors.New("lookup secret") },
		}, http.StatusInternalServerError, "internal"},
		{"bad stored address", fakeBidReader{
			bidderAddress: func(context.Context, int64) (string, error) { return "bad", nil },
		}, http.StatusInternalServerError, "internal"},
		{"already banned", fakeBidReader{
			bidderAddress: func(context.Context, int64) (string, error) { return bidModerationAddress, nil },
			createBan: func(context.Context, int64, string, time.Time) (cgmodel.CGBannedBidRec, error) {
				return cgmodel.CGBannedBidRec{}, store.ErrConflict
			},
		}, http.StatusConflict, "bid-already-banned"},
		{"write error", fakeBidReader{
			bidderAddress: func(context.Context, int64) (string, error) { return bidModerationAddress, nil },
			createBan: func(context.Context, int64, string, time.Time) (cgmodel.CGBannedBidRec, error) {
				return cgmodel.CGBannedBidRec{}, errors.New("write secret")
			},
		}, http.StatusInternalServerError, "internal"},
		{"wrong returned identity", fakeBidReader{
			bidderAddress: func(context.Context, int64) (string, error) { return bidModerationAddress, nil },
			createBan: func(context.Context, int64, string, time.Time) (cgmodel.CGBannedBidRec, error) {
				return cgmodel.CGBannedBidRec{
					Id: 1, BidId: 2002, UserAddr: bidModerationAddress, CreatedAt: 1,
				}, nil
			},
		}, http.StatusInternalServerError, "internal"},
		{"invalid returned record", fakeBidReader{
			bidderAddress: func(context.Context, int64) (string, error) { return bidModerationAddress, nil },
			createBan: func(context.Context, int64, string, time.Time) (cgmodel.CGBannedBidRec, error) {
				return cgmodel.CGBannedBidRec{
					Id: 0, BidId: 2001, UserAddr: bidModerationAddress, CreatedAt: 1,
				}, nil
			},
		}, http.StatusInternalServerError, "internal"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			response := servePost(t, newBidModerationTestServer(t, tc.bids, config),
				bannedBidsInstance, `{"bidId":2001}`, moderationHeaders())
			assertProblemKind(t, response, tc.status, tc.kind)
			if strings.Contains(response.Body.String(), "secret") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}
}

func TestDeleteCosmicGameBidBan(t *testing.T) {
	t.Parallel()
	var removedID int64
	server := newBidModerationTestServer(t, fakeBidReader{
		removeBan: func(_ context.Context, bidID int64) (bool, error) {
			removedID = bidID
			return true, nil
		},
	}, moderationAdminConfig())
	response := serveDeleteBidBan(t, server, "2001", moderationHeaders())
	if response.Code != http.StatusNoContent || response.Body.Len() != 0 || removedID != 2001 {
		t.Fatalf("delete response = %d %q, removed=%d", response.Code, response.Body.String(), removedID)
	}

	for name, tc := range map[string]struct {
		id     string
		remove func(context.Context, int64) (bool, error)
		status int
		kind   string
	}{
		"invalid":   {"0", nil, http.StatusBadRequest, "invalid-bid"},
		"malformed": {"not-a-number", nil, http.StatusBadRequest, "invalid-request"},
		"missing":   {"2001", func(context.Context, int64) (bool, error) { return false, nil }, http.StatusNotFound, "bid-ban-not-found"},
		"failure":   {"2001", func(context.Context, int64) (bool, error) { return false, errors.New("delete secret") }, http.StatusInternalServerError, "internal"},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := serveDeleteBidBan(t,
				newBidModerationTestServer(t, fakeBidReader{removeBan: tc.remove}, moderationAdminConfig()),
				tc.id, moderationHeaders())
			assertProblemKind(t, got, tc.status, tc.kind)
			if strings.Contains(got.Body.String(), "secret") {
				t.Fatalf("internal detail leaked: %s", got.Body.String())
			}
		})
	}
}

func TestBidModerationAdminAuthentication(t *testing.T) {
	t.Parallel()
	refuse := fakeBidReader{
		bidderAddress: func(context.Context, int64) (string, error) {
			t.Error("handler ran without valid authentication")
			return "", nil
		},
	}
	body := `{"bidId":2001}`

	t.Run("unconfigured fails closed", func(t *testing.T) {
		t.Parallel()
		server := newBidModerationTestServer(t, refuse, AdminConfig{
			AdminKeys: []common.AdminKey{{Name: "ADMIN_API_KEY", Value: "  "}},
		})
		response := servePost(t, server, bannedBidsInstance, body, moderationHeaders())
		problem := assertProblemKind(t, response, http.StatusServiceUnavailable, "admin-disabled")
		if problem.Detail == nil || !strings.Contains(*problem.Detail, "ADMIN_API_KEY") {
			t.Fatalf("problem = %+v", problem)
		}
	})
	t.Run("missing key", func(t *testing.T) {
		t.Parallel()
		response := servePost(t, newBidModerationTestServer(t, refuse, moderationAdminConfig()),
			bannedBidsInstance, body, nil)
		assertProblemKind(t, response, http.StatusUnauthorized, "unauthorized")
	})
	t.Run("wrong key", func(t *testing.T) {
		t.Parallel()
		response := servePost(t, newBidModerationTestServer(t, refuse, moderationAdminConfig()),
			bannedBidsInstance, body, map[string]string{"X-Admin-Key": "wrong"})
		assertProblemKind(t, response, http.StatusUnauthorized, "unauthorized")
	})
	t.Run("delete uses the same scheme", func(t *testing.T) {
		t.Parallel()
		configured := newBidModerationTestServer(t, fakeBidReader{
			removeBan: func(context.Context, int64) (bool, error) {
				t.Error("delete handler ran without valid authentication")
				return false, nil
			},
		}, moderationAdminConfig())
		assertProblemKind(t, serveDeleteBidBan(t, configured, "2001", nil),
			http.StatusUnauthorized, "unauthorized")

		keyless := newBidModerationTestServer(t, fakeBidReader{}, AdminConfig{
			AdminKeys: []common.AdminKey{{Name: "ADMIN_API_KEY", Value: ""}},
		})
		assertProblemKind(t, serveDeleteBidBan(t, keyless, "2001", moderationHeaders()),
			http.StatusServiceUnavailable, "admin-disabled")
	})
	t.Run("public list bypasses auth", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newBidModerationTestServer(t, fakeBidReader{}, AdminConfig{}), bannedBidsInstance)
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d; body=%s", response.Code, response.Body.String())
		}
	})
}

func TestBidModerationRateLimits(t *testing.T) {
	t.Parallel()
	config := moderationAdminConfig()
	config.WriteLimits = &AdminWriteLimits{
		CreateBidBan: RateLimitSpec{PerSecond: 0.001, Burst: 1},
		DeleteBidBan: RateLimitSpec{PerSecond: 10_000, Burst: 10_000},
	}
	server := newBidModerationTestServer(t, fakeBidReader{
		bidderAddress: func(context.Context, int64) (string, error) { return bidModerationAddress, nil },
		createBan: func(_ context.Context, bidID int64, address string, at time.Time) (cgmodel.CGBannedBidRec, error) {
			return cgmodel.CGBannedBidRec{Id: bidID, BidId: bidID, UserAddr: address, CreatedAt: at.Unix()}, nil
		},
		removeBan: func(context.Context, int64) (bool, error) { return true, nil },
	}, config)
	router := httpx.NewRouter()
	server.RegisterRoutes(router)
	do := func(method, target, body string) *httptest.ResponseRecorder {
		var reader *strings.Reader
		if body != "" {
			reader = strings.NewReader(body)
		} else {
			reader = strings.NewReader("")
		}
		request := httptest.NewRequest(method, target, reader)
		request.Header.Set("X-Admin-Key", bidModerationKey)
		if body != "" {
			request.Header.Set("Content-Type", "application/json")
		}
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		return response
	}
	if first := do(http.MethodPost, bannedBidsInstance, `{"bidId":2001}`); first.Code != http.StatusCreated {
		t.Fatalf("first create = %d %s", first.Code, first.Body.String())
	}
	second := do(http.MethodPost, bannedBidsInstance, `{"bidId":2002}`)
	assertProblemKind(t, second, http.StatusTooManyRequests, "rate-limited")
	if second.Header().Get("Retry-After") != "1" {
		t.Fatalf("Retry-After = %q", second.Header().Get("Retry-After"))
	}
	// The delete operation has an independent bucket.
	if deleted := do(http.MethodDelete, bannedBidsInstance+"/2001", ""); deleted.Code != http.StatusNoContent {
		t.Fatalf("delete = %d %s", deleted.Code, deleted.Body.String())
	}

	defaults := defaultAdminWriteLimits()
	if defaults.CreateBidBan != (RateLimitSpec{PerSecond: 2, Burst: 5}) ||
		defaults.DeleteBidBan != (RateLimitSpec{PerSecond: 2, Burst: 5}) ||
		(AdminConfig{}).writeLimits() != defaults {
		t.Fatalf("default moderation limits = %+v", defaults)
	}
}

func TestConfiguredAdminSecret(t *testing.T) {
	t.Parallel()
	secret, names := configuredAdminSecret([]common.AdminKey{
		{Name: "PRIMARY", Value: " "},
		{Name: "FALLBACK", Value: " value "},
		{Name: "IGNORED", Value: "other"},
	})
	if secret != "value" || strings.Join(names, ",") != "PRIMARY,FALLBACK,IGNORED" {
		t.Fatalf("configuredAdminSecret = %q, %v", secret, names)
	}
	if ethcommon.HexToAddress(strings.ToLower(bidModerationAddress)).Hex() != bidModerationAddress {
		t.Fatal("test address is not a stable checksum fixture")
	}
}

// TestBidModerationGeneratedMiddlewareResponses pins the generated response
// visitors for statuses normally written directly by strict middleware. This
// keeps the OpenAPI contract executable even though auth/rate middleware
// short-circuits before a handler can return these response objects.
func TestBidModerationGeneratedMiddlewareResponses(t *testing.T) {
	t.Parallel()
	problem := newProblem(
		http.StatusUnauthorized,
		"unauthorized",
		"Unauthorized",
		"The admin key is missing or wrong.",
		bannedBidsInstance,
	)
	rateLimited := RateLimitedApplicationProblemPlusJSONResponse{
		Body:    rateLimitedProblem(bannedBidsInstance),
		Headers: RateLimitedResponseHeaders{RetryAfter: 1},
	}
	disabled := AdminDisabledApplicationProblemPlusJSONResponse(newProblem(
		http.StatusServiceUnavailable,
		"admin-disabled",
		"Admin operation disabled",
		"Admin authentication is not configured.",
		bannedBidsInstance,
	))

	createCases := []struct {
		name   string
		status int
		visit  func(http.ResponseWriter) error
	}{
		{
			"unauthorized",
			http.StatusUnauthorized,
			func(w http.ResponseWriter) error {
				return (CreateCosmicGameBidBan401ApplicationProblemPlusJSONResponse{
					UnauthorizedApplicationProblemPlusJSONResponse: UnauthorizedApplicationProblemPlusJSONResponse(problem),
				}).VisitCreateCosmicGameBidBanResponse(w)
			},
		},
		{
			"rate limited",
			http.StatusTooManyRequests,
			func(w http.ResponseWriter) error {
				return (CreateCosmicGameBidBan429ApplicationProblemPlusJSONResponse{
					RateLimitedApplicationProblemPlusJSONResponse: rateLimited,
				}).VisitCreateCosmicGameBidBanResponse(w)
			},
		},
		{
			"disabled",
			http.StatusServiceUnavailable,
			func(w http.ResponseWriter) error {
				return (CreateCosmicGameBidBan503ApplicationProblemPlusJSONResponse{
					AdminDisabledApplicationProblemPlusJSONResponse: disabled,
				}).VisitCreateCosmicGameBidBanResponse(w)
			},
		},
	}
	for _, tc := range createCases {
		t.Run("create "+tc.name, func(t *testing.T) {
			t.Parallel()
			response := httptest.NewRecorder()
			if err := tc.visit(response); err != nil {
				t.Fatal(err)
			}
			if response.Code != tc.status || response.Header().Get("Content-Type") != "application/problem+json" {
				t.Fatalf("response = %d %v", response.Code, response.Header())
			}
			if tc.status == http.StatusTooManyRequests && response.Header().Get("Retry-After") != "1" {
				t.Fatalf("Retry-After = %q", response.Header().Get("Retry-After"))
			}
		})
	}

	deleteCases := []struct {
		name   string
		status int
		visit  func(http.ResponseWriter) error
	}{
		{
			"unauthorized",
			http.StatusUnauthorized,
			func(w http.ResponseWriter) error {
				return (DeleteCosmicGameBidBan401ApplicationProblemPlusJSONResponse{
					UnauthorizedApplicationProblemPlusJSONResponse: UnauthorizedApplicationProblemPlusJSONResponse(problem),
				}).VisitDeleteCosmicGameBidBanResponse(w)
			},
		},
		{
			"rate limited",
			http.StatusTooManyRequests,
			func(w http.ResponseWriter) error {
				return (DeleteCosmicGameBidBan429ApplicationProblemPlusJSONResponse{
					RateLimitedApplicationProblemPlusJSONResponse: rateLimited,
				}).VisitDeleteCosmicGameBidBanResponse(w)
			},
		},
		{
			"disabled",
			http.StatusServiceUnavailable,
			func(w http.ResponseWriter) error {
				return (DeleteCosmicGameBidBan503ApplicationProblemPlusJSONResponse{
					AdminDisabledApplicationProblemPlusJSONResponse: disabled,
				}).VisitDeleteCosmicGameBidBanResponse(w)
			},
		},
	}
	for _, tc := range deleteCases {
		t.Run("delete "+tc.name, func(t *testing.T) {
			t.Parallel()
			response := httptest.NewRecorder()
			if err := tc.visit(response); err != nil {
				t.Fatal(err)
			}
			if response.Code != tc.status || response.Header().Get("Content-Type") != "application/problem+json" {
				t.Fatalf("response = %d %v", response.Code, response.Header())
			}
			if tc.status == http.StatusTooManyRequests && response.Header().Get("Retry-After") != "1" {
				t.Fatalf("Retry-After = %q", response.Header().Get("Retry-After"))
			}
		})
	}
}
