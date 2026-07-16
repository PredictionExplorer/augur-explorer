package v2

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func newGlobalDirectoryTestServer(t *testing.T, directories globalDirectoryReader) *Server {
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
		fakeUserActivityReader{},
		directories,
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func globalTokenAt(tokenID int64) cgstore.GlobalTokenRecord {
	record := validGlobalTokenRecord()
	record.TokenID = tokenID
	return record
}

func TestListGlobalTokensPaginates(t *testing.T) {
	t.Parallel()

	var gotFilter cgstore.GlobalTokenFilter
	var gotAfter *cgstore.GlobalTokenPageCursor
	var gotLimit int
	server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
		tokens: func(_ context.Context, filter cgstore.GlobalTokenFilter, after *cgstore.GlobalTokenPageCursor, limit int) ([]cgstore.GlobalTokenRecord, bool, error) {
			gotFilter, gotAfter, gotLimit = filter, after, limit
			return []cgstore.GlobalTokenRecord{globalTokenAt(9), globalTokenAt(4)}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotFilter != (cgstore.GlobalTokenFilter{}) || gotAfter != nil || gotLimit != 2 {
		t.Fatalf("repository args = (%+v, %+v, %d)", gotFilter, gotAfter, gotLimit)
	}
	var page CosmicGameCosmicSignatureTokenPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	if page.Data[0].NftTokenId != 9 || page.Data[1].NftTokenId != 4 {
		t.Fatalf("token ids = %d, %d", page.Data[0].NftTokenId, page.Data[1].NftTokenId)
	}
	cursor, err := decodeGlobalTokenCursor(*page.Meta.NextCursor, globalTokenFilterScope{})
	if err != nil || cursor.TokenID != 4 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}

	server = newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
		tokens: func(_ context.Context, _ cgstore.GlobalTokenFilter, after *cgstore.GlobalTokenPageCursor, _ int) ([]cgstore.GlobalTokenRecord, bool, error) {
			gotAfter = after
			return []cgstore.GlobalTokenRecord{}, false, nil
		},
	})
	response = serve(t, server,
		"/api/v2/cosmicgame/cosmic-signature-tokens?cursor="+*page.Meta.NextCursor)
	if response.Code != http.StatusOK {
		t.Fatalf("continuation status = %d", response.Code)
	}
	if gotAfter == nil || gotAfter.TokenID != 4 {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
}

func TestListGlobalTokensFilters(t *testing.T) {
	t.Parallel()

	var gotFilter cgstore.GlobalTokenFilter
	server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
		tokens: func(_ context.Context, filter cgstore.GlobalTokenFilter, _ *cgstore.GlobalTokenPageCursor, _ int) ([]cgstore.GlobalTokenRecord, bool, error) {
			gotFilter = filter
			return []cgstore.GlobalTokenRecord{}, false, nil
		},
	})

	response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens?named=true")
	if response.Code != http.StatusOK || !gotFilter.NamedOnly || gotFilter.NameContains != "" {
		t.Fatalf("named filter = %+v, status=%d", gotFilter, response.Code)
	}

	response = serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens?name=Genesis")
	if response.Code != http.StatusOK || gotFilter.NamedOnly || gotFilter.NameContains != "Genesis" {
		t.Fatalf("name filter = %+v, status=%d", gotFilter, response.Code)
	}

	// named=false behaves like the unfiltered directory.
	response = serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens?named=false")
	if response.Code != http.StatusOK || gotFilter != (cgstore.GlobalTokenFilter{}) {
		t.Fatalf("named=false filter = %+v, status=%d", gotFilter, response.Code)
	}
}

func TestListGlobalTokensRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	namedCursor, err := encodeGlobalTokenCursor(globalTokenCursor{
		Version: globalTokenCursorVersion,
		Filter:  globalTokenFilterScope{Named: true},
		TokenID: 5,
	})
	if err != nil {
		t.Fatal(err)
	}
	server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
		tokens: func(context.Context, cgstore.GlobalTokenFilter, *cgstore.GlobalTokenPageCursor, int) ([]cgstore.GlobalTokenRecord, bool, error) {
			t.Error("repository reached with invalid input")
			return nil, false, nil
		},
	})

	cases := map[string]string{
		"contradictory filters": "?named=true&name=x",
		"empty name":            "?name=",
		"oversized name":        "?name=" + strings.Repeat("a", maxTokenNameSearchLength+1),
		"zero limit":            "?limit=0",
		"oversized limit":       "?limit=201",
		"malformed cursor":      "?cursor=%25%25",
		"cross-filter cursor":   "?cursor=" + namedCursor,
	}
	for name, query := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens"+query)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListGlobalTokensInternalFailures(t *testing.T) {
	t.Parallel()

	secret := errors.New("secret store failure")
	unordered := []cgstore.GlobalTokenRecord{globalTokenAt(4), globalTokenAt(9)}
	corrupt := globalTokenAt(3)
	corrupt.Seed = ""

	cases := map[string]fakeGlobalDirectoryReader{
		"store failure": {
			tokens: func(context.Context, cgstore.GlobalTokenFilter, *cgstore.GlobalTokenPageCursor, int) ([]cgstore.GlobalTokenRecord, bool, error) {
				return nil, false, secret
			},
		},
		"unordered page": {
			tokens: func(context.Context, cgstore.GlobalTokenFilter, *cgstore.GlobalTokenPageCursor, int) ([]cgstore.GlobalTokenRecord, bool, error) {
				return unordered, false, nil
			},
		},
		"over-cardinality page": {
			tokens: func(_ context.Context, _ cgstore.GlobalTokenFilter, _ *cgstore.GlobalTokenPageCursor, limit int) ([]cgstore.GlobalTokenRecord, bool, error) {
				records := make([]cgstore.GlobalTokenRecord, limit+1)
				for i := range records {
					records[i] = globalTokenAt(int64(1000 - i))
				}
				return records, true, nil
			},
		},
		"corrupt row": {
			tokens: func(context.Context, cgstore.GlobalTokenFilter, *cgstore.GlobalTokenPageCursor, int) ([]cgstore.GlobalTokenRecord, bool, error) {
				return []cgstore.GlobalTokenRecord{corrupt}, false, nil
			},
		},
		"hasMore without rows": {
			tokens: func(context.Context, cgstore.GlobalTokenFilter, *cgstore.GlobalTokenPageCursor, int) ([]cgstore.GlobalTokenRecord, bool, error) {
				return []cgstore.GlobalTokenRecord{}, true, nil
			},
		},
	}
	for name, directories := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			server := newGlobalDirectoryTestServer(t, directories)
			response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens")
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "secret") {
				t.Fatal("internal detail leaked")
			}
		})
	}
}

func TestGetGlobalTokenDetail(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenDetail: func(_ context.Context, tokenID int64) (cgstore.GlobalTokenDetailRecord, error) {
				record := validGlobalTokenDetailRecord()
				record.TokenID = tokenID
				return record, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var detail CosmicGameCosmicSignatureTokenDetail
		decodeResponse(t, response, &detail)
		if detail.NftTokenId != 5 || !detail.Staked || detail.CurrentStake == nil {
			t.Fatalf("detail = %+v", detail)
		}
	})

	t.Run("unknown token is 404", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/404")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("negative token id is 400", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/-1")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("wrong token from the repository is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenDetail: func(context.Context, int64) (cgstore.GlobalTokenDetailRecord, error) {
				return validGlobalTokenDetailRecord(), nil // token 5
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/6")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("corrupt detail record is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenDetail: func(_ context.Context, tokenID int64) (cgstore.GlobalTokenDetailRecord, error) {
				record := validGlobalTokenDetailRecord()
				record.TokenID = tokenID
				record.Seed = ""
				return record, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("store failure is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenDetail: func(context.Context, int64) (cgstore.GlobalTokenDetailRecord, error) {
				return cgstore.GlobalTokenDetailRecord{}, errors.New("secret failure")
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "secret") {
			t.Fatal("internal detail leaked")
		}
	})
}

func nameChangeAt(eventLogID int64) cgstore.TokenNameChangeRecord {
	record := validTokenNameChangeRecord()
	record.Tx.EvtLogId = eventLogID
	return record
}

