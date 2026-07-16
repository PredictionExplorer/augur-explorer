package v2

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func newGlobalStakingTestServer(t *testing.T, staking globalStakingReader) *Server {
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
		fakeGlobalDirectoryReader{},
		staking,
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func validCstGlobalAction() cgstore.GlobalStakingActionRecord {
	return validGlobalStakingAction(cgstore.UserStakingActionStake)
}

func validRwalkGlobalAction(kind cgstore.UserStakingActionKind) cgstore.GlobalStakingActionRecord {
	record := validGlobalStakingAction(kind)
	record.RewardWei = ""
	record.RewardPerTokenWei = ""
	return record
}

func TestGlobalStakingActionPages(t *testing.T) {
	t.Parallel()
	t.Run("CST page and continuation", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			cstActions: func(
				_ context.Context,
				after *cgstore.GlobalStakingActionPageCursor,
				limit int,
			) ([]cgstore.GlobalStakingActionRecord, bool, error) {
				if after != nil || limit != 1 {
					t.Fatalf("after=%+v limit=%d", after, limit)
				}
				return []cgstore.GlobalStakingActionRecord{
					validCstGlobalAction(),
				}, true, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader),
			globalCstActionsInstance+"?limit=1")
		if response.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", response.Code, response.Body)
		}
		var page CosmicGameGlobalCstStakingActionPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 1 || page.Meta.NextCursor == nil {
			t.Fatalf("page=%+v", page)
		}
		cursor, err := decodeGlobalStakingEventCursor(
			*page.Meta.NextCursor,
			globalStakingEventCstActions,
		)
		if err != nil || cursor.EventLogID != page.Data[0].EventLogId {
			t.Fatalf("cursor=%+v err=%v", cursor, err)
		}
	})

	t.Run("RandomWalk page", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			rwalkActions: func(
				_ context.Context,
				_ *cgstore.GlobalStakingActionPageCursor,
				_ int,
			) ([]cgstore.GlobalStakingActionRecord, bool, error) {
				return []cgstore.GlobalStakingActionRecord{
					validRwalkGlobalAction(cgstore.UserStakingActionUnstake),
				}, false, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalRwalkActionsInstance)
		if response.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", response.Code, response.Body)
		}
		var page CosmicGameGlobalRandomWalkStakingActionPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 1 || page.Data[0].ActionType != Unstake {
			t.Fatalf("page=%+v", page)
		}
	})

	cross, err := encodeGlobalStakingEventCursor(globalStakingEventCursor{
		Version:    globalStakingEventCursorVersion,
		Resource:   globalStakingEventCstActions,
		EventLogID: 100,
	})
	if err != nil {
		t.Fatal(err)
	}
	for name, target := range map[string]string{
		"cross-resource cursor": globalRwalkActionsInstance + "?cursor=" + cross,
		"malformed cursor":      globalCstActionsInstance + "?cursor=nope",
		"invalid limit":         globalCstActionsInstance + "?limit=0",
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, fakeGlobalStakingReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}

	errorCases := map[string]fakeGlobalStakingReader{
		"store error": {
			cstActions: func(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error) {
				return nil, false, errors.New("secret")
			},
		},
		"over cardinality": {
			cstActions: func(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error) {
				record := validCstGlobalAction()
				record2 := record
				record2.Tx.EvtLogId--
				return []cgstore.GlobalStakingActionRecord{record, record2}, false, nil
			},
		},
		"mapping error": {
			cstActions: func(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error) {
				record := validCstGlobalAction()
				record.StakerAddress = "secret-invalid-address"
				return []cgstore.GlobalStakingActionRecord{record}, false, nil
			},
		},
		"has more without row": {
			cstActions: func(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error) {
				return []cgstore.GlobalStakingActionRecord{}, true, nil
			},
		},
	}
	for name, reader := range errorCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, reader),
				globalCstActionsInstance+"?limit=1")
			assertOpaqueInternalProblem(t, response)
		})
	}

	t.Run("unordered continuation", func(t *testing.T) {
		t.Parallel()
		cursor, err := encodeGlobalStakingEventCursor(globalStakingEventCursor{
			Version:    globalStakingEventCursorVersion,
			Resource:   globalStakingEventCstActions,
			EventLogID: 100,
		})
		if err != nil {
			t.Fatal(err)
		}
		reader := fakeGlobalStakingReader{
			cstActions: func(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error) {
				record := validCstGlobalAction()
				record.Tx.EvtLogId = 100
				return []cgstore.GlobalStakingActionRecord{record}, false, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader),
			globalCstActionsInstance+"?cursor="+cursor)
		assertOpaqueInternalProblem(t, response)
	})
}

