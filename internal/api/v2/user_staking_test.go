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

func newUserStakingTestServer(t *testing.T, staking userStakingReader) *Server {
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
		staking,
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func userStakingActionAt(kind cgstore.UserStakingActionKind, eventLogID int64) cgstore.UserStakingActionRecord {
	record := validUserStakingActionRecord(kind)
	record.Tx.EvtLogId = eventLogID
	return record
}

func TestListUserCstStakingActionsPaginates(t *testing.T) {
	t.Parallel()

	var gotAfter *cgstore.UserEventPageCursor
	var gotLimit int
	server := newUserStakingTestServer(t, fakeUserStakingReader{
		cstActions: func(_ context.Context, userAid int64, after *cgstore.UserEventPageCursor, limit int) ([]cgstore.UserStakingActionRecord, bool, error) {
			if userAid != 1 {
				t.Errorf("user aid = %d", userAid)
			}
			gotAfter, gotLimit = after, limit
			if after != nil {
				return []cgstore.UserStakingActionRecord{}, false, nil
			}
			return []cgstore.UserStakingActionRecord{
				userStakingActionAt(cgstore.UserStakingActionUnstake, 300),
				userStakingActionAt(cgstore.UserStakingActionStake, 290),
			}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/actions?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter != nil || gotLimit != 2 {
		t.Fatalf("repository args = (%+v, %d)", gotAfter, gotLimit)
	}
	var page CosmicGameUserCstStakingActionPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	if page.Data[0].ActionType != Unstake || page.Data[0].RewardWei == nil ||
		*page.Data[0].RewardWei != "1000000000000000000" {
		t.Fatalf("unstake item = %+v", page.Data[0])
	}
	if page.Data[1].ActionType != Stake || page.Data[1].RewardWei != nil {
		t.Fatalf("stake item = %+v", page.Data[1])
	}
	cursor, err := decodeUserEventCursor(
		*page.Meta.NextCursor, userCursorAlice, userEventResourceCstStakingActions)
	if err != nil || cursor.EventLogID != 290 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}

	// The continuation cursor decodes into the repository cursor and cannot
	// cross into the RandomWalk actions resource.
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/actions?cursor="+*page.Meta.NextCursor)
	if response.Code != http.StatusOK || gotAfter == nil || gotAfter.EventLogID != 290 {
		t.Fatalf("continuation: status=%d cursor=%+v", response.Code, gotAfter)
	}
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/random-walk/actions?cursor="+*page.Meta.NextCursor)
	assertProblem(t, response, http.StatusBadRequest)
}

func TestListUserRandomWalkStakingActions(t *testing.T) {
	t.Parallel()

	rwalkAction := userStakingActionAt(cgstore.UserStakingActionUnstake, 300)
	rwalkAction.RewardWei = ""
	server := newUserStakingTestServer(t, fakeUserStakingReader{
		rwalkActions: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error) {
			return []cgstore.UserStakingActionRecord{rwalkAction}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/random-walk/actions")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page CosmicGameUserRandomWalkStakingActionPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 1 || page.Data[0].ActionType != Unstake || page.Meta.NextCursor != nil {
		t.Fatalf("page = %+v", page)
	}
	if strings.Contains(response.Body.String(), "rewardWei") {
		t.Fatalf("RandomWalk actions leaked a reward field: %s", response.Body.String())
	}

	// A reward-carrying RandomWalk row is an internal inconsistency.
	server = newUserStakingTestServer(t, fakeUserStakingReader{
		rwalkActions: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error) {
			return []cgstore.UserStakingActionRecord{userStakingActionAt(cgstore.UserStakingActionUnstake, 300)}, false, nil
		},
	})
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/random-walk/actions")
	assertProblem(t, response, http.StatusInternalServerError)
}

func TestUserStakingActionsFailClosedOnScopeViolations(t *testing.T) {
	t.Parallel()

	t.Run("foreign staker", func(t *testing.T) {
		t.Parallel()
		record := userStakingActionAt(cgstore.UserStakingActionStake, 300)
		record.StakerAid = 77
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			cstActions: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error) {
				return []cgstore.UserStakingActionRecord{record}, false, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/actions")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("unordered events", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			cstActions: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error) {
				return []cgstore.UserStakingActionRecord{
					userStakingActionAt(cgstore.UserStakingActionStake, 290),
					userStakingActionAt(cgstore.UserStakingActionUnstake, 300),
				}, false, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/actions")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("repository failure stays opaque", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			cstActions: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error) {
				return nil, false, errors.New("password=super-secret")
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/actions")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "super-secret") {
			t.Fatalf("internal error leaked: %s", response.Body.String())
		}
	})
}

func userStakedCstTokenAt(tokenID int64) cgstore.UserStakedCstTokenRecord {
	record := validUserStakedCstTokenRecord()
	record.TokenID = tokenID
	return record
}

func TestListUserStakedCstTokensPaginates(t *testing.T) {
	t.Parallel()

	var gotAfter *cgstore.UserStakingTokenPageCursor
	server := newUserStakingTestServer(t, fakeUserStakingReader{
		cstStaked: func(_ context.Context, _ int64, after *cgstore.UserStakingTokenPageCursor, _ int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
			gotAfter = after
			if after != nil {
				return []cgstore.UserStakedCstTokenRecord{}, false, nil
			}
			return []cgstore.UserStakedCstTokenRecord{
				userStakedCstTokenAt(5),
				userStakedCstTokenAt(9),
			}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/staked-tokens?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page CosmicGameUserStakedCstTokenPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeUserStakingTokenCursor(
		*page.Meta.NextCursor, userCursorAlice, userStakingTokenResourceCstStakedTokens)
	if err != nil || cursor.TokenID != 9 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}

	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/staked-tokens?cursor="+*page.Meta.NextCursor)
	if response.Code != http.StatusOK || gotAfter == nil || gotAfter.TokenID != 9 {
		t.Fatalf("continuation: status=%d cursor=%+v", response.Code, gotAfter)
	}

	// The cursor is rejected by the two sibling token-keyed resources.
	for _, target := range []string{
		"/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/random-walk/staked-tokens?cursor=" + *page.Meta.NextCursor,
		"/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/token-rewards?cursor=" + *page.Meta.NextCursor,
	} {
		response = serve(t, newUserStakingTestServer(t, fakeUserStakingReader{}), target)
		assertProblem(t, response, http.StatusBadRequest)
	}
}

func TestUserStakingTokenPagesFailClosedOnViolations(t *testing.T) {
	t.Parallel()

	tests := map[string]fakeUserStakingReader{
		"address resolution failure": {
			addressID: func(context.Context, string) (int64, error) {
				return 0, errors.New("secret db detail")
			},
		},
		"repository failure": {
			cstStaked: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
				return nil, false, errors.New("secret db detail")
			},
		},
		"over cardinality": {
			cstStaked: func(_ context.Context, _ int64, _ *cgstore.UserStakingTokenPageCursor, limit int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
				records := make([]cgstore.UserStakedCstTokenRecord, limit+1)
				for i := range records {
					records[i] = userStakedCstTokenAt(int64(i))
				}
				return records, false, nil
			},
		},
		"foreign staker": {
			cstStaked: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
				record := userStakedCstTokenAt(5)
				record.StakerAid = 77
				return []cgstore.UserStakedCstTokenRecord{record}, false, nil
			},
		},
		"unordered tokens": {
			cstStaked: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
				return []cgstore.UserStakedCstTokenRecord{
					userStakedCstTokenAt(9),
					userStakedCstTokenAt(5),
				}, false, nil
			},
		},
		"duplicate tokens": {
			cstStaked: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
				return []cgstore.UserStakedCstTokenRecord{
					userStakedCstTokenAt(5),
					userStakedCstTokenAt(5),
				}, false, nil
			},
		},
		"has more without rows": {
			cstStaked: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
				return []cgstore.UserStakedCstTokenRecord{}, true, nil
			},
		},
		"unmappable row": {
			cstStaked: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
				record := userStakedCstTokenAt(5)
				record.Seed = ""
				return []cgstore.UserStakedCstTokenRecord{record}, false, nil
			},
		},
	}
	for name, reader := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserStakingTestServer(t, reader),
				"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/staked-tokens")
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "secret db detail") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}

	t.Run("random-walk repository failure", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			rwalkStaked: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedRwalkTokenRecord, bool, error) {
				return nil, false, errors.New("secret db detail")
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/random-walk/staked-tokens")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "secret db detail") {
			t.Fatalf("internal detail leaked: %s", response.Body.String())
		}
	})
}

