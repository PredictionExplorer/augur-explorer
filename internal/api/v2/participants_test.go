package v2

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestParticipantDirectoriesReturnTypedPages(t *testing.T) {
	t.Parallel()
	server := newParticipantTestServer(t, fakeParticipantReader{
		bidders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error) {
			return []cgstore.BidderParticipantRecord{{
				BidderAid: 21, Address: participantTestAddress, BidCount: 3, MaxBidWei: "-1",
			}}, false, nil
		},
		winners: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.WinnerParticipantRecord, bool, error) {
			return []cgstore.WinnerParticipantRecord{validWinnerParticipantRecord()}, false, nil
		},
		donors: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DonorParticipantRecord, bool, error) {
			return []cgstore.DonorParticipantRecord{{
				DonorAid: 21, Address: participantTestAddress, DonationCount: 1, TotalDonatedWei: "10",
			}}, false, nil
		},
		cstStakers: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CSTStakerParticipantRecord, bool, error) {
			return []cgstore.CSTStakerParticipantRecord{{
				StakerAid: 21, Address: participantTestAddress, StakedTokenCount: 1,
				StakeActionCount: 1, TotalRewardWei: "10", UnclaimedRewardWei: "5",
			}}, false, nil
		},
		randomWalkStakers: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.RandomWalkStakerParticipantRecord, bool, error) {
			return []cgstore.RandomWalkStakerParticipantRecord{{
				StakerAid: 21, Address: participantTestAddress,
				StakedTokenCount: 1, StakeActionCount: 1,
			}}, false, nil
		},
		dualStakers: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DualStakerParticipantRecord, bool, error) {
			return []cgstore.DualStakerParticipantRecord{validDualStakerParticipantRecord()}, false, nil
		},
	})
	paths := []string{
		"/api/v2/cosmicgame/statistics/participants/bidders",
		"/api/v2/cosmicgame/statistics/participants/winners",
		"/api/v2/cosmicgame/statistics/participants/donors",
		"/api/v2/cosmicgame/statistics/participants/stakers/cst",
		"/api/v2/cosmicgame/statistics/participants/stakers/random-walk",
		"/api/v2/cosmicgame/statistics/participants/stakers/both",
	}
	for _, path := range paths {
		response := serve(t, server, path)
		if response.Code != http.StatusOK {
			t.Errorf("%s = %d %s", path, response.Code, response.Body.String())
		}
		if !bytes.Contains(response.Body.Bytes(), []byte(`"data":[{`)) ||
			!bytes.Contains(response.Body.Bytes(), []byte(`"limit":50`)) {
			t.Errorf("%s returned an invalid page: %s", path, response.Body.String())
		}
	}
}