func TestGlobalStakingActionDetails(t *testing.T) {
	t.Parallel()
	t.Run("closed CST lifecycle", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			cstAction: func(_ context.Context, actionID int64) (cgmodel.CGStakeUnstakeCombined, error) {
				if actionID != 7 {
					t.Fatalf("action=%d", actionID)
				}
				return validStakeUnstakeCombined(true, true), nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalCstActionsInstance+"/7")
		if response.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", response.Code, response.Body)
		}
		var detail GlobalCstStakingActionDetail
		decodeResponse(t, response, &detail)
		if detail.Unstake == nil || detail.Unstake.RewardWei == nil {
			t.Fatalf("detail=%+v", detail)
		}
	})

	t.Run("open RandomWalk lifecycle", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			rwalkAction: func(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error) {
				return validStakeUnstakeCombined(false, false), nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalRwalkActionsInstance+"/7")
		if response.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", response.Code, response.Body)
		}
		var detail GlobalRandomWalkStakingActionDetail
		decodeResponse(t, response, &detail)
		if detail.Unstake != nil {
			t.Fatalf("detail=%+v", detail)
		}
	})

	for name, target := range map[string]string{
		"CST missing":         globalCstActionsInstance + "/999",
		"RandomWalk missing":  globalRwalkActionsInstance + "/999",
		"CST negative":        globalCstActionsInstance + "/-1",
		"RandomWalk negative": globalRwalkActionsInstance + "/-1",
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, fakeGlobalStakingReader{}), target)
			status := http.StatusNotFound
			if target[len(target)-2:] == "-1" {
				status = http.StatusBadRequest
			}
			assertProblem(t, response, status)
		})
	}

	t.Run("store error is opaque", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			cstAction: func(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error) {
				return cgmodel.CGStakeUnstakeCombined{}, errors.New("secret")
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalCstActionsInstance+"/7")
		assertOpaqueInternalProblem(t, response)
	})

	t.Run("CST mapping error is opaque", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			cstAction: func(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error) {
				record := validStakeUnstakeCombined(false, true)
				record.Stake.StakerAddr = "secret-invalid-address"
				return record, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalCstActionsInstance+"/7")
		assertOpaqueInternalProblem(t, response)
	})

	t.Run("RandomWalk store error is opaque", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			rwalkAction: func(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error) {
				return cgmodel.CGStakeUnstakeCombined{}, errors.New("secret")
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalRwalkActionsInstance+"/7")
		assertOpaqueInternalProblem(t, response)
	})

	t.Run("mapping error is opaque", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			rwalkAction: func(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error) {
				record := validStakeUnstakeCombined(false, false)
				record.Stake.StakerAddr = "secret-invalid-address"
				return record, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalRwalkActionsInstance+"/7")
		assertOpaqueInternalProblem(t, response)
	})
}

func validGlobalStakedCstRecord(tokenID int64) cgstore.GlobalStakedCstTokenRecord {
	return cgstore.GlobalStakedCstTokenRecord{
		StakeTx:       validDonationTransaction(),
		StakerAid:     21,
		StakerAddress: userCursorAlice,
		ActionID:      7,
		TokenID:       tokenID,
		MintRound:     2,
		Seed:          "seed",
	}
}

func validGlobalStakedRwalkRecord(tokenID int64) cgstore.GlobalStakedRwalkTokenRecord {
	return cgstore.GlobalStakedRwalkTokenRecord{
		StakeTx:       validDonationTransaction(),
		StakerAid:     21,
		StakerAddress: userCursorAlice,
		ActionID:      7,
		TokenID:       tokenID,
	}
}

