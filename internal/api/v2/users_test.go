package v2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const userTestBidAddress = "0x2300000000000000000000000000000000000023"

func TestGetCosmicGameUser(t *testing.T) {
	t.Parallel()
	address := userMappingAddress
	server := newUserTestServer(t, fakeUserReader{
		addressID: func(_ context.Context, got string) (int64, error) {
			if got != "0xABcdEFABcdEFabcdEfAbCdefabcdeFABcDEFabCD" {
				t.Fatalf("lookup address = %q", got)
			}
			return 42, nil
		},
		profile: func(_ context.Context, userAid int64) (cgstore.UserProfileRecord, error) {
			if userAid != 42 {
				t.Fatalf("user aid = %d", userAid)
			}
			return validUserProfileRecord(), nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/users/"+address)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d body=%s", response.Code, response.Body.String())
	}
	var profile CosmicGameUserProfile
	decodeResponse(t, response, &profile)
	if profile.Address != "0xABcdEFABcdEFabcdEfAbCdefabcdeFABcDEFabCD" ||
		profile.Bidding.BidCount != 5 ||
		profile.Prizes.PrizeCount != 7 {
		t.Fatalf("profile = %+v", profile)
	}
}

func TestGetCosmicGameUserReturnsZeroProfileWhenUnindexed(t *testing.T) {
	t.Parallel()
	server := newUserTestServer(t, fakeUserReader{
		addressID: func(context.Context, string) (int64, error) {
			return 0, store.ErrNotFound
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/users/"+userMappingAddress)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d body=%s", response.Code, response.Body.String())
	}
	var profile CosmicGameUserProfile
	decodeResponse(t, response, &profile)
	if profile.Address != "0xABcdEFABcdEFabcdEfAbCdefabcdeFABcDEFabCD" ||
		profile.Bidding.BidCount != 0 ||
		profile.Bidding.TotalEthSpentWei != "0" ||
		profile.Bidding.MaxEthBidWei != nil {
		t.Fatalf("zero profile = %+v", profile)
	}
}

func TestGetCosmicGameUserRejectsInvalidAddress(t *testing.T) {
	t.Parallel()
	for _, address := range []string{
		"bad",
		"0x1234",
		"0xzz00000000000000000000000000000000000000",
		"2100000000000000000000000000000000000021",
	} {
		response := serve(t, newUserTestServer(t, fakeUserReader{}),
			"/api/v2/cosmicgame/users/"+address)
		assertProblem(t, response, http.StatusBadRequest)
	}
}

func TestGetCosmicGameUserHidesInternalFailures(t *testing.T) {
	t.Parallel()
	tests := map[string]fakeUserReader{
		"address lookup": {
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("private lookup detail")
			},
		},
		"profile lookup": {
			profile: func(context.Context, int64) (cgstore.UserProfileRecord, error) {
				return cgstore.UserProfileRecord{}, errors.New("private profile detail")
			},
		},
		"malformed profile": {
			profile: func(context.Context, int64) (cgstore.UserProfileRecord, error) {
				record := validUserProfileRecord()
				record.TotalETHWonWei = "-1"
				return record, nil
			},
		},
		"wrong profile identity": {
			profile: func(context.Context, int64) (cgstore.UserProfileRecord, error) {
				record := validUserProfileRecord()
				record.Address = userCursorBob
				return record, nil
			},
		},
	}
	for name, users := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserTestServer(t, users),
				"/api/v2/cosmicgame/users/"+userMappingAddress)
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "private") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}
}

