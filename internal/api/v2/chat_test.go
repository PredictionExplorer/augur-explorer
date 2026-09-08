package v2

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func chatRecord(id int64) cgstore.ChatMessageRecord {
	return cgstore.ChatMessageRecord{
		ChatPosition: cgstore.ChatPosition{OccurredAt: time.Unix(1000, 0), EventLogID: id},
		Round:        7, Position: id, BidderAddress: "0x1111111111111111111111111111111111111111", Message: "hello",
		BidType: 0, EthPriceWei: "100000000000000001", CstPriceWei: "-1", RandomWalkTokenID: -1, TransactionHash: "0x" + strings.Repeat("a", 64),
	}
}

func chatToken(t *testing.T, direction string, id, revision int64) string {
	t.Helper()
	position := chatRecord(id).ChatPosition
	if id == 0 {
		position.OccurredAt = time.Unix(0, 0)
	}
	token, err := encodeChatCursor(7, direction, position, revision)
	if err != nil {
		t.Fatal(err)
	}
	return token
}

func TestChatMessagePagesAndSync(t *testing.T) {
	ctx := context.Background()
	limit := 2
	s := newTestServer(t, fakeBidReader{chatMessages: func(_ context.Context, round int64, cursor *cgstore.ChatPosition, after bool, n int) (cgstore.ChatMessagePage, error) {
		if round != 7 || n != 2 {
			t.Fatalf("query scope = %d/%d", round, n)
		}
		out := cgstore.ChatMessagePage{Revision: 3, Latest: chatRecord(5).ChatPosition}
		switch {
		case after:
			if cursor == nil || cursor.EventLogID != 2 {
				t.Fatalf("after boundary = %+v", cursor)
			}
			out.Records = []cgstore.ChatMessageRecord{chatRecord(3), chatRecord(4)}
			out.HasMore = true
		case cursor != nil:
			if cursor.EventLogID != 4 {
				t.Fatalf("older boundary = %+v", cursor)
			}
			out.Records = []cgstore.ChatMessageRecord{chatRecord(3), chatRecord(2)}
		default:
			out.Records = []cgstore.ChatMessageRecord{chatRecord(5), chatRecord(4)}
			out.HasMore = true
		}
		return out, nil
	}})
	first, err := s.ListRoundMessages(ctx, ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{Limit: &limit}})
	if err != nil {
		t.Fatal(err)
	}
	page, ok := first.(ListRoundMessages200JSONResponse)
	if !ok || page.Meta.NextCursor == nil || page.Meta.HasMore || page.Meta.Revision != "3" || page.Data[0].EventLogId != 5 {
		t.Fatalf("first = %#v", first)
	}
	if page.Data[0].EthPriceWei == nil || *page.Data[0].EthPriceWei != "100000000000000001" {
		t.Fatal("lost exact amount")
	}
	if page.Data[0].CstPriceWei != nil || page.Data[0].RandomWalkTokenId != nil {
		t.Fatal("sentinel leaked")
	}
	older, err := s.ListRoundMessages(ctx, ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{Limit: &limit, Cursor: page.Meta.NextCursor}})
	if err != nil {
		t.Fatal(err)
	}
	olderPage := older.(ListRoundMessages200JSONResponse)
	if olderPage.Data[0].EventLogId != 3 || olderPage.Meta.NextCursor != nil {
		t.Fatalf("older = %#v", olderPage)
	}
	after := chatToken(t, "after", 2, 3)
	result, err := s.ListRoundMessages(ctx, ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{Limit: &limit, After: &after}})
	if err != nil {
		t.Fatal(err)
	}
	catchup := result.(ListRoundMessages200JSONResponse)
	next, err := decodeChatCursor(catchup.Meta.SyncCursor, 7, "after")
	if err != nil || next.EventLogID != 4 || !catchup.Meta.HasMore || catchup.Meta.NextCursor != nil {
		t.Fatalf("catchup skipped rows: %#v / %+v / %v", catchup, next, err)
	}
}

func TestChatEmptyAndUnchangedSync(t *testing.T) {
	s := newTestServer(t, fakeBidReader{chatMessages: func(_ context.Context, _ int64, cursor *cgstore.ChatPosition, after bool, limit int) (cgstore.ChatMessagePage, error) {
		if limit != 50 {
			t.Fatalf("default = %d", limit)
		}
		return cgstore.ChatMessagePage{Revision: 1, Latest: cgstore.ChatPosition{OccurredAt: time.Unix(0, 0)}}, nil
	}})
	response, _ := s.ListRoundMessages(context.Background(), ListRoundMessagesRequestObject{Round: 7})
	page := response.(ListRoundMessages200JSONResponse)
	if page.Data == nil || len(page.Data) != 0 || page.Meta.NextCursor != nil {
		t.Fatalf("empty = %#v", page)
	}
	response, _ = s.ListRoundMessages(context.Background(), ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{After: &page.Meta.SyncCursor}})
	if got := response.(ListRoundMessages200JSONResponse); got.Meta.SyncCursor != page.Meta.SyncCursor || got.Meta.HasMore {
		t.Fatalf("empty advanced: %#v", got)
	}
}