// TestTokenEventPageFailureMatrix drives every opaque-500 arm of the shared
// token-event page flow through the name-history resource.
func TestTokenEventPageFailureMatrix(t *testing.T) {
	t.Parallel()

	corruptRename := nameChangeAt(900)
	corruptRename.ChangedBy = "nope"

	internalCases := map[string]fakeGlobalDirectoryReader{
		"store failure": {
			nameHistory: func(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenNameChangeRecord, bool, error) {
				return nil, false, errors.New("secret store failure")
			},
		},
		"over-cardinality page": {
			nameHistory: func(_ context.Context, _ int64, _ *cgstore.TokenEventPageCursor, limit int) ([]cgstore.TokenNameChangeRecord, bool, error) {
				records := make([]cgstore.TokenNameChangeRecord, limit+1)
				for i := range records {
					records[i] = nameChangeAt(int64(10_000 - i))
				}
				return records, true, nil
			},
		},
		"corrupt row": {
			nameHistory: func(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenNameChangeRecord, bool, error) {
				return []cgstore.TokenNameChangeRecord{corruptRename}, false, nil
			},
		},
		"hasMore without rows": {
			nameHistory: func(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenNameChangeRecord, bool, error) {
				return []cgstore.TokenNameChangeRecord{}, true, nil
			},
		},
	}
	for name, directories := range internalCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			server := newGlobalDirectoryTestServer(t, directories)
			response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5/name-history")
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "secret") {
				t.Fatal("internal detail leaked")
			}
		})
	}

	t.Run("invalid limit", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-signature-tokens/5/name-history?limit=0")
		assertProblem(t, response, http.StatusBadRequest)
	})
}

