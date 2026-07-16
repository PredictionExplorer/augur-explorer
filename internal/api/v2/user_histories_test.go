package v2

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func newUserHistoryTestServer(t *testing.T, histories userHistoryReader) *Server {
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
		histories,
		fakeUserStakingReader{},
		fakeUserActivityReader{},
		fakeGlobalDirectoryReader{},
		fakeGlobalStakingReader{},
		fakeRandomWalkReader{},
		fakeRankingRepository{},
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func userHistoryPrizeRecord(round, prizeType, winnerIndex, eventLogID int64) cgmodel.CGPrizeHistory {
	record := cgmodel.CGPrizeHistory{
		RecordType:  prizeType,
		RoundNum:    round,
		WinnerIndex: winnerIndex,
		Amount:      "50000000000000000",
		TokenId:     -1,
		WinnerAddr:  userCursorAlice,
		WinnerAid:   1,
	}
	record.Tx = validDonationTransaction()
	record.Tx.EvtLogId = eventLogID
	return record
}

func TestListUserPrizesPaginatesWithTupleCursor(t *testing.T) {
	t.Parallel()

	var gotAfter *cgstore.UserPrizePageCursor
	var gotLimit int
	first := userHistoryPrizeRecord(5, 0, 0, 300)
	second := userHistoryPrizeRecord(4, 7, 2, 290)
	server := newUserHistoryTestServer(t, fakeUserHistoryReader{
		prizes: func(_ context.Context, userAid int64, after *cgstore.UserPrizePageCursor, limit int) ([]cgmodel.CGPrizeHistory, bool, error) {
			if userAid != 1 {
				t.Errorf("user aid = %d", userAid)
			}
			gotAfter, gotLimit = after, limit
			return []cgmodel.CGPrizeHistory{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/users/"+userCursorAlice+"/prizes?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter != nil || gotLimit != 2 {
		t.Fatalf("repository args = (%+v, %d)", gotAfter, gotLimit)
	}
	var page CosmicGameUserPrizePage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	if page.Data[0].Type != MainPrizeEth || page.Data[1].Type != ChronoWarriorEth {
		t.Fatalf("prize types = %v, %v", page.Data[0].Type, page.Data[1].Type)
	}
	cursor, err := decodeUserPrizeCursor(*page.Meta.NextCursor, userCursorAlice)
	if err != nil || cursor.Round != 4 || cursor.PrizeType != 7 || cursor.WinnerIndex != 2 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}

	// The continuation cursor decodes into the repository cursor.
	server = newUserHistoryTestServer(t, fakeUserHistoryReader{
		prizes: func(_ context.Context, _ int64, after *cgstore.UserPrizePageCursor, _ int) ([]cgmodel.CGPrizeHistory, bool, error) {
			gotAfter = after
			return []cgmodel.CGPrizeHistory{}, false, nil
		},
	})
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/prizes?cursor="+*page.Meta.NextCursor)
	if response.Code != http.StatusOK {
		t.Fatalf("continuation status = %d", response.Code)
	}
	if gotAfter == nil || gotAfter.Round != 4 || gotAfter.PrizeType != 7 || gotAfter.WinnerIndex != 2 {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
}

func TestListUserPrizesEmptyForUnindexedWallet(t *testing.T) {
	t.Parallel()

	server := newUserHistoryTestServer(t, fakeUserHistoryReader{
		addressID: func(context.Context, string) (int64, error) {
			return 0, store.ErrNotFound
		},
		prizes: func(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
			t.Error("prize page fetched for an unindexed wallet")
			return nil, false, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/users/"+userCursorAlice+"/prizes")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	var page CosmicGameUserPrizePage
	decodeResponse(t, response, &page)
	if len(page.Data) != 0 || page.Meta.NextCursor != nil {
		t.Fatalf("page = %+v", page)
	}
}

func TestListUserPrizesRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	bobCursor, err := encodeUserPrizeCursor(userPrizeCursor{
		Version: userPrizeCursorVersion, Address: userCursorBob,
		Round: 1, PrizeType: 0, WinnerIndex: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]string{
		"invalid address":   "/api/v2/cosmicgame/users/not-an-address/prizes",
		"invalid limit":     "/api/v2/cosmicgame/users/" + userCursorAlice + "/prizes?limit=0",
		"bind limit":        "/api/v2/cosmicgame/users/" + userCursorAlice + "/prizes?limit=wat",
		"malformed cursor":  "/api/v2/cosmicgame/users/" + userCursorAlice + "/prizes?cursor=bad",
		"cross-user cursor": "/api/v2/cosmicgame/users/" + userCursorAlice + "/prizes?cursor=" + bobCursor,
	}
	for name, target := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserHistoryTestServer(t, fakeUserHistoryReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListUserPrizesFailsClosedOnRepositoryViolations(t *testing.T) {
	t.Parallel()

	tests := map[string]fakeUserHistoryReader{
		"address resolution failure": {
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("secret db detail")
			},
		},
		"repository failure": {
			prizes: func(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
				return nil, false, errors.New("secret db detail")
			},
		},
		"over cardinality": {
			prizes: func(_ context.Context, _ int64, _ *cgstore.UserPrizePageCursor, limit int) ([]cgmodel.CGPrizeHistory, bool, error) {
				records := make([]cgmodel.CGPrizeHistory, limit+1)
				for i := range records {
					records[i] = userHistoryPrizeRecord(int64(100-i), 0, 0, int64(400-i))
				}
				return records, false, nil
			},
		},
		"foreign winner": {
			prizes: func(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
				record := userHistoryPrizeRecord(5, 0, 0, 300)
				record.WinnerAid = 77
				return []cgmodel.CGPrizeHistory{record}, false, nil
			},
		},
		"foreign winner address": {
			prizes: func(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
				record := userHistoryPrizeRecord(5, 0, 0, 300)
				record.WinnerAddr = userCursorBob
				return []cgmodel.CGPrizeHistory{record}, false, nil
			},
		},
		"unordered rows": {
			prizes: func(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
				return []cgmodel.CGPrizeHistory{
					userHistoryPrizeRecord(4, 7, 2, 290),
					userHistoryPrizeRecord(5, 0, 0, 300),
				}, false, nil
			},
		},
		"has more without rows": {
			prizes: func(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
				return []cgmodel.CGPrizeHistory{}, true, nil
			},
		},
		"unmappable row": {
			prizes: func(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
				record := userHistoryPrizeRecord(5, 0, 0, 300)
				record.Amount = "not-a-number"
				return []cgmodel.CGPrizeHistory{record}, false, nil
			},
		},
	}
	for name, reader := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserHistoryTestServer(t, reader),
				"/api/v2/cosmicgame/users/"+userCursorAlice+"/prizes")
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "secret db detail") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}
}

func userHistoryDepositRecord(eventLogID int64) cgstore.UserRaffleEthDepositRecord {
	record := validUserRaffleEthDepositRecord()
	record.WinnerAid = 1
	record.WinnerAddr = userCursorAlice
	record.Tx.EvtLogId = eventLogID
	return record
}

func TestListUserRaffleEthDepositsPaginatesAndFilters(t *testing.T) {
	t.Parallel()

	var gotClaimed *bool
	var gotAfter *cgstore.UserEventPageCursor
	first := userHistoryDepositRecord(300)
	first.Claimed = true
	withdrawal := validUserDepositWithdrawalRecord()
	first.Withdrawal = &withdrawal
	second := userHistoryDepositRecord(290)
	second.Claimed = true
	withdrawalCopy := validUserDepositWithdrawalRecord()
	second.Withdrawal = &withdrawalCopy
	server := newUserHistoryTestServer(t, fakeUserHistoryReader{
		deposits: func(_ context.Context, _ int64, claimed *bool, after *cgstore.UserEventPageCursor, _ int) ([]cgstore.UserRaffleEthDepositRecord, bool, error) {
			gotClaimed, gotAfter = claimed, after
			return []cgstore.UserRaffleEthDepositRecord{first, second}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/raffle-eth-deposits?claimed=true&limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotClaimed == nil || !*gotClaimed || gotAfter != nil {
		t.Fatalf("repository args = claimed=%v after=%+v", gotClaimed, gotAfter)
	}
	var page CosmicGameUserRaffleEthDepositPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	if !page.Data[0].Claimed || page.Data[0].Withdrawal == nil {
		t.Fatalf("deposit = %+v", page.Data[0])
	}
	cursor, err := decodeUserEventCursor(
		*page.Meta.NextCursor, userCursorAlice, userEventResourceRaffleEthDeposits)
	if err != nil || cursor.EventLogID != 290 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}
}

func TestListUserRaffleEthDepositsRejectsFilterViolation(t *testing.T) {
	t.Parallel()

	// A repository row that contradicts the requested claimed filter is an
	// internal inconsistency, never data to serve.
	server := newUserHistoryTestServer(t, fakeUserHistoryReader{
		deposits: func(context.Context, int64, *bool, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleEthDepositRecord, bool, error) {
			return []cgstore.UserRaffleEthDepositRecord{userHistoryDepositRecord(300)}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/raffle-eth-deposits?claimed=true")
	assertProblem(t, response, http.StatusInternalServerError)
}

func TestUserEventResourcesShareUniformSemantics(t *testing.T) {
	t.Parallel()

	// Every event-keyed history behaves identically at the edge; drive the
	// shared generic through each generated operation once.
	nftWin := func(eventLogID int64) cgstore.UserRaffleNftWinRecord {
		record := validUserRaffleNftWinRecord()
		record.WinnerAid = 1
		record.WinnerAddr = userCursorAlice
		record.Tx.EvtLogId = eventLogID
		return record
	}
	ethDonation := func(eventLogID int64) cgstore.RoundEthDonationRecord {
		record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
		record.DonorAddr = userCursorAlice
		record.Tx.EvtLogId = eventLogID
		return record
	}
	ercDonation := func(eventLogID int64) cgstore.RoundERC20DonationRecord {
		record := validRoundERC20DonationRecord()
		record.DonorAddr = userCursorAlice
		record.Tx.EvtLogId = eventLogID
		return record
	}
	nftDonation := func(eventLogID int64) cgstore.RoundNFTDonationRecord {
		record := validRoundNFTDonationRecord()
		record.DonorAddr = userCursorAlice
		record.Tx.EvtLogId = eventLogID
		return record
	}
	donatedNft := func(eventLogID int64) cgstore.UserDonatedNftRecord {
		record := validUserDonatedNftRecord()
		record.RoundWinnerAid = 1
		record.Tx.EvtLogId = eventLogID
		return record
	}

	resources := map[string]struct {
		pathSegment string
		resource    userEventResource
		reader      fakeUserHistoryReader
	}{
		"raffle NFT wins": {
			pathSegment: "raffle-nft-wins",
			resource:    userEventResourceRaffleNftWins,
			reader: fakeUserHistoryReader{
				nftWins: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleNftWinRecord, bool, error) {
					return []cgstore.UserRaffleNftWinRecord{nftWin(300)}, true, nil
				},
			},
		},
		"eth donations": {
			pathSegment: "eth-donations",
			resource:    userEventResourceEthDonations,
			reader: fakeUserHistoryReader{
				ethDonations: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
					return []cgstore.RoundEthDonationRecord{ethDonation(300)}, true, nil
				},
			},
		},
		"erc20 donations": {
			pathSegment: "erc20-donations",
			resource:    userEventResourceErc20Donations,
			reader: fakeUserHistoryReader{
				ercDonations: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error) {
					return []cgstore.RoundERC20DonationRecord{ercDonation(300)}, true, nil
				},
			},
		},
		"nft donations": {
			pathSegment: "nft-donations",
			resource:    userEventResourceNftDonations,
			reader: fakeUserHistoryReader{
				nftDonations: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error) {
					return []cgstore.RoundNFTDonationRecord{nftDonation(300)}, true, nil
				},
			},
		},
		"donated nfts": {
			pathSegment: "donated-nfts",
			resource:    userEventResourceDonatedNfts,
			reader: fakeUserHistoryReader{
				donatedNfts: func(context.Context, int64, *bool, *cgstore.UserEventPageCursor, int) ([]cgstore.UserDonatedNftRecord, bool, error) {
					return []cgstore.UserDonatedNftRecord{donatedNft(300)}, true, nil
				},
			},
		},
	}
	for name, test := range resources {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			base := "/api/v2/cosmicgame/users/" + userCursorAlice + "/" + test.pathSegment

			response := serve(t, newUserHistoryTestServer(t, test.reader), base+"?limit=1")
			if response.Code != http.StatusOK {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			body := response.Body.String()
			if !strings.Contains(body, `"nextCursor":"`) {
				t.Fatalf("page did not include a continuation cursor: %s", body)
			}

			// The advertised cursor is scoped to this resource and wallet.
			var page struct {
				Meta PageMeta `json:"meta"`
			}
			decodeResponse(t, response, &page)
			cursor, err := decodeUserEventCursor(*page.Meta.NextCursor, userCursorAlice, test.resource)
			if err != nil || cursor.EventLogID != 300 {
				t.Fatalf("next cursor = %+v, err=%v", cursor, err)
			}
			foreign, err := encodeUserEventCursor(userEventCursor{
				Version:    userEventCursorVersion,
				Address:    userCursorAlice,
				Resource:   foreignUserEventResource(test.resource),
				EventLogID: 300,
			})
			if err != nil {
				t.Fatal(err)
			}
			response = serve(t, newUserHistoryTestServer(t, test.reader), base+"?cursor="+foreign)
			assertProblem(t, response, http.StatusBadRequest)

			// Unindexed wallets answer an empty page.
			empty := test.reader
			empty.addressID = func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			}
			response = serve(t, newUserHistoryTestServer(t, empty), base)
			if response.Code != http.StatusOK || !strings.Contains(response.Body.String(), `"data":[]`) {
				t.Fatalf("unindexed wallet = %d %s", response.Code, response.Body.String())
			}

			// Malformed addresses are client errors.
			response = serve(t, newUserHistoryTestServer(t, test.reader),
				"/api/v2/cosmicgame/users/not-an-address/"+test.pathSegment)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func foreignUserEventResource(resource userEventResource) userEventResource {
	if resource == userEventResourceRaffleNftWins {
		return userEventResourceEthDonations
	}
	return userEventResourceRaffleNftWins
}

func TestUserEventResourcesFailClosedOnScopeViolations(t *testing.T) {
	t.Parallel()

	t.Run("foreign donor", func(t *testing.T) {
		t.Parallel()
		record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
		record.DonorAddr = userCursorBob
		server := newUserHistoryTestServer(t, fakeUserHistoryReader{
			ethDonations: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
				return []cgstore.RoundEthDonationRecord{record}, false, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/eth-donations")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("unordered events", func(t *testing.T) {
		t.Parallel()
		first := validRoundNFTDonationRecord()
		first.DonorAddr = userCursorAlice
		first.Tx.EvtLogId = 290
		second := validRoundNFTDonationRecord()
		second.DonorAddr = userCursorAlice
		second.Tx.EvtLogId = 300
		server := newUserHistoryTestServer(t, fakeUserHistoryReader{
			nftDonations: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error) {
				return []cgstore.RoundNFTDonationRecord{first, second}, false, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/nft-donations")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("donated nft outside both branches", func(t *testing.T) {
		t.Parallel()
		record := validUserDonatedNftRecord()
		record.RoundWinnerAid = 77
		server := newUserHistoryTestServer(t, fakeUserHistoryReader{
			donatedNfts: func(context.Context, int64, *bool, *cgstore.UserEventPageCursor, int) ([]cgstore.UserDonatedNftRecord, bool, error) {
				return []cgstore.UserDonatedNftRecord{record}, false, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/donated-nfts")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("repository failure stays opaque", func(t *testing.T) {
		t.Parallel()
		server := newUserHistoryTestServer(t, fakeUserHistoryReader{
			nftWins: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleNftWinRecord, bool, error) {
				return nil, false, errors.New("password=super-secret")
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/raffle-nft-wins")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "super-secret") {
			t.Fatalf("internal error leaked: %s", response.Body.String())
		}
	})
}

func TestUserEventResourcesFailClosedOnRepositoryContractViolations(t *testing.T) {
	t.Parallel()

	base := "/api/v2/cosmicgame/users/" + userCursorAlice + "/raffle-nft-wins"
	inScopeWin := func(eventLogID int64) cgstore.UserRaffleNftWinRecord {
		record := validUserRaffleNftWinRecord()
		record.WinnerAid = 1
		record.WinnerAddr = userCursorAlice
		record.Tx.EvtLogId = eventLogID
		return record
	}
	tests := map[string]fakeUserHistoryReader{
		"address resolution failure": {
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("secret db detail")
			},
		},
		"over cardinality": {
			nftWins: func(_ context.Context, _ int64, _ *cgstore.UserEventPageCursor, limit int) ([]cgstore.UserRaffleNftWinRecord, bool, error) {
				records := make([]cgstore.UserRaffleNftWinRecord, limit+1)
				for i := range records {
					records[i] = inScopeWin(int64(1000 - i))
				}
				return records, false, nil
			},
		},
		"unmappable row": {
			nftWins: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleNftWinRecord, bool, error) {
				record := inScopeWin(300)
				record.CstAmountWei = "not-a-number"
				return []cgstore.UserRaffleNftWinRecord{record}, false, nil
			},
		},
		"has more without rows": {
			nftWins: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleNftWinRecord, bool, error) {
				return []cgstore.UserRaffleNftWinRecord{}, true, nil
			},
		},
	}
	for name, reader := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserHistoryTestServer(t, reader), base)
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "secret db detail") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}
}

func TestListUserDonatedErc20RejectsInvalidInput(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"invalid address": "/api/v2/cosmicgame/users/not-an-address/donated-erc20",
		"invalid limit":   "/api/v2/cosmicgame/users/" + userCursorAlice + "/donated-erc20?limit=201",
	}
	for name, target := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserHistoryTestServer(t, fakeUserHistoryReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestBiddingActivityAndTypeRatioRejectInvertedWindows(t *testing.T) {
	t.Parallel()

	for _, target := range []string{
		"/api/v2/cosmicgame/statistics/bidding/activity?from=100&to=50",
		"/api/v2/cosmicgame/statistics/bidding/type-ratio?from=100&to=50",
	} {
		t.Run(target, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserHistoryTestServer(t, fakeUserHistoryReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestUserHistoriesRejectDuplicateQueryParameters(t *testing.T) {
	t.Parallel()

	base := "/api/v2/cosmicgame/users/" + userCursorAlice
	targets := []string{
		base + "/prizes?cursor=a&cursor=b",
		base + "/raffle-eth-deposits?cursor=a&cursor=b",
		base + "/raffle-eth-deposits?claimed=true&claimed=false",
		base + "/raffle-nft-wins?cursor=a&cursor=b",
		base + "/eth-donations?cursor=a&cursor=b",
		base + "/erc20-donations?cursor=a&cursor=b",
		base + "/nft-donations?cursor=a&cursor=b",
		base + "/donated-nfts?cursor=a&cursor=b",
		base + "/donated-nfts?status=claimed&status=unclaimed",
		base + "/donated-erc20?cursor=a&cursor=b",
		base + "/prizes?limit=1&limit=2",
	}
	for _, target := range targets {
		t.Run(target, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserHistoryTestServer(t, fakeUserHistoryReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListUserDonatedNftsAppliesStatusFilter(t *testing.T) {
	t.Parallel()

	var gotClaimed *bool
	timeoutClaim := validUserDonatedNftRecord()
	timeoutClaim.RoundWinnerAid = 77 // someone else won the round
	timeoutClaim.Claimed = true
	claim := validUserDonatedNftClaimRecord()
	claim.ClaimerAid = 1 // ... but the requested wallet claimed the NFT
	timeoutClaim.Claim = &claim
	server := newUserHistoryTestServer(t, fakeUserHistoryReader{
		donatedNfts: func(_ context.Context, _ int64, claimed *bool, _ *cgstore.UserEventPageCursor, _ int) ([]cgstore.UserDonatedNftRecord, bool, error) {
			gotClaimed = claimed
			return []cgstore.UserDonatedNftRecord{timeoutClaim}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/donated-nfts?status=claimed")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotClaimed == nil || !*gotClaimed {
		t.Fatalf("repository claimed filter = %v", gotClaimed)
	}
	var page CosmicGameUserDonatedNftPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 1 || !page.Data[0].Claimed || page.Data[0].Claim == nil {
		t.Fatalf("page = %+v", page)
	}

	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/donated-nfts?status=unclaimed")
	if gotClaimed == nil || *gotClaimed {
		t.Fatalf("repository unclaimed filter = %v", gotClaimed)
	}
	// The fake still returns a claimed row: a filter contradiction fails closed.
	assertProblem(t, response, http.StatusInternalServerError)

	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/donated-nfts?status=wat")
	assertProblem(t, response, http.StatusBadRequest)
}

func userHistoryDonatedErc20Record(round, tokenAid int64) cgstore.UserDonatedErc20Record {
	record := validUserDonatedErc20Record()
	record.RoundNum = round
	record.TokenAid = tokenAid
	return record
}

func TestListUserDonatedErc20PaginatesWithRoundTokenCursor(t *testing.T) {
	t.Parallel()

	var gotAfter *cgstore.UserDonatedErc20PageCursor
	server := newUserHistoryTestServer(t, fakeUserHistoryReader{
		donatedTokens: func(_ context.Context, _ int64, after *cgstore.UserDonatedErc20PageCursor, _ int) ([]cgstore.UserDonatedErc20Record, bool, error) {
			gotAfter = after
			if after != nil {
				return []cgstore.UserDonatedErc20Record{
					userHistoryDonatedErc20Record(3, 26),
				}, false, nil
			}
			return []cgstore.UserDonatedErc20Record{
				userHistoryDonatedErc20Record(5, 26),
				userHistoryDonatedErc20Record(5, 30),
				userHistoryDonatedErc20Record(4, 26),
			}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/donated-erc20?limit=3")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page CosmicGameUserDonatedErc20Page
	decodeResponse(t, response, &page)
	if len(page.Data) != 3 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeUserDonatedErc20Cursor(*page.Meta.NextCursor, userCursorAlice)
	if err != nil || cursor.Round != 4 || cursor.TokenID != 26 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}

	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/donated-erc20?cursor="+*page.Meta.NextCursor)
	if response.Code != http.StatusOK {
		t.Fatalf("continuation status = %d", response.Code)
	}
	if gotAfter == nil || gotAfter.Round != 4 || gotAfter.TokenAid != 26 {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
}

func TestListUserDonatedErc20FailsClosedOnViolations(t *testing.T) {
	t.Parallel()

	tests := map[string]fakeUserHistoryReader{
		"address resolution failure": {
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("secret db detail")
			},
		},
		"over cardinality": {
			donatedTokens: func(_ context.Context, _ int64, _ *cgstore.UserDonatedErc20PageCursor, limit int) ([]cgstore.UserDonatedErc20Record, bool, error) {
				records := make([]cgstore.UserDonatedErc20Record, limit+1)
				for i := range records {
					records[i] = userHistoryDonatedErc20Record(int64(1000-i), 26)
				}
				return records, false, nil
			},
		},
		"has more without rows": {
			donatedTokens: func(context.Context, int64, *cgstore.UserDonatedErc20PageCursor, int) ([]cgstore.UserDonatedErc20Record, bool, error) {
				return []cgstore.UserDonatedErc20Record{}, true, nil
			},
		},
		"unordered summaries": {
			donatedTokens: func(context.Context, int64, *cgstore.UserDonatedErc20PageCursor, int) ([]cgstore.UserDonatedErc20Record, bool, error) {
				return []cgstore.UserDonatedErc20Record{
					userHistoryDonatedErc20Record(4, 26),
					userHistoryDonatedErc20Record(5, 26),
				}, false, nil
			},
		},
		"duplicate keys": {
			donatedTokens: func(context.Context, int64, *cgstore.UserDonatedErc20PageCursor, int) ([]cgstore.UserDonatedErc20Record, bool, error) {
				return []cgstore.UserDonatedErc20Record{
					userHistoryDonatedErc20Record(5, 26),
					userHistoryDonatedErc20Record(5, 26),
				}, false, nil
			},
		},
		"invalid token id": {
			donatedTokens: func(context.Context, int64, *cgstore.UserDonatedErc20PageCursor, int) ([]cgstore.UserDonatedErc20Record, bool, error) {
				return []cgstore.UserDonatedErc20Record{userHistoryDonatedErc20Record(5, 0)}, false, nil
			},
		},
		"inconsistent totals": {
			donatedTokens: func(context.Context, int64, *cgstore.UserDonatedErc20PageCursor, int) ([]cgstore.UserDonatedErc20Record, bool, error) {
				record := userHistoryDonatedErc20Record(5, 26)
				record.RemainingBaseUnits = "999"
				return []cgstore.UserDonatedErc20Record{record}, false, nil
			},
		},
		"repository failure": {
			donatedTokens: func(context.Context, int64, *cgstore.UserDonatedErc20PageCursor, int) ([]cgstore.UserDonatedErc20Record, bool, error) {
				return nil, false, errors.New("secret db detail")
			},
		},
	}
	for name, reader := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserHistoryTestServer(t, reader),
				"/api/v2/cosmicgame/users/"+userCursorAlice+"/donated-erc20")
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "secret db detail") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}

	t.Run("unindexed wallet answers empty page", func(t *testing.T) {
		t.Parallel()
		server := newUserHistoryTestServer(t, fakeUserHistoryReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/donated-erc20")
		if response.Code != http.StatusOK || !strings.Contains(response.Body.String(), `"data":[]`) {
			t.Fatalf("unindexed wallet = %d %s", response.Code, response.Body.String())
		}
	})
}