func TestChatCursorScopeAndMalformedInputs(t *testing.T) {
	valid := chatToken(t, "older", 1, 3)
	if _, err := decodeChatCursor(valid, 8, "older"); err == nil {
		t.Fatal("accepted different round")
	}
	if _, err := decodeChatCursor(valid, 7, "after"); err == nil {
		t.Fatal("accepted different direction")
	}
	for _, input := range []string{"", "!", strings.Repeat("a", 513), "e30"} {
		if _, err := decodeChatCursor(input, 7, "older"); err == nil {
			t.Errorf("accepted %q", input)
		}
	}
	for _, c := range []chatCursor{
		{Version: 2, Round: 7, Direction: "after", OccurredAt: time.Unix(0, 0), Revision: 1},
		{Version: 1, Round: 7, Direction: "older", OccurredAt: time.Unix(0, 0), Revision: 1},
		{Version: 1, Round: 7, Direction: "after", OccurredAt: time.Unix(1, 0), Revision: 1},
		{Version: 1, Round: 7, Direction: "after", OccurredAt: time.Unix(0, 0), Revision: 0},
	} {
		if validChatCursor(c) {
			t.Errorf("accepted invalid fields: %+v", c)
		}
	}
}

func TestChatProblems(t *testing.T) {
	limit0, limit201 := 0, 201
	older, after := chatToken(t, "older", 1, 1), chatToken(t, "after", 1, 1)
	cases := []struct {
		name    string
		request ListRoundMessagesRequestObject
		page    cgstore.ChatMessagePage
		err     error
		status  int
	}{
		{name: "negative round", request: ListRoundMessagesRequestObject{Round: -1}, status: 400},
		{name: "zero limit", request: ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{Limit: &limit0}}, status: 400},
		{name: "oversized limit", request: ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{Limit: &limit201}}, status: 400},
		{name: "two cursors", request: ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{Cursor: &older, After: &after}}, status: 400},
		{name: "wrong direction", request: ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{After: &older}}, status: 400},
		{name: "stale revision", request: ListRoundMessagesRequestObject{Round: 7, Params: ListRoundMessagesParams{After: &after}}, page: cgstore.ChatMessagePage{Revision: 2}, status: 409},
		{name: "database failure", request: ListRoundMessagesRequestObject{Round: 7}, err: errors.New("private connection details"), status: 500},
		{name: "missing revision", request: ListRoundMessagesRequestObject{Round: 7}, status: 500},
		{name: "unordered", request: ListRoundMessagesRequestObject{Round: 7}, page: cgstore.ChatMessagePage{Revision: 1, Latest: chatRecord(3).ChatPosition, Records: []cgstore.ChatMessageRecord{chatRecord(1), chatRecord(2)}}, status: 500},
		{name: "invalid snapshot", request: ListRoundMessagesRequestObject{Round: 7}, page: cgstore.ChatMessagePage{Revision: 1, Latest: chatRecord(1).ChatPosition, Records: []cgstore.ChatMessageRecord{chatRecord(2)}}, status: 500},
		{name: "empty continuation", request: ListRoundMessagesRequestObject{Round: 7}, page: cgstore.ChatMessagePage{Revision: 1, Latest: chatRecord(1).ChatPosition, HasMore: true}, status: 500},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestServer(t, fakeBidReader{chatMessages: func(context.Context, int64, *cgstore.ChatPosition, bool, int) (cgstore.ChatMessagePage, error) {
				return tc.page, tc.err
			}})
			response, err := s.ListRoundMessages(context.Background(), tc.request)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			if err := response.VisitListRoundMessagesResponse(rr); err != nil {
				t.Fatal(err)
			}
			if rr.Code != tc.status || strings.Contains(rr.Body.String(), "private connection") {
				t.Fatalf("response=%d %s", rr.Code, rr.Body.String())
			}
			if tc.status == 409 && !strings.Contains(rr.Body.String(), "feed-reset-required") {
				t.Fatal("missing reset problem")
			}
		})
	}
}