func TestGlobalStakedTokenPages(t *testing.T) {
	t.Parallel()
	t.Run("CST page", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			cstStaked: func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedCstTokenRecord, bool, error) {
				return []cgstore.GlobalStakedCstTokenRecord{validGlobalStakedCstRecord(5)}, true, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader),
			globalCstStakedInstance+"?limit=1")
		if response.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", response.Code, response.Body)
		}
		var page CosmicGameGlobalStakedCstTokenPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 1 || page.Meta.NextCursor == nil {
			t.Fatalf("page=%+v", page)
		}
	})
	t.Run("RandomWalk page", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			rwalkStaked: func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedRwalkTokenRecord, bool, error) {
				return []cgstore.GlobalStakedRwalkTokenRecord{validGlobalStakedRwalkRecord(5)}, false, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalRwalkStakedInstance)
		if response.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", response.Code, response.Body)
		}
	})

	cross, err := encodeGlobalStakedTokenCursor(globalStakedTokenCursor{
		Version:  globalStakedTokenCursorVersion,
		Resource: globalStakedTokenCst,
		TokenID:  5,
	})
	if err != nil {
		t.Fatal(err)
	}
	for name, target := range map[string]string{
		"cross resource": globalRwalkStakedInstance + "?cursor=" + cross,
		"bad cursor":     globalCstStakedInstance + "?cursor=nope",
		"bad limit":      globalCstStakedInstance + "?limit=201",
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, fakeGlobalStakingReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}

	for name, reader := range map[string]fakeGlobalStakingReader{
		"store error": {
			cstStaked: func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedCstTokenRecord, bool, error) {
				return nil, false, errors.New("secret")
			},
		},
		"over cardinality": {
			cstStaked: func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedCstTokenRecord, bool, error) {
				return []cgstore.GlobalStakedCstTokenRecord{
					validGlobalStakedCstRecord(1),
					validGlobalStakedCstRecord(2),
				}, false, nil
			},
		},
		"mapping error": {
			cstStaked: func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedCstTokenRecord, bool, error) {
				record := validGlobalStakedCstRecord(1)
				record.Seed = ""
				return []cgstore.GlobalStakedCstTokenRecord{record}, false, nil
			},
		},
		"has more without row": {
			cstStaked: func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedCstTokenRecord, bool, error) {
				return []cgstore.GlobalStakedCstTokenRecord{}, true, nil
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, reader),
				globalCstStakedInstance+"?limit=1")
			assertOpaqueInternalProblem(t, response)
		})
	}

	t.Run("unordered continuation", func(t *testing.T) {
		t.Parallel()
		cursor, err := encodeGlobalStakedTokenCursor(globalStakedTokenCursor{
			Version:  globalStakedTokenCursorVersion,
			Resource: globalStakedTokenCst,
			TokenID:  5,
		})
		if err != nil {
			t.Fatal(err)
		}
		reader := fakeGlobalStakingReader{
			cstStaked: func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedCstTokenRecord, bool, error) {
				return []cgstore.GlobalStakedCstTokenRecord{validGlobalStakedCstRecord(5)}, false, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader),
			globalCstStakedInstance+"?cursor="+cursor)
		assertOpaqueInternalProblem(t, response)
	})

	t.Run("RandomWalk store error is opaque", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			rwalkStaked: func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedRwalkTokenRecord, bool, error) {
				return nil, false, errors.New("secret")
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), globalRwalkStakedInstance)
		assertOpaqueInternalProblem(t, response)
	})
}