func TestListTokenNameHistory(t *testing.T) {
	t.Parallel()

	t.Run("paginates newest first", func(t *testing.T) {
		t.Parallel()
		var gotAfter *cgstore.TokenEventPageCursor
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			nameHistory: func(_ context.Context, tokenID int64, after *cgstore.TokenEventPageCursor, limit int) ([]cgstore.TokenNameChangeRecord, bool, error) {
				if tokenID != 5 || limit != 2 {
					t.Errorf("args = (%d, %d)", tokenID, limit)
				}
				gotAfter = after
				return []cgstore.TokenNameChangeRecord{nameChangeAt(900), nameChangeAt(800)}, true, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-signature-tokens/5/name-history?limit=2")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		if gotAfter != nil {
			t.Fatalf("first page cursor = %+v", gotAfter)
		}
		var page CosmicGameTokenNameChangePage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Meta.NextCursor == nil {
			t.Fatalf("page = %+v", page)
		}
		cursor, err := decodeTokenEventCursor(*page.Meta.NextCursor, 5, tokenEventResourceNameHistory)
		if err != nil || cursor.EventLogID != 800 {
			t.Fatalf("next cursor = %+v, err=%v", cursor, err)
		}
	})

	t.Run("unknown token is 404 before paging", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenExists: func(context.Context, int64) (bool, error) { return false, nil },
			nameHistory: func(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenNameChangeRecord, bool, error) {
				t.Error("page fetched for an unknown token")
				return nil, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5/name-history")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("existence failure is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenExists: func(context.Context, int64) (bool, error) {
				return false, errors.New("secret failure")
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5/name-history")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "secret") {
			t.Fatal("internal detail leaked")
		}
	})

	t.Run("cross-resource cursor is rejected", func(t *testing.T) {
		t.Parallel()
		transfersCursor, err := encodeTokenEventCursor(tokenEventCursor{
			Version:    tokenEventCursorVersion,
			Resource:   tokenEventResourceTransfers,
			TokenID:    5,
			EventLogID: 900,
		})
		if err != nil {
			t.Fatal(err)
		}
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-signature-tokens/5/name-history?cursor="+transfersCursor)
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("cross-token cursor is rejected", func(t *testing.T) {
		t.Parallel()
		otherToken, err := encodeTokenEventCursor(tokenEventCursor{
			Version:    tokenEventCursorVersion,
			Resource:   tokenEventResourceNameHistory,
			TokenID:    6,
			EventLogID: 900,
		})
		if err != nil {
			t.Fatal(err)
		}
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-signature-tokens/5/name-history?cursor="+otherToken)
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("out-of-scope row is opaque", func(t *testing.T) {
		t.Parallel()
		foreign := nameChangeAt(900)
		foreign.TokenID = 6
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			nameHistory: func(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenNameChangeRecord, bool, error) {
				return []cgstore.TokenNameChangeRecord{foreign}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5/name-history")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func tokenTransferAt(eventLogID int64) cgstore.TokenTransferRecord {
	record := validTokenTransferRecord()
	record.Tx.EvtLogId = eventLogID
	return record
}

func TestListTokenTransfers(t *testing.T) {
	t.Parallel()

	t.Run("paginates newest first", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenTransfers: func(_ context.Context, tokenID int64, _ *cgstore.TokenEventPageCursor, _ int) ([]cgstore.TokenTransferRecord, bool, error) {
				if tokenID != 5 {
					t.Errorf("token id = %d", tokenID)
				}
				return []cgstore.TokenTransferRecord{tokenTransferAt(700), tokenTransferAt(600)}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5/transfers")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var page CosmicGameCosmicSignatureTokenTransferPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Meta.NextCursor == nil {
			t.Fatalf("page = %+v", page)
		}
		cursor, err := decodeTokenEventCursor(*page.Meta.NextCursor, 5, tokenEventResourceTransfers)
		if err != nil || cursor.EventLogID != 600 {
			t.Fatalf("next cursor = %+v, err=%v", cursor, err)
		}
	})

	t.Run("unknown token is 404", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenExists: func(context.Context, int64) (bool, error) { return false, nil },
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5/transfers")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("unordered page is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenTransfers: func(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenTransferRecord, bool, error) {
				return []cgstore.TokenTransferRecord{tokenTransferAt(600), tokenTransferAt(700)}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/5/transfers")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func TestListCosmicSignatureHolders(t *testing.T) {
	t.Parallel()

	t.Run("paginates by holdings", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			csHolders: func(_ context.Context, after *cgstore.ParticipantPageCursor, limit int) ([]cgstore.CosmicSignatureHolderRecord, bool, error) {
				if after != nil || limit != 2 {
					t.Errorf("args = (%+v, %d)", after, limit)
				}
				return []cgstore.CosmicSignatureHolderRecord{
					{OwnerAid: 1, Address: userCursorAlice, TokenCount: 4},
					{OwnerAid: 2, Address: userCursorBob, TokenCount: 2},
				}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/holders?limit=2")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var page CosmicGameCosmicSignatureHolderPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Meta.NextCursor == nil ||
			page.Data[0].TokenCount != 4 || page.Data[1].TokenCount != 2 {
			t.Fatalf("page = %+v", page)
		}
		cursor, err := decodeParticipantCursor(*page.Meta.NextCursor, cgstore.ParticipantCsTokenHolders)
		if err != nil || cursor.SortValue != "2" || cursor.AddressID != 2 {
			t.Fatalf("next cursor = %+v, err=%v", cursor, err)
		}
	})

	t.Run("cross-directory cursor is rejected", func(t *testing.T) {
		t.Parallel()
		bidders, err := encodeParticipantCursor(participantCursor{
			Version: participantCursorVersion, Kind: cgstore.ParticipantBidders,
			SortValue: "1", AddressID: 1,
		})
		if err != nil {
			t.Fatal(err)
		}
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-signature-tokens/holders?cursor="+bidders)
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("unordered page is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			csHolders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CosmicSignatureHolderRecord, bool, error) {
				return []cgstore.CosmicSignatureHolderRecord{
					{OwnerAid: 2, Address: userCursorBob, TokenCount: 2},
					{OwnerAid: 1, Address: userCursorAlice, TokenCount: 4},
				}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/holders")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("store failure is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			csHolders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CosmicSignatureHolderRecord, bool, error) {
				return nil, false, errors.New("secret store failure")
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-signature-tokens/holders")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "secret") {
			t.Fatal("internal detail leaked")
		}
	})
}

func TestListCosmicTokenHolders(t *testing.T) {
	t.Parallel()

	t.Run("paginates by exact balance", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			ctHolders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CosmicTokenHolderRecord, bool, error) {
				return []cgstore.CosmicTokenHolderRecord{
					{OwnerAid: 1, Address: userCursorAlice, BalanceWei: "999999999999999999999999"},
					{OwnerAid: 2, Address: userCursorBob, BalanceWei: "1"},
				}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/holders")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var page CosmicGameCosmicTokenHolderPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Data[0].BalanceWei != "999999999999999999999999" {
			t.Fatalf("page = %+v", page)
		}
		cursor, err := decodeParticipantCursor(*page.Meta.NextCursor, cgstore.ParticipantCosmicTokenHolders)
		if err != nil || cursor.SortValue != "1" {
			t.Fatalf("next cursor = %+v, err=%v", cursor, err)
		}
	})

	t.Run("zero-balance row is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			ctHolders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CosmicTokenHolderRecord, bool, error) {
				return []cgstore.CosmicTokenHolderRecord{
					{OwnerAid: 1, Address: userCursorAlice, BalanceWei: "0"},
				}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/holders")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("invalid limit is rejected", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/holders?limit=201")
		assertProblem(t, response, http.StatusBadRequest)
	})
}

func TestGetCosmicTokenStatistics(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenStatistics: func(context.Context) (cgstore.CosmicTokenStatisticsRecord, error) {
				return validCosmicTokenStatisticsRecord(), nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/statistics")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var statistics CosmicGameCosmicTokenStatistics
		decodeResponse(t, response, &statistics)
		if statistics.TotalSupplyWei != "1000" || len(statistics.TopHolders) != 3 ||
			statistics.TopHolders[0].ShareOfSupply != "60" {
			t.Fatalf("statistics = %+v", statistics)
		}
	})

	t.Run("store failure is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenStatistics: func(context.Context) (cgstore.CosmicTokenStatisticsRecord, error) {
				return cgstore.CosmicTokenStatisticsRecord{}, errors.New("secret failure")
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/statistics")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "secret") {
			t.Fatal("internal detail leaked")
		}
	})

	t.Run("invariant violation is opaque", func(t *testing.T) {
		t.Parallel()
		record := validCosmicTokenStatisticsRecord()
		record.NetWei = "1"
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			tokenStatistics: func(context.Context) (cgstore.CosmicTokenStatisticsRecord, error) {
				return record, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/statistics")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func supplyChangeRow(eventLogID, previousTotal, net int64) cgstore.SupplyChangeRecord {
	record := validSupplyChangeRecord()
	record.Tx.EvtLogId = eventLogID
	record.MintedWei = strconv.FormatInt(net, 10)
	record.BurnedWei = "0"
	record.NetWei = strconv.FormatInt(net, 10)
	record.TotalSupplyWei = strconv.FormatInt(previousTotal+net, 10)
	return record
}

func TestListSupplyByBid(t *testing.T) {
	t.Parallel()

	t.Run("paginates oldest first with running totals", func(t *testing.T) {
		t.Parallel()
		var gotAfter *cgstore.SupplyChangePageCursor
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			supplyByBid: func(_ context.Context, after *cgstore.SupplyChangePageCursor, _ int) ([]cgstore.SupplyChangeRecord, bool, error) {
				gotAfter = after
				return []cgstore.SupplyChangeRecord{
					supplyChangeRow(100, 0, 40),
					supplyChangeRow(200, 40, 60),
				}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/supply-history/by-bid")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		if gotAfter != nil {
			t.Fatalf("first page cursor = %+v", gotAfter)
		}
		var page CosmicGameCosmicTokenSupplyByBidPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Meta.NextCursor == nil ||
			page.Data[1].TotalSupplyWei != "100" {
			t.Fatalf("page = %+v", page)
		}
		cursor, err := decodeSupplyChangeCursor(*page.Meta.NextCursor)
		if err != nil || cursor.EventLogID != 200 {
			t.Fatalf("next cursor = %+v, err=%v", cursor, err)
		}
	})

	t.Run("diverging running total is opaque", func(t *testing.T) {
		t.Parallel()
		second := supplyChangeRow(200, 40, 60)
		second.TotalSupplyWei = "999"
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			supplyByBid: func(context.Context, *cgstore.SupplyChangePageCursor, int) ([]cgstore.SupplyChangeRecord, bool, error) {
				return []cgstore.SupplyChangeRecord{supplyChangeRow(100, 0, 40), second}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/supply-history/by-bid")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("unordered page is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			supplyByBid: func(context.Context, *cgstore.SupplyChangePageCursor, int) ([]cgstore.SupplyChangeRecord, bool, error) {
				return []cgstore.SupplyChangeRecord{
					supplyChangeRow(200, 0, 40),
					supplyChangeRow(100, 40, 60),
				}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/supply-history/by-bid")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("malformed cursor is rejected", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-token/supply-history/by-bid?cursor=nope")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("invalid limit is rejected", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-token/supply-history/by-bid?limit=0")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("over-cardinality page is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			supplyByBid: func(_ context.Context, _ *cgstore.SupplyChangePageCursor, limit int) ([]cgstore.SupplyChangeRecord, bool, error) {
				records := make([]cgstore.SupplyChangeRecord, limit+1)
				total := int64(0)
				for i := range records {
					records[i] = supplyChangeRow(int64(100+i), total, 10)
					total += 10
				}
				return records, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/supply-history/by-bid")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("corrupt row is opaque", func(t *testing.T) {
		t.Parallel()
		corrupt := supplyChangeRow(100, 0, 40)
		corrupt.BidderAddr = "nope"
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			supplyByBid: func(context.Context, *cgstore.SupplyChangePageCursor, int) ([]cgstore.SupplyChangeRecord, bool, error) {
				return []cgstore.SupplyChangeRecord{corrupt}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/supply-history/by-bid")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("hasMore without rows is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			supplyByBid: func(context.Context, *cgstore.SupplyChangePageCursor, int) ([]cgstore.SupplyChangeRecord, bool, error) {
				return []cgstore.SupplyChangeRecord{}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/cosmic-token/supply-history/by-bid")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func dailySupplyRow(date string) cgstore.DailySupplyRecord {
	return cgstore.DailySupplyRecord{
		Date:           date,
		BidCount:       2,
		MintedWei:      "10",
		BurnedWei:      "0",
		NetWei:         "10",
		TotalSupplyWei: "10",
	}
}

func TestListSupplyDaily(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		var gotFrom, gotTo time.Time
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			supplyDaily: func(_ context.Context, from, to time.Time) ([]cgstore.DailySupplyRecord, error) {
				gotFrom, gotTo = from, to
				return []cgstore.DailySupplyRecord{dailySupplyRow("2026-01-02")}, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-token/supply-history/daily?from=2026-01-01&to=2026-01-10")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		if gotFrom.Format("2006-01-02") != "2026-01-01" || gotTo.Format("2006-01-02") != "2026-01-10" {
			t.Fatalf("window = (%v, %v)", gotFrom, gotTo)
		}
		var list CosmicGameCosmicTokenSupplyDailyList
		decodeResponse(t, response, &list)
		if len(list.Data) != 1 || list.Data[0].BidCount != 2 {
			t.Fatalf("list = %+v", list)
		}
	})

	windowCases := map[string]string{
		"missing from":  "?to=2026-01-10",
		"missing to":    "?from=2026-01-01",
		"empty window":  "?from=2026-01-10&to=2026-01-10",
		"reversed":      "?from=2026-01-10&to=2026-01-01",
		"oversized":     "?from=2020-01-01&to=2026-01-01",
		"malformed day": "?from=notaday&to=2026-01-10",
	}
	for name, query := range windowCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
				supplyDaily: func(context.Context, time.Time, time.Time) ([]cgstore.DailySupplyRecord, error) {
					t.Error("repository reached with an invalid window")
					return nil, nil
				},
			})
			response := serve(t, server,
				"/api/v2/cosmicgame/cosmic-token/supply-history/daily"+query)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}

	tooManyRows := make([]cgstore.DailySupplyRecord, cgstore.MaxSupplyDailyWindowDays+1)
	for i := range tooManyRows {
		day := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, i)
		tooManyRows[i] = dailySupplyRow(day.Format("2006-01-02"))
	}
	invalidRows := map[string][]cgstore.DailySupplyRecord{
		"unordered days": {dailySupplyRow("2026-01-03"), dailySupplyRow("2026-01-02")},
		"day out of window": {
			dailySupplyRow("2026-01-10"),
		},
		"corrupt row": {{
			Date: "2026-01-02", BidCount: 0,
			MintedWei: "0", BurnedWei: "0", NetWei: "0", TotalSupplyWei: "0",
		}},
		"more rows than the window allows": tooManyRows,
	}
	for name, rows := range invalidRows {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
				supplyDaily: func(context.Context, time.Time, time.Time) ([]cgstore.DailySupplyRecord, error) {
					return rows, nil
				},
			})
			response := serve(t, server,
				"/api/v2/cosmicgame/cosmic-token/supply-history/daily?from=2026-01-01&to=2026-01-10")
			assertProblem(t, response, http.StatusInternalServerError)
		})
	}

	t.Run("store failure is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			supplyDaily: func(context.Context, time.Time, time.Time) ([]cgstore.DailySupplyRecord, error) {
				return nil, errors.New("secret failure")
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/cosmic-token/supply-history/daily?from=2026-01-01&to=2026-01-10")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "secret") {
			t.Fatal("internal detail leaked")
		}
	})
}