func TestChatMessageMappingValidation(t *testing.T) {
	for name, change := range map[string]func(*cgstore.ChatMessageRecord){
		"identity":    func(r *cgstore.ChatMessageRecord) { r.EventLogID = 0 },
		"address":     func(r *cgstore.ChatMessageRecord) { r.BidderAddress = "bad" },
		"transaction": func(r *cgstore.ChatMessageRecord) { r.TransactionHash = "bad" },
		"blank":       func(r *cgstore.ChatMessageRecord) { r.Message = "\ufeff\u3000\n" },
		"amount":      func(r *cgstore.ChatMessageRecord) { r.CstPriceWei = "0.1" },
		"ETH amount":  func(r *cgstore.ChatMessageRecord) { r.EthPriceWei = "0.1" },
		"token":       func(r *cgstore.ChatMessageRecord) { r.RandomWalkTokenID = -2 },
	} {
		t.Run(name, func(t *testing.T) {
			r := chatRecord(1)
			change(&r)
			if _, err := mapChatMessage(r); err == nil {
				t.Fatal("accepted invalid record")
			}
		})
	}
	if !chatHasText("\u0085") || chatHasText("\ufeff\t\n \u00a0\u2001\u2028\u3000") {
		t.Fatal("ECMAScript whitespace mismatch")
	}
	r := chatRecord(1)
	r.BidType = 2
	r.EthPriceWei = "-1"
	r.CstPriceWei = "123"
	r.RandomWalkTokenID = 0
	got, err := mapChatMessage(r)
	if err != nil || got.BidType != Cst || got.CstPriceWei == nil || got.RandomWalkTokenId == nil {
		t.Fatalf("mapping=%+v,%v", got, err)
	}
}

func TestChatContextSnapshot(t *testing.T) {
	r := chatRecord(1)
	entry := cgstore.ChatContextRecord{ChatPosition: r.ChatPosition, Round: 7, Position: 1, BidderAddress: r.BidderAddress, BidType: 2, PrizeAt: time.Unix(2000, 0), CstDutchAuctionDurationSeconds: 0}
	s := newTestServer(t, fakeBidReader{chatContext: func(context.Context, int64) (cgstore.ChatContextSnapshot, error) {
		return cgstore.ChatContextSnapshot{Revision: 2, Records: []cgstore.ChatContextRecord{entry}}, nil
	}})
	response, err := s.GetRoundChatContext(context.Background(), GetRoundChatContextRequestObject{Round: 7})
	if err != nil {
		t.Fatal(err)
	}
	page := response.(GetRoundChatContext200JSONResponse)
	if page.Meta.Revision != "2" || len(page.Data) != 1 || page.Data[0].CstDutchAuctionDurationSeconds == nil {
		t.Fatalf("context=%#v", page)
	}
	rr := httptest.NewRecorder()
	_ = response.VisitGetRoundChatContextResponse(rr)
	if rr.Code != http.StatusOK || strings.Contains(rr.Body.String(), "message") || strings.Contains(rr.Body.String(), "reward") {
		t.Fatalf("context leaked fields: %s", rr.Body.String())
	}
	for _, round := range []int64{-1, 8} {
		response, _ := s.GetRoundChatContext(context.Background(), GetRoundChatContextRequestObject{Round: round})
		rr := httptest.NewRecorder()
		_ = response.VisitGetRoundChatContextResponse(rr)
		want := 500
		if round < 0 {
			want = 400
		}
		if rr.Code != want {
			t.Fatalf("round %d got%d", round, rr.Code)
		}
	}
}

func FuzzChatCursor(f *testing.F) {
	f.Add("")
	f.Add("e30")
	f.Add("!")
	f.Fuzz(func(t *testing.T, input string) {
		cursor, err := decodeChatCursor(input, 7, "after")
		if err != nil {
			return
		}
		encoded, err := encodeChatCursor(cursor.Round, cursor.Direction, cgstore.ChatPosition{OccurredAt: cursor.OccurredAt, EventLogID: cursor.EventLogID}, cursor.Revision)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := decodeChatCursor(encoded, 7, "after"); err != nil {
			t.Fatal(err)
		}
	})
}