func TestGlobalStakingDepositPage(t *testing.T) {
	t.Parallel()
	t.Run("happy path and continuation", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			deposits: func(context.Context, *cgstore.GlobalStakingDepositPageCursor, int) ([]cgstore.GlobalStakingDepositRecord, bool, error) {
				return []cgstore.GlobalStakingDepositRecord{validGlobalStakingDepositRecord()}, true, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader),
			globalStakingDepositsInstance+"?limit=1")
		if response.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", response.Code, response.Body)
		}
		var page CosmicGameGlobalStakingDepositPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 1 || page.Meta.NextCursor == nil {
			t.Fatalf("page=%+v", page)
		}
	})
	for name, target := range map[string]string{
		"bad cursor": globalStakingDepositsInstance + "?cursor=nope",
		"bad limit":  globalStakingDepositsInstance + "?limit=0",
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, fakeGlobalStakingReader{}), target)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
	for name, reader := range map[string]fakeGlobalStakingReader{
		"store error": {
			deposits: func(context.Context, *cgstore.GlobalStakingDepositPageCursor, int) ([]cgstore.GlobalStakingDepositRecord, bool, error) {
				return nil, false, errors.New("secret")
			},
		},
		"over cardinality": {
			deposits: func(context.Context, *cgstore.GlobalStakingDepositPageCursor, int) ([]cgstore.GlobalStakingDepositRecord, bool, error) {
				record := validGlobalStakingDepositRecord()
				record2 := record
				record2.DepositID--
				return []cgstore.GlobalStakingDepositRecord{record, record2}, false, nil
			},
		},
		"mapping error": {
			deposits: func(context.Context, *cgstore.GlobalStakingDepositPageCursor, int) ([]cgstore.GlobalStakingDepositRecord, bool, error) {
				record := validGlobalStakingDepositRecord()
				record.TotalDepositWei = "secret"
				return []cgstore.GlobalStakingDepositRecord{record}, false, nil
			},
		},
		"has more without row": {
			deposits: func(context.Context, *cgstore.GlobalStakingDepositPageCursor, int) ([]cgstore.GlobalStakingDepositRecord, bool, error) {
				return []cgstore.GlobalStakingDepositRecord{}, true, nil
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, reader),
				globalStakingDepositsInstance+"?limit=1")
			assertOpaqueInternalProblem(t, response)
		})
	}
	t.Run("unordered continuation", func(t *testing.T) {
		t.Parallel()
		cursor, err := encodeGlobalStakingDepositCursor(globalStakingDepositCursor{
			Version:   globalStakingDepositCursorVersion,
			DepositID: 8,
		})
		if err != nil {
			t.Fatal(err)
		}
		reader := fakeGlobalStakingReader{
			deposits: func(context.Context, *cgstore.GlobalStakingDepositPageCursor, int) ([]cgstore.GlobalStakingDepositRecord, bool, error) {
				return []cgstore.GlobalStakingDepositRecord{validGlobalStakingDepositRecord()}, false, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader),
			globalStakingDepositsInstance+"?cursor="+cursor)
		assertOpaqueInternalProblem(t, response)
	})
}

func validRoundStakingRewardRecord() cgstore.RoundStakingRewardRecord {
	return cgstore.RoundStakingRewardRecord{
		DepositID:      8,
		RoundNum:       2,
		StakerAid:      21,
		StakerAddress:  userCursorAlice,
		StakedNftCount: 2,
		RewardWei:      "100",
		CollectedWei:   "60",
		PendingWei:     "40",
	}
}