func userStakingDepositAt(depositID int64) cgstore.UserStakingDepositRecord {
	record := validUserStakingDepositRecord()
	record.DepositID = depositID
	return record
}

func TestListUserStakingDepositsPaginatesAndFilters(t *testing.T) {
	t.Parallel()

	var gotClaimed *bool
	var gotAfter *cgstore.UserStakingDepositPageCursor
	server := newUserStakingTestServer(t, fakeUserStakingReader{
		deposits: func(_ context.Context, _ int64, claimed *bool, after *cgstore.UserStakingDepositPageCursor, _ int) ([]cgstore.UserStakingDepositRecord, bool, error) {
			gotClaimed, gotAfter = claimed, after
			if after != nil {
				return []cgstore.UserStakingDepositRecord{}, false, nil
			}
			return []cgstore.UserStakingDepositRecord{
				userStakingDepositAt(502),
				userStakingDepositAt(501),
			}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits?claimed=true&limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotClaimed == nil || !*gotClaimed || gotAfter != nil {
		t.Fatalf("repository args = claimed=%v after=%+v", gotClaimed, gotAfter)
	}
	var page CosmicGameUserStakingDepositPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil ||
		!page.Data[0].FullyClaimed || page.Data[0].DepositId != 502 {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeUserStakingDepositCursor(*page.Meta.NextCursor, userCursorAlice)
	if err != nil || cursor.DepositID != 501 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}

	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits?cursor="+*page.Meta.NextCursor)
	if response.Code != http.StatusOK || gotAfter == nil || gotAfter.DepositID != 501 {
		t.Fatalf("continuation: status=%d cursor=%+v", response.Code, gotAfter)
	}
}

func TestListUserStakingDepositsFailsClosedOnViolations(t *testing.T) {
	t.Parallel()

	pendingDeposit := func(depositID int64) cgstore.UserStakingDepositRecord {
		record := userStakingDepositAt(depositID)
		record.StakedNftCount = 2
		record.AmountDepositedWei = "2000000000000000000"
		record.AmountToClaimWei = "1000000000000000000"
		record.PendingRewardWei = "1000000000000000000"
		record.PendingNftCount = 1
		return record
	}
	tests := map[string]struct {
		target string
		reader fakeUserStakingReader
	}{
		"filter contradiction": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits?claimed=true",
			reader: fakeUserStakingReader{
				deposits: func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error) {
					return []cgstore.UserStakingDepositRecord{pendingDeposit(501)}, false, nil
				},
			},
		},
		"address resolution failure": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits",
			reader: fakeUserStakingReader{
				addressID: func(context.Context, string) (int64, error) {
					return 0, errors.New("secret db detail")
				},
			},
		},
		"repository failure": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits",
			reader: fakeUserStakingReader{
				deposits: func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error) {
					return nil, false, errors.New("secret db detail")
				},
			},
		},
		"over cardinality": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits",
			reader: fakeUserStakingReader{
				deposits: func(_ context.Context, _ int64, _ *bool, _ *cgstore.UserStakingDepositPageCursor, limit int) ([]cgstore.UserStakingDepositRecord, bool, error) {
					records := make([]cgstore.UserStakingDepositRecord, limit+1)
					for i := range records {
						records[i] = userStakingDepositAt(int64(1000 - i))
					}
					return records, false, nil
				},
			},
		},
		"unmappable row": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits",
			reader: fakeUserStakingReader{
				deposits: func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error) {
					record := userStakingDepositAt(501)
					record.Tx.TxHash = "0xnope"
					return []cgstore.UserStakingDepositRecord{record}, false, nil
				},
			},
		},
		"foreign staker": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits",
			reader: fakeUserStakingReader{
				deposits: func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error) {
					record := userStakingDepositAt(501)
					record.StakerAid = 77
					return []cgstore.UserStakingDepositRecord{record}, false, nil
				},
			},
		},
		"unordered deposits": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits",
			reader: fakeUserStakingReader{
				deposits: func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error) {
					return []cgstore.UserStakingDepositRecord{
						userStakingDepositAt(501),
						userStakingDepositAt(502),
					}, false, nil
				},
			},
		},
		"has more without rows": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits",
			reader: fakeUserStakingReader{
				deposits: func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error) {
					return []cgstore.UserStakingDepositRecord{}, true, nil
				},
			},
		},
		"inconsistent accumulators": {
			target: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits",
			reader: fakeUserStakingReader{
				deposits: func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error) {
					record := userStakingDepositAt(501)
					record.ClaimedRewardWei = "1"
					return []cgstore.UserStakingDepositRecord{record}, false, nil
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserStakingTestServer(t, test.reader), test.target)
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "secret db detail") {
				t.Fatalf("internal detail leaked: %s", response.Body.String())
			}
		})
	}

	t.Run("unindexed wallet answers an empty filtered page", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits?claimed=false")
		if response.Code != http.StatusOK || !strings.Contains(response.Body.String(), `"data":[]`) {
			t.Fatalf("unindexed wallet = %d %s", response.Code, response.Body.String())
		}
	})
}