func TestBidderParticipantsPaginatesAcrossTies(t *testing.T) {
	t.Parallel()
	var gotAfter *cgstore.ParticipantPageCursor
	var gotLimit int
	server := newParticipantTestServer(t, fakeParticipantReader{
		bidders: func(
			_ context.Context,
			after *cgstore.ParticipantPageCursor,
			limit int,
		) ([]cgstore.BidderParticipantRecord, bool, error) {
			gotAfter, gotLimit = after, limit
			return []cgstore.BidderParticipantRecord{
				{BidderAid: 21, Address: participantTestAddress, BidCount: 5, MaxBidWei: "10"},
				{BidderAid: 22, Address: "0x2200000000000000000000000000000000000022", BidCount: 5, MaxBidWei: "9"},
			}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/statistics/participants/bidders?limit=2")
	if response.Code != http.StatusOK || gotAfter != nil || gotLimit != 2 {
		t.Fatalf("response=%d args=%+v,%d body=%s", response.Code, gotAfter, gotLimit, response.Body.String())
	}
	var page BidderParticipantPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeParticipantCursor(*page.Meta.NextCursor, cgstore.ParticipantBidders)
	if err != nil || cursor.SortValue != "5" || cursor.AddressID != 22 {
		t.Fatalf("next cursor = %+v, %v", cursor, err)
	}
}

func TestParticipantDirectoryDecodesScopedCursor(t *testing.T) {
	t.Parallel()
	encoded, err := encodeParticipantCursor(participantCursor{
		Version: 1, Kind: cgstore.ParticipantDonors, SortValue: "100", AddressID: 25,
	})
	if err != nil {
		t.Fatal(err)
	}
	var got *cgstore.ParticipantPageCursor
	server := newParticipantTestServer(t, fakeParticipantReader{
		donors: func(
			_ context.Context,
			after *cgstore.ParticipantPageCursor,
			limit int,
		) ([]cgstore.DonorParticipantRecord, bool, error) {
			got = after
			if limit != defaultPageLimit {
				t.Errorf("limit = %d", limit)
			}
			return []cgstore.DonorParticipantRecord{}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/statistics/participants/donors?cursor="+encoded)
	if response.Code != http.StatusOK || got == nil ||
		got.Kind != cgstore.ParticipantDonors ||
		got.SortValue != "100" || got.AddressID != 25 {
		t.Fatalf("response=%d cursor=%+v body=%s", response.Code, got, response.Body.String())
	}
	if !bytes.Contains(response.Body.Bytes(), []byte(`"data":[]`)) {
		t.Fatalf("empty data is not []: %s", response.Body.String())
	}
}

func TestParticipantDirectoriesRejectInvalidInput(t *testing.T) {
	t.Parallel()
	crossKind, err := encodeParticipantCursor(participantCursor{
		Version: 1, Kind: cgstore.ParticipantWinners, SortValue: "1", AddressID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	base := "/api/v2/cosmicgame/statistics/participants/bidders"
	tests := map[string]string{
		"zero limit":       base + "?limit=0",
		"excessive limit":  base + "?limit=201",
		"bind limit":       base + "?limit=bad",
		"duplicate limit":  base + "?limit=1&limit=2",
		"malformed cursor": base + "?cursor=bad",
		"oversized cursor": base + "?cursor=" + strings.Repeat("a", maxCursorLength+1),
		"cross directory":  base + "?cursor=" + crossKind,
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assertProblem(t, serve(t, newParticipantTestServer(t, fakeParticipantReader{}), path),
				http.StatusBadRequest)
		})
	}
}

func TestParticipantDirectoriesHideRepositoryErrors(t *testing.T) {
	t.Parallel()
	secret := errors.New("password=participant-secret")
	server := newParticipantTestServer(t, fakeParticipantReader{
		bidders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error) {
			return nil, false, secret
		},
		winners: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.WinnerParticipantRecord, bool, error) {
			return nil, false, secret
		},
		donors: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DonorParticipantRecord, bool, error) {
			return nil, false, secret
		},
		cstStakers: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CSTStakerParticipantRecord, bool, error) {
			return nil, false, secret
		},
		randomWalkStakers: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.RandomWalkStakerParticipantRecord, bool, error) {
			return nil, false, secret
		},
		dualStakers: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DualStakerParticipantRecord, bool, error) {
			return nil, false, secret
		},
	})
	for _, path := range []string{
		"/api/v2/cosmicgame/statistics/participants/bidders",
		"/api/v2/cosmicgame/statistics/participants/winners",
		"/api/v2/cosmicgame/statistics/participants/donors",
		"/api/v2/cosmicgame/statistics/participants/stakers/cst",
		"/api/v2/cosmicgame/statistics/participants/stakers/random-walk",
		"/api/v2/cosmicgame/statistics/participants/stakers/both",
	} {
		response := serve(t, server, path)
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "participant-secret") {
			t.Fatalf("%s leaked an internal error: %s", path, response.Body.String())
		}
	}
}

func TestParticipantPageRejectsRepositoryInvariantViolations(t *testing.T) {
	t.Parallel()
	valid := cgstore.BidderParticipantRecord{
		BidderAid: 21, Address: participantTestAddress, BidCount: 5, MaxBidWei: "1",
	}
	tests := map[string]fakeParticipantReader{
		"has more without row": {
			bidders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error) {
				return []cgstore.BidderParticipantRecord{}, true, nil
			},
		},
		"out of order": {
			bidders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error) {
				second := valid
				second.BidderAid, second.BidCount = 22, 6
				return []cgstore.BidderParticipantRecord{valid, second}, false, nil
			},
		},
		"invalid mapped record": {
			bidders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error) {
				record := valid
				record.Address = "not-an-address"
				return []cgstore.BidderParticipantRecord{record}, false, nil
			},
		},
	}
	for name, participants := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newParticipantTestServer(t, participants),
				"/api/v2/cosmicgame/statistics/participants/bidders")
			assertProblem(t, response, http.StatusInternalServerError)
		})
	}
	t.Run("more rows than requested", func(t *testing.T) {
		t.Parallel()
		second := valid
		second.BidderAid = 22
		server := newParticipantTestServer(t, fakeParticipantReader{
			bidders: func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error) {
				return []cgstore.BidderParticipantRecord{valid, second}, false, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/statistics/participants/bidders?limit=1")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func newParticipantTestServer(t *testing.T, participants participantReader) *Server {
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
		participants,
		fakeUserReader{},
		fakeUserHistoryReader{},
		fakeUserStakingReader{},
		fakeUserActivityReader{},
		fakeGlobalDirectoryReader{},
		fakeGlobalStakingReader{},
		fakeRandomWalkReader{},
		fakeRankingRepository{},
		fakeContractState{},
		slog.New(slog.DiscardHandler))
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}