func TestRoundStakingRewardPage(t *testing.T) {
	t.Parallel()
	instance := "/api/v2/cosmicgame/rounds/2/staking-rewards"
	t.Run("happy path and continuation", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			roundRewards: func(
				_ context.Context,
				round int64,
				after *cgstore.RoundStakingRewardPageCursor,
				limit int,
			) ([]cgstore.RoundStakingRewardRecord, bool, error) {
				if round != 2 || after != nil || limit != 1 {
					t.Fatalf("round=%d after=%+v limit=%d", round, after, limit)
				}
				return []cgstore.RoundStakingRewardRecord{validRoundStakingRewardRecord()}, true, nil
			},
		}
		response := serve(t, newGlobalStakingTestServer(t, reader), instance+"?limit=1")
		if response.Code != http.StatusOK {
			t.Fatalf("status=%d body=%s", response.Code, response.Body)
		}
		var page CosmicGameRoundStakingRewardPage
		decodeResponse(t, response, &page)
		if len(page.Data) != 1 || page.Meta.NextCursor == nil {
			t.Fatalf("page=%+v", page)
		}
	})
	t.Run("missing round", func(t *testing.T) {
		t.Parallel()
		reader := fakeGlobalStakingReader{
			roundExists: func(context.Context, int64) (bool, error) { return false, nil },
		}
		assertProblem(t, serve(t, newGlobalStakingTestServer(t, reader), instance), http.StatusNotFound)
	})
	for name, target := range map[string]string{
		"negative round": "/api/v2/cosmicgame/rounds/-1/staking-rewards",
		"bad cursor":     instance + "?cursor=nope",
		"bad limit":      instance + "?limit=201",
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assertProblem(t,
				serve(t, newGlobalStakingTestServer(t, fakeGlobalStakingReader{}), target),
				http.StatusBadRequest)
		})
	}
	cross, err := encodeRoundStakingRewardCursor(roundStakingRewardCursor{
		Version:   roundStakingRewardCursorVersion,
		Round:     3,
		DepositID: 8,
		StakerAid: 21,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Run("cross-round cursor", func(t *testing.T) {
		t.Parallel()
		assertProblem(t,
			serve(t, newGlobalStakingTestServer(t, fakeGlobalStakingReader{}),
				instance+"?cursor="+cross),
			http.StatusBadRequest)
	})

	for name, reader := range map[string]fakeGlobalStakingReader{
		"existence error": {
			roundExists: func(context.Context, int64) (bool, error) {
				return false, errors.New("secret")
			},
		},
		"store error": {
			roundRewards: func(context.Context, int64, *cgstore.RoundStakingRewardPageCursor, int) ([]cgstore.RoundStakingRewardRecord, bool, error) {
				return nil, false, errors.New("secret")
			},
		},
		"over cardinality": {
			roundRewards: func(context.Context, int64, *cgstore.RoundStakingRewardPageCursor, int) ([]cgstore.RoundStakingRewardRecord, bool, error) {
				record := validRoundStakingRewardRecord()
				record2 := record
				record2.StakerAid++
				return []cgstore.RoundStakingRewardRecord{record, record2}, false, nil
			},
		},
		"out of scope": {
			roundRewards: func(context.Context, int64, *cgstore.RoundStakingRewardPageCursor, int) ([]cgstore.RoundStakingRewardRecord, bool, error) {
				record := validRoundStakingRewardRecord()
				record.RoundNum = 3
				return []cgstore.RoundStakingRewardRecord{record}, false, nil
			},
		},
		"mapping error": {
			roundRewards: func(context.Context, int64, *cgstore.RoundStakingRewardPageCursor, int) ([]cgstore.RoundStakingRewardRecord, bool, error) {
				record := validRoundStakingRewardRecord()
				record.RewardWei = "secret"
				return []cgstore.RoundStakingRewardRecord{record}, false, nil
			},
		},
		"has more without row": {
			roundRewards: func(context.Context, int64, *cgstore.RoundStakingRewardPageCursor, int) ([]cgstore.RoundStakingRewardRecord, bool, error) {
				return []cgstore.RoundStakingRewardRecord{}, true, nil
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, reader), instance+"?limit=1")
			assertOpaqueInternalProblem(t, response)
		})
	}
}

func validGlobalStakerRaffleRecord(isRwalk bool) cgmodel.CGRaffleNFTWinnerRec {
	return cgmodel.CGRaffleNFTWinnerRec{
		Tx:          validDonationTransaction(),
		WinnerAddr:  userCursorAlice,
		WinnerAid:   21,
		RoundNum:    2,
		TokenId:     5,
		CstAmount:   "100",
		WinnerIndex: 0,
		IsRWalk:     isRwalk,
		IsStaker:    true,
	}
}