func TestListUserStakingDepositRewards(t *testing.T) {
	t.Parallel()

	reward := func(actionID int64) cgstore.UserStakingDepositRewardRecord {
		return cgstore.UserStakingDepositRewardRecord{
			StakerAid: 1,
			ActionID:  actionID,
			TokenID:   5,
			RewardWei: "1000000000000000000",
		}
	}

	t.Run("paginates within one deposit", func(t *testing.T) {
		t.Parallel()
		var gotDeposit int64
		var gotAfter *cgstore.UserStakingRewardPageCursor
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			depositReward: func(_ context.Context, _ int64, depositID int64, after *cgstore.UserStakingRewardPageCursor, _ int) ([]cgstore.UserStakingDepositRewardRecord, bool, error) {
				gotDeposit, gotAfter = depositID, after
				if after != nil {
					return []cgstore.UserStakingDepositRewardRecord{}, false, nil
				}
				return []cgstore.UserStakingDepositRewardRecord{reward(2), reward(7)}, true, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits/501/rewards?limit=2")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		if gotDeposit != 501 || gotAfter != nil {
			t.Fatalf("repository args = (%d, %+v)", gotDeposit, gotAfter)
		}
		var page CosmicGameUserStakingDepositRewardPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Meta.NextCursor == nil {
			t.Fatalf("page = %+v", page)
		}
		cursor, err := decodeUserStakingDepositRewardCursor(*page.Meta.NextCursor, userCursorAlice, 501)
		if err != nil || cursor.ActionID != 7 {
			t.Fatalf("next cursor = %+v, err=%v", cursor, err)
		}

		response = serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits/501/rewards?cursor="+*page.Meta.NextCursor)
		if response.Code != http.StatusOK || gotAfter == nil || gotAfter.ActionID != 7 {
			t.Fatalf("continuation: status=%d cursor=%+v", response.Code, gotAfter)
		}
		// The same cursor cannot continue another deposit's page.
		response = serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits/502/rewards?cursor="+*page.Meta.NextCursor)
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("missing deposit answers 404", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			depositExists: func(_ context.Context, depositID int64) (bool, error) {
				return depositID == 501, nil
			},
			depositReward: func(context.Context, int64, int64, *cgstore.UserStakingRewardPageCursor, int) ([]cgstore.UserStakingDepositRewardRecord, bool, error) {
				t.Error("reward page fetched for a missing deposit")
				return nil, false, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits/999/rewards")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("existence check failure stays opaque", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			depositExists: func(context.Context, int64) (bool, error) {
				return false, errors.New("secret db detail")
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits/501/rewards")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "secret db detail") {
			t.Fatalf("internal detail leaked: %s", response.Body.String())
		}
	})

	t.Run("unindexed wallet gets an empty page for an existing deposit", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			addressID: func(context.Context, string) (int64, error) {
				return 0, store.ErrNotFound
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits/501/rewards")
		if response.Code != http.StatusOK || !strings.Contains(response.Body.String(), `"data":[]`) {
			t.Fatalf("unindexed wallet = %d %s", response.Code, response.Body.String())
		}
	})

	t.Run("scope and order violations fail closed", func(t *testing.T) {
		t.Parallel()
		violations := map[string][]cgstore.UserStakingDepositRewardRecord{
			"foreign staker": {func() cgstore.UserStakingDepositRewardRecord {
				record := reward(2)
				record.StakerAid = 77
				return record
			}()},
			"unordered actions": {reward(7), reward(2)},
			"duplicate actions": {reward(2), reward(2)},
			"unmappable row": {func() cgstore.UserStakingDepositRewardRecord {
				record := reward(2)
				record.RewardWei = "not-a-number"
				return record
			}()},
		}
		for name, records := range violations {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				server := newUserStakingTestServer(t, fakeUserStakingReader{
					depositReward: func(context.Context, int64, int64, *cgstore.UserStakingRewardPageCursor, int) ([]cgstore.UserStakingDepositRewardRecord, bool, error) {
						return records, false, nil
					},
				})
				response := serve(t, server,
					"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits/501/rewards")
				assertProblem(t, response, http.StatusInternalServerError)
			})
		}
	})

	t.Run("repository contract violations fail closed", func(t *testing.T) {
		t.Parallel()
		readers := map[string]fakeUserStakingReader{
			"address resolution failure": {
				addressID: func(context.Context, string) (int64, error) {
					return 0, errors.New("secret db detail")
				},
			},
			"repository failure": {
				depositReward: func(context.Context, int64, int64, *cgstore.UserStakingRewardPageCursor, int) ([]cgstore.UserStakingDepositRewardRecord, bool, error) {
					return nil, false, errors.New("secret db detail")
				},
			},
			"over cardinality": {
				depositReward: func(_ context.Context, _ int64, _ int64, _ *cgstore.UserStakingRewardPageCursor, limit int) ([]cgstore.UserStakingDepositRewardRecord, bool, error) {
					records := make([]cgstore.UserStakingDepositRewardRecord, limit+1)
					for i := range records {
						records[i] = reward(int64(i + 1))
					}
					return records, false, nil
				},
			},
			"has more without rows": {
				depositReward: func(context.Context, int64, int64, *cgstore.UserStakingRewardPageCursor, int) ([]cgstore.UserStakingDepositRewardRecord, bool, error) {
					return []cgstore.UserStakingDepositRewardRecord{}, true, nil
				},
			},
		}
		for name, reader := range readers {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				response := serve(t, newUserStakingTestServer(t, reader),
					"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/deposits/501/rewards")
				assertProblem(t, response, http.StatusInternalServerError)
				if strings.Contains(response.Body.String(), "secret db detail") {
					t.Fatalf("internal detail leaked: %s", response.Body.String())
				}
			})
		}
	})
}