func TestGlobalTokenNegativePathIdentifiersRejected(t *testing.T) {
	t.Parallel()

	// The OpenAPI minimum cannot be enforced by the stdlib binder, so the
	// handlers reject negative path identifiers themselves.
	for _, target := range []string{
		"/api/v2/cosmicgame/cosmic-signature-tokens/-1/name-history",
		"/api/v2/cosmicgame/cosmic-signature-tokens/-1/transfers",
	} {
		t.Run(target, func(t *testing.T) {
			t.Parallel()
			server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
				tokenExists: func(context.Context, int64) (bool, error) {
					t.Error("existence checked for a negative token id")
					return false, nil
				},
			})
			response := serve(t, server, target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListGlobalMarketingRewards(t *testing.T) {
	t.Parallel()

	reward := func(eventLogID int64) cgstore.MarketingRewardRecord {
		record := cgstore.MarketingRewardRecord{
			Tx:           validDonationTransaction(),
			MarketerAid:  1,
			MarketerAddr: userCursorAlice,
			AmountWei:    "5",
		}
		record.Tx.EvtLogId = eventLogID
		return record
	}

	t.Run("paginates newest first", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			mktRewards: func(_ context.Context, after *cgstore.UserEventPageCursor, limit int) ([]cgstore.MarketingRewardRecord, bool, error) {
				if after != nil || limit != 2 {
					t.Errorf("args = (%+v, %d)", after, limit)
				}
				return []cgstore.MarketingRewardRecord{reward(900), reward(800)}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/marketing-rewards?limit=2")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var page CosmicGameMarketingRewardPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Meta.NextCursor == nil ||
			!strings.EqualFold(page.Data[0].MarketerAddress, userCursorAlice) {
			t.Fatalf("page = %+v", page)
		}
		cursor, err := decodeGlobalMarketingCursor(*page.Meta.NextCursor)
		if err != nil || cursor.EventLogID != 800 {
			t.Fatalf("next cursor = %+v, err=%v", cursor, err)
		}
	})

	t.Run("unordered page is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			mktRewards: func(context.Context, *cgstore.UserEventPageCursor, int) ([]cgstore.MarketingRewardRecord, bool, error) {
				return []cgstore.MarketingRewardRecord{reward(800), reward(900)}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/marketing-rewards")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("cross-resource cursor is rejected", func(t *testing.T) {
		t.Parallel()
		supply, err := encodeSupplyChangeCursor(supplyChangeCursor{
			Version:    supplyChangeCursorVersion,
			EventLogID: 5,
		})
		if err != nil {
			t.Fatal(err)
		}
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server, "/api/v2/cosmicgame/marketing-rewards?cursor="+supply)
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("invalid limit is rejected", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{})
		response := serve(t, server, "/api/v2/cosmicgame/marketing-rewards?limit=201")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("over-cardinality page is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			mktRewards: func(_ context.Context, _ *cgstore.UserEventPageCursor, limit int) ([]cgstore.MarketingRewardRecord, bool, error) {
				records := make([]cgstore.MarketingRewardRecord, limit+1)
				for i := range records {
					records[i] = reward(int64(10_000 - i))
				}
				return records, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/marketing-rewards")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("corrupt row is opaque", func(t *testing.T) {
		t.Parallel()
		corrupt := reward(900)
		corrupt.MarketerAddr = "nope"
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			mktRewards: func(context.Context, *cgstore.UserEventPageCursor, int) ([]cgstore.MarketingRewardRecord, bool, error) {
				return []cgstore.MarketingRewardRecord{corrupt}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/marketing-rewards")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("hasMore without rows is opaque", func(t *testing.T) {
		t.Parallel()
		server := newGlobalDirectoryTestServer(t, fakeGlobalDirectoryReader{
			mktRewards: func(context.Context, *cgstore.UserEventPageCursor, int) ([]cgstore.MarketingRewardRecord, bool, error) {
				return []cgstore.MarketingRewardRecord{}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/marketing-rewards")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}