func TestGlobalStakerRaffleNftWinPage(t *testing.T) {
	t.Parallel()
	for _, pool := range []struct {
		value   string
		isRwalk bool
	}{
		{value: "cst"},
		{value: "randomWalk", isRwalk: true},
	} {
		t.Run(pool.value, func(t *testing.T) {
			t.Parallel()
			reader := fakeGlobalStakingReader{
				raffleWins: func(
					_ context.Context,
					isRwalk bool,
					after *cgstore.GlobalStakerRafflePageCursor,
					limit int,
				) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
					if isRwalk != pool.isRwalk || after != nil || limit != 1 {
						t.Fatalf("isRwalk=%v after=%+v limit=%d", isRwalk, after, limit)
					}
					return []cgmodel.CGRaffleNFTWinnerRec{
						validGlobalStakerRaffleRecord(isRwalk),
					}, true, nil
				},
			}
			response := serve(t, newGlobalStakingTestServer(t, reader),
				globalStakerRaffleInstance+"?pool="+pool.value+"&limit=1")
			if response.Code != http.StatusOK {
				t.Fatalf("status=%d body=%s", response.Code, response.Body)
			}
			var page CosmicGameGlobalStakerRaffleNftWinPage
			decodeResponse(t, response, &page)
			if len(page.Data) != 1 || page.Meta.NextCursor == nil ||
				page.Data[0].IsRandomWalk != pool.isRwalk {
				t.Fatalf("page=%+v", page)
			}
		})
	}
	cross, err := encodeGlobalStakingEventCursor(globalStakingEventCursor{
		Version:    globalStakingEventCursorVersion,
		Resource:   globalStakingEventCstRaffle,
		EventLogID: 100,
	})
	if err != nil {
		t.Fatal(err)
	}
	for name, target := range map[string]string{
		"missing pool":      globalStakerRaffleInstance,
		"unknown pool":      globalStakerRaffleInstance + "?pool=unknown",
		"cross-pool cursor": globalStakerRaffleInstance + "?pool=randomWalk&cursor=" + cross,
		"malformed cursor":  globalStakerRaffleInstance + "?pool=cst&cursor=nope",
		"invalid limit":     globalStakerRaffleInstance + "?pool=cst&limit=0",
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assertProblem(t,
				serve(t, newGlobalStakingTestServer(t, fakeGlobalStakingReader{}), target),
				http.StatusBadRequest)
		})
	}
	for name, reader := range map[string]fakeGlobalStakingReader{
		"store error": {
			raffleWins: func(context.Context, bool, *cgstore.GlobalStakerRafflePageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
				return nil, false, errors.New("secret")
			},
		},
		"over cardinality": {
			raffleWins: func(context.Context, bool, *cgstore.GlobalStakerRafflePageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
				record := validGlobalStakerRaffleRecord(false)
				record2 := record
				record2.Tx.EvtLogId--
				return []cgmodel.CGRaffleNFTWinnerRec{record, record2}, false, nil
			},
		},
		"wrong pool": {
			raffleWins: func(context.Context, bool, *cgstore.GlobalStakerRafflePageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
				return []cgmodel.CGRaffleNFTWinnerRec{
					validGlobalStakerRaffleRecord(true),
				}, false, nil
			},
		},
		"mapping error": {
			raffleWins: func(context.Context, bool, *cgstore.GlobalStakerRafflePageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
				record := validGlobalStakerRaffleRecord(false)
				record.CstAmount = "secret"
				return []cgmodel.CGRaffleNFTWinnerRec{record}, false, nil
			},
		},
		"has more without row": {
			raffleWins: func(context.Context, bool, *cgstore.GlobalStakerRafflePageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
				return []cgmodel.CGRaffleNFTWinnerRec{}, true, nil
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newGlobalStakingTestServer(t, reader),
				globalStakerRaffleInstance+"?pool=cst&limit=1")
			assertOpaqueInternalProblem(t, response)
		})
	}
}

func assertOpaqueInternalProblem(t *testing.T, response *httptest.ResponseRecorder) {
	t.Helper()
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("status=%d, want 500", response.Code)
	}
	var problem Problem
	if err := json.Unmarshal(response.Body.Bytes(), &problem); err != nil {
		t.Fatal(err)
	}
	if problem.Type != problemTypeBase+"internal" ||
		(problem.Detail != nil && (*problem.Detail == "secret" ||
			*problem.Detail == "secret-invalid-address")) {
		t.Fatalf("problem=%+v", problem)
	}
}