func TestListUserStakingTokenRewards(t *testing.T) {
	t.Parallel()

	tokenReward := func(tokenID int64) cgstore.UserStakingTokenRewardRecord {
		return cgstore.UserStakingTokenRewardRecord{
			TokenID:      tokenID,
			TotalWei:     "4000000000000000000",
			CollectedWei: "1000000000000000000",
			PendingWei:   "3000000000000000000",
		}
	}
	var gotAfter *cgstore.UserStakingTokenPageCursor
	server := newUserStakingTestServer(t, fakeUserStakingReader{
		tokenRewards: func(_ context.Context, _ int64, after *cgstore.UserStakingTokenPageCursor, _ int) ([]cgstore.UserStakingTokenRewardRecord, bool, error) {
			gotAfter = after
			if after != nil {
				return []cgstore.UserStakingTokenRewardRecord{}, false, nil
			}
			return []cgstore.UserStakingTokenRewardRecord{tokenReward(1), tokenReward(5)}, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var page CosmicGameUserStakingTokenRewardPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.NextCursor == nil ||
		page.Data[0].TotalWei != "4000000000000000000" {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeUserStakingTokenCursor(
		*page.Meta.NextCursor, userCursorAlice, userStakingTokenResourceCstTokenRewards)
	if err != nil || cursor.TokenID != 5 {
		t.Fatalf("next cursor = %+v, err=%v", cursor, err)
	}
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards?cursor="+*page.Meta.NextCursor)
	if response.Code != http.StatusOK || gotAfter == nil || gotAfter.TokenID != 5 {
		t.Fatalf("continuation: status=%d cursor=%+v", response.Code, gotAfter)
	}

	// Inconsistent totals fail closed.
	server = newUserStakingTestServer(t, fakeUserStakingReader{
		tokenRewards: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakingTokenRewardRecord, bool, error) {
			record := tokenReward(1)
			record.TotalWei = "5"
			return []cgstore.UserStakingTokenRewardRecord{record}, false, nil
		},
	})
	response = serve(t, server,
		"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards")
	assertProblem(t, response, http.StatusInternalServerError)
}

func TestListUserStakingTokenRewardDeposits(t *testing.T) {
	t.Parallel()

	tokenDeposit := func(depositID int64) cgstore.UserStakingTokenRewardDepositRecord {
		return cgstore.UserStakingTokenRewardDepositRecord{
			Tx:        validDonationTransaction(),
			StakerAid: 1,
			DepositID: depositID,
			RoundNum:  0,
			RewardWei: "1000000000000000000",
			Claimed:   true,
		}
	}

	t.Run("paginates within one token", func(t *testing.T) {
		t.Parallel()
		var gotToken int64
		var gotAfter *cgstore.UserStakingTokenDepositPageCursor
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			tokenDeposits: func(_ context.Context, _ int64, tokenID int64, after *cgstore.UserStakingTokenDepositPageCursor, _ int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
				gotToken, gotAfter = tokenID, after
				if after != nil {
					return []cgstore.UserStakingTokenRewardDepositRecord{}, false, nil
				}
				return []cgstore.UserStakingTokenRewardDepositRecord{
					tokenDeposit(501),
					tokenDeposit(502),
				}, true, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards/5/deposits?limit=2")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		if gotToken != 5 || gotAfter != nil {
			t.Fatalf("repository args = (%d, %+v)", gotToken, gotAfter)
		}
		var page CosmicGameUserStakingTokenRewardDepositPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 2 || page.Meta.NextCursor == nil {
			t.Fatalf("page = %+v", page)
		}
		cursor, err := decodeUserStakingTokenDepositCursor(*page.Meta.NextCursor, userCursorAlice, 5)
		if err != nil || cursor.DepositID != 502 {
			t.Fatalf("next cursor = %+v, err=%v", cursor, err)
		}
		response = serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards/5/deposits?cursor="+*page.Meta.NextCursor)
		if response.Code != http.StatusOK || gotAfter == nil || gotAfter.DepositID != 502 {
			t.Fatalf("continuation: status=%d cursor=%+v", response.Code, gotAfter)
		}
		// The same cursor cannot continue another token's page.
		response = serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards/9/deposits?cursor="+*page.Meta.NextCursor)
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("unminted token answers 404", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			tokenExists: func(_ context.Context, tokenID int64) (bool, error) {
				return tokenID == 5, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards/999/deposits")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("unordered deposits fail closed", func(t *testing.T) {
		t.Parallel()
		server := newUserStakingTestServer(t, fakeUserStakingReader{
			tokenDeposits: func(context.Context, int64, int64, *cgstore.UserStakingTokenDepositPageCursor, int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
				return []cgstore.UserStakingTokenRewardDepositRecord{
					tokenDeposit(502),
					tokenDeposit(501),
				}, false, nil
			},
		})
		response := serve(t, server,
			"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards/5/deposits")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("repository contract violations fail closed", func(t *testing.T) {
		t.Parallel()
		readers := map[string]fakeUserStakingReader{
			"existence check failure": {
				tokenExists: func(context.Context, int64) (bool, error) {
					return false, errors.New("secret db detail")
				},
			},
			"address resolution failure": {
				addressID: func(context.Context, string) (int64, error) {
					return 0, errors.New("secret db detail")
				},
			},
			"repository failure": {
				tokenDeposits: func(context.Context, int64, int64, *cgstore.UserStakingTokenDepositPageCursor, int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
					return nil, false, errors.New("secret db detail")
				},
			},
			"over cardinality": {
				tokenDeposits: func(_ context.Context, _ int64, _ int64, _ *cgstore.UserStakingTokenDepositPageCursor, limit int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
					records := make([]cgstore.UserStakingTokenRewardDepositRecord, limit+1)
					for i := range records {
						records[i] = tokenDeposit(int64(i + 1))
					}
					return records, false, nil
				},
			},
			"has more without rows": {
				tokenDeposits: func(context.Context, int64, int64, *cgstore.UserStakingTokenDepositPageCursor, int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
					return []cgstore.UserStakingTokenRewardDepositRecord{}, true, nil
				},
			},
			"foreign staker": {
				tokenDeposits: func(context.Context, int64, int64, *cgstore.UserStakingTokenDepositPageCursor, int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
					record := tokenDeposit(501)
					record.StakerAid = 77
					return []cgstore.UserStakingTokenRewardDepositRecord{record}, false, nil
				},
			},
			"unmappable row": {
				tokenDeposits: func(context.Context, int64, int64, *cgstore.UserStakingTokenDepositPageCursor, int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
					record := tokenDeposit(501)
					record.RewardWei = "not-a-number"
					return []cgstore.UserStakingTokenRewardDepositRecord{record}, false, nil
				},
			},
		}
		for name, reader := range readers {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				response := serve(t, newUserStakingTestServer(t, reader),
					"/api/v2/cosmicgame/users/"+userCursorAlice+"/staking/cst/token-rewards/5/deposits")
				assertProblem(t, response, http.StatusInternalServerError)
				if strings.Contains(response.Body.String(), "secret db detail") {
					t.Fatalf("internal detail leaked: %s", response.Body.String())
				}
			})
		}
	})
}

func TestUserStakingResourcesRejectInvalidInput(t *testing.T) {
	t.Parallel()

	base := "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking"
	// #nosec G101 -- request paths under test, not credentials.
	targets := map[string]string{
		"actions invalid address":        "/api/v2/cosmicgame/users/not-an-address/staking/cst/actions",
		"actions invalid limit":          base + "/cst/actions?limit=0",
		"actions bind limit":             base + "/cst/actions?limit=wat",
		"actions malformed cursor":       base + "/cst/actions?cursor=bad",
		"rwalk actions invalid address":  "/api/v2/cosmicgame/users/nope/staking/random-walk/actions",
		"staked invalid limit":           base + "/cst/staked-tokens?limit=201",
		"staked malformed cursor":        base + "/cst/staked-tokens?cursor=bad",
		"rwalk staked invalid address":   "/api/v2/cosmicgame/users/nope/staking/random-walk/staked-tokens",
		"deposits invalid address":       "/api/v2/cosmicgame/users/nope/staking/cst/deposits",
		"deposits invalid limit":         base + "/cst/deposits?limit=201",
		"deposits malformed cursor":      base + "/cst/deposits?cursor=bad",
		"deposits duplicate filter":      base + "/cst/deposits?claimed=true&claimed=false",
		"deposit rewards bind id":        base + "/cst/deposits/wat/rewards",
		"deposit rewards invalid limit":  base + "/cst/deposits/501/rewards?limit=0",
		"deposit rewards bad cursor":     base + "/cst/deposits/501/rewards?cursor=bad",
		"deposit rewards bad address":    "/api/v2/cosmicgame/users/nope/staking/cst/deposits/501/rewards",
		"token rewards malformed cursor": base + "/cst/token-rewards?cursor=bad",
		"token deposits bind id":         base + "/cst/token-rewards/wat/deposits",
		"token deposits invalid limit":   base + "/cst/token-rewards/5/deposits?limit=0",
		"token deposits bad cursor":      base + "/cst/token-rewards/5/deposits?cursor=bad",
		"token deposits invalid address": "/api/v2/cosmicgame/users/nope/staking/cst/token-rewards/5/deposits",
	}
	for name, target := range targets {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserStakingTestServer(t, fakeUserStakingReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestUserStakingResourcesAnswerEmptyPagesForUnindexedWallets(t *testing.T) {
	t.Parallel()

	unindexed := fakeUserStakingReader{
		addressID: func(context.Context, string) (int64, error) {
			return 0, store.ErrNotFound
		},
	}
	base := "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking"
	targets := []string{
		base + "/cst/actions",
		base + "/random-walk/actions",
		base + "/cst/staked-tokens",
		base + "/random-walk/staked-tokens",
		base + "/cst/deposits",
		base + "/cst/deposits/501/rewards",
		base + "/cst/token-rewards",
		base + "/cst/token-rewards/5/deposits",
	}
	for _, target := range targets {
		t.Run(target, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserStakingTestServer(t, unindexed), target)
			if response.Code != http.StatusOK || !strings.Contains(response.Body.String(), `"data":[]`) {
				t.Fatalf("unindexed wallet = %d %s", response.Code, response.Body.String())
			}
		})
	}
}

func TestUserStakingBindingEdgesStaySecretFree(t *testing.T) {
	t.Parallel()

	// Every generated binding arm of the new operations answers the same
	// stable RFC 9457 problem and never echoes the offending value.
	base := "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking"
	endpoints := []string{
		base + "/cst/actions",
		base + "/cst/staked-tokens",
		base + "/cst/deposits",
		base + "/cst/deposits/501/rewards",
		base + "/cst/token-rewards",
		base + "/cst/token-rewards/5/deposits",
		base + "/random-walk/actions",
		base + "/random-walk/staked-tokens",
	}
	edges := map[string]string{
		"duplicate cursor": "?cursor=password-super-secret&cursor=b",
		"duplicate limit":  "?limit=1&limit=2",
		"malformed limit":  "?limit=password-super-secret",
	}
	for _, endpoint := range endpoints {
		for name, query := range edges {
			t.Run(name+" "+endpoint, func(t *testing.T) {
				t.Parallel()
				response := serve(t, newUserStakingTestServer(t, fakeUserStakingReader{}), endpoint+query)
				assertProblem(t, response, http.StatusBadRequest)
				var problem Problem
				decodeResponse(t, response, &problem)
				if problem.Detail == nil || !strings.Contains(*problem.Detail, `Parameter "`) {
					t.Fatalf("problem = %+v", problem)
				}
				if strings.Contains(response.Body.String(), "password-super-secret") {
					t.Fatalf("parameter value leaked: %s", response.Body.String())
				}
			})
		}
	}

	t.Run("malformed claimed filter", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newUserStakingTestServer(t, fakeUserStakingReader{}),
			base+"/cst/deposits?claimed=password-super-secret")
		assertProblem(t, response, http.StatusBadRequest)
		if strings.Contains(response.Body.String(), "password-super-secret") {
			t.Fatalf("parameter value leaked: %s", response.Body.String())
		}
	})
}

func TestUserStakingNegativePathIdentifiersRejected(t *testing.T) {
	t.Parallel()

	// The OpenAPI minimum cannot be enforced by the stdlib binder, so the
	// handlers reject negative path identifiers themselves.
	base := "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking"
	for _, target := range []string{
		base + "/cst/deposits/-1/rewards",
		base + "/cst/token-rewards/-1/deposits",
	} {
		t.Run(target, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newUserStakingTestServer(t, fakeUserStakingReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}