// A successful repository call still needs validation: malformed snapshots must
// produce a sanitized error rather than a partial page or unusable cursor.
func TestChatRejectsInvalidRepositorySnapshots(t *testing.T) {
	limit := 1
	for _, tc := range []struct {
		name string
		page cgstore.ChatMessagePage
	}{
		{
			name: "more rows than the requested limit",
			page: cgstore.ChatMessagePage{Revision: 1, Latest: chatRecord(2).ChatPosition,
				Records: []cgstore.ChatMessageRecord{chatRecord(2), chatRecord(1)}},
		},
		{
			name: "negative latest boundary",
			page: cgstore.ChatMessagePage{Revision: 1,
				Latest: cgstore.ChatPosition{OccurredAt: time.Unix(1000, 0), EventLogID: -1}},
		},
		{
			name: "message mapping failure",
			page: func() cgstore.ChatMessagePage {
				record := chatRecord(1)
				record.TransactionHash = "private invalid transaction"
				return cgstore.ChatMessagePage{Revision: 1, Latest: record.ChatPosition,
					Records: []cgstore.ChatMessageRecord{record}}
			}(),
		},
		{
			// PostgreSQL timestamps allow years outside JSON's RFC3339 range.
			// Reject the snapshot if its sync boundary cannot be serialized.
			name: "latest boundary outside cursor timestamp range",
			page: cgstore.ChatMessagePage{Revision: 1, Latest: cgstore.ChatPosition{
				OccurredAt: time.Date(10000, time.January, 1, 0, 0, 0, 0, time.UTC), EventLogID: 1}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestServer(t, fakeBidReader{chatMessages: func(context.Context, int64, *cgstore.ChatPosition, bool, int) (cgstore.ChatMessagePage, error) {
				return tc.page, nil
			}})
			response, err := s.ListRoundMessages(context.Background(), ListRoundMessagesRequestObject{
				Round: 7, Params: ListRoundMessagesParams{Limit: &limit},
			})
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			if err := response.VisitListRoundMessagesResponse(rr); err != nil {
				t.Fatal(err)
			}
			if rr.Code != http.StatusInternalServerError || strings.Contains(rr.Body.String(), "private") || strings.Contains(rr.Body.String(), "syncCursor") {
				t.Fatalf("invalid repository snapshot response = %d %s", rr.Code, rr.Body.String())
			}
		})
	}
}

func TestChatContextRejectsCorruptSnapshots(t *testing.T) {
	valid := cgstore.ChatContextRecord{
		ChatPosition: chatRecord(1).ChatPosition, Round: 7, Position: 1,
		BidderAddress: chatRecord(1).BidderAddress, PrizeAt: time.Unix(2000, 0),
		CstDutchAuctionDurationSeconds: -1,
	}
	for _, tc := range []struct {
		name     string
		snapshot cgstore.ChatContextSnapshot
		err      error
	}{
		{name: "repository failure", err: errors.New("private database details")},
		{name: "missing revision", snapshot: cgstore.ChatContextSnapshot{Records: []cgstore.ChatContextRecord{valid}}},
		{name: "invalid identity", snapshot: cgstore.ChatContextSnapshot{Revision: 1, Records: []cgstore.ChatContextRecord{func() cgstore.ChatContextRecord {
			record := valid
			record.BidderAddress = "private invalid address"
			return record
		}()}}},
		{name: "duplicate identity", snapshot: cgstore.ChatContextSnapshot{Revision: 1, Records: []cgstore.ChatContextRecord{valid, valid}}},
		{name: "missing allocation deadline", snapshot: cgstore.ChatContextSnapshot{Revision: 1, Records: []cgstore.ChatContextRecord{func() cgstore.ChatContextRecord {
			record := valid
			record.PrizeAt = time.Time{}
			return record
		}()}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestServer(t, fakeBidReader{chatContext: func(context.Context, int64) (cgstore.ChatContextSnapshot, error) {
				return tc.snapshot, tc.err
			}})
			response, err := s.GetRoundChatContext(context.Background(), GetRoundChatContextRequestObject{Round: 7})
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			if err := response.VisitGetRoundChatContextResponse(rr); err != nil {
				t.Fatal(err)
			}
			if rr.Code != http.StatusInternalServerError || strings.Contains(rr.Body.String(), "private") {
				t.Fatalf("invalid context response = %d %s", rr.Code, rr.Body.String())
			}
		})
	}
}

func TestEncodeChatCursorRejectsInvalidBoundary(t *testing.T) {
	for _, tc := range []struct {
		name      string
		direction string
		position  cgstore.ChatPosition
		revision  int64
	}{
		{name: "empty older boundary", direction: "older", position: cgstore.ChatPosition{OccurredAt: time.Unix(0, 0)}, revision: 1},
		{name: "zero identity after nonempty time", direction: "after", position: cgstore.ChatPosition{OccurredAt: time.Unix(1, 0)}, revision: 1},
		{name: "missing generation", direction: "after", position: chatRecord(1).ChatPosition},
	} {
		t.Run(tc.name, func(t *testing.T) {
			encoded, err := encodeChatCursor(7, tc.direction, tc.position, tc.revision)
			if encoded != "" || !errors.Is(err, errInvalidCursor) {
				t.Fatalf("invalid cursor encoded as %q with error %v", encoded, err)
			}
		})
	}
}