func TestListCosmicGameUserBidsPaginates(t *testing.T) {
	t.Parallel()
	const (
		address = userTestBidAddress
		userAid = int64(23)
	)
	first := userBidRecord(5008)
	second := userBidRecord(5007)
	last := userBidRecord(5006)
	server := newUserTestServer(t, fakeUserReader{
		addressID: func(_ context.Context, got string) (int64, error) {
			if got != address {
				t.Fatalf("lookup address = %q", got)
			}
			return userAid, nil
		},
		bids: func(
			_ context.Context,
			gotAid int64,
			after *cgstore.UserBidPageCursor,
			limit int,
		) ([]cgmodel.CGBidRec, bool, error) {
			if gotAid != userAid || limit != 2 {
				t.Fatalf("page args = aid %d limit %d", gotAid, limit)
			}
			if after == nil {
				return []cgmodel.CGBidRec{first, second}, true, nil
			}
			if after.EventLogID != second.Tx.EvtLogId {
				t.Fatalf("after = %+v", after)
			}
			return []cgmodel.CGBidRec{last}, false, nil
		},
	})

	target := "/api/v2/cosmicgame/users/" + address + "/bids?limit=2"
	firstResponse := serve(t, server, target)
	if firstResponse.Code != http.StatusOK {
		t.Fatalf("first status = %d body=%s", firstResponse.Code, firstResponse.Body.String())
	}
	var firstPage CosmicGameUserBidPage
	decodeResponse(t, firstResponse, &firstPage)
	if len(firstPage.Data) != 2 || firstPage.Meta.NextCursor == nil ||
		firstPage.Data[0].EventLogId != 5008 || firstPage.Data[1].EventLogId != 5007 {
		t.Fatalf("first page = %+v", firstPage)
	}
	cursor, err := decodeUserBidCursor(*firstPage.Meta.NextCursor, address)
	if err != nil || cursor.EventLogID != 5007 {
		t.Fatalf("next cursor = %+v, %v", cursor, err)
	}

	secondResponse := serve(t, server, target+"&cursor="+*firstPage.Meta.NextCursor)
	if secondResponse.Code != http.StatusOK {
		t.Fatalf("second status = %d body=%s", secondResponse.Code, secondResponse.Body.String())
	}
	var secondPage CosmicGameUserBidPage
	decodeResponse(t, secondResponse, &secondPage)
	if len(secondPage.Data) != 1 || secondPage.Data[0].EventLogId != 5006 ||
		secondPage.Meta.NextCursor != nil {
		t.Fatalf("second page = %+v", secondPage)
	}
}

