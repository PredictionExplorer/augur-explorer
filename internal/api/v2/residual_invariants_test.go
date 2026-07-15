package v2

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestPaginatedHandlersRejectOversizedRepositoryPages(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		path      string
		configure func(*Server)
	}{
		{
			name: "round bids",
			path: "/api/v2/cosmicgame/rounds/1/bids?limit=1",
			configure: func(server *Server) {
				first := validBidRecord()
				first.RoundNum, first.BidPosition, first.Tx.EvtLogId = 1, 1, 10
				second := validBidRecord()
				second.RoundNum, second.BidPosition, second.Tx.EvtLogId = 1, 2, 11
				server.bids = fakeBidReader{
					page: func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgmodel.CGBidRec, bool, error) {
						return []cgmodel.CGBidRec{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "completed rounds",
			path: "/api/v2/cosmicgame/rounds?limit=1",
			configure: func(server *Server) {
				first := validRoundRecord()
				first.RoundNum, first.ClaimPrizeTx.Tx.EvtLogId = 2, 20
				second := validRoundRecord()
				second.RoundNum, second.ClaimPrizeTx.Tx.EvtLogId = 1, 10
				server.rounds = fakeRoundReader{
					page: func(context.Context, *cgstore.RoundPageCursor, int) ([]cgmodel.CGRoundRec, bool, error) {
						return []cgmodel.CGRoundRec{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "round prizes",
			path: "/api/v2/cosmicgame/rounds/1/prizes?limit=1",
			configure: func(server *Server) {
				first := validRoundPrizeRecord(0)
				first.RoundNum, first.WinnerIndex = 1, 0
				second := validRoundPrizeRecord(1)
				second.RoundNum, second.WinnerIndex = 1, 0
				server.prizes = fakeRoundPrizeReader{
					page: func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
						return []cgmodel.CGPrizeHistory{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "raffle ETH deposits",
			path: "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits?limit=1",
			configure: func(server *Server) {
				first := validRaffleEthDepositRecord()
				first.RoundNum, first.WinnerIndex, first.Tx.EvtLogId = 1, 0, 10
				second := validRaffleEthDepositRecord()
				second.RoundNum, second.WinnerIndex, second.Tx.EvtLogId = 1, 1, 11
				server.raffles = fakeRoundRaffleReader{
					eth: func(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error) {
						return []cgstore.RaffleEthDepositRecord{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "raffle NFT winners",
			path: "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder&limit=1",
			configure: func(server *Server) {
				first := validRaffleNftWinnerRecord(false)
				first.RoundNum, first.WinnerIndex, first.Tx.EvtLogId = 1, 0, 10
				second := validRaffleNftWinnerRecord(false)
				second.RoundNum, second.WinnerIndex, second.Tx.EvtLogId = 1, 1, 11
				server.raffles = fakeRoundRaffleReader{
					nft: func(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
						return []cgmodel.CGRaffleNFTWinnerRec{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "round donations",
			path: "/api/v2/cosmicgame/rounds/1/eth-donations?limit=1",
			configure: func(server *Server) {
				first := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
				first.RoundNum, first.Tx.EvtLogId = 1, 11
				second := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
				second.RoundNum, second.Tx.EvtLogId = 1, 10
				server.donations = fakeRoundDonationReader{
					eth: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
						return []cgstore.RoundEthDonationRecord{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "ROI leaderboard",
			path: "/api/v2/cosmicgame/statistics/leaderboard/roi?limit=1",
			configure: func(server *Server) {
				first := validROILeaderboardRecord()
				first.BidderAid, first.NetProfitWei = 2, "20"
				second := validROILeaderboardRecord()
				second.BidderAid, second.NetProfitWei = 3, "10"
				server.statistics = fakeStatisticsReader{
					roi: func(context.Context, int, cgstore.ROILeaderboardSort, *cgstore.ROILeaderboardPageCursor, int) ([]cgstore.ROILeaderboardRecord, bool, error) {
						return []cgstore.ROILeaderboardRecord{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "claim summaries",
			path: "/api/v2/cosmicgame/statistics/claims?limit=1",
			configure: func(server *Server) {
				server.statistics = fakeStatisticsReader{
					claims: func(context.Context, *cgstore.ClaimSummaryCursor, int) ([]cgstore.ClaimSummaryRecord, bool, error) {
						return []cgstore.ClaimSummaryRecord{
							validClaimSummaryRecord(2),
							validClaimSummaryRecord(1),
						}, false, nil
					},
				}
			},
		},
		{
			name: "claim transactions",
			path: "/api/v2/cosmicgame/rounds/0/claims?limit=1",
			configure: func(server *Server) {
				amount := "1"
				first := validClaimTransactionRecord(cgstore.ClaimAssetETH, func(record *cgstore.ClaimTransactionRecord) {
					record.RoundNum, record.EventLogID, record.EthAmountWei = 0, 10, &amount
				})
				second := first
				second.EventLogID = 11
				server.statistics = fakeStatisticsReader{
					transactions: func(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.ClaimTransactionRecord, bool, error) {
						return []cgstore.ClaimTransactionRecord{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "attached tokens",
			path: "/api/v2/cosmicgame/rounds/0/claims?limit=1",
			configure: func(server *Server) {
				tokenID := int64(7)
				first := validAttachedTokenRecord(cgstore.ClaimAssetERC721)
				first.RoundNum, first.EventLogID, first.TokenID = 0, 10, &tokenID
				second := first
				second.EventLogID = 11
				server.statistics = fakeStatisticsReader{
					attached: func(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.AttachedTokenRecord, bool, error) {
						return []cgstore.AttachedTokenRecord{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "unclaimed items",
			path: "/api/v2/cosmicgame/rounds/0/claims?limit=1",
			configure: func(server *Server) {
				amount := "1"
				first := validUnclaimedItemRecord(cgstore.ClaimAssetETH, func(record *cgstore.UnclaimedItemRecord) {
					record.RoundNum, record.Segment, record.Key, record.EthAmountWei = 0, 0, 10, &amount
				})
				second := first
				second.Key = 11
				server.statistics = fakeStatisticsReader{
					unclaimed: func(context.Context, int64, *cgstore.UnclaimedItemCursor, int) ([]cgstore.UnclaimedItemRecord, bool, error) {
						return []cgstore.UnclaimedItemRecord{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "user staking actions",
			path: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/actions?limit=1",
			configure: func(server *Server) {
				first := userStakingActionAt(cgstore.UserStakingActionUnstake, 300)
				second := userStakingActionAt(cgstore.UserStakingActionStake, 290)
				server.userStaking = fakeUserStakingReader{
					cstActions: func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error) {
						return []cgstore.UserStakingActionRecord{first, second}, false, nil
					},
				}
			},
		},
		{
			name: "user staked tokens",
			path: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/staked-tokens?limit=1",
			configure: func(server *Server) {
				server.userStaking = fakeUserStakingReader{
					cstStaked: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
						return []cgstore.UserStakedCstTokenRecord{
							userStakedCstTokenAt(5),
							userStakedCstTokenAt(9),
						}, false, nil
					},
				}
			},
		},
		{
			name: "user staking deposits",
			path: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits?limit=1",
			configure: func(server *Server) {
				server.userStaking = fakeUserStakingReader{
					deposits: func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error) {
						return []cgstore.UserStakingDepositRecord{
							userStakingDepositAt(502),
							userStakingDepositAt(501),
						}, false, nil
					},
				}
			},
		},
		{
			name: "user staking deposit rewards",
			path: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/deposits/501/rewards?limit=1",
			configure: func(server *Server) {
				reward := cgstore.UserStakingDepositRewardRecord{
					StakerAid: 1, ActionID: 2, TokenID: 5, RewardWei: "1",
				}
				second := reward
				second.ActionID = 7
				server.userStaking = fakeUserStakingReader{
					depositReward: func(context.Context, int64, int64, *cgstore.UserStakingRewardPageCursor, int) ([]cgstore.UserStakingDepositRewardRecord, bool, error) {
						return []cgstore.UserStakingDepositRewardRecord{reward, second}, false, nil
					},
				}
			},
		},
		{
			name: "user staking token rewards",
			path: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/token-rewards?limit=1",
			configure: func(server *Server) {
				reward := cgstore.UserStakingTokenRewardRecord{
					TokenID: 1, TotalWei: "2", CollectedWei: "1", PendingWei: "1",
				}
				second := reward
				second.TokenID = 5
				server.userStaking = fakeUserStakingReader{
					tokenRewards: func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakingTokenRewardRecord, bool, error) {
						return []cgstore.UserStakingTokenRewardRecord{reward, second}, false, nil
					},
				}
			},
		},
		{
			name: "user staking token reward deposits",
			path: "/api/v2/cosmicgame/users/" + userCursorAlice + "/staking/cst/token-rewards/5/deposits?limit=1",
			configure: func(server *Server) {
				deposit := cgstore.UserStakingTokenRewardDepositRecord{
					Tx:        validDonationTransaction(),
					StakerAid: 1, DepositID: 501, RoundNum: 0, RewardWei: "1",
				}
				second := deposit
				second.DepositID = 502
				server.userStaking = fakeUserStakingReader{
					tokenDeposits: func(context.Context, int64, int64, *cgstore.UserStakingTokenDepositPageCursor, int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
						return []cgstore.UserStakingTokenRewardDepositRecord{deposit, second}, false, nil
					},
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := newTestServer(t, fakeBidReader{})
			test.configure(server)

			response := serve(t, server, test.path)
			assertProblem(t, response, http.StatusInternalServerError)

			var problem Problem
			decodeResponse(t, response, &problem)
			if problem.Type != problemTypeBase+"internal" ||
				problem.Title != "Internal server error" ||
				problem.Detail == nil ||
				*problem.Detail != "The request could not be completed." ||
				problem.Instance == nil ||
				*problem.Instance != strings.Split(test.path, "?")[0] {
				t.Fatalf("problem = %+v", problem)
			}
			for _, internal := range []string{"repository", "rows", "requested"} {
				if strings.Contains(response.Body.String(), internal) {
					t.Fatalf("internal cardinality detail %q leaked: %s", internal, response.Body.String())
				}
			}
		})
	}
}

func TestPaginatedHandlersRejectCursorBoundaryReplay(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		setup func(*testing.T) (*Server, string)
	}{
		{
			name: "completed rounds",
			setup: func(t *testing.T) (*Server, string) {
				t.Helper()
				cursor, err := encodeRoundCursor(roundCursor{
					Version: roundCursorVersion, RoundNum: 7, EventLogID: 88,
				})
				if err != nil {
					t.Fatal(err)
				}
				record := validRoundRecord()
				record.RoundNum, record.ClaimPrizeTx.Tx.EvtLogId = 7, 88
				server := newTestServer(t, fakeBidReader{})
				server.rounds = fakeRoundReader{
					page: func(context.Context, *cgstore.RoundPageCursor, int) ([]cgmodel.CGRoundRec, bool, error) {
						return []cgmodel.CGRoundRec{record}, false, nil
					},
				}
				return server, "/api/v2/cosmicgame/rounds?cursor=" + cursor
			},
		},
		{
			name: "round bids",
			setup: func(t *testing.T) (*Server, string) {
				t.Helper()
				cursor, err := encodeBidCursor(bidCursor{
					Version: bidCursorVersion, Round: 1, BidPosition: 2, EventLogID: 88,
				})
				if err != nil {
					t.Fatal(err)
				}
				record := validBidRecord()
				record.RoundNum, record.BidPosition, record.Tx.EvtLogId = 1, 2, 88
				server := newTestServer(t, fakeBidReader{
					page: func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgmodel.CGBidRec, bool, error) {
						return []cgmodel.CGBidRec{record}, false, nil
					},
				})
				return server, "/api/v2/cosmicgame/rounds/1/bids?cursor=" + cursor
			},
		},
		{
			name: "round prizes",
			setup: func(t *testing.T) (*Server, string) {
				t.Helper()
				cursor, err := encodePrizeCursor(prizeCursor{
					Version: prizeCursorVersion, Round: 1, PrizeType: 2, WinnerIndex: 3,
				})
				if err != nil {
					t.Fatal(err)
				}
				record := validRoundPrizeRecord(2)
				record.RoundNum, record.WinnerIndex = 1, 3
				server := newTestServer(t, fakeBidReader{})
				server.prizes = fakeRoundPrizeReader{
					page: func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error) {
						return []cgmodel.CGPrizeHistory{record}, false, nil
					},
				}
				return server, "/api/v2/cosmicgame/rounds/1/prizes?cursor=" + cursor
			},
		},
		{
			name: "raffle ETH deposits",
			setup: func(t *testing.T) (*Server, string) {
				t.Helper()
				cursor, err := encodeRaffleEthDepositCursor(raffleEthDepositCursor{
					Version: raffleEthDepositCursorVersion, Round: 1, WinnerIndex: 2, EventLogID: 88,
				})
				if err != nil {
					t.Fatal(err)
				}
				record := validRaffleEthDepositRecord()
				record.RoundNum, record.WinnerIndex, record.Tx.EvtLogId = 1, 2, 88
				server := newTestServer(t, fakeBidReader{})
				server.raffles = fakeRoundRaffleReader{
					eth: func(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error) {
						return []cgstore.RaffleEthDepositRecord{record}, false, nil
					},
				}
				return server, "/api/v2/cosmicgame/rounds/1/raffle-eth-deposits?cursor=" + cursor
			},
		},
		{
			name: "raffle NFT winners",
			setup: func(t *testing.T) (*Server, string) {
				t.Helper()
				cursor, err := encodeRaffleNftWinnerCursor(raffleNftWinnerCursor{
					Version: raffleNftWinnerCursorVersion,
					Round:   1, Pool: Bidder, WinnerIndex: 2, EventLogID: 88,
				})
				if err != nil {
					t.Fatal(err)
				}
				record := validRaffleNftWinnerRecord(false)
				record.RoundNum, record.WinnerIndex, record.Tx.EvtLogId = 1, 2, 88
				server := newTestServer(t, fakeBidReader{})
				server.raffles = fakeRoundRaffleReader{
					nft: func(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
						return []cgmodel.CGRaffleNFTWinnerRec{record}, false, nil
					},
				}
				return server, "/api/v2/cosmicgame/rounds/1/raffle-nft-winners?pool=bidder&cursor=" + cursor
			},
		},
		{
			name: "round donations",
			setup: func(t *testing.T) (*Server, string) {
				t.Helper()
				cursor, err := encodeRoundDonationCursor(roundDonationCursor{
					Version: roundDonationCursorVersion,
					Round:   1, Resource: roundDonationResourceETH, EventLogID: 88,
				})
				if err != nil {
					t.Fatal(err)
				}
				record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
				record.RoundNum, record.Tx.EvtLogId = 1, 88
				server := newTestServer(t, fakeBidReader{})
				server.donations = fakeRoundDonationReader{
					eth: func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error) {
						return []cgstore.RoundEthDonationRecord{record}, false, nil
					},
				}
				return server, "/api/v2/cosmicgame/rounds/1/eth-donations?cursor=" + cursor
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server, path := test.setup(t)
			response := serve(t, server, path)
			assertProblem(t, response, http.StatusInternalServerError)
			if strings.Contains(response.Body.String(), "unordered") ||
				strings.Contains(response.Body.String(), "cursor") {
				t.Fatalf("ordering detail leaked: %s", response.Body.String())
			}
		})
	}
}

func TestRequestBindingErrorsUseStableSecretFreeDetails(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		path       string
		wantDetail string
	}{
		{
			name:       "invalid format",
			path:       "/api/v2/cosmicgame/rounds?limit=password-super-secret",
			wantDetail: `Parameter "limit" has an invalid value.`,
		},
		{
			name:       "duplicate value",
			path:       "/api/v2/cosmicgame/rounds?limit=1&limit=2",
			wantDetail: `Parameter "limit" has an invalid value.`,
		},
		{
			name:       "missing required value",
			path:       "/api/v2/cosmicgame/rounds/1/raffle-nft-winners",
			wantDetail: `Parameter "pool" is required.`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newTestServer(t, fakeBidReader{}), test.path)
			assertProblem(t, response, http.StatusBadRequest)

			var problem Problem
			decodeResponse(t, response, &problem)
			if problem.Type != problemTypeBase+"invalid-request" ||
				problem.Title != "Invalid request" ||
				problem.Detail == nil ||
				*problem.Detail != test.wantDetail {
				t.Fatalf("problem = %+v", problem)
			}
			if strings.Contains(response.Body.String(), "password-super-secret") {
				t.Fatalf("invalid parameter value leaked: %s", response.Body.String())
			}
		})
	}

	t.Run("too many values error", func(t *testing.T) {
		t.Parallel()
		server := newTestServer(t, fakeBidReader{})
		request := httptest.NewRequest(http.MethodGet, "/api/v2/cosmicgame/rounds", nil)
		response := httptest.NewRecorder()
		server.writeRequestError(response, request, &TooManyValuesForParamError{
			ParamName: "limit",
			Count:     2,
		})

		assertProblem(t, response, http.StatusBadRequest)
		var problem Problem
		decodeResponse(t, response, &problem)
		if problem.Detail == nil || *problem.Detail != `Parameter "limit" must be provided once.` {
			t.Fatalf("problem = %+v", problem)
		}
	})
}
