package v2

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestListRoundEthDonationsPaginates(t *testing.T) {
	t.Parallel()

	first := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
	first.RoundNum, first.Tx.EvtLogId = 9, 100
	second := validRoundEthDonationRecord(cgstore.RoundEthDonationWithInfo)
	second.RoundNum, second.Tx.EvtLogId = 9, 90
	recordID := int64(7)
	data := `{"title":"hello"}`
	second.ContractRecordID, second.Data = &recordID, &data

	var (
		gotRound int64
		gotAfter *cgstore.DonationPageCursor
		gotLimit int
	)
	server := newDonationTestServer(t, fakeRoundDonationReader{
		eth: func(_ context.Context, round int64, after *cgstore.DonationPageCursor, limit int) ([]cgstore.RoundEthDonationRecord, bool, error) {
			gotRound, gotAfter, gotLimit = round, after, limit
			return []cgstore.RoundEthDonationRecord{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/9/eth-donations?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotRound != 9 || gotAfter != nil || gotLimit != 2 {
		t.Fatalf("repository args = (%d,%+v,%d)", gotRound, gotAfter, gotLimit)
	}
	var page RoundEthDonationPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Data[0].Kind != Plain || page.Data[1].Kind != WithInfo ||
		page.Meta.Limit != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeRoundDonationCursor(*page.Meta.NextCursor, 9, roundDonationResourceETH)
	if err != nil {
		t.Fatalf("decode next cursor: %v", err)
	}
	if cursor.EventLogID != 90 {
		t.Fatalf("next cursor = %+v", cursor)
	}
}

func TestListRoundDonationsMapsEveryResource(t *testing.T) {
	t.Parallel()

	erc20 := validRoundERC20DonationRecord()
	erc20.RoundNum = 3
	nft := validRoundNFTDonationRecord()
	nft.RoundNum = 3
	server := newDonationTestServer(t, fakeRoundDonationReader{
		erc20: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error) {
			return []cgstore.RoundERC20DonationRecord{erc20}, false, nil
		},
		nft: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error) {
			return []cgstore.RoundNFTDonationRecord{nft}, false, nil
		},
	})

	erc20Response := serve(t, server, "/api/v2/cosmicgame/rounds/3/erc20-donations")
	if erc20Response.Code != http.StatusOK {
		t.Fatalf("ERC-20 status = %d, body=%s", erc20Response.Code, erc20Response.Body.String())
	}
	var erc20Page RoundErc20DonationPage
	decodeResponse(t, erc20Response, &erc20Page)
	if len(erc20Page.Data) != 1 || erc20Page.Data[0].AmountBaseUnits != "42" ||
		erc20Page.Meta.NextCursor != nil {
		t.Fatalf("ERC-20 page = %+v", erc20Page)
	}

	nftResponse := serve(t, server, "/api/v2/cosmicgame/rounds/3/nft-donations")
	if nftResponse.Code != http.StatusOK {
		t.Fatalf("NFT status = %d, body=%s", nftResponse.Code, nftResponse.Body.String())
	}
	var nftPage RoundNftDonationPage
	decodeResponse(t, nftResponse, &nftPage)
	if len(nftPage.Data) != 1 || nftPage.Data[0].TokenId != 777 ||
		nftPage.Meta.NextCursor != nil {
		t.Fatalf("NFT page = %+v", nftPage)
	}
}

func TestListRoundDonationsDecodesContinuationCursor(t *testing.T) {
	t.Parallel()

	encoded, err := encodeRoundDonationCursor(roundDonationCursor{
		Version:    roundDonationCursorVersion,
		Round:      3,
		Resource:   roundDonationResourceERC20,
		EventLogID: 88,
	})
	if err != nil {
		t.Fatal(err)
	}
	var gotAfter *cgstore.DonationPageCursor
	server := newDonationTestServer(t, fakeRoundDonationReader{
		erc20: func(_ context.Context, _ int64, after *cgstore.DonationPageCursor, limit int) ([]cgstore.RoundERC20DonationRecord, bool, error) {
			gotAfter = after
			if limit != defaultPageLimit {
				t.Errorf("default limit = %d, want %d", limit, defaultPageLimit)
			}
			return []cgstore.RoundERC20DonationRecord{}, false, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/3/erc20-donations?cursor="+encoded)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter == nil || gotAfter.EventLogID != 88 {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
	if !bytes.Contains(response.Body.Bytes(), []byte(`"data":[]`)) {
		t.Fatalf("empty data was not encoded as []: %s", response.Body.String())
	}
}

func TestListRoundDonationsAllowEmptyOpenRound(t *testing.T) {
	t.Parallel()

	server := newDonationTestServer(t, fakeRoundDonationReader{})
	for _, resource := range []string{"eth-donations", "erc20-donations", "nft-donations"} {
		t.Run(resource, func(t *testing.T) {
			t.Parallel()
			response := serve(t, server, "/api/v2/cosmicgame/rounds/3/"+resource)
			if response.Code != http.StatusOK || !bytes.Contains(response.Body.Bytes(), []byte(`"data":[]`)) {
				t.Fatalf("response = %d %s", response.Code, response.Body.String())
			}
		})
	}
}

func TestListRoundDonationsRejectInvalidInput(t *testing.T) {
	t.Parallel()

	crossRound, err := encodeRoundDonationCursor(roundDonationCursor{
		Version:    roundDonationCursorVersion,
		Round:      2,
		Resource:   roundDonationResourceETH,
		EventLogID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	crossResource, err := encodeRoundDonationCursor(roundDonationCursor{
		Version:    roundDonationCursorVersion,
		Round:      1,
		Resource:   roundDonationResourceNFT,
		EventLogID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]string{
		"negative round":   "/api/v2/cosmicgame/rounds/-1/eth-donations",
		"zero limit":       "/api/v2/cosmicgame/rounds/1/eth-donations?limit=0",
		"excessive limit":  "/api/v2/cosmicgame/rounds/1/erc20-donations?limit=201",
		"duplicate limit":  "/api/v2/cosmicgame/rounds/1/nft-donations?limit=1&limit=2",
		"malformed cursor": "/api/v2/cosmicgame/rounds/1/eth-donations?cursor=not-a-cursor",
		"oversized cursor": "/api/v2/cosmicgame/rounds/1/erc20-donations?cursor=" + strings.Repeat("a", maxCursorLength+1),
		"cross round":      "/api/v2/cosmicgame/rounds/1/eth-donations?cursor=" + crossRound,
		"cross resource":   "/api/v2/cosmicgame/rounds/1/eth-donations?cursor=" + crossResource,
		"bind round":       "/api/v2/cosmicgame/rounds/not-a-number/nft-donations",
		"bind limit":       "/api/v2/cosmicgame/rounds/1/erc20-donations?limit=wat",
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newDonationTestServer(t, fakeRoundDonationReader{}), path)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListRoundDonationsHideRepositoryErrors(t *testing.T) {
	t.Parallel()

	secretErr := errors.New("password=super-secret")
	tests := map[string]struct {
		path   string
		reader fakeRoundDonationReader
	}{
		"eth": {
			path: "/api/v2/cosmicgame/rounds/1/eth-donations",
			reader: fakeRoundDonationReader{eth: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
				return nil, false, secretErr
			}},
		},
		"erc20": {
			path: "/api/v2/cosmicgame/rounds/1/erc20-donations",
			reader: fakeRoundDonationReader{erc20: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error) {
				return nil, false, secretErr
			}},
		},
		"nft": {
			path: "/api/v2/cosmicgame/rounds/1/nft-donations",
			reader: fakeRoundDonationReader{nft: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error) {
				return nil, false, secretErr
			}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newDonationTestServer(t, tc.reader), tc.path)
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "super-secret") {
				t.Fatalf("internal error leaked: %s", response.Body.String())
			}
		})
	}
}

func TestListRoundDonationsRejectInconsistentRepositoryPages(t *testing.T) {
	t.Parallel()

	tests := map[string]func() fakeRoundDonationReader{
		"has more without row": func() fakeRoundDonationReader {
			return fakeRoundDonationReader{eth: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
				return []cgstore.RoundEthDonationRecord{}, true, nil
			}}
		},
		"wrong round": func() fakeRoundDonationReader {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			record.RoundNum = 2
			return fakeRoundDonationReader{eth: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
				return []cgstore.RoundEthDonationRecord{record}, false, nil
			}}
		},
		"out of order": func() fakeRoundDonationReader {
			first := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			first.RoundNum, first.Tx.EvtLogId = 1, 10
			second := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			second.RoundNum, second.Tx.EvtLogId = 1, 11
			return fakeRoundDonationReader{eth: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
				return []cgstore.RoundEthDonationRecord{first, second}, false, nil
			}}
		},
		"invalid mapped record": func() fakeRoundDonationReader {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			record.RoundNum = 1
			record.DonorAddr = "invalid"
			return fakeRoundDonationReader{eth: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
				return []cgstore.RoundEthDonationRecord{record}, false, nil
			}}
		},
	}
	for name, build := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newDonationTestServer(t, build()), "/api/v2/cosmicgame/rounds/1/eth-donations")
			assertProblem(t, response, http.StatusInternalServerError)
		})
	}
}

func TestListRoundDonationsPropagateCancelledContextAsOpaqueError(t *testing.T) {
	t.Parallel()

	server := newDonationTestServer(t, fakeRoundDonationReader{
		nft: func(ctx context.Context, _ int64, _ *cgstore.DonationPageCursor, _ int) ([]cgstore.RoundNFTDonationRecord, bool, error) {
			return nil, false, ctx.Err()
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	response := serveDonationContext(t, server, ctx, "/api/v2/cosmicgame/rounds/1/nft-donations")
	assertProblem(t, response, http.StatusInternalServerError)
}

func newDonationTestServer(t *testing.T, donations roundDonationReader) *Server {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		donations,
		fakeStatisticsReader{},
		fakeContractState{},
		logger,
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func serveDonationContext(
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