func TestListCosmicGameUserBidsReturnsEmptyForUnindexedWallet(t *testing.T) {
	t.Parallel()
	server := newUserTestServer(t, fakeUserReader{
		addressID: func(context.Context, string) (int64, error) {
			return 0, store.ErrNotFound
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/bids?limit=3")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d body=%s", response.Code, response.Body.String())
	}
	var page CosmicGameUserBidPage
	decodeResponse(t, response, &page)
	if page.Data == nil || len(page.Data) != 0 || page.Meta.Limit != 3 ||
		page.Meta.NextCursor != nil {
		t.Fatalf("page = %+v", page)
	}
}

func TestListCosmicGameUserBidsRejectsInput(t *testing.T) {
	t.Parallel()
	crossUser, err := encodeUserBidCursor(userBidCursor{
		Version: userBidCursorVersion, Address: userCursorBob, EventLogID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	tests := []string{
		"/api/v2/cosmicgame/users/bad/bids",
		"/api/v2/cosmicgame/users/" + userCursorAlice + "/bids?limit=0",
		"/api/v2/cosmicgame/users/" + userCursorAlice + "/bids?limit=201",
		"/api/v2/cosmicgame/users/" + userCursorAlice + "/bids?limit=wat",
		"/api/v2/cosmicgame/users/" + userCursorAlice + "/bids?cursor=bad",
		"/api/v2/cosmicgame/users/" + userCursorAlice + "/bids?cursor=" + crossUser,
	}
	for _, target := range tests {
		response := serve(t, newUserTestServer(t, fakeUserReader{}), target)
		assertProblem(t, response, http.StatusBadRequest)
	}
}

func TestListCosmicGameUserBidsRejectsRepositoryViolations(t *testing.T) {
	t.Parallel()
	const address = userTestBidAddress
	valid := userBidRecord(5008)
	wrongUser := valid
	wrongUser.BidderAid = 99
	wrongAddress := valid
	wrongAddress.BidderAddr = userCursorBob
	malformed := valid
	malformed.Tx.TxHash = "bad"
	tests := map[string]struct {
		records []cgmodel.CGBidRec
		hasMore bool
		err     error
		limit   int
	}{
		"repository error": {err: errors.New("private bid detail")},
		"too many":         {records: []cgmodel.CGBidRec{valid, valid}},
		"wrong user ID":    {records: []cgmodel.CGBidRec{wrongUser}},
		"wrong address":    {records: []cgmodel.CGBidRec{wrongAddress}},
		"unordered": {
			records: []cgmodel.CGBidRec{
				userBidRecord(5007),
				userBidRecord(5008),
			},
			limit: 2,
		},
		"malformed bid":        {records: []cgmodel.CGBidRec{malformed}},
		"empty page with more": {records: []cgmodel.CGBidRec{}, hasMore: true},
	}
	for name, result := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			limit := result.limit
			if limit == 0 {
				limit = 1
			}
			server := newUserTestServer(t, fakeUserReader{
				addressID: func(context.Context, string) (int64, error) {
					return 23, nil
				},
				bids: func(
					context.Context,
					int64,
					*cgstore.UserBidPageCursor,
					int,
				) ([]cgmodel.CGBidRec, bool, error) {
					return result.records, result.hasMore, result.err
				},
			})
			response := serve(t, server,
				fmt.Sprintf("/api/v2/cosmicgame/users/%s/bids?limit=%d", address, limit))
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "private") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}
}

func TestUserHandlersHonorCancelledContext(t *testing.T) {
	t.Parallel()
	users := fakeUserReader{
		addressID: func(ctx context.Context, _ string) (int64, error) {
			return 0, ctx.Err()
		},
	}
	server := newUserTestServer(t, users)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	for _, target := range []string{
		"/api/v2/cosmicgame/users/" + userCursorAlice,
		"/api/v2/cosmicgame/users/" + userCursorAlice + "/bids",
	} {
		response := serveUserContext(t, server, ctx, target)
		assertProblem(t, response, http.StatusInternalServerError)
	}
}

func TestUserResponsesAreDeterministic(t *testing.T) {
	t.Parallel()
	server := newUserTestServer(t, fakeUserReader{
		profile: func(context.Context, int64) (cgstore.UserProfileRecord, error) {
			return validUserProfileRecord(), nil
		},
	})
	target := "/api/v2/cosmicgame/users/" + userMappingAddress
	first := serve(t, server, target)
	second := serve(t, server, target)
	if first.Code != second.Code || !bytes.Equal(first.Body.Bytes(), second.Body.Bytes()) {
		t.Fatalf("nondeterministic responses: %d %s / %d %s",
			first.Code, first.Body.String(), second.Code, second.Body.String())
	}
}

func newUserTestServer(t *testing.T, users userReader) *Server {
	t.Helper()
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeStatisticsReader{},
		fakeBiddingAnalyticsReader{},
		fakeContractAddressReader{},
		fakeParticipantReader{},
		users,
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func userBidRecord(eventLogID int64) cgmodel.CGBidRec {
	record := validBidRecord()
	record.BidderAid = 23
	record.BidderAddr = userTestBidAddress
	record.Tx.EvtLogId = eventLogID
	return record
}

func serveUserContext(
	t *testing.T,
	server *Server,
	ctx context.Context,
	target string,
) *httptest.ResponseRecorder {
	t.Helper()
	router := httpx.NewRouter()
	server.RegisterRoutes(router)
	request := httptest.NewRequest(http.MethodGet, target, nil).WithContext(ctx)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}
