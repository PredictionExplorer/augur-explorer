package v2

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func newUserActivityTestServer(t *testing.T, activity userActivityReader) *Server {
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
		fakeUserReader{},
		fakeUserHistoryReader{},
		fakeUserStakingReader{},
		activity,
		fakeGlobalDirectoryReader{},
		fakeGlobalStakingReader{},
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func ownedTokenAt(tokenID int64) cgstore.UserOwnedTokenRecord {
	record := validUserOwnedTokenRecord()
	record.TokenID = tokenID
	return record
}

func TestListUserOwnedTokensPaginates(t *testing.T) {
	t.Parallel()

	var gotAfter *cgstore.UserTokenPageCursor
	var gotLimit int
	server := newUserActivityTestServer(t, fakeUserActivityReader{
		ownedTokens: func(_ context.Context, userAid int64, after *cgstore.UserTokenPageCursor, limit int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
			if userAid != 1 {
				t.Errorf("user aid = %d", userAid)
			}
			gotAfter, gotLimit = after, limit
			return []cgstore.UserOwnedTokenRecord{ownedTokenAt(1), ownedTokenAt(6)}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-signature-tokens?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter != nil || gotLimit != 2 {
		t.Fatalf("repository args = (%+v, %d)", gotAfter, gotLimit)
	}
	var page CosmicGameUserCosmicSignatureTokenPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	if page.Data[0].NftTokenId != 1 || page.Data[1].NftTokenId != 6 {
		t.Fatalf("token ids = %d, %d", page.Data[0].NftTokenId, page.Data[1].NftTokenId)
	}
	cursor, err := decodeUserOwnedTokenCursor(*page.Meta.NextCursor, userCursorAlice)
	if err != nil || cursor.TokenID != 6 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}

	// The continuation cursor decodes into the repository cursor.
	server = newUserActivityTestServer(t, fakeUserActivityReader{
		ownedTokens: func(_ context.Context, _ int64, after *cgstore.UserTokenPageCursor, _ int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
			gotAfter = after
			return []cgstore.UserOwnedTokenRecord{}, false, nil
		},
	})
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-signature-tokens?cursor="+*page.Meta.NextCursor)
	if response.Code != http.StatusOK {
		t.Fatalf("continuation status = %d", response.Code)
	}
	if gotAfter == nil || gotAfter.TokenID != 6 {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
}

func TestListUserOwnedTokensEmptyForUnindexedWallet(t *testing.T) {
	t.Parallel()

	server := newUserActivityTestServer(t, fakeUserActivityReader{
		addressID: func(context.Context, string) (int64, error) {
			return 0, store.ErrNotFound
		},
		ownedTokens: func(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
			t.Error("owned tokens fetched for an unindexed wallet")
			return nil, false, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-signature-tokens")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	var page CosmicGameUserCosmicSignatureTokenPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 0 || page.Meta.NextCursor != nil {
		t.Fatalf("page = %+v", page)
	}
}

func TestListUserOwnedTokensRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	bobCursor, err := encodeUserOwnedTokenCursor(userOwnedTokenCursor{
		Version: userOwnedTokenCursorVersion, Address: userCursorBob, TokenID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	base := "/api/v2/cosmicgame/users/" + userCursorAlice + "/cosmic-signature-tokens"
	tests := map[string]string{
		"invalid address":   "/api/v2/cosmicgame/users/not-an-address/cosmic-signature-tokens",
		"invalid limit":     base + "?limit=201",
		"bind limit":        base + "?limit=wat",
		"malformed cursor":  base + "?cursor=bad",
		"cross-user cursor": base + "?cursor=" + bobCursor,
	}
	server := newUserActivityTestServer(t, fakeUserActivityReader{})
	for name, target := range tests {
		response := serve(t, server, target)
		if response.Code != http.StatusBadRequest {
			t.Errorf("%s: status = %d", name, response.Code)
		}
		if contentType := response.Header().Get("Content-Type"); contentType != "application/problem+json" {
			t.Errorf("%s: content type = %q", name, contentType)
		}
	}
}

func TestListUserOwnedTokensInternalFailures(t *testing.T) {
	t.Parallel()

	target := "/api/v2/cosmicgame/users/" + userCursorAlice + "/cosmic-signature-tokens"
	cases := map[string]fakeUserActivityReader{
		"address resolution fails": {
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("boom")
			},
		},
		"fetch fails": {
			ownedTokens: func(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
				return nil, false, errors.New("boom")
			},
		},
		"cardinality violation": {
			ownedTokens: func(_ context.Context, _ int64, _ *cgstore.UserTokenPageCursor, limit int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
				records := make([]cgstore.UserOwnedTokenRecord, limit+1)
				for i := range records {
					records[i] = ownedTokenAt(int64(i + 1))
				}
				return records, false, nil
			},
		},
		"out-of-scope row": {
			ownedTokens: func(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
				record := ownedTokenAt(1)
				record.OwnerAid = 99
				return []cgstore.UserOwnedTokenRecord{record}, false, nil
			},
		},
		"unordered rows": {
			ownedTokens: func(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
				return []cgstore.UserOwnedTokenRecord{ownedTokenAt(6), ownedTokenAt(1)}, false, nil
			},
		},
		"corrupt row": {
			ownedTokens: func(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
				record := ownedTokenAt(1)
				record.Seed = ""
				return []cgstore.UserOwnedTokenRecord{record}, false, nil
			},
		},
		"more without rows": {
			ownedTokens: func(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
				return []cgstore.UserOwnedTokenRecord{}, true, nil
			},
		},
	}
	for name, reader := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserActivityTestServer(t, reader), target)
			if response.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			if strings.Contains(response.Body.String(), "boom") {
				t.Fatal("internal error details leaked")
			}
		})
	}
}

func TestListUserOwnedTokensCursorScopedToPreviousPage(t *testing.T) {
	t.Parallel()

	// A repository row at or below the continuation cursor is unordered.
	cursor, err := encodeUserOwnedTokenCursor(userOwnedTokenCursor{
		Version: userOwnedTokenCursorVersion, Address: userCursorAlice, TokenID: 5,
	})
	if err != nil {
		t.Fatal(err)
	}
	server := newUserActivityTestServer(t, fakeUserActivityReader{
		ownedTokens: func(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error) {
			return []cgstore.UserOwnedTokenRecord{ownedTokenAt(5)}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-signature-tokens?cursor="+cursor)
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d", response.Code)
	}
}

func csTransferAt(eventLogID int64, direction cgstore.UserTransferDirection) cgstore.UserCosmicSignatureTransferRecord {
	record := validCsTransferRecord()
	record.Tx.EvtLogId = eventLogID
	record.Direction = direction
	switch direction {
	case cgstore.UserTransferOut:
		record.FromAid = 1
		record.ToAid = 2
	case cgstore.UserTransferIn:
		record.FromAid = 2
		record.ToAid = 1
	case cgstore.UserTransferSelf:
		record.FromAid = 1
		record.ToAid = 1
	}
	return record
}

func TestListUserCosmicSignatureTransfersPaginates(t *testing.T) {
	t.Parallel()

	server := newUserActivityTestServer(t, fakeUserActivityReader{
		csTransfers: func(_ context.Context, userAid int64, after *cgstore.UserEventPageCursor, limit int) ([]cgstore.UserCosmicSignatureTransferRecord, bool, error) {
			if userAid != 1 || after != nil || limit != 2 {
				t.Errorf("repository args = (%d, %+v, %d)", userAid, after, limit)
			}
			return []cgstore.UserCosmicSignatureTransferRecord{
				csTransferAt(300, cgstore.UserTransferIn),
				csTransferAt(290, cgstore.UserTransferSelf),
			}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-signature-transfers?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page CosmicGameUserCosmicSignatureTransferPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	if page.Data[0].Direction != In || page.Data[1].Direction != Self {
		t.Fatalf("directions = %v, %v", page.Data[0].Direction, page.Data[1].Direction)
	}
	cursor, err := decodeUserEventCursor(*page.Meta.NextCursor, userCursorAlice, userEventResourceCsTransfers)
	if err != nil || cursor.EventLogID != 290 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}
}

func TestListUserCosmicSignatureTransfersRejectsInconsistentDirection(t *testing.T) {
	t.Parallel()

	record := csTransferAt(300, cgstore.UserTransferIn)
	record.Direction = cgstore.UserTransferOut // contradicts from/to
	server := newUserActivityTestServer(t, fakeUserActivityReader{
		csTransfers: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserCosmicSignatureTransferRecord, bool, error) {
			return []cgstore.UserCosmicSignatureTransferRecord{record}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-signature-transfers")
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d", response.Code)
	}
}

func TestListUserCosmicTokenTransfersPaginates(t *testing.T) {
	t.Parallel()

	transfer := validCtTransferRecord()
	transfer.Tx.EvtLogId = 5049
	server := newUserActivityTestServer(t, fakeUserActivityReader{
		ctTransfers: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserCosmicTokenTransferRecord, bool, error) {
			return []cgstore.UserCosmicTokenTransferRecord{transfer}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-token-transfers?limit=1")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page CosmicGameUserCosmicTokenTransferPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 1 || page.Meta.NextCursor == nil ||
		page.Data[0].AmountWei != "10000000000000000000" {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeUserEventCursor(*page.Meta.NextCursor, userCursorAlice, userEventResourceCtTransfers)
	if err != nil || cursor.EventLogID != 5049 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}

	// A row that does not involve the wallet is out of scope.
	foreign := validCtTransferRecord()
	foreign.FromAid, foreign.ToAid = 7, 8
	server = newUserActivityTestServer(t, fakeUserActivityReader{
		ctTransfers: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserCosmicTokenTransferRecord, bool, error) {
			return []cgstore.UserCosmicTokenTransferRecord{foreign}, false, nil
		},
	})
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-token-transfers")
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("out-of-scope status = %d", response.Code)
	}
}

func TestListUserMarketingRewardsPaginates(t *testing.T) {
	t.Parallel()

	reward := cgstore.UserMarketingRewardRecord{
		Tx:          validDonationTransaction(),
		MarketerAid: 1,
		AmountWei:   "50000000000000000000",
	}
	server := newUserActivityTestServer(t, fakeUserActivityReader{
		mktRewards: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserMarketingRewardRecord, bool, error) {
			return []cgstore.UserMarketingRewardRecord{reward}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/marketing-rewards")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page CosmicGameUserMarketingRewardPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 1 || page.Meta.NextCursor != nil ||
		page.Data[0].AmountWei != "50000000000000000000" {
		t.Fatalf("page = %+v", page)
	}

	// A reward paid to another wallet is out of scope.
	foreign := reward
	foreign.MarketerAid = 9
	server = newUserActivityTestServer(t, fakeUserActivityReader{
		mktRewards: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserMarketingRewardRecord, bool, error) {
			return []cgstore.UserMarketingRewardRecord{foreign}, false, nil
		},
	})
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/marketing-rewards")
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("out-of-scope status = %d", response.Code)
	}
}

func TestGetUserCosmicTokenSummary(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		server := newUserActivityTestServer(t, fakeUserActivityReader{
			tokenSummary: func(_ context.Context, userAid int64) (cgstore.UserCosmicTokenSummaryRecord, error) {
				if userAid != 1 {
					t.Errorf("user aid = %d", userAid)
				}
				return validUserCosmicTokenSummaryRecord(), nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-token-summary")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var summary CosmicGameUserCosmicTokenSummary
		decodeResponse(t, response, &summary)
		if summary.BalanceWei != "290000000000000000000" ||
			summary.Earned.TotalWei != "635000000000000000000" ||
			!strings.EqualFold(summary.Address, userCursorAlice) {
			t.Fatalf("summary = %+v", summary)
		}
	})

	t.Run("unindexed wallet gets the zero shape", func(t *testing.T) {
		t.Parallel()
		server := newUserActivityTestServer(t, fakeUserActivityReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			},
			tokenSummary: func(context.Context, int64) (cgstore.UserCosmicTokenSummaryRecord, error) {
				t.Error("summary fetched for an unindexed wallet")
				return cgstore.UserCosmicTokenSummaryRecord{}, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-token-summary")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d", response.Code)
		}
		var summary CosmicGameUserCosmicTokenSummary
		decodeResponse(t, response, &summary)
		if summary.BalanceWei != "0" || summary.Earned.TotalWei != "0" || summary.NetWei != "0" {
			t.Fatalf("zero summary = %+v", summary)
		}
	})

	t.Run("invalid address", func(t *testing.T) {
		t.Parallel()
		server := newUserActivityTestServer(t, fakeUserActivityReader{})
		response := serve(t, server, "/api/v2/cosmicgame/users/not-an-address/cosmic-token-summary")
		if response.Code != http.StatusBadRequest {
			t.Fatalf("status = %d", response.Code)
		}
	})

	failures := map[string]fakeUserActivityReader{
		"address resolution fails": {
			addressID: func(context.Context, string) (int64, error) { return 0, errors.New("boom") },
		},
		"store fails": {
			tokenSummary: func(context.Context, int64) (cgstore.UserCosmicTokenSummaryRecord, error) {
				return cgstore.UserCosmicTokenSummaryRecord{}, errors.New("boom")
			},
		},
		"inconsistent record": {
			tokenSummary: func(context.Context, int64) (cgstore.UserCosmicTokenSummaryRecord, error) {
				record := validUserCosmicTokenSummaryRecord()
				record.NetWei = "1"
				return record, nil
			},
		},
	}
	for name, reader := range failures {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserActivityTestServer(t, reader),
				"/api/v2/cosmicgame/users/"+userCursorAlice+"/cosmic-token-summary")
			if response.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d", response.Code)
			}
			if strings.Contains(response.Body.String(), "boom") {
				t.Fatal("internal error details leaked")
			}
		})
	}
}

func TestGetUserPendingWinnings(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		server := newUserActivityTestServer(t, fakeUserActivityReader{
			pending: func(context.Context, int64) (cgstore.UserPendingWinningsRecord, error) {
				return cgstore.UserPendingWinningsRecord{
					RaffleEthWei:           "60000000000000000",
					ChronoWarriorEthWei:    "80000000000000000",
					DonatedNftCount:        1,
					StakingRewardWei:       "0",
					DonatedErc20TokenCount: 0,
				}, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/pending-winnings")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var winnings CosmicGameUserPendingWinnings
		decodeResponse(t, response, &winnings)
		if winnings.RaffleEthWei != "60000000000000000" ||
			winnings.ChronoWarriorEthWei != "80000000000000000" ||
			winnings.DonatedNftCount != 1 {
			t.Fatalf("winnings = %+v", winnings)
		}
	})

	t.Run("unindexed wallet gets the zero shape", func(t *testing.T) {
		t.Parallel()
		server := newUserActivityTestServer(t, fakeUserActivityReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/pending-winnings")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d", response.Code)
		}
		var winnings CosmicGameUserPendingWinnings
		decodeResponse(t, response, &winnings)
		if winnings.RaffleEthWei != "0" || winnings.StakingRewardWei != "0" ||
			winnings.DonatedNftCount != 0 {
			t.Fatalf("zero winnings = %+v", winnings)
		}
	})

	t.Run("invalid address", func(t *testing.T) {
		t.Parallel()
		server := newUserActivityTestServer(t, fakeUserActivityReader{})
		response := serve(t, server, "/api/v2/cosmicgame/users/nope/pending-winnings")
		if response.Code != http.StatusBadRequest {
			t.Fatalf("status = %d", response.Code)
		}
	})

	failures := map[string]fakeUserActivityReader{
		"address resolution fails": {
			addressID: func(context.Context, string) (int64, error) { return 0, errors.New("boom") },
		},
		"store fails": {
			pending: func(context.Context, int64) (cgstore.UserPendingWinningsRecord, error) {
				return cgstore.UserPendingWinningsRecord{}, errors.New("boom")
			},
		},
		"corrupt record": {
			pending: func(context.Context, int64) (cgstore.UserPendingWinningsRecord, error) {
				return cgstore.UserPendingWinningsRecord{
					RaffleEthWei:        "-1",
					ChronoWarriorEthWei: "0",
					StakingRewardWei:    "0",
				}, nil
			},
		},
	}
	for name, reader := range failures {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserActivityTestServer(t, reader),
				"/api/v2/cosmicgame/users/"+userCursorAlice+"/pending-winnings")
			if response.Code != http.StatusInternalServerError {
				t.Fatalf("status = %d", response.Code)
			}
			if strings.Contains(response.Body.String(), "boom") {
				t.Fatal("internal error details leaked")
			}
		})
	}
}
